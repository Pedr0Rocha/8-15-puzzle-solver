all: test build

build:
	go build -o ./bin/puzzle .

test:
	go test -v ./...

clean:
	rm -f ./bin/puzzle

run: build
	./bin/puzzle

fmt:
	go fmt ./...

.PHONY: all build test clean run fmt
