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

# Overwrite biglake protos with the updated version that has IcebergCatalog
mkdir -p google/cloud/biglake/v1
cp ${REPO_ROOT}/mockgcp/apis/google/cloud/biglake/v1/* google/cloud/biglake/v1/

# Overwrite catalog.proto with the updated version that has MetadataFeed
mkdir -p google/cloud/dataplex/v1
cp ${REPO_ROOT}/mockgcp/apis/google/cloud/dataplex/v1/catalog.proto google/cloud/dataplex/v1/catalog.proto

# Overwrite backupdr protos with the updated version that has AutoProtectionPolicy and RestoreVerificationPlan
mkdir -p google/cloud/backupdr/v1
cp ${REPO_ROOT}/mockgcp/apis/google/cloud/backupdr/v1/autoprotectionpolicy.proto google/cloud/backupdr/v1/autoprotectionpolicy.proto
cp ${REPO_ROOT}/mockgcp/apis/google/cloud/backupdr/v1/backupdr.proto google/cloud/backupdr/v1/backupdr.proto
cp ${REPO_ROOT}/mockgcp/apis/google/cloud/backupdr/v1/restore_verification_plan.proto google/cloud/backupdr/v1/restore_verification_plan.proto

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

# Patch compute v1 compute.proto to include HealthSource and RegionHealthSources (which exists in apiv1/computepb) if missing
if ! grep -q "message HealthSource " google/cloud/compute/v1/compute.proto; then
cat >> google/cloud/compute/v1/compute.proto <<EOF

// The HealthSource message
message HealthSource {
  enum SourceType {
    UNDEFINED_SOURCE_TYPE = 0;
    BACKEND_SERVICE = 285792202;
  }
  optional string creation_timestamp = 30525366;
  optional string description = 422937596;
  optional string fingerprint = 234678500;
  optional string health_aggregation_policy = 240314354;
  optional uint64 id = 3355;
  optional string kind = 3292052;
  optional string name = 3373707;
  optional string region = 138946292;
  optional string self_link = 456214797;
  optional string self_link_with_id = 44520962;
  optional string source_type = 452245726;
  repeated string sources = 125606840;
}

message DeleteRegionHealthSourceRequest {
  string health_source = 376521566 [(google.api.field_behavior) = REQUIRED];
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

message GetRegionHealthSourceRequest {
  string health_source = 376521566 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];
}

message InsertRegionHealthSourceRequest {
  HealthSource health_source_resource = 357735375 [(google.api.field_behavior) = REQUIRED];
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

message PatchRegionHealthSourceRequest {
  string health_source = 376521566 [(google.api.field_behavior) = REQUIRED];
  HealthSource health_source_resource = 357735375 [(google.api.field_behavior) = REQUIRED];
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

message ListRegionHealthSourcesRequest {
  optional string filter = 336120696;
  optional uint32 max_results = 54715419;
  optional string order_by = 160562920;
  optional string page_token = 19994697;
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];
  optional bool return_partial_success = 517198390;
}

message HealthSourceList {
  optional string id = 3355;
  repeated HealthSource items = 100526016;
  optional string kind = 3292052;
  optional string next_page_token = 79797525;
  optional string self_link = 456214797;
  optional Warning warning = 50704284;
}

message TestIamPermissionsRegionHealthSourceRequest {
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];
  string resource = 195806222 [(google.api.field_behavior) = REQUIRED];
  TestPermissionsRequest test_permissions_request_resource = 439214758 [(google.api.field_behavior) = REQUIRED];
}

service RegionHealthSources {
  option (google.api.default_host) = "compute.googleapis.com";
  option (google.api.oauth_scopes) =
    "https://www.googleapis.com/auth/compute,"
    "https://www.googleapis.com/auth/cloud-platform";

  rpc Delete(DeleteRegionHealthSourceRequest) returns (Operation) {
    option (google.api.http) = {
      delete: "/compute/v1/projects/{project}/regions/{region}/healthSources/{health_source}"
    };
    option (google.api.method_signature) = "project,region,health_source";
    option (google.cloud.operation_service) = "RegionOperations";
  }

  rpc Get(GetRegionHealthSourceRequest) returns (HealthSource) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/regions/{region}/healthSources/{health_source}"
    };
    option (google.api.method_signature) = "project,region,health_source";
  }

  rpc Insert(InsertRegionHealthSourceRequest) returns (Operation) {
    option (google.api.http) = {
      body: "health_source_resource"
      post: "/compute/v1/projects/{project}/regions/{region}/healthSources"
    };
    option (google.api.method_signature) = "project,region,health_source_resource";
    option (google.cloud.operation_service) = "RegionOperations";
  }

  rpc List(ListRegionHealthSourcesRequest) returns (HealthSourceList) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/regions/{region}/healthSources"
    };
    option (google.api.method_signature) = "project,region";
  }

  rpc Patch(PatchRegionHealthSourceRequest) returns (Operation) {
    option (google.api.http) = {
      body: "health_source_resource"
      patch: "/compute/v1/projects/{project}/regions/{region}/healthSources/{health_source}"
    };
    option (google.api.method_signature) = "project,region,health_source,health_source_resource";
    option (google.cloud.operation_service) = "RegionOperations";
  }

  rpc TestIamPermissions(TestIamPermissionsRegionHealthSourceRequest) returns (TestPermissionsResponse) {
    option (google.api.http) = {
      body: "test_permissions_request_resource"
      post: "/compute/v1/projects/{project}/regions/{region}/healthSources/{resource}/testIamPermissions"
    };
    option (google.api.method_signature) = "project,region,resource,test_permissions_request_resource";
  }
}
EOF
fi

# Patch compute v1 compute.proto to include InstantSnapshotGroup (which exists in apiv1/computepb) if missing
if ! grep -q "message InstantSnapshotGroup " google/cloud/compute/v1/compute.proto; then
cat >> google/cloud/compute/v1/compute.proto <<EOF

// Represents an InstantSnapshotGroup resource.
//
// An instant snapshot group is a set of instant snapshots that represents a
// point in time state of a consistency group.
message InstantSnapshotGroup {
  enum Status {
    UNDEFINED_STATUS = 0;
    CREATING = 455564985;
    DELETING = 528602024;
    FAILED = 455706685;
    INVALID = 530283991;
    READY = 77848963;
    UNKNOWN = 433141802;
  }

  // Output only. [Output Only] Creation timestamp in RFC3339 text format.
  optional string creation_timestamp = 30525366;

  // Optional. An optional description of this resource. Provide this property when you create the resource.
  optional string description = 422937596;

  // Output only. [Output Only] The unique identifier for the resource. This identifier is defined by the server.
  optional uint64 id = 3355;

  // Output only. [Output Only] Type of the resource. Always compute#instantSnapshotGroup for InstantSnapshotGroup resources.
  optional string kind = 3292052;

  // Identifier. Name of the resource; provided by the client when the resource is created.
  // The name must be 1-63 characters long, and comply with RFC1035.
  // Specifically, the name must be 1-63 characters long and match the regular
  // expression \`[a-z]([-a-z0-9]*[a-z0-9])?\` which means the first
  // character must be a lowercase letter, and all following characters must be
  // a dash, lowercase letter, or digit, except the last character, which cannot
  // be a dash.
  optional string name = 3373707;

  // Output only. [Output Only] URL of the region where the instant snapshot group resides.
  // You must specify this field as part of the HTTP request URL. It is
  // not settable as a field in the request body.
  optional string region = 138946292;

  optional InstantSnapshotGroupResourceStatus resource_status = 249429315;

  // Output only. [Output Only] Server-defined URL for the resource.
  optional string self_link = 456214797;

  // Output only. [Output Only] Server-defined URL for this resource's resource id.
  optional string self_link_with_id = 44520962;

  optional string source_consistency_group = 531359348;

  // Output only. [Output Only] Check the Status enum for the list of possible values.
  optional string status = 181260274;

  // Output only. [Output Only] URL of the zone where the instant snapshot group resides.
  // You must specify this field as part of the HTTP request URL. It is
  // not settable as a field in the request body.
  optional string zone = 3744684;
}

message InstantSnapshotGroupParameters {
  // The source instant snapshot group used to create disks. You can provide
  // this as a partial or full URL to the resource. For example, the following
  // are valid values:
  //
  //   - https://www.googleapis.com/compute/v1/projects/project/zones/zone/instantSnapshotGroups/instantSnapshotGroup
  //   - projects/project/zones/zone/instantSnapshotGroups/instantSnapshotGroup
  //   - zones/zone/instantSnapshotGroups/instantSnapshotGroup
  optional string source_instant_snapshot_group = 525991686;
}

message InstantSnapshotGroupResourceStatus {
  // Output only. [Output Only]
  optional string consistency_membership_resolution_time = 346675870;

  // Output only. [Output Only]
  optional InstantSnapshotGroupSourceInfo source_info = 451907154;
}

message InstantSnapshotGroupSourceInfo {
  optional string consistency_group = 508183960;
  optional string consistency_group_id = 85599522;
}

message ListInstantSnapshotGroups {
  optional string etag = 3123477;
  optional string id = 3355;
  repeated InstantSnapshotGroup items = 100526016;
  optional string kind = 3292052;
  optional string next_page_token = 79797525;
  optional string self_link = 456214797;
  repeated string unreachables = 243372063;
  optional Warning warning = 50704284;
}

message DeleteInstantSnapshotGroupRequest {
  string instant_snapshot_group = 223180386 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  optional string request_id = 37109963;
  string zone = 3744684 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "zone"
  ];
}

message GetInstantSnapshotGroupRequest {
  string instant_snapshot_group = 223180386 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  string zone = 3744684 [(google.api.field_behavior) = REQUIRED];
}

message InsertInstantSnapshotGroupRequest {
  InstantSnapshotGroup instant_snapshot_group_resource = 145067339 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  optional string request_id = 37109963;
  optional string source_consistency_group = 531359348;
  string zone = 3744684 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "zone"
  ];
}

message ListInstantSnapshotGroupsRequest {
  optional string filter = 336120696;
  optional uint32 max_results = 54715419;
  optional string order_by = 160562920;
  optional string page_token = 19994697;
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  optional bool return_partial_success = 517198390;
  string zone = 3744684 [(google.api.field_behavior) = REQUIRED];
}

service InstantSnapshotGroups {
  option (google.api.default_host) = "compute.googleapis.com";
  option (google.api.oauth_scopes) =
    "https://www.googleapis.com/auth/compute,"
    "https://www.googleapis.com/auth/cloud-platform";

  rpc Delete(DeleteInstantSnapshotGroupRequest) returns (Operation) {
    option (google.api.http) = {
      delete: "/compute/v1/projects/{project}/zones/{zone}/instantSnapshotGroups/{instant_snapshot_group}"
    };
    option (google.api.method_signature) = "project,zone,instant_snapshot_group";
    option (google.cloud.operation_service) = "ZoneOperations";
  }

  rpc Get(GetInstantSnapshotGroupRequest) returns (InstantSnapshotGroup) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/zones/{zone}/instantSnapshotGroups/{instant_snapshot_group}"
    };
    option (google.api.method_signature) = "project,zone,instant_snapshot_group";
  }

  rpc Insert(InsertInstantSnapshotGroupRequest) returns (Operation) {
    option (google.api.http) = {
      body: "instant_snapshot_group_resource"
      post: "/compute/v1/projects/{project}/zones/{zone}/instantSnapshotGroups"
    };
    option (google.api.method_signature) = "project,zone,instant_snapshot_group_resource";
    option (google.cloud.operation_service) = "ZoneOperations";
  }

  rpc List(ListInstantSnapshotGroupsRequest) returns (ListInstantSnapshotGroups) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/zones/{zone}/instantSnapshotGroups"
    };
    option (google.api.method_signature) = "project,zone";
  }
}

message DeleteRegionInstantSnapshotGroupRequest {
  string instant_snapshot_group = 223180386 [(google.api.field_behavior) = REQUIRED];
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

message GetRegionInstantSnapshotGroupRequest {
  string instant_snapshot_group = 223180386 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];
}

message InsertRegionInstantSnapshotGroupRequest {
  InstantSnapshotGroup instant_snapshot_group_resource = 145067339 [(google.api.field_behavior) = REQUIRED];
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];
  string region = 138946292 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "region"
  ];
  optional string request_id = 37109963;
  optional string source_consistency_group = 531359348;
}

message ListRegionInstantSnapshotGroupsRequest {
  optional string filter = 336120696;
  optional uint32 max_results = 54715419;
  optional string order_by = 160562920;
  optional string page_token = 19994697;
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];
  optional bool return_partial_success = 517198390;
}

service RegionInstantSnapshotGroups {
  option (google.api.default_host) = "compute.googleapis.com";
  option (google.api.oauth_scopes) =
    "https://www.googleapis.com/auth/compute,"
    "https://www.googleapis.com/auth/cloud-platform";

  rpc Delete(DeleteRegionInstantSnapshotGroupRequest) returns (Operation) {
    option (google.api.http) = {
      delete: "/compute/v1/projects/{project}/regions/{region}/instantSnapshotGroups/{instant_snapshot_group}"
    };
    option (google.api.method_signature) = "project,region,instant_snapshot_group";
    option (google.cloud.operation_service) = "RegionOperations";
  }

  rpc Get(GetRegionInstantSnapshotGroupRequest) returns (InstantSnapshotGroup) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/regions/{region}/instantSnapshotGroups/{instant_snapshot_group}"
    };
    option (google.api.method_signature) = "project,region,instant_snapshot_group";
  }

  rpc Insert(InsertRegionInstantSnapshotGroupRequest) returns (Operation) {
    option (google.api.http) = {
      body: "instant_snapshot_group_resource"
      post: "/compute/v1/projects/{project}/regions/{region}/instantSnapshotGroups"
    };
    option (google.api.method_signature) = "project,region,instant_snapshot_group_resource";
    option (google.cloud.operation_service) = "RegionOperations";
  }

  rpc List(ListRegionInstantSnapshotGroupsRequest) returns (ListInstantSnapshotGroups) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/regions/{region}/instantSnapshotGroups"
    };
    option (google.api.method_signature) = "project,region";
  }
}
EOF
fi


# Patch compute v1 compute.proto to include RolloutPlan and RolloutPlans if missing
if ! grep -q "message RolloutPlan " google/cloud/compute/v1/compute.proto; then
cat >> google/cloud/compute/v1/compute.proto <<'EOF'

// A request message for RolloutPlans.Delete. See the method description for details.
message DeleteRolloutPlanRequest {
  // Project ID for this request.
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];

  // An optional request ID to identify requests. Specify a unique request ID so
  // that if you must retry your request, the server will know to ignore the
  // request if it has already been completed.
  //
  // For example, consider a situation where you make an initial request and
  // the request times out. If you make the request again with the same
  // request ID, the server can check if original operation with the same
  // request ID was received, and if so, will ignore the second request. This
  // prevents clients from accidentally creating duplicate commitments.
  //
  // The request ID must be
  // a valid UUID with the exception that zero UUID is not supported
  // (00000000-0000-0000-0000-000000000000).
  optional string request_id = 37109963;

  // Name of the RolloutPlan resource to delete.
  string rollout_plan = 459057399 [(google.api.field_behavior) = REQUIRED];

}

// A request message for RolloutPlans.Get. See the method description for details.
message GetRolloutPlanRequest {
  // Project ID for this request.
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];

  // Name of the persistent rollout plan to return.
  string rollout_plan = 459057399 [(google.api.field_behavior) = REQUIRED];

}

// A request message for RolloutPlans.Insert. See the method description for details.
message InsertRolloutPlanRequest {
  // Project ID for this request.
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];

  // An optional request ID to identify requests. Specify a unique request ID so
  // that if you must retry your request, the server will know to ignore the
  // request if it has already been completed.
  //
  // For example, consider a situation where you make an initial request and
  // the request times out. If you make the request again with the same
  // request ID, the server can check if original operation with the same
  // request ID was received, and if so, will ignore the second request. This
  // prevents clients from accidentally creating duplicate commitments.
  //
  // The request ID must be
  // a valid UUID with the exception that zero UUID is not supported
  // (00000000-0000-0000-0000-000000000000).
  optional string request_id = 37109963;

  // The body resource for this request
  RolloutPlan rollout_plan_resource = 38629974 [(google.api.field_behavior) = REQUIRED];

}

// A request message for RolloutPlans.List. See the method description for details.
message ListRolloutPlansRequest {
  // A filter expression that filters resources listed in the response. Most
  // Compute resources support two types of filter expressions:
  // expressions that support regular expressions and expressions that follow
  // API improvement proposal AIP-160.
  // These two types of filter expressions cannot be mixed in one request.
  //
  // If you want to use AIP-160, your expression must specify the field name, an
  // operator, and the value that you want to use for filtering. The value
  // must be a string, a number, or a boolean. The operator
  // must be either `=`, `!=`, `>`, `<`, `<=`, `>=` or `:`.
  //
  // For example, if you are filtering Compute Engine instances, you can
  // exclude instances named `example-instance` by specifying
  // `name != example-instance`.
  //
  // The `:*` comparison can be used to test whether a key has been defined.
  // For example, to find all objects with `owner` label use:
  // ```
  // labels.owner:*
  // ```
  //
  // You can also filter nested fields. For example, you could specify
  // `scheduling.automaticRestart = false` to include instances only
  // if they are not scheduled for automatic restarts. You can use filtering
  // on nested fields to filter based onresource labels.
  //
  // To filter on multiple expressions, provide each separate expression within
  // parentheses. For example:
  // ```
  // (scheduling.automaticRestart = true)
  // (cpuPlatform = "Intel Skylake")
  // ```
  // By default, each expression is an `AND` expression. However, you
  // can include `AND` and `OR` expressions explicitly.
  // For example:
  // ```
  // (cpuPlatform = "Intel Skylake") OR
  // (cpuPlatform = "Intel Broadwell") AND
  // (scheduling.automaticRestart = true)
  // ```
  //
  // If you want to use a regular expression, use the `eq` (equal) or `ne`
  // (not equal) operator against a single un-parenthesized expression with or
  // without quotes or against multiple parenthesized expressions. Examples:
  //
  // `fieldname eq unquoted literal`
  // `fieldname eq 'single quoted literal'`
  // `fieldname eq "double quoted literal"`
  // `(fieldname1 eq literal) (fieldname2 ne "literal")`
  //
  // The literal value is interpreted as a regular expression using GoogleRE2 library syntax.
  // The literal value must match the entire field.
  //
  // For example, to filter for instances that do not end with name "instance",
  // you would use `name ne .*instance`.
  //
  // You cannot combine constraints on multiple fields using regular
  // expressions.
  optional string filter = 336120696;

  // The maximum number of results per page that should be returned.
  // If the number of available results is larger than `maxResults`,
  // Compute Engine returns a `nextPageToken` that can be used to get
  // the next page of results in subsequent list requests. Acceptable values are
  // `0` to `500`, inclusive. (Default: `500`)
  optional uint32 max_results = 54715419;

  // Sorts list results by a certain order. By default, results
  // are returned in alphanumerical order based on the resource name.
  //
  // You can also sort results in descending order based on the creation
  // timestamp using `orderBy="creationTimestamp desc"`. This sorts
  // results based on the `creationTimestamp` field in
  // reverse chronological order (newest result first). Use this to sort
  // resources like operations so that the newest operation is returned first.
  //
  // Currently, only sorting by `name` or
  // `creationTimestamp desc` is supported.
  optional string order_by = 160562920;

  // Specifies a page token to use. Set `pageToken` to the
  // `nextPageToken` returned by a previous list request to get
  // the next page of results.
  optional string page_token = 19994697;

  // Project ID for this request.
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];

  // Opt-in for partial success behavior which provides partial results in case
  // of failure. The default value is false.
  //
  // For example, when partial success behavior is enabled, aggregatedList for a
  // single zone scope either returns all resources in the zone or no resources,
  // with an error code.
  optional bool return_partial_success = 517198390;

}

// RolloutPlan resource.
//
// A RolloutPlan is the customer-defined strategy to divide a large-scale change
// into smaller increments, referred to as "waves". Each wave targets a specific
// portion of the overall affected area and defines criteria that must be met
// before progressing to the subsequent wave.
message RolloutPlan {
  // The location scope of the rollout plan. If not specified, the location
  // scope is considered as ZONAL.
  enum LocationScope {
    // A value indicating that the enum field is not set.
    UNDEFINED_LOCATION_SCOPE = 0;

    // Unspecified value. Considered as ZONAL.
    LOCATION_SCOPE_UNSPECIFIED = 416410498;

    // Regional scope.
    REGIONAL = 92288543;

    // Zonal scope.
    ZONAL = 85547428;

  }

  // Output only. [Output Only] Creation timestamp inRFC3339
  // text format.
  optional string creation_timestamp = 30525366;

  // An optional description of this resource. Provide this property when you
  // create the resource.
  optional string description = 422937596;

  // Output only. [Output Only] The unique identifier for the resource. This identifier is
  // defined by the server.
  optional uint64 id = 3355;

  // Output only. [Output Only] Type of the resource. Always compute#rolloutPlan
  // for rolloutPlans.
  optional string kind = 3292052;

  // The location scope of the rollout plan. If not specified, the location
  // scope is considered as ZONAL.
  // Check the LocationScope enum for the list of possible values.
  optional string location_scope = 138667338;

  // Name of the resource. Provided by the client when the resource is created.
  // The name must be 1-63 characters long, and comply withRFC1035.
  // Specifically, the name must be 1-63 characters long and match the regular
  // expression `[a-z]([-a-z0-9]*[a-z0-9])?`
  // which means the first character must be a lowercase letter, and all
  // following characters must be a dash, lowercase letter, or digit, except
  // the last character, which cannot be a dash.
  optional string name = 3373707;

  // Output only. [Output Only] Server-defined fully-qualified URL for this resource.
  optional string self_link = 456214797;

  // Output only. [Output Only] Server-defined URL for this resource's resource id.
  optional string self_link_with_id = 44520962;

  // Required. The waves included in this rollout plan.
  repeated RolloutPlanWave waves = 112905370;

}

// A single wave in a rollout plan.
message RolloutPlanWave {
  // Optional. The display name of this wave of the rollout plan.
  optional string display_name = 4473832;

  // Output only. The wave number.
  optional int64 number = 39377737;

  // Optional. The orchestration options for this wave.
  optional RolloutPlanWaveOrchestrationOptions orchestration_options = 93218464;

  // Required. The selectors for this wave. There is a logical AND between each selector
  // defined in a wave, so a resource must satisfy the criteria of *all* the
  // specified selectors to be in scope for the wave.
  repeated RolloutPlanWaveSelector selectors = 431523956;

  // Required. The validation to be performed at the end of this wave.
  optional RolloutPlanWaveValidation validation = 493308025;

}

// Options to control the pace of orchestration of a wave. These options are
// required only if the resource being rolled out follows the Orchestrated
// pattern.
message RolloutPlanWaveOrchestrationOptions {
  // Optional. Delays, if any, to be added between batches of projects. We allow
  // multiple Delays to be specified, letting users set separate delays
  // between batches of projects corresponding to different locations and
  // batches of projects corresponding to the same location.
  repeated RolloutPlanWaveOrchestrationOptionsDelay delays = 275150672;

  // Optional. Maximum number of locations to be orchestrated in parallel.
  optional int64 max_concurrent_locations = 493763281;

  // Optional. Maximum number of resources to be orchestrated per location in
  // parallel.
  optional int64 max_concurrent_resources_per_location = 178401246;

}

// Options to control the delay, if any, between batches of projects.
message RolloutPlanWaveOrchestrationOptionsDelay {
  // Optional. Controls whether the delay should only be added between batches of
  // projects corresponding to different locations, or also between
  // batches of projects corresponding to the same location.
  //
  // Must be set to DELIMITER_UNSPECIFIED if no delay is to be added.
  enum Delimiter {
    // A value indicating that the enum field is not set.
    UNDEFINED_DELIMITER = 0;

    // The delay will also be added between batches of projects
    // corresponding to the same location.
    DELIMITER_BATCH = 253390530;

    // The delay will only be added between batches of projects
    // corresponding to different locations.
    DELIMITER_LOCATION = 189398797;

    // No delay will be added between batches of projects. Processing will
    // continue with the next batch as soon as the previous batch of LROs
    // is done.
    DELIMITER_UNSPECIFIED = 372199615;

  }

  // Optional. Controls whether the specified duration is to be added at the end of
  // each batch, or if the total processing time for each batch will be
  // padded if needed to meet the specified duration.
  //
  // Must be set to TYPE_UNSPECIFIED if no delay is to be added.
  enum Type {
    // A value indicating that the enum field is not set.
    UNDEFINED_TYPE = 0;

    // The total processing time for each batch of projects will be padded
    // if needed to meet the specified delay duration.
    TYPE_MINIMUM = 234909673;

    // The specified delay will directly be added after each batch of
    // projects as specified by the delimiter.
    TYPE_OFFSET = 165747000;

    // No delay will be added between batches of projects. Processing will
    // continue with the next batch as soon as the previous batch of LROs
    // is done.
    TYPE_UNSPECIFIED = 437714322;

  }

  // Optional. Controls whether the delay should only be added between batches of
  // projects corresponding to different locations, or also between
  // batches of projects corresponding to the same location.
  //
  // Must be set to DELIMITER_UNSPECIFIED if no delay is to be added.
  // Check the Delimiter enum for the list of possible values.
  optional string delimiter = 286352903;

  // Optional. The duration of the delay, if any, to be added between batches of
  // projects. A zero duration corresponds to no delay.
  optional string duration = 155471252;

  // Optional. Controls whether the specified duration is to be added at the end of
  // each batch, or if the total processing time for each batch will be
  // padded if needed to meet the specified duration.
  //
  // Must be set to TYPE_UNSPECIFIED if no delay is to be added.
  // Check the Type enum for the list of possible values.
  optional string type = 3575610;

}

// A selector which specifies what resource(s) are included in a given wave.
message RolloutPlanWaveSelector {
  // Optional. Roll out to resources by Cloud locations.
  optional RolloutPlanWaveSelectorLocationSelector location_selector = 422273321;

  // Optional. Roll out to resources by Cloud Resource Manager resource hierarchy.
  optional RolloutPlanWaveSelectorResourceHierarchySelector resource_hierarchy_selector = 66779770;

}

// Roll out to resources by location.
message RolloutPlanWaveSelectorLocationSelector {
  // Optional. Example: "us-central1-a"
  repeated string included_locations = 78239643;

}

// Roll out to resources by Cloud Resource Manager resource hierarchy
// nodes such as projects, folders, orgs.
message RolloutPlanWaveSelectorResourceHierarchySelector {
  // Optional. Format: "folders/{folder_id}"
  repeated string included_folders = 372825410;

  // Optional. Format: "organizations/{organization_id}"
  repeated string included_organizations = 318362461;

  // Optional. Format: "projects/{project_id}"
  repeated string included_projects = 64232605;

}

// The validation to be performed before progressing to the next wave.
message RolloutPlanWaveValidation {
  // Optional. Metadata required if type = "time".
  optional RolloutPlanWaveValidationTimeBasedValidationMetadata time_based_validation_metadata = 69499575;

  // Required. The type of the validation. If a type of validation is associated with
  // a metadata object, the appropriate metadata field mapping to the
  // validation type must be provided in the validation message. Possible
  // values are in quotes below alongside an explanation:
  //   "manual": The system waits for an end-user approval API before
  //     progressing to the next wave.
  //   "time": The system waits for a user specified duration before
  //     progressing to the next wave. TimeBasedValidation must be provided.
  optional string type = 3575610;

}

// Metadata required if type = "time".
message RolloutPlanWaveValidationTimeBasedValidationMetadata {
  // Optional. The duration that the system waits in between waves. This wait starts
  // after all changes in the wave are rolled out.
  optional string wait_duration = 452313374;

}

// Contains a list of RolloutPlan resources.
message RolloutPlansListResponse {
  optional string etag = 3123477;

  // [Output Only] Unique identifier for the resource; defined by the server.
  optional string id = 3355;

  // A list of RolloutPlan resources.
  repeated RolloutPlan items = 100526016;

  // [Output Only] This token allows you to get the next page of results for
  // list requests. If the number of results is larger thanmaxResults, use the nextPageToken as a value for
  // the query parameter pageToken in the next list request.
  // Subsequent list requests will have their own nextPageToken to
  // continue paging through the results.
  optional string next_page_token = 79797525;

  // Output only. [Output Only] Server-defined URL for this resource.
  optional string self_link = 456214797;

  // Output only. [Output Only] Unreachable resources.
  // end_interface: MixerListResponseWithEtagBuilder
  repeated string unreachables = 243372063;

  // [Output Only] Informational warning message.
  optional Warning warning = 50704284;

}

// The RolloutPlans API.
service RolloutPlans {
  option (google.api.default_host) =
    "compute.googleapis.com";

  option (google.api.oauth_scopes) =
    "https://www.googleapis.com/auth/compute,"
    "https://www.googleapis.com/auth/cloud-platform";

  // Deletes a RolloutPlan.
  rpc Delete(DeleteRolloutPlanRequest) returns (Operation) {
    option (google.api.http) = {
      delete: "/compute/v1/projects/{project}/global/rolloutPlans/{rollout_plan}"
    };
    option (google.api.method_signature) = "project,rollout_plan";
    option (google.cloud.operation_service) = "GlobalOperations";
  }

  // Gets details of a single project-scoped RolloutPlan.
  rpc Get(GetRolloutPlanRequest) returns (RolloutPlan) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/global/rolloutPlans/{rollout_plan}"
    };
    option (google.api.method_signature) = "project,rollout_plan";
  }

  // Creates a new RolloutPlan in a given project and location.
  rpc Insert(InsertRolloutPlanRequest) returns (Operation) {
    option (google.api.http) = {
      body: "rollout_plan_resource"
      post: "/compute/v1/projects/{project}/global/rolloutPlans"
    };
    option (google.api.method_signature) = "project,rollout_plan_resource";
    option (google.cloud.operation_service) = "GlobalOperations";
  }

  // Lists RolloutPlans in a given project and location.
  rpc List(ListRolloutPlansRequest) returns (RolloutPlansListResponse) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/global/rolloutPlans"
    };
    option (google.api.method_signature) = "project";
  }

}
EOF
fi

# Patch compute v1 compute.proto to include CompositeHealthCheck and RegionCompositeHealthChecks if missing
if ! grep -q "message CompositeHealthCheck " google/cloud/compute/v1/compute.proto; then
cat >> google/cloud/compute/v1/compute.proto <<'EOF'

// Represents a composite health check.
//
// A composite health check resource specifies the health source resources and
// the health destination resource to which the aggregated health result from
// the health source resources is delivered.
message CompositeHealthCheck {
  // Output only. [Output Only] Creation timestamp inRFC3339
  // text format.
  optional string creation_timestamp = 30525366;

  // An optional description of this resource. Provide this property when you
  // create the resource.
  optional string description = 422937596;

  // Fingerprint of this resource. A hash of the contents stored in this object.
  // This field is used in optimistic locking. This field will be ignored when
  // inserting a CompositeHealthCheck. An up-to-date fingerprint
  // must be provided in order to patch the CompositeHealthCheck; Otherwise,
  // the request will fail with error 412 conditionNotMet. To see
  // the latest fingerprint, make a get() request to retrieve the
  // CompositeHealthCheck.
  optional string fingerprint = 234678500;

  // URL to the destination resource. Must be set. Must be aForwardingRule. The ForwardingRule must have
  // load balancing scheme INTERNAL orINTERNAL_MANAGED and must be regional and in the same region
  // as the CompositeHealthCheck (cross-region deployment forINTERNAL_MANAGED is not supported). Can be mutated.
  optional string health_destination = 311098027;

  // URLs to the HealthSource resources whose results are AND'ed.
  // I.e. he aggregated result is is HEALTHY only if all sources
  // are HEALTHY. Must have at least 1. Must not have more than 10.
  // Must be regional and in the same region as theCompositeHealthCheck. Can be mutated.
  repeated string health_sources = 397879509;

  // Output only. [Output Only] A unique identifier for this resource type. The server
  // generates this identifier.
  optional uint64 id = 3355;

  // Output only. [Output Only] Type of the resource. Alwayscompute#compositeHealthCheck for composite health checks.
  optional string kind = 3292052;

  // Name of the resource. Provided by the client when the resource is created.
  // The name must be 1-63 characters long, and comply withRFC1035.
  // Specifically, the name must be 1-63 characters long and match the regular
  // expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first
  // character must be a lowercase letter, and all following characters must
  // be a dash, lowercase letter, or digit, except the last character, which
  // cannot be a dash.
  optional string name = 3373707;

  // Output only. [Output Only] URL of the region where the composite health check resides.
  // This field applies only to the regional resource. You must specify this
  // field as part of the HTTP request URL. It is not settable as a field in
  // the request body.
  optional string region = 138946292;

  // Output only. [Output Only] Server-defined URL for the resource.
  optional string self_link = 456214797;

  // Output only. [Output Only] Server-defined URL with id for the resource.
  optional string self_link_with_id = 44520962;

}

// Contains a list of CompositeHealthChecksScopedList.
message CompositeHealthCheckAggregatedList {
  // [Output Only] Unique identifier for the resource; defined by the server.
  optional string id = 3355;

  // A list of CompositeHealthChecksScopedList resources.
  map<string, CompositeHealthChecksScopedList> items = 100526016;

  // Output only. Type of resource.
  optional string kind = 3292052;

  // [Output Only] This token allows you to get the next page of results for
  // list requests. If the number of results is larger thanmaxResults, use the nextPageToken as a value for
  // the query parameter pageToken in the next list request.
  // Subsequent list requests will have their own nextPageToken to
  // continue paging through the results.
  optional string next_page_token = 79797525;

  // Output only. [Output Only] Server-defined URL for this resource.
  optional string self_link = 456214797;

  // Output only. [Output Only] Unreachable resources.
  repeated string unreachables = 243372063;

  // [Output Only] Informational warning message.
  optional Warning warning = 50704284;

}

// Response message for RegionCompositeHealthChecks.GetHealth
message CompositeHealthCheckHealth {
  // Health state of the CompositeHealthCheck.
  enum HealthState {
    // A value indicating that the enum field is not set.
    UNDEFINED_HEALTH_STATE = 0;

    HEALTHY = 439801213;

    UNHEALTHY = 462118084;

    UNKNOWN = 433141802;

  }

  // Health sources and their corresponding health states.
  repeated CompositeHealthChecksGetHealthResponseHealthSourceHealth health_sources = 397879509;

  // Health state of the CompositeHealthCheck.
  // Check the HealthState enum for the list of possible values.
  optional string health_state = 324007150;

  // Output only. [Output Only] Type of resource. Alwayscompute#compositeHealthCheckHealth for the health of
  // composite health checks.
  optional string kind = 3292052;

}

//
message CompositeHealthCheckList {
  // [Output Only] Unique identifier for the resource; defined by the server.
  optional string id = 3355;

  // A list of CompositeHealthCheck resources.
  repeated CompositeHealthCheck items = 100526016;

  // Output only. [Output Only] Type of the resource. Alwayscompute#compositeHealthCheck for composite health checks.
  optional string kind = 3292052;

  // [Output Only] This token allows you to get the next page of results for
  // list requests. If the number of results is larger thanmaxResults, use the nextPageToken as a value for
  // the query parameter pageToken in the next list request.
  // Subsequent list requests will have their own nextPageToken to
  // continue paging through the results.
  optional string next_page_token = 79797525;

  // [Output Only] Server-defined URL for this resource.
  optional string self_link = 456214797;

  // [Output Only] Informational warning message.
  optional Warning warning = 50704284;

}

//
message CompositeHealthChecksGetHealthResponseHealthSourceHealth {
  // Health state of the associated HealthSource resource.
  enum HealthState {
    // A value indicating that the enum field is not set.
    UNDEFINED_HEALTH_STATE = 0;

    HEALTHY = 439801213;

    UNHEALTHY = 462118084;

    UNKNOWN = 433141802;

  }

  // Health state of the associated HealthSource resource.
  // Check the HealthState enum for the list of possible values.
  optional string health_state = 324007150;

  // Fully qualified URL of the associated HealthSource resource.
  optional string source = 177235995;

}

//
message CompositeHealthChecksScopedList {
  // A list of CompositeHealthChecks contained in this scope.
  repeated CompositeHealthCheck composite_health_checks = 517891062;

  // Informational warning which replaces the list of composite health checks
  // when the list is empty.
  optional Warning warning = 50704284;

}

// A request message for RegionCompositeHealthChecks.AggregatedList. See the method description for details.
message AggregatedListRegionCompositeHealthChecksRequest {
  optional string filter = 336120696;
  optional bool include_all_scopes = 391327988;
  optional uint32 max_results = 54715419;
  optional string order_by = 160562920;
  optional string page_token = 19994697;
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  optional bool return_partial_success = 517198390;
  optional int64 service_project_number = 316757497;
}

// A request message for RegionCompositeHealthChecks.Delete. See the method description for details.
message DeleteRegionCompositeHealthCheckRequest {
  // Name of the CompositeHealthCheck resource to delete.
  string composite_health_check = 466984989 [(google.api.field_behavior) = REQUIRED];

  // Project ID for this request.
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];

  // Name of the region scoping this request.
  string region = 138946292 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "region"
  ];

  // An optional request ID to identify requests. Specify a unique request ID so
  // that if you must retry your request, the server will know to ignore the
  // request if it has already been completed.
  optional string request_id = 37109963;
}

// A request message for RegionCompositeHealthChecks.GetHealth. See the method description for details.
message GetHealthRegionCompositeHealthCheckRequest {
  // Name of the CompositeHealthCheck resource to get health for.
  string composite_health_check = 466984989 [(google.api.field_behavior) = REQUIRED];

  // Name of the project scoping this request.
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];

  // Name of the region scoping this request.
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];
}

// A request message for RegionCompositeHealthChecks.Get. See the method description for details.
message GetRegionCompositeHealthCheckRequest {
  // Name of the CompositeHealthCheck resource to return.
  string composite_health_check = 466984989 [(google.api.field_behavior) = REQUIRED];

  // Project ID for this request.
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];

  // Name of the region scoping this request.
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];
}

// A request message for RegionCompositeHealthChecks.Insert. See the method description for details.
message InsertRegionCompositeHealthCheckRequest {
  // The body resource for this request
  CompositeHealthCheck composite_health_check_resource = 132195824 [(google.api.field_behavior) = REQUIRED];

  // Project ID for this request.
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];

  // Name of the region scoping this request.
  string region = 138946292 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "region"
  ];

  // An optional request ID to identify requests. Specify a unique request ID so
  // that if you must retry your request, the server will know to ignore the
  // request if it has already been completed.
  optional string request_id = 37109963;
}

// A request message for RegionCompositeHealthChecks.List. See the method description for details.
message ListRegionCompositeHealthChecksRequest {
  optional string filter = 336120696;
  optional uint32 max_results = 54715419;
  optional string order_by = 160562920;
  optional string page_token = 19994697;
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];
  optional bool return_partial_success = 517198390;
}

// A request message for RegionCompositeHealthChecks.Patch. See the method description for details.
message PatchRegionCompositeHealthCheckRequest {
  // Name of the CompositeHealthCheck to update. The name
  // must be 1-63 characters long, and comply with RFC1035.
  string composite_health_check = 466984989 [(google.api.field_behavior) = REQUIRED];

  // The body resource for this request
  CompositeHealthCheck composite_health_check_resource = 132195824 [(google.api.field_behavior) = REQUIRED];

  // Project ID for this request.
  string project = 227560217 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "project"
  ];

  // Name of the region scoping this request.
  string region = 138946292 [
    (google.api.field_behavior) = REQUIRED,
    (google.cloud.operation_request_field) = "region"
  ];

  // An optional request ID to identify requests. Specify a unique request ID so
  // that if you must retry your request, the server will know to ignore the
  // request if it has already been completed.
  optional string request_id = 37109963;
}

// A request message for RegionCompositeHealthChecks.TestIamPermissions. See the method description for details.
message TestIamPermissionsRegionCompositeHealthCheckRequest {
  // Project ID for this request.
  string project = 227560217 [(google.api.field_behavior) = REQUIRED];

  // The name of the region for this request.
  string region = 138946292 [(google.api.field_behavior) = REQUIRED];

  // Name or id of the resource for this request.
  string resource = 195806222 [(google.api.field_behavior) = REQUIRED];

  // The body resource for this request
  TestPermissionsRequest test_permissions_request_resource = 439214758 [(google.api.field_behavior) = REQUIRED];
}

// The RegionCompositeHealthChecks API.
service RegionCompositeHealthChecks {
  option (google.api.default_host) =
    "compute.googleapis.com";

  option (google.api.oauth_scopes) =
    "https://www.googleapis.com/auth/compute,"
    "https://www.googleapis.com/auth/cloud-platform";

  // Retrieves the list of all CompositeHealthCheck resources (all
  // regional) available to the specified project.
  rpc AggregatedList(AggregatedListRegionCompositeHealthChecksRequest) returns (CompositeHealthCheckAggregatedList) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/aggregated/compositeHealthChecks"
    };
    option (google.api.method_signature) = "project";
  }

  // Deletes the specified CompositeHealthCheck in the given region
  rpc Delete(DeleteRegionCompositeHealthCheckRequest) returns (Operation) {
    option (google.api.http) = {
      delete: "/compute/v1/projects/{project}/regions/{region}/compositeHealthChecks/{composite_health_check}"
    };
    option (google.api.method_signature) = "project,region,composite_health_check";
    option (google.cloud.operation_service) = "RegionOperations";
  }

  // Returns the specified CompositeHealthCheck resource in the given region.
  rpc Get(GetRegionCompositeHealthCheckRequest) returns (CompositeHealthCheck) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/regions/{region}/compositeHealthChecks/{composite_health_check}"
    };
    option (google.api.method_signature) = "project,region,composite_health_check";
  }

  // Gets the most recent health check results for this
  // regional CompositeHealthCheck.
  rpc GetHealth(GetHealthRegionCompositeHealthCheckRequest) returns (CompositeHealthCheckHealth) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/regions/{region}/compositeHealthChecks/{composite_health_check}/getHealth"
    };
    option (google.api.method_signature) = "project,region,composite_health_check";
  }

  // Create a CompositeHealthCheck in the specified project in the given region
  // using the parameters that are included in the request.
  rpc Insert(InsertRegionCompositeHealthCheckRequest) returns (Operation) {
    option (google.api.http) = {
      body: "composite_health_check_resource"
      post: "/compute/v1/projects/{project}/regions/{region}/compositeHealthChecks"
    };
    option (google.api.method_signature) = "project,region,composite_health_check_resource";
    option (google.cloud.operation_service) = "RegionOperations";
  }

  // Lists the CompositeHealthChecks for a project in the given region.
  rpc List(ListRegionCompositeHealthChecksRequest) returns (CompositeHealthCheckList) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/regions/{region}/compositeHealthChecks"
    };
    option (google.api.method_signature) = "project,region";
  }

  // Updates the specified regional CompositeHealthCheck resource
  // with the data included in the request.  This method supportsPATCH
  // semantics and uses theJSON merge
  // patch format and processing rules.
  rpc Patch(PatchRegionCompositeHealthCheckRequest) returns (Operation) {
    option (google.api.http) = {
      body: "composite_health_check_resource"
      patch: "/compute/v1/projects/{project}/regions/{region}/compositeHealthChecks/{composite_health_check}"
    };
    option (google.api.method_signature) = "project,region,composite_health_check,composite_health_check_resource";
    option (google.cloud.operation_service) = "RegionOperations";
  }

  // Returns permissions that a caller has on the specified resource.
  rpc TestIamPermissions(TestIamPermissionsRegionCompositeHealthCheckRequest) returns (TestPermissionsResponse) {
    option (google.api.http) = {
      body: "test_permissions_request_resource"
      post: "/compute/v1/projects/{project}/regions/{region}/compositeHealthChecks/{resource}/testIamPermissions"
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
