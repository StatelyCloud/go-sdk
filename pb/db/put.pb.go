// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: db/put.proto

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

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// store_id is a globally unique Store ID, which can be looked up from the
	// console or CLI.
	StoreId uint64 `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	// puts is up to 50 items to be put into the Store.
	Puts []*PutItem `protobuf:"bytes,2,rep,name=puts,proto3" json:"puts,omitempty"`
	// schema_version_id refers to the item version to return.
	//
	// If the store's schema does not have this version, the operation will error
	// with SchemaVersionNotFound error. You should not have to set this manually
	// as your generated SDK should know its schema version and wire this in for
	// you.
	SchemaVersionId uint32 `protobuf:"varint,3,opt,name=schema_version_id,json=schemaVersionId,proto3" json:"schema_version_id,omitempty"`
	// schema_id refers to the schema to use for this operation.
	// If the store_id does not have a schema with this ID, the operation will
	// error with SchemaNotFound error. You should not have to set this manually
	// as your generated SDK should know its schema and wire this in for you.
	SchemaId uint64 `protobuf:"varint,4,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"` // [(buf.validate.field).required = true]; (after clients have been regen'd and updated)
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_put_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_put_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_db_put_proto_rawDescGZIP(), []int{0}
}

func (x *PutRequest) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *PutRequest) GetPuts() []*PutItem {
	if x != nil {
		return x.Puts
	}
	return nil
}

func (x *PutRequest) GetSchemaVersionId() uint32 {
	if x != nil {
		return x.SchemaVersionId
	}
	return 0
}

func (x *PutRequest) GetSchemaId() uint64 {
	if x != nil {
		return x.SchemaId
	}
	return 0
}

type PutItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// item is the data to be put, including its item_type.
	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	// overwrite_metadata_timestamps indicates that any "fromMetadata" timestamp
	// fields in the incoming payload should be saved as provided in the database.
	// Normally these would be ignored as they are automatically maintained, but
	// this flag can be useful for migrations from other systems. Note that this
	// only works for timestamps (createdAtTime and lastModifiedAtTime) - versions
	// cannot be overridden.
	OverwriteMetadataTimestamps bool `protobuf:"varint,2,opt,name=overwrite_metadata_timestamps,json=overwriteMetadataTimestamps,proto3" json:"overwrite_metadata_timestamps,omitempty"`
	// must_not_exist is a condition that indicates this item must not already
	// exist at any of its key paths. If there is already an item at one of those
	// paths, the Put operation will fail with a ConditionalCheckFailed error.
	// Note that if the item has an `initialValue` field in its key, that initial
	// value will automatically be chosen not to conflict with existing items, so
	// this condition only applies to key paths that do not contain the
	// `initialValue` field.
	MustNotExist bool `protobuf:"varint,3,opt,name=must_not_exist,json=mustNotExist,proto3" json:"must_not_exist,omitempty"`
}

func (x *PutItem) Reset() {
	*x = PutItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_put_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutItem) ProtoMessage() {}

func (x *PutItem) ProtoReflect() protoreflect.Message {
	mi := &file_db_put_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutItem.ProtoReflect.Descriptor instead.
func (*PutItem) Descriptor() ([]byte, []int) {
	return file_db_put_proto_rawDescGZIP(), []int{1}
}

func (x *PutItem) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *PutItem) GetOverwriteMetadataTimestamps() bool {
	if x != nil {
		return x.OverwriteMetadataTimestamps
	}
	return false
}

func (x *PutItem) GetMustNotExist() bool {
	if x != nil {
		return x.MustNotExist
	}
	return false
}

type PutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// items is the full result of each put operation. The response items are in
	// the same order as the request items. Each item is fully "filled out" - for
	// example, `initialValue` and `fromMetadata` fields are resolved.
	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *PutResponse) Reset() {
	*x = PutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_put_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResponse) ProtoMessage() {}

func (x *PutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_put_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResponse.ProtoReflect.Descriptor instead.
func (*PutResponse) Descriptor() ([]byte, []int) {
	return file_db_put_proto_rawDescGZIP(), []int{2}
}

func (x *PutResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_db_put_proto protoreflect.FileDescriptor

var file_db_put_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x62, 0x2f, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x64, 0x62, 0x2f, 0x69, 0x74, 0x65, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x70, 0x75, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79,
	0x2e, 0x64, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x0a, 0xba, 0x48, 0x07,
	0x92, 0x01, 0x04, 0x08, 0x01, 0x10, 0x32, 0x52, 0x04, 0x70, 0x75, 0x74, 0x73, 0x12, 0x32, 0x0a,
	0x11, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x22, 0xa1,
	0x01, 0x0a, 0x07, 0x50, 0x75, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2c, 0x0a, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x6c, 0x79, 0x2e, 0x64, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x42, 0x0a, 0x1d, 0x6f, 0x76, 0x65, 0x72,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x1b, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x12, 0x24, 0x0a, 0x0e,
	0x6d, 0x75, 0x73, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6d, 0x75, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x22, 0x3d, 0x0a, 0x0b, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x42, 0x89, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c,
	0x79, 0x2e, 0x64, 0x62, 0x42, 0x08, 0x50, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
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
	file_db_put_proto_rawDescOnce sync.Once
	file_db_put_proto_rawDescData = file_db_put_proto_rawDesc
)

func file_db_put_proto_rawDescGZIP() []byte {
	file_db_put_proto_rawDescOnce.Do(func() {
		file_db_put_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_put_proto_rawDescData)
	})
	return file_db_put_proto_rawDescData
}

var file_db_put_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_db_put_proto_goTypes = []interface{}{
	(*PutRequest)(nil),  // 0: stately.db.PutRequest
	(*PutItem)(nil),     // 1: stately.db.PutItem
	(*PutResponse)(nil), // 2: stately.db.PutResponse
	(*Item)(nil),        // 3: stately.db.Item
}
var file_db_put_proto_depIdxs = []int32{
	1, // 0: stately.db.PutRequest.puts:type_name -> stately.db.PutItem
	3, // 1: stately.db.PutItem.item:type_name -> stately.db.Item
	3, // 2: stately.db.PutResponse.items:type_name -> stately.db.Item
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_db_put_proto_init() }
func file_db_put_proto_init() {
	if File_db_put_proto != nil {
		return
	}
	file_db_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_db_put_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequest); i {
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
		file_db_put_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutItem); i {
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
		file_db_put_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutResponse); i {
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
			RawDescriptor: file_db_put_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_put_proto_goTypes,
		DependencyIndexes: file_db_put_proto_depIdxs,
		MessageInfos:      file_db_put_proto_msgTypes,
	}.Build()
	File_db_put_proto = out.File
	file_db_put_proto_rawDesc = nil
	file_db_put_proto_goTypes = nil
	file_db_put_proto_depIdxs = nil
}
