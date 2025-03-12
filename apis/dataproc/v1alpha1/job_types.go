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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataprocJobGVK = GroupVersion.WithKind("DataprocJob")

// DataprocJobSpec defines the desired state of DataprocJob
// +kcc:proto=google.cloud.dataproc.v1.Job
type DataprocJobSpec struct {
	// Optional. The fully qualified reference to the job, which can be used to
	//  obtain the equivalent REST path of the job resource. If this property
	//  is not specified when a job is created, the server generates a
	//  <code>job_id</code>.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.reference
	Reference *JobReference `json:"reference,omitempty"`

	// Required. Job information, including how, when, and where to
	//  run the job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.placement
	Placement *JobPlacement `json:"placement,omitempty"`

	// Optional. Job is a Hadoop job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.hadoop_job
	HadoopJob *HadoopJob `json:"hadoopJob,omitempty"`

	// Optional. Job is a Spark job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.spark_job
	SparkJob *SparkJob `json:"sparkJob,omitempty"`

	// Optional. Job is a PySpark job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.pyspark_job
	PysparkJob *PySparkJob `json:"pysparkJob,omitempty"`

	// Optional. Job is a Hive job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.hive_job
	HiveJob *HiveJob `json:"hiveJob,omitempty"`

	// Optional. Job is a Pig job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.pig_job
	PigJob *PigJob `json:"pigJob,omitempty"`

	// Optional. Job is a SparkR job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.spark_r_job
	SparkRJob *SparkRJob `json:"sparkRJob,omitempty"`

	// Optional. Job is a SparkSql job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.spark_sql_job
	SparkSQLJob *SparkSQLJob `json:"sparkSQLJob,omitempty"`

	// Optional. Job is a Presto job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.presto_job
	PrestoJob *PrestoJob `json:"prestoJob,omitempty"`

	// Optional. Job is a Trino job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.trino_job
	TrinoJob *TrinoJob `json:"trinoJob,omitempty"`

	// Optional. Job is a Flink job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.flink_job
	FlinkJob *FlinkJob `json:"flinkJob,omitempty"`

	// Optional. The labels to associate with this job.
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Label **values** can be empty, but, if present, must contain 1 to 63
	//  characters, and must conform to [RFC
	//  1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	//  associated with a job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Job scheduling configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.scheduling
	Scheduling *JobScheduling `json:"scheduling,omitempty"`

	// Optional. Driver scheduling configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.driver_scheduling_config
	DriverSchedulingConfig *DriverSchedulingConfig `json:"driverSchedulingConfig,omitempty"`

	// Required.
	Parent *DataprocJobParent `json:"parent,omitempty"`

	// The DataprocJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

type DataprocJobParent struct {

	// Required. The ID of the Google Cloud Platform project that the job belongs to.
	projectRef *refv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Required. The Dataproc region in which to handle the request.
	region string `json:"region,omitempty"`
}

// DataprocJobStatus defines the config connector machine state of DataprocJob
type DataprocJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataprocJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataprocJobObservedState `json:"observedState,omitempty"`
}

// DataprocJobObservedState is the state of the DataprocJob resource as most recently observed in GCP.
// +kcc:proto=google.cloud.dataproc.v1.Job
type DataprocJobObservedState struct {
	// Required. Job information, including how, when, and where to
	//  run the job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.placement
	Placement *JobPlacementObservedState `json:"placement,omitempty"`

	// Output only. The job status. Additional application-specific
	//  status information might be contained in the <code>type_job</code>
	//  and <code>yarn_applications</code> fields.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.status
	Status *JobStatus `json:"status,omitempty"`

	// Output only. The previous job status.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.status_history
	StatusHistory []JobStatus `json:"statusHistory,omitempty"`

	// Output only. The collection of YARN applications spun up by this job.
	//
	//  **Beta** Feature: This report is available for testing purposes only. It
	//  might be changed before final release.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.yarn_applications
	YarnApplications []YarnApplication `json:"yarnApplications,omitempty"`

	// Output only. A URI pointing to the location of the stdout of the job's
	//  driver program.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.driver_output_resource_uri
	DriverOutputResourceURI *string `json:"driverOutputResourceURI,omitempty"`

	// Output only. If present, the location of miscellaneous control files
	//  which can be used as part of job setup and handling. If not present,
	//  control files might be placed in the same location as `driver_output_uri`.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.driver_control_files_uri
	DriverControlFilesURI *string `json:"driverControlFilesURI,omitempty"`

	// Output only. A UUID that uniquely identifies a job within the project
	//  over time. This is in contrast to a user-settable reference.job_id that
	//  might be reused over time.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.job_uuid
	JobUuid *string `json:"jobUuid,omitempty"`

	// Output only. Indicates whether the job is completed. If the value is
	//  `false`, the job is still in progress. If `true`, the job is completed, and
	//  `status.state` field will indicate if it was successful, failed,
	//  or cancelled.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.done
	Done *bool `json:"done,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdataprocjob;gcpdataprocjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataprocJob is the Schema for the DataprocJob API
// +k8s:openapi-gen=true
type DataprocJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataprocJobSpec   `json:"spec,omitempty"`
	Status DataprocJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataprocJobList contains a list of DataprocJob
type DataprocJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocJob{}, &DataprocJobList{})
}
