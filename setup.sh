#!/bin/bash

source ./init_env.sh

docker build --tag img-ws ./src/websocket

pushd ./src || exit 1
docker compose up -d --no-recreate
goose up
popd 

./feed.sh
