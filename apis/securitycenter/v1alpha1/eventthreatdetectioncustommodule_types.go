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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecurityCenterEventThreatDetectionCustomModuleGVK = GroupVersion.WithKind("SecurityCenterEventThreatDetectionCustomModule")

// SecurityCenterEventThreatDetectionCustomModuleSpec defines the desired state of SecurityCenterEventThreatDetectionCustomModule
// +kcc:spec:proto=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule
type SecurityCenterEventThreatDetectionCustomModuleSpec struct {
	// The organization, folder, or project that this resource belongs to.
	// +optional
	OrganizationRef *refs.OrganizationRef `json:"organizationRef,omitempty"`

	// +optional
	FolderRef *refs.FolderRef `json:"folderRef,omitempty"`

	// +optional
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Immutable. The SecurityCenterEventThreatDetectionCustomModule name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Config for the module. For the resident module, its config value is defined
	// at this level. For the inherited module, its config value is inherited from
	// the ancestor module.
	// +optional
	Config *apiextensionsv1.JSON `json:"config,omitempty"`

	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED", "INHERITED"]
	// +optional
	EnablementState *string `json:"enablementState,omitempty"`

	// Type for the module. e.g. CONFIGURABLE_BAD_IP.
	// +optional
	Type *string `json:"type,omitempty"`

	// The human readable name to be displayed for the module.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// The description for the module.
	// +optional
	Description *string `json:"description,omitempty"`
}

// SecurityCenterEventThreatDetectionCustomModuleStatus defines the observed state of SecurityCenterEventThreatDetectionCustomModule
type SecurityCenterEventThreatDetectionCustomModuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecurityCenterEventThreatDetectionCustomModule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecurityCenterEventThreatDetectionCustomModuleObservedState `json:"observedState,omitempty"`
}

// SecurityCenterEventThreatDetectionCustomModuleObservedState defines the observed state of SecurityCenterEventThreatDetectionCustomModule
// +kcc:observedstate:proto=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule
type SecurityCenterEventThreatDetectionCustomModuleObservedState struct {
	// Output only. The time at which the custom module was last updated.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The editor that last updated the custom module.
	// +optional
	LastEditor *string `json:"lastEditor,omitempty"`

	// Output only. The closest ancestor module that this module inherits the
	// enablement state from.
	// +optional
	AncestorModule *string `json:"ancestorModule,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuritycentereventthreatdetectioncustommodule;gcpsecuritycentereventthreatdetectioncustommodules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecurityCenterEventThreatDetectionCustomModule is the Schema for the SecurityCenterEventThreatDetectionCustomModule API
// +k8s:openapi-gen=true
type SecurityCenterEventThreatDetectionCustomModule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityCenterEventThreatDetectionCustomModuleSpec   `json:"spec,omitempty"`
	Status SecurityCenterEventThreatDetectionCustomModuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecurityCenterEventThreatDetectionCustomModuleList contains a list of SecurityCenterEventThreatDetectionCustomModule
type SecurityCenterEventThreatDetectionCustomModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterEventThreatDetectionCustomModule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecurityCenterEventThreatDetectionCustomModule{}, &SecurityCenterEventThreatDetectionCustomModuleList{})
}
