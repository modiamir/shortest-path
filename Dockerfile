FROM golang:1.15-alpine

WORKDIR /app
COPY . /app

RUN go build main.go
ENTRYPOINT "/app/main"