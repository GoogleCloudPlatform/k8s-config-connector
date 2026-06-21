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

# Dataform patches

go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/dataform/v1/dataform.proto --service "Dataform" --mode "append" <<EOF
  // Fetches a single TeamFolder.
  rpc GetTeamFolder(GetTeamFolderRequest) returns (TeamFolder) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/*/teamFolders/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Creates a new TeamFolder in a given project and location.
  rpc CreateTeamFolder(CreateTeamFolderRequest) returns (TeamFolder) {
    option (google.api.http) = {
      post: "/v1/{parent=projects/*/locations/*}/teamFolders"
      body: "team_folder"
    };
    option (google.api.method_signature) = "parent,team_folder";
  }

  // Updates a single TeamFolder.
  rpc UpdateTeamFolder(UpdateTeamFolderRequest) returns (TeamFolder) {
    option (google.api.http) = {
      patch: "/v1/{team_folder.name=projects/*/locations/*/teamFolders/*}"
      body: "team_folder"
    };
    option (google.api.method_signature) = "team_folder,update_mask";
  }

  // Deletes a single TeamFolder.
  rpc DeleteTeamFolder(DeleteTeamFolderRequest)
      returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v1/{name=projects/*/locations/*/teamFolders/*}"
    };
    option (google.api.method_signature) = "name";
  }
EOF

cat >> ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/dataform/v1/dataform.proto <<EOF

// \`GetTeamFolder\` request message.
message GetTeamFolderRequest {
  // Required. The team folder's name.
  string name = 1 [
    (google.api.field_behavior) = REQUIRED,
    (google.api.resource_reference) = {
      type: "dataform.googleapis.com/TeamFolder"
    }
  ];
}

// \`CreateTeamFolder\` request message.
message CreateTeamFolderRequest {
  // Required. The location in which to create the TeamFolder. Must be in the
  // format \`projects/*/locations/*\`.
  string parent = 1 [
    (google.api.field_behavior) = REQUIRED,
    (google.api.resource_reference) = {
      type: "locations.googleapis.com/Location"
    }
  ];

  // Required. The TeamFolder to create.
  TeamFolder team_folder = 2 [(google.api.field_behavior) = REQUIRED];
}

// \`UpdateTeamFolder\` request message.
message UpdateTeamFolderRequest {
  // Optional. Specifies the fields to be updated in the Folder. If left unset,
  // all fields will be updated.
  google.protobuf.FieldMask update_mask = 1
      [(google.api.field_behavior) = OPTIONAL];

  // Required. The updated TeamFolder.
  TeamFolder team_folder = 2 [(google.api.field_behavior) = REQUIRED];
}

// \`DeleteTeamFolder\` request message.
message DeleteTeamFolderRequest {
  // Required. The team folder's name.
  string name = 1 [
    (google.api.field_behavior) = REQUIRED,
    (google.api.resource_reference) = {
      type: "dataform.googleapis.com/TeamFolder"
    }
  ];
}

// Represents a Dataform team folder.
message TeamFolder {
  option (google.api.resource) = {
    type: "dataform.googleapis.com/TeamFolder"
    pattern: "projects/{project}/locations/{location}/teamFolders/{team_folder}"
  };

  // Identifier. The TeamFolder's name.
  string name = 1 [(google.api.field_behavior) = IDENTIFIER];

  // Required. The TeamFolder's user-friendly name.
  string display_name = 2 [(google.api.field_behavior) = REQUIRED];

  // Output only. The timestamp of when the TeamFolder was created.
  google.protobuf.Timestamp create_time = 3 [(google.api.field_behavior) = OUTPUT_ONLY];

  // Output only. The timestamp of when the TeamFolder was last updated.
  google.protobuf.Timestamp update_time = 4 [(google.api.field_behavior) = OUTPUT_ONLY];

  // Output only. Real-time metadata about the TeamFolder.
  string internal_metadata = 5 [(google.api.field_behavior) = OUTPUT_ONLY];

  // Output only. The IAM principal of the creator of the TeamFolder.
  string creator_iam_principal = 6 [(google.api.field_behavior) = OUTPUT_ONLY];
}
EOF

# NetworkServices EdgeCacheService patches
NS_PROTO=${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/networkservices/v1/network_services.proto
sed -i 's|option go_package = "cloud.google.com/go/networkservices/apiv1/networkservicespb;networkservicespb";|option go_package = "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkservices/v1;networkservicespb";|' ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/networkservices/v1/*.proto
sed -i 's|import "google/protobuf/empty.proto";|import "google/protobuf/empty.proto";\nimport "google/protobuf/field_mask.proto";\nimport "google/protobuf/struct.proto";|' ${NS_PROTO}

go run . --file ${NS_PROTO} --service "NetworkServices" --mode "append" <<EOF
  // Lists EdgeCacheServices in a given project and location.
  rpc ListEdgeCacheServices(ListEdgeCacheServicesRequest)
      returns (ListEdgeCacheServicesResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/global}/edgeCacheServices"
    };
    option (google.api.method_signature) = "parent";
  }

  // Gets details of a single EdgeCacheService.
  rpc GetEdgeCacheService(GetEdgeCacheServiceRequest) returns (EdgeCacheService) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/global/edgeCacheServices/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Creates a new EdgeCacheService in a given project and location.
  rpc CreateEdgeCacheService(CreateEdgeCacheServiceRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{parent=projects/*/locations/global}/edgeCacheServices"
      body: "edge_cache_service"
    };
    option (google.api.method_signature) =
        "parent,edge_cache_service,edge_cache_service_id";
    option (google.longrunning.operation_info) = {
      response_type: "EdgeCacheService"
      metadata_type: "google.cloud.networkservices.v1.OperationMetadata"
    };
  }

  // Updates the parameters of a single EdgeCacheService.
  rpc UpdateEdgeCacheService(UpdateEdgeCacheServiceRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v1/{edge_cache_service.name=projects/*/locations/global/edgeCacheServices/*}"
      body: "edge_cache_service"
    };
    option (google.api.method_signature) = "edge_cache_service,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "EdgeCacheService"
      metadata_type: "google.cloud.networkservices.v1.OperationMetadata"
    };
  }

  // Deletes a single EdgeCacheService.
  rpc DeleteEdgeCacheService(DeleteEdgeCacheServiceRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      delete: "/v1/{name=projects/*/locations/global/edgeCacheServices/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "google.protobuf.Empty"
      metadata_type: "google.cloud.networkservices.v1.OperationMetadata"
    };
  }

  // Lists EdgeCacheOrigins in a given project and location.
  rpc ListEdgeCacheOrigins(ListEdgeCacheOriginsRequest)
      returns (ListEdgeCacheOriginsResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/global}/edgeCacheOrigins"
    };
    option (google.api.method_signature) = "parent";
  }

  // Gets details of a single EdgeCacheOrigin.
  rpc GetEdgeCacheOrigin(GetEdgeCacheOriginRequest) returns (EdgeCacheOrigin) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/global/edgeCacheOrigins/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Creates a new EdgeCacheOrigin in a given project and location.
  rpc CreateEdgeCacheOrigin(CreateEdgeCacheOriginRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{parent=projects/*/locations/global}/edgeCacheOrigins"
      body: "edge_cache_origin"
    };
    option (google.api.method_signature) =
        "parent,edge_cache_origin,edge_cache_origin_id";
    option (google.longrunning.operation_info) = {
      response_type: "EdgeCacheOrigin"
      metadata_type: "google.cloud.networkservices.v1.OperationMetadata"
    };
  }

  // Updates the parameters of a single EdgeCacheOrigin.
  rpc UpdateEdgeCacheOrigin(UpdateEdgeCacheOriginRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v1/{edge_cache_origin.name=projects/*/locations/global/edgeCacheOrigins/*}"
      body: "edge_cache_origin"
    };
    option (google.api.method_signature) = "edge_cache_origin,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "EdgeCacheOrigin"
      metadata_type: "google.cloud.networkservices.v1.OperationMetadata"
    };
  }

  // Deletes a single EdgeCacheOrigin.
  rpc DeleteEdgeCacheOrigin(DeleteEdgeCacheOriginRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      delete: "/v1/{name=projects/*/locations/global/edgeCacheOrigins/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "google.protobuf.Empty"
      metadata_type: "google.cloud.networkservices.v1.OperationMetadata"
    };
  }
EOF

cat <<EOF >> ${NS_PROTO}

message EdgeCacheOrigin {
  string name = 1;
  string description = 2;
  google.protobuf.Struct labels = 3;
  string origin_address = 4;
  int32 max_attempts = 5;
  google.protobuf.Struct timeout = 6;
  string protocol = 7;
  string port = 8;
  google.protobuf.Struct retry_conditions = 9;
}

message ListEdgeCacheOriginsRequest {
  string parent = 1;
  int32 page_size = 2;
  string page_token = 3;
  string filter = 4;
  string order_by = 5;
}

message ListEdgeCacheOriginsResponse {
  repeated EdgeCacheOrigin edge_cache_origins = 1;
  string next_page_token = 2;
}

message GetEdgeCacheOriginRequest {
  string name = 1;
}

message CreateEdgeCacheOriginRequest {
  string parent = 1;
  string edge_cache_origin_id = 2;
  EdgeCacheOrigin edge_cache_origin = 3;
}

message UpdateEdgeCacheOriginRequest {
  google.protobuf.FieldMask update_mask = 1;
  EdgeCacheOrigin edge_cache_origin = 2;
}

message DeleteEdgeCacheOriginRequest {
  string name = 1;
}

message EdgeCacheService {
  string name = 1;
  string description = 2;
  google.protobuf.Struct labels = 3;
  bool disable_http2 = 4;
  bool disable_quic = 5;
  string edge_security_policy = 6;
  repeated string edge_ssl_certificates = 7;
  google.protobuf.Struct log_config = 8;
  bool require_tls = 9;
  google.protobuf.Struct routing = 10;
  string ssl_policy = 11;
}

message ListEdgeCacheServicesRequest {
  string parent = 1;
  int32 page_size = 2;
  string page_token = 3;
  string filter = 4;
  string order_by = 5;
}

message ListEdgeCacheServicesResponse {
  repeated EdgeCacheService edge_cache_services = 1;
  string next_page_token = 2;
}

message GetEdgeCacheServiceRequest {
  string name = 1;
}

message CreateEdgeCacheServiceRequest {
  string parent = 1;
  string edge_cache_service_id = 2;
  EdgeCacheService edge_cache_service = 3;
}

message UpdateEdgeCacheServiceRequest {
  google.protobuf.FieldMask update_mask = 1;
  EdgeCacheService edge_cache_service = 2;
}

message DeleteEdgeCacheServiceRequest {
  string name = 1;
}
EOF
