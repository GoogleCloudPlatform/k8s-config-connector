#!/bin/bash

set -e

# Ensure yq is installed
if ! command -v yq &> /dev/null
then
    echo "yq could not be found, please install it first"
    exit 1
fi

# Get all resource kinds from CRD files
CRD_KINDS=$(find config/crds -name "*.yaml" -print0 | xargs -0 yq '.spec.names.kind' | sort -u)

# Get all resource kinds from static_config.go
CONTROLLER_KINDS=$(awk -F'"' '/Kind: / {print $4}' pkg/controller/resourceconfig/static_config.go | sort -u)

# Find the difference
UNCONTROLLED_KINDS=$(comm -23 <(echo "${CRD_KINDS}") <(echo "${CONTROLLER_KINDS}"))

# Save the list to a file
echo "${UNCONTROLLED_KINDS}" > experiments/resources/uncontrolled_resources.txt

echo "Found the following uncontrolled resources:"
cat experiments/resources/uncontrolled_resources.txt