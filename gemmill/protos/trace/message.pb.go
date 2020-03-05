// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gemmill/protos/trace/message.proto

package trace

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MsgType int32

const (
	MsgType_None     MsgType = 0
	MsgType_Request  MsgType = 1
	MsgType_Responce MsgType = 2
)

var MsgType_name = map[int32]string{
	0: "None",
	1: "Request",
	2: "Responce",
}

var MsgType_value = map[string]int32{
	"None":     0,
	"Request":  1,
	"Responce": 2,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6e1ec0e473c7137, []int{0}
}

type TraceRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceRequest) Reset()         { *m = TraceRequest{} }
func (m *TraceRequest) String() string { return proto.CompactTextString(m) }
func (*TraceRequest) ProtoMessage()    {}
func (*TraceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6e1ec0e473c7137, []int{0}
}
func (m *TraceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceRequest.Merge(m, src)
}
func (m *TraceRequest) XXX_Size() int {
	return m.Size()
}
func (m *TraceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TraceRequest proto.InternalMessageInfo

func (m *TraceRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type TraceResponse struct {
	RequestHash          []byte   `protobuf:"bytes,1,opt,name=RequestHash,proto3" json:"RequestHash,omitempty"`
	Resp                 []byte   `protobuf:"bytes,2,opt,name=Resp,proto3" json:"Resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceResponse) Reset()         { *m = TraceResponse{} }
func (m *TraceResponse) String() string { return proto.CompactTextString(m) }
func (*TraceResponse) ProtoMessage()    {}
func (*TraceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6e1ec0e473c7137, []int{1}
}
func (m *TraceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceResponse.Merge(m, src)
}
func (m *TraceResponse) XXX_Size() int {
	return m.Size()
}
func (m *TraceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TraceResponse proto.InternalMessageInfo

func (m *TraceResponse) GetRequestHash() []byte {
	if m != nil {
		return m.RequestHash
	}
	return nil
}

func (m *TraceResponse) GetResp() []byte {
	if m != nil {
		return m.Resp
	}
	return nil
}

type TraceMessage struct {
	Type                 MsgType  `protobuf:"varint,1,opt,name=Type,proto3,enum=trace.MsgType" json:"Type,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceMessage) Reset()         { *m = TraceMessage{} }
func (m *TraceMessage) String() string { return proto.CompactTextString(m) }
func (*TraceMessage) ProtoMessage()    {}
func (*TraceMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6e1ec0e473c7137, []int{2}
}
func (m *TraceMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceMessage.Merge(m, src)
}
func (m *TraceMessage) XXX_Size() int {
	return m.Size()
}
func (m *TraceMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TraceMessage proto.InternalMessageInfo

func (m *TraceMessage) GetType() MsgType {
	if m != nil {
		return m.Type
	}
	return MsgType_None
}

func (m *TraceMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("trace.MsgType", MsgType_name, MsgType_value)
	proto.RegisterType((*TraceRequest)(nil), "trace.TraceRequest")
	proto.RegisterType((*TraceResponse)(nil), "trace.TraceResponse")
	proto.RegisterType((*TraceMessage)(nil), "trace.TraceMessage")
}

func init() { proto.RegisterFile("gemmill/protos/trace/message.proto", fileDescriptor_a6e1ec0e473c7137) }

var fileDescriptor_a6e1ec0e473c7137 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0xeb, 0x2a, 0xd0, 0xea, 0x35, 0x54, 0x91, 0xa7, 0x4e, 0x51, 0xe5, 0x09, 0x31, 0xc4,
	0x12, 0x3d, 0x01, 0x7f, 0x8a, 0x58, 0xca, 0x10, 0x75, 0x62, 0x73, 0xd3, 0x27, 0x27, 0x52, 0x62,
	0x9b, 0xd8, 0x1d, 0xb8, 0x09, 0x47, 0x62, 0xe4, 0x08, 0x28, 0x5c, 0x04, 0xd9, 0xb1, 0x80, 0x81,
	0xed, 0xf3, 0xe7, 0x9f, 0xdf, 0xfb, 0xc9, 0xc0, 0x24, 0x76, 0x5d, 0xd3, 0xb6, 0xdc, 0xf4, 0xda,
	0x69, 0xcb, 0x5d, 0x2f, 0x2a, 0xe4, 0x1d, 0x5a, 0x2b, 0x24, 0x16, 0xa1, 0xa4, 0x67, 0xa1, 0x64,
	0x0c, 0xd2, 0xbd, 0x0f, 0x25, 0xbe, 0x9c, 0xd0, 0x3a, 0x4a, 0x21, 0xb9, 0x17, 0x4e, 0xac, 0xc8,
	0x9a, 0x5c, 0xa6, 0x65, 0xc8, 0x6c, 0x0b, 0x17, 0x91, 0xb1, 0x46, 0x2b, 0x8b, 0x74, 0x0d, 0x8b,
	0xc8, 0x3f, 0x0a, 0x5b, 0x47, 0xf6, 0x6f, 0xe5, 0xc7, 0x78, 0x7a, 0x35, 0x1d, 0xc7, 0xf8, 0xcc,
	0x1e, 0xe2, 0xaa, 0xdd, 0xe8, 0x41, 0x19, 0x24, 0xfb, 0x57, 0x83, 0xe1, 0xf9, 0xf2, 0x7a, 0x59,
	0x04, 0xa1, 0x62, 0x67, 0xa5, 0x6f, 0xcb, 0x70, 0xf7, 0xa3, 0x33, 0xfd, 0xd5, 0xb9, 0x2a, 0x60,
	0x16, 0x21, 0x3a, 0x87, 0xe4, 0x49, 0x2b, 0xcc, 0x26, 0x74, 0x01, 0xb3, 0xb8, 0x3f, 0x23, 0x34,
	0x85, 0xf9, 0xe8, 0x5a, 0x61, 0x36, 0xbd, 0xdd, 0xbe, 0x0f, 0x39, 0xf9, 0x18, 0x72, 0xf2, 0x39,
	0xe4, 0xe4, 0xed, 0x2b, 0x9f, 0x3c, 0x6f, 0x64, 0xe3, 0xea, 0xd3, 0xa1, 0xa8, 0x74, 0xc7, 0x8f,
	0xc2, 0x98, 0x16, 0x8f, 0x12, 0x7b, 0x7e, 0xa3, 0xd4, 0x5d, 0x2d, 0x1a, 0xc5, 0xff, 0xfb, 0xbe,
	0xc3, 0x79, 0x38, 0x6d, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x28, 0x83, 0xda, 0xf2, 0x5d, 0x01,
	0x00, 0x00,
}

func (m *TraceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TraceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Resp) > 0 {
		i -= len(m.Resp)
		copy(dAtA[i:], m.Resp)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Resp)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RequestHash) > 0 {
		i -= len(m.RequestHash)
		copy(dAtA[i:], m.RequestHash)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.RequestHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TraceMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TraceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TraceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestHash)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Resp)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TraceMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMessage(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TraceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: TraceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *TraceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: TraceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestHash = append(m.RequestHash[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestHash == nil {
				m.RequestHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resp", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resp = append(m.Resp[:0], dAtA[iNdEx:postIndex]...)
			if m.Resp == nil {
				m.Resp = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *TraceMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: TraceMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= MsgType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)