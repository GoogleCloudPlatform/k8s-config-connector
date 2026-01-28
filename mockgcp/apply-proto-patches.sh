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

# SQL patches

go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/sql/v1beta4/cloud_sql_resources.proto --message BackupConfiguration --mode append <<EOF

  // The tier of the backup.

  optional string backup_tier = 14;

EOF

go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/sql/v1beta4/cloud_sql_resources.proto --message IpConfiguration --mode append <<EOF

  // The server certificate rotation mode.

  optional string server_certificate_rotation_mode = 10;

EOF

go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/sql/v1beta4/cloud_sql_resources.proto --message Settings --mode append <<EOF

  // Maximum replication lag in seconds.

  optional int64 replication_lag_max_seconds = 100;

EOF

go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/sql/v1beta4/cloud_sql_resources.proto --message DatabaseInstance --mode append <<EOF

  // Whether to include replicas for major version upgrade.

  optional bool include_replicas_for_major_version_upgrade = 56;



  // Whether the instance satisfies PZI.

  optional bool satisfies_pzi = 57;

EOF

go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/sql/v1beta4/cloud_sql_users.proto --message User --mode append <<EOF

  // The status of the user's IAM authentication.

  optional string iam_status = 16;

EOF



# AlloyDB patches

sed -i 's/map<string, string> flags = 13 \[(google.api.field_behavior) = OPTIONAL\];/map<string, string> flags = 13 [(google.api.field_behavior) = OPTIONAL];\n    int32 pooler_count = 14;/' ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/alloydb/v1beta/resources.proto
