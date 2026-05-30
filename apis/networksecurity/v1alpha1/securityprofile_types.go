// Copyright 2026 Google LLC
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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkSecuritySecurityProfileGVK = GroupVersion.WithKind("NetworkSecuritySecurityProfile")

// +kcc:proto=google.cloud.networksecurity.v1.CustomInterceptProfile
type CustomInterceptProfile struct {
	// Required. The target InterceptEndpointGroup.
	//  When a firewall rule with this security profile attached matches a packet,
	//  the packet will be intercepted to the location-local target in this group.
	// +kcc:proto:field=google.cloud.networksecurity.v1.CustomInterceptProfile.intercept_endpoint_group
	InterceptEndpointGroupRef *refsv1beta1.NetworkSecurityInterceptEndpointGroupRef `json:"interceptEndpointGroupRef,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.CustomMirroringProfile
type CustomMirroringProfile struct {
	// Required. Immutable. The target MirroringEndpointGroup.
	//  When a mirroring rule with this security profile attached matches a packet,
	//  a replica will be mirrored to the location-local target in this group.
	// +kcc:proto:field=google.cloud.networksecurity.v1.CustomMirroringProfile.mirroring_endpoint_group
	MirroringEndpointGroupRef *refsv1beta1.NetworkSecurityMirroringEndpointGroupRef `json:"mirroringEndpointGroupRef,omitempty"`
}

// NetworkSecuritySecurityProfileSpec defines the desired state of NetworkSecuritySecurityProfile
// +kcc:spec:proto=google.cloud.networksecurity.v1.SecurityProfile
type NetworkSecuritySecurityProfileSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Optional
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// The organization that this resource belongs to.
	// +kubebuilder:validation:Optional
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef,omitempty"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The NetworkSecuritySecurityProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The single ProfileType that the SecurityProfile resource configures.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=THREAT_PREVENTION;CUSTOM_MIRRORING;CUSTOM_INTERCEPT;URL_FILTERING
	// +kcc:proto:field=google.cloud.networksecurity.v1.SecurityProfile.type
	Type *string `json:"type,omitempty"`

	// Optional. An optional description of the profile. Max length 512 characters.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.SecurityProfile.description
	Description *string `json:"description,omitempty"`

	// Optional. Labels as key value pairs
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.SecurityProfile.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The threat prevention configuration for the SecurityProfile.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.SecurityProfile.threat_prevention_profile
	ThreatPreventionProfile *ThreatPreventionProfile `json:"threatPreventionProfile,omitempty"`

	// The custom Packet Mirroring v2 configuration for the SecurityProfile.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.SecurityProfile.custom_mirroring_profile
	CustomMirroringProfile *CustomMirroringProfile `json:"customMirroringProfile,omitempty"`

	// The custom TPPI configuration for the SecurityProfile.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.SecurityProfile.custom_intercept_profile
	CustomInterceptProfile *CustomInterceptProfile `json:"customInterceptProfile,omitempty"`

	// The URL filtering configuration for the SecurityProfile.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.SecurityProfile.url_filtering_profile
	URLFilteringProfile *URLFilteringProfile `json:"urlFilteringProfile,omitempty"`
}

// NetworkSecuritySecurityProfileStatus defines the config connector machine state of NetworkSecuritySecurityProfile
type NetworkSecuritySecurityProfileStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecuritySecurityProfile resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecuritySecurityProfileObservedState `json:"observedState,omitempty"`
}

// NetworkSecuritySecurityProfileObservedState is the state of the NetworkSecuritySecurityProfile resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.SecurityProfile
type NetworkSecuritySecurityProfileObservedState struct {
	// Output only. Resource creation timestamp.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SecurityProfile.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last resource update timestamp.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SecurityProfile.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. This checksum is computed by the server based on the value of other fields.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SecurityProfile.etag
	Etag *string `json:"etag,omitempty"`

	// The threat prevention configuration for the SecurityProfile.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SecurityProfile.threat_prevention_profile
	ThreatPreventionProfile *ThreatPreventionProfileObservedState `json:"threatPreventionProfile,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritysecurityprofile;gcpnetworksecuritysecurityprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecuritySecurityProfile is the Schema for the NetworkSecuritySecurityProfile API
// +k8s:openapi-gen=true
type NetworkSecuritySecurityProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecuritySecurityProfileSpec   `json:"spec,omitempty"`
	Status NetworkSecuritySecurityProfileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecuritySecurityProfileList contains a list of NetworkSecuritySecurityProfile
type NetworkSecuritySecurityProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecuritySecurityProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecuritySecurityProfile{}, &NetworkSecuritySecurityProfileList{})
}
