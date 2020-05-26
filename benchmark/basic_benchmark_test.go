package benchmark

import (
	"encoding/json"
	"testing"
)

var s = &BasicTypes{
	BoolValue:    true,
	IntValue:     -1,
	Int8Value:    -123,
	Int16Value:   -1234,
	Int32Value:   -12345,
	Int64Value:   -123456,
	UintValue:    1,
	Uint8Value:   123,
	Uint16Value:  1234,
	Uint32Value:  12345,
	Uint64Value:  123456,
	Float32Value: 123.456,
	Float64Value: 1234.5678,
	StringValue:  `hello "world"`,
}

func BenchmarkMarshal_EncodingJSON(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(s)
	}
}

func BenchmarkMarshal_JSONIce(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MarshalBasicTypesAsJSON(s)
	}
}
