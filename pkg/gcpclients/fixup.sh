#!/bin/bash

# Copyright 2026 Google LLC
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

REPO_ROOT=$(git rev-parse --show-toplevel)
GCPCLIENTS=${REPO_ROOT}/pkg/gcpclients
cd ${GCPCLIENTS}/.build/googleapis

# We need to rewrite various proto import paths, to avoid conflicts with the official Go libraries.
# We rename "google/" to "gcpclients/" in the proto paths and packages.

rm -rf gcpclients/ && mkdir -p gcpclients/

mv google/cloud/ gcpclients/
mv google/firestore/ gcpclients/

# Also move google/api/apikeys if it existed (though not in Makefile yet)
if [ -d google/api/apikeys ]; then
  mkdir -p gcpclients/api
  mv google/api/apikeys/ gcpclients/api/
fi

cd gcpclients

# Rewrite import paths: google/cloud/ -> gcpclients/cloud/ etc
find . -type f -print0 | xargs -0 sed -i \
  -e "s@google/cloud/@gcpclients/cloud/@g" \
  -e "s@google\.cloud@gcpclients.cloud@g" \
  -e "s@google/firestore/@gcpclients/firestore/@g" \
  -e "s@google\.firestore@gcpclients.firestore@g" \
  -e "s@google/api/apikeys/@gcpclients/api/apikeys/@g" \
  -e "s@google\.api\.apikeys@gcpclients.api.apikeys@g"
