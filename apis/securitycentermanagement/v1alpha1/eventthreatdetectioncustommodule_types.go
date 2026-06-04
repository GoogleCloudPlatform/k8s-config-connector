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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecurityCenterManagementEventThreatDetectionCustomModuleGVK = GroupVersion.WithKind("SecurityCenterManagementEventThreatDetectionCustomModule")

// SecurityCenterManagementEventThreatDetectionCustomModuleSpec defines the desired state of SecurityCenterManagementEventThreatDetectionCustomModule
// +kcc:spec:proto=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule
type SecurityCenterManagementEventThreatDetectionCustomModuleSpec struct {
	// The organization that this resource belongs to.
	// +kubebuilder:validation:Optional
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef,omitempty"`

	// The folder that this resource belongs to.
	// +kubebuilder:validation:Optional
	FolderRef *refsv1beta1.FolderRef `json:"folderRef,omitempty"`

	// The project that this resource belongs to.
	// +kubebuilder:validation:Optional
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The SecurityCenterManagementEventThreatDetectionCustomModule name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Configuration for the module. For the resident module, its
	// configuration value is defined at this level. For the inherited module, its
	// configuration value is inherited from the ancestor module.
	// +kubebuilder:validation:Optional
	Config *apiextensionsv1.JSON `json:"config,omitempty"`

	// Optional. The state of enablement for the module at the given level of the
	// hierarchy.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=ENABLEMENT_STATE_UNSPECIFIED;ENABLED;DISABLED;INHERITED
	EnablementState *string `json:"enablementState,omitempty"`

	// Optional. Type for the module. For example, `CONFIGURABLE_BAD_IP`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty"`

	// Optional. The human-readable name of the module.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty"`
}

// SecurityCenterManagementEventThreatDetectionCustomModuleStatus defines the config connector machine state of SecurityCenterManagementEventThreatDetectionCustomModule
type SecurityCenterManagementEventThreatDetectionCustomModuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecurityCenterManagementEventThreatDetectionCustomModule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecurityCenterManagementEventThreatDetectionCustomModuleObservedState `json:"observedState,omitempty"`
}

// SecurityCenterManagementEventThreatDetectionCustomModuleObservedState is the state of the SecurityCenterManagementEventThreatDetectionCustomModule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule
type SecurityCenterManagementEventThreatDetectionCustomModuleObservedState struct {
	// Output only. The closest ancestor module that this module inherits the
	// enablement state from. If empty, indicates that the custom module was
	// created in the requesting parent organization, folder, or project. The
	// format is the same as the custom module's resource name.
	AncestorModule *string `json:"ancestorModule,omitempty"`

	// Output only. The time the module was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The editor the module was last updated by.
	LastEditor *string `json:"lastEditor,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuritycentermanagementeventthreatdetectioncustommodule;gcpsecuritycentermanagementeventthreatdetectioncustommodules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecurityCenterManagementEventThreatDetectionCustomModule is the Schema for the SecurityCenterManagementEventThreatDetectionCustomModule API
// +k8s:openapi-gen=true
type SecurityCenterManagementEventThreatDetectionCustomModule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecurityCenterManagementEventThreatDetectionCustomModuleSpec   `json:"spec,omitempty"`
	Status SecurityCenterManagementEventThreatDetectionCustomModuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecurityCenterManagementEventThreatDetectionCustomModuleList contains a list of SecurityCenterManagementEventThreatDetectionCustomModule
type SecurityCenterManagementEventThreatDetectionCustomModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterManagementEventThreatDetectionCustomModule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecurityCenterManagementEventThreatDetectionCustomModule{}, &SecurityCenterManagementEventThreatDetectionCustomModuleList{})
}
