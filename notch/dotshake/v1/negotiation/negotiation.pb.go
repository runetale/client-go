// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.1
// source: notch/dotshake/v1/negotiation.proto

package negotiation

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

type NegotiationType int32

const (
	NegotiationType_OFFER     NegotiationType = 0
	NegotiationType_ANSWER    NegotiationType = 1
	NegotiationType_CANDIDATE NegotiationType = 2
)

// Enum value maps for NegotiationType.
var (
	NegotiationType_name = map[int32]string{
		0: "OFFER",
		1: "ANSWER",
		2: "CANDIDATE",
	}
	NegotiationType_value = map[string]int32{
		"OFFER":     0,
		"ANSWER":    1,
		"CANDIDATE": 2,
	}
)

func (x NegotiationType) Enum() *NegotiationType {
	p := new(NegotiationType)
	*p = x
	return p
}

func (x NegotiationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NegotiationType) Descriptor() protoreflect.EnumDescriptor {
	return file_notch_dotshake_v1_negotiation_proto_enumTypes[0].Descriptor()
}

func (NegotiationType) Type() protoreflect.EnumType {
	return &file_notch_dotshake_v1_negotiation_proto_enumTypes[0]
}

func (x NegotiationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NegotiationType.Descriptor instead.
func (NegotiationType) EnumDescriptor() ([]byte, []int) {
	return file_notch_dotshake_v1_negotiation_proto_rawDescGZIP(), []int{0}
}

type NegotiationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                 NegotiationType `protobuf:"varint,1,opt,name=type,proto3,enum=protos.NegotiationType" json:"type,omitempty"`
	RemotePeerMachineKey string          `protobuf:"bytes,2,opt,name=remotePeerMachineKey,proto3" json:"remotePeerMachineKey,omitempty"`
}

func (x *NegotiationRequest) Reset() {
	*x = NegotiationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notch_dotshake_v1_negotiation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NegotiationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NegotiationRequest) ProtoMessage() {}

func (x *NegotiationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notch_dotshake_v1_negotiation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NegotiationRequest.ProtoReflect.Descriptor instead.
func (*NegotiationRequest) Descriptor() ([]byte, []int) {
	return file_notch_dotshake_v1_negotiation_proto_rawDescGZIP(), []int{0}
}

func (x *NegotiationRequest) GetType() NegotiationType {
	if x != nil {
		return x.Type
	}
	return NegotiationType_OFFER
}

func (x *NegotiationRequest) GetRemotePeerMachineKey() string {
	if x != nil {
		return x.RemotePeerMachineKey
	}
	return ""
}

type NegotiationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          NegotiationType `protobuf:"varint,1,opt,name=type,proto3,enum=protos.NegotiationType" json:"type,omitempty"`
	IsNegotiation bool            `protobuf:"varint,2,opt,name=isNegotiation,proto3" json:"isNegotiation,omitempty"`
}

func (x *NegotiationResponse) Reset() {
	*x = NegotiationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notch_dotshake_v1_negotiation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NegotiationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NegotiationResponse) ProtoMessage() {}

func (x *NegotiationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notch_dotshake_v1_negotiation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NegotiationResponse.ProtoReflect.Descriptor instead.
func (*NegotiationResponse) Descriptor() ([]byte, []int) {
	return file_notch_dotshake_v1_negotiation_proto_rawDescGZIP(), []int{1}
}

func (x *NegotiationResponse) GetType() NegotiationType {
	if x != nil {
		return x.Type
	}
	return NegotiationType_OFFER
}

func (x *NegotiationResponse) GetIsNegotiation() bool {
	if x != nil {
		return x.IsNegotiation
	}
	return false
}

var File_notch_dotshake_v1_negotiation_proto protoreflect.FileDescriptor

var file_notch_dotshake_v1_negotiation_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x6f, 0x74, 0x63, 0x68, 0x2f, 0x64, 0x6f, 0x74, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x12, 0x4e, 0x65,
	0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a,
	0x14, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4b, 0x65,
	0x79, 0x22, 0x68, 0x0a, 0x13, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x73, 0x4e, 0x65, 0x67, 0x6f, 0x74,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73,
	0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x37, 0x0a, 0x0f, 0x4e,
	0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x4f, 0x46, 0x46, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4e, 0x53,
	0x57, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x02, 0x32, 0x9e, 0x02, 0x0a, 0x0b, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x67, 0x6f,
	0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x6e, 0x65, 0x67, 0x6f, 0x74,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notch_dotshake_v1_negotiation_proto_rawDescOnce sync.Once
	file_notch_dotshake_v1_negotiation_proto_rawDescData = file_notch_dotshake_v1_negotiation_proto_rawDesc
)

func file_notch_dotshake_v1_negotiation_proto_rawDescGZIP() []byte {
	file_notch_dotshake_v1_negotiation_proto_rawDescOnce.Do(func() {
		file_notch_dotshake_v1_negotiation_proto_rawDescData = protoimpl.X.CompressGZIP(file_notch_dotshake_v1_negotiation_proto_rawDescData)
	})
	return file_notch_dotshake_v1_negotiation_proto_rawDescData
}

var file_notch_dotshake_v1_negotiation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notch_dotshake_v1_negotiation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_notch_dotshake_v1_negotiation_proto_goTypes = []interface{}{
	(NegotiationType)(0),        // 0: protos.NegotiationType
	(*NegotiationRequest)(nil),  // 1: protos.NegotiationRequest
	(*NegotiationResponse)(nil), // 2: protos.NegotiationResponse
	(*emptypb.Empty)(nil),       // 3: google.protobuf.Empty
}
var file_notch_dotshake_v1_negotiation_proto_depIdxs = []int32{
	0, // 0: protos.NegotiationRequest.type:type_name -> protos.NegotiationType
	0, // 1: protos.NegotiationResponse.type:type_name -> protos.NegotiationType
	1, // 2: protos.Negotiation.Offer:input_type -> protos.NegotiationRequest
	1, // 3: protos.Negotiation.Answer:input_type -> protos.NegotiationRequest
	1, // 4: protos.Negotiation.Candidate:input_type -> protos.NegotiationRequest
	1, // 5: protos.Negotiation.StartConnect:input_type -> protos.NegotiationRequest
	3, // 6: protos.Negotiation.Offer:output_type -> google.protobuf.Empty
	3, // 7: protos.Negotiation.Answer:output_type -> google.protobuf.Empty
	3, // 8: protos.Negotiation.Candidate:output_type -> google.protobuf.Empty
	2, // 9: protos.Negotiation.StartConnect:output_type -> protos.NegotiationResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notch_dotshake_v1_negotiation_proto_init() }
func file_notch_dotshake_v1_negotiation_proto_init() {
	if File_notch_dotshake_v1_negotiation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notch_dotshake_v1_negotiation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NegotiationRequest); i {
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
		file_notch_dotshake_v1_negotiation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NegotiationResponse); i {
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
			RawDescriptor: file_notch_dotshake_v1_negotiation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notch_dotshake_v1_negotiation_proto_goTypes,
		DependencyIndexes: file_notch_dotshake_v1_negotiation_proto_depIdxs,
		EnumInfos:         file_notch_dotshake_v1_negotiation_proto_enumTypes,
		MessageInfos:      file_notch_dotshake_v1_negotiation_proto_msgTypes,
	}.Build()
	File_notch_dotshake_v1_negotiation_proto = out.File
	file_notch_dotshake_v1_negotiation_proto_rawDesc = nil
	file_notch_dotshake_v1_negotiation_proto_goTypes = nil
	file_notch_dotshake_v1_negotiation_proto_depIdxs = nil
}
