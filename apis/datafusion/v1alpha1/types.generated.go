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


// +kcc:proto=google.cloud.datafusion.v1.Accelerator
type Accelerator struct {
	// The type of an accelator for a CDF instance.
	// +kcc:proto:field=google.cloud.datafusion.v1.Accelerator.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// The state of the accelerator
	// +kcc:proto:field=google.cloud.datafusion.v1.Accelerator.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.datafusion.v1.CryptoKeyConfig
type CryptoKeyConfig struct {
	// The name of the key which is used to encrypt/decrypt customer data. For key
	//  in Cloud KMS, the key should be in the format of
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.cloud.datafusion.v1.CryptoKeyConfig.key_reference
	KeyReference *string `json:"keyReference,omitempty"`
}

// +kcc:proto=google.cloud.datafusion.v1.Instance
type Instance struct {

	// A description of this instance.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.description
	Description *string `json:"description,omitempty"`

	// Required. Instance type.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.type
	Type *string `json:"type,omitempty"`

	// Option to enable Stackdriver Logging.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.enable_stackdriver_logging
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty"`

	// Option to enable Stackdriver Monitoring.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.enable_stackdriver_monitoring
	EnableStackdriverMonitoring *bool `json:"enableStackdriverMonitoring,omitempty"`

	// Specifies whether the Data Fusion instance should be private. If set to
	//  true, all Data Fusion nodes will have private IP addresses and will not be
	//  able to access the public internet.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.private_instance
	PrivateInstance *bool `json:"privateInstance,omitempty"`

	// Network configuration options. These are required when a private Data
	//  Fusion instance is to be created.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.network_config
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`

	// The resource labels for instance to use to annotate any related underlying
	//  resources such as Compute Engine VMs. The character '=' is not allowed to
	//  be used within the labels.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Map of additional options used to configure the behavior of
	//  Data Fusion instance.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.options
	Options map[string]string `json:"options,omitempty"`

	// Name of the zone in which the Data Fusion instance will be created. Only
	//  DEVELOPER instances use this field.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.zone
	Zone *string `json:"zone,omitempty"`

	// Current version of the Data Fusion. Only specifiable in Update.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.version
	Version *string `json:"version,omitempty"`

	// Display name for an instance.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Available versions that the instance can be upgraded to using
	//  UpdateInstanceRequest.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.available_version
	AvailableVersion []Version `json:"availableVersion,omitempty"`

	// List of accelerators enabled for this CDF instance.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.accelerators
	Accelerators []Accelerator `json:"accelerators,omitempty"`

	// User-managed service account to set on Dataproc when Cloud Data Fusion
	//  creates Dataproc to run data processing pipelines.
	//
	//  This allows users to have fine-grained access control on Dataproc's
	//  accesses to cloud resources.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.dataproc_service_account
	DataprocServiceAccount *string `json:"dataprocServiceAccount,omitempty"`

	// Option to enable granular role-based access control.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.enable_rbac
	EnableRbac *bool `json:"enableRbac,omitempty"`

	// The crypto key configuration. This field is used by the Customer-Managed
	//  Encryption Keys (CMEK) feature.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.crypto_key_config
	CryptoKeyConfig *CryptoKeyConfig `json:"cryptoKeyConfig,omitempty"`
}

// +kcc:proto=google.cloud.datafusion.v1.NetworkConfig
type NetworkConfig struct {
	// Name of the network in the customer project with which the Tenant Project
	//  will be peered for executing pipelines. In case of shared VPC where the
	//  network resides in another host project the network should specified in
	//  the form of projects/{host-project-id}/global/networks/{network}
	// +kcc:proto:field=google.cloud.datafusion.v1.NetworkConfig.network
	Network *string `json:"network,omitempty"`

	// The IP range in CIDR notation to use for the managed Data Fusion instance
	//  nodes. This range must not overlap with any other ranges used in the
	//  customer network.
	// +kcc:proto:field=google.cloud.datafusion.v1.NetworkConfig.ip_allocation
	IPAllocation *string `json:"ipAllocation,omitempty"`
}

// +kcc:proto=google.cloud.datafusion.v1.Version
type Version struct {
	// The version number of the Data Fusion instance, such as '6.0.1.0'.
	// +kcc:proto:field=google.cloud.datafusion.v1.Version.version_number
	VersionNumber *string `json:"versionNumber,omitempty"`

	// Whether this is currently the default version for Cloud Data Fusion
	// +kcc:proto:field=google.cloud.datafusion.v1.Version.default_version
	DefaultVersion *bool `json:"defaultVersion,omitempty"`

	// Represents a list of available feature names for a given version.
	// +kcc:proto:field=google.cloud.datafusion.v1.Version.available_features
	AvailableFeatures []string `json:"availableFeatures,omitempty"`

	// Type represents the release availability of the version
	// +kcc:proto:field=google.cloud.datafusion.v1.Version.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.datafusion.v1.Instance
type InstanceObservedState struct {
	// Output only. The name of this instance is in the form of
	//  projects/{project}/locations/{location}/instances/{instance}.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.name
	Name *string `json:"name,omitempty"`

	// Output only. The time the instance was created.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the instance was last updated.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The current state of this Data Fusion instance.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the current state of this Data
	//  Fusion instance if available.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. Endpoint on which the Data Fusion UI is accessible.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.service_endpoint
	ServiceEndpoint *string `json:"serviceEndpoint,omitempty"`

	// Output only. Deprecated. Use tenant_project_id instead to extract the tenant project ID.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. Endpoint on which the REST APIs is accessible.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.api_endpoint
	ApiEndpoint *string `json:"apiEndpoint,omitempty"`

	// Output only. Cloud Storage bucket generated by Data Fusion in the customer project.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.gcs_bucket
	GcsBucket *string `json:"gcsBucket,omitempty"`

	// Output only. P4 service account for the customer project.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.p4_service_account
	P4ServiceAccount *string `json:"p4ServiceAccount,omitempty"`

	// Output only. The name of the tenant project.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.tenant_project_id
	TenantProjectID *string `json:"tenantProjectID,omitempty"`

	// Output only. If the instance state is DISABLED, the reason for disabling the instance.
	// +kcc:proto:field=google.cloud.datafusion.v1.Instance.disabled_reason
	DisabledReason []string `json:"disabledReason,omitempty"`
}
