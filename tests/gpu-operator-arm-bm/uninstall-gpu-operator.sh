#!/bin/bash

set -e
set -x

subscription="$(oc get subscriptions.operators.coreos.com -o json | jq -r '.items[].metadata.name' | grep gpu-operator)" || exit 0
if [ -z ${subscription} ]; then
    echo Subscription is empty.
    exit -1
fi
currentCSV=$(oc get subscriptions.operators.coreos.com ${subscription} -o json | jq -r '.status.currentCSV')
if [ -z ${currentCSV} ]; then
    echo CSV is empty.
    exit -1
fi

echo Subscription: ${subscription}
echo CSV: ${currentCSV}

oc delete subscription ${subscription}
oc delete clusterserviceversion ${currentCSV}
oc delete crd clusterpolicies.nvidia.com

