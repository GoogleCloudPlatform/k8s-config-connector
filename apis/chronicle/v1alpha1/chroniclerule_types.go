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

// API sources for ChronicleRule, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/chronicle/v1/rule.proto
// +kcc:source:service-docs=https://cloud.google.com/chronicle/docs/secops/secops-overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/chronicle/docs/reference/rest/v1/projects.locations.instances.rules

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ChronicleRuleGVK = GroupVersion.WithKind("ChronicleRule")

// ChronicleRuleSpec defines the desired state of ChronicleRule
// +kcc:spec:proto=google.cloud.chronicle.v1.Rule
// +kcc:required-from-proto
type ChronicleRuleSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ChronicleRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The YARA-L content of the rule.
	//  Populated in FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.text
	Text *string `json:"text,omitempty"`

	// The etag for this rule.
	//  If this is provided on update, the request will succeed if and only if it
	//  matches the server-computed value, and will fail with an ABORTED error
	//  otherwise.
	//  Populated in BASIC view and FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.etag
	Etag *string `json:"etag,omitempty"`

	// Resource name of the DataAccessScope bound to this rule.
	//  Populated in BASIC view and FULL view.
	//  If reference lists are used in the rule, validations will be performed
	//  against this scope to ensure that the reference lists are compatible with
	//  both the user's and the rule's scopes.
	//  The scope should be in the format:
	//  `projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope}`.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.scope
	Scope *string `json:"scope,omitempty"`
}

// ChronicleRuleStatus defines the config connector machine state of ChronicleRule
type ChronicleRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ChronicleRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ChronicleRuleObservedState `json:"observedState,omitempty"`
}

// ChronicleRuleObservedState is the state of the ChronicleRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.chronicle.v1.Rule
type ChronicleRuleObservedState struct {
	// Output only. The revision ID of the rule.
	//  A new revision is created whenever the rule text is changed in any way.
	//  Format: `v_{10 digits}_{9 digits}`
	//  Populated in REVISION_METADATA_ONLY view and FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. Display name of the rule.
	//  Populated in BASIC view and FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The author of the rule. Extracted from the meta section of
	//  text. Populated in BASIC view and FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.author
	Author *string `json:"author,omitempty"`

	// Output only. The severity of the rule as specified in the meta section of
	//  text. Populated in BASIC view and FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.severity
	Severity *Severity `json:"severity,omitempty"`

	// Output only. Additional metadata specified in the meta section of text.
	//  Populated in FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Output only. The timestamp of when the rule was created.
	//  Populated in FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp of when the rule revision was created.
	//  Populated in FULL, REVISION_METADATA_ONLY views.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. The current compilation state of the rule.
	//  Populated in FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.compilation_state
	CompilationState *string `json:"compilationState,omitempty"`

	// Output only. User-facing type of the rule. Extracted from the events
	//  section of rule text. Populated in BASIC view and FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.type
	Type *string `json:"type,omitempty"`

	// Output only. Resource names of the reference lists used in this rule.
	//  Populated in FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.reference_lists
	ReferenceLists []string `json:"referenceLists,omitempty"`

	// Output only. The run frequencies that are allowed for the rule.
	//  Populated in BASIC view and FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.allowed_run_frequencies
	AllowedRunFrequencies []string `json:"allowedRunFrequencies,omitempty"`

	// Output only. A list of a rule's corresponding compilation diagnostic
	//  messages such as compilation errors and compilation warnings. Populated in
	//  FULL view.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.compilation_diagnostics
	CompilationDiagnostics []CompilationDiagnosticObservedState `json:"compilationDiagnostics,omitempty"`

	// Output only. Indicate the rule can run in near real time live rule.
	//  If this is true, the rule uses the near real time live rule when the run
	//  frequency is set to LIVE.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.near_real_time_live_rule_eligible
	NearRealTimeLiveRuleEligible *bool `json:"nearRealTimeLiveRuleEligible,omitempty"`

	// Output only. The set of inputs used in the rule. For example, if the rule
	//  uses $e.principal.hostname, then the uses_udm field will be true.
	// +kcc:proto:field=google.cloud.chronicle.v1.Rule.inputs_used
	InputsUsed *InputsUsed `json:"inputsUsed,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpchroniclerule;gcpchroniclerules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ChronicleRule is the Schema for the ChronicleRule API
// +k8s:openapi-gen=true
type ChronicleRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ChronicleRuleSpec   `json:"spec,omitempty"`
	Status ChronicleRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ChronicleRuleList contains a list of ChronicleRule
type ChronicleRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronicleRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronicleRule{}, &ChronicleRuleList{})
}
