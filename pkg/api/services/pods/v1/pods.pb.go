// Code generated by protoc-gen-go.
// source: services/pods/v1/pods.proto
// DO NOT EDIT!

/*
Package pods is a generated protocol buffer package.

It is generated from these files:
	services/pods/v1/pods.proto

It has these top-level messages:
	CreatePodRequest
	CreatePodStreamResponse
	ImageFetch
	ImageLayerStatus
	StartPodRequest
	StartPodResponse
	DeletePodRequest
	DeletePodResponse
	ListPodsRequest
	ListPodsResponse
	Pod
	PodSpec
	PodStatus
*/
package pods

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import cand_core "github.com/ernoaapa/can/pkg/api/core"
import cand_services_containers_v1 "github.com/ernoaapa/can/pkg/api/services/containers/v1"

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
	Tty bool `protobuf:"varint,2,opt,name=tty" json:"tty,omitempty"`
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

func (m *CreatePodRequest) GetTty() bool {
	if m != nil {
		return m.Tty
	}
	return false
}

type CreatePodStreamResponse struct {
	Images []*ImageFetch `protobuf:"bytes,1,rep,name=images" json:"images,omitempty"`
}

func (m *CreatePodStreamResponse) Reset()                    { *m = CreatePodStreamResponse{} }
func (m *CreatePodStreamResponse) String() string            { return proto.CompactTextString(m) }
func (*CreatePodStreamResponse) ProtoMessage()               {}
func (*CreatePodStreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreatePodStreamResponse) GetImages() []*ImageFetch {
	if m != nil {
		return m.Images
	}
	return nil
}

type ImageFetch struct {
	ContainerID string              `protobuf:"bytes,1,opt,name=containerID" json:"containerID,omitempty"`
	Image       string              `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	Resolved    bool                `protobuf:"varint,3,opt,name=resolved" json:"resolved,omitempty"`
	Layers      []*ImageLayerStatus `protobuf:"bytes,4,rep,name=layers" json:"layers,omitempty"`
}

func (m *ImageFetch) Reset()                    { *m = ImageFetch{} }
func (m *ImageFetch) String() string            { return proto.CompactTextString(m) }
func (*ImageFetch) ProtoMessage()               {}
func (*ImageFetch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ImageFetch) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *ImageFetch) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ImageFetch) GetResolved() bool {
	if m != nil {
		return m.Resolved
	}
	return false
}

func (m *ImageFetch) GetLayers() []*ImageLayerStatus {
	if m != nil {
		return m.Layers
	}
	return nil
}

type ImageLayerStatus struct {
	Ref    string `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	Digest string `protobuf:"bytes,2,opt,name=digest" json:"digest,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Offset int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	Total  int64  `protobuf:"varint,5,opt,name=total" json:"total,omitempty"`
}

func (m *ImageLayerStatus) Reset()                    { *m = ImageLayerStatus{} }
func (m *ImageLayerStatus) String() string            { return proto.CompactTextString(m) }
func (*ImageLayerStatus) ProtoMessage()               {}
func (*ImageLayerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ImageLayerStatus) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *ImageLayerStatus) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *ImageLayerStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ImageLayerStatus) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ImageLayerStatus) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type StartPodRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *StartPodRequest) Reset()                    { *m = StartPodRequest{} }
func (m *StartPodRequest) String() string            { return proto.CompactTextString(m) }
func (*StartPodRequest) ProtoMessage()               {}
func (*StartPodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StartPodRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *StartPodRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StartPodResponse struct {
	Pod *Pod `protobuf:"bytes,1,opt,name=pod" json:"pod,omitempty"`
}

func (m *StartPodResponse) Reset()                    { *m = StartPodResponse{} }
func (m *StartPodResponse) String() string            { return proto.CompactTextString(m) }
func (*StartPodResponse) ProtoMessage()               {}
func (*StartPodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StartPodResponse) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

type DeletePodRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *DeletePodRequest) Reset()                    { *m = DeletePodRequest{} }
func (m *DeletePodRequest) String() string            { return proto.CompactTextString(m) }
func (*DeletePodRequest) ProtoMessage()               {}
func (*DeletePodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeletePodRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeletePodRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeletePodResponse struct {
	Pod *Pod `protobuf:"bytes,1,opt,name=pod" json:"pod,omitempty"`
}

func (m *DeletePodResponse) Reset()                    { *m = DeletePodResponse{} }
func (m *DeletePodResponse) String() string            { return proto.CompactTextString(m) }
func (*DeletePodResponse) ProtoMessage()               {}
func (*DeletePodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeletePodResponse) GetPod() *Pod {
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
func (*ListPodsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
func (*ListPodsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListPodsResponse) GetPods() []*Pod {
	if m != nil {
		return m.Pods
	}
	return nil
}

type Pod struct {
	Metadata *cand_core.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *PodSpec                    `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *PodStatus                  `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *Pod) Reset()                    { *m = Pod{} }
func (m *Pod) String() string            { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()               {}
func (*Pod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Pod) GetMetadata() *cand_core.ResourceMetadata {
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
	Containers  []*cand_services_containers_v1.Container `protobuf:"bytes,1,rep,name=containers" json:"containers,omitempty"`
	HostNetwork bool                                     `protobuf:"varint,2,opt,name=hostNetwork" json:"hostNetwork,omitempty"`
	HostPID     bool                                     `protobuf:"varint,3,opt,name=hostPID" json:"hostPID,omitempty"`
}

func (m *PodSpec) Reset()                    { *m = PodSpec{} }
func (m *PodSpec) String() string            { return proto.CompactTextString(m) }
func (*PodSpec) ProtoMessage()               {}
func (*PodSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PodSpec) GetContainers() []*cand_services_containers_v1.Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *PodSpec) GetHostNetwork() bool {
	if m != nil {
		return m.HostNetwork
	}
	return false
}

func (m *PodSpec) GetHostPID() bool {
	if m != nil {
		return m.HostPID
	}
	return false
}

type PodStatus struct {
	ContainerStatuses []*cand_services_containers_v1.ContainerStatus `protobuf:"bytes,1,rep,name=containerStatuses" json:"containerStatuses,omitempty"`
}

func (m *PodStatus) Reset()                    { *m = PodStatus{} }
func (m *PodStatus) String() string            { return proto.CompactTextString(m) }
func (*PodStatus) ProtoMessage()               {}
func (*PodStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PodStatus) GetContainerStatuses() []*cand_services_containers_v1.ContainerStatus {
	if m != nil {
		return m.ContainerStatuses
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePodRequest)(nil), "cand.services.pods.v1.CreatePodRequest")
	proto.RegisterType((*CreatePodStreamResponse)(nil), "cand.services.pods.v1.CreatePodStreamResponse")
	proto.RegisterType((*ImageFetch)(nil), "cand.services.pods.v1.ImageFetch")
	proto.RegisterType((*ImageLayerStatus)(nil), "cand.services.pods.v1.ImageLayerStatus")
	proto.RegisterType((*StartPodRequest)(nil), "cand.services.pods.v1.StartPodRequest")
	proto.RegisterType((*StartPodResponse)(nil), "cand.services.pods.v1.StartPodResponse")
	proto.RegisterType((*DeletePodRequest)(nil), "cand.services.pods.v1.DeletePodRequest")
	proto.RegisterType((*DeletePodResponse)(nil), "cand.services.pods.v1.DeletePodResponse")
	proto.RegisterType((*ListPodsRequest)(nil), "cand.services.pods.v1.ListPodsRequest")
	proto.RegisterType((*ListPodsResponse)(nil), "cand.services.pods.v1.ListPodsResponse")
	proto.RegisterType((*Pod)(nil), "cand.services.pods.v1.Pod")
	proto.RegisterType((*PodSpec)(nil), "cand.services.pods.v1.PodSpec")
	proto.RegisterType((*PodStatus)(nil), "cand.services.pods.v1.PodStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pods service

type PodsClient interface {
	Create(ctx context.Context, in *CreatePodRequest, opts ...grpc.CallOption) (Pods_CreateClient, error)
	Start(ctx context.Context, in *StartPodRequest, opts ...grpc.CallOption) (*StartPodResponse, error)
	Delete(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error)
	List(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error)
}

type podsClient struct {
	cc *grpc.ClientConn
}

func NewPodsClient(cc *grpc.ClientConn) PodsClient {
	return &podsClient{cc}
}

func (c *podsClient) Create(ctx context.Context, in *CreatePodRequest, opts ...grpc.CallOption) (Pods_CreateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Pods_serviceDesc.Streams[0], c.cc, "/cand.services.pods.v1.Pods/Create", opts...)
	if err != nil {
		return nil, err
	}
	x := &podsCreateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Pods_CreateClient interface {
	Recv() (*CreatePodStreamResponse, error)
	grpc.ClientStream
}

type podsCreateClient struct {
	grpc.ClientStream
}

func (x *podsCreateClient) Recv() (*CreatePodStreamResponse, error) {
	m := new(CreatePodStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *podsClient) Start(ctx context.Context, in *StartPodRequest, opts ...grpc.CallOption) (*StartPodResponse, error) {
	out := new(StartPodResponse)
	err := grpc.Invoke(ctx, "/cand.services.pods.v1.Pods/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podsClient) Delete(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error) {
	out := new(DeletePodResponse)
	err := grpc.Invoke(ctx, "/cand.services.pods.v1.Pods/Delete", in, out, c.cc, opts...)
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

// Server API for Pods service

type PodsServer interface {
	Create(*CreatePodRequest, Pods_CreateServer) error
	Start(context.Context, *StartPodRequest) (*StartPodResponse, error)
	Delete(context.Context, *DeletePodRequest) (*DeletePodResponse, error)
	List(context.Context, *ListPodsRequest) (*ListPodsResponse, error)
}

func RegisterPodsServer(s *grpc.Server, srv PodsServer) {
	s.RegisterService(&_Pods_serviceDesc, srv)
}

func _Pods_Create_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreatePodRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PodsServer).Create(m, &podsCreateServer{stream})
}

type Pods_CreateServer interface {
	Send(*CreatePodStreamResponse) error
	grpc.ServerStream
}

type podsCreateServer struct {
	grpc.ServerStream
}

func (x *podsCreateServer) Send(m *CreatePodStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Pods_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodsServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cand.services.pods.v1.Pods/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodsServer).Start(ctx, req.(*StartPodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pods_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cand.services.pods.v1.Pods/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodsServer).Delete(ctx, req.(*DeletePodRequest))
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

var _Pods_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cand.services.pods.v1.Pods",
	HandlerType: (*PodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Pods_Start_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Pods_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Pods_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Create",
			Handler:       _Pods_Create_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/pods/v1/pods.proto",
}

func init() { proto.RegisterFile("services/pods/v1/pods.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0x34, 0x6d, 0x26, 0x87, 0xa6, 0xcb, 0x97, 0xe5, 0x22, 0x14, 0x7c, 0x68, 0x73,
	0xa8, 0x6c, 0x1a, 0x84, 0x0a, 0xe2, 0x00, 0xb4, 0x51, 0xa5, 0x4a, 0x05, 0x55, 0x1b, 0x24, 0x50,
	0x39, 0x6d, 0xed, 0x69, 0x6a, 0x35, 0xf1, 0x9a, 0xdd, 0x4d, 0x50, 0xaf, 0xdc, 0xf9, 0x11, 0xdc,
	0xf9, 0x87, 0x5c, 0xd0, 0xae, 0x37, 0x4e, 0x6a, 0x70, 0x3f, 0x38, 0x79, 0x67, 0xfc, 0xe6, 0xed,
	0x5b, 0xef, 0xbc, 0x31, 0x6c, 0x48, 0x14, 0xd3, 0x24, 0x42, 0x19, 0x66, 0x3c, 0x96, 0xe1, 0x74,
	0xc7, 0x3c, 0x83, 0x4c, 0x70, 0xc5, 0xc9, 0x83, 0x88, 0xa5, 0x71, 0x30, 0x43, 0x04, 0xe6, 0xcd,
	0x74, 0xc7, 0xbb, 0x17, 0x71, 0x81, 0xe1, 0x18, 0x15, 0x8b, 0x99, 0x62, 0x39, 0xd6, 0xdb, 0x2a,
	0x88, 0x22, 0x9e, 0x2a, 0x96, 0xa4, 0x28, 0x0c, 0xdd, 0x3c, 0xca, 0x81, 0x3e, 0x85, 0xf6, 0xbe,
	0x40, 0xa6, 0xf0, 0x98, 0xc7, 0x14, 0xbf, 0x4e, 0x50, 0x2a, 0xb2, 0x0d, 0xb5, 0x8c, 0xc7, 0xae,
	0xd3, 0x71, 0xba, 0xad, 0x9e, 0x17, 0xfc, 0x73, 0xdb, 0x40, 0xe3, 0x35, 0x8c, 0xb4, 0xa1, 0xa6,
	0xd4, 0xa5, 0xbb, 0xd4, 0x71, 0xba, 0xab, 0x54, 0x2f, 0xfd, 0x8f, 0xf0, 0xa8, 0xe0, 0x1c, 0x28,
	0x81, 0x6c, 0x4c, 0x51, 0x66, 0x3c, 0x95, 0x48, 0x5e, 0x41, 0x23, 0x19, 0xb3, 0x21, 0x4a, 0xd7,
	0xe9, 0xd4, 0xba, 0xad, 0xde, 0xd3, 0x0a, 0xf6, 0x43, 0x0d, 0x3a, 0x40, 0x15, 0x9d, 0x53, 0x5b,
	0xe0, 0xff, 0x74, 0x00, 0xe6, 0x69, 0xd2, 0x81, 0x56, 0x71, 0x98, 0xc3, 0xbe, 0x11, 0xdb, 0xa4,
	0x8b, 0x29, 0x72, 0x1f, 0x96, 0x4d, 0xa9, 0x91, 0xd6, 0xa4, 0x79, 0x40, 0x3c, 0x58, 0x15, 0x28,
	0xf9, 0x68, 0x8a, 0xb1, 0x5b, 0x33, 0x9a, 0x8b, 0x98, 0xbc, 0x81, 0xc6, 0x88, 0x5d, 0xa2, 0x90,
	0x6e, 0xdd, 0xa8, 0xdb, 0xba, 0x4e, 0xdd, 0x91, 0x46, 0x0e, 0x14, 0x53, 0x13, 0x49, 0x6d, 0x99,
	0xff, 0xdd, 0x81, 0x76, 0xf9, 0xa5, 0xfe, 0x40, 0x02, 0xcf, 0xac, 0x42, 0xbd, 0x24, 0x0f, 0xa1,
	0x11, 0x27, 0x43, 0x94, 0xca, 0x4a, 0xb3, 0x91, 0xce, 0x4b, 0x53, 0x63, 0x94, 0x35, 0xa9, 0x8d,
	0x74, 0x9e, 0x9f, 0x9d, 0x49, 0x54, 0x6e, 0xbd, 0xe3, 0x74, 0x6b, 0xd4, 0x46, 0xfa, 0x84, 0x8a,
	0x2b, 0x36, 0x72, 0x97, 0x4d, 0x3a, 0x0f, 0xfc, 0x7d, 0x58, 0x1b, 0x28, 0x26, 0xd4, 0xc2, 0x8d,
	0x3e, 0x86, 0x66, 0xca, 0xc6, 0x28, 0x33, 0x16, 0xa1, 0x15, 0x32, 0x4f, 0x10, 0x02, 0x75, 0x1d,
	0x58, 0x31, 0x66, 0xed, 0xbf, 0x85, 0xf6, 0x9c, 0xc4, 0x5e, 0xde, 0x9d, 0xfa, 0xc2, 0xef, 0x43,
	0xbb, 0x8f, 0x23, 0xbc, 0xd2, 0x59, 0x77, 0xd7, 0xf1, 0x0e, 0xd6, 0x17, 0x58, 0xfe, 0x4b, 0x48,
	0x08, 0x6b, 0x47, 0x89, 0xd4, 0x27, 0x91, 0xb7, 0xd2, 0xe1, 0xef, 0x41, 0x7b, 0x5e, 0x60, 0xb7,
	0x0c, 0xa0, 0xae, 0x89, 0x6d, 0xdb, 0x5e, 0xb7, 0xa7, 0xc1, 0xf9, 0xbf, 0x1c, 0xa8, 0x1d, 0xf3,
	0x98, 0xec, 0xc2, 0xea, 0xcc, 0x9a, 0x56, 0xef, 0x46, 0x5e, 0xab, 0x5d, 0x1b, 0x50, 0x94, 0x7c,
	0x22, 0x22, 0x7c, 0x6f, 0x21, 0xb4, 0x00, 0x93, 0x1e, 0xd4, 0x65, 0x86, 0x91, 0xf9, 0x18, 0xad,
	0xde, 0x93, 0xea, 0x0d, 0x07, 0x19, 0x46, 0xd4, 0x60, 0xc9, 0xcb, 0x2b, 0xfd, 0xd3, 0xea, 0x75,
	0xae, 0xa9, 0xb2, 0x8d, 0x9b, 0xe3, 0xfd, 0x1f, 0x0e, 0xac, 0x58, 0x2e, 0x72, 0x00, 0x30, 0x1f,
	0x13, 0xf6, 0xc0, 0x9b, 0x25, 0xa6, 0x85, 0x39, 0x32, 0xdd, 0x09, 0xf6, 0x67, 0x11, 0x5d, 0xa8,
	0xd4, 0x0e, 0x3d, 0xe7, 0x52, 0x7d, 0x40, 0xf5, 0x8d, 0x8b, 0x0b, 0x3b, 0x20, 0x16, 0x53, 0xc4,
	0x85, 0x15, 0x1d, 0x1e, 0x1f, 0xf6, 0xad, 0x15, 0x67, 0xa1, 0x3f, 0x84, 0x66, 0x21, 0x92, 0x9c,
	0xc0, 0x7a, 0x41, 0x9b, 0xa7, 0x8a, 0xf9, 0xb1, 0x7d, 0x3b, 0x5d, 0xf6, 0xb4, 0x7f, 0xd3, 0xf4,
	0x7e, 0x2f, 0x41, 0x5d, 0x5f, 0x34, 0x89, 0xa0, 0x91, 0x0f, 0x2d, 0x52, 0xe5, 0xfa, 0xf2, 0x9c,
	0xf4, 0x82, 0x9b, 0x80, 0x57, 0x87, 0xdf, 0x33, 0x87, 0x7c, 0x86, 0x65, 0xe3, 0x2a, 0xb2, 0x59,
	0x51, 0x5a, 0x32, 0xae, 0xb7, 0x75, 0x23, 0xce, 0xf6, 0xe7, 0x17, 0x68, 0xe4, 0x3e, 0xa9, 0x94,
	0x5f, 0x36, 0xa3, 0xd7, 0xbd, 0x19, 0x68, 0xc9, 0x3f, 0x41, 0x5d, 0x1b, 0xa2, 0x52, 0x75, 0xc9,
	0x5e, 0x95, 0xaa, 0xcb, 0xae, 0xda, 0xdb, 0x3d, 0x79, 0x31, 0x4c, 0xd4, 0xf9, 0xe4, 0x34, 0x88,
	0xf8, 0x38, 0x44, 0x91, 0x72, 0xc6, 0x32, 0x16, 0x46, 0x2c, 0x0d, 0xb3, 0x8b, 0x61, 0xc8, 0xb2,
	0x24, 0x2c, 0xff, 0x11, 0x5f, 0xeb, 0xe7, 0x69, 0xc3, 0xfc, 0xbd, 0x9e, 0xff, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0x0f, 0x95, 0xd3, 0x31, 0x07, 0x00, 0x00,
}
