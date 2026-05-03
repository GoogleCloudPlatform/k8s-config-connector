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

REPO_ROOT="$(git rev-parse --show-toplevel)"
source "${REPO_ROOT}/dev/tools/goimports.sh"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

./generate-proto.sh

go run . generate-types \
  --service google.cloud.redis.cluster.v1 \
  --api-version redis.cnrm.cloud.google.com/v1beta1 \
  --include-skipped-output \
  --resource RedisCluster:google.cloud.redis.cluster.v1.Cluster

sed -i 's/PSCConfig \*PSCConfig `json:"pscConfig,omitempty"`/PSCConfig \*PSCConfigObservedState `json:"pscConfig,omitempty"`/' ${REPO_ROOT}/apis/redis/v1beta1/types.generated.go

# Temporarily copy mapper.go so generate-mapper can see the manually implemented functions and skip generating them
cp ${REPO_ROOT}/pkg/controller/direct/redis/cluster/mapper.go ${REPO_ROOT}/pkg/controller/direct/redis/mapper.go
sed -i 's/package cluster/package redis/' ${REPO_ROOT}/pkg/controller/direct/redis/mapper.go

go run . generate-mapper \
  --service google.cloud.redis.cluster.v1 \
  --api-version redis.cnrm.cloud.google.com/v1beta1 \
  --include-skipped-output

rm ${REPO_ROOT}/pkg/controller/direct/redis/mapper.go

mv ${REPO_ROOT}/pkg/controller/direct/redis/mapper.generated.go ${REPO_ROOT}/pkg/controller/direct/redis/cluster/mapper.generated.go
sed -i 's/package redis/package cluster/' ${REPO_ROOT}/pkg/controller/direct/redis/cluster/mapper.generated.go

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w  pkg/controller/direct/redis/cluster/
