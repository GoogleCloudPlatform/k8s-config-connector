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

// API sources for TelcoAutomationBlueprint, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/telcoautomation/v1/telcoautomation.proto
// +kcc:source:service-docs=https://cloud.google.com/telecom-network-automation
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/telecom-network-automation/docs/reference/rest/v1/projects.locations.orchestrationClusters.blueprints

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TelcoAutomationBlueprintGVK = GroupVersion.WithKind("TelcoAutomationBlueprint")

// TelcoAutomationBlueprintSpec defines the desired state of TelcoAutomationBlueprint
// +kcc:spec:proto=google.cloud.telcoautomation.v1.Blueprint
// +kcc:required-from-proto
type TelcoAutomationBlueprintSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The TelcoAutomationBlueprint name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Immutable. The public blueprint ID from which this blueprint was
	//  created.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.source_blueprint
	// +required
	SourceBlueprint *string `json:"sourceBlueprint,omitempty"`

	// Optional. Human readable name of a Blueprint.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Files present in a blueprint.
	//  When invoking UpdateBlueprint API, only the modified files should be
	//  included in this. Files that are not included in the update of a blueprint
	//  will not be changed.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.files
	Files []File `json:"files,omitempty"`

	// Optional. Labels are key-value attributes that can be set on a blueprint
	//  resource by the user.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// TelcoAutomationBlueprintStatus defines the config connector machine state of TelcoAutomationBlueprint
type TelcoAutomationBlueprintStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TelcoAutomationBlueprint resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TelcoAutomationBlueprintObservedState `json:"observedState,omitempty"`
}

// TelcoAutomationBlueprintObservedState is the state of the TelcoAutomationBlueprint resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.telcoautomation.v1.Blueprint
type TelcoAutomationBlueprintObservedState struct {
	// Output only. Immutable. The revision ID of the blueprint.
	//  A new revision is committed whenever a blueprint is approved.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. The timestamp that the revision was created.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. Approval state of the blueprint (DRAFT, PROPOSED, APPROVED)
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.approval_state
	ApprovalState *string `json:"approvalState,omitempty"`

	// Output only. Name of the repository where the blueprint files are stored.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.repository
	Repository *string `json:"repository,omitempty"`

	// Output only. Blueprint creation time.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the blueprint was updated.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Source provider is the author of a public blueprint, from
	//  which this blueprint is created.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.source_provider
	SourceProvider *string `json:"sourceProvider,omitempty"`

	// Output only. DeploymentLevel of a blueprint signifies where the blueprint
	//  will be applied. e.g. [HYDRATION, SINGLE_DEPLOYMENT, MULTI_DEPLOYMENT]
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.deployment_level
	DeploymentLevel *string `json:"deploymentLevel,omitempty"`

	// Output only. Indicates if the deployment created from this blueprint can be
	//  rolled back.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Blueprint.rollback_support
	RollbackSupport *bool `json:"rollbackSupport,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptelcoautomationblueprint;gcptelcoautomationblueprints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TelcoAutomationBlueprint is the Schema for the TelcoAutomationBlueprint API
// +k8s:openapi-gen=true
type TelcoAutomationBlueprint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TelcoAutomationBlueprintSpec   `json:"spec,omitempty"`
	Status TelcoAutomationBlueprintStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TelcoAutomationBlueprintList contains a list of TelcoAutomationBlueprint
type TelcoAutomationBlueprintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TelcoAutomationBlueprint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TelcoAutomationBlueprint{}, &TelcoAutomationBlueprintList{})
}
