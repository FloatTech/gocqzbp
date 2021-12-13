FROM golang:1.17-alpine AS builder

RUN go env -w GO111MODULE=on \
  && go env -w CGO_ENABLED=0 \
  && go env

WORKDIR /build

COPY ./ .

RUN set -ex \
    && cd /build \
    && go mod tidy\
    && chmod 755 `go env GOMODCACHE`/github.com/wdvxdr1123/!zero!bot@v1.4.1/driver\
    && mv funcall.txt `go env GOMODCACHE`/github.com/wdvxdr1123/!zero!bot@v1.4.1/driver/funcall.go\
    && go build -ldflags "-s -w" -o cqhttp -trimpath

FROM alpine:latest

RUN apk add --no-cache ffmpeg

COPY --from=builder /build/cqhttp /usr/bin/cqhttp
RUN chmod +x /usr/bin/cqhttp

WORKDIR /data

ENTRYPOINT [ "/usr/bin/cqhttp" ]
