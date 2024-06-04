init_hook:
	git config core.hooksPath ./hooks

install:
	@echo "Installing dependencies..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	# go install -v github.com/incu6us/goimports-reviser/v3@latest
	@echo "Dependencies installed successfully!"

init: init_hook install
	@echo "Init done!"

lint:
	# 简化 go mod 文件
	go mod tidy
	# 校验代码规范
	golangci-lint run --fix

build:
	@echo "Building..."
	go build -o server ./internal
run :
	@echo "Running..."
	./server -config ./config/config.yaml -casbin ./config/casbin/model.conf
.PHONY: init