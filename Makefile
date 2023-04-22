.PHONY: build gen_example
build:
	go build -o build/ .

gen_example: build
	PATH=$$PWD/build:$$PATH && cd example && buf generate .