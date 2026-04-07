#!/bin/bash

# Copyright 2019 The Kubernetes Authors.
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

set -o errexit
set -o nounset
set -o pipefail

CI_ROOT=$(dirname "${BASH_SOURCE}")
REPO_ROOT=$(dirname "${BASH_SOURCE}")/../..

source "${CI_ROOT}/fetch_kubebuilder_release_bin.sh"

cd "${REPO_ROOT}"
# Make sure REPO_ROOT is an absolute path
REPO_ROOT=$(pwd)

# Ensure we run with a known version of kubectl
if [[ ! -f "bin/kubectl" ]]; then
  echo "Downloading kubectl to bin/kubectl"
  mkdir -p bin/
  curl -L -o bin/kubectl https://dl.k8s.io/release/v1.32.2/bin/linux/amd64/kubectl
fi
chmod +x bin/kubectl
export PATH="${REPO_ROOT}/bin:$PATH"
echo "kubectl version is $(kubectl version --client)"

dev/test
