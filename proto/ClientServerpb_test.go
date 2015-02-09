// Code generated by protoc-gen-gogo.
// source: ClientServer.proto
// DO NOT EDIT!

package proto

import testing "testing"
import math_rand "math/rand"
import time "time"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import encoding_json "encoding/json"

func TestServerToClientProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedServerToClient(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &ServerToClient{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestServerToClientMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedServerToClient(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &ServerToClient{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkServerToClientProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ServerToClient, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedServerToClient(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_gogo_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkServerToClientProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_gogo_protobuf_proto.Marshal(NewPopulatedServerToClient(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &ServerToClient{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_gogo_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestClientToServerProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &ClientToServer{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestClientToServerMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &ClientToServer{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkClientToServerProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ClientToServer, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedClientToServer(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_gogo_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkClientToServerProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_gogo_protobuf_proto.Marshal(NewPopulatedClientToServer(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &ClientToServer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_gogo_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestClientToServer_DeliverEnvelopeProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer_DeliverEnvelope(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &ClientToServer_DeliverEnvelope{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestClientToServer_DeliverEnvelopeMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer_DeliverEnvelope(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &ClientToServer_DeliverEnvelope{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkClientToServer_DeliverEnvelopeProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ClientToServer_DeliverEnvelope, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedClientToServer_DeliverEnvelope(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_gogo_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkClientToServer_DeliverEnvelopeProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_gogo_protobuf_proto.Marshal(NewPopulatedClientToServer_DeliverEnvelope(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &ClientToServer_DeliverEnvelope{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_gogo_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestServerToClientJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedServerToClient(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &ServerToClient{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestClientToServerJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &ClientToServer{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestClientToServer_DeliverEnvelopeJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer_DeliverEnvelope(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &ClientToServer_DeliverEnvelope{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestServerToClientProtoText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedServerToClient(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &ServerToClient{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestServerToClientProtoCompactText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedServerToClient(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &ServerToClient{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestClientToServerProtoText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &ClientToServer{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestClientToServerProtoCompactText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &ClientToServer{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestClientToServer_DeliverEnvelopeProtoText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer_DeliverEnvelope(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &ClientToServer_DeliverEnvelope{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestClientToServer_DeliverEnvelopeProtoCompactText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer_DeliverEnvelope(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &ClientToServer_DeliverEnvelope{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestServerToClientSize(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedServerToClient(popr, true)
	size2 := github_com_gogo_protobuf_proto.Size(p)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
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
	size3 := github_com_gogo_protobuf_proto.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkServerToClientSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ServerToClient, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedServerToClient(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestClientToServerSize(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer(popr, true)
	size2 := github_com_gogo_protobuf_proto.Size(p)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
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
	size3 := github_com_gogo_protobuf_proto.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkClientToServerSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ClientToServer, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedClientToServer(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestClientToServer_DeliverEnvelopeSize(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedClientToServer_DeliverEnvelope(popr, true)
	size2 := github_com_gogo_protobuf_proto.Size(p)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
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
	size3 := github_com_gogo_protobuf_proto.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkClientToServer_DeliverEnvelopeSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ClientToServer_DeliverEnvelope, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedClientToServer_DeliverEnvelope(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
