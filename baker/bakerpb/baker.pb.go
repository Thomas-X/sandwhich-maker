// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bakerpb/baker.proto

package bakerpb

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

type Bread struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Amount               int32    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	GrainUsed            int32    `protobuf:"varint,4,opt,name=grainUsed,proto3" json:"grainUsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bread) Reset()         { *m = Bread{} }
func (m *Bread) String() string { return proto.CompactTextString(m) }
func (*Bread) ProtoMessage()    {}
func (*Bread) Descriptor() ([]byte, []int) {
	return fileDescriptor_0409a05eb2d91315, []int{0}
}

func (m *Bread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bread.Unmarshal(m, b)
}
func (m *Bread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bread.Marshal(b, m, deterministic)
}
func (m *Bread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bread.Merge(m, src)
}
func (m *Bread) XXX_Size() int {
	return xxx_messageInfo_Bread.Size(m)
}
func (m *Bread) XXX_DiscardUnknown() {
	xxx_messageInfo_Bread.DiscardUnknown(m)
}

var xxx_messageInfo_Bread proto.InternalMessageInfo

func (m *Bread) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Bread) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Bread) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Bread) GetGrainUsed() int32 {
	if m != nil {
		return m.GrainUsed
	}
	return 0
}

type GetBreadRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Amount               int32    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBreadRequest) Reset()         { *m = GetBreadRequest{} }
func (m *GetBreadRequest) String() string { return proto.CompactTextString(m) }
func (*GetBreadRequest) ProtoMessage()    {}
func (*GetBreadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0409a05eb2d91315, []int{1}
}

func (m *GetBreadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreadRequest.Unmarshal(m, b)
}
func (m *GetBreadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreadRequest.Marshal(b, m, deterministic)
}
func (m *GetBreadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreadRequest.Merge(m, src)
}
func (m *GetBreadRequest) XXX_Size() int {
	return xxx_messageInfo_GetBreadRequest.Size(m)
}
func (m *GetBreadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreadRequest proto.InternalMessageInfo

func (m *GetBreadRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetBreadRequest) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type GetBreadResponse struct {
	Bread                *Bread   `protobuf:"bytes,1,opt,name=bread,proto3" json:"bread,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBreadResponse) Reset()         { *m = GetBreadResponse{} }
func (m *GetBreadResponse) String() string { return proto.CompactTextString(m) }
func (*GetBreadResponse) ProtoMessage()    {}
func (*GetBreadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0409a05eb2d91315, []int{2}
}

func (m *GetBreadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreadResponse.Unmarshal(m, b)
}
func (m *GetBreadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreadResponse.Marshal(b, m, deterministic)
}
func (m *GetBreadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreadResponse.Merge(m, src)
}
func (m *GetBreadResponse) XXX_Size() int {
	return xxx_messageInfo_GetBreadResponse.Size(m)
}
func (m *GetBreadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreadResponse proto.InternalMessageInfo

func (m *GetBreadResponse) GetBread() *Bread {
	if m != nil {
		return m.Bread
	}
	return nil
}

func init() {
	proto.RegisterType((*Bread)(nil), "baker.Bread")
	proto.RegisterType((*GetBreadRequest)(nil), "baker.GetBreadRequest")
	proto.RegisterType((*GetBreadResponse)(nil), "baker.GetBreadResponse")
}

func init() { proto.RegisterFile("bakerpb/baker.proto", fileDescriptor_0409a05eb2d91315) }

var fileDescriptor_0409a05eb2d91315 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4a, 0xcc, 0x4e,
	0x2d, 0x2a, 0x48, 0xd2, 0x07, 0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e,
	0x52, 0x3a, 0x17, 0xab, 0x53, 0x51, 0x6a, 0x62, 0x8a, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e,
	0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc6, 0xc5, 0x96, 0x98, 0x9b,
	0x5f, 0x9a, 0x57, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0xe5, 0x09, 0x89, 0x70, 0xb1,
	0x16, 0x14, 0x65, 0x26, 0xa7, 0x4a, 0x30, 0x83, 0x85, 0x21, 0x1c, 0x21, 0x19, 0x2e, 0xce, 0xf4,
	0xa2, 0xc4, 0xcc, 0xbc, 0xd0, 0xe2, 0xd4, 0x14, 0x09, 0x16, 0xb0, 0x0c, 0x42, 0x40, 0xc9, 0x96,
	0x8b, 0xdf, 0x3d, 0xb5, 0x04, 0x6c, 0x57, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x09, 0x29, 0x56,
	0x2a, 0x99, 0x71, 0x09, 0x20, 0xb4, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x29, 0x71, 0xb1,
	0x26, 0x81, 0x04, 0xc0, 0x06, 0x70, 0x1b, 0xf1, 0xe8, 0x41, 0xfc, 0x07, 0x51, 0x04, 0x91, 0x32,
	0xf2, 0xe6, 0xe2, 0x71, 0x02, 0x89, 0x06, 0xa7, 0x16, 0x95, 0x81, 0x1c, 0x69, 0xcd, 0xc5, 0x01,
	0x33, 0x47, 0x48, 0x0c, 0xaa, 0x01, 0xcd, 0x5d, 0x52, 0xe2, 0x18, 0xe2, 0x10, 0x0b, 0x9d, 0x38,
	0xa3, 0xd8, 0xa1, 0x41, 0x99, 0xc4, 0x06, 0x0e, 0x45, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb5, 0xc6, 0x43, 0x34, 0x5c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BakerServiceClient is the client API for BakerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BakerServiceClient interface {
	GetBread(ctx context.Context, in *GetBreadRequest, opts ...grpc.CallOption) (*GetBreadResponse, error)
}

type bakerServiceClient struct {
	cc *grpc.ClientConn
}

func NewBakerServiceClient(cc *grpc.ClientConn) BakerServiceClient {
	return &bakerServiceClient{cc}
}

func (c *bakerServiceClient) GetBread(ctx context.Context, in *GetBreadRequest, opts ...grpc.CallOption) (*GetBreadResponse, error) {
	out := new(GetBreadResponse)
	err := c.cc.Invoke(ctx, "/baker.BakerService/GetBread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BakerServiceServer is the server API for BakerService service.
type BakerServiceServer interface {
	GetBread(context.Context, *GetBreadRequest) (*GetBreadResponse, error)
}

// UnimplementedBakerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBakerServiceServer struct {
}

func (*UnimplementedBakerServiceServer) GetBread(ctx context.Context, req *GetBreadRequest) (*GetBreadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBread not implemented")
}

func RegisterBakerServiceServer(s *grpc.Server, srv BakerServiceServer) {
	s.RegisterService(&_BakerService_serviceDesc, srv)
}

func _BakerService_GetBread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BakerServiceServer).GetBread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/baker.BakerService/GetBread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BakerServiceServer).GetBread(ctx, req.(*GetBreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BakerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "baker.BakerService",
	HandlerType: (*BakerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBread",
			Handler:    _BakerService_GetBread_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bakerpb/baker.proto",
}
