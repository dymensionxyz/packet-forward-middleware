// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: router/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the router genesis state
type GenesisState struct {
	Params          Params                    `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	InFlightPackets map[string]InFlightPacket `protobuf:"bytes,2,rep,name=in_flight_packets,json=inFlightPackets,proto3" json:"in_flight_packets" yaml:"in_flight_packets" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4940b763c55c4e0b, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetInFlightPackets() map[string]InFlightPacket {
	if m != nil {
		return m.InFlightPackets
	}
	return nil
}

// Params defines the set of IBC router parameters.
type Params struct {
	FeePercentage github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=fee_percentage,json=feePercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_percentage" yaml:"fee_percentage"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4940b763c55c4e0b, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

type InFlightPacket struct {
	OriginalSenderAddress string `protobuf:"bytes,1,opt,name=original_sender_address,json=originalSenderAddress,proto3" json:"original_sender_address,omitempty"`
	RefundChannelId       string `protobuf:"bytes,2,opt,name=refund_channel_id,json=refundChannelId,proto3" json:"refund_channel_id,omitempty"`
	RefundPortId          string `protobuf:"bytes,3,opt,name=refund_port_id,json=refundPortId,proto3" json:"refund_port_id,omitempty"`
	RetriesRemaining      int32  `protobuf:"varint,4,opt,name=retries_remaining,json=retriesRemaining,proto3" json:"retries_remaining,omitempty"`
	Timeout               uint64 `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (m *InFlightPacket) Reset()         { *m = InFlightPacket{} }
func (m *InFlightPacket) String() string { return proto.CompactTextString(m) }
func (*InFlightPacket) ProtoMessage()    {}
func (*InFlightPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_4940b763c55c4e0b, []int{2}
}
func (m *InFlightPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InFlightPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InFlightPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InFlightPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InFlightPacket.Merge(m, src)
}
func (m *InFlightPacket) XXX_Size() int {
	return m.Size()
}
func (m *InFlightPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_InFlightPacket.DiscardUnknown(m)
}

var xxx_messageInfo_InFlightPacket proto.InternalMessageInfo

func (m *InFlightPacket) GetOriginalSenderAddress() string {
	if m != nil {
		return m.OriginalSenderAddress
	}
	return ""
}

func (m *InFlightPacket) GetRefundChannelId() string {
	if m != nil {
		return m.RefundChannelId
	}
	return ""
}

func (m *InFlightPacket) GetRefundPortId() string {
	if m != nil {
		return m.RefundPortId
	}
	return ""
}

func (m *InFlightPacket) GetRetriesRemaining() int32 {
	if m != nil {
		return m.RetriesRemaining
	}
	return 0
}

func (m *InFlightPacket) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "router.v1.GenesisState")
	proto.RegisterMapType((map[string]InFlightPacket)(nil), "router.v1.GenesisState.InFlightPacketsEntry")
	proto.RegisterType((*Params)(nil), "router.v1.Params")
	proto.RegisterType((*InFlightPacket)(nil), "router.v1.InFlightPacket")
}

func init() { proto.RegisterFile("router/v1/genesis.proto", fileDescriptor_4940b763c55c4e0b) }

var fileDescriptor_4940b763c55c4e0b = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0xf3, 0x42, 0x99, 0x96, 0xb4, 0xb1, 0x5a, 0xd5, 0x74, 0xe1, 0x44, 0x11, 0x42, 0x11,
	0x10, 0x5b, 0x6d, 0x25, 0x84, 0xba, 0x23, 0x3c, 0xaa, 0xec, 0x22, 0x77, 0x87, 0x84, 0xac, 0xa9,
	0xe7, 0xc6, 0x19, 0xc5, 0x9e, 0xb1, 0x66, 0xc6, 0x2e, 0xe1, 0x2b, 0x58, 0xf3, 0x45, 0x5d, 0x76,
	0x89, 0x58, 0x44, 0x90, 0xfc, 0x41, 0xbf, 0x00, 0xd9, 0xe3, 0x40, 0x22, 0x58, 0x79, 0xe6, 0x9c,
	0x73, 0x8f, 0xef, 0xb9, 0xba, 0x83, 0x4e, 0x04, 0x4f, 0x15, 0x08, 0x37, 0x3b, 0x73, 0x43, 0x60,
	0x20, 0xa9, 0x74, 0x12, 0xc1, 0x15, 0x37, 0x5b, 0x9a, 0x70, 0xb2, 0xb3, 0xd3, 0xa3, 0x90, 0x87,
	0xbc, 0x40, 0xdd, 0xfc, 0xa4, 0x05, 0xfd, 0x6f, 0x55, 0xb4, 0x7f, 0xa5, 0x4b, 0xae, 0x15, 0x56,
	0x60, 0xba, 0xa8, 0x99, 0x60, 0x81, 0x63, 0x69, 0x19, 0x3d, 0x63, 0xb0, 0x77, 0xde, 0x71, 0xfe,
	0x58, 0x38, 0x93, 0x82, 0x18, 0xd5, 0xef, 0x96, 0xdd, 0x8a, 0x57, 0xca, 0xcc, 0x2f, 0xa8, 0x43,
	0x99, 0x3f, 0x8d, 0x68, 0x38, 0x53, 0x7e, 0x82, 0x83, 0x39, 0x28, 0x69, 0x55, 0x7b, 0xb5, 0xc1,
	0xde, 0xf9, 0xcb, 0xad, 0xda, 0xed, 0x9f, 0x38, 0x63, 0xf6, 0xa1, 0xd0, 0x4f, 0xb4, 0xfc, 0x3d,
	0x53, 0x62, 0x31, 0xea, 0xe5, 0xb6, 0x0f, 0xcb, 0xae, 0xb5, 0xc0, 0x71, 0x74, 0xd9, 0xff, 0xc7,
	0xb4, 0xef, 0x1d, 0xd0, 0xdd, 0xba, 0xd3, 0x4f, 0xe8, 0xe8, 0x7f, 0x56, 0xe6, 0x21, 0xaa, 0xcd,
	0x61, 0x51, 0x24, 0x68, 0x79, 0xf9, 0xd1, 0x74, 0x51, 0x23, 0xc3, 0x51, 0x0a, 0x56, 0xb5, 0x48,
	0xf5, 0x64, 0xab, 0xb3, 0x5d, 0x07, 0x4f, 0xeb, 0x2e, 0xab, 0xaf, 0x8d, 0xfe, 0x67, 0xd4, 0xd4,
	0x91, 0x4d, 0x86, 0xda, 0x53, 0x00, 0x3f, 0x01, 0x11, 0x00, 0x53, 0x38, 0x04, 0xed, 0x3d, 0xba,
	0xca, 0x7b, 0xfe, 0xb1, 0xec, 0x3e, 0x0b, 0xa9, 0x9a, 0xa5, 0x37, 0x4e, 0xc0, 0x63, 0x37, 0xe0,
	0x32, 0xe6, 0xb2, 0xfc, 0x0c, 0x25, 0x99, 0xbb, 0x6a, 0x91, 0x80, 0x74, 0xde, 0x41, 0xf0, 0xb0,
	0xec, 0x1e, 0xeb, 0x74, 0xbb, 0x6e, 0x7d, 0xef, 0xf1, 0x14, 0x60, 0xf2, 0xf7, 0xfe, 0xcb, 0x40,
	0xed, 0xdd, 0xbe, 0xcc, 0x57, 0xe8, 0x84, 0x0b, 0x1a, 0x52, 0x86, 0x23, 0x5f, 0x02, 0x23, 0x20,
	0x7c, 0x4c, 0x88, 0x00, 0x29, 0xcb, 0x9c, 0xc7, 0x1b, 0xfa, 0xba, 0x60, 0xdf, 0x68, 0xd2, 0x7c,
	0x8e, 0x3a, 0x02, 0xa6, 0x29, 0x23, 0x7e, 0x30, 0xc3, 0x8c, 0x41, 0xe4, 0x53, 0x52, 0x4c, 0xa1,
	0xe5, 0x1d, 0x68, 0xe2, 0xad, 0xc6, 0xc7, 0xc4, 0x7c, 0x8a, 0xda, 0xa5, 0x36, 0xe1, 0x42, 0xe5,
	0xc2, 0x5a, 0x21, 0xdc, 0xd7, 0xe8, 0x84, 0x0b, 0x35, 0x26, 0xe6, 0x8b, 0xdc, 0x51, 0x09, 0x0a,
	0xd2, 0x17, 0x10, 0x63, 0xca, 0x28, 0x0b, 0xad, 0x7a, 0xcf, 0x18, 0x34, 0xbc, 0xc3, 0x92, 0xf0,
	0x36, 0xb8, 0x69, 0xa1, 0x47, 0x8a, 0xc6, 0xc0, 0x53, 0x65, 0x35, 0x7a, 0xc6, 0xa0, 0xee, 0x6d,
	0xae, 0xa3, 0xe0, 0x6e, 0x65, 0x1b, 0xf7, 0x2b, 0xdb, 0xf8, 0xb9, 0xb2, 0x8d, 0xaf, 0x6b, 0xbb,
	0x72, 0xbf, 0xb6, 0x2b, 0xdf, 0xd7, 0x76, 0xe5, 0xe3, 0x78, 0x6b, 0x9a, 0x52, 0x09, 0xcc, 0x42,
	0x88, 0x78, 0x06, 0xc3, 0x0c, 0x98, 0x4a, 0x05, 0x48, 0x57, 0x6f, 0xc4, 0x70, 0xca, 0xc5, 0x2d,
	0x16, 0x64, 0x18, 0x53, 0x42, 0x22, 0xb8, 0xc5, 0x02, 0xdc, 0xec, 0xc2, 0x2d, 0xdf, 0x42, 0x31,
	0xf4, 0x9b, 0x66, 0xb1, 0xe6, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x43, 0x50, 0x11, 0xd0,
	0x22, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InFlightPackets) > 0 {
		for k := range m.InFlightPackets {
			v := m.InFlightPackets[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeePercentage.Size()
		i -= size
		if _, err := m.FeePercentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *InFlightPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InFlightPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InFlightPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timeout != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x28
	}
	if m.RetriesRemaining != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RetriesRemaining))
		i--
		dAtA[i] = 0x20
	}
	if len(m.RefundPortId) > 0 {
		i -= len(m.RefundPortId)
		copy(dAtA[i:], m.RefundPortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.RefundPortId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RefundChannelId) > 0 {
		i -= len(m.RefundChannelId)
		copy(dAtA[i:], m.RefundChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.RefundChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OriginalSenderAddress) > 0 {
		i -= len(m.OriginalSenderAddress)
		copy(dAtA[i:], m.OriginalSenderAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.OriginalSenderAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.InFlightPackets) > 0 {
		for k, v := range m.InFlightPackets {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + 1 + l + sovGenesis(uint64(l))
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.FeePercentage.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *InFlightPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OriginalSenderAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.RefundChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.RefundPortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.RetriesRemaining != 0 {
		n += 1 + sovGenesis(uint64(m.RetriesRemaining))
	}
	if m.Timeout != 0 {
		n += 1 + sovGenesis(uint64(m.Timeout))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InFlightPackets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InFlightPackets == nil {
				m.InFlightPackets = make(map[string]InFlightPacket)
			}
			var mapkey string
			mapvalue := &InFlightPacket{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenesis
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenesis
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &InFlightPacket{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.InFlightPackets[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeePercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *InFlightPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: InFlightPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InFlightPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalSenderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalSenderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefundChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefundChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefundPortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefundPortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetriesRemaining", wireType)
			}
			m.RetriesRemaining = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RetriesRemaining |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
