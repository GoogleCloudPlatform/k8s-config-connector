// Copyright 2024 Google LLC
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

package v1beta1

import (
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

var DataflowFlexTemplateJobGVK = GroupVersion.WithKind("DataflowFlexTemplateJob")

// DataflowFlexTemplateJobSpec defines the desired state of DataflowFlexTemplateJob
// +kcc:spec:proto=google.dataflow.v1beta3.FlexTemplateRuntimeEnvironment
type DataflowFlexTemplateJobSpec struct {

	/* NOTYET
	// The DataflowFlexTemplateJob name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
	*/

	// Immutable. The region in which the created job should run.
	// +optional
	Region *string `json:"region"`

	// Cloud Storage path to a file with json serialized ContainerSpec as
	//  content.
	// +required
	ContainerSpecGcsPath *string `json:"containerSpecGcsPath,omitempty"`

	// The parameters for FlexTemplate.
	//  Ex. {"num_workers":"5"}
	Parameters *runtime.RawExtension `json:"parameters,omitempty"`

	// Map of transform name prefixes of the job to be replaced
	// with the corresponding name prefixes of the new job.
	// Only applicable when updating a pipeline.
	TransformNameMapping *runtime.RawExtension `json:"transformNameMapping,omitempty"`

	// The initial number of Google Compute Engine instances for the job.
	NumWorkers *int32 `json:"numWorkers,omitempty"`

	// The maximum number of Google Compute Engine instances to be made
	//  available to your pipeline during execution, from 1 to 1000.
	MaxWorkers *int32 `json:"maxWorkers,omitempty"`

	/* NOTYET
	// The Compute Engine [availability
	//  zone](https://cloud.google.com/compute/docs/regions-zones/regions-zones)
	//  for launching worker instances to run your pipeline.
	//  In the future, worker_zone will take precedence.
	Zone *string `json:"zone,omitempty"`
	*/

	// The email address of the service account to run the job as.
	ServiceAccountEmailRef *refs.IAMServiceAccountRef `json:"serviceAccountEmailRef,omitempty"`

	// The Cloud Storage path to use for temporary files.
	//  Must be a valid Cloud Storage URL, beginning with `gs://`.
	TempLocation *string `json:"tempLocation,omitempty"`

	// The machine type to use for the job. Defaults to the value from the
	//  template if not specified.
	MachineType *string `json:"machineType,omitempty"`

	// Additional experiment flags for the job.
	AdditionalExperiments []string `json:"additionalExperiments,omitempty"`

	// Network to which VMs will be assigned.  If empty or unspecified,
	//  the service will use the network "default".
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Subnetwork to which VMs will be assigned, if desired. You can specify a
	//  subnetwork using either a complete URL or an abbreviated path. Expected to
	//  be of the form
	//  "https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK"
	//  or "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in
	//  a Shared VPC network, you must use the complete URL.
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	/* NOTYET
	// Additional user labels to be specified for the job.
	//  Keys and values must follow the restrictions specified in the [labeling
	//  restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
	//  page.
	//  An object containing a list of "key": value pairs.
	//  Example: { "name": "wrench", "mass": "1kg", "count": "3" }.
	AdditionalUserLabels map[string]string `json:"additionalUserLabels,omitempty"`
	*/

	// The Cloud KMS key for the job.
	KMSKeyNameRef *kmsv1beta1.KMSKeyRef_OneOf `json:"kmsKeyNameRef,omitempty"`

	// Configuration for VM IPs.
	IPConfiguration *string `json:"ipConfiguration,omitempty"`

	/* NOTYET
	// The Compute Engine region
	//  (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in
	//  which worker processing should occur, e.g. "us-west1". Mutually exclusive
	//  with worker_zone. If neither worker_region nor worker_zone is specified,
	//  default to the control plane's region.
	WorkerRegion *string `json:"workerRegion,omitempty"`
	*/

	/* NOTYET
	// The Compute Engine zone
	//  (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in
	//  which worker processing should occur, e.g. "us-west1-a". Mutually exclusive
	//  with worker_region. If neither worker_region nor worker_zone is specified,
	//  a zone in the control plane's region is chosen based on available capacity.
	//  If both `worker_zone` and `zone` are set, `worker_zone` takes precedence.
	WorkerZone *string `json:"workerZone,omitempty"`
	*/

	// Whether to enable Streaming Engine for the job.
	EnableStreamingEngine *bool `json:"enableStreamingEngine,omitempty"`

	/*NOTYET
	// Set FlexRS goal for the job.
	//  https://cloud.google.com/dataflow/docs/guides/flexrs
	FlexrsGoal *string `json:"flexrsGoal,omitempty"`
	*/

	// The Cloud Storage path for staging local files.
	//  Must be a valid Cloud Storage URL, beginning with `gs://`.
	StagingLocation *string `json:"stagingLocation,omitempty"`

	// Docker registry location of container image to use for the 'worker harness.
	//  Default is the container for the version of the SDK. Note this field is
	//  only valid for portable pipelines.
	SDKContainerImage *string `json:"sdkContainerImage,omitempty"`

	/* NOTYET
	// Worker disk size, in gigabytes.
	DiskSizeGb *int32 `json:"diskSizeGb,omitempty"`
	*/

	// The algorithm to use for autoscaling
	AutoscalingAlgorithm *string `json:"autoscalingAlgorithm,omitempty"`

	/* NOTYET
	// If true, save a heap dump before killing a thread or process which is GC
	//  thrashing or out of memory. The location of the heap file will either be
	//  echoed back to the user, or the user will be given the opportunity to
	//  download the heap file.
	DumpHeapOnOom *bool `json:"dumpHeapOnOom,omitempty"`
	*/

	/* NOTYET
	// Cloud Storage bucket (directory) to upload heap dumps to the given
	//  location. Enabling this implies that heap dumps should be generated on OOM
	//  (dump_heap_on_oom is set to true).
	SaveHeapDumpsToGcsPath *string `json:"saveHeapDumpsToGcsPath,omitempty"`
	*/

	// The machine type to use for launching the job. The default is
	//  n1-standard-1.
	LauncherMachineType *string `json:"launcherMachineType,omitempty"`
}

// DataflowFlexTemplateJobStatus defines the config connector machine state of DataflowFlexTemplateJob
type DataflowFlexTemplateJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* NOTYET - keeping compatibility with terraform
	// A unique specifier for the DataflowFlexTemplateJob resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
	*/

	/* NOTYET - keeping compatibility with terraform
	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *DataflowFlexTemplateJobObservedState `json:"observedState,omitempty"`
	*/

	// TODO: Description for JobID?

	JobID string `json:"jobId,omitempty"`

	// The current state of the job.
	//
	//  Jobs are created in the `JOB_STATE_STOPPED` state unless otherwise
	//  specified.
	//
	//  A job in the `JOB_STATE_RUNNING` state may asynchronously enter a
	//  terminal state. After a job has reached a terminal state, no
	//  further state updates may be made.
	//
	//  This field may be mutated by the Cloud Dataflow service;
	//  callers cannot mutate it.
	CurrentState *string `json:"state,omitempty"`

	// The type of Cloud Dataflow job.
	Type *string `json:"type,omitempty"`
}

// DataflowFlexTemplateJobSpec defines the desired state of DataflowFlexTemplateJob
type DataflowFlexTemplateJobObservedState struct {
	/* NOTYET - maintaining compatibility with terraform for now

	// The current state of the job.
	//
	//  Jobs are created in the `JOB_STATE_STOPPED` state unless otherwise
	//  specified.
	//
	//  A job in the `JOB_STATE_RUNNING` state may asynchronously enter a
	//  terminal state. After a job has reached a terminal state, no
	//  further state updates may be made.
	//
	//  This field may be mutated by the Cloud Dataflow service;
	//  callers cannot mutate it.
	CurrentState *string `json:"currentState,omitempty"`

	// The timestamp associated with the current state.
	CurrentStateTime *string `json:"currentStateTime,omitempty"`

	// The timestamp when the job was initially created. Immutable and set by the
	//  Cloud Dataflow service.
	CreateTime *string `json:"createTime,omitempty"`

	// If this job is an update of an existing job, this field is the job ID
	//  of the job it replaced.
	//
	//  When sending a `CreateJobRequest`, you can update a job by specifying it
	//  here. The job named here is stopped, and its intermediate state is
	//  transferred to this job.
	ReplaceJobID *string `json:"replaceJobID,omitempty"`

	// If another job is an update of this job (and thus, this job is in
	//  `JOB_STATE_UPDATED`), this field contains the ID of that job.
	ReplacedByJobID *string `json:"replacedByJobID,omitempty"`

	// The timestamp when the job was started (transitioned to JOB_STATE_PENDING).
	//  Flexible resource scheduling jobs are started with some delay after job
	//  creation, so start_time is unset before start and is updated when the
	//  job is started by the Cloud Dataflow service. For other jobs, start_time
	//  always equals create_time and is immutable and set by the Cloud Dataflow
	//  service.
	StartTime *string `json:"startTime,omitempty"`

	// If this is specified, the job's initial state is populated from the given
	//  snapshot.
	CreatedFromSnapshotID *string `json:"createdFromSnapshotID,omitempty"`
	*/
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataflowflextemplatejob;gcpdataflowflextemplatejobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataflowFlexTemplateJob is the Schema for the DataflowFlexTemplateJob API
// +k8s:openapi-gen=true
type DataflowFlexTemplateJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec DataflowFlexTemplateJobSpec `json:"spec,omitempty"`

	Status DataflowFlexTemplateJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataflowFlexTemplateJobList contains a list of DataflowFlexTemplateJob
type DataflowFlexTemplateJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataflowFlexTemplateJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataflowFlexTemplateJob{}, &DataflowFlexTemplateJobList{})
}
