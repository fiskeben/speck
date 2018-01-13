.PHONY: build

build:
	go build -o mcro *.go

install:
	go install

clean: rm -f mcro
