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

package v1beta1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataflowJobGVK = GroupVersion.WithKind("DataflowJob")

// DataflowJobSpec defines the desired state of DataflowJob
// +kcc:spec:proto=google.dataflow.v1beta3.Job
type DataflowJobSpec struct {
	// Additional experiment flags for the job.
	AdditionalExperiments []string `json:"additionalExperiments,omitempty"`

	// Whether to enable Streaming Engine for the job.
	EnableStreamingEngine *bool `json:"enableStreamingEngine,omitempty"`

	// The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".
	IPConfiguration *string `json:"ipConfiguration,omitempty"`

	// The Cloud KMS key for the job.
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// The machine type to use for the job.
	MachineType *string `json:"machineType,omitempty"`

	// Immutable. The number of workers permitted to work on the job.
	// More workers may improve processing speed at additional cost.
	MaxWorkers *int32 `json:"maxWorkers,omitempty"`

	// The network to which VMs will be assigned.
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	// +kubebuilder:validation:Type=object
	Parameters *apiextensionsv1.JSON `json:"parameters,omitempty"`

	// Immutable. The region in which the created job should run.
	Region *string `json:"region,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition.
	// When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// The service account to run the job as.
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// The subnetwork to which VMs will be assigned.
	SubnetworkRef *refsv1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	// A writeable location on Google Cloud Storage for the Dataflow job to dump its temporary data.
	TempGcsLocation string `json:"tempGcsLocation"`

	// The Google Cloud Storage path to the Dataflow job template.
	TemplateGcsPath string `json:"templateGcsPath"`

	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job.
	// +kubebuilder:validation:Type=object
	TransformNameMapping *apiextensionsv1.JSON `json:"transformNameMapping,omitempty"`

	// Immutable. The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone *string `json:"zone,omitempty"`
}

// DataflowJobStatus defines the config connector machine state of DataflowJob
type DataflowJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The unique ID of this job.
	JobID *string `json:"jobId,omitempty"`

	// The current state of the resource, selected from the JobState enum.
	State *string `json:"state,omitempty"`

	// The type of this job, selected from the JobType enum.
	Type *string `json:"type,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataflowjob;gcpdataflowjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataflowJob is the Schema for the DataflowJob API
// +k8s:openapi-gen=true
type DataflowJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataflowJobSpec   `json:"spec,omitempty"`
	Status DataflowJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataflowJobList contains a list of DataflowJob
type DataflowJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataflowJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataflowJob{}, &DataflowJobList{})
}
