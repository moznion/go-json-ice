package tests

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/json-ice_test --type=SliceTypes"
type SliceTypes struct {
	StringSlice        []string  `json:"string_slice"`
	StringPointerSlice []*string `json:"string_pointer_slice"`
	EmptySlice         []string  `json:"empty_slice"`
}
