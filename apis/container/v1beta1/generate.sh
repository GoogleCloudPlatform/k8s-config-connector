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
  --service google.container.v1beta1 \
  --api-version container.cnrm.cloud.google.com/v1beta1 \
  --include-skipped-output \
  --skip-scaffold-files \
  --resource ContainerCluster:Cluster \
  --resource ContainerNodePool:NodePool

go run . generate-mapper \
  --service google.container.v1beta1 \
  --api-version container.cnrm.cloud.google.com/v1beta1 \
  --include-skipped-output

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/container/





# Fix capitalization differences between KRM and protobuf
python3 -c '
import os
import re
with open("pkg/controller/direct/container/mapper.generated.go", "r") as f: s = f.read()

s = s.replace("in.GetEnableK8sBetaApis()", "in.GetEnableK8SBetaApis()")
s = s.replace("out.EnableK8sBetaApis = K8sBetaAPIConfig_ToProto", "out.EnableK8SBetaApis = K8sBetaAPIConfig_ToProto")
s = s.replace("out.EnableK8SBetaApis = K8sBetaAPIConfig_FromProto(mapCtx, in.GetEnableK8SBetaApis())", "out.EnableK8sBetaApis = K8sBetaAPIConfig_FromProto(mapCtx, in.GetEnableK8SBetaApis())")
s = s.replace("pb.K8sBetaAPIConfig", "pb.K8SBetaAPIConfig")
s = s.replace("K8sBetaAPIConfig_FromProto", "K8SBetaAPIConfig_FromProto")
s = s.replace("K8sBetaAPIConfig_ToProto", "K8SBetaAPIConfig_ToProto")

s = s.replace("in.GetEnableL4ilbSubsetting()", "in.GetEnableL4IlbSubsetting()")
s = s.replace("out.EnableL4ilbSubsetting = direct.ValueOf(in.EnableL4ilbSubsetting)", "out.EnableL4IlbSubsetting = direct.ValueOf(in.EnableL4ilbSubsetting)")
s = s.replace("out.EnableL4IlbSubsetting = direct.LazyPtr(in.GetEnableL4IlbSubsetting())", "out.EnableL4ilbSubsetting = direct.LazyPtr(in.GetEnableL4IlbSubsetting())")

s = re.sub(r"func LinuxNodeConfig_HugepagesConfig_FromProto.*?return out\n\}", "func LinuxNodeConfig_HugepagesConfig_FromProto(mapCtx *direct.MapContext, in *pb.LinuxNodeConfig_HugepagesConfig) *krm.LinuxNodeConfig_HugepagesConfig {\n\tif in == nil {\n\t\treturn nil\n\t}\n\tout := &krm.LinuxNodeConfig_HugepagesConfig{}\n\tout.HugepageSize2m = in.HugepageSize2M\n\tout.HugepageSize1g = in.HugepageSize1G\n\treturn out\n}", s, flags=re.DOTALL)

s = re.sub(r"func LinuxNodeConfig_HugepagesConfig_ToProto.*?return out\n\}", "func LinuxNodeConfig_HugepagesConfig_ToProto(mapCtx *direct.MapContext, in *krm.LinuxNodeConfig_HugepagesConfig) *pb.LinuxNodeConfig_HugepagesConfig {\n\tif in == nil {\n\t\treturn nil\n\t}\n\tout := &pb.LinuxNodeConfig_HugepagesConfig{}\n\tout.HugepageSize2M = in.HugepageSize2m\n\tout.HugepageSize1G = in.HugepageSize1g\n\treturn out\n}", s, flags=re.DOTALL)


with open("pkg/controller/direct/container/mapper.generated.go", "w") as f: f.write(s)
'
