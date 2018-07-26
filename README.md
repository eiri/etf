# etf
[![Build Status](https://travis-ci.com/eiri/etf.svg?branch=master)](https://travis-ci.com/eiri/etf)

Package etf implements encoding and decoding of Erlang's External Term Format

## Type mapping

Erlang              | Tag | Go
------------------- | --- | -------
SMALL_INTEGER_EXT   |  97 | int
INTEGER_EXT         |  98 | int
FLOAT_EXT           |  99 | ---
NEW_FLOAT_EXT       |  70 | float64
SMALL_ATOM_EXT      | 115 | ---
ATOM_EXT            | 100 | string or bool
SMALL_ATOM_UTF8_EXT | 119 | string
ATOM_UTF8_EXT       | 118 | string
BINARY_EXT          | 109 | []byte
