// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: protos/negotiation.proto

package negotiation

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

type Body_Type int32

const (
	Body_OFFER     Body_Type = 0
	Body_ANSWER    Body_Type = 1
	Body_CANDIDATE Body_Type = 2
)

// Enum value maps for Body_Type.
var (
	Body_Type_name = map[int32]string{
		0: "OFFER",
		1: "ANSWER",
		2: "CANDIDATE",
	}
	Body_Type_value = map[string]int32{
		"OFFER":     0,
		"ANSWER":    1,
		"CANDIDATE": 2,
	}
)

func (x Body_Type) Enum() *Body_Type {
	p := new(Body_Type)
	*p = x
	return p
}

func (x Body_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Body_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_negotiation_proto_enumTypes[0].Descriptor()
}

func (Body_Type) Type() protoreflect.EnumType {
	return &file_protos_negotiation_proto_enumTypes[0]
}

func (x Body_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Body_Type.Descriptor instead.
func (Body_Type) EnumDescriptor() ([]byte, []int) {
	return file_protos_negotiation_proto_rawDescGZIP(), []int{0, 0}
}

type Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKey       string    `protobuf:"bytes,1,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	ClientMachineKey string    `protobuf:"bytes,2,opt,name=clientMachineKey,proto3" json:"clientMachineKey,omitempty"`
	Type             Body_Type `protobuf:"varint,3,opt,name=type,proto3,enum=protos.Body_Type" json:"type,omitempty"`
	Payload          string    `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	Key              string    `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Remotekey        string    `protobuf:"bytes,6,opt,name=remotekey,proto3" json:"remotekey,omitempty"`
	UFlag            string    `protobuf:"bytes,7,opt,name=uFlag,proto3" json:"uFlag,omitempty"`
	Pwd              string    `protobuf:"bytes,8,opt,name=pwd,proto3" json:"pwd,omitempty"`
}

func (x *Body) Reset() {
	*x = Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_negotiation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Body) ProtoMessage() {}

func (x *Body) ProtoReflect() protoreflect.Message {
	mi := &file_protos_negotiation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Body.ProtoReflect.Descriptor instead.
func (*Body) Descriptor() ([]byte, []int) {
	return file_protos_negotiation_proto_rawDescGZIP(), []int{0}
}

func (x *Body) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *Body) GetClientMachineKey() string {
	if x != nil {
		return x.ClientMachineKey
	}
	return ""
}

func (x *Body) GetType() Body_Type {
	if x != nil {
		return x.Type
	}
	return Body_OFFER
}

func (x *Body) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *Body) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Body) GetRemotekey() string {
	if x != nil {
		return x.Remotekey
	}
	return ""
}

func (x *Body) GetUFlag() string {
	if x != nil {
		return x.UFlag
	}
	return ""
}

func (x *Body) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

var File_protos_negotiation_proto protoreflect.FileDescriptor

var file_protos_negotiation_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x22, 0x99, 0x02, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42,
	0x6f, 0x64, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x46, 0x6c, 0x61,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x77, 0x64,
	0x22, 0x2c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x46, 0x46, 0x45,
	0x52, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x32, 0x66,
	0x0a, 0x0b, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a,
	0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42,
	0x6f, 0x64, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x64,
	0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f,
	0x64, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x64, 0x79,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x6e, 0x65, 0x67, 0x6f,
	0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_negotiation_proto_rawDescOnce sync.Once
	file_protos_negotiation_proto_rawDescData = file_protos_negotiation_proto_rawDesc
)

func file_protos_negotiation_proto_rawDescGZIP() []byte {
	file_protos_negotiation_proto_rawDescOnce.Do(func() {
		file_protos_negotiation_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_negotiation_proto_rawDescData)
	})
	return file_protos_negotiation_proto_rawDescData
}

var file_protos_negotiation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_negotiation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_negotiation_proto_goTypes = []interface{}{
	(Body_Type)(0), // 0: protos.Body.Type
	(*Body)(nil),   // 1: protos.Body
}
var file_protos_negotiation_proto_depIdxs = []int32{
	0, // 0: protos.Body.type:type_name -> protos.Body.Type
	1, // 1: protos.Negotiation.Send:input_type -> protos.Body
	1, // 2: protos.Negotiation.ConnectStream:input_type -> protos.Body
	1, // 3: protos.Negotiation.Send:output_type -> protos.Body
	1, // 4: protos.Negotiation.ConnectStream:output_type -> protos.Body
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_negotiation_proto_init() }
func file_protos_negotiation_proto_init() {
	if File_protos_negotiation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_negotiation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Body); i {
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
			RawDescriptor: file_protos_negotiation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_negotiation_proto_goTypes,
		DependencyIndexes: file_protos_negotiation_proto_depIdxs,
		EnumInfos:         file_protos_negotiation_proto_enumTypes,
		MessageInfos:      file_protos_negotiation_proto_msgTypes,
	}.Build()
	File_protos_negotiation_proto = out.File
	file_protos_negotiation_proto_rawDesc = nil
	file_protos_negotiation_proto_goTypes = nil
	file_protos_negotiation_proto_depIdxs = nil
}
