#!/bin/bash

set -e
set -x

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}
. ${SCRIPT_DIR}/config.sh

subscription="$(${oc_command} get subscriptions.operators.coreos.com -o json | jq -r '.items[].metadata.name' | grep gpu-operator)" || exit 0
if [ -z ${subscription} ]; then
    echo Subscription is empty.
    exit -1
fi
currentCSV=$(${oc_command} get subscriptions.operators.coreos.com ${subscription} -o json | jq -r '.status.currentCSV')
if [ -z ${currentCSV} ]; then
    echo CSV is empty.
    exit -1
fi

echo Subscription: ${subscription}
echo CSV: ${currentCSV}

${oc_command} delete subscription ${subscription}
${oc_command} delete clusterserviceversion ${currentCSV}
${oc_command} delete crd clusterpolicies.nvidia.com

