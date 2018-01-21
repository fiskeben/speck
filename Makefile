.PHONY: build 

build:
	@go build -o speck *.go

test:
	@go test ./...

install:
	@go install

clean:
	@rm -f speck
