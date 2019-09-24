// Code generated by protoc-gen-go. DO NOT EDIT.
// source: output.proto

package output

import (
	context "context"
	fmt "fmt"
	schema "github.com/falcosecurity/client-go/pkg/api/schema"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Request struct {
	Keepalive            bool     `protobuf:"varint,1,opt,name=keepalive,proto3" json:"keepalive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b2b3ae2e703b013, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetKeepalive() bool {
	if m != nil {
		return m.Keepalive
	}
	return false
}

type Response struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Priority             schema.Priority      `protobuf:"varint,2,opt,name=priority,proto3,enum=falco.schema.Priority" json:"priority,omitempty"`
	Source               schema.Source        `protobuf:"varint,3,opt,name=source,proto3,enum=falco.schema.Source" json:"source,omitempty"`
	Rule                 string               `protobuf:"bytes,4,opt,name=rule,proto3" json:"rule,omitempty"`
	Output               string               `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
	OutputFields         map[string]string    `protobuf:"bytes,6,rep,name=output_fields,json=outputFields,proto3" json:"output_fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b2b3ae2e703b013, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Response) GetPriority() schema.Priority {
	if m != nil {
		return m.Priority
	}
	return schema.Priority_EMERGENCY
}

func (m *Response) GetSource() schema.Source {
	if m != nil {
		return m.Source
	}
	return schema.Source_SYSCALL
}

func (m *Response) GetRule() string {
	if m != nil {
		return m.Rule
	}
	return ""
}

func (m *Response) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *Response) GetOutputFields() map[string]string {
	if m != nil {
		return m.OutputFields
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "falco.output.request")
	proto.RegisterType((*Response)(nil), "falco.output.response")
	proto.RegisterMapType((map[string]string)(nil), "falco.output.response.OutputFieldsEntry")
}

func init() { proto.RegisterFile("output.proto", fileDescriptor_0b2b3ae2e703b013) }

var fileDescriptor_0b2b3ae2e703b013 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcb, 0x6b, 0xdb, 0x40,
	0x10, 0xc6, 0x91, 0x1f, 0xb2, 0xb5, 0x76, 0x4b, 0xbb, 0xb8, 0x46, 0x88, 0x42, 0x85, 0x2f, 0xd5,
	0xa1, 0x5d, 0xb5, 0xf2, 0xa5, 0x94, 0x40, 0x20, 0x90, 0xe4, 0x14, 0x02, 0x4b, 0x4e, 0xb9, 0x04,
	0x69, 0x33, 0x96, 0x17, 0x4b, 0xde, 0xcd, 0x3e, 0x0c, 0xfe, 0x4f, 0xf3, 0xe7, 0x04, 0xef, 0xda,
	0xce, 0xf3, 0x36, 0xf3, 0xcd, 0x6f, 0x34, 0x9f, 0xbe, 0x45, 0x63, 0x61, 0x8d, 0xb4, 0x86, 0x48,
	0x25, 0x8c, 0xc0, 0xe3, 0x45, 0xd9, 0x30, 0x41, 0xbc, 0x96, 0xfc, 0xa8, 0x85, 0xa8, 0x1b, 0xc8,
	0xdd, 0xac, 0xb2, 0x8b, 0xdc, 0xf0, 0x16, 0xb4, 0x29, 0x5b, 0xe9, 0xf1, 0x64, 0xac, 0xd9, 0x12,
	0xda, 0xd2, 0x77, 0xb3, 0x9f, 0x68, 0xa0, 0xe0, 0xc1, 0x82, 0x36, 0xf8, 0x3b, 0x8a, 0x56, 0x00,
	0xb2, 0x6c, 0xf8, 0x06, 0xe2, 0x20, 0x0d, 0xb2, 0x21, 0x7d, 0x16, 0x66, 0x8f, 0x1d, 0x34, 0x54,
	0xa0, 0xa5, 0x58, 0x6b, 0xc0, 0x04, 0xf5, 0x76, 0x9f, 0x75, 0xd4, 0xa8, 0x48, 0x88, 0xbf, 0x49,
	0x0e, 0x37, 0xc9, 0xcd, 0xe1, 0x26, 0x75, 0x1c, 0x2e, 0xd0, 0x50, 0x2a, 0x2e, 0x14, 0x37, 0xdb,
	0xb8, 0x93, 0x06, 0xd9, 0xe7, 0x62, 0x4a, 0xbc, 0xeb, 0xa3, 0x19, 0x3f, 0xa5, 0x47, 0x0e, 0xff,
	0x42, 0xa1, 0x16, 0x56, 0x31, 0x88, 0xbb, 0x6e, 0x63, 0xf2, 0x7a, 0xc3, 0xcf, 0xe8, 0x9e, 0xc1,
	0x18, 0xf5, 0x94, 0x6d, 0x20, 0xee, 0xa5, 0x41, 0x16, 0x51, 0x57, 0xe3, 0x29, 0x0a, 0x7d, 0x28,
	0x71, 0xdf, 0xa9, 0xfb, 0x0e, 0x5f, 0xa1, 0x4f, 0xbe, 0xba, 0x5b, 0x70, 0x68, 0xee, 0x75, 0x1c,
	0xa6, 0xdd, 0x6c, 0x54, 0x64, 0xe4, 0x65, 0x90, 0xe4, 0xf0, 0xb3, 0xe4, 0xda, 0xf5, 0x17, 0x0e,
	0x3d, 0x5f, 0x1b, 0xb5, 0xa5, 0xfb, 0xfc, 0xbd, 0x94, 0x9c, 0xa2, 0xaf, 0xef, 0x10, 0xfc, 0x05,
	0x75, 0x57, 0xb0, 0x75, 0x01, 0x45, 0x74, 0x57, 0xe2, 0x09, 0xea, 0x6f, 0xca, 0xc6, 0x82, 0x0b,
	0x20, 0xa2, 0xbe, 0xf9, 0xdf, 0xf9, 0x17, 0x14, 0x97, 0x68, 0xa0, 0x41, 0x6d, 0x38, 0x03, 0x7c,
	0x82, 0x22, 0x6d, 0x2b, 0xcd, 0x14, 0xaf, 0x00, 0x7f, 0x7b, 0x6b, 0xc8, 0xbd, 0x53, 0x32, 0xfd,
	0xd8, 0xe7, 0x9f, 0xe0, 0x6c, 0x7e, 0xfb, 0xb7, 0xe6, 0x66, 0x69, 0x2b, 0xc2, 0x44, 0x9b, 0x3b,
	0x4a, 0x03, 0xb3, 0xbb, 0x38, 0x73, 0xd6, 0x70, 0x58, 0x9b, 0xdf, 0xb5, 0xc8, 0xe5, 0xaa, 0xce,
	0x4b, 0xc9, 0x73, 0xbf, 0x5f, 0x85, 0xee, 0xd5, 0xe6, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x97,
	0x95, 0xea, 0x3b, 0x55, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (Service_SubscribeClient, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (Service_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[0], "/falco.output.service/subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_SubscribeClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type serviceSubscribeClient struct {
	grpc.ClientStream
}

func (x *serviceSubscribeClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Subscribe(*Request, Service_SubscribeServer) error
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Subscribe(req *Request, srv Service_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Subscribe(m, &serviceSubscribeServer{stream})
}

type Service_SubscribeServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type serviceSubscribeServer struct {
	grpc.ServerStream
}

func (x *serviceSubscribeServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "falco.output.service",
	HandlerType: (*ServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "subscribe",
			Handler:       _Service_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "output.proto",
}