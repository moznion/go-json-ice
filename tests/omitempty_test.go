package tests

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOmitemptyWithoutFields(t *testing.T) {
	given := &OmitemptyStruct{}

	serialized, err := given.MarshalJSON()
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

	serialized, err := given.MarshalJSON()
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	assert.Equal(t, `{"not_empty_string":""}`, string(serialized))
}
