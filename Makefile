.PHONY: build 

build:
	@go build -o speck *.go

test:
	@go test ./...

install:
	@go install

release: test
	@go build -o speck-${shell git describe --tags --always} *.go

clean:
	@rm -f speck
