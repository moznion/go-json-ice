PKGS := $(shell go list ./...)
PKGS_WITHOUT_TEST := $(shell go list ./... | grep -v "tests" | grep -v "benchmark")

check: test lint vet fmt-check

build4test:
	go build -o dist/json-ice_test cmd/json-ice/main.go

gen4test:
	go generate ./...

test: build4test gen4test
	go test -v ./...

bench: build4test gen4test
	@(cd benchmark && \
		../dist/json-ice_test --type=BasicTypes && \
		echo "=== auto capsize ===" && \
		go test -bench . -benchmem)
	@(cd benchmark && \
		../dist/json-ice_test --type=BasicTypes --cap-size=1 && \
		echo "=== minimum capsize => 1 ===" && \
		go test -bench . -benchmem)

lint:
	golint -set_exit_status $(PKGS_WITHOUT_TEST)

vet:
	go vet $(PKGS)

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

clean:
	rm -rf dist/json-ice*

