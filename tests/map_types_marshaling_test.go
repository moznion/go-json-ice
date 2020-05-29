package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalMap(t *testing.T) {
	pointerStr := "pointer"
	given := &MapTypes{
		Map: map[string]string{
			"foo": "bar",
			"buz": "qux",
		},
		MapSlice: []map[string]string{
			{
				"foo": "bar",
			},
			{
				"buz": "qux",
			},
		},
		NestedMap: map[string]map[string]string{
			"map1": {
				"foo": "bar",
				"buz": "qux",
			},
			"map2": {
				"hoge": "fuga",
			},
		},
		IntKeyMap: map[int]string{
			1: "yo",
		},
		SliceValueMap: map[string][]string{
			"foo": {"bar", "buz"},
			"qux": {},
		},
		PointerValueMap: map[string]*string{
			"foo": &pointerStr,
			"bar": nil,
		},
		NullMap: nil,
	}

	serialized, err := MarshalMapTypesAsJSON(given)
	assert.NoError(t, err)

	log.Printf("[debug] %s", serialized)

	var got MapTypes
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, *given, got)
}
