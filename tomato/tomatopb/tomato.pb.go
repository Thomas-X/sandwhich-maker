// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tomatopb/tomato.proto

package tomatopb

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

type GetTomatoRequest struct {
	Tomatoes             []string `protobuf:"bytes,1,rep,name=tomatoes,proto3" json:"tomatoes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTomatoRequest) Reset()         { *m = GetTomatoRequest{} }
func (m *GetTomatoRequest) String() string { return proto.CompactTextString(m) }
func (*GetTomatoRequest) ProtoMessage()    {}
func (*GetTomatoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6c6d9b0efa3f328, []int{0}
}

func (m *GetTomatoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTomatoRequest.Unmarshal(m, b)
}
func (m *GetTomatoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTomatoRequest.Marshal(b, m, deterministic)
}
func (m *GetTomatoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTomatoRequest.Merge(m, src)
}
func (m *GetTomatoRequest) XXX_Size() int {
	return xxx_messageInfo_GetTomatoRequest.Size(m)
}
func (m *GetTomatoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTomatoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTomatoRequest proto.InternalMessageInfo

func (m *GetTomatoRequest) GetTomatoes() []string {
	if m != nil {
		return m.Tomatoes
	}
	return nil
}

type Tomato struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                int32    `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tomato) Reset()         { *m = Tomato{} }
func (m *Tomato) String() string { return proto.CompactTextString(m) }
func (*Tomato) ProtoMessage()    {}
func (*Tomato) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6c6d9b0efa3f328, []int{1}
}

func (m *Tomato) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tomato.Unmarshal(m, b)
}
func (m *Tomato) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tomato.Marshal(b, m, deterministic)
}
func (m *Tomato) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tomato.Merge(m, src)
}
func (m *Tomato) XXX_Size() int {
	return xxx_messageInfo_Tomato.Size(m)
}
func (m *Tomato) XXX_DiscardUnknown() {
	xxx_messageInfo_Tomato.DiscardUnknown(m)
}

var xxx_messageInfo_Tomato proto.InternalMessageInfo

func (m *Tomato) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tomato) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type GetTomatoResponse struct {
	Tomatoes             []*Tomato `protobuf:"bytes,1,rep,name=tomatoes,proto3" json:"tomatoes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetTomatoResponse) Reset()         { *m = GetTomatoResponse{} }
func (m *GetTomatoResponse) String() string { return proto.CompactTextString(m) }
func (*GetTomatoResponse) ProtoMessage()    {}
func (*GetTomatoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6c6d9b0efa3f328, []int{2}
}

func (m *GetTomatoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTomatoResponse.Unmarshal(m, b)
}
func (m *GetTomatoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTomatoResponse.Marshal(b, m, deterministic)
}
func (m *GetTomatoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTomatoResponse.Merge(m, src)
}
func (m *GetTomatoResponse) XXX_Size() int {
	return xxx_messageInfo_GetTomatoResponse.Size(m)
}
func (m *GetTomatoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTomatoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTomatoResponse proto.InternalMessageInfo

func (m *GetTomatoResponse) GetTomatoes() []*Tomato {
	if m != nil {
		return m.Tomatoes
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTomatoRequest)(nil), "tomato.GetTomatoRequest")
	proto.RegisterType((*Tomato)(nil), "tomato.Tomato")
	proto.RegisterType((*GetTomatoResponse)(nil), "tomato.GetTomatoResponse")
}

func init() { proto.RegisterFile("tomatopb/tomato.proto", fileDescriptor_f6c6d9b0efa3f328) }

var fileDescriptor_f6c6d9b0efa3f328 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0xc9, 0xcf, 0x4d,
	0x2c, 0xc9, 0x2f, 0x48, 0xd2, 0x87, 0x30, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0x20,
	0x3c, 0x25, 0x3d, 0x2e, 0x01, 0xf7, 0xd4, 0x92, 0x10, 0x30, 0x27, 0x28, 0xb5, 0xb0, 0x34, 0xb5,
	0xb8, 0x44, 0x48, 0x8a, 0x8b, 0x03, 0x22, 0x9b, 0x5a, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x19,
	0x04, 0xe7, 0x2b, 0x19, 0x71, 0xb1, 0x41, 0x14, 0x0b, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0x05, 0x45, 0x99,
	0xc9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x92, 0x3d, 0x97, 0x20, 0x92,
	0x1d, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x5a, 0x68, 0x96, 0x70, 0x1b, 0xf1, 0xe9, 0x41,
	0x5d, 0x08, 0x55, 0x09, 0x97, 0x37, 0x0a, 0xe6, 0xe2, 0x85, 0x88, 0x05, 0xa7, 0x16, 0x95, 0x65,
	0x26, 0xa7, 0x0a, 0x39, 0x71, 0x71, 0xc2, 0x4d, 0x14, 0x92, 0x80, 0xe9, 0x43, 0xf7, 0x88, 0x94,
	0x24, 0x16, 0x19, 0x88, 0xf5, 0x4a, 0x0c, 0x4e, 0x5c, 0x51, 0x1c, 0xb0, 0xa0, 0x49, 0x62, 0x03,
	0x07, 0x8a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x43, 0xb6, 0xdb, 0x1c, 0x2d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TomatoServiceClient is the client API for TomatoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TomatoServiceClient interface {
	GetTomato(ctx context.Context, in *GetTomatoRequest, opts ...grpc.CallOption) (*GetTomatoResponse, error)
}

type tomatoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTomatoServiceClient(cc *grpc.ClientConn) TomatoServiceClient {
	return &tomatoServiceClient{cc}
}

func (c *tomatoServiceClient) GetTomato(ctx context.Context, in *GetTomatoRequest, opts ...grpc.CallOption) (*GetTomatoResponse, error) {
	out := new(GetTomatoResponse)
	err := c.cc.Invoke(ctx, "/tomato.TomatoService/GetTomato", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TomatoServiceServer is the server API for TomatoService service.
type TomatoServiceServer interface {
	GetTomato(context.Context, *GetTomatoRequest) (*GetTomatoResponse, error)
}

// UnimplementedTomatoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTomatoServiceServer struct {
}

func (*UnimplementedTomatoServiceServer) GetTomato(ctx context.Context, req *GetTomatoRequest) (*GetTomatoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTomato not implemented")
}

func RegisterTomatoServiceServer(s *grpc.Server, srv TomatoServiceServer) {
	s.RegisterService(&_TomatoService_serviceDesc, srv)
}

func _TomatoService_GetTomato_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTomatoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TomatoServiceServer).GetTomato(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tomato.TomatoService/GetTomato",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TomatoServiceServer).GetTomato(ctx, req.(*GetTomatoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TomatoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tomato.TomatoService",
	HandlerType: (*TomatoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTomato",
			Handler:    _TomatoService_GetTomato_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tomatopb/tomato.proto",
}