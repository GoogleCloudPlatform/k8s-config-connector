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

// API sources for DataplexAsset, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dataplex/v1/resources.proto
// +kcc:source:service-docs=https://cloud.google.com/dataplex/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataplex/docs/reference/rest/v1/projects.locations.lakes.zones.assets

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexAssetGVK = GroupVersion.WithKind("DataplexAsset")

// DataplexAssetSpec defines the desired state of DataplexAsset
// +kcc:spec:proto=google.cloud.dataplex.v1.Asset
// +kcc:required-from-proto
type DataplexAssetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DataplexAsset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined labels for the asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.description
	Description *string `json:"description,omitempty"`

	// Required. Specification of the resource that is referenced by this asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.resource_spec
	// +required
	ResourceSpec *Asset_ResourceSpec `json:"resourceSpec,omitempty"`

	// Optional. Specification of the discovery feature applied to data referenced
	//  by this asset. When this spec is left unset, the asset will use the spec
	//  set on the parent zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.discovery_spec
	DiscoverySpec *Asset_DiscoverySpec `json:"discoverySpec,omitempty"`
}

// DataplexAssetStatus defines the config connector machine state of DataplexAsset
type DataplexAssetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexAsset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexAssetObservedState `json:"observedState,omitempty"`
}

// DataplexAssetObservedState is the state of the DataplexAsset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.Asset
type DataplexAssetObservedState struct {
	// Output only. System generated globally unique ID for the asset. This ID
	//  will be different if the asset is deleted and re-created with the same
	//  name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the asset was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the asset was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.state
	State *string `json:"state,omitempty"`

	// Output only. Status of the resource referenced by this asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.resource_status
	ResourceStatus *Asset_ResourceStatusObservedState `json:"resourceStatus,omitempty"`

	// Output only. Status of the security policy applied to resource referenced
	//  by this asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.security_status
	SecurityStatus *Asset_SecurityStatus `json:"securityStatus,omitempty"`

	// Output only. Status of the discovery feature applied to data referenced by
	//  this asset.
	// +kcc:proto:field=google.cloud.dataplex.v1.Asset.discovery_status
	DiscoveryStatus *Asset_DiscoveryStatus `json:"discoveryStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexasset;gcpdataplexassets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexAsset is the Schema for the DataplexAsset API
// +k8s:openapi-gen=true
type DataplexAsset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexAssetSpec   `json:"spec,omitempty"`
	Status DataplexAssetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexAssetList contains a list of DataplexAsset
type DataplexAssetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexAsset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexAsset{}, &DataplexAssetList{})
}
