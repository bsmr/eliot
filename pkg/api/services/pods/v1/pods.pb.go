// Code generated by protoc-gen-go.
// source: pods.proto
// DO NOT EDIT!

/*
Package pods is a generated protocol buffer package.

It is generated from these files:
	pods.proto

It has these top-level messages:
	CreatePodRequest
	CreatePodResponse
	ListPodsRequest
	ListPodsResponse
	GetLogsRequest
	GetLogsResponse
	Pod
	PodSpec
	Container
	PodStatus
	ContainerStatus
*/
package pods

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

type GetLogsResponse_Type int32

const (
	GetLogsResponse_STDOUT GetLogsResponse_Type = 0
	GetLogsResponse_STDERR GetLogsResponse_Type = 1
)

var GetLogsResponse_Type_name = map[int32]string{
	0: "STDOUT",
	1: "STDERR",
}
var GetLogsResponse_Type_value = map[string]int32{
	"STDOUT": 0,
	"STDERR": 1,
}

func (x GetLogsResponse_Type) String() string {
	return proto.EnumName(GetLogsResponse_Type_name, int32(x))
}
func (GetLogsResponse_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type CreatePodRequest struct {
	Pod *Pod `protobuf:"bytes,1,opt,name=pod" json:"pod,omitempty"`
}

func (m *CreatePodRequest) Reset()                    { *m = CreatePodRequest{} }
func (m *CreatePodRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePodRequest) ProtoMessage()               {}
func (*CreatePodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreatePodRequest) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

type CreatePodResponse struct {
	Pod *Pod `protobuf:"bytes,1,opt,name=pod" json:"pod,omitempty"`
}

func (m *CreatePodResponse) Reset()                    { *m = CreatePodResponse{} }
func (m *CreatePodResponse) String() string            { return proto.CompactTextString(m) }
func (*CreatePodResponse) ProtoMessage()               {}
func (*CreatePodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreatePodResponse) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

type ListPodsRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ListPodsRequest) Reset()                    { *m = ListPodsRequest{} }
func (m *ListPodsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPodsRequest) ProtoMessage()               {}
func (*ListPodsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListPodsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ListPodsResponse struct {
	Pods []*Pod `protobuf:"bytes,1,rep,name=pods" json:"pods,omitempty"`
}

func (m *ListPodsResponse) Reset()                    { *m = ListPodsResponse{} }
func (m *ListPodsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPodsResponse) ProtoMessage()               {}
func (*ListPodsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListPodsResponse) GetPods() []*Pod {
	if m != nil {
		return m.Pods
	}
	return nil
}

type GetLogsRequest struct {
	Namespace   string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	ContainerID string `protobuf:"bytes,2,opt,name=containerID" json:"containerID,omitempty"`
}

func (m *GetLogsRequest) Reset()                    { *m = GetLogsRequest{} }
func (m *GetLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLogsRequest) ProtoMessage()               {}
func (*GetLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetLogsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetLogsRequest) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

type GetLogsResponse struct {
	Type GetLogsResponse_Type `protobuf:"varint,1,opt,name=type,enum=cand.services.pods.v1.GetLogsResponse_Type" json:"type,omitempty"`
	Line []byte               `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
}

func (m *GetLogsResponse) Reset()                    { *m = GetLogsResponse{} }
func (m *GetLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLogsResponse) ProtoMessage()               {}
func (*GetLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetLogsResponse) GetType() GetLogsResponse_Type {
	if m != nil {
		return m.Type
	}
	return GetLogsResponse_STDOUT
}

func (m *GetLogsResponse) GetLine() []byte {
	if m != nil {
		return m.Line
	}
	return nil
}

type Pod struct {
	Metadata map[string]string `protobuf:"bytes,1,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Spec     *PodSpec          `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *PodStatus        `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *Pod) Reset()                    { *m = Pod{} }
func (m *Pod) String() string            { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()               {}
func (*Pod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Pod) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Pod) GetSpec() *PodSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Pod) GetStatus() *PodStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type PodSpec struct {
	Containers []*Container `protobuf:"bytes,1,rep,name=containers" json:"containers,omitempty"`
}

func (m *PodSpec) Reset()                    { *m = PodSpec{} }
func (m *PodSpec) String() string            { return proto.CompactTextString(m) }
func (*PodSpec) ProtoMessage()               {}
func (*PodSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PodSpec) GetAllContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

type Container struct {
	ID    string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Image string `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Container) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type PodStatus struct {
	ContainerStatuses []*ContainerStatus `protobuf:"bytes,1,rep,name=containerStatuses" json:"containerStatuses,omitempty"`
}

func (m *PodStatus) Reset()                    { *m = PodStatus{} }
func (m *PodStatus) String() string            { return proto.CompactTextString(m) }
func (*PodStatus) ProtoMessage()               {}
func (*PodStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PodStatus) GetContainerStatuses() []*ContainerStatus {
	if m != nil {
		return m.ContainerStatuses
	}
	return nil
}

type ContainerStatus struct {
	ContainerID string `protobuf:"bytes,1,opt,name=containerID" json:"containerID,omitempty"`
	Image       string `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	State       string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
}

func (m *ContainerStatus) Reset()                    { *m = ContainerStatus{} }
func (m *ContainerStatus) String() string            { return proto.CompactTextString(m) }
func (*ContainerStatus) ProtoMessage()               {}
func (*ContainerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ContainerStatus) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *ContainerStatus) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ContainerStatus) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*CreatePodRequest)(nil), "cand.services.pods.v1.CreatePodRequest")
	proto.RegisterType((*CreatePodResponse)(nil), "cand.services.pods.v1.CreatePodResponse")
	proto.RegisterType((*ListPodsRequest)(nil), "cand.services.pods.v1.ListPodsRequest")
	proto.RegisterType((*ListPodsResponse)(nil), "cand.services.pods.v1.ListPodsResponse")
	proto.RegisterType((*GetLogsRequest)(nil), "cand.services.pods.v1.GetLogsRequest")
	proto.RegisterType((*GetLogsResponse)(nil), "cand.services.pods.v1.GetLogsResponse")
	proto.RegisterType((*Pod)(nil), "cand.services.pods.v1.Pod")
	proto.RegisterType((*PodSpec)(nil), "cand.services.pods.v1.PodSpec")
	proto.RegisterType((*Container)(nil), "cand.services.pods.v1.Container")
	proto.RegisterType((*PodStatus)(nil), "cand.services.pods.v1.PodStatus")
	proto.RegisterType((*ContainerStatus)(nil), "cand.services.pods.v1.ContainerStatus")
	proto.RegisterEnum("cand.services.pods.v1.GetLogsResponse_Type", GetLogsResponse_Type_name, GetLogsResponse_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pods service

type PodsClient interface {
	Create(ctx context.Context, in *CreatePodRequest, opts ...grpc.CallOption) (*CreatePodResponse, error)
	List(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error)
	Logs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (Pods_LogsClient, error)
}

type podsClient struct {
	cc *grpc.ClientConn
}

func NewPodsClient(cc *grpc.ClientConn) PodsClient {
	return &podsClient{cc}
}

func (c *podsClient) Create(ctx context.Context, in *CreatePodRequest, opts ...grpc.CallOption) (*CreatePodResponse, error) {
	out := new(CreatePodResponse)
	err := grpc.Invoke(ctx, "/cand.services.pods.v1.Pods/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podsClient) List(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error) {
	out := new(ListPodsResponse)
	err := grpc.Invoke(ctx, "/cand.services.pods.v1.Pods/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podsClient) Logs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (Pods_LogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Pods_serviceDesc.Streams[0], c.cc, "/cand.services.pods.v1.Pods/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &podsLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Pods_LogsClient interface {
	Recv() (*GetLogsResponse, error)
	grpc.ClientStream
}

type podsLogsClient struct {
	grpc.ClientStream
}

func (x *podsLogsClient) Recv() (*GetLogsResponse, error) {
	m := new(GetLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Pods service

type PodsServer interface {
	Create(context.Context, *CreatePodRequest) (*CreatePodResponse, error)
	List(context.Context, *ListPodsRequest) (*ListPodsResponse, error)
	Logs(*GetLogsRequest, Pods_LogsServer) error
}

func RegisterPodsServer(s *grpc.Server, srv PodsServer) {
	s.RegisterService(&_Pods_serviceDesc, srv)
}

func _Pods_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cand.services.pods.v1.Pods/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodsServer).Create(ctx, req.(*CreatePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pods_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cand.services.pods.v1.Pods/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodsServer).List(ctx, req.(*ListPodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pods_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PodsServer).Logs(m, &podsLogsServer{stream})
}

type Pods_LogsServer interface {
	Send(*GetLogsResponse) error
	grpc.ServerStream
}

type podsLogsServer struct {
	grpc.ServerStream
}

func (x *podsLogsServer) Send(m *GetLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Pods_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cand.services.pods.v1.Pods",
	HandlerType: (*PodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Pods_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Pods_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _Pods_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pods.proto",
}

func init() { proto.RegisterFile("pods.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x25, 0x69, 0x28, 0xf4, 0x0e, 0xda, 0xce, 0x02, 0xa9, 0xaa, 0x60, 0x54, 0x96, 0xd8, 0x2a,
	0x81, 0x52, 0x56, 0x5e, 0x26, 0xf6, 0xc0, 0xd8, 0x5a, 0xa1, 0x8a, 0x21, 0x2a, 0xaf, 0x08, 0x09,
	0x1e, 0x90, 0x49, 0xae, 0xaa, 0x68, 0x6b, 0x6c, 0x6a, 0xb7, 0x52, 0xff, 0x00, 0xff, 0x80, 0xbf,
	0xca, 0x33, 0xb2, 0xe3, 0xa6, 0xa5, 0x23, 0x74, 0x3c, 0xc5, 0xb1, 0xcf, 0x39, 0xf7, 0xf8, 0x7e,
	0x18, 0x40, 0x8a, 0x58, 0x85, 0x72, 0x2a, 0xb4, 0x20, 0x0f, 0x23, 0x9e, 0xc6, 0xa1, 0xc2, 0xe9,
	0x3c, 0x89, 0x50, 0x85, 0xf6, 0x64, 0x7e, 0x48, 0x4f, 0xa0, 0x7e, 0x36, 0x45, 0xae, 0x71, 0x28,
	0x62, 0x86, 0xdf, 0x67, 0xa8, 0x34, 0x79, 0x0e, 0x25, 0x29, 0xe2, 0x86, 0xd7, 0xf2, 0xda, 0x3b,
	0xdd, 0x66, 0xf8, 0x57, 0x62, 0x68, 0xf0, 0x06, 0x46, 0xdf, 0xc0, 0xee, 0x9a, 0x82, 0x92, 0x22,
	0x55, 0xf8, 0x9f, 0x12, 0x1d, 0xa8, 0x9d, 0x27, 0x4a, 0x0f, 0x45, 0xac, 0x96, 0x1e, 0x1e, 0x41,
	0x25, 0xe5, 0x13, 0x54, 0x92, 0x47, 0x68, 0x65, 0x2a, 0x6c, 0xb5, 0x41, 0x4f, 0xa1, 0xbe, 0x22,
	0xb8, 0x90, 0x21, 0x04, 0x46, 0xb8, 0xe1, 0xb5, 0x4a, 0x5b, 0x62, 0x5a, 0x1c, 0x1d, 0x42, 0xf5,
	0x2d, 0xea, 0x73, 0x31, 0xbe, 0x59, 0x4c, 0xd2, 0x82, 0x9d, 0x48, 0xa4, 0x9a, 0x27, 0x29, 0x4e,
	0x07, 0xbd, 0x86, 0x6f, 0xcf, 0xd7, 0xb7, 0xe8, 0x0f, 0x0f, 0x6a, 0xb9, 0xa4, 0x73, 0xf5, 0x1a,
	0x02, 0xbd, 0x90, 0x99, 0x5c, 0xb5, 0xfb, 0xac, 0xc0, 0xd5, 0x06, 0x2b, 0x1c, 0x2d, 0x24, 0x32,
	0x4b, 0x24, 0x04, 0x82, 0xab, 0x24, 0x45, 0x1b, 0xef, 0x1e, 0xb3, 0x6b, 0xba, 0x07, 0x81, 0x41,
	0x10, 0x80, 0xf2, 0xc5, 0xa8, 0xf7, 0xe1, 0xe3, 0xa8, 0x7e, 0xcb, 0xad, 0xfb, 0x8c, 0xd5, 0x3d,
	0xfa, 0xcb, 0x83, 0xd2, 0x50, 0xc4, 0xa4, 0x07, 0x77, 0x27, 0xa8, 0x79, 0xcc, 0x35, 0x77, 0x69,
	0x69, 0x17, 0xa7, 0x25, 0x7c, 0xef, 0xa0, 0xfd, 0x54, 0x4f, 0x17, 0x2c, 0x67, 0x92, 0x2e, 0x04,
	0x4a, 0x62, 0x64, 0x1d, 0xec, 0x74, 0xf7, 0x8a, 0x15, 0x2e, 0x24, 0x46, 0xcc, 0x62, 0xc9, 0x11,
	0x94, 0x95, 0xe6, 0x7a, 0xa6, 0x1a, 0x25, 0xcb, 0x6a, 0xfd, 0x83, 0x65, 0x71, 0xcc, 0xe1, 0x9b,
	0xc7, 0x70, 0xff, 0x0f, 0x23, 0xa4, 0x0e, 0xa5, 0x4b, 0x5c, 0xb8, 0x7a, 0x98, 0x25, 0x79, 0x00,
	0xb7, 0xe7, 0xfc, 0x6a, 0x86, 0xae, 0x06, 0xd9, 0xcf, 0x2b, 0xff, 0xc8, 0xa3, 0xef, 0xe0, 0x8e,
	0xf3, 0x41, 0x4e, 0x00, 0xf2, 0xda, 0x2c, 0x9b, 0xa2, 0xc8, 0xc5, 0xd9, 0x12, 0xc8, 0xd6, 0x38,
	0xb4, 0x0f, 0x95, 0xfc, 0x80, 0x54, 0xc1, 0x1f, 0xf4, 0x9c, 0x09, 0x7f, 0xd0, 0x33, 0x65, 0x31,
	0xad, 0xe1, 0x2c, 0xd8, 0xb5, 0xf1, 0x95, 0x4c, 0xf8, 0x18, 0xed, 0x9d, 0x2b, 0x2c, 0xfb, 0xa1,
	0x1c, 0x2a, 0xf9, 0x2d, 0xc9, 0x08, 0x76, 0xf3, 0x08, 0xd9, 0x16, 0x2e, 0xcd, 0xed, 0x6f, 0x33,
	0xe7, 0x12, 0x75, 0x5d, 0x80, 0x7e, 0x85, 0xda, 0x06, 0x6a, 0xb3, 0x5b, 0xbd, 0x6b, 0xdd, 0xba,
	0x72, 0xeb, 0xaf, 0xb9, 0x35, 0xbb, 0xa6, 0x10, 0xf9, 0x1d, 0xec, 0x4f, 0xf7, 0xa7, 0x0f, 0x81,
	0x19, 0x36, 0xf2, 0x05, 0xca, 0xd9, 0xb0, 0x93, 0x83, 0x22, 0xbb, 0x1b, 0xaf, 0x49, 0xb3, 0xbd,
	0x1d, 0xe8, 0x66, 0xe5, 0x13, 0x04, 0x66, 0xaa, 0x49, 0x51, 0x26, 0x36, 0xde, 0x88, 0xe6, 0xc1,
	0x56, 0xdc, 0x9a, 0xb0, 0x18, 0x2b, 0xf2, 0x74, 0xdb, 0xf8, 0x65, 0xba, 0xfb, 0x37, 0x9b, 0xd2,
	0x17, 0xde, 0xe9, 0x93, 0xcf, 0x8f, 0xe5, 0xe5, 0xb8, 0xc3, 0x65, 0xd2, 0x59, 0xa2, 0x3b, 0x06,
	0xdd, 0x99, 0x1f, 0x1e, 0x9b, 0xef, 0xb7, 0xb2, 0x7d, 0x7c, 0x5f, 0xfe, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x67, 0xf4, 0x35, 0xee, 0x8a, 0x05, 0x00, 0x00,
}