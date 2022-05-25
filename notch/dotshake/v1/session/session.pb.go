// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.1
// source: notch/dotshake/v1/session.proto

package session

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

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// auth0で登録したメールアドレス
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notch_dotshake_v1_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notch_dotshake_v1_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_notch_dotshake_v1_session_proto_rawDescGZIP(), []int{0}
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 初回ログインがどうかを判断するフラグ
	IsFirst bool `protobuf:"varint,1,opt,name=isFirst,proto3" json:"isFirst,omitempty"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notch_dotshake_v1_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notch_dotshake_v1_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_notch_dotshake_v1_session_proto_rawDescGZIP(), []int{1}
}

func (x *SignInResponse) GetIsFirst() bool {
	if x != nil {
		return x.IsFirst
	}
	return false
}

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// auth0のuserID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// 端末の名前
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// 端末のOS
	Os string `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notch_dotshake_v1_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notch_dotshake_v1_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_notch_dotshake_v1_session_proto_rawDescGZIP(), []int{2}
}

func (x *SignUpRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *SignUpRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SignUpRequest) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 使用するwireguardのIPアドレス
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// 使用するwireguardのIPアドレスのCIDR。今は/24がデフォルトで返ってくる
	Cidr string `protobuf:"bytes,2,opt,name=cidr,proto3" json:"cidr,omitempty"`
	// UDP Hole Punchingするために必要なSignalServerのホストURL
	SignalServerHost string `protobuf:"bytes,3,opt,name=signalServerHost,proto3" json:"signalServerHost,omitempty"`
	SignalServerPort uint64 `protobuf:"varint,4,opt,name=signalServerPort,proto3" json:"signalServerPort,omitempty"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notch_dotshake_v1_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notch_dotshake_v1_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_notch_dotshake_v1_session_proto_rawDescGZIP(), []int{3}
}

func (x *SignUpResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *SignUpResponse) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

func (x *SignUpResponse) GetSignalServerHost() string {
	if x != nil {
		return x.SignalServerHost
	}
	return ""
}

func (x *SignUpResponse) GetSignalServerPort() uint64 {
	if x != nil {
		return x.SignalServerPort
	}
	return 0
}

type VerifyPeerLoginSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// jwtの中に入っているユニークなid
	SessionID string `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
}

func (x *VerifyPeerLoginSessionRequest) Reset() {
	*x = VerifyPeerLoginSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notch_dotshake_v1_session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPeerLoginSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPeerLoginSessionRequest) ProtoMessage() {}

func (x *VerifyPeerLoginSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notch_dotshake_v1_session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPeerLoginSessionRequest.ProtoReflect.Descriptor instead.
func (*VerifyPeerLoginSessionRequest) Descriptor() ([]byte, []int) {
	return file_notch_dotshake_v1_session_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyPeerLoginSessionRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

type VerifyPeerLoginSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 使用するwireguardのIPアドレス
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// 端末の名前
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// 端末のOS
	Os string `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
}

func (x *VerifyPeerLoginSessionResponse) Reset() {
	*x = VerifyPeerLoginSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notch_dotshake_v1_session_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPeerLoginSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPeerLoginSessionResponse) ProtoMessage() {}

func (x *VerifyPeerLoginSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notch_dotshake_v1_session_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPeerLoginSessionResponse.ProtoReflect.Descriptor instead.
func (*VerifyPeerLoginSessionResponse) Descriptor() ([]byte, []int) {
	return file_notch_dotshake_v1_session_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyPeerLoginSessionResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *VerifyPeerLoginSessionResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *VerifyPeerLoginSessionResponse) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

var File_notch_dotshake_v1_session_proto protoreflect.FileDescriptor

var file_notch_dotshake_v1_session_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x6f, 0x74, 0x63, 0x68, 0x2f, 0x64, 0x6f, 0x74, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x0d,
	0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x64, 0x72,
	0x12, 0x2a, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x48, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x3d, 0x0a, 0x1d, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x54, 0x0a, 0x1e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x32, 0xf1, 0x01,
	0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x39, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x16, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x50, 0x65, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x50, 0x65, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notch_dotshake_v1_session_proto_rawDescOnce sync.Once
	file_notch_dotshake_v1_session_proto_rawDescData = file_notch_dotshake_v1_session_proto_rawDesc
)

func file_notch_dotshake_v1_session_proto_rawDescGZIP() []byte {
	file_notch_dotshake_v1_session_proto_rawDescOnce.Do(func() {
		file_notch_dotshake_v1_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_notch_dotshake_v1_session_proto_rawDescData)
	})
	return file_notch_dotshake_v1_session_proto_rawDescData
}

var file_notch_dotshake_v1_session_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_notch_dotshake_v1_session_proto_goTypes = []interface{}{
	(*SignInRequest)(nil),                  // 0: protos.SignInRequest
	(*SignInResponse)(nil),                 // 1: protos.SignInResponse
	(*SignUpRequest)(nil),                  // 2: protos.SignUpRequest
	(*SignUpResponse)(nil),                 // 3: protos.SignUpResponse
	(*VerifyPeerLoginSessionRequest)(nil),  // 4: protos.VerifyPeerLoginSessionRequest
	(*VerifyPeerLoginSessionResponse)(nil), // 5: protos.VerifyPeerLoginSessionResponse
}
var file_notch_dotshake_v1_session_proto_depIdxs = []int32{
	0, // 0: protos.SessionService.SignIn:input_type -> protos.SignInRequest
	2, // 1: protos.SessionService.SignUp:input_type -> protos.SignUpRequest
	4, // 2: protos.SessionService.VerifyPeerLoginSession:input_type -> protos.VerifyPeerLoginSessionRequest
	1, // 3: protos.SessionService.SignIn:output_type -> protos.SignInResponse
	3, // 4: protos.SessionService.SignUp:output_type -> protos.SignUpResponse
	5, // 5: protos.SessionService.VerifyPeerLoginSession:output_type -> protos.VerifyPeerLoginSessionResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notch_dotshake_v1_session_proto_init() }
func file_notch_dotshake_v1_session_proto_init() {
	if File_notch_dotshake_v1_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notch_dotshake_v1_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
		file_notch_dotshake_v1_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
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
		file_notch_dotshake_v1_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
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
		file_notch_dotshake_v1_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
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
		file_notch_dotshake_v1_session_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPeerLoginSessionRequest); i {
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
		file_notch_dotshake_v1_session_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPeerLoginSessionResponse); i {
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
			RawDescriptor: file_notch_dotshake_v1_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notch_dotshake_v1_session_proto_goTypes,
		DependencyIndexes: file_notch_dotshake_v1_session_proto_depIdxs,
		MessageInfos:      file_notch_dotshake_v1_session_proto_msgTypes,
	}.Build()
	File_notch_dotshake_v1_session_proto = out.File
	file_notch_dotshake_v1_session_proto_rawDesc = nil
	file_notch_dotshake_v1_session_proto_goTypes = nil
	file_notch_dotshake_v1_session_proto_depIdxs = nil
}
