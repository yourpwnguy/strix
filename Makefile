.PHONY: build clean test-err test-suc

VERSION ?=1.0
LDFLAGS := -s -w \
		-X 'strix/cmd.Version=$(VERSION)'

build:
	GOOS=linux GOARCH=amd64 go build ./

build-optimized:
	GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" ./

test-err:
	-./strix info Makefile

test-suc:
	-./strix info ./ls

deps:
	go mod download
	go mod tidy
