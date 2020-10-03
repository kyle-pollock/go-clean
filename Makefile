.POSIX:

all:
	mkdir -p dist
	go build -o dist/ ./cmd/...

clean:
	rm -rf ./dist

run:
	go run ./cmd/go-clean

test:
	go test ./... -p 1

test-ci: test-short

test-short:
	go test ./... --short

.PHONY: run test test-ci test-short
