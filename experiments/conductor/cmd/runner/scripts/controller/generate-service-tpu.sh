#!/bin/bash
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

export SERVICE=tpu
export HTTP_HOST=tpu.googleapis.com
export PROTO_PACKAGE=google.cloud.tpu.v1
export PROTO_SERVICE=google.cloud.tpu.v1.Tpu
export CRD_GROUP=${SERVICE}.cnrm.cloud.google.com

# Note tpu v2 library is not published!!!

export PROTO_MESSAGE=google.cloud.tpu.v1.Node
export PROTO_RESOURCE=Node
export RESOURCE=node 
export CRD_KIND=TPUNode
export CRD_VERSION=v1alpha1
./_generate-resource.sh || true
