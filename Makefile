.PHONY: build
build: go-modules-tidy
	go build -o bin/felix cmd/felix/main.go

.PHONY: go-modules-tidy
go-modules-tidy:
	go mod tidy
