// Code generated by protoc-gen-go. DO NOT EDIT.
// source: icho.proto

package icho

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetEchoRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEchoRequest) Reset()         { *m = GetEchoRequest{} }
func (m *GetEchoRequest) String() string { return proto.CompactTextString(m) }
func (*GetEchoRequest) ProtoMessage()    {}
func (*GetEchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_556f94487289403c, []int{0}
}

func (m *GetEchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEchoRequest.Unmarshal(m, b)
}
func (m *GetEchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEchoRequest.Marshal(b, m, deterministic)
}
func (m *GetEchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEchoRequest.Merge(m, src)
}
func (m *GetEchoRequest) XXX_Size() int {
	return xxx_messageInfo_GetEchoRequest.Size(m)
}
func (m *GetEchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEchoRequest proto.InternalMessageInfo

func (m *GetEchoRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type HealthResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_556f94487289403c, []int{1}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type PostEchoRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostEchoRequest) Reset()         { *m = PostEchoRequest{} }
func (m *PostEchoRequest) String() string { return proto.CompactTextString(m) }
func (*PostEchoRequest) ProtoMessage()    {}
func (*PostEchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_556f94487289403c, []int{2}
}

func (m *PostEchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostEchoRequest.Unmarshal(m, b)
}
func (m *PostEchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostEchoRequest.Marshal(b, m, deterministic)
}
func (m *PostEchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostEchoRequest.Merge(m, src)
}
func (m *PostEchoRequest) XXX_Size() int {
	return xxx_messageInfo_PostEchoRequest.Size(m)
}
func (m *PostEchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostEchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostEchoRequest proto.InternalMessageInfo

func (m *PostEchoRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type PostEchoResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostEchoResponse) Reset()         { *m = PostEchoResponse{} }
func (m *PostEchoResponse) String() string { return proto.CompactTextString(m) }
func (*PostEchoResponse) ProtoMessage()    {}
func (*PostEchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_556f94487289403c, []int{3}
}

func (m *PostEchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostEchoResponse.Unmarshal(m, b)
}
func (m *PostEchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostEchoResponse.Marshal(b, m, deterministic)
}
func (m *PostEchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostEchoResponse.Merge(m, src)
}
func (m *PostEchoResponse) XXX_Size() int {
	return xxx_messageInfo_PostEchoResponse.Size(m)
}
func (m *PostEchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostEchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostEchoResponse proto.InternalMessageInfo

func (m *PostEchoResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*GetEchoRequest)(nil), "icho.GetEchoRequest")
	proto.RegisterType((*HealthResponse)(nil), "icho.HealthResponse")
	proto.RegisterType((*PostEchoRequest)(nil), "icho.PostEchoRequest")
	proto.RegisterType((*PostEchoResponse)(nil), "icho.PostEchoResponse")
}

func init() { proto.RegisterFile("icho.proto", fileDescriptor_556f94487289403c) }

var fileDescriptor_556f94487289403c = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0x41, 0x4e, 0xeb, 0x30,
	0x10, 0x55, 0xbe, 0xfa, 0x9b, 0x66, 0x50, 0x5b, 0x64, 0x20, 0x42, 0x81, 0x05, 0xb2, 0x40, 0xea,
	0xca, 0x91, 0x60, 0xc7, 0xbe, 0x14, 0x58, 0x20, 0x14, 0x4e, 0x90, 0x46, 0x43, 0x1d, 0xa9, 0x64,
	0x42, 0x3c, 0x41, 0xb0, 0xe5, 0x0a, 0x1c, 0x8d, 0x2b, 0x70, 0x01, 0x6e, 0x80, 0x62, 0xbb, 0xa0,
	0x76, 0x01, 0xbb, 0x79, 0xf6, 0x9b, 0xf7, 0xde, 0x3c, 0x80, 0xb2, 0xd0, 0xa4, 0xea, 0x86, 0x98,
	0x44, 0xaf, 0x9b, 0x93, 0xc3, 0x05, 0xd1, 0x62, 0x89, 0x69, 0x5e, 0x97, 0x69, 0x5e, 0x55, 0xc4,
	0x39, 0x97, 0x54, 0x19, 0xc7, 0x49, 0x0e, 0xfc, 0xaf, 0x45, 0xf3, 0xf6, 0x3e, 0xc5, 0x87, 0x9a,
	0x5f, 0xdc, 0xa7, 0x3c, 0x86, 0xd1, 0x0c, 0x79, 0x5a, 0x68, 0xca, 0xf0, 0xb1, 0x45, 0xc3, 0x42,
	0x40, 0x8f, 0xf1, 0x99, 0xf7, 0x83, 0xa3, 0x60, 0x12, 0x65, 0x76, 0x96, 0x13, 0x18, 0x5d, 0x62,
	0xbe, 0x64, 0x9d, 0xa1, 0xa9, 0xa9, 0x32, 0x28, 0x62, 0xe8, 0x1b, 0xce, 0xb9, 0x35, 0x9e, 0xe7,
	0x91, 0x3c, 0x81, 0xf1, 0x2d, 0x99, 0x3f, 0x05, 0x15, 0x6c, 0xff, 0xd0, 0xbc, 0x64, 0x02, 0x83,
	0xc6, 0xcf, 0x9e, 0xfb, 0x8d, 0x4f, 0x3f, 0x03, 0xd8, 0xba, 0x2a, 0x34, 0xdd, 0x61, 0xf3, 0x54,
	0x16, 0x28, 0x2e, 0x20, 0xf4, 0xb1, 0xc5, 0xae, 0xb2, 0x7d, 0xac, 0x5f, 0x91, 0xc4, 0xee, 0x75,
	0xd3, 0x44, 0x0e, 0x5f, 0xdf, 0x3f, 0xde, 0xfe, 0x85, 0xe2, 0x7f, 0x8a, 0xdd, 0xf2, 0x35, 0x44,
	0x33, 0x64, 0x77, 0x9b, 0x88, 0x95, 0x6b, 0x4a, 0xad, 0x9a, 0x52, 0xd3, 0xae, 0xa9, 0xc4, 0x3b,
	0xac, 0x37, 0x20, 0xc7, 0x56, 0x29, 0x12, 0x61, 0xaa, 0xdd, 0xfa, 0x0d, 0x0c, 0x56, 0x76, 0x62,
	0x6f, 0xd3, 0xfe, 0xf7, 0x54, 0x3b, 0x56, 0x6b, 0x28, 0x5d, 0xaa, 0x73, 0xdb, 0xd1, 0xbc, 0x6f,
	0x63, 0x9c, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x77, 0x29, 0x96, 0x7c, 0xf0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IchoServiceClient is the client API for IchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IchoServiceClient interface {
	// Echo request body
	//
	// Echo request
	GetEcho(ctx context.Context, in *GetEchoRequest, opts ...grpc.CallOption) (*PostEchoResponse, error)
	// health
	//
	// Health of service
	GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResponse, error)
	// Post for Echo
	//
	// Post for Echo
	PostEcho(ctx context.Context, in *PostEchoRequest, opts ...grpc.CallOption) (*PostEchoResponse, error)
}

type ichoServiceClient struct {
	cc *grpc.ClientConn
}

func NewIchoServiceClient(cc *grpc.ClientConn) IchoServiceClient {
	return &ichoServiceClient{cc}
}

func (c *ichoServiceClient) GetEcho(ctx context.Context, in *GetEchoRequest, opts ...grpc.CallOption) (*PostEchoResponse, error) {
	out := new(PostEchoResponse)
	err := c.cc.Invoke(ctx, "/icho.IchoService/GetEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ichoServiceClient) GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/icho.IchoService/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ichoServiceClient) PostEcho(ctx context.Context, in *PostEchoRequest, opts ...grpc.CallOption) (*PostEchoResponse, error) {
	out := new(PostEchoResponse)
	err := c.cc.Invoke(ctx, "/icho.IchoService/PostEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IchoServiceServer is the server API for IchoService service.
type IchoServiceServer interface {
	// Echo request body
	//
	// Echo request
	GetEcho(context.Context, *GetEchoRequest) (*PostEchoResponse, error)
	// health
	//
	// Health of service
	GetHealth(context.Context, *empty.Empty) (*HealthResponse, error)
	// Post for Echo
	//
	// Post for Echo
	PostEcho(context.Context, *PostEchoRequest) (*PostEchoResponse, error)
}

// UnimplementedIchoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIchoServiceServer struct {
}

func (*UnimplementedIchoServiceServer) GetEcho(ctx context.Context, req *GetEchoRequest) (*PostEchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEcho not implemented")
}
func (*UnimplementedIchoServiceServer) GetHealth(ctx context.Context, req *empty.Empty) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (*UnimplementedIchoServiceServer) PostEcho(ctx context.Context, req *PostEchoRequest) (*PostEchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostEcho not implemented")
}

func RegisterIchoServiceServer(s *grpc.Server, srv IchoServiceServer) {
	s.RegisterService(&_IchoService_serviceDesc, srv)
}

func _IchoService_GetEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IchoServiceServer).GetEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/icho.IchoService/GetEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IchoServiceServer).GetEcho(ctx, req.(*GetEchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IchoService_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IchoServiceServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/icho.IchoService/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IchoServiceServer).GetHealth(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _IchoService_PostEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostEchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IchoServiceServer).PostEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/icho.IchoService/PostEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IchoServiceServer).PostEcho(ctx, req.(*PostEchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "icho.IchoService",
	HandlerType: (*IchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEcho",
			Handler:    _IchoService_GetEcho_Handler,
		},
		{
			MethodName: "GetHealth",
			Handler:    _IchoService_GetHealth_Handler,
		},
		{
			MethodName: "PostEcho",
			Handler:    _IchoService_PostEcho_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "icho.proto",
}
