// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: no-service.proto

package example

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_no_service_proto protoreflect.FileDescriptor

var file_no_service_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6e, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x9a, 0x01, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x0e, 0x4e,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x61, 0x72,
	0x6b, 0x79, 0x2d, 0x70, 0x75, 0x70, 0x70, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x66, 0x69, 0x78, 0x74,
	0x75, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0xa2, 0x02, 0x03, 0x50, 0x58,
	0x58, 0xaa, 0x02, 0x08, 0x50, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xca, 0x02, 0x08, 0x50,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xe2, 0x02, 0x14, 0x50, 0x65, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x08, 0x50, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_no_service_proto_goTypes = []interface{}{}
var file_no_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_no_service_proto_init() }
func file_no_service_proto_init() {
	if File_no_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_no_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_no_service_proto_goTypes,
		DependencyIndexes: file_no_service_proto_depIdxs,
	}.Build()
	File_no_service_proto = out.File
	file_no_service_proto_rawDesc = nil
	file_no_service_proto_goTypes = nil
	file_no_service_proto_depIdxs = nil
}
