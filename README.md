# etf
[![Build Status](https://travis-ci.com/eiri/etf.svg?branch=master)](https://travis-ci.com/eiri/etf)

Package etf implements encoding and decoding of Erlang's External Term Format

## Type mapping

Erlang            | Tag | Go
----------------- | --- | -----
SMALL_INTEGER_EXT |  97 | uint8
INTEGER_EXT       |  98 | int32
