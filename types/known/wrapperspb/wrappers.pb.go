// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/wrappers.proto

package wrapperspb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoregistry "google.golang.org/protobuf/reflect/protoregistry"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

// Wrapper message for `double`.
//
// The JSON representation for `DoubleValue` is JSON number.
type DoubleValue struct {
	// The double value.
	Value                float64                 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *DoubleValue) Reset() {
	*x = DoubleValue{}
}

func (x *DoubleValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleValue) ProtoMessage() {}

func (x *DoubleValue) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_wrappers_proto_msgTypes[0].MessageOf(x)
}

func (m *DoubleValue) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_wrappers_proto_msgTypes[0].Methods()
}

// Deprecated: Use DoubleValue.ProtoReflect.Type instead.
func (*DoubleValue) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{0}
}

func (*DoubleValue) XXX_WellKnownType() string { return "DoubleValue" }

func (x *DoubleValue) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `float`.
//
// The JSON representation for `FloatValue` is JSON number.
type FloatValue struct {
	// The float value.
	Value                float32                 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *FloatValue) Reset() {
	*x = FloatValue{}
}

func (x *FloatValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatValue) ProtoMessage() {}

func (x *FloatValue) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_wrappers_proto_msgTypes[1].MessageOf(x)
}

func (m *FloatValue) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_wrappers_proto_msgTypes[1].Methods()
}

// Deprecated: Use FloatValue.ProtoReflect.Type instead.
func (*FloatValue) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{1}
}

func (*FloatValue) XXX_WellKnownType() string { return "FloatValue" }

func (x *FloatValue) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `int64`.
//
// The JSON representation for `Int64Value` is JSON string.
type Int64Value struct {
	// The int64 value.
	Value                int64                   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *Int64Value) Reset() {
	*x = Int64Value{}
}

func (x *Int64Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Value) ProtoMessage() {}

func (x *Int64Value) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_wrappers_proto_msgTypes[2].MessageOf(x)
}

func (m *Int64Value) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_wrappers_proto_msgTypes[2].Methods()
}

// Deprecated: Use Int64Value.ProtoReflect.Type instead.
func (*Int64Value) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{2}
}

func (*Int64Value) XXX_WellKnownType() string { return "Int64Value" }

func (x *Int64Value) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `uint64`.
//
// The JSON representation for `UInt64Value` is JSON string.
type UInt64Value struct {
	// The uint64 value.
	Value                uint64                  `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *UInt64Value) Reset() {
	*x = UInt64Value{}
}

func (x *UInt64Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt64Value) ProtoMessage() {}

func (x *UInt64Value) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_wrappers_proto_msgTypes[3].MessageOf(x)
}

func (m *UInt64Value) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_wrappers_proto_msgTypes[3].Methods()
}

// Deprecated: Use UInt64Value.ProtoReflect.Type instead.
func (*UInt64Value) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{3}
}

func (*UInt64Value) XXX_WellKnownType() string { return "UInt64Value" }

func (x *UInt64Value) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `int32`.
//
// The JSON representation for `Int32Value` is JSON number.
type Int32Value struct {
	// The int32 value.
	Value                int32                   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *Int32Value) Reset() {
	*x = Int32Value{}
}

func (x *Int32Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Value) ProtoMessage() {}

func (x *Int32Value) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_wrappers_proto_msgTypes[4].MessageOf(x)
}

func (m *Int32Value) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_wrappers_proto_msgTypes[4].Methods()
}

// Deprecated: Use Int32Value.ProtoReflect.Type instead.
func (*Int32Value) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{4}
}

func (*Int32Value) XXX_WellKnownType() string { return "Int32Value" }

func (x *Int32Value) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `uint32`.
//
// The JSON representation for `UInt32Value` is JSON number.
type UInt32Value struct {
	// The uint32 value.
	Value                uint32                  `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *UInt32Value) Reset() {
	*x = UInt32Value{}
}

func (x *UInt32Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt32Value) ProtoMessage() {}

func (x *UInt32Value) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_wrappers_proto_msgTypes[5].MessageOf(x)
}

func (m *UInt32Value) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_wrappers_proto_msgTypes[5].Methods()
}

// Deprecated: Use UInt32Value.ProtoReflect.Type instead.
func (*UInt32Value) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{5}
}

func (*UInt32Value) XXX_WellKnownType() string { return "UInt32Value" }

func (x *UInt32Value) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `bool`.
//
// The JSON representation for `BoolValue` is JSON `true` and `false`.
type BoolValue struct {
	// The bool value.
	Value                bool                    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *BoolValue) Reset() {
	*x = BoolValue{}
}

func (x *BoolValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolValue) ProtoMessage() {}

func (x *BoolValue) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_wrappers_proto_msgTypes[6].MessageOf(x)
}

func (m *BoolValue) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_wrappers_proto_msgTypes[6].Methods()
}

// Deprecated: Use BoolValue.ProtoReflect.Type instead.
func (*BoolValue) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{6}
}

func (*BoolValue) XXX_WellKnownType() string { return "BoolValue" }

func (x *BoolValue) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

// Wrapper message for `string`.
//
// The JSON representation for `StringValue` is JSON string.
type StringValue struct {
	// The string value.
	Value                string                  `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *StringValue) Reset() {
	*x = StringValue{}
}

func (x *StringValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringValue) ProtoMessage() {}

func (x *StringValue) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_wrappers_proto_msgTypes[7].MessageOf(x)
}

func (m *StringValue) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_wrappers_proto_msgTypes[7].Methods()
}

// Deprecated: Use StringValue.ProtoReflect.Type instead.
func (*StringValue) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{7}
}

func (*StringValue) XXX_WellKnownType() string { return "StringValue" }

func (x *StringValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Wrapper message for `bytes`.
//
// The JSON representation for `BytesValue` is JSON string.
type BytesValue struct {
	// The bytes value.
	Value                []byte                  `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *BytesValue) Reset() {
	*x = BytesValue{}
}

func (x *BytesValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesValue) ProtoMessage() {}

func (x *BytesValue) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_wrappers_proto_msgTypes[8].MessageOf(x)
}

func (m *BytesValue) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_wrappers_proto_msgTypes[8].Methods()
}

// Deprecated: Use BytesValue.ProtoReflect.Type instead.
func (*BytesValue) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{8}
}

func (*BytesValue) XXX_WellKnownType() string { return "BytesValue" }

func (x *BytesValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_google_protobuf_wrappers_proto protoreflect.FileDescriptor

var file_google_protobuf_wrappers_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x22, 0x23, 0x0a, 0x0b, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23,
	0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x09,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x23, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x83, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x42, 0x0d, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x42, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_protobuf_wrappers_proto_rawDescOnce sync.Once
	file_google_protobuf_wrappers_proto_rawDescData = file_google_protobuf_wrappers_proto_rawDesc
)

func file_google_protobuf_wrappers_proto_rawDescGZIP() []byte {
	file_google_protobuf_wrappers_proto_rawDescOnce.Do(func() {
		file_google_protobuf_wrappers_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_wrappers_proto_rawDescData)
	})
	return file_google_protobuf_wrappers_proto_rawDescData
}

var file_google_protobuf_wrappers_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_google_protobuf_wrappers_proto_goTypes = []interface{}{
	(*DoubleValue)(nil), // 0: google.protobuf.DoubleValue
	(*FloatValue)(nil),  // 1: google.protobuf.FloatValue
	(*Int64Value)(nil),  // 2: google.protobuf.Int64Value
	(*UInt64Value)(nil), // 3: google.protobuf.UInt64Value
	(*Int32Value)(nil),  // 4: google.protobuf.Int32Value
	(*UInt32Value)(nil), // 5: google.protobuf.UInt32Value
	(*BoolValue)(nil),   // 6: google.protobuf.BoolValue
	(*StringValue)(nil), // 7: google.protobuf.StringValue
	(*BytesValue)(nil),  // 8: google.protobuf.BytesValue
}
var file_google_protobuf_wrappers_proto_depIdxs = []int32{}

func init() { file_google_protobuf_wrappers_proto_init() }
func file_google_protobuf_wrappers_proto_init() {
	if File_google_protobuf_wrappers_proto != nil {
		return
	}
	File_google_protobuf_wrappers_proto = protoimpl.FileBuilder{
		RawDescriptor:      file_google_protobuf_wrappers_proto_rawDesc,
		GoTypes:            file_google_protobuf_wrappers_proto_goTypes,
		DependencyIndexes:  file_google_protobuf_wrappers_proto_depIdxs,
		MessageOutputTypes: file_google_protobuf_wrappers_proto_msgTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	file_google_protobuf_wrappers_proto_rawDesc = nil
	file_google_protobuf_wrappers_proto_goTypes = nil
	file_google_protobuf_wrappers_proto_depIdxs = nil
}
