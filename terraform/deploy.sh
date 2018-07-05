#!/bin/bash
set -o pipefail

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

varPath=$(find ./ -name "00-variables.tf" | grep modules)
invPath=$(find ./ -name "09-inventory.tf" | grep modules)

cp ./tmp/00-variables.tf $varPath && cp ./tmp/09-inventory.tf $invPath

terraform_run

inventoryPath=$(find ./ -name "inventory.cfg")

#Must have ssh keys 
eval `ssh-agent -s`
ssh-add ~/.ssh/id_rsa
scp $inventoryPath ec2-user@$(terraform output bastion-public_dns):~
#install openshift
echo "Installing openshift now"
cat ./scripts/openshiftInstall.sh | ssh -A ec2-user@$(terraform output bastion-public_dns)
if [ $? -eq 0 ]; then
    echo "Openshift installed successfully."
else 
    error_exit "Failed to install openshift" 4
fi

yum install httpd-tools