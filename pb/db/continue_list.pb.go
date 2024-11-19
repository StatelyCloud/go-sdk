// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: db/continue_list.proto

package db

import (
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

// ContinueListDirection is used to indicate whether we are expanding the result
// set (paginating) forward (in the direction of the original List operation) or
// backward (in the opposite direction).
type ContinueListDirection int32

const (
	ContinueListDirection_CONTINUE_LIST_FORWARD  ContinueListDirection = 0 // this is the default
	ContinueListDirection_CONTINUE_LIST_BACKWARD ContinueListDirection = 1
)

// Enum value maps for ContinueListDirection.
var (
	ContinueListDirection_name = map[int32]string{
		0: "CONTINUE_LIST_FORWARD",
		1: "CONTINUE_LIST_BACKWARD",
	}
	ContinueListDirection_value = map[string]int32{
		"CONTINUE_LIST_FORWARD":  0,
		"CONTINUE_LIST_BACKWARD": 1,
	}
)

func (x ContinueListDirection) Enum() *ContinueListDirection {
	p := new(ContinueListDirection)
	*p = x
	return p
}

func (x ContinueListDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContinueListDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_db_continue_list_proto_enumTypes[0].Descriptor()
}

func (ContinueListDirection) Type() protoreflect.EnumType {
	return &file_db_continue_list_proto_enumTypes[0]
}

func (x ContinueListDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContinueListDirection.Descriptor instead.
func (ContinueListDirection) EnumDescriptor() ([]byte, []int) {
	return file_db_continue_list_proto_rawDescGZIP(), []int{0}
}

type ContinueListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// token_data is an opaque list continuation token returned by a previous call to
	// BeginList, ContinueList, or SyncList.
	TokenData []byte `protobuf:"bytes,1,opt,name=token_data,json=tokenData,proto3" json:"token_data,omitempty"`
	// direction indicates whether we are expanding the result set (paginating)
	// forward (in the direction of the original List operation) or backward (in
	// the opposite direction). The default is to expand forward.
	Direction ContinueListDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=stately.db.ContinueListDirection" json:"direction,omitempty"`
	// schema_version_id is the schema version id the client is using. This is
	// used to ensure the same schema version is used across each page of the list.
	// If the version the token was created with is different from the client's
	// current version, the operation will error with SchemaVersionMismatch error,
	// in which case you should start over with a fresh BeginList call.
	// TODO this will soon be required.
	SchemaVersionId uint32 `protobuf:"varint,5,opt,name=schema_version_id,json=schemaVersionId,proto3" json:"schema_version_id,omitempty"`
}

func (x *ContinueListRequest) Reset() {
	*x = ContinueListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_continue_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContinueListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContinueListRequest) ProtoMessage() {}

func (x *ContinueListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_continue_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContinueListRequest.ProtoReflect.Descriptor instead.
func (*ContinueListRequest) Descriptor() ([]byte, []int) {
	return file_db_continue_list_proto_rawDescGZIP(), []int{0}
}

func (x *ContinueListRequest) GetTokenData() []byte {
	if x != nil {
		return x.TokenData
	}
	return nil
}

func (x *ContinueListRequest) GetDirection() ContinueListDirection {
	if x != nil {
		return x.Direction
	}
	return ContinueListDirection_CONTINUE_LIST_FORWARD
}

func (x *ContinueListRequest) GetSchemaVersionId() uint32 {
	if x != nil {
		return x.SchemaVersionId
	}
	return 0
}

var File_db_continue_list_proto protoreflect.FileDescriptor

var file_db_continue_list_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c,
	0x79, 0x2e, 0x64, 0x62, 0x22, 0xa9, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0a,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79,
	0x2e, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x2a, 0x4e, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e,
	0x54, 0x49, 0x4e, 0x55, 0x45, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41,
	0x52, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45,
	0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x57, 0x41, 0x52, 0x44, 0x10, 0x01,
	0x42, 0x92, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79,
	0x2e, 0x64, 0x62, 0x42, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x62, 0xa2, 0x02,
	0x03, 0x53, 0x44, 0x58, 0xaa, 0x02, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x44,
	0x62, 0xca, 0x02, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x5c, 0x44, 0x62, 0xe2, 0x02,
	0x16, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x5c, 0x44, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c,
	0x79, 0x3a, 0x3a, 0x44, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_continue_list_proto_rawDescOnce sync.Once
	file_db_continue_list_proto_rawDescData = file_db_continue_list_proto_rawDesc
)

func file_db_continue_list_proto_rawDescGZIP() []byte {
	file_db_continue_list_proto_rawDescOnce.Do(func() {
		file_db_continue_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_continue_list_proto_rawDescData)
	})
	return file_db_continue_list_proto_rawDescData
}

var file_db_continue_list_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_db_continue_list_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_db_continue_list_proto_goTypes = []interface{}{
	(ContinueListDirection)(0),  // 0: stately.db.ContinueListDirection
	(*ContinueListRequest)(nil), // 1: stately.db.ContinueListRequest
}
var file_db_continue_list_proto_depIdxs = []int32{
	0, // 0: stately.db.ContinueListRequest.direction:type_name -> stately.db.ContinueListDirection
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_db_continue_list_proto_init() }
func file_db_continue_list_proto_init() {
	if File_db_continue_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_continue_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContinueListRequest); i {
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
			RawDescriptor: file_db_continue_list_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_continue_list_proto_goTypes,
		DependencyIndexes: file_db_continue_list_proto_depIdxs,
		EnumInfos:         file_db_continue_list_proto_enumTypes,
		MessageInfos:      file_db_continue_list_proto_msgTypes,
	}.Build()
	File_db_continue_list_proto = out.File
	file_db_continue_list_proto_rawDesc = nil
	file_db_continue_list_proto_goTypes = nil
	file_db_continue_list_proto_depIdxs = nil
}
