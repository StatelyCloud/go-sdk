// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.2
// source: data/delete.proto

package data

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	common "github.com/StatelyCloud/go-sdk/pb/common"
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

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// store_id is a globally unique Store ID, which can be looked up from the
	// console or CLI.
	StoreId uint64 `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	// deletes is one or more items to be deleted from the Group.
	Deletes []*DeleteItem `protobuf:"bytes,3,rep,name=deletes,proto3" json:"deletes,omitempty"`
	// atomic indicates that all deletes must succeed or none will (i.e. that they
	// are applied in a transaction), and that other operations will be serialized
	// ahead or behind this operation. Some store configurations may ignore this
	// option and will always apply the whole batch in a transaction (such as in
	// version-tracking stores). Note that this has no effect if there is only one
	// delete. Enabling this option increases cost and latency, and may result in
	// the operation failing if it conflicts with another atomic operation.
	Atomic bool `protobuf:"varint,4,opt,name=atomic,proto3" json:"atomic,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_delete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_data_delete_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteRequest) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *DeleteRequest) GetDeletes() []*DeleteItem {
	if x != nil {
		return x.Deletes
	}
	return nil
}

func (x *DeleteRequest) GetAtomic() bool {
	if x != nil {
		return x.Atomic
	}
	return false
}

type DeleteItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key_path is the full path to the item. See Item#key_path for more details.
	KeyPath string `protobuf:"bytes,1,opt,name=key_path,json=keyPath,proto3" json:"key_path,omitempty"`
}

func (x *DeleteItem) Reset() {
	*x = DeleteItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_delete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItem) ProtoMessage() {}

func (x *DeleteItem) ProtoReflect() protoreflect.Message {
	mi := &file_data_delete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItem.ProtoReflect.Descriptor instead.
func (*DeleteItem) Descriptor() ([]byte, []int) {
	return file_data_delete_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteItem) GetKeyPath() string {
	if x != nil {
		return x.KeyPath
	}
	return ""
}

type DeleteResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key_path of the item that was deleted.
	KeyPath string `protobuf:"bytes,1,opt,name=key_path,json=keyPath,proto3" json:"key_path,omitempty"`
	// error is the error that occurred while deleting this item, if any. error is
	// not set if the item was successfully deleted (or didn't exist in the first
	// place).
	Error *common.OperationError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteResult) Reset() {
	*x = DeleteResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_delete_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResult) ProtoMessage() {}

func (x *DeleteResult) ProtoReflect() protoreflect.Message {
	mi := &file_data_delete_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResult.ProtoReflect.Descriptor instead.
func (*DeleteResult) Descriptor() ([]byte, []int) {
	return file_data_delete_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteResult) GetKeyPath() string {
	if x != nil {
		return x.KeyPath
	}
	return ""
}

func (x *DeleteResult) GetError() *common.OperationError {
	if x != nil {
		return x.Error
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// results is the result of each delete operation, whether it succeeded or failed.
	Results []*DeleteResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_delete_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_delete_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_data_delete_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteResponse) GetResults() []*DeleteResult {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_data_delete_proto protoreflect.FileDescriptor

var file_data_delete_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x92, 0x01, 0x04,
	0x08, 0x01, 0x10, 0x19, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61,
	0x74, 0x6f, 0x6d, 0x69, 0x63, 0x22, 0x2f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x21, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x50, 0x61, 0x74, 0x68, 0x22, 0x60, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x07, 0x6b, 0x65, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x6c, 0x79, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x41, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_data_delete_proto_rawDescOnce sync.Once
	file_data_delete_proto_rawDescData = file_data_delete_proto_rawDesc
)

func file_data_delete_proto_rawDescGZIP() []byte {
	file_data_delete_proto_rawDescOnce.Do(func() {
		file_data_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_delete_proto_rawDescData)
	})
	return file_data_delete_proto_rawDescData
}

var file_data_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_data_delete_proto_goTypes = []interface{}{
	(*DeleteRequest)(nil),         // 0: stately.DeleteRequest
	(*DeleteItem)(nil),            // 1: stately.DeleteItem
	(*DeleteResult)(nil),          // 2: stately.DeleteResult
	(*DeleteResponse)(nil),        // 3: stately.DeleteResponse
	(*common.OperationError)(nil), // 4: stately.OperationError
}
var file_data_delete_proto_depIdxs = []int32{
	1, // 0: stately.DeleteRequest.deletes:type_name -> stately.DeleteItem
	4, // 1: stately.DeleteResult.error:type_name -> stately.OperationError
	2, // 2: stately.DeleteResponse.results:type_name -> stately.DeleteResult
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_data_delete_proto_init() }
func file_data_delete_proto_init() {
	if File_data_delete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_data_delete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItem); i {
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
		file_data_delete_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResult); i {
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
		file_data_delete_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_data_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_delete_proto_goTypes,
		DependencyIndexes: file_data_delete_proto_depIdxs,
		MessageInfos:      file_data_delete_proto_msgTypes,
	}.Build()
	File_data_delete_proto = out.File
	file_data_delete_proto_rawDesc = nil
	file_data_delete_proto_goTypes = nil
	file_data_delete_proto_depIdxs = nil
}
