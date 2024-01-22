FROM golang:alpine as builder

WORKDIR /go/src/github.com/yaoyaochil/bodo-admin-server/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -ldflags "-s -w" -o server .

FROM alpine:latest

RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

LABEL MAINTAINER="wangrui19970405@gmail.com"

WORKDIR /go/src/github.com/yaoyaochil/bodo-admin-server/server

COPY --from=0 /go/src/github.com/yaoyaochil/bodo-admin-server/server/server ./
COPY --from=0 /go/src/github.com/yaoyaochil/bodo-admin-server/server/resource ./resource/
COPY --from=0 /go/src/github.com/yaoyaochil/bodo-admin-server/server/config.docker.yaml ./

EXPOSE 5100
ENTRYPOINT ./server -c config.docker.yaml
