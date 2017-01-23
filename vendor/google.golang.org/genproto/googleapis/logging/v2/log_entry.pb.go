// Code generated by protoc-gen-go.
// source: google/logging/v2/log_entry.proto
// DO NOT EDIT!

/*
Package logging is a generated protocol buffer package.

It is generated from these files:
	google/logging/v2/log_entry.proto
	google/logging/v2/logging.proto
	google/logging/v2/logging_config.proto
	google/logging/v2/logging_metrics.proto

It has these top-level messages:
	LogEntry
	LogEntryOperation
	LogEntrySourceLocation
	DeleteLogRequest
	WriteLogEntriesRequest
	WriteLogEntriesResponse
	ListLogEntriesRequest
	ListLogEntriesResponse
	ListMonitoredResourceDescriptorsRequest
	ListMonitoredResourceDescriptorsResponse
	ListLogsRequest
	ListLogsResponse
	LogSink
	ListSinksRequest
	ListSinksResponse
	GetSinkRequest
	CreateSinkRequest
	UpdateSinkRequest
	DeleteSinkRequest
	LogMetric
	ListLogMetricsRequest
	ListLogMetricsResponse
	GetLogMetricRequest
	CreateLogMetricRequest
	UpdateLogMetricRequest
	DeleteLogMetricRequest
*/
package logging

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_api3 "google.golang.org/genproto/googleapis/api/monitoredres"
import google_logging_type "google.golang.org/genproto/googleapis/logging/type"
import google_logging_type1 "google.golang.org/genproto/googleapis/logging/type"
import google_protobuf2 "github.com/golang/protobuf/ptypes/any"
import google_protobuf3 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf4 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An individual entry in a log.
type LogEntry struct {
	// Required. The resource name of the log to which this log entry belongs:
	//
	//     "projects/[PROJECT_ID]/logs/[LOG_ID]"
	//     "organizations/[ORGANIZATION_ID]/logs/[LOG_ID]"
	//
	// `[LOG_ID]` must be URL-encoded within `log_name`. Example:
	// `"organizations/1234567890/logs/cloudresourcemanager.googleapis.com%2Factivity"`.
	// `[LOG_ID]` must be less than 512 characters long and can only include the
	// following characters: upper and lower case alphanumeric characters,
	// forward-slash, underscore, hyphen, and period.
	//
	// For backward compatibility, if `log_name` begins with a forward-slash, such
	// as `/projects/...`, then the log entry is ingested as usual but the
	// forward-slash is removed. Listing the log entry will not show the leading
	// slash and filtering for a log name with a leading slash will never return
	// any results.
	LogName string `protobuf:"bytes,12,opt,name=log_name,json=logName" json:"log_name,omitempty"`
	// Required. The monitored resource associated with this log entry.
	// Example: a log entry that reports a database error would be
	// associated with the monitored resource designating the particular
	// database that reported the error.
	Resource *google_api3.MonitoredResource `protobuf:"bytes,8,opt,name=resource" json:"resource,omitempty"`
	// Optional. The log entry payload, which can be one of multiple types.
	//
	// Types that are valid to be assigned to Payload:
	//	*LogEntry_ProtoPayload
	//	*LogEntry_TextPayload
	//	*LogEntry_JsonPayload
	Payload isLogEntry_Payload `protobuf_oneof:"payload"`
	// Optional. The time the event described by the log entry occurred.  If
	// omitted, Stackdriver Logging will use the time the log entry is received.
	Timestamp *google_protobuf4.Timestamp `protobuf:"bytes,9,opt,name=timestamp" json:"timestamp,omitempty"`
	// Optional. The severity of the log entry. The default value is
	// `LogSeverity.DEFAULT`.
	Severity google_logging_type1.LogSeverity `protobuf:"varint,10,opt,name=severity,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// Optional. A unique ID for the log entry. If you provide this
	// field, the logging service considers other log entries in the
	// same project with the same ID as duplicates which can be removed.  If
	// omitted, Stackdriver Logging will generate a unique ID for this
	// log entry.
	InsertId string `protobuf:"bytes,4,opt,name=insert_id,json=insertId" json:"insert_id,omitempty"`
	// Optional. Information about the HTTP request associated with this
	// log entry, if applicable.
	HttpRequest *google_logging_type.HttpRequest `protobuf:"bytes,7,opt,name=http_request,json=httpRequest" json:"http_request,omitempty"`
	// Optional. A set of user-defined (key, value) data that provides additional
	// information about the log entry.
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Optional. Information about an operation associated with the log entry, if
	// applicable.
	Operation *LogEntryOperation `protobuf:"bytes,15,opt,name=operation" json:"operation,omitempty"`
	// Optional. Resource name of the trace associated with the log entry, if any.
	// If it contains a relative resource name, the name is assumed to be relative
	// to `//tracing.googleapis.com`. Example:
	// `projects/my-projectid/traces/06796866738c859f2f19b7cfb3214824`
	Trace string `protobuf:"bytes,22,opt,name=trace" json:"trace,omitempty"`
	// Optional. Source code location information associated with the log entry,
	// if any.
	SourceLocation *LogEntrySourceLocation `protobuf:"bytes,23,opt,name=source_location,json=sourceLocation" json:"source_location,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isLogEntry_Payload interface {
	isLogEntry_Payload()
}

type LogEntry_ProtoPayload struct {
	ProtoPayload *google_protobuf2.Any `protobuf:"bytes,2,opt,name=proto_payload,json=protoPayload,oneof"`
}
type LogEntry_TextPayload struct {
	TextPayload string `protobuf:"bytes,3,opt,name=text_payload,json=textPayload,oneof"`
}
type LogEntry_JsonPayload struct {
	JsonPayload *google_protobuf3.Struct `protobuf:"bytes,6,opt,name=json_payload,json=jsonPayload,oneof"`
}

func (*LogEntry_ProtoPayload) isLogEntry_Payload() {}
func (*LogEntry_TextPayload) isLogEntry_Payload()  {}
func (*LogEntry_JsonPayload) isLogEntry_Payload()  {}

func (m *LogEntry) GetPayload() isLogEntry_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *LogEntry) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *LogEntry) GetResource() *google_api3.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *LogEntry) GetProtoPayload() *google_protobuf2.Any {
	if x, ok := m.GetPayload().(*LogEntry_ProtoPayload); ok {
		return x.ProtoPayload
	}
	return nil
}

func (m *LogEntry) GetTextPayload() string {
	if x, ok := m.GetPayload().(*LogEntry_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (m *LogEntry) GetJsonPayload() *google_protobuf3.Struct {
	if x, ok := m.GetPayload().(*LogEntry_JsonPayload); ok {
		return x.JsonPayload
	}
	return nil
}

func (m *LogEntry) GetTimestamp() *google_protobuf4.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogEntry) GetSeverity() google_logging_type1.LogSeverity {
	if m != nil {
		return m.Severity
	}
	return google_logging_type1.LogSeverity_DEFAULT
}

func (m *LogEntry) GetInsertId() string {
	if m != nil {
		return m.InsertId
	}
	return ""
}

func (m *LogEntry) GetHttpRequest() *google_logging_type.HttpRequest {
	if m != nil {
		return m.HttpRequest
	}
	return nil
}

func (m *LogEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LogEntry) GetOperation() *LogEntryOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *LogEntry) GetTrace() string {
	if m != nil {
		return m.Trace
	}
	return ""
}

func (m *LogEntry) GetSourceLocation() *LogEntrySourceLocation {
	if m != nil {
		return m.SourceLocation
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LogEntry_OneofMarshaler, _LogEntry_OneofUnmarshaler, _LogEntry_OneofSizer, []interface{}{
		(*LogEntry_ProtoPayload)(nil),
		(*LogEntry_TextPayload)(nil),
		(*LogEntry_JsonPayload)(nil),
	}
}

func _LogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProtoPayload); err != nil {
			return err
		}
	case *LogEntry_TextPayload:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TextPayload)
	case *LogEntry_JsonPayload:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JsonPayload); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LogEntry.Payload has unexpected type %T", x)
	}
	return nil
}

func _LogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogEntry)
	switch tag {
	case 2: // payload.proto_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf2.Any)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_ProtoPayload{msg}
		return true, err
	case 3: // payload.text_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &LogEntry_TextPayload{x}
		return true, err
	case 6: // payload.json_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf3.Struct)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_JsonPayload{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		s := proto.Size(x.ProtoPayload)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LogEntry_TextPayload:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TextPayload)))
		n += len(x.TextPayload)
	case *LogEntry_JsonPayload:
		s := proto.Size(x.JsonPayload)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Additional information about a potentially long-running operation with which
// a log entry is associated.
type LogEntryOperation struct {
	// Optional. An arbitrary operation identifier. Log entries with the
	// same identifier are assumed to be part of the same operation.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Optional. An arbitrary producer identifier. The combination of
	// `id` and `producer` must be globally unique.  Examples for `producer`:
	// `"MyDivision.MyBigCompany.com"`, `"github.com/MyProject/MyApplication"`.
	Producer string `protobuf:"bytes,2,opt,name=producer" json:"producer,omitempty"`
	// Optional. Set this to True if this is the first log entry in the operation.
	First bool `protobuf:"varint,3,opt,name=first" json:"first,omitempty"`
	// Optional. Set this to True if this is the last log entry in the operation.
	Last bool `protobuf:"varint,4,opt,name=last" json:"last,omitempty"`
}

func (m *LogEntryOperation) Reset()                    { *m = LogEntryOperation{} }
func (m *LogEntryOperation) String() string            { return proto.CompactTextString(m) }
func (*LogEntryOperation) ProtoMessage()               {}
func (*LogEntryOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogEntryOperation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogEntryOperation) GetProducer() string {
	if m != nil {
		return m.Producer
	}
	return ""
}

func (m *LogEntryOperation) GetFirst() bool {
	if m != nil {
		return m.First
	}
	return false
}

func (m *LogEntryOperation) GetLast() bool {
	if m != nil {
		return m.Last
	}
	return false
}

// Additional information about the source code location that produced the log
// entry.
type LogEntrySourceLocation struct {
	// Optional. Source file name. Depending on the runtime environment, this
	// might be a simple name or a fully-qualified name.
	File string `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	// Optional. Line within the source file. 1-based; 0 indicates no line number
	// available.
	Line int64 `protobuf:"varint,2,opt,name=line" json:"line,omitempty"`
	// Optional. Human-readable name of the function or method being invoked, with
	// optional context such as the class or package name. This information may be
	// used in contexts such as the logs viewer, where a file and line number are
	// less meaningful. The format can vary by language. For example:
	// `qual.if.ied.Class.method` (Java), `dir/package.func` (Go), `function`
	// (Python).
	Function string `protobuf:"bytes,3,opt,name=function" json:"function,omitempty"`
}

func (m *LogEntrySourceLocation) Reset()                    { *m = LogEntrySourceLocation{} }
func (m *LogEntrySourceLocation) String() string            { return proto.CompactTextString(m) }
func (*LogEntrySourceLocation) ProtoMessage()               {}
func (*LogEntrySourceLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogEntrySourceLocation) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *LogEntrySourceLocation) GetLine() int64 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *LogEntrySourceLocation) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "google.logging.v2.LogEntry")
	proto.RegisterType((*LogEntryOperation)(nil), "google.logging.v2.LogEntryOperation")
	proto.RegisterType((*LogEntrySourceLocation)(nil), "google.logging.v2.LogEntrySourceLocation")
}

func init() { proto.RegisterFile("google/logging/v2/log_entry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0x97, 0x76, 0x5f, 0x97, 0xba, 0xdd, 0xf6, 0xcd, 0x1a, 0x5b, 0x56, 0x86, 0x28, 0x1b,
	0x82, 0x72, 0x93, 0x4a, 0xe5, 0x66, 0x63, 0x93, 0x10, 0x9d, 0x10, 0x43, 0x2a, 0x30, 0x79, 0x88,
	0x0b, 0x84, 0x54, 0x79, 0x89, 0x9b, 0x19, 0x52, 0x3b, 0x38, 0x4e, 0x45, 0xaf, 0x78, 0x04, 0xde,
	0x83, 0x27, 0xe4, 0x12, 0xf9, 0xd8, 0xe9, 0x4a, 0x3b, 0xed, 0xee, 0x9c, 0xfa, 0xf7, 0x3f, 0xff,
	0xe3, 0xe3, 0x93, 0xa2, 0x47, 0x89, 0x94, 0x49, 0xca, 0xba, 0xa9, 0x4c, 0x12, 0x2e, 0x92, 0xee,
	0xa4, 0x67, 0xc2, 0x21, 0x13, 0x5a, 0x4d, 0xc3, 0x4c, 0x49, 0x2d, 0xf1, 0x96, 0x45, 0x42, 0x87,
	0x84, 0x93, 0x5e, 0x6b, 0xdf, 0xa9, 0x68, 0xc6, 0xbb, 0x54, 0x08, 0xa9, 0xa9, 0xe6, 0x52, 0xe4,
	0x56, 0xd0, 0x3a, 0x9c, 0x3b, 0x1d, 0x4b, 0xc1, 0xb5, 0x54, 0x2c, 0x1e, 0x2a, 0x96, 0xcb, 0x42,
	0x45, 0xcc, 0x41, 0x4f, 0x16, 0x8c, 0xf5, 0x34, 0x63, 0xdd, 0x6b, 0xad, 0xb3, 0xa1, 0x62, 0xdf,
	0x0b, 0x96, 0xeb, 0xbb, 0x38, 0xd3, 0x62, 0xce, 0x26, 0x4c, 0x71, 0xed, 0xba, 0x6c, 0xed, 0x39,
	0x0e, 0xb2, 0xab, 0x62, 0xd4, 0xa5, 0xa2, 0x3c, 0xda, 0x5f, 0x3c, 0xca, 0xb5, 0x2a, 0xa2, 0xd2,
	0xe0, 0xe1, 0xe2, 0xa9, 0xe6, 0x63, 0x96, 0x6b, 0x3a, 0xce, 0x2c, 0x70, 0xf0, 0xab, 0x86, 0xfc,
	0x81, 0x4c, 0x5e, 0x9b, 0x91, 0xe0, 0x3d, 0xe4, 0x1b, 0x73, 0x41, 0xc7, 0x2c, 0x68, 0xb6, 0xbd,
	0x4e, 0x9d, 0xac, 0xa5, 0x32, 0x79, 0x4f, 0xc7, 0x0c, 0x1f, 0x23, 0xbf, 0xbc, 0x63, 0xe0, 0xb7,
	0xbd, 0x4e, 0xa3, 0xf7, 0x20, 0x74, 0xa3, 0xa3, 0x19, 0x0f, 0xdf, 0x95, 0x93, 0x20, 0x0e, 0x22,
	0x33, 0x1c, 0x9f, 0xa0, 0x75, 0xf0, 0x1a, 0x66, 0x74, 0x9a, 0x4a, 0x1a, 0x07, 0x15, 0xd0, 0x6f,
	0x97, 0xfa, 0xb2, 0xb7, 0xf0, 0x95, 0x98, 0x9e, 0xaf, 0x90, 0x26, 0xe4, 0x17, 0x96, 0xc5, 0x87,
	0xa8, 0xa9, 0xd9, 0x0f, 0x3d, 0xd3, 0x56, 0x4d, 0x5b, 0xe7, 0x2b, 0xa4, 0x61, 0x7e, 0x2d, 0xa1,
	0x53, 0xd4, 0xfc, 0x9a, 0x4b, 0x31, 0x83, 0x6a, 0x60, 0xb0, 0xbb, 0x64, 0x70, 0x09, 0xa3, 0x31,
	0x6a, 0x83, 0x97, 0xea, 0x23, 0x54, 0x9f, 0x4d, 0x25, 0xa8, 0x83, 0xb4, 0xb5, 0x24, 0xfd, 0x58,
	0x12, 0xe4, 0x06, 0xc6, 0xa7, 0xc8, 0x2f, 0x1f, 0x2a, 0x40, 0x6d, 0xaf, 0xb3, 0xd1, 0x6b, 0x87,
	0x0b, 0xfb, 0x64, 0x5e, 0x34, 0x1c, 0xc8, 0xe4, 0xd2, 0x71, 0x64, 0xa6, 0xc0, 0xf7, 0x51, 0x9d,
	0x8b, 0x9c, 0x29, 0x3d, 0xe4, 0x71, 0xb0, 0x0a, 0xe3, 0xf6, 0xed, 0x0f, 0x6f, 0x63, 0x7c, 0x86,
	0x9a, 0xf3, 0xfb, 0x12, 0xac, 0x41, 0x5f, 0xb7, 0x97, 0x3f, 0xd7, 0x3a, 0x23, 0x96, 0x23, 0x8d,
	0xeb, 0x9b, 0x04, 0xbf, 0x44, 0xb5, 0x94, 0x5e, 0xb1, 0x34, 0x0f, 0x1a, 0xed, 0x6a, 0xa7, 0xd1,
	0x7b, 0x1a, 0x2e, 0x6d, 0x7b, 0x58, 0x3e, 0x7e, 0x38, 0x00, 0x12, 0x62, 0xe2, 0x64, 0xb8, 0x8f,
	0xea, 0x32, 0x63, 0x0a, 0x3e, 0x80, 0x60, 0x13, 0x5a, 0x78, 0x7c, 0x47, 0x8d, 0x0f, 0x25, 0x4b,
	0x6e, 0x64, 0x78, 0x1b, 0xfd, 0xa7, 0x15, 0x8d, 0x58, 0xb0, 0x03, 0x57, 0xb4, 0x09, 0x26, 0x68,
	0xd3, 0xae, 0xc7, 0x30, 0x95, 0x91, 0xad, 0xbf, 0x0b, 0xf5, 0x9f, 0xdd, 0x51, 0xff, 0x12, 0x14,
	0x03, 0x27, 0x20, 0x1b, 0xf9, 0x3f, 0x79, 0xeb, 0x18, 0x35, 0xe6, 0x2e, 0x81, 0xff, 0x47, 0xd5,
	0x6f, 0x6c, 0x1a, 0x78, 0x60, 0x6b, 0x42, 0xd3, 0xca, 0x84, 0xa6, 0x05, 0x83, 0x0d, 0xac, 0x13,
	0x9b, 0xbc, 0xa8, 0x1c, 0x79, 0xfd, 0x3a, 0x5a, 0x73, 0xcb, 0x73, 0xc0, 0xd1, 0xd6, 0xd2, 0x7d,
	0xf0, 0x06, 0xaa, 0xf0, 0xd8, 0x95, 0xaa, 0xf0, 0x18, 0xb7, 0x90, 0x9f, 0x29, 0x19, 0x17, 0x11,
	0x53, 0xae, 0xd8, 0x2c, 0x37, 0x2e, 0x23, 0xae, 0x72, 0x0d, 0xbb, 0xea, 0x13, 0x9b, 0x60, 0x8c,
	0x56, 0x53, 0x9a, 0x6b, 0x78, 0x68, 0x9f, 0x40, 0x7c, 0xf0, 0x05, 0xed, 0xdc, 0x7e, 0x35, 0x43,
	0x8f, 0x78, 0xca, 0x9c, 0x23, 0xc4, 0x50, 0x81, 0x0b, 0xdb, 0x7c, 0x95, 0x40, 0x6c, 0xfa, 0x18,
	0x15, 0x22, 0x82, 0xf9, 0x55, 0x6d, 0x1f, 0x65, 0xde, 0xff, 0x89, 0xee, 0x45, 0x72, 0xbc, 0x3c,
	0xce, 0xfe, 0x7a, 0x69, 0x7a, 0x01, 0x5f, 0x9a, 0xf7, 0xf9, 0xc8, 0x31, 0x89, 0x4c, 0xa9, 0x48,
	0x42, 0xa9, 0x92, 0x6e, 0xc2, 0x04, 0xec, 0x7e, 0xd7, 0x1e, 0xd1, 0x8c, 0xe7, 0x73, 0x7f, 0xa3,
	0x27, 0x2e, 0xfc, 0xe3, 0x79, 0xbf, 0x2b, 0xbb, 0x6f, 0xac, 0xfa, 0x2c, 0x95, 0x45, 0x6c, 0xde,
	0x0a, 0x7c, 0x3e, 0xf5, 0xae, 0x6a, 0x50, 0xe1, 0xf9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34,
	0x32, 0x8a, 0x87, 0x87, 0x05, 0x00, 0x00,
}
