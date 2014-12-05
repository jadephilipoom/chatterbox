// Code generated by protoc-gen-gogo.
// source: ConversationMetadata.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import github_com_andres_erbsen_dename_protocol "github.com/andres-erbsen/dename/protocol"

import io1 "io"
import fmt1 "fmt"
import github_com_gogo_protobuf_proto1 "github.com/gogo/protobuf/proto"

import bytes1 "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

type ConversationMetadata struct {
	Participants        []string                                         `protobuf:"bytes,1,rep,name=participants" json:"participants"`
	Subject             string                                           `protobuf:"bytes,2,req,name=subject" json:"subject"`
	SenderDenameProfile github_com_andres_erbsen_dename_protocol.Profile `protobuf:"bytes,3,req,name=sender_dename_profile,customtype=github.com/andres-erbsen/dename/protocol.Profile" json:"sender_dename_profile"`
	XXX_unrecognized    []byte                                           `json:"-"`
}

func (m *ConversationMetadata) Reset()         { *m = ConversationMetadata{} }
func (m *ConversationMetadata) String() string { return proto1.CompactTextString(m) }
func (*ConversationMetadata) ProtoMessage()    {}

func init() {
}
func (m *ConversationMetadata) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt1.Errorf("proto: wrong wireType = %d for field Participants", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			m.Participants = append(m.Participants, string(data[index:postIndex]))
			index = postIndex
		case 2:
			if wireType != 2 {
				return fmt1.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			m.Subject = string(data[index:postIndex])
			index = postIndex
		case 3:
			if wireType != 2 {
				return fmt1.Errorf("proto: wrong wireType = %d for field SenderDenameProfile", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			if err := m.SenderDenameProfile.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *ConversationMetadata) Size() (n int) {
	var l int
	_ = l
	if len(m.Participants) > 0 {
		for _, s := range m.Participants {
			l = len(s)
			n += 1 + l + sovConversationMetadata(uint64(l))
		}
	}
	l = len(m.Subject)
	n += 1 + l + sovConversationMetadata(uint64(l))
	l = m.SenderDenameProfile.Size()
	n += 1 + l + sovConversationMetadata(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovConversationMetadata(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConversationMetadata(x uint64) (n int) {
	return sovConversationMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConversationMetadata) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ConversationMetadata) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Participants) > 0 {
		for _, s := range m.Participants {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	data[i] = 0x12
	i++
	i = encodeVarintConversationMetadata(data, i, uint64(len(m.Subject)))
	i += copy(data[i:], m.Subject)
	data[i] = 0x1a
	i++
	i = encodeVarintConversationMetadata(data, i, uint64(m.SenderDenameProfile.Size()))
	n1, err := m.SenderDenameProfile.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64ConversationMetadata(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ConversationMetadata(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConversationMetadata(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *ConversationMetadata) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ConversationMetadata)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Participants) != len(that1.Participants) {
		return false
	}
	for i := range this.Participants {
		if this.Participants[i] != that1.Participants[i] {
			return false
		}
	}
	if this.Subject != that1.Subject {
		return false
	}
	if !this.SenderDenameProfile.Equal(that1.SenderDenameProfile) {
		return false
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
