#!/bin/bash
master_fqdn=$1
userName=$2
privateKeyLocation=$3
openshiftVersion=$4
NUM_MASTER_NODE=$5
NUM_DB_NODE=$6
NUM_ROUTER_NODE=$7

declare -i mastercount=-1
declare -i dbcount=-1
declare -i routercount=-1

#Simply check number of db nodes
if [ "$NUM_DB_NODE" -lt 1 ] || [ "$NUM_MASTER_NODE" -lt 1 ] || [ "$NUM_ROUTER_NODE" -lt 1 ]; then
    echo "ERROR: The number of nodes is not correct."
    exit 220
fi

((NUM_DB_NODE--))
((NUM_MASTER_NODE--))
((NUM_ROUTER_NODE--))


function export_master_nodes(){
    ((mastercount++))
    if [ $mastercount -le $NUM_MASTER_NODE ]; then
        if [ "$1" -eq 1 ]; then
            if [ $mastercount -eq 0 ]; then
                export MASTER_HOST_1="master$mastercount"
            else
                export MASTER_HOST_1="$MASTER_HOST_1"$'\n'"master$mastercount"
            fi
            
        else
            if [ $mastercount -eq 0 ]; then
                export MASTER_HOST_2="master$mastercount openshift_schedulable=false"
            else
                export MASTER_HOST_2="$MASTER_HOST_2"$'\n'"master$mastercount openshift_schedulable=false"
            fi
        fi
        export_master_nodes $1
    fi
}

function export_router_nodes(){
    ((routercount++))
    if [ $routercount -le $NUM_ROUTER_NODE ]; then
        if [ $routercount -eq 0 ]; then
            export ROUTER_HOST="router$routercount openshift_node_labels=\"{'region': 'infra', 'purpose': 'infra'}\""
        else
            export ROUTER_HOST="$ROUTER_HOST"$'\n'"router$routercount openshift_node_labels=\"{'region': 'infra', 'purpose': 'infra'}\""
        fi
        export_router_nodes
    fi
}

function export_db_nodes(){
    ((dbcount++))
    if [ $dbcount -le $NUM_DB_NODE ]; then
        if [ $dbcount -eq 0 ]; then
            export DB_HOST="db$dbcount openshift_node_labels=\"{'region': 'primary', 'purpose': '"db$dbcount"'}\""
        else
            export DB_HOST="$DB_HOST"$'\n'"db$dbcount openshift_node_labels=\"{'region': 'primary', 'purpose': '"db$dbcount"'}\""
        fi
        export_db_nodes
    fi   
}

cd /home/$userName

# Create Ansible Hosts File
echo $(date) " - Create Ansible Hosts file"

#Generate master node host configuration
export_master_nodes 1
#Reset variable
declare -i mastercount=-1
export_master_nodes 2
#Generate router node host configuration
export_router_nodes
#Generate db node host configuration
export_db_nodes

cat > /home/$userName/hosts <<EOF
[OSEv3:children]
masters
etcd
nodes
 
# Set variables common for all OSEv3 hosts
[OSEv3:vars]
ansible_ssh_user=<ssh_user>
ansible_become=true
deployment_type=origin
ansible_ssh_private_key_file=<path_to_ssh_key>
osm_cluster_network_cidr=10.128.0.0/14
openshift_master_identity_providers=[{'name': 'htpasswd_auth', 'login': 'true', 'challenge': 'true','kind': 'HTPasswdPasswordIdentityProvider','filename': '/etc/origin/master/htpasswd'}]
osm_host_subnet_length=12
openshift_disable_check=disk_availability,docker_storage,memory_availability,package_version
openshift_master_cluster_method=native
openshift_master_cluster_hostname=<master_dns_endpoint>
openshift_master_cluster_public_hostname=<master_dns_endpoint>

# enable ntp on masters to ensure proper failover
openshift_clock_enabled=true

# enables masters rolling restart with full system
openshift_rolling_restart_mode=true

#enables API service auditing
openshift_master_audit_config={"enabled": true, "auditFilePath": "/var/log/audit-ocp.log", "maximumFileRetentionDays": 14, "maximumFileSizeMegabytes": 200, "maximumRetainedFiles": 10}

#set enterprise sdn plugin
os_sdn_network_plugin_name=redhat/openshift-ovs-multitenant

#set firewall configuration
os_firewall_use_firewalld=true

#add Artifactory registry
openshift_docker_additional_registries=docker.artifactory.zcloudcentral.com
openshift_docker_insecure_registries=docker.artifactory.zcloudcentral.com

#don't add repo to /etc/yum.repos.d
openshift_enable_origin_repo=false

# host group for masters
[masters]
$MASTER_HOST_1

# host group for etcd
[etcd]
$MASTER_HOST_1

# host group for nodes, includes region info
[nodes]
$MASTER_HOST_2
$DB_HOST
$ROUTER_HOST
app0 openshift_node_labels="{'region': 'primary', 'purpose': 'app'}"
app1 openshift_node_labels="{'region': 'primary', 'purpose': 'app'}"
sftp openshift_node_labels="{'region': 'primary', 'purpose': 'sftp-etl'}"
EOF

echo "Updating ansible hosts file..."

sed -i -e "s/<master_dns_endpoint>/${master_fqdn}/g" -e "s/<master1 ip or dns name>/${master1}/g" -e "s/<master2 ip or dns name>/${master2}/g" -e "s/<db1 ip or dns name>/${db1}/g" -e "s/<db2 ip or dns name>/${db2}/g" -e "s/<db3 ip or dns name>/${db3}/g" -e "s/<appnode1 ip or dns name>/${appNode1}/g" -e "s/<appnode2 ip or dns name>/${appNode2}/g" -e "s/<sftpnode ip or dns name>/${sftpNode}/g" -e "s/<router1 ip or dns name>/${routerNode1}/g" -e "s/<router2 ip or dns name>/${routerNode2}/g" -e "s/<ssh_user>/${userName}/g" -e "s:<path_to_ssh_key>:${privateKeyLocation}:g" /home/$userName/hosts

sudo cp /home/$userName/hosts /etc/ansible/hosts
