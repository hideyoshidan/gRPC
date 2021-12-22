FROM php:8.1-fpm

RUN mkdir /proto

# install packages
RUN apt-get update \
    && apt-get install -y \
    git \
    zip \
    unzip \
    vim \
    libpq-dev \
    autoconf \
    automake \
    libtool \
    curl \
    make \
    g++ \
    libprotobuf-dev \
    libprotoc-dev \
    protobuf-compiler 

# install need things for php
RUN docker-php-ext-install pdo_mysql pdo_pgsql

# grpc
RUN apt-get install libz-dev
RUN pecl install grpc

# install gRPC PHP Protoc Plugin
RUN mkdir /plugin
RUN cd /plugin

RUN git clone -b v1.27.2 https://github.com/grpc/grpc
RUN cd grpc
RUN git submodule update --init
RUN mkdir grpc_php_plugin

# install composer
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
RUN php -r "if (hash_file('sha384', 'composer-setup.php') === '906a84df04cea2aa72f40b5f787e49f22d4c2f19492ac310e8cba5b96ac8b64115ac402c8cd292b8a03482574915d1a8') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;"
RUN php composer-setup.php
RUN php -r "unlink('composer-setup.php');"

# path
RUN mv composer.phar /usr/local/bin/composer

# set WORKDIR
WORKDIR /grpc