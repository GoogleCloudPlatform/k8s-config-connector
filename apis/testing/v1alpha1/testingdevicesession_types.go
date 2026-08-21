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

var TestingDeviceSessionGVK = GroupVersion.WithKind("TestingDeviceSession")

// TestingDeviceSessionSpec defines the desired state of TestingDeviceSession
// +kcc:spec:proto=google.devtools.testing.v1.DeviceSession
type TestingDeviceSessionSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The TestingDeviceSession name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The amount of time that a device will be initially allocated
	//  for. This can eventually be extended with the UpdateDeviceSession RPC.
	//  Default: 15 minutes.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.ttl
	TTL *string `json:"ttl,omitempty"`

	// Optional. If the device is still in use at this time, any connections
	//  will be ended and the SessionState will transition from ACTIVE to
	//  FINISHED.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Required. The requested device
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.android_device
	AndroidDevice *AndroidDevice `json:"androidDevice"`
}

// TestingDeviceSessionStatus defines the config connector machine state of TestingDeviceSession
type TestingDeviceSessionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TestingDeviceSession resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TestingDeviceSessionObservedState `json:"observedState,omitempty"`
}

// TestingDeviceSessionObservedState is the state of the TestingDeviceSession resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devtools.testing.v1.DeviceSession
type TestingDeviceSessionObservedState struct {
	// Output only. The title of the DeviceSession to be presented in the UI.
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Current state of the DeviceSession.
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.state
	State *string `json:"state,omitempty"`

	// Output only. The historical state transitions of the session_state message
	//  including the current session state.
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.state_histories
	StateHistories []DeviceSession_SessionStateEventObservedState `json:"stateHistories,omitempty"`

	// Output only. The interval of time that this device must be interacted with
	//  before it transitions from ACTIVE to TIMEOUT_INACTIVITY.
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.inactivity_timeout
	InactivityTimeout *string `json:"inactivityTimeout,omitempty"`

	// Output only. The time that the Session was created.
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp that the session first became ACTIVE.
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.active_start_time
	ActiveStartTime *string `json:"activeStartTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptestingdevicesession;gcptestingdevicesessions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TestingDeviceSession is the Schema for the TestingDeviceSession API
// +k8s:openapi-gen=true
type TestingDeviceSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TestingDeviceSessionSpec   `json:"spec,omitempty"`
	Status TestingDeviceSessionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TestingDeviceSessionList contains a list of TestingDeviceSession
type TestingDeviceSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestingDeviceSession `json:"items"`
}

// +kcc:proto=google.devtools.testing.v1.AndroidDevice
type AndroidDevice struct {
	// Required. The id of the Android device to be used.
	//  Use the TestEnvironmentDiscoveryService to get supported options.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.testing.v1.AndroidDevice.android_model_id
	AndroidModelID *string `json:"androidModelID"`

	// Required. The id of the Android OS version to be used.
	//  Use the TestEnvironmentDiscoveryService to get supported options.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.testing.v1.AndroidDevice.android_version_id
	AndroidVersionID *string `json:"androidVersionID"`

	// Required. The locale the test device used for testing.
	//  Use the TestEnvironmentDiscoveryService to get supported options.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.testing.v1.AndroidDevice.locale
	Locale *string `json:"locale"`

	// Required. How the device is oriented during the test.
	//  Use the TestEnvironmentDiscoveryService to get supported options.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.testing.v1.AndroidDevice.orientation
	Orientation *string `json:"orientation"`
}

// +kcc:observedstate:proto=google.devtools.testing.v1.DeviceSession.SessionStateEvent
type DeviceSession_SessionStateEventObservedState struct {
	// Output only. The session_state tracked by this event
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.SessionStateEvent.session_state
	SessionState *string `json:"sessionState,omitempty"`

	// Output only. The time that the session_state first encountered that
	//  state.
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.SessionStateEvent.event_time
	EventTime *string `json:"eventTime,omitempty"`

	// Output only. A human-readable message to explain the state.
	// +kcc:proto:field=google.devtools.testing.v1.DeviceSession.SessionStateEvent.state_message
	StateMessage *string `json:"stateMessage,omitempty"`
}

func init() {
	SchemeBuilder.Register(&TestingDeviceSession{}, &TestingDeviceSessionList{})
}
