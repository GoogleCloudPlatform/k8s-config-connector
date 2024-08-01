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


set -e
set -x

# RedisCluster

go run . generate-types  \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.redis.cluster.v1 \
    --version redis.cnrm.cloud.google.com/v1alpha1  \
    --output-api ~/kcc/k8s-config-connector/apis/ \
    --kinds RedisCluster 

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.redis.cluster.v1 \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ~/kcc/k8s-config-connector/pkg/controller/direct/ \
    --api-dir ~/kcc/k8s-config-connector/apis/

# Bigtable

go run . generate-types  \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.bigtable.admin.v2 \
    --version bigtable.cnrm.cloud.google.com/v1beta1  \
    --output-api ~/kcc/k8s-config-connector/apis/ \
    --kinds BigtableInstance

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.bigtable.admin.v2 \
    --version bigtable.cnrm.cloud.google.com/v1beta1  \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ~/kcc/k8s-config-connector/pkg/controller/direct/ \
    --api-dir ~/kcc/k8s-config-connector/apis/

# NetworkConnectivity
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service mockgcp.cloud.networkconnectivity.v1 \
    --version networkconnectivity.cnrm.cloud.google.com/v1alpha1 \
    --output-api ~/kcc/k8s-config-connector/apis \
    --kinds NetworkConnectivityServiceConnectionPolicy

# TODO: mappers
# go run . generate-mapper \
#     --proto-source-path ../proto-to-mapper/build/googleapis.pb \
#     --service mockgcp.cloud.networkconnectivity.v1 \
#     --version networkconnectivity.cnrm.cloud.google.com/v1alpha1 \
#     --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
#     --output-dir ~/kcc/k8s-config-connector/pkg/controller/direct/ \
#     --api-dir ~/kcc/k8s-config-connector/apis/
