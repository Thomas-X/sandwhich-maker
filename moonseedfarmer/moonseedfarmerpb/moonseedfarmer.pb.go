// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moonseedfarmerpb/moonseedfarmer.proto

package moonseedfarmerpb

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

type GetMoonSeedRequest struct {
	BreadName            string   `protobuf:"bytes,1,opt,name=breadName,proto3" json:"breadName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMoonSeedRequest) Reset()         { *m = GetMoonSeedRequest{} }
func (m *GetMoonSeedRequest) String() string { return proto.CompactTextString(m) }
func (*GetMoonSeedRequest) ProtoMessage()    {}
func (*GetMoonSeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_468f641641730a40, []int{0}
}

func (m *GetMoonSeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMoonSeedRequest.Unmarshal(m, b)
}
func (m *GetMoonSeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMoonSeedRequest.Marshal(b, m, deterministic)
}
func (m *GetMoonSeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMoonSeedRequest.Merge(m, src)
}
func (m *GetMoonSeedRequest) XXX_Size() int {
	return xxx_messageInfo_GetMoonSeedRequest.Size(m)
}
func (m *GetMoonSeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMoonSeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMoonSeedRequest proto.InternalMessageInfo

func (m *GetMoonSeedRequest) GetBreadName() string {
	if m != nil {
		return m.BreadName
	}
	return ""
}

type GetMoonSeedResponse struct {
	Price                int32    `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMoonSeedResponse) Reset()         { *m = GetMoonSeedResponse{} }
func (m *GetMoonSeedResponse) String() string { return proto.CompactTextString(m) }
func (*GetMoonSeedResponse) ProtoMessage()    {}
func (*GetMoonSeedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_468f641641730a40, []int{1}
}

func (m *GetMoonSeedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMoonSeedResponse.Unmarshal(m, b)
}
func (m *GetMoonSeedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMoonSeedResponse.Marshal(b, m, deterministic)
}
func (m *GetMoonSeedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMoonSeedResponse.Merge(m, src)
}
func (m *GetMoonSeedResponse) XXX_Size() int {
	return xxx_messageInfo_GetMoonSeedResponse.Size(m)
}
func (m *GetMoonSeedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMoonSeedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMoonSeedResponse proto.InternalMessageInfo

func (m *GetMoonSeedResponse) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*GetMoonSeedRequest)(nil), "moonseedfarmer.GetMoonSeedRequest")
	proto.RegisterType((*GetMoonSeedResponse)(nil), "moonseedfarmer.GetMoonSeedResponse")
}

func init() {
	proto.RegisterFile("moonseedfarmerpb/moonseedfarmer.proto", fileDescriptor_468f641641730a40)
}

var fileDescriptor_468f641641730a40 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0xcd, 0xcf, 0xcf,
	0x2b, 0x4e, 0x4d, 0x4d, 0x49, 0x4b, 0x2c, 0xca, 0x4d, 0x2d, 0x2a, 0x48, 0xd2, 0x47, 0x15, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x43, 0x15, 0x55, 0x32, 0xe2, 0x12, 0x72, 0x4f, 0x2d,
	0xf1, 0xcd, 0xcf, 0xcf, 0x0b, 0x4e, 0x4d, 0x4d, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x92, 0xe1, 0xe2, 0x4c, 0x2a, 0x4a, 0x4d, 0x4c, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x28, 0x69, 0x73, 0x09, 0xa3, 0xe8, 0x29, 0x2e, 0x00, 0x19, 0x29,
	0x24, 0xc2, 0xc5, 0x5a, 0x50, 0x94, 0x99, 0x0c, 0xd1, 0xc0, 0x1a, 0x04, 0xe1, 0x18, 0xe5, 0x73,
	0x89, 0xc2, 0x54, 0xba, 0x81, 0xad, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x0a, 0xe3,
	0xe2, 0x46, 0x32, 0x45, 0x48, 0x49, 0x0f, 0xcd, 0xbd, 0x98, 0xce, 0x92, 0x52, 0xc6, 0xab, 0x06,
	0xe2, 0x0c, 0x27, 0xa1, 0x28, 0x01, 0xf4, 0xa0, 0x48, 0x62, 0x03, 0x7b, 0xde, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x73, 0x15, 0x9b, 0x7d, 0x25, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MoonSeedFarmerServiceClient is the client API for MoonSeedFarmerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoonSeedFarmerServiceClient interface {
	GetMoonSeed(ctx context.Context, in *GetMoonSeedRequest, opts ...grpc.CallOption) (*GetMoonSeedResponse, error)
}

type moonSeedFarmerServiceClient struct {
	cc *grpc.ClientConn
}

func NewMoonSeedFarmerServiceClient(cc *grpc.ClientConn) MoonSeedFarmerServiceClient {
	return &moonSeedFarmerServiceClient{cc}
}

func (c *moonSeedFarmerServiceClient) GetMoonSeed(ctx context.Context, in *GetMoonSeedRequest, opts ...grpc.CallOption) (*GetMoonSeedResponse, error) {
	out := new(GetMoonSeedResponse)
	err := c.cc.Invoke(ctx, "/moonseedfarmer.MoonSeedFarmerService/GetMoonSeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoonSeedFarmerServiceServer is the server API for MoonSeedFarmerService service.
type MoonSeedFarmerServiceServer interface {
	GetMoonSeed(context.Context, *GetMoonSeedRequest) (*GetMoonSeedResponse, error)
}

// UnimplementedMoonSeedFarmerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMoonSeedFarmerServiceServer struct {
}

func (*UnimplementedMoonSeedFarmerServiceServer) GetMoonSeed(ctx context.Context, req *GetMoonSeedRequest) (*GetMoonSeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoonSeed not implemented")
}

func RegisterMoonSeedFarmerServiceServer(s *grpc.Server, srv MoonSeedFarmerServiceServer) {
	s.RegisterService(&_MoonSeedFarmerService_serviceDesc, srv)
}

func _MoonSeedFarmerService_GetMoonSeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMoonSeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoonSeedFarmerServiceServer).GetMoonSeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moonseedfarmer.MoonSeedFarmerService/GetMoonSeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoonSeedFarmerServiceServer).GetMoonSeed(ctx, req.(*GetMoonSeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MoonSeedFarmerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moonseedfarmer.MoonSeedFarmerService",
	HandlerType: (*MoonSeedFarmerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMoonSeed",
			Handler:    _MoonSeedFarmerService_GetMoonSeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moonseedfarmerpb/moonseedfarmer.proto",
}
