// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/duration.proto

package known_proto

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoiface "github.com/golang/protobuf/v2/runtime/protoiface"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

// A Duration represents a signed, fixed-length span of time represented
// as a count of seconds and fractions of seconds at nanosecond
// resolution. It is independent of any calendar and concepts like "day"
// or "month". It is related to Timestamp in that the difference between
// two Timestamp values is a Duration and it can be added or subtracted
// from a Timestamp. Range is approximately +-10,000 years.
//
// # Examples
//
// Example 1: Compute Duration from two Timestamps in pseudo code.
//
//     Timestamp start = ...;
//     Timestamp end = ...;
//     Duration duration = ...;
//
//     duration.seconds = end.seconds - start.seconds;
//     duration.nanos = end.nanos - start.nanos;
//
//     if (duration.seconds < 0 && duration.nanos > 0) {
//       duration.seconds += 1;
//       duration.nanos -= 1000000000;
//     } else if (durations.seconds > 0 && duration.nanos < 0) {
//       duration.seconds -= 1;
//       duration.nanos += 1000000000;
//     }
//
// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
//
//     Timestamp start = ...;
//     Duration duration = ...;
//     Timestamp end = ...;
//
//     end.seconds = start.seconds + duration.seconds;
//     end.nanos = start.nanos + duration.nanos;
//
//     if (end.nanos < 0) {
//       end.seconds -= 1;
//       end.nanos += 1000000000;
//     } else if (end.nanos >= 1000000000) {
//       end.seconds += 1;
//       end.nanos -= 1000000000;
//     }
//
// Example 3: Compute Duration from datetime.timedelta in Python.
//
//     td = datetime.timedelta(days=3, minutes=10)
//     duration = Duration()
//     duration.FromTimedelta(td)
//
// # JSON Mapping
//
// In JSON format, the Duration type is encoded as a string rather than an
// object, where the string ends in the suffix "s" (indicating seconds) and
// is preceded by the number of seconds, with nanoseconds expressed as
// fractional seconds. For example, 3 seconds with 0 nanoseconds should be
// encoded in JSON format as "3s", while 3 seconds and 1 nanosecond should
// be expressed in JSON format as "3.000000001s", and 3 seconds and 1
// microsecond should be expressed in JSON format as "3.000001s".
//
//
type Duration struct {
	// Signed seconds of the span of time. Must be from -315,576,000,000
	// to +315,576,000,000 inclusive. Note: these bounds are computed from:
	// 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Signed fractions of a second at nanosecond resolution of the span
	// of time. Durations less than one second are represented with a 0
	// `seconds` field and a positive or negative `nanos` field. For durations
	// of one second or more, a non-zero value for the `nanos` field must be
	// of the same sign as the `seconds` field. Must be from -999,999,999
	// to +999,999,999 inclusive.
	Nanos                int32                   `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *Duration) Reset() {
	*x = Duration{}
}

func (x *Duration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Duration) ProtoMessage() {}

func (x *Duration) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_duration_proto_msgTypes[0].MessageOf(x)
}

func (m *Duration) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_duration_proto_msgTypes[0].Methods()
}

// Deprecated: Use Duration.ProtoReflect.Type instead.
func (*Duration) Descriptor() ([]byte, []int) {
	return file_google_protobuf_duration_proto_rawDescGZIP(), []int{0}
}

func (*Duration) XXX_WellKnownType() string { return "Duration" }

func (x *Duration) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Duration) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

var File_google_protobuf_duration_proto protoreflect.FileDescriptor

var file_google_protobuf_duration_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x22, 0x3a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x42, 0x87, 0x01,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x0d, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x3b, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x03, 0x47, 0x50, 0x42, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_protobuf_duration_proto_rawDescOnce sync.Once
	file_google_protobuf_duration_proto_rawDescData = file_google_protobuf_duration_proto_rawDesc
)

func file_google_protobuf_duration_proto_rawDescGZIP() []byte {
	file_google_protobuf_duration_proto_rawDescOnce.Do(func() {
		file_google_protobuf_duration_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_duration_proto_rawDescData)
	})
	return file_google_protobuf_duration_proto_rawDescData
}

var file_google_protobuf_duration_proto_msgTypes = make([]protoimpl.MessageType, 1)
var file_google_protobuf_duration_proto_goTypes = []interface{}{
	(*Duration)(nil), // 0: google.protobuf.Duration
}
var file_google_protobuf_duration_proto_depIdxs = []int32{}

func init() { file_google_protobuf_duration_proto_init() }
func file_google_protobuf_duration_proto_init() {
	if File_google_protobuf_duration_proto != nil {
		return
	}
	File_google_protobuf_duration_proto = protoimpl.FileBuilder{
		RawDescriptor:      file_google_protobuf_duration_proto_rawDesc,
		GoTypes:            file_google_protobuf_duration_proto_goTypes,
		DependencyIndexes:  file_google_protobuf_duration_proto_depIdxs,
		MessageOutputTypes: file_google_protobuf_duration_proto_msgTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	file_google_protobuf_duration_proto_rawDesc = nil
	file_google_protobuf_duration_proto_goTypes = nil
	file_google_protobuf_duration_proto_depIdxs = nil
}
