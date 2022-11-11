FROM golang:1.19-alpine AS builder

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

WORKDIR /github.com/paw1a/transaction-service/
COPY . /github.com/paw1a/transaction-service/

RUN go mod download

RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /github.com/paw1a/transaction-service/.bin/ ./.bin/
COPY --from=builder /github.com/paw1a/transaction-service/migrations/ ./migrations/

EXPOSE 8080

CMD ["./.bin/app"]
