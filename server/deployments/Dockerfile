FROM golang:1.15.6-alpine3.12 AS build-stage

LABEL maintainer="Mario Raspiantoro<raspiantoro@gmail.com>"
LABEL REPO="https://github.com/raspiantoro/realtime-web-apps"

RUN apk add git && apk add --update make

WORKDIR /opt/streamserver

COPY ./server .

RUN go get -v ./...

RUN go build -o bin/streamserver


FROM alpine

LABEL maintainer="Mario Raspiantoro<raspiantoro@gmail.com>"
LABEL REPO="https://github.com/raspiantoro/realtime-web-apps"

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true 

COPY --from=build-stage /opt/streamserver/bin/streamserver /opt/streamserver/bin/

WORKDIR /opt/streamserver/bin

RUN adduser -D -S streamserver
USER streamserver

CMD ["/opt/streamserver/bin/streamserver"]
