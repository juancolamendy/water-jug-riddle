FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git && apk add ca-certificates && rm -rf /var/cache/apk/* && mkdir -p $GOPATH/src/github.com/juancolamendy/water-jug-riddle

WORKDIR /app
COPY . .

RUN go build -v -i -o app.bin cmd/main.go

FROM alpine:latest
MAINTAINER JC

WORKDIR /app
COPY --from=builder /app .
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

EXPOSE 3001

ENTRYPOINT ["./app.bin"]