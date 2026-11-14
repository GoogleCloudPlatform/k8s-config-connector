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
CALLER_DIR="$(pwd)"

cd ${REPO_ROOT}/dev/tools/controllerbuilder

# We share the version with mockgcp, which is maybe a boundary violation, but is convenient.
# (It would be confusing if these were out of sync!)
# Extract the default git commit SHA version of the googleapis repository from apis/git.versions
DEFAULT_GOOGLE_API_VERSION=$(grep https://github.com/googleapis/googleapis ${REPO_ROOT}/apis/git.versions | awk '{print $2}')

# Check for --force flag or env var FORCE_GENERATE_PROTOS=1
FORCE_GENERATE="${FORCE_GENERATE_PROTOS:-0}"
if [[ "${1:-}" == "--force" || "${1:-}" == "-f" ]]; then
    FORCE_GENERATE=1
    shift
fi

# Take googleapi version as parameter, default to version from git.versions.
# Use "HEAD" to get the latest from remote.
GOOGLEAPI_VERSION=${1:-$DEFAULT_GOOGLE_API_VERSION}

# Take output path as parameter, default to .build/googleapis.pb
OUTPUT_PATH=${2:-"${REPO_ROOT}/.build/googleapis.pb"}
if [[ "${OUTPUT_PATH}" != /* ]]; then
    OUTPUT_PATH="$(realpath -m "${CALLER_DIR}/${OUTPUT_PATH}")"
fi

THIRD_PARTY="${REPO_ROOT}/.build/third_party"
mkdir -p "${THIRD_PARTY}/"

DEFAULT_GOOGLEAPI_DIR="${THIRD_PARTY}/googleapis"
if [ ! -d "${DEFAULT_GOOGLEAPI_DIR}" ]; then
    git clone --depth 1 https://github.com/googleapis/googleapis.git "${DEFAULT_GOOGLEAPI_DIR}"
fi

if [ "${GOOGLEAPI_VERSION}" == "HEAD" ]; then
    echo "Fetching latest googleapis for HEAD version"
    cd "${DEFAULT_GOOGLEAPI_DIR}"
    GOOGLEAPI_VERSION=$(git ls-remote https://github.com/googleapis/googleapis.git refs/heads/master | awk '{print $1}')
fi

if [ -n "${2:-}" ]; then
    # Explicitly provided output path, use it directly
    VERSIONED_OUTPUT_PATH="${OUTPUT_PATH}"
else
    # Default output path, version it with the SHA.
    # ${OUTPUT_PATH%.pb} strips the suffix '.pb' from the path, then we append the SHA.
    VERSIONED_OUTPUT_PATH="${OUTPUT_PATH%.pb}-${GOOGLEAPI_VERSION}.pb"
fi

# Fast-path check: if versioned pb file already exists, matches script hash, and force is not set, exit immediately
SCRIPT_HASH=$(sha256sum "${REPO_ROOT}/dev/tools/controllerbuilder/generate-proto.sh" | awk '{print $1}')
if [[ "${FORCE_GENERATE}" != "1" ]] && [ -f "${VERSIONED_OUTPUT_PATH}" ] && [ -f "${VERSIONED_OUTPUT_PATH}.hash" ]; then
    if [ "$(cat "${VERSIONED_OUTPUT_PATH}.hash" 2>/dev/null)" == "${SCRIPT_HASH}" ]; then
        if [ "${VERSIONED_OUTPUT_PATH}" != "${OUTPUT_PATH}" ]; then
            cp "${VERSIONED_OUTPUT_PATH}" "${OUTPUT_PATH}"
        fi
        exit 0
    fi
fi

if [[ "${FORCE_GENERATE}" != "1" ]] && [[ "${SKIP_GENERATE_PROTOS:-0}" == "1" ]] && [ -f "${VERSIONED_OUTPUT_PATH}" ]; then
    echo "Skipping generate-proto.sh as requested by SKIP_GENERATE_PROTOS=1 and output file exists: ${VERSIONED_OUTPUT_PATH}"
    if [ "${VERSIONED_OUTPUT_PATH}" != "${OUTPUT_PATH}" ]; then
        cp "${VERSIONED_OUTPUT_PATH}" "${OUTPUT_PATH}"
    fi
    exit 0
fi

if [ "${GOOGLEAPI_VERSION}" == "${DEFAULT_GOOGLE_API_VERSION}" ]; then
    VERSION_DIR="${DEFAULT_GOOGLEAPI_DIR}"
else
    VERSION_DIR="${THIRD_PARTY}/googleapis-${GOOGLEAPI_VERSION}"
    if [ ! -d "${VERSION_DIR}" ]; then
        git clone --depth 1 https://github.com/googleapis/googleapis.git "${VERSION_DIR}"
    fi
fi

cd "${VERSION_DIR}"

if ! git cat-file -e ${GOOGLEAPI_VERSION}^{commit} 2> /dev/null; then
    echo "Fetching googleapis git objects to find version ${GOOGLEAPI_VERSION}"
    git fetch origin ${GOOGLEAPI_VERSION}
fi

git reset --hard ${GOOGLEAPI_VERSION}

# Overwrite config.proto with the updated version that has DeploymentGroup
mkdir -p google/cloud/config/v1
cp ${REPO_ROOT}/mockgcp/apis/google/cloud/config/v1/config.proto google/cloud/config/v1/config.proto

# Overwrite catalog.proto with the updated version that has MetadataFeed
mkdir -p google/cloud/dataplex/v1
cp ${REPO_ROOT}/mockgcp/apis/google/cloud/dataplex/v1/catalog.proto google/cloud/dataplex/v1/catalog.proto

# Overwrite logging_config.proto with the updated version that has LogScope
mkdir -p google/logging/v2
cp ${REPO_ROOT}/mockgcp/apis/google/logging/v2/logging_config.proto google/logging/v2/logging_config.proto

# Overwrite dlp.proto with the updated version that has ContentPolicy
mkdir -p google/privacy/dlp/v2
cp ${REPO_ROOT}/mockgcp/apis/google/privacy/dlp/v2/dlp.proto google/privacy/dlp/v2/dlp.proto

# Patch monitoring dashboard metrics.proto to include OpsAnalyticsQuery
sed -i '/string prometheus_query = 6;/a \ \ \ \ // A query used to fetch time series with SQL.\n    OpsAnalyticsQuery ops_analytics_query = 8;' google/monitoring/dashboard/v1/metrics.proto
cat >> google/monitoring/dashboard/v1/metrics.proto <<EOF

// A query that produces an aggregated response and supporting data.
message OpsAnalyticsQuery {
  // A SQL query to fetch time series, category series, or numeric series data.
  string sql = 2;
}
EOF

# Patch compute v1 compute.proto to include SecurityPolicyAssociation (which exists in apiv1/computepb) if missing
if ! grep -q "message SecurityPolicyAssociation" google/cloud/compute/v1/compute.proto; then
cat >> google/cloud/compute/v1/compute.proto <<EOF

// The SecurityPolicyAssociation message
message SecurityPolicyAssociation {
  // The resource that the security policy is attached to.
  optional string attachment_id = 103768503;

  // Output only. [Output Only] The display name of the security policy of the association.
  optional string display_name = 4473832;

  // A list of folders to exclude from the security policy.
  repeated string excluded_folders = 310611024;

  // A list of projects to exclude from the security policy.
  repeated string excluded_projects = 283070287;

  // The name for an association.
  optional string name = 3373707;

  // Output only. [Output Only] The security policy ID of the association.
  optional string security_policy_id = 203671817;

  // Output only. [Output Only] The short name of the security policy of the association.
  optional string short_name = 492051566;
}
EOF
fi

# Patch compute v1 compute.proto to include VmExtensionPolicy (which exists in apiv1/computepb) if missing
if ! grep -q "message VmExtensionPolicy " google/cloud/compute/v1/compute.proto; then
cat >> google/cloud/compute/v1/compute.proto <<EOF

// Configuration for a specific VM extension.
message VmExtensionPolicyExtensionPolicy {
  // Optional. The specific version of the extension to install. If not set, the latest
  // version is used.
  optional string pinned_version = 446267249;

  // Optional. String-based configuration data for the extension.
  optional string string_config = 318905136;
}

// A LabelSelector is applied to a VM only if it matches all the specified
// labels.
message VmExtensionPolicyLabelSelector {
  // Optional. A map of key-value pairs representing VM labels.
  map<string, string> inclusion_labels = 501312642;
}

// Defines how to select VMs to apply a zone VM extension policy.
message VmExtensionPolicyInstanceSelector {
  // Optional. LabelSelector selects VMs based on their labels.
  optional VmExtensionPolicyLabelSelector label_selector = 349780170;
}

// Represents a VM Extension Policy resource.
message VmExtensionPolicy {
  // Optional. Output only. [Output Only] Current state of the policy: ACTIVE or DELETING.
  enum State {
    // A value indicating that the enum field is not set.
    UNDEFINED_STATE = 0;

    // The policy is active and applied to matching VMs.
    ACTIVE = 314733318;

    // The policy is in the process of being deleted.
    DELETING = 528602024;

    // Default value. Do not use.
    STATE_UNSPECIFIED = 470755401;
  }

  // Output only. [Output Only] Creation timestamp inRFC3339 text format.
  optional string creation_timestamp = 30525366;

  // An optional description of this resource.
  optional string description = 422937596;

  // Required. A map of extension names (for example, "ops-agent") to their corresponding
  // policy configurations.
  map<string, VmExtensionPolicyExtensionPolicy> extension_policies = 227761488;

  // Optional. Output only. [Output Only] Link to the global policy that manages this zone policy, if
  // applicable.
  optional string global_resource_link = 516569583;

  // Output only. [Output Only] The unique identifier for the resource. This identifier is
  // defined by the server.
  optional uint64 id = 3355;

  // Optional. Selectors to target VMs for this policy.
  repeated VmExtensionPolicyInstanceSelector instance_selectors = 520298826;

  // Output only. [Output Only] Type of the resource. Alwayscompute#vmExtensionPolicy.
  optional string kind = 3292052;

  // Optional. Output only. [Output Only] Indicates if this policy is managed by a global policy.
  optional bool managed_by_global = 88605675;

  // Name of the resource. Provided by the client when the resource is created.
  optional string name = 3373707;

  // Optional. Priority of this policy.
  optional int32 priority = 445151652;

  // Output only. [Output Only] Server-defined fully-qualified URL for this resource.
  optional string self_link = 456214797;

  // Output only. [Output Only] Server-defined URL for this resource's resource id.
  optional string self_link_with_id = 44520962;

  // Optional. Output only. [Output Only] Current state of the policy: ACTIVE or DELETING.
  optional string state = 109757585;

  // Output only. [Output Only] Update timestamp inRFC3339 text format.
  optional string update_timestamp = 120894752;
}

message VmExtensionPolicyList {
  optional string etag = 3123477;
  optional string id = 3355;
  repeated VmExtensionPolicy items = 100526016;
  optional string kind = 3292052;
  optional string next_page_token = 79797525;
  optional string self_link = 456214797;
  repeated string unreachables = 243372063;
  optional Warning warning = 50704284;
}

message DeleteZoneVmExtensionPolicyRequest {
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  optional string request_id = 37109963;
  string vm_extension_policy = 331532730 [(google.api.field_behavior) = REQUIRED];
  string zone = 3744684 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "zone"
  ];
}

message GetZoneVmExtensionPolicyRequest {
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  string vm_extension_policy = 331532730 [(google.api.field_behavior) = REQUIRED];
  string zone = 3744684 [(google.api.field_behavior) = REQUIRED];
}

message InsertZoneVmExtensionPolicyRequest {
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  optional string request_id = 37109963;
  VmExtensionPolicy vm_extension_policy_resource = 515669235 [(google.api.field_behavior) = REQUIRED];
  string zone = 3744684 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "zone"
  ];
}

message ListZoneVmExtensionPoliciesRequest {
  optional string filter = 336120696;
  optional uint32 max_results = 54715419;
  optional string order_by = 160562920;
  optional string page_token = 19994697;
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  optional bool return_partial_success = 517198390;
  string zone = 3744684 [(google.api.field_behavior) = REQUIRED];
}

message UpdateZoneVmExtensionPolicyRequest {
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  optional string request_id = 37109963;
  string vm_extension_policy = 331532730 [(google.api.field_behavior) = REQUIRED];
  VmExtensionPolicy vm_extension_policy_resource = 515669235 [(google.api.field_behavior) = REQUIRED];
  string zone = 3744684 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "zone"
  ];
}

service ZoneVmExtensionPolicies {
  option (google.api.default_host) = "compute.googleapis.com";
  option (google.api.oauth_scopes) =
    "https://www.googleapis.com/auth/compute,"
    "https://www.googleapis.com/auth/cloud-platform";

  rpc Delete(DeleteZoneVmExtensionPolicyRequest) returns (Operation) {
    option (google.api.http) = {
      delete: "/compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies/{vm_extension_policy}"
    };
    option (google.api.method_signature) = "project,zone,vm_extension_policy";
    option (google.cloud.operation_service) = "ZoneOperations";
  }

  rpc Get(GetZoneVmExtensionPolicyRequest) returns (VmExtensionPolicy) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies/{vm_extension_policy}"
    };
    option (google.api.method_signature) = "project,zone,vm_extension_policy";
  }

  rpc Insert(InsertZoneVmExtensionPolicyRequest) returns (Operation) {
    option (google.api.http) = {
      body: "vm_extension_policy_resource"
      post: "/compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies"
    };
    option (google.api.method_signature) = "project,zone,vm_extension_policy_resource";
    option (google.cloud.operation_service) = "ZoneOperations";
  }

  rpc List(ListZoneVmExtensionPoliciesRequest) returns (VmExtensionPolicyList) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies"
    };
    option (google.api.method_signature) = "project,zone";
  }

  rpc Update(UpdateZoneVmExtensionPolicyRequest) returns (Operation) {
    option (google.api.http) = {
      body: "vm_extension_policy_resource"
      patch: "/compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies/{vm_extension_policy}"
    };
    option (google.api.method_signature) = "project,zone,vm_extension_policy,vm_extension_policy_resource";
    option (google.cloud.operation_service) = "ZoneOperations";
  }
}
EOF
fi

# Patch compute v1 compute.proto to include CrossSiteNetwork (which exists in apiv1/computepb) if missing
if ! grep -q "message CrossSiteNetwork " google/cloud/compute/v1/compute.proto; then
cat >> google/cloud/compute/v1/compute.proto <<EOF

// A resource that represents a cross-site network.
message CrossSiteNetwork {
  // Output only. [Output Only] Creation timestamp in RFC3339 text format.
  optional string creation_timestamp = 30525366;

  // An optional description of the cross-site network.
  optional string description = 422937596;

  // Output only. [Output Only] The unique identifier for the resource type. The server generates this identifier.
  optional uint64 id = 3355;

  // Output only. [Output Only] Type of the resource. Always compute#crossSiteNetwork for cross-site networks.
  optional string kind = 3292052;

  // Name of the resource. Provided by the client when the resource is created.
  optional string name = 3373707;

  // Output only. [Output Only] Server-defined URL for the resource.
  optional string self_link = 456214797;
}

// A request message for CrossSiteNetworks.Delete.
message DeleteCrossSiteNetworkRequest {
  // Name of the cross-site network to delete.
  string cross_site_network = 108192469 [(google.api.field_behavior) = REQUIRED];

  // Project ID for this request.
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];

  optional string request_id = 37109963;
}

// A request message for CrossSiteNetworks.Get.
message GetCrossSiteNetworkRequest {
  // Name of the cross-site network to return.
  string cross_site_network = 108192469 [(google.api.field_behavior) = REQUIRED];

  // Project ID for this request.
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
}

// A request message for CrossSiteNetworks.Insert.
message InsertCrossSiteNetworkRequest {
  CrossSiteNetwork cross_site_network_resource = 30896696 [(google.api.field_behavior) = REQUIRED];

  // Project ID for this request.
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];

  optional string request_id = 37109963;
}

// Response to the list request that contains a list of cross-site networks.
message CrossSiteNetworkList {
  optional string etag = 3123477;
  optional string id = 3355;
  repeated CrossSiteNetwork items = 100526016;
  optional string kind = 3292052;
  optional string next_page_token = 79797525;
  optional string self_link = 456214797;
  repeated string unreachables = 243372063;
  optional Warning warning = 50704284;
}

// A request message for CrossSiteNetworks.List.
message ListCrossSiteNetworksRequest {
  optional string filter = 336120696;
  optional uint32 max_results = 54715419;
  optional string order_by = 160562920;
  optional string page_token = 19994697;
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  optional bool return_partial_success = 517198390;
}

// A request message for CrossSiteNetworks.Patch.
message PatchCrossSiteNetworkRequest {
  string cross_site_network = 108192469 [(google.api.field_behavior) = REQUIRED];
  CrossSiteNetwork cross_site_network_resource = 30896696 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  optional string request_id = 37109963;
  optional string update_mask = 500078867;
}

// The CrossSiteNetworks API.
service CrossSiteNetworks {
  option (google.api.default_host) = "compute.googleapis.com";

  rpc Delete(DeleteCrossSiteNetworkRequest) returns (Operation) {
    option (google.api.http) = {
      delete: "/compute/v1/projects/{project}/global/crossSiteNetworks/{cross_site_network}"
    };
    option (google.api.method_signature) = "project,cross_site_network";
    option (google.cloud.operation_service) = "GlobalOperations";
  }

  rpc Get(GetCrossSiteNetworkRequest) returns (CrossSiteNetwork) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/global/crossSiteNetworks/{cross_site_network}"
    };
    option (google.api.method_signature) = "project,cross_site_network";
  }

  rpc Insert(InsertCrossSiteNetworkRequest) returns (Operation) {
    option (google.api.http) = {
      body: "cross_site_network_resource"
      post: "/compute/v1/projects/{project}/global/crossSiteNetworks"
    };
    option (google.api.method_signature) = "project,cross_site_network_resource";
    option (google.cloud.operation_service) = "GlobalOperations";
  }

  rpc List(ListCrossSiteNetworksRequest) returns (CrossSiteNetworkList) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/global/crossSiteNetworks"
    };
    option (google.api.method_signature) = "project";
  }

  rpc Patch(PatchCrossSiteNetworkRequest) returns (Operation) {
    option (google.api.http) = {
      body: "cross_site_network_resource"
      patch: "/compute/v1/projects/{project}/global/crossSiteNetworks/{cross_site_network}"
    };
    option (google.api.method_signature) = "project,cross_site_network,cross_site_network_resource";
    option (google.cloud.operation_service) = "GlobalOperations";
  }
}
EOF
fi

# Patch compute v1 compute.proto to include GlobalVmExtensionPolicy (which exists in discovery) if missing
if ! grep -q "message GlobalVmExtensionPolicy " google/cloud/compute/v1/compute.proto; then
cat >> google/cloud/compute/v1/compute.proto <<EOF

// Configuration for a specific VM extension.
message GlobalVmExtensionPolicyExtensionPolicy {
  // Optional. The specific version of the extension to install. If not set, the latest
  // version is used.
  optional string pinned_version = 446267249;

  // Optional. String-based configuration data for the extension.
  optional string string_config = 318905136;
}

// A LabelSelector is applied to a VM only if it matches all the specified
// labels.
message GlobalVmExtensionPolicyLabelSelector {
  // Optional. A map of key-value pairs representing VM labels.
  map<string, string> inclusion_labels = 501312642;
}

// Defines how to select VMs to apply a global VM extension policy.
message GlobalVmExtensionPolicyInstanceSelector {
  // Optional. LabelSelector selects VMs based on their labels.
  optional GlobalVmExtensionPolicyLabelSelector label_selector = 349780170;
}

// Status of the rollout for a single location.
message GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus {
  // Output only. [Output Only] The state of the location rollout.
  optional string state = 109757585;
}

// Metadata for a rollout operation.
message GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata {
  // Output only. [Output Only] The rollout status for each location.
  map<string, GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus> location_rollout_status = 278036490;

  // Output only. [Output Only] The name of the rollout.
  optional string rollout = 303366577;

  // Output only. [Output Only] The name of the rollout plan.
  optional string rollout_plan = 459057399;

  // Output only. [Output Only] The overall state of the rollout.
  optional string state = 109757585;
}

// Status of the rollout operation.
message GlobalVmExtensionPolicyRolloutOperationRolloutStatus {
  // Output only. [Output Only] The current rollouts for the latest version of the resource.
  repeated GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata current_rollouts = 39829736;

  // Output only. [Output Only] The last completed rollout resource.
  optional GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata previous_rollout = 375749737;
}

// Input for a rollout operation.
message GlobalVmExtensionPolicyRolloutOperationRolloutInput {
  // Optional. Specifies the behavior of the rollout if a conflict is detected.
  optional string conflict_behavior = 115729439;

  // Optional. The name of the rollout plan.
  optional string name = 3373707;

  // Optional. Specifies the predefined rollout plan for the policy.
  optional string predefined_rollout_plan = 364023440;

  // Optional. The UUID that identifies a policy rollout retry attempt.
  optional string retry_uuid = 472838898;
}

// Represents the rollout operation.
message GlobalVmExtensionPolicyRolloutOperation {
  // Required. The rollout input which defines the rollout plan.
  optional GlobalVmExtensionPolicyRolloutOperationRolloutInput rollout_input = 265745340;

  // Output only. [Output Only] The rollout status of the policy.
  optional GlobalVmExtensionPolicyRolloutOperationRolloutStatus rollout_status = 476426816;
}

// Represents a Global VM Extension Policy resource.
message GlobalVmExtensionPolicy {
  // Output only. [Output Only] Creation timestamp inRFC3339 text format.
  optional string creation_timestamp = 30525366;

  // An optional description of this resource.
  optional string description = 422937596;

  // Required. A map of extension names (for example, "ops-agent") to their corresponding
  // policy configurations.
  map<string, GlobalVmExtensionPolicyExtensionPolicy> extension_policies = 227761488;

  // Output only. [Output Only] The unique identifier for the resource. This identifier is
  // defined by the server.
  optional uint64 id = 3355;

  // Optional. Selectors to target VMs for this policy.
  repeated GlobalVmExtensionPolicyInstanceSelector instance_selectors = 520298826;

  // Output only. [Output Only] Type of the resource. Alwayscompute#globalVmExtensionPolicy.
  optional string kind = 3292052;

  // Name of the resource. Provided by the client when the resource is created.
  optional string name = 3373707;

  // Optional. Priority of this policy.
  optional int32 priority = 445151652;

  // Required. The rollout strategy and status.
  optional GlobalVmExtensionPolicyRolloutOperation rollout_operation = 486746201;

  // Output only. [Output Only] The scoped resource status. It's only for tracking the
  // purging status of the policy.
  optional string scoped_resource_status = 181976980;

  // Output only. [Output Only] Server-defined fully-qualified URL for this resource.
  optional string self_link = 456214797;

  // Output only. [Output Only] Server-defined URL for this resource's resource id.
  optional string self_link_with_id = 44520962;

  // Output only. [Output Only] Update timestamp inRFC3339 text format.
  optional string update_timestamp = 120894752;
}

message GlobalVmExtensionPolicyList {
  optional string etag = 3123477;
  optional string id = 3355;
  repeated GlobalVmExtensionPolicy items = 100526016;
  optional string kind = 3292052;
  optional string next_page_token = 79797525;
  optional string self_link = 456214797;
  repeated string unreachables = 243372063;
  optional Warning warning = 50704284;
}

message DeleteGlobalVmExtensionPolicyRequest {
  string global_vm_extension_policy = 8334110 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  optional string request_id = 37109963;
}

message GetGlobalVmExtensionPolicyRequest {
  string global_vm_extension_policy = 8334110 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
}

message InsertGlobalVmExtensionPolicyRequest {
  GlobalVmExtensionPolicy global_vm_extension_policy_resource = 101542927 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  optional string request_id = 37109963;
}

message ListGlobalVmExtensionPoliciesRequest {
  optional string filter = 336120696;
  optional uint32 max_results = 54715419;
  optional string order_by = 160562920;
  optional string page_token = 19994697;
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  optional bool return_partial_success = 517198390;
}

message UpdateGlobalVmExtensionPolicyRequest {
  string global_vm_extension_policy = 8334110 [(google.api.field_behavior) = REQUIRED];
  GlobalVmExtensionPolicy global_vm_extension_policy_resource = 101542927 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  optional string request_id = 37109963;
}

service GlobalVmExtensionPolicies {
  option (google.api.default_host) = "compute.googleapis.com";
  option (google.api.oauth_scopes) =
    "https://www.googleapis.com/auth/compute,"
    "https://www.googleapis.com/auth/cloud-platform";

  rpc Delete(DeleteGlobalVmExtensionPolicyRequest) returns (Operation) {
    option (google.api.http) = {
      delete: "/compute/v1/projects/{project}/global/vmExtensionPolicies/{global_vm_extension_policy}"
    };
    option (google.api.method_signature) = "project,global_vm_extension_policy";
    option (google.cloud.operation_service) = "GlobalOperations";
  }

  rpc Get(GetGlobalVmExtensionPolicyRequest) returns (GlobalVmExtensionPolicy) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/global/vmExtensionPolicies/{global_vm_extension_policy}"
    };
    option (google.api.method_signature) = "project,global_vm_extension_policy";
  }

  rpc Insert(InsertGlobalVmExtensionPolicyRequest) returns (Operation) {
    option (google.api.http) = {
      body: "global_vm_extension_policy_resource"
      post: "/compute/v1/projects/{project}/global/vmExtensionPolicies"
    };
    option (google.api.method_signature) = "project,global_vm_extension_policy_resource";
    option (google.cloud.operation_service) = "GlobalOperations";
  }

  rpc List(ListGlobalVmExtensionPoliciesRequest) returns (GlobalVmExtensionPolicyList) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/global/vmExtensionPolicies"
    };
    option (google.api.method_signature) = "project";
  }

  rpc Update(UpdateGlobalVmExtensionPolicyRequest) returns (Operation) {
    option (google.api.http) = {
      body: "global_vm_extension_policy_resource"
      patch: "/compute/v1/projects/{project}/global/vmExtensionPolicies/{global_vm_extension_policy}"
    };
    option (google.api.method_signature) = "project,global_vm_extension_policy,global_vm_extension_policy_resource";
    option (google.cloud.operation_service) = "GlobalOperations";
  }
}
EOF
fi

# Patch compute v1 compute.proto to include HealthAggregationPolicy and RegionHealthAggregationPolicies (which exists in apiv1/computepb) if missing
if ! grep -q "message HealthAggregationPolicy" google/cloud/compute/v1/compute.proto; then
cat >> google/cloud/compute/v1/compute.proto <<EOF

// The HealthAggregationPolicy message
message HealthAggregationPolicy {
  enum PolicyType {
    UNDEFINED_POLICY_TYPE = 0;
    BACKEND_SERVICE_POLICY = 415582023;
    DNS_PUBLIC_IP_POLICY = 366354794;
  }
  optional string creation_timestamp = 30525366;
  optional string description = 422937596;
  optional string fingerprint = 234678500;
  optional uint32 healthy_percent_threshold = 72242895;
  optional uint64 id = 3355;
  optional string kind = 3292052;
  optional uint32 min_healthy_threshold = 132310940;
  optional string name = 3373707;
  optional string policy_type = 18158119;
  optional string region = 138946292;
  optional string self_link = 456214797;
  optional string self_link_with_id = 44520962;
}

message DeleteRegionHealthAggregationPolicyRequest {
  string health_aggregation_policy = 240314354 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  string region = 138946292 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "region"
  ];
  optional string request_id = 37109963;
}

message GetRegionHealthAggregationPolicyRequest {
  string health_aggregation_policy = 240314354 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];
}

message InsertRegionHealthAggregationPolicyRequest {
  HealthAggregationPolicy health_aggregation_policy_resource = 293526971 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  string region = 138946292 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "region"
  ];
  optional string request_id = 37109963;
}

message PatchRegionHealthAggregationPolicyRequest {
  string health_aggregation_policy = 240314354 [(google.api.field_behavior) = REQUIRED];
  HealthAggregationPolicy health_aggregation_policy_resource = 293526971 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  string region = 138946292 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "region"
  ];
  optional string request_id = 37109963;
}

message ListRegionHealthAggregationPoliciesRequest {
  optional string filter = 3361206;
  optional uint32 max_results = 54715419;
  optional string order_by = 160562920;
  optional string page_token = 19994697;
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];
  optional bool return_partial_success = 517198390;
}

message HealthAggregationPolicyList {
  optional string id = 3355;
  repeated HealthAggregationPolicy items = 100526016;
  optional string kind = 3292052;
  optional string next_page_token = 79797525;
  optional string self_link = 456214797;
  optional Warning warning = 50704284;
}

message TestIamPermissionsRegionHealthAggregationPolicyRequest {
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];
  string resource = 195806222 [(google.api.field_behavior) = REQUIRED];
  TestPermissionsRequest test_permissions_request_resource = 439214758 [(google.api.field_behavior) = REQUIRED];
}

service RegionHealthAggregationPolicies {
  option (google.api.default_host) = "compute.googleapis.com";
  option (google.api.oauth_scopes) =
    "https://www.googleapis.com/auth/compute,"
    "https://www.googleapis.com/auth/cloud-platform";

  rpc Delete(DeleteRegionHealthAggregationPolicyRequest) returns (Operation) {
    option (google.api.http) = {
      delete: "/compute/v1/projects/{project}/regions/{region}/healthAggregationPolicies/{health_aggregation_policy}"
    };
    option (google.api.method_signature) = "project,region,health_aggregation_policy";
    option (google.cloud.operation_service) = "RegionOperations";
  }

  rpc Get(GetRegionHealthAggregationPolicyRequest) returns (HealthAggregationPolicy) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/regions/{region}/healthAggregationPolicies/{health_aggregation_policy}"
    };
    option (google.api.method_signature) = "project,region,health_aggregation_policy";
  }

  rpc Insert(InsertRegionHealthAggregationPolicyRequest) returns (Operation) {
    option (google.api.http) = {
      body: "health_aggregation_policy_resource"
      post: "/compute/v1/projects/{project}/regions/{region}/healthAggregationPolicies"
    };
    option (google.api.method_signature) = "project,region,health_aggregation_policy_resource";
    option (google.cloud.operation_service) = "RegionOperations";
  }

  rpc List(ListRegionHealthAggregationPoliciesRequest) returns (HealthAggregationPolicyList) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/regions/{region}/healthAggregationPolicies"
    };
    option (google.api.method_signature) = "project,region";
  }

  rpc Patch(PatchRegionHealthAggregationPolicyRequest) returns (Operation) {
    option (google.api.http) = {
      body: "health_aggregation_policy_resource"
      patch: "/compute/v1/projects/{project}/regions/{region}/healthAggregationPolicies/{health_aggregation_policy}"
    };
    option (google.api.method_signature) = "project,region,health_aggregation_policy,health_aggregation_policy_resource";
    option (google.cloud.operation_service) = "RegionOperations";
  }

  rpc TestIamPermissions(TestIamPermissionsRegionHealthAggregationPolicyRequest) returns (TestPermissionsResponse) {
    option (google.api.http) = {
      body: "test_permissions_request_resource"
      post: "/compute/v1/projects/{project}/regions/{region}/healthAggregationPolicies/{resource}/testIamPermissions"
    };
    option (google.api.method_signature) = "project,region,resource,test_permissions_request_resource";
  }
}
EOF
fi


if (which protoc); then
    echo "Found protoc version $(protoc --version)"
else
    echo "Installing protoc"
    if [ "$(uname)" == "Darwin" ]; then
      brew install protobuf
    else
      echo "apt update..."
      sudo apt update
      echo "apt install..."
      sudo apt install -y protobuf-compiler
    fi
fi

# Enable nullglob shell option so that unmatched glob patterns (e.g. if certain subdirectories don't contain any .proto files)
# expand to an empty string instead of the literal wildcard string, avoiding protoc errors.
shopt -s nullglob
PROTO_FILES=(
    ${REPO_ROOT}/mockgcp/apis/google/apps/cloudidentity/*/*.proto
    ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/apigee/*/*.proto
    ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/networkconnectivity/*/*.proto
    ${REPO_ROOT}/mockgcp/apis/mockgcp/cloud/servicenetworking/*/*.proto
    ${REPO_ROOT}/mockgcp/apis/google/cloud/binaryauthorization/*/*.proto
    ${REPO_ROOT}/mockgcp/apis/google/iam/*/*.proto
    ${VERSION_DIR}/google/*/*.proto
    ${VERSION_DIR}/google/analytics/*/*/*.proto
    ${VERSION_DIR}/google/privacy/dlp/v2/*.proto
    ${VERSION_DIR}/google/api/*.proto
    ${VERSION_DIR}/google/api/*/*/*.proto
    ${VERSION_DIR}/google/appengine/*/*.proto
    ${VERSION_DIR}/google/bigtable/*/*/*.proto
    ${VERSION_DIR}/google/cloud/bigquery/*/*.proto
    ${VERSION_DIR}/google/cloud/*/*/*.proto
    ${VERSION_DIR}/google/cloud/*/*/*/*.proto
    ${VERSION_DIR}/google/cloud/*/*/*/*/*.proto
    ${VERSION_DIR}/google/dataflow/*/*.proto
    ${VERSION_DIR}/google/firestore/*/*.proto
    ${VERSION_DIR}/google/firestore/*/*/*.proto
    ${VERSION_DIR}/google/iam/v1/*.proto
    ${VERSION_DIR}/google/iam/admin/v1/*.proto
    ${VERSION_DIR}/google/logging/v2/*.proto
    ${VERSION_DIR}/google/monitoring/v3/*.proto
    ${VERSION_DIR}/google/monitoring/metricsscope/v1/*.proto
    ${VERSION_DIR}/google/monitoring/dashboard/v1/*.proto
    ${VERSION_DIR}/google/devtools/cloudbuild/*/*.proto
    ${VERSION_DIR}/google/devtools/artifactregistry/*/*.proto
    ${VERSION_DIR}/google/devtools/testing/*/*.proto
    ${VERSION_DIR}/google/spanner/admin/instance/v1/*.proto
    ${VERSION_DIR}/google/spanner/admin/database/v1/*.proto
    ${VERSION_DIR}/google/storage/control/v2/*.proto
    ${VERSION_DIR}/google/storage/v1/*.proto
    ${VERSION_DIR}/google/pubsub/v1/*.proto
    ${VERSION_DIR}/google/maps/mapmanagement/*/*.proto
    ${VERSION_DIR}/google/cloud/memorystore/v1/*.proto
    ${VERSION_DIR}/google/container/*/*.proto
    ${VERSION_DIR}/google/privacy/dlp/v2/*.proto
    ${VERSION_DIR}/grafeas/v1/*.proto
    ${VERSION_DIR}/google/iam/v3/*.proto
    ${VERSION_DIR}/google/identity/*/*/*.proto
)
# Disable nullglob shell option to restore the default shell globbing behavior
shopt -u nullglob

# Compile the protocols to a single descriptor set binary file (.pb).
# 2> >(...) redirects stderr into process substitution to filter out noisy "Import ... is unused" warnings.
protoc --include_imports --include_source_info \
    --experimental_allow_proto3_optional \
    -I ${VERSION_DIR}/ \
    -I ${REPO_ROOT}/mockgcp/apis \
    "${PROTO_FILES[@]}" \
    -o ${VERSIONED_OUTPUT_PATH} 2> >(grep -v "Import .* is unused" >&2)

if [ "${VERSIONED_OUTPUT_PATH}" != "${OUTPUT_PATH}" ]; then
    cp "${VERSIONED_OUTPUT_PATH}" "${OUTPUT_PATH}"
fi
echo "${SCRIPT_HASH}" > "${VERSIONED_OUTPUT_PATH}.hash"
