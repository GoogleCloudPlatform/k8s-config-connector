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

// API sources for SecuritycenterEventThreatDetectionCustomModule, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/securitycenter/v1/event_threat_detection_custom_module.proto
// +kcc:source:service-docs=https://cloud.google.com/security-command-center
// +kcc:source:resource-docs=https://docs.cloud.google.com/security-command-center/docs/reference/rest/v1/organizations.eventThreatDetectionSettings.customModules

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecuritycenterEventThreatDetectionCustomModuleGVK = GroupVersion.WithKind("SecuritycenterEventThreatDetectionCustomModule")

// SecuritycenterEventThreatDetectionCustomModuleSpec defines the desired state of SecuritycenterEventThreatDetectionCustomModule
// +kcc:spec:proto=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule
// +kcc:required-from-proto
type SecuritycenterEventThreatDetectionCustomModuleSpec struct {
	// The organization that this resource belongs to.
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The SecuritycenterEventThreatDetectionCustomModule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Config for the module. For the resident module, its config value is defined
	//  at this level. For the inherited module, its config value is inherited from
	//  the ancestor module.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.config
	Config apiextensionsv1.JSON `json:"config,omitempty"`

	// The state of enablement for the module at the given level of the hierarchy.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.enablement_state
	EnablementState *string `json:"enablementState,omitempty"`

	// Type for the module. e.g. CONFIGURABLE_BAD_IP.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.type
	Type *string `json:"type,omitempty"`

	// The human readable name to be displayed for the module.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description for the module.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.description
	Description *string `json:"description,omitempty"`
}

// SecuritycenterEventThreatDetectionCustomModuleStatus defines the config connector machine state of SecuritycenterEventThreatDetectionCustomModule
type SecuritycenterEventThreatDetectionCustomModuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecuritycenterEventThreatDetectionCustomModule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecuritycenterEventThreatDetectionCustomModuleObservedState `json:"observedState,omitempty"`
}

// SecuritycenterEventThreatDetectionCustomModuleObservedState is the state of the SecuritycenterEventThreatDetectionCustomModule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule
type SecuritycenterEventThreatDetectionCustomModuleObservedState struct {
	// Output only. The closest ancestor module that this module inherits the
	//  enablement state from. The format is the same as the
	//  EventThreatDetectionCustomModule resource name.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.ancestor_module
	AncestorModule *string `json:"ancestorModule,omitempty"`

	// Output only. The time the module was last updated.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The editor the module was last updated by.
	// +kcc:proto:field=google.cloud.securitycenter.v1.EventThreatDetectionCustomModule.last_editor
	LastEditor *string `json:"lastEditor,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuritycentereventthreatdetectioncustommodule;gcpsecuritycentereventthreatdetectioncustommodules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecuritycenterEventThreatDetectionCustomModule is the Schema for the SecuritycenterEventThreatDetectionCustomModule API
// +k8s:openapi-gen=true
type SecuritycenterEventThreatDetectionCustomModule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecuritycenterEventThreatDetectionCustomModuleSpec   `json:"spec,omitempty"`
	Status SecuritycenterEventThreatDetectionCustomModuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecuritycenterEventThreatDetectionCustomModuleList contains a list of SecuritycenterEventThreatDetectionCustomModule
type SecuritycenterEventThreatDetectionCustomModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecuritycenterEventThreatDetectionCustomModule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecuritycenterEventThreatDetectionCustomModule{}, &SecuritycenterEventThreatDetectionCustomModuleList{})
}
