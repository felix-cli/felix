.PHONY: all build install lint proto race run test vendor

all: build test

PROTO_PACKAGES := services/*/*.proto

# Build

build: proto
	go build ./...

install:
	go install ./...

lint:
	golint ./...

proto:
	$(foreach package,$(PROTO_PACKAGES), \
		protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		$(package);)

race:
	go test -race ./...

run:
	go run main.go

test:
	go test -cover ./...

vendor:
	go mod tidy
	go mod vendor
