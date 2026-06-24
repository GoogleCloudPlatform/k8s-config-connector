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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowSecuritySettingsGVK = GroupVersion.WithKind("DialogflowSecuritySettings")

type Parent struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// +required
	Location *string `json:"location"`
}

// DialogflowSecuritySettingsSpec defines the desired state of DialogflowSecuritySettings
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.SecuritySettings
type DialogflowSecuritySettingsSpec struct {
	Parent `json:",inline"`

	// Required. The human-readable name of the security settings, unique within
	// the location.
	// +required
	DisplayName *string `json:"displayName"`

	// Strategy that defines how we do redaction.
	// +optional
	RedactionStrategy *string `json:"redactionStrategy,omitempty"`

	// Defines the data for which Dialogflow applies redaction. Dialogflow does
	// not redact data that it does not have access to – for example, Cloud
	// logging.
	// +optional
	RedactionScope *string `json:"redactionScope,omitempty"`

	// [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this
	// template to define inspect base settings.
	// +optional
	InspectTemplate *string `json:"inspectTemplate,omitempty"`

	// [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this
	// template to define de-identification configuration for the content.
	// +optional
	DeidentifyTemplate *string `json:"deidentifyTemplate,omitempty"`

	// Retains the data for the specified number of days.
	// +optional
	RetentionWindowDays *int32 `json:"retentionWindowDays,omitempty"`

	// Specifies the retention behavior defined by SecuritySettings.RetentionStrategy.
	// +optional
	RetentionStrategy *string `json:"retentionStrategy,omitempty"`

	// List of types of data to remove when retention settings triggers purge.
	// +optional
	PurgeDataTypes []string `json:"purgeDataTypes,omitempty"`

	// Controls audio export settings for post-conversation analytics.
	// +optional
	AudioExportSettings *SecuritySettings_AudioExportSettings `json:"audioExportSettings,omitempty"`

	// Controls conversation exporting settings to Insights after conversation is completed.
	// +optional
	InsightsExportSettings *SecuritySettings_InsightsExportSettings `json:"insightsExportSettings,omitempty"`

	// The DialogflowSecuritySettings name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// DialogflowSecuritySettingsStatus defines the config connector machine state of DialogflowSecuritySettings
type DialogflowSecuritySettingsStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowSecuritySettings resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowSecuritySettingsObservedState `json:"observedState,omitempty"`
}

// DialogflowSecuritySettingsObservedState is the state of the DialogflowSecuritySettings resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.SecuritySettings
type DialogflowSecuritySettingsObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowsecuritysetting;gcpdialogflowsecuritysettings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowSecuritySettings is the Schema for the DialogflowSecuritySettings API
// +k8s:openapi-gen=true
type DialogflowSecuritySettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowSecuritySettingsSpec   `json:"spec,omitempty"`
	Status DialogflowSecuritySettingsStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowSecuritySettingsList contains a list of DialogflowSecuritySettings
type DialogflowSecuritySettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowSecuritySettings `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowSecuritySettings{}, &DialogflowSecuritySettingsList{})
}
