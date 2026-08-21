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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineSitemapGVK = GroupVersion.WithKind("DiscoveryEngineSitemap")

// DiscoveryEngineSitemapSpec defines the desired state of DiscoveryEngineSitemap
// +kcc:spec:proto=google.cloud.discoveryengine.v1beta.Sitemap
type DiscoveryEngineSitemapSpec struct {
	// Immutable. The DataStore this sitemap belongs to.
	// +required
	DataStoreRef *DiscoveryEngineDataStoreRef `json:"dataStoreRef"`

	// The DiscoveryEngineSitemap name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Public URI for the sitemap, e.g. `www.example.com/sitemap.xml`.
	// +required
	URI *string `json:"uri,omitempty"`
}

// DiscoveryEngineSitemapStatus defines the config connector machine state of DiscoveryEngineSitemap
type DiscoveryEngineSitemapStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineSitemap resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineSitemapObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineSitemapObservedState is the state of the DiscoveryEngineSitemap resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1beta.Sitemap
type DiscoveryEngineSitemapObservedState struct {
	// Output only. The fully qualified resource name of the sitemap.
	// `projects/*/locations/*/collections/*/dataStores/*/siteSearchEngine/sitemaps/*`
	// The `sitemap_id` suffix is system-generated.
	Name *string `json:"name,omitempty"`

	// Output only. The sitemap's creation time.
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginesitemap;gcpdiscoveryenginesitemaps
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineSitemap is the Schema for the DiscoveryEngineSitemap API
// +k8s:openapi-gen=true
type DiscoveryEngineSitemap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineSitemapSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineSitemapStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineSitemapList contains a list of DiscoveryEngineSitemap
type DiscoveryEngineSitemapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineSitemap `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineSitemap{}, &DiscoveryEngineSitemapList{})
}
