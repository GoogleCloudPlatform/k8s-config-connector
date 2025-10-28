# Copyright 2024 Google LLC
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

#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"

if [[ -z "${VERSION:-}" ]]; then
  echo "VERSION must be set"
  exit 1
fi

TARGET_CRD_DIR=${REPO_ROOT}/crds
TEMP_DIR=$(mktemp -td sync-crds-folder.XXXXXXXX)
CRDS_FILE=${REPO_ROOT}/operator/channels/packages/configconnector/${VERSION}/crds.yaml

if [ ! -f "${CRDS_FILE}" ]; then
    echo "Error: CRDs file not found at ${CRDS_FILE}"
    exit 1
fi

# Parse CRDs into individual files
echo "Splitting manifest ${CRDS_FILE}"
cd ${REPO_ROOT} && go run -mod=readonly ${REPO_ROOT}/scripts/parse-crds/parse-crds.go \
    -file ${CRDS_FILE} \
    -output-dir ${TEMP_DIR}

echo "Writing split files to ${TARGET_CRD_DIR}"
rm -rf ${TARGET_CRD_DIR}
mkdir ${TARGET_CRD_DIR}
for filepath in ${TEMP_DIR}/*.yaml; do

    filename=$(basename -- ${filepath})
    echo "base file name: '${filename}'"

    # Robustly parse the metadata.name field to get the correct plural and group names.
    # This is the most reliable way to reconstruct the original filename.
    metadata_name=$(awk '/^metadata:/ { found=1 } found && /name:/ { print $2; exit }' "${filepath}")
    if [[ -z "${metadata_name}" ]]; then
      echo "Error: Could not parse metadata.name from '${filepath}'. Aborting."
      exit 1
    fi

    plural_from_file=$(echo "${metadata_name}" | cut -d'.' -f1)
    group_name=$(echo "${metadata_name}" | cut -d'.' -f2)

    if [[ -z "${plural_from_file}" || -z "${group_name}" ]]; then
      echo "Error: Failed to extract plural or group from metadata.name '${metadata_name}' in file '${filepath}'. Aborting."
      exit 1
    fi

    long_filename="apiextensions.k8s.io_v1_customresourcedefinition_${plural_from_file}.${group_name}.cnrm.cloud.google.com.yaml"
    original_filepath=${REPO_ROOT}/config/crds/resources/${long_filename}
    
    echo "Checking for file: '${original_filepath}'"
    if [ -f "${original_filepath}" ]; then
      commit_year=$(git log --reverse -n 1 --format=%ad --date=format:%Y -- "${original_filepath}")
      if [[ -z "${commit_year}" ]]; then
        echo "Error: Could not determine commit year for '${original_filepath}'. Aborting."
        exit 1
      fi
    else
      commit_year=$(date +%Y)
    fi
    echo "Processing ${filename} with commit year ${commit_year}"
    # Add license header to the file
    LICENSE_HEADER=$(cat << EOF
# Copyright ${commit_year} Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
EOF
)
    echo "${LICENSE_HEADER}" > ${TARGET_CRD_DIR}/${filename}
    echo "" >> ${TARGET_CRD_DIR}/${filename} # Add a empty line after the license header
    cat ${filepath} >> ${TARGET_CRD_DIR}/${filename}
done


rm -rf ${TEMP_DIR}
