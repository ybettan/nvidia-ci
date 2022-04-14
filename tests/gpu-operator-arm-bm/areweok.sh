#!/bin/bash

set -e
set -x

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}
. ${SCRIPT_DIR}/config.sh

status=$(${oc_command} get clusterpolicy gpu-cluster-policy -o yaml | yq '.status.state')

${oc_command} wait --for=condition=ready pod -l app=nvidia-dcgm  --timeout=5m

if [ "${status}" != "ready" ]; then
    echo "GPU Operator not ready"
    exit -2
fi
