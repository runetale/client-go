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

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/.env .env
COPY --from=builder /app/cmd/server/migrations migrations

EXPOSE 443
ENTRYPOINT ["./server"]
