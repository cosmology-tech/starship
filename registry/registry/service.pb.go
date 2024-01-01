// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: registry/service.proto

package registry

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
		mi := &file_registry_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseChains) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseChains) ProtoMessage() {}

func (x *ResponseChains) ProtoReflect() protoreflect.Message {
	mi := &file_registry_service_proto_msgTypes[0]
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
	return file_registry_service_proto_rawDescGZIP(), []int{0}
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
		mi := &file_registry_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseChainIDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseChainIDs) ProtoMessage() {}

func (x *ResponseChainIDs) ProtoReflect() protoreflect.Message {
	mi := &file_registry_service_proto_msgTypes[1]
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
	return file_registry_service_proto_rawDescGZIP(), []int{1}
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
		mi := &file_registry_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestChain) ProtoMessage() {}

func (x *RequestChain) ProtoReflect() protoreflect.Message {
	mi := &file_registry_service_proto_msgTypes[2]
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
	return file_registry_service_proto_rawDescGZIP(), []int{2}
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
		mi := &file_registry_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseChainAssets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseChainAssets) ProtoMessage() {}

func (x *ResponseChainAssets) ProtoReflect() protoreflect.Message {
	mi := &file_registry_service_proto_msgTypes[3]
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
	return file_registry_service_proto_rawDescGZIP(), []int{3}
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
		mi := &file_registry_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseListIBC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseListIBC) ProtoMessage() {}

func (x *ResponseListIBC) ProtoReflect() protoreflect.Message {
	mi := &file_registry_service_proto_msgTypes[4]
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
	return file_registry_service_proto_rawDescGZIP(), []int{4}
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
		mi := &file_registry_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestIBCInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestIBCInfo) ProtoMessage() {}

func (x *RequestIBCInfo) ProtoReflect() protoreflect.Message {
	mi := &file_registry_service_proto_msgTypes[5]
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
	return file_registry_service_proto_rawDescGZIP(), []int{5}
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

var File_registry_service_proto protoreflect.FileDescriptor

var file_registry_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x69, 0x62, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x30, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x22, 0x24, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x7c, 0x0a,
	0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x24, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x22, 0x38, 0x0a, 0x0f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x42, 0x43, 0x12, 0x25,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x49, 0x42, 0x43, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x42, 0x43, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x31, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x32, 0x32, 0xfb, 0x06, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x56, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x73, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73,
	0x12, 0x4f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x73,
	0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x73, 0x12, 0x54, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f,
	0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x12, 0x54, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a,
	0x0e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73,
	0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x58, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12,
	0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x12, 0x15, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x7d, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x55, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x41, 0x50, 0x49, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x41, 0x50, 0x49, 0x73,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x12, 0x67,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x12, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12,
	0x16, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x42, 0x43, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x42, 0x43, 0x22, 0x0c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x06, 0x12, 0x04, 0x2f,
	0x69, 0x62, 0x63, 0x12, 0x57, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x49, 0x42, 0x43, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x19, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x42, 0x43, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c,
	0x2f, 0x69, 0x62, 0x63, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x12, 0x5b, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x49, 0x42, 0x43, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x42, 0x43,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x11, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x49, 0x42, 0x43, 0x44, 0x61, 0x74, 0x61, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12,
	0x18, 0x2f, 0x69, 0x62, 0x63, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x31, 0x7d, 0x2f,
	0x7b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x32, 0x7d, 0x42, 0x89, 0x01, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0xa2, 0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0xca, 0x02, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0xe2, 0x02, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_registry_service_proto_rawDescOnce sync.Once
	file_registry_service_proto_rawDescData = file_registry_service_proto_rawDesc
)

func file_registry_service_proto_rawDescGZIP() []byte {
	file_registry_service_proto_rawDescOnce.Do(func() {
		file_registry_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_service_proto_rawDescData)
	})
	return file_registry_service_proto_rawDescData
}

var file_registry_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_registry_service_proto_goTypes = []interface{}{
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
var file_registry_service_proto_depIdxs = []int32{
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

func init() { file_registry_service_proto_init() }
func file_registry_service_proto_init() {
	if File_registry_service_proto != nil {
		return
	}
	file_registry_chain_proto_init()
	file_registry_ibc_proto_init()
	file_registry_mnemonic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_registry_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_registry_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_registry_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_registry_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_registry_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_registry_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_registry_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_registry_service_proto_goTypes,
		DependencyIndexes: file_registry_service_proto_depIdxs,
		MessageInfos:      file_registry_service_proto_msgTypes,
	}.Build()
	File_registry_service_proto = out.File
	file_registry_service_proto_rawDesc = nil
	file_registry_service_proto_goTypes = nil
	file_registry_service_proto_depIdxs = nil
}
