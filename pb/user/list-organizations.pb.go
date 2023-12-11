// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: user/list-organizations.proto

package user

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type ListOrganizationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListOrganizationsRequest) Reset() {
	*x = ListOrganizationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_list_organizations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrganizationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrganizationsRequest) ProtoMessage() {}

func (x *ListOrganizationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_list_organizations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrganizationsRequest.ProtoReflect.Descriptor instead.
func (*ListOrganizationsRequest) Descriptor() ([]byte, []int) {
	return file_user_list_organizations_proto_rawDescGZIP(), []int{0}
}

type ListOrganizationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// organizations is a list of all organizations this user has access to. The
	// user can subsequently call DescribeOrganization to get more info about an
	// individual oeganization.
	Organizations []*ListOrganizationsEntry `protobuf:"bytes,1,rep,name=organizations,proto3" json:"organizations,omitempty"`
}

func (x *ListOrganizationsResponse) Reset() {
	*x = ListOrganizationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_list_organizations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrganizationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrganizationsResponse) ProtoMessage() {}

func (x *ListOrganizationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_list_organizations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrganizationsResponse.ProtoReflect.Descriptor instead.
func (*ListOrganizationsResponse) Descriptor() ([]byte, []int) {
	return file_user_list_organizations_proto_rawDescGZIP(), []int{1}
}

func (x *ListOrganizationsResponse) GetOrganizations() []*ListOrganizationsEntry {
	if x != nil {
		return x.Organizations
	}
	return nil
}

// ListOrganizationsEntry provides a sample of information about an organization.
type ListOrganizationsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// organization_id is a globally unique identifier.
	OrganizationId uint64 `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	// name is a required human readable name that will be displayed in the UI.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListOrganizationsEntry) Reset() {
	*x = ListOrganizationsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_list_organizations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrganizationsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrganizationsEntry) ProtoMessage() {}

func (x *ListOrganizationsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_user_list_organizations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrganizationsEntry.ProtoReflect.Descriptor instead.
func (*ListOrganizationsEntry) Descriptor() ([]byte, []int) {
	return file_user_list_organizations_proto_rawDescGZIP(), []int{2}
}

func (x *ListOrganizationsEntry) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *ListOrganizationsEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_user_list_organizations_proto protoreflect.FileDescriptor

var file_user_list_organizations_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x62, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45,
	0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6b, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x2f, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xba, 0x48, 0x09, 0xd0, 0x01, 0x01, 0x72, 0x04, 0x20, 0x01, 0x28, 0x40, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_list_organizations_proto_rawDescOnce sync.Once
	file_user_list_organizations_proto_rawDescData = file_user_list_organizations_proto_rawDesc
)

func file_user_list_organizations_proto_rawDescGZIP() []byte {
	file_user_list_organizations_proto_rawDescOnce.Do(func() {
		file_user_list_organizations_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_list_organizations_proto_rawDescData)
	})
	return file_user_list_organizations_proto_rawDescData
}

var file_user_list_organizations_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_list_organizations_proto_goTypes = []interface{}{
	(*ListOrganizationsRequest)(nil),  // 0: stately.ListOrganizationsRequest
	(*ListOrganizationsResponse)(nil), // 1: stately.ListOrganizationsResponse
	(*ListOrganizationsEntry)(nil),    // 2: stately.ListOrganizationsEntry
}
var file_user_list_organizations_proto_depIdxs = []int32{
	2, // 0: stately.ListOrganizationsResponse.organizations:type_name -> stately.ListOrganizationsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_list_organizations_proto_init() }
func file_user_list_organizations_proto_init() {
	if File_user_list_organizations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_list_organizations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrganizationsRequest); i {
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
		file_user_list_organizations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrganizationsResponse); i {
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
		file_user_list_organizations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrganizationsEntry); i {
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
			RawDescriptor: file_user_list_organizations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_list_organizations_proto_goTypes,
		DependencyIndexes: file_user_list_organizations_proto_depIdxs,
		MessageInfos:      file_user_list_organizations_proto_msgTypes,
	}.Build()
	File_user_list_organizations_proto = out.File
	file_user_list_organizations_proto_rawDesc = nil
	file_user_list_organizations_proto_goTypes = nil
	file_user_list_organizations_proto_depIdxs = nil
}
