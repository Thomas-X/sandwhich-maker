// Code generated by protoc-gen-go. DO NOT EDIT.
// source: toppingspb/toppings.proto

package toppingspb

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

type ProvideToppingRequest struct {
	Meats                []string `protobuf:"bytes,1,rep,name=meats,proto3" json:"meats,omitempty"`
	Lettuces             []string `protobuf:"bytes,2,rep,name=lettuces,proto3" json:"lettuces,omitempty"`
	Tomatoes             []string `protobuf:"bytes,3,rep,name=tomatoes,proto3" json:"tomatoes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProvideToppingRequest) Reset()         { *m = ProvideToppingRequest{} }
func (m *ProvideToppingRequest) String() string { return proto.CompactTextString(m) }
func (*ProvideToppingRequest) ProtoMessage()    {}
func (*ProvideToppingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c01fe464c717444d, []int{0}
}

func (m *ProvideToppingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProvideToppingRequest.Unmarshal(m, b)
}
func (m *ProvideToppingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProvideToppingRequest.Marshal(b, m, deterministic)
}
func (m *ProvideToppingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvideToppingRequest.Merge(m, src)
}
func (m *ProvideToppingRequest) XXX_Size() int {
	return xxx_messageInfo_ProvideToppingRequest.Size(m)
}
func (m *ProvideToppingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvideToppingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProvideToppingRequest proto.InternalMessageInfo

func (m *ProvideToppingRequest) GetMeats() []string {
	if m != nil {
		return m.Meats
	}
	return nil
}

func (m *ProvideToppingRequest) GetLettuces() []string {
	if m != nil {
		return m.Lettuces
	}
	return nil
}

func (m *ProvideToppingRequest) GetTomatoes() []string {
	if m != nil {
		return m.Tomatoes
	}
	return nil
}

type Meat struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                int32    `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meat) Reset()         { *m = Meat{} }
func (m *Meat) String() string { return proto.CompactTextString(m) }
func (*Meat) ProtoMessage()    {}
func (*Meat) Descriptor() ([]byte, []int) {
	return fileDescriptor_c01fe464c717444d, []int{1}
}

func (m *Meat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meat.Unmarshal(m, b)
}
func (m *Meat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meat.Marshal(b, m, deterministic)
}
func (m *Meat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meat.Merge(m, src)
}
func (m *Meat) XXX_Size() int {
	return xxx_messageInfo_Meat.Size(m)
}
func (m *Meat) XXX_DiscardUnknown() {
	xxx_messageInfo_Meat.DiscardUnknown(m)
}

var xxx_messageInfo_Meat proto.InternalMessageInfo

func (m *Meat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Meat) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type Lettuce struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                int32    `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lettuce) Reset()         { *m = Lettuce{} }
func (m *Lettuce) String() string { return proto.CompactTextString(m) }
func (*Lettuce) ProtoMessage()    {}
func (*Lettuce) Descriptor() ([]byte, []int) {
	return fileDescriptor_c01fe464c717444d, []int{2}
}

func (m *Lettuce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lettuce.Unmarshal(m, b)
}
func (m *Lettuce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lettuce.Marshal(b, m, deterministic)
}
func (m *Lettuce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lettuce.Merge(m, src)
}
func (m *Lettuce) XXX_Size() int {
	return xxx_messageInfo_Lettuce.Size(m)
}
func (m *Lettuce) XXX_DiscardUnknown() {
	xxx_messageInfo_Lettuce.DiscardUnknown(m)
}

var xxx_messageInfo_Lettuce proto.InternalMessageInfo

func (m *Lettuce) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Lettuce) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
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
	return fileDescriptor_c01fe464c717444d, []int{3}
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

type ProvideToppingResponse struct {
	Meats                []*Meat    `protobuf:"bytes,1,rep,name=meats,proto3" json:"meats,omitempty"`
	Lettuces             []*Lettuce `protobuf:"bytes,2,rep,name=lettuces,proto3" json:"lettuces,omitempty"`
	Tomatoes             []*Tomato  `protobuf:"bytes,3,rep,name=tomatoes,proto3" json:"tomatoes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProvideToppingResponse) Reset()         { *m = ProvideToppingResponse{} }
func (m *ProvideToppingResponse) String() string { return proto.CompactTextString(m) }
func (*ProvideToppingResponse) ProtoMessage()    {}
func (*ProvideToppingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c01fe464c717444d, []int{4}
}

func (m *ProvideToppingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProvideToppingResponse.Unmarshal(m, b)
}
func (m *ProvideToppingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProvideToppingResponse.Marshal(b, m, deterministic)
}
func (m *ProvideToppingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvideToppingResponse.Merge(m, src)
}
func (m *ProvideToppingResponse) XXX_Size() int {
	return xxx_messageInfo_ProvideToppingResponse.Size(m)
}
func (m *ProvideToppingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvideToppingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProvideToppingResponse proto.InternalMessageInfo

func (m *ProvideToppingResponse) GetMeats() []*Meat {
	if m != nil {
		return m.Meats
	}
	return nil
}

func (m *ProvideToppingResponse) GetLettuces() []*Lettuce {
	if m != nil {
		return m.Lettuces
	}
	return nil
}

func (m *ProvideToppingResponse) GetTomatoes() []*Tomato {
	if m != nil {
		return m.Tomatoes
	}
	return nil
}

func init() {
	proto.RegisterType((*ProvideToppingRequest)(nil), "toppings.ProvideToppingRequest")
	proto.RegisterType((*Meat)(nil), "toppings.Meat")
	proto.RegisterType((*Lettuce)(nil), "toppings.Lettuce")
	proto.RegisterType((*Tomato)(nil), "toppings.Tomato")
	proto.RegisterType((*ProvideToppingResponse)(nil), "toppings.ProvideToppingResponse")
}

func init() { proto.RegisterFile("toppingspb/toppings.proto", fileDescriptor_c01fe464c717444d) }

var fileDescriptor_c01fe464c717444d = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4e, 0xc3, 0x30,
	0x0c, 0xc6, 0xc9, 0xfe, 0xb1, 0xb9, 0xa8, 0x82, 0x08, 0x50, 0xd8, 0x85, 0xaa, 0xe2, 0xd0, 0x03,
	0x0c, 0x94, 0xbd, 0x01, 0x67, 0x90, 0x50, 0x19, 0x17, 0x6e, 0x5d, 0xb1, 0xa6, 0x4a, 0xb4, 0x09,
	0x8d, 0xb7, 0x97, 0xe1, 0x65, 0x51, 0x9b, 0x65, 0x19, 0x15, 0x1c, 0x76, 0xf3, 0x57, 0xdb, 0xb5,
	0xbf, 0x5f, 0x0c, 0x57, 0xa4, 0xb4, 0x2e, 0xaa, 0x95, 0xd1, 0xcb, 0x7b, 0x17, 0xce, 0x74, 0xad,
	0x48, 0xf1, 0xb1, 0xd3, 0x31, 0xc2, 0xc5, 0x4b, 0xad, 0x36, 0xc5, 0x07, 0x2e, 0xec, 0xa7, 0x14,
	0xbf, 0xd6, 0x68, 0x88, 0x9f, 0xc3, 0xb0, 0xc4, 0x8c, 0x8c, 0x60, 0x51, 0x3f, 0x99, 0xa4, 0x56,
	0xf0, 0x29, 0x8c, 0x3f, 0x91, 0x68, 0x9d, 0xa3, 0x11, 0xbd, 0x36, 0xb1, 0xd3, 0x4d, 0x8e, 0x54,
	0x99, 0x91, 0x42, 0x23, 0xfa, 0x36, 0xe7, 0x74, 0xfc, 0x00, 0x83, 0x67, 0xcc, 0x88, 0x73, 0x18,
	0x54, 0x59, 0x89, 0x82, 0x45, 0x2c, 0x99, 0xa4, 0x6d, 0xdc, 0x4c, 0xd2, 0x75, 0x91, 0xa3, 0xe8,
	0x45, 0x2c, 0x19, 0xa6, 0x56, 0xc4, 0x73, 0x38, 0x7e, 0xb2, 0x7f, 0x3e, 0xa0, 0x49, 0xc2, 0x68,
	0xd1, 0x8e, 0x3c, 0xa0, 0xe7, 0x9b, 0xc1, 0x65, 0x17, 0x81, 0xd1, 0xaa, 0x32, 0xc8, 0x6f, 0xf6,
	0x19, 0x04, 0x32, 0x9c, 0xed, 0x30, 0x36, 0x66, 0x1c, 0x93, 0xbb, 0x0e, 0x93, 0x40, 0x9e, 0xf9,
	0xc2, 0xad, 0x87, 0x3d, 0x4c, 0xb7, 0x1d, 0x4c, 0x81, 0x3c, 0xf5, 0xe5, 0x76, 0x7b, 0x0f, 0x4e,
	0xae, 0x20, 0xdc, 0x6e, 0xf5, 0x8a, 0xf5, 0xa6, 0xc8, 0x91, 0xbf, 0x41, 0xf8, 0x7b, 0x5d, 0x7e,
	0xed, 0xfb, 0xff, 0x7c, 0xcb, 0x69, 0xf4, 0x7f, 0x81, 0x75, 0x1a, 0x1f, 0x3d, 0x9e, 0xbc, 0x83,
	0xbf, 0x97, 0xe5, 0xa8, 0xbd, 0x93, 0xf9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0x51, 0x9a,
	0x04, 0x44, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ToppingServiceClient is the client API for ToppingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ToppingServiceClient interface {
	ProvideTopping(ctx context.Context, in *ProvideToppingRequest, opts ...grpc.CallOption) (*ProvideToppingResponse, error)
}

type toppingServiceClient struct {
	cc *grpc.ClientConn
}

func NewToppingServiceClient(cc *grpc.ClientConn) ToppingServiceClient {
	return &toppingServiceClient{cc}
}

func (c *toppingServiceClient) ProvideTopping(ctx context.Context, in *ProvideToppingRequest, opts ...grpc.CallOption) (*ProvideToppingResponse, error) {
	out := new(ProvideToppingResponse)
	err := c.cc.Invoke(ctx, "/toppings.ToppingService/ProvideTopping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToppingServiceServer is the server API for ToppingService service.
type ToppingServiceServer interface {
	ProvideTopping(context.Context, *ProvideToppingRequest) (*ProvideToppingResponse, error)
}

// UnimplementedToppingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedToppingServiceServer struct {
}

func (*UnimplementedToppingServiceServer) ProvideTopping(ctx context.Context, req *ProvideToppingRequest) (*ProvideToppingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvideTopping not implemented")
}

func RegisterToppingServiceServer(s *grpc.Server, srv ToppingServiceServer) {
	s.RegisterService(&_ToppingService_serviceDesc, srv)
}

func _ToppingService_ProvideTopping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvideToppingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToppingServiceServer).ProvideTopping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toppings.ToppingService/ProvideTopping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToppingServiceServer).ProvideTopping(ctx, req.(*ProvideToppingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToppingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "toppings.ToppingService",
	HandlerType: (*ToppingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProvideTopping",
			Handler:    _ToppingService_ProvideTopping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "toppingspb/toppings.proto",
}
