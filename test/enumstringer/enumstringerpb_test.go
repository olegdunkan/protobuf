// Code generated by protoc-gen-gogo.
// source: enumstringer.proto
// DO NOT EDIT!

/*
Package enumstringer is a generated protocol buffer package.

It is generated from these files:
	enumstringer.proto

It has these top-level messages:
	NidOptEnum
	NinOptEnum
	NidRepEnum
	NinRepEnum
*/
package enumstringer

import testing "testing"
import math_rand "math/rand"
import time "time"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import encoding_json "encoding/json"

func TestNidOptEnumProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNidOptEnum(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &NidOptEnum{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNinOptEnumProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNinOptEnum(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &NinOptEnum{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNidRepEnumProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNidRepEnum(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &NidRepEnum{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNinRepEnumProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNinRepEnum(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &NinRepEnum{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNidOptEnumJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNidOptEnum(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		if _, ok := err.(*encoding_json.UnsupportedTypeError); ok {
			t.Skip(err)
		}
		t.Fatal(err)
	}
	msg := &NidOptEnum{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestNinOptEnumJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNinOptEnum(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		if _, ok := err.(*encoding_json.UnsupportedTypeError); ok {
			t.Skip(err)
		}
		t.Fatal(err)
	}
	msg := &NinOptEnum{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestNidRepEnumJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNidRepEnum(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		if _, ok := err.(*encoding_json.UnsupportedTypeError); ok {
			t.Skip(err)
		}
		t.Fatal(err)
	}
	msg := &NidRepEnum{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestNinRepEnumJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNinRepEnum(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		if _, ok := err.(*encoding_json.UnsupportedTypeError); ok {
			t.Skip(err)
		}
		t.Fatal(err)
	}
	msg := &NinRepEnum{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestNidOptEnumProtoText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNidOptEnum(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &NidOptEnum{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNidOptEnumProtoCompactText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNidOptEnum(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &NidOptEnum{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNinOptEnumProtoText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNinOptEnum(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &NinOptEnum{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNinOptEnumProtoCompactText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNinOptEnum(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &NinOptEnum{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNidRepEnumProtoText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNidRepEnum(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &NidRepEnum{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNidRepEnumProtoCompactText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNidRepEnum(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &NidRepEnum{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNinRepEnumProtoText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNinRepEnum(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &NinRepEnum{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNinRepEnumProtoCompactText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNinRepEnum(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &NinRepEnum{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNidOptEnumVerboseEqual(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNidOptEnum(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &NidOptEnum{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestNinOptEnumVerboseEqual(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNinOptEnum(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &NinOptEnum{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestNidRepEnumVerboseEqual(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNidRepEnum(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &NidRepEnum{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestNinRepEnumVerboseEqual(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNinRepEnum(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &NinRepEnum{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
