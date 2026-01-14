#!/bin/bash

# INFO: Enable secure guards.
set -u
set -e

# INFO: Clean-up after previous tests.
go \
    clean \
    -testcache

# INFO: Perform fresh tests.
go \
    test \
    -v \
    ./...