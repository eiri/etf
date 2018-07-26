// Copyright (c) 2018 Eric Avdey

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//   http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package etf

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

type any interface{}

type decode func(*Decoder) (any, error)

type decoderTest struct {
	name   string
	decode decode
	expect any
}

var decodeInt = func(d *Decoder) (any, error) {
	var v int
	err := d.Decode(&v)
	return v, err
}

var decodeFloat64 = func(d *Decoder) (any, error) {
	var v float64
	err := d.Decode(&v)
	return v, err
}

var decodeString = func(d *Decoder) (any, error) {
	var v string
	err := d.Decode(&v)
	return v, err
}

var decodeBool = func(d *Decoder) (any, error) {
	var v bool
	err := d.Decode(&v)
	return v, err
}

var decodeBytes = func(d *Decoder) (any, error) {
	var v []byte
	err := d.Decode(&v)
	return v, err
}

var tests = []decoderTest{
	{"uint8", decodeInt, 42},
	{"int32", decodeInt, 523124044},
	{"negint32", decodeInt, -42},
	{"float", decodeFloat64, 3.14159},
	{"atom", decodeString, "cat"},
	{"atomutf8", decodeString, strings.Repeat("😀", 64)},
	{"smallatomutf8", decodeString, "猫"},
	{"booltrue", decodeBool, true},
	{"boolfalse", decodeBool, false},
	{"binary", decodeBytes, []byte{23, 198, 181, 53, 145, 254, 7}},
}

// TestDecode to make sure we cover all types decoding
func TestDecode(t *testing.T) {
	for _, tt := range tests {
		path := filepath.Join("testdata", tt.name+".golden")
		r, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		d := NewDecoder(r)
		v, err := tt.decode(d)
		r.Close()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(v, tt.expect) {
			t.Fatalf("Expecting %#v, got %#v", tt.expect, v)
		}
		if testing.Verbose() {
			t.Logf("%-16s - ok", tt.name)
		}
	}
}
