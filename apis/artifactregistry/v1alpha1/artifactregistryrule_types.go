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

// API sources for ArtifactRegistryRule, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/devtools/artifactregistry/v1/rule.proto
// +kcc:source:service-docs=https://cloud.google.com/artifacts/docs/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/artifacts/docs/reference/rest/v1/projects.locations.repositories.rules

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ArtifactRegistryRuleGVK = GroupVersion.WithKind("ArtifactRegistryRule")

// ArtifactRegistryRuleSpec defines the desired state of ArtifactRegistryRule
// +kcc:spec:proto=google.devtools.artifactregistry.v1.Rule
// +kcc:required-from-proto
type ArtifactRegistryRuleSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ArtifactRegistryRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The action this rule takes.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Rule.action
	Action *string `json:"action,omitempty"`

	// +kcc:proto:field=google.devtools.artifactregistry.v1.Rule.operation
	Operation *string `json:"operation,omitempty"`

	// Optional. A CEL expression for conditions that must be met in order for the
	//  rule to apply. If not provided, the rule matches all objects.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Rule.condition
	Condition *Expr `json:"condition,omitempty"`

	// The package ID the rule applies to.
	//  If empty, this rule applies to all packages inside the repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Rule.package_id
	PackageID *string `json:"packageID,omitempty"`
}

// ArtifactRegistryRuleStatus defines the config connector machine state of ArtifactRegistryRule
type ArtifactRegistryRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ArtifactRegistryRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ArtifactRegistryRuleObservedState `json:"observedState,omitempty"`
}

// ArtifactRegistryRuleObservedState is the state of the ArtifactRegistryRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devtools.artifactregistry.v1.Rule
type ArtifactRegistryRuleObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpartifactregistryrule;gcpartifactregistryrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ArtifactRegistryRule is the Schema for the ArtifactRegistryRule API
// +k8s:openapi-gen=true
type ArtifactRegistryRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ArtifactRegistryRuleSpec   `json:"spec,omitempty"`
	Status ArtifactRegistryRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ArtifactRegistryRuleList contains a list of ArtifactRegistryRule
type ArtifactRegistryRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArtifactRegistryRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArtifactRegistryRule{}, &ArtifactRegistryRuleList{})
}
