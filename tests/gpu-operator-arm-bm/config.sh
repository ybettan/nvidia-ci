#!/bin/bash

set -e
set -x

config_file=config.json

bundle=$(cat ${config_file} | jq -r '.bundle_url')
driver=$(cat ${config_file} | jq -r '.driver_image')
channel=$(cat ${config_file} | jq -r '.channel')

oc_command="oc --insecure-skip-tls-verify --kubeconfig /var/run/secrets/armsnokubeconfig"

ls -lah /var/run/secrets
cat /var/run/secrets/version
