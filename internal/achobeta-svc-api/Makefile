lint:
	golangci-lint run --fix
	@echo "lint achobeta-svc-authz done"

init_hook:
	git config core.hooksPath ./hooks
	
wire:
	@wire  ./internal/server/service
	@echo "wire successfully!"
	
install:
	@echo "Installing dependencies..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	# go install -v github.com/incu6us/goimports-reviser/v3@latest
	@echo "Dependencies installed successfully!"

init: init_hook install
	@echo "Init done!"

arch ?= $(shell uname -m)
build:
	@echo "Building..."
	CGO_ENABLED=0  GOOS=linux  GOARCH=$(arch) go build -o svc-api-server ./internal

run :
	@echo "Running..."
	./svc-api-server -config ./config/config.yaml

start: build run
.PHONY: build