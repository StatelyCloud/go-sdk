// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: data/query.proto

package data

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

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// store_id is a globally unique Store ID, which can be looked up from the
	// console or CLI.
	StoreId uint64 `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	// key_path_prefix is the a prefix that limits what items we will return. This
	// must contain at least a root segment. See Item#key_path for more details.
	KeyPathPrefix string `protobuf:"bytes,2,opt,name=key_path_prefix,json=keyPathPrefix,proto3" json:"key_path_prefix,omitempty"`
	// pagination_token is an optional token to continue retrieving the next page of results.
	// This value must be read from a QueryResponse and passed with a clone of the
	// previous request to fetch the next page of data
	PaginationToken []byte `protobuf:"bytes,3,opt,name=pagination_token,json=paginationToken,proto3" json:"pagination_token,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_data_query_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRequest) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *QueryRequest) GetKeyPathPrefix() string {
	if x != nil {
		return x.KeyPathPrefix
	}
	return ""
}

func (x *QueryRequest) GetPaginationToken() []byte {
	if x != nil {
		return x.PaginationToken
	}
	return nil
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// results is a list that contains one entry for each Item that was found.
	Results []*QueryItemsResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// This field is optional and will be set if there are more query results to fetch.
	// To fetch the next page of results you must make the exact same QueryRequest as before
	// but set QueryRequest.pagination_token to the value returned here.
	PaginationToken []byte `protobuf:"bytes,2,opt,name=pagination_token,json=paginationToken,proto3" json:"pagination_token,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_data_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryResponse) GetResults() []*QueryItemsResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *QueryResponse) GetPaginationToken() []byte {
	if x != nil {
		return x.PaginationToken
	}
	return nil
}

// We return a list of these result messages because someday we may want to
// return something other than just an item.
type QueryItemsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Item is the full item, with metadata, returned from the server.
	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *QueryItemsResult) Reset() {
	*x = QueryItemsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryItemsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryItemsResult) ProtoMessage() {}

func (x *QueryItemsResult) ProtoReflect() protoreflect.Message {
	mi := &file_data_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryItemsResult.ProtoReflect.Descriptor instead.
func (*QueryItemsResult) Descriptor() ([]byte, []int) {
	return file_data_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryItemsResult) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_data_query_proto protoreflect.FileDescriptor

var file_data_query_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x1a, 0x0f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x08, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x0f, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0d,
	0x6b, 0x65, 0x79, 0x50, 0x61, 0x74, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x29, 0x0a,
	0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6f, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6c, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x29,
	0x0a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3d, 0x0a, 0x10, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x29, 0x0a,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_query_proto_rawDescOnce sync.Once
	file_data_query_proto_rawDescData = file_data_query_proto_rawDesc
)

func file_data_query_proto_rawDescGZIP() []byte {
	file_data_query_proto_rawDescOnce.Do(func() {
		file_data_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_query_proto_rawDescData)
	})
	return file_data_query_proto_rawDescData
}

var file_data_query_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_data_query_proto_goTypes = []interface{}{
	(*QueryRequest)(nil),     // 0: stately.QueryRequest
	(*QueryResponse)(nil),    // 1: stately.QueryResponse
	(*QueryItemsResult)(nil), // 2: stately.QueryItemsResult
	(*Item)(nil),             // 3: stately.Item
}
var file_data_query_proto_depIdxs = []int32{
	2, // 0: stately.QueryResponse.results:type_name -> stately.QueryItemsResult
	3, // 1: stately.QueryItemsResult.item:type_name -> stately.Item
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_data_query_proto_init() }
func file_data_query_proto_init() {
	if File_data_query_proto != nil {
		return
	}
	file_data_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_data_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_data_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
		file_data_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryItemsResult); i {
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
			RawDescriptor: file_data_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_query_proto_goTypes,
		DependencyIndexes: file_data_query_proto_depIdxs,
		MessageInfos:      file_data_query_proto_msgTypes,
	}.Build()
	File_data_query_proto = out.File
	file_data_query_proto_rawDesc = nil
	file_data_query_proto_goTypes = nil
	file_data_query_proto_depIdxs = nil
}
