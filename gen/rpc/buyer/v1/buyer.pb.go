// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: rpc/buyer/v1/buyer.proto

package buyerv1

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

type CreateBuyerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Phone   string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *CreateBuyerRequest) Reset() {
	*x = CreateBuyerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBuyerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBuyerRequest) ProtoMessage() {}

func (x *CreateBuyerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBuyerRequest.ProtoReflect.Descriptor instead.
func (*CreateBuyerRequest) Descriptor() ([]byte, []int) {
	return file_rpc_buyer_v1_buyer_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBuyerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBuyerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateBuyerRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateBuyerRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type CreateBuyerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateBuyerResponse) Reset() {
	*x = CreateBuyerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBuyerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBuyerResponse) ProtoMessage() {}

func (x *CreateBuyerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBuyerResponse.ProtoReflect.Descriptor instead.
func (*CreateBuyerResponse) Descriptor() ([]byte, []int) {
	return file_rpc_buyer_v1_buyer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBuyerResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type UpdateBuyerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *UpdateBuyerRequest) Reset() {
	*x = UpdateBuyerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBuyerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBuyerRequest) ProtoMessage() {}

func (x *UpdateBuyerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBuyerRequest.ProtoReflect.Descriptor instead.
func (*UpdateBuyerRequest) Descriptor() ([]byte, []int) {
	return file_rpc_buyer_v1_buyer_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBuyerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBuyerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateBuyerRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type UpdateBuyerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateBuyerResponse) Reset() {
	*x = UpdateBuyerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBuyerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBuyerResponse) ProtoMessage() {}

func (x *UpdateBuyerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBuyerResponse.ProtoReflect.Descriptor instead.
func (*UpdateBuyerResponse) Descriptor() ([]byte, []int) {
	return file_rpc_buyer_v1_buyer_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBuyerResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type GetBuyerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBuyerRequest) Reset() {
	*x = GetBuyerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBuyerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuyerRequest) ProtoMessage() {}

func (x *GetBuyerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuyerRequest.ProtoReflect.Descriptor instead.
func (*GetBuyerRequest) Descriptor() ([]byte, []int) {
	return file_rpc_buyer_v1_buyer_proto_rawDescGZIP(), []int{4}
}

type GetBuyerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetBuyerResponse) Reset() {
	*x = GetBuyerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBuyerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuyerResponse) ProtoMessage() {}

func (x *GetBuyerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuyerResponse.ProtoReflect.Descriptor instead.
func (*GetBuyerResponse) Descriptor() ([]byte, []int) {
	return file_rpc_buyer_v1_buyer_proto_rawDescGZIP(), []int{5}
}

func (x *GetBuyerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBuyerResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetBuyerResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetBuyerResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DeactivateBuyerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeactivateBuyerRequest) Reset() {
	*x = DeactivateBuyerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateBuyerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateBuyerRequest) ProtoMessage() {}

func (x *DeactivateBuyerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateBuyerRequest.ProtoReflect.Descriptor instead.
func (*DeactivateBuyerRequest) Descriptor() ([]byte, []int) {
	return file_rpc_buyer_v1_buyer_proto_rawDescGZIP(), []int{6}
}

func (x *DeactivateBuyerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeactivateBuyerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeactivateBuyerResponse) Reset() {
	*x = DeactivateBuyerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateBuyerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateBuyerResponse) ProtoMessage() {}

func (x *DeactivateBuyerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_buyer_v1_buyer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateBuyerResponse.ProtoReflect.Descriptor instead.
func (*DeactivateBuyerResponse) Descriptor() ([]byte, []int) {
	return file_rpc_buyer_v1_buyer_proto_rawDescGZIP(), []int{7}
}

func (x *DeactivateBuyerResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_rpc_buyer_v1_buyer_proto protoreflect.FileDescriptor

var file_rpc_buyer_v1_buyer_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x75, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x75, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x75, 0x79, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x22, 0x6e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x58, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2d, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x11, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x66, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x31, 0x0a, 0x17, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0xc9, 0x02, 0x0a, 0x0c, 0x42, 0x75, 0x79, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x79, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x62, 0x75, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x75, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x79,
	0x65, 0x72, 0x12, 0x1c, 0x2e, 0x62, 0x75, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x62, 0x75, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x75, 0x79, 0x65, 0x72, 0x12, 0x19, 0x2e,
	0x62, 0x75, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x75, 0x79, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0f, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x42, 0x75, 0x79, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x62, 0x75, 0x79, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x75,
	0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6d, 0x6b, 0x72, 0x69, 0x73, 0x68, 0x37, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x68, 0x69, 0x70, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x75, 0x79, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x75, 0x79, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_buyer_v1_buyer_proto_rawDescOnce sync.Once
	file_rpc_buyer_v1_buyer_proto_rawDescData = file_rpc_buyer_v1_buyer_proto_rawDesc
)

func file_rpc_buyer_v1_buyer_proto_rawDescGZIP() []byte {
	file_rpc_buyer_v1_buyer_proto_rawDescOnce.Do(func() {
		file_rpc_buyer_v1_buyer_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_buyer_v1_buyer_proto_rawDescData)
	})
	return file_rpc_buyer_v1_buyer_proto_rawDescData
}

var file_rpc_buyer_v1_buyer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_buyer_v1_buyer_proto_goTypes = []interface{}{
	(*CreateBuyerRequest)(nil),      // 0: buyer.v1.CreateBuyerRequest
	(*CreateBuyerResponse)(nil),     // 1: buyer.v1.CreateBuyerResponse
	(*UpdateBuyerRequest)(nil),      // 2: buyer.v1.UpdateBuyerRequest
	(*UpdateBuyerResponse)(nil),     // 3: buyer.v1.UpdateBuyerResponse
	(*GetBuyerRequest)(nil),         // 4: buyer.v1.GetBuyerRequest
	(*GetBuyerResponse)(nil),        // 5: buyer.v1.GetBuyerResponse
	(*DeactivateBuyerRequest)(nil),  // 6: buyer.v1.DeactivateBuyerRequest
	(*DeactivateBuyerResponse)(nil), // 7: buyer.v1.DeactivateBuyerResponse
}
var file_rpc_buyer_v1_buyer_proto_depIdxs = []int32{
	0, // 0: buyer.v1.BuyerService.CreateBuyer:input_type -> buyer.v1.CreateBuyerRequest
	2, // 1: buyer.v1.BuyerService.UpdateBuyer:input_type -> buyer.v1.UpdateBuyerRequest
	4, // 2: buyer.v1.BuyerService.GetBuyer:input_type -> buyer.v1.GetBuyerRequest
	6, // 3: buyer.v1.BuyerService.DeactivateBuyer:input_type -> buyer.v1.DeactivateBuyerRequest
	1, // 4: buyer.v1.BuyerService.CreateBuyer:output_type -> buyer.v1.CreateBuyerResponse
	3, // 5: buyer.v1.BuyerService.UpdateBuyer:output_type -> buyer.v1.UpdateBuyerResponse
	5, // 6: buyer.v1.BuyerService.GetBuyer:output_type -> buyer.v1.GetBuyerResponse
	7, // 7: buyer.v1.BuyerService.DeactivateBuyer:output_type -> buyer.v1.DeactivateBuyerResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_buyer_v1_buyer_proto_init() }
func file_rpc_buyer_v1_buyer_proto_init() {
	if File_rpc_buyer_v1_buyer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_buyer_v1_buyer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBuyerRequest); i {
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
		file_rpc_buyer_v1_buyer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBuyerResponse); i {
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
		file_rpc_buyer_v1_buyer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBuyerRequest); i {
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
		file_rpc_buyer_v1_buyer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBuyerResponse); i {
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
		file_rpc_buyer_v1_buyer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBuyerRequest); i {
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
		file_rpc_buyer_v1_buyer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBuyerResponse); i {
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
		file_rpc_buyer_v1_buyer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeactivateBuyerRequest); i {
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
		file_rpc_buyer_v1_buyer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeactivateBuyerResponse); i {
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
			RawDescriptor: file_rpc_buyer_v1_buyer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_buyer_v1_buyer_proto_goTypes,
		DependencyIndexes: file_rpc_buyer_v1_buyer_proto_depIdxs,
		MessageInfos:      file_rpc_buyer_v1_buyer_proto_msgTypes,
	}.Build()
	File_rpc_buyer_v1_buyer_proto = out.File
	file_rpc_buyer_v1_buyer_proto_rawDesc = nil
	file_rpc_buyer_v1_buyer_proto_goTypes = nil
	file_rpc_buyer_v1_buyer_proto_depIdxs = nil
}