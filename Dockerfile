FROM golang:1.19-alpine AS builder

RUN go env -w GO111MODULE=on \
  && go env -w CGO_ENABLED=0 \
  && go env

RUN apk update && apk add git

WORKDIR /build

COPY ./ .

RUN set -ex \
    && go mod tidy \
    && go build -ldflags "-s -w" -o cqhttp -trimpath

FROM alpine:latest

RUN apk add --no-cache ffmpeg

COPY --from=builder /build/cqhttp /usr/bin/cqhttp
RUN chmod +x /usr/bin/cqhttp

WORKDIR /data

ENTRYPOINT [ "/usr/bin/cqhttp" ]
