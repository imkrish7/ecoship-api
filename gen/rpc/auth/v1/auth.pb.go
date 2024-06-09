// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: rpc/auth/v1/auth.proto

package authv1

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

type SignupServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Type     string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *SignupServiceRequest) Reset() {
	*x = SignupServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_auth_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupServiceRequest) ProtoMessage() {}

func (x *SignupServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_auth_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupServiceRequest.ProtoReflect.Descriptor instead.
func (*SignupServiceRequest) Descriptor() ([]byte, []int) {
	return file_rpc_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SignupServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignupServiceRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupServiceRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SignupServiceRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignupServiceRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type SignupServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SignupServiceResponse) Reset() {
	*x = SignupServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_auth_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupServiceResponse) ProtoMessage() {}

func (x *SignupServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_auth_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupServiceResponse.ProtoReflect.Descriptor instead.
func (*SignupServiceResponse) Descriptor() ([]byte, []int) {
	return file_rpc_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SignupServiceResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type LoginServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginServiceRequest) Reset() {
	*x = LoginServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_auth_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginServiceRequest) ProtoMessage() {}

func (x *LoginServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_auth_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginServiceRequest.ProtoReflect.Descriptor instead.
func (*LoginServiceRequest) Descriptor() ([]byte, []int) {
	return file_rpc_auth_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginServiceRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginServiceRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	NextStep string `protobuf:"bytes,3,opt,name=next_step,json=nextStep,proto3" json:"next_step,omitempty"`
	Auth     string `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *LoginServiceResponse) Reset() {
	*x = LoginServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_auth_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginServiceResponse) ProtoMessage() {}

func (x *LoginServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_auth_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginServiceResponse.ProtoReflect.Descriptor instead.
func (*LoginServiceResponse) Descriptor() ([]byte, []int) {
	return file_rpc_auth_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginServiceResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *LoginServiceResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LoginServiceResponse) GetNextStep() string {
	if x != nil {
		return x.NextStep
	}
	return ""
}

func (x *LoginServiceResponse) GetAuth() string {
	if x != nil {
		return x.Auth
	}
	return ""
}

var File_rpc_auth_v1_auth_proto protoreflect.FileDescriptor

var file_rpc_auth_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2f, 0x0a, 0x15, 0x53, 0x69,
	0x67, 0x6e, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x47, 0x0a, 0x13, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x73, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x78,
	0x74, 0x53, 0x74, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x32, 0xae, 0x01, 0x0a, 0x0b, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x6b, 0x72, 0x69, 0x73, 0x68,
	0x37, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x68, 0x69, 0x70, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75,
	0x74, 0x68, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_auth_v1_auth_proto_rawDescOnce sync.Once
	file_rpc_auth_v1_auth_proto_rawDescData = file_rpc_auth_v1_auth_proto_rawDesc
)

func file_rpc_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_rpc_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_rpc_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_auth_v1_auth_proto_rawDescData)
	})
	return file_rpc_auth_v1_auth_proto_rawDescData
}

var file_rpc_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_auth_v1_auth_proto_goTypes = []interface{}{
	(*SignupServiceRequest)(nil),  // 0: auth.v1.SignupServiceRequest
	(*SignupServiceResponse)(nil), // 1: auth.v1.SignupServiceResponse
	(*LoginServiceRequest)(nil),   // 2: auth.v1.LoginServiceRequest
	(*LoginServiceResponse)(nil),  // 3: auth.v1.LoginServiceResponse
}
var file_rpc_auth_v1_auth_proto_depIdxs = []int32{
	2, // 0: auth.v1.AuthService.LoginService:input_type -> auth.v1.LoginServiceRequest
	0, // 1: auth.v1.AuthService.SignupService:input_type -> auth.v1.SignupServiceRequest
	3, // 2: auth.v1.AuthService.LoginService:output_type -> auth.v1.LoginServiceResponse
	1, // 3: auth.v1.AuthService.SignupService:output_type -> auth.v1.SignupServiceResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_auth_v1_auth_proto_init() }
func file_rpc_auth_v1_auth_proto_init() {
	if File_rpc_auth_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_auth_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupServiceRequest); i {
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
		file_rpc_auth_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupServiceResponse); i {
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
		file_rpc_auth_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginServiceRequest); i {
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
		file_rpc_auth_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginServiceResponse); i {
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
			RawDescriptor: file_rpc_auth_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_rpc_auth_v1_auth_proto_depIdxs,
		MessageInfos:      file_rpc_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_rpc_auth_v1_auth_proto = out.File
	file_rpc_auth_v1_auth_proto_rawDesc = nil
	file_rpc_auth_v1_auth_proto_goTypes = nil
	file_rpc_auth_v1_auth_proto_depIdxs = nil
}