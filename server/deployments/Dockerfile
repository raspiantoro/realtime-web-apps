FROM golang:1.12-alpine3.10 AS build-stage

ARG ENVIRONMENT
ARG GOPROXY

LABEL maintainer="Mario Raspiantoro<raspiantoro@gmail.com>"
LABEL REPO="https://github.com/raspiantoro/realtime-web-apps"

RUN apk add git && apk add --update make

WORKDIR /opt/traffic-count

COPY ./server .

RUN go get -v ./...

RUN go build -o bin/trafficcount


FROM alpine

LABEL maintainer="Mario Raspiantoro<raspiantoro@gmail.com>"
LABEL REPO="https://github.com/raspiantoro/realtime-web-apps"

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

COPY --from=build-stage /opt/traffic-count/bin/trafficcount /opt/traffic-count/bin/

WORKDIR /opt/traffic-count/bin

RUN adduser -D -S trafficcount
USER trafficcount

CMD ["/opt/traffic-count/bin/trafficcount"]