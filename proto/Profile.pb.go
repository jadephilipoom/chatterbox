// Code generated by protoc-gen-gogo.
// source: Profile.proto
// DO NOT EDIT!

package proto

import proto1 "code.google.com/p/gogoprotobuf/proto"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import io1 "io"
import fmt1 "fmt"
import code_google_com_p_gogoprotobuf_proto1 "code.google.com/p/gogoprotobuf/proto"

import bytes1 "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

type Profile struct {
	ServerAddressTCP string `protobuf:"bytes,1,req" json:"ServerAddressTCP"`
	// allowing the port to be specified might open a DOS vector, so not now
	ServerTransportPK Byte32 `protobuf:"bytes,2,req,customtype=Byte32" json:"ServerTransportPK"`
	UserIDAtServer    Byte32 `protobuf:"bytes,3,req,customtype=Byte32" json:"UserIDAtServer"`
	KeySigningKey     Byte32 `protobuf:"bytes,4,req,customtype=Byte32" json:"KeySigningKey"`
	MessageAuthKey    Byte32 `protobuf:"bytes,5,req,customtype=Byte32" json:"MessageAuthKey"`
	XXX_unrecognized  []byte `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto1.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}

func init() {
}
func (m *Profile) Unmarshal(data []byte) error {
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
				return fmt1.Errorf("proto: wrong wireType = %d for field ServerAddressTCP", wireType)
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
			m.ServerAddressTCP = string(data[index:postIndex])
			index = postIndex
		case 2:
			if wireType != 2 {
				return fmt1.Errorf("proto: wrong wireType = %d for field ServerTransportPK", wireType)
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
			if err := m.ServerTransportPK.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 3:
			if wireType != 2 {
				return fmt1.Errorf("proto: wrong wireType = %d for field UserIDAtServer", wireType)
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
			if err := m.UserIDAtServer.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 4:
			if wireType != 2 {
				return fmt1.Errorf("proto: wrong wireType = %d for field KeySigningKey", wireType)
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
			if err := m.KeySigningKey.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 5:
			if wireType != 2 {
				return fmt1.Errorf("proto: wrong wireType = %d for field MessageAuthKey", wireType)
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
			if err := m.MessageAuthKey.Unmarshal(data[index:postIndex]); err != nil {
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
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
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
func (m *Profile) Size() (n int) {
	var l int
	_ = l
	l = len(m.ServerAddressTCP)
	n += 1 + l + sovProfile(uint64(l))
	l = m.ServerTransportPK.Size()
	n += 1 + l + sovProfile(uint64(l))
	l = m.UserIDAtServer.Size()
	n += 1 + l + sovProfile(uint64(l))
	l = m.KeySigningKey.Size()
	n += 1 + l + sovProfile(uint64(l))
	l = m.MessageAuthKey.Size()
	n += 1 + l + sovProfile(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProfile(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProfile(x uint64) (n int) {
	return sovProfile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedProfile(r randyProfile, easy bool) *Profile {
	this := &Profile{}
	this.ServerAddressTCP = randStringProfile(r)
	v1 := NewPopulatedByte32(r)
	this.ServerTransportPK = *v1
	v2 := NewPopulatedByte32(r)
	this.UserIDAtServer = *v2
	v3 := NewPopulatedByte32(r)
	this.KeySigningKey = *v3
	v4 := NewPopulatedByte32(r)
	this.MessageAuthKey = *v4
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedProfile(r, 6)
	}
	return this
}

type randyProfile interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneProfile(r randyProfile) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringProfile(r randyProfile) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneProfile(r)
	}
	return string(tmps)
}
func randUnrecognizedProfile(r randyProfile, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldProfile(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldProfile(data []byte, r randyProfile, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateProfile(data, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		data = encodeVarintPopulateProfile(data, uint64(v6))
	case 1:
		data = encodeVarintPopulateProfile(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateProfile(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateProfile(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateProfile(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateProfile(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *Profile) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Profile) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintProfile(data, i, uint64(len(m.ServerAddressTCP)))
	i += copy(data[i:], m.ServerAddressTCP)
	data[i] = 0x12
	i++
	i = encodeVarintProfile(data, i, uint64(m.ServerTransportPK.Size()))
	n1, err := m.ServerTransportPK.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x1a
	i++
	i = encodeVarintProfile(data, i, uint64(m.UserIDAtServer.Size()))
	n2, err := m.UserIDAtServer.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x22
	i++
	i = encodeVarintProfile(data, i, uint64(m.KeySigningKey.Size()))
	n3, err := m.KeySigningKey.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	data[i] = 0x2a
	i++
	i = encodeVarintProfile(data, i, uint64(m.MessageAuthKey.Size()))
	n4, err := m.MessageAuthKey.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64Profile(data []byte, offset int, v uint64) int {
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
func encodeFixed32Profile(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintProfile(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *Profile) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Profile)
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
	if this.ServerAddressTCP != that1.ServerAddressTCP {
		return false
	}
	if !this.ServerTransportPK.Equal(that1.ServerTransportPK) {
		return false
	}
	if !this.UserIDAtServer.Equal(that1.UserIDAtServer) {
		return false
	}
	if !this.KeySigningKey.Equal(that1.KeySigningKey) {
		return false
	}
	if !this.MessageAuthKey.Equal(that1.MessageAuthKey) {
		return false
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
