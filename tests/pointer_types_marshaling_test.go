package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalPointerTypesWithValues(t *testing.T) {
	boolValue := true
	intValue := int(-1)
	int8Value := int8(-123)
	int16Value := int16(-1234)
	int32Value := int32(-12345)
	int64Value := int64(-123456)
	uintValue := uint(1)
	uint8Value := uint8(123)
	uint16Value := uint16(1234)
	uint32Value := uint32(12345)
	uint64Value := uint64(123456)
	float32Value := float32(123.456)
	float64Value := float64(1234.5678)
	stringValue := `hello "world"`

	given := &PointerTypes{
		BoolValue:    &boolValue,
		IntValue:     &intValue,
		Int8Value:    &int8Value,
		Int16Value:   &int16Value,
		Int32Value:   &int32Value,
		Int64Value:   &int64Value,
		UintValue:    &uintValue,
		Uint8Value:   &uint8Value,
		Uint16Value:  &uint16Value,
		Uint32Value:  &uint32Value,
		Uint64Value:  &uint64Value,
		Float32Value: &float32Value,
		Float64Value: &float64Value,
		StringValue:  &stringValue,
	}

	serialized, err := given.MarshalJSON()
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	var got PointerTypes
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, *given, got)
}

func TestMarshalPointerTypesWithNulls(t *testing.T) {
	given := &PointerTypes{
		BoolValue:    nil,
		IntValue:     nil,
		Int8Value:    nil,
		Int16Value:   nil,
		Int32Value:   nil,
		Int64Value:   nil,
		UintValue:    nil,
		Uint8Value:   nil,
		Uint16Value:  nil,
		Uint32Value:  nil,
		Uint64Value:  nil,
		Float32Value: nil,
		Float64Value: nil,
		StringValue:  nil,
	}

	serialized, err := given.MarshalJSON()
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	var got PointerTypes
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, *given, got)
}
