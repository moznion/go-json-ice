package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalBasicTypes(t *testing.T) {
	given := &BasicTypes{
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

	serialized, err := MarshalBasicTypesAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	var got BasicTypes
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, *given, got)
}
