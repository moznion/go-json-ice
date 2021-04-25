package tests

// see: https://github.com/moznion/go-json-ice/issues/24
//
//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/json-ice_test --type=Issue24Struct"
type Issue24Struct struct {
	Foo string `json:"foo"`
	Bar string
}
