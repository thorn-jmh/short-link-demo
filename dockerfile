FROM golang:1.20-bullseye AS builder

ENV \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /workdir/
COPY . ./

RUN go build -o /short-link-demo


FROM debian:bullseye

ENV GIN_MODE=release

COPY --from=builder /short-link-demo /short-link-demo

# 拷贝静态文件依赖
COPY config.yml /

EXPOSE 8080
ENTRYPOINT ["/short-link-demo"]