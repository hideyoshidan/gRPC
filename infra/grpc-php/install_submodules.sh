#!/bin/sh
cd ./infra/grpc-php/plugins

git clone -b v1.27.2 https://github.com/grpc/grpc

cd grpc

git submodule update --init

docker exec grpc-php bash -c 'cd /plugins/grpc; make grpc_php_plugin'