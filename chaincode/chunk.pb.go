// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: chunk.proto

package chaincode

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Chunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	mi := &file_chunk_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_chunk_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_chunk_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_chunk_proto protoreflect.FileDescriptor

var file_chunk_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1b, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chunk_proto_rawDescOnce sync.Once
	file_chunk_proto_rawDescData = file_chunk_proto_rawDesc
)

func file_chunk_proto_rawDescGZIP() []byte {
	file_chunk_proto_rawDescOnce.Do(func() {
		file_chunk_proto_rawDescData = protoimpl.X.CompressGZIP(file_chunk_proto_rawDescData)
	})
	return file_chunk_proto_rawDescData
}

var file_chunk_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chunk_proto_goTypes = []any{
	(*Chunk)(nil), // 0: chaincode.Chunk
}
var file_chunk_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chunk_proto_init() }
func file_chunk_proto_init() {
	if File_chunk_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chunk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chunk_proto_goTypes,
		DependencyIndexes: file_chunk_proto_depIdxs,
		MessageInfos:      file_chunk_proto_msgTypes,
	}.Build()
	File_chunk_proto = out.File
	file_chunk_proto_rawDesc = nil
	file_chunk_proto_goTypes = nil
	file_chunk_proto_depIdxs = nil
}
