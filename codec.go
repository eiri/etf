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

type codec interface {
	Decode(*Decoder, reflect.Value) error
}

func getCodecForTag(tag byte) (c codec, err error) {
	switch tag {
	case 97, 98:
		c = new(intCodec)
	case 100, 118:
		c = new(atomCodec)
	case 119:
		c = new(smallAtomCodec)
	case 70:
		c = new(floatCodec)
	case 109:
		c = new(binaryCodec)
	default:
		err = errors.New("unknown tag")
	}
	return
}

// Codec to int
type intCodec struct{}

func (c *intCodec) Decode(d *Decoder, rv reflect.Value) error {
	if rv.Type().Kind() != reflect.Int {
		return errors.New("invalid type")
	}
	buf := make([]byte, 4)
	n, err := d.Read(buf)
	if err != nil {
		return err
	}
	var val int64
	switch n {
	case 1:
		val = int64(buf[0])
	case 4:
		var i int32
		bbuf := bytes.NewReader(buf)
		err := binary.Read(bbuf, binary.BigEndian, &i)
		if err != nil {
			return err
		}
		val = int64(i)
	default:
		return errors.New("bad argument")
	}
	rv.SetInt(val)
	return nil
}

// Codec to float64
type floatCodec struct{}

func (c *floatCodec) Decode(d *Decoder, rv reflect.Value) error {
	if rv.Type().Kind() != reflect.Float64 {
		return errors.New("invalid type")
	}
	buf := make([]byte, 8)
	if _, err := io.ReadFull(d, buf); err != nil {
		return err
	}
	var f float64
	bbuf := bytes.NewReader(buf)
	err := binary.Read(bbuf, binary.BigEndian, &f)
	if err != nil {
		return err
	}
	rv.SetFloat(f)
	return nil
}

// Codec to atom
type atomCodec struct{}

func (c *atomCodec) Decode(d *Decoder, rv reflect.Value) error {
	if rv.Type().Kind() != reflect.String && rv.Type().Kind() != reflect.Bool {
		return errors.New("invalid type")
	}
	lnBuf := make([]byte, 2)
	if _, err := io.ReadFull(d, lnBuf); err != nil {
		return err
	}
	ln := binary.BigEndian.Uint16(lnBuf)
	buf := make([]byte, ln)
	if _, err := io.ReadFull(d, buf); err != nil {
		return err
	}
	val := string(buf)
	if rv.Type().Kind() == reflect.Bool && val == "true" {
		rv.SetBool(true)
	} else if rv.Type().Kind() == reflect.Bool && val == "false" {
		rv.SetBool(false)
	} else if rv.Type().Kind() == reflect.String {
		rv.SetString(val)
	} else {
		return errors.New("invalid type")
	}
	return nil
}

// Codec to small atom
type smallAtomCodec struct{}

func (c *smallAtomCodec) Decode(d *Decoder, rv reflect.Value) error {
	if rv.Type().Kind() != reflect.String {
		return errors.New("invalid type")
	}
	var ln [1]byte
	if _, err := io.ReadFull(d, ln[:]); err != nil {
		return err
	}
	buf := make([]byte, ln[0])
	if _, err := io.ReadFull(d, buf); err != nil {
		return err
	}
	val := string(buf)
	rv.SetString(val)
	return nil
}

// Codec to small atom
type binaryCodec struct{}

func (c *binaryCodec) Decode(d *Decoder, rv reflect.Value) error {
	if rv.Type().Kind() != reflect.Slice {
		return errors.New("invalid type")
	} else if rv.Type().Elem().Kind() != reflect.Uint8 {
		return errors.New("invalid type")
	}
	lnBuf := make([]byte, 4)
	if _, err := io.ReadFull(d, lnBuf); err != nil {
		return err
	}
	ln := binary.BigEndian.Uint32(lnBuf)
	val := make([]byte, ln)
	if _, err := io.ReadFull(d, val); err != nil {
		return err
	}
	rv.SetBytes(val)
	return nil
}
