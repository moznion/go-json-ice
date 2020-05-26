package benchmark

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/json-ice_test --type=BasicTypes"
type BasicTypes struct {
	BoolValue    bool    `json:"bool_value"`
	IntValue     int     `json:"int_value"`
	Int8Value    int8    `json:"int8_value"`
	Int16Value   int16   `json:"int16_value"`
	Int32Value   int32   `json:"int32_value"`
	Int64Value   int64   `json:"int64_value"`
	UintValue    uint    `json:"uint_value"`
	Uint8Value   uint8   `json:"uint8_value"`
	Uint16Value  uint16  `json:"uint16_value"`
	Uint32Value  uint32  `json:"uint32_value"`
	Uint64Value  uint64  `json:"uint64_value"`
	Float32Value float32 `json:"float32_value"`
	Float64Value float64 `json:"float64_value"`
	StringValue  string  `json:"string_value"`
}
