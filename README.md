# protoc-gen-go-grpc-fixture

`protoc-gen-go-grpc-fixture` generates gRPC client stubs which 
read from a fixture file instead of making a real gRPC call.

## Installation

```
go install github.com/snarky-puppy/protoc-gen-go-grpc-fixture@latest
```

Also required:

- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://github.com/golang/protobuf)

## Usage

### with protoc

```shell
protoc --go_out=. --go-grpc_out=. --go-grpc-fixture_out=. petstore.proto 
```

### with buf

```yaml
version: v1

plugins:
  - name: go
    out: .
    opt: paths=source_relative

  - name: go-grpc
    out: .
    opt: paths=source_relative

  - name: go-grpc-fixture
    out: .
    opt: paths=source_relative
```