// Code generated by protoc-gen-gogo.
// source: Profile.proto
// DO NOT EDIT!

package proto

import testing4 "testing"
import math_rand4 "math/rand"
import time4 "time"
import code_google_com_p_gogoprotobuf_proto3 "code.google.com/p/gogoprotobuf/proto"
import testing5 "testing"
import math_rand5 "math/rand"
import time5 "time"
import encoding_json1 "encoding/json"
import testing6 "testing"
import math_rand6 "math/rand"
import time6 "time"
import code_google_com_p_gogoprotobuf_proto4 "code.google.com/p/gogoprotobuf/proto"
import math_rand7 "math/rand"
import time7 "time"
import testing7 "testing"
import code_google_com_p_gogoprotobuf_proto5 "code.google.com/p/gogoprotobuf/proto"

func TestProfileProto(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedProfile(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Profile{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestProfileMarshalTo(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedProfile(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &Profile{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkProfileProtoMarshal(b *testing4.B) {
	popr := math_rand4.New(math_rand4.NewSource(616))
	total := 0
	pops := make([]*Profile, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedProfile(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkProfileProtoUnmarshal(b *testing4.B) {
	popr := math_rand4.New(math_rand4.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(NewPopulatedProfile(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &Profile{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestProfileJSON(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedProfile(popr, true)
	jsondata, err := encoding_json1.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Profile{}
	err = encoding_json1.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestProfileProtoText(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedProfile(popr, true)
	data := code_google_com_p_gogoprotobuf_proto4.MarshalTextString(p)
	msg := &Profile{}
	if err := code_google_com_p_gogoprotobuf_proto4.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestProfileProtoCompactText(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedProfile(popr, true)
	data := code_google_com_p_gogoprotobuf_proto4.CompactTextString(p)
	msg := &Profile{}
	if err := code_google_com_p_gogoprotobuf_proto4.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestProfileSize(t *testing7.T) {
	popr := math_rand7.New(math_rand7.NewSource(time7.Now().UnixNano()))
	p := NewPopulatedProfile(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto5.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto5.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto5.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkProfileSize(b *testing7.B) {
	popr := math_rand7.New(math_rand7.NewSource(616))
	total := 0
	pops := make([]*Profile, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedProfile(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

//These tests are generated by code.google.com/p/gogoprotobuf/plugin/testgen
