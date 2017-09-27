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
	StdOutputStreamResponse
	Pod
	PodSpec
	Container
	PodStatus
	ContainerStatus
	ResourceMetadata
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

type StdOutputStreamResponse struct {
	Line []byte `protobuf:"bytes,1,opt,name=line,proto3" json:"line,omitempty"`
	// Is this stderr(=true) or stdout(=false)
	Stderr bool `protobuf:"varint,2,opt,name=stderr" json:"stderr,omitempty"`
}

func (m *StdOutputStreamResponse) Reset()                    { *m = StdOutputStreamResponse{} }
func (m *StdOutputStreamResponse) String() string            { return proto.CompactTextString(m) }
func (*StdOutputStreamResponse) ProtoMessage()               {}
func (*StdOutputStreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StdOutputStreamResponse) GetLine() []byte {
	if m != nil {
		return m.Line
	}
	return nil
}

func (m *StdOutputStreamResponse) GetStderr() bool {
	if m != nil {
		return m.Stderr
	}
	return false
}

type Pod struct {
	Metadata *ResourceMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *PodSpec          `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *PodStatus        `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *Pod) Reset()                    { *m = Pod{} }
func (m *Pod) String() string            { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()               {}
func (*Pod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Pod) GetMetadata() *ResourceMetadata {
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

func (m *PodSpec) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

type Container struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Image string `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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

// ResourceMetadata is metadata that all resources must have
type ResourceMetadata struct {
	// Name must be unique within a namespace.
	// Cannot be updated.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Namespace defines space within each name must be unique.
	// An empty namespace is equivalent to the default namespace.
	// Cannot be updated.
	Namespace string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ResourceMetadata) Reset()                    { *m = ResourceMetadata{} }
func (m *ResourceMetadata) String() string            { return proto.CompactTextString(m) }
func (*ResourceMetadata) ProtoMessage()               {}
func (*ResourceMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ResourceMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceMetadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*CreatePodRequest)(nil), "cand.services.pods.v1.CreatePodRequest")
	proto.RegisterType((*CreatePodResponse)(nil), "cand.services.pods.v1.CreatePodResponse")
	proto.RegisterType((*ListPodsRequest)(nil), "cand.services.pods.v1.ListPodsRequest")
	proto.RegisterType((*ListPodsResponse)(nil), "cand.services.pods.v1.ListPodsResponse")
	proto.RegisterType((*GetLogsRequest)(nil), "cand.services.pods.v1.GetLogsRequest")
	proto.RegisterType((*StdOutputStreamResponse)(nil), "cand.services.pods.v1.StdOutputStreamResponse")
	proto.RegisterType((*Pod)(nil), "cand.services.pods.v1.Pod")
	proto.RegisterType((*PodSpec)(nil), "cand.services.pods.v1.PodSpec")
	proto.RegisterType((*Container)(nil), "cand.services.pods.v1.Container")
	proto.RegisterType((*PodStatus)(nil), "cand.services.pods.v1.PodStatus")
	proto.RegisterType((*ContainerStatus)(nil), "cand.services.pods.v1.ContainerStatus")
	proto.RegisterType((*ResourceMetadata)(nil), "cand.services.pods.v1.ResourceMetadata")
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
	Recv() (*StdOutputStreamResponse, error)
	grpc.ClientStream
}

type podsLogsClient struct {
	grpc.ClientStream
}

func (x *podsLogsClient) Recv() (*StdOutputStreamResponse, error) {
	m := new(StdOutputStreamResponse)
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
	Send(*StdOutputStreamResponse) error
	grpc.ServerStream
}

type podsLogsServer struct {
	grpc.ServerStream
}

func (x *podsLogsServer) Send(m *StdOutputStreamResponse) error {
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
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x6d, 0x8b, 0xd3, 0x40,
	0x10, 0x26, 0x6d, 0xac, 0x97, 0xa9, 0x78, 0xbd, 0xc5, 0x97, 0x52, 0x7c, 0x29, 0x0b, 0x7a, 0xf7,
	0x41, 0x12, 0xaf, 0x22, 0x08, 0x7e, 0x39, 0xaf, 0x27, 0x22, 0x9e, 0x58, 0xb6, 0x82, 0xa0, 0x1f,
	0x8e, 0x35, 0x19, 0x4a, 0xd0, 0x66, 0xd7, 0xec, 0xa6, 0xff, 0xc6, 0xff, 0xe1, 0xcf, 0x93, 0xdd,
	0x6c, 0x72, 0xbd, 0xb4, 0xb1, 0xde, 0xa7, 0x64, 0x67, 0x9f, 0xe7, 0x99, 0xc9, 0x33, 0x33, 0x01,
	0x90, 0x22, 0x51, 0xa1, 0xcc, 0x85, 0x16, 0xe4, 0x6e, 0xcc, 0xb3, 0x24, 0x54, 0x98, 0xaf, 0xd2,
	0x18, 0x55, 0x68, 0x6f, 0x56, 0xc7, 0xf4, 0x04, 0x06, 0xd3, 0x1c, 0xb9, 0xc6, 0x99, 0x48, 0x18,
	0xfe, 0x2a, 0x50, 0x69, 0xf2, 0x0c, 0xba, 0x52, 0x24, 0x43, 0x6f, 0xec, 0x1d, 0xf5, 0x27, 0xa3,
	0x70, 0x2b, 0x31, 0x34, 0x78, 0x03, 0xa3, 0x6f, 0xe0, 0x60, 0x4d, 0x41, 0x49, 0x91, 0x29, 0xbc,
	0xa6, 0x44, 0x04, 0xfb, 0xe7, 0xa9, 0xd2, 0x33, 0x91, 0xa8, 0xaa, 0x86, 0x07, 0x10, 0x64, 0x7c,
	0x89, 0x4a, 0xf2, 0x18, 0xad, 0x4c, 0xc0, 0x2e, 0x03, 0xf4, 0x14, 0x06, 0x97, 0x04, 0x97, 0x32,
	0x04, 0xdf, 0x08, 0x0f, 0xbd, 0x71, 0x77, 0x47, 0x4e, 0x8b, 0xa3, 0x33, 0xb8, 0xfd, 0x0e, 0xf5,
	0xb9, 0x58, 0xfc, 0x5f, 0x4e, 0x32, 0x86, 0x7e, 0x2c, 0x32, 0xcd, 0xd3, 0x0c, 0xf3, 0xf7, 0x67,
	0xc3, 0x8e, 0xbd, 0x5f, 0x0f, 0xd1, 0xb7, 0x70, 0x7f, 0xae, 0x93, 0x4f, 0x85, 0x96, 0x85, 0x9e,
	0xeb, 0x1c, 0xf9, 0xb2, 0x2e, 0x8e, 0x80, 0xff, 0x33, 0xcd, 0x4a, 0xd5, 0x5b, 0xcc, 0xbe, 0x93,
	0x7b, 0xd0, 0x53, 0x3a, 0xc1, 0x3c, 0xb7, 0x5a, 0x7b, 0xcc, 0x9d, 0xe8, 0x1f, 0x0f, 0xba, 0x33,
	0x91, 0x90, 0x29, 0xec, 0x2d, 0x51, 0xf3, 0x84, 0x6b, 0xee, 0x8c, 0x3c, 0x6c, 0xf9, 0x28, 0x86,
	0x4a, 0x14, 0x79, 0x8c, 0x1f, 0x1d, 0x9c, 0xd5, 0x44, 0x32, 0x01, 0x5f, 0x49, 0x8c, 0x6d, 0x8a,
	0xfe, 0xe4, 0x51, 0xbb, 0x2b, 0x73, 0x89, 0x31, 0xb3, 0x58, 0xf2, 0xca, 0x14, 0xc6, 0x75, 0xa1,
	0x86, 0x5d, 0xcb, 0x1a, 0xff, 0x83, 0x65, 0x71, 0xcc, 0xe1, 0xe9, 0x07, 0xb8, 0xe9, 0xa4, 0xc8,
	0x09, 0x40, 0xed, 0x4d, 0xd5, 0x94, 0x36, 0xa1, 0x69, 0x05, 0x64, 0x6b, 0x1c, 0xfa, 0x12, 0x82,
	0xfa, 0xc2, 0x18, 0x68, 0x5a, 0xe1, 0xda, 0x62, 0xdf, 0xc9, 0x1d, 0xb8, 0x91, 0x2e, 0xf9, 0x02,
	0x5d, 0x2f, 0xca, 0x03, 0xe5, 0x10, 0xd4, 0x85, 0x91, 0xcf, 0x70, 0x50, 0x2b, 0x96, 0x21, 0xac,
	0x8a, 0x79, 0xba, 0xab, 0x18, 0xf7, 0x6d, 0x9b, 0x02, 0xf4, 0x02, 0xf6, 0x1b, 0xa8, 0xe6, 0x74,
	0x78, 0x1b, 0xd3, 0xb1, 0xbd, 0x5a, 0x13, 0x35, 0xde, 0xa1, 0xb5, 0x3a, 0x60, 0xe5, 0x81, 0x9e,
	0xc1, 0xa0, 0xd9, 0xd3, 0xad, 0x0e, 0x5c, 0x99, 0xd8, 0x4e, 0x63, 0x62, 0x27, 0xbf, 0x3b, 0xe0,
	0x9b, 0x15, 0x21, 0xdf, 0xa0, 0x57, 0xae, 0x28, 0x69, 0x9b, 0xa0, 0xe6, 0x3f, 0x60, 0x74, 0xb4,
	0x1b, 0xe8, 0x46, 0xfb, 0x0b, 0xf8, 0x66, 0x17, 0x49, 0x9b, 0x9f, 0x8d, 0xcd, 0x1e, 0x1d, 0xee,
	0xc4, 0x39, 0xe1, 0x0b, 0xf0, 0xcd, 0x76, 0x92, 0x27, 0x2d, 0x84, 0xab, 0xdb, 0x3b, 0x0a, 0x5b,
	0x60, 0x2d, 0x2b, 0xf9, 0xdc, 0x3b, 0x7d, 0xfc, 0xf5, 0xa1, 0xfc, 0xb1, 0x88, 0xb8, 0x4c, 0xa3,
	0x8a, 0x15, 0x19, 0x56, 0xb4, 0x3a, 0x7e, 0x6d, 0x9e, 0xdf, 0x7b, 0xf6, 0xd7, 0xf9, 0xe2, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x71, 0xbf, 0x42, 0x48, 0x05, 0x00, 0x00,
}
