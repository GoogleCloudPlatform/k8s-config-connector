#!/usr/bin/env bash
# Copyright 2025 Google LLC
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


REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}

if [[ -z "${VERSION:-}" ]]; then
  echo "VERSION must be set"
  exit 1
fi

# find the STALE_VERSION
folder_path="${REPO_ROOT}/operator/channels/packages/configconnector"
if [ ! -d "$folder_path" ]; then
    echo "Error: Directory $folder_path does not exist"
    exit 1
fi

# List all directories in the specified path and store them in an array
versions=($(ls -d "$folder_path"/*/ 2>/dev/null | xargs -n 1 basename))

if [ ${#versions[@]} -eq 0 ]; then
    echo "Error: No version directories found in $folder_path"
    exit 1
fi

echo "Found versions:"
printf '%s\n' "${versions[@]}"

# Update the latest CRD and RBAC manifests to previous KCC versions
echo "Update the latest CRD and RBAC manifests to previous KCC versions"
for version in "${versions[@]}"; do
  # Skip if version is the same as VERSION
  if [ "${version}" = "${VERSION}" ]; then
    continue
  fi
  # Update autopilot manifests
  cp ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${VERSION}/crds.yaml ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${version}/crds.yaml
  cp ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${VERSION}/cluster/gcp-identity/0-cnrm-system.yaml ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${version}/cluster/gcp-identity/0-cnrm-system.yaml
  cp ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${VERSION}/cluster/workload-identity/0-cnrm-system.yaml ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${version}/cluster/workload-identity/0-cnrm-system.yaml
  cp ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${VERSION}/namespaced/0-cnrm-system.yaml ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${version}/namespaced/0-cnrm-system.yaml
  
  # Keep the controller version and update the rest of the manifest
  yq 'select(.kind == "StatefulSet")' ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${version}/namespaced/per-namespace-components.yaml > temp_statefulset.yaml
  cp ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${VERSION}/namespaced/per-namespace-components.yaml ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${version}/namespaced/per-namespace-components.yaml
  yq 'select(.kind != "StatefulSet")' -i ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${version}/namespaced/per-namespace-components.yaml 
  echo "---" >> ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${version}/namespaced/per-namespace-components.yaml
  cat temp_statefulset.yaml >> ${REPO_ROOT}/operator/autopilot-channels/packages/configconnector/${version}/namespaced/per-namespace-components.yaml
  rm temp_statefulset.yaml
  
  # Update standard manifests
  cp ${REPO_ROOT}/operator/channels/packages/configconnector/${VERSION}/crds.yaml ${REPO_ROOT}/operator/channels/packages/configconnector/${version}/crds.yaml
  cp ${REPO_ROOT}/operator/channels/packages/configconnector/${VERSION}/cluster/gcp-identity/0-cnrm-system.yaml ${REPO_ROOT}/operator/channels/packages/configconnector/${version}/cluster/gcp-identity/0-cnrm-system.yaml
  cp ${REPO_ROOT}/operator/channels/packages/configconnector/${VERSION}/cluster/workload-identity/0-cnrm-system.yaml ${REPO_ROOT}/operator/channels/packages/configconnector/${version}/cluster/workload-identity/0-cnrm-system.yaml
  cp ${REPO_ROOT}/operator/channels/packages/configconnector/${VERSION}/namespaced/0-cnrm-system.yaml ${REPO_ROOT}/operator/channels/packages/configconnector/${version}/namespaced/0-cnrm-system.yaml

    # Keep the controller version and update the rest of the manifest
  yq 'select(.kind == "StatefulSet")' ${REPO_ROOT}/operator/channels/packages/configconnector/${version}/namespaced/per-namespace-components.yaml > temp_statefulset.yaml
  cp ${REPO_ROOT}/operator/channels/packages/configconnector/${VERSION}/namespaced/per-namespace-components.yaml ${REPO_ROOT}/operator/channels/packages/configconnector/${version}/namespaced/per-namespace-components.yaml
  yq 'select(.kind != "StatefulSet")' -i ${REPO_ROOT}/operator/channels/packages/configconnector/${version}/namespaced/per-namespace-components.yaml 
  echo "---" >> ${REPO_ROOT}/operator/channels/packages/configconnector/${version}/namespaced/per-namespace-components.yaml
  cat temp_statefulset.yaml >> ${REPO_ROOT}/operator/channels/packages/configconnector/${version}/namespaced/per-namespace-components.yaml
  rm temp_statefulset.yaml

done

make fmt