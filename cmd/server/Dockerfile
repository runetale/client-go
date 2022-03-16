FROM golang:1.17-alpine as dev

# add sudo user
RUN apk add sudo

RUN adduser -D -u 1001 -s /bin/sh -G wheel alpine-sudoer
RUN echo '%wheel ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers

USER alpine-sudoer

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

EXPOSE 443

CMD ["sudo", "go", "run", "server.go"]
