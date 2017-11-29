package controlapi

import (
	"crypto/x509"
	"encoding/pem"

	"github.com/docker/swarmkit/api"
	"github.com/docker/swarmkit/manager/state/raft/membership"
	"github.com/docker/swarmkit/manager/state/store"
	gogotypes "github.com/gogo/protobuf/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func validateNodeSpec(spec *api.NodeSpec) error {
	if spec == nil {
		return grpc.Errorf(codes.InvalidArgument, errInvalidArgument.Error())
	}
	return nil
}

// GetNode returns a Node given a NodeID.
// - Returns `InvalidArgument` if NodeID is not provided.
// - Returns `NotFound` if the Node is not found.
func (s *Server) GetNode(ctx context.Context, request *api.GetNodeRequest) (*api.GetNodeResponse, error) {
	if request.NodeID == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, errInvalidArgument.Error())
	}

	var node *api.Node
	s.store.View(func(tx store.ReadTx) {
		node = store.GetNode(tx, request.NodeID)
	})
	if node == nil {
		return nil, grpc.Errorf(codes.NotFound, "node %s not found", request.NodeID)
	}

	if s.raft != nil {
		memberlist := s.raft.GetMemberlist()
		for _, member := range memberlist {
			if member.NodeID == node.ID {
				node.ManagerStatus = &api.ManagerStatus{
					RaftID:       member.RaftID,
					Addr:         member.Addr,
					Leader:       member.Status.Leader,
					Reachability: member.Status.Reachability,
				}
				break
			}
		}
	}

	return &api.GetNodeResponse{
		Node: node,
	}, nil
}

func filterNodes(candidates []*api.Node, filters ...func(*api.Node) bool) []*api.Node {
	result := []*api.Node{}

	for _, c := range candidates {
		match := true
		for _, f := range filters {
			if !f(c) {
				match = false
				break
			}
		}
		if match {
			result = append(result, c)
		}
	}

	return result
}

// ListNodes returns a list of all nodes.
func (s *Server) ListNodes(ctx context.Context, request *api.ListNodesRequest) (*api.ListNodesResponse, error) {
	var (
		nodes []*api.Node
		err   error
	)
	s.store.View(func(tx store.ReadTx) {
		switch {
		case request.Filters != nil && len(request.Filters.Names) > 0:
			nodes, err = store.FindNodes(tx, buildFilters(store.ByName, request.Filters.Names))
		case request.Filters != nil && len(request.Filters.NamePrefixes) > 0:
			nodes, err = store.FindNodes(tx, buildFilters(store.ByNamePrefix, request.Filters.NamePrefixes))
		case request.Filters != nil && len(request.Filters.IDPrefixes) > 0:
			nodes, err = store.FindNodes(tx, buildFilters(store.ByIDPrefix, request.Filters.IDPrefixes))
		case request.Filters != nil && len(request.Filters.Roles) > 0:
			filters := make([]store.By, 0, len(request.Filters.Roles))
			for _, v := range request.Filters.Roles {
				filters = append(filters, store.ByRole(v))
			}
			nodes, err = store.FindNodes(tx, store.Or(filters...))
		case request.Filters != nil && len(request.Filters.Memberships) > 0:
			filters := make([]store.By, 0, len(request.Filters.Memberships))
			for _, v := range request.Filters.Memberships {
				filters = append(filters, store.ByMembership(v))
			}
			nodes, err = store.FindNodes(tx, store.Or(filters...))
		default:
			nodes, err = store.FindNodes(tx, store.All)
		}
	})
	if err != nil {
		return nil, err
	}

	if request.Filters != nil {
		nodes = filterNodes(nodes,
			func(e *api.Node) bool {
				if len(request.Filters.Names) == 0 {
					return true
				}
				if e.Description == nil {
					return false
				}
				return filterContains(e.Description.Hostname, request.Filters.Names)
			},
			func(e *api.Node) bool {
				if len(request.Filters.NamePrefixes) == 0 {
					return true
				}
				if e.Description == nil {
					return false
				}
				return filterContainsPrefix(e.Description.Hostname, request.Filters.NamePrefixes)
			},
			func(e *api.Node) bool {
				return filterContainsPrefix(e.ID, request.Filters.IDPrefixes)
			},
			func(e *api.Node) bool {
				if len(request.Filters.Labels) == 0 {
					return true
				}
				if e.Description == nil {
					return false
				}
				return filterMatchLabels(e.Description.Engine.Labels, request.Filters.Labels)
			},
			func(e *api.Node) bool {
				if len(request.Filters.Roles) == 0 {
					return true
				}
				for _, c := range request.Filters.Roles {
					if c == e.Role {
						return true
					}
				}
				return false
			},
			func(e *api.Node) bool {
				if len(request.Filters.Memberships) == 0 {
					return true
				}
				for _, c := range request.Filters.Memberships {
					if c == e.Spec.Membership {
						return true
					}
				}
				return false
			},
		)
	}

	// Add in manager information on nodes that are managers
	if s.raft != nil {
		memberlist := s.raft.GetMemberlist()

		for _, node := range nodes {
			for _, member := range memberlist {
				if member.NodeID == node.ID {
					node.ManagerStatus = &api.ManagerStatus{
						RaftID:       member.RaftID,
						Addr:         member.Addr,
						Leader:       member.Status.Leader,
						Reachability: member.Status.Reachability,
					}
					break
				}
			}
		}
	}

	return &api.ListNodesResponse{
		Nodes: nodes,
	}, nil
}

// UpdateNode updates a Node referenced by NodeID with the given NodeSpec.
// - Returns `NotFound` if the Node is not found.
// - Returns `InvalidArgument` if the NodeSpec is malformed.
// - Returns an error if the update fails.
func (s *Server) UpdateNode(ctx context.Context, request *api.UpdateNodeRequest) (*api.UpdateNodeResponse, error) {
	if request.NodeID == "" || request.NodeVersion == nil {
		return nil, grpc.Errorf(codes.InvalidArgument, errInvalidArgument.Error())
	}
	if err := validateNodeSpec(request.Spec); err != nil {
		return nil, err
	}

	var (
		node   *api.Node
		member *membership.Member
	)

	err := s.store.Update(func(tx store.Tx) error {
		node = store.GetNode(tx, request.NodeID)
		if node == nil {
			return grpc.Errorf(codes.NotFound, "node %s not found", request.NodeID)
		}

		// Demotion sanity checks.
		if node.Spec.DesiredRole == api.NodeRoleManager && request.Spec.DesiredRole == api.NodeRoleWorker {
			// Check for manager entries in Store.
			managers, err := store.FindNodes(tx, store.ByRole(api.NodeRoleManager))
			if err != nil {
				return grpc.Errorf(codes.Internal, "internal store error: %v", err)
			}
			if len(managers) == 1 && managers[0].ID == node.ID {
				return grpc.Errorf(codes.FailedPrecondition, "attempting to demote the last manager of the swarm")
			}

			// Check for node in memberlist
			if member = s.raft.GetMemberByNodeID(request.NodeID); member == nil {
				return grpc.Errorf(codes.NotFound, "can't find manager in raft memberlist")
			}

			// Quorum safeguard
			if !s.raft.CanRemoveMember(member.RaftID) {
				return grpc.Errorf(codes.FailedPrecondition, "can't remove member from the raft: this would result in a loss of quorum")
			}
		}

		node.Meta.Version = *request.NodeVersion
		node.Spec = *request.Spec.Copy()
		return store.UpdateNode(tx, node)
	})
	if err != nil {
		return nil, err
	}

	return &api.UpdateNodeResponse{
		Node: node,
	}, nil
}

func removeNodeAttachments(tx store.Tx, nodeID string) error {
	// orphan the node's attached containers. if we don't do this, the
	// network these attachments are connected to will never be removeable
	tasks, err := store.FindTasks(tx, store.ByNodeID(nodeID))
	if err != nil {
		return err
	}
	for _, task := range tasks {
		// if the task is an attachment, then we just delete it. the allocator
		// will do the heavy lifting. basically, GetAttachment will return the
		// attachment if that's the kind of runtime, or nil if it's not.
		if task.Spec.GetAttachment() != nil {
			// don't delete the task. instead, update it to `ORPHANED` so that
			// the taskreaper will clean it up.
			task.Status.State = api.TaskStateOrphaned
			if err := store.UpdateTask(tx, task); err != nil {
				return err
			}
		}
	}
	return nil
}

// RemoveNode removes a Node referenced by NodeID with the given NodeSpec.
// - Returns NotFound if the Node is not found.
// - Returns FailedPrecondition if the Node has manager role (and is part of the memberlist) or is not shut down.
// - Returns InvalidArgument if NodeID or NodeVersion is not valid.
// - Returns an error if the delete fails.
func (s *Server) RemoveNode(ctx context.Context, request *api.RemoveNodeRequest) (*api.RemoveNodeResponse, error) {
	if request.NodeID == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, errInvalidArgument.Error())
	}

	err := s.store.Update(func(tx store.Tx) error {
		node := store.GetNode(tx, request.NodeID)
		if node == nil {
			return grpc.Errorf(codes.NotFound, "node %s not found", request.NodeID)
		}
		if node.Spec.DesiredRole == api.NodeRoleManager {
			if s.raft == nil {
				return grpc.Errorf(codes.FailedPrecondition, "node %s is a manager but cannot access node information from the raft memberlist", request.NodeID)
			}
			if member := s.raft.GetMemberByNodeID(request.NodeID); member != nil {
				return grpc.Errorf(codes.FailedPrecondition, "node %s is a cluster manager and is a member of the raft cluster. It must be demoted to worker before removal", request.NodeID)
			}
		}
		if !request.Force && node.Status.State == api.NodeStatus_READY {
			return grpc.Errorf(codes.FailedPrecondition, "node %s is not down and can't be removed", request.NodeID)
		}

		// lookup the cluster
		clusters, err := store.FindClusters(tx, store.ByName(store.DefaultClusterName))
		if err != nil {
			return err
		}
		if len(clusters) != 1 {
			return grpc.Errorf(codes.Internal, "could not fetch cluster object")
		}
		cluster := clusters[0]

		blacklistedCert := &api.BlacklistedCertificate{}

		// Set an expiry time for this RemovedNode if a certificate
		// exists and can be parsed.
		if len(node.Certificate.Certificate) != 0 {
			certBlock, _ := pem.Decode(node.Certificate.Certificate)
			if certBlock != nil {
				X509Cert, err := x509.ParseCertificate(certBlock.Bytes)
				if err == nil && !X509Cert.NotAfter.IsZero() {
					expiry, err := gogotypes.TimestampProto(X509Cert.NotAfter)
					if err == nil {
						blacklistedCert.Expiry = expiry
					}
				}
			}
		}

		if cluster.BlacklistedCertificates == nil {
			cluster.BlacklistedCertificates = make(map[string]*api.BlacklistedCertificate)
		}
		cluster.BlacklistedCertificates[node.ID] = blacklistedCert

		expireBlacklistedCerts(cluster)

		if err := store.UpdateCluster(tx, cluster); err != nil {
			return err
		}

		if err := removeNodeAttachments(tx, request.NodeID); err != nil {
			return err
		}

		return store.DeleteNode(tx, request.NodeID)
	})
	if err != nil {
		return nil, err
	}
	return &api.RemoveNodeResponse{}, nil
}
