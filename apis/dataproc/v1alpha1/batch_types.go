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

package v1alpha1

import (
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataprocBatchGVK = GroupVersion.WithKind("DataprocBatch")

// DataprocBatchSpec defines the desired state of DataprocBatch
// +kcc:proto=google.cloud.dataproc.v1.Batch
type DataprocBatchSpec struct {
	// Optional. PySpark batch config.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.pyspark_batch
	PysparkBatch *PySparkBatch `json:"pysparkBatch,omitempty"`

	// Optional. Spark batch config.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.spark_batch
	SparkBatch *SparkBatch `json:"sparkBatch,omitempty"`

	// Optional. SparkR batch config.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.spark_r_batch
	SparkRBatch *SparkRBatch `json:"sparkRBatch,omitempty"`

	// Optional. SparkSql batch config.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.spark_sql_batch
	SparkSQLBatch *SparkSQLBatch `json:"sparkSQLBatch,omitempty"`

	// Optional. The labels to associate with this batch.
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Label **values** may be empty, but, if present, must contain 1 to 63
	//  characters, and must conform to [RFC
	//  1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	//  associated with a batch.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Runtime configuration for the batch execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.runtime_config
	RuntimeConfig *RuntimeConfig `json:"runtimeConfig,omitempty"`

	// Optional. Environment configuration for the batch execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.environment_config
	EnvironmentConfig *EnvironmentConfig `json:"environmentConfig,omitempty"`

	// Required.
	Location string `json:"location,omitempty"`

	commonv1alpha1.CommonSpec `json:",inline"`
}

// DataprocBatchStatus defines the config connector machine state of DataprocBatch
type DataprocBatchStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataprocBatch resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataprocBatchObservedState `json:"observedState,omitempty"`
}

// DataprocBatchObservedState is the state of the DataprocBatch resource as most recently observed in GCP.
// +kcc:proto=google.cloud.dataproc.v1.Batch
type DataprocBatchObservedState struct {
	// Output only. A batch UUID (Unique Universal Identifier). The service
	//  generates this value when it creates the batch.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.uuid
	Uuid *string `json:"uuid,omitempty"`

	// Output only. The time when the batch was created.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Runtime information about batch execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.runtime_info
	RuntimeInfo *RuntimeInfo `json:"runtimeInfo,omitempty"`

	// Output only. The state of the batch.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.state
	State *string `json:"state,omitempty"`

	// Output only. Batch state details, such as a failure
	//  description if the state is `FAILED`.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. The time when the batch entered a current state.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.state_time
	StateTime *string `json:"stateTime,omitempty"`

	// Output only. The email address of the user who created the batch.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. The resource name of the operation associated with this batch.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.operation
	Operation *string `json:"operation,omitempty"`

	// Output only. Historical state information for the batch.
	// +kcc:proto:field=google.cloud.dataproc.v1.Batch.state_history
	StateHistory []Batch_StateHistory `json:"stateHistory,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdataprocbatch;gcpdataprocbatchs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataprocBatch is the Schema for the DataprocBatch API
// +k8s:openapi-gen=true
type DataprocBatch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataprocBatchSpec   `json:"spec,omitempty"`
	Status DataprocBatchStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataprocBatchList contains a list of DataprocBatch
type DataprocBatchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocBatch `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocBatch{}, &DataprocBatchList{})
}
