#!/bin/bash
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


set -o errexit
set -o nounset
set -o pipefail

# Get the top-level directory of the git repository
TOP_LEVEL=$(git rev-parse --show-toplevel)

if kubectl wait --for=condition=Ready storagebuckets/storagebucket-apply -n storagebucket-apply --timeout=30s; then
    exit 0
else
    exit 1
fi 


if gcloud storage buckets describe gs://storagebucket-apply; then
    exit 0
else
    exit 1
fi 