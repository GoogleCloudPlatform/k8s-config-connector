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

APIS_DIR=${REPO_ROOT}/apis/
OUTPUT_MAPPER=${REPO_ROOT}/pkg/controller/direct/

# DiscoveryEngine
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.discoveryengine.v1 \
    --api-version discoveryengine.cnrm.cloud.google.com/v1alpha1 \
    --output-api ${APIS_DIR} \
    --resource DiscoveryEngineDataStore:DataStore

# go run . prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/dev/tools/proto-to-mapper/third_party/googleapis/ <<EOF
# // +kcc:proto=google.cloud.discoveryengine.v1.Engine
# EOF

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.discoveryengine.v1 \
    --api-version discoveryengine.cnrm.cloud.google.com/v1alpha1 \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# DataFlow
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.dataflow.v1beta3 \
    --api-version dataflow.cnrm.cloud.google.com/v1beta1 \
    --output-api ${APIS_DIR} \
    --resource DataflowFlexTemplateJob:FlexTemplateRuntimeEnvironment

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.dataflow.v1beta3 \
    --api-version dataflow.cnrm.cloud.google.com/v1alpha1 \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# SecureSourceManagerInstance
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.securesourcemanager.v1 \
    --api-version securesourcemanager.cnrm.cloud.google.com/v1alpha1 \
    --output-api ${APIS_DIR} \
    --resource SecureSourceManagerInstance:Instance

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.securesourcemanager.v1 \
    --api-version securesourcemanager.cnrm.cloud.google.com/v1alpha1 \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# RedisCluster
go run . generate-types  \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.redis.cluster.v1 \
    --api-version redis.cnrm.cloud.google.com/v1alpha1  \
    --output-api ${APIS_DIR} \
    --resource RedisCluster:Cluster

go run . generate-types  \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.redis.cluster.v1 \
    --api-version redis.cnrm.cloud.google.com/v1beta1  \
    --output-api ${APIS_DIR} \
    --resource RedisCluster:Cluster

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.redis.cluster.v1 \
    --api-version redis.cnrm.cloud.google.com/v1beta1  \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# Bigtable

go run . generate-types  \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.bigtable.admin.v2 \
    --api-version bigtable.cnrm.cloud.google.com/v1beta1  \
    --output-api ${APIS_DIR} \
    --resource BigtableInstance:Instance

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.bigtable.admin.v2 \
    --api-version bigtable.cnrm.cloud.google.com/v1beta1  \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# NetworkConnectivity
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service mockgcp.cloud.networkconnectivity.v1 \
    --api-version networkconnectivity.cnrm.cloud.google.com/v1alpha1 \
    --output-api ${APIS_DIR} \
    --resource NetworkConnectivityServiceConnectionPolicy:ServiceConnectionPolicy

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service mockgcp.cloud.networkconnectivity.v1 \
    --api-version networkconnectivity.cnrm.cloud.google.com/v1alpha1 \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# BigQueryDataset
go run . generate-types  \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.bigquery.v2 \
    --api-version bigquery.cnrm.cloud.google.com/v1beta1  \
    --output-api ${APIS_DIR} \
    --resource BigQueryDataset:Dataset

# go run . generate-mapper \
#     --proto-source-path ../proto-to-mapper/build/googleapis.pb \
#     --service google.cloud.bigquery.v2 \
#     --api-version bigquery.cnrm.cloud.google.com/v1beta1 \
#     --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
#     --output-dir ${OUTPUT_MAPPER} \
#     --api-dir ${APIS_DIR}

# BigQueryDataTransferConfig
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.bigquery.datatransfer.v1 \
    --api-version bigquerydatatransfer.cnrm.cloud.google.com/v1alpha1 \
    --output-api ${APIS_DIR} \
    --resource BigQueryDataTransferConfig:TransferConfig

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.bigquery.datatransfer.v1 \
    --api-version bigquerydatatransfer.cnrm.cloud.google.com/v1alpha1 \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# Firestore
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.firestore.admin.v1 \
    --api-version firestore.cnrm.cloud.google.com/v1alpha1 \
    --output-api ${APIS_DIR} \
    --resource FirestoreDatabase:Database

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.firestore.admin.v1 \
    --api-version firestore.cnrm.cloud.google.com/v1alpha1 \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# Certificate Manager DNSAuthorization
go run . generate-types \
    --service google.cloud.certificatemanager.v1  \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --output-api $REPO_ROOT/apis \
    --resource CertificateManagerDNSAuthorization:DnsAuthorization \
    --api-version "certificatemanager.cnrm.cloud.google.com/v1beta1"

go run . generate-types \
    --service google.cloud.certificatemanager.v1  \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --output-api $REPO_ROOT/apis \
    --resource CertificateManagerDNSAuthorization:DnsAuthorization \
    --api-version "certificatemanager.cnrm.cloud.google.com/v1alpha1"

# Workstations
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.workstations.v1 \
    --api-version workstations.cnrm.cloud.google.com/v1beta1 \
    --output-api ${APIS_DIR} \
    --resource WorkstationCluster:WorkstationCluster

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.workstations.v1 \
    --api-version workstations.cnrm.cloud.google.com/v1beta1 \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# SecretManager
go run main.go generate-types \
     --service google.cloud.secretmanager.v1 \
     --proto-source-path ../proto-to-mapper/build/googleapis.pb \
     --output-api ${APIS_DIR} \
     --resource SecretManagerSecret:Secret \
     --api-version "secretmanager.cnrm.cloud.google.com/v1beta1"

go run . generate-mapper \
   --proto-source-path ../proto-to-mapper/build/googleapis.pb \
   --service google.cloud.secretmanager.v1 \
   --api-version "secretmanager.cnrm.cloud.google.com/v1beta1" \
   --api-go-package-path  $REPO_ROOT/apis/ \
   --output-dir $REPO_ROOT/pkg/controller/direct/ \
   --api-dir $REPO_ROOT/apis/


go run main.go generate-types \
    --service google.spanner.admin.instance.v1 \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --output-api $REPO_ROOT/apis \
    --resource SpannerInstance:Instance \
    --api-version "spanner.cnrm.cloud.google.com/v1beta1"

# Fix up formatting
${REPO_ROOT}/dev/tasks/fix-gofmt
