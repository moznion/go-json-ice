test: build4test gen
	go test -v ./...

bench: build4test gen
	@(cd benchmark && \
		../dist/json-ice_test --type=BasicTypes && \
		echo "=== auto capsize ===" && \
		go test -bench . -benchmem)
	@(cd benchmark && \
		../dist/json-ice_test --type=BasicTypes --cap-size=1 && \
		echo "=== minimum capsize => 1 ===" && \
		go test -bench . -benchmem)

build4test:
	go build -o dist/json-ice_test cmd/json-ice/main.go

gen:
	go generate ./...

fmt-check:
	gofmt -l -s **/*.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi; \
	goimports -l **/*.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi \

fmt:
	gofmt -w -s **/*.go
	goimports -w **/*.go

