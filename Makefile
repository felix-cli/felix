.PHONY: build test test-short vendor

build: vendor
	go build -o bin/felix cmd/felix/main.go

vendor:
	go mod tidy
	go mod vendor

test:
	go test -count=1 ./...

test-short:
	go test -short -count=1 ./...

test-update:
	go test ./... -update
