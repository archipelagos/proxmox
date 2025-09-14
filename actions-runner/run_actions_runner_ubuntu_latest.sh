#!/bin/bash

# INFO: Build containers.
docker \
    -v \
    -count=1 \
    compose \
    build

# INFO: Launch containers.
docker \
    compose \
    up \
    --remove-orphans