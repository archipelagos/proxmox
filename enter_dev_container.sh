#!/bin/bash

CONTAINER_NAME=$(basename $(pwd))-dev

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
    -v /etc/passwd:/etc/passwd:ro \
    -v /etc/group:/etc/group:ro \
    -v ./ssh:/home/${USER}/.ssh:ro \
    -v $(pwd):/home/${USER} \
    -w /home/${USER} \
    -u $(id -u):$(id -g) \
    -t \
    ${CONTAINER_NAME} \
    bash