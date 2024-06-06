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

# Script to set up golang on cloudtop
set -o errexit

# This will pull the current go version used in the repository
REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}
GO_VERSION=`cat go.mod | grep -P -o "(?<=^go ).*"`
sudo apt-get install wget
wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz
sudo tar -xvzf go${GO_VERSION}.linux-amd64.tar.gz -C /usr/local
rm go${GO_VERSION}.linux-amd64.tar.gz

echo 'PATH=/usr/local/go/bin:$PATH' >> ~/.profile
source ~/.profile
go version

sudo apt install jq

echo "export GOPATH=${HOME}/go" >> ~/.profile
source ~/.profile

GREEN='\033[0;32m'
NC='\033[0m'
echo -e "${GREEN}GOLANG SETUP SUCCESSFUL${NC}"
