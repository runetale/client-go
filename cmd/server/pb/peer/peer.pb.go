// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: protos/peer.proto

package peer

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

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKey       string `protobuf:"bytes,1,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	ClientMachineKey string `protobuf:"bytes,2,opt,name=clientMachineKey,proto3" json:"clientMachineKey,omitempty"`
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_protos_peer_proto_rawDescGZIP(), []int{0}
}

func (x *SyncRequest) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *SyncRequest) GetClientMachineKey() string {
	if x != nil {
		return x.ClientMachineKey
	}
	return ""
}

type SyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerConfig        *PeerConfig     `protobuf:"bytes,1,opt,name=peerConfig,proto3" json:"peerConfig,omitempty"`
	RemotePeers       []*RemotePeer   `protobuf:"bytes,2,rep,name=remotePeers,proto3" json:"remotePeers,omitempty"`
	RemotePeerIsEmpty bool            `protobuf:"varint,3,opt,name=remotePeerIsEmpty,proto3" json:"remotePeerIsEmpty,omitempty"`
	StunTurnConfig    *StunTurnConfig `protobuf:"bytes,4,opt,name=stunTurnConfig,proto3" json:"stunTurnConfig,omitempty"`
}

func (x *SyncResponse) Reset() {
	*x = SyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncResponse) ProtoMessage() {}

func (x *SyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncResponse.ProtoReflect.Descriptor instead.
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return file_protos_peer_proto_rawDescGZIP(), []int{1}
}

func (x *SyncResponse) GetPeerConfig() *PeerConfig {
	if x != nil {
		return x.PeerConfig
	}
	return nil
}

func (x *SyncResponse) GetRemotePeers() []*RemotePeer {
	if x != nil {
		return x.RemotePeers
	}
	return nil
}

func (x *SyncResponse) GetRemotePeerIsEmpty() bool {
	if x != nil {
		return x.RemotePeerIsEmpty
	}
	return false
}

func (x *SyncResponse) GetStunTurnConfig() *StunTurnConfig {
	if x != nil {
		return x.StunTurnConfig
	}
	return nil
}

type RemotePeer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientMachineKey string   `protobuf:"bytes,1,opt,name=clientMachineKey,proto3" json:"clientMachineKey,omitempty"`
	WgPubKey         string   `protobuf:"bytes,2,opt,name=wgPubKey,proto3" json:"wgPubKey,omitempty"`
	AllowedIps       []string `protobuf:"bytes,3,rep,name=allowedIps,proto3" json:"allowedIps,omitempty"`
}

func (x *RemotePeer) Reset() {
	*x = RemotePeer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_peer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemotePeer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemotePeer) ProtoMessage() {}

func (x *RemotePeer) ProtoReflect() protoreflect.Message {
	mi := &file_protos_peer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemotePeer.ProtoReflect.Descriptor instead.
func (*RemotePeer) Descriptor() ([]byte, []int) {
	return file_protos_peer_proto_rawDescGZIP(), []int{2}
}

func (x *RemotePeer) GetClientMachineKey() string {
	if x != nil {
		return x.ClientMachineKey
	}
	return ""
}

func (x *RemotePeer) GetWgPubKey() string {
	if x != nil {
		return x.WgPubKey
	}
	return ""
}

func (x *RemotePeer) GetAllowedIps() []string {
	if x != nil {
		return x.AllowedIps
	}
	return nil
}

type PeerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Dns     string `protobuf:"bytes,2,opt,name=dns,proto3" json:"dns,omitempty"`
}

func (x *PeerConfig) Reset() {
	*x = PeerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_peer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerConfig) ProtoMessage() {}

func (x *PeerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_protos_peer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerConfig.ProtoReflect.Descriptor instead.
func (*PeerConfig) Descriptor() ([]byte, []int) {
	return file_protos_peer_proto_rawDescGZIP(), []int{3}
}

func (x *PeerConfig) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PeerConfig) GetDns() string {
	if x != nil {
		return x.Dns
	}
	return ""
}

type StunTurnConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stuns           []*Host           `protobuf:"bytes,1,rep,name=stuns,proto3" json:"stuns,omitempty"`
	TurnCredentials []*TurnCredential `protobuf:"bytes,2,rep,name=turnCredentials,proto3" json:"turnCredentials,omitempty"`
	Signal          *Host             `protobuf:"bytes,3,opt,name=signal,proto3" json:"signal,omitempty"`
}

func (x *StunTurnConfig) Reset() {
	*x = StunTurnConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_peer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StunTurnConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StunTurnConfig) ProtoMessage() {}

func (x *StunTurnConfig) ProtoReflect() protoreflect.Message {
	mi := &file_protos_peer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StunTurnConfig.ProtoReflect.Descriptor instead.
func (*StunTurnConfig) Descriptor() ([]byte, []int) {
	return file_protos_peer_proto_rawDescGZIP(), []int{4}
}

func (x *StunTurnConfig) GetStuns() []*Host {
	if x != nil {
		return x.Stuns
	}
	return nil
}

func (x *StunTurnConfig) GetTurnCredentials() []*TurnCredential {
	if x != nil {
		return x.TurnCredentials
	}
	return nil
}

func (x *StunTurnConfig) GetSignal() *Host {
	if x != nil {
		return x.Signal
	}
	return nil
}

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_peer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_protos_peer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_protos_peer_proto_rawDescGZIP(), []int{5}
}

func (x *Host) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TurnCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	User     string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *TurnCredential) Reset() {
	*x = TurnCredential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_peer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TurnCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TurnCredential) ProtoMessage() {}

func (x *TurnCredential) ProtoReflect() protoreflect.Message {
	mi := &file_protos_peer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TurnCredential.ProtoReflect.Descriptor instead.
func (*TurnCredential) Descriptor() ([]byte, []int) {
	return file_protos_peer_proto_rawDescGZIP(), []int{6}
}

func (x *TurnCredential) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *TurnCredential) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *TurnCredential) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_protos_peer_proto protoreflect.FileDescriptor

var file_protos_peer_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x59, 0x0a, 0x0b, 0x53,
	0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x22, 0xe6, 0x01, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0a, 0x70, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x0b, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x50, 0x65, 0x65, 0x72, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72,
	0x73, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x49,
	0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x49, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x3e, 0x0a, 0x0e, 0x73, 0x74, 0x75, 0x6e, 0x54, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x53, 0x74, 0x75, 0x6e, 0x54, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0e, 0x73, 0x74, 0x75, 0x6e, 0x54, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x74, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x67, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x67, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x49, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x49, 0x70, 0x73, 0x22, 0x38, 0x0a, 0x0a, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6e, 0x73, 0x22,
	0x9c, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x6e, 0x54, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x74, 0x75, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x05, 0x73, 0x74, 0x75, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x0f, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x75, 0x72, 0x6e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0f, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x22, 0x18,
	0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x52, 0x0a, 0x0e, 0x54, 0x75, 0x72, 0x6e,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x44, 0x0a, 0x0b,
	0x50, 0x65, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x53,
	0x79, 0x6e, 0x63, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x79, 0x6e,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_peer_proto_rawDescOnce sync.Once
	file_protos_peer_proto_rawDescData = file_protos_peer_proto_rawDesc
)

func file_protos_peer_proto_rawDescGZIP() []byte {
	file_protos_peer_proto_rawDescOnce.Do(func() {
		file_protos_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_peer_proto_rawDescData)
	})
	return file_protos_peer_proto_rawDescData
}

var file_protos_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_peer_proto_goTypes = []interface{}{
	(*SyncRequest)(nil),    // 0: protos.SyncRequest
	(*SyncResponse)(nil),   // 1: protos.SyncResponse
	(*RemotePeer)(nil),     // 2: protos.RemotePeer
	(*PeerConfig)(nil),     // 3: protos.PeerConfig
	(*StunTurnConfig)(nil), // 4: protos.StunTurnConfig
	(*Host)(nil),           // 5: protos.Host
	(*TurnCredential)(nil), // 6: protos.TurnCredential
}
var file_protos_peer_proto_depIdxs = []int32{
	3, // 0: protos.SyncResponse.peerConfig:type_name -> protos.PeerConfig
	2, // 1: protos.SyncResponse.remotePeers:type_name -> protos.RemotePeer
	4, // 2: protos.SyncResponse.stunTurnConfig:type_name -> protos.StunTurnConfig
	5, // 3: protos.StunTurnConfig.stuns:type_name -> protos.Host
	6, // 4: protos.StunTurnConfig.turnCredentials:type_name -> protos.TurnCredential
	5, // 5: protos.StunTurnConfig.signal:type_name -> protos.Host
	0, // 6: protos.PeerService.Sync:input_type -> protos.SyncRequest
	1, // 7: protos.PeerService.Sync:output_type -> protos.SyncResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protos_peer_proto_init() }
func file_protos_peer_proto_init() {
	if File_protos_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_protos_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncResponse); i {
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
		file_protos_peer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemotePeer); i {
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
		file_protos_peer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerConfig); i {
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
		file_protos_peer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StunTurnConfig); i {
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
		file_protos_peer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_protos_peer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TurnCredential); i {
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
			RawDescriptor: file_protos_peer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_peer_proto_goTypes,
		DependencyIndexes: file_protos_peer_proto_depIdxs,
		MessageInfos:      file_protos_peer_proto_msgTypes,
	}.Build()
	File_protos_peer_proto = out.File
	file_protos_peer_proto_rawDesc = nil
	file_protos_peer_proto_goTypes = nil
	file_protos_peer_proto_depIdxs = nil
}
