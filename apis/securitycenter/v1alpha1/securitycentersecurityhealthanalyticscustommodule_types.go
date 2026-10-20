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

// API sources for SecuritycenterSecurityHealthAnalyticsCustomModule, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/securitycenter/v1/security_health_analytics_custom_module.proto
// +kcc:source:service-docs=https://cloud.google.com/security-command-center
// +kcc:source:resource-docs=https://docs.cloud.google.com/security-command-center/docs/reference/rest/v1/organizations.securityHealthAnalyticsSettings.customModules

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecuritycenterSecurityHealthAnalyticsCustomModuleGVK = GroupVersion.WithKind("SecuritycenterSecurityHealthAnalyticsCustomModule")

// SecuritycenterSecurityHealthAnalyticsCustomModuleSpec defines the desired state of SecuritycenterSecurityHealthAnalyticsCustomModule
// +kcc:spec:proto=google.cloud.securitycenter.v1.SecurityHealthAnalyticsCustomModule
// +kcc:required-from-proto
type SecuritycenterSecurityHealthAnalyticsCustomModuleSpec struct {
	// The organization that this resource belongs to.
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The SecuritycenterSecurityHealthAnalyticsCustomModule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="the field is immutable"
	// The display name of the Security Health Analytics custom module. This
	//  display name becomes the finding category for all findings that are
	//  returned by this custom module. The display name must be between 1 and
	//  128 characters, start with a lowercase letter, and contain alphanumeric
	//  characters or underscores only.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityHealthAnalyticsCustomModule.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The enablement state of the custom module.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityHealthAnalyticsCustomModule.enablement_state
	EnablementState *string `json:"enablementState,omitempty"`

	// The user specified custom configuration for the module.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityHealthAnalyticsCustomModule.custom_config
	CustomConfig *CustomConfig `json:"customConfig,omitempty"`
}

// SecuritycenterSecurityHealthAnalyticsCustomModuleStatus defines the config connector machine state of SecuritycenterSecurityHealthAnalyticsCustomModule
type SecuritycenterSecurityHealthAnalyticsCustomModuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecuritycenterSecurityHealthAnalyticsCustomModule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecuritycenterSecurityHealthAnalyticsCustomModuleObservedState `json:"observedState,omitempty"`
}

// SecuritycenterSecurityHealthAnalyticsCustomModuleObservedState is the state of the SecuritycenterSecurityHealthAnalyticsCustomModule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.securitycenter.v1.SecurityHealthAnalyticsCustomModule
type SecuritycenterSecurityHealthAnalyticsCustomModuleObservedState struct {
	// Output only. The time at which the custom module was last updated.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityHealthAnalyticsCustomModule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The editor that last updated the custom module.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityHealthAnalyticsCustomModule.last_editor
	LastEditor *string `json:"lastEditor,omitempty"`

	// Output only. If empty, indicates that the custom module was created in the
	//  organization, folder, or project in which you are viewing the custom
	//  module. Otherwise, `ancestor_module` specifies the organization or folder
	//  from which the custom module is inherited.
	// +kcc:proto:field=google.cloud.securitycenter.v1.SecurityHealthAnalyticsCustomModule.ancestor_module
	AncestorModule *string `json:"ancestorModule,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuritycentersecurityhealthanalyticscustommodule;gcpsecuritycentersecurityhealthanalyticscustommodules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecuritycenterSecurityHealthAnalyticsCustomModule is the Schema for the SecuritycenterSecurityHealthAnalyticsCustomModule API
// +k8s:openapi-gen=true
type SecuritycenterSecurityHealthAnalyticsCustomModule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecuritycenterSecurityHealthAnalyticsCustomModuleSpec   `json:"spec,omitempty"`
	Status SecuritycenterSecurityHealthAnalyticsCustomModuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecuritycenterSecurityHealthAnalyticsCustomModuleList contains a list of SecuritycenterSecurityHealthAnalyticsCustomModule
type SecuritycenterSecurityHealthAnalyticsCustomModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecuritycenterSecurityHealthAnalyticsCustomModule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecuritycenterSecurityHealthAnalyticsCustomModule{}, &SecuritycenterSecurityHealthAnalyticsCustomModuleList{})
}
