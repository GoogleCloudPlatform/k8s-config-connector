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

# To regenerate the types, uncomment the lines below and run this script.
# NOTE: Requires googleapis SHA pointing to a version containing GKEHub Fleet (e.g. HEAD).
#
# REPO_ROOT="$(git rev-parse --show-toplevel)"
# cd ${REPO_ROOT}/dev/tools/controllerbuilder
#
# ./generate-proto.sh HEAD
#
# go run . generate-types \
#   --service google.cloud.gkehub.v1 \
#   --api-version gkehub.cnrm.cloud.google.com/v1alpha1 \
#   --resource GKEHubFleet:Fleet
