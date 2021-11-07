.PHONY: build test test-short test-update vendor version

latest_tag := $(shell git describe --tags --abbrev=0)
version := ${latest_tag}

commits_since_tag := $(shell git rev-list ${latest_tag}..HEAD --count)
ifneq (${commits_since_tag}, 0)
	version = ${latest_tag}-${commits_since_tag}
endif

go_ldflags := "-X main.Version=${version}"

build: vendor
	go build -ldflags=${go_ldflags} -o bin/felix cmd/felix/main.go

test:
	go test -count=1 ./...

test-short:
	go test -short -count=1 ./...

test-update:
	go test ./... -update

version:
	@echo ${version}

vendor:
	go mod tidy
	go mod vendor
