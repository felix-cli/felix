.PHONY: build test test-short vendor

latest_tag := $(shell git describe --tags --abbrev=0)
commits_since_tag := $(shell git rev-list ${latest_tag}..HEAD --count)
# TODO: append commits_since_tag to the version if commits_since_tag is not 0
version := ${latest_tag}
go_ldflags := "-X main.Version=${version}"

build: vendor
	go build -ldflags=${go_ldflags} -o bin/felix cmd/felix/main.go

version:
	echo ${version}

vendor:
	go mod tidy
	go mod vendor

test:
	go test -count=1 ./...

test-short:
	go test -short -count=1 ./...

test-update:
	go test ./... -update
