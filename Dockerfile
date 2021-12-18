FROM golang:alpine as builder

WORKDIR /go/src/github.com/fast-crud/fast-auth
COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env
RUN go mod tidy
RUN go build -tags mysql -o server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

ENV GFVA_CONFIG = 'config/config.docker.yaml'

WORKDIR /go/src/github.com/fast-crud/fast-auth

COPY --from=0 /go/src/github.com/fast-crud/fast-auth ./

EXPOSE 8888

ENTRYPOINT ./server
