FROM golang:1.22.1

# # 描述
# LABEL description="achobeta svc api"
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release
EXPOSE 4397
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download
RUN  go build -o svc-authz-server ./internal/achobeta-svc-authz/internal
# 把文件复制到当前工作目录
ADD ./internal/achobeta-svc-authz/config/config.yaml ./config/authz/config.yaml
ADD ./internal/achobeta-svc-authz/config/model.conf ./config/authz/model.conf
# 设置默认命令
ENTRYPOINT ["./svc-authz-server", "-config", "./config/authz/config.yaml", "-casbin", "./config/authz/model.conf"]