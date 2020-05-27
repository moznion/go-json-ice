package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalSliceTypes(t *testing.T) {
	stringSlice := []string{"foo", "bar", "buz"}

	s1 := "hoge"
	s2 := "fuga"
	stringPointerSlice := []*string{&s1, &s2}

	given := &SliceTypes{
		StringSlice:        stringSlice,
		StringPointerSlice: stringPointerSlice,
		EmptySlice:         []string{},
	}

	serialized, err := MarshalSliceTypesAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	var got SliceTypes
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, *given, got)
}
