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

var DataprocSessionTemplateGVK = GroupVersion.WithKind("DataprocSessionTemplate")

// +kcc:proto=google.cloud.dataproc.v1.SparkConnectConfig
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:Schemaless
type SparkConnectConfig struct {
}

// DataprocSessionTemplateSpec defines the desired state of DataprocSessionTemplate
// +kcc:spec:proto=google.cloud.dataproc.v1.SessionTemplate
type DataprocSessionTemplateSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// Optional. Brief description of the template.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.description
	Description *string `json:"description,omitempty"`

	// Optional. Jupyter session config.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.jupyter_session
	JupyterSession *JupyterConfig `json:"jupyterSession,omitempty"`

	// Optional. Spark Connect session config.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.spark_connect_session
	SparkConnectSession *SparkConnectConfig `json:"sparkConnectSession,omitempty"`

	// Optional. Labels to associate with sessions created using this template.
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Label **values** can be empty, but, if present, must contain 1 to 63
	//  characters and conform to [RFC
	//  1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	//  associated with a session.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Runtime configuration for session execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.runtime_config
	RuntimeConfig *RuntimeConfig `json:"runtimeConfig,omitempty"`

	// Optional. Environment configuration for session execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.environment_config
	EnvironmentConfig *EnvironmentConfig `json:"environmentConfig,omitempty"`

	// The DataprocSessionTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// DataprocSessionTemplateStatus defines the config connector machine state of DataprocSessionTemplate
type DataprocSessionTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataprocSessionTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataprocSessionTemplateObservedState `json:"observedState,omitempty"`
}

// DataprocSessionTemplateObservedState is the state of the DataprocSessionTemplate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataproc.v1.SessionTemplate
type DataprocSessionTemplateObservedState struct {
	// Output only. The time when the template was created.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The email address of the user who created the template.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. The time the template was last updated.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A session template UUID (Unique Universal Identifier). The
	//  service generates this value when it creates the session template.
	// +kcc:proto:field=google.cloud.dataproc.v1.SessionTemplate.uuid
	Uuid *string `json:"uuid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataprocsessiontemplate;gcpdataprocsessiontemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataprocSessionTemplate is the Schema for the DataprocSessionTemplate API
// +k8s:openapi-gen=true
type DataprocSessionTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataprocSessionTemplateSpec   `json:"spec,omitempty"`
	Status DataprocSessionTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataprocSessionTemplateList contains a list of DataprocSessionTemplate
type DataprocSessionTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocSessionTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocSessionTemplate{}, &DataprocSessionTemplateList{})
}
