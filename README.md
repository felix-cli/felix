# felix
Felix is a CLI to quickly create service based on templates. For now it uses a default template but soon it will be able to accept all.

## installing
`go get -u github.com/scottcrawford03/felix/...`

## templates
* `https://github.com/scottcrawford03/grpc-service.felix`

# Quick Start

## Environment set up

Ensure Go is installed and `/usr/local/go/bin` and `$HOME/go/bin` is in your `PATH` environment variable.

Ref = https://golang.org/doc/install

Example:
`export PATH="$PATH:/usr/local/go/bin:$HOME/go/bin"`

Ensure your `GOPATH` is set with something like:

`export GOPATH="$GOPATH:$HOME/go"`


## Install Protobuf and GRPC Gateway

Install protobuf for MacOS using brew:

`brew install protobuf`

Install GRPC Gateway:

https://github.com/grpc-ecosystem/grpc-gateway#installation
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Install Felix

`go get -u github.com/scottcrawford03/felix/...`

## Example usage

`felix version`

Output:

`felix version felix0.1.2%`


Make a test directory that will house your new Go GRPC project:
```
mkdir ~/test-felix
cd ~/test-felix`
```

Initialize a project

`felix fixit`

Output:
```
Proj [update]: test
Org [update_me]: tester
All done!
```

Run main.go

`make run`

Output:
```
. . . 
{"level":"info","ts":1572414251.498708,"caller":"felix/main.go:64","msg":"Starting gRPC server on port 8080"}
{"level":"info","ts":1572414251.4988759,"caller":"felix/main.go:74","msg":"Starting http server on port 8000"}
```

In another terminal window: 

`curl localhost:8000/hello_world`

Output:

`{"response":{"msg":"Welcome to the future!"}}%`

Press control-c or close the terminal windown you are running the `make run` to stop the server.

Now you can make some changes. For example, edit the Server and change the greeting from `Welcome to the future!` to `Yo, Welcome to the future!`

`vi server/handler.go`

Import and compile with the annotation.proto via the Makefile

`make proto`

Output:
```
protoc -I/usr/local/include -I. -I/Users/tester/go/src -I/Users/tester/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. services/*/*.proto;
```

Run main.go

`make run`

Output:
```
go run main.go
{"level":"info","ts":1572416470.9507,"caller":"felix/main.go:64","msg":"Starting gRPC server on port 8080"}
{"level":"info","ts":1572416470.9508128,"caller":"felix/main.go:74","msg":"Starting http server on port 8000"}
```

In another terminal window run:

`curl localhost:8000/hello_world`

Output:

```{"response":{"msg":"Yo, Welcome to the future!"}}%```

Press control-c or close the terminal windown you are running the `make run` to stop the server.
