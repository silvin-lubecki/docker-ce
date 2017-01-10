// Code generated by protoc-gen-gogo.
// source: health.proto
// DO NOT EDIT!

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/docker/swarmkit/protobuf/plugin"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import raftselector "github.com/docker/swarmkit/manager/raftselector"
import codes "google.golang.org/grpc/codes"
import metadata "google.golang.org/grpc/metadata"
import transport "google.golang.org/grpc/transport"
import time "time"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorHealth, []int{1, 0}
}

type HealthCheckRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptorHealth, []int{0} }

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=docker.swarmkit.v1.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptorHealth, []int{1} }

func init() {
	proto.RegisterType((*HealthCheckRequest)(nil), "docker.swarmkit.v1.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "docker.swarmkit.v1.HealthCheckResponse")
	proto.RegisterEnum("docker.swarmkit.v1.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

type authenticatedWrapperHealthServer struct {
	local     HealthServer
	authorize func(context.Context, []string) error
}

func NewAuthenticatedWrapperHealthServer(local HealthServer, authorize func(context.Context, []string) error) HealthServer {
	return &authenticatedWrapperHealthServer{
		local:     local,
		authorize: authorize,
	}
}

func (p *authenticatedWrapperHealthServer) Check(ctx context.Context, r *HealthCheckRequest) (*HealthCheckResponse, error) {

	if err := p.authorize(ctx, []string{"swarm-manager"}); err != nil {
		return nil, err
	}
	return p.local.Check(ctx, r)
}

func (m *HealthCheckRequest) Copy() *HealthCheckRequest {
	if m == nil {
		return nil
	}

	o := &HealthCheckRequest{
		Service: m.Service,
	}

	return o
}

func (m *HealthCheckResponse) Copy() *HealthCheckResponse {
	if m == nil {
		return nil
	}

	o := &HealthCheckResponse{
		Status: m.Status,
	}

	return o
}

func (this *HealthCheckRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&api.HealthCheckRequest{")
	s = append(s, "Service: "+fmt.Sprintf("%#v", this.Service)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HealthCheckResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&api.HealthCheckResponse{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringHealth(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringHealth(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Health service

type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/docker.swarmkit.v1.Health/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.swarmkit.v1.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "docker.swarmkit.v1.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorHealth,
}

func (m *HealthCheckRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *HealthCheckRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Service) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintHealth(data, i, uint64(len(m.Service)))
		i += copy(data[i:], m.Service)
	}
	return i, nil
}

func (m *HealthCheckResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *HealthCheckResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintHealth(data, i, uint64(m.Status))
	}
	return i, nil
}

func encodeFixed64Health(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Health(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHealth(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}

type raftProxyHealthServer struct {
	local                       HealthServer
	connSelector                raftselector.ConnProvider
	localCtxMods, remoteCtxMods []func(context.Context) (context.Context, error)
}

func NewRaftProxyHealthServer(local HealthServer, connSelector raftselector.ConnProvider, localCtxMod, remoteCtxMod func(context.Context) (context.Context, error)) HealthServer {
	redirectChecker := func(ctx context.Context) (context.Context, error) {
		s, ok := transport.StreamFromContext(ctx)
		if !ok {
			return ctx, grpc.Errorf(codes.InvalidArgument, "remote addr is not found in context")
		}
		addr := s.ServerTransport().RemoteAddr().String()
		md, ok := metadata.FromContext(ctx)
		if ok && len(md["redirect"]) != 0 {
			return ctx, grpc.Errorf(codes.ResourceExhausted, "more than one redirect to leader from: %s", md["redirect"])
		}
		if !ok {
			md = metadata.New(map[string]string{})
		}
		md["redirect"] = append(md["redirect"], addr)
		return metadata.NewContext(ctx, md), nil
	}
	remoteMods := []func(context.Context) (context.Context, error){redirectChecker}
	remoteMods = append(remoteMods, remoteCtxMod)

	var localMods []func(context.Context) (context.Context, error)
	if localCtxMod != nil {
		localMods = []func(context.Context) (context.Context, error){localCtxMod}
	}

	return &raftProxyHealthServer{
		local:         local,
		connSelector:  connSelector,
		localCtxMods:  localMods,
		remoteCtxMods: remoteMods,
	}
}
func (p *raftProxyHealthServer) runCtxMods(ctx context.Context, ctxMods []func(context.Context) (context.Context, error)) (context.Context, error) {
	var err error
	for _, mod := range ctxMods {
		ctx, err = mod(ctx)
		if err != nil {
			return ctx, err
		}
	}
	return ctx, nil
}
func (p *raftProxyHealthServer) pollNewLeaderConn(ctx context.Context) (*grpc.ClientConn, error) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			conn, err := p.connSelector.LeaderConn(ctx)
			if err != nil {
				return nil, err
			}

			client := NewHealthClient(conn)

			resp, err := client.Check(ctx, &HealthCheckRequest{Service: "Raft"})
			if err != nil || resp.Status != HealthCheckResponse_SERVING {
				continue
			}
			return conn, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func (p *raftProxyHealthServer) Check(ctx context.Context, r *HealthCheckRequest) (*HealthCheckResponse, error) {

	conn, err := p.connSelector.LeaderConn(ctx)
	if err != nil {
		if err == raftselector.ErrIsLeader {
			ctx, err = p.runCtxMods(ctx, p.localCtxMods)
			if err != nil {
				return nil, err
			}
			return p.local.Check(ctx, r)
		}
		return nil, err
	}
	modCtx, err := p.runCtxMods(ctx, p.remoteCtxMods)
	if err != nil {
		return nil, err
	}

	resp, err := NewHealthClient(conn).Check(modCtx, r)
	if err != nil {
		if !strings.Contains(err.Error(), "is closing") && !strings.Contains(err.Error(), "the connection is unavailable") && !strings.Contains(err.Error(), "connection error") {
			return resp, err
		}
		conn, err := p.pollNewLeaderConn(ctx)
		if err != nil {
			if err == raftselector.ErrIsLeader {
				return p.local.Check(ctx, r)
			}
			return nil, err
		}
		return NewHealthClient(conn).Check(modCtx, r)
	}
	return resp, err
}

func (m *HealthCheckRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sovHealth(uint64(l))
	}
	return n
}

func (m *HealthCheckResponse) Size() (n int) {
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovHealth(uint64(m.Status))
	}
	return n
}

func sovHealth(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHealth(x uint64) (n int) {
	return sovHealth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HealthCheckRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HealthCheckRequest{`,
		`Service:` + fmt.Sprintf("%v", this.Service) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HealthCheckResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HealthCheckResponse{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHealth(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HealthCheckRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HealthCheckRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheckRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHealth(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HealthCheckResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HealthCheckResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheckResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Status |= (HealthCheckResponse_ServingStatus(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHealth(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHealth(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHealth
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHealth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHealth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHealth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHealth
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHealth(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHealth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHealth   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("health.proto", fileDescriptorHealth) }

var fileDescriptorHealth = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0xc9, 0x4f, 0xce, 0x4e, 0x2d,
	0xd2, 0x2b, 0x2e, 0x4f, 0x2c, 0xca, 0xcd, 0xce, 0x2c, 0xd1, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0x4b, 0xeb, 0x83, 0x58, 0x10, 0x95, 0x52, 0xc2, 0x05, 0x39, 0xa5, 0xe9, 0x99,
	0x79, 0xfa, 0x10, 0x0a, 0x22, 0xa8, 0xa4, 0xc7, 0x25, 0xe4, 0x01, 0x36, 0xce, 0x39, 0x23, 0x35,
	0x39, 0x3b, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x82, 0x8b, 0xbd, 0x38, 0xb5, 0xa8,
	0x2c, 0x33, 0x39, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x55, 0x5a, 0xc0, 0xc8,
	0x25, 0x8c, 0xa2, 0xa1, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x97, 0x8b, 0xad, 0xb8, 0x24,
	0xb1, 0xa4, 0xb4, 0x18, 0xac, 0x81, 0xcf, 0xc8, 0x54, 0x0f, 0xd3, 0x5d, 0x7a, 0x58, 0x34, 0xea,
	0x05, 0x83, 0x0c, 0xce, 0x4b, 0x0f, 0x06, 0x6b, 0x0e, 0x82, 0x1a, 0xa2, 0x64, 0xc5, 0xc5, 0x8b,
	0x22, 0x21, 0xc4, 0xcd, 0xc5, 0x1e, 0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x27, 0xc0, 0x00, 0xe2,
	0x04, 0xbb, 0x06, 0x85, 0x79, 0xfa, 0xb9, 0x0b, 0x30, 0x0a, 0xf1, 0x73, 0x71, 0xfb, 0xf9, 0x87,
	0xc4, 0xc3, 0x04, 0x98, 0x8c, 0x2a, 0xb9, 0xd8, 0x20, 0x16, 0x09, 0xe5, 0x73, 0xb1, 0x82, 0x2d,
	0x13, 0x52, 0x23, 0xe8, 0x1a, 0xb0, 0xbf, 0xa5, 0xd4, 0x89, 0x74, 0xb5, 0x92, 0xe8, 0xa9, 0x75,
	0xef, 0x66, 0x30, 0xf1, 0x73, 0xf1, 0x82, 0x15, 0xea, 0xe6, 0x26, 0xe6, 0x25, 0xa6, 0xa7, 0x16,
	0x39, 0xc9, 0x9c, 0x78, 0x28, 0xc7, 0x70, 0xe3, 0xa1, 0x1c, 0xc3, 0x87, 0x87, 0x72, 0x8c, 0x0d,
	0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x24, 0x36, 0x70, 0x90, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x14, 0x7c, 0x23, 0xc1,
	0x01, 0x00, 0x00,
}
