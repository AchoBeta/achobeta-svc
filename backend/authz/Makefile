init_hook:
	git config core.hooksPath ./hooks

install:
	@echo "Installing dependencies..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	# go install -v github.com/incu6us/goimports-reviser/v3@latest
	@echo "Dependencies installed successfully!"

init: init_hook install
	@echo "Init done!"
wire:
	@wire  ./internal/service
	@echo "wire successfully!"
	
lint:
	@echo "authz Linting..."
	golangci-lint run --fix
	@echo "lint achobeta-svc-authz done"

arch ?= $(shell uname -m)
build:
	@echo "Building svc-authz..."
	CGO_ENABLED=0  GOOS=linux  GOARCH=$(arch) go build -o svc-authz-server ./internal

run :
	@echo "Running..."
	./svc-authz-server  -config ./config/config.yaml -casbin ./config/model.conf
.PHONY: init