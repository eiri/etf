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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type any interface{}

type decoderTest func(*Decoder) (any, error)

func testDecode(fn string, dt decoderTest, expect any) error {
	path := filepath.Join("testdata", fn+".bin")
	r, err := os.Open(path)
	if err != nil {
		return err
	}
	defer r.Close()
	d := NewDecoder(r)
	v, err := dt(d)
	if err == nil && v != expect {
		err = fmt.Errorf("Expecting %#v, got %#v", expect, v)
	}
	return err
}

// TestDecodeSmallInt to make sure we can decode uint8
func TestDecodeSmallInt(t *testing.T) {
	dt := func(d *Decoder) (any, error) {
		var v int
		err := d.Decode(&v)
		return v, err
	}
	err := testDecode("uint8", dt, 42)
	if err != nil {
		t.Fatal(err)
	}
}

// TestDecodeInt to make sure we can decode int32
func TestDecodeInt(t *testing.T) {
	dt := func(d *Decoder) (any, error) {
		var v int
		err := d.Decode(&v)
		return v, err
	}
	err := testDecode("int32", dt, 523124044)
	if err != nil {
		t.Fatal(err)
	}
}

// TestDecodeNegInt to make sure we can decode negative int32
func TestDecodeNegInt(t *testing.T) {
	dt := func(d *Decoder) (any, error) {
		var v int
		err := d.Decode(&v)
		return v, err
	}
	err := testDecode("negint32", dt, -42)
	if err != nil {
		t.Fatal(err)
	}
}

// TestDecodeFloat64 to make sure we can decode float
func TestDecodeFloat64(t *testing.T) {
	dt := func(d *Decoder) (any, error) {
		var v float64
		err := d.Decode(&v)
		return v, err
	}
	err := testDecode("float", dt, 3.14159)
	if err != nil {
		t.Fatal(err)
	}
}

// TestDecodeAtom to make sure we can decode atom
func TestDecodeAtom(t *testing.T) {
	dt := func(d *Decoder) (any, error) {
		var v string
		err := d.Decode(&v)
		return v, err
	}
	err := testDecode("atom", dt, "cat")
	if err != nil {
		t.Fatal(err)
	}
}

// TestDecodeUTF8Atom to make sure we can decode UTF8 encoded atom
func TestDecodeUTF8Atom(t *testing.T) {
	dt := func(d *Decoder) (any, error) {
		var v string
		err := d.Decode(&v)
		return v, err
	}
	expect := strings.Repeat("😀", 64)
	err := testDecode("atomutf8", dt, expect)
	if err != nil {
		t.Fatal(err)
	}
}

// TestDecodeSmallUTF8Atom to make sure we can decode UTF8 encoded atom
func TestDecodeSmallUTF8Atom(t *testing.T) {
	dt := func(d *Decoder) (any, error) {
		var v string
		err := d.Decode(&v)
		return v, err
	}
	err := testDecode("smallatomutf8", dt, "猫")
	if err != nil {
		t.Fatal(err)
	}
}

// TestDecodeBoolTrue to make sure we can decode boolean true
func TestDecodeBoolTrue(t *testing.T) {
	dt := func(d *Decoder) (any, error) {
		var v bool
		err := d.Decode(&v)
		return v, err
	}
	err := testDecode("booltrue", dt, true)
	if err != nil {
		t.Fatal(err)
	}
}

// TestDecodeBoolFalse to make sure we can decode boolean false
func TestDecodeBoolFalse(t *testing.T) {
	dt := func(d *Decoder) (any, error) {
		var v bool
		err := d.Decode(&v)
		return v, err
	}
	err := testDecode("boolfalse", dt, false)
	if err != nil {
		t.Fatal(err)
	}
}
