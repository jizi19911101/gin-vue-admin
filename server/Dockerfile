FROM golang:alpine as builder

WORKDIR /go/src/github.com/jizi19911101/gin-vue-admin/server
COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env
RUN go mod tidy
RUN go build -o server .
RUN cat > /etc/apk/repositories << EOF
    http://mirrors.aliyun.com/alpine/v3.12/main/
    http://mirrors.aliyun.com/alpine/v3.12/community
    EOF
RUN apk update && apk add curl openssh-client busybox-extras vim

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/github.com/jizi19911101/gin-vue-admin/server

COPY --from=0 /go/src/github.com/jizi19911101/gin-vue-admin/server ./
COPY ./config.docker.yaml ./config.yaml

EXPOSE 8888

ENTRYPOINT ./server