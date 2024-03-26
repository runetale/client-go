// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: runetale/runetale/v1/oidc.proto

package oidc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub        string `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	TenantID   string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	Doamin     string `protobuf:"bytes,3,opt,name=doamin,proto3" json:"doamin,omitempty"`
	ProviderID string `protobuf:"bytes,4,opt,name=providerID,proto3" json:"providerID,omitempty"`
	Email      string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Username   string `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	Picture    string `protobuf:"bytes,7,opt,name=picture,proto3" json:"picture,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_oidc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_oidc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_oidc_proto_rawDescGZIP(), []int{0}
}

func (x *LoginResponse) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *LoginResponse) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *LoginResponse) GetDoamin() string {
	if x != nil {
		return x.Doamin
	}
	return ""
}

func (x *LoginResponse) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *LoginResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginResponse) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub        string `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	TenantID   string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	Doamin     string `protobuf:"bytes,3,opt,name=doamin,proto3" json:"doamin,omitempty"`
	ProviderID string `protobuf:"bytes,4,opt,name=providerID,proto3" json:"providerID,omitempty"`
	Email      string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Username   string `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	Picture    string `protobuf:"bytes,7,opt,name=picture,proto3" json:"picture,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_oidc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_oidc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_oidc_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *LoginRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *LoginRequest) GetDoamin() string {
	if x != nil {
		return x.Doamin
	}
	return ""
}

func (x *LoginRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

type AuthenticateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doamin       string `protobuf:"bytes,1,opt,name=doamin,proto3" json:"doamin,omitempty"`
	Email        string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username     string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Sub          string `protobuf:"bytes,4,opt,name=sub,proto3" json:"sub,omitempty"`
	IsRegistered bool   `protobuf:"varint,5,opt,name=isRegistered,proto3" json:"isRegistered,omitempty"`
}

func (x *AuthenticateResponse) Reset() {
	*x = AuthenticateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_oidc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateResponse) ProtoMessage() {}

func (x *AuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_oidc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_oidc_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticateResponse) GetDoamin() string {
	if x != nil {
		return x.Doamin
	}
	return ""
}

func (x *AuthenticateResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthenticateResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthenticateResponse) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *AuthenticateResponse) GetIsRegistered() bool {
	if x != nil {
		return x.IsRegistered
	}
	return false
}

var File_runetale_runetale_v1_oidc_proto protoreflect.FileDescriptor

var file_runetale_runetale_v1_oidc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x75, 0x6e, 0x65, 0x74, 0x61, 0x6c, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x65, 0x74,
	0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x61, 0x6d, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x61, 0x6d, 0x69, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x61,
	0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x61, 0x6d, 0x69,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0x96, 0x01,
	0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x61, 0x6d, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x61, 0x6d, 0x69, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x75, 0x62, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x32, 0xcf, 0x01, 0x0a, 0x0b, 0x4f, 0x49, 0x44, 0x43, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6f, 0x69,
	0x64, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runetale_runetale_v1_oidc_proto_rawDescOnce sync.Once
	file_runetale_runetale_v1_oidc_proto_rawDescData = file_runetale_runetale_v1_oidc_proto_rawDesc
)

func file_runetale_runetale_v1_oidc_proto_rawDescGZIP() []byte {
	file_runetale_runetale_v1_oidc_proto_rawDescOnce.Do(func() {
		file_runetale_runetale_v1_oidc_proto_rawDescData = protoimpl.X.CompressGZIP(file_runetale_runetale_v1_oidc_proto_rawDescData)
	})
	return file_runetale_runetale_v1_oidc_proto_rawDescData
}

var file_runetale_runetale_v1_oidc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_runetale_runetale_v1_oidc_proto_goTypes = []interface{}{
	(*LoginResponse)(nil),        // 0: protos.LoginResponse
	(*LoginRequest)(nil),         // 1: protos.LoginRequest
	(*AuthenticateResponse)(nil), // 2: protos.AuthenticateResponse
	(*emptypb.Empty)(nil),        // 3: google.protobuf.Empty
}
var file_runetale_runetale_v1_oidc_proto_depIdxs = []int32{
	1, // 0: protos.OIDCService.Login:input_type -> protos.LoginRequest
	3, // 1: protos.OIDCService.Authenticate:input_type -> google.protobuf.Empty
	3, // 2: protos.OIDCService.RefreshToken:input_type -> google.protobuf.Empty
	0, // 3: protos.OIDCService.Login:output_type -> protos.LoginResponse
	2, // 4: protos.OIDCService.Authenticate:output_type -> protos.AuthenticateResponse
	3, // 5: protos.OIDCService.RefreshToken:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_runetale_runetale_v1_oidc_proto_init() }
func file_runetale_runetale_v1_oidc_proto_init() {
	if File_runetale_runetale_v1_oidc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runetale_runetale_v1_oidc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_runetale_runetale_v1_oidc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_runetale_runetale_v1_oidc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateResponse); i {
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
			RawDescriptor: file_runetale_runetale_v1_oidc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runetale_runetale_v1_oidc_proto_goTypes,
		DependencyIndexes: file_runetale_runetale_v1_oidc_proto_depIdxs,
		MessageInfos:      file_runetale_runetale_v1_oidc_proto_msgTypes,
	}.Build()
	File_runetale_runetale_v1_oidc_proto = out.File
	file_runetale_runetale_v1_oidc_proto_rawDesc = nil
	file_runetale_runetale_v1_oidc_proto_goTypes = nil
	file_runetale_runetale_v1_oidc_proto_depIdxs = nil
}
