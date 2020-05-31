package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalToplevelSlice(t *testing.T) {
	given := []string{"foo", "bar", "buz"}
	serialized, err := MarshalStringArrayAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	var got []string
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, given, got)
}
