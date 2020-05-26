package tests

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/json-ice_test --type=OmitemptyStruct"
type OmitemptyStruct struct {
	EmptyBool      bool    `json:"empty_bool,omitempty"`
	EmptyInt       int     `json:"empty_int,omitempty"`
	EmptyUint      uint    `json:"empty_uint,omitempty"`
	EmptyFloat     float32 `json:"empty_float,omitempty"`
	EmptyString    string  `json:"empty_string,omitempty"`
	NotEmptyString string  `json:"not_empty_string"`
}

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/json-ice_test --type=OmitemptyPointerStruct"
type OmitemptyPointerStruct struct {
	EmptyBool      *bool    `json:"empty_bool,omitempty"`
	EmptyInt       *int     `json:"empty_int,omitempty"`
	EmptyUint      *uint    `json:"empty_uint,omitempty"`
	EmptyFloat     *float32 `json:"empty_float,omitempty"`
	EmptyString    *string  `json:"empty_string,omitempty"`
	NotEmptyString *string  `json:"not_empty_string"`
}
