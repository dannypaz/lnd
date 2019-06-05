// Code generated by protoc-gen-go. DO NOT EDIT.
// source: autopilotrpc/autopilot.proto

/*
Package autopilotrpc is a generated protocol buffer package.

It is generated from these files:
	autopilotrpc/autopilot.proto

It has these top-level messages:
	StatusRequest
	StatusResponse
	ModifyStatusRequest
	ModifyStatusResponse
	QueryScoresRequest
	QueryScoresResponse
	SetScoresRequest
	SetScoresResponse
*/
package autopilotrpc

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

type StatusRequest struct {
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StatusResponse struct {
	// / Indicates whether the autopilot is active or not.
	Active bool `protobuf:"varint,1,opt,name=active" json:"active,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatusResponse) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type ModifyStatusRequest struct {
	// / Whether the autopilot agent should be enabled or not.
	Enable bool `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
}

func (m *ModifyStatusRequest) Reset()                    { *m = ModifyStatusRequest{} }
func (m *ModifyStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyStatusRequest) ProtoMessage()               {}
func (*ModifyStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ModifyStatusRequest) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

type ModifyStatusResponse struct {
}

func (m *ModifyStatusResponse) Reset()                    { *m = ModifyStatusResponse{} }
func (m *ModifyStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*ModifyStatusResponse) ProtoMessage()               {}
func (*ModifyStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type QueryScoresRequest struct {
	Pubkeys []string `protobuf:"bytes,1,rep,name=pubkeys" json:"pubkeys,omitempty"`
	// / If set, we will ignore the local channel state when calculating scores.
	IgnoreLocalState bool `protobuf:"varint,2,opt,name=ignore_local_state,json=no_state" json:"ignore_local_state,omitempty"`
}

func (m *QueryScoresRequest) Reset()                    { *m = QueryScoresRequest{} }
func (m *QueryScoresRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryScoresRequest) ProtoMessage()               {}
func (*QueryScoresRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryScoresRequest) GetPubkeys() []string {
	if m != nil {
		return m.Pubkeys
	}
	return nil
}

func (m *QueryScoresRequest) GetIgnoreLocalState() bool {
	if m != nil {
		return m.IgnoreLocalState
	}
	return false
}

type QueryScoresResponse struct {
	Results []*QueryScoresResponse_HeuristicResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *QueryScoresResponse) Reset()                    { *m = QueryScoresResponse{} }
func (m *QueryScoresResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryScoresResponse) ProtoMessage()               {}
func (*QueryScoresResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *QueryScoresResponse) GetResults() []*QueryScoresResponse_HeuristicResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type QueryScoresResponse_HeuristicResult struct {
	Heuristic string             `protobuf:"bytes,1,opt,name=heuristic" json:"heuristic,omitempty"`
	Scores    map[string]float64 `protobuf:"bytes,2,rep,name=scores" json:"scores,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
}

func (m *QueryScoresResponse_HeuristicResult) Reset()         { *m = QueryScoresResponse_HeuristicResult{} }
func (m *QueryScoresResponse_HeuristicResult) String() string { return proto.CompactTextString(m) }
func (*QueryScoresResponse_HeuristicResult) ProtoMessage()    {}
func (*QueryScoresResponse_HeuristicResult) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

func (m *QueryScoresResponse_HeuristicResult) GetHeuristic() string {
	if m != nil {
		return m.Heuristic
	}
	return ""
}

func (m *QueryScoresResponse_HeuristicResult) GetScores() map[string]float64 {
	if m != nil {
		return m.Scores
	}
	return nil
}

type SetScoresRequest struct {
	// / The name of the heuristic to provide scores to.
	Heuristic string `protobuf:"bytes,1,opt,name=heuristic" json:"heuristic,omitempty"`
	// *
	// A map from hex-encoded public keys to scores. Scores must be in the range
	// [0.0, 1.0].
	Scores map[string]float64 `protobuf:"bytes,2,rep,name=scores" json:"scores,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
}

func (m *SetScoresRequest) Reset()                    { *m = SetScoresRequest{} }
func (m *SetScoresRequest) String() string            { return proto.CompactTextString(m) }
func (*SetScoresRequest) ProtoMessage()               {}
func (*SetScoresRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SetScoresRequest) GetHeuristic() string {
	if m != nil {
		return m.Heuristic
	}
	return ""
}

func (m *SetScoresRequest) GetScores() map[string]float64 {
	if m != nil {
		return m.Scores
	}
	return nil
}

type SetScoresResponse struct {
}

func (m *SetScoresResponse) Reset()                    { *m = SetScoresResponse{} }
func (m *SetScoresResponse) String() string            { return proto.CompactTextString(m) }
func (*SetScoresResponse) ProtoMessage()               {}
func (*SetScoresResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*StatusRequest)(nil), "autopilotrpc.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "autopilotrpc.StatusResponse")
	proto.RegisterType((*ModifyStatusRequest)(nil), "autopilotrpc.ModifyStatusRequest")
	proto.RegisterType((*ModifyStatusResponse)(nil), "autopilotrpc.ModifyStatusResponse")
	proto.RegisterType((*QueryScoresRequest)(nil), "autopilotrpc.QueryScoresRequest")
	proto.RegisterType((*QueryScoresResponse)(nil), "autopilotrpc.QueryScoresResponse")
	proto.RegisterType((*QueryScoresResponse_HeuristicResult)(nil), "autopilotrpc.QueryScoresResponse.HeuristicResult")
	proto.RegisterType((*SetScoresRequest)(nil), "autopilotrpc.SetScoresRequest")
	proto.RegisterType((*SetScoresResponse)(nil), "autopilotrpc.SetScoresResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Autopilot service

type AutopilotClient interface {
	// *
	// Status returns whether the daemon's autopilot agent is active.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	// *
	// ModifyStatus is used to modify the status of the autopilot agent, like
	// enabling or disabling it.
	ModifyStatus(ctx context.Context, in *ModifyStatusRequest, opts ...grpc.CallOption) (*ModifyStatusResponse, error)
	// *
	// QueryScores queries all available autopilot heuristics, in addition to any
	// active combination of these heruristics, for the scores they would give to
	// the given nodes.
	QueryScores(ctx context.Context, in *QueryScoresRequest, opts ...grpc.CallOption) (*QueryScoresResponse, error)
	// *
	// SetScores attempts to set the scores used by the running autopilot agent,
	// if the external scoring heuristic is enabled.
	SetScores(ctx context.Context, in *SetScoresRequest, opts ...grpc.CallOption) (*SetScoresResponse, error)
}

type autopilotClient struct {
	cc *grpc.ClientConn
}

func NewAutopilotClient(cc *grpc.ClientConn) AutopilotClient {
	return &autopilotClient{cc}
}

func (c *autopilotClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/autopilotrpc.Autopilot/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autopilotClient) ModifyStatus(ctx context.Context, in *ModifyStatusRequest, opts ...grpc.CallOption) (*ModifyStatusResponse, error) {
	out := new(ModifyStatusResponse)
	err := grpc.Invoke(ctx, "/autopilotrpc.Autopilot/ModifyStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autopilotClient) QueryScores(ctx context.Context, in *QueryScoresRequest, opts ...grpc.CallOption) (*QueryScoresResponse, error) {
	out := new(QueryScoresResponse)
	err := grpc.Invoke(ctx, "/autopilotrpc.Autopilot/QueryScores", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autopilotClient) SetScores(ctx context.Context, in *SetScoresRequest, opts ...grpc.CallOption) (*SetScoresResponse, error) {
	out := new(SetScoresResponse)
	err := grpc.Invoke(ctx, "/autopilotrpc.Autopilot/SetScores", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Autopilot service

type AutopilotServer interface {
	// *
	// Status returns whether the daemon's autopilot agent is active.
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	// *
	// ModifyStatus is used to modify the status of the autopilot agent, like
	// enabling or disabling it.
	ModifyStatus(context.Context, *ModifyStatusRequest) (*ModifyStatusResponse, error)
	// *
	// QueryScores queries all available autopilot heuristics, in addition to any
	// active combination of these heruristics, for the scores they would give to
	// the given nodes.
	QueryScores(context.Context, *QueryScoresRequest) (*QueryScoresResponse, error)
	// *
	// SetScores attempts to set the scores used by the running autopilot agent,
	// if the external scoring heuristic is enabled.
	SetScores(context.Context, *SetScoresRequest) (*SetScoresResponse, error)
}

func RegisterAutopilotServer(s *grpc.Server, srv AutopilotServer) {
	s.RegisterService(&_Autopilot_serviceDesc, srv)
}

func _Autopilot_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutopilotServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autopilotrpc.Autopilot/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutopilotServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Autopilot_ModifyStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutopilotServer).ModifyStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autopilotrpc.Autopilot/ModifyStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutopilotServer).ModifyStatus(ctx, req.(*ModifyStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Autopilot_QueryScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutopilotServer).QueryScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autopilotrpc.Autopilot/QueryScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutopilotServer).QueryScores(ctx, req.(*QueryScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Autopilot_SetScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutopilotServer).SetScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autopilotrpc.Autopilot/SetScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutopilotServer).SetScores(ctx, req.(*SetScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Autopilot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "autopilotrpc.Autopilot",
	HandlerType: (*AutopilotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Autopilot_Status_Handler,
		},
		{
			MethodName: "ModifyStatus",
			Handler:    _Autopilot_ModifyStatus_Handler,
		},
		{
			MethodName: "QueryScores",
			Handler:    _Autopilot_QueryScores_Handler,
		},
		{
			MethodName: "SetScores",
			Handler:    _Autopilot_SetScores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "autopilotrpc/autopilot.proto",
}

func init() { proto.RegisterFile("autopilotrpc/autopilot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0x4c, 0x74, 0xcb, 0xeb, 0x60, 0xc3, 0x9d, 0xa6, 0x28, 0x54, 0xd0, 0x45, 0x1c,
	0x2a, 0x24, 0x52, 0x51, 0x38, 0x00, 0x12, 0x07, 0x86, 0x90, 0x90, 0x80, 0x03, 0x2e, 0xbb, 0x70,
	0x99, 0xd2, 0xcc, 0xb4, 0x56, 0x8d, 0x1d, 0xec, 0xe7, 0xa1, 0xfc, 0x43, 0x5c, 0xf9, 0x1b, 0x38,
	0xf2, 0x5f, 0xa1, 0xc5, 0x49, 0x49, 0xaa, 0x12, 0x84, 0xb4, 0x5b, 0xde, 0x8f, 0x7c, 0x9e, 0xdf,
	0xd7, 0x5f, 0x19, 0x86, 0xa9, 0x45, 0x95, 0x73, 0xa1, 0x50, 0xe7, 0xd9, 0x64, 0x1d, 0x24, 0xb9,
	0x56, 0xa8, 0xc8, 0x7e, 0xb3, 0x1a, 0x1f, 0xc0, 0xcd, 0x19, 0xa6, 0x68, 0x0d, 0x65, 0x5f, 0x2d,
	0x33, 0x18, 0x8f, 0xe1, 0x56, 0x9d, 0x30, 0xb9, 0x92, 0x86, 0x91, 0x63, 0xe8, 0xa5, 0x19, 0xf2,
	0x4b, 0x16, 0x7a, 0x23, 0x6f, 0xbc, 0x47, 0xab, 0x28, 0x7e, 0x08, 0x83, 0xf7, 0xea, 0x82, 0x7f,
	0x2e, 0x5a, 0x80, 0xab, 0x76, 0x26, 0xd3, 0xb9, 0x58, 0xb7, 0xbb, 0x28, 0x3e, 0x86, 0xa3, 0x76,
	0xbb, 0xc3, 0xc7, 0x1f, 0x81, 0x7c, 0xb0, 0x4c, 0x17, 0xb3, 0x4c, 0x69, 0xb6, 0xa6, 0x84, 0xb0,
	0x9b, 0xdb, 0xf9, 0x8a, 0x15, 0x26, 0xf4, 0x46, 0x3b, 0xe3, 0x80, 0xd6, 0x21, 0xb9, 0x0f, 0x84,
	0x2f, 0xa4, 0xd2, 0xec, 0x5c, 0xa8, 0x2c, 0x15, 0xe7, 0x06, 0x53, 0x64, 0xa1, 0x5f, 0xce, 0xda,
	0x93, 0xca, 0xc5, 0xf1, 0x77, 0x1f, 0x06, 0x2d, 0x6c, 0xb5, 0xcc, 0x5b, 0xd8, 0xd5, 0xcc, 0x58,
	0x81, 0x8e, 0xdb, 0x9f, 0x3e, 0x4a, 0x9a, 0x7a, 0x24, 0x5b, 0xfe, 0x49, 0xde, 0x30, 0xab, 0xb9,
	0x41, 0x9e, 0xd1, 0xf2, 0x4f, 0x5a, 0x13, 0xa2, 0x9f, 0x1e, 0x1c, 0x6c, 0x14, 0xc9, 0x10, 0x82,
	0x65, 0x9d, 0x2a, 0x15, 0x08, 0xe8, 0x9f, 0x04, 0x39, 0x83, 0x9e, 0x29, 0xe1, 0xa1, 0x5f, 0x4e,
	0x7f, 0xf1, 0xdf, 0xd3, 0x13, 0x57, 0x7e, 0x2d, 0x51, 0x17, 0xb4, 0x82, 0x45, 0xcf, 0xa0, 0xdf,
	0x48, 0x93, 0x43, 0xd8, 0x59, 0xb1, 0xa2, 0x9a, 0x7e, 0xf5, 0x49, 0x8e, 0xe0, 0xc6, 0x65, 0x2a,
	0xac, 0xd3, 0xc9, 0xa3, 0x2e, 0x78, 0xee, 0x3f, 0xf5, 0xe2, 0x1f, 0x1e, 0x1c, 0xce, 0x18, 0xb6,
	0xd5, 0xef, 0x5e, 0xe2, 0x74, 0x63, 0x89, 0x07, 0xed, 0x25, 0x36, 0x69, 0xd7, 0x7d, 0xe2, 0x01,
	0xdc, 0x6e, 0x8c, 0x70, 0x2a, 0x4d, 0x7f, 0xf9, 0x10, 0xbc, 0xac, 0x4f, 0x41, 0x5e, 0x41, 0xcf,
	0xb9, 0x8c, 0xdc, 0xd9, 0x38, 0x5b, 0xd3, 0xaa, 0xd1, 0x70, 0x7b, 0xb1, 0xb2, 0xca, 0x19, 0xec,
	0x37, 0x0d, 0x4b, 0x4e, 0xda, 0xdd, 0x5b, 0xbc, 0x1f, 0xc5, 0x5d, 0x2d, 0x15, 0x96, 0x42, 0xbf,
	0x71, 0xcd, 0x64, 0xd4, 0xe1, 0x00, 0x07, 0x3d, 0xf9, 0xa7, 0x47, 0xc8, 0x3b, 0x08, 0xd6, 0x92,
	0x90, 0xbb, 0xdd, 0xd7, 0x11, 0xdd, 0xfb, 0x6b, 0xdd, 0xd1, 0x4e, 0x9f, 0x7c, 0x9a, 0x2e, 0x38,
	0x2e, 0xed, 0x3c, 0xc9, 0xd4, 0x97, 0x89, 0xe0, 0x8b, 0x25, 0x4a, 0x2e, 0x17, 0x92, 0xe1, 0x37,
	0xa5, 0x57, 0x13, 0x21, 0x2f, 0x26, 0x42, 0xb6, 0x9e, 0x16, 0x9d, 0x67, 0xf3, 0x5e, 0xf9, 0xbc,
	0x3c, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x4e, 0x74, 0x27, 0x7e, 0x04, 0x00, 0x00,
}
