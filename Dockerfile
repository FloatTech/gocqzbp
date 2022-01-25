FROM golang:1.17-alpine AS builder

RUN go env -w GO111MODULE=on \
  && go env -w CGO_ENABLED=0 \
  && go env

RUN apk update && apk add git

WORKDIR /build

COPY ./ .

RUN set -ex \
    && mkdir -p `go env GOMODCACHE`/github.com \
    && chmod 755 `go env GOMODCACHE`/github.com \
    && git clone --depth=1 -b dev https://github.com/FloatTech/gocq.git `go env GOMODCACHE`/github.com/!mrs4s/go-cqhttp@v1.0.0-beta8-fix2 \
    && git clone --depth=1 https://github.com/wdvxdr1123/ZeroBot.git `go env GOMODCACHE`/github.com/wdvxdr1123/!zero!bot@v1.4.2-0.20220122162257-bc71c479f3d1 \
    && mv funcall.txt `go env GOMODCACHE`/github.com/wdvxdr1123/!zero!bot@v1.4.2-0.20220122162257-bc71c479f3d1/driver/funcall.go \
    && go mod tidy -compat=1.17 \
    && go build -ldflags "-s -w" -o cqhttp -trimpath

FROM alpine:latest

RUN apk add --no-cache ffmpeg

COPY --from=builder /build/cqhttp /usr/bin/cqhttp
RUN chmod +x /usr/bin/cqhttp

WORKDIR /data

ENTRYPOINT [ "/usr/bin/cqhttp" ]
