// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: checker.proto

package checker

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

var file_checker_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1130,
		Name:          "checker.disable_file_validate",
		Tag:           "varint,1130,opt,name=disable_file_validate",
		Filename:      "checker.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1131,
		Name:          "checker.disable_field_validate",
		Tag:           "varint,1131,opt,name=disable_field_validate",
		Filename:      "checker.proto",
	},
}

// Extension fields to descriptor.FileOptions.
var (
	// optional bool disable_file_validate = 1130;
	E_DisableFileValidate = &file_checker_proto_extTypes[0]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional bool disable_field_validate = 1131;
	E_DisableFieldValidate = &file_checker_proto_extTypes[1]
)

var File_checker_proto protoreflect.FileDescriptor

var file_checker_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x54, 0x0a, 0x15, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xea, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01,
	0x3a, 0x57, 0x0a, 0x16, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xeb, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x65,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_checker_proto_goTypes = []interface{}{
	(*descriptor.FileOptions)(nil),  // 0: google.protobuf.FileOptions
	(*descriptor.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_checker_proto_depIdxs = []int32{
	0, // 0: checker.disable_file_validate:extendee -> google.protobuf.FileOptions
	1, // 1: checker.disable_field_validate:extendee -> google.protobuf.FieldOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_checker_proto_init() }
func file_checker_proto_init() {
	if File_checker_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_checker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_checker_proto_goTypes,
		DependencyIndexes: file_checker_proto_depIdxs,
		ExtensionInfos:    file_checker_proto_extTypes,
	}.Build()
	File_checker_proto = out.File
	file_checker_proto_rawDesc = nil
	file_checker_proto_goTypes = nil
	file_checker_proto_depIdxs = nil
}
