ARG GO_VERSION=1.18.1
ARG ALPINE_VERSION=3.15

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build
ENV APP_ROOT /go/src/app

WORKDIR ${APP_ROOT}
COPY . .
RUN go mod download

RUN apk update && \
    apk add --no-cache git gcc libc-dev && \
    git config --global http.postBuffer 524288000 && \
    rm -rf /var/cache/apk/*
