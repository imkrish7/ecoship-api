// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: rpc/persona/v1/persona.proto

package personav1

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

type CreatePersonaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Persona *Persona `protobuf:"bytes,2,opt,name=persona,proto3" json:"persona,omitempty"`
}

func (x *CreatePersonaRequest) Reset() {
	*x = CreatePersonaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_persona_v1_persona_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonaRequest) ProtoMessage() {}

func (x *CreatePersonaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_persona_v1_persona_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonaRequest.ProtoReflect.Descriptor instead.
func (*CreatePersonaRequest) Descriptor() ([]byte, []int) {
	return file_rpc_persona_v1_persona_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePersonaRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreatePersonaRequest) GetPersona() *Persona {
	if x != nil {
		return x.Persona
	}
	return nil
}

type Persona struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone   string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Persona) Reset() {
	*x = Persona{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_persona_v1_persona_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Persona) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Persona) ProtoMessage() {}

func (x *Persona) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_persona_v1_persona_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Persona.ProtoReflect.Descriptor instead.
func (*Persona) Descriptor() ([]byte, []int) {
	return file_rpc_persona_v1_persona_proto_rawDescGZIP(), []int{1}
}

func (x *Persona) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Persona) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Persona) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Persona) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CreatePersonaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	NextStep string `protobuf:"bytes,2,opt,name=next_step,json=nextStep,proto3" json:"next_step,omitempty"`
}

func (x *CreatePersonaResponse) Reset() {
	*x = CreatePersonaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_persona_v1_persona_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonaResponse) ProtoMessage() {}

func (x *CreatePersonaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_persona_v1_persona_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonaResponse.ProtoReflect.Descriptor instead.
func (*CreatePersonaResponse) Descriptor() ([]byte, []int) {
	return file_rpc_persona_v1_persona_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePersonaResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *CreatePersonaResponse) GetNextStep() string {
	if x != nil {
		return x.NextStep
	}
	return ""
}

type VerifyAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VerifyAccountRequest) Reset() {
	*x = VerifyAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_persona_v1_persona_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyAccountRequest) ProtoMessage() {}

func (x *VerifyAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_persona_v1_persona_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyAccountRequest.ProtoReflect.Descriptor instead.
func (*VerifyAccountRequest) Descriptor() ([]byte, []int) {
	return file_rpc_persona_v1_persona_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *VerifyAccountRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	NextStep string `protobuf:"bytes,2,opt,name=next_step,json=nextStep,proto3" json:"next_step,omitempty"`
}

func (x *VerifyAccountResponse) Reset() {
	*x = VerifyAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_persona_v1_persona_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyAccountResponse) ProtoMessage() {}

func (x *VerifyAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_persona_v1_persona_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyAccountResponse.ProtoReflect.Descriptor instead.
func (*VerifyAccountResponse) Descriptor() ([]byte, []int) {
	return file_rpc_persona_v1_persona_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyAccountResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *VerifyAccountResponse) GetNextStep() string {
	if x != nil {
		return x.NextStep
	}
	return ""
}

type ResendVerifyLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ResendVerifyLinkRequest) Reset() {
	*x = ResendVerifyLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_persona_v1_persona_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResendVerifyLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResendVerifyLinkRequest) ProtoMessage() {}

func (x *ResendVerifyLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_persona_v1_persona_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResendVerifyLinkRequest.ProtoReflect.Descriptor instead.
func (*ResendVerifyLinkRequest) Descriptor() ([]byte, []int) {
	return file_rpc_persona_v1_persona_proto_rawDescGZIP(), []int{5}
}

func (x *ResendVerifyLinkRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResendVerifyLinkRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ResendVerifyLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResendVerifyLinkResponse) Reset() {
	*x = ResendVerifyLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_persona_v1_persona_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResendVerifyLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResendVerifyLinkResponse) ProtoMessage() {}

func (x *ResendVerifyLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_persona_v1_persona_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResendVerifyLinkResponse.ProtoReflect.Descriptor instead.
func (*ResendVerifyLinkResponse) Descriptor() ([]byte, []int) {
	return file_rpc_persona_v1_persona_proto_rawDescGZIP(), []int{6}
}

func (x *ResendVerifyLinkResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_rpc_persona_v1_persona_proto protoreflect.FileDescriptor

var file_rpc_persona_v1_persona_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x22, 0x59, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x52, 0x07, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x22, 0x63, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4c, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x65, 0x70, 0x22, 0x42, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4c, 0x0a, 0x15,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x65, 0x70, 0x22, 0x4b, 0x0a, 0x17, 0x52, 0x65,
	0x73, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x32, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x65, 0x6e,
	0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xa1, 0x02, 0x0a, 0x0e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x12,
	0x20, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f,
	0x0a, 0x10, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x23, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d,
	0x6b, 0x72, 0x69, 0x73, 0x68, 0x37, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x68, 0x69, 0x70, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x65, 0x72, 0x6f, 0x73,
	0x6f, 0x6e, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_persona_v1_persona_proto_rawDescOnce sync.Once
	file_rpc_persona_v1_persona_proto_rawDescData = file_rpc_persona_v1_persona_proto_rawDesc
)

func file_rpc_persona_v1_persona_proto_rawDescGZIP() []byte {
	file_rpc_persona_v1_persona_proto_rawDescOnce.Do(func() {
		file_rpc_persona_v1_persona_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_persona_v1_persona_proto_rawDescData)
	})
	return file_rpc_persona_v1_persona_proto_rawDescData
}

var file_rpc_persona_v1_persona_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_persona_v1_persona_proto_goTypes = []interface{}{
	(*CreatePersonaRequest)(nil),     // 0: persona.v1.CreatePersonaRequest
	(*Persona)(nil),                  // 1: persona.v1.Persona
	(*CreatePersonaResponse)(nil),    // 2: persona.v1.CreatePersonaResponse
	(*VerifyAccountRequest)(nil),     // 3: persona.v1.VerifyAccountRequest
	(*VerifyAccountResponse)(nil),    // 4: persona.v1.VerifyAccountResponse
	(*ResendVerifyLinkRequest)(nil),  // 5: persona.v1.ResendVerifyLinkRequest
	(*ResendVerifyLinkResponse)(nil), // 6: persona.v1.ResendVerifyLinkResponse
}
var file_rpc_persona_v1_persona_proto_depIdxs = []int32{
	1, // 0: persona.v1.CreatePersonaRequest.persona:type_name -> persona.v1.Persona
	0, // 1: persona.v1.PersonaService.CreatePersona:input_type -> persona.v1.CreatePersonaRequest
	3, // 2: persona.v1.PersonaService.VerifyAccount:input_type -> persona.v1.VerifyAccountRequest
	5, // 3: persona.v1.PersonaService.ResendVerifyLink:input_type -> persona.v1.ResendVerifyLinkRequest
	2, // 4: persona.v1.PersonaService.CreatePersona:output_type -> persona.v1.CreatePersonaResponse
	4, // 5: persona.v1.PersonaService.VerifyAccount:output_type -> persona.v1.VerifyAccountResponse
	6, // 6: persona.v1.PersonaService.ResendVerifyLink:output_type -> persona.v1.ResendVerifyLinkResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_persona_v1_persona_proto_init() }
func file_rpc_persona_v1_persona_proto_init() {
	if File_rpc_persona_v1_persona_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_persona_v1_persona_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonaRequest); i {
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
		file_rpc_persona_v1_persona_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Persona); i {
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
		file_rpc_persona_v1_persona_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonaResponse); i {
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
		file_rpc_persona_v1_persona_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyAccountRequest); i {
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
		file_rpc_persona_v1_persona_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyAccountResponse); i {
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
		file_rpc_persona_v1_persona_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendVerifyLinkRequest); i {
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
		file_rpc_persona_v1_persona_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendVerifyLinkResponse); i {
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
			RawDescriptor: file_rpc_persona_v1_persona_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_persona_v1_persona_proto_goTypes,
		DependencyIndexes: file_rpc_persona_v1_persona_proto_depIdxs,
		MessageInfos:      file_rpc_persona_v1_persona_proto_msgTypes,
	}.Build()
	File_rpc_persona_v1_persona_proto = out.File
	file_rpc_persona_v1_persona_proto_rawDesc = nil
	file_rpc_persona_v1_persona_proto_goTypes = nil
	file_rpc_persona_v1_persona_proto_depIdxs = nil
}