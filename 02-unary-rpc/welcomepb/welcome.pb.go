// Code generated by protoc-gen-go. DO NOT EDIT.
// source: welcomepb/welcome.proto

package welcomepb

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

type Person struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_46790a53682c3a6a, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type WelcomeRequest struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WelcomeRequest) Reset()         { *m = WelcomeRequest{} }
func (m *WelcomeRequest) String() string { return proto.CompactTextString(m) }
func (*WelcomeRequest) ProtoMessage()    {}
func (*WelcomeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46790a53682c3a6a, []int{1}
}

func (m *WelcomeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WelcomeRequest.Unmarshal(m, b)
}
func (m *WelcomeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WelcomeRequest.Marshal(b, m, deterministic)
}
func (m *WelcomeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WelcomeRequest.Merge(m, src)
}
func (m *WelcomeRequest) XXX_Size() int {
	return xxx_messageInfo_WelcomeRequest.Size(m)
}
func (m *WelcomeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WelcomeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WelcomeRequest proto.InternalMessageInfo

func (m *WelcomeRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type WelcomeResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WelcomeResponse) Reset()         { *m = WelcomeResponse{} }
func (m *WelcomeResponse) String() string { return proto.CompactTextString(m) }
func (*WelcomeResponse) ProtoMessage()    {}
func (*WelcomeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46790a53682c3a6a, []int{2}
}

func (m *WelcomeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WelcomeResponse.Unmarshal(m, b)
}
func (m *WelcomeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WelcomeResponse.Marshal(b, m, deterministic)
}
func (m *WelcomeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WelcomeResponse.Merge(m, src)
}
func (m *WelcomeResponse) XXX_Size() int {
	return xxx_messageInfo_WelcomeResponse.Size(m)
}
func (m *WelcomeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WelcomeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WelcomeResponse proto.InternalMessageInfo

func (m *WelcomeResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "welcome.Person")
	proto.RegisterType((*WelcomeRequest)(nil), "welcome.WelcomeRequest")
	proto.RegisterType((*WelcomeResponse)(nil), "welcome.WelcomeResponse")
}

func init() { proto.RegisterFile("welcomepb/welcome.proto", fileDescriptor_46790a53682c3a6a) }

var fileDescriptor_46790a53682c3a6a = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4f, 0xcd, 0x49,
	0xce, 0xcf, 0x4d, 0x2d, 0x48, 0xd2, 0x87, 0xb2, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8,
	0xa1, 0x5c, 0x25, 0x17, 0x2e, 0xb6, 0x80, 0xd4, 0xa2, 0xe2, 0xfc, 0x3c, 0x21, 0x59, 0x2e, 0xae,
	0xb4, 0xcc, 0xa2, 0xe2, 0x92, 0xf8, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x4e, 0xb0, 0x88, 0x5f, 0x62, 0x6e, 0xaa, 0x90, 0x34, 0x17, 0x67, 0x4e, 0x22, 0x4c, 0x96,
	0x09, 0x2c, 0xcb, 0x01, 0x12, 0x00, 0x49, 0x2a, 0x59, 0x72, 0xf1, 0x85, 0x43, 0x0c, 0x0c, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe7, 0x62, 0x2b, 0x00, 0x9b, 0x0b, 0x36, 0x89, 0xdb,
	0x88, 0x5f, 0x0f, 0xe6, 0x00, 0x88, 0x75, 0x41, 0x50, 0x69, 0x25, 0x4d, 0x2e, 0x7e, 0xb8, 0xd6,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12,
	0xa8, 0x2b, 0xa0, 0x3c, 0xa3, 0x00, 0xb8, 0x2d, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42,
	0x76, 0x5c, 0xec, 0x50, 0x11, 0x21, 0x71, 0xb8, 0x05, 0xa8, 0x2e, 0x91, 0x92, 0xc0, 0x94, 0x80,
	0xd8, 0xa3, 0xc4, 0xe0, 0xc4, 0x1d, 0xc5, 0x09, 0x0f, 0xa1, 0x24, 0x36, 0x70, 0xd0, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x30, 0x20, 0x90, 0x35, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WelcomeServiceClient is the client API for WelcomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WelcomeServiceClient interface {
	Welcome(ctx context.Context, in *WelcomeRequest, opts ...grpc.CallOption) (*WelcomeResponse, error)
}

type welcomeServiceClient struct {
	cc *grpc.ClientConn
}

func NewWelcomeServiceClient(cc *grpc.ClientConn) WelcomeServiceClient {
	return &welcomeServiceClient{cc}
}

func (c *welcomeServiceClient) Welcome(ctx context.Context, in *WelcomeRequest, opts ...grpc.CallOption) (*WelcomeResponse, error) {
	out := new(WelcomeResponse)
	err := c.cc.Invoke(ctx, "/welcome.WelcomeService/Welcome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WelcomeServiceServer is the server API for WelcomeService service.
type WelcomeServiceServer interface {
	Welcome(context.Context, *WelcomeRequest) (*WelcomeResponse, error)
}

// UnimplementedWelcomeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWelcomeServiceServer struct {
}

func (*UnimplementedWelcomeServiceServer) Welcome(ctx context.Context, req *WelcomeRequest) (*WelcomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Welcome not implemented")
}

func RegisterWelcomeServiceServer(s *grpc.Server, srv WelcomeServiceServer) {
	s.RegisterService(&_WelcomeService_serviceDesc, srv)
}

func _WelcomeService_Welcome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WelcomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WelcomeServiceServer).Welcome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/welcome.WelcomeService/Welcome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WelcomeServiceServer).Welcome(ctx, req.(*WelcomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WelcomeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "welcome.WelcomeService",
	HandlerType: (*WelcomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Welcome",
			Handler:    _WelcomeService_Welcome_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "welcomepb/welcome.proto",
}
