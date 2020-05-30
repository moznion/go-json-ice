# go-json-ice

Simple, less runtime dependencies and fast JSON marshaler for go and tinygo.

## Motivation

### for "restricted" tinygo environment

[tinygo](https://github.com/tinygo-org/tinygo) doesn't support `encoding/json` package because reflection support is not enough: https://tinygo.org/lang-support/stdlib/#encoding-json

This library aims to marshal a struct to JSON on tinygo environment. And it also supposes to run this library in restricted environment; for example, wasm sandbox runtime.

So this library has the following features:

- don't depend on `encoding/json` package
- don't depend on `js` package
  - i.e. don't depend on browser wasm runtime

This library has only one dependency for now, `strconv`.

### Better performance

`encoding/json` uses runtime reflection to marshal and unmarshal, but this library doesn't depend on it. Basically, code generation a marshaller statically by without reflection makes the performance be better.

## Synopsis

`go generate` with the following struct that uses `json-ice`,

```go
//go:generate json-ice --type=AwesomeStruct
type AwesomeStruct struct {
	Foo string `json:"foo"`
	Bar string `json:"bar,omitempty"`
}
```

then it generates a marshaler as `MarshalAwesomeStructAsJSON(s *AwesomeStruct) ([]byte, error)`; you can use that like the following:

```go
marshaled, err := MarshalAwesomeStructAsJSON(&AwesomeStruct{
	Foo: "buz",
	Bar: "",
})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%s\n", marshaled) // => {"foo":"buz"}
```

## Usage

```
json-ice:
  -cap-size int
        [optional] a cap-size of a byte slice buffer for marshaling; by default, it calculates this value automatically
  -output string
        [optional] output file name (default "srcdir/<type>_gen.go")
  -type string
        [mandatory] a type name
  -version
        show version information
```

## How to install

```
$ go get -u github.com/moznion/go-json-ice/cmd/json-ice
```

also you can get the pre-built binaries on [Releases](https://github.com/moznion/go-json-ice/releases).

## How it works

Parses given struct and generates a marshaler according to the parsed result; it means resolve fields statically without runtime reflection.

## Performance

It calculates a cap-size of a byte slice buffer for marshaling automatically, but you can specify it by yourself.

If you need more throughput, it is a good idea to specify it as a bigger cap-size, on the other hand, if you would like to save the amount of memory, you can put it as a smaller (e.g. `1`) cap-size.

### Benchmark

```
=== auto capsize ===
goos: darwin
goarch: amd64
pkg: github.com/moznion/go-json-ice/benchmark
BenchmarkMarshal_EncodingJSON-12         1179334              1010 ns/op             320 B/op          1 allocs/op
BenchmarkMarshal_JSONIce-12              2014201               629 ns/op             384 B/op          1 allocs/op
PASS
ok      github.com/moznion/go-json-ice/benchmark        4.099s
=== minimum capsize => 1 ===
goos: darwin
goarch: amd64
pkg: github.com/moznion/go-json-ice/benchmark
BenchmarkMarshal_EncodingJSON-12         1136454              1003 ns/op             320 B/op          1 allocs/op
BenchmarkMarshal_JSONIce-12              1421254               846 ns/op            1009 B/op          7 allocs/op
PASS
ok      github.com/moznion/go-json-ice/benchmark        4.287s
```

`auto capsize` benchmarks it the default situation, and `minimum capsize` benchmarks it with the minimum cap-size (i.e. in worst performance case: `1`).

`make bench` command benchmarks the performance.

## FAQ

### How to marshal non-primitive types

For example,

```go
type Foo struct {
	Hoge string `json:"hoge"`
}

//go:generate json-ice --type=Bar
type Bar struct {
	Foo *Foo `json:"foo"`
}
```

In this case, a marshaler for `Bar` tries to marshal `Foo` by `Foo.MarshalJSON() ([]byte, error)`. It means it expects the target type implements `json.Marshaler` when it tries to marshal a non-primitive type.

The easiest way to marshal the struct like the above is the follwoing:

```go
//go:generate json-ice --type=Foo
type Foo struct {
	Hoge string `json:"hoge"`
}

//go:generate json-ice --type=Bar
type Bar struct {
	Foo *Foo `json:"foo"`
}

func (f *Foo) MarshalJSON() ([]byte, error) {
	return MarshalversionAsJSON(f)
}
```

## Restrictions / Known issues

- it cannot marshal `named type` and `type alias` types automatically
  - if you would like to marshal such them, please implement `json.Marshaler` on the type as like as the above FAQ answer.

## How to build binaries

Binaries are built and uploaded by [goreleaser](https://goreleaser.com/). Please refer to the configuration file: [.goreleaser.yml](./.goreleaser.yml)

## Author

moznion (<moznion@gmail.com>)

