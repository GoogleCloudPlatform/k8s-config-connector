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

var CloudNumberRegistryRealmGVK = GroupVersion.WithKind("CloudNumberRegistryRealm")

// CloudNumberRegistryRealmSpec defines the desired state of CloudNumberRegistryRealm
// +kcc:spec:proto=google.cloud.numberregistry.v1alpha.Realm
type CloudNumberRegistryRealmSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Required. The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The CloudNumberRegistryRealm name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Traffic type of the Realm.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=TRAFFIC_TYPE_UNSPECIFIED;UNSET;INTERNET;PRIVATE;LINKLOCAL
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.Realm.traffic_type
	TrafficType *string `json:"trafficType,omitempty"`

	// Optional. Management type of the Realm.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=MANAGEMENT_TYPE_UNSPECIFIED;CNR;USER
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.Realm.management_type
	ManagementType *string `json:"managementType,omitempty"`

	// Required. Reference to the RegistryBook that claims the Realm.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.Realm.registry_book
	RegistryBookRef *CloudNumberRegistryRegistryBookRef `json:"registryBookRef,omitempty"`

	// Optional. User-defined labels.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.Realm.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. IP version of the Realm.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=IP_VERSION_UNSPECIFIED;IPV4;IPV6
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.Realm.ip_version
	IPVersion *string `json:"ipVersion,omitempty"`
}

// CloudNumberRegistryRealmStatus defines the config connector machine state of CloudNumberRegistryRealm
type CloudNumberRegistryRealmStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudNumberRegistryRealm resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudNumberRegistryRealmObservedState `json:"observedState,omitempty"`
}

// CloudNumberRegistryRealmObservedState is the state of the CloudNumberRegistryRealm resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.numberregistry.v1alpha.Realm
type CloudNumberRegistryRealmObservedState struct {
	// Output only. Discovery metadata of the Realm.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.Realm.discovery_metadata
	DiscoveryMetadata *DiscoveryMetadataObservedState `json:"discoveryMetadata,omitempty"`

	// Output only. The time at which the Realm was created.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.Realm.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the Realm was last updated.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.Realm.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Aggregated data for the Realm. Populated only when the view is AGGREGATE.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.Realm.aggregated_data
	AggregatedData *Realm_RealmAggregatedDataObservedState `json:"aggregatedData,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudnumberregistryrealm;gcpcloudnumberregistryrealms
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudNumberRegistryRealm is the Schema for the CloudNumberRegistryRealm API
// +k8s:openapi-gen=true
type CloudNumberRegistryRealm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudNumberRegistryRealmSpec   `json:"spec,omitempty"`
	Status CloudNumberRegistryRealmStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudNumberRegistryRealmList contains a list of CloudNumberRegistryRealm
type CloudNumberRegistryRealmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudNumberRegistryRealm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudNumberRegistryRealm{}, &CloudNumberRegistryRealmList{})
}
