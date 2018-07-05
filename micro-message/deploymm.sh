#!/bin/bash

set -o pipefail

function error_exit {
    echo "$1" 1>&2
    exit $2
}

HOST_NAME=$1

oc process -f message.yaml \
    -p HOST_NAME=$HOST_NAME > message.json

if [ $? -ne 0 ]; then
    error_exit "cannot process message.yaml file" 1
fi

oc create -f message.json