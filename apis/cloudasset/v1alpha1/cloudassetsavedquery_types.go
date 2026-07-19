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

var CloudAssetSavedQueryGVK = GroupVersion.WithKind("CloudAssetSavedQuery")

// CloudAssetSavedQuerySpec defines the desired state of CloudAssetSavedQuery
// +kcc:spec:proto=google.cloud.asset.v1.SavedQuery
type CloudAssetSavedQuerySpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The CloudAssetSavedQuery name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// The description of this saved query. This value should be fewer than 255
	//  characters.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.description
	// +optional
	Description *string `json:"description,omitempty"`

	// Labels applied on the resource.
	//  This value should not contain more than 10 entries. The key and value of
	//  each entry must be non-empty and fewer than 64 characters.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.labels
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// The query content.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.content
	// +optional
	Content *SavedQuery_QueryContent `json:"content,omitempty"`
}

// CloudAssetSavedQueryStatus defines the config connector machine state of CloudAssetSavedQuery
type CloudAssetSavedQueryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudAssetSavedQuery resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudAssetSavedQueryObservedState `json:"observedState,omitempty"`
}

// CloudAssetSavedQueryObservedState is the state of the CloudAssetSavedQuery resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.asset.v1.SavedQuery
type CloudAssetSavedQueryObservedState struct {
	// Output only. The create time of this saved query.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The account's email address who has created this saved query.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. The last update time of this saved query.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// Output only. The account's email address who has updated this saved query
	//  most recently.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.last_updater
	LastUpdater *string `json:"lastUpdater,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudassetsavedquery;gcpcloudassetsavedqueries
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudAssetSavedQuery is the Schema for the CloudAssetSavedQuery API
// +k8s:openapi-gen=true
type CloudAssetSavedQuery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudAssetSavedQuerySpec   `json:"spec,omitempty"`
	Status CloudAssetSavedQueryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudAssetSavedQueryList contains a list of CloudAssetSavedQuery
type CloudAssetSavedQueryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudAssetSavedQuery `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudAssetSavedQuery{}, &CloudAssetSavedQueryList{})
}
