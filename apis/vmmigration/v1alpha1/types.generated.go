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

package v1alpha1


// +kcc:proto=google.cloud.vmmigration.v1.ApplianceVersion
type ApplianceVersion struct {
	// The appliance version.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ApplianceVersion.version
	Version *string `json:"version,omitempty"`

	// A link for downloading the version.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ApplianceVersion.uri
	URI *string `json:"uri,omitempty"`

	// Determine whether it's critical to upgrade the appliance to this version.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ApplianceVersion.critical
	Critical *bool `json:"critical,omitempty"`

	// Link to a page that contains the version release notes.
	// +kcc:proto:field=google.cloud.vmmigration.v1.ApplianceVersion.release_notes_uri
	ReleaseNotesURI *string `json:"releaseNotesURI,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.AvailableUpdates
type AvailableUpdates struct {
	// The newest deployable version of the appliance.
	//  The current appliance can't be updated into this version, and the owner
	//  must manually deploy this OVA to a new appliance.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AvailableUpdates.new_deployable_appliance
	NewDeployableAppliance *ApplianceVersion `json:"newDeployableAppliance,omitempty"`

	// The latest version for in place update.
	//  The current appliance can be updated to this version using the API or m4c
	//  CLI.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AvailableUpdates.in_place_update
	InPlaceUpdate *ApplianceVersion `json:"inPlaceUpdate,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.DatacenterConnector
type DatacenterConnector struct {

	// Immutable. A unique key for this connector. This key is internal to the OVA
	//  connector and is supplied with its creation during the registration process
	//  and can not be modified.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.registration_id
	RegistrationID *string `json:"registrationID,omitempty"`

	// The service account to use in the connector when communicating with the
	//  cloud.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The version running in the DatacenterConnector. This is supplied by the OVA
	//  connector during the registration process and can not be modified.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.UpgradeStatus
type UpgradeStatus struct {
	// The version to upgrade to.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UpgradeStatus.version
	Version *string `json:"version,omitempty"`

	// The state of the upgradeAppliance operation.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UpgradeStatus.state
	State *string `json:"state,omitempty"`

	// Provides details on the state of the upgrade operation in case of an error.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UpgradeStatus.error
	Error *Status `json:"error,omitempty"`

	// The time the operation was started.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UpgradeStatus.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The version from which we upgraded.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UpgradeStatus.previous_version
	PreviousVersion *string `json:"previousVersion,omitempty"`
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
	//  type.googleapis.com.
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

// +kcc:proto=google.cloud.vmmigration.v1.DatacenterConnector
type DatacenterConnectorObservedState struct {
	// Output only. The time the connector was created (as an API call, not when
	//  it was actually installed).
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last time the connector was updated with an API call.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The connector's name.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.name
	Name *string `json:"name,omitempty"`

	// Output only. The communication channel between the datacenter connector and
	//  Google Cloud.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Output only. State of the DatacenterConnector, as determined by the health
	//  checks.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.state
	State *string `json:"state,omitempty"`

	// Output only. The time the state was last set.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.state_time
	StateTime *string `json:"stateTime,omitempty"`

	// Output only. Provides details on the state of the Datacenter Connector in
	//  case of an error.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.error
	Error *Status `json:"error,omitempty"`

	// Output only. Appliance OVA version.
	//  This is the OVA which is manually installed by the user and contains the
	//  infrastructure for the automatically updatable components on the appliance.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.appliance_infrastructure_version
	ApplianceInfrastructureVersion *string `json:"applianceInfrastructureVersion,omitempty"`

	// Output only. Appliance last installed update bundle version.
	//  This is the version of the automatically updatable components on the
	//  appliance.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.appliance_software_version
	ApplianceSoftwareVersion *string `json:"applianceSoftwareVersion,omitempty"`

	// Output only. The available versions for updating this appliance.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.available_versions
	AvailableVersions *AvailableUpdates `json:"availableVersions,omitempty"`

	// Output only. The status of the current / last upgradeAppliance operation.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.upgrade_status
	UpgradeStatus *UpgradeStatus `json:"upgradeStatus,omitempty"`
}
