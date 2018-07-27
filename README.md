# etf
[![Build Status](https://travis-ci.com/eiri/etf.svg?branch=master)](https://travis-ci.com/eiri/etf)

Package etf implements encoding and decoding of Erlang's External Term Format

## Type mapping

Erlang              | Tag | Go
------------------- | --- | -------
SMALL_INTEGER_EXT   |  97 | int
INTEGER_EXT         |  98 | int
FLOAT_EXT           |  99 | ---
REFERENCE_EXT       | 101 | ---
PORT_EXT            | 102 | ---
PID_EXT             | 103 | ---
SMALL_TUPLE_EXT     | 104 | TBD
LARGE_TUPLE_EXT     | 105 | TBD
MAP_EXT             | 116 | TBD
NIL_EXT             | 106 | TBD
STRING_EXT          | 107 | TBD
LIST_EXT            | 108 | TBD
BINARY_EXT          | 109 | []byte
SMALL_BIG_EXT       | 110 | ---
LARGE_BIG_EXT       | 111 | ---
NEW_REFERENCE_EXT   | 114 | ---
FUN_EXT             | 117 | ---
NEW_FUN_EXT         | 112 | ---
EXPORT_EXT          | 113 | ---
BIT_BINARY_EXT      |  77 | ---
NEW_FLOAT_EXT       |  70 | float64
ATOM_UTF8_EXT       | 118 | string
SMALL_ATOM_UTF8_EXT | 119 | string
ATOM_EXT            | 100 | string or bool
SMALL_ATOM_EXT      | 115 | ---

## Remainder to myself

Since I mostly need this only for [Herbal](https://github.com/eiri/herbal)
and (probably) SASL log reader let's go for a minimum viable product
and do ony necessary bits of decoder.

Ok, maybe an encoder too later, for sheer fun of it.
