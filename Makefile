TEST?=$$(go list ./... | grep -v 'vendor')
SHELL := /bin/bash
GOOS=linux
GOARCH=arm64
VERSION=test

# List of targets the `readme` target should call before generating the readme
export README_DEPS ?= docs/targets.md

-include $(shell curl -sSL -o .build-harness "https://cloudposse.tools/build-harness"; echo .build-harness)

## Lint terraform code
lint:
	$(SELF) terraform/install terraform/get-modules terraform/get-plugins terraform/lint terraform/validate

build-all: build-listener

build-listener:
	cd lambdas && GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 go build -v -ldflags -o build/listener/bootstrap -tags lambda.norpc ./cmd/listener
	cd lambdas/build/listener/ && zip listener-lambda.zip bootstrap

deps:
	go mod download

# Run acceptance tests
testacc: deps
	go test $(TEST) -v $(TESTARGS) -timeout 2m

.PHONY: lint build-listener build-all deps version testacc
