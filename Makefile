.DEFAULT_TARGET=help

## help: Display list of commands
.PHONY: help
help: Makefile
	@sed -n 's|^##||p' $< | column -t -s ':' | sed -e 's|^| |'

## build: Build pira binary
.PHONY: build
build: fmt vet
	go build -o bin/pira

## fmt: Format source code
.PHONY: fmt
fmt:
	go fmt ./...

## vet: Vet source code
.PHONY: vet
vet:
	go vet ./...

## test: Run unit tests
.PHONY: test
test:
	go test ./...
