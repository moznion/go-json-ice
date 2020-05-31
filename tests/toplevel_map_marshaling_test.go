package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalToplevelMap(t *testing.T) {
	given := map[string]string{
		"foo": "bar",
		"buz": "qux",
	}
	serialized, err := MarshalStringToStringMapAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	var got map[string]string
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, given, got)
}
