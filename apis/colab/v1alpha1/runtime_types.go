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

var ColabRuntimeGVK = GroupVersion.WithKind("ColabRuntime")

type Runtime_Parent struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. The name of the location where the Runtime will be created.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location is immutable."
	// Required.
	Location string `json:"location"`
}

// ColabRuntimeSpec defines the desired state of ColabRuntime
// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookRuntime
type ColabRuntimeSpec struct {
	Runtime_Parent `json:",inline"`

	// The ColabRuntime name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The pointer to NotebookRuntimeTemplate this NotebookRuntime is
	//  created from.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.notebook_runtime_template_ref
	NotebookRuntimeTemplateRef *NotebookRuntimeTemplateRef `json:"notebookRuntimeTemplateRef,omitempty"`

	// Required. The user email of the NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.runtime_user
	RuntimeUser *string `json:"runtimeUser,omitempty"`

	// Required. The display name of the NotebookRuntime.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.description
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
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The Compute Engine tags to add to runtime (see [Tagging
	//  instances](https://cloud.google.com/vpc/docs/add-remove-network-tags)).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.network_tags
	NetworkTags []string `json:"networkTags,omitempty"`
}

// ColabRuntimeStatus defines the config connector machine state of ColabRuntime
type ColabRuntimeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ColabRuntime resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ColabRuntimeObservedState `json:"observedState,omitempty"`
}

// ColabRuntimeObservedState is the state of the ColabRuntime resource as most recently observed in GCP.
// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookRuntime
type ColabRuntimeObservedState struct {
	// Output only. The proxy endpoint used to access the NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.proxy_uri
	ProxyURI *string `json:"proxyURI,omitempty"`

	// Output only. Timestamp when this NotebookRuntime was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this NotebookRuntime was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The health state of the NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.health_state
	HealthState *string `json:"healthState,omitempty"`

	// Output only. The service account that the NotebookRuntime workload runs as.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. The runtime (instance) state of the NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.runtime_state
	RuntimeState *string `json:"runtimeState,omitempty"`

	// Output only. Whether NotebookRuntime is upgradable.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.is_upgradable
	IsUpgradable *bool `json:"isUpgradable,omitempty"`

	// Output only. Timestamp when this NotebookRuntime will be expired:
	//  1. System Predefined NotebookRuntime: 24 hours after creation. After
	//  expiration, system predifined runtime will be deleted.
	//  2. User created NotebookRuntime: 6 months after last upgrade. After
	//  expiration, user created runtime will be stopped and allowed for upgrade.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.expiration_time
	ExpirationTime *string `json:"expirationTime,omitempty"`

	// Output only. The VM os image version of NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.version
	Version *string `json:"version,omitempty"`

	// Output only. The type of the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.notebook_runtime_type
	NotebookRuntimeType *string `json:"notebookRuntimeType,omitempty"`

	// Output only. The idle shutdown configuration of the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.idle_shutdown_config
	IdleShutdownConfig *NotebookIdleShutdownConfig `json:"idleShutdownConfig,omitempty"`

	// Output only. Customer-managed encryption key spec for the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.encryption_spec
	EncryptionSpec *EncryptionSpecObservedState `json:"encryptionSpec,omitempty"`

	/*
		// Output only. Reserved for future use.
		// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.satisfies_pzs
		SatisfiesPZS *bool `json:"satisfiesPZS,omitempty"`

		// Output only. Reserved for future use.
		// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookRuntime.satisfies_pzi
		SatisfiesPZI *bool `json:"satisfiesPZI,omitempty"`
	*/
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.EncryptionSpec
type EncryptionSpecObservedState struct {
	// The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EncryptionSpec.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcolabruntime;gcpcolabruntimes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ColabRuntime is the Schema for the ColabRuntime API
// +k8s:openapi-gen=true
type ColabRuntime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ColabRuntimeSpec   `json:"spec,omitempty"`
	Status ColabRuntimeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ColabRuntimeList contains a list of ColabRuntime
type ColabRuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ColabRuntime `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ColabRuntime{}, &ColabRuntimeList{})
}
