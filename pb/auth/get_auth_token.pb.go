// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: auth/get_auth_token.proto

package auth

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

type GetAuthTokenRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Identity:
	//
	//	*GetAuthTokenRequest_AccessKey
	Identity      isGetAuthTokenRequest_Identity `protobuf_oneof:"identity"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthTokenRequest) Reset() {
	*x = GetAuthTokenRequest{}
	mi := &file_auth_get_auth_token_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthTokenRequest) ProtoMessage() {}

func (x *GetAuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_get_auth_token_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthTokenRequest.ProtoReflect.Descriptor instead.
func (*GetAuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_get_auth_token_proto_rawDescGZIP(), []int{0}
}

func (x *GetAuthTokenRequest) GetIdentity() isGetAuthTokenRequest_Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *GetAuthTokenRequest) GetAccessKey() string {
	if x != nil {
		if x, ok := x.Identity.(*GetAuthTokenRequest_AccessKey); ok {
			return x.AccessKey
		}
	}
	return ""
}

type isGetAuthTokenRequest_Identity interface {
	isGetAuthTokenRequest_Identity()
}

type GetAuthTokenRequest_AccessKey struct {
	// access_key is an access key that has been added to an organization via
	// stately.user.UserService.CreateAccessKey. It is a sensitive secret used
	// for programmatic access to Stately APIs, equivalent to a
	// username+password.
	AccessKey string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3,oneof"`
}

func (*GetAuthTokenRequest_AccessKey) isGetAuthTokenRequest_Identity() {}

type GetAuthTokenResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// auth_token is the token that can be used to authenticate requests to the
	// StatelyDB API.
	AuthToken string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	// expires_in_s is the duration in seconds that this access token is
	// valid. After this time has passed, the access token will be rejected - you
	// should refresh the token via another call to GetAuthToken before this
	// happens.
	ExpiresInS    uint64 `protobuf:"varint,2,opt,name=expires_in_s,json=expiresInS,proto3" json:"expires_in_s,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthTokenResponse) Reset() {
	*x = GetAuthTokenResponse{}
	mi := &file_auth_get_auth_token_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthTokenResponse) ProtoMessage() {}

func (x *GetAuthTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_get_auth_token_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthTokenResponse.ProtoReflect.Descriptor instead.
func (*GetAuthTokenResponse) Descriptor() ([]byte, []int) {
	return file_auth_get_auth_token_proto_rawDescGZIP(), []int{1}
}

func (x *GetAuthTokenResponse) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *GetAuthTokenResponse) GetExpiresInS() uint64 {
	if x != nil {
		return x.ExpiresInS
	}
	return 0
}

var File_auth_get_auth_token_proto protoreflect.FileDescriptor

const file_auth_get_auth_token_proto_rawDesc = "" +
	"\n" +
	"\x19auth/get_auth_token.proto\x12\fstately.auth\x1a\x1bbuf/validate/validate.proto\"I\n" +
	"\x13GetAuthTokenRequest\x12\x1f\n" +
	"\n" +
	"access_key\x18\x01 \x01(\tH\x00R\taccessKeyB\x11\n" +
	"\bidentity\x12\x05\xbaH\x02\b\x01\"g\n" +
	"\x14GetAuthTokenResponse\x12%\n" +
	"\n" +
	"auth_token\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\tauthToken\x12(\n" +
	"\fexpires_in_s\x18\x02 \x01(\x04B\x06\xbaH\x03\xc8\x01\x01R\n" +
	"expiresInSB\x9e\x01\n" +
	"\x10com.stately.authB\x11GetAuthTokenProtoP\x01Z&github.com/StatelyCloud/go-sdk/pb/auth\xa2\x02\x03SAX\xaa\x02\fStately.Auth\xca\x02\fStately\\Auth\xe2\x02\x18Stately\\Auth\\GPBMetadata\xea\x02\rStately::Authb\x06proto3"

var (
	file_auth_get_auth_token_proto_rawDescOnce sync.Once
	file_auth_get_auth_token_proto_rawDescData []byte
)

func file_auth_get_auth_token_proto_rawDescGZIP() []byte {
	file_auth_get_auth_token_proto_rawDescOnce.Do(func() {
		file_auth_get_auth_token_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_auth_get_auth_token_proto_rawDesc), len(file_auth_get_auth_token_proto_rawDesc)))
	})
	return file_auth_get_auth_token_proto_rawDescData
}

var file_auth_get_auth_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_get_auth_token_proto_goTypes = []any{
	(*GetAuthTokenRequest)(nil),  // 0: stately.auth.GetAuthTokenRequest
	(*GetAuthTokenResponse)(nil), // 1: stately.auth.GetAuthTokenResponse
}
var file_auth_get_auth_token_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_get_auth_token_proto_init() }
func file_auth_get_auth_token_proto_init() {
	if File_auth_get_auth_token_proto != nil {
		return
	}
	file_auth_get_auth_token_proto_msgTypes[0].OneofWrappers = []any{
		(*GetAuthTokenRequest_AccessKey)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auth_get_auth_token_proto_rawDesc), len(file_auth_get_auth_token_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_get_auth_token_proto_goTypes,
		DependencyIndexes: file_auth_get_auth_token_proto_depIdxs,
		MessageInfos:      file_auth_get_auth_token_proto_msgTypes,
	}.Build()
	File_auth_get_auth_token_proto = out.File
	file_auth_get_auth_token_proto_goTypes = nil
	file_auth_get_auth_token_proto_depIdxs = nil
}
