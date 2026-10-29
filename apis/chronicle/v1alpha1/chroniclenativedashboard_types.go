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

// API sources for ChronicleNativeDashboard, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/chronicle/v1/native_dashboard.proto
// +kcc:source:service-docs=https://cloud.google.com/chronicle/docs/secops/secops-overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/chronicle/docs/reference/rest/v1/projects.locations.instances.nativeDashboards

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ChronicleNativeDashboardGVK = GroupVersion.WithKind("ChronicleNativeDashboard")

// ChronicleNativeDashboardSpec defines the desired state of ChronicleNativeDashboard
// +kcc:spec:proto=google.cloud.chronicle.v1.NativeDashboard
// +kcc:required-from-proto
type ChronicleNativeDashboardSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ChronicleNativeDashboard name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Dashboard display name/title visible to users.
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the dashboard.
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.description
	Description *string `json:"description,omitempty"`

	// Optional. Definition of the dashboard like metadata, visualization and
	//  datasource configuration etc.
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.definition
	Definition *DashboardDefinition `json:"definition,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.etag
	Etag *string `json:"etag,omitempty"`
}

// ChronicleNativeDashboardStatus defines the config connector machine state of ChronicleNativeDashboard
type ChronicleNativeDashboardStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ChronicleNativeDashboard resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ChronicleNativeDashboardObservedState `json:"observedState,omitempty"`
}

// ChronicleNativeDashboardObservedState is the state of the ChronicleNativeDashboard resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.chronicle.v1.NativeDashboard
type ChronicleNativeDashboardObservedState struct {
	// Optional. Definition of the dashboard like metadata, visualization and
	//  datasource configuration etc.
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.definition
	Definition *DashboardDefinitionObservedState `json:"definition,omitempty"`

	// Output only. Whether it's an out of the box or custom created dashboard.
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.type
	Type *string `json:"type,omitempty"`

	// Output only. Creation time of dashboard.
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time the dashboard was last edited.
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. User who created the dashboard.
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.create_user_id
	CreateUserID *string `json:"createUserID,omitempty"`

	// Output only. User who last edited the dashboard.
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.update_user_id
	UpdateUserID *string `json:"updateUserID,omitempty"`

	// Output only. User Preferences for a dashboard
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.dashboard_user_data
	DashboardUserData *DashboardUserData `json:"dashboardUserData,omitempty"`

	// Output only. Access of the dashboard
	// +kcc:proto:field=google.cloud.chronicle.v1.NativeDashboard.access
	Access *string `json:"access,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpchroniclenativedashboard;gcpchroniclenativedashboards
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ChronicleNativeDashboard is the Schema for the ChronicleNativeDashboard API
// +k8s:openapi-gen=true
type ChronicleNativeDashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ChronicleNativeDashboardSpec   `json:"spec,omitempty"`
	Status ChronicleNativeDashboardStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ChronicleNativeDashboardList contains a list of ChronicleNativeDashboard
type ChronicleNativeDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronicleNativeDashboard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronicleNativeDashboard{}, &ChronicleNativeDashboardList{})
}
