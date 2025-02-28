// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: db/scan.proto

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

type FilterCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*FilterCondition_ItemType
	Value isFilterCondition_Value `protobuf_oneof:"value"`
}

func (x *FilterCondition) Reset() {
	*x = FilterCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_scan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterCondition) ProtoMessage() {}

func (x *FilterCondition) ProtoReflect() protoreflect.Message {
	mi := &file_db_scan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterCondition.ProtoReflect.Descriptor instead.
func (*FilterCondition) Descriptor() ([]byte, []int) {
	return file_db_scan_proto_rawDescGZIP(), []int{0}
}

func (m *FilterCondition) GetValue() isFilterCondition_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *FilterCondition) GetItemType() string {
	if x, ok := x.GetValue().(*FilterCondition_ItemType); ok {
		return x.ItemType
	}
	return ""
}

type isFilterCondition_Value interface {
	isFilterCondition_Value()
}

type FilterCondition_ItemType struct {
	// item_type is the type of item to filter by.
	ItemType string `protobuf:"bytes,1,opt,name=item_type,json=itemType,proto3,oneof"`
}

func (*FilterCondition_ItemType) isFilterCondition_Value() {}

type BeginScanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// store_id is a globally unique Store ID, which can be looked up from the
	// console or CLI.
	StoreId uint64 `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	// filter_condition is a set of conditions to filter the scan result by.
	// If no conditions are provided, all items in the store will be returned.
	// Filter conditions are combined with OR.
	FilterCondition []*FilterCondition `protobuf:"bytes,2,rep,name=filter_condition,json=filterCondition,proto3" json:"filter_condition,omitempty"`
	// limit is the maximum number of items to return. If this is not specified or
	// set to 0, it will default to unlimited. Fewer items than the limit may be
	// returned even if there are more items to get - make sure to check
	// token.can_continue.
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// segmentation_params is used to enable parallelization of the list operation.
	// This is useful for bulk processing of large stores.
	// If this field is set then the list operation will be split into segment_params.total_segments
	// and this request will process the segment defined in segmentation_params.segment_index.
	// See SegmentationParams for more details.
	// Please be warned that parallelization will result on increase throughput to your store
	// which may result in throttling.
	SegmentationParams *SegmentationParams `protobuf:"bytes,4,opt,name=segmentation_params,json=segmentationParams,proto3" json:"segmentation_params,omitempty"`
	// schema_version_id refers to the item version to return.
	// If the store's schema does not have this version, the operation
	// will error with SchemaVersionNotFound error. You should not have to
	// set this manually as your generated SDK should know its schema version
	// and wire this in for you.
	SchemaVersionId uint32 `protobuf:"varint,5,opt,name=schema_version_id,json=schemaVersionId,proto3" json:"schema_version_id,omitempty"`
}

func (x *BeginScanRequest) Reset() {
	*x = BeginScanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_scan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginScanRequest) ProtoMessage() {}

func (x *BeginScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_scan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginScanRequest.ProtoReflect.Descriptor instead.
func (*BeginScanRequest) Descriptor() ([]byte, []int) {
	return file_db_scan_proto_rawDescGZIP(), []int{1}
}

func (x *BeginScanRequest) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *BeginScanRequest) GetFilterCondition() []*FilterCondition {
	if x != nil {
		return x.FilterCondition
	}
	return nil
}

func (x *BeginScanRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *BeginScanRequest) GetSegmentationParams() *SegmentationParams {
	if x != nil {
		return x.SegmentationParams
	}
	return nil
}

func (x *BeginScanRequest) GetSchemaVersionId() uint32 {
	if x != nil {
		return x.SchemaVersionId
	}
	return 0
}

type SegmentationParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// total_segments is used to determine the number of segments the store will be split into.
	// The segment_index field will then be used to determine which segment to process.
	// For example, if total_segments is set to 5 then the store will be split into 5 segments
	// and by setting segment_index to 0, 1, 2, 3, or 4 the request will process the corresponding
	// segment.
	TotalSegments uint32 `protobuf:"varint,5,opt,name=total_segments,json=totalSegments,proto3" json:"total_segments,omitempty"`
	// segment_index is used to determine which segment of the store to process
	// with this request.
	// Segments are zero-indexed so the value of segment_index must be less than total_segments.
	SegmentIndex uint32 `protobuf:"varint,6,opt,name=segment_index,json=segmentIndex,proto3" json:"segment_index,omitempty"`
}

func (x *SegmentationParams) Reset() {
	*x = SegmentationParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_scan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmentationParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmentationParams) ProtoMessage() {}

func (x *SegmentationParams) ProtoReflect() protoreflect.Message {
	mi := &file_db_scan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmentationParams.ProtoReflect.Descriptor instead.
func (*SegmentationParams) Descriptor() ([]byte, []int) {
	return file_db_scan_proto_rawDescGZIP(), []int{2}
}

func (x *SegmentationParams) GetTotalSegments() uint32 {
	if x != nil {
		return x.TotalSegments
	}
	return 0
}

func (x *SegmentationParams) GetSegmentIndex() uint32 {
	if x != nil {
		return x.SegmentIndex
	}
	return 0
}

var File_db_scan_proto protoreflect.FileDescriptor

var file_db_scan_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x62, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x09, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0e, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x22, 0x98, 0x02, 0x0a, 0x10, 0x42,
	0x65, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x49, 0x64, 0x12, 0x46, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x4f, 0x0a, 0x13, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x12, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x32, 0x0a, 0x11, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xdb, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x32, 0x0a, 0x0e,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x0b, 0xba, 0x48, 0x08, 0x2a, 0x06, 0x18, 0xc0, 0x84, 0x3d, 0x28,
	0x01, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x6c, 0xba, 0x48, 0x69, 0x1a, 0x67, 0x0a, 0x0b, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2e, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65,
	0x20, 0x6c, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x61, 0x6e, 0x20, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e,
	0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x20, 0x3c, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x8a, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x42, 0x09, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x62, 0xa2, 0x02, 0x03, 0x53, 0x44, 0x58, 0xaa,
	0x02, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x44, 0x62, 0xca, 0x02, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x5c, 0x44, 0x62, 0xe2, 0x02, 0x16, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x5c, 0x44, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x3a, 0x3a, 0x44, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_scan_proto_rawDescOnce sync.Once
	file_db_scan_proto_rawDescData = file_db_scan_proto_rawDesc
)

func file_db_scan_proto_rawDescGZIP() []byte {
	file_db_scan_proto_rawDescOnce.Do(func() {
		file_db_scan_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_scan_proto_rawDescData)
	})
	return file_db_scan_proto_rawDescData
}

var file_db_scan_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_db_scan_proto_goTypes = []interface{}{
	(*FilterCondition)(nil),    // 0: stately.db.FilterCondition
	(*BeginScanRequest)(nil),   // 1: stately.db.BeginScanRequest
	(*SegmentationParams)(nil), // 2: stately.db.SegmentationParams
}
var file_db_scan_proto_depIdxs = []int32{
	0, // 0: stately.db.BeginScanRequest.filter_condition:type_name -> stately.db.FilterCondition
	2, // 1: stately.db.BeginScanRequest.segmentation_params:type_name -> stately.db.SegmentationParams
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_db_scan_proto_init() }
func file_db_scan_proto_init() {
	if File_db_scan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_scan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterCondition); i {
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
		file_db_scan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginScanRequest); i {
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
		file_db_scan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmentationParams); i {
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
	file_db_scan_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FilterCondition_ItemType)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_db_scan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_scan_proto_goTypes,
		DependencyIndexes: file_db_scan_proto_depIdxs,
		MessageInfos:      file_db_scan_proto_msgTypes,
	}.Build()
	File_db_scan_proto = out.File
	file_db_scan_proto_rawDesc = nil
	file_db_scan_proto_goTypes = nil
	file_db_scan_proto_depIdxs = nil
}
