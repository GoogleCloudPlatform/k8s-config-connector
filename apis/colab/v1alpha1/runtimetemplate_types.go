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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ColabRuntimeTemplateGVK = GroupVersion.WithKind("ColabRuntimeTemplate")

type Parent struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. The name of the location where the RuntimeTemplate will be created.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location is immutable."
	// Required.
	Location string `json:"location"`
}

// ColabRuntimeTemplateSpec defines the desired state of ColabRuntimeTemplate
// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate
type ColabRuntimeTemplateSpec struct {
	Parent `json:",inline"`

	// Required. The display name of the NotebookRuntimeTemplate.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.description
	Description *string `json:"description,omitempty"`

	// Optional. Immutable. The specification of a single machine for the
	//  template.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// Optional. The specification of [persistent
	//  disk][https://cloud.google.com/compute/docs/disks/persistent-disks]
	//  attached to the runtime as data disk storage.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.data_persistent_disk_spec
	DataPersistentDiskSpec *PersistentDiskSpec `json:"dataPersistentDiskSpec,omitempty"`

	// Optional. Network spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.network_spec
	NetworkSpec *NetworkSpec `json:"networkSpec,omitempty"`

	// The service account that the runtime workload runs as.
	//  You can use any service account within the same project, but you
	//  must have the service account user permission to use the instance.
	//
	//  If not specified, the [Compute Engine default service
	//  account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account)
	//  is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.service_account
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// The labels with user-defined metadata to organize the
	//  NotebookRuntimeTemplates.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The idle shutdown configuration of NotebookRuntimeTemplate. This config
	//  will only be set when idle shutdown is enabled.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.idle_shutdown_config
	IdleShutdownConfig *NotebookIdleShutdownConfig `json:"idleShutdownConfig,omitempty"`

	// EUC configuration of the NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.euc_config
	EUCConfig *NotebookEUCConfig `json:"eucConfig,omitempty"`

	// Optional. Immutable. The type of the notebook runtime template.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.notebook_runtime_type
	NotebookRuntimeType *string `json:"notebookRuntimeType,omitempty"`

	// Optional. Immutable. Runtime Shielded VM spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.shielded_vm_config
	ShieldedVMConfig *ShieldedVMConfig `json:"shieldedVMConfig,omitempty"`

	// Optional. The Compute Engine tags to add to runtime (see [Tagging
	//  instances](https://cloud.google.com/vpc/docs/add-remove-network-tags)).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.network_tags
	NetworkTags []string `json:"networkTags,omitempty"`

	// Customer-managed encryption key spec for the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// The ColabRuntimeTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EncryptionSpec.kms_key_name
	KMSKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NetworkSpec
type NetworkSpec struct {
	// Whether to enable public internet access. Default false.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NetworkSpec.enable_internet_access
	EnableInternetAccess *bool `json:"enableInternetAccess,omitempty"`

	// The full name of the Google Compute Engine
	//  [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks)
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NetworkSpec.network
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// The name of the subnet that this instance is in.
	//  Format:
	//  `projects/{project_id_or_number}/regions/{region}/subnetworks/{subnetwork_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NetworkSpec.subnetwork
	SubnetworkRef refs.ComputeSubnetworkRef `json:"subnetworkRef"`
}

// ColabRuntimeTemplateStatus defines the config connector machine state of ColabRuntimeTemplate
type ColabRuntimeTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ColabRuntimeTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ColabRuntimeTemplateObservedState `json:"observedState,omitempty"`
}

// ColabRuntimeTemplateObservedState is the state of the ColabRuntimeTemplate resource as most recently observed in GCP.
// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate
type ColabRuntimeTemplateObservedState struct {
	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.etag
	Etag *string `json:"etag,omitempty"`

	// EUC configuration of the NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.euc_config
	EUCConfig *NotebookEUCConfigObservedState `json:"eucConfig,omitempty"`

	// Output only. Timestamp when this NotebookRuntimeTemplate was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this NotebookRuntimeTemplate was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpcolabruntimetemplate;gcpcolabruntimetemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ColabRuntimeTemplate is the Schema for the ColabRuntimeTemplate API
// +k8s:openapi-gen=true
type ColabRuntimeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ColabRuntimeTemplateSpec   `json:"spec,omitempty"`
	Status ColabRuntimeTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ColabRuntimeTemplateList contains a list of ColabRuntimeTemplate
type ColabRuntimeTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ColabRuntimeTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ColabRuntimeTemplate{}, &ColabRuntimeTemplateList{})
}
