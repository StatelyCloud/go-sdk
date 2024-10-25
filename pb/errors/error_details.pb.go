// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: errors/error_details.proto

package errors

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

// StatelyErrorDetails is a message containing detailed error information.
// This is returned from the Stately API via Connect error details:
//
//	https://connectrpc.com/docs/go/errors#error-details
//
// Note: As a customer, you should not need to handle this message directly unless writing
// a custom low-level SDK. Instead, language-specific SDKs will provide a more user-friendly
// error object that wraps this message.
type StatelyErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	Code uint32 `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *StatelyErrorDetails) Reset() {
	*x = StatelyErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_error_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatelyErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatelyErrorDetails) ProtoMessage() {}

func (x *StatelyErrorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_errors_error_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_errors_error_details_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x9d, 0x01, 0x0a,
	0x13, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x29, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x61,
	0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0xaa, 0x01, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x42, 0x11, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xa2, 0x02, 0x03, 0x53, 0x45, 0x58, 0xaa, 0x02, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x6c, 0x79, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x0e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xe2, 0x02, 0x1a, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x6c, 0x79, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c,
	0x79, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_errors_error_details_proto_rawDescOnce sync.Once
	file_errors_error_details_proto_rawDescData = file_errors_error_details_proto_rawDesc
)

func file_errors_error_details_proto_rawDescGZIP() []byte {
	file_errors_error_details_proto_rawDescOnce.Do(func() {
		file_errors_error_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_error_details_proto_rawDescData)
	})
	return file_errors_error_details_proto_rawDescData
}

var file_errors_error_details_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_error_details_proto_goTypes = []interface{}{
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
	if !protoimpl.UnsafeEnabled {
		file_errors_error_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatelyErrorDetails); i {
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
			RawDescriptor: file_errors_error_details_proto_rawDesc,
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
	file_errors_error_details_proto_rawDesc = nil
	file_errors_error_details_proto_goTypes = nil
	file_errors_error_details_proto_depIdxs = nil
}
