CURRENT_DIR := .
INTERNAL_DIR := $(CURRENT_DIR)/backend
# sub dirs
COMMON_DIR := $(INTERNAL_DIR)/common
PROTO_DIR := $(INTERNAL_DIR)/proto
# exclude dirs
EXCLUDE_DIRS := $(COMMON_DIR) $(PROTO_DIR)
# service dirs, exclude common and proto
SERVICE_DIRS := $(filter-out $(EXCLUDE_DIRS), $(wildcard $(INTERNAL_DIR)/*))
#define params
arch ?= $(shell uname -m)
ifeq ($(arch),x86_64)
	arch := amd64
endif
# define function
define FOREACH_SERVICE
	@for dir in $(SERVICE_DIRS); do \
		$(MAKE) -C $$dir $1 arch=$(arch) || exit "$$?"; \
		echo "$1 $$dir completed"; \
	done
endef
# 编译的时候一定要先编译proto，因为其他服务依赖proto
build: proto dir-build
init: install build
# 编译proto
proto:
	@echo "Compiling proto..."
	$(MAKE) -C $(PROTO_DIR) gen

# 编译所有目录
dir-build:
	@echo $(SERVICE_DIRS)
	$(call FOREACH_SERVICE, build)

# 下载依赖
install:
	@GO111MODULE=on \
        GOBIN=/usr/local/bin \
    		go install github.com/bufbuild/buf/cmd/buf@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@$(MAKE) -C $(PROTO_DIR) install
# 校验检查
lint: 
	@echo "Linting proto..."
	@$(MAKE) -C $(PROTO_DIR) lint
	$(call FOREACH_SERVICE, lint)

# docker 启动服务
docker-run: build
	@docker-compose up --build