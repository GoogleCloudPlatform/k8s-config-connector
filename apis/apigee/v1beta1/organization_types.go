// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApigeeOrganizationGVK = GroupVersion.WithKind("ApigeeOrganization")

// ApigeeOrganizationSpec defines the desired state of ApigeeOrganization
// +kcc:spec:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization
type ApigeeOrganizationSpec struct {
	// The ApigeeOrganization name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Addon configurations of the Apigee organization.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.addons_config
	AddonsConfig *GoogleCloudApigeeV1AddonsConfig `json:"addonsConfig,omitempty"`

	// Required. DEPRECATED: This field will eventually be deprecated and replaced with a differently-named field. Primary Google Cloud region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.analytics_region
	// +required
	AnalyticsRegion *string `json:"analyticsRegion,omitempty"`

	/* NOTYET: Add this once direct controller is implemented
	// Cloud KMS key name used for encrypting API consumer data. If not specified or [BillingType](#BillingType) is `EVALUATION`, a Google-Managed encryption key will be used. Format: `projects/{{projectID}}/locations/{{locationID}}/keyRings/{{keyRingID}}/cryptoKeys/{{cryptoKeyID}}`
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.api_consumer_data_encryption_key_name
	ApiConsumerDataEncryptionKeyName *string `json:"apiConsumerDataEncryptionKeyName,omitempty"`
	c

	/* NOTYET: Add this once direct controller is implemented
	// This field is needed only for customers using non-default data residency regions. Apigee stores some control plane data only in single region. This field determines which single region Apigee should use. For example: "us-west1" when control plane is in US or "europe-west2" when control plane is in EU.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.api_consumer_data_location
	ApiConsumerDataLocation *string `json:"apiConsumerDataLocation,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Not used by Apigee.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.attributes
	Attributes []string `json:"attributes,omitempty"`
	*/

	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started). Valid only when [RuntimeType](#RuntimeType) is set to `CLOUD`. The value must be set before the creation of a runtime instance and can be updated only when there are no runtime instances. For example: `default`. Apigee also supports shared VPC (that is, the host network project is not the same as the one that is peering with Apigee). See [Shared VPC overview](https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.authorized_network
	AuthorizedNetworkRef *refs.ComputeNetworkRef `json:"authorizedNetworkRef,omitempty"`

	/* NOTYET: Add this once direct controller is implemented
	// Cloud KMS key name used for encrypting control plane data that is stored in a multi region. Only used for the data residency region "US" or "EU". If not specified or [BillingType](#BillingType) is `EVALUATION`, a Google-Managed encryption key will be used. Format: `projects/{{projectID}}/locations/{{locationID}}/keyRings/{{keyRingID}}/cryptoKeys/{{cryptoKeyID}}`
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.control_plane_encryption_key_name
	ControlPlaneEncryptionKeyName *string `json:"controlPlaneEncryptionKeyName,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Not used by Apigee.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.customer_name
	CustomerName *string `json:"customerName,omitempty"`
	*/

	// Description of the Apigee organization.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.description
	Description *string `json:"description,omitempty"`

	/* NOTYET: Add this once direct controller is implemented
	// Optional. Flag that specifies whether the VPC Peering through Private Google Access should be disabled between the consumer network and Apigee. Valid only when RuntimeType is set to CLOUD. Required if an authorizedNetwork on the consumer project is not provided, in which case the flag should be set to true. The value must be set before the creation of any Apigee runtime instance and can be updated only when there are no runtime instances. **Note:** Apigee will be deprecating the vpc peering model that requires you to provide 'authorizedNetwork', by making the non-peering model as the default way of provisioning Apigee organization in future. So, this will be a temporary flag to enable the transition. Not supported for Apigee hybrid.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.disable_vpc_peering
	DisableVpcPeering *bool `json:"disableVpcPeering,omitempty"`
	*/

	// Display name for the Apigee organization. Unused, but reserved for future use.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Name of the GCP project in which to associate the Apigee organization.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.project_id
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	/* NOTYET: Add this once direct controller is implemented
	// Configuration for the Portals settings.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.portal_disabled
	PortalDisabled *bool `json:"portalDisabled,omitempty"`
	*/

	// Properties defined in the Apigee organization profile.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.properties
	Properties *map[string]string `json:"properties,omitempty"`

	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. If not specified or [RuntimeType](#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: "projects/foo/locations/us/keyRings/bar/cryptoKeys/baz". **Note:** Not supported for Apigee hybrid.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.runtime_database_encryption_key_name
	RuntimeDatabaseEncryptionKeyRef *refs.KMSCryptoKeyRef `json:"runtimeDatabaseEncryptionKeyRef,omitempty"`

	// Required. Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.runtime_type
	// +required
	RuntimeType *string `json:"runtimeType,omitempty"`

	/* NOTYET: Add this once direct controller is implemented
	// Not used by Apigee.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.type
	Type *string `json:"type,omitempty"`
	*/
}

// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig
type GoogleCloudApigeeV1AddonsConfig struct {
	// Configuration for the Advanced API Ops add-on.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig.advanced_api_ops_config
	AdvancedApiOpsConfig *GoogleCloudApigeeV1AdvancedApiOpsConfig `json:"advancedApiOpsConfig,omitempty"`

	/* NOTYET: Add this once direct controller is implemented
	// Configuration for the Analytics add-on. Only used in organizations.environments.addonsConfig.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig.analytics_config
	AnalyticsConfig *GoogleCloudApigeeV1AnalyticsConfig `json:"analyticsConfig,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Configuration for the API Security add-on.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig.api_security_config
	ApiSecurityConfig *GoogleCloudApigeeV1ApiSecurityConfig `json:"apiSecurityConfig,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Configuration for the Connectors Platform add-on.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig.connectors_platform_config
	ConnectorsPlatformConfig *GoogleCloudApigeeV1ConnectorsPlatformConfig `json:"connectorsPlatformConfig,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Configuration for the Integration add-on.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig.integration_config
	IntegrationConfig *GoogleCloudApigeeV1IntegrationConfig `json:"integrationConfig,omitempty"`
	*/

	// Configuration for the Monetization add-on.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig.monetization_config
	MonetizationConfig *GoogleCloudApigeeV1MonetizationConfig `json:"monetizationConfig,omitempty"`
}

// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AdvancedApiOpsConfig
type GoogleCloudApigeeV1AdvancedApiOpsConfig struct {
	// Flag that specifies whether the Advanced API Ops add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AdvancedApiOpsConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

/* NOTYET: Add this once direct controller is implemented
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AnalyticsConfig
type GoogleCloudApigeeV1AnalyticsConfig struct {
	// Whether the Analytics add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AnalyticsConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Output only. Time at which the Analytics add-on expires in milliseconds since epoch. If unspecified, the add-on will never expire.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AnalyticsConfig.expire_time_millis
	ExpireTimeMillis *int64 `json:"expireTimeMillis,omitempty"`

	// Output only. The state of the Analytics add-on.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AnalyticsConfig.state
	State *string `json:"state,omitempty"`

	// Output only. The latest update time.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AnalyticsConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
*/

/* NOTYET: Add this once direct controller is implemented
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiSecurityConfig
type GoogleCloudApigeeV1ApiSecurityConfig struct {
	// Flag that specifies whether the API security add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiSecurityConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Output only. Time at which the API Security add-on expires in in milliseconds since epoch. If unspecified, the add-on will never expire.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiSecurityConfig.expires_at
	ExpiresAt *int64 `json:"expiresAt,omitempty"`
}
*/

/* NOTYET: Add this once direct controller is implemented
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ConnectorsPlatformConfig
type GoogleCloudApigeeV1ConnectorsPlatformConfig struct {
	// Flag that specifies whether the Connectors Platform add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ConnectorsPlatformConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Output only. Time at which the Connectors Platform add-on expires in milliseconds since epoch. If unspecified, the add-on will never expire.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ConnectorsPlatformConfig.expires_at
	ExpiresAt *int64 `json:"expiresAt,omitempty"`
}
*/

/* NOTYET: Add this once direct controller is implemented
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1IntegrationConfig
type GoogleCloudApigeeV1IntegrationConfig struct {
	// Flag that specifies whether the Integration add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1IntegrationConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}
*/

// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1MonetizationConfig
type GoogleCloudApigeeV1MonetizationConfig struct {
	// Flag that specifies whether the Monetization add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1MonetizationConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// ApigeeOrganizationStatus defines the config connector machine state of ApigeeOrganization
type ApigeeOrganizationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeOrganization resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApigeeOrganizationObservedState `json:"observedState,omitempty"`

	/* NOTYET: Perhaps add this once direct controller is implemented. Or, we may only add it to observedState.
	// Output only. Apigee Project ID associated with the organization. Use this project to allowlist Apigee in the Service Attachment when using private service connect with Apigee.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.apigee_project_id
	ApigeeProjectID *string `json:"apigeeProjectID,omitempty"`
	*/

	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.billing_type
	BillingType *string `json:"billingType,omitempty"`

	// Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when [RuntimeType](#RuntimeType) is `CLOUD`.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.ca_certificate
	CaCertificate []byte `json:"caCertificate,omitempty"`

	// Output only. Time that the Apigee organization was created in milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.created_at
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// Output only. List of environments in the Apigee organization.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.environments
	Environments []string `json:"environments,omitempty"`

	// Output only. Time that the Apigee organization is scheduled for deletion.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.expires_at
	ExpiresAt *int64 `json:"expiresAt,omitempty"`

	// Output only. Time that the Apigee organization was last modified in milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.last_modified_at
	LastModifiedAt *int64 `json:"lastModifiedAt,omitempty"`

	// Output only. Project ID associated with the Apigee organization.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.project_id
	ProjectID *string `json:"projectId,omitempty"`

	// Output only. State of the organization. Values other than ACTIVE means the resource is not ready to use.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.state
	State *string `json:"state,omitempty"`

	/* NOTYET: Perhaps add this once direct controller is implemented. Or, we may only add it to observedState.
	// Output only. Subscription plan that the customer has purchased. Output only.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.subscription_plan
	SubscriptionPlan *string `json:"subscriptionPlan,omitempty"`
	*/

	// Output only. DEPRECATED: This will eventually be replaced by BillingType. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation purposes only) or paid (full subscription has been purchased). See [Apigee pricing](https://cloud.google.com/apigee/pricing/).
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.subscription_type
	SubscriptionType *string `json:"subscriptionType,omitempty"`
}

// ApigeeOrganizationObservedState is the state of the ApigeeOrganization resource as most recently observed in GCP.
// +kcc:observedstate:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization
type ApigeeOrganizationObservedState struct {
	/* NOTYET: Add this once direct controller is implemented
	// Output only. Apigee Project ID associated with the organization. Use this project to allowlist Apigee in the Service Attachment when using private service connect with Apigee.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.apigee_project_id
	ApigeeProjectID *string `json:"apigeeProjectID,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.billing_type
	BillingType *string `json:"billingType,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when [RuntimeType](#RuntimeType) is `CLOUD`.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.ca_certificate
	CaCertificate []byte `json:"caCertificate,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Output only. Time that the Apigee organization was created in milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.created_at
	CreatedAt *int64 `json:"createdAt,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Output only. List of environments in the Apigee organization.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.environments
	Environments []string `json:"environments,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Output only. Time that the Apigee organization is scheduled for deletion.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.expires_at
	ExpiresAt *int64 `json:"expiresAt,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Output only. Time that the Apigee organization was last modified in milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.last_modified_at
	LastModifiedAt *int64 `json:"lastModifiedAt,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Output only. Project ID associated with the Apigee organization.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.project_id
	ProjectID *string `json:"projectID,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Output only. State of the organization. Values other than ACTIVE means the resource is not ready to use.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.state
	State *string `json:"state,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Output only. Subscription plan that the customer has purchased. Output only.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.subscription_plan
	SubscriptionPlan *string `json:"subscriptionPlan,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Output only. DEPRECATED: This will eventually be replaced by BillingType. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation purposes only) or paid (full subscription has been purchased). See [Apigee pricing](https://cloud.google.com/apigee/pricing/).
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Organization.subscription_type
	SubscriptionType *string `json:"subscriptionType,omitempty"`
	*/
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeorganization;gcpapigeeorganizations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion

// ApigeeOrganization is the Schema for the ApigeeOrganization API
// +k8s:openapi-gen=true
type ApigeeOrganization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeOrganizationSpec   `json:"spec,omitempty"`
	Status ApigeeOrganizationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeOrganizationList contains a list of ApigeeOrganization
type ApigeeOrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeOrganization `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeOrganization{}, &ApigeeOrganizationList{})
}
