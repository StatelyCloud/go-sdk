// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: schemaservice/create.proto

package schemaservice

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project_id is the project within which the schema will be created. Project
	// ID is a globally unique ID tied to a project. We do not need User ID
	// because projects exist in only a single organization.
	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// name is a user-facing, memorable name for the schema.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// description is a longer-form, user-facing explanation of what the schema is
	// used for - it is used in the console and in generated documentation. We
	// should encourage users to provide a meaningful description.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemaservice_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schemaservice_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_schemaservice_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// schema_id is a globally unique identifier for a schema. Schema IDs are
	// assigned by the system when a schema is created.
	SchemaId uint64 `protobuf:"varint,1,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemaservice_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schemaservice_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_schemaservice_create_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetSchemaId() uint64 {
	if x != nil {
		return x.SchemaId
	}
	return 0
}

var File_schemaservice_create_proto protoreflect.FileDescriptor

var file_schemaservice_create_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72,
	0x04, 0x20, 0x01, 0x28, 0x40, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x28, 0x80, 0x80, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x09, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x42,
	0xce, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x03,
	0x53, 0x53, 0x58, 0xaa, 0x02, 0x15, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x15, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x5c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0xe2, 0x02, 0x21, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x5c, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c,
	0x79, 0x3a, 0x3a, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schemaservice_create_proto_rawDescOnce sync.Once
	file_schemaservice_create_proto_rawDescData = file_schemaservice_create_proto_rawDesc
)

func file_schemaservice_create_proto_rawDescGZIP() []byte {
	file_schemaservice_create_proto_rawDescOnce.Do(func() {
		file_schemaservice_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_schemaservice_create_proto_rawDescData)
	})
	return file_schemaservice_create_proto_rawDescData
}

var file_schemaservice_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_schemaservice_create_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),  // 0: stately.schemaservice.CreateRequest
	(*CreateResponse)(nil), // 1: stately.schemaservice.CreateResponse
}
var file_schemaservice_create_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schemaservice_create_proto_init() }
func file_schemaservice_create_proto_init() {
	if File_schemaservice_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schemaservice_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_schemaservice_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_schemaservice_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schemaservice_create_proto_goTypes,
		DependencyIndexes: file_schemaservice_create_proto_depIdxs,
		MessageInfos:      file_schemaservice_create_proto_msgTypes,
	}.Build()
	File_schemaservice_create_proto = out.File
	file_schemaservice_create_proto_rawDesc = nil
	file_schemaservice_create_proto_goTypes = nil
	file_schemaservice_create_proto_depIdxs = nil
}
