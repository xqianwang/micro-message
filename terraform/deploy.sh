#!/bin/bash
set -o pipefail

function error_exit {
    echo "$1" 1>&2
    exit $2
}

#function to run terraform application
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
#first initiate to download terraform dependencies: terraform aws   
terraform init 2>/dev/null

varPath=$(find ./ -name "00-variables.tf" | grep modules)
invPath=$(find ./ -name "09-inventory.tf" | grep modules)

cp ./tmp/00-variables.tf $varPath && cp ./tmp/09-inventory.tf $invPath

terraform_run

inventoryPath=$(find ./ -name "inventory.cfg")

#Must have ssh keys 
eval `ssh-agent -s`
ssh-add ~/.ssh/id_rsa
scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null $inventoryPath ec2-user@$(terraform output bastion-public_dns):~
#install openshift
echo "Installing openshift now"
cat ./scripts/openshiftInstall.sh | ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -A ec2-user@$(terraform output bastion-public_dns)
if [ $? -eq 0 ]; then
    echo "Openshift installed successfully."
else 
    error_exit "Failed to install openshift" 4
fi

echo "Configuring Openshift deployment now"
MASTER_HOST=$(terraform output master-private_dns)
NODE1_HOST=$(terraform output node1-private_dns)
NODE2_HOST=$(terraform output node2-private_dns)
cp ../envs.sh ./tmp && sed -i "s/masterHost/$MASTER_HOST/g;s/node1Host/$NODE1_HOST/g;s/node2Host/$NODE2_HOST/g" ./tmp/envs.sh

scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ./scripts/openshiftConfig.sh ~/.ssh/id_rsa ./tmp/envs.sh ec2-user@$(terraform output bastion-public_dns):~
ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ec2-user@$(terraform output bastion-public_dns) 'mv /home/ec2-user/id_rsa /home/ec2-user/.ssh/ && chmod 0600 /home/ec2-user/.ssh/id_rsa'
ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ec2-user@$(terraform output bastion-public_dns) 'sh ~/openshiftConfig.sh | ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i ~/inventory.cfg ~/openShiftConfig.yml'

