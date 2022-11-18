// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eve/claim/v1beta1/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AirdropProposal struct {
	Title             string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description       string        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	AirdropStartTime  time.Time     `protobuf:"bytes,3,opt,name=airdrop_start_time,json=airdropStartTime,proto3,stdtime" json:"airdrop_start_time" yaml:"airdrop_start_time"`
	DurationOfAirdrop time.Duration `protobuf:"bytes,4,opt,name=duration_of_airdrop,json=durationOfAirdrop,proto3,stdduration" json:"duration_of_airdrop,omitempty" yaml:"duration_of_airdrop"`
	ClaimDenom        string        `protobuf:"bytes,6,opt,name=claim_denom,json=claimDenom,proto3" json:"claim_denom,omitempty"`
	AirdropList       []Claimable   `protobuf:"bytes,7,rep,name=airdrop_list,json=airdropList,proto3" json:"airdrop_list" yaml:"airdrop_list"`
}

func (m *AirdropProposal) Reset()         { *m = AirdropProposal{} }
func (m *AirdropProposal) String() string { return proto.CompactTextString(m) }
func (*AirdropProposal) ProtoMessage()    {}
func (*AirdropProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_606c3441e279e3e6, []int{0}
}
func (m *AirdropProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AirdropProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AirdropProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AirdropProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AirdropProposal.Merge(m, src)
}
func (m *AirdropProposal) XXX_Size() int {
	return m.Size()
}
func (m *AirdropProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AirdropProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AirdropProposal proto.InternalMessageInfo

func (m *AirdropProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AirdropProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AirdropProposal) GetAirdropStartTime() time.Time {
	if m != nil {
		return m.AirdropStartTime
	}
	return time.Time{}
}

func (m *AirdropProposal) GetDurationOfAirdrop() time.Duration {
	if m != nil {
		return m.DurationOfAirdrop
	}
	return 0
}

func (m *AirdropProposal) GetClaimDenom() string {
	if m != nil {
		return m.ClaimDenom
	}
	return ""
}

func (m *AirdropProposal) GetAirdropList() []Claimable {
	if m != nil {
		return m.AirdropList
	}
	return nil
}

func init() {
	proto.RegisterType((*AirdropProposal)(nil), "evenetwork.eve.claim.v1beta1.AirdropProposal")
}

func init() { proto.RegisterFile("eve/claim/v1beta1/proposal.proto", fileDescriptor_606c3441e279e3e6) }

var fileDescriptor_606c3441e279e3e6 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xb1, 0x6e, 0xd4, 0x40,
	0x10, 0x86, 0xcf, 0x1c, 0x04, 0xb1, 0x46, 0x02, 0x9c, 0x14, 0xce, 0x09, 0x6c, 0xcb, 0x02, 0x71,
	0x12, 0xb0, 0x56, 0x42, 0x47, 0x87, 0x89, 0x44, 0x83, 0x04, 0x3a, 0xa8, 0x68, 0xac, 0xb5, 0x3d,
	0x67, 0x56, 0xd8, 0x37, 0xd6, 0x7a, 0xcf, 0xe1, 0xde, 0x22, 0xa2, 0xe2, 0x91, 0x52, 0xa6, 0xa4,
	0x32, 0xe8, 0xae, 0x82, 0xf2, 0x9e, 0x00, 0xed, 0x7a, 0x2d, 0x11, 0x2e, 0xa2, 0xf3, 0xcc, 0xff,
	0xcf, 0xcc, 0xe7, 0x5f, 0x4b, 0x02, 0x68, 0x21, 0xca, 0x4a, 0xc6, 0xab, 0xa8, 0x3d, 0x4a, 0x41,
	0xb2, 0xa3, 0xa8, 0x16, 0x58, 0x63, 0xc3, 0x4a, 0x5a, 0x0b, 0x94, 0xe8, 0xdc, 0x87, 0x16, 0x16,
	0x20, 0x4f, 0x51, 0x7c, 0xa6, 0xd0, 0x02, 0xd5, 0x66, 0x6a, 0xcc, 0x93, 0x83, 0x02, 0x0b, 0xd4,
	0xc6, 0x48, 0x7d, 0xf5, 0x33, 0x13, 0xaf, 0x40, 0x2c, 0x4a, 0x88, 0x74, 0x95, 0x2e, 0xe7, 0x51,
	0xbe, 0x14, 0x4c, 0x72, 0x5c, 0x18, 0xdd, 0xff, 0x57, 0x97, 0xbc, 0x82, 0x46, 0xb2, 0xaa, 0x36,
	0x86, 0x87, 0xbb, 0x58, 0xba, 0x4a, 0x04, 0x64, 0x28, 0xf2, 0xde, 0x15, 0xfe, 0x1a, 0x93, 0x3b,
	0x2f, 0xb9, 0xc8, 0x05, 0xd6, 0xef, 0x0c, 0xb4, 0x73, 0x40, 0x6e, 0x48, 0x2e, 0x4b, 0x70, 0xad,
	0xc0, 0x9a, 0xde, 0x9a, 0xf5, 0x85, 0x13, 0x10, 0x3b, 0x87, 0x26, 0x13, 0xbc, 0x56, 0x14, 0xee,
	0x35, 0xad, 0xfd, 0xdd, 0x72, 0x90, 0x38, 0xac, 0x5f, 0x95, 0x34, 0x92, 0x09, 0x99, 0x28, 0x24,
	0x77, 0x1c, 0x58, 0x53, 0xfb, 0x78, 0x42, 0x7b, 0x5e, 0x3a, 0xf0, 0xd2, 0x0f, 0x03, 0x6f, 0xfc,
	0xe8, 0xbc, 0xf3, 0x47, 0xdb, 0xce, 0x3f, 0x5c, 0xb1, 0xaa, 0x7c, 0x11, 0xee, 0xee, 0x08, 0xcf,
	0x7e, 0xf8, 0xd6, 0xec, 0xae, 0x11, 0xde, 0xab, 0xbe, 0x9a, 0x76, 0xbe, 0x5a, 0x64, 0x7f, 0x88,
	0x25, 0xc1, 0x79, 0x62, 0x0c, 0xee, 0x75, 0x7d, 0xf2, 0x70, 0xe7, 0xe4, 0x89, 0xf1, 0xc6, 0xaf,
	0xd5, 0xc5, 0xdf, 0x9d, 0xff, 0xe0, 0x8a, 0xe9, 0xa7, 0x58, 0x71, 0x09, 0x55, 0x2d, 0x57, 0xdb,
	0xce, 0x9f, 0xf4, 0x48, 0x57, 0xd8, 0xc2, 0x6f, 0x8a, 0xe9, 0xde, 0xa0, 0xbc, 0x9d, 0x9b, 0x14,
	0x1d, 0x9f, 0xd8, 0x7d, 0xce, 0x39, 0x2c, 0xb0, 0x72, 0xf7, 0x74, 0x4e, 0x44, 0xb7, 0x4e, 0x54,
	0xc7, 0x39, 0x25, 0xb7, 0x87, 0x5f, 0x2c, 0x79, 0x23, 0xdd, 0x9b, 0xc1, 0x78, 0x6a, 0x1f, 0x3f,
	0xa6, 0xff, 0x7b, 0x24, 0xf4, 0x95, 0xaa, 0x58, 0x5a, 0x42, 0xfc, 0xc4, 0xb0, 0x5f, 0x5a, 0xb2,
	0xed, 0xfc, 0xfd, 0xcb, 0xe9, 0xa9, 0x6e, 0x38, 0xb3, 0x4d, 0xf9, 0x86, 0x37, 0x32, 0x8e, 0xcf,
	0xd7, 0x9e, 0x75, 0xb1, 0xf6, 0xac, 0x9f, 0x6b, 0xcf, 0x3a, 0xdb, 0x78, 0xa3, 0x8b, 0x8d, 0x37,
	0xfa, 0xbe, 0xf1, 0x46, 0x1f, 0xa7, 0x05, 0x97, 0x9f, 0x96, 0x29, 0xcd, 0xb0, 0x8a, 0xa0, 0x85,
	0x67, 0x86, 0x43, 0x7d, 0x47, 0x5f, 0xcc, 0x23, 0x92, 0xab, 0x1a, 0x9a, 0x74, 0x4f, 0x87, 0xf9,
	0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xb1, 0xd6, 0x37, 0xf5, 0x02, 0x00, 0x00,
}

func (m *AirdropProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AirdropProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AirdropProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AirdropList) > 0 {
		for iNdEx := len(m.AirdropList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AirdropList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ClaimDenom) > 0 {
		i -= len(m.ClaimDenom)
		copy(dAtA[i:], m.ClaimDenom)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.ClaimDenom)))
		i--
		dAtA[i] = 0x32
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationOfAirdrop, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfAirdrop):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintProposal(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintProposal(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AirdropProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime)
	n += 1 + l + sovProposal(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfAirdrop)
	n += 1 + l + sovProposal(uint64(l))
	l = len(m.ClaimDenom)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.AirdropList) > 0 {
		for _, e := range m.AirdropList {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AirdropProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: AirdropProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AirdropProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationOfAirdrop", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationOfAirdrop, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AirdropList = append(m.AirdropList, Claimable{})
			if err := m.AirdropList[len(m.AirdropList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
