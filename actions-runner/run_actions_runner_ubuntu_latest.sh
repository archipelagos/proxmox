#!/bin/bash

# INFO: Build containers.
COMPOSE_BAKE=true \
docker \
    compose \
    build

# INFO: Launch containers.
docker \
    compose \
    down \
    --remove-orphans

# INFO: Launch containers.
docker \
    compose \
    up \
    --detach
