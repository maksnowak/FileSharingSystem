BINARY_NAME=main
OS := $(shell uname | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m | sed 's/x86_64/amd64/')


.DEFAULT_GOAL := help

.PHONY: help
help:             ## Show the help message.
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: 					## Build the application.
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

.PHONY: run
run: build 				## Run the application.
	./${BINARY_NAME}-${OS}

.PHONY: test
test: 					## Run tests.
	go test -v ./... -coverprofile=coverage.out

.PHONY: dep
dep: 					## Install dependencies.
	go mod download

.PHONY: clean
clean: 					## Clean project by deleting files in .gitignore.
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows
	git clean -Xdf

.PHONY: clean
clean:            ## Clean project by deleting files in .gitignore.

