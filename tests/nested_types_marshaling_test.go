package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalingNestedTypes(t *testing.T) {
	given := &NestedTypes{
		Field1: OmitemptyStruct{
			NotEmptyString: "foo",
		},
		Field2: &OmitemptyStruct{
			NotEmptyString: "bar",
		},
		OmittableField: &OmitemptyStruct{
			NotEmptyString: "buz",
		},
	}

	serialized, err := MarshalNestedTypesAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	var got NestedTypes
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, *given, got)
}

func TestMarshlingEmptyNestedTypes(t *testing.T) {
	given := &NestedTypes{}

	serialized, err := MarshalNestedTypesAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	var got NestedTypes
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, *given, got)
}
