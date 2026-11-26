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

// API sources for AgentRegistrySkill, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/25fb8ea2d4bfa8fd41e24b28029e34a4adc35db0/google/cloud/agentregistry/v1/skill.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/agent-registry/overview
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/agent-registry/overview/docs/reference/rest/v1/projects.locations.skills

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AgentRegistrySkillGVK = GroupVersion.WithKind("AgentRegistrySkill")

// AgentRegistrySkillSpec defines the desired state of AgentRegistrySkill
// +kcc:spec:proto=google.cloud.agentregistry.v1.Skill
// +kcc:required-from-proto
type AgentRegistrySkillSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AgentRegistrySkill name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Human-readable display name of the skill.
	//  Maximum length is 128 characters.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Brief summary describing the capabilities of the skill.
	//  Maximum length is 2048 characters.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.description
	Description *string `json:"description,omitempty"`

	// Required. Structural deployment type (SIMPLE leaf vs COMPOSITE bundle).
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.type
	// +required
	Type *string `json:"type,omitempty"`

	// Required. User-managed target state of the skill.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.target_state
	// +required
	TargetState *string `json:"targetState,omitempty"`

	// Optional. The full resource name of the revision currently served by
	//  default (floating track).
	//  Format:
	//  `projects/{project}/locations/{location}/skills/{skill}/revisions/{revision}`
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.default_revision
	DefaultRevision *string `json:"defaultRevision,omitempty"`

	// Optional. The publisher resource associated with this skill.
	//  Format: `projects/{project}/locations/{location}/publishers/{publisher}`
	//  The publisher dictates the allowed namespace prefixes for the skill's name
	//  and logical `skill_id` (e.g., Publisher `google` authorizes the `google-*`
	//  prefix).
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.publisher
	Publisher *string `json:"publisher,omitempty"`

	// Optional. Input only. Optional nested initial revision payload to support
	//  standard one-shot creation. The server processes this field on input during
	//  creation but must never return it in responses.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.initial_revision
	InitialRevision *SkillRevision `json:"initialRevision,omitempty"`

	// Optional. Optional. Scanning configuration for the skill.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.scanning_config
	ScanningConfig *ScanningConfig `json:"scanningConfig,omitempty"`
}

// AgentRegistrySkillStatus defines the config connector machine state of AgentRegistrySkill
type AgentRegistrySkillStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AgentRegistrySkill resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AgentRegistrySkillObservedState `json:"observedState,omitempty"`
}

// AgentRegistrySkillObservedState is the state of the AgentRegistrySkill resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.agentregistry.v1.Skill
type AgentRegistrySkillObservedState struct {
	// Output only. Universally unique identifier (UUID4) for the logical
	//  container.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The system-managed state of the skill.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.state
	State *string `json:"state,omitempty"`

	// Output only. Lightweight frontmatter metadata attributes copied from the
	//  default revision.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.frontmatter
	Frontmatter *FrontmatterObservedState `json:"frontmatter,omitempty"`

	// Output only. A stable, globally unique logical identifier for the skill.
	//  It is securely constructed by the backend by combining the associated
	//  `publisher`'s verified namespace and the skill's resource ID to enforce
	//  strict ownership. For example, the prefix `google-` is reserved exclusively
	//  for first-party Google publishers to prevent namespace squatting.
	//  Example: `urn:skill:google-workspace:create-docs`
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.skill_id
	SkillID *string `json:"skillID,omitempty"`

	// Output only. Create time.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Input only. Optional nested initial revision payload to support
	//  standard one-shot creation. The server processes this field on input during
	//  creation but must never return it in responses.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Skill.initial_revision
	InitialRevision *SkillRevisionObservedState `json:"initialRevision,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpagentregistryskill;gcpagentregistryskills
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AgentRegistrySkill is the Schema for the AgentRegistrySkill API
// +k8s:openapi-gen=true
type AgentRegistrySkill struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AgentRegistrySkillSpec   `json:"spec,omitempty"`
	Status AgentRegistrySkillStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AgentRegistrySkillList contains a list of AgentRegistrySkill
type AgentRegistrySkillList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentRegistrySkill `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AgentRegistrySkill{}, &AgentRegistrySkillList{})
}
