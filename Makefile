SHELL := /bin/bash

GO111MODULE = on

test:
	go test -v ./...
.PHONY: test

lint:
	golangci-lint run -v
.PHONY: lint

install-linter:
	# install linter
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.17.1
.PHONY: install-linter

vendor:
	go mod tidy
	go mod vendor
.PHONY: vendor

generate:
	go generate ./...
.PHONY: generate
