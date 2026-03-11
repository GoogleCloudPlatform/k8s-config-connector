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
${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.cloud.discoveryengine.v1 \
    --api-version discoveryengine.cnrm.cloud.google.com/v1alpha1 \
    --resource DiscoveryEngineDataStore:DataStore \
    --resource DiscoveryEngineDataStoreTargetSite:TargetSite \
    --resource DiscoveryEngineEngine:Engine

# ${CONTROLLERBUILDER:-go run .} prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF
# // +kcc:proto=google.cloud.discoveryengine.v1.Engine
# EOF

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.cloud.discoveryengine.v1 \
    --api-version discoveryengine.cnrm.cloud.google.com/v1alpha1

# DataFlow
${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.dataflow.v1beta3 \
    --api-version dataflow.cnrm.cloud.google.com/v1beta1 \
    --resource DataflowFlexTemplateJob:FlexTemplateRuntimeEnvironment

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.dataflow.v1beta3 \
    --api-version dataflow.cnrm.cloud.google.com/v1alpha1

# SecureSourceManagerInstance
${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.cloud.securesourcemanager.v1 \
    --api-version securesourcemanager.cnrm.cloud.google.com/v1alpha1 \
    --resource SecureSourceManagerInstance:Instance

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.cloud.securesourcemanager.v1 \
    --api-version securesourcemanager.cnrm.cloud.google.com/v1alpha1

# RedisCluster
${CONTROLLERBUILDER:-go run .} generate-types  \
    --service google.cloud.redis.cluster.v1 \
    --api-version redis.cnrm.cloud.google.com/v1alpha1 \
    --resource RedisCluster:Cluster

${CONTROLLERBUILDER:-go run .} generate-types  \
    --service google.cloud.redis.cluster.v1 \
    --api-version redis.cnrm.cloud.google.com/v1beta1  \
    --resource RedisCluster:Cluster

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.cloud.redis.cluster.v1 \
    --api-version redis.cnrm.cloud.google.com/v1beta1

# Bigtable

${CONTROLLERBUILDER:-go run .} generate-types  \
    --service google.bigtable.admin.v2 \
    --api-version bigtable.cnrm.cloud.google.com/v1beta1  \
    --resource BigtableInstance:Instance

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.bigtable.admin.v2 \
    --api-version bigtable.cnrm.cloud.google.com/v1beta1

# NetworkConnectivity
${CONTROLLERBUILDER:-go run .} generate-types \
    --service mockgcp.cloud.networkconnectivity.v1 \
    --api-version networkconnectivity.cnrm.cloud.google.com/v1alpha1 \
    --resource NetworkConnectivityServiceConnectionPolicy:ServiceConnectionPolicy

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service mockgcp.cloud.networkconnectivity.v1 \
    --api-version networkconnectivity.cnrm.cloud.google.com/v1alpha1

# BigQueryDataset
${CONTROLLERBUILDER:-go run .} generate-types  \
    --service google.cloud.bigquery.v2 \
    --api-version bigquery.cnrm.cloud.google.com/v1beta1  \
    --resource BigQueryDataset:Dataset

# ${CONTROLLERBUILDER:-go run .} generate-mapper \
#     --service google.cloud.bigquery.v2 \
#     --api-version bigquery.cnrm.cloud.google.com/v1beta1

# BigQueryDataTransferConfig
${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.cloud.bigquery.datatransfer.v1 \
    --api-version bigquerydatatransfer.cnrm.cloud.google.com/v1beta1 \
    --resource BigQueryDataTransferConfig:TransferConfig \
    --skip-scaffold-files # skipping because the files were generated using a previous pattern, making them incompatible with the new scaffolding approach.

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.cloud.bigquery.datatransfer.v1 \
    --api-version bigquerydatatransfer.cnrm.cloud.google.com/v1beta1

# Firestore
${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.firestore.admin.v1 \
    --api-version firestore.cnrm.cloud.google.com/v1beta1 \
    --resource FirestoreDatabase:Database

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.firestore.admin.v1 \
    --api-version firestore.cnrm.cloud.google.com/v1beta1

# Certificate Manager DNSAuthorization
${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.cloud.certificatemanager.v1  \
    --resource CertificateManagerDNSAuthorization:DnsAuthorization \
    --api-version "certificatemanager.cnrm.cloud.google.com/v1beta1"

${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.cloud.certificatemanager.v1  \
    --resource CertificateManagerDNSAuthorization:DnsAuthorization \
    --api-version "certificatemanager.cnrm.cloud.google.com/v1alpha1"

# Workstations
${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.cloud.workstations.v1 \
    --api-version workstations.cnrm.cloud.google.com/v1beta1 \
    --resource WorkstationCluster:WorkstationCluster

${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.cloud.workstations.v1 \
    --api-version workstations.cnrm.cloud.google.com/v1beta1 \
    --resource WorkstationConfig:WorkstationConfig

${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.cloud.workstations.v1 \
    --api-version workstations.cnrm.cloud.google.com/v1beta1 \
    --resource Workstation:Workstation

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.cloud.workstations.v1 \
    --api-version workstations.cnrm.cloud.google.com/v1beta1

# SecretManager
${CONTROLLERBUILDER:-go run .} generate-types \
     --service google.cloud.secretmanager.v1 \
     --resource SecretManagerSecret:Secret \
     --api-version "secretmanager.cnrm.cloud.google.com/v1beta1"

${CONTROLLERBUILDER:-go run .} generate-mapper \
   --service google.cloud.secretmanager.v1 \
   --api-version "secretmanager.cnrm.cloud.google.com/v1beta1"

# Spanner
${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.spanner.admin.instance.v1 \
    --resource SpannerInstance:Instance \
    --api-version "spanner.cnrm.cloud.google.com/v1beta1"

${CONTROLLERBUILDER:-go run .} generate-mapper \
   --service google.spanner.admin.instance.v1  \
   --api-version "spanner.cnrm.cloud.google.com/v1beta1"

# IAPSettings
${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.cloud.iap.v1 \
    --api-version iap.cnrm.cloud.google.com/v1beta1 \
    --resource IAPSettings:IapSettings

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.cloud.iap.v1 \
    --api-version iap.cnrm.cloud.google.com/v1beta1

# ManagedKafka
${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.cloud.managedkafka.v1 \
    --api-version managedkafka.cnrm.cloud.google.com/v1alpha1 \
    --resource ManagedKafkaCluster:Cluster \
    --resource ManagedKafkaTopic:Topic

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.cloud.managedkafka.v1 \
    --api-version managedkafka.cnrm.cloud.google.com/v1alpha1

# PrivilegedAccessManager
${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.cloud.privilegedaccessmanager.v1 \
    --api-version privilegedaccessmanager.cnrm.cloud.google.com/v1beta1

# Apigee
${CONTROLLERBUILDER:-go run .} generate-types \
    --service mockgcp.cloud.apigee.v1 \
    --api-version apigee.cnrm.cloud.google.com/v1alpha1 \
    --resource ApigeeInstance:GoogleCloudApigeeV1Instance

# CloudIdentity
${CONTROLLERBUILDER:-go run .} generate-types \
     --service google.apps.cloudidentity.v1beta1 \
     --resource CloudIdentityGroup:Group \
     --resource CloudIdentityMembership:Membership \
     --api-version "cloudidentity.cnrm.cloud.google.com/v1beta1"

${CONTROLLERBUILDER:-go run .} generate-mapper \
     --service google.apps.cloudidentity.v1beta1 \
     --api-version cloudidentity.cnrm.cloud.google.com/v1beta1

# Workflow : Workflow
${CONTROLLERBUILDER:-go run .} generate-types \
     --service google.cloud.workflows.v1 \
     --resource WorkflowsWorkflow:Workflow \
     --api-version "workflows.cnrm.cloud.google.com/v1alpha1"

${CONTROLLERBUILDER:-go run .} generate-mapper \
     --service google.cloud.workflows.v1 \
     --api-version workflows.cnrm.cloud.google.com/v1alpha1

# DocumentAI
${CONTROLLERBUILDER:-go run .} generate-types \
    --service google.cloud.documentai.v1 \
    --api-version documentai.cnrm.cloud.google.com/v1alpha1 \
    --resource DocumentAIProcessor:Processor

${CONTROLLERBUILDER:-go run .} generate-mapper \
    --service google.cloud.documentai.v1 \
    --api-version documentai.cnrm.cloud.google.com/v1alpha1

# AlloyDB
${CONTROLLERBUILDER:-go run .} generate-types \
     --service google.cloud.alloydb.v1beta \
     --api-version alloydb.cnrm.cloud.google.com/v1beta1 \
     --resource AlloyDBCluster:Cluster \
     --resource AlloyDBInstance:Instance

${CONTROLLERBUILDER:-go run .} generate-mapper \
   --service google.cloud.alloydb.v1beta  \
   --api-version alloydb.cnrm.cloud.google.com/v1alpha1

# Fix up formatting
${REPO_ROOT}/dev/tasks/fix-gofmt
