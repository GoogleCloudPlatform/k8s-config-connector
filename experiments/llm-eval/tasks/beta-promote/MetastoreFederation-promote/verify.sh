#!/bin/bash
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http{
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o nounset
set -o errexit
set -o pipefail

# This script is run from the task directory and must find the cloned repo.
REPO_DIR_NAME="MetastoreFederation-promote-metastore"
REPO_DIR=$(find . -mindepth 1 -maxdepth 1 -type d -name "${REPO_DIR_NAME}")

if [ -z "$REPO_DIR" ]; then
    echo "Error: Could not find the cloned repository directory: ${REPO_DIR_NAME}"
    exit 1
fi

# Change into the cloned repository to run verification checks.
cd "$REPO_DIR"
echo "Running verification checks inside $(pwd)"

CREATE_YAML_PATH="pkg/test/resourcefixture/testdata/basic/metastore/v1beta1/metastorefederation/create.yaml"
EXPECTED_APIVERSION="apiVersion: metastore.cnrm.cloud.google.com/v1beta1"

# Check if the create.yaml file exists
if [ ! -f "$CREATE_YAML_PATH" ]; then
    echo "Error: $CREATE_YAML_PATH not found."
    exit 1
fi

# Check if the apiVersion is correct
if ! grep -q "$EXPECTED_APIVERSION" "$CREATE_YAML_PATH"; then
    echo "Error: Did not find '$EXPECTED_APIVERSION' in $CREATE_YAML_PATH."
    exit 1
fi

echo "Verified apiVersion in $CREATE_YAML_PATH."

# Proceed with the mock comparison
hack/compare-mock fixtures/metastorefederation

echo "Verification successful."
exit 0
