FROM golang:1.22.1 AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release
    
WORKDIR /app
COPY . .
# RUN  go build -o svc-api-server ./internal/achobeta-svc-api/internal

FROM alpine:latest
WORKDIR /app
COPY ./internal/achobeta-svc-api/svc-api-server .
COPY ./internal/achobeta-svc-api/config/config.yaml ./config/api/config.yaml
EXPOSE 4397
# run the binary
ENTRYPOINT ["./svc-api-server", "-config", "./config/api/config.yaml"]