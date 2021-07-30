// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onos/mho/mho.proto

package mho

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MhoParamType is enumeration type of MHO parameters
type MhoParamType int32

const (
	MhoParamType_ALL           MhoParamType = 0
	MhoParamType_A3OFFSET      MhoParamType = 1
	MhoParamType_HYSTERESIS    MhoParamType = 2
	MhoParamType_TIMETOTRIGGER MhoParamType = 3
)

var MhoParamType_name = map[int32]string{
	0: "ALL",
	1: "A3OFFSET",
	2: "HYSTERESIS",
	3: "TIMETOTRIGGER",
}

var MhoParamType_value = map[string]int32{
	"ALL":           0,
	"A3OFFSET":      1,
	"HYSTERESIS":    2,
	"TIMETOTRIGGER": 3,
}

func (x MhoParamType) String() string {
	return proto.EnumName(MhoParamType_name, int32(x))
}

func (MhoParamType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2241cd1f7f71b11c, []int{0}
}

type GetMhoParamRequest struct {
	// hoParamType is a type of handover parameter
	HoParamType MhoParamType `protobuf:"varint,1,opt,name=hoParamType,proto3,enum=onos.mho.MhoParamType" json:"hoParamType,omitempty"`
}

func (m *GetMhoParamRequest) Reset()         { *m = GetMhoParamRequest{} }
func (m *GetMhoParamRequest) String() string { return proto.CompactTextString(m) }
func (*GetMhoParamRequest) ProtoMessage()    {}
func (*GetMhoParamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2241cd1f7f71b11c, []int{0}
}
func (m *GetMhoParamRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMhoParamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMhoParamRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetMhoParamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMhoParamRequest.Merge(m, src)
}
func (m *GetMhoParamRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetMhoParamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMhoParamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMhoParamRequest proto.InternalMessageInfo

func (m *GetMhoParamRequest) GetHoParamType() MhoParamType {
	if m != nil {
		return m.HoParamType
	}
	return MhoParamType_ALL
}

type GetMhoParamResponse struct {
	// hoParamType is a type of handover parameter
	HoParamType MhoParamType `protobuf:"varint,1,opt,name=hoParamType,proto3,enum=onos.mho.MhoParamType" json:"hoParamType,omitempty"`
	// A3-Offset value
	A3Offset int32 `protobuf:"varint,2,opt,name=a3Offset,proto3" json:"a3Offset,omitempty"`
	// Hysteresis value
	Hysteresis int32 `protobuf:"varint,3,opt,name=hysteresis,proto3" json:"hysteresis,omitempty"`
	// Time-to-Trigger value
	TimeToTrigger int32 `protobuf:"varint,4,opt,name=timeToTrigger,proto3" json:"timeToTrigger,omitempty"`
}

func (m *GetMhoParamResponse) Reset()         { *m = GetMhoParamResponse{} }
func (m *GetMhoParamResponse) String() string { return proto.CompactTextString(m) }
func (*GetMhoParamResponse) ProtoMessage()    {}
func (*GetMhoParamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2241cd1f7f71b11c, []int{1}
}
func (m *GetMhoParamResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMhoParamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMhoParamResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetMhoParamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMhoParamResponse.Merge(m, src)
}
func (m *GetMhoParamResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetMhoParamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMhoParamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMhoParamResponse proto.InternalMessageInfo

func (m *GetMhoParamResponse) GetHoParamType() MhoParamType {
	if m != nil {
		return m.HoParamType
	}
	return MhoParamType_ALL
}

func (m *GetMhoParamResponse) GetA3Offset() int32 {
	if m != nil {
		return m.A3Offset
	}
	return 0
}

func (m *GetMhoParamResponse) GetHysteresis() int32 {
	if m != nil {
		return m.Hysteresis
	}
	return 0
}

func (m *GetMhoParamResponse) GetTimeToTrigger() int32 {
	if m != nil {
		return m.TimeToTrigger
	}
	return 0
}

type SetMhoParamRequest struct {
	// hoParamType is a type of handover parameter
	HoParamType MhoParamType `protobuf:"varint,1,opt,name=hoParamType,proto3,enum=onos.mho.MhoParamType" json:"hoParamType,omitempty"`
	// A3-Offset value
	A3Offset int32 `protobuf:"varint,2,opt,name=a3Offset,proto3" json:"a3Offset,omitempty"`
	// Hysteresis value
	Hysteresis int32 `protobuf:"varint,3,opt,name=hysteresis,proto3" json:"hysteresis,omitempty"`
	// Time-to-Trigger value
	TimeToTrigger int32 `protobuf:"varint,4,opt,name=timeToTrigger,proto3" json:"timeToTrigger,omitempty"`
}

func (m *SetMhoParamRequest) Reset()         { *m = SetMhoParamRequest{} }
func (m *SetMhoParamRequest) String() string { return proto.CompactTextString(m) }
func (*SetMhoParamRequest) ProtoMessage()    {}
func (*SetMhoParamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2241cd1f7f71b11c, []int{2}
}
func (m *SetMhoParamRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetMhoParamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetMhoParamRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetMhoParamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMhoParamRequest.Merge(m, src)
}
func (m *SetMhoParamRequest) XXX_Size() int {
	return m.Size()
}
func (m *SetMhoParamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMhoParamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetMhoParamRequest proto.InternalMessageInfo

func (m *SetMhoParamRequest) GetHoParamType() MhoParamType {
	if m != nil {
		return m.HoParamType
	}
	return MhoParamType_ALL
}

func (m *SetMhoParamRequest) GetA3Offset() int32 {
	if m != nil {
		return m.A3Offset
	}
	return 0
}

func (m *SetMhoParamRequest) GetHysteresis() int32 {
	if m != nil {
		return m.Hysteresis
	}
	return 0
}

func (m *SetMhoParamRequest) GetTimeToTrigger() int32 {
	if m != nil {
		return m.TimeToTrigger
	}
	return 0
}

type SetMhoParamResponse struct {
	// success is a result whether MHO param is set successfully or not
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *SetMhoParamResponse) Reset()         { *m = SetMhoParamResponse{} }
func (m *SetMhoParamResponse) String() string { return proto.CompactTextString(m) }
func (*SetMhoParamResponse) ProtoMessage()    {}
func (*SetMhoParamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2241cd1f7f71b11c, []int{3}
}
func (m *SetMhoParamResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetMhoParamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetMhoParamResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetMhoParamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMhoParamResponse.Merge(m, src)
}
func (m *SetMhoParamResponse) XXX_Size() int {
	return m.Size()
}
func (m *SetMhoParamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMhoParamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetMhoParamResponse proto.InternalMessageInfo

func (m *SetMhoParamResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterEnum("onos.mho.MhoParamType", MhoParamType_name, MhoParamType_value)
	proto.RegisterType((*GetMhoParamRequest)(nil), "onos.mho.GetMhoParamRequest")
	proto.RegisterType((*GetMhoParamResponse)(nil), "onos.mho.GetMhoParamResponse")
	proto.RegisterType((*SetMhoParamRequest)(nil), "onos.mho.SetMhoParamRequest")
	proto.RegisterType((*SetMhoParamResponse)(nil), "onos.mho.SetMhoParamResponse")
}

func init() { proto.RegisterFile("onos/mho/mho.proto", fileDescriptor_2241cd1f7f71b11c) }

var fileDescriptor_2241cd1f7f71b11c = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xbd, 0x4e, 0xc2, 0x50,
	0x18, 0xed, 0xa5, 0x2a, 0xcd, 0x67, 0x21, 0xf5, 0x23, 0x31, 0x0d, 0xd1, 0x86, 0x10, 0x07, 0xe2,
	0x50, 0x12, 0x58, 0x5c, 0x31, 0x29, 0x3f, 0x81, 0x8a, 0xe9, 0xbd, 0x8b, 0x23, 0x92, 0x0b, 0x65,
	0x28, 0x17, 0x7b, 0xcb, 0xc0, 0x5b, 0xf8, 0x06, 0xbe, 0x82, 0xbe, 0x85, 0x23, 0xa3, 0xa3, 0x81,
	0x17, 0x31, 0x34, 0xa0, 0x25, 0x82, 0x83, 0x31, 0x0e, 0x77, 0xf8, 0xce, 0xf9, 0xee, 0xc9, 0xc9,
	0xf9, 0x0e, 0xa0, 0x18, 0x0b, 0x59, 0x0e, 0x7c, 0xb1, 0x7a, 0xf6, 0x24, 0x14, 0x91, 0x40, 0x6d,
	0x85, 0xd9, 0x81, 0x2f, 0x8a, 0x37, 0x80, 0x0d, 0x1e, 0xb9, 0xbe, 0xb8, 0xed, 0x85, 0xbd, 0xc0,
	0xe3, 0x0f, 0x53, 0x2e, 0x23, 0xbc, 0x82, 0xe3, 0x35, 0xc2, 0x66, 0x13, 0x6e, 0x92, 0x02, 0x29,
	0x65, 0x2b, 0xa7, 0xf6, 0xe6, 0x97, 0xed, 0x26, 0x58, 0x2f, 0xb9, 0x5a, 0x7c, 0x21, 0x90, 0xdb,
	0x12, 0x94, 0x13, 0x31, 0x96, 0xfc, 0xf7, 0x8a, 0x98, 0x07, 0xad, 0x57, 0xed, 0x0e, 0x06, 0x92,
	0x47, 0x66, 0xaa, 0x40, 0x4a, 0x87, 0xde, 0xe7, 0x8c, 0x16, 0x80, 0x3f, 0x93, 0x11, 0x0f, 0xb9,
	0x1c, 0x49, 0x53, 0x8d, 0xd9, 0x04, 0x82, 0x17, 0x90, 0x89, 0x46, 0x01, 0x67, 0x82, 0x85, 0xa3,
	0xe1, 0x90, 0x87, 0xe6, 0x41, 0xbc, 0xb2, 0x0d, 0x16, 0x9f, 0x09, 0x20, 0xfd, 0xc3, 0x10, 0xfe,
	0xc1, 0x72, 0x19, 0x72, 0x74, 0x47, 0xca, 0x26, 0xa4, 0xe5, 0xb4, 0xdf, 0xe7, 0x52, 0xc6, 0x76,
	0x35, 0x6f, 0x33, 0x5e, 0x36, 0x41, 0x4f, 0xfa, 0xc5, 0x34, 0xa8, 0xb5, 0x4e, 0xc7, 0x50, 0x50,
	0x07, 0xad, 0x56, 0xed, 0xd6, 0xeb, 0xd4, 0x61, 0x06, 0xc1, 0x2c, 0x40, 0xf3, 0x8e, 0x32, 0xc7,
	0x73, 0x68, 0x8b, 0x1a, 0x29, 0x3c, 0x81, 0x0c, 0x6b, 0xb9, 0x0e, 0xeb, 0x32, 0xaf, 0xd5, 0x68,
	0x38, 0x9e, 0xa1, 0x56, 0x9e, 0x08, 0xa8, 0xae, 0x2f, 0xb0, 0x0d, 0x7a, 0xe2, 0xd0, 0x12, 0xcf,
	0xbe, 0x92, 0xf9, 0xde, 0xa8, 0xfc, 0xf9, 0x1e, 0x76, 0x6d, 0xbc, 0x0d, 0x3a, 0xdd, 0x23, 0x46,
	0x7f, 0x14, 0xdb, 0x91, 0xc2, 0xb5, 0xf9, 0xba, 0xb0, 0xc8, 0x7c, 0x61, 0x91, 0xf7, 0x85, 0x45,
	0x1e, 0x97, 0x96, 0x32, 0x5f, 0x5a, 0xca, 0xdb, 0xd2, 0x52, 0xee, 0x8f, 0xe2, 0xfa, 0x57, 0x3f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x02, 0xa6, 0xdb, 0x55, 0x14, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MhoClient is the client API for Mho service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MhoClient interface {
	// To get MHO parameters
	GetMhoParams(ctx context.Context, in *GetMhoParamRequest, opts ...grpc.CallOption) (*GetMhoParamResponse, error)
	// To set MHO parameters
	SetMhoParams(ctx context.Context, in *SetMhoParamRequest, opts ...grpc.CallOption) (*SetMhoParamResponse, error)
}

type mhoClient struct {
	cc *grpc.ClientConn
}

func NewMhoClient(cc *grpc.ClientConn) MhoClient {
	return &mhoClient{cc}
}

func (c *mhoClient) GetMhoParams(ctx context.Context, in *GetMhoParamRequest, opts ...grpc.CallOption) (*GetMhoParamResponse, error) {
	out := new(GetMhoParamResponse)
	err := c.cc.Invoke(ctx, "/onos.mho.Mho/GetMhoParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mhoClient) SetMhoParams(ctx context.Context, in *SetMhoParamRequest, opts ...grpc.CallOption) (*SetMhoParamResponse, error) {
	out := new(SetMhoParamResponse)
	err := c.cc.Invoke(ctx, "/onos.mho.Mho/SetMhoParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MhoServer is the server API for Mho service.
type MhoServer interface {
	// To get MHO parameters
	GetMhoParams(context.Context, *GetMhoParamRequest) (*GetMhoParamResponse, error)
	// To set MHO parameters
	SetMhoParams(context.Context, *SetMhoParamRequest) (*SetMhoParamResponse, error)
}

// UnimplementedMhoServer can be embedded to have forward compatible implementations.
type UnimplementedMhoServer struct {
}

func (*UnimplementedMhoServer) GetMhoParams(ctx context.Context, req *GetMhoParamRequest) (*GetMhoParamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMhoParams not implemented")
}
func (*UnimplementedMhoServer) SetMhoParams(ctx context.Context, req *SetMhoParamRequest) (*SetMhoParamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMhoParams not implemented")
}

func RegisterMhoServer(s *grpc.Server, srv MhoServer) {
	s.RegisterService(&_Mho_serviceDesc, srv)
}

func _Mho_GetMhoParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMhoParamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MhoServer).GetMhoParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onos.mho.Mho/GetMhoParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MhoServer).GetMhoParams(ctx, req.(*GetMhoParamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mho_SetMhoParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMhoParamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MhoServer).SetMhoParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onos.mho.Mho/SetMhoParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MhoServer).SetMhoParams(ctx, req.(*SetMhoParamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mho_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onos.mho.Mho",
	HandlerType: (*MhoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMhoParams",
			Handler:    _Mho_GetMhoParams_Handler,
		},
		{
			MethodName: "SetMhoParams",
			Handler:    _Mho_SetMhoParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onos/mho/mho.proto",
}

func (m *GetMhoParamRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMhoParamRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetMhoParamRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HoParamType != 0 {
		i = encodeVarintMho(dAtA, i, uint64(m.HoParamType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetMhoParamResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMhoParamResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetMhoParamResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeToTrigger != 0 {
		i = encodeVarintMho(dAtA, i, uint64(m.TimeToTrigger))
		i--
		dAtA[i] = 0x20
	}
	if m.Hysteresis != 0 {
		i = encodeVarintMho(dAtA, i, uint64(m.Hysteresis))
		i--
		dAtA[i] = 0x18
	}
	if m.A3Offset != 0 {
		i = encodeVarintMho(dAtA, i, uint64(m.A3Offset))
		i--
		dAtA[i] = 0x10
	}
	if m.HoParamType != 0 {
		i = encodeVarintMho(dAtA, i, uint64(m.HoParamType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SetMhoParamRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetMhoParamRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetMhoParamRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeToTrigger != 0 {
		i = encodeVarintMho(dAtA, i, uint64(m.TimeToTrigger))
		i--
		dAtA[i] = 0x20
	}
	if m.Hysteresis != 0 {
		i = encodeVarintMho(dAtA, i, uint64(m.Hysteresis))
		i--
		dAtA[i] = 0x18
	}
	if m.A3Offset != 0 {
		i = encodeVarintMho(dAtA, i, uint64(m.A3Offset))
		i--
		dAtA[i] = 0x10
	}
	if m.HoParamType != 0 {
		i = encodeVarintMho(dAtA, i, uint64(m.HoParamType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SetMhoParamResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetMhoParamResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetMhoParamResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMho(dAtA []byte, offset int, v uint64) int {
	offset -= sovMho(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetMhoParamRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HoParamType != 0 {
		n += 1 + sovMho(uint64(m.HoParamType))
	}
	return n
}

func (m *GetMhoParamResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HoParamType != 0 {
		n += 1 + sovMho(uint64(m.HoParamType))
	}
	if m.A3Offset != 0 {
		n += 1 + sovMho(uint64(m.A3Offset))
	}
	if m.Hysteresis != 0 {
		n += 1 + sovMho(uint64(m.Hysteresis))
	}
	if m.TimeToTrigger != 0 {
		n += 1 + sovMho(uint64(m.TimeToTrigger))
	}
	return n
}

func (m *SetMhoParamRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HoParamType != 0 {
		n += 1 + sovMho(uint64(m.HoParamType))
	}
	if m.A3Offset != 0 {
		n += 1 + sovMho(uint64(m.A3Offset))
	}
	if m.Hysteresis != 0 {
		n += 1 + sovMho(uint64(m.Hysteresis))
	}
	if m.TimeToTrigger != 0 {
		n += 1 + sovMho(uint64(m.TimeToTrigger))
	}
	return n
}

func (m *SetMhoParamResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovMho(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMho(x uint64) (n int) {
	return sovMho(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetMhoParamRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMho
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetMhoParamRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMhoParamRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HoParamType", wireType)
			}
			m.HoParamType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HoParamType |= MhoParamType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMho(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMho
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetMhoParamResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMho
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetMhoParamResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMhoParamResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HoParamType", wireType)
			}
			m.HoParamType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HoParamType |= MhoParamType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A3Offset", wireType)
			}
			m.A3Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.A3Offset |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hysteresis", wireType)
			}
			m.Hysteresis = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hysteresis |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeToTrigger", wireType)
			}
			m.TimeToTrigger = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeToTrigger |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMho(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMho
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetMhoParamRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMho
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetMhoParamRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetMhoParamRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HoParamType", wireType)
			}
			m.HoParamType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HoParamType |= MhoParamType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A3Offset", wireType)
			}
			m.A3Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.A3Offset |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hysteresis", wireType)
			}
			m.Hysteresis = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hysteresis |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeToTrigger", wireType)
			}
			m.TimeToTrigger = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeToTrigger |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMho(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMho
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetMhoParamResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMho
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetMhoParamResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetMhoParamResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMho(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMho
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMho(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMho
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMho
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMho
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMho
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMho
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMho
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMho        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMho          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMho = fmt.Errorf("proto: unexpected end of group")
)
