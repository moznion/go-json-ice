package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// see: https://github.com/moznion/go-json-ice/issues/24
func TestMarshalIssue24StructAsJSON(t *testing.T) {
	given := &Issue24Struct{
		Foo: "foo",
		Bar: "bar",
	}

	serialized, err := MarshalIssue24StructAsJSON(given)
	assert.NoError(t, err)

	var got Issue24Struct
	err = json.Unmarshal(serialized, &got)
	assert.NoError(t, err)
	assert.EqualValues(t, Issue24Struct{
		Foo: "foo",
		// `Bar` should be dropped because this property doesn't have a tag for JSON serialization.
	}, got)
}
