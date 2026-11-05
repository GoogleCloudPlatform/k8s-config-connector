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

REPO_ROOT="$(git rev-parse --show-toplevel)"
PB_DIR="${REPO_ROOT}/pkg/controller/direct/aiplatform/pb"
GOOGLEAPIS_COMMIT="e7513c828b8db536fa12fc5493d541f323aedf90"

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "${TMP_DIR}"' EXIT

GOOGLEAPIS_DIR="${TMP_DIR}/googleapis"
git clone --depth 1 https://github.com/googleapis/googleapis.git "${GOOGLEAPIS_DIR}"
(cd "${GOOGLEAPIS_DIR}" && git fetch --depth 1 origin "${GOOGLEAPIS_COMMIT}" && git checkout "${GOOGLEAPIS_COMMIT}")

GOBIN="${TMP_DIR}/bin"
mkdir -p "${GOBIN}"
export PATH="${GOBIN}:${PATH}"

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

PROTOC_ZIP="${TMP_DIR}/protoc.zip"
wget -q -O "${PROTOC_ZIP}" https://github.com/protocolbuffers/protobuf/releases/download/v3.12.4/protoc-3.12.4-linux-x86_64.zip
python3 -c "import zipfile, sys; zipfile.ZipFile(sys.argv[1]).extractall(sys.argv[2])" "${PROTOC_ZIP}" "${TMP_DIR}/protoc"
PROTOC="${TMP_DIR}/protoc/bin/protoc"
chmod +x "${PROTOC}"

# Prepare proto source directory with package rename to google.cloud.aiplatform.v1beta1.onlineevaluator
# to prevent GlobalTypes / GlobalFiles collisions with GAPIC apiv1beta1/aiplatformpb.
PROTO_WORK_DIR="${TMP_DIR}/work"
mkdir -p "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1"
cp -r "${GOOGLEAPIS_DIR}/google/cloud/aiplatform/v1beta1"/* "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/"

# Rename evaluation_*.proto to online_evaluator_evaluation_*.proto to avoid filename clashes
mv "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/evaluation_service.proto" "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/online_evaluator_evaluation_service.proto"
mv "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/evaluation_rubric.proto" "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/online_evaluator_evaluation_rubric.proto"
mv "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/evaluation_agent_data.proto" "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/online_evaluator_evaluation_agent_data.proto"

# Rewrite imports and package name
TARGET_PROTOS=(
  "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/online_evaluator.proto"
  "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/online_evaluator_service.proto"
  "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/online_evaluator_evaluation_service.proto"
  "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/online_evaluator_evaluation_rubric.proto"
  "${PROTO_WORK_DIR}/google/cloud/aiplatform/v1beta1/online_evaluator_evaluation_agent_data.proto"
)

for f in "${TARGET_PROTOS[@]}"; do
  sed -i 's|google/cloud/aiplatform/v1beta1/evaluation_|google/cloud/aiplatform/v1beta1/online_evaluator_evaluation_|g' "${f}"
  sed -i 's|package google.cloud.aiplatform.v1beta1;|package google.cloud.aiplatform.v1beta1.onlineevaluator;|g' "${f}"
done

"${PROTOC}" -I "${PROTO_WORK_DIR}" -I "${GOOGLEAPIS_DIR}" \
  --experimental_allow_proto3_optional \
  --go_out="${REPO_ROOT}" \
  --go_opt=module=github.com/GoogleCloudPlatform/k8s-config-connector \
  --go_opt="Mgoogle/cloud/aiplatform/v1beta1/online_evaluator.proto=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb;pb" \
  --go_opt="Mgoogle/cloud/aiplatform/v1beta1/online_evaluator_service.proto=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb;pb" \
  --go_opt="Mgoogle/cloud/aiplatform/v1beta1/online_evaluator_evaluation_service.proto=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb;pb" \
  --go_opt="Mgoogle/cloud/aiplatform/v1beta1/online_evaluator_evaluation_rubric.proto=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb;pb" \
  --go_opt="Mgoogle/cloud/aiplatform/v1beta1/online_evaluator_evaluation_agent_data.proto=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb;pb" \
  --go-grpc_out="${REPO_ROOT}" \
  --go-grpc_opt=module=github.com/GoogleCloudPlatform/k8s-config-connector \
  --go-grpc_opt="Mgoogle/cloud/aiplatform/v1beta1/online_evaluator.proto=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb;pb" \
  --go-grpc_opt="Mgoogle/cloud/aiplatform/v1beta1/online_evaluator_service.proto=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb;pb" \
  --go-grpc_opt="Mgoogle/cloud/aiplatform/v1beta1/online_evaluator_evaluation_service.proto=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb;pb" \
  --go-grpc_opt="Mgoogle/cloud/aiplatform/v1beta1/online_evaluator_evaluation_rubric.proto=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb;pb" \
  --go-grpc_opt="Mgoogle/cloud/aiplatform/v1beta1/online_evaluator_evaluation_agent_data.proto=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb;pb" \
  "${TARGET_PROTOS[@]}"

# Fix gRPC service full method names to retain the canonical GCP service name (/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/...)
sed -i 's|/google.cloud.aiplatform.v1beta1.onlineevaluator.OnlineEvaluatorService/|/google.cloud.aiplatform.v1beta1.OnlineEvaluatorService/|g' "${PB_DIR}/online_evaluator_service_grpc.pb.go"

go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w "${PB_DIR}/"

echo "Protos in ${PB_DIR} successfully generated."
