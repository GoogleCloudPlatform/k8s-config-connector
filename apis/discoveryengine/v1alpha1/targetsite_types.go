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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineDataStoreTargetSiteGVK = GroupVersion.WithKind("DiscoveryEngineDataStoreTargetSite")

// DiscoveryEngineDataStoreTargetSiteSpec defines the desired state of DiscoveryEngineDataStoreTargetSite
// +kcc:spec:proto=google.cloud.discoveryengine.v1.TargetSite
type DiscoveryEngineDataStoreTargetSiteSpec struct {
	// The DataStore this target site should be part of.
	DataStoreRef *DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`

	// The resource ID is server-generated, so no ResourceID field

	// Required. Input only. The user provided URI pattern from which the
	// `generated_uri_pattern` is generated.
	ProvidedURIPattern *string `json:"providedURIPattern,omitempty"`

	// The type of the target site, e.g., whether the site is to be included or
	// excluded.
	Type *string `json:"type,omitempty"`

	// Input only. If set to false, a uri_pattern is generated to include all
	// pages whose address contains the provided_uri_pattern. If set to true, an
	// uri_pattern is generated to try to be an exact match of the
	// provided_uri_pattern or just the specific page if the provided_uri_pattern
	// is a specific one. provided_uri_pattern is always normalized to
	// generate the URI pattern to be used by the search engine.
	ExactMatch *bool `json:"exactMatch,omitempty"`
}

// DiscoveryEngineDataStoreTargetSiteStatus defines the config connector machine state of DiscoveryEngineDataStoreTargetSite
type DiscoveryEngineDataStoreTargetSiteStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineDataStoreTargetSite resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineDataStoreTargetSiteObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineDataStoreTargetSiteObservedState is the state of the DiscoveryEngineDataStoreTargetSite resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.TargetSite
type DiscoveryEngineDataStoreTargetSiteObservedState struct {
	// Output only. This is system-generated based on the provided_uri.
	GeneratedURIPattern *string `json:"generatedURIPattern,omitempty"`

	// Output only. Root domain of the provided_uri.
	RootDomainURI *string `json:"rootDomainURI,omitempty"`

	// Output only. Site ownership and validity verification status.
	SiteVerificationInfo *SiteVerificationInfo `json:"siteVerificationInfo,omitempty"`

	// Output only. Indexing status.
	IndexingStatus *string `json:"indexingStatus,omitempty"`

	// Output only. The target site's last updated time.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Failure reason.
	FailureReason *TargetSite_FailureReason `json:"failureReason,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginedatastoretargetsite;gcpdiscoveryenginedatastoretargetsites
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineDataStoreTargetSite is the Schema for the DiscoveryEngineDataStoreTargetSite API
// +k8s:openapi-gen=true
type DiscoveryEngineDataStoreTargetSite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineDataStoreTargetSiteSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineDataStoreTargetSiteStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineDataStoreTargetSiteList contains a list of DiscoveryEngineDataStoreTargetSite
type DiscoveryEngineDataStoreTargetSiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineDataStoreTargetSite `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineDataStoreTargetSite{}, &DiscoveryEngineDataStoreTargetSiteList{})
}
