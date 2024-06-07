lint:
	# 简化 go mod 文件
	go mod tidy
	# 校验代码规范
	golangci-lint run --fix

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


build:
	@echo "Building..."
	go build -o svc-api-server ./internal

run :
	@echo "Running..."
	./svc-api-server -config ./config/config.yaml

start: build run
.PHONY: build