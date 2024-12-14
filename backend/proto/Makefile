lint:
	@buf lint ./proto
	@echo "Check done"

clean:
	@rm -rf ./gen
	@echo "Clean done"

.PHONY: install
install:

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/bufbuild/buf/cmd/buf@latest

	go get -u github.com/envoyproxy/protoc-gen-validate/templates/java
	go get -u github.com/lyft/protoc-gen-star/v2
	go get -u github.com/iancoleman/strcase
	
.PHONY: gen
gen: 
	@go mod tidy
	@rm -rf ./gen
	@buf dep update proto
	@buf build proto
	@buf generate
	@go mod tidy
	@echo "Generate done"
 
.PHONY: buf
buf:
# 或许需要网络代理
	curl -sSL && \
    	"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" && \
    	-o "$(shell go env GOPATH)/bin/buf" && && \
  	chmod +x "$(shell go env GOPATH)/bin/buf"
 
all: install gen
