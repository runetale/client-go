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

	RemoteClientNodeKey string   `protobuf:"bytes,1,opt,name=remoteClientNodeKey,proto3" json:"remoteClientNodeKey,omitempty"`
	RemoteWgPubKey      string   `protobuf:"bytes,2,opt,name=remoteWgPubKey,proto3" json:"remoteWgPubKey,omitempty"`
	AllowedIPs          []string `protobuf:"bytes,3,rep,name=allowedIPs,proto3" json:"allowedIPs,omitempty"`
	Ip                  string   `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Cidr                string   `protobuf:"bytes,5,opt,name=cidr,proto3" json:"cidr,omitempty"`
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

func (x *Node) GetRemoteClientNodeKey() string {
	if x != nil {
		return x.RemoteClientNodeKey
	}
	return ""
}

func (x *Node) GetRemoteWgPubKey() string {
	if x != nil {
		return x.RemoteWgPubKey
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
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x64, 0x72, 0x22, 0xa4, 0x01, 0x0a, 0x04, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x57,
	0x67, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x57, 0x67, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x50, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x50, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x64,
	0x72, 0x22, 0x39, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x64, 0x72, 0x32, 0xa1, 0x01, 0x0a,
	0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x15,
	0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_runetale_runetale_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_runetale_runetale_v1_node_proto_goTypes = []interface{}{
	(*SyncNodesResponse)(nil),   // 0: protos.SyncNodesResponse
	(*Node)(nil),                // 1: protos.Node
	(*ComposeNodeResponse)(nil), // 2: protos.ComposeNodeResponse
	(*emptypb.Empty)(nil),       // 3: google.protobuf.Empty
}
var file_runetale_runetale_v1_node_proto_depIdxs = []int32{
	1, // 0: protos.SyncNodesResponse.remoteNodes:type_name -> protos.Node
	3, // 1: protos.NodeService.SyncRemoteNodesConfig:input_type -> google.protobuf.Empty
	3, // 2: protos.NodeService.ComposeNode:input_type -> google.protobuf.Empty
	0, // 3: protos.NodeService.SyncRemoteNodesConfig:output_type -> protos.SyncNodesResponse
	2, // 4: protos.NodeService.ComposeNode:output_type -> protos.ComposeNodeResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_runetale_runetale_v1_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
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
