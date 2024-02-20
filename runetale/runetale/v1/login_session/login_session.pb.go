// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.3
// source: runetale/runetale/v1/login_session.proto

package login_session

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

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsRegistered bool   `protobuf:"varint,1,opt,name=isRegistered,proto3" json:"isRegistered,omitempty"`
	LoginUrl     string `protobuf:"bytes,2,opt,name=loginUrl,proto3" json:"loginUrl,omitempty"`
	Ip           string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Cidr         string `protobuf:"bytes,4,opt,name=cidr,proto3" json:"cidr,omitempty"`
	SignalHost   string `protobuf:"bytes,5,opt,name=signalHost,proto3" json:"signalHost,omitempty"`
	SignalPort   uint64 `protobuf:"varint,6,opt,name=signalPort,proto3" json:"signalPort,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_login_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_login_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_login_session_proto_rawDescGZIP(), []int{0}
}

func (x *JoinResponse) GetIsRegistered() bool {
	if x != nil {
		return x.IsRegistered
	}
	return false
}

func (x *JoinResponse) GetLoginUrl() string {
	if x != nil {
		return x.LoginUrl
	}
	return ""
}

func (x *JoinResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *JoinResponse) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

func (x *JoinResponse) GetSignalHost() string {
	if x != nil {
		return x.SignalHost
	}
	return ""
}

func (x *JoinResponse) GetSignalPort() uint64 {
	if x != nil {
		return x.SignalPort
	}
	return 0
}

type PeerLoginSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// host ip
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// host wireguard cidr
	Cidr string `protobuf:"bytes,2,opt,name=cidr,proto3" json:"cidr,omitempty"`
	// host name
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	// host os
	Os               string `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	SignalServerHost string `protobuf:"bytes,5,opt,name=signalServerHost,proto3" json:"signalServerHost,omitempty"`
	SignalServerPort uint64 `protobuf:"varint,6,opt,name=signalServerPort,proto3" json:"signalServerPort,omitempty"`
}

func (x *PeerLoginSessionResponse) Reset() {
	*x = PeerLoginSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runetale_runetale_v1_login_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerLoginSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerLoginSessionResponse) ProtoMessage() {}

func (x *PeerLoginSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runetale_runetale_v1_login_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerLoginSessionResponse.ProtoReflect.Descriptor instead.
func (*PeerLoginSessionResponse) Descriptor() ([]byte, []int) {
	return file_runetale_runetale_v1_login_session_proto_rawDescGZIP(), []int{1}
}

func (x *PeerLoginSessionResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *PeerLoginSessionResponse) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

func (x *PeerLoginSessionResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PeerLoginSessionResponse) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *PeerLoginSessionResponse) GetSignalServerHost() string {
	if x != nil {
		return x.SignalServerHost
	}
	return ""
}

func (x *PeerLoginSessionResponse) GetSignalServerPort() uint64 {
	if x != nil {
		return x.SignalServerPort
	}
	return 0
}

var File_runetale_runetale_v1_login_session_proto protoreflect.FileDescriptor

var file_runetale_runetale_v1_login_session_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x75, 0x6e, 0x65, 0x74, 0x61, 0x6c, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x65, 0x74,
	0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb2, 0x01, 0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x64, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x48, 0x6f,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x50, 0x6f,
	0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x50, 0x6f, 0x72, 0x74, 0x22, 0xba, 0x01, 0x0a, 0x18, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72,
	0x74, 0x32, 0xa7, 0x01, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x4a, 0x6f, 0x69,
	0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x58, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x65, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runetale_runetale_v1_login_session_proto_rawDescOnce sync.Once
	file_runetale_runetale_v1_login_session_proto_rawDescData = file_runetale_runetale_v1_login_session_proto_rawDesc
)

func file_runetale_runetale_v1_login_session_proto_rawDescGZIP() []byte {
	file_runetale_runetale_v1_login_session_proto_rawDescOnce.Do(func() {
		file_runetale_runetale_v1_login_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_runetale_runetale_v1_login_session_proto_rawDescData)
	})
	return file_runetale_runetale_v1_login_session_proto_rawDescData
}

var file_runetale_runetale_v1_login_session_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_runetale_runetale_v1_login_session_proto_goTypes = []interface{}{
	(*JoinResponse)(nil),             // 0: protos.JoinResponse
	(*PeerLoginSessionResponse)(nil), // 1: protos.PeerLoginSessionResponse
	(*emptypb.Empty)(nil),            // 2: google.protobuf.Empty
}
var file_runetale_runetale_v1_login_session_proto_depIdxs = []int32{
	2, // 0: protos.LoginSessionService.Join:input_type -> google.protobuf.Empty
	2, // 1: protos.LoginSessionService.StreamPeerLoginSession:input_type -> google.protobuf.Empty
	0, // 2: protos.LoginSessionService.Join:output_type -> protos.JoinResponse
	1, // 3: protos.LoginSessionService.StreamPeerLoginSession:output_type -> protos.PeerLoginSessionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_runetale_runetale_v1_login_session_proto_init() }
func file_runetale_runetale_v1_login_session_proto_init() {
	if File_runetale_runetale_v1_login_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runetale_runetale_v1_login_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_runetale_runetale_v1_login_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerLoginSessionResponse); i {
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
			RawDescriptor: file_runetale_runetale_v1_login_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runetale_runetale_v1_login_session_proto_goTypes,
		DependencyIndexes: file_runetale_runetale_v1_login_session_proto_depIdxs,
		MessageInfos:      file_runetale_runetale_v1_login_session_proto_msgTypes,
	}.Build()
	File_runetale_runetale_v1_login_session_proto = out.File
	file_runetale_runetale_v1_login_session_proto_rawDesc = nil
	file_runetale_runetale_v1_login_session_proto_goTypes = nil
	file_runetale_runetale_v1_login_session_proto_depIdxs = nil
}
