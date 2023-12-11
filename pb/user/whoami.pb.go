// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: user/whoami.proto

package user

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	dbmanagement "github.com/StatelyCloud/go-sdk/pb/dbmanagement"
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

type WhoamiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WhoamiRequest) Reset() {
	*x = WhoamiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_whoami_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoamiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoamiRequest) ProtoMessage() {}

func (x *WhoamiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_whoami_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoamiRequest.ProtoReflect.Descriptor instead.
func (*WhoamiRequest) Descriptor() ([]byte, []int) {
	return file_user_whoami_proto_rawDescGZIP(), []int{0}
}

type WhoamiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// oauth_subject is the user's ID in Auth0. We already have this from the auth token
	// but it can't hurt to return it so we can validate that what we have locally
	// is still what we thought.
	OauthSubject string `protobuf:"bytes,1,opt,name=oauth_subject,json=oauthSubject,proto3" json:"oauth_subject,omitempty"`
	// user_id is the user's ID in Stately. This is generated by us during enrollment
	// and can may also be referred to as member ID within Stately.
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// enrollment_time is the UTC epoch of when the User was created.
	EnrollmentTime uint64 `protobuf:"varint,4,opt,name=enrollment_time,json=enrollmentTime,proto3" json:"enrollment_time,omitempty"`
	// organizations is a tree of the organizations, their projects and their stores.
	Organizations []*OrganizationNode `protobuf:"bytes,6,rep,name=organizations,proto3" json:"organizations,omitempty"`
}

func (x *WhoamiResponse) Reset() {
	*x = WhoamiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_whoami_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoamiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoamiResponse) ProtoMessage() {}

func (x *WhoamiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_whoami_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoamiResponse.ProtoReflect.Descriptor instead.
func (*WhoamiResponse) Descriptor() ([]byte, []int) {
	return file_user_whoami_proto_rawDescGZIP(), []int{1}
}

func (x *WhoamiResponse) GetOauthSubject() string {
	if x != nil {
		return x.OauthSubject
	}
	return ""
}

func (x *WhoamiResponse) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *WhoamiResponse) GetEnrollmentTime() uint64 {
	if x != nil {
		return x.EnrollmentTime
	}
	return 0
}

func (x *WhoamiResponse) GetOrganizations() []*OrganizationNode {
	if x != nil {
		return x.Organizations
	}
	return nil
}

type OrganizationNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// organization contains details about an organization.
	Organization *Organization `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	// projects is a list of projects that belong to this organization.
	Projects []*ProjectNode `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *OrganizationNode) Reset() {
	*x = OrganizationNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_whoami_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationNode) ProtoMessage() {}

func (x *OrganizationNode) ProtoReflect() protoreflect.Message {
	mi := &file_user_whoami_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationNode.ProtoReflect.Descriptor instead.
func (*OrganizationNode) Descriptor() ([]byte, []int) {
	return file_user_whoami_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationNode) GetOrganization() *Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

func (x *OrganizationNode) GetProjects() []*ProjectNode {
	if x != nil {
		return x.Projects
	}
	return nil
}

type ProjectNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project contains details about a project.
	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// stores is a list of stores that belong to this project.
	Stores []*StoreNode `protobuf:"bytes,2,rep,name=stores,proto3" json:"stores,omitempty"`
}

func (x *ProjectNode) Reset() {
	*x = ProjectNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_whoami_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectNode) ProtoMessage() {}

func (x *ProjectNode) ProtoReflect() protoreflect.Message {
	mi := &file_user_whoami_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectNode.ProtoReflect.Descriptor instead.
func (*ProjectNode) Descriptor() ([]byte, []int) {
	return file_user_whoami_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectNode) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *ProjectNode) GetStores() []*StoreNode {
	if x != nil {
		return x.Stores
	}
	return nil
}

type StoreNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// store contains details about a store.
	Store *dbmanagement.StoreInfo `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
}

func (x *StoreNode) Reset() {
	*x = StoreNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_whoami_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreNode) ProtoMessage() {}

func (x *StoreNode) ProtoReflect() protoreflect.Message {
	mi := &file_user_whoami_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreNode.ProtoReflect.Descriptor instead.
func (*StoreNode) Descriptor() ([]byte, []int) {
	return file_user_whoami_proto_rawDescGZIP(), []int{4}
}

func (x *StoreNode) GetStore() *dbmanagement.StoreInfo {
	if x != nil {
		return x.Store
	}
	return nil
}

var File_user_whoami_proto protoreflect.FileDescriptor

var file_user_whoami_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x77, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x64, 0x62, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x57, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc0, 0x01, 0x0a, 0x0e, 0x57, 0x68, 0x6f, 0x61, 0x6d,
	0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0d, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0c, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0d, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7f, 0x0a, 0x10, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a,
	0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6c, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x65, 0x0a, 0x0b, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6c, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x73, 0x22, 0x35, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x28,
	0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_whoami_proto_rawDescOnce sync.Once
	file_user_whoami_proto_rawDescData = file_user_whoami_proto_rawDesc
)

func file_user_whoami_proto_rawDescGZIP() []byte {
	file_user_whoami_proto_rawDescOnce.Do(func() {
		file_user_whoami_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_whoami_proto_rawDescData)
	})
	return file_user_whoami_proto_rawDescData
}

var file_user_whoami_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_whoami_proto_goTypes = []interface{}{
	(*WhoamiRequest)(nil),          // 0: stately.WhoamiRequest
	(*WhoamiResponse)(nil),         // 1: stately.WhoamiResponse
	(*OrganizationNode)(nil),       // 2: stately.OrganizationNode
	(*ProjectNode)(nil),            // 3: stately.ProjectNode
	(*StoreNode)(nil),              // 4: stately.StoreNode
	(*Organization)(nil),           // 5: stately.Organization
	(*Project)(nil),                // 6: stately.Project
	(*dbmanagement.StoreInfo)(nil), // 7: stately.StoreInfo
}
var file_user_whoami_proto_depIdxs = []int32{
	2, // 0: stately.WhoamiResponse.organizations:type_name -> stately.OrganizationNode
	5, // 1: stately.OrganizationNode.organization:type_name -> stately.Organization
	3, // 2: stately.OrganizationNode.projects:type_name -> stately.ProjectNode
	6, // 3: stately.ProjectNode.project:type_name -> stately.Project
	4, // 4: stately.ProjectNode.stores:type_name -> stately.StoreNode
	7, // 5: stately.StoreNode.store:type_name -> stately.StoreInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_user_whoami_proto_init() }
func file_user_whoami_proto_init() {
	if File_user_whoami_proto != nil {
		return
	}
	file_user_organization_proto_init()
	file_user_project_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_whoami_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoamiRequest); i {
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
		file_user_whoami_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoamiResponse); i {
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
		file_user_whoami_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationNode); i {
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
		file_user_whoami_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectNode); i {
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
		file_user_whoami_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreNode); i {
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
			RawDescriptor: file_user_whoami_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_whoami_proto_goTypes,
		DependencyIndexes: file_user_whoami_proto_depIdxs,
		MessageInfos:      file_user_whoami_proto_msgTypes,
	}.Build()
	File_user_whoami_proto = out.File
	file_user_whoami_proto_rawDesc = nil
	file_user_whoami_proto_goTypes = nil
	file_user_whoami_proto_depIdxs = nil
}
