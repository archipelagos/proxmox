#!/bin/bash

# INFO: Load secrets into env.
source /run/secrets/actions-runner-creds

trap 'cleanup' SIGINT

# INFO: Clean after runner.
cleanup() {
    ./config.sh \
        remove \
        --token \
        ${RUNNER_TOKEN}
}

# INFO: Perform runner configuration.
echo \
    -en \
    \\n${RUNNER_NAME}\\n${RUNNER_EXTRA_TAGS}\\n${RUNNER_WORKDIR}\\n | \
    ./config.sh \
        --url \
        ${RUNNER_URL} \
        --token \
        ${RUNNER_TOKEN}

# INFO: Launch runner.
./run.sh