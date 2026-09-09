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

package v1beta1

import (
	monitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ErrorReportingNotificationSettingsGVK = GroupVersion.WithKind("ErrorReportingNotificationSettings")

// ErrorReportingNotificationSettingsSpec defines the desired state of ErrorReportingNotificationSettings
// +kcc:spec:proto=google.devtools.clouderrorreporting.v1beta1.NotificationSettings
type ErrorReportingNotificationSettingsSpec struct {
	/* The project that this resource belongs to. */
	ProjectRef refsv1beta1.ProjectRef `json:"projectRef"`

	/* The notification channels to send error reports to. */
	// +kcc:proto:field=google.devtools.clouderrorreporting.v1beta1.NotificationSettings.notification_channels
	NotificationChannels []monitoringv1beta1.MonitoringNotificationChannelRef `json:"notificationChannels,omitempty"`

	/* Optional. The delay periods for version skew reports. */
	// +kcc:proto:field=google.devtools.clouderrorreporting.v1beta1.NotificationSettings.version_skew_report_delays
	VersionSkewReportDelays []string `json:"versionSkewReportDelays,omitempty"`
}

// ErrorReportingNotificationSettingsStatus defines the observed state of ErrorReportingNotificationSettings
type ErrorReportingNotificationSettingsStatus struct {
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to the metadata's generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A list of the condition Taints that this resource has. */
	// +optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* Observed state of the resource. */
	// +optional
	ObservedState *ErrorReportingNotificationSettingsObservedState `json:"observedState,omitempty"`
}

// ErrorReportingNotificationSettingsObservedState defines the observed state of ErrorReportingNotificationSettings
// +kcc:observedstate:proto=google.devtools.clouderrorreporting.v1beta1.NotificationSettings
type ErrorReportingNotificationSettingsObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpernotificationsettings;gcpernotificationsetting
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].message",type="string"

// ErrorReportingNotificationSettings is the Schema for the ErrorReportingNotificationSettings API
type ErrorReportingNotificationSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ErrorReportingNotificationSettingsSpec   `json:"spec,omitempty"`
	Status ErrorReportingNotificationSettingsStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ErrorReportingNotificationSettingsList contains a list of ErrorReportingNotificationSettings
type ErrorReportingNotificationSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ErrorReportingNotificationSettings `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ErrorReportingNotificationSettings{}, &ErrorReportingNotificationSettingsList{})
}
