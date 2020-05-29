package tests

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/json-ice_test --type=MapTypes"
type MapTypes struct {
	Map             map[string]string            `json:"map"`
	MapSlice        []map[string]string          `json:"map_slice"`
	NestedMap       map[string]map[string]string `json:"nested_map"`
	IntKeyMap       map[int]string               `json:"int_key_map"`
	SliceValueMap   map[string][]string          `json:"slice_value_map"`
	PointerValueMap map[string]*string           `json:"pointer_value_map"`
	NullMap         map[string]string            `json:"null_map"`
}
