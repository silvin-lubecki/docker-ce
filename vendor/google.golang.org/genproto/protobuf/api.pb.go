// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/protobuf/api.proto
// DO NOT EDIT!

/*
Package descriptor is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/protobuf/api.proto
	google.golang.org/genproto/protobuf/descriptor.proto
	google.golang.org/genproto/protobuf/field_mask.proto
	google.golang.org/genproto/protobuf/source_context.proto
	google.golang.org/genproto/protobuf/type.proto

It has these top-level messages:
	Api
	Method
	Mixin
	FileDescriptorSet
	FileDescriptorProto
	DescriptorProto
	FieldDescriptorProto
	OneofDescriptorProto
	EnumDescriptorProto
	EnumValueDescriptorProto
	ServiceDescriptorProto
	MethodDescriptorProto
	FileOptions
	MessageOptions
	FieldOptions
	OneofOptions
	EnumOptions
	EnumValueOptions
	ServiceOptions
	MethodOptions
	UninterpretedOption
	SourceCodeInfo
	GeneratedCodeInfo
	FieldMask
	SourceContext
	Type
	Field
	Enum
	EnumValue
	Option
*/
package descriptor // import "google.golang.org/genproto/protobuf"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Api is a light-weight descriptor for a protocol buffer service.
type Api struct {
	// The fully qualified name of this api, including package name
	// followed by the api's simple name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The methods of this api, in unspecified order.
	Methods []*Method `protobuf:"bytes,2,rep,name=methods" json:"methods,omitempty"`
	// Any metadata attached to the API.
	Options []*Option `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
	// A version string for this api. If specified, must have the form
	// `major-version.minor-version`, as in `1.10`. If the minor version
	// is omitted, it defaults to zero. If the entire version field is
	// empty, the major version is derived from the package name, as
	// outlined below. If the field is not empty, the version in the
	// package name will be verified to be consistent with what is
	// provided here.
	//
	// The versioning schema uses [semantic
	// versioning](http://semver.org) where the major version number
	// indicates a breaking change and the minor version an additive,
	// non-breaking change. Both version numbers are signals to users
	// what to expect from different versions, and should be carefully
	// chosen based on the product plan.
	//
	// The major version is also reflected in the package name of the
	// API, which must end in `v<major-version>`, as in
	// `google.feature.v1`. For major versions 0 and 1, the suffix can
	// be omitted. Zero major versions must only be used for
	// experimental, none-GA apis.
	//
	//
	Version string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	// Source context for the protocol buffer service represented by this
	// message.
	SourceContext *SourceContext `protobuf:"bytes,5,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	// Included APIs. See [Mixin][].
	Mixins []*Mixin `protobuf:"bytes,6,rep,name=mixins" json:"mixins,omitempty"`
	// The source syntax of the service.
	Syntax Syntax `protobuf:"varint,7,opt,name=syntax,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
}

func (m *Api) Reset()                    { *m = Api{} }
func (m *Api) String() string            { return proto.CompactTextString(m) }
func (*Api) ProtoMessage()               {}
func (*Api) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Api) GetMethods() []*Method {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *Api) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Api) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

func (m *Api) GetMixins() []*Mixin {
	if m != nil {
		return m.Mixins
	}
	return nil
}

// Method represents a method of an api.
type Method struct {
	// The simple name of this method.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A URL of the input message type.
	RequestTypeUrl string `protobuf:"bytes,2,opt,name=request_type_url,json=requestTypeUrl" json:"request_type_url,omitempty"`
	// If true, the request is streamed.
	RequestStreaming bool `protobuf:"varint,3,opt,name=request_streaming,json=requestStreaming" json:"request_streaming,omitempty"`
	// The URL of the output message type.
	ResponseTypeUrl string `protobuf:"bytes,4,opt,name=response_type_url,json=responseTypeUrl" json:"response_type_url,omitempty"`
	// If true, the response is streamed.
	ResponseStreaming bool `protobuf:"varint,5,opt,name=response_streaming,json=responseStreaming" json:"response_streaming,omitempty"`
	// Any metadata attached to the method.
	Options []*Option `protobuf:"bytes,6,rep,name=options" json:"options,omitempty"`
	// The source syntax of this method.
	Syntax Syntax `protobuf:"varint,7,opt,name=syntax,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
}

func (m *Method) Reset()                    { *m = Method{} }
func (m *Method) String() string            { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()               {}
func (*Method) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Method) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

// Declares an API to be included in this API. The including API must
// redeclare all the methods from the included API, but documentation
// and options are inherited as follows:
//
// - If after comment and whitespace stripping, the documentation
//   string of the redeclared method is empty, it will be inherited
//   from the original method.
//
// - Each annotation belonging to the service config (http,
//   visibility) which is not set in the redeclared method will be
//   inherited.
//
// - If an http annotation is inherited, the path pattern will be
//   modified as follows. Any version prefix will be replaced by the
//   version of the including API plus the [root][] path if specified.
//
// Example of a simple mixin:
//
//     package google.acl.v1;
//     service AccessControl {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v1/{resource=**}:getAcl";
//       }
//     }
//
//     package google.storage.v2;
//     service Storage {
//       rpc GetAcl(GetAclRequest) returns (Acl);
//
//       // Get a data record.
//       rpc GetData(GetDataRequest) returns (Data) {
//         option (google.api.http).get = "/v2/{resource=**}";
//       }
//     }
//
// Example of a mixin configuration:
//
//     apis:
//     - name: google.storage.v2.Storage
//       mixins:
//       - name: google.acl.v1.AccessControl
//
// The mixin construct implies that all methods in `AccessControl` are
// also declared with same name and request/response types in
// `Storage`. A documentation generator or annotation processor will
// see the effective `Storage.GetAcl` method after inherting
// documentation and annotations as follows:
//
//     service Storage {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v2/{resource=**}:getAcl";
//       }
//       ...
//     }
//
// Note how the version in the path pattern changed from `v1` to `v2`.
//
// If the `root` field in the mixin is specified, it should be a
// relative path under which inherited HTTP paths are placed. Example:
//
//     apis:
//     - name: google.storage.v2.Storage
//       mixins:
//       - name: google.acl.v1.AccessControl
//         root: acls
//
// This implies the following inherited HTTP annotation:
//
//     service Storage {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v2/acls/{resource=**}:getAcl";
//       }
//       ...
//     }
type Mixin struct {
	// The fully qualified name of the API which is included.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// If non-empty specifies a path under which inherited HTTP paths
	// are rooted.
	Root string `protobuf:"bytes,2,opt,name=root" json:"root,omitempty"`
}

func (m *Mixin) Reset()                    { *m = Mixin{} }
func (m *Mixin) String() string            { return proto.CompactTextString(m) }
func (*Mixin) ProtoMessage()               {}
func (*Mixin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Api)(nil), "google.protobuf.Api")
	proto.RegisterType((*Method)(nil), "google.protobuf.Method")
	proto.RegisterType((*Mixin)(nil), "google.protobuf.Mixin")
}

func init() { proto.RegisterFile("google.golang.org/genproto/protobuf/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x4f, 0xe2, 0x40,
	0x14, 0x4f, 0x5b, 0x28, 0xec, 0x90, 0x85, 0xdd, 0xd9, 0x64, 0xb7, 0xe1, 0x40, 0x08, 0xa7, 0x66,
	0x37, 0xb4, 0x59, 0xbc, 0x78, 0x15, 0x63, 0x38, 0x10, 0x63, 0x53, 0x34, 0x1e, 0x49, 0xc1, 0xb1,
	0x36, 0x69, 0x67, 0xea, 0xcc, 0x54, 0xe1, 0xdb, 0x18, 0x8f, 0x1e, 0xfd, 0x06, 0x7e, 0x33, 0xa7,
	0xd3, 0x0e, 0x20, 0x60, 0x82, 0x97, 0x66, 0xde, 0xfb, 0xfd, 0x79, 0xf3, 0x7e, 0x53, 0xd0, 0x0f,
	0x09, 0x09, 0x63, 0xe4, 0x84, 0x24, 0x0e, 0x70, 0xe8, 0x10, 0x1a, 0xba, 0x21, 0xc2, 0x29, 0x25,
	0x9c, 0xb8, 0xf2, 0x3b, 0xcb, 0x6e, 0xdd, 0x20, 0x8d, 0x1c, 0x59, 0xc0, 0x56, 0x49, 0x57, 0x50,
	0xfb, 0xf8, 0x10, 0x3d, 0x23, 0x19, 0x9d, 0xa3, 0xe9, 0x9c, 0x60, 0x8e, 0x16, 0xbc, 0x10, 0xb7,
	0x9d, 0x43, 0x94, 0x7c, 0x99, 0x96, 0xc3, 0x7a, 0x6f, 0x3a, 0x30, 0x4e, 0xd2, 0x08, 0x42, 0x50,
	0xc1, 0x41, 0x82, 0x2c, 0xad, 0xab, 0xd9, 0xdf, 0x7c, 0x79, 0x86, 0xff, 0x41, 0x2d, 0x41, 0xfc,
	0x8e, 0xdc, 0x30, 0x4b, 0xef, 0x1a, 0x76, 0x63, 0xf0, 0xc7, 0xd9, 0xba, 0xa8, 0x73, 0x2e, 0x71,
	0x5f, 0xf1, 0x72, 0x09, 0x49, 0x79, 0x44, 0x30, 0xb3, 0x8c, 0x4f, 0x24, 0x17, 0x12, 0xf7, 0x15,
	0x0f, 0x5a, 0xa0, 0xf6, 0x80, 0x28, 0x13, 0x67, 0xab, 0x22, 0x87, 0xab, 0x12, 0x9e, 0x81, 0xe6,
	0xc7, 0x1d, 0xad, 0xaa, 0x20, 0x34, 0x06, 0x9d, 0x1d, 0xcf, 0x89, 0xa4, 0x9d, 0x16, 0x2c, 0xff,
	0x3b, 0xdb, 0x2c, 0xa1, 0x03, 0xcc, 0x24, 0x5a, 0x44, 0xe2, 0x4a, 0xa6, 0xbc, 0xd2, 0xef, 0xdd,
	0x2d, 0x72, 0xd8, 0x2f, 0x59, 0xd0, 0x05, 0x26, 0x5b, 0x62, 0x1e, 0x2c, 0xac, 0x9a, 0x18, 0xd7,
	0xdc, 0xb3, 0xc2, 0x44, 0xc2, 0x7e, 0x49, 0xeb, 0xbd, 0xea, 0xc0, 0x2c, 0x82, 0xd8, 0x1b, 0xa3,
	0x0d, 0x7e, 0x50, 0x74, 0x9f, 0x21, 0xc6, 0xa7, 0x79, 0xf0, 0xd3, 0x8c, 0xc6, 0x22, 0xcf, 0x1c,
	0x6f, 0x96, 0xfd, 0x4b, 0xd1, 0xbe, 0xa2, 0x31, 0xfc, 0x07, 0x7e, 0x2a, 0x26, 0xe3, 0x14, 0x05,
	0x49, 0x84, 0x43, 0x91, 0xa3, 0x66, 0xd7, 0x7d, 0x65, 0x31, 0x51, 0x7d, 0xf8, 0x37, 0x27, 0xb3,
	0x54, 0x44, 0x88, 0xd6, 0xbe, 0x45, 0x82, 0x2d, 0x05, 0x28, 0xe3, 0x3e, 0x80, 0x2b, 0xee, 0xda,
	0xb9, 0x2a, 0x9d, 0x57, 0x2e, 0x6b, 0xeb, 0x8d, 0x57, 0x34, 0x0f, 0x7c, 0xc5, 0x2f, 0x87, 0xe6,
	0x82, 0xaa, 0x8c, 0x7d, 0x6f, 0x64, 0xa2, 0x47, 0x09, 0xe1, 0x65, 0x4c, 0xf2, 0x3c, 0x1c, 0x83,
	0x5f, 0x73, 0x92, 0x6c, 0xdb, 0x0e, 0xeb, 0xe2, 0xef, 0xf5, 0xf2, 0xc2, 0xd3, 0x9e, 0x34, 0xed,
	0x59, 0x37, 0x46, 0xde, 0xf0, 0x45, 0xef, 0x8c, 0x0a, 0x9a, 0xa7, 0xa6, 0x5f, 0xa3, 0x38, 0x1e,
	0x63, 0xf2, 0x88, 0xf3, 0x48, 0xd8, 0xcc, 0x94, 0xfa, 0xa3, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x97, 0x07, 0xcf, 0x1c, 0xa9, 0x03, 0x00, 0x00,
}
