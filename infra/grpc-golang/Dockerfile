# new version of golang
FROM golang:latest

# make working dir
RUN mkdir /grpc

# install packages to use gRPC
RUN apt-get update \
    && apt-get install -y vim zip \
    && apt install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
RUN export PATH="$PATH:$(go env GOPATH)/bin"

# setting working dir "/grpc"
WORKDIR /grpc
EXPOSE 1323