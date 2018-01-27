.PHONY: build

BINARY=speck
VERSION=$(shell git describe --tags --always)

LD_FLAGS=-ldflags "-X 'github.com/fiskeben/speck/command.version=${VERSION}'"

build:
	go build ${LD_FLAGS} -o speck *.go

test:
	go test ./...

install:
	go install

release: test
	go build ${LD_FLAGS} -o speck-${VERSION} *.go

clean:
	@rm -f speck
