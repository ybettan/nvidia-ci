#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}
. ${SCRIPT_DIR}/config.sh

set -e
set -x

if [ "${channel}" = "stable" ]; then
    #todo: better differentiate between types of deployments (catalogsource bundle, marketplace etc)
    currentCSV=$(oc get packagemanifests/gpu-operator-certified -n openshift-marketplace -ojson | jq -r '.status.channels[] | select(.name == "stable") | .currentCSV')
    cp subscription.yaml _subscription.yaml
    echo "  startingCSV: ${currentCSV}" >> _subscription.yaml
    echo "  channel: ${channel}" >> _subscription.yaml
    oc apply -f operatorgroup.yaml
    oc apply -f _subscription.yaml
else
    operator-sdk run bundle --timeout=1m -n nvidia-gpu-operator --install-mode OwnNamespace "${bundle}"
fi

#todo: write a proper wait for the correct resource
sleep 60
oc wait --for=condition=ready pod -l app=gpu-operator --timeout=5m

while [ "$(oc get -n nvidia-gpu-operator ClusterServiceVersion ${currentCSV} -o json | jq -r '.status.phase')" != "Succeeded" ]; do sleep 10; done

oc get csv -n nvidia-gpu-operator "${currentCSV}" -ojsonpath={.metadata.annotations.alm-examples} | jq .[0] | oc apply -f -

#todo: proper wait
sleep 60
