// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/sub2/a.proto

package sub2

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoiface "github.com/golang/protobuf/v2/runtime/protoiface"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type Sub2Message struct {
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *Sub2Message) Reset() {
	*x = Sub2Message{}
}

func (x *Sub2Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sub2Message) ProtoMessage() {}

func (x *Sub2Message) ProtoReflect() protoreflect.Message {
	return file_import_public_sub2_a_proto_msgTypes[0].MessageOf(x)
}

func (m *Sub2Message) XXX_Methods() *protoiface.Methods {
	return file_import_public_sub2_a_proto_msgTypes[0].Methods()
}

// Deprecated: Use Sub2Message.ProtoReflect.Type instead.
func (*Sub2Message) Descriptor() ([]byte, []int) {
	return file_import_public_sub2_a_proto_rawDescGZIP(), []int{0}
}

var File_import_public_sub2_a_proto protoreflect.FileDescriptor

var file_import_public_sub2_a_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x73, 0x75, 0x62, 0x32, 0x2f, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x75, 0x62, 0x32, 0x22,
	0x0d, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x4d,
	0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f,
	0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x75, 0x62, 0x32,
}

var (
	file_import_public_sub2_a_proto_rawDescOnce sync.Once
	file_import_public_sub2_a_proto_rawDescData = file_import_public_sub2_a_proto_rawDesc
)

func file_import_public_sub2_a_proto_rawDescGZIP() []byte {
	file_import_public_sub2_a_proto_rawDescOnce.Do(func() {
		file_import_public_sub2_a_proto_rawDescData = protoimpl.X.CompressGZIP(file_import_public_sub2_a_proto_rawDescData)
	})
	return file_import_public_sub2_a_proto_rawDescData
}

var file_import_public_sub2_a_proto_msgTypes = make([]protoimpl.MessageType, 1)
var file_import_public_sub2_a_proto_goTypes = []interface{}{
	(*Sub2Message)(nil), // 0: goproto.protoc.import_public.sub2.Sub2Message
}
var file_import_public_sub2_a_proto_depIdxs = []int32{}

func init() { file_import_public_sub2_a_proto_init() }
func file_import_public_sub2_a_proto_init() {
	if File_import_public_sub2_a_proto != nil {
		return
	}
	File_import_public_sub2_a_proto = protoimpl.FileBuilder{
		RawDescriptor:      file_import_public_sub2_a_proto_rawDesc,
		GoTypes:            file_import_public_sub2_a_proto_goTypes,
		DependencyIndexes:  file_import_public_sub2_a_proto_depIdxs,
		MessageOutputTypes: file_import_public_sub2_a_proto_msgTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	file_import_public_sub2_a_proto_rawDesc = nil
	file_import_public_sub2_a_proto_goTypes = nil
	file_import_public_sub2_a_proto_depIdxs = nil
}
