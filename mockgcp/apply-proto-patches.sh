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

set -x

REPO_ROOT=$(git rev-parse --show-toplevel)
cd ${REPO_ROOT}/mockgcp

cd tools/patch-proto

go run . --file ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/apigee/v1/service.proto --service "ProjectsServer" --mode "replace" <<EOF
  // Provisions a new Apigee organization with a functioning runtime. This is the standard way to create trial organizations for a free Apigee trial.
  rpc ProvisionOrganizationProject(ProvisionOrganizationProjectRequest) returns (.google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{name=projects/*}:provisionOrganization"
      body: "project"
    };
  };
EOF

# ResourceManager v1 patches - temporarily switching to proto3 because patch-proto has issues with proto2
API_PROTO=${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/resourcemanager/v1/api.proto
sed -i 's/^syntax = "proto2";/syntax = "proto3";/' ${API_PROTO}

go run . --file ${API_PROTO} --message "TestIamPermissionsRequest" --mode "replace" <<EOF
  // REQUIRED: The resource for which the permission checking is to be performed.
  optional string resource = 1;

  // The set of permissions to check for the resource.
  repeated string permissions = 2 [json_name="permissions"];
EOF

go run . --file ${API_PROTO} --service "FoldersServer" --mode "append" <<EOF
  // Returns permissions that a caller has on the specified project.
  rpc TestIamPermissions(TestIamPermissionsRequest) returns (TestIamPermissionsResponse) {
    option (google.api.http) = {
      post: "/v1/{resource=folders/*}:testIamPermissions"
      body: "*"
    };
  };
EOF

go run . --file ${API_PROTO} --service "OrganizationsServer" --mode "append" <<EOF
  // Returns permissions that a caller has on the specified project.
  rpc TestIamPermissions(TestIamPermissionsRequest) returns (TestIamPermissionsResponse) {
    option (google.api.http) = {
      post: "/v1/{resource=organizations/*}:testIamPermissions"
      body: "*"
    };
  };
EOF

go run . --file ${API_PROTO} --service "ProjectsServer" --mode "append" <<EOF
  // Returns permissions that a caller has on the specified project.
  rpc TestIamPermissions(TestIamPermissionsRequest) returns (TestIamPermissionsResponse) {
    option (google.api.http) = {
      post: "/v1/{resource=projects/*}:testIamPermissions"
      body: "*"
    };
  };
EOF

sed -i 's/^syntax = "proto3";/syntax = "proto2";/' ${API_PROTO}
