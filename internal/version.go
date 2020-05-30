package internal

import (
	"fmt"
)

var ver string
var rev string

//go:generate sh -c "go run $(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/cmd/json-ice/main.go --type=version --cap-size=74"
type version struct {
	Version  string `json:"version"`
	Revision string `json:"revision"`
}

// ShowVersion shows the version and revision as JSON string.
func ShowVersion() {
	versionJSON, _ := MarshalversionAsJSON(&version{
		Version:  ver,
		Revision: rev,
	})
	fmt.Printf("%s\n", versionJSON)
}
