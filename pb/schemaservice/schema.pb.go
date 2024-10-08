// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: schemaservice/schema.proto

package schemaservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SchemaModel is the publicly exposed representation of a StatelyDB Schema.
// Each SchemaModel is fully self-contained and includes all nested messages and enums
// within the file_descriptor field along with other supporting metadata fields.
type SchemaModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// file_descriptor is the representation of the complete Schema in Protobuf as generated by the Stately Schema CLI.
	// All nested messages and enums must be defined in the this one file_descriptor so it is fully self-contained.
	// file_descriptor must be a valid FileDescriptorProto which contains no imports.
	FileDescriptor *descriptorpb.FileDescriptorProto `protobuf:"bytes,1,opt,name=file_descriptor,json=fileDescriptor,proto3" json:"file_descriptor,omitempty"`
	// last_modified_at_micros is the time at which this Schema was last modified,
	// as a Unix microsecond timestamp.
	LastModifiedAtMicros uint64 `protobuf:"varint,2,opt,name=last_modified_at_micros,json=lastModifiedAtMicros,proto3" json:"last_modified_at_micros,omitempty"`
	// created_at_micros is the time at which this Schema was created,
	// as a Unix microsecond timestamp.
	CreatedAtMicros uint64 `protobuf:"varint,3,opt,name=created_at_micros,json=createdAtMicros,proto3" json:"created_at_micros,omitempty"`
	// formatted_schema is a human-readable formatted version of the file descriptor.
	FormattedSchema string `protobuf:"bytes,4,opt,name=formatted_schema,json=formattedSchema,proto3" json:"formatted_schema,omitempty"`
	// schema_version_id is the version of the schema this schema represents.
	SchemaVersionId uint32 `protobuf:"varint,5,opt,name=schema_version_id,json=schemaVersionId,proto3" json:"schema_version_id,omitempty"`
	// schema_id is the schema this specific version is bound to. You can lookup which
	// schema your store is bound to in the web console.
	SchemaId uint64 `protobuf:"varint,6,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
}

func (x *SchemaModel) Reset() {
	*x = SchemaModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemaservice_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchemaModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaModel) ProtoMessage() {}

func (x *SchemaModel) ProtoReflect() protoreflect.Message {
	mi := &file_schemaservice_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaModel.ProtoReflect.Descriptor instead.
func (*SchemaModel) Descriptor() ([]byte, []int) {
	return file_schemaservice_schema_proto_rawDescGZIP(), []int{0}
}

func (x *SchemaModel) GetFileDescriptor() *descriptorpb.FileDescriptorProto {
	if x != nil {
		return x.FileDescriptor
	}
	return nil
}

func (x *SchemaModel) GetLastModifiedAtMicros() uint64 {
	if x != nil {
		return x.LastModifiedAtMicros
	}
	return 0
}

func (x *SchemaModel) GetCreatedAtMicros() uint64 {
	if x != nil {
		return x.CreatedAtMicros
	}
	return 0
}

func (x *SchemaModel) GetFormattedSchema() string {
	if x != nil {
		return x.FormattedSchema
	}
	return ""
}

func (x *SchemaModel) GetSchemaVersionId() uint32 {
	if x != nil {
		return x.SchemaVersionId
	}
	return 0
}

func (x *SchemaModel) GetSchemaId() uint64 {
	if x != nil {
		return x.SchemaId
	}
	return 0
}

var File_schemaservice_schema_proto protoreflect.FileDescriptor

var file_schemaservice_schema_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x02, 0x0a, 0x0b, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x55, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0e, 0x66, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x3d, 0x0a, 0x17,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x41, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x32, 0x0a, 0x11, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12,
	0x31, 0x0a, 0x10, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x0f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x42, 0xce, 0x01, 0x0a, 0x19,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0b, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa,
	0x02, 0x15, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x15, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c,
	0x79, 0x5c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2,
	0x02, 0x21, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x5c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x3a, 0x3a, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schemaservice_schema_proto_rawDescOnce sync.Once
	file_schemaservice_schema_proto_rawDescData = file_schemaservice_schema_proto_rawDesc
)

func file_schemaservice_schema_proto_rawDescGZIP() []byte {
	file_schemaservice_schema_proto_rawDescOnce.Do(func() {
		file_schemaservice_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schemaservice_schema_proto_rawDescData)
	})
	return file_schemaservice_schema_proto_rawDescData
}

var file_schemaservice_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schemaservice_schema_proto_goTypes = []interface{}{
	(*SchemaModel)(nil),                      // 0: stately.schemaservice.SchemaModel
	(*descriptorpb.FileDescriptorProto)(nil), // 1: google.protobuf.FileDescriptorProto
}
var file_schemaservice_schema_proto_depIdxs = []int32{
	1, // 0: stately.schemaservice.SchemaModel.file_descriptor:type_name -> google.protobuf.FileDescriptorProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_schemaservice_schema_proto_init() }
func file_schemaservice_schema_proto_init() {
	if File_schemaservice_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schemaservice_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchemaModel); i {
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
			RawDescriptor: file_schemaservice_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schemaservice_schema_proto_goTypes,
		DependencyIndexes: file_schemaservice_schema_proto_depIdxs,
		MessageInfos:      file_schemaservice_schema_proto_msgTypes,
	}.Build()
	File_schemaservice_schema_proto = out.File
	file_schemaservice_schema_proto_rawDesc = nil
	file_schemaservice_schema_proto_goTypes = nil
	file_schemaservice_schema_proto_depIdxs = nil
}
