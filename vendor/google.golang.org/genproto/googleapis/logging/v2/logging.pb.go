// Code generated by protoc-gen-go.
// source: google/logging/v2/logging.proto
// DO NOT EDIT!

package logging

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_api3 "google.golang.org/genproto/googleapis/api/monitoredres"
import _ "github.com/golang/protobuf/ptypes/duration"
import google_protobuf5 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The parameters to DeleteLog.
type DeleteLogRequest struct {
	// Required. The resource name of the log to delete:
	//
	//     "projects/[PROJECT_ID]/logs/[LOG_ID]"
	//     "organizations/[ORGANIZATION_ID]/logs/[LOG_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/logs/[LOG_ID]"
	//     "folders/[FOLDER_ID]/logs/[LOG_ID]"
	//
	// `[LOG_ID]` must be URL-encoded. For example,
	// `"projects/my-project-id/logs/syslog"`,
	// `"organizations/1234567890/logs/cloudresourcemanager.googleapis.com%2Factivity"`.
	// For more information about log names, see
	// [LogEntry][google.logging.v2.LogEntry].
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName" json:"log_name,omitempty"`
}

func (m *DeleteLogRequest) Reset()                    { *m = DeleteLogRequest{} }
func (m *DeleteLogRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogRequest) ProtoMessage()               {}
func (*DeleteLogRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *DeleteLogRequest) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

// The parameters to WriteLogEntries.
type WriteLogEntriesRequest struct {
	// Optional. A default log resource name that is assigned to all log entries
	// in `entries` that do not specify a value for `log_name`:
	//
	//     "projects/[PROJECT_ID]/logs/[LOG_ID]"
	//     "organizations/[ORGANIZATION_ID]/logs/[LOG_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/logs/[LOG_ID]"
	//     "folders/[FOLDER_ID]/logs/[LOG_ID]"
	//
	// `[LOG_ID]` must be URL-encoded. For example,
	// `"projects/my-project-id/logs/syslog"` or
	// `"organizations/1234567890/logs/cloudresourcemanager.googleapis.com%2Factivity"`.
	// For more information about log names, see
	// [LogEntry][google.logging.v2.LogEntry].
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName" json:"log_name,omitempty"`
	// Optional. A default monitored resource object that is assigned to all log
	// entries in `entries` that do not specify a value for `resource`. Example:
	//
	//     { "type": "gce_instance",
	//       "labels": {
	//         "zone": "us-central1-a", "instance_id": "00000000000000000000" }}
	//
	// See [LogEntry][google.logging.v2.LogEntry].
	Resource *google_api3.MonitoredResource `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	// Optional. Default labels that are added to the `labels` field of all log
	// entries in `entries`. If a log entry already has a label with the same key
	// as a label in this parameter, then the log entry's label is not changed.
	// See [LogEntry][google.logging.v2.LogEntry].
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Required.  The log entries to write. Values supplied for the fields
	// `log_name`, `resource`, and `labels` in this `entries.write` request are
	// inserted into those log entries in this list that do not provide their own
	// values.
	//
	// Stackdriver Logging also creates and inserts values for `timestamp` and
	// `insert_id` if the entries do not provide them. The created `insert_id` for
	// the N'th entry in this list will be greater than earlier entries and less
	// than later entries.  Otherwise, the order of log entries in this list does
	// not matter.
	//
	// To improve throughput and to avoid exceeding the
	// [quota limit](/logging/quota-policy) for calls to `entries.write`,
	// you should write multiple log entries at once rather than
	// calling this method for each individual log entry.
	Entries []*LogEntry `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
	// Optional. Whether valid entries should be written even if some other
	// entries fail due to INVALID_ARGUMENT or PERMISSION_DENIED errors. If any
	// entry is not written, then the response status is the error associated
	// with one of the failed entries and the response includes error details
	// keyed by the entries' zero-based index in the `entries.write` method.
	PartialSuccess bool `protobuf:"varint,5,opt,name=partial_success,json=partialSuccess" json:"partial_success,omitempty"`
}

func (m *WriteLogEntriesRequest) Reset()                    { *m = WriteLogEntriesRequest{} }
func (m *WriteLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesRequest) ProtoMessage()               {}
func (*WriteLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *WriteLogEntriesRequest) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *WriteLogEntriesRequest) GetResource() *google_api3.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetPartialSuccess() bool {
	if m != nil {
		return m.PartialSuccess
	}
	return false
}

// Result returned from WriteLogEntries.
// empty
type WriteLogEntriesResponse struct {
}

func (m *WriteLogEntriesResponse) Reset()                    { *m = WriteLogEntriesResponse{} }
func (m *WriteLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesResponse) ProtoMessage()               {}
func (*WriteLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// Error details for WriteLogEntries with partial success.
type WriteLogEntriesPartialErrors struct {
	// When `WriteLogEntriesRequest.partial_success` is true, records the error
	// status for entries that were not written due to a permanent error, keyed
	// by the entry's zero-based index in `WriteLogEntriesRequest.entries`.
	//
	// Failed requests for which no entries are written will not include
	// per-entry errors.
	LogEntryErrors map[int32]*google_rpc.Status `protobuf:"bytes,1,rep,name=log_entry_errors,json=logEntryErrors" json:"log_entry_errors,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *WriteLogEntriesPartialErrors) Reset()                    { *m = WriteLogEntriesPartialErrors{} }
func (m *WriteLogEntriesPartialErrors) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesPartialErrors) ProtoMessage()               {}
func (*WriteLogEntriesPartialErrors) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *WriteLogEntriesPartialErrors) GetLogEntryErrors() map[int32]*google_rpc.Status {
	if m != nil {
		return m.LogEntryErrors
	}
	return nil
}

// The parameters to `ListLogEntries`.
type ListLogEntriesRequest struct {
	// Deprecated. Use `resource_names` instead.  One or more project identifiers
	// or project numbers from which to retrieve log entries.  Example:
	// `"my-project-1A"`. If present, these project identifiers are converted to
	// resource name format and added to the list of resources in
	// `resource_names`.
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds" json:"project_ids,omitempty"`
	// Required. Names of one or more parent resources from which to
	// retrieve log entries:
	//
	//     "projects/[PROJECT_ID]"
	//     "organizations/[ORGANIZATION_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]"
	//     "folders/[FOLDER_ID]"
	//
	// Projects listed in the `project_ids` field are added to this list.
	ResourceNames []string `protobuf:"bytes,8,rep,name=resource_names,json=resourceNames" json:"resource_names,omitempty"`
	// Optional. A filter that chooses which log entries to return.  See [Advanced
	// Logs Filters](/logging/docs/view/advanced_filters).  Only log entries that
	// match the filter are returned.  An empty filter matches all log entries in
	// the resources listed in `resource_names`. Referencing a parent resource
	// that is not listed in `resource_names` will cause the filter to return no
	// results.
	// The maximum length of the filter is 20000 characters.
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// Optional. How the results should be sorted.  Presently, the only permitted
	// values are `"timestamp asc"` (default) and `"timestamp desc"`. The first
	// option returns entries in order of increasing values of
	// `LogEntry.timestamp` (oldest first), and the second option returns entries
	// in order of decreasing timestamps (newest first).  Entries with equal
	// timestamps are returned in order of their `insert_id` values.
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy" json:"order_by,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored.  The presence of `next_page_token` in the
	// response indicates that more results might be available.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the
	// preceding call to this method.  `page_token` must be the value of
	// `next_page_token` from the previous response.  The values of other method
	// parameters should be identical to those in the previous call.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListLogEntriesRequest) Reset()                    { *m = ListLogEntriesRequest{} }
func (m *ListLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesRequest) ProtoMessage()               {}
func (*ListLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ListLogEntriesRequest) GetProjectIds() []string {
	if m != nil {
		return m.ProjectIds
	}
	return nil
}

func (m *ListLogEntriesRequest) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func (m *ListLogEntriesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListLogEntriesRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *ListLogEntriesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListLogEntriesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Result returned from `ListLogEntries`.
type ListLogEntriesResponse struct {
	// A list of log entries.
	Entries []*LogEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	// If there might be more results than those appearing in this response, then
	// `nextPageToken` is included.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	//
	// If a value for `next_page_token` appears and the `entries` field is empty,
	// it means that the search found no log entries so far but it did not have
	// time to search all the possible log entries.  Retry the method with this
	// value for `page_token` to continue the search.  Alternatively, consider
	// speeding up the search by changing your filter to specify a single log name
	// or resource type, or to narrow the time range of the search.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLogEntriesResponse) Reset()                    { *m = ListLogEntriesResponse{} }
func (m *ListLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesResponse) ProtoMessage()               {}
func (*ListLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ListLogEntriesResponse) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *ListLogEntriesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The parameters to ListMonitoredResourceDescriptors
type ListMonitoredResourceDescriptorsRequest struct {
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored.  The presence of `nextPageToken` in the
	// response indicates that more results might be available.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the
	// preceding call to this method.  `pageToken` must be the value of
	// `nextPageToken` from the previous response.  The values of other method
	// parameters should be identical to those in the previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsRequest) Reset() {
	*m = ListMonitoredResourceDescriptorsRequest{}
}
func (m *ListMonitoredResourceDescriptorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsRequest) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{6}
}

func (m *ListMonitoredResourceDescriptorsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListMonitoredResourceDescriptorsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Result returned from ListMonitoredResourceDescriptors.
type ListMonitoredResourceDescriptorsResponse struct {
	// A list of resource descriptors.
	ResourceDescriptors []*google_api3.MonitoredResourceDescriptor `protobuf:"bytes,1,rep,name=resource_descriptors,json=resourceDescriptors" json:"resource_descriptors,omitempty"`
	// If there might be more results than those appearing in this response, then
	// `nextPageToken` is included.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsResponse) Reset() {
	*m = ListMonitoredResourceDescriptorsResponse{}
}
func (m *ListMonitoredResourceDescriptorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsResponse) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{7}
}

func (m *ListMonitoredResourceDescriptorsResponse) GetResourceDescriptors() []*google_api3.MonitoredResourceDescriptor {
	if m != nil {
		return m.ResourceDescriptors
	}
	return nil
}

func (m *ListMonitoredResourceDescriptorsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The parameters to ListLogs.
type ListLogsRequest struct {
	// Required. The resource name that owns the logs:
	//
	//     "projects/[PROJECT_ID]"
	//     "organizations/[ORGANIZATION_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]"
	//     "folders/[FOLDER_ID]"
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored.  The presence of `nextPageToken` in the
	// response indicates that more results might be available.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the
	// preceding call to this method.  `pageToken` must be the value of
	// `nextPageToken` from the previous response.  The values of other method
	// parameters should be identical to those in the previous call.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListLogsRequest) Reset()                    { *m = ListLogsRequest{} }
func (m *ListLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogsRequest) ProtoMessage()               {}
func (*ListLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *ListLogsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListLogsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListLogsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Result returned from ListLogs.
type ListLogsResponse struct {
	// A list of log names. For example,
	// `"projects/my-project/syslog"` or
	// `"organizations/123/cloudresourcemanager.googleapis.com%2Factivity"`.
	LogNames []string `protobuf:"bytes,3,rep,name=log_names,json=logNames" json:"log_names,omitempty"`
	// If there might be more results than those appearing in this response, then
	// `nextPageToken` is included.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLogsResponse) Reset()                    { *m = ListLogsResponse{} }
func (m *ListLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogsResponse) ProtoMessage()               {}
func (*ListLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *ListLogsResponse) GetLogNames() []string {
	if m != nil {
		return m.LogNames
	}
	return nil
}

func (m *ListLogsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteLogRequest)(nil), "google.logging.v2.DeleteLogRequest")
	proto.RegisterType((*WriteLogEntriesRequest)(nil), "google.logging.v2.WriteLogEntriesRequest")
	proto.RegisterType((*WriteLogEntriesResponse)(nil), "google.logging.v2.WriteLogEntriesResponse")
	proto.RegisterType((*WriteLogEntriesPartialErrors)(nil), "google.logging.v2.WriteLogEntriesPartialErrors")
	proto.RegisterType((*ListLogEntriesRequest)(nil), "google.logging.v2.ListLogEntriesRequest")
	proto.RegisterType((*ListLogEntriesResponse)(nil), "google.logging.v2.ListLogEntriesResponse")
	proto.RegisterType((*ListMonitoredResourceDescriptorsRequest)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsRequest")
	proto.RegisterType((*ListMonitoredResourceDescriptorsResponse)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsResponse")
	proto.RegisterType((*ListLogsRequest)(nil), "google.logging.v2.ListLogsRequest")
	proto.RegisterType((*ListLogsResponse)(nil), "google.logging.v2.ListLogsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoggingServiceV2 service

type LoggingServiceV2Client interface {
	// Deletes all the log entries in a log.
	// The log reappears if it receives new entries.
	// Log entries written shortly before the delete operation might not be
	// deleted.
	DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*google_protobuf5.Empty, error)
	// Writes log entries to Stackdriver Logging.
	WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries from
	// Stackdriver Logging.  For ways to export log entries, see
	// [Exporting Logs](/logging/docs/export).
	ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error)
	// Lists the descriptors for monitored resource types used by Stackdriver
	// Logging.
	ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error)
	// Lists the logs in projects, organizations, folders, or billing accounts.
	// Only logs that have entries are listed.
	ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error)
}

type loggingServiceV2Client struct {
	cc *grpc.ClientConn
}

func NewLoggingServiceV2Client(cc *grpc.ClientConn) LoggingServiceV2Client {
	return &loggingServiceV2Client{cc}
}

func (c *loggingServiceV2Client) DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*google_protobuf5.Empty, error) {
	out := new(google_protobuf5.Empty)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/DeleteLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error) {
	out := new(WriteLogEntriesResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/WriteLogEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error) {
	out := new(ListLogEntriesResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListLogEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error) {
	out := new(ListMonitoredResourceDescriptorsResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error) {
	out := new(ListLogsResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoggingServiceV2 service

type LoggingServiceV2Server interface {
	// Deletes all the log entries in a log.
	// The log reappears if it receives new entries.
	// Log entries written shortly before the delete operation might not be
	// deleted.
	DeleteLog(context.Context, *DeleteLogRequest) (*google_protobuf5.Empty, error)
	// Writes log entries to Stackdriver Logging.
	WriteLogEntries(context.Context, *WriteLogEntriesRequest) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries from
	// Stackdriver Logging.  For ways to export log entries, see
	// [Exporting Logs](/logging/docs/export).
	ListLogEntries(context.Context, *ListLogEntriesRequest) (*ListLogEntriesResponse, error)
	// Lists the descriptors for monitored resource types used by Stackdriver
	// Logging.
	ListMonitoredResourceDescriptors(context.Context, *ListMonitoredResourceDescriptorsRequest) (*ListMonitoredResourceDescriptorsResponse, error)
	// Lists the logs in projects, organizations, folders, or billing accounts.
	// Only logs that have entries are listed.
	ListLogs(context.Context, *ListLogsRequest) (*ListLogsResponse, error)
}

func RegisterLoggingServiceV2Server(s *grpc.Server, srv LoggingServiceV2Server) {
	s.RegisterService(&_LoggingServiceV2_serviceDesc, srv)
}

func _LoggingServiceV2_DeleteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/DeleteLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, req.(*DeleteLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_WriteLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/WriteLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, req.(*WriteLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, req.(*ListLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMonitoredResourceDescriptorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, req.(*ListMonitoredResourceDescriptorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListLogs(ctx, req.(*ListLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoggingServiceV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.logging.v2.LoggingServiceV2",
	HandlerType: (*LoggingServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteLog",
			Handler:    _LoggingServiceV2_DeleteLog_Handler,
		},
		{
			MethodName: "WriteLogEntries",
			Handler:    _LoggingServiceV2_WriteLogEntries_Handler,
		},
		{
			MethodName: "ListLogEntries",
			Handler:    _LoggingServiceV2_ListLogEntries_Handler,
		},
		{
			MethodName: "ListMonitoredResourceDescriptors",
			Handler:    _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler,
		},
		{
			MethodName: "ListLogs",
			Handler:    _LoggingServiceV2_ListLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/logging/v2/logging.proto",
}

func init() { proto.RegisterFile("google/logging/v2/logging.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 975 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x06, 0xa5, 0xd8, 0x91, 0x46, 0x8d, 0xad, 0x6c, 0x62, 0x59, 0x91, 0x9c, 0x58, 0xa5, 0x9b,
	0x5a, 0x11, 0x10, 0x12, 0x55, 0x11, 0x20, 0x71, 0xd0, 0x8b, 0x13, 0xa3, 0x28, 0xe0, 0x14, 0x06,
	0xdd, 0x26, 0x40, 0x2e, 0x02, 0x25, 0x4d, 0x88, 0x6d, 0x28, 0x2e, 0xbb, 0xbb, 0x92, 0xab, 0x04,
	0xe9, 0x21, 0x87, 0xbe, 0x40, 0xdf, 0xa2, 0x87, 0xbe, 0x45, 0xaf, 0xbd, 0xf4, 0xd2, 0x43, 0x8f,
	0x79, 0x88, 0x1e, 0x0b, 0xee, 0x2e, 0x65, 0xea, 0x27, 0xb2, 0xdc, 0x9b, 0x76, 0xe6, 0xdb, 0x99,
	0xf9, 0x86, 0xdf, 0xcc, 0x0a, 0x76, 0x03, 0xc6, 0x82, 0x10, 0xdd, 0x90, 0x05, 0x01, 0x8d, 0x02,
	0x77, 0xd4, 0x4e, 0x7f, 0x3a, 0x31, 0x67, 0x92, 0x91, 0xeb, 0x1a, 0xe0, 0xa4, 0xd6, 0x51, 0xbb,
	0xb6, 0x63, 0xee, 0xf8, 0x31, 0x75, 0xfd, 0x28, 0x62, 0xd2, 0x97, 0x94, 0x45, 0x42, 0x5f, 0xa8,
	0xed, 0x65, 0xbc, 0x03, 0x16, 0x51, 0xc9, 0x38, 0xf6, 0x3b, 0x1c, 0x05, 0x1b, 0xf2, 0x1e, 0x1a,
	0xd0, 0xa7, 0x0b, 0xd3, 0x76, 0x30, 0x92, 0x7c, 0x6c, 0x20, 0x77, 0x0c, 0x44, 0x9d, 0xba, 0xc3,
	0x57, 0x6e, 0x7f, 0xc8, 0x55, 0x22, 0xe3, 0xaf, 0xcf, 0xfa, 0x71, 0x10, 0xcb, 0xf4, 0xf2, 0xee,
	0xac, 0x53, 0xd2, 0x01, 0x0a, 0xe9, 0x0f, 0x62, 0x03, 0xd8, 0x36, 0x00, 0x1e, 0xf7, 0x5c, 0x21,
	0x7d, 0x39, 0x34, 0xe5, 0xdb, 0xf7, 0xa1, 0xfc, 0x14, 0x43, 0x94, 0x78, 0xcc, 0x02, 0x0f, 0x7f,
	0x1c, 0xa2, 0x90, 0xe4, 0x16, 0x14, 0x92, 0xea, 0x22, 0x7f, 0x80, 0x55, 0xab, 0x61, 0x35, 0x8b,
	0xde, 0xd5, 0x90, 0x05, 0xdf, 0xfa, 0x03, 0xb4, 0xff, 0xce, 0x41, 0xe5, 0x05, 0xa7, 0x0a, 0x7e,
	0x14, 0x49, 0x4e, 0x51, 0x5c, 0x7c, 0x8b, 0x3c, 0x82, 0x42, 0xda, 0x90, 0x6a, 0xae, 0x61, 0x35,
	0x4b, 0xed, 0xdb, 0x8e, 0xe9, 0xb3, 0x1f, 0x53, 0xe7, 0x59, 0xda, 0x36, 0xcf, 0x80, 0xbc, 0x09,
	0x9c, 0x3c, 0x83, 0xf5, 0xd0, 0xef, 0x62, 0x28, 0xaa, 0xf9, 0x46, 0xbe, 0x59, 0x6a, 0x3f, 0x70,
	0xe6, 0x3e, 0x90, 0xb3, 0xb8, 0x20, 0xe7, 0x58, 0xdd, 0x4b, 0x8c, 0x63, 0xcf, 0x04, 0x21, 0x0f,
	0xe0, 0x2a, 0x6a, 0x54, 0xf5, 0x8a, 0x8a, 0x57, 0x5f, 0x10, 0xcf, 0x84, 0x1a, 0x7b, 0x29, 0x96,
	0xec, 0xc3, 0x66, 0xec, 0x73, 0x49, 0xfd, 0xb0, 0x23, 0x86, 0xbd, 0x1e, 0x0a, 0x51, 0x5d, 0x6b,
	0x58, 0xcd, 0x82, 0xb7, 0x61, 0xcc, 0xa7, 0xda, 0x5a, 0x7b, 0x04, 0xa5, 0x4c, 0x5a, 0x52, 0x86,
	0xfc, 0x6b, 0x1c, 0x9b, 0x76, 0x24, 0x3f, 0xc9, 0x4d, 0x58, 0x1b, 0xf9, 0xe1, 0x50, 0xf7, 0xa1,
	0xe8, 0xe9, 0xc3, 0x41, 0xee, 0xa1, 0x65, 0xdf, 0x82, 0xed, 0x39, 0x22, 0x22, 0x66, 0x91, 0x40,
	0xfb, 0x83, 0x05, 0x3b, 0x33, 0xbe, 0x13, 0x9d, 0xf7, 0x88, 0x73, 0xc6, 0x05, 0x19, 0x40, 0x79,
	0xa2, 0xa7, 0x0e, 0x2a, 0x5b, 0xd5, 0x52, 0xfc, 0x9e, 0x5c, 0xdc, 0xaf, 0xa9, 0x50, 0x13, 0xf2,
	0xfa, 0xa8, 0xfb, 0xb0, 0x11, 0x4e, 0x19, 0x6b, 0xdf, 0xc3, 0x8d, 0x05, 0xb0, 0x2c, 0xdb, 0x35,
	0xcd, 0xb6, 0x99, 0x65, 0x5b, 0x6a, 0x93, 0xb4, 0x18, 0x1e, 0xf7, 0x9c, 0x53, 0x25, 0xc3, 0x6c,
	0x07, 0xfe, 0xb4, 0x60, 0xeb, 0x98, 0x0a, 0x39, 0xaf, 0xad, 0x5d, 0x28, 0xc5, 0x9c, 0xfd, 0x80,
	0x3d, 0xd9, 0xa1, 0x7d, 0x4d, 0xad, 0xe8, 0x81, 0x31, 0x7d, 0xd3, 0x17, 0xe4, 0x2e, 0x6c, 0xa4,
	0x92, 0x51, 0x0a, 0x14, 0xd5, 0x82, 0xc2, 0x5c, 0x4b, 0xad, 0x89, 0x0e, 0x05, 0xa9, 0xc0, 0xfa,
	0x2b, 0x1a, 0x4a, 0xe4, 0xa6, 0xfd, 0xe6, 0x94, 0x68, 0x97, 0xf1, 0x3e, 0xf2, 0x4e, 0x77, 0x5c,
	0xcd, 0x6b, 0xed, 0xaa, 0xf3, 0xe1, 0x98, 0xd4, 0xa1, 0x18, 0xfb, 0x01, 0x76, 0x04, 0x7d, 0x83,
	0xd5, 0x2b, 0x8a, 0x5a, 0x21, 0x31, 0x9c, 0xd2, 0x37, 0x48, 0x6e, 0x03, 0x28, 0xa7, 0x64, 0xaf,
	0x31, 0x52, 0x92, 0x28, 0x7a, 0x0a, 0xfe, 0x5d, 0x62, 0xb0, 0xcf, 0xa0, 0x32, 0xcb, 0x47, 0x7f,
	0xd1, 0xac, 0x0e, 0xad, 0x4b, 0xe8, 0xf0, 0x73, 0xd8, 0x8c, 0xf0, 0x27, 0xd9, 0xc9, 0x24, 0xd5,
	0x44, 0xae, 0x25, 0xe6, 0x93, 0x49, 0x62, 0x84, 0xfd, 0x24, 0xf1, 0xdc, 0x60, 0x3d, 0x45, 0xd1,
	0xe3, 0x34, 0x96, 0x8c, 0x4f, 0x5a, 0x3b, 0xc5, 0xcf, 0x5a, 0xca, 0x2f, 0x37, 0xcb, 0xef, 0x77,
	0x0b, 0x9a, 0x17, 0xe7, 0x31, 0x94, 0x5f, 0xc2, 0xcd, 0xc9, 0x27, 0xea, 0x9f, 0xfb, 0x0d, 0xff,
	0xfd, 0xa5, 0x0b, 0xe1, 0x3c, 0x9e, 0x77, 0x83, 0xcf, 0xe7, 0xb8, 0x44, 0x5f, 0x36, 0xcd, 0x07,
	0x99, 0xf0, 0xaf, 0xc0, 0x7a, 0xec, 0x73, 0x8c, 0xa4, 0x99, 0x52, 0x73, 0x9a, 0xee, 0x4b, 0x6e,
	0x69, 0x5f, 0xf2, 0xb3, 0x7d, 0x79, 0x01, 0xe5, 0xf3, 0x34, 0x86, 0x7e, 0x1d, 0x8a, 0xe9, 0x7a,
	0xd4, 0xbb, 0xac, 0xe8, 0x15, 0xcc, 0x7e, 0x5c, 0xb9, 0xfe, 0xf6, 0x3f, 0x6b, 0x50, 0x3e, 0xd6,
	0x02, 0x39, 0x45, 0x3e, 0xa2, 0x3d, 0x7c, 0xde, 0x26, 0x67, 0x50, 0x9c, 0xac, 0x70, 0xb2, 0xb7,
	0x40, 0x47, 0xb3, 0x0b, 0xbe, 0x56, 0x49, 0x41, 0xe9, 0x7b, 0xe1, 0x1c, 0x25, 0x8f, 0x89, 0x7d,
	0xff, 0xfd, 0x5f, 0x1f, 0x7e, 0xcd, 0xed, 0xb7, 0xee, 0xba, 0xa3, 0x76, 0x17, 0xa5, 0xff, 0x85,
	0xfb, 0x36, 0xad, 0xf9, 0x2b, 0x33, 0x6c, 0xc2, 0x6d, 0x25, 0x4f, 0x97, 0x70, 0x5b, 0xef, 0xc8,
	0x2f, 0x16, 0x6c, 0xce, 0xec, 0x12, 0x72, 0x6f, 0xe5, 0xfd, 0x5c, 0x6b, 0xad, 0x02, 0x35, 0x1b,
	0x70, 0x47, 0x55, 0x56, 0xb1, 0xaf, 0x27, 0x4f, 0xa7, 0x99, 0x86, 0x83, 0xb3, 0x04, 0x7c, 0x60,
	0xb5, 0xc8, 0x7b, 0x0b, 0x36, 0xa6, 0x07, 0x8d, 0x34, 0x17, 0xcd, 0xd3, 0xa2, 0xdd, 0x52, 0xbb,
	0xb7, 0x02, 0xd2, 0x54, 0x51, 0x57, 0x55, 0x6c, 0xd9, 0xe5, 0x6c, 0x15, 0x21, 0x15, 0x32, 0x29,
	0xe2, 0x0f, 0x0b, 0x1a, 0x17, 0x0d, 0x03, 0x39, 0xf8, 0x48, 0xb2, 0x15, 0x26, 0xb5, 0xf6, 0xf8,
	0x7f, 0xdd, 0x35, 0xa5, 0x37, 0x55, 0xe9, 0x36, 0x69, 0x24, 0xa5, 0x0f, 0x96, 0x95, 0x38, 0x86,
	0x42, 0x2a, 0x5e, 0x62, 0x7f, 0xbc, 0x37, 0x93, 0xb2, 0xf6, 0x96, 0x62, 0x4c, 0xfa, 0xcf, 0x54,
	0xfa, 0x3b, 0x64, 0x27, 0x49, 0xff, 0x56, 0x8f, 0x58, 0x46, 0x52, 0xef, 0x94, 0xa6, 0x0e, 0x7f,
	0x86, 0xad, 0x1e, 0x1b, 0xcc, 0xc7, 0x3b, 0xfc, 0xc4, 0x88, 0xfe, 0x24, 0xd1, 0xeb, 0x89, 0xf5,
	0xf2, 0xa1, 0x81, 0x04, 0x2c, 0xf4, 0xa3, 0xc0, 0x61, 0x3c, 0x70, 0x03, 0x8c, 0x94, 0x9a, 0x5d,
	0xed, 0xf2, 0x63, 0x2a, 0x32, 0x7f, 0xb7, 0x1e, 0x9b, 0x9f, 0xff, 0x5a, 0xd6, 0x6f, 0xb9, 0xed,
	0xaf, 0xf5, 0xed, 0x27, 0x21, 0x1b, 0xf6, 0x1d, 0x13, 0xda, 0x79, 0xde, 0xee, 0xae, 0xab, 0x08,
	0x5f, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xc4, 0xaa, 0x91, 0x26, 0x0a, 0x00, 0x00,
}
