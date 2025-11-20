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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MonitoringAlertPolicyGVK = GroupVersion.WithKind("MonitoringAlertPolicy")

// MonitoringAlertPolicySpec defines the desired state of MonitoringAlertPolicy
// +kcc:spec:proto=google.monitoring.v3.AlertPolicy
type MonitoringAlertPolicySpec struct {
	// // Required. Defines the parent path of the resource.
	// *parent.ProjectAndLocationRef `json:",inline"`

	// The MonitoringAlertPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// A short name or phrase used to identify the policy in dashboards,
	//  notifications, and incidents. To avoid confusion, don't use the same
	//  display name for multiple policies in the same project. The name is
	//  limited to 512 Unicode characters.
	//
	//  The convention for the display_name of a PrometheusQueryLanguageCondition
	//  is "{rule group name}/{alert name}", where the {rule group name} and
	//  {alert name} should be taken from the corresponding Prometheus
	//  configuration file. This convention is not enforced.
	//  In any case the display_name is not a unique key of the AlertPolicy.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Documentation that is included with notifications and incidents related to
	//  this policy. Best practice is for the documentation to include information
	//  to help responders understand, mitigate, escalate, and correct the
	//  underlying problems detected by the alerting policy. Notification channels
	//  that have limited capacity might not show this documentation.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.documentation
	Documentation *AlertPolicy_Documentation `json:"documentation,omitempty"`

	// // User-supplied key/value data to be used for organizing and
	// //  identifying the `AlertPolicy` objects.
	// //
	// //  The field can contain up to 64 entries. Each key and value is limited to
	// //  63 Unicode characters or 128 bytes, whichever is smaller. Labels and
	// //  values can contain only lowercase letters, numerals, underscores, and
	// //  dashes. Keys must begin with a letter.
	// //
	// //  Note that Prometheus {alert name} is a
	// //  [valid Prometheus label
	// //  names](https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels),
	// //  whereas Prometheus {rule group} is an unrestricted UTF-8 string.
	// //  This means that they cannot be stored as-is in user labels, because
	// //  they may contain characters that are not allowed in user-label values.
	// // +kcc:proto:field=google.monitoring.v3.AlertPolicy.user_labels
	// UserLabels map[string]string `json:"userLabels,omitempty"`

	// A list of conditions for the policy. The conditions are combined by AND or
	//  OR according to the `combiner` field. If the combined conditions evaluate
	//  to true, then an incident is created. A policy can have from one to six
	//  conditions.
	//  If `condition_time_series_query_language` is present, it must be the only
	//  `condition`.
	//  If `condition_monitoring_query_language` is present, it must be the only
	//  `condition`.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.conditions
	Conditions []AlertPolicy_Condition `json:"conditions,omitempty"`

	// How to combine the results of multiple conditions to determine if an
	//  incident should be opened.
	//  If `condition_time_series_query_language` is present, this must be
	//  `COMBINE_UNSPECIFIED`.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.combiner
	Combiner *string `json:"combiner,omitempty"`

	// Whether or not the policy is enabled. On write, the default interpretation
	//  if unset is that the policy is enabled. On read, clients should not make
	//  any assumption about the state if it has not been populated. The
	//  field should always be populated on List and Get operations, unless
	//  a field projection has been specified that strips it out.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// // Read-only description of how the alerting policy is invalid. This field is
	// //  only set when the alerting policy is invalid. An invalid alerting policy
	// //  will not generate incidents.
	// // +kcc:proto:field=google.monitoring.v3.AlertPolicy.validity
	// Validity *Status `json:"validity,omitempty"`

	// Identifies the notification channels to which notifications should be sent
	//  when incidents are opened or closed or when new violations occur on
	//  an already opened incident. Each element of this array corresponds to
	//  the `name` field in each of the
	//  [`NotificationChannel`][google.monitoring.v3.NotificationChannel]
	//  objects that are returned from the [`ListNotificationChannels`]
	//  [google.monitoring.v3.NotificationChannelService.ListNotificationChannels]
	//  method. The format of the entries in this field is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID]
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.notification_channels
	NotificationChannels []string `json:"notificationChannels,omitempty"`

	// Control over how this alerting policy's notification channels are notified.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.alert_strategy
	AlertStrategy *AlertPolicy_AlertStrategy `json:"alertStrategy,omitempty"`

	// // Optional. The severity of an alerting policy indicates how important
	// //  incidents generated by that policy are. The severity level will be
	// //  displayed on the Incident detail page and in notifications.
	// // +kcc:proto:field=google.monitoring.v3.AlertPolicy.severity
	// Severity *string `json:"severity,omitempty"`

}

// MonitoringAlertPolicyStatus defines the config connector machine state of MonitoringAlertPolicy
type MonitoringAlertPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MonitoringAlertPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// NOTYET(TERRAFORM)
	// // ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *MonitoringAlertPolicyObservedState `json:"observedState,omitempty"`

	// A read-only record of the creation of the alerting policy. If provided
	//  in a call to create or update, this field will be ignored.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.creation_record
	CreationRecord *MutationRecord `json:"creationRecord,omitempty"`

	// // A read-only record of the most recent change to the alerting policy. If
	// //  provided in a call to create or update, this field will be ignored.
	// // +kcc:proto:field=google.monitoring.v3.AlertPolicy.mutation_record
	// MutationRecord *MutationRecord `json:"mutationRecord,omitempty"`

	// Identifier. Required if the policy exists. The resource name for this
	//  policy. The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID]
	//
	//  `[ALERT_POLICY_ID]` is assigned by Cloud Monitoring when the policy
	//  is created. When calling the
	//  [alertPolicies.create][google.monitoring.v3.AlertPolicyService.CreateAlertPolicy]
	//  method, do not include the `name` field in the alerting policy passed as
	//  part of the request.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.name
	Name *string `json:"name,omitempty"`
}

// MonitoringAlertPolicyObservedState is the state of the MonitoringAlertPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.monitoring.v3.AlertPolicy
type MonitoringAlertPolicyObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringalertpolicy;gcpmonitoringalertpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringAlertPolicy is the Schema for the MonitoringAlertPolicy API
// +k8s:openapi-gen=true
type MonitoringAlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MonitoringAlertPolicySpec   `json:"spec,omitempty"`
	Status MonitoringAlertPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MonitoringAlertPolicyList contains a list of MonitoringAlertPolicy
type MonitoringAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringAlertPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringAlertPolicy{}, &MonitoringAlertPolicyList{})
}
