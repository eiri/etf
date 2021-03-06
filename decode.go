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

// Read because our decoder is itself io.Reader
func (d *Decoder) Read(p []byte) (n int, err error) {
	n, err = d.Reader.Read(p)
	return
}

// Decode reads from input and stores value in value or v pointer
func (d *Decoder) Decode(v interface{}) (err error) {
	// don't bother if we can't set value
	rv := reflect.Indirect(reflect.ValueOf(v))
	if !rv.CanAddr() {
		return errors.New("not a pointer")
	}
	// grep header and tag
	probe := make([]byte, 2)
	if _, err := io.ReadFull(d, probe); err != nil {
		return err
	}
	if probe[0] != 131 {
		return errors.New("bad argument")
	}
	c, err := getCodecForTag(probe[1])
	if err == nil {
		err = c.Decode(d, rv)
	}
	return
}
