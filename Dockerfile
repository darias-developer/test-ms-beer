# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY config ./config
COPY controller ./controller
COPY data ./data
COPY external ./external
COPY handler ./handler
COPY external ./external
COPY middleware ./middleware
COPY model ./model
COPY router ./router
COPY service ./service
COPY util ./util

RUN go mod download
RUN go build -o /test-ms-beer

EXPOSE 8080

CMD [ "/test-ms-beer" ]