// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: db/list.proto

package db

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
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
	return file_db_list_proto_enumTypes[0].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_db_list_proto_enumTypes[0]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_db_list_proto_rawDescGZIP(), []int{0}
}

type Operator int32

const (
	Operator_OPERATOR_UNSPECIFIED Operator = 0 // This is the default
	// The key must be greater than the specified value based on lexicographic ordering.
	Operator_OPERATOR_GREATER_THAN Operator = 4
	// The key must be greater than or equal to the specified value based on lexicographic ordering.
	Operator_OPERATOR_GREATER_THAN_OR_EQUAL Operator = 5
	// The key must be less than the specified value based on lexicographic ordering.
	Operator_OPERATOR_LESS_THAN Operator = 6
	// The key must be less than or equal to the specified value based on lexicographic ordering.
	Operator_OPERATOR_LESS_THAN_OR_EQUAL Operator = 7
)

// Enum value maps for Operator.
var (
	Operator_name = map[int32]string{
		0: "OPERATOR_UNSPECIFIED",
		4: "OPERATOR_GREATER_THAN",
		5: "OPERATOR_GREATER_THAN_OR_EQUAL",
		6: "OPERATOR_LESS_THAN",
		7: "OPERATOR_LESS_THAN_OR_EQUAL",
	}
	Operator_value = map[string]int32{
		"OPERATOR_UNSPECIFIED":           0,
		"OPERATOR_GREATER_THAN":          4,
		"OPERATOR_GREATER_THAN_OR_EQUAL": 5,
		"OPERATOR_LESS_THAN":             6,
		"OPERATOR_LESS_THAN_OR_EQUAL":    7,
	}
)

func (x Operator) Enum() *Operator {
	p := new(Operator)
	*p = x
	return p
}

func (x Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_db_list_proto_enumTypes[1].Descriptor()
}

func (Operator) Type() protoreflect.EnumType {
	return &file_db_list_proto_enumTypes[1]
}

func (x Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operator.Descriptor instead.
func (Operator) EnumDescriptor() ([]byte, []int) {
	return file_db_list_proto_rawDescGZIP(), []int{1}
}

type BeginListRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
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
	SortProperty SortableProperty `protobuf:"varint,5,opt,name=sort_property,json=sortProperty,proto3,enum=stately.db.SortableProperty" json:"sort_property,omitempty"`
	// sort_direction is the direction to sort the results in. If this is not set,
	// we will sort in ascending order.
	SortDirection SortDirection `protobuf:"varint,6,opt,name=sort_direction,json=sortDirection,proto3,enum=stately.db.SortDirection" json:"sort_direction,omitempty"`
	// schema_version_id is the version of the store's schema to use to interpret
	// items. If there is no version with this ID, the operation will error with
	// SchemaVersionNotFound error. You should not have to set this manually as
	// your generated SDK should know its schema version and wire this in for you.
	SchemaVersionId uint32 `protobuf:"varint,7,opt,name=schema_version_id,json=schemaVersionId,proto3" json:"schema_version_id,omitempty"`
	// schema_id refers to the schema to use for this operation.
	// If the store_id does not have a schema with this ID, the operation will
	// error with SchemaNotFound error. You should not have to set this manually
	// as your generated SDK should know its schema and wire this in for you.
	SchemaId uint64 `protobuf:"varint,8,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"` // [(buf.validate.field).required = true]; (after clients have been regen'd and updated)
	// filter_conditions are a set of conditions to filter the list result by.
	// If no conditions are provided, all items in the store will be returned.
	// Filter conditions are combined with OR.
	FilterConditions []*FilterCondition `protobuf:"bytes,9,rep,name=filter_conditions,json=filterConditions,proto3" json:"filter_conditions,omitempty"`
	// key_conditions are a set of conditions to apply to the list operation.
	// Wherever possible, Stately will apply these key conditions at the DB layer
	// to optimize the list operation cost.
	//
	// A maximum of two key conditions are allowed: one with a GREATER_THAN (or equal to)
	// operator and one with a LESS_THAN (or equal to) operator. Together these amount to
	// a "between" condition on the key path.
	//
	// If these conditions are provided they must share the same prefix as the
	// key_path_prefix. For example, the following is valid:
	//
	//	key_path_prefix: "/group-:groupID/namespace"
	//	key_conditions:
	//	  - key_path: "/group-:groupID/namespace-44"
	//	    operator: GREATER_THAN_OR_EQUAL
	//	  - key_path: "/group-:groupID/namespace-100"
	//	    operator: LESS_THAN_OR_EQUAL
	//
	// A key_path_prefix of "/group-:groupID" would also be valid above, as the prefix is shared
	// with the key conditions.
	//
	// The following is NOT valid because the key_path_prefix does not
	// share the same prefix as the key conditions:
	//
	//	key_path_prefix: "/group-:groupID/namespace"
	//	key_conditions:
	//	  - key_path: "/group-:groupID/beatles-1984"
	//	    operator: GREATER_THAN_OR_EQUAL
	KeyConditions []*KeyCondition `protobuf:"bytes,10,rep,name=key_conditions,json=keyConditions,proto3" json:"key_conditions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BeginListRequest) Reset() {
	*x = BeginListRequest{}
	mi := &file_db_list_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BeginListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginListRequest) ProtoMessage() {}

func (x *BeginListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_list_proto_msgTypes[0]
	if x != nil {
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
	return file_db_list_proto_rawDescGZIP(), []int{0}
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

func (x *BeginListRequest) GetSchemaVersionId() uint32 {
	if x != nil {
		return x.SchemaVersionId
	}
	return 0
}

func (x *BeginListRequest) GetSchemaId() uint64 {
	if x != nil {
		return x.SchemaId
	}
	return 0
}

func (x *BeginListRequest) GetFilterConditions() []*FilterCondition {
	if x != nil {
		return x.FilterConditions
	}
	return nil
}

func (x *BeginListRequest) GetKeyConditions() []*KeyCondition {
	if x != nil {
		return x.KeyConditions
	}
	return nil
}

// These are stream messages, so multiple responses may be sent.
type ListResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Response:
	//
	//	*ListResponse_Result
	//	*ListResponse_Finished
	Response      isListResponse_Response `protobuf_oneof:"response"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_db_list_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_list_proto_msgTypes[1]
	if x != nil {
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
	return file_db_list_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetResponse() isListResponse_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ListResponse) GetResult() *ListPartialResult {
	if x != nil {
		if x, ok := x.Response.(*ListResponse_Result); ok {
			return x.Result
		}
	}
	return nil
}

func (x *ListResponse) GetFinished() *ListFinished {
	if x != nil {
		if x, ok := x.Response.(*ListResponse_Finished); ok {
			return x.Finished
		}
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// results is a list that contains one entry for each Item that was found.
	Items         []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPartialResult) Reset() {
	*x = ListPartialResult{}
	mi := &file_db_list_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPartialResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPartialResult) ProtoMessage() {}

func (x *ListPartialResult) ProtoReflect() protoreflect.Message {
	mi := &file_db_list_proto_msgTypes[2]
	if x != nil {
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
	return file_db_list_proto_rawDescGZIP(), []int{2}
}

func (x *ListPartialResult) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListFinished struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// token is always set and represents an updated list continuation token that
	// can be used in subsequent calls to ContinueList or SyncList.
	Token         *ListToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFinished) Reset() {
	*x = ListFinished{}
	mi := &file_db_list_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFinished) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFinished) ProtoMessage() {}

func (x *ListFinished) ProtoReflect() protoreflect.Message {
	mi := &file_db_list_proto_msgTypes[3]
	if x != nil {
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
	return file_db_list_proto_rawDescGZIP(), []int{3}
}

func (x *ListFinished) GetToken() *ListToken {
	if x != nil {
		return x.Token
	}
	return nil
}

// A KeyCondition is an additional constraint to be applied to the list
// operation. It is used to filter the results based on a specific key path
// and an operator.
// Wherever possible, stately will apply these key conditions at the DB layer
// to optimize the list operation latency and cost.
// Key conditions may be combined with a key_path_prefix to further
// optimize the list operation. HOWEVER Key conditions must share the
// same prefix as the key_path_prefix.
type KeyCondition struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// key_path is a valid key prefix (or full key) used to filter or optimize the list
	// operation based on the operator specified below.
	KeyPath string `protobuf:"bytes,1,opt,name=key_path,json=keyPath,proto3" json:"key_path,omitempty"`
	// Operator indicates how to apply key_path condition to the list operation.
	// Valid options are:
	// - GREATER_THAN: key_path must be greater than the specified value
	// - GREATER_THAN_OR_EQUAL: key_path must be greater than or equal to the specified value
	// - LESS_THAN: key_path must be less than the specified value
	// - LESS_THAN_OR_EQUAL: key_path must be less than or equal to the specified value
	//
	// Note: Operators are strictly evaluated they do not change meaning based on sort direction.
	// For example, regardless of sort direction, a GREATER_THAN operator
	// will still mean that a key_path must be greater than the specified value in order
	// to be included in the result set.
	Operator      Operator `protobuf:"varint,2,opt,name=operator,proto3,enum=stately.db.Operator" json:"operator,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KeyCondition) Reset() {
	*x = KeyCondition{}
	mi := &file_db_list_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyCondition) ProtoMessage() {}

func (x *KeyCondition) ProtoReflect() protoreflect.Message {
	mi := &file_db_list_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyCondition.ProtoReflect.Descriptor instead.
func (*KeyCondition) Descriptor() ([]byte, []int) {
	return file_db_list_proto_rawDescGZIP(), []int{4}
}

func (x *KeyCondition) GetKeyPath() string {
	if x != nil {
		return x.KeyPath
	}
	return ""
}

func (x *KeyCondition) GetOperator() Operator {
	if x != nil {
		return x.Operator
	}
	return Operator_OPERATOR_UNSPECIFIED
}

var File_db_list_proto protoreflect.FileDescriptor

const file_db_list_proto_rawDesc = "" +
	"\n" +
	"\rdb/list.proto\x12\n" +
	"stately.db\x1a\x1bbuf/validate/validate.proto\x1a\rdb/item.proto\x1a\x16db/item_property.proto\x1a\x15db/list_filters.proto\x1a\x13db/list_token.proto\"\x89\x04\n" +
	"\x10BeginListRequest\x12!\n" +
	"\bstore_id\x18\x01 \x01(\x04B\x06\xbaH\x03\xc8\x01\x01R\astoreId\x12.\n" +
	"\x0fkey_path_prefix\x18\x02 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\rkeyPathPrefix\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\rR\x05limit\x12\x1f\n" +
	"\vallow_stale\x18\x04 \x01(\bR\n" +
	"allowStale\x12A\n" +
	"\rsort_property\x18\x05 \x01(\x0e2\x1c.stately.db.SortablePropertyR\fsortProperty\x12@\n" +
	"\x0esort_direction\x18\x06 \x01(\x0e2\x19.stately.db.SortDirectionR\rsortDirection\x122\n" +
	"\x11schema_version_id\x18\a \x01(\rB\x06\xbaH\x03\xc8\x01\x01R\x0fschemaVersionId\x12\x1b\n" +
	"\tschema_id\x18\b \x01(\x04R\bschemaId\x12H\n" +
	"\x11filter_conditions\x18\t \x03(\v2\x1b.stately.db.FilterConditionR\x10filterConditions\x12K\n" +
	"\x0ekey_conditions\x18\n" +
	" \x03(\v2\x18.stately.db.KeyConditionB\n" +
	"\xbaH\a\x92\x01\x04\b\x00\x10\x02R\rkeyConditions\"\x92\x01\n" +
	"\fListResponse\x127\n" +
	"\x06result\x18\x01 \x01(\v2\x1d.stately.db.ListPartialResultH\x00R\x06result\x126\n" +
	"\bfinished\x18\x02 \x01(\v2\x18.stately.db.ListFinishedH\x00R\bfinishedB\x11\n" +
	"\bresponse\x12\x05\xbaH\x02\b\x01\"C\n" +
	"\x11ListPartialResult\x12.\n" +
	"\x05items\x18\x01 \x03(\v2\x10.stately.db.ItemB\x06\xbaH\x03\xc8\x01\x01R\x05items\"C\n" +
	"\fListFinished\x123\n" +
	"\x05token\x18\x01 \x01(\v2\x15.stately.db.ListTokenB\x06\xbaH\x03\xc8\x01\x01R\x05token\"k\n" +
	"\fKeyCondition\x12!\n" +
	"\bkey_path\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\akeyPath\x128\n" +
	"\boperator\x18\x02 \x01(\x0e2\x14.stately.db.OperatorB\x06\xbaH\x03\xc8\x01\x01R\boperator*8\n" +
	"\rSortDirection\x12\x12\n" +
	"\x0eSORT_ASCENDING\x10\x00\x12\x13\n" +
	"\x0fSORT_DESCENDING\x10\x01*\x9c\x01\n" +
	"\bOperator\x12\x18\n" +
	"\x14OPERATOR_UNSPECIFIED\x10\x00\x12\x19\n" +
	"\x15OPERATOR_GREATER_THAN\x10\x04\x12\"\n" +
	"\x1eOPERATOR_GREATER_THAN_OR_EQUAL\x10\x05\x12\x16\n" +
	"\x12OPERATOR_LESS_THAN\x10\x06\x12\x1f\n" +
	"\x1bOPERATOR_LESS_THAN_OR_EQUAL\x10\aB\x8a\x01\n" +
	"\x0ecom.stately.dbB\tListProtoP\x01Z$github.com/StatelyCloud/go-sdk/pb/db\xa2\x02\x03SDX\xaa\x02\n" +
	"Stately.Db\xca\x02\n" +
	"Stately\\Db\xe2\x02\x16Stately\\Db\\GPBMetadata\xea\x02\vStately::Dbb\x06proto3"

var (
	file_db_list_proto_rawDescOnce sync.Once
	file_db_list_proto_rawDescData []byte
)

func file_db_list_proto_rawDescGZIP() []byte {
	file_db_list_proto_rawDescOnce.Do(func() {
		file_db_list_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_db_list_proto_rawDesc), len(file_db_list_proto_rawDesc)))
	})
	return file_db_list_proto_rawDescData
}

var file_db_list_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_db_list_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_db_list_proto_goTypes = []any{
	(SortDirection)(0),        // 0: stately.db.SortDirection
	(Operator)(0),             // 1: stately.db.Operator
	(*BeginListRequest)(nil),  // 2: stately.db.BeginListRequest
	(*ListResponse)(nil),      // 3: stately.db.ListResponse
	(*ListPartialResult)(nil), // 4: stately.db.ListPartialResult
	(*ListFinished)(nil),      // 5: stately.db.ListFinished
	(*KeyCondition)(nil),      // 6: stately.db.KeyCondition
	(SortableProperty)(0),     // 7: stately.db.SortableProperty
	(*FilterCondition)(nil),   // 8: stately.db.FilterCondition
	(*Item)(nil),              // 9: stately.db.Item
	(*ListToken)(nil),         // 10: stately.db.ListToken
}
var file_db_list_proto_depIdxs = []int32{
	7,  // 0: stately.db.BeginListRequest.sort_property:type_name -> stately.db.SortableProperty
	0,  // 1: stately.db.BeginListRequest.sort_direction:type_name -> stately.db.SortDirection
	8,  // 2: stately.db.BeginListRequest.filter_conditions:type_name -> stately.db.FilterCondition
	6,  // 3: stately.db.BeginListRequest.key_conditions:type_name -> stately.db.KeyCondition
	4,  // 4: stately.db.ListResponse.result:type_name -> stately.db.ListPartialResult
	5,  // 5: stately.db.ListResponse.finished:type_name -> stately.db.ListFinished
	9,  // 6: stately.db.ListPartialResult.items:type_name -> stately.db.Item
	10, // 7: stately.db.ListFinished.token:type_name -> stately.db.ListToken
	1,  // 8: stately.db.KeyCondition.operator:type_name -> stately.db.Operator
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_db_list_proto_init() }
func file_db_list_proto_init() {
	if File_db_list_proto != nil {
		return
	}
	file_db_item_proto_init()
	file_db_item_property_proto_init()
	file_db_list_filters_proto_init()
	file_db_list_token_proto_init()
	file_db_list_proto_msgTypes[1].OneofWrappers = []any{
		(*ListResponse_Result)(nil),
		(*ListResponse_Finished)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_db_list_proto_rawDesc), len(file_db_list_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_list_proto_goTypes,
		DependencyIndexes: file_db_list_proto_depIdxs,
		EnumInfos:         file_db_list_proto_enumTypes,
		MessageInfos:      file_db_list_proto_msgTypes,
	}.Build()
	File_db_list_proto = out.File
	file_db_list_proto_goTypes = nil
	file_db_list_proto_depIdxs = nil
}
