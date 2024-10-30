// Copyright 2022 The N42 Authors
// This file is part of the N42 library.
//
// The N42 library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The N42 library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the N42 library. If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: consensus.proto

package consensus_proto

import (
	types_pb "github.com/N42world/ast/api/protocol/types_pb"
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

type PBSigners struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer []*PBSigner `protobuf:"bytes,1,rep,name=signer,proto3" json:"signer,omitempty"`
}

func (x *PBSigners) Reset() {
	*x = PBSigners{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBSigners) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBSigners) ProtoMessage() {}

func (x *PBSigners) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBSigners.ProtoReflect.Descriptor instead.
func (*PBSigners) Descriptor() ([]byte, []int) {
	return file_consensus_proto_rawDescGZIP(), []int{0}
}

func (x *PBSigners) GetSigner() []*PBSigner {
	if x != nil {
		return x.Signer
	}
	return nil
}

type PBSigner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Public  string         `protobuf:"bytes,1,opt,name=public,proto3" json:"public,omitempty"`
	Address *types_pb.H160 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Version string         `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *PBSigner) Reset() {
	*x = PBSigner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBSigner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBSigner) ProtoMessage() {}

func (x *PBSigner) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBSigner.ProtoReflect.Descriptor instead.
func (*PBSigner) Descriptor() ([]byte, []int) {
	return file_consensus_proto_rawDescGZIP(), []int{1}
}

func (x *PBSigner) GetPublic() string {
	if x != nil {
		return x.Public
	}
	return ""
}

func (x *PBSigner) GetAddress() *types_pb.H160 {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *PBSigner) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type PBVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PBVote) Reset() {
	*x = PBVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBVote) ProtoMessage() {}

func (x *PBVote) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBVote.ProtoReflect.Descriptor instead.
func (*PBVote) Descriptor() ([]byte, []int) {
	return file_consensus_proto_rawDescGZIP(), []int{2}
}

type PBPoaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Public string `protobuf:"bytes,1,opt,name=public,proto3" json:"public,omitempty"`
	Sign   []byte `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
	Type   int64  `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PBPoaInfo) Reset() {
	*x = PBPoaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBPoaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBPoaInfo) ProtoMessage() {}

func (x *PBPoaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBPoaInfo.ProtoReflect.Descriptor instead.
func (*PBPoaInfo) Descriptor() ([]byte, []int) {
	return file_consensus_proto_rawDescGZIP(), []int{3}
}

func (x *PBPoaInfo) GetPublic() string {
	if x != nil {
		return x.Public
	}
	return ""
}

func (x *PBPoaInfo) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

func (x *PBPoaInfo) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_consensus_proto protoreflect.FileDescriptor

var file_consensus_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x70, 0x62, 0x1a,
	0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x09, 0x50, 0x42, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x72, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x70,
	0x62, 0x2e, 0x50, 0x42, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x22, 0x66, 0x0a, 0x08, 0x50, 0x42, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f,
	0x70, 0x62, 0x2e, 0x48, 0x31, 0x36, 0x30, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x08, 0x0a, 0x06, 0x50, 0x42,
	0x56, 0x6f, 0x74, 0x65, 0x22, 0x4b, 0x0a, 0x09, 0x50, 0x42, 0x50, 0x6f, 0x61, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6d, 0x61, 0x7a, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x6d, 0x63, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_consensus_proto_rawDescOnce sync.Once
	file_consensus_proto_rawDescData = file_consensus_proto_rawDesc
)

func file_consensus_proto_rawDescGZIP() []byte {
	file_consensus_proto_rawDescOnce.Do(func() {
		file_consensus_proto_rawDescData = protoimpl.X.CompressGZIP(file_consensus_proto_rawDescData)
	})
	return file_consensus_proto_rawDescData
}

var file_consensus_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_consensus_proto_goTypes = []interface{}{
	(*PBSigners)(nil),     // 0: consensus_pb.PBSigners
	(*PBSigner)(nil),      // 1: consensus_pb.PBSigner
	(*PBVote)(nil),        // 2: consensus_pb.PBVote
	(*PBPoaInfo)(nil),     // 3: consensus_pb.PBPoaInfo
	(*types_pb.H160)(nil), // 4: types_pb.H160
}
var file_consensus_proto_depIdxs = []int32{
	1, // 0: consensus_pb.PBSigners.signer:type_name -> consensus_pb.PBSigner
	4, // 1: consensus_pb.PBSigner.address:type_name -> types_pb.H160
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_consensus_proto_init() }
func file_consensus_proto_init() {
	if File_consensus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consensus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBSigners); i {
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
		file_consensus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBSigner); i {
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
		file_consensus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBVote); i {
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
		file_consensus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBPoaInfo); i {
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
			RawDescriptor: file_consensus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_consensus_proto_goTypes,
		DependencyIndexes: file_consensus_proto_depIdxs,
		MessageInfos:      file_consensus_proto_msgTypes,
	}.Build()
	File_consensus_proto = out.File
	file_consensus_proto_rawDesc = nil
	file_consensus_proto_goTypes = nil
	file_consensus_proto_depIdxs = nil
}
