// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: errors/error_details.proto

package errors

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

// StatelyErrorDetails is a message containing detailed error information.
// This is returned from the Stately API via Connect error details:
//
//	https://connectrpc.com/docs/go/errors#error-details
//
// Note: As a customer, you should not need to handle this message directly unless writing
// a custom low-level SDK. Instead, language-specific SDKs will provide a more user-friendly
// error object that wraps this message.
type StatelyErrorDetails struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// stately_code is the error code that was returned by the Stately API.
	// The full list of codes is available at https://docs.stately.cloud/api/error-codes/
	// and documentation on a specific code can be found at https://docs.stately.cloud/api/error-codes/#{stately_code}
	StatelyCode string `protobuf:"bytes,1,opt,name=stately_code,json=statelyCode,proto3" json:"stately_code,omitempty"`
	// message is a human-readable error message that can be displayed to the user that
	// provides more context about the error.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// upstream_cause is additional information about the error that can be used to help debug the error,
	// This field will only optionally be supplied by the Stately API.
	// Note: This may row over as the error is passed through multiple services.
	UpstreamCause string `protobuf:"bytes,3,opt,name=upstream_cause,json=upstreamCause,proto3" json:"upstream_cause,omitempty"`
	// The Connect/gRPC code associated with this error. This generally isn't set,
	// because the overall API response has an error code. But this can be used in
	// the case that we're returning multiple different errors, or communicating
	// errors across non-Connect APIs.
	Code          uint32 `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatelyErrorDetails) Reset() {
	*x = StatelyErrorDetails{}
	mi := &file_errors_error_details_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatelyErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatelyErrorDetails) ProtoMessage() {}

func (x *StatelyErrorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_errors_error_details_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatelyErrorDetails.ProtoReflect.Descriptor instead.
func (*StatelyErrorDetails) Descriptor() ([]byte, []int) {
	return file_errors_error_details_proto_rawDescGZIP(), []int{0}
}

func (x *StatelyErrorDetails) GetStatelyCode() string {
	if x != nil {
		return x.StatelyCode
	}
	return ""
}

func (x *StatelyErrorDetails) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StatelyErrorDetails) GetUpstreamCause() string {
	if x != nil {
		return x.UpstreamCause
	}
	return ""
}

func (x *StatelyErrorDetails) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_errors_error_details_proto protoreflect.FileDescriptor

const file_errors_error_details_proto_rawDesc = "" +
	"\n" +
	"\x1aerrors/error_details.proto\x12\x0estately.errors\x1a\x1bbuf/validate/validate.proto\"\x9d\x01\n" +
	"\x13StatelyErrorDetails\x12)\n" +
	"\fstately_code\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\vstatelyCode\x12 \n" +
	"\amessage\x18\x02 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\amessage\x12%\n" +
	"\x0eupstream_cause\x18\x03 \x01(\tR\rupstreamCause\x12\x12\n" +
	"\x04code\x18\x04 \x01(\rR\x04codeB\xaa\x01\n" +
	"\x12com.stately.errorsB\x11ErrorDetailsProtoP\x01Z(github.com/StatelyCloud/go-sdk/pb/errors\xa2\x02\x03SEX\xaa\x02\x0eStately.Errors\xca\x02\x0eStately\\Errors\xe2\x02\x1aStately\\Errors\\GPBMetadata\xea\x02\x0fStately::Errorsb\x06proto3"

var (
	file_errors_error_details_proto_rawDescOnce sync.Once
	file_errors_error_details_proto_rawDescData []byte
)

func file_errors_error_details_proto_rawDescGZIP() []byte {
	file_errors_error_details_proto_rawDescOnce.Do(func() {
		file_errors_error_details_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_errors_error_details_proto_rawDesc), len(file_errors_error_details_proto_rawDesc)))
	})
	return file_errors_error_details_proto_rawDescData
}

var file_errors_error_details_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_error_details_proto_goTypes = []any{
	(*StatelyErrorDetails)(nil), // 0: stately.errors.StatelyErrorDetails
}
var file_errors_error_details_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_error_details_proto_init() }
func file_errors_error_details_proto_init() {
	if File_errors_error_details_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_errors_error_details_proto_rawDesc), len(file_errors_error_details_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_error_details_proto_goTypes,
		DependencyIndexes: file_errors_error_details_proto_depIdxs,
		MessageInfos:      file_errors_error_details_proto_msgTypes,
	}.Build()
	File_errors_error_details_proto = out.File
	file_errors_error_details_proto_goTypes = nil
	file_errors_error_details_proto_depIdxs = nil
}
