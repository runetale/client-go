# build stage
FROM golang:1.17-alpine as builder
ARG version 
RUN apk update
RUN apk add sudo build-base
WORKDIR /app
COPY . .
ENV GO111MODULE=auto

RUN GOOS=linux CGO_ENABLED=1 go build -o server cmd/server/server.go

# launch stage
FROM alpine:3.14.3

#ARG USER=default
#RUN adduser -D $USER \
#        && mkdir -p /etc/sudoers.d \
#        && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
#        && chmod 0440 /etc/sudoers.d/$USER
#
#USER $USER

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/.env .env
COPY --from=builder /app/cmd/server/migrations migrations

EXPOSE 443
#RUN ["chmod", "+x", "./server"]
ENTRYPOINT ["./server"]
