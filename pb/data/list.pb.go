// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: data/list.proto

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

// SortDirection represents the direction of iteration.
type SortDirection int32

const (
	SortDirection_SORT_ASCENDING  SortDirection = 0 // This is the default
	SortDirection_SORT_DESCENDING SortDirection = 1
)

// Enum value maps for SortDirection.
var (
	SortDirection_name = map[int32]string{
		0: "SORT_ASCENDING",
		1: "SORT_DESCENDING",
	}
	SortDirection_value = map[string]int32{
		"SORT_ASCENDING":  0,
		"SORT_DESCENDING": 1,
	}
)

func (x SortDirection) Enum() *SortDirection {
	p := new(SortDirection)
	*p = x
	return p
}

func (x SortDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_data_list_proto_enumTypes[0].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_data_list_proto_enumTypes[0]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_data_list_proto_rawDescGZIP(), []int{0}
}

type BeginListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// store_id is a globally unique Store ID, which can be looked up from the
	// console or CLI.
	StoreId uint64 `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	// key_path_prefix is the a prefix that limits what items we will return. This
	// must contain at least a root segment. See Item#key_path for more details.
	KeyPathPrefix string `protobuf:"bytes,2,opt,name=key_path_prefix,json=keyPathPrefix,proto3" json:"key_path_prefix,omitempty"`
	// limit is the maximum number of items to return. If this is not specified or
	// set to 0, it will default to unlimited. Fewer items than the limit may be
	// returned even if there are more items to get - make sure to check
	// token.can_continue.
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// allow_stale indicates that you're okay with getting slightly stale items -
	// that is, if you had just changed an item and then call a List operation,
	// you might get the old version of the item. This can result in improved
	// performance, availability, and cost.
	AllowStale bool `protobuf:"varint,4,opt,name=allow_stale,json=allowStale,proto3" json:"allow_stale,omitempty"`
	// sort_property is the property of the item to sort the results by. If this
	// is not set, we will sort by key path.
	SortProperty SortableProperty `protobuf:"varint,5,opt,name=sort_property,json=sortProperty,proto3,enum=stately.data.SortableProperty" json:"sort_property,omitempty"`
	// sort_direction is the direction to sort the results in. If this is not set,
	// we will sort in ascending order.
	SortDirection SortDirection `protobuf:"varint,6,opt,name=sort_direction,json=sortDirection,proto3,enum=stately.data.SortDirection" json:"sort_direction,omitempty"`
}

func (x *BeginListRequest) Reset() {
	*x = BeginListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginListRequest) ProtoMessage() {}

func (x *BeginListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginListRequest.ProtoReflect.Descriptor instead.
func (*BeginListRequest) Descriptor() ([]byte, []int) {
	return file_data_list_proto_rawDescGZIP(), []int{0}
}

func (x *BeginListRequest) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *BeginListRequest) GetKeyPathPrefix() string {
	if x != nil {
		return x.KeyPathPrefix
	}
	return ""
}

func (x *BeginListRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *BeginListRequest) GetAllowStale() bool {
	if x != nil {
		return x.AllowStale
	}
	return false
}

func (x *BeginListRequest) GetSortProperty() SortableProperty {
	if x != nil {
		return x.SortProperty
	}
	return SortableProperty_SORTABLE_PROPERTY_KEY_PATH
}

func (x *BeginListRequest) GetSortDirection() SortDirection {
	if x != nil {
		return x.SortDirection
	}
	return SortDirection_SORT_ASCENDING
}

// These are stream messages, so multiple responses may be sent.
type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*ListResponse_Result
	//	*ListResponse_Finished
	Response isListResponse_Response `protobuf_oneof:"response"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_data_list_proto_rawDescGZIP(), []int{1}
}

func (m *ListResponse) GetResponse() isListResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *ListResponse) GetResult() *ListPartialResult {
	if x, ok := x.GetResponse().(*ListResponse_Result); ok {
		return x.Result
	}
	return nil
}

func (x *ListResponse) GetFinished() *ListFinished {
	if x, ok := x.GetResponse().(*ListResponse_Finished); ok {
		return x.Finished
	}
	return nil
}

type isListResponse_Response interface {
	isListResponse_Response()
}

type ListResponse_Result struct {
	// Result is a segment of the result set - multiple of these may be returned
	// in a stream before the final "finished" message.
	Result *ListPartialResult `protobuf:"bytes,1,opt,name=result,proto3,oneof"`
}

type ListResponse_Finished struct {
	// Finished is sent when there are no more results in this operation, and
	// there will only be one.
	Finished *ListFinished `protobuf:"bytes,2,opt,name=finished,proto3,oneof"`
}

func (*ListResponse_Result) isListResponse_Response() {}

func (*ListResponse_Finished) isListResponse_Response() {}

type ListPartialResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// results is a list that contains one entry for each Item that was found.
	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListPartialResult) Reset() {
	*x = ListPartialResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPartialResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPartialResult) ProtoMessage() {}

func (x *ListPartialResult) ProtoReflect() protoreflect.Message {
	mi := &file_data_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPartialResult.ProtoReflect.Descriptor instead.
func (*ListPartialResult) Descriptor() ([]byte, []int) {
	return file_data_list_proto_rawDescGZIP(), []int{2}
}

func (x *ListPartialResult) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListFinished struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// token is always set and represents an updated list continuation token that
	// can be used in subsequent calls to ContinueList or SyncList.
	Token *ListToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ListFinished) Reset() {
	*x = ListFinished{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFinished) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFinished) ProtoMessage() {}

func (x *ListFinished) ProtoReflect() protoreflect.Message {
	mi := &file_data_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFinished.ProtoReflect.Descriptor instead.
func (*ListFinished) Descriptor() ([]byte, []int) {
	return file_data_list_proto_rawDescGZIP(), []int{3}
}

func (x *ListFinished) GetToken() *ListToken {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_data_list_proto protoreflect.FileDescriptor

var file_data_list_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5,
	0x02, 0x0a, 0x10, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0f, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0d, 0x6b, 0x65, 0x79, 0x50, 0x61, 0x74, 0x68,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x6c, 0x65, 0x12, 0x43, 0x0a,
	0x0d, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x52, 0x0c, 0x73, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x12, 0x42, 0x0a, 0x0e, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x96, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c,
	0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x42, 0x11, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x22,
	0x45, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x45, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0x38, 0x0a,
	0x0d, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x0e, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x42, 0x96, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0xa2, 0x02, 0x03, 0x53, 0x44, 0x58, 0xaa, 0x02, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c,
	0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0xca, 0x02, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79,
	0x5c, 0x44, 0x61, 0x74, 0x61, 0xe2, 0x02, 0x18, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x5c,
	0x44, 0x61, 0x74, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_list_proto_rawDescOnce sync.Once
	file_data_list_proto_rawDescData = file_data_list_proto_rawDesc
)

func file_data_list_proto_rawDescGZIP() []byte {
	file_data_list_proto_rawDescOnce.Do(func() {
		file_data_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_list_proto_rawDescData)
	})
	return file_data_list_proto_rawDescData
}

var file_data_list_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_data_list_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_data_list_proto_goTypes = []interface{}{
	(SortDirection)(0),        // 0: stately.data.SortDirection
	(*BeginListRequest)(nil),  // 1: stately.data.BeginListRequest
	(*ListResponse)(nil),      // 2: stately.data.ListResponse
	(*ListPartialResult)(nil), // 3: stately.data.ListPartialResult
	(*ListFinished)(nil),      // 4: stately.data.ListFinished
	(SortableProperty)(0),     // 5: stately.data.SortableProperty
	(*Item)(nil),              // 6: stately.data.Item
	(*ListToken)(nil),         // 7: stately.data.ListToken
}
var file_data_list_proto_depIdxs = []int32{
	5, // 0: stately.data.BeginListRequest.sort_property:type_name -> stately.data.SortableProperty
	0, // 1: stately.data.BeginListRequest.sort_direction:type_name -> stately.data.SortDirection
	3, // 2: stately.data.ListResponse.result:type_name -> stately.data.ListPartialResult
	4, // 3: stately.data.ListResponse.finished:type_name -> stately.data.ListFinished
	6, // 4: stately.data.ListPartialResult.items:type_name -> stately.data.Item
	7, // 5: stately.data.ListFinished.token:type_name -> stately.data.ListToken
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_data_list_proto_init() }
func file_data_list_proto_init() {
	if File_data_list_proto != nil {
		return
	}
	file_data_item_proto_init()
	file_data_item_property_proto_init()
	file_data_list_token_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_data_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginListRequest); i {
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
		file_data_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_data_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPartialResult); i {
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
		file_data_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFinished); i {
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
	file_data_list_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ListResponse_Result)(nil),
		(*ListResponse_Finished)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_list_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_list_proto_goTypes,
		DependencyIndexes: file_data_list_proto_depIdxs,
		EnumInfos:         file_data_list_proto_enumTypes,
		MessageInfos:      file_data_list_proto_msgTypes,
	}.Build()
	File_data_list_proto = out.File
	file_data_list_proto_rawDesc = nil
	file_data_list_proto_goTypes = nil
	file_data_list_proto_depIdxs = nil
}
