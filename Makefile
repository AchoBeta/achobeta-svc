CURRENT_DIR := .
INTERNAL_DIR := $(CURRENT_DIR)/internal
# sub dirs
COMMON_DIR := $(INTERNAL_DIR)/achobeta-svc-common
PROTO_DIR := $(INTERNAL_DIR)/achobeta-svc-proto
# exclude dirs
EXCLUDE_DIRS := $(COMMON_DIR) $(PROTO_DIR)
# service dirs, exclude common and proto
SERVICE_DIRS := $(filter-out $(EXCLUDE_DIRS), $(wildcard $(INTERNAL_DIR)/*))

# 编译的时候一定要先编译proto，因为其他服务依赖proto
build: proto sbuild

# 启动的时候就比较随意了, 只要api最后启动就行
run: srun api

install:
	@GO111MODULE=on \
        GOBIN=/usr/local/bin \
    		go install github.com/bufbuild/buf/cmd/buf@v1.33.0
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: 
	@echo "Linting proto..."
	@$(MAKE) -C $(PROTO_DIR) lint
	@for dir in $(SERVICE_DIRS); do \
		$(MAKE) -C $$dir lint || exit "$$?"; \
		echo "Compile $$(basename $$dir) done"; \
	done



proto:
	@echo "Compiling proto..."
	$(MAKE) -C $(PROTO_DIR) gen

# sub build
sbuild:
	@echo $(SERVICE_DIRS)
	@for dir in $(SERVICE_DIRS); do \
		$(MAKE) -C $$dir build || exit "$$?"; \
		echo "Compile $$(basename $$dir) done"; \
	done
