// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.15.5
// source: proto/registry/service.proto

package registry

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
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

type ResponseChains struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chains []*ChainRegistry `protobuf:"bytes,1,rep,name=chains,proto3" json:"chains,omitempty"`
}

func (x *ResponseChains) Reset() {
	*x = ResponseChains{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_registry_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseChains) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseChains) ProtoMessage() {}

func (x *ResponseChains) ProtoReflect() protoreflect.Message {
	mi := &file_proto_registry_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseChains.ProtoReflect.Descriptor instead.
func (*ResponseChains) Descriptor() ([]byte, []int) {
	return file_proto_registry_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseChains) GetChains() []*ChainRegistry {
	if x != nil {
		return x.Chains
	}
	return nil
}

type ResponseChainIDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainIds []string `protobuf:"bytes,1,rep,name=chain_ids,proto3" json:"chain_ids,omitempty"`
}

func (x *ResponseChainIDs) Reset() {
	*x = ResponseChainIDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_registry_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseChainIDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseChainIDs) ProtoMessage() {}

func (x *ResponseChainIDs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_registry_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseChainIDs.ProtoReflect.Descriptor instead.
func (*ResponseChainIDs) Descriptor() ([]byte, []int) {
	return file_proto_registry_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseChainIDs) GetChainIds() []string {
	if x != nil {
		return x.ChainIds
	}
	return nil
}

type RequestChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (x *RequestChain) Reset() {
	*x = RequestChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_registry_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestChain) ProtoMessage() {}

func (x *RequestChain) ProtoReflect() protoreflect.Message {
	mi := &file_proto_registry_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestChain.ProtoReflect.Descriptor instead.
func (*RequestChain) Descriptor() ([]byte, []int) {
	return file_proto_registry_service_proto_rawDescGZIP(), []int{2}
}

func (x *RequestChain) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

type ResponseChainAssets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema    string        `protobuf:"bytes,1,opt,name=schema,json=$schema,proto3" json:"schema,omitempty"`
	ChainName string        `protobuf:"bytes,2,opt,name=chain_name,proto3" json:"chain_name,omitempty"`
	Assets    []*ChainAsset `protobuf:"bytes,3,rep,name=assets,proto3" json:"assets,omitempty"`
}

func (x *ResponseChainAssets) Reset() {
	*x = ResponseChainAssets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_registry_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseChainAssets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseChainAssets) ProtoMessage() {}

func (x *ResponseChainAssets) ProtoReflect() protoreflect.Message {
	mi := &file_proto_registry_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseChainAssets.ProtoReflect.Descriptor instead.
func (*ResponseChainAssets) Descriptor() ([]byte, []int) {
	return file_proto_registry_service_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseChainAssets) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *ResponseChainAssets) GetChainName() string {
	if x != nil {
		return x.ChainName
	}
	return ""
}

func (x *ResponseChainAssets) GetAssets() []*ChainAsset {
	if x != nil {
		return x.Assets
	}
	return nil
}

type ResponseListIBC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*IBCData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseListIBC) Reset() {
	*x = ResponseListIBC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_registry_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseListIBC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseListIBC) ProtoMessage() {}

func (x *ResponseListIBC) ProtoReflect() protoreflect.Message {
	mi := &file_proto_registry_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseListIBC.ProtoReflect.Descriptor instead.
func (*ResponseListIBC) Descriptor() ([]byte, []int) {
	return file_proto_registry_service_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseListIBC) GetData() []*IBCData {
	if x != nil {
		return x.Data
	}
	return nil
}

type RequestIBCInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain_1 string `protobuf:"bytes,1,opt,name=chain_1,proto3" json:"chain_1,omitempty"`
	Chain_2 string `protobuf:"bytes,2,opt,name=chain_2,proto3" json:"chain_2,omitempty"`
}

func (x *RequestIBCInfo) Reset() {
	*x = RequestIBCInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_registry_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestIBCInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestIBCInfo) ProtoMessage() {}

func (x *RequestIBCInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_registry_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestIBCInfo.ProtoReflect.Descriptor instead.
func (*RequestIBCInfo) Descriptor() ([]byte, []int) {
	return file_proto_registry_service_proto_rawDescGZIP(), []int{5}
}

func (x *RequestIBCInfo) GetChain_1() string {
	if x != nil {
		return x.Chain_1
	}
	return ""
}

func (x *RequestIBCInfo) GetChain_2() string {
	if x != nil {
		return x.Chain_2
	}
	return ""
}

var File_proto_registry_service_proto protoreflect.FileDescriptor

var file_proto_registry_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x69, 0x62, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x6d, 0x6e, 0x65, 0x6d,
	0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x06,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x30, 0x0a,
	0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x22,
	0x24, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x7c, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x24, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x06, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x22, 0x38, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x42, 0x43, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x49, 0x42, 0x43, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x44, 0x0a,
	0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x42, 0x43, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x32, 0x32, 0xfb, 0x06, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x12, 0x56, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x49, 0x44, 0x73, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x12, 0x4f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09,
	0x12, 0x07, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x54, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x17, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f,
	0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x12,
	0x54, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x12,
	0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12,
	0x14, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d,
	0x2f, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x58, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a,
	0x0f, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12,
	0x55, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x41, 0x50, 0x49, 0x73,
	0x12, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x41, 0x50, 0x49, 0x73, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x12, 0x14, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x7d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x12, 0x67, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73,
	0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12,
	0x4a, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x42, 0x43, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x42, 0x43, 0x22, 0x0c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x06, 0x12, 0x04, 0x2f, 0x69, 0x62, 0x63, 0x12, 0x57, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x42, 0x43, 0x12, 0x16, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x42, 0x43, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x69, 0x62, 0x63, 0x2f, 0x7b, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x7d, 0x12, 0x5b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49, 0x42, 0x43, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x42, 0x43, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x11, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x49, 0x42, 0x43, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x69, 0x62, 0x63, 0x2f, 0x7b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x31, 0x7d, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x32,
	0x7d, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73,
	0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_registry_service_proto_rawDescOnce sync.Once
	file_proto_registry_service_proto_rawDescData = file_proto_registry_service_proto_rawDesc
)

func file_proto_registry_service_proto_rawDescGZIP() []byte {
	file_proto_registry_service_proto_rawDescOnce.Do(func() {
		file_proto_registry_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_registry_service_proto_rawDescData)
	})
	return file_proto_registry_service_proto_rawDescData
}

var file_proto_registry_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_registry_service_proto_goTypes = []interface{}{
	(*ResponseChains)(nil),      // 0: registry.ResponseChains
	(*ResponseChainIDs)(nil),    // 1: registry.ResponseChainIDs
	(*RequestChain)(nil),        // 2: registry.RequestChain
	(*ResponseChainAssets)(nil), // 3: registry.ResponseChainAssets
	(*ResponseListIBC)(nil),     // 4: registry.ResponseListIBC
	(*RequestIBCInfo)(nil),      // 5: registry.RequestIBCInfo
	(*ChainRegistry)(nil),       // 6: registry.ChainRegistry
	(*ChainAsset)(nil),          // 7: registry.ChainAsset
	(*IBCData)(nil),             // 8: registry.IBCData
	(*emptypb.Empty)(nil),       // 9: google.protobuf.Empty
	(*Keys)(nil),                // 10: registry.Keys
	(*Peers)(nil),               // 11: registry.Peers
	(*APIs)(nil),                // 12: registry.APIs
}
var file_proto_registry_service_proto_depIdxs = []int32{
	6,  // 0: registry.ResponseChains.chains:type_name -> registry.ChainRegistry
	7,  // 1: registry.ResponseChainAssets.assets:type_name -> registry.ChainAsset
	8,  // 2: registry.ResponseListIBC.data:type_name -> registry.IBCData
	9,  // 3: registry.Registry.ListChainIDs:input_type -> google.protobuf.Empty
	9,  // 4: registry.Registry.ListChains:input_type -> google.protobuf.Empty
	2,  // 5: registry.Registry.GetChain:input_type -> registry.RequestChain
	2,  // 6: registry.Registry.GetChainKeys:input_type -> registry.RequestChain
	2,  // 7: registry.Registry.ListChainPeers:input_type -> registry.RequestChain
	2,  // 8: registry.Registry.ListChainAPIs:input_type -> registry.RequestChain
	2,  // 9: registry.Registry.GetChainAssets:input_type -> registry.RequestChain
	9,  // 10: registry.Registry.ListIBC:input_type -> google.protobuf.Empty
	2,  // 11: registry.Registry.ListChainIBC:input_type -> registry.RequestChain
	5,  // 12: registry.Registry.GetIBCInfo:input_type -> registry.RequestIBCInfo
	1,  // 13: registry.Registry.ListChainIDs:output_type -> registry.ResponseChainIDs
	0,  // 14: registry.Registry.ListChains:output_type -> registry.ResponseChains
	6,  // 15: registry.Registry.GetChain:output_type -> registry.ChainRegistry
	10, // 16: registry.Registry.GetChainKeys:output_type -> registry.Keys
	11, // 17: registry.Registry.ListChainPeers:output_type -> registry.Peers
	12, // 18: registry.Registry.ListChainAPIs:output_type -> registry.APIs
	3,  // 19: registry.Registry.GetChainAssets:output_type -> registry.ResponseChainAssets
	4,  // 20: registry.Registry.ListIBC:output_type -> registry.ResponseListIBC
	4,  // 21: registry.Registry.ListChainIBC:output_type -> registry.ResponseListIBC
	8,  // 22: registry.Registry.GetIBCInfo:output_type -> registry.IBCData
	13, // [13:23] is the sub-list for method output_type
	3,  // [3:13] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_registry_service_proto_init() }
func file_proto_registry_service_proto_init() {
	if File_proto_registry_service_proto != nil {
		return
	}
	file_proto_registry_chain_proto_init()
	file_proto_registry_ibc_proto_init()
	file_proto_registry_mnemonic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_registry_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseChains); i {
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
		file_proto_registry_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseChainIDs); i {
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
		file_proto_registry_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestChain); i {
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
		file_proto_registry_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseChainAssets); i {
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
		file_proto_registry_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseListIBC); i {
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
		file_proto_registry_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestIBCInfo); i {
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
			RawDescriptor: file_proto_registry_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_registry_service_proto_goTypes,
		DependencyIndexes: file_proto_registry_service_proto_depIdxs,
		MessageInfos:      file_proto_registry_service_proto_msgTypes,
	}.Build()
	File_proto_registry_service_proto = out.File
	file_proto_registry_service_proto_rawDesc = nil
	file_proto_registry_service_proto_goTypes = nil
	file_proto_registry_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryClient interface {
	ListChainIDs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseChainIDs, error)
	ListChains(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseChains, error)
	GetChain(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*ChainRegistry, error)
	GetChainKeys(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*Keys, error)
	ListChainPeers(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*Peers, error)
	ListChainAPIs(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*APIs, error)
	GetChainAssets(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*ResponseChainAssets, error)
	ListIBC(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseListIBC, error)
	ListChainIBC(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*ResponseListIBC, error)
	GetIBCInfo(ctx context.Context, in *RequestIBCInfo, opts ...grpc.CallOption) (*IBCData, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) ListChainIDs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseChainIDs, error) {
	out := new(ResponseChainIDs)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListChainIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListChains(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseChains, error) {
	out := new(ResponseChains)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListChains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetChain(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*ChainRegistry, error) {
	out := new(ChainRegistry)
	err := c.cc.Invoke(ctx, "/registry.Registry/GetChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetChainKeys(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*Keys, error) {
	out := new(Keys)
	err := c.cc.Invoke(ctx, "/registry.Registry/GetChainKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListChainPeers(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*Peers, error) {
	out := new(Peers)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListChainPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListChainAPIs(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*APIs, error) {
	out := new(APIs)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListChainAPIs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetChainAssets(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*ResponseChainAssets, error) {
	out := new(ResponseChainAssets)
	err := c.cc.Invoke(ctx, "/registry.Registry/GetChainAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListIBC(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseListIBC, error) {
	out := new(ResponseListIBC)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListIBC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListChainIBC(ctx context.Context, in *RequestChain, opts ...grpc.CallOption) (*ResponseListIBC, error) {
	out := new(ResponseListIBC)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListChainIBC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetIBCInfo(ctx context.Context, in *RequestIBCInfo, opts ...grpc.CallOption) (*IBCData, error) {
	out := new(IBCData)
	err := c.cc.Invoke(ctx, "/registry.Registry/GetIBCInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
type RegistryServer interface {
	ListChainIDs(context.Context, *emptypb.Empty) (*ResponseChainIDs, error)
	ListChains(context.Context, *emptypb.Empty) (*ResponseChains, error)
	GetChain(context.Context, *RequestChain) (*ChainRegistry, error)
	GetChainKeys(context.Context, *RequestChain) (*Keys, error)
	ListChainPeers(context.Context, *RequestChain) (*Peers, error)
	ListChainAPIs(context.Context, *RequestChain) (*APIs, error)
	GetChainAssets(context.Context, *RequestChain) (*ResponseChainAssets, error)
	ListIBC(context.Context, *emptypb.Empty) (*ResponseListIBC, error)
	ListChainIBC(context.Context, *RequestChain) (*ResponseListIBC, error)
	GetIBCInfo(context.Context, *RequestIBCInfo) (*IBCData, error)
}

// UnimplementedRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (*UnimplementedRegistryServer) ListChainIDs(context.Context, *emptypb.Empty) (*ResponseChainIDs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChainIDs not implemented")
}
func (*UnimplementedRegistryServer) ListChains(context.Context, *emptypb.Empty) (*ResponseChains, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChains not implemented")
}
func (*UnimplementedRegistryServer) GetChain(context.Context, *RequestChain) (*ChainRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChain not implemented")
}
func (*UnimplementedRegistryServer) GetChainKeys(context.Context, *RequestChain) (*Keys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChainKeys not implemented")
}
func (*UnimplementedRegistryServer) ListChainPeers(context.Context, *RequestChain) (*Peers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChainPeers not implemented")
}
func (*UnimplementedRegistryServer) ListChainAPIs(context.Context, *RequestChain) (*APIs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChainAPIs not implemented")
}
func (*UnimplementedRegistryServer) GetChainAssets(context.Context, *RequestChain) (*ResponseChainAssets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChainAssets not implemented")
}
func (*UnimplementedRegistryServer) ListIBC(context.Context, *emptypb.Empty) (*ResponseListIBC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIBC not implemented")
}
func (*UnimplementedRegistryServer) ListChainIBC(context.Context, *RequestChain) (*ResponseListIBC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChainIBC not implemented")
}
func (*UnimplementedRegistryServer) GetIBCInfo(context.Context, *RequestIBCInfo) (*IBCData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIBCInfo not implemented")
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_ListChainIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListChainIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListChainIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListChainIDs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListChains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListChains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListChains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListChains(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/GetChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetChain(ctx, req.(*RequestChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetChainKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetChainKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/GetChainKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetChainKeys(ctx, req.(*RequestChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListChainPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListChainPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListChainPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListChainPeers(ctx, req.(*RequestChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListChainAPIs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListChainAPIs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListChainAPIs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListChainAPIs(ctx, req.(*RequestChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetChainAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetChainAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/GetChainAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetChainAssets(ctx, req.(*RequestChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListIBC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListIBC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListIBC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListIBC(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListChainIBC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListChainIBC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListChainIBC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListChainIBC(ctx, req.(*RequestChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetIBCInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestIBCInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetIBCInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/GetIBCInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetIBCInfo(ctx, req.(*RequestIBCInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListChainIDs",
			Handler:    _Registry_ListChainIDs_Handler,
		},
		{
			MethodName: "ListChains",
			Handler:    _Registry_ListChains_Handler,
		},
		{
			MethodName: "GetChain",
			Handler:    _Registry_GetChain_Handler,
		},
		{
			MethodName: "GetChainKeys",
			Handler:    _Registry_GetChainKeys_Handler,
		},
		{
			MethodName: "ListChainPeers",
			Handler:    _Registry_ListChainPeers_Handler,
		},
		{
			MethodName: "ListChainAPIs",
			Handler:    _Registry_ListChainAPIs_Handler,
		},
		{
			MethodName: "GetChainAssets",
			Handler:    _Registry_GetChainAssets_Handler,
		},
		{
			MethodName: "ListIBC",
			Handler:    _Registry_ListIBC_Handler,
		},
		{
			MethodName: "ListChainIBC",
			Handler:    _Registry_ListChainIBC_Handler,
		},
		{
			MethodName: "GetIBCInfo",
			Handler:    _Registry_GetIBCInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/registry/service.proto",
}
