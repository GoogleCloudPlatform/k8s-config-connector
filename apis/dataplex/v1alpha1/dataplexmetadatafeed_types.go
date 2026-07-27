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
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexMetadataFeedGVK = GroupVersion.WithKind("DataplexMetadataFeed")

// DataplexMetadataFeedSpec defines the desired state of DataplexMetadataFeed
// +kcc:spec:proto=google.cloud.dataplex.v1.MetadataFeed
type DataplexMetadataFeedSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The DataplexMetadataFeed name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The scope of the metadata feed.
	// Only the in scope changes are published.
	// +kubebuilder:validation:Required
	Scope *MetadataFeed_Scope `json:"scope"`

	// Optional. The filters of the metadata feed.
	// Only the changes that match the filters are published.
	Filters *MetadataFeed_Filters `json:"filters,omitempty"`

	// Optional. User-defined labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The pubsub topic that you want the metadata feed messages to
	// publish to. Please grant Dataplex service account the permission to
	// publish messages to the topic. The service account is:
	// service-{PROJECT_NUMBER}@gcp-sa-dataplex.iam.gserviceaccount.com.
	PubsubTopicRef *pubsubv1beta1.PubSubTopicRef `json:"pubsubTopicRef,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataFeed.Scope
type MetadataFeed_Scope struct {
	// Optional. Whether the metadata feed is at the organization-level.
	//
	// - If `true`, all changes happened to the entries in the same
	// organization as the feed are published.
	// - If `false`, you must specify a list of projects or a list of entry
	// groups whose entries you want to listen to.
	//
	// The default is `false`.
	OrganizationLevel *bool `json:"organizationLevel,omitempty"`

	// Optional. The projects whose entries you want to listen to.
	// Must be in the same organization as the feed.
	// Must be in the format: `projects/{project_id_or_number}`.
	ProjectRefs []refsv1beta1.ProjectRef `json:"projectRefs,omitempty"`

	// Optional. The entry groups whose entries you want to listen to.
	// Must be in the format:
	// `projects/{project_id_or_number}/locations/{location_id}/entryGroups/{entry_group_id}`.
	EntryGroupRefs []EntryGroupRef `json:"entryGroupRefs,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataFeed.Filters
type MetadataFeed_Filters struct {
	// Optional. The entry types that you want to listen to, specified as
	// relative resource names in the format
	// `projects/{project_id_or_number}/locations/{location}/entryTypes/{entry_type_id}`.
	// Only entries that belong to the specified entry types are published.
	EntryTypeRefs []EntryTypeRef `json:"entryTypeRefs,omitempty"`

	// Optional. The aspect types that you want to listen to. Depending on how
	// the aspect is attached to the entry, in the format:
	// `projects/{project_id_or_number}/locations/{location}/aspectTypes/{aspect_type_id}`.
	AspectTypeRefs []AspectTypeRef `json:"aspectTypeRefs,omitempty"`

	// Optional. The type of change that you want to listen to.
	// If not specified, all changes are published.
	// +kubebuilder:validation:items:Enum=CREATE;UPDATE;DELETE
	ChangeTypes []string `json:"changeTypes,omitempty"`
}

// DataplexMetadataFeedStatus defines the config connector machine state of DataplexMetadataFeed
type DataplexMetadataFeedStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexMetadataFeed resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexMetadataFeedObservedState `json:"observedState,omitempty"`
}

// DataplexMetadataFeedObservedState is the state of the DataplexMetadataFeed resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.MetadataFeed
type DataplexMetadataFeedObservedState struct {
	// Output only. System generated globally unique ID for the feed. This ID will
	// be different if the feed is deleted and re-created with the same name.
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the feed was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the feed was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexmetadatafeed;gcpdataplexmetadatafeeds
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexMetadataFeed is the Schema for the DataplexMetadataFeed API
// +k8s:openapi-gen=true
type DataplexMetadataFeed struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexMetadataFeedSpec   `json:"spec,omitempty"`
	Status DataplexMetadataFeedStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexMetadataFeedList contains a list of DataplexMetadataFeed
type DataplexMetadataFeedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexMetadataFeed `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexMetadataFeed{}, &DataplexMetadataFeedList{})
}
