package tests

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/json-ice_test --type=NestedTypes"
type NestedTypes struct {
	Field1         OmitemptyStruct  `json:"field1"`
	Field2         *OmitemptyStruct `json:"field2"`
	OmittableField *OmitemptyStruct `json:"omittable_field,omitempty"`
}

func (o *OmitemptyStruct) MarshalJSON() ([]byte, error) {
	return MarshalOmitemptyStructAsJSON(o)
}
