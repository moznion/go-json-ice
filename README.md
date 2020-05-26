# go-json-ice

Simple, minimal runtime dependencies and fast JSON marshaler for go/tinygo.

The status of this library is under development. API and all other has the possibility of changing without notice.

## Motivation

### for "restricted" tinygo environment

[tinygo](https://github.com/tinygo-org/tinygo) doesn't support `encoding/json` package because reflection support is not enough: https://tinygo.org/lang-support/stdlib/#encoding-json

This library aims to marshal a struct to JSON on tinygo environment. And it also suppose to run this library in restricted environment; for example, wasm sandbox runtime.

So this library has the following features:

- don't depend on `encoding/json` package
- don't depend on `js` package
  - don't depend on browser wasm runtime

This library has only one dependency for now, `strconv`.

## Synopsis

TBD

## How it works

TBD

## Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/moznion/go-json-ice/benchmark
BenchmarkMarsha_EncodingJSON-12          1148733              1022 ns/op             320 B/op          1 allocs/op
BenchmarkMarshal_JSONIce-12              1932854               619 ns/op             512 B/op          1 allocs/op
PASS
ok      github.com/moznion/go-json-ice/benchmark        4.520s
```

`make bench` command benchmarks the performance.

## Author

moznion (<moznion@gmail.com>)

