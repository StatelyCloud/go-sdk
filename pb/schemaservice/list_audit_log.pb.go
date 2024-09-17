// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: schemaservice/list_audit_log.proto

package schemaservice

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

type ListAuditLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// store_id is a globally unique Store ID, which can be looked up from the
	// console or CLI.
	StoreId uint64 `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	// limit is the maximum number of schema audit log entries to return.
	// If limit is not set, the default limit is 10.
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListAuditLogRequest) Reset() {
	*x = ListAuditLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemaservice_list_audit_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuditLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditLogRequest) ProtoMessage() {}

func (x *ListAuditLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schemaservice_list_audit_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuditLogRequest.ProtoReflect.Descriptor instead.
func (*ListAuditLogRequest) Descriptor() ([]byte, []int) {
	return file_schemaservice_list_audit_log_proto_rawDescGZIP(), []int{0}
}

func (x *ListAuditLogRequest) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *ListAuditLogRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListAuditLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// entries is a list of schema audit log entries sorted by recency with the most recent first.
	// If the schema exists there will always be at least one entry for it's creation.
	// If there is no schema on the provided storeID then the list will be empty.
	Entries []*SchemaAuditLogEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ListAuditLogResponse) Reset() {
	*x = ListAuditLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemaservice_list_audit_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuditLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditLogResponse) ProtoMessage() {}

func (x *ListAuditLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schemaservice_list_audit_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuditLogResponse.ProtoReflect.Descriptor instead.
func (*ListAuditLogResponse) Descriptor() ([]byte, []int) {
	return file_schemaservice_list_audit_log_proto_rawDescGZIP(), []int{1}
}

func (x *ListAuditLogResponse) GetEntries() []*SchemaAuditLogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

// SchemaAuditLogEntry is a log entry that represents a change to a Schema.
// Each time a Schema is modified, a new SchemaAuditLogEntry is created.
// allowing for a full audit log of all changes to a Schema.
// A list of many SchemaAuditLogEntry objects constitutes an audit log.
type SchemaAuditLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// schema is the full SchemaModel that was put via the update.
	// This is the full state of the schema after the update.
	Schema *SchemaModel `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	// modified_at_micros is the time at which this Schema update occurred,
	// as a Unix microsecond timestamp.
	ModifiedAtMicros uint64 `protobuf:"varint,2,opt,name=modified_at_micros,json=modifiedAtMicros,proto3" json:"modified_at_micros,omitempty"`
	// modified_by is the display name of the user who performed the update.
	ModifiedBy string `protobuf:"bytes,3,opt,name=modified_by,json=modifiedBy,proto3" json:"modified_by,omitempty"`
	// change_description is a human-readable description of the change.
	// This field is optional and entries are limited to 2000 characters.
	ChangeDescription string `protobuf:"bytes,4,opt,name=change_description,json=changeDescription,proto3" json:"change_description,omitempty"`
}

func (x *SchemaAuditLogEntry) Reset() {
	*x = SchemaAuditLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemaservice_list_audit_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchemaAuditLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaAuditLogEntry) ProtoMessage() {}

func (x *SchemaAuditLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_schemaservice_list_audit_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaAuditLogEntry.ProtoReflect.Descriptor instead.
func (*SchemaAuditLogEntry) Descriptor() ([]byte, []int) {
	return file_schemaservice_list_audit_log_proto_rawDescGZIP(), []int{2}
}

func (x *SchemaAuditLogEntry) GetSchema() *SchemaModel {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *SchemaAuditLogEntry) GetModifiedAtMicros() uint64 {
	if x != nil {
		return x.ModifiedAtMicros
	}
	return 0
}

func (x *SchemaAuditLogEntry) GetModifiedBy() string {
	if x != nil {
		return x.ModifiedBy
	}
	return ""
}

func (x *SchemaAuditLogEntry) GetChangeDescription() string {
	if x != nil {
		return x.ChangeDescription
	}
	return ""
}

var File_schemaservice_list_audit_log_proto protoreflect.FileDescriptor

var file_schemaservice_list_audit_log_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x08, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xba,
	0x48, 0x04, 0x2a, 0x02, 0x10, 0x64, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5c, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0xf1, 0x01, 0x0a, 0x13,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x42, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x34, 0x0a, 0x12, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x10, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x27, 0x0a,
	0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x42, 0x79, 0x12, 0x37, 0x0a, 0x12, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x18, 0xd0, 0x0f, 0x52, 0x11, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0xd4, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x15, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0xca, 0x02, 0x15, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x5c, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x21, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x5c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x3a, 0x3a, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schemaservice_list_audit_log_proto_rawDescOnce sync.Once
	file_schemaservice_list_audit_log_proto_rawDescData = file_schemaservice_list_audit_log_proto_rawDesc
)

func file_schemaservice_list_audit_log_proto_rawDescGZIP() []byte {
	file_schemaservice_list_audit_log_proto_rawDescOnce.Do(func() {
		file_schemaservice_list_audit_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_schemaservice_list_audit_log_proto_rawDescData)
	})
	return file_schemaservice_list_audit_log_proto_rawDescData
}

var file_schemaservice_list_audit_log_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schemaservice_list_audit_log_proto_goTypes = []interface{}{
	(*ListAuditLogRequest)(nil),  // 0: stately.schemaservice.ListAuditLogRequest
	(*ListAuditLogResponse)(nil), // 1: stately.schemaservice.ListAuditLogResponse
	(*SchemaAuditLogEntry)(nil),  // 2: stately.schemaservice.SchemaAuditLogEntry
	(*SchemaModel)(nil),          // 3: stately.schemaservice.SchemaModel
}
var file_schemaservice_list_audit_log_proto_depIdxs = []int32{
	2, // 0: stately.schemaservice.ListAuditLogResponse.entries:type_name -> stately.schemaservice.SchemaAuditLogEntry
	3, // 1: stately.schemaservice.SchemaAuditLogEntry.schema:type_name -> stately.schemaservice.SchemaModel
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_schemaservice_list_audit_log_proto_init() }
func file_schemaservice_list_audit_log_proto_init() {
	if File_schemaservice_list_audit_log_proto != nil {
		return
	}
	file_schemaservice_schema_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schemaservice_list_audit_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuditLogRequest); i {
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
		file_schemaservice_list_audit_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuditLogResponse); i {
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
		file_schemaservice_list_audit_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchemaAuditLogEntry); i {
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
			RawDescriptor: file_schemaservice_list_audit_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schemaservice_list_audit_log_proto_goTypes,
		DependencyIndexes: file_schemaservice_list_audit_log_proto_depIdxs,
		MessageInfos:      file_schemaservice_list_audit_log_proto_msgTypes,
	}.Build()
	File_schemaservice_list_audit_log_proto = out.File
	file_schemaservice_list_audit_log_proto_rawDesc = nil
	file_schemaservice_list_audit_log_proto_goTypes = nil
	file_schemaservice_list_audit_log_proto_depIdxs = nil
}