// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: dbmanagement/config_item.proto

package dbmanagement

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	schemamodel "github.com/StatelyCloud/stately/schema/lib/schemamodel"
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

// ItemTypeConfig is the specific configuration of a group within a store.
// Note: This is primarily derived from item schema.
type ItemTypeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// item_type is the item type (last segment of a key path) that this configuration is for.
	// The item_type "*" is the default configuration used by un-configured item types.
	ItemType string `protobuf:"bytes,1,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	// key_configs is the configuration of keys for this item type.
	KeyConfigs []*KeyConfig `protobuf:"bytes,2,rep,name=key_configs,json=keyConfigs,proto3" json:"key_configs,omitempty"`
	// ttl_config is the configuration for TimeToLive.
	TtlConfig *TTLConfig `protobuf:"bytes,3,opt,name=ttl_config,json=ttlConfig,proto3" json:"ttl_config,omitempty"`
	// indexes are configurations for optional local indexes.
	Indexes []*GroupLocalIndexConfig `protobuf:"bytes,4,rep,name=indexes,proto3" json:"indexes,omitempty"`
}

func (x *ItemTypeConfig) Reset() {
	*x = ItemTypeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmanagement_config_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemTypeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemTypeConfig) ProtoMessage() {}

func (x *ItemTypeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_dbmanagement_config_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemTypeConfig.ProtoReflect.Descriptor instead.
func (*ItemTypeConfig) Descriptor() ([]byte, []int) {
	return file_dbmanagement_config_item_proto_rawDescGZIP(), []int{0}
}

func (x *ItemTypeConfig) GetItemType() string {
	if x != nil {
		return x.ItemType
	}
	return ""
}

func (x *ItemTypeConfig) GetKeyConfigs() []*KeyConfig {
	if x != nil {
		return x.KeyConfigs
	}
	return nil
}

func (x *ItemTypeConfig) GetTtlConfig() *TTLConfig {
	if x != nil {
		return x.TtlConfig
	}
	return nil
}

func (x *ItemTypeConfig) GetIndexes() []*GroupLocalIndexConfig {
	if x != nil {
		return x.Indexes
	}
	return nil
}

// TTLConfig is the configuration an item uses for TimeToLive.
type TTLConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source          schemamodel.Ttl_TtlSource     `protobuf:"varint,1,opt,name=source,proto3,enum=stately.schemamodel.Ttl_TtlSource" json:"source,omitempty"`
	Path            *FieldPath                    `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	DurationSeconds uint64                        `protobuf:"varint,3,opt,name=duration_seconds,json=durationSeconds,proto3" json:"duration_seconds,omitempty"`
	InterpretAs     schemamodel.NumberInterpretAs `protobuf:"varint,4,opt,name=interpret_as,json=interpretAs,proto3,enum=stately.schemamodel.NumberInterpretAs" json:"interpret_as,omitempty"`
}

func (x *TTLConfig) Reset() {
	*x = TTLConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmanagement_config_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TTLConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TTLConfig) ProtoMessage() {}

func (x *TTLConfig) ProtoReflect() protoreflect.Message {
	mi := &file_dbmanagement_config_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TTLConfig.ProtoReflect.Descriptor instead.
func (*TTLConfig) Descriptor() ([]byte, []int) {
	return file_dbmanagement_config_item_proto_rawDescGZIP(), []int{1}
}

func (x *TTLConfig) GetSource() schemamodel.Ttl_TtlSource {
	if x != nil {
		return x.Source
	}
	return schemamodel.Ttl_TtlSource(0)
}

func (x *TTLConfig) GetPath() *FieldPath {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *TTLConfig) GetDurationSeconds() uint64 {
	if x != nil {
		return x.DurationSeconds
	}
	return 0
}

func (x *TTLConfig) GetInterpretAs() schemamodel.NumberInterpretAs {
	if x != nil {
		return x.InterpretAs
	}
	return schemamodel.NumberInterpretAs(0)
}

// KeyConfig is the configuration of an item Key. An item may have one or more of these for aliasing.
type KeyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyPath []*NamedFieldPathSegment `protobuf:"bytes,1,rep,name=key_path,json=keyPath,proto3" json:"key_path,omitempty"`
}

func (x *KeyConfig) Reset() {
	*x = KeyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmanagement_config_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyConfig) ProtoMessage() {}

func (x *KeyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_dbmanagement_config_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyConfig.ProtoReflect.Descriptor instead.
func (*KeyConfig) Descriptor() ([]byte, []int) {
	return file_dbmanagement_config_item_proto_rawDescGZIP(), []int{2}
}

func (x *KeyConfig) GetKeyPath() []*NamedFieldPathSegment {
	if x != nil {
		return x.KeyPath
	}
	return nil
}

// GroupLocalIndexConfig is the configuration of an index.
type GroupLocalIndexConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of index.
	GroupLocalIndex uint32 `protobuf:"varint,1,opt,name=group_local_index,json=groupLocalIndex,proto3" json:"group_local_index,omitempty"`
	// The fields that are indexed in descending order.
	Path []*NamedFieldPathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *GroupLocalIndexConfig) Reset() {
	*x = GroupLocalIndexConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmanagement_config_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupLocalIndexConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupLocalIndexConfig) ProtoMessage() {}

func (x *GroupLocalIndexConfig) ProtoReflect() protoreflect.Message {
	mi := &file_dbmanagement_config_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupLocalIndexConfig.ProtoReflect.Descriptor instead.
func (*GroupLocalIndexConfig) Descriptor() ([]byte, []int) {
	return file_dbmanagement_config_item_proto_rawDescGZIP(), []int{3}
}

func (x *GroupLocalIndexConfig) GetGroupLocalIndex() uint32 {
	if x != nil {
		return x.GroupLocalIndex
	}
	return 0
}

func (x *GroupLocalIndexConfig) GetPath() []*NamedFieldPathSegment {
	if x != nil {
		return x.Path
	}
	return nil
}

var File_dbmanagement_config_item_proto protoreflect.FileDescriptor

var file_dbmanagement_config_item_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x64, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02,
	0x0a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x23, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x69, 0x74, 0x65,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xba, 0x48, 0x05,
	0x92, 0x01, 0x02, 0x10, 0x05, 0x52, 0x0a, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x74, 0x74, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e,
	0x64, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x54, 0x4c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x74, 0x74, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x4f, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x10, 0x04, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x65, 0x73, 0x22, 0xad, 0x03, 0x0a, 0x09, 0x54, 0x54, 0x4c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3a, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x74, 0x6c, 0x2e, 0x54, 0x74, 0x6c, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x49, 0x0a, 0x0c,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x5f, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x41, 0x73, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x65, 0x74, 0x41, 0x73, 0x3a, 0xb8, 0x01, 0xba, 0x48, 0xb4, 0x01, 0x1a, 0xb1,
	0x01, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x45, 0x69, 0x66, 0x20, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x20, 0x69, 0x73, 0x20, 0x54, 0x54, 0x4c, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x41,
	0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x2c, 0x20, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x20, 0x69, 0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x1a, 0x54, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x3d, 0x3d, 0x20, 0x33, 0x20, 0x26,
	0x26, 0x20, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x29,
	0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20,
	0x21, 0x3d, 0x20, 0x33, 0x20, 0x26, 0x26, 0x20, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x29, 0x22, 0x5b, 0x0a, 0x09, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x4e, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x8d, 0x01, 0x0a, 0x15, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x33, 0x0a, 0x11, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xba, 0x48, 0x04, 0x2a, 0x02, 0x18, 0x04, 0x52, 0x0f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x3f,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x64, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61,
	0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42,
	0xcc, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e,
	0x64, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0f, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x70, 0x62, 0x2f, 0x64, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0xa2,
	0x02, 0x03, 0x53, 0x44, 0x58, 0xaa, 0x02, 0x14, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e,
	0x44, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x14, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x5c, 0x44, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0xe2, 0x02, 0x20, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x5c, 0x44, 0x62,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79,
	0x3a, 0x3a, 0x44, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dbmanagement_config_item_proto_rawDescOnce sync.Once
	file_dbmanagement_config_item_proto_rawDescData = file_dbmanagement_config_item_proto_rawDesc
)

func file_dbmanagement_config_item_proto_rawDescGZIP() []byte {
	file_dbmanagement_config_item_proto_rawDescOnce.Do(func() {
		file_dbmanagement_config_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_dbmanagement_config_item_proto_rawDescData)
	})
	return file_dbmanagement_config_item_proto_rawDescData
}

var file_dbmanagement_config_item_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dbmanagement_config_item_proto_goTypes = []interface{}{
	(*ItemTypeConfig)(nil),             // 0: stately.dbmanagement.ItemTypeConfig
	(*TTLConfig)(nil),                  // 1: stately.dbmanagement.TTLConfig
	(*KeyConfig)(nil),                  // 2: stately.dbmanagement.KeyConfig
	(*GroupLocalIndexConfig)(nil),      // 3: stately.dbmanagement.GroupLocalIndexConfig
	(schemamodel.Ttl_TtlSource)(0),     // 4: stately.schemamodel.Ttl.TtlSource
	(*FieldPath)(nil),                  // 5: stately.dbmanagement.FieldPath
	(schemamodel.NumberInterpretAs)(0), // 6: stately.schemamodel.NumberInterpretAs
	(*NamedFieldPathSegment)(nil),      // 7: stately.dbmanagement.NamedFieldPathSegment
}
var file_dbmanagement_config_item_proto_depIdxs = []int32{
	2, // 0: stately.dbmanagement.ItemTypeConfig.key_configs:type_name -> stately.dbmanagement.KeyConfig
	1, // 1: stately.dbmanagement.ItemTypeConfig.ttl_config:type_name -> stately.dbmanagement.TTLConfig
	3, // 2: stately.dbmanagement.ItemTypeConfig.indexes:type_name -> stately.dbmanagement.GroupLocalIndexConfig
	4, // 3: stately.dbmanagement.TTLConfig.source:type_name -> stately.schemamodel.Ttl.TtlSource
	5, // 4: stately.dbmanagement.TTLConfig.path:type_name -> stately.dbmanagement.FieldPath
	6, // 5: stately.dbmanagement.TTLConfig.interpret_as:type_name -> stately.schemamodel.NumberInterpretAs
	7, // 6: stately.dbmanagement.KeyConfig.key_path:type_name -> stately.dbmanagement.NamedFieldPathSegment
	7, // 7: stately.dbmanagement.GroupLocalIndexConfig.path:type_name -> stately.dbmanagement.NamedFieldPathSegment
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_dbmanagement_config_item_proto_init() }
func file_dbmanagement_config_item_proto_init() {
	if File_dbmanagement_config_item_proto != nil {
		return
	}
	file_dbmanagement_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dbmanagement_config_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemTypeConfig); i {
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
		file_dbmanagement_config_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TTLConfig); i {
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
		file_dbmanagement_config_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyConfig); i {
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
		file_dbmanagement_config_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupLocalIndexConfig); i {
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
			RawDescriptor: file_dbmanagement_config_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dbmanagement_config_item_proto_goTypes,
		DependencyIndexes: file_dbmanagement_config_item_proto_depIdxs,
		MessageInfos:      file_dbmanagement_config_item_proto_msgTypes,
	}.Build()
	File_dbmanagement_config_item_proto = out.File
	file_dbmanagement_config_item_proto_rawDesc = nil
	file_dbmanagement_config_item_proto_goTypes = nil
	file_dbmanagement_config_item_proto_depIdxs = nil
}
