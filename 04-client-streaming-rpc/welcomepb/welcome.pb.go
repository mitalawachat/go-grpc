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

type LongWelcomeRequest struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongWelcomeRequest) Reset()         { *m = LongWelcomeRequest{} }
func (m *LongWelcomeRequest) String() string { return proto.CompactTextString(m) }
func (*LongWelcomeRequest) ProtoMessage()    {}
func (*LongWelcomeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46790a53682c3a6a, []int{1}
}

func (m *LongWelcomeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongWelcomeRequest.Unmarshal(m, b)
}
func (m *LongWelcomeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongWelcomeRequest.Marshal(b, m, deterministic)
}
func (m *LongWelcomeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongWelcomeRequest.Merge(m, src)
}
func (m *LongWelcomeRequest) XXX_Size() int {
	return xxx_messageInfo_LongWelcomeRequest.Size(m)
}
func (m *LongWelcomeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LongWelcomeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LongWelcomeRequest proto.InternalMessageInfo

func (m *LongWelcomeRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type LongWelcomeResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongWelcomeResponse) Reset()         { *m = LongWelcomeResponse{} }
func (m *LongWelcomeResponse) String() string { return proto.CompactTextString(m) }
func (*LongWelcomeResponse) ProtoMessage()    {}
func (*LongWelcomeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46790a53682c3a6a, []int{2}
}

func (m *LongWelcomeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongWelcomeResponse.Unmarshal(m, b)
}
func (m *LongWelcomeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongWelcomeResponse.Marshal(b, m, deterministic)
}
func (m *LongWelcomeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongWelcomeResponse.Merge(m, src)
}
func (m *LongWelcomeResponse) XXX_Size() int {
	return xxx_messageInfo_LongWelcomeResponse.Size(m)
}
func (m *LongWelcomeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LongWelcomeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LongWelcomeResponse proto.InternalMessageInfo

func (m *LongWelcomeResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "welcome.Person")
	proto.RegisterType((*LongWelcomeRequest)(nil), "welcome.LongWelcomeRequest")
	proto.RegisterType((*LongWelcomeResponse)(nil), "welcome.LongWelcomeResponse")
}

func init() { proto.RegisterFile("welcomepb/welcome.proto", fileDescriptor_46790a53682c3a6a) }

var fileDescriptor_46790a53682c3a6a = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4f, 0xcd, 0x49,
	0xce, 0xcf, 0x4d, 0x2d, 0x48, 0xd2, 0x87, 0xb2, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8,
	0xa1, 0x5c, 0x25, 0x17, 0x2e, 0xb6, 0x80, 0xd4, 0xa2, 0xe2, 0xfc, 0x3c, 0x21, 0x59, 0x2e, 0xae,
	0xb4, 0xcc, 0xa2, 0xe2, 0x92, 0xf8, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x4e, 0xb0, 0x88, 0x5f, 0x62, 0x6e, 0xaa, 0x90, 0x34, 0x17, 0x67, 0x4e, 0x22, 0x4c, 0x96,
	0x09, 0x2c, 0xcb, 0x01, 0x12, 0x00, 0x49, 0x2a, 0xd9, 0x72, 0x09, 0xf9, 0xe4, 0xe7, 0xa5, 0x87,
	0x43, 0x0c, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe7, 0x62, 0x2b, 0x00, 0x9b,
	0x0d, 0x36, 0x8d, 0xdb, 0x88, 0x5f, 0x0f, 0xe6, 0x08, 0x88, 0x95, 0x41, 0x50, 0x69, 0x25, 0x5d,
	0x2e, 0x61, 0x14, 0xed, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c, 0x45, 0xa9,
	0xc5, 0xa5, 0x39, 0x25, 0x50, 0xd7, 0x40, 0x79, 0x46, 0x71, 0x5c, 0x7c, 0x50, 0xa5, 0xc1, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x3e, 0x5c, 0xdc, 0x48, 0x06, 0x08, 0x49, 0xc3, 0x2d, 0xc2,
	0x74, 0x95, 0x94, 0x0c, 0x76, 0x49, 0x88, 0x9d, 0x4a, 0x0c, 0x1a, 0x8c, 0x4e, 0xdc, 0x51, 0x9c,
	0xf0, 0x70, 0x4b, 0x62, 0x03, 0x07, 0x98, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x69, 0x71, 0x49,
	0x39, 0x4b, 0x01, 0x00, 0x00,
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
	LongWelcome(ctx context.Context, opts ...grpc.CallOption) (WelcomeService_LongWelcomeClient, error)
}

type welcomeServiceClient struct {
	cc *grpc.ClientConn
}

func NewWelcomeServiceClient(cc *grpc.ClientConn) WelcomeServiceClient {
	return &welcomeServiceClient{cc}
}

func (c *welcomeServiceClient) LongWelcome(ctx context.Context, opts ...grpc.CallOption) (WelcomeService_LongWelcomeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WelcomeService_serviceDesc.Streams[0], "/welcome.WelcomeService/LongWelcome", opts...)
	if err != nil {
		return nil, err
	}
	x := &welcomeServiceLongWelcomeClient{stream}
	return x, nil
}

type WelcomeService_LongWelcomeClient interface {
	Send(*LongWelcomeRequest) error
	CloseAndRecv() (*LongWelcomeResponse, error)
	grpc.ClientStream
}

type welcomeServiceLongWelcomeClient struct {
	grpc.ClientStream
}

func (x *welcomeServiceLongWelcomeClient) Send(m *LongWelcomeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *welcomeServiceLongWelcomeClient) CloseAndRecv() (*LongWelcomeResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongWelcomeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WelcomeServiceServer is the server API for WelcomeService service.
type WelcomeServiceServer interface {
	LongWelcome(WelcomeService_LongWelcomeServer) error
}

// UnimplementedWelcomeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWelcomeServiceServer struct {
}

func (*UnimplementedWelcomeServiceServer) LongWelcome(srv WelcomeService_LongWelcomeServer) error {
	return status.Errorf(codes.Unimplemented, "method LongWelcome not implemented")
}

func RegisterWelcomeServiceServer(s *grpc.Server, srv WelcomeServiceServer) {
	s.RegisterService(&_WelcomeService_serviceDesc, srv)
}

func _WelcomeService_LongWelcome_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WelcomeServiceServer).LongWelcome(&welcomeServiceLongWelcomeServer{stream})
}

type WelcomeService_LongWelcomeServer interface {
	SendAndClose(*LongWelcomeResponse) error
	Recv() (*LongWelcomeRequest, error)
	grpc.ServerStream
}

type welcomeServiceLongWelcomeServer struct {
	grpc.ServerStream
}

func (x *welcomeServiceLongWelcomeServer) SendAndClose(m *LongWelcomeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *welcomeServiceLongWelcomeServer) Recv() (*LongWelcomeRequest, error) {
	m := new(LongWelcomeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _WelcomeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "welcome.WelcomeService",
	HandlerType: (*WelcomeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LongWelcome",
			Handler:       _WelcomeService_LongWelcome_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "welcomepb/welcome.proto",
}