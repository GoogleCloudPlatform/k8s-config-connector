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

// API sources for SecureSourceManagerHook, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/securesourcemanager/v1/secure_source_manager.proto
// +kcc:source:service-docs=https://cloud.google.com/secure-source-manager
// +kcc:source:resource-docs=https://docs.cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories.hooks

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecureSourceManagerHookGVK = GroupVersion.WithKind("SecureSourceManagerHook")

// SecureSourceManagerHookSpec defines the desired state of SecureSourceManagerHook
// +kcc:spec:proto=google.cloud.securesourcemanager.v1.Hook
// +kcc:required-from-proto
type SecureSourceManagerHookSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The SecureSourceManagerHook name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The target URI to which the payloads will be delivered.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Hook.target_uri
	// +required
	TargetURI *string `json:"targetURI,omitempty"`

	// Optional. Determines if the hook disabled or not.
	//  Set to true to stop sending traffic.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Hook.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Optional. The events that trigger hook on.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Hook.events
	Events []string `json:"events,omitempty"`

	// Optional. The trigger option for push events.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Hook.push_option
	PushOption *Hook_PushOption `json:"pushOption,omitempty"`

	// Optional. The sensitive query string to be appended to the target URI.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Hook.sensitive_query_string
	SensitiveQueryString *string `json:"sensitiveQueryString,omitempty"`
}

// SecureSourceManagerHookStatus defines the config connector machine state of SecureSourceManagerHook
type SecureSourceManagerHookStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecureSourceManagerHook resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecureSourceManagerHookObservedState `json:"observedState,omitempty"`
}

// SecureSourceManagerHookObservedState is the state of the SecureSourceManagerHook resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.securesourcemanager.v1.Hook
type SecureSourceManagerHookObservedState struct {
	// Output only. Create timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Hook.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Hook.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Unique identifier of the hook.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Hook.uid
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuresourcemanagerhook;gcpsecuresourcemanagerhooks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecureSourceManagerHook is the Schema for the SecureSourceManagerHook API
// +k8s:openapi-gen=true
type SecureSourceManagerHook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecureSourceManagerHookSpec   `json:"spec,omitempty"`
	Status SecureSourceManagerHookStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecureSourceManagerHookList contains a list of SecureSourceManagerHook
type SecureSourceManagerHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecureSourceManagerHook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecureSourceManagerHook{}, &SecureSourceManagerHookList{})
}
