#!/bin/bash
oc delete configmap database-conf
oc delete is pgprimary
oc delete service pgprimary
oc delete dc pgprimary