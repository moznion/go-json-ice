package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOmitemptyWithoutFields(t *testing.T) {
	given := &OmitemptyStruct{}

	serialized, err := MarshalOmitemptyStructAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	assert.Equal(t, `{"not_empty_string":""}`, string(serialized))
}

func TestOmitemptyWithEmptyValues(t *testing.T) {
	given := &OmitemptyStruct{
		EmptyBool:      false,
		EmptyInt:       0,
		EmptyUint:      0,
		EmptyFloat:     0,
		EmptyString:    "",
		NotEmptyString: "",
	}

	serialized, err := MarshalOmitemptyStructAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	assert.Equal(t, `{"not_empty_string":""}`, string(serialized))
}

func TestOmitemptyPointerWithNullValues(t *testing.T) {
	given := &OmitemptyPointerStruct{
		EmptyBool:      nil,
		EmptyInt:       nil,
		EmptyUint:      nil,
		EmptyFloat:     nil,
		EmptyString:    nil,
		EmptySlice:     nil,
		EmptyMap:       nil,
		NotEmptyString: nil,
	}

	serialized, err := MarshalOmitemptyPointerStructAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	assert.Equal(t, `{"not_empty_string":null}`, string(serialized))
}

func TestOmitemptyPointerWithEmptyValues(t *testing.T) {
	emptyBool := false
	emptyInt := int(0)
	emptyUint := uint(0)
	emptyFloat := float32(0)
	emptyString := ""
	notEmptyString := ""

	given := &OmitemptyPointerStruct{
		EmptyBool:      &emptyBool,
		EmptyInt:       &emptyInt,
		EmptyUint:      &emptyUint,
		EmptyFloat:     &emptyFloat,
		EmptyString:    &emptyString,
		EmptySlice:     []string{},
		EmptyMap:       map[string]string{},
		NotEmptyString: &notEmptyString,
	}

	serialized, err := MarshalOmitemptyPointerStructAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	assert.Equal(t, `{"not_empty_string":""}`, string(serialized))
}

func TestOmitemptyWithEmptyStruct(t *testing.T) {
	given := &OmitemptyStruct{}

	serialized, err := MarshalOmitemptyStructAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	var got OmitemptyStruct
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, *given, got)
}

func TestOmitemptyPointerWithEmptyStruct(t *testing.T) {
	given := &OmitemptyPointerStruct{}

	serialized, err := MarshalOmitemptyPointerStructAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	var got OmitemptyPointerStruct
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, *given, got)
}
