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

// API sources for AIPlatformSession, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/aiplatform/v1/session.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.reasoningEngines.sessions

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformSessionGVK = GroupVersion.WithKind("AIPlatformSession")

// AIPlatformSessionSpec defines the desired state of AIPlatformSession
// +kcc:spec:proto=google.cloud.aiplatform.v1.Session
// +kcc:required-from-proto
type AIPlatformSessionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/reasoningEngines/{reasoning_engine} this resource belongs to.
	// +kcc:guess
	// ReasoningEngineRef *AIPlatformReasoningEngineRef `json:"reasoningEngineRef,omitempty"`

	// The AIPlatformSession name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Timestamp of when this session is considered expired.
	//  This is *always* provided on output, regardless of what was sent
	//  on input.
	//  The minimum value is 24 hours from the time of creation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Session.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Optional. Input only. The TTL for this session.
	//  The minimum value is 24 hours.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Session.ttl
	TTL *string `json:"ttl,omitempty"`

	// Optional. The display name of the session.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Session.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The labels with user-defined metadata to organize your Sessions.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Session.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Session specific memory which stores key conversation points.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Session.session_state
	SessionState apiextensionsv1.JSON `json:"sessionState,omitempty"`

	// Required. Immutable. String id provided by the user
	// +kcc:proto:field=google.cloud.aiplatform.v1.Session.user_id
	// +required
	UserID *string `json:"userID,omitempty"`
}

// AIPlatformSessionStatus defines the config connector machine state of AIPlatformSession
type AIPlatformSessionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformSession resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformSessionObservedState `json:"observedState,omitempty"`
}

// AIPlatformSessionObservedState is the state of the AIPlatformSession resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.Session
type AIPlatformSessionObservedState struct {
	// Output only. Timestamp when the session was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Session.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the session was updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Session.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformsession;gcpaiplatformsessions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformSession is the Schema for the AIPlatformSession API
// +k8s:openapi-gen=true
type AIPlatformSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformSessionSpec   `json:"spec,omitempty"`
	Status AIPlatformSessionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformSessionList contains a list of AIPlatformSession
type AIPlatformSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformSession `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformSession{}, &AIPlatformSessionList{})
}
