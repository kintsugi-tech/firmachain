// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: firmachain/contract/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
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

type MsgCreateContractFile struct {
	Creator            string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	FileHash           string   `protobuf:"bytes,2,opt,name=fileHash,proto3" json:"fileHash,omitempty"`
	TimeStamp          uint64   `protobuf:"varint,3,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	OwnerList          []string `protobuf:"bytes,4,rep,name=ownerList,proto3" json:"ownerList,omitempty"`
	MetaDataJsonString string   `protobuf:"bytes,5,opt,name=metaDataJsonString,proto3" json:"metaDataJsonString,omitempty"`
}

func (m *MsgCreateContractFile) Reset()         { *m = MsgCreateContractFile{} }
func (m *MsgCreateContractFile) String() string { return proto.CompactTextString(m) }
func (*MsgCreateContractFile) ProtoMessage()    {}
func (*MsgCreateContractFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a32ed0eb42a3cae9, []int{0}
}
func (m *MsgCreateContractFile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateContractFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateContractFile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateContractFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateContractFile.Merge(m, src)
}
func (m *MsgCreateContractFile) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateContractFile) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateContractFile.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateContractFile proto.InternalMessageInfo

func (m *MsgCreateContractFile) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateContractFile) GetFileHash() string {
	if m != nil {
		return m.FileHash
	}
	return ""
}

func (m *MsgCreateContractFile) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *MsgCreateContractFile) GetOwnerList() []string {
	if m != nil {
		return m.OwnerList
	}
	return nil
}

func (m *MsgCreateContractFile) GetMetaDataJsonString() string {
	if m != nil {
		return m.MetaDataJsonString
	}
	return ""
}

type MsgCreateContractFileResponse struct {
}

func (m *MsgCreateContractFileResponse) Reset()         { *m = MsgCreateContractFileResponse{} }
func (m *MsgCreateContractFileResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateContractFileResponse) ProtoMessage()    {}
func (*MsgCreateContractFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a32ed0eb42a3cae9, []int{1}
}
func (m *MsgCreateContractFileResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateContractFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateContractFileResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateContractFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateContractFileResponse.Merge(m, src)
}
func (m *MsgCreateContractFileResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateContractFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateContractFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateContractFileResponse proto.InternalMessageInfo

type MsgAddContractLog struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ContractHash string `protobuf:"bytes,2,opt,name=contractHash,proto3" json:"contractHash,omitempty"`
	TimeStamp    uint64 `protobuf:"varint,3,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	EventName    string `protobuf:"bytes,4,opt,name=eventName,proto3" json:"eventName,omitempty"`
	OwnerAddress string `protobuf:"bytes,5,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
	JsonString   string `protobuf:"bytes,6,opt,name=jsonString,proto3" json:"jsonString,omitempty"`
}

func (m *MsgAddContractLog) Reset()         { *m = MsgAddContractLog{} }
func (m *MsgAddContractLog) String() string { return proto.CompactTextString(m) }
func (*MsgAddContractLog) ProtoMessage()    {}
func (*MsgAddContractLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_a32ed0eb42a3cae9, []int{2}
}
func (m *MsgAddContractLog) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddContractLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddContractLog.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddContractLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddContractLog.Merge(m, src)
}
func (m *MsgAddContractLog) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddContractLog) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddContractLog.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddContractLog proto.InternalMessageInfo

func (m *MsgAddContractLog) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAddContractLog) GetContractHash() string {
	if m != nil {
		return m.ContractHash
	}
	return ""
}

func (m *MsgAddContractLog) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *MsgAddContractLog) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *MsgAddContractLog) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *MsgAddContractLog) GetJsonString() string {
	if m != nil {
		return m.JsonString
	}
	return ""
}

type MsgAddContractLogResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgAddContractLogResponse) Reset()         { *m = MsgAddContractLogResponse{} }
func (m *MsgAddContractLogResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddContractLogResponse) ProtoMessage()    {}
func (*MsgAddContractLogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a32ed0eb42a3cae9, []int{3}
}
func (m *MsgAddContractLogResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddContractLogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddContractLogResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddContractLogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddContractLogResponse.Merge(m, src)
}
func (m *MsgAddContractLogResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddContractLogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddContractLogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddContractLogResponse proto.InternalMessageInfo

func (m *MsgAddContractLogResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreateContractFile)(nil), "firmachain.firmachain.contract.MsgCreateContractFile")
	proto.RegisterType((*MsgCreateContractFileResponse)(nil), "firmachain.firmachain.contract.MsgCreateContractFileResponse")
	proto.RegisterType((*MsgAddContractLog)(nil), "firmachain.firmachain.contract.MsgAddContractLog")
	proto.RegisterType((*MsgAddContractLogResponse)(nil), "firmachain.firmachain.contract.MsgAddContractLogResponse")
}

func init() { proto.RegisterFile("firmachain/contract/tx.proto", fileDescriptor_a32ed0eb42a3cae9) }

var fileDescriptor_a32ed0eb42a3cae9 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcf, 0xee, 0xd2, 0x40,
	0x10, 0x66, 0x4b, 0x45, 0x99, 0x18, 0x12, 0x37, 0x31, 0xa9, 0x0d, 0xae, 0xa4, 0x07, 0x25, 0x31,
	0x69, 0xfd, 0x13, 0x0e, 0x1e, 0x3c, 0x20, 0xc6, 0x18, 0x03, 0x9a, 0x94, 0x9b, 0x17, 0xb3, 0xb4,
	0x4b, 0x59, 0x43, 0xbb, 0xcd, 0xee, 0x8a, 0x78, 0xf1, 0x01, 0x3c, 0x19, 0xdf, 0xc6, 0x37, 0xf0,
	0x88, 0x37, 0x8f, 0x06, 0x5e, 0xc4, 0xb4, 0x50, 0x5a, 0xb4, 0x41, 0xfd, 0xdd, 0x66, 0xbf, 0x6f,
	0xbe, 0x99, 0xfd, 0x76, 0x76, 0xa0, 0x3b, 0xe7, 0x32, 0xa6, 0xc1, 0x82, 0xf2, 0xc4, 0x0b, 0x44,
	0xa2, 0x25, 0x0d, 0xb4, 0xa7, 0xd7, 0x6e, 0x2a, 0x85, 0x16, 0x98, 0x94, 0xac, 0x5b, 0x09, 0x8b,
	0x44, 0xfb, 0x4e, 0x9d, 0xba, 0x08, 0xde, 0xcc, 0xf9, 0x92, 0xed, 0x0b, 0xd9, 0xb7, 0xcf, 0x26,
	0x2e, 0x45, 0xb4, 0xcf, 0x73, 0xbe, 0x22, 0xb8, 0x3e, 0x51, 0xd1, 0x48, 0x32, 0xaa, 0xd9, 0xe8,
	0xc0, 0x3f, 0xe3, 0x4b, 0x86, 0x2d, 0xb8, 0x1c, 0x64, 0xa8, 0x90, 0x16, 0xea, 0xa1, 0x7e, 0xdb,
	0x2f, 0x8e, 0xd8, 0x86, 0x2b, 0x59, 0xa7, 0xe7, 0x54, 0x2d, 0x2c, 0x23, 0xa7, 0x8e, 0x67, 0xdc,
	0x85, 0xb6, 0xe6, 0x31, 0x9b, 0x6a, 0x1a, 0xa7, 0x56, 0xb3, 0x87, 0xfa, 0xa6, 0x5f, 0x02, 0x19,
	0x2b, 0xde, 0x27, 0x4c, 0x8e, 0xb9, 0xd2, 0x96, 0xd9, 0x6b, 0xf6, 0xdb, 0x7e, 0x09, 0x60, 0x17,
	0x70, 0xcc, 0x34, 0x7d, 0x4a, 0x35, 0x7d, 0xa1, 0x44, 0x32, 0xd5, 0x92, 0x27, 0x91, 0x75, 0x29,
	0xef, 0x50, 0xc3, 0x38, 0xb7, 0xe0, 0x66, 0xed, 0xd5, 0x7d, 0xa6, 0x52, 0x91, 0x28, 0xe6, 0x7c,
	0x47, 0x70, 0x6d, 0xa2, 0xa2, 0x61, 0x18, 0x16, 0xf4, 0x58, 0x44, 0x67, 0x8c, 0x39, 0x70, 0xb5,
	0x78, 0xa2, 0x8a, 0xb9, 0x13, 0xec, 0xef, 0x06, 0xd9, 0x8a, 0x25, 0xfa, 0x25, 0x8d, 0x99, 0x65,
	0xe6, 0xf2, 0x12, 0xc8, 0xea, 0xe7, 0x6e, 0x87, 0x61, 0x28, 0x99, 0x52, 0x07, 0x6b, 0x27, 0x18,
	0x26, 0x00, 0x6f, 0x4b, 0xf3, 0xad, 0x3c, 0xa3, 0x82, 0x38, 0x77, 0xe1, 0xc6, 0x1f, 0x96, 0x0a,
	0xc3, 0xb8, 0x03, 0x06, 0x0f, 0x73, 0x57, 0xa6, 0x6f, 0xf0, 0xf0, 0xc1, 0x17, 0x03, 0x9a, 0x13,
	0x15, 0xe1, 0x4f, 0x08, 0x70, 0xcd, 0x88, 0x07, 0xee, 0xf9, 0xef, 0xe6, 0xd6, 0x3e, 0xaf, 0xfd,
	0xf8, 0x42, 0xb2, 0xe3, 0x25, 0x3f, 0x42, 0xe7, 0xb7, 0x89, 0xdc, 0xff, 0x87, 0x82, 0xa7, 0x12,
	0xfb, 0xd1, 0x7f, 0x4b, 0x8a, 0xfe, 0x4f, 0x5e, 0x7d, 0xdb, 0x12, 0xb4, 0xd9, 0x12, 0xf4, 0x73,
	0x4b, 0xd0, 0xe7, 0x1d, 0x69, 0x6c, 0x76, 0xa4, 0xf1, 0x63, 0x47, 0x1a, 0xaf, 0x07, 0x11, 0xd7,
	0x8b, 0x77, 0x33, 0x37, 0x10, 0xb1, 0x57, 0xd9, 0x9f, 0x4a, 0xb8, 0xba, 0x37, 0xf0, 0xd6, 0x95,
	0xbd, 0xfd, 0x90, 0x32, 0x35, 0x6b, 0xe5, 0xab, 0xf4, 0xf0, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc6, 0x0c, 0xab, 0xc3, 0xdb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateContractFile(ctx context.Context, in *MsgCreateContractFile, opts ...grpc.CallOption) (*MsgCreateContractFileResponse, error)
	AddContractLog(ctx context.Context, in *MsgAddContractLog, opts ...grpc.CallOption) (*MsgAddContractLogResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateContractFile(ctx context.Context, in *MsgCreateContractFile, opts ...grpc.CallOption) (*MsgCreateContractFileResponse, error) {
	out := new(MsgCreateContractFileResponse)
	err := c.cc.Invoke(ctx, "/firmachain.firmachain.contract.Msg/CreateContractFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddContractLog(ctx context.Context, in *MsgAddContractLog, opts ...grpc.CallOption) (*MsgAddContractLogResponse, error) {
	out := new(MsgAddContractLogResponse)
	err := c.cc.Invoke(ctx, "/firmachain.firmachain.contract.Msg/AddContractLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateContractFile(context.Context, *MsgCreateContractFile) (*MsgCreateContractFileResponse, error)
	AddContractLog(context.Context, *MsgAddContractLog) (*MsgAddContractLogResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateContractFile(ctx context.Context, req *MsgCreateContractFile) (*MsgCreateContractFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContractFile not implemented")
}
func (*UnimplementedMsgServer) AddContractLog(ctx context.Context, req *MsgAddContractLog) (*MsgAddContractLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContractLog not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateContractFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateContractFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateContractFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmachain.firmachain.contract.Msg/CreateContractFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateContractFile(ctx, req.(*MsgCreateContractFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddContractLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddContractLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddContractLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmachain.firmachain.contract.Msg/AddContractLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddContractLog(ctx, req.(*MsgAddContractLog))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "firmachain.firmachain.contract.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContractFile",
			Handler:    _Msg_CreateContractFile_Handler,
		},
		{
			MethodName: "AddContractLog",
			Handler:    _Msg_AddContractLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "firmachain/contract/tx.proto",
}

func (m *MsgCreateContractFile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateContractFile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateContractFile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MetaDataJsonString) > 0 {
		i -= len(m.MetaDataJsonString)
		copy(dAtA[i:], m.MetaDataJsonString)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MetaDataJsonString)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OwnerList) > 0 {
		for iNdEx := len(m.OwnerList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.OwnerList[iNdEx])
			copy(dAtA[i:], m.OwnerList[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.OwnerList[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.TimeStamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.FileHash) > 0 {
		i -= len(m.FileHash)
		copy(dAtA[i:], m.FileHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FileHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateContractFileResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateContractFileResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateContractFileResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgAddContractLog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddContractLog) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddContractLog) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.JsonString) > 0 {
		i -= len(m.JsonString)
		copy(dAtA[i:], m.JsonString)
		i = encodeVarintTx(dAtA, i, uint64(len(m.JsonString)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.EventName) > 0 {
		i -= len(m.EventName)
		copy(dAtA[i:], m.EventName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EventName)))
		i--
		dAtA[i] = 0x22
	}
	if m.TimeStamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ContractHash) > 0 {
		i -= len(m.ContractHash)
		copy(dAtA[i:], m.ContractHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddContractLogResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddContractLogResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddContractLogResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateContractFile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.FileHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeStamp != 0 {
		n += 1 + sovTx(uint64(m.TimeStamp))
	}
	if len(m.OwnerList) > 0 {
		for _, s := range m.OwnerList {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.MetaDataJsonString)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateContractFileResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgAddContractLog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ContractHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeStamp != 0 {
		n += 1 + sovTx(uint64(m.TimeStamp))
	}
	l = len(m.EventName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.JsonString)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddContractLogResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateContractFile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateContractFile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateContractFile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			m.TimeStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeStamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerList = append(m.OwnerList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaDataJsonString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetaDataJsonString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateContractFileResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateContractFileResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateContractFileResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddContractLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddContractLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddContractLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			m.TimeStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeStamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JsonString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JsonString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddContractLogResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddContractLogResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddContractLogResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
