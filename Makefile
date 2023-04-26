.PHONY: build gen_example
build:
	go build -o build/ ./cmd/protoc-gen-go-grpc-fixture

gen_example: build
	PATH=$$PWD/build:$$PATH && cd example && buf generate .
