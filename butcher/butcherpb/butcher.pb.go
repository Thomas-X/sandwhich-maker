// Code generated by protoc-gen-go. DO NOT EDIT.
// source: butcherpb/butcher.proto

package butcher

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetMeatRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMeatRequest) Reset()         { *m = GetMeatRequest{} }
func (m *GetMeatRequest) String() string { return proto.CompactTextString(m) }
func (*GetMeatRequest) ProtoMessage()    {}
func (*GetMeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98c07088cbbc2c32, []int{0}
}

func (m *GetMeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMeatRequest.Unmarshal(m, b)
}
func (m *GetMeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMeatRequest.Marshal(b, m, deterministic)
}
func (m *GetMeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMeatRequest.Merge(m, src)
}
func (m *GetMeatRequest) XXX_Size() int {
	return xxx_messageInfo_GetMeatRequest.Size(m)
}
func (m *GetMeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMeatRequest proto.InternalMessageInfo

func (m *GetMeatRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetMeatResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                int32    `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMeatResponse) Reset()         { *m = GetMeatResponse{} }
func (m *GetMeatResponse) String() string { return proto.CompactTextString(m) }
func (*GetMeatResponse) ProtoMessage()    {}
func (*GetMeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98c07088cbbc2c32, []int{1}
}

func (m *GetMeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMeatResponse.Unmarshal(m, b)
}
func (m *GetMeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMeatResponse.Marshal(b, m, deterministic)
}
func (m *GetMeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMeatResponse.Merge(m, src)
}
func (m *GetMeatResponse) XXX_Size() int {
	return xxx_messageInfo_GetMeatResponse.Size(m)
}
func (m *GetMeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMeatResponse proto.InternalMessageInfo

func (m *GetMeatResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetMeatResponse) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*GetMeatRequest)(nil), "butcher.GetMeatRequest")
	proto.RegisterType((*GetMeatResponse)(nil), "butcher.GetMeatResponse")
}

func init() { proto.RegisterFile("butcherpb/butcher.proto", fileDescriptor_98c07088cbbc2c32) }

var fileDescriptor_98c07088cbbc2c32 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2a, 0x2d, 0x49,
	0xce, 0x48, 0x2d, 0x2a, 0x48, 0xd2, 0x87, 0xb2, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8,
	0xa1, 0x5c, 0x25, 0x15, 0x2e, 0x3e, 0xf7, 0xd4, 0x12, 0xdf, 0xd4, 0xc4, 0x92, 0xa0, 0xd4, 0xc2,
	0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x9a, 0x8b, 0x1f, 0xae, 0xaa, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x15, 0x9b, 0x32, 0x21, 0x11, 0x2e, 0xd6, 0x82, 0xa2, 0xcc, 0xe4, 0x54, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0xd6, 0x20, 0x08, 0xc7, 0x28, 0x80, 0x8b, 0xcf, 0x09, 0x62, 0x5b, 0x70, 0x6a, 0x51, 0x59,
	0x66, 0x72, 0xaa, 0x90, 0x1d, 0x17, 0x3b, 0xd4, 0x38, 0x21, 0x71, 0x3d, 0x98, 0xc3, 0x50, 0x9d,
	0x21, 0x25, 0x81, 0x29, 0x01, 0xb1, 0x59, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x09, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x9b, 0x1d, 0x66, 0xd0, 0xdf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ButcherServiceClient is the client API for ButcherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ButcherServiceClient interface {
	GetMeat(ctx context.Context, in *GetMeatRequest, opts ...grpc.CallOption) (*GetMeatResponse, error)
}

type butcherServiceClient struct {
	cc *grpc.ClientConn
}

func NewButcherServiceClient(cc *grpc.ClientConn) ButcherServiceClient {
	return &butcherServiceClient{cc}
}

func (c *butcherServiceClient) GetMeat(ctx context.Context, in *GetMeatRequest, opts ...grpc.CallOption) (*GetMeatResponse, error) {
	out := new(GetMeatResponse)
	err := c.cc.Invoke(ctx, "/butcher.ButcherService/GetMeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ButcherServiceServer is the server API for ButcherService service.
type ButcherServiceServer interface {
	GetMeat(context.Context, *GetMeatRequest) (*GetMeatResponse, error)
}

// UnimplementedButcherServiceServer can be embedded to have forward compatible implementations.
type UnimplementedButcherServiceServer struct {
}

func (*UnimplementedButcherServiceServer) GetMeat(ctx context.Context, req *GetMeatRequest) (*GetMeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeat not implemented")
}

func RegisterButcherServiceServer(s *grpc.Server, srv ButcherServiceServer) {
	s.RegisterService(&_ButcherService_serviceDesc, srv)
}

func _ButcherService_GetMeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ButcherServiceServer).GetMeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/butcher.ButcherService/GetMeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ButcherServiceServer).GetMeat(ctx, req.(*GetMeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ButcherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "butcher.ButcherService",
	HandlerType: (*ButcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMeat",
			Handler:    _ButcherService_GetMeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "butcherpb/butcher.proto",
}
