// Code generated by protoc-gen-gogo.
// source: LocalAccountConfig.proto
// DO NOT EDIT!

package proto

import testing8 "testing"
import math_rand8 "math/rand"
import time8 "time"
import github_com_gogo_protobuf_proto6 "github.com/gogo/protobuf/proto"
import testing9 "testing"
import math_rand9 "math/rand"
import time9 "time"
import encoding_json2 "encoding/json"
import testing10 "testing"
import math_rand10 "math/rand"
import time10 "time"
import github_com_gogo_protobuf_proto7 "github.com/gogo/protobuf/proto"
import math_rand11 "math/rand"
import time11 "time"
import testing11 "testing"
import github_com_gogo_protobuf_proto8 "github.com/gogo/protobuf/proto"

func TestLocalAccountConfigProto(t *testing8.T) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, false)
	data, err := github_com_gogo_protobuf_proto6.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &LocalAccountConfig{}
	if err := github_com_gogo_protobuf_proto6.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestLocalAccountConfigMarshalTo(t *testing8.T) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &LocalAccountConfig{}
	if err := github_com_gogo_protobuf_proto6.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkLocalAccountConfigProtoMarshal(b *testing8.B) {
	popr := math_rand8.New(math_rand8.NewSource(616))
	total := 0
	pops := make([]*LocalAccountConfig, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedLocalAccountConfig(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_gogo_protobuf_proto6.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkLocalAccountConfigProtoUnmarshal(b *testing8.B) {
	popr := math_rand8.New(math_rand8.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_gogo_protobuf_proto6.Marshal(NewPopulatedLocalAccountConfig(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &LocalAccountConfig{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_gogo_protobuf_proto6.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestLocalAccountConfigJSON(t *testing9.T) {
	popr := math_rand9.New(math_rand9.NewSource(time9.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, true)
	jsondata, err := encoding_json2.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &LocalAccountConfig{}
	err = encoding_json2.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestLocalAccountConfigProtoText(t *testing10.T) {
	popr := math_rand10.New(math_rand10.NewSource(time10.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, true)
	data := github_com_gogo_protobuf_proto7.MarshalTextString(p)
	msg := &LocalAccountConfig{}
	if err := github_com_gogo_protobuf_proto7.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestLocalAccountConfigProtoCompactText(t *testing10.T) {
	popr := math_rand10.New(math_rand10.NewSource(time10.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, true)
	data := github_com_gogo_protobuf_proto7.CompactTextString(p)
	msg := &LocalAccountConfig{}
	if err := github_com_gogo_protobuf_proto7.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestLocalAccountConfigSize(t *testing11.T) {
	popr := math_rand11.New(math_rand11.NewSource(time11.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, true)
	size2 := github_com_gogo_protobuf_proto8.Size(p)
	data, err := github_com_gogo_protobuf_proto8.Marshal(p)
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
	size3 := github_com_gogo_protobuf_proto8.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkLocalAccountConfigSize(b *testing11.B) {
	popr := math_rand11.New(math_rand11.NewSource(616))
	total := 0
	pops := make([]*LocalAccountConfig, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedLocalAccountConfig(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
