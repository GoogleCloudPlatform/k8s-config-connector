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

// API sources for AIPlatformNotebookRuntimeTemplate, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/aiplatform/v1/notebook_runtime.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.notebookRuntimeTemplates

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformNotebookRuntimeTemplateGVK = GroupVersion.WithKind("AIPlatformNotebookRuntimeTemplate")

// AIPlatformNotebookRuntimeTemplateSpec defines the desired state of AIPlatformNotebookRuntimeTemplate
// +kcc:spec:proto=google.cloud.aiplatform.v1.NotebookRuntimeTemplate
// +kcc:required-from-proto
type AIPlatformNotebookRuntimeTemplateSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformNotebookRuntimeTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the NotebookRuntimeTemplate.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.description
	Description *string `json:"description,omitempty"`

	// Optional. Immutable. The specification of a single machine for the
	//  template.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// Optional. The specification of [persistent
	//  disk][https://cloud.google.com/compute/docs/disks/persistent-disks]
	//  attached to the runtime as data disk storage.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.data_persistent_disk_spec
	DataPersistentDiskSpec *PersistentDiskSpec `json:"dataPersistentDiskSpec,omitempty"`

	// Optional. Network spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.network_spec
	NetworkSpec *NetworkSpec `json:"networkSpec,omitempty"`

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize the
	//  NotebookRuntimeTemplates.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The idle shutdown configuration of NotebookRuntimeTemplate. This config
	//  will only be set when idle shutdown is enabled.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.idle_shutdown_config
	IdleShutdownConfig *NotebookIdleShutdownConfig `json:"idleShutdownConfig,omitempty"`

	// EUC configuration of the NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.euc_config
	EUCConfig *NotebookEUCConfig `json:"eucConfig,omitempty"`

	// Optional. Immutable. The type of the notebook runtime template.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.notebook_runtime_type
	NotebookRuntimeType *string `json:"notebookRuntimeType,omitempty"`

	// Optional. Immutable. Runtime Shielded VM spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.shielded_vm_config
	ShieldedVMConfig *ShieldedVMConfig `json:"shieldedVMConfig,omitempty"`

	// Optional. The Compute Engine tags to add to runtime (see [Tagging
	//  instances](https://cloud.google.com/vpc/docs/add-remove-network-tags)).
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.network_tags
	NetworkTags []string `json:"networkTags,omitempty"`

	// Customer-managed encryption key spec for the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Optional. The notebook software configuration of the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.software_config
	SoftwareConfig *NotebookSoftwareConfig `json:"softwareConfig,omitempty"`
}

// AIPlatformNotebookRuntimeTemplateStatus defines the config connector machine state of AIPlatformNotebookRuntimeTemplate
type AIPlatformNotebookRuntimeTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformNotebookRuntimeTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformNotebookRuntimeTemplateObservedState `json:"observedState,omitempty"`
}

// AIPlatformNotebookRuntimeTemplateObservedState is the state of the AIPlatformNotebookRuntimeTemplate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.NotebookRuntimeTemplate
type AIPlatformNotebookRuntimeTemplateObservedState struct {
	// Output only. Deprecated: This field has no behavior. Use
	//  notebook_runtime_type = 'ONE_CLICK' instead.
	//
	//  The default template to use if not specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.is_default
	IsDefault *bool `json:"isDefault,omitempty"`

	// EUC configuration of the NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.euc_config
	EUCConfig *NotebookEUCConfigObservedState `json:"eucConfig,omitempty"`

	// Output only. Timestamp when this NotebookRuntimeTemplate was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this NotebookRuntimeTemplate was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. The notebook software configuration of the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntimeTemplate.software_config
	SoftwareConfig *NotebookSoftwareConfigObservedState `json:"softwareConfig,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformnotebookruntimetemplate;gcpaiplatformnotebookruntimetemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformNotebookRuntimeTemplate is the Schema for the AIPlatformNotebookRuntimeTemplate API
// +k8s:openapi-gen=true
type AIPlatformNotebookRuntimeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformNotebookRuntimeTemplateSpec   `json:"spec,omitempty"`
	Status AIPlatformNotebookRuntimeTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformNotebookRuntimeTemplateList contains a list of AIPlatformNotebookRuntimeTemplate
type AIPlatformNotebookRuntimeTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformNotebookRuntimeTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformNotebookRuntimeTemplate{}, &AIPlatformNotebookRuntimeTemplateList{})
}
