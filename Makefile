sources = $(shell find . -name "*.go")

./bin/qurtc: $(sources)
	mkdir -p ./bin
	go mod tidy
	go build -o ./bin/qurtc .
