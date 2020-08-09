.PHONY: all deps test clean
all: test
	go build -v 

deps:
	#go get -v ./...
	go mod tidy

test: deps
	go test -v ./...

clean:
	go clean
