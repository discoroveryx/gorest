FROM golang:1.19-alpine

WORKDIR /home/root/project

RUN apk add build-base

COPY ./src/go.mod ./
COPY ./src/go.sum ./

RUN go mod tidy
# RUN go mod download
RUN go install github.com/cosmtrek/air@v1.40.4

EXPOSE 8080
