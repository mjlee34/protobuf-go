// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2/proto2.proto

package proto2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoregistry "google.golang.org/protobuf/reflect/protoregistry"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type Message struct {
	I32                  *int32                  `protobuf:"varint,1,opt,name=i32" json:"i32,omitempty"`
	M                    *Message                `protobuf:"bytes,2,opt,name=m" json:"m,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *Message) Reset() {
	*x = Message{}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	return file_proto2_proto2_proto_msgTypes[0].MessageOf(x)
}

func (m *Message) XXX_Methods() *protoiface.Methods {
	return file_proto2_proto2_proto_msgTypes[0].Methods()
}

// Deprecated: Use Message.ProtoReflect.Type instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto2_proto2_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetI32() int32 {
	if x != nil && x.I32 != nil {
		return *x.I32
	}
	return 0
}

func (x *Message) GetM() *Message {
	if x != nil {
		return x.M
	}
	return nil
}

var File_proto2_proto2_proto protoreflect.FileDescriptor

var file_proto2_proto2_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x22, 0x49, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x33, 0x32, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x33, 0x32, 0x12, 0x2c, 0x0a, 0x01, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x01, 0x6d, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
}

var (
	file_proto2_proto2_proto_rawDescOnce sync.Once
	file_proto2_proto2_proto_rawDescData = file_proto2_proto2_proto_rawDesc
)

func file_proto2_proto2_proto_rawDescGZIP() []byte {
	file_proto2_proto2_proto_rawDescOnce.Do(func() {
		file_proto2_proto2_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto2_proto2_proto_rawDescData)
	})
	return file_proto2_proto2_proto_rawDescData
}

var file_proto2_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto2_proto2_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: goproto.protoc.proto2.Message
}
var file_proto2_proto2_proto_depIdxs = []int32{
	0, // goproto.protoc.proto2.Message.m:type_name -> goproto.protoc.proto2.Message
}

func init() { file_proto2_proto2_proto_init() }
func file_proto2_proto2_proto_init() {
	if File_proto2_proto2_proto != nil {
		return
	}
	File_proto2_proto2_proto = protoimpl.FileBuilder{
		RawDescriptor:      file_proto2_proto2_proto_rawDesc,
		GoTypes:            file_proto2_proto2_proto_goTypes,
		DependencyIndexes:  file_proto2_proto2_proto_depIdxs,
		MessageOutputTypes: file_proto2_proto2_proto_msgTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	file_proto2_proto2_proto_rawDesc = nil
	file_proto2_proto2_proto_goTypes = nil
	file_proto2_proto2_proto_depIdxs = nil
}
