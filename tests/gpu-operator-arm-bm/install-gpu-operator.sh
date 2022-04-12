#!/bin/bash

. config.sh

operator-sdk run bundle --timeout=1m -n nvidia-gpu-operator --install-mode OwnNamespace "${bundle}"
