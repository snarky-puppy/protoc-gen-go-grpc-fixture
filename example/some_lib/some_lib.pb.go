// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: some_lib/some_lib.proto

package some_lib

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

type SomeObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SomeOtherField string `protobuf:"bytes,1,opt,name=some_other_field,json=someOtherField,proto3" json:"some_other_field,omitempty"`
}

func (x *SomeObject) Reset() {
	*x = SomeObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_some_lib_some_lib_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeObject) ProtoMessage() {}

func (x *SomeObject) ProtoReflect() protoreflect.Message {
	mi := &file_some_lib_some_lib_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeObject.ProtoReflect.Descriptor instead.
func (*SomeObject) Descriptor() ([]byte, []int) {
	return file_some_lib_some_lib_proto_rawDescGZIP(), []int{0}
}

func (x *SomeObject) GetSomeOtherField() string {
	if x != nil {
		return x.SomeOtherField
	}
	return ""
}

type SomeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SomeField  string      `protobuf:"bytes,1,opt,name=some_field,json=someField,proto3" json:"some_field,omitempty"`
	SomeObject *SomeObject `protobuf:"bytes,2,opt,name=some_object,json=someObject,proto3" json:"some_object,omitempty"`
}

func (x *SomeMessage) Reset() {
	*x = SomeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_some_lib_some_lib_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeMessage) ProtoMessage() {}

func (x *SomeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_some_lib_some_lib_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeMessage.ProtoReflect.Descriptor instead.
func (*SomeMessage) Descriptor() ([]byte, []int) {
	return file_some_lib_some_lib_proto_rawDescGZIP(), []int{1}
}

func (x *SomeMessage) GetSomeField() string {
	if x != nil {
		return x.SomeField
	}
	return ""
}

func (x *SomeMessage) GetSomeObject() *SomeObject {
	if x != nil {
		return x.SomeObject
	}
	return nil
}

var File_some_lib_some_lib_proto protoreflect.FileDescriptor

var file_some_lib_some_lib_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x62, 0x2f, 0x73, 0x6f, 0x6d, 0x65, 0x5f,
	0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x53, 0x6f, 0x6d, 0x65, 0x2e,
	0x4c, 0x69, 0x62, 0x22, 0x36, 0x0a, 0x0a, 0x53, 0x6f, 0x6d, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f, 0x6d,
	0x65, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x63, 0x0a, 0x0b, 0x53,
	0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f,
	0x6d, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x6f, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x6f, 0x6d,
	0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x2e, 0x4c, 0x69, 0x62, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x0a, 0x73, 0x6f, 0x6d, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x42, 0xa2, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x2e, 0x4c, 0x69,
	0x62, 0x42, 0x0c, 0x53, 0x6f, 0x6d, 0x65, 0x4c, 0x69, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e,
	0x61, 0x72, 0x6b, 0x79, 0x2d, 0x70, 0x75, 0x70, 0x70, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x66, 0x69,
	0x78, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x6f,
	0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x62, 0xa2, 0x02, 0x03, 0x53, 0x4c, 0x58, 0xaa, 0x02, 0x08, 0x53,
	0x6f, 0x6d, 0x65, 0x2e, 0x4c, 0x69, 0x62, 0xca, 0x02, 0x08, 0x53, 0x6f, 0x6d, 0x65, 0x5c, 0x4c,
	0x69, 0x62, 0xe2, 0x02, 0x14, 0x53, 0x6f, 0x6d, 0x65, 0x5c, 0x4c, 0x69, 0x62, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x53, 0x6f, 0x6d, 0x65,
	0x3a, 0x3a, 0x4c, 0x69, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_some_lib_some_lib_proto_rawDescOnce sync.Once
	file_some_lib_some_lib_proto_rawDescData = file_some_lib_some_lib_proto_rawDesc
)

func file_some_lib_some_lib_proto_rawDescGZIP() []byte {
	file_some_lib_some_lib_proto_rawDescOnce.Do(func() {
		file_some_lib_some_lib_proto_rawDescData = protoimpl.X.CompressGZIP(file_some_lib_some_lib_proto_rawDescData)
	})
	return file_some_lib_some_lib_proto_rawDescData
}

var file_some_lib_some_lib_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_some_lib_some_lib_proto_goTypes = []interface{}{
	(*SomeObject)(nil),  // 0: Some.Lib.SomeObject
	(*SomeMessage)(nil), // 1: Some.Lib.SomeMessage
}
var file_some_lib_some_lib_proto_depIdxs = []int32{
	0, // 0: Some.Lib.SomeMessage.some_object:type_name -> Some.Lib.SomeObject
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_some_lib_some_lib_proto_init() }
func file_some_lib_some_lib_proto_init() {
	if File_some_lib_some_lib_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_some_lib_some_lib_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeObject); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_some_lib_some_lib_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_some_lib_some_lib_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_some_lib_some_lib_proto_goTypes,
		DependencyIndexes: file_some_lib_some_lib_proto_depIdxs,
		MessageInfos:      file_some_lib_some_lib_proto_msgTypes,
	}.Build()
	File_some_lib_some_lib_proto = out.File
	file_some_lib_some_lib_proto_rawDesc = nil
	file_some_lib_some_lib_proto_goTypes = nil
	file_some_lib_some_lib_proto_depIdxs = nil
}
