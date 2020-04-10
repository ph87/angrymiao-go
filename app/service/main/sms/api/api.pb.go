// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SendReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendReply) Reset()         { *m = SendReply{} }
func (m *SendReply) String() string { return proto.CompactTextString(m) }
func (*SendReply) ProtoMessage()    {}
func (*SendReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *SendReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendReply.Merge(m, src)
}
func (m *SendReply) XXX_Size() int {
	return m.Size()
}
func (m *SendReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SendReply.DiscardUnknown(m)
}

var xxx_messageInfo_SendReply proto.InternalMessageInfo

type SendReq struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty" form:"mobile"`
	RegionId             string   `protobuf:"bytes,2,opt,name=regionId,proto3" json:"regionId,omitempty" form:"regionId"`
	Country              string   `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty" form:"country"`
	SignName             string   `protobuf:"bytes,4,opt,name=signName,proto3" json:"signName,omitempty" form:"signName"`
	Tcode                string   `protobuf:"bytes,5,opt,name=tcode,proto3" json:"tcode,omitempty" form:"tcode"`
	Tparam               string   `protobuf:"bytes,6,opt,name=tparam,proto3" json:"tparam,omitempty" form:"tparam"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendReq) Reset()         { *m = SendReq{} }
func (m *SendReq) String() string { return proto.CompactTextString(m) }
func (*SendReq) ProtoMessage()    {}
func (*SendReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *SendReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendReq.Merge(m, src)
}
func (m *SendReq) XXX_Size() int {
	return m.Size()
}
func (m *SendReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendReq proto.InternalMessageInfo

func (m *SendReq) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *SendReq) GetRegionId() string {
	if m != nil {
		return m.RegionId
	}
	return ""
}

func (m *SendReq) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *SendReq) GetSignName() string {
	if m != nil {
		return m.SignName
	}
	return ""
}

func (m *SendReq) GetTcode() string {
	if m != nil {
		return m.Tcode
	}
	return ""
}

func (m *SendReq) GetTparam() string {
	if m != nil {
		return m.Tparam
	}
	return ""
}

func init() {
	proto.RegisterType((*SendReply)(nil), "sms.service.v1.SendReply")
	proto.RegisterType((*SendReq)(nil), "sms.service.v1.SendReq")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd1, 0xbb, 0x4e, 0xf3, 0x30,
	0x14, 0x07, 0xf0, 0x2f, 0xbd, 0x7e, 0x35, 0x50, 0xc0, 0x20, 0x61, 0x3a, 0x24, 0xc8, 0x03, 0x02,
	0x09, 0x5c, 0x15, 0xb6, 0x6e, 0x74, 0x63, 0x61, 0x70, 0x37, 0xb6, 0x24, 0x75, 0x83, 0xa5, 0x3a,
	0x0e, 0xce, 0x45, 0xca, 0x93, 0xc0, 0x23, 0x31, 0xf2, 0x04, 0x11, 0x0a, 0x6f, 0x90, 0x27, 0x40,
	0xb1, 0x93, 0x4a, 0x54, 0x6c, 0x39, 0xe7, 0xff, 0xcb, 0x91, 0xce, 0x31, 0x18, 0xb9, 0x11, 0x27,
	0x91, 0x92, 0x89, 0x84, 0xe3, 0x58, 0xc4, 0x24, 0x66, 0x2a, 0xe3, 0x3e, 0x23, 0xd9, 0x6c, 0x72,
	0x1b, 0xf0, 0xe4, 0x25, 0xf5, 0x88, 0x2f, 0xc5, 0x34, 0x90, 0x81, 0x9c, 0x6a, 0xe6, 0xa5, 0x6b,
	0x5d, 0xe9, 0x42, 0x7f, 0x99, 0xdf, 0xf1, 0x1e, 0x18, 0x2d, 0x59, 0xb8, 0xa2, 0x2c, 0xda, 0xe4,
	0xf8, 0xad, 0x03, 0x86, 0xa6, 0x7a, 0x85, 0xd7, 0x60, 0x20, 0xa4, 0xc7, 0x37, 0x0c, 0x59, 0x17,
	0xd6, 0xd5, 0x68, 0x71, 0x5c, 0x15, 0xce, 0xc1, 0x5a, 0x2a, 0x31, 0xc7, 0xa6, 0x8f, 0x69, 0x03,
	0xe0, 0x14, 0xfc, 0x57, 0x2c, 0xe0, 0x32, 0x7c, 0x5c, 0xa1, 0x8e, 0xc6, 0x27, 0x55, 0xe1, 0x1c,
	0x1a, 0xdc, 0x26, 0x98, 0x6e, 0x11, 0xbc, 0x01, 0x43, 0x5f, 0xa6, 0x61, 0xa2, 0x72, 0xd4, 0xd5,
	0x1e, 0x56, 0x85, 0x33, 0x36, 0xbe, 0x09, 0x30, 0x6d, 0x49, 0x3d, 0x3e, 0xe6, 0x41, 0xf8, 0xe4,
	0x0a, 0x86, 0x7a, 0xbb, 0xe3, 0xdb, 0x04, 0xd3, 0x2d, 0x82, 0x97, 0xa0, 0x9f, 0xf8, 0x72, 0xc5,
	0x50, 0x5f, 0xeb, 0xa3, 0xaa, 0x70, 0xf6, 0x8d, 0xd6, 0x6d, 0x4c, 0x4d, 0x5c, 0xaf, 0x98, 0x44,
	0xae, 0x72, 0x05, 0x1a, 0xec, 0xae, 0x68, 0xfa, 0x98, 0x36, 0xe0, 0xee, 0x01, 0x74, 0x97, 0x22,
	0x86, 0x73, 0xd0, 0xab, 0xef, 0x03, 0xcf, 0xc8, 0xef, 0xab, 0x93, 0xe6, 0x6a, 0x93, 0xf3, 0xbf,
	0x83, 0x68, 0x93, 0x2f, 0x4e, 0x3f, 0x4a, 0xdb, 0xfa, 0x2c, 0x6d, 0xeb, 0xab, 0xb4, 0xad, 0xf7,
	0x6f, 0xfb, 0xdf, 0x73, 0x27, 0x9b, 0x79, 0x03, 0xfd, 0x0c, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xcb, 0x65, 0xb4, 0xe2, 0xd2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// smsClient is the client API for Sms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmsClient interface {
	// Send send sms
	Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendReply, error)
}

type smsClient struct {
	cc *grpc.ClientConn
}

func NewSmsClient(cc *grpc.ClientConn) SmsClient {
	return &smsClient{cc}
}

func (c *smsClient) Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendReply, error) {
	out := new(SendReply)
	err := c.cc.Invoke(ctx, "/sms.service.v1.Sms/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServer is the server API for Sms service.
type SmsServer interface {
	// Send send sms
	Send(context.Context, *SendReq) (*SendReply, error)
}

// UnimplementedSmsServer can be embedded to have forward compatible implementations.
type UnimplementedSmsServer struct {
}

func (*UnimplementedSmsServer) Send(ctx context.Context, req *SendReq) (*SendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterSmsServer(s *grpc.Server, srv SmsServer) {
	s.RegisterService(&_Sms_serviceDesc, srv)
}

func _Sms_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.service.v1.Sms/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).Send(ctx, req.(*SendReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sms_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sms.service.v1.Sms",
	HandlerType: (*SmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Sms_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *SendReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *SendReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tparam) > 0 {
		i -= len(m.Tparam)
		copy(dAtA[i:], m.Tparam)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Tparam)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Tcode) > 0 {
		i -= len(m.Tcode)
		copy(dAtA[i:], m.Tcode)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Tcode)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SignName) > 0 {
		i -= len(m.SignName)
		copy(dAtA[i:], m.SignName)
		i = encodeVarintApi(dAtA, i, uint64(len(m.SignName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Country) > 0 {
		i -= len(m.Country)
		copy(dAtA[i:], m.Country)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Country)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RegionId) > 0 {
		i -= len(m.RegionId)
		copy(dAtA[i:], m.RegionId)
		i = encodeVarintApi(dAtA, i, uint64(len(m.RegionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Mobile) > 0 {
		i -= len(m.Mobile)
		copy(dAtA[i:], m.Mobile)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Mobile)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SendReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Mobile)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.RegionId)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Country)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.SignName)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Tcode)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Tparam)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: SendReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SendReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: SendReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mobile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mobile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Country", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Country = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tcode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tcode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tparam", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tparam = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
