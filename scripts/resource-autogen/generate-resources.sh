#!/bin/bash
# Copyright 2023 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

while [[ $# -gt 0 ]]; do
  case "${1}" in
    --cnrm-repo)  LOCAL_CNRM_REPO="${2:-}"; shift ;;
    --mm-repo)    LOCAL_MM_REPO="${2:-}"; shift ;;
    *)            break
  esac
  shift
done

echo "Ensure you have set up magic-modules locally following the guide: https://googlecloudplatform.github.io/magic-modules/docs/getting-started/setup/ !"
echo "This script usually takes a few minutes to finish."
LOCAL_CNRM_REPO="${LOCAL_CNRM_REPO:-}"
LOCAL_MM_REPO="${LOCAL_MM_REPO:-}"

if [[ "${LOCAL_CNRM_REPO}" == "" ]] && [[ "${CNRM_REPO:-}" != "" ]]; then
    echo "Using environment variable CNRM_REPO for LOCAL_CNRM_REPO..."
    LOCAL_CNRM_REPO="${CNRM_REPO}"
fi
if [[ "${LOCAL_CNRM_REPO}" == "" ]]; then
    echo "Absolute path to local CNRM repository is required. You should either set it via --cnrm-repo flag, or using environment variable CNRM_REPO ('export CNRM_REPO=[path/to/cnrm/repo]')."
    exit 1
fi
if [[ "${LOCAL_MM_REPO}" == "" ]] && [[ "${MM_REPO:-}" != "" ]]; then
    echo "Using environment variable MM_REPO for LOCAL_MM_REPO..."
    LOCAL_MM_REPO="${MM_REPO}"
fi
if [[ "${LOCAL_MM_REPO}" == "" ]]; then
    echo "Absolute path to local magic-modules repository is required. You should either set it via --mm-repo flag, or using environment variable MM_REPO ('export MM_REPO=[path/to/magic-modules/repo]')."
    exit 1
fi

echo "Variables..."
echo "LOCAL_CNRM_REPO=${LOCAL_CNRM_REPO}"
echo "LOCAL_MM_REPO=${LOCAL_MM_REPO}"

echo "Generating service mappings using the KCC provider in magic-modules..."
CD_MM_COMMAND="cd ${LOCAL_MM_REPO}/mmv1"
echo "${CD_MM_COMMAND}"
eval "${CD_MM_COMMAND}"

# If generating resources based on services/products, the patches won't be
# applied correctly.
EXEC_COMMAND="bundle exec compiler -e terraform -a -f kcc -o ${LOCAL_CNRM_REPO}/scripts/resource-autogen/generated -v beta"
echo "${EXEC_COMMAND}"
eval "${EXEC_COMMAND}"

echo "Applying patches to the generated service mappings in cnrm..."
CD_AUTOGEN_COMMAND="cd ${LOCAL_CNRM_REPO}/scripts/resource-autogen"
echo "${CD_AUTOGEN_COMMAND}"
eval "${CD_AUTOGEN_COMMAND}"

APPLY_PATCHES_COMMAND="make apply-autogen-patches"
echo "${APPLY_PATCHES_COMMAND}"
eval "${APPLY_PATCHES_COMMAND}"

echo "Generating CRDs for allowlisted resources..."
CD_CNRM_COMMAND="cd ${LOCAL_CNRM_REPO}"
echo "${CD_CNRM_COMMAND}"
eval "${CD_CNRM_COMMAND}"

GENERATE_CRDS_COMMAND="make manifests"
echo "${GENERATE_CRDS_COMMAND}"
eval "${GENERATE_CRDS_COMMAND}"

echo "Converting TF samples to KCC testdata for allowlisted resources..."
CONVERT_SAMPLES_COMMAND="go run scripts/resource-autogen/main.go"
echo "${CONVERT_SAMPLES_COMMAND}"
eval "${CONVERT_SAMPLES_COMMAND}"
