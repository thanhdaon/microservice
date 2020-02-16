// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/calculator.proto

package pb

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

type SumRequest struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e2d3384312a837, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *SumRequest) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e2d3384312a837, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeDecomposeRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeDecomposeRequest) Reset()         { *m = PrimeDecomposeRequest{} }
func (m *PrimeDecomposeRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeDecomposeRequest) ProtoMessage()    {}
func (*PrimeDecomposeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e2d3384312a837, []int{2}
}

func (m *PrimeDecomposeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeDecomposeRequest.Unmarshal(m, b)
}
func (m *PrimeDecomposeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeDecomposeRequest.Marshal(b, m, deterministic)
}
func (m *PrimeDecomposeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeDecomposeRequest.Merge(m, src)
}
func (m *PrimeDecomposeRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeDecomposeRequest.Size(m)
}
func (m *PrimeDecomposeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeDecomposeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeDecomposeRequest proto.InternalMessageInfo

func (m *PrimeDecomposeRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeDecomposeResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeDecomposeResponse) Reset()         { *m = PrimeDecomposeResponse{} }
func (m *PrimeDecomposeResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeDecomposeResponse) ProtoMessage()    {}
func (*PrimeDecomposeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e2d3384312a837, []int{3}
}

func (m *PrimeDecomposeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeDecomposeResponse.Unmarshal(m, b)
}
func (m *PrimeDecomposeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeDecomposeResponse.Marshal(b, m, deterministic)
}
func (m *PrimeDecomposeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeDecomposeResponse.Merge(m, src)
}
func (m *PrimeDecomposeResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeDecomposeResponse.Size(m)
}
func (m *PrimeDecomposeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeDecomposeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeDecomposeResponse proto.InternalMessageInfo

func (m *PrimeDecomposeResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ComputeAverageRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageRequest) Reset()         { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()    {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e2d3384312a837, []int{4}
}

func (m *ComputeAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageRequest.Unmarshal(m, b)
}
func (m *ComputeAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageRequest.Merge(m, src)
}
func (m *ComputeAverageRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageRequest.Size(m)
}
func (m *ComputeAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageRequest proto.InternalMessageInfo

func (m *ComputeAverageRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ComputeAverageResponse struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageResponse) Reset()         { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()    {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e2d3384312a837, []int{5}
}

func (m *ComputeAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageResponse.Unmarshal(m, b)
}
func (m *ComputeAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageResponse.Merge(m, src)
}
func (m *ComputeAverageResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageResponse.Size(m)
}
func (m *ComputeAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageResponse proto.InternalMessageInfo

func (m *ComputeAverageResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type FindMaxRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxRequest) Reset()         { *m = FindMaxRequest{} }
func (m *FindMaxRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaxRequest) ProtoMessage()    {}
func (*FindMaxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e2d3384312a837, []int{6}
}

func (m *FindMaxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxRequest.Unmarshal(m, b)
}
func (m *FindMaxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxRequest.Marshal(b, m, deterministic)
}
func (m *FindMaxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxRequest.Merge(m, src)
}
func (m *FindMaxRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaxRequest.Size(m)
}
func (m *FindMaxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxRequest proto.InternalMessageInfo

func (m *FindMaxRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type FindMaxResponse struct {
	Max                  int32    `protobuf:"varint,1,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxResponse) Reset()         { *m = FindMaxResponse{} }
func (m *FindMaxResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaxResponse) ProtoMessage()    {}
func (*FindMaxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e2d3384312a837, []int{7}
}

func (m *FindMaxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxResponse.Unmarshal(m, b)
}
func (m *FindMaxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxResponse.Marshal(b, m, deterministic)
}
func (m *FindMaxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxResponse.Merge(m, src)
}
func (m *FindMaxResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaxResponse.Size(m)
}
func (m *FindMaxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxResponse proto.InternalMessageInfo

func (m *FindMaxResponse) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

type SquareRootRequest struct {
	Number               float64  `protobuf:"fixed64,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e2d3384312a837, []int{8}
}

func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetNumber() float64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootResponse struct {
	NumberRoot           float64  `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootResponse) Reset()         { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()    {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e2d3384312a837, []int{9}
}

func (m *SquareRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootResponse.Unmarshal(m, b)
}
func (m *SquareRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootResponse.Marshal(b, m, deterministic)
}
func (m *SquareRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootResponse.Merge(m, src)
}
func (m *SquareRootResponse) XXX_Size() int {
	return xxx_messageInfo_SquareRootResponse.Size(m)
}
func (m *SquareRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootResponse proto.InternalMessageInfo

func (m *SquareRootResponse) GetNumberRoot() float64 {
	if m != nil {
		return m.NumberRoot
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeDecomposeRequest)(nil), "calculator.PrimeDecomposeRequest")
	proto.RegisterType((*PrimeDecomposeResponse)(nil), "calculator.PrimeDecomposeResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calculator.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calculator.ComputeAverageResponse")
	proto.RegisterType((*FindMaxRequest)(nil), "calculator.FindMaxRequest")
	proto.RegisterType((*FindMaxResponse)(nil), "calculator.FindMaxResponse")
	proto.RegisterType((*SquareRootRequest)(nil), "calculator.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "calculator.SquareRootResponse")
}

func init() { proto.RegisterFile("pb/calculator.proto", fileDescriptor_82e2d3384312a837) }

var fileDescriptor_82e2d3384312a837 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdd, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xd7, 0xf5, 0x7d, 0xf7, 0xc2, 0xd9, 0xde, 0x39, 0x23, 0xab, 0x52, 0xf1, 0x63, 0x19,
	0xc2, 0x44, 0xd8, 0x86, 0x22, 0x78, 0xeb, 0x07, 0xe2, 0xcd, 0x44, 0x36, 0xaf, 0xf4, 0x62, 0xb4,
	0x5d, 0x94, 0xc1, 0xd2, 0x74, 0x69, 0x22, 0xfb, 0x57, 0xfc, 0x6f, 0xa5, 0x6d, 0xfa, 0x91, 0xd9,
	0xe9, 0x5d, 0xf3, 0x9c, 0xe7, 0xfc, 0x4e, 0x38, 0x4f, 0x0a, 0x3b, 0x81, 0x3b, 0xf0, 0x9c, 0x85,
	0x27, 0x17, 0x8e, 0x60, 0xbc, 0x1f, 0x70, 0x26, 0x18, 0x82, 0x5c, 0xc1, 0xcf, 0x00, 0x13, 0x49,
	0xc7, 0x64, 0x29, 0x49, 0x28, 0x50, 0x07, 0x1a, 0x6f, 0x73, 0x1e, 0x8a, 0xa9, 0x2f, 0xa9, 0x4b,
	0xf8, 0x9e, 0x71, 0x6c, 0xf4, 0xfe, 0x8e, 0xeb, 0xb1, 0xf6, 0x18, 0x4b, 0xa8, 0x0b, 0xff, 0x43,
	0xe2, 0x31, 0x7f, 0x96, 0x7a, 0xaa, 0xb1, 0xa7, 0x91, 0x88, 0x89, 0x09, 0x9f, 0x40, 0x3d, 0xa6,
	0x86, 0x01, 0xf3, 0x43, 0x82, 0x2c, 0xa8, 0x71, 0x12, 0xca, 0x85, 0x50, 0x40, 0x75, 0xc2, 0x03,
	0x68, 0x3f, 0xf1, 0x39, 0x25, 0x77, 0xc4, 0x63, 0x34, 0x60, 0x21, 0x49, 0xef, 0x61, 0x41, 0x4d,
	0xbb, 0x81, 0x3a, 0xe1, 0x21, 0x58, 0xeb, 0x0d, 0xbf, 0x8c, 0x38, 0x85, 0xf6, 0x2d, 0xa3, 0x81,
	0x14, 0xe4, 0xfa, 0x83, 0x70, 0xe7, 0x3d, 0x1b, 0xd1, 0x02, 0xd3, 0x97, 0x54, 0xb9, 0xa3, 0xcf,
	0x08, 0xbe, 0x6e, 0x2d, 0x85, 0x57, 0x33, 0x38, 0x86, 0xe6, 0xfd, 0xdc, 0x9f, 0x8d, 0x9c, 0xd5,
	0x66, 0x6a, 0x17, 0xb6, 0x32, 0x8f, 0xc2, 0xb5, 0xc0, 0xa4, 0xce, 0x2a, 0x35, 0x51, 0x67, 0x85,
	0xcf, 0x60, 0x7b, 0xb2, 0x94, 0x0e, 0x27, 0x63, 0xc6, 0x44, 0xf9, 0x12, 0x8c, 0x6c, 0x09, 0x97,
	0x80, 0x8a, 0x66, 0x05, 0x3d, 0x82, 0x7a, 0x52, 0x9f, 0x72, 0xc6, 0x84, 0x6a, 0x81, 0x44, 0x8a,
	0x8c, 0xe7, 0x9f, 0x26, 0x14, 0x82, 0x47, 0x57, 0x60, 0x4e, 0x24, 0x45, 0x56, 0xbf, 0xf0, 0x3c,
	0xf2, 0x97, 0x60, 0xef, 0x7e, 0xd3, 0x93, 0x39, 0xb8, 0x82, 0x5e, 0xa1, 0xa9, 0x87, 0x80, 0x3a,
	0x45, 0x73, 0x69, 0xa2, 0x36, 0xfe, 0xc9, 0x92, 0xa2, 0x87, 0x46, 0x04, 0xd7, 0x43, 0xd0, 0xe1,
	0xa5, 0x59, 0xea, 0xf0, 0xf2, 0x0c, 0x71, 0xa5, 0x67, 0xa0, 0x07, 0xf8, 0xa7, 0xb2, 0x40, 0x76,
	0xb1, 0x45, 0x0f, 0xd1, 0xde, 0x2f, 0xad, 0xe5, 0x9c, 0xa1, 0x81, 0x46, 0x00, 0x79, 0x06, 0xe8,
	0x40, 0x5b, 0xd6, 0x7a, 0x90, 0xf6, 0xe1, 0xa6, 0x72, 0x8a, 0xbc, 0xf9, 0xf3, 0x52, 0x0d, 0x5c,
	0xb7, 0x16, 0xff, 0x9e, 0x17, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xb4, 0xca, 0x8e, 0xb5,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeDecompose(ctx context.Context, in *PrimeDecomposeRequest, opts ...grpc.CallOption) (Calculator_PrimeDecomposeClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (Calculator_ComputeAverageClient, error)
	FindMax(ctx context.Context, opts ...grpc.CallOption) (Calculator_FindMaxClient, error)
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.calculator/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) PrimeDecompose(ctx context.Context, in *PrimeDecomposeRequest, opts ...grpc.CallOption) (Calculator_PrimeDecomposeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Calculator_serviceDesc.Streams[0], "/calculator.calculator/PrimeDecompose", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorPrimeDecomposeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calculator_PrimeDecomposeClient interface {
	Recv() (*PrimeDecomposeResponse, error)
	grpc.ClientStream
}

type calculatorPrimeDecomposeClient struct {
	grpc.ClientStream
}

func (x *calculatorPrimeDecomposeClient) Recv() (*PrimeDecomposeResponse, error) {
	m := new(PrimeDecomposeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (Calculator_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Calculator_serviceDesc.Streams[1], "/calculator.calculator/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorComputeAverageClient{stream}
	return x, nil
}

type Calculator_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) FindMax(ctx context.Context, opts ...grpc.CallOption) (Calculator_FindMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Calculator_serviceDesc.Streams[2], "/calculator.calculator/FindMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorFindMaxClient{stream}
	return x, nil
}

type Calculator_FindMaxClient interface {
	Send(*FindMaxRequest) error
	Recv() (*FindMaxResponse, error)
	grpc.ClientStream
}

type calculatorFindMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorFindMaxClient) Send(m *FindMaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorFindMaxClient) Recv() (*FindMaxResponse, error) {
	m := new(FindMaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calculator.calculator/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeDecompose(*PrimeDecomposeRequest, Calculator_PrimeDecomposeServer) error
	ComputeAverage(Calculator_ComputeAverageServer) error
	FindMax(Calculator_FindMaxServer) error
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServer) PrimeDecompose(req *PrimeDecomposeRequest, srv Calculator_PrimeDecomposeServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeDecompose not implemented")
}
func (*UnimplementedCalculatorServer) ComputeAverage(srv Calculator_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (*UnimplementedCalculatorServer) FindMax(srv Calculator_FindMaxServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMax not implemented")
}
func (*UnimplementedCalculatorServer) SquareRoot(ctx context.Context, req *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.calculator/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_PrimeDecompose_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeDecomposeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServer).PrimeDecompose(m, &calculatorPrimeDecomposeServer{stream})
}

type Calculator_PrimeDecomposeServer interface {
	Send(*PrimeDecomposeResponse) error
	grpc.ServerStream
}

type calculatorPrimeDecomposeServer struct {
	grpc.ServerStream
}

func (x *calculatorPrimeDecomposeServer) Send(m *PrimeDecomposeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Calculator_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).ComputeAverage(&calculatorComputeAverageServer{stream})
}

type Calculator_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calculator_FindMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).FindMax(&calculatorFindMaxServer{stream})
}

type Calculator_FindMaxServer interface {
	Send(*FindMaxResponse) error
	Recv() (*FindMaxRequest, error)
	grpc.ServerStream
}

type calculatorFindMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorFindMaxServer) Send(m *FindMaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorFindMaxServer) Recv() (*FindMaxRequest, error) {
	m := new(FindMaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calculator_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.calculator/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Calculator_Sum_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _Calculator_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeDecompose",
			Handler:       _Calculator_PrimeDecompose_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _Calculator_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMax",
			Handler:       _Calculator_FindMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/calculator.proto",
}
