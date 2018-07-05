#!/bin/bash
set -o pipefail

function error_exit {
    echo "$1" 1>&2
    exit $2
}
#get current dir
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
#get envs
source ./envvars.sh
#create configuration
oc create configmap database-conf --from-file=$DIR/setup.sql --from-file=$DIR/postgresql.conf --from-file=$DIR/pg_hba.conf
#deploy database container
echo $PG_LOCALE  $CCP_IMAGE_TAG  $PG_PRIMARY_USER $PG_USER $PG_DATABASE $NODE_LABEL
oc process -f $DIR/postgresql.yaml \
        -p PG_LOCALE=$PG_LOCALE \
        -p CCP_IMAGE_TAG=$CCP_IMAGE_TAG \
        -p PG_PRIMARY_USER=$PG_PRIMARY_USER \
        -p PG_USER=$PG_USER \
        -p PG_DATABASE=$PG_DATABASE \
        -p NODE_LABEL=$NODE_LABEL \
        -p PG_PASSWORD=$PG_PASSWORD \
        -p PG_ROOT_PASSWORD=$PG_ROOT_PASSWORD > postgresql.json

if [ $? -ne 0 ]; then
    error_exit "cannot process postgresql.yaml file" 1
fi

oc create -f postgresql.json