#!/usr/bin/env bash
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

# This script documents how the protobuf files in pkg/controller/direct/aiplatform/pb were compiled
# from the pinned googleapis repository SHA e7513c828b8db536fa12fc5493d541f323aedf90.
#
# Package rename:
# The proto package for OnlineEvaluator is compiled as package google.cloud.aiplatform.v1beta1.onlineevaluator
# (Mgoogle/cloud/aiplatform/v1beta1/online_evaluator.proto=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb)
# to avoid GlobalTypes/GlobalFiles protobuf registration collisions with GAPIC apiv1beta1/aiplatformpb.

REPO_ROOT="$(git rev-parse --show-toplevel)"
PB_DIR="${REPO_ROOT}/pkg/controller/direct/aiplatform/pb"

echo "Protos in ${PB_DIR} were compiled from googleapis commit e7513c828b8db536fa12fc5493d541f323aedf90."
