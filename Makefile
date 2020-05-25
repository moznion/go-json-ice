test: build4test gen
	go test -v ./...

build4test:
	go build -o dist/json-ice_test cmd/json-ice/main.go

gen:
	go generate ./...
