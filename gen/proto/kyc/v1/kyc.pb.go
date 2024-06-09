// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: proto/kyc/v1/kyc.proto

package kyc1

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

type CompleteKycRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname  string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Age       int64  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	BirthDate string `protobuf:"bytes,3,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Aadhar    string `protobuf:"bytes,4,opt,name=aadhar,proto3" json:"aadhar,omitempty"`
	Pan       string `protobuf:"bytes,5,opt,name=pan,proto3" json:"pan,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *CompleteKycRequest) Reset() {
	*x = CompleteKycRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kyc_v1_kyc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteKycRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteKycRequest) ProtoMessage() {}

func (x *CompleteKycRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kyc_v1_kyc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteKycRequest.ProtoReflect.Descriptor instead.
func (*CompleteKycRequest) Descriptor() ([]byte, []int) {
	return file_proto_kyc_v1_kyc_proto_rawDescGZIP(), []int{0}
}

func (x *CompleteKycRequest) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *CompleteKycRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CompleteKycRequest) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *CompleteKycRequest) GetAadhar() string {
	if x != nil {
		return x.Aadhar
	}
	return ""
}

func (x *CompleteKycRequest) GetPan() string {
	if x != nil {
		return x.Pan
	}
	return ""
}

func (x *CompleteKycRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type CompleteKycResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CompleteKycResponse) Reset() {
	*x = CompleteKycResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kyc_v1_kyc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteKycResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteKycResponse) ProtoMessage() {}

func (x *CompleteKycResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kyc_v1_kyc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteKycResponse.ProtoReflect.Descriptor instead.
func (*CompleteKycResponse) Descriptor() ([]byte, []int) {
	return file_proto_kyc_v1_kyc_proto_rawDescGZIP(), []int{1}
}

func (x *CompleteKycResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type UpdateKycRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname  string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Age       int64  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	BirthDate string `protobuf:"bytes,3,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Aadhar    string `protobuf:"bytes,4,opt,name=aadhar,proto3" json:"aadhar,omitempty"`
	Pan       string `protobuf:"bytes,5,opt,name=pan,proto3" json:"pan,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UpdateKycRequest) Reset() {
	*x = UpdateKycRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kyc_v1_kyc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateKycRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateKycRequest) ProtoMessage() {}

func (x *UpdateKycRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kyc_v1_kyc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateKycRequest.ProtoReflect.Descriptor instead.
func (*UpdateKycRequest) Descriptor() ([]byte, []int) {
	return file_proto_kyc_v1_kyc_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateKycRequest) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *UpdateKycRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UpdateKycRequest) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *UpdateKycRequest) GetAadhar() string {
	if x != nil {
		return x.Aadhar
	}
	return ""
}

func (x *UpdateKycRequest) GetPan() string {
	if x != nil {
		return x.Pan
	}
	return ""
}

func (x *UpdateKycRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type UpdateKycResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateKycResponse) Reset() {
	*x = UpdateKycResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kyc_v1_kyc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateKycResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateKycResponse) ProtoMessage() {}

func (x *UpdateKycResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kyc_v1_kyc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateKycResponse.ProtoReflect.Descriptor instead.
func (*UpdateKycResponse) Descriptor() ([]byte, []int) {
	return file_proto_kyc_v1_kyc_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateKycResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type DeleteKycRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteKycRequest) Reset() {
	*x = DeleteKycRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kyc_v1_kyc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKycRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKycRequest) ProtoMessage() {}

func (x *DeleteKycRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kyc_v1_kyc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKycRequest.ProtoReflect.Descriptor instead.
func (*DeleteKycRequest) Descriptor() ([]byte, []int) {
	return file_proto_kyc_v1_kyc_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteKycRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteKycResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteKycResponse) Reset() {
	*x = DeleteKycResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kyc_v1_kyc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKycResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKycResponse) ProtoMessage() {}

func (x *DeleteKycResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kyc_v1_kyc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKycResponse.ProtoReflect.Descriptor instead.
func (*DeleteKycResponse) Descriptor() ([]byte, []int) {
	return file_proto_kyc_v1_kyc_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteKycResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_proto_kyc_v1_kyc_proto protoreflect.FileDescriptor

var file_proto_kyc_v1_kyc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x79, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6b,
	0x79, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x22, 0xa1, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x61,
	0x64, 0x68, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x61, 0x64, 0x68,
	0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x70, 0x61, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x61, 0x64, 0x68, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x61, 0x64,
	0x68, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x70, 0x61, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x2b, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x82, 0x02, 0x0a, 0x0a, 0x4b, 0x79,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4b,
	0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x12, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x12, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x6b,
	0x72, 0x69, 0x73, 0x68, 0x37, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x68, 0x69, 0x70, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6b, 0x79, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x79, 0x63,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_kyc_v1_kyc_proto_rawDescOnce sync.Once
	file_proto_kyc_v1_kyc_proto_rawDescData = file_proto_kyc_v1_kyc_proto_rawDesc
)

func file_proto_kyc_v1_kyc_proto_rawDescGZIP() []byte {
	file_proto_kyc_v1_kyc_proto_rawDescOnce.Do(func() {
		file_proto_kyc_v1_kyc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_kyc_v1_kyc_proto_rawDescData)
	})
	return file_proto_kyc_v1_kyc_proto_rawDescData
}

var file_proto_kyc_v1_kyc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_kyc_v1_kyc_proto_goTypes = []interface{}{
	(*CompleteKycRequest)(nil),  // 0: proto.kyc.v1.CompleteKycRequest
	(*CompleteKycResponse)(nil), // 1: proto.kyc.v1.CompleteKycResponse
	(*UpdateKycRequest)(nil),    // 2: proto.kyc.v1.UpdateKycRequest
	(*UpdateKycResponse)(nil),   // 3: proto.kyc.v1.UpdateKycResponse
	(*DeleteKycRequest)(nil),    // 4: proto.kyc.v1.DeleteKycRequest
	(*DeleteKycResponse)(nil),   // 5: proto.kyc.v1.DeleteKycResponse
}
var file_proto_kyc_v1_kyc_proto_depIdxs = []int32{
	0, // 0: proto.kyc.v1.KycService.CompleteKyc:input_type -> proto.kyc.v1.CompleteKycRequest
	2, // 1: proto.kyc.v1.KycService.UpdateKyc:input_type -> proto.kyc.v1.UpdateKycRequest
	4, // 2: proto.kyc.v1.KycService.DeleteKyc:input_type -> proto.kyc.v1.DeleteKycRequest
	1, // 3: proto.kyc.v1.KycService.CompleteKyc:output_type -> proto.kyc.v1.CompleteKycResponse
	3, // 4: proto.kyc.v1.KycService.UpdateKyc:output_type -> proto.kyc.v1.UpdateKycResponse
	5, // 5: proto.kyc.v1.KycService.DeleteKyc:output_type -> proto.kyc.v1.DeleteKycResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_kyc_v1_kyc_proto_init() }
func file_proto_kyc_v1_kyc_proto_init() {
	if File_proto_kyc_v1_kyc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_kyc_v1_kyc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteKycRequest); i {
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
		file_proto_kyc_v1_kyc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteKycResponse); i {
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
		file_proto_kyc_v1_kyc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateKycRequest); i {
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
		file_proto_kyc_v1_kyc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateKycResponse); i {
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
		file_proto_kyc_v1_kyc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKycRequest); i {
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
		file_proto_kyc_v1_kyc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKycResponse); i {
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
			RawDescriptor: file_proto_kyc_v1_kyc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kyc_v1_kyc_proto_goTypes,
		DependencyIndexes: file_proto_kyc_v1_kyc_proto_depIdxs,
		MessageInfos:      file_proto_kyc_v1_kyc_proto_msgTypes,
	}.Build()
	File_proto_kyc_v1_kyc_proto = out.File
	file_proto_kyc_v1_kyc_proto_rawDesc = nil
	file_proto_kyc_v1_kyc_proto_goTypes = nil
	file_proto_kyc_v1_kyc_proto_depIdxs = nil
}
