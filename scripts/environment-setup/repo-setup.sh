#!/bin/bash
# Copyright 2022 Google LLC
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
[[ ":$PATH:" != *":${GOPATH}/bin:"* ]] && echo "PATH=\"${GOPATH}/bin:\$PATH\"" >> ~/.profile
source ~/.profile
REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

# Downloads and configures kustomize
VERSION=$(source scripts/shared-vars-public.sh && echo "${KUSTOMIZE_VERSION}")
INSTALL_DIR=/usr/local/kustomize/bin
curl -O -L https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv${VERSION}/kustomize_v${VERSION}_${GOOS}_${GOARCH}.tar.gz
tar -zxvf kustomize_v${VERSION}_${GOOS}_${GOARCH}.tar.gz
chmod a+x kustomize
sudo mkdir -p "${INSTALL_DIR}"
sudo mv kustomize ${INSTALL_DIR}/kustomize
[[ ":$PATH:" != *":${INSTALL_DIR}:"* ]] && echo "PATH=\"${INSTALL_DIR}:\$PATH\"" >> ~/.profile
source ~/.profile
rm -f kustomize_v${VERSION}_${GOOS}_${GOARCH}.tar.gz

# Checks to make sure you have all the tools you need
kustomize version

GREEN='\033[0;32m'
NC='\033[0m'
echo -e "${GREEN}Config Connector REPO SETUP SUCCESSFUL${NC}"
