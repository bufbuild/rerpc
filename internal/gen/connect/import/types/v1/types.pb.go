// Copyright 2021-2022 Buf Technologies, Inc.
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
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: connect/import/types/v1/types.proto

package typesv1

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

type ImportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ImportRequest) Reset() {
	*x = ImportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_import_types_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportRequest) ProtoMessage() {}

func (x *ImportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connect_import_types_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportRequest.ProtoReflect.Descriptor instead.
func (*ImportRequest) Descriptor() ([]byte, []int) {
	return file_connect_import_types_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *ImportRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ImportRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ImportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ImportResponse) Reset() {
	*x = ImportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_import_types_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResponse) ProtoMessage() {}

func (x *ImportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connect_import_types_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportResponse.ProtoReflect.Descriptor instead.
func (*ImportResponse) Descriptor() ([]byte, []int) {
	return file_connect_import_types_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *ImportResponse) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ImportResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_connect_import_types_v1_types_proto protoreflect.FileDescriptor

var file_connect_import_types_v1_types_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x3b,
	0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3c, 0x0a, 0x0e, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0xf5, 0x01, 0x0a, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x49, 0x54, 0xaa, 0x02, 0x17, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x23, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5c,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x3a, 0x3a,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connect_import_types_v1_types_proto_rawDescOnce sync.Once
	file_connect_import_types_v1_types_proto_rawDescData = file_connect_import_types_v1_types_proto_rawDesc
)

func file_connect_import_types_v1_types_proto_rawDescGZIP() []byte {
	file_connect_import_types_v1_types_proto_rawDescOnce.Do(func() {
		file_connect_import_types_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_connect_import_types_v1_types_proto_rawDescData)
	})
	return file_connect_import_types_v1_types_proto_rawDescData
}

var file_connect_import_types_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_connect_import_types_v1_types_proto_goTypes = []interface{}{
	(*ImportRequest)(nil),  // 0: connect.import.types.v1.ImportRequest
	(*ImportResponse)(nil), // 1: connect.import.types.v1.ImportResponse
}
var file_connect_import_types_v1_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_connect_import_types_v1_types_proto_init() }
func file_connect_import_types_v1_types_proto_init() {
	if File_connect_import_types_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connect_import_types_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportRequest); i {
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
		file_connect_import_types_v1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportResponse); i {
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
			RawDescriptor: file_connect_import_types_v1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connect_import_types_v1_types_proto_goTypes,
		DependencyIndexes: file_connect_import_types_v1_types_proto_depIdxs,
		MessageInfos:      file_connect_import_types_v1_types_proto_msgTypes,
	}.Build()
	File_connect_import_types_v1_types_proto = out.File
	file_connect_import_types_v1_types_proto_rawDesc = nil
	file_connect_import_types_v1_types_proto_goTypes = nil
	file_connect_import_types_v1_types_proto_depIdxs = nil
}
