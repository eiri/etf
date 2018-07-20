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
	"testing"
)

// TestDecodeSmallInt to make sure we can decode into int8
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

// TestDecodeInt to make sure we can decode into int32
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

// TestDecodeNegInt to make sure we can decode into negative int32
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

// TestFloat64 to make sure we can decode into float
func TestFloat64(t *testing.T) {
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
