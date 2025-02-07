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


// +kcc:proto=google.cloud.assuredworkloads.v1beta1.Workload
type Workload struct {
	// Optional. The resource name of the workload.
	//  Format:
	//  organizations/{organization}/locations/{location}/workloads/{workload}
	//
	//  Read-only.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.name
	Name *string `json:"name,omitempty"`

	// Required. The user-assigned display name of the Workload.
	//  When present it must be between 4 to 30 characters.
	//  Allowed characters are: lowercase and uppercase letters, numbers,
	//  hyphen, and spaces.
	//
	//  Example: My Workload
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. Compliance Regime associated with this workload.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.compliance_regime
	ComplianceRegime *string `json:"complianceRegime,omitempty"`

	// Input only. Immutable. Settings specific to resources needed for IL4.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.il4_settings
	Il4Settings *Workload_IL4Settings `json:"il4Settings,omitempty"`

	// Input only. Immutable. Settings specific to resources needed for CJIS.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.cjis_settings
	CjisSettings *Workload_CJISSettings `json:"cjisSettings,omitempty"`

	// Input only. Immutable. Settings specific to resources needed for FedRAMP High.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.fedramp_high_settings
	FedrampHighSettings *Workload_FedrampHighSettings `json:"fedrampHighSettings,omitempty"`

	// Input only. Immutable. Settings specific to resources needed for FedRAMP Moderate.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.fedramp_moderate_settings
	FedrampModerateSettings *Workload_FedrampModerateSettings `json:"fedrampModerateSettings,omitempty"`

	// Optional. ETag of the workload, it is calculated on the basis
	//  of the Workload contents. It will be used in Update & Delete operations.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Labels applied to the workload.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Input only. The parent resource for the resources managed by this Assured Workload. May
	//  be either empty or a folder resource which is a child of the
	//  Workload parent. If not specified all resources are created under the
	//  parent organization.
	//  Format:
	//  folders/{folder_id}
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.provisioned_resources_parent
	ProvisionedResourcesParent *string `json:"provisionedResourcesParent,omitempty"`

	// Input only. Settings used to create a CMEK crypto key. When set, a project with a KMS
	//  CMEK key is provisioned.
	//  This field is deprecated as of Feb 28, 2022.
	//  In order to create a Keyring, callers should specify,
	//  ENCRYPTION_KEYS_PROJECT or KEYRING in ResourceSettings.resource_type field.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.kms_settings
	KMSSettings *Workload_KMSSettings `json:"kmsSettings,omitempty"`

	// Input only. Resource properties that are used to customize workload resources.
	//  These properties (such as custom project id) will be used to create
	//  workload resources if possible. This field is optional.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.resource_settings
	ResourceSettings []Workload_ResourceSettings `json:"resourceSettings,omitempty"`

	// Optional. Indicates the sovereignty status of the given workload.
	//  Currently meant to be used by Europe/Canada customers.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.enable_sovereign_controls
	EnableSovereignControls *bool `json:"enableSovereignControls,omitempty"`
}

// +kcc:proto=google.cloud.assuredworkloads.v1beta1.Workload.CJISSettings
type Workload_CJISSettings struct {
	// Input only. Immutable. Settings used to create a CMEK crypto key.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.CJISSettings.kms_settings
	KMSSettings *Workload_KMSSettings `json:"kmsSettings,omitempty"`
}

// +kcc:proto=google.cloud.assuredworkloads.v1beta1.Workload.FedrampHighSettings
type Workload_FedrampHighSettings struct {
	// Input only. Immutable. Settings used to create a CMEK crypto key.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.FedrampHighSettings.kms_settings
	KMSSettings *Workload_KMSSettings `json:"kmsSettings,omitempty"`
}

// +kcc:proto=google.cloud.assuredworkloads.v1beta1.Workload.FedrampModerateSettings
type Workload_FedrampModerateSettings struct {
	// Input only. Immutable. Settings used to create a CMEK crypto key.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.FedrampModerateSettings.kms_settings
	KMSSettings *Workload_KMSSettings `json:"kmsSettings,omitempty"`
}

// +kcc:proto=google.cloud.assuredworkloads.v1beta1.Workload.IL4Settings
type Workload_IL4Settings struct {
	// Input only. Immutable. Settings used to create a CMEK crypto key.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.IL4Settings.kms_settings
	KMSSettings *Workload_KMSSettings `json:"kmsSettings,omitempty"`
}

// +kcc:proto=google.cloud.assuredworkloads.v1beta1.Workload.KMSSettings
type Workload_KMSSettings struct {
	// Required. Input only. Immutable. The time at which the Key Management Service will automatically create a
	//  new version of the crypto key and mark it as the primary.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.KMSSettings.next_rotation_time
	NextRotationTime *string `json:"nextRotationTime,omitempty"`

	// Required. Input only. Immutable. [next_rotation_time] will be advanced by this period when the Key
	//  Management Service automatically rotates a key. Must be at least 24 hours
	//  and at most 876,000 hours.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.KMSSettings.rotation_period
	RotationPeriod *string `json:"rotationPeriod,omitempty"`
}

// +kcc:proto=google.cloud.assuredworkloads.v1beta1.Workload.ResourceInfo
type Workload_ResourceInfo struct {
	// Resource identifier.
	//  For a project this represents project_number.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.ResourceInfo.resource_id
	ResourceID *int64 `json:"resourceID,omitempty"`

	// Indicates the type of resource.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.ResourceInfo.resource_type
	ResourceType *string `json:"resourceType,omitempty"`
}

// +kcc:proto=google.cloud.assuredworkloads.v1beta1.Workload.ResourceSettings
type Workload_ResourceSettings struct {
	// Resource identifier.
	//  For a project this represents project_id. If the project is already
	//  taken, the workload creation will fail.
	//  For KeyRing, this represents the keyring_id.
	//  For a folder, don't set this value as folder_id is assigned by Google.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.ResourceSettings.resource_id
	ResourceID *string `json:"resourceID,omitempty"`

	// Indicates the type of resource. This field should be specified to
	//  correspond the id to the right project type (CONSUMER_PROJECT or
	//  ENCRYPTION_KEYS_PROJECT)
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.ResourceSettings.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// User-assigned resource display name.
	//  If not empty it will be used to create a resource with the specified
	//  name.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.ResourceSettings.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.assuredworkloads.v1beta1.Workload.SaaEnrollmentResponse
type Workload_SaaEnrollmentResponse struct {
	// Indicates SAA enrollment status of a given workload.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.SaaEnrollmentResponse.setup_status
	SetupStatus *string `json:"setupStatus,omitempty"`

	// Indicates SAA enrollment setup error if any.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.SaaEnrollmentResponse.setup_errors
	SetupErrors []string `json:"setupErrors,omitempty"`
}

// +kcc:proto=google.cloud.assuredworkloads.v1beta1.Workload
type WorkloadObservedState struct {
	// Output only. The resources associated with this workload.
	//  These resources will be created when creating the workload.
	//  If any of the projects already exist, the workload creation will fail.
	//  Always read only.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.resources
	Resources []Workload_ResourceInfo `json:"resources,omitempty"`

	// Output only. Immutable. The Workload creation timestamp.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The billing account used for the resources which are
	//  direct children of workload. This billing account is initially associated
	//  with the resources created as part of Workload creation.
	//  After the initial creation of these resources, the customer can change
	//  the assigned billing account.
	//  The resource name has the form
	//  `billingAccounts/{billing_account_id}`. For example,
	//  `billingAccounts/012345-567890-ABCDEF`.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.billing_account
	BillingAccount *string `json:"billingAccount,omitempty"`

	// Output only. Represents the KAJ enrollment state of the given workload.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.kaj_enrollment_state
	KajEnrollmentState *string `json:"kajEnrollmentState,omitempty"`

	// Output only. Represents the SAA enrollment response of the given workload.
	//  SAA enrollment response is queried during GetWorkload call.
	//  In failure cases, user friendly error message is shown in SAA details page.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.saa_enrollment_response
	SaaEnrollmentResponse *Workload_SaaEnrollmentResponse `json:"saaEnrollmentResponse,omitempty"`

	// Output only. Urls for services which are compliant for this Assured Workload, but which
	//  are currently disallowed by the ResourceUsageRestriction org policy.
	//  Invoke RestrictAllowedResources endpoint to allow your project developers
	//  to use these services in their environment."
	// +kcc:proto:field=google.cloud.assuredworkloads.v1beta1.Workload.compliant_but_disallowed_services
	CompliantButDisallowedServices []string `json:"compliantButDisallowedServices,omitempty"`
}
