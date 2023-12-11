// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: dbmanagement/describe-store.proto

package dbmanagement

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

type DescribeStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identifier:
	//
	//	*DescribeStoreRequest_StoreId
	//	*DescribeStoreRequest_StoreLookup
	Identifier isDescribeStoreRequest_Identifier `protobuf_oneof:"identifier"`
}

func (x *DescribeStoreRequest) Reset() {
	*x = DescribeStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmanagement_describe_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeStoreRequest) ProtoMessage() {}

func (x *DescribeStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dbmanagement_describe_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeStoreRequest.ProtoReflect.Descriptor instead.
func (*DescribeStoreRequest) Descriptor() ([]byte, []int) {
	return file_dbmanagement_describe_store_proto_rawDescGZIP(), []int{0}
}

func (m *DescribeStoreRequest) GetIdentifier() isDescribeStoreRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *DescribeStoreRequest) GetStoreId() uint64 {
	if x, ok := x.GetIdentifier().(*DescribeStoreRequest_StoreId); ok {
		return x.StoreId
	}
	return 0
}

func (x *DescribeStoreRequest) GetStoreLookup() *StoreLookup {
	if x, ok := x.GetIdentifier().(*DescribeStoreRequest_StoreLookup); ok {
		return x.StoreLookup
	}
	return nil
}

type isDescribeStoreRequest_Identifier interface {
	isDescribeStoreRequest_Identifier()
}

type DescribeStoreRequest_StoreId struct {
	// store_id is a globally unique identifier for a store. Users will of course
	// be able to name their stores with friendly names, but for efficiency's sake
	// the API talks in terms of IDs, which the user can find in the console or
	// via a DescribeStore command. Store IDs are assigned by the system when a
	// store is created.
	StoreId uint64 `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3,oneof"`
}

type DescribeStoreRequest_StoreLookup struct {
	// In some cases, for example in the CLI, the user may not already have the
	// store ID, but they do know the store's name. This provides an option for
	// directly looking up the store by name without having to make two calls,
	// e.g. calling ListStores and then DescribeStore.
	StoreLookup *StoreLookup `protobuf:"bytes,2,opt,name=store_lookup,json=storeLookup,proto3,oneof"`
}

func (*DescribeStoreRequest_StoreId) isDescribeStoreRequest_Identifier() {}

func (*DescribeStoreRequest_StoreLookup) isDescribeStoreRequest_Identifier() {}

type DescribeStoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All info about the requested store. This object can be modified and sent
	// back via UpdateStore to update the store's configuration.
	Store *StoreInfo `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
}

func (x *DescribeStoreResponse) Reset() {
	*x = DescribeStoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmanagement_describe_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeStoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeStoreResponse) ProtoMessage() {}

func (x *DescribeStoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dbmanagement_describe_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeStoreResponse.ProtoReflect.Descriptor instead.
func (*DescribeStoreResponse) Descriptor() ([]byte, []int) {
	return file_dbmanagement_describe_store_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeStoreResponse) GetStore() *StoreInfo {
	if x != nil {
		return x.Store
	}
	return nil
}

// StoreLookup is an option for looking up stores when you do not have their
// Store ID.
type StoreLookup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project_id is the project within which the store will looked up by name.
	// Project ID is a globally unique ID tied to a project. We do not need
	// User ID because projects exist in only a single organization. We assume
	// the ID will be available from the CLI or website, which will know it from
	// having logged in / selected an active project.
	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// name is a user-facing, memorable name for the store. While most APIs deal
	// strictly in store IDs, the name will be shown in the console and usable
	// within the API.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StoreLookup) Reset() {
	*x = StoreLookup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmanagement_describe_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreLookup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreLookup) ProtoMessage() {}

func (x *StoreLookup) ProtoReflect() protoreflect.Message {
	mi := &file_dbmanagement_describe_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreLookup.ProtoReflect.Descriptor instead.
func (*StoreLookup) Descriptor() ([]byte, []int) {
	return file_dbmanagement_describe_store_proto_rawDescGZIP(), []int{2}
}

func (x *StoreLookup) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *StoreLookup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_dbmanagement_describe_store_proto protoreflect.FileDescriptor

var file_dbmanagement_describe_store_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x64, 0x62, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x42, 0x13, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x22, 0x49,
	0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x53, 0x0a, 0x0b, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x25, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba,
	0x48, 0x06, 0x72, 0x04, 0x20, 0x01, 0x28, 0x40, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dbmanagement_describe_store_proto_rawDescOnce sync.Once
	file_dbmanagement_describe_store_proto_rawDescData = file_dbmanagement_describe_store_proto_rawDesc
)

func file_dbmanagement_describe_store_proto_rawDescGZIP() []byte {
	file_dbmanagement_describe_store_proto_rawDescOnce.Do(func() {
		file_dbmanagement_describe_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_dbmanagement_describe_store_proto_rawDescData)
	})
	return file_dbmanagement_describe_store_proto_rawDescData
}

var file_dbmanagement_describe_store_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dbmanagement_describe_store_proto_goTypes = []interface{}{
	(*DescribeStoreRequest)(nil),  // 0: stately.DescribeStoreRequest
	(*DescribeStoreResponse)(nil), // 1: stately.DescribeStoreResponse
	(*StoreLookup)(nil),           // 2: stately.StoreLookup
	(*StoreInfo)(nil),             // 3: stately.StoreInfo
}
var file_dbmanagement_describe_store_proto_depIdxs = []int32{
	2, // 0: stately.DescribeStoreRequest.store_lookup:type_name -> stately.StoreLookup
	3, // 1: stately.DescribeStoreResponse.store:type_name -> stately.StoreInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dbmanagement_describe_store_proto_init() }
func file_dbmanagement_describe_store_proto_init() {
	if File_dbmanagement_describe_store_proto != nil {
		return
	}
	file_dbmanagement_store_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dbmanagement_describe_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeStoreRequest); i {
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
		file_dbmanagement_describe_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeStoreResponse); i {
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
		file_dbmanagement_describe_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreLookup); i {
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
	file_dbmanagement_describe_store_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DescribeStoreRequest_StoreId)(nil),
		(*DescribeStoreRequest_StoreLookup)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dbmanagement_describe_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dbmanagement_describe_store_proto_goTypes,
		DependencyIndexes: file_dbmanagement_describe_store_proto_depIdxs,
		MessageInfos:      file_dbmanagement_describe_store_proto_msgTypes,
	}.Build()
	File_dbmanagement_describe_store_proto = out.File
	file_dbmanagement_describe_store_proto_rawDesc = nil
	file_dbmanagement_describe_store_proto_goTypes = nil
	file_dbmanagement_describe_store_proto_depIdxs = nil
}
