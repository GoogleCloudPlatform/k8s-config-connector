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

var AIPlatformNotebookRuntimeGVK = GroupVersion.WithKind("AIPlatformNotebookRuntime")

// AIPlatformNotebookRuntimeSpec defines the desired state of AIPlatformNotebookRuntime
// +kcc:spec:proto=google.cloud.aiplatform.v1.NotebookRuntime
type AIPlatformNotebookRuntimeSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The AIPlatformNotebookRuntime name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The user email of the NotebookRuntime.
	// +required
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.runtime_user
	RuntimeUser *string `json:"runtimeUser,omitempty"`

	// Required. The display name of the NotebookRuntime.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +required
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.description
	Description *string `json:"description,omitempty"`

	// The labels with user-defined metadata to organize your
	//  NotebookRuntime.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one NotebookRuntime
	//  (System labels are excluded).
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	//  and are immutable. Following system labels exist for NotebookRuntime:
	//
	//  * "aiplatform.googleapis.com/notebook_runtime_gce_instance_id": output
	//  only, its value is the Compute Engine instance id.
	//  * "aiplatform.googleapis.com/colab_enterprise_entry_service": its value is
	//  either "bigquery" or "vertex"; if absent, it should be "vertex". This is to
	//  describe the entry service, either BigQuery or Vertex.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The Compute Engine tags to add to runtime (see [Tagging
	//  instances](https://cloud.google.com/vpc/docs/add-remove-network-tags)).
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.network_tags
	NetworkTags []string `json:"networkTags,omitempty"`
}

// AIPlatformNotebookRuntimeStatus defines the config connector machine state of AIPlatformNotebookRuntime
type AIPlatformNotebookRuntimeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformNotebookRuntime resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformNotebookRuntimeObservedState `json:"observedState,omitempty"`
}

// AIPlatformNotebookRuntimeObservedState is the state of the AIPlatformNotebookRuntime resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.NotebookRuntime
type AIPlatformNotebookRuntimeObservedState struct {
	// Output only. The resource name of the NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.name
	Name *string `json:"name,omitempty"`

	// Output only. The pointer to NotebookRuntimeTemplate this NotebookRuntime is
	//  created from.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.notebook_runtime_template_ref
	NotebookRuntimeTemplateRef *NotebookRuntimeTemplateRef `json:"notebookRuntimeTemplateRef,omitempty"`

	// Output only. The proxy endpoint used to access the NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.proxy_uri
	ProxyURI *string `json:"proxyURI,omitempty"`

	// Output only. Timestamp when this NotebookRuntime was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this NotebookRuntime was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The health state of the NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.health_state
	HealthState *string `json:"healthState,omitempty"`

	// Output only. Deprecated: This field is no longer used and the "Vertex AI
	//  Notebook Service Account"
	//  (service-PROJECT_NUMBER@gcp-sa-aiplatform-vm.iam.gserviceaccount.com) is
	//  used for the runtime workload identity.
	//  See
	//  https://cloud.google.com/iam/docs/service-agents#vertex-ai-notebook-service-account
	//  for more details.
	//
	//  The service account that the NotebookRuntime workload runs as.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. The runtime (instance) state of the NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.runtime_state
	RuntimeState *string `json:"runtimeState,omitempty"`

	// Output only. Whether NotebookRuntime is upgradable.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.is_upgradable
	IsUpgradable *bool `json:"isUpgradable,omitempty"`

	// Output only. Timestamp when this NotebookRuntime will be expired:
	//  1. System Predefined NotebookRuntime: 24 hours after creation. After
	//  expiration, system predifined runtime will be deleted.
	//  2. User created NotebookRuntime: 6 months after last upgrade. After
	//  expiration, user created runtime will be stopped and allowed for upgrade.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.expiration_time
	ExpirationTime *string `json:"expirationTime,omitempty"`

	// Output only. The VM os image version of NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.version
	Version *string `json:"version,omitempty"`

	// Output only. The type of the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.notebook_runtime_type
	NotebookRuntimeType *string `json:"notebookRuntimeType,omitempty"`

	// Output only. The specification of a single machine used by the notebook
	//  runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// Output only. The specification of [persistent
	//  disk][https://cloud.google.com/compute/docs/disks/persistent-disks]
	//  attached to the notebook runtime as data disk storage.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.data_persistent_disk_spec
	DataPersistentDiskSpec *PersistentDiskSpec `json:"dataPersistentDiskSpec,omitempty"`

	// Output only. Network spec of the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.network_spec
	NetworkSpec *NetworkSpec `json:"networkSpec,omitempty"`

	// Output only. The idle shutdown configuration of the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.idle_shutdown_config
	IdleShutdownConfig *NotebookIdleShutdownConfig `json:"idleShutdownConfig,omitempty"`

	// Output only. EUC configuration of the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.euc_config
	EUCConfig *NotebookEUCConfigObservedState `json:"eucConfig,omitempty"`

	// Output only. Runtime Shielded VM spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.shielded_vm_config
	ShieldedVMConfig *ShieldedVMConfig `json:"shieldedVMConfig,omitempty"`

	// Output only. Software config of the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.software_config
	SoftwareConfig *NotebookSoftwareConfigObservedState `json:"softwareConfig,omitempty"`

	// Output only. Customer-managed encryption key spec for the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformnotebookruntime;gcpaiplatformnotebookruntimes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformNotebookRuntime is the Schema for the AIPlatformNotebookRuntime API
// +k8s:openapi-gen=true
type AIPlatformNotebookRuntime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformNotebookRuntimeSpec   `json:"spec,omitempty"`
	Status AIPlatformNotebookRuntimeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformNotebookRuntimeList contains a list of AIPlatformNotebookRuntime
type AIPlatformNotebookRuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformNotebookRuntime `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformNotebookRuntime{}, &AIPlatformNotebookRuntimeList{})
}
