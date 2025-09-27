sources = $(shell find . -name "*.go")

all: ./bin/qurtc ./bin/qurtc.wasm

./bin/qurtc: $(sources)
	mkdir -p ./bin
	go mod tidy
	go build -o ./bin/qurtc .

./bin/qurtc.wasm: $(sources)
	mkdir -p ./bin
	go mod tidy
	GOOS=js GOARCH=wasm go build -o ./bin/qurtc.wasm .
