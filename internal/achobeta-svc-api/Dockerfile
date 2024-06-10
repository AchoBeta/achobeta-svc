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
RUN  go build -o svc-api-server ./internal/achobeta-svc-api/internal
# 把文件复制到当前工作目录
ADD ./internal/achobeta-svc-api/config/config.yaml ./config/api/config.yaml
# 设置默认命令
ENTRYPOINT ["./svc-api-server", "-config", "./config/api/config.yaml"]