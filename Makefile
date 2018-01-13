.PHONY: build 

build:
	go build -o speck *.go

install:
	go install

clean:
	rm -f speck
