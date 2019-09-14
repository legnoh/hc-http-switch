NAME := hc-http-switch
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)'
GO ?= GO111MODULE=on go
.DEFAULT_GOAL := help

.PHONY: build
build: main.go  ## Build a binary.
	$(GO) build -ldflags "$(LDFLAGS)"

.PHONY: cross
cross: main.go  ## Build binaries for cross platform.
	mkdir -p pkg
	@# darwin
	@for arch in "amd64" "386"; do \
		GOOS=darwin GOARCH=$${arch} make build; \
		zip pkg/$(NAME)_$(VERSION)_darwin_$${arch}.zip $(NAME); \
	done;
	@# linux
	@for arch in "amd64" "386" "arm64"; do \
		GOOS=linux GOARCH=$${arch} make build; \
		zip pkg/$(NAME)_$(VERSION)_linux_$${arch}.zip $(NAME); \
	done;
	@# linux(raspberry pi)
	@for arm in "5" "6" "7"; do \
		GOOS=linux GOARM=$${arm} GOARCH=arm make build; \
		zip pkg/$(NAME)_$(VERSION)_linux_arm$${arm}.zip $(NAME); \
	done;

.PHONY: help
help: ## Show help text
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'
