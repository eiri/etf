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
	"strings"
	"testing"
)

// TestDecodeSmallInt to make sure we can decode uint8
func TestDecodeSmallInt(t *testing.T) {
	r, err := os.Open("testdata/uint8.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	var v int
	err = NewDecoder(r).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	if v != 42 {
		t.Fatalf("Expecting 42, got %v", v)
	}
}

// TestDecodeInt to make sure we can decode int32
func TestDecodeInt(t *testing.T) {
	r, err := os.Open("testdata/int32.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	var v int
	err = NewDecoder(r).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	if v != 523124044 {
		t.Fatalf("Expecting 523124044, got %v", v)
	}
}

// TestDecodeNegInt to make sure we can decode negative int32
func TestDecodeNegInt(t *testing.T) {
	r, err := os.Open("testdata/negint32.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	var v int
	err = NewDecoder(r).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	if v != -42 {
		t.Fatalf("Expecting -42, got %v", v)
	}
}

// TestDecodeFloat64 to make sure we can decode float
func TestDecodeFloat64(t *testing.T) {
	r, err := os.Open("testdata/float.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	var v float64
	err = NewDecoder(r).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	if v != 3.14159 {
		t.Fatalf("Expecting 3.14159, got %v", v)
	}
}

// TestDecodeAtom to make sure we can decode atom
func TestDecodeAtom(t *testing.T) {
	r, err := os.Open("testdata/atom.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	var v string
	err = NewDecoder(r).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	if v != "cat" {
		t.Fatalf("Expecting \"cat\", got %v", v)
	}
}

// TestDecodeUTF8Atom to make sure we can decode UTF8 encoded atom
func TestDecodeUTF8Atom(t *testing.T) {
	r, err := os.Open("testdata/atomutf8.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	var v string
	err = NewDecoder(r).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	expect := strings.Repeat("😀", 64)
	if v != expect {
		t.Fatalf("Expecting \"%s\", got %v", expect, v)
	}
}

// TestDecodeSmallUTF8Atom to make sure we can decode UTF8 encoded atom
func TestDecodeSmallUTF8Atom(t *testing.T) {
	r, err := os.Open("testdata/smallatomutf8.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	var v string
	err = NewDecoder(r).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	if v != "猫" {
		t.Fatalf("Expecting \"猫\", got %v", v)
	}
}

// TestDecodeBoolTrue to make sure we can decode boolean true
func TestDecodeBoolTrue(t *testing.T) {
	r, err := os.Open("testdata/booltrue.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	var v bool
	err = NewDecoder(r).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	if !v {
		t.Fatalf("Expecting true, got %v", v)
	}
}

// TestDecodeBoolFalse to make sure we can decode boolean false
func TestDecodeBoolFalse(t *testing.T) {
	r, err := os.Open("testdata/boolfalse.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	var v bool
	err = NewDecoder(r).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	if v {
		t.Fatalf("Expecting false, got %v", v)
	}
}
