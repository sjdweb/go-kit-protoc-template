// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sprint.proto

/*
Package sprint is a generated protocol buffer package.

It is generated from these files:
	sprint.proto

It has these top-level messages:
	AddSprintRequest
	AddSprintResponse
	CloseSprintRequest
	CloseSprintResponse
	GetSprintRequest
	GetSprintResponse
	Sprint
*/
package sprint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddSprintRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AddSprintRequest) Reset()                    { *m = AddSprintRequest{} }
func (m *AddSprintRequest) String() string            { return proto.CompactTextString(m) }
func (*AddSprintRequest) ProtoMessage()               {}
func (*AddSprintRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddSprintRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddSprintResponse struct {
	Sprint *Sprint `protobuf:"bytes,1,opt,name=sprint" json:"sprint,omitempty"`
	Err    string  `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *AddSprintResponse) Reset()                    { *m = AddSprintResponse{} }
func (m *AddSprintResponse) String() string            { return proto.CompactTextString(m) }
func (*AddSprintResponse) ProtoMessage()               {}
func (*AddSprintResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddSprintResponse) GetSprint() *Sprint {
	if m != nil {
		return m.Sprint
	}
	return nil
}

func (m *AddSprintResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type CloseSprintRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CloseSprintRequest) Reset()                    { *m = CloseSprintRequest{} }
func (m *CloseSprintRequest) String() string            { return proto.CompactTextString(m) }
func (*CloseSprintRequest) ProtoMessage()               {}
func (*CloseSprintRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CloseSprintRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CloseSprintResponse struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *CloseSprintResponse) Reset()                    { *m = CloseSprintResponse{} }
func (m *CloseSprintResponse) String() string            { return proto.CompactTextString(m) }
func (*CloseSprintResponse) ProtoMessage()               {}
func (*CloseSprintResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CloseSprintResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetSprintRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetSprintRequest) Reset()                    { *m = GetSprintRequest{} }
func (m *GetSprintRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSprintRequest) ProtoMessage()               {}
func (*GetSprintRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetSprintRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetSprintResponse struct {
	Sprint *Sprint `protobuf:"bytes,1,opt,name=sprint" json:"sprint,omitempty"`
	Err    string  `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *GetSprintResponse) Reset()                    { *m = GetSprintResponse{} }
func (m *GetSprintResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSprintResponse) ProtoMessage()               {}
func (*GetSprintResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetSprintResponse) GetSprint() *Sprint {
	if m != nil {
		return m.Sprint
	}
	return nil
}

func (m *GetSprintResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type Sprint struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt uint32 `protobuf:"varint,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Sprint) Reset()                    { *m = Sprint{} }
func (m *Sprint) String() string            { return proto.CompactTextString(m) }
func (*Sprint) ProtoMessage()               {}
func (*Sprint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Sprint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Sprint) GetCreatedAt() uint32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Sprint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*AddSprintRequest)(nil), "sprint.AddSprintRequest")
	proto.RegisterType((*AddSprintResponse)(nil), "sprint.AddSprintResponse")
	proto.RegisterType((*CloseSprintRequest)(nil), "sprint.CloseSprintRequest")
	proto.RegisterType((*CloseSprintResponse)(nil), "sprint.CloseSprintResponse")
	proto.RegisterType((*GetSprintRequest)(nil), "sprint.GetSprintRequest")
	proto.RegisterType((*GetSprintResponse)(nil), "sprint.GetSprintResponse")
	proto.RegisterType((*Sprint)(nil), "sprint.Sprint")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SprintService service

type SprintServiceClient interface {
	AddSprint(ctx context.Context, in *AddSprintRequest, opts ...grpc.CallOption) (*AddSprintResponse, error)
	CloseSprint(ctx context.Context, in *CloseSprintRequest, opts ...grpc.CallOption) (*CloseSprintResponse, error)
	GetSprint(ctx context.Context, in *GetSprintRequest, opts ...grpc.CallOption) (*GetSprintResponse, error)
}

type sprintServiceClient struct {
	cc *grpc.ClientConn
}

func NewSprintServiceClient(cc *grpc.ClientConn) SprintServiceClient {
	return &sprintServiceClient{cc}
}

func (c *sprintServiceClient) AddSprint(ctx context.Context, in *AddSprintRequest, opts ...grpc.CallOption) (*AddSprintResponse, error) {
	out := new(AddSprintResponse)
	err := grpc.Invoke(ctx, "/sprint.SprintService/AddSprint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintServiceClient) CloseSprint(ctx context.Context, in *CloseSprintRequest, opts ...grpc.CallOption) (*CloseSprintResponse, error) {
	out := new(CloseSprintResponse)
	err := grpc.Invoke(ctx, "/sprint.SprintService/CloseSprint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintServiceClient) GetSprint(ctx context.Context, in *GetSprintRequest, opts ...grpc.CallOption) (*GetSprintResponse, error) {
	out := new(GetSprintResponse)
	err := grpc.Invoke(ctx, "/sprint.SprintService/GetSprint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SprintService service

type SprintServiceServer interface {
	AddSprint(context.Context, *AddSprintRequest) (*AddSprintResponse, error)
	CloseSprint(context.Context, *CloseSprintRequest) (*CloseSprintResponse, error)
	GetSprint(context.Context, *GetSprintRequest) (*GetSprintResponse, error)
}

func RegisterSprintServiceServer(s *grpc.Server, srv SprintServiceServer) {
	s.RegisterService(&_SprintService_serviceDesc, srv)
}

func _SprintService_AddSprint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSprintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SprintServiceServer).AddSprint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sprint.SprintService/AddSprint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SprintServiceServer).AddSprint(ctx, req.(*AddSprintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SprintService_CloseSprint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSprintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SprintServiceServer).CloseSprint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sprint.SprintService/CloseSprint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SprintServiceServer).CloseSprint(ctx, req.(*CloseSprintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SprintService_GetSprint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSprintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SprintServiceServer).GetSprint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sprint.SprintService/GetSprint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SprintServiceServer).GetSprint(ctx, req.(*GetSprintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SprintService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sprint.SprintService",
	HandlerType: (*SprintServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSprint",
			Handler:    _SprintService_AddSprint_Handler,
		},
		{
			MethodName: "CloseSprint",
			Handler:    _SprintService_CloseSprint_Handler,
		},
		{
			MethodName: "GetSprint",
			Handler:    _SprintService_GetSprint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sprint.proto",
}

func init() { proto.RegisterFile("sprint.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x0b, 0x35, 0x24, 0xfc, 0xb5, 0x0d, 0x1d, 0x2f, 0x88, 0x31, 0x31, 0x1b, 0x53, 0x3d,
	0xf5, 0x50, 0x9f, 0xa0, 0xf5, 0xa0, 0x89, 0xe9, 0x85, 0x3e, 0x80, 0xc1, 0x32, 0x07, 0x12, 0x05,
	0xdc, 0x5d, 0x7d, 0x5f, 0xdf, 0xc4, 0xb8, 0x6c, 0x59, 0x4a, 0x1b, 0x2f, 0xbd, 0xcd, 0x30, 0x3f,
	0xff, 0x3f, 0xf3, 0x01, 0xce, 0x55, 0x2d, 0x8b, 0x52, 0xcf, 0x6b, 0x59, 0xe9, 0x8a, 0x82, 0xa6,
	0x13, 0x33, 0x44, 0xcb, 0x3c, 0xdf, 0x98, 0x26, 0xe5, 0xcf, 0x2f, 0x56, 0x9a, 0x08, 0x67, 0x65,
	0xf6, 0xc1, 0xb1, 0x77, 0xe3, 0xdd, 0x87, 0xa9, 0xa9, 0xc5, 0x1a, 0xd3, 0x8e, 0x4e, 0xd5, 0x55,
	0xa9, 0x98, 0x66, 0xb0, 0x36, 0x46, 0x3a, 0x5a, 0x4c, 0xe6, 0x36, 0xc3, 0xea, 0xec, 0x94, 0x22,
	0x0c, 0x59, 0xca, 0xd8, 0x37, 0x7e, 0x7f, 0xa5, 0xb8, 0x05, 0x3d, 0xbe, 0x57, 0x8a, 0xf7, 0x83,
	0x27, 0xf0, 0x8b, 0xdc, 0xc6, 0xfa, 0x45, 0x2e, 0xee, 0x70, 0xb1, 0xa7, 0xb2, 0xb1, 0xd6, 0xce,
	0x73, 0x76, 0x02, 0xd1, 0x13, 0xeb, 0xff, 0xcd, 0xd6, 0x98, 0x76, 0x34, 0x27, 0x5f, 0xf0, 0x82,
	0xa0, 0xd1, 0xf4, 0x83, 0xe8, 0x1a, 0xd8, 0x4a, 0xce, 0x34, 0xe7, 0xaf, 0x99, 0x36, 0xaf, 0x8c,
	0xd3, 0xd0, 0x3e, 0x59, 0x3a, 0xba, 0x43, 0x47, 0x77, 0xf1, 0xe3, 0x61, 0xdc, 0xb8, 0x6d, 0x58,
	0x7e, 0x17, 0x5b, 0xa6, 0x15, 0xc2, 0x96, 0x37, 0xc5, 0xbb, 0xad, 0xfa, 0x9f, 0x2a, 0xb9, 0x3c,
	0x32, 0x69, 0x4e, 0x13, 0x03, 0x7a, 0xc6, 0xa8, 0x83, 0x8f, 0x92, 0x9d, 0xf6, 0x90, 0x7c, 0x72,
	0x75, 0x74, 0xd6, 0x3a, 0xad, 0x10, 0xb6, 0xec, 0xdc, 0x36, 0x7d, 0xe4, 0x6e, 0x9b, 0x03, 0xd0,
	0x62, 0xf0, 0x16, 0x98, 0x1f, 0xef, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xef, 0x45, 0x79, 0x37,
	0x88, 0x02, 0x00, 0x00,
}