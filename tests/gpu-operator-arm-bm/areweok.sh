#!/bin/bash

set -e
set -x

status=$(oc get clusterpolicy gpu-cluster-policy -o yaml | yq '.status.state')

oc wait --for=condition=ready pod -l app=nvidia-dcgm  --timeout=5m

if [ "${status}" != "ready" ]; then
    echo "GPU Operator not ready"
    exit -2
fi
