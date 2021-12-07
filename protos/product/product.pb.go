// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: test-proto/product.proto

package __

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

type ProductNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductNameRequest) Reset() {
	*x = ProductNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductNameRequest) ProtoMessage() {}

func (x *ProductNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductNameRequest.ProtoReflect.Descriptor instead.
func (*ProductNameRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductNameRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProductNameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProductNameReply) Reset() {
	*x = ProductNameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductNameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductNameReply) ProtoMessage() {}

func (x *ProductNameReply) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductNameReply.ProtoReflect.Descriptor instead.
func (*ProductNameReply) Descriptor() ([]byte, []int) {
	return file_test_proto_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductNameReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_test_proto_product_proto protoreflect.FileDescriptor

var file_test_proto_product_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x26, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x4d, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0f, 0x53, 0x65,
	0x6e, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_product_proto_rawDescOnce sync.Once
	file_test_proto_product_proto_rawDescData = file_test_proto_product_proto_rawDesc
)

func file_test_proto_product_proto_rawDescGZIP() []byte {
	file_test_proto_product_proto_rawDescOnce.Do(func() {
		file_test_proto_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_product_proto_rawDescData)
	})
	return file_test_proto_product_proto_rawDescData
}

var file_test_proto_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_product_proto_goTypes = []interface{}{
	(*ProductNameRequest)(nil), // 0: ProductNameRequest
	(*ProductNameReply)(nil),   // 1: ProductNameReply
}
var file_test_proto_product_proto_depIdxs = []int32{
	0, // 0: ProductService.SendProductName:input_type -> ProductNameRequest
	1, // 1: ProductService.SendProductName:output_type -> ProductNameReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_product_proto_init() }
func file_test_proto_product_proto_init() {
	if File_test_proto_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductNameRequest); i {
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
		file_test_proto_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductNameReply); i {
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
			RawDescriptor: file_test_proto_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_product_proto_goTypes,
		DependencyIndexes: file_test_proto_product_proto_depIdxs,
		MessageInfos:      file_test_proto_product_proto_msgTypes,
	}.Build()
	File_test_proto_product_proto = out.File
	file_test_proto_product_proto_rawDesc = nil
	file_test_proto_product_proto_goTypes = nil
	file_test_proto_product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	// Sends a greeting
	SendProductName(ctx context.Context, in *ProductNameRequest, opts ...grpc.CallOption) (*ProductNameReply, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) SendProductName(ctx context.Context, in *ProductNameRequest, opts ...grpc.CallOption) (*ProductNameReply, error) {
	out := new(ProductNameReply)
	err := c.cc.Invoke(ctx, "/ProductService/SendProductName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	// Sends a greeting
	SendProductName(context.Context, *ProductNameRequest) (*ProductNameReply, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) SendProductName(context.Context, *ProductNameRequest) (*ProductNameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendProductName not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_SendProductName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SendProductName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/SendProductName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SendProductName(ctx, req.(*ProductNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendProductName",
			Handler:    _ProductService_SendProductName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test-proto/product.proto",
}
