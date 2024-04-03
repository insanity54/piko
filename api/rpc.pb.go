// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: api/rpc.proto

package api

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProxyHttpStatus int32

const (
	ProxyHttpStatus_OK                   ProxyHttpStatus = 0
	ProxyHttpStatus_UPSTREAM_TIMEOUT     ProxyHttpStatus = 1
	ProxyHttpStatus_UPSTREAM_UNREACHABLE ProxyHttpStatus = 2
	ProxyHttpStatus_INTERNAL_ERROR       ProxyHttpStatus = 3
)

// Enum value maps for ProxyHttpStatus.
var (
	ProxyHttpStatus_name = map[int32]string{
		0: "OK",
		1: "UPSTREAM_TIMEOUT",
		2: "UPSTREAM_UNREACHABLE",
		3: "INTERNAL_ERROR",
	}
	ProxyHttpStatus_value = map[string]int32{
		"OK":                   0,
		"UPSTREAM_TIMEOUT":     1,
		"UPSTREAM_UNREACHABLE": 2,
		"INTERNAL_ERROR":       3,
	}
)

func (x ProxyHttpStatus) Enum() *ProxyHttpStatus {
	p := new(ProxyHttpStatus)
	*p = x
	return p
}

func (x ProxyHttpStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProxyHttpStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_rpc_proto_enumTypes[0].Descriptor()
}

func (ProxyHttpStatus) Type() protoreflect.EnumType {
	return &file_api_rpc_proto_enumTypes[0]
}

func (x ProxyHttpStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProxyHttpStatus.Descriptor instead.
func (ProxyHttpStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_rpc_proto_rawDescGZIP(), []int{0}
}

type ProxyHttpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains the encoded HTTP request.
	HttpReq []byte `protobuf:"bytes,1,opt,name=http_req,json=httpReq,proto3" json:"http_req,omitempty"`
}

func (x *ProxyHttpReq) Reset() {
	*x = ProxyHttpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyHttpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyHttpReq) ProtoMessage() {}

func (x *ProxyHttpReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyHttpReq.ProtoReflect.Descriptor instead.
func (*ProxyHttpReq) Descriptor() ([]byte, []int) {
	return file_api_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyHttpReq) GetHttpReq() []byte {
	if x != nil {
		return x.HttpReq
	}
	return nil
}

type ProxyHttpError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  ProxyHttpStatus `protobuf:"varint,1,opt,name=status,proto3,enum=api.ProxyHttpStatus" json:"status,omitempty"`
	Message string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ProxyHttpError) Reset() {
	*x = ProxyHttpError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyHttpError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyHttpError) ProtoMessage() {}

func (x *ProxyHttpError) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyHttpError.ProtoReflect.Descriptor instead.
func (*ProxyHttpError) Descriptor() ([]byte, []int) {
	return file_api_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyHttpError) GetStatus() ProxyHttpStatus {
	if x != nil {
		return x.Status
	}
	return ProxyHttpStatus_OK
}

func (x *ProxyHttpError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ProxyHttpResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains the encoded HTTP response, or is empty if an error occured.
	HttpResp []byte          `protobuf:"bytes,1,opt,name=http_resp,json=httpResp,proto3" json:"http_resp,omitempty"`
	Error    *ProxyHttpError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ProxyHttpResp) Reset() {
	*x = ProxyHttpResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyHttpResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyHttpResp) ProtoMessage() {}

func (x *ProxyHttpResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyHttpResp.ProtoReflect.Descriptor instead.
func (*ProxyHttpResp) Descriptor() ([]byte, []int) {
	return file_api_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *ProxyHttpResp) GetHttpResp() []byte {
	if x != nil {
		return x.HttpResp
	}
	return nil
}

func (x *ProxyHttpResp) GetError() *ProxyHttpError {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_api_rpc_proto protoreflect.FileDescriptor

var file_api_rpc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x29, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x48, 0x74, 0x74,
	0x70, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x22,
	0x58, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x48, 0x74, 0x74, 0x70, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x48, 0x74, 0x74,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x57, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x68,
	0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x48, 0x74, 0x74, 0x70, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2a, 0x5d, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x48, 0x74, 0x74, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x55, 0x50, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x50, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f,
	0x55, 0x4e, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a,
	0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x03, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6e, 0x64, 0x79, 0x64, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x2f, 0x70, 0x69, 0x63,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rpc_proto_rawDescOnce sync.Once
	file_api_rpc_proto_rawDescData = file_api_rpc_proto_rawDesc
)

func file_api_rpc_proto_rawDescGZIP() []byte {
	file_api_rpc_proto_rawDescOnce.Do(func() {
		file_api_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rpc_proto_rawDescData)
	})
	return file_api_rpc_proto_rawDescData
}

var file_api_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_rpc_proto_goTypes = []interface{}{
	(ProxyHttpStatus)(0),   // 0: api.ProxyHttpStatus
	(*ProxyHttpReq)(nil),   // 1: api.ProxyHttpReq
	(*ProxyHttpError)(nil), // 2: api.ProxyHttpError
	(*ProxyHttpResp)(nil),  // 3: api.ProxyHttpResp
}
var file_api_rpc_proto_depIdxs = []int32{
	0, // 0: api.ProxyHttpError.status:type_name -> api.ProxyHttpStatus
	2, // 1: api.ProxyHttpResp.error:type_name -> api.ProxyHttpError
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_rpc_proto_init() }
func file_api_rpc_proto_init() {
	if File_api_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyHttpReq); i {
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
		file_api_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyHttpError); i {
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
		file_api_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyHttpResp); i {
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
			RawDescriptor: file_api_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_rpc_proto_goTypes,
		DependencyIndexes: file_api_rpc_proto_depIdxs,
		EnumInfos:         file_api_rpc_proto_enumTypes,
		MessageInfos:      file_api_rpc_proto_msgTypes,
	}.Build()
	File_api_rpc_proto = out.File
	file_api_rpc_proto_rawDesc = nil
	file_api_rpc_proto_goTypes = nil
	file_api_rpc_proto_depIdxs = nil
}
