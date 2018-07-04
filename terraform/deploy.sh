#!/bin/bash
function error_exit {
    echo "$1" 1>&2
    exit $2
}

function terraform_run {
    terraform init
    if [ $? -ne 0 ]; then
        error_exit "ERROR: Cannot initiate terraform." 1
    fi
    terraform plan
    if [ $? -ne 0 ]; then
        error_exit "ERROR: Cannot do terraform plan." 2
    fi
    terraform apply -parallelism=20 -lock-timeout=50s -auto-approve
    if [ $? -ne 0 ]; then
        error_exit "ERROR: Cannot apply terraform changes." 3
    fi
}

terraform init 2>/dev/null

varpath=$(find ./ -name "00-variables.tf" | grep modules)
invpath=$(find ./ -name "09-inventory.tf" | grep modules)

cp ./tmp/00-variables.tf $varpath && cp ./tmp/09-inventory.tf $invpath

terraform_run