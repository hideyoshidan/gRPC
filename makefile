VERSION=1.0
CONTAINER_GO=grpc-golang
CONTAINER_PHP=grpc-php
CONTAINER_NGINX=grpc-webserver

#########################################
# commands for docker containers
######################################### 
# build docker containers
build:
	docker compsoe -f ./docker-compose.yml build --no-cache

# make docker containers
up:
	docker compose -f ./docker-compose.yml up -d

# start docker containers
start:
	docker compose -f ./docker-compose.yml start

# stop docker containers
stop:
	docker compose -f ./docker-compose.yml stop

# restart doker containers
restart:
	docker compose -f ./docker-compose.yml restart

#########################################
# into containers
######################################### 
# into golang containers
go:
	docker compose -f ./docker-compose.yml exec ${CONTAINER_GO} bash

# into php containers
php:
	docker compose -f ./docker-compose.yml exec ${CONTAINER_PHP} bash

# into nginx containers
nginx:
	docker compose -f ./docker-compose.yml exec ${CONTAINER_NGINX} bash


#########################################
# anything 
######################################### 
# run golang
go-run:
	docker exec ${CONTAINER_GO} bash -c 'go run main.go'

# reload nginx
nginx-reload:
	docker exec ${CONTAINER_NGINX} bash -c 'nginx -s reload'