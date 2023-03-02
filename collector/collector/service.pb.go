// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: collector/service.proto

package collector

import (
	_ "github.com/bufbuild/buf-tour/gen/google/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponseNodeID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *ResponseNodeID) Reset() {
	*x = ResponseNodeID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseNodeID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseNodeID) ProtoMessage() {}

func (x *ResponseNodeID) ProtoReflect() protoreflect.Message {
	mi := &file_collector_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseNodeID.ProtoReflect.Descriptor instead.
func (*ResponseNodeID) Descriptor() ([]byte, []int) {
	return file_collector_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseNodeID) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type ResponsePubKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ResponsePubKey) Reset() {
	*x = ResponsePubKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePubKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePubKey) ProtoMessage() {}

func (x *ResponsePubKey) ProtoReflect() protoreflect.Message {
	mi := &file_collector_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePubKey.ProtoReflect.Descriptor instead.
func (*ResponsePubKey) Descriptor() ([]byte, []int) {
	return file_collector_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResponsePubKey) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ResponsePubKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ResponseFileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseFileData) Reset() {
	*x = ResponseFileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseFileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFileData) ProtoMessage() {}

func (x *ResponseFileData) ProtoReflect() protoreflect.Message {
	mi := &file_collector_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFileData.ProtoReflect.Descriptor instead.
func (*ResponseFileData) Descriptor() ([]byte, []int) {
	return file_collector_service_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseFileData) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_collector_service_proto protoreflect.FileDescriptor

var file_collector_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x3c, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xe1, 0x06,
	0x0a, 0x09, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x50, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x7a, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x7d, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x7d, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x7d, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x7d, 0x2f, 0x65, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7d, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x35, 0x22, 0x33, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x7d, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x7d, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7e, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x38, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x7d, 0x2f, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x37, 0x12, 0x35, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x7d, 0x2f, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x82, 0x01, 0x0a,
	0x11, 0x53, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x3d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x22, 0x35, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x7d, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x42, 0x8d, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x74, 0x6f, 0x75,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0xa2,
	0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0xca, 0x02, 0x09, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0xe2, 0x02, 0x15,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collector_service_proto_rawDescOnce sync.Once
	file_collector_service_proto_rawDescData = file_collector_service_proto_rawDesc
)

func file_collector_service_proto_rawDescGZIP() []byte {
	file_collector_service_proto_rawDescOnce.Do(func() {
		file_collector_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_collector_service_proto_rawDescData)
	})
	return file_collector_service_proto_rawDescData
}

var file_collector_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_collector_service_proto_goTypes = []interface{}{
	(*ResponseNodeID)(nil),   // 0: collector.ResponseNodeID
	(*ResponsePubKey)(nil),   // 1: collector.ResponsePubKey
	(*ResponseFileData)(nil), // 2: collector.ResponseFileData
	(*anypb.Any)(nil),        // 3: google.protobuf.Any
	(*emptypb.Empty)(nil),    // 4: google.protobuf.Empty
}
var file_collector_service_proto_depIdxs = []int32{
	3, // 0: collector.ResponseFileData.data:type_name -> google.protobuf.Any
	4, // 1: collector.Collector.GetChains:input_type -> google.protobuf.Empty
	4, // 2: collector.Collector.ListChainExports:input_type -> google.protobuf.Empty
	4, // 3: collector.Collector.GetChainExport:input_type -> google.protobuf.Empty
	4, // 4: collector.Collector.SetChainExport:input_type -> google.protobuf.Empty
	4, // 5: collector.Collector.ListChainSnapshots:input_type -> google.protobuf.Empty
	4, // 6: collector.Collector.GetChainSnapshots:input_type -> google.protobuf.Empty
	4, // 7: collector.Collector.SetChainSnapshots:input_type -> google.protobuf.Empty
	0, // 8: collector.Collector.GetChains:output_type -> collector.ResponseNodeID
	4, // 9: collector.Collector.ListChainExports:output_type -> google.protobuf.Empty
	4, // 10: collector.Collector.GetChainExport:output_type -> google.protobuf.Empty
	4, // 11: collector.Collector.SetChainExport:output_type -> google.protobuf.Empty
	4, // 12: collector.Collector.ListChainSnapshots:output_type -> google.protobuf.Empty
	4, // 13: collector.Collector.GetChainSnapshots:output_type -> google.protobuf.Empty
	4, // 14: collector.Collector.SetChainSnapshots:output_type -> google.protobuf.Empty
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_collector_service_proto_init() }
func file_collector_service_proto_init() {
	if File_collector_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collector_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseNodeID); i {
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
		file_collector_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePubKey); i {
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
		file_collector_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseFileData); i {
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
			RawDescriptor: file_collector_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collector_service_proto_goTypes,
		DependencyIndexes: file_collector_service_proto_depIdxs,
		MessageInfos:      file_collector_service_proto_msgTypes,
	}.Build()
	File_collector_service_proto = out.File
	file_collector_service_proto_rawDesc = nil
	file_collector_service_proto_goTypes = nil
	file_collector_service_proto_depIdxs = nil
}