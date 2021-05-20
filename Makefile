all: run

clean:
	rm -rf bin

run: build
	./bin/mapper

build:
	go build -o bin/mapper ./cmd

test: test_capitalize test_cmd

test_capitalize:
	go test ./capitalize/...

test_cmd:
	go test ./cmd/...
