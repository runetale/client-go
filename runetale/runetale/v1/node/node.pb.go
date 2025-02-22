// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: runetale/runetale/v1/node.proto

package node

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

type SyncNodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsEmpty     bool    `protobuf:"varint,1,opt,name=isEmpty,proto3" json:"isEmpty,omitempty"`
	RemoteNodes []*Node `protobuf:"bytes,2,rep,name=remoteNodes,proto3" json:"remoteNodes,omitempty"`
	Ip          string  `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`     // host ip
	Cidr        string  `protobuf:"bytes,4,opt,name=cidr,proto3" json:"cidr,omitempty"` // host cidr
}

func (x *SyncNodesResponse) Reset() {
	*x = SyncNodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncNodesResponse) ProtoMessage() {}

func (x *SyncNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncNodesResponse.ProtoReflect.Descriptor instead.
func (*SyncNodesResponse) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_node_proto_rawDescGZIP(), []int{0}
}

func (x *SyncNodesResponse) GetIsEmpty() bool {
	if x != nil {
		return x.IsEmpty
	}
	return false
}

func (x *SyncNodesResponse) GetRemoteNodes() []*Node {
	if x != nil {
		return x.RemoteNodes
	}
	return nil
}

func (x *SyncNodesResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *SyncNodesResponse) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NodeId     uint64   `protobuf:"varint,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	NodeKey    string   `protobuf:"bytes,3,opt,name=nodeKey,proto3" json:"nodeKey,omitempty"`
	WgPubKey   string   `protobuf:"bytes,4,opt,name=wgPubKey,proto3" json:"wgPubKey,omitempty"`
	AllowedIPs []string `protobuf:"bytes,5,rep,name=allowedIPs,proto3" json:"allowedIPs,omitempty"`
	Ip         string   `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`
	Cidr       string   `protobuf:"bytes,7,opt,name=cidr,proto3" json:"cidr,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_node_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetNodeId() uint64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *Node) GetNodeKey() string {
	if x != nil {
		return x.NodeKey
	}
	return ""
}

func (x *Node) GetWgPubKey() string {
	if x != nil {
		return x.WgPubKey
	}
	return ""
}

func (x *Node) GetAllowedIPs() []string {
	if x != nil {
		return x.AllowedIPs
	}
	return nil
}

func (x *Node) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Node) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

type ComposeNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Cidr string `protobuf:"bytes,2,opt,name=cidr,proto3" json:"cidr,omitempty"`
}

func (x *ComposeNodeResponse) Reset() {
	*x = ComposeNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeNodeResponse) ProtoMessage() {}

func (x *ComposeNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeNodeResponse.ProtoReflect.Descriptor instead.
func (*ComposeNodeResponse) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_node_proto_rawDescGZIP(), []int{2}
}

func (x *ComposeNodeResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ComposeNodeResponse) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

type NetPortRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 以下のような形式
	// - IPv4 or IPv6の単一のIPアドレス
	// - "*" は全て許可
	// - "192.168.0.0/16" cidrが含まれたipの範囲
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// portのフォーマットは
	// - 全て指定の `*` か
	// - 単一指定の `22` か
	// - 複数指定の `80, 443` か `2つまで`
	// - 範囲指定の `100-200“
	// - 単一のportの場合はlastにも同じポート番号が入る
	Ports *NetPortRangePortRange `protobuf:"bytes,2,opt,name=ports,proto3" json:"ports,omitempty"`
	// advertiseすることが許可されたIP範囲
	// 1.2.3.4/16のIP+Maskの形
	// "10.0.0.0/8,192.172.0.0/24"のようにcommaで区切る
	AdvertisedRoutes string `protobuf:"bytes,3,opt,name=advertisedRoutes,proto3" json:"advertisedRoutes,omitempty"`
}

func (x *NetPortRange) Reset() {
	*x = NetPortRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetPortRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetPortRange) ProtoMessage() {}

func (x *NetPortRange) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetPortRange.ProtoReflect.Descriptor instead.
func (*NetPortRange) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_node_proto_rawDescGZIP(), []int{3}
}

func (x *NetPortRange) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *NetPortRange) GetPorts() *NetPortRangePortRange {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *NetPortRange) GetAdvertisedRoutes() string {
	if x != nil {
		return x.AdvertisedRoutes
	}
	return ""
}

type FilterRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// source ips,
	// - "192.168.0.0/16" cidrが含まれたipの範囲
	SrcIps []string `protobuf:"bytes,1,rep,name=srcIps,proto3" json:"srcIps,omitempty"`
	// dstのpeerのリスト
	Dsts []*NetPortRange `protobuf:"bytes,2,rep,name=dsts,proto3" json:"dsts,omitempty"`
	// https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	// protocol numbers
	// Unknown = 0x00
	// ICMPv4  = 0x01
	// ICMPv6  = 0x3a
	// TCP     = 0x06
	// UDP     = 0x11
	IPProto []uint32 `protobuf:"varint,3,rep,packed,name=iPProto,proto3" json:"iPProto,omitempty"`
}

func (x *FilterRule) Reset() {
	*x = FilterRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterRule) ProtoMessage() {}

func (x *FilterRule) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterRule.ProtoReflect.Descriptor instead.
func (*FilterRule) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_node_proto_rawDescGZIP(), []int{4}
}

func (x *FilterRule) GetSrcIps() []string {
	if x != nil {
		return x.SrcIps
	}
	return nil
}

func (x *FilterRule) GetDsts() []*NetPortRange {
	if x != nil {
		return x.Dsts
	}
	return nil
}

func (x *FilterRule) GetIPProto() []uint32 {
	if x != nil {
		return x.IPProto
	}
	return nil
}

type NetworkMapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// このmapのsequential id
	Seq uint64 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	// このNodeの情報
	Node *Node `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	// このNodeがアクセスするpeers, つまりremote nodesの情報が含まれている
	Peers []*Node `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty"`
	// 変更があった場合のPeers
	// serverで差分更新される
	PeersChanged []*Node `protobuf:"bytes,4,rep,name=peersChanged,proto3" json:"peersChanged,omitempty"`
	// 消された場合のPeersのNodeID
	PeersRemoved []uint64 `protobuf:"varint,5,rep,packed,name=peersRemoved,proto3" json:"peersRemoved,omitempty"`
	// Firewall Rules
	PacketFilter []*FilterRule `protobuf:"bytes,6,rep,name=packetFilter,proto3" json:"packetFilter,omitempty"`
	// このnodeがadvertiseするIPアドレス
	// 1.2.3.4/16のIP+Maskの形
	// "10.0.0.0/8,192.172.0.0/24"のようにcommaで区切る
	AdvertisedRoute string `protobuf:"bytes,7,opt,name=advertisedRoute,proto3" json:"advertisedRoute,omitempty"`
	// jailedがtrueの場合全てのパケットを拒否する
	// defaultの状態はこの状態である
	Jailed   bool    `protobuf:"varint,8,opt,name=Jailed,proto3" json:"Jailed,omitempty"`
	IceTable []*Node `protobuf:"bytes,9,rep,name=iceTable,proto3" json:"iceTable,omitempty"`
}

func (x *NetworkMapResponse) Reset() {
	*x = NetworkMapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkMapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkMapResponse) ProtoMessage() {}

func (x *NetworkMapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkMapResponse.ProtoReflect.Descriptor instead.
func (*NetworkMapResponse) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_node_proto_rawDescGZIP(), []int{5}
}

func (x *NetworkMapResponse) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *NetworkMapResponse) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *NetworkMapResponse) GetPeers() []*Node {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *NetworkMapResponse) GetPeersChanged() []*Node {
	if x != nil {
		return x.PeersChanged
	}
	return nil
}

func (x *NetworkMapResponse) GetPeersRemoved() []uint64 {
	if x != nil {
		return x.PeersRemoved
	}
	return nil
}

func (x *NetworkMapResponse) GetPacketFilter() []*FilterRule {
	if x != nil {
		return x.PacketFilter
	}
	return nil
}

func (x *NetworkMapResponse) GetAdvertisedRoute() string {
	if x != nil {
		return x.AdvertisedRoute
	}
	return ""
}

func (x *NetworkMapResponse) GetJailed() bool {
	if x != nil {
		return x.Jailed
	}
	return false
}

func (x *NetworkMapResponse) GetIceTable() []*Node {
	if x != nil {
		return x.IceTable
	}
	return nil
}

type NetPortRangePortRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First uint64 `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	Last  uint64 `protobuf:"varint,2,opt,name=last,proto3" json:"last,omitempty"`
}

func (x *NetPortRangePortRange) Reset() {
	*x = NetPortRangePortRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetPortRangePortRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetPortRangePortRange) ProtoMessage() {}

func (x *NetPortRangePortRange) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetPortRangePortRange.ProtoReflect.Descriptor instead.
func (*NetPortRangePortRange) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_node_proto_rawDescGZIP(), []int{3, 0}
}

func (x *NetPortRangePortRange) GetFirst() uint64 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *NetPortRangePortRange) GetLast() uint64 {
	if x != nil {
		return x.Last
	}
	return 0
}

var File_runetale_runetale_v1_node_proto protoreflect.FileDescriptor

var file_runetale_runetale_v1_node_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x75, 0x6e, 0x65, 0x74, 0x61, 0x6c, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x65, 0x74,
	0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x64, 0x72, 0x22, 0xac, 0x01, 0x0a, 0x04, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x67, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x67, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x49, 0x50, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x49, 0x50, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x64, 0x72, 0x22, 0x39, 0x0a, 0x13, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x64, 0x72, 0x22, 0xb7, 0x01, 0x0a, 0x0c, 0x4e, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x34, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x61,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65,
	0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x35, 0x0a, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x22, 0x68,
	0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x72, 0x63, 0x49, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x72,
	0x63, 0x49, 0x70, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x04, 0x64, 0x73, 0x74, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x50, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x07, 0x69, 0x50, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x02, 0x0a, 0x12, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x65,
	0x71, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x0c, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0c, 0x70, 0x65, 0x65,
	0x72, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x65, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x0c, 0x70, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x36, 0x0a,
	0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x4a, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x4a, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x08, 0x69, 0x63, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x69, 0x63, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x32, 0xf2, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x44, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x70, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56,
	0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x4d, 0x61, 0x70, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runetale_runetale_v1_node_proto_rawDescOnce sync.Once
	file_runetale_runetale_v1_node_proto_rawDescData = file_runetale_runetale_v1_node_proto_rawDesc
)

func file_runetale_runetale_v1_node_proto_rawDescGZIP() []byte {
	file_runetale_runetale_v1_node_proto_rawDescOnce.Do(func() {
		file_runetale_runetale_v1_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_runetale_runetale_v1_node_proto_rawDescData)
	})
	return file_runetale_runetale_v1_node_proto_rawDescData
}

var file_runetale_runetale_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_runetale_runetale_v1_node_proto_goTypes = []interface{}{
	(*SyncNodesResponse)(nil),     // 0: protos.SyncNodesResponse
	(*Node)(nil),                  // 1: protos.Node
	(*ComposeNodeResponse)(nil),   // 2: protos.ComposeNodeResponse
	(*NetPortRange)(nil),          // 3: protos.NetPortRange
	(*FilterRule)(nil),            // 4: protos.FilterRule
	(*NetworkMapResponse)(nil),    // 5: protos.NetworkMapResponse
	(*NetPortRangePortRange)(nil), // 6: protos.NetPortRange.portRange
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_runetale_runetale_v1_node_proto_depIdxs = []int32{
	1,  // 0: protos.SyncNodesResponse.remoteNodes:type_name -> protos.Node
	6,  // 1: protos.NetPortRange.ports:type_name -> protos.NetPortRange.portRange
	3,  // 2: protos.FilterRule.dsts:type_name -> protos.NetPortRange
	1,  // 3: protos.NetworkMapResponse.node:type_name -> protos.Node
	1,  // 4: protos.NetworkMapResponse.peers:type_name -> protos.Node
	1,  // 5: protos.NetworkMapResponse.peersChanged:type_name -> protos.Node
	4,  // 6: protos.NetworkMapResponse.packetFilter:type_name -> protos.FilterRule
	1,  // 7: protos.NetworkMapResponse.iceTable:type_name -> protos.Node
	7,  // 8: protos.NodeService.ComposeNode:input_type -> google.protobuf.Empty
	7,  // 9: protos.NodeService.GetNetworkMap:input_type -> google.protobuf.Empty
	5,  // 10: protos.NodeService.ConnectNetworkMapTable:input_type -> protos.NetworkMapResponse
	2,  // 11: protos.NodeService.ComposeNode:output_type -> protos.ComposeNodeResponse
	5,  // 12: protos.NodeService.GetNetworkMap:output_type -> protos.NetworkMapResponse
	5,  // 13: protos.NodeService.ConnectNetworkMapTable:output_type -> protos.NetworkMapResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_runetale_runetale_v1_node_proto_init() }
func file_runetale_runetale_v1_node_proto_init() {
	if File_runetale_runetale_v1_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runetale_runetale_v1_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncNodesResponse); i {
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
		file_runetale_runetale_v1_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_runetale_runetale_v1_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeNodeResponse); i {
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
		file_runetale_runetale_v1_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetPortRange); i {
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
		file_runetale_runetale_v1_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterRule); i {
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
		file_runetale_runetale_v1_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkMapResponse); i {
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
		file_runetale_runetale_v1_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetPortRangePortRange); i {
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
			RawDescriptor: file_runetale_runetale_v1_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runetale_runetale_v1_node_proto_goTypes,
		DependencyIndexes: file_runetale_runetale_v1_node_proto_depIdxs,
		MessageInfos:      file_runetale_runetale_v1_node_proto_msgTypes,
	}.Build()
	File_runetale_runetale_v1_node_proto = out.File
	file_runetale_runetale_v1_node_proto_rawDesc = nil
	file_runetale_runetale_v1_node_proto_goTypes = nil
	file_runetale_runetale_v1_node_proto_depIdxs = nil
}
