package main

import (
	"github.com/moznion/go-json-ice/tests"
)

func main() {
	s := &tests.BasicTypes{
		BoolValue:    true,
		IntValue:     1,
		Int8Value:    12,
		Int16Value:   123,
		Int32Value:   1234,
		Int64Value:   12345,
		UintValue:    2,
		Uint8Value:   23,
		Uint16Value:  234,
		Uint32Value:  2345,
		Uint64Value:  23456,
		Float32Value: 1.234,
		Float64Value: 2.345,
		StringValue:  "foobar",
	}
	_, err := tests.MarshalBasicTypesAsJSON(s)
	if err != nil {
		panic(err)
	}
}
