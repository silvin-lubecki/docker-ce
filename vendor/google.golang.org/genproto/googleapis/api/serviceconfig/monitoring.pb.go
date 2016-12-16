// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/monitoring.proto
// DO NOT EDIT!

package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Monitoring configuration of the service.
//
// The example below shows how to configure monitored resources and metrics
// for monitoring. In the example, a monitored resource and two metrics are
// defined. The `library.googleapis.com/book/returned_count` metric is sent
// to both producer and consumer projects, whereas the
// `library.googleapis.com/book/overdue_count` metric is only sent to the
// consumer project.
//
//     monitored_resources:
//     - type: library.googleapis.com/branch
//       labels:
//       - key: /city
//         description: The city where the library branch is located in.
//       - key: /name
//         description: The name of the branch.
//     metrics:
//     - name: library.googleapis.com/book/returned_count
//       metric_kind: DELTA
//       value_type: INT64
//       labels:
//       - key: /customer_id
//     - name: library.googleapis.com/book/overdue_count
//       metric_kind: GAUGE
//       value_type: INT64
//       labels:
//       - key: /customer_id
//     monitoring:
//       producer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         metrics:
//         - library.googleapis.com/book/returned_count
//       consumer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         metrics:
//         - library.googleapis.com/book/returned_count
//         - library.googleapis.com/book/overdue_count
type Monitoring struct {
	// Monitoring configurations for sending metrics to the producer project.
	// There can be multiple producer destinations, each one must have a
	// different monitored resource type. A metric can be used in at most
	// one producer destination.
	ProducerDestinations []*Monitoring_MonitoringDestination `protobuf:"bytes,1,rep,name=producer_destinations,json=producerDestinations" json:"producer_destinations,omitempty"`
	// Monitoring configurations for sending metrics to the consumer project.
	// There can be multiple consumer destinations, each one must have a
	// different monitored resource type. A metric can be used in at most
	// one consumer destination.
	ConsumerDestinations []*Monitoring_MonitoringDestination `protobuf:"bytes,2,rep,name=consumer_destinations,json=consumerDestinations" json:"consumer_destinations,omitempty"`
}

func (m *Monitoring) Reset()                    { *m = Monitoring{} }
func (m *Monitoring) String() string            { return proto.CompactTextString(m) }
func (*Monitoring) ProtoMessage()               {}
func (*Monitoring) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *Monitoring) GetProducerDestinations() []*Monitoring_MonitoringDestination {
	if m != nil {
		return m.ProducerDestinations
	}
	return nil
}

func (m *Monitoring) GetConsumerDestinations() []*Monitoring_MonitoringDestination {
	if m != nil {
		return m.ConsumerDestinations
	}
	return nil
}

// Configuration of a specific monitoring destination (the producer project
// or the consumer project).
type Monitoring_MonitoringDestination struct {
	// The monitored resource type. The type must be defined in
	// [Service.monitored_resources][google.api.Service.monitored_resources] section.
	MonitoredResource string `protobuf:"bytes,1,opt,name=monitored_resource,json=monitoredResource" json:"monitored_resource,omitempty"`
	// Names of the metrics to report to this monitoring destination.
	// Each name must be defined in [Service.metrics][google.api.Service.metrics] section.
	Metrics []string `protobuf:"bytes,2,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *Monitoring_MonitoringDestination) Reset()         { *m = Monitoring_MonitoringDestination{} }
func (m *Monitoring_MonitoringDestination) String() string { return proto.CompactTextString(m) }
func (*Monitoring_MonitoringDestination) ProtoMessage()    {}
func (*Monitoring_MonitoringDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{0, 0}
}

func init() {
	proto.RegisterType((*Monitoring)(nil), "google.api.Monitoring")
	proto.RegisterType((*Monitoring_MonitoringDestination)(nil), "google.api.Monitoring.MonitoringDestination")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/monitoring.proto", fileDescriptor12)
}

var fileDescriptor12 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x69, 0x15, 0x65, 0x47, 0x50, 0x0c, 0x0e, 0xc6, 0xae, 0xc4, 0xab, 0x21, 0xda, 0x80,
	0x3e, 0x81, 0x43, 0xd0, 0x5d, 0x08, 0xa5, 0x2f, 0x30, 0x63, 0x7a, 0x0c, 0x81, 0x35, 0x67, 0x9c,
	0xa4, 0x3e, 0x90, 0xcf, 0xe0, 0x03, 0x9a, 0xad, 0xed, 0x5a, 0xc4, 0xab, 0xde, 0x84, 0xe4, 0x9c,
	0xff, 0xfc, 0xdf, 0x9f, 0x03, 0xaf, 0x86, 0xc8, 0x6c, 0x30, 0x33, 0xb4, 0x51, 0xce, 0x64, 0xc4,
	0x46, 0x1a, 0x74, 0x5b, 0xa6, 0x40, 0xb2, 0x69, 0xa9, 0xad, 0xf5, 0x32, 0x1e, 0xd2, 0x23, 0x7f,
	0x59, 0x8d, 0x9a, 0xdc, 0xa7, 0x35, 0xb2, 0x22, 0x67, 0x03, 0xb1, 0x8d, 0x43, 0x7b, 0xb5, 0x80,
	0xd6, 0x29, 0x4a, 0xe7, 0xab, 0xb1, 0xae, 0xca, 0x39, 0x0a, 0x2a, 0x58, 0x72, 0xbe, 0xb1, 0xbd,
	0xf9, 0x49, 0x01, 0xde, 0x0e, 0x2c, 0xa1, 0x60, 0x1a, 0xeb, 0x65, 0xad, 0x91, 0xd7, 0x25, 0xfa,
	0x60, 0x5d, 0xa3, 0x9e, 0x25, 0xd7, 0x47, 0x8b, 0xb3, 0x87, 0xbb, 0xac, 0x4f, 0x91, 0xf5, 0x63,
	0x83, 0xeb, 0x73, 0x3f, 0x54, 0x5c, 0x75, 0x56, 0x83, 0xa2, 0xdf, 0x21, 0x62, 0x1a, 0x5f, 0x57,
	0x7f, 0x11, 0xe9, 0x18, 0x44, 0x67, 0x35, 0x44, 0xcc, 0xdf, 0x61, 0xfa, 0xaf, 0x5c, 0xdc, 0x83,
	0x68, 0x17, 0x8b, 0xe5, 0x9a, 0xd1, 0x53, 0xcd, 0x1a, 0xe3, 0xdf, 0x92, 0xc5, 0xa4, 0xb8, 0x3c,
	0x74, 0x8a, 0xb6, 0x21, 0x66, 0x70, 0x5a, 0x61, 0x60, 0xab, 0x9b, 0x70, 0x93, 0xa2, 0x7b, 0x2e,
	0x6f, 0xe1, 0x5c, 0x53, 0x35, 0x88, 0xba, 0xbc, 0xe8, 0x89, 0xf9, 0x6e, 0xb3, 0x79, 0xf2, 0x9d,
	0x1e, 0xbf, 0x3c, 0xe5, 0xab, 0x8f, 0x93, 0xfd, 0xa6, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x1a, 0x02, 0x76, 0xbb, 0x0c, 0x02, 0x00, 0x00,
}
