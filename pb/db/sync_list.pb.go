// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: db/sync_list.proto

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

type SyncListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// token_data is an opaque list continuation token returned by a previous call to
	// List, ContinueList, or SyncList.
	TokenData []byte `protobuf:"bytes,1,opt,name=token_data,json=tokenData,proto3" json:"token_data,omitempty"`
}

func (x *SyncListRequest) Reset() {
	*x = SyncListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_sync_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncListRequest) ProtoMessage() {}

func (x *SyncListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_sync_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncListRequest.ProtoReflect.Descriptor instead.
func (*SyncListRequest) Descriptor() ([]byte, []int) {
	return file_db_sync_list_proto_rawDescGZIP(), []int{0}
}

func (x *SyncListRequest) GetTokenData() []byte {
	if x != nil {
		return x.TokenData
	}
	return nil
}

// These are stream messages, so multiple responses may be sent.
type SyncListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*SyncListResponse_Reset_
	//	*SyncListResponse_Result
	//	*SyncListResponse_Finished
	Response isSyncListResponse_Response `protobuf_oneof:"response"`
}

func (x *SyncListResponse) Reset() {
	*x = SyncListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_sync_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncListResponse) ProtoMessage() {}

func (x *SyncListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_sync_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncListResponse.ProtoReflect.Descriptor instead.
func (*SyncListResponse) Descriptor() ([]byte, []int) {
	return file_db_sync_list_proto_rawDescGZIP(), []int{1}
}

func (m *SyncListResponse) GetResponse() isSyncListResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *SyncListResponse) GetReset_() *SyncListReset {
	if x, ok := x.GetResponse().(*SyncListResponse_Reset_); ok {
		return x.Reset_
	}
	return nil
}

func (x *SyncListResponse) GetResult() *SyncListPartialResponse {
	if x, ok := x.GetResponse().(*SyncListResponse_Result); ok {
		return x.Result
	}
	return nil
}

func (x *SyncListResponse) GetFinished() *ListFinished {
	if x, ok := x.GetResponse().(*SyncListResponse_Finished); ok {
		return x.Finished
	}
	return nil
}

type isSyncListResponse_Response interface {
	isSyncListResponse_Response()
}

type SyncListResponse_Reset_ struct {
	// SyncListReset is returned if the provided token is too far behind to be able to
	// report deleted items, and subsequent results will start over with a fresh result
	// set. Clients should discard any cached data from this result set and start re-populating it.
	Reset_ *SyncListReset `protobuf:"bytes,1,opt,name=reset,proto3,oneof"`
}

type SyncListResponse_Result struct {
	// Result is a segment of sync results - multiple of these may be returned.
	Result *SyncListPartialResponse `protobuf:"bytes,2,opt,name=result,proto3,oneof"`
}

type SyncListResponse_Finished struct {
	// Finished is sent when the sync is complete, and there will only be one.
	Finished *ListFinished `protobuf:"bytes,3,opt,name=finished,proto3,oneof"`
}

func (*SyncListResponse_Reset_) isSyncListResponse_Response() {}

func (*SyncListResponse_Result) isSyncListResponse_Response() {}

func (*SyncListResponse_Finished) isSyncListResponse_Response() {}

// SyncListReset is returned if the provided token is too far behind to be able to
// report deleted items, and subsequent results will start over with a fresh result
// set. Clients should discard any cached data from this result set and start re-populating it.
type SyncListReset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncListReset) Reset() {
	*x = SyncListReset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_sync_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncListReset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncListReset) ProtoMessage() {}

func (x *SyncListReset) ProtoReflect() protoreflect.Message {
	mi := &file_db_sync_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncListReset.ProtoReflect.Descriptor instead.
func (*SyncListReset) Descriptor() ([]byte, []int) {
	return file_db_sync_list_proto_rawDescGZIP(), []int{2}
}

type SyncListPartialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Items in the token window that were added or updated since the last
	// sync/list.
	ChangedItems []*Item `protobuf:"bytes,1,rep,name=changed_items,json=changedItems,proto3" json:"changed_items,omitempty"`
	// Items in the token window that were deleted since the last sync/list.
	DeletedItems []*DeletedItem `protobuf:"bytes,2,rep,name=deleted_items,json=deletedItems,proto3" json:"deleted_items,omitempty"`
	// Keys of items that were updated but Stately cannot tell if they were in the
	// sync window. Treat these as deleted in most cases. For more information
	// see: https://docs.stately.cloud/api/sync
	UpdatedItemKeysOutsideListWindow []string `protobuf:"bytes,3,rep,name=updated_item_keys_outside_list_window,json=updatedItemKeysOutsideListWindow,proto3" json:"updated_item_keys_outside_list_window,omitempty"`
}

func (x *SyncListPartialResponse) Reset() {
	*x = SyncListPartialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_sync_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncListPartialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncListPartialResponse) ProtoMessage() {}

func (x *SyncListPartialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_sync_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncListPartialResponse.ProtoReflect.Descriptor instead.
func (*SyncListPartialResponse) Descriptor() ([]byte, []int) {
	return file_db_sync_list_proto_rawDescGZIP(), []int{3}
}

func (x *SyncListPartialResponse) GetChangedItems() []*Item {
	if x != nil {
		return x.ChangedItems
	}
	return nil
}

func (x *SyncListPartialResponse) GetDeletedItems() []*DeletedItem {
	if x != nil {
		return x.DeletedItems
	}
	return nil
}

func (x *SyncListPartialResponse) GetUpdatedItemKeysOutsideListWindow() []string {
	if x != nil {
		return x.UpdatedItemKeysOutsideListWindow
	}
	return nil
}

type DeletedItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Since the item was deleted, only the key is provided.
	KeyPath string `protobuf:"bytes,1,opt,name=key_path,json=keyPath,proto3" json:"key_path,omitempty"`
}

func (x *DeletedItem) Reset() {
	*x = DeletedItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_sync_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletedItem) ProtoMessage() {}

func (x *DeletedItem) ProtoReflect() protoreflect.Message {
	mi := &file_db_sync_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletedItem.ProtoReflect.Descriptor instead.
func (*DeletedItem) Descriptor() ([]byte, []int) {
	return file_db_sync_list_proto_rawDescGZIP(), []int{4}
}

func (x *DeletedItem) GetKeyPath() string {
	if x != nil {
		return x.KeyPath
	}
	return ""
}

var File_db_sync_list_proto protoreflect.FileDescriptor

var file_db_sync_list_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x62, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x64,
	0x62, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x64, 0x62,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0f, 0x53,
	0x79, 0x6e, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0xcf, 0x01, 0x0a, 0x10, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x3d, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x08,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x42, 0x11, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x79, 0x6e, 0x63, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x22, 0xdf, 0x01, 0x0a, 0x17, 0x53, 0x79, 0x6e,
	0x63, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x4f, 0x0a, 0x25, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x6f,
	0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x77, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x4b, 0x65, 0x79, 0x73, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x22, 0x30, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x50, 0x61, 0x74, 0x68, 0x42, 0x8e, 0x01, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x42,
	0x0d, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x6c, 0x79, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x62, 0x2f, 0x64, 0x62, 0xa2, 0x02, 0x03, 0x53, 0x44, 0x58, 0xaa, 0x02, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x44, 0x62, 0xca, 0x02, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x5c, 0x44, 0x62, 0xe2, 0x02, 0x16, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79,
	0x5c, 0x44, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x3a, 0x3a, 0x44, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_sync_list_proto_rawDescOnce sync.Once
	file_db_sync_list_proto_rawDescData = file_db_sync_list_proto_rawDesc
)

func file_db_sync_list_proto_rawDescGZIP() []byte {
	file_db_sync_list_proto_rawDescOnce.Do(func() {
		file_db_sync_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_sync_list_proto_rawDescData)
	})
	return file_db_sync_list_proto_rawDescData
}

var file_db_sync_list_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_db_sync_list_proto_goTypes = []interface{}{
	(*SyncListRequest)(nil),         // 0: stately.db.SyncListRequest
	(*SyncListResponse)(nil),        // 1: stately.db.SyncListResponse
	(*SyncListReset)(nil),           // 2: stately.db.SyncListReset
	(*SyncListPartialResponse)(nil), // 3: stately.db.SyncListPartialResponse
	(*DeletedItem)(nil),             // 4: stately.db.DeletedItem
	(*ListFinished)(nil),            // 5: stately.db.ListFinished
	(*Item)(nil),                    // 6: stately.db.Item
}
var file_db_sync_list_proto_depIdxs = []int32{
	2, // 0: stately.db.SyncListResponse.reset:type_name -> stately.db.SyncListReset
	3, // 1: stately.db.SyncListResponse.result:type_name -> stately.db.SyncListPartialResponse
	5, // 2: stately.db.SyncListResponse.finished:type_name -> stately.db.ListFinished
	6, // 3: stately.db.SyncListPartialResponse.changed_items:type_name -> stately.db.Item
	4, // 4: stately.db.SyncListPartialResponse.deleted_items:type_name -> stately.db.DeletedItem
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_db_sync_list_proto_init() }
func file_db_sync_list_proto_init() {
	if File_db_sync_list_proto != nil {
		return
	}
	file_db_item_proto_init()
	file_db_list_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_db_sync_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncListRequest); i {
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
		file_db_sync_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncListResponse); i {
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
		file_db_sync_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncListReset); i {
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
		file_db_sync_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncListPartialResponse); i {
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
		file_db_sync_list_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletedItem); i {
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
	file_db_sync_list_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SyncListResponse_Reset_)(nil),
		(*SyncListResponse_Result)(nil),
		(*SyncListResponse_Finished)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_db_sync_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_sync_list_proto_goTypes,
		DependencyIndexes: file_db_sync_list_proto_depIdxs,
		MessageInfos:      file_db_sync_list_proto_msgTypes,
	}.Build()
	File_db_sync_list_proto = out.File
	file_db_sync_list_proto_rawDesc = nil
	file_db_sync_list_proto_goTypes = nil
	file_db_sync_list_proto_depIdxs = nil
}
