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

// API sources for DLPContentPolicy, recorded by generate-types:
// +kcc:source:proto=https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/google/privacy/dlp/v2/dlp.proto
// +kcc:source:service-docs=https://cloud.google.com/sensitive-data-protection/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/sensitive-data-protection/docs/reference/rest/v2/projects.locations.contentPolicies

package v1alpha1

import (
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DLPContentPolicyGVK = GroupVersion.WithKind("DLPContentPolicy")

// DLPContentPolicySpec defines the desired state of DLPContentPolicy
// +kcc:spec:proto=google.privacy.dlp.v2.ContentPolicy
// +kcc:required-from-proto
type DLPContentPolicySpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// The DLPContentPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Display name (max 63 chars)
	// +kcc:proto:field=google.privacy.dlp.v2.ContentPolicy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. InspectConfig to use to produce findings.
	// +kcc:proto:field=google.privacy.dlp.v2.ContentPolicy.inspect_config
	InspectConfig *InspectConfig `json:"inspectConfig,omitempty"`

	// Required. Policies to apply, based on the findings returned by inspection.
	//  The first rule to match applies.
	// +kcc:proto:field=google.privacy.dlp.v2.ContentPolicy.rules
	// +required
	Rules []ContentPolicy_PolicyRule `json:"rules,omitempty"`

	// Optional. Action to take if the content is an unsupported file type.
	// +kcc:proto:field=google.privacy.dlp.v2.ContentPolicy.unsupported_file_type
	UnsupportedFileType *ContentPolicy_PolicyAction `json:"unsupportedFileType,omitempty"`

	// Optional. Action to take if the content is a supported file type but is too
	//  large to be scanned.
	// +kcc:proto:field=google.privacy.dlp.v2.ContentPolicy.input_too_large
	InputTooLarge *ContentPolicy_PolicyAction `json:"inputTooLarge,omitempty"`

	// Optional. Action to take if the content is a supported file type and size
	//  but fails to be scanned, for example because the file is encrypted or
	//  corrupted.
	// +kcc:proto:field=google.privacy.dlp.v2.ContentPolicy.failed_to_scan_supported_file_type
	FailedToScanSupportedFileType *ContentPolicy_PolicyAction `json:"failedToScanSupportedFileType,omitempty"`

	// Action to take if the content is scanned and no rules match.
	//  Defaults to returning an ALLOW verdict if not set.
	// +kcc:proto:field=google.privacy.dlp.v2.ContentPolicy.default_action
	DefaultAction *ContentPolicy_PolicyAction `json:"defaultAction,omitempty"`

	// Optional. Log the actions taken by the content policy to external systems.
	// +kcc:proto:field=google.privacy.dlp.v2.ContentPolicy.logging_configs
	LoggingConfigs []ContentPolicy_LoggingConfig `json:"loggingConfigs,omitempty"`
}

// DLPContentPolicyStatus defines the config connector machine state of DLPContentPolicy
type DLPContentPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DLPContentPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DLPContentPolicyObservedState `json:"observedState,omitempty"`
}

// DLPContentPolicyObservedState is the state of the DLPContentPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.privacy.dlp.v2.ContentPolicy
type DLPContentPolicyObservedState struct {
	// Output only. The creation timestamp of a contentPolicy; output-only field.
	// +kcc:proto:field=google.privacy.dlp.v2.ContentPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of a contentPolicy; output-only
	//  field.
	// +kcc:proto:field=google.privacy.dlp.v2.ContentPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A stream of errors encountered when the policy was applied.
	//  Output only field. Will return the last 100 errors. Whenever the policy is
	//  modified this list will be cleared.
	// +kcc:proto:field=google.privacy.dlp.v2.ContentPolicy.errors
	Errors []Error `json:"errors,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdlpcontentpolicy;gcpdlpcontentpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DLPContentPolicy is the Schema for the DLPContentPolicy API
// +k8s:openapi-gen=true
type DLPContentPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DLPContentPolicySpec   `json:"spec,omitempty"`
	Status DLPContentPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DLPContentPolicyList contains a list of DLPContentPolicy
type DLPContentPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DLPContentPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DLPContentPolicy{}, &DLPContentPolicyList{})
}
