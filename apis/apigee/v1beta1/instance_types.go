// Copyright 2024 Google LLC
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

var ApigeeInstanceGVK = GroupVersion.WithKind("ApigeeInstance")

// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AccessLoggingConfig
type AccessLoggingConfig struct {
	// Optional. Boolean flag that specifies whether the customer access log feature is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AccessLoggingConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. Ship the access log entries that match the status_code defined in the filter. The status_code is the only expected/supported filter field. (Ex: status_code) The filter will parse it to the Common Expression Language semantics for expression evaluation to build the filter condition. (Ex: "filter": status_code >= 200 && status_code < 300 )
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AccessLoggingConfig.filter
	Filter *string `json:"filter,omitempty"`
}

// ApigeeInstanceSpec defines the desired state of ApigeeInstance
// +kcc:spec:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance
type ApigeeInstanceSpec struct {
	// Reference to parent Apigee Organization.
	// +required
	OrganizationRef *ApigeeOrganizationRef `json:"organizationRef"`

	// The ApigeeInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Access logging configuration enables the access logging feature at the instance. Apigee customers can enable access logging to ship the access logs to their own project's cloud logging.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.access_logging_config
	AccessLoggingConfig *AccessLoggingConfig `json:"accessLoggingConfig,omitempty"`

	// Optional. Customer accept list represents the list of projects (id/number) on customer side that can privately connect to the service attachment. It is an optional field which the customers can provide during the instance creation. By default, the customer project associated with the Apigee organization will be included to the list.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.consumer_accept_list
	ConsumerAcceptList []string `json:"consumerAcceptList,omitempty"`

	// Optional. Description of the instance.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.description
	Description *string `json:"description,omitempty"`

	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. If not specified, a Google-Managed encryption key will be used.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.disk_encryption_key_name
	DiskEncryptionKMSCryptoKeyRef *refs.KMSCryptoKeyRef `json:"diskEncryptionKMSCryptoKeyRef,omitempty"`

	// Optional. Display name for the instance.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Comma-separated list of CIDR blocks of length 22 and/or 28 used to create the Apigee instance. Providing CIDR ranges is optional. You can provide just /22 or /28 or both (or neither). Ranges you provide should be freely available as part of a larger named range you have allocated to the Service Networking peering. If this parameter is not provided, Apigee automatically requests an available /22 and /28 CIDR block from Service Networking. Use the /22 CIDR block for configuring your firewall needs to allow traffic from Apigee. Input formats: `a.b.c.d/22` or `e.f.g.h/28` or `a.b.c.d/22,e.f.g.h/28`
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.ip_range
	IPRange *string `json:"ipRange,omitempty"`

	// Required. Compute Engine location where the instance resides.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.location
	// +required
	Location *string `json:"location,omitempty"`

	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.peering_cidr_range
	PeeringCIDRRange *string `json:"peeringCIDRRange,omitempty"`
}

// ApigeeInstanceStatus defines the config connector machine state of ApigeeInstance
type ApigeeInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApigeeInstanceObservedState `json:"observedState,omitempty"`
}

// ApigeeInstanceSpec defines the desired state of ApigeeInstance
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance
// ApigeeInstanceObservedState is the state of the ApigeeInstance resource as most recently observed in GCP.
type ApigeeInstanceObservedState struct {
	// Output only. Time the instance was created in milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.created_at
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// Output only. Internal hostname or IP address of the Apigee endpoint used by clients to connect to the service.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.host
	Host *string `json:"host,omitempty"`

	// Output only. Time the instance was last modified in milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.last_modified_at
	LastModifiedAt *int64 `json:"lastModifiedAt,omitempty"`

	// Output only. Port number of the exposed Apigee endpoint.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.port
	Port *string `json:"port,omitempty"`

	// Output only. Version of the runtime system running in the instance. The runtime system is the set of components that serve the API Proxy traffic in your Environments.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.runtime_version
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`

	// Output only. Resource name of the service attachment created for the instance in the format: `projects/{{project-id}}/regions/{{region-id}}/serviceAttachments/{{service-attachment-id}}` Apigee customers can privately forward traffic to this service attachment using the PSC endpoints.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.service_attachment
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`

	// Output only. State of the instance. Values other than `ACTIVE` means the resource is not ready to use.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Instance.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeinstance;gcpapigeeinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion

// ApigeeInstance is the Schema for the ApigeeInstance API
// +k8s:openapi-gen=true
type ApigeeInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeInstanceSpec   `json:"spec,omitempty"`
	Status ApigeeInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeInstanceList contains a list of ApigeeInstance
type ApigeeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeInstance{}, &ApigeeInstanceList{})
}
