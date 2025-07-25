// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +generated:types
// krm.group: backupdr.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.backupdr.v1
// resource: BackupDRManagementServer:ManagementServer
// resource: BackupDRBackupVault:BackupVault
// resource: BackupDRBackupPlanAssociation:BackupPlanAssociation

package v1alpha1


// +kcc:proto=google.cloud.backupdr.v1.RuleConfigInfo
type RuleConfigInfo struct {
}

// +kcc:proto=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI
type WorkforceIdentityBasedManagementURI struct {
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com. As of May 2023, there are no widely used type server
	//  implementations and no plans to implement one.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.backupdr.v1.RuleConfigInfo
type RuleConfigInfoObservedState struct {
	// Output only. Output Only.
	//
	//  Backup Rule id fetched from backup plan.
	// +kcc:proto:field=google.cloud.backupdr.v1.RuleConfigInfo.rule_id
	RuleID *string `json:"ruleID,omitempty"`

	// Output only. The last backup state for rule.
	// +kcc:proto:field=google.cloud.backupdr.v1.RuleConfigInfo.last_backup_state
	LastBackupState *string `json:"lastBackupState,omitempty"`

	// Output only. Output Only.
	//
	//  google.rpc.Status object to store the last backup error.
	// +kcc:proto:field=google.cloud.backupdr.v1.RuleConfigInfo.last_backup_error
	LastBackupError *Status `json:"lastBackupError,omitempty"`

	// Output only. The point in time when the last successful backup was captured
	//  from the source.
	// +kcc:proto:field=google.cloud.backupdr.v1.RuleConfigInfo.last_successful_backup_consistency_time
	LastSuccessfulBackupConsistencyTime *string `json:"lastSuccessfulBackupConsistencyTime,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI
type WorkforceIdentityBasedManagementURIObservedState struct {
	// Output only. First party Management URI for Google Identities.
	// +kcc:proto:field=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI.first_party_management_uri
	FirstPartyManagementURI *string `json:"firstPartyManagementURI,omitempty"`

	// Output only. Third party Management URI for External Identity Providers.
	// +kcc:proto:field=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI.third_party_management_uri
	ThirdPartyManagementURI *string `json:"thirdPartyManagementURI,omitempty"`
}
