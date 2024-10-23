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


# DiscoveryEngine
go run . generate-types \
    --service google.cloud.discoveryengine.v1 \
    --api-version discoveryengine.cnrm.cloud.google.com/v1alpha1 \
    --resource DiscoveryEngineDataStore:DataStore \
    --resource DiscoveryEngineDataStoreTargetSite:TargetSite \
    --resource DiscoveryEngineEngine:Engine

# go run . prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF
# // +kcc:proto=google.cloud.discoveryengine.v1.Engine
# EOF

go run . generate-mapper \
    --service google.cloud.discoveryengine.v1 \
    --api-version discoveryengine.cnrm.cloud.google.com/v1alpha1

# DiscoveryEngine
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.discoveryengine.v1 \
    --api-version discoveryengine.cnrm.cloud.google.com/v1alpha1 \
    --output-api ${APIS_DIR} \
    --kind DiscoveryEngineEngine \
    --proto-resource Engine

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
    --service google.dataflow.v1beta3 \
    --api-version dataflow.cnrm.cloud.google.com/v1beta1 \
    --resource DataflowFlexTemplateJob:FlexTemplateRuntimeEnvironment

go run . generate-mapper \
    --service google.dataflow.v1beta3 \
    --api-version dataflow.cnrm.cloud.google.com/v1alpha1

# SecureSourceManagerInstance
go run . generate-types \
    --service google.cloud.securesourcemanager.v1 \
    --api-version securesourcemanager.cnrm.cloud.google.com/v1alpha1 \
    --resource SecureSourceManagerInstance:Instance

go run . generate-mapper \
    --service google.cloud.securesourcemanager.v1 \
    --api-version securesourcemanager.cnrm.cloud.google.com/v1alpha1

# RedisCluster
go run . generate-types  \
    --service google.cloud.redis.cluster.v1 \
    --api-version redis.cnrm.cloud.google.com/v1alpha1 \
    --resource RedisCluster:Cluster

go run . generate-types  \
    --service google.cloud.redis.cluster.v1 \
    --api-version redis.cnrm.cloud.google.com/v1beta1  \
    --resource RedisCluster:Cluster

go run . generate-mapper \
    --service google.cloud.redis.cluster.v1 \
    --api-version redis.cnrm.cloud.google.com/v1beta1

# Bigtable

go run . generate-types  \
    --service google.bigtable.admin.v2 \
    --api-version bigtable.cnrm.cloud.google.com/v1beta1  \
    --resource BigtableInstance:Instance

go run . generate-mapper \
    --service google.bigtable.admin.v2 \
    --api-version bigtable.cnrm.cloud.google.com/v1beta1

# NetworkConnectivity
go run . generate-types \
    --service mockgcp.cloud.networkconnectivity.v1 \
    --api-version networkconnectivity.cnrm.cloud.google.com/v1alpha1 \
    --resource NetworkConnectivityServiceConnectionPolicy:ServiceConnectionPolicy

go run . generate-mapper \
    --service mockgcp.cloud.networkconnectivity.v1 \
    --api-version networkconnectivity.cnrm.cloud.google.com/v1alpha1

# BigQueryDataset
go run . generate-types  \
    --service google.cloud.bigquery.v2 \
    --api-version bigquery.cnrm.cloud.google.com/v1beta1  \
    --resource BigQueryDataset:Dataset

# go run . generate-mapper \
#     --service google.cloud.bigquery.v2 \
#     --api-version bigquery.cnrm.cloud.google.com/v1beta1

# BigQueryDataTransferConfig
go run . generate-types \
    --service google.cloud.bigquery.datatransfer.v1 \
    --api-version bigquerydatatransfer.cnrm.cloud.google.com/v1beta1 \
    --resource BigQueryDataTransferConfig:TransferConfig \
    --skip-scaffold-files # skipping because the files were generated using a previous pattern, making them incompatible with the new scaffolding approach.

go run . generate-mapper \
    --service google.cloud.bigquery.datatransfer.v1 \
    --api-version bigquerydatatransfer.cnrm.cloud.google.com/v1beta1

# Firestore
go run . generate-types \
    --service google.firestore.admin.v1 \
    --api-version firestore.cnrm.cloud.google.com/v1beta1 \
    --resource FirestoreDatabase:Database

go run . generate-mapper \
    --service google.firestore.admin.v1 \
    --api-version firestore.cnrm.cloud.google.com/v1beta1

# Certificate Manager DNSAuthorization
go run . generate-types \
    --service google.cloud.certificatemanager.v1  \
    --resource CertificateManagerDNSAuthorization:DnsAuthorization \
    --api-version "certificatemanager.cnrm.cloud.google.com/v1beta1"

go run . generate-types \
    --service google.cloud.certificatemanager.v1  \
    --resource CertificateManagerDNSAuthorization:DnsAuthorization \
    --api-version "certificatemanager.cnrm.cloud.google.com/v1alpha1"

# Workstations
go run . generate-types \
    --service google.cloud.workstations.v1 \
    --api-version workstations.cnrm.cloud.google.com/v1beta1 \
    --resource WorkstationCluster:WorkstationCluster

go run . generate-types \
    --service google.cloud.workstations.v1 \
    --api-version workstations.cnrm.cloud.google.com/v1beta1 \
    --resource WorkstationConfig:WorkstationConfig

go run . generate-types \
    --service google.cloud.workstations.v1 \
    --api-version workstations.cnrm.cloud.google.com/v1beta1 \
    --resource Workstation:Workstation

go run . generate-mapper \
    --service google.cloud.workstations.v1 \
    --api-version workstations.cnrm.cloud.google.com/v1beta1

# SecretManager
go run main.go generate-types \
     --service google.cloud.secretmanager.v1 \
     --resource SecretManagerSecret:Secret \
     --api-version "secretmanager.cnrm.cloud.google.com/v1beta1"

go run . generate-mapper \
   --service google.cloud.secretmanager.v1 \
   --api-version "secretmanager.cnrm.cloud.google.com/v1beta1"

# Spanner
go run main.go generate-types \
    --service google.spanner.admin.instance.v1 \
    --resource SpannerInstance:Instance \
    --api-version "spanner.cnrm.cloud.google.com/v1beta1"

go run . generate-mapper \
   --service google.spanner.admin.instance.v1  \
   --api-version "spanner.cnrm.cloud.google.com/v1beta1"

# IAPSettings
go run . generate-types \
    --service google.cloud.iap.v1 \
    --api-version iap.cnrm.cloud.google.com/v1alpha1 \
    --resource IAPSettings:IapSettings

go run . generate-mapper \
    --service google.cloud.iap.v1 \
    --api-version iap.cnrm.cloud.google.com/v1alpha1

# ManagedKafka
go run . generate-types \
    --service google.cloud.managedkafka.v1 \
    --api-version managedkafka.cnrm.cloud.google.com/v1alpha1 \
    --resource ManagedKafkaCluster:Cluster \
    --resource ManagedKafkaTopic:Topic

go run . generate-mapper \
    --service google.cloud.managedkafka.v1 \
    --api-version managedkafka.cnrm.cloud.google.com/v1alpha1

# PrivilegedAccessManager
go run . generate-mapper \
    --service google.cloud.privilegedaccessmanager.v1 \
    --api-version privilegedaccessmanager.cnrm.cloud.google.com/v1beta1

# Apigee
go run . generate-types \
    --service mockgcp.cloud.apigee.v1 \
    --api-version apigee.cnrm.cloud.google.com/v1alpha1 \
    --resource ApigeeInstance:GoogleCloudApigeeV1Instance

# CloudIdentity : Group
go run . generate-types \
     --service google.apps.cloudidentity.groups.v1beta1 \
     --resource CloudIdentityGroup:Group\
     --api-version "cloudidentity.cnrm.cloud.google.com/v1beta1"

go run . generate-mapper \
     --service google.apps.cloudidentity.groups.v1beta1 \
     --api-version cloudidentity.cnrm.cloud.google.com/v1beta1

# CloudIdentity : Membership
go run . generate-types \
     --service mockgcp.cloud.cloudidentity.groups.v1beta1 \
     --resource CloudIdentityMembership:Membership \
     --api-version "cloudidentity.cnrm.cloud.google.com/v1beta1"

go run . generate-mapper \
     --service mockgcp.cloud.cloudidentity.groups.v1beta1 \
     --api-version cloudidentity.cnrm.cloud.google.com/v1beta1

# Fix up formatting
${REPO_ROOT}/dev/tasks/fix-gofmt
