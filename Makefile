.PHONY: all deps test clean
all: test
	go build -v 

deps:
	go get gopkg.in/go-playground/assert.v1
	go get -v ./...

test: deps
	go test -v ./...

clean:
	rm nanoweb
