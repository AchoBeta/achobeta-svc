FROM golang:1.23 AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release

WORKDIR /app
COPY . .
COPY ./backend/authz/svc-authz-server .
ADD ./backend/authz/config/config.yaml ./config/authz/config.yaml
ADD ./backend/authz/config/model.conf ./config/authz/model.conf
EXPOSE 4396
# 设置默认命令
ENTRYPOINT ["./svc-authz-server", "-config", "./config/authz/config.yaml", "-casbin", "./config/authz/model.conf"]