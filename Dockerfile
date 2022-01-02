FROM golang:1.17-alpine AS builder

RUN go env -w GO111MODULE=on \
  && go env -w CGO_ENABLED=0 \
  && go env

RUN apk update && apk add git

WORKDIR /build

COPY ./ .

RUN set -ex \
    && cd /build \
    && mkdir -p `go env GOMODCACHE`/github.com \
    && chmod 755 `go env GOMODCACHE`/github.com \
    && git clone --depth=1 -b dev https://github.com/FloatTech/gocq.git `go env GOMODCACHE`/github.com/!mrs4s/go-cqhttp@v1.0.0-beta8-fix2 \
    && go mod tidy -compat=1.17 \
    && chmod 755 `go env GOMODCACHE`/github.com/wdvxdr1123/!zero!bot@v1.4.1/driver\
    && mv funcall.txt `go env GOMODCACHE`/github.com/wdvxdr1123/!zero!bot@v1.4.1/driver/funcall.go\
    && go build -ldflags "-s -w" -o cqhttp -trimpath

FROM alpine:latest

RUN apk add --no-cache ffmpeg

COPY --from=builder /build/cqhttp /usr/bin/cqhttp
RUN chmod +x /usr/bin/cqhttp

WORKDIR /data

ENTRYPOINT [ "/usr/bin/cqhttp" ]
