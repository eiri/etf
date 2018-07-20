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
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"reflect"
)

// Decoder reads and decodes serialized ETF from a given input stream
type Decoder struct {
	io.Reader
}

// NewDecoder returns a new decoder that read from given reader
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r}
}

// Read because or decoder is io.Reader itself
func (d *Decoder) Read(p []byte) (n int, err error) {
	n, err = d.Reader.Read(p)
	return
}

// Decode reads from input and stores value in value or v pointer
func (d *Decoder) Decode(v interface{}) error {
	// don't bother if we can't set value
	rv := reflect.Indirect(reflect.ValueOf(v))
	if !rv.CanAddr() {
		return errors.New("not a pointer")
	}
	// grep header and tag
	var probe [2]byte
	if _, err := io.ReadFull(d, probe[:]); err != nil {
		return err
	}
	if probe[0] != 131 {
		return errors.New("bad argument")
	}
	// read based on tag
	switch probe[1] {
	case 97:
		err := d.decodeToSmallInt(rv)
		return err
	case 98:
		err := d.decodeToInt(rv)
		return err
	case 70:
		err := d.decodeToFloat64(rv)
		return err
	default:
		return errors.New("unknown tag")
	}
}

func (d *Decoder) decodeToSmallInt(rv reflect.Value) error {
	if rv.Type().Kind() != reflect.Int {
		return errors.New("invalid type")
	}
	var buf [1]byte
	if _, err := io.ReadFull(d, buf[:]); err != nil {
		return err
	}
	val := int64(buf[0])
	rv.SetInt(val)
	return nil
}

func (d *Decoder) decodeToInt(rv reflect.Value) error {
	if rv.Type().Kind() != reflect.Int {
		return errors.New("invalid type")
	}
	var buf [4]byte
	if _, err := io.ReadFull(d, buf[:]); err != nil {
		return err
	}
	var i int32
	bbuf := bytes.NewReader(buf[:])
	err := binary.Read(bbuf, binary.BigEndian, &i)
	if err != nil {
		return err
	}
	val := int64(i)
	rv.SetInt(val)
	return nil
}

func (d *Decoder) decodeToFloat64(rv reflect.Value) error {
	if rv.Type().Kind() != reflect.Float64 {
		return errors.New("invalid type")
	}
	var buf [8]byte
	if _, err := io.ReadFull(d, buf[:]); err != nil {
		return err
	}
	var f float64
	bbuf := bytes.NewReader(buf[:])
	err := binary.Read(bbuf, binary.BigEndian, &f)
	if err != nil {
		return err
	}
	rv.SetFloat(f)
	return nil
}
