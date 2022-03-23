// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: protos/admin_network.proto

package admin_network

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

type CreateDefaultAdminNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyName string `protobuf:"bytes,1,opt,name=companyName,proto3" json:"companyName,omitempty"`
	UserID      string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateDefaultAdminNetworkRequest) Reset() {
	*x = CreateDefaultAdminNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_admin_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDefaultAdminNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDefaultAdminNetworkRequest) ProtoMessage() {}

func (x *CreateDefaultAdminNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_admin_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDefaultAdminNetworkRequest.ProtoReflect.Descriptor instead.
func (*CreateDefaultAdminNetworkRequest) Descriptor() ([]byte, []int) {
	return file_protos_admin_network_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDefaultAdminNetworkRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *CreateDefaultAdminNetworkRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateDefaultAdminNetworkRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateDefaultAdminNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationID string `protobuf:"bytes,1,opt,name=organizationID,proto3" json:"organizationID,omitempty"`
}

func (x *CreateDefaultAdminNetworkResponse) Reset() {
	*x = CreateDefaultAdminNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_admin_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDefaultAdminNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDefaultAdminNetworkResponse) ProtoMessage() {}

func (x *CreateDefaultAdminNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_admin_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDefaultAdminNetworkResponse.ProtoReflect.Descriptor instead.
func (*CreateDefaultAdminNetworkResponse) Descriptor() ([]byte, []int) {
	return file_protos_admin_network_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDefaultAdminNetworkResponse) GetOrganizationID() string {
	if x != nil {
		return x.OrganizationID
	}
	return ""
}

type LoginNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkName string `protobuf:"bytes,1,opt,name=networkName,proto3" json:"networkName,omitempty"`
	UserID      string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *LoginNetworkRequest) Reset() {
	*x = LoginNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_admin_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginNetworkRequest) ProtoMessage() {}

func (x *LoginNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_admin_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginNetworkRequest.ProtoReflect.Descriptor instead.
func (*LoginNetworkRequest) Descriptor() ([]byte, []int) {
	return file_protos_admin_network_proto_rawDescGZIP(), []int{2}
}

func (x *LoginNetworkRequest) GetNetworkName() string {
	if x != nil {
		return x.NetworkName
	}
	return ""
}

func (x *LoginNetworkRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type LoginNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationID string `protobuf:"bytes,1,opt,name=organizationID,proto3" json:"organizationID,omitempty"`
}

func (x *LoginNetworkResponse) Reset() {
	*x = LoginNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_admin_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginNetworkResponse) ProtoMessage() {}

func (x *LoginNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_admin_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginNetworkResponse.ProtoReflect.Descriptor instead.
func (*LoginNetworkResponse) Descriptor() ([]byte, []int) {
	return file_protos_admin_network_proto_rawDescGZIP(), []int{3}
}

func (x *LoginNetworkResponse) GetOrganizationID() string {
	if x != nil {
		return x.OrganizationID
	}
	return ""
}

var File_protos_admin_network_proto protoreflect.FileDescriptor

var file_protos_admin_network_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x72, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4b, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x4f, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3e, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x32, 0xd1, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_admin_network_proto_rawDescOnce sync.Once
	file_protos_admin_network_proto_rawDescData = file_protos_admin_network_proto_rawDesc
)

func file_protos_admin_network_proto_rawDescGZIP() []byte {
	file_protos_admin_network_proto_rawDescOnce.Do(func() {
		file_protos_admin_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_admin_network_proto_rawDescData)
	})
	return file_protos_admin_network_proto_rawDescData
}

var file_protos_admin_network_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_admin_network_proto_goTypes = []interface{}{
	(*CreateDefaultAdminNetworkRequest)(nil),  // 0: protos.CreateDefaultAdminNetworkRequest
	(*CreateDefaultAdminNetworkResponse)(nil), // 1: protos.CreateDefaultAdminNetworkResponse
	(*LoginNetworkRequest)(nil),               // 2: protos.LoginNetworkRequest
	(*LoginNetworkResponse)(nil),              // 3: protos.LoginNetworkResponse
}
var file_protos_admin_network_proto_depIdxs = []int32{
	0, // 0: protos.AdminNetworkService.CreateDefaultNetwork:input_type -> protos.CreateDefaultAdminNetworkRequest
	2, // 1: protos.AdminNetworkService.LoginNetwork:input_type -> protos.LoginNetworkRequest
	1, // 2: protos.AdminNetworkService.CreateDefaultNetwork:output_type -> protos.CreateDefaultAdminNetworkResponse
	3, // 3: protos.AdminNetworkService.LoginNetwork:output_type -> protos.LoginNetworkResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_admin_network_proto_init() }
func file_protos_admin_network_proto_init() {
	if File_protos_admin_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_admin_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDefaultAdminNetworkRequest); i {
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
		file_protos_admin_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDefaultAdminNetworkResponse); i {
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
		file_protos_admin_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginNetworkRequest); i {
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
		file_protos_admin_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginNetworkResponse); i {
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
			RawDescriptor: file_protos_admin_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_admin_network_proto_goTypes,
		DependencyIndexes: file_protos_admin_network_proto_depIdxs,
		MessageInfos:      file_protos_admin_network_proto_msgTypes,
	}.Build()
	File_protos_admin_network_proto = out.File
	file_protos_admin_network_proto_rawDesc = nil
	file_protos_admin_network_proto_goTypes = nil
	file_protos_admin_network_proto_depIdxs = nil
}
