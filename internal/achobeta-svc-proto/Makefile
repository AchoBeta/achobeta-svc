lint:
	@buf lint ./proto
	@echo "Check done"

clean:
	@rm -rf ./gen
	@echo "Clean done"

.PHONY: install
install:
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get github.com/envoyproxy/protoc-gen-validate@v1.0.4
	go get github.com/envoyproxy/protoc-gen-validate/templates/java@v1.0.4
	go get github.com/lyft/protoc-gen-star/v2@latest
	go get github.com/iancoleman/strcase@latest
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/envoyproxy/protoc-gen-validate;
 
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
