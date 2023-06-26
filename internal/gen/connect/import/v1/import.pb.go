// Copyright 2021-2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: connect/import/v1/import.proto

package importv1

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

var File_connect_import_v1_import_proto protoreflect.FileDescriptor

var file_connect_import_v1_import_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x76, 0x31, 0x32, 0x0f, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x42, 0xd2, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x49, 0x58, 0xaa, 0x02, 0x11, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_connect_import_v1_import_proto_goTypes = []interface{}{}
var file_connect_import_v1_import_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_connect_import_v1_import_proto_init() }
func file_connect_import_v1_import_proto_init() {
	if File_connect_import_v1_import_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_connect_import_v1_import_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connect_import_v1_import_proto_goTypes,
		DependencyIndexes: file_connect_import_v1_import_proto_depIdxs,
	}.Build()
	File_connect_import_v1_import_proto = out.File
	file_connect_import_v1_import_proto_rawDesc = nil
	file_connect_import_v1_import_proto_goTypes = nil
	file_connect_import_v1_import_proto_depIdxs = nil
}
