// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: application/grpc/proto/v1/currency_conversion_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ConvertCurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Code   string  `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ConvertCurrencyRequest) Reset() {
	*x = ConvertCurrencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_grpc_proto_v1_currency_conversion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertCurrencyRequest) ProtoMessage() {}

func (x *ConvertCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_grpc_proto_v1_currency_conversion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertCurrencyRequest.ProtoReflect.Descriptor instead.
func (*ConvertCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_application_grpc_proto_v1_currency_conversion_service_proto_rawDescGZIP(), []int{0}
}

func (x *ConvertCurrencyRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ConvertCurrencyRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ConvertCurrencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values map[string]float64 `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *ConvertCurrencyResponse) Reset() {
	*x = ConvertCurrencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_grpc_proto_v1_currency_conversion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertCurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertCurrencyResponse) ProtoMessage() {}

func (x *ConvertCurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_application_grpc_proto_v1_currency_conversion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertCurrencyResponse.ProtoReflect.Descriptor instead.
func (*ConvertCurrencyResponse) Descriptor() ([]byte, []int) {
	return file_application_grpc_proto_v1_currency_conversion_service_proto_rawDescGZIP(), []int{1}
}

func (x *ConvertCurrencyResponse) GetValues() map[string]float64 {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_application_grpc_proto_v1_currency_conversion_service_proto protoreflect.FileDescriptor

var file_application_grpc_proto_v1_currency_conversion_service_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x31, 0x22, 0x44, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0x67, 0x0a, 0x19, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0f,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_application_grpc_proto_v1_currency_conversion_service_proto_rawDescOnce sync.Once
	file_application_grpc_proto_v1_currency_conversion_service_proto_rawDescData = file_application_grpc_proto_v1_currency_conversion_service_proto_rawDesc
)

func file_application_grpc_proto_v1_currency_conversion_service_proto_rawDescGZIP() []byte {
	file_application_grpc_proto_v1_currency_conversion_service_proto_rawDescOnce.Do(func() {
		file_application_grpc_proto_v1_currency_conversion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_application_grpc_proto_v1_currency_conversion_service_proto_rawDescData)
	})
	return file_application_grpc_proto_v1_currency_conversion_service_proto_rawDescData
}

var file_application_grpc_proto_v1_currency_conversion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_application_grpc_proto_v1_currency_conversion_service_proto_goTypes = []interface{}{
	(*ConvertCurrencyRequest)(nil),  // 0: v1.ConvertCurrencyRequest
	(*ConvertCurrencyResponse)(nil), // 1: v1.ConvertCurrencyResponse
	nil,                             // 2: v1.ConvertCurrencyResponse.ValuesEntry
}
var file_application_grpc_proto_v1_currency_conversion_service_proto_depIdxs = []int32{
	2, // 0: v1.ConvertCurrencyResponse.values:type_name -> v1.ConvertCurrencyResponse.ValuesEntry
	0, // 1: v1.CurrencyConversionService.ConvertCurrency:input_type -> v1.ConvertCurrencyRequest
	1, // 2: v1.CurrencyConversionService.ConvertCurrency:output_type -> v1.ConvertCurrencyResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_application_grpc_proto_v1_currency_conversion_service_proto_init() }
func file_application_grpc_proto_v1_currency_conversion_service_proto_init() {
	if File_application_grpc_proto_v1_currency_conversion_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_application_grpc_proto_v1_currency_conversion_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertCurrencyRequest); i {
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
		file_application_grpc_proto_v1_currency_conversion_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertCurrencyResponse); i {
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
			RawDescriptor: file_application_grpc_proto_v1_currency_conversion_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_application_grpc_proto_v1_currency_conversion_service_proto_goTypes,
		DependencyIndexes: file_application_grpc_proto_v1_currency_conversion_service_proto_depIdxs,
		MessageInfos:      file_application_grpc_proto_v1_currency_conversion_service_proto_msgTypes,
	}.Build()
	File_application_grpc_proto_v1_currency_conversion_service_proto = out.File
	file_application_grpc_proto_v1_currency_conversion_service_proto_rawDesc = nil
	file_application_grpc_proto_v1_currency_conversion_service_proto_goTypes = nil
	file_application_grpc_proto_v1_currency_conversion_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyConversionServiceClient is the client API for CurrencyConversionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyConversionServiceClient interface {
	ConvertCurrency(ctx context.Context, in *ConvertCurrencyRequest, opts ...grpc.CallOption) (*ConvertCurrencyResponse, error)
}

type currencyConversionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyConversionServiceClient(cc grpc.ClientConnInterface) CurrencyConversionServiceClient {
	return &currencyConversionServiceClient{cc}
}

func (c *currencyConversionServiceClient) ConvertCurrency(ctx context.Context, in *ConvertCurrencyRequest, opts ...grpc.CallOption) (*ConvertCurrencyResponse, error) {
	out := new(ConvertCurrencyResponse)
	err := c.cc.Invoke(ctx, "/v1.CurrencyConversionService/ConvertCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyConversionServiceServer is the server API for CurrencyConversionService service.
type CurrencyConversionServiceServer interface {
	ConvertCurrency(context.Context, *ConvertCurrencyRequest) (*ConvertCurrencyResponse, error)
}

// UnimplementedCurrencyConversionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyConversionServiceServer struct {
}

func (*UnimplementedCurrencyConversionServiceServer) ConvertCurrency(context.Context, *ConvertCurrencyRequest) (*ConvertCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertCurrency not implemented")
}

func RegisterCurrencyConversionServiceServer(s *grpc.Server, srv CurrencyConversionServiceServer) {
	s.RegisterService(&_CurrencyConversionService_serviceDesc, srv)
}

func _CurrencyConversionService_ConvertCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyConversionServiceServer).ConvertCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CurrencyConversionService/ConvertCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyConversionServiceServer).ConvertCurrency(ctx, req.(*ConvertCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CurrencyConversionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.CurrencyConversionService",
	HandlerType: (*CurrencyConversionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertCurrency",
			Handler:    _CurrencyConversionService_ConvertCurrency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application/grpc/proto/v1/currency_conversion_service.proto",
}