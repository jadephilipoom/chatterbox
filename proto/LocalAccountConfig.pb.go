// Code generated by protoc-gen-gogo.
// source: LocalAccountConfig.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io3 "io"
import fmt3 "fmt"
import github_com_gogo_protobuf_proto3 "github.com/gogo/protobuf/proto"

import bytes3 "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

type LocalAccountConfig struct {
	ServerAddressTCP            string `protobuf:"bytes,1,req" json:"ServerAddressTCP"`
	ServerPortTCP               int32  `protobuf:"varint,2,opt" json:"ServerPortTCP"`
	ServerTransportPK           Byte32 `protobuf:"bytes,3,req,customtype=Byte32" json:"ServerTransportPK"`
	TransportSecretKeyForServer Byte32 `protobuf:"bytes,4,req,customtype=Byte32" json:"TransportSecretKeyForServer"`
	KeySigningSecretKey         []byte `protobuf:"bytes,5,req" json:"KeySigningSecretKey"`
	MessageAuthSecretKey        Byte32 `protobuf:"bytes,6,req,customtype=Byte32" json:"MessageAuthSecretKey"`
	Dename                      string `protobuf:"bytes,7,req" json:"Dename"`
	XXX_unrecognized            []byte `json:"-"`
}

func (m *LocalAccountConfig) Reset()         { *m = LocalAccountConfig{} }
func (m *LocalAccountConfig) String() string { return proto1.CompactTextString(m) }
func (*LocalAccountConfig) ProtoMessage()    {}

func init() {
}
func (m *LocalAccountConfig) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io3.ErrUnexpectedEOF
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
				return fmt3.Errorf("proto: wrong wireType = %d for field ServerAddressTCP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io3.ErrUnexpectedEOF
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
				return io3.ErrUnexpectedEOF
			}
			m.ServerAddressTCP = string(data[index:postIndex])
			index = postIndex
		case 2:
			if wireType != 0 {
				return fmt3.Errorf("proto: wrong wireType = %d for field ServerPortTCP", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io3.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.ServerPortTCP |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt3.Errorf("proto: wrong wireType = %d for field ServerTransportPK", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io3.ErrUnexpectedEOF
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
				return io3.ErrUnexpectedEOF
			}
			if err := m.ServerTransportPK.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 4:
			if wireType != 2 {
				return fmt3.Errorf("proto: wrong wireType = %d for field TransportSecretKeyForServer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io3.ErrUnexpectedEOF
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
				return io3.ErrUnexpectedEOF
			}
			if err := m.TransportSecretKeyForServer.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 5:
			if wireType != 2 {
				return fmt3.Errorf("proto: wrong wireType = %d for field KeySigningSecretKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io3.ErrUnexpectedEOF
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
				return io3.ErrUnexpectedEOF
			}
			m.KeySigningSecretKey = append(m.KeySigningSecretKey, data[index:postIndex]...)
			index = postIndex
		case 6:
			if wireType != 2 {
				return fmt3.Errorf("proto: wrong wireType = %d for field MessageAuthSecretKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io3.ErrUnexpectedEOF
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
				return io3.ErrUnexpectedEOF
			}
			if err := m.MessageAuthSecretKey.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 7:
			if wireType != 2 {
				return fmt3.Errorf("proto: wrong wireType = %d for field Dename", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io3.ErrUnexpectedEOF
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
				return io3.ErrUnexpectedEOF
			}
			m.Dename = string(data[index:postIndex])
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
			skippy, err := github_com_gogo_protobuf_proto3.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io3.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *LocalAccountConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.ServerAddressTCP)
	n += 1 + l + sovLocalAccountConfig(uint64(l))
	n += 1 + sovLocalAccountConfig(uint64(uint32(m.ServerPortTCP)))
	l = m.ServerTransportPK.Size()
	n += 1 + l + sovLocalAccountConfig(uint64(l))
	l = m.TransportSecretKeyForServer.Size()
	n += 1 + l + sovLocalAccountConfig(uint64(l))
	l = len(m.KeySigningSecretKey)
	n += 1 + l + sovLocalAccountConfig(uint64(l))
	l = m.MessageAuthSecretKey.Size()
	n += 1 + l + sovLocalAccountConfig(uint64(l))
	l = len(m.Dename)
	n += 1 + l + sovLocalAccountConfig(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLocalAccountConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLocalAccountConfig(x uint64) (n int) {
	return sovLocalAccountConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedLocalAccountConfig(r randyLocalAccountConfig, easy bool) *LocalAccountConfig {
	this := &LocalAccountConfig{}
	this.ServerAddressTCP = randStringLocalAccountConfig(r)
	this.ServerPortTCP = r.Int31()
	if r.Intn(2) == 0 {
		this.ServerPortTCP *= -1
	}
	v1 := NewPopulatedByte32(r)
	this.ServerTransportPK = *v1
	v2 := NewPopulatedByte32(r)
	this.TransportSecretKeyForServer = *v2
	v3 := r.Intn(100)
	this.KeySigningSecretKey = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.KeySigningSecretKey[i] = byte(r.Intn(256))
	}
	v4 := NewPopulatedByte32(r)
	this.MessageAuthSecretKey = *v4
	this.Dename = randStringLocalAccountConfig(r)
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedLocalAccountConfig(r, 8)
	}
	return this
}

type randyLocalAccountConfig interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneLocalAccountConfig(r randyLocalAccountConfig) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringLocalAccountConfig(r randyLocalAccountConfig) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneLocalAccountConfig(r)
	}
	return string(tmps)
}
func randUnrecognizedLocalAccountConfig(r randyLocalAccountConfig, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldLocalAccountConfig(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldLocalAccountConfig(data []byte, r randyLocalAccountConfig, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateLocalAccountConfig(data, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		data = encodeVarintPopulateLocalAccountConfig(data, uint64(v6))
	case 1:
		data = encodeVarintPopulateLocalAccountConfig(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateLocalAccountConfig(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateLocalAccountConfig(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateLocalAccountConfig(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateLocalAccountConfig(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *LocalAccountConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *LocalAccountConfig) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintLocalAccountConfig(data, i, uint64(len(m.ServerAddressTCP)))
	i += copy(data[i:], m.ServerAddressTCP)
	data[i] = 0x10
	i++
	i = encodeVarintLocalAccountConfig(data, i, uint64(uint32(m.ServerPortTCP)))
	data[i] = 0x1a
	i++
	i = encodeVarintLocalAccountConfig(data, i, uint64(m.ServerTransportPK.Size()))
	n1, err := m.ServerTransportPK.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x22
	i++
	i = encodeVarintLocalAccountConfig(data, i, uint64(m.TransportSecretKeyForServer.Size()))
	n2, err := m.TransportSecretKeyForServer.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x2a
	i++
	i = encodeVarintLocalAccountConfig(data, i, uint64(len(m.KeySigningSecretKey)))
	i += copy(data[i:], m.KeySigningSecretKey)
	data[i] = 0x32
	i++
	i = encodeVarintLocalAccountConfig(data, i, uint64(m.MessageAuthSecretKey.Size()))
	n3, err := m.MessageAuthSecretKey.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	data[i] = 0x3a
	i++
	i = encodeVarintLocalAccountConfig(data, i, uint64(len(m.Dename)))
	i += copy(data[i:], m.Dename)
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64LocalAccountConfig(data []byte, offset int, v uint64) int {
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
func encodeFixed32LocalAccountConfig(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLocalAccountConfig(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *LocalAccountConfig) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*LocalAccountConfig)
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
	if this.ServerPortTCP != that1.ServerPortTCP {
		return false
	}
	if !this.ServerTransportPK.Equal(that1.ServerTransportPK) {
		return false
	}
	if !this.TransportSecretKeyForServer.Equal(that1.TransportSecretKeyForServer) {
		return false
	}
	if !bytes3.Equal(this.KeySigningSecretKey, that1.KeySigningSecretKey) {
		return false
	}
	if !this.MessageAuthSecretKey.Equal(that1.MessageAuthSecretKey) {
		return false
	}
	if this.Dename != that1.Dename {
		return false
	}
	if !bytes3.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
