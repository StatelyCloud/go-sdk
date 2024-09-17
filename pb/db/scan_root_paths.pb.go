// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: db/scan_root_paths.proto

package db

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

type ScanRootPathsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// store_id is a globally unique Store ID, which can be looked up from the
	// console or CLI.
	StoreId uint64 `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	// Limit sets an upper bound on how many root paths to return.
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// pagination_token is an optional token to continue retrieving the next page of results.
	// This value must be read from a ScanRootPathsResponse and passed with a clone of the
	// previous request to fetch the next page of data
	PaginationToken []byte `protobuf:"bytes,3,opt,name=pagination_token,json=paginationToken,proto3" json:"pagination_token,omitempty"`
}

func (x *ScanRootPathsRequest) Reset() {
	*x = ScanRootPathsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_scan_root_paths_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRootPathsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRootPathsRequest) ProtoMessage() {}

func (x *ScanRootPathsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_scan_root_paths_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRootPathsRequest.ProtoReflect.Descriptor instead.
func (*ScanRootPathsRequest) Descriptor() ([]byte, []int) {
	return file_db_scan_root_paths_proto_rawDescGZIP(), []int{0}
}

func (x *ScanRootPathsRequest) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *ScanRootPathsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ScanRootPathsRequest) GetPaginationToken() []byte {
	if x != nil {
		return x.PaginationToken
	}
	return nil
}

type ScanRootPathsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// results is a list that contains one entry for each root path that was found.
	Results []*ScanRootPathResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// This field is optional and will be set if there are more query results to fetch.
	// To fetch the next page of results you must make the exact same ScanRootPathsRequest as before
	// but set ScanRootPathsRequest.pagination_token to the value returned here.
	PaginationToken []byte `protobuf:"bytes,2,opt,name=pagination_token,json=paginationToken,proto3" json:"pagination_token,omitempty"`
}

func (x *ScanRootPathsResponse) Reset() {
	*x = ScanRootPathsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_scan_root_paths_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRootPathsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRootPathsResponse) ProtoMessage() {}

func (x *ScanRootPathsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_scan_root_paths_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRootPathsResponse.ProtoReflect.Descriptor instead.
func (*ScanRootPathsResponse) Descriptor() ([]byte, []int) {
	return file_db_scan_root_paths_proto_rawDescGZIP(), []int{1}
}

func (x *ScanRootPathsResponse) GetResults() []*ScanRootPathResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ScanRootPathsResponse) GetPaginationToken() []byte {
	if x != nil {
		return x.PaginationToken
	}
	return nil
}

type ScanRootPathResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key_path is a single root key path. Users can Query by this root path to
	// get all of the items under it.
	KeyPath string `protobuf:"bytes,1,opt,name=key_path,json=keyPath,proto3" json:"key_path,omitempty"`
}

func (x *ScanRootPathResult) Reset() {
	*x = ScanRootPathResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_scan_root_paths_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRootPathResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRootPathResult) ProtoMessage() {}

func (x *ScanRootPathResult) ProtoReflect() protoreflect.Message {
	mi := &file_db_scan_root_paths_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRootPathResult.ProtoReflect.Descriptor instead.
func (*ScanRootPathResult) Descriptor() ([]byte, []int) {
	return file_db_scan_root_paths_proto_rawDescGZIP(), []int{2}
}

func (x *ScanRootPathResult) GetKeyPath() string {
	if x != nil {
		return x.KeyPath
	}
	return ""
}

var File_db_scan_root_paths_proto protoreflect.FileDescriptor

var file_db_scan_root_paths_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x62, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x14, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x6f, 0x6f, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x08, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x7c, 0x0a, 0x15, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x6f, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x68, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x6f, 0x6f, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x37, 0x0a,
	0x12, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x6f, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x50, 0x61, 0x74, 0x68, 0x42, 0x93, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x42, 0x12, 0x53, 0x63, 0x61, 0x6e, 0x52,
	0x6f, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x68, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x70, 0x62, 0x2f, 0x64, 0x62, 0xa2, 0x02, 0x03, 0x53, 0x44, 0x58, 0xaa, 0x02, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x44, 0x62, 0xca, 0x02, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x6c, 0x79, 0x5c, 0x44, 0x62, 0xe2, 0x02, 0x16, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x5c,
	0x44, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x3a, 0x3a, 0x44, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_scan_root_paths_proto_rawDescOnce sync.Once
	file_db_scan_root_paths_proto_rawDescData = file_db_scan_root_paths_proto_rawDesc
)

func file_db_scan_root_paths_proto_rawDescGZIP() []byte {
	file_db_scan_root_paths_proto_rawDescOnce.Do(func() {
		file_db_scan_root_paths_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_scan_root_paths_proto_rawDescData)
	})
	return file_db_scan_root_paths_proto_rawDescData
}

var file_db_scan_root_paths_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_db_scan_root_paths_proto_goTypes = []interface{}{
	(*ScanRootPathsRequest)(nil),  // 0: stately.db.ScanRootPathsRequest
	(*ScanRootPathsResponse)(nil), // 1: stately.db.ScanRootPathsResponse
	(*ScanRootPathResult)(nil),    // 2: stately.db.ScanRootPathResult
}
var file_db_scan_root_paths_proto_depIdxs = []int32{
	2, // 0: stately.db.ScanRootPathsResponse.results:type_name -> stately.db.ScanRootPathResult
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_db_scan_root_paths_proto_init() }
func file_db_scan_root_paths_proto_init() {
	if File_db_scan_root_paths_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_scan_root_paths_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRootPathsRequest); i {
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
		file_db_scan_root_paths_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRootPathsResponse); i {
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
		file_db_scan_root_paths_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRootPathResult); i {
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
			RawDescriptor: file_db_scan_root_paths_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_scan_root_paths_proto_goTypes,
		DependencyIndexes: file_db_scan_root_paths_proto_depIdxs,
		MessageInfos:      file_db_scan_root_paths_proto_msgTypes,
	}.Build()
	File_db_scan_root_paths_proto = out.File
	file_db_scan_root_paths_proto_rawDesc = nil
	file_db_scan_root_paths_proto_goTypes = nil
	file_db_scan_root_paths_proto_depIdxs = nil
}