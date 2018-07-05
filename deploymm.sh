#!/bin/bash

set -o pipefail

HOST_NAME=$1

oc process -f message.yaml \
    -p HOST_NAME=$HOST_NAME