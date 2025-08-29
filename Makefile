sources = $(shell find . -name "*.go")

./bin/qurtc: $(sources)
	go mod tidy
	go build -o ./bin/qurtc ./cmd/qurtc/
