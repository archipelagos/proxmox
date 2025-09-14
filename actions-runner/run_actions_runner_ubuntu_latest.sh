#!/bin/bash

CONTAINER_NAME=$(basename $(pwd))

UID=1000
GID=1000

docker \
    build \
    --cache-from \
    ${CONTAINER_NAME} \
    -t \
    ${CONTAINER_NAME} \
    .
    
docker \
    run \
    -it \
    --rm \
    -h ${CONTAINER_NAME} \
    -w /home/ubuntu \
    -u ${UID}:${GID} \
    -t \
    ${CONTAINER_NAME} \
    bash