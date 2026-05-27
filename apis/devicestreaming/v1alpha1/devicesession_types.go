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

var DeviceStreamingSessionGVK = GroupVersion.WithKind("DeviceStreamingSession")

// DeviceStreamingSessionSpec defines the desired state of DeviceStreamingSession
// +kcc:spec:proto=google.cloud.devicestreaming.v1.DeviceSession
type DeviceStreamingSessionSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The DeviceStreamingSession name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The amount of time that a device will be initially allocated
	// for. This can eventually be extended with the UpdateDeviceSession RPC.
	// Default: 15 minutes.
	// +kubebuilder:validation:Optional
	TTL *string `json:"ttl,omitempty"`

	// Optional. If the device is still in use at this time, any connections
	// will be ended and the SessionState will transition from ACTIVE to
	// FINISHED.
	// +kubebuilder:validation:Optional
	ExpireTime *string `json:"expireTime,omitempty"`

	// Required. The requested device
	// +kubebuilder:validation:Required
	AndroidDevice *AndroidDevice `json:"androidDevice"`
}

// DeviceStreamingSessionStatus defines the config connector machine state of DeviceStreamingSession
type DeviceStreamingSessionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DeviceStreamingSession resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DeviceStreamingSessionObservedState `json:"observedState,omitempty"`
}

// DeviceStreamingSessionObservedState is the state of the DeviceStreamingSession resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.devicestreaming.v1.DeviceSession
type DeviceStreamingSessionObservedState struct {
	// Output only. The title of the DeviceSession to be presented in the UI.
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Current state of the DeviceSession.
	State *string `json:"state,omitempty"`

	// Output only. The historical state transitions of the session_state message
	// including the current session state.
	StateHistories []DeviceSession_SessionStateEvent `json:"stateHistories,omitempty"`

	// Output only. The interval of time that this device must be interacted with
	// before it transitions from ACTIVE to TIMEOUT_INACTIVITY.
	InactivityTimeout *string `json:"inactivityTimeout,omitempty"`

	// Output only. The time that the Session was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp that the session first became ACTIVE.
	ActiveStartTime *string `json:"activeStartTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdevicestreamingsession;gcpdevicestreamingsessions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DeviceStreamingSession is the Schema for the DeviceStreamingSession API
// +k8s:openapi-gen=true
type DeviceStreamingSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DeviceStreamingSessionSpec   `json:"spec,omitempty"`
	Status DeviceStreamingSessionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DeviceStreamingSessionList contains a list of DeviceStreamingSession
type DeviceStreamingSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeviceStreamingSession `json:"items"`
}

// +kcc:proto=google.cloud.devicestreaming.v1.AndroidDevice
type AndroidDevice struct {
	// Required. The id of the Android device to be used.
	// Use the TestEnvironmentDiscoveryService to get supported options.
	// +kubebuilder:validation:Required
	AndroidModelID *string `json:"androidModelID"`

	// Required. The id of the Android OS version to be used.
	// Use the TestEnvironmentDiscoveryService to get supported options.
	// +kubebuilder:validation:Required
	AndroidVersionID *string `json:"androidVersionID"`

	// Optional. The locale the test device used for testing.
	// Use the TestEnvironmentDiscoveryService to get supported options.
	// +kubebuilder:validation:Optional
	Locale *string `json:"locale,omitempty"`

	// Optional. How the device is oriented during the test.
	// Use the TestEnvironmentDiscoveryService to get supported options.
	// +kubebuilder:validation:Optional
	Orientation *string `json:"orientation,omitempty"`
}

// +kcc:proto=google.cloud.devicestreaming.v1.DeviceSession.SessionStateEvent
type DeviceSession_SessionStateEvent struct {
	// Output only. The session_state tracked by this event
	SessionState *string `json:"sessionState,omitempty"`

	// Output only. The time that the session_state first encountered that
	// state.
	EventTime *string `json:"eventTime,omitempty"`

	// Output only. A human-readable message to explain the state.
	StateMessage *string `json:"stateMessage,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DeviceStreamingSession{}, &DeviceStreamingSessionList{})
}
