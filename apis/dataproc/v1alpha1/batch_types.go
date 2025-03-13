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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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

	*Parent `json:",inline"`

	// The DataprocBatch name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

type Parent struct {
	// Required.
	Location string `json:"location,omitempty"`

	// Required.
	ProjectRef v1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SparkBatch
type SparkBatch struct {
	// Optional. The HCFS URI of the jar file that contains the main class.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkBatch.main_jar_file_uri
	MainJarFileURI *string `json:"mainJarFileURI,omitempty"`

	// Optional. The name of the driver main class. The jar file that contains
	//  the class must be in the classpath or specified in `jar_file_uris`.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkBatch.main_class
	MainClass *string `json:"mainClass,omitempty"`

	// Optional. The arguments to pass to the driver. Do not include arguments
	//  that can be set as batch properties, such as `--conf`, since a collision
	//  can occur that causes an incorrect batch submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkBatch.args
	Args []string `json:"args,omitempty"`

	// Optional. HCFS URIs of jar files to add to the classpath of the
	//  Spark driver and tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkBatch.jar_file_uris
	JarFileURIs []string `json:"jarFileURIs,omitempty"`

	// Optional. HCFS URIs of files to be placed in the working directory of
	//  each executor.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkBatch.file_uris
	FileURIs []string `json:"fileURIs,omitempty"`

	// Optional. HCFS URIs of archives to be extracted into the working directory
	//  of each executor. Supported file types:
	//  `.jar`, `.tar`, `.tar.gz`, `.tgz`, and `.zip`.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkBatch.archive_uris
	ArchiveURIs []string `json:"archiveURIs,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.PySparkBatch
type PySparkBatch struct {
	// Required. The HCFS URI of the main Python file to use as the Spark driver.
	//  Must be a .py file.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkBatch.main_python_file_uri
	MainPythonFileURI *string `json:"mainPythonFileURI,omitempty"`

	// Optional. The arguments to pass to the driver. Do not include arguments
	//  that can be set as batch properties, such as `--conf`, since a collision
	//  can occur that causes an incorrect batch submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkBatch.args
	Args []string `json:"args,omitempty"`

	// Optional. HCFS file URIs of Python files to pass to the PySpark
	//  framework. Supported file types: `.py`, `.egg`, and `.zip`.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkBatch.python_file_uris
	PythonFileURIs []string `json:"pythonFileURIs,omitempty"`

	// Optional. HCFS URIs of jar files to add to the classpath of the
	//  Spark driver and tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkBatch.jar_file_uris
	JarFileURIs []string `json:"jarFileURIs,omitempty"`

	// Optional. HCFS URIs of files to be placed in the working directory of
	//  each executor.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkBatch.file_uris
	FileURIs []string `json:"fileURIs,omitempty"`

	// Optional. HCFS URIs of archives to be extracted into the working directory
	//  of each executor. Supported file types:
	//  `.jar`, `.tar`, `.tar.gz`, `.tgz`, and `.zip`.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkBatch.archive_uris
	ArchiveURIs []string `json:"archiveURIs,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SparkRBatch
type SparkRBatch struct {
	// Required. The HCFS URI of the main R file to use as the driver.
	//  Must be a `.R` or `.r` file.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRBatch.main_r_file_uri
	MainRFileURI *string `json:"mainRFileURI,omitempty"`

	// Optional. The arguments to pass to the Spark driver. Do not include
	//  arguments that can be set as batch properties, such as `--conf`, since a
	//  collision can occur that causes an incorrect batch submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRBatch.args
	Args []string `json:"args,omitempty"`

	// Optional. HCFS URIs of files to be placed in the working directory of
	//  each executor.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRBatch.file_uris
	FileURIs []string `json:"fileURIs,omitempty"`

	// Optional. HCFS URIs of archives to be extracted into the working directory
	//  of each executor. Supported file types:
	//  `.jar`, `.tar`, `.tar.gz`, `.tgz`, and `.zip`.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRBatch.archive_uris
	ArchiveURIs []string `json:"archiveURIs,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SparkSqlBatch
type SparkSQLBatch struct {
	// Required. The HCFS URI of the script that contains Spark SQL queries to
	//  execute.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlBatch.query_file_uri
	QueryFileURI *string `json:"queryFileURI,omitempty"`

	// Optional. Mapping of query variable names to values (equivalent to the
	//  Spark SQL command: `SET name="value";`).
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlBatch.query_variables
	QueryVariables map[string]string `json:"queryVariables,omitempty"`

	// Optional. HCFS URIs of jar files to be added to the Spark CLASSPATH.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlBatch.jar_file_uris
	JarFileURIs []string `json:"jarFileURIs,omitempty"`
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
	RuntimeInfo *RuntimeInfoObservedState `json:"runtimeInfo,omitempty"`

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
	StateHistory []Batch_StateHistoryObservedState `json:"stateHistory,omitempty"`
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
