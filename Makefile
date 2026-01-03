BINARY := strix
BUILD_DIR := bin
MODULE := github.com/yourpwnguy/strix

VERSION := $(shell git describe --tags --dirty --always 2>/dev/null || echo dev)

.PHONY: build run clean test install uninstall fmt vet check

build:
	@mkdir -p $(BUILD_DIR)
	@go build \
		-ldflags "-s -w -X $(MODULE)/cmd.Version=$(VERSION)" \
		-o $(BUILD_DIR)/$(BINARY) \
		.
	@echo "✓ Built $(BUILD_DIR)/$(BINARY) ($(VERSION))"

run: build
	@./$(BUILD_DIR)/$(BINARY)

install: build
	@cp $(BUILD_DIR)/$(BINARY) $(GOPATH)/bin/ 2>/dev/null || cp $(BUILD_DIR)/$(BINARY) $(HOME)/go/bin/
	@echo "✓ Installed"

uninstall:
	@rm -f $(GOPATH)/bin/$(BINARY) $(HOME)/go/bin/$(BINARY)
	@echo "✓ Uninstalled"

clean:
	@rm -rf $(BUILD_DIR)
	@echo "✓ Cleaned"

test:
	@go test -v ./...

fmt:
	@go fmt ./...

vet:
	@go vet ./...

check: fmt vet test
	@echo "✓ All good"
