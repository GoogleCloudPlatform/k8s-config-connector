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

  optional int32 replication_lag_max_seconds = 100;

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

go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/alloydb/v1beta/resources.proto --message ConnectionPoolConfig --mode append <<EOF

  // The number of pooler instances.

  int32 pooler_count = 14;

EOF

# Container/GKE patches

sed -i '/message NodeNetworkConfig {/,/}/ s/string subnetwork = 19 \[(google.api.field_behavior) = OUTPUT_ONLY\];/string subnetwork = 19;/' ${REPO_ROOT}/mockgcp/third_party/googleapis/google/container/v1beta1/cluster_service.proto

go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/container/v1beta1/cluster_service.proto --message AdditionalIPRangesConfig --mode append <<EOF

  // Output only. The status of the additional pod ranges.
  string status = 3;

EOF

go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/container/v1beta1/cluster_service.proto --message DNSEndpointConfig --mode append <<EOF

  // Controls whether the k8s token auth is allowed via DNS.

  optional bool enable_k8s_tokens_via_dns = 5;

EOF

go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/container/v1beta1/cluster_service.proto --message LinuxNodeConfig --mode append <<'EOF'


  // Configuration for swap memory on a node pool.
  message SwapConfig {
    // Defines encryption settings for the swap space.
    message EncryptionConfig {
      // Optional. If true, swap space will not be encrypted.
      // Defaults to false (encrypted).
      optional bool disabled = 1;
    }

    // Swap on the node's boot disk.
    message BootDiskProfile {
      // Optional. Specifies the size of the swap space. If omitted, GKE
      // determines an optimal size based on node memory.
      oneof swap_size {
        // Specifies the size of the swap space in gibibytes (GiB).
        int64 swap_size_gib = 1;

        // Specifies the size of the swap space as a percentage of the boot disk
        // size.
        int32 swap_size_percent = 2;
      }
    }

    // Swap on the local SSD shared with pod ephemeral storage.
    message EphemeralLocalSsdProfile {
      // Specifies the size of the swap space to be provisioned.
      oneof swap_size {
        // Specifies the size of the swap space in gibibytes (GiB).
        int64 swap_size_gib = 1;

        // Specifies the size of the swap space as a percentage of the ephemeral
        // local SSD capacity.
        int32 swap_size_percent = 2;
      }
    }

    // Provisions a new, separate local NVMe SSD exclusively for swap.
    message DedicatedLocalSsdProfile {
      // The number of physical local NVMe SSD disks to attach.
      int64 disk_count = 1;
    }

    // Optional. Enables or disables swap for the node pool.
    optional bool enabled = 1;

    // Optional. If omitted, swap space is encrypted by default.
    optional EncryptionConfig encryption_config = 2;

    // Optional. Defines the backing storage for the swap space.
    // If omitted, defaults to the 'boot_disk_profile'.
    oneof performance_profile {
      // Swap on the node's boot disk.
      BootDiskProfile boot_disk_profile = 3;

      // Swap on the local SSD shared with pod ephemeral storage.
      EphemeralLocalSsdProfile ephemeral_local_ssd_profile = 4;

      // Provisions a new, separate local NVMe SSD exclusively for swap.
      DedicatedLocalSsdProfile dedicated_local_ssd_profile = 5;
    }
  }

  // Optional. Enables and configures swap space on nodes.
  // If omitted, swap is disabled.
  optional SwapConfig swap_config = 12;

EOF

go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/container/v1beta1/cluster_service.proto --message DatabaseEncryption --mode replace <<'EOF'
  // State of etcd encryption.
  enum State {
    // Should never be set
    UNKNOWN = 0;

    // Secrets in etcd are encrypted.
    ENCRYPTED = 1;

    // Secrets in etcd are stored in plain text (at etcd level) - this is
    // unrelated to Compute Engine level full disk encryption.
    DECRYPTED = 2;

    // Secrets in etcd are encrypted.
    ALL_OBJECTS_ENCRYPTION_ENABLED = 3;
  }

  // Current State of etcd encryption.
  enum CurrentState {
    // Should never be set
    CURRENT_STATE_UNSPECIFIED = 0;

    // Secrets in etcd are encrypted.
    CURRENT_STATE_ENCRYPTED = 7;

    // Secrets in etcd are stored in plain text (at etcd level) - this is
    // unrelated to Compute Engine level full disk encryption.
    CURRENT_STATE_DECRYPTED = 2;

    // Encryption (or re-encryption with a different CloudKMS key)
    // of Secrets is in progress.
    CURRENT_STATE_ENCRYPTION_PENDING = 3;

    // Encryption (or re-encryption with a different CloudKMS key) of Secrets in
    // etcd encountered an error.
    CURRENT_STATE_ENCRYPTION_ERROR = 4;

    // De-crypting Secrets to plain text in etcd is in progress.
    CURRENT_STATE_DECRYPTION_PENDING = 5;

    // De-crypting Secrets to plain text in etcd encountered an error.
    CURRENT_STATE_DECRYPTION_ERROR = 6;
  }

  // OperationError records errors seen from CloudKMS keys
  // encountered during updates to DatabaseEncryption configuration.
  message OperationError {
    // CloudKMS key resource that had the error.
    string key_name = 1;

    // Description of the error seen during the operation.
    string error_message = 2;

    // Time when the CloudKMS error was seen.
    google.protobuf.Timestamp timestamp = 3;
  }

  // Name of CloudKMS key to use for the encryption of secrets in etcd.
  // Ex. projects/my-project/locations/global/keyRings/my-ring/cryptoKeys/my-key
  string key_name = 1;

  // The desired state of etcd encryption.
  State state = 2;

  // Output only. The current state of etcd encryption.
  optional CurrentState current_state = 3
      [(google.api.field_behavior) = OUTPUT_ONLY];

  // Output only. Keys in use by the cluster for decrypting
  // existing objects, in addition to the key in `key_name`.
  //
  // Each item is a CloudKMS key resource.
  repeated string decryption_keys = 4
      [(google.api.field_behavior) = OUTPUT_ONLY];

  // Output only. Records errors seen during DatabaseEncryption update
  // operations.
  repeated OperationError last_operation_errors = 5
      [(google.api.field_behavior) = OUTPUT_ONLY];
EOF



# ResourceManager v1 patches - temporarily switching to proto3 because patch-proto has issues with proto2
API_PROTO=${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/resourcemanager/v1/api.proto

if ! grep -q "FoldersServer" "${API_PROTO}" || ! grep -q "TestIamPermissions" "${API_PROTO}"; then
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
fi
