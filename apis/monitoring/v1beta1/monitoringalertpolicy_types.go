// Copyright 2024 Google LLC
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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	MonitoringAlertPolicyGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "MonitoringAlertPolicy",
	}
)

// +kcc:proto=google.monitoring.v3.AlertPolicy
type MonitoringAlertPolicySpec struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
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
	DisplayName *string `json:"displayName,omitempty"`

	// // Required if the policy exists. The resource name for this policy. The
	// //  format is:
	// //
	// //      projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID]
	// //
	// //  `[ALERT_POLICY_ID]` is assigned by Cloud Monitoring when the policy
	// //  is created. When calling the
	// //  [alertPolicies.create][google.monitoring.v3.AlertPolicyService.CreateAlertPolicy]
	// //  method, do not include the `name` field in the alerting policy passed as
	// //  part of the request.
	// Name *string `json:"name,omitempty"`

	// Documentation that is included with notifications and incidents related to
	//  this policy. Best practice is for the documentation to include information
	//  to help responders understand, mitigate, escalate, and correct the
	//  underlying problems detected by the alerting policy. Notification channels
	//  that have limited capacity might not show this documentation.
	Documentation *AlertPolicy_Documentation `json:"documentation,omitempty"`

	// User-supplied key/value data to be used for organizing and
	//  identifying the `AlertPolicy` objects.
	//
	//  The field can contain up to 64 entries. Each key and value is limited to
	//  63 Unicode characters or 128 bytes, whichever is smaller. Labels and
	//  values can contain only lowercase letters, numerals, underscores, and
	//  dashes. Keys must begin with a letter.
	//
	//  Note that Prometheus {alert name} is a
	//  [valid Prometheus label
	//  names](https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels),
	//  whereas Prometheus {rule group} is an unrestricted UTF-8 string.
	//  This means that they cannot be stored as-is in user labels, because
	//  they may contain characters that are not allowed in user-label values.
	UserLabels map[string]string `json:"userLabels,omitempty"`

	// A list of conditions for the policy. The conditions are combined by AND or
	//  OR according to the `combiner` field. If the combined conditions evaluate
	//  to true, then an incident is created. A policy can have from one to six
	//  conditions.
	//  If `condition_time_series_query_language` is present, it must be the only
	//  `condition`.
	//  If `condition_monitoring_query_language` is present, it must be the only
	//  `condition`.
	Conditions []*AlertPolicy_Condition `json:"conditions,omitempty"`

	// How to combine the results of multiple conditions to determine if an
	//  incident should be opened.
	//  If `condition_time_series_query_language` is present, this must be
	//  `COMBINE_UNSPECIFIED`.
	Combiner *string `json:"combiner,omitempty"`

	// Whether or not the policy is enabled. On write, the default interpretation
	//  if unset is that the policy is enabled. On read, clients should not make
	//  any assumption about the state if it has not been populated. The
	//  field should always be populated on List and Get operations, unless
	//  a field projection has been specified that strips it out.
	Enabled *bool `json:"enabled,omitempty"`

	/*NOTYET
	// Read-only description of how the alert policy is invalid. This field is
	//  only set when the alert policy is invalid. An invalid alert policy will not
	//  generate incidents.
	Validity *Status `json:"validity,omitempty"`
	*/

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
	NotificationChannels []string `json:"notificationChannels,omitempty"`

	// Control over how this alert policy's notification channels are notified.
	AlertStrategy *AlertPolicy_AlertStrategy `json:"alertStrategy,omitempty"`

	// Optional. The severity of an alert policy indicates how important incidents
	//  generated by that policy are. The severity level will be displayed on the
	//  Incident detail page and in notifications.
	Severity *string `json:"severity,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy
type MonitoringAlertPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   MonitoringAlertPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A read-only record of the creation of the alerting policy. If provided
	//  in a call to create or update, this field will be ignored.
	CreationRecord *MutationRecord `json:"creationRecord,omitempty"`

	// A read-only record of the most recent change to the alerting policy. If
	//  provided in a call to create or update, this field will be ignored.
	MutationRecord *MutationRecord `json:"mutationRecord,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringalertpolicy;gcpmonitoringalertpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringAlertPolicy is the Schema for the monitoring API
// +k8s:openapi-gen=true
type MonitoringAlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec MonitoringAlertPolicySpec `json:"spec,omitempty"`

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
