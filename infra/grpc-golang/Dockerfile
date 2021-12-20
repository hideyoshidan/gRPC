# new version of golang
FROM golang:latest

# make working dir
RUN mkdir /grpc

# install packages to use gRPC
RUN apt-get update \
    && apt-get install -y vim zip \
    && apt install -y protobuf-compiler

# setting working dir "/grpc"
WORKDIR /grpc
EXPOSE 1323