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

REPO_ROOT=$(git rev-parse --show-toplevel)
MOCKGCP=${REPO_ROOT}/mockgcp
cd ${MOCKGCP}/third_party/googleapis

GOOGLEAPIS=$(pwd)

# We need to rewrite various proto import paths, to avoid conflicts with the version that is in use in KCC itself.
# We also need to be a little careful not to rewrite the google.api.http annotations, which are used by grpc-go.

rm -rf mockgcp/ && mkdir -p mockgcp/api/ && mkdir -p mockgcp/devtools

mv google/cloud/ mockgcp/
mv google/container/ mockgcp/
mv google/dataflow/ mockgcp/
mv google/firestore/ mockgcp/
mv google/iam/ mockgcp/
mv google/logging/ mockgcp/
mv google/pubsub/ mockgcp/
mv google/spanner/ mockgcp/
mv google/storage/ mockgcp/
mv google/monitoring/ mockgcp/

mv google/api/apikeys/ mockgcp/api/
mv google/api/serviceusage/ mockgcp/api/
mv google/api/cloudquotas/ mockgcp/api/
mv google/devtools/artifactregistry mockgcp/devtools
mv google/devtools/cloudbuild mockgcp/devtools

rm -rf mockgrafeas/
mv grafeas mockgrafeas

cd mockgcp

# Rewrite import paths: google.cloud. -> mockgcp.cloud. etc
find . -type f -print0 | xargs -0 sed -i \
  -e "s@google/cloud/@mockgcp/cloud/@g" \
  -e "s@google\.cloud@mockgcp.cloud@g" \
  -e "s@google/container/@mockgcp/container/@g" \
  -e "s@google\.container@mockgcp.container@g" \
  -e "s@google/dataflow/@mockgcp/dataflow/@g" \
  -e "s@google\.dataflow@mockgcp.dataflow@g" \
  -e "s@google/firestore/@mockgcp/firestore/@g" \
  -e "s@google\.firestore@mockgcp.firestore@g" \
  -e "s@google/iam/@mockgcp/iam/@g" \
  -e "s@google\.iam@mockgcp.iam@g" \
  -e "s@google/logging/@mockgcp/logging/@g" \
  -e "s@google\.logging@mockgcp.logging@g" \
  -e "s@google/pubsub/@mockgcp/pubsub/@g" \
  -e "s@google\.pubsub@mockgcp.pubsub@g" \
  -e "s@google/monitoring/@mockgcp/monitoring/@g" \
  -e "s@google\.monitoring@mockgcp.monitoring@g" \
  -e "s@google/storage/@mockgcp/storage/@g" \
  -e "s@google\.storage@mockgcp.storage@g" \
  -e "s@google/spanner/@mockgcp/spanner/@g" \
  -e "s@google\.spanner@mockgcp.spanner@g" \
  -e "s@google/api/apikeys/@mockgcp/api/apikeys/@g" \
  -e "s@google\.api\.apikeys@mockgcp.api.apikeys@g" \
  -e "s@google/api/serviceusage/@mockgcp/api/serviceusage/@g" \
  -e "s@google\.api\.serviceusage@mockgcp.api.serviceusage@g" \
  -e "s@google/api/cloudquotas/@mockgcp/api/cloudquotas/@g" \
  -e "s@google\.api\.cloudquotas@mockgcp.api.cloudquotas@g" \
  -e "s@google/devtools/artifactregistry/@mockgcp/devtools/artifactregistry/@g" \
  -e "s@google\.devtools\.artifactregistry@mockgcp.devtools.artifactregistry@g" \
  -e "s@google/devtools/cloudbuild/@mockgcp/devtools/cloudbuild/@g" \
  -e "s@google\.devtools\.cloudbuild@mockgcp.devtools.cloudbuild@g" \
  -e 's@option go_package = "cloud.google.com/go/gkehub/configmanagement/apiv1beta/configmanagementpb;configmanagementpb"@option go_package = "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1beta/configmanagement;configmanagementpb"@g' \
  -e 's@option go_package = "cloud.google.com/go/gkehub/metering/apiv1beta/meteringpb;meteringpb"@option go_package = "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1beta/metering;meteringpb"@g' \
  -e 's@option go_package = "cloud.google.com/go/gkehub/multiclusteringress/apiv1beta/multiclusteringresspb;multiclusteringresspb"@option go_package = "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1beta/multiclusteringress;multiclusteringresspb"@g' \
  -e 's@option go_package = "cloud.google.com/go/gkehub/servicemesh/apiv1beta/servicemeshpb;servicemeshpb"@option go_package = "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/servicemesh/v1beta;servicemeshpb"@g' \
  -e 's@option go_package = "cloud.google.com/go/gkehub/policycontroller/apiv1beta/policycontrollerpb;policycontrollerpb"@option go_package = "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/policycontroller/v1beta;policycontrollerpb"@g'


cd ${GOOGLEAPIS}/mockgrafeas
# Rewrite grafeas -> mockgrafeas
find . -type f -print0 | xargs -0 sed -i \
  -e "s@grafeas/@mockgrafeas/@g" \
  -e "s@grafeas\.@mockgrafeas.@g"
