.PHONY: build
build: go-modules-tidy
	go build -o bin/felix cmd/felix/main.go

.PHONY: go-modules-tidy
go-modules-tidy:
	go mod tidy
	go mod vendor

.PHONY: test
test:
	go test -count=1 ./...

.PHONY: test-short
test-short:
	go test -short -count=1 ./...
