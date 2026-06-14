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

var AIPlatformBatchPredictionJobGVK = GroupVersion.WithKind("AIPlatformBatchPredictionJob")

// AIPlatformBatchPredictionJobSpec defines the desired state of AIPlatformBatchPredictionJob
// +kcc:spec:proto=google.cloud.aiplatform.v1.BatchPredictionJob
type AIPlatformBatchPredictionJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The AIPlatformBatchPredictionJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The name of the Model resource that produces the predictions via this job,
	//  must share the same ancestor Location.
	//  Starting this job has no impact on any existing deployments of the Model
	//  and their resources.
	//  Exactly one of model and unmanaged_container_model must be set.
	//
	//  The model resource name may contain version id or version alias to specify
	//  the version.
	//   Example: `projects/{project}/locations/{location}/models/{model}@2`
	//               or
	//             `projects/{project}/locations/{location}/models/{model}@golden`
	//  if no version is specified, the default version will be deployed.
	//
	//  The model resource could also be a publisher model.
	//   Example: `publishers/{publisher}/models/{model}`
	//               or
	//            `projects/{project}/locations/{location}/publishers/{publisher}/models/{model}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.model
	ModelRef *AIPlatformModelRef `json:"modelRef,omitempty"`

	// The service account that the DeployedModel's container runs as. If not
	//  specified, a system generated one will be used, which
	//  has minimal permissions and the custom container, if used, may not have
	//  enough permission to access other Google Cloud resources.
	//
	//  Users deploying the Model must have the `iam.serviceAccounts.actAs`
	//  permission on this service account.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Customer-managed encryption key options for a BatchPredictionJob. If this
	//  is set, then all resources created by the BatchPredictionJob will be
	//  encrypted with the provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BatchPredictionJob.encryption_spec
	EncryptionSpec *BatchPredictionJobEncryptionSpec `json:"encryptionSpec,omitempty"`
}

type BatchPredictionJobEncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name
	KMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// AIPlatformBatchPredictionJobStatus defines the config connector machine state of AIPlatformBatchPredictionJob
type AIPlatformBatchPredictionJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformBatchPredictionJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformBatchPredictionJobObservedState `json:"observedState,omitempty"`
}

// AIPlatformBatchPredictionJobObservedState is the state of the AIPlatformBatchPredictionJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.BatchPredictionJob
type AIPlatformBatchPredictionJobObservedState struct {
	// Output only. Time when the BatchPredictionJob was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the BatchPredictionJob was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformbatchpredictionjob;gcpaiplatformbatchpredictionjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformBatchPredictionJob is the Schema for the AIPlatformBatchPredictionJob API
// +k8s:openapi-gen=true
type AIPlatformBatchPredictionJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformBatchPredictionJobSpec   `json:"spec,omitempty"`
	Status AIPlatformBatchPredictionJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformBatchPredictionJobList contains a list of AIPlatformBatchPredictionJob
type AIPlatformBatchPredictionJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformBatchPredictionJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformBatchPredictionJob{}, &AIPlatformBatchPredictionJobList{})
}
