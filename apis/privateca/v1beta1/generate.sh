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

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

./generate-proto.sh

go run . generate-types \
  --service google.cloud.security.privateca.v1 \
  --api-version privateca.cnrm.cloud.google.com/v1beta1  \
  --resource PrivateCACAPool:CaPool

# Remove duplicated structs from types.generated.go (overridden in capool_types.go)
# Using sed to delete from annotation to closing brace.
TYPES_FILE="${REPO_ROOT}/apis/privateca/v1beta1/types.generated.go"
sed -i '/\/\/ \+kcc:proto=google\.cloud\.security\.privateca\.v1\.CaPool\.IssuancePolicy/,/^}/d' "${TYPES_FILE}"
sed -i '/\/\/ \+kcc:proto=google\.cloud\.security\.privateca\.v1\.CaPool\.PublishingOptions/,/^}/d' "${TYPES_FILE}"
sed -i '/\/\/ \+kcc:proto=google\.cloud\.security\.privateca\.v1\.X509Extension/,/^}/d' "${TYPES_FILE}"
sed -i '/\/\/ \+kcc:proto=google\.cloud\.security\.privateca\.v1\.ObjectId/,/^}/d' "${TYPES_FILE}"
# Use $ to match end of line to avoid prefix matching CaOptions
sed -i '/\/\/ \+kcc:proto=google\.cloud\.security\.privateca\.v1\.X509Parameters$/,/^}/d' "${TYPES_FILE}"
sed -i '/\/\/ \+kcc:proto=google\.cloud\.security\.privateca\.v1\.X509Parameters\.CaOptions/,/^}/d' "${TYPES_FILE}"
sed -i '/\/\/ \+kcc:proto=google\.cloud\.security\.privateca\.v1\.X509Parameters\.NameConstraints/,/^}/d' "${TYPES_FILE}"

go run . generate-mapper \
  --service google.cloud.security.privateca.v1 \
  --api-version privateca.cnrm.cloud.google.com/v1beta1

cd ${REPO_ROOT}

# Fix Value conversion (byte[] <-> string mismatch) in mapper.generated.go
# FromProto: generated code uses direct.LazyPtr(in.GetValue()) or in.GetValue(), we want string(in.GetValue())
sed -i 's/out.Value = direct.LazyPtr(in.GetValue())/out.Value = string(in.GetValue())/g' pkg/controller/direct/privateca/mapper.generated.go
sed -i 's/out.Value = in.GetValue()/out.Value = string(in.GetValue())/g' pkg/controller/direct/privateca/mapper.generated.go

# ToProto: generated code uses in.Value, we want []byte(in.Value)
sed -i 's/out.Value = in.Value/out.Value = []byte(in.Value)/g' pkg/controller/direct/privateca/mapper.generated.go

# Remove ZeroMaxIssuerPathLength mapping if generated (SDK mismatch)
sed -i '/ZeroMaxIssuerPathLength/d' pkg/controller/direct/privateca/mapper.generated.go

dev/tasks/generate-crds

# Format files
go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w  pkg/controller/direct/privateca/
go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w  apis/privateca/v1beta1/
