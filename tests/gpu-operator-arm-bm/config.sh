#!/bin/bash

set -e

config_file=config.json

bundle=$(cat ${config_file} | jq -r '.bundle_url')
driver=$(cat ${config_file} | jq -r '.driver_image')

