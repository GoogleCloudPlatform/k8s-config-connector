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

var DataprocSessionGVK = GroupVersion.WithKind("DataprocSession")

// +kcc:proto=google.cloud.dataproc.v1.JupyterConfig
type JupyterConfig struct {
	// Optional. Kernel
	// +kcc:proto:field=google.cloud.dataproc.v1.JupyterConfig.kernel
	Kernel *string `json:"kernel,omitempty"`

	// Optional. Display name, shown in the Jupyter kernelspec card.
	// +kcc:proto:field=google.cloud.dataproc.v1.JupyterConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SparkConnectConfig
type SparkConnectConfig struct {
}

// DataprocSessionSpec defines the desired state of DataprocSession
// +kcc:spec:proto=google.cloud.dataproc.v1.Session
type DataprocSessionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// Optional. Jupyter session config.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.jupyter_session
	JupyterSession *JupyterConfig `json:"jupyterSession,omitempty"`

	// Optional. Spark Connect session config.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.spark_connect_session
	SparkConnectSession *SparkConnectConfig `json:"sparkConnectSession,omitempty"`

	// Optional. The labels to associate with the session.
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Label **values** may be empty, but, if present, must contain 1 to 63
	//  characters, and must conform to [RFC
	//  1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	//  associated with a session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Runtime configuration for the session execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.runtime_config
	RuntimeConfig *RuntimeConfig `json:"runtimeConfig,omitempty"`

	// Optional. Environment configuration for the session execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.environment_config
	EnvironmentConfig *EnvironmentConfig `json:"environmentConfig,omitempty"`

	// Optional. The email address of the user who owns the session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.user
	User *string `json:"user,omitempty"`

	// Optional. The session template used by the session.
	//
	//  Only resource names, including project ID and location, are valid.
	//
	//  Example:
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/locations/[dataproc_region]/sessionTemplates/[template_id]`
	//  * `projects/[project_id]/locations/[dataproc_region]/sessionTemplates/[template_id]`
	//
	//  The template must be in the same project and Dataproc region as the
	//  session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.session_template
	SessionTemplate *string `json:"sessionTemplate,omitempty"`

	// The DataprocSession name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// DataprocSessionStatus defines the config connector machine state of DataprocSession
type DataprocSessionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataprocSession resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataprocSessionObservedState `json:"observedState,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.Session.SessionStateHistory
type Session_SessionStateHistoryObservedState struct {
	// Output only. The state of the session at this point in the session
	//  history.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.SessionStateHistory.state
	State *string `json:"state,omitempty"`

	// Output only. Details about the state at this point in the session
	//  history.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.SessionStateHistory.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. The time when the session entered the historical state.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.SessionStateHistory.state_start_time
	StateStartTime *string `json:"stateStartTime,omitempty"`
}

// DataprocSessionObservedState is the state of the DataprocSession resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataproc.v1.Session
type DataprocSessionObservedState struct {
	// Output only. A session UUID (Unique Universal Identifier). The service
	//  generates this value when it creates the session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.uuid
	Uuid *string `json:"uuid,omitempty"`

	// Output only. The time when the session was created.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Runtime information about session execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.runtime_info
	RuntimeInfo *RuntimeInfoObservedState `json:"runtimeInfo,omitempty"`

	// Output only. A state of the session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.state
	State *string `json:"state,omitempty"`

	// Output only. Session state details, such as the failure
	//  description if the state is `FAILED`.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. The time when the session entered the current state.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.state_time
	StateTime *string `json:"stateTime,omitempty"`

	// Output only. The email address of the user who created the session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. Historical state information for the session.
	// +kcc:proto:field=google.cloud.dataproc.v1.Session.state_history
	StateHistory []Session_SessionStateHistoryObservedState `json:"stateHistory,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataprocsession;gcpdataprocsessions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataprocSession is the Schema for the DataprocSession API
// +k8s:openapi-gen=true
type DataprocSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataprocSessionSpec   `json:"spec,omitempty"`
	Status DataprocSessionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataprocSessionList contains a list of DataprocSession
type DataprocSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocSession `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocSession{}, &DataprocSessionList{})
}
