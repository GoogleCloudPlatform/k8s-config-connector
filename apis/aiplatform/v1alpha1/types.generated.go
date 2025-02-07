// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.aiplatform.v1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookIdleShutdownConfig
type NotebookIdleShutdownConfig struct {
	// Required. Duration is accurate to the second. In Notebook, Idle Timeout is
	//  accurate to minute so the range of idle_timeout (second) is: 10 * 60 ~ 1440
	//  * 60.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookIdleShutdownConfig.idle_timeout
	IdleTimeout *string `json:"idleTimeout,omitempty"`

	// Whether Idle Shutdown is disabled in this NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookIdleShutdownConfig.idle_shutdown_disabled
	IdleShutdownDisabled *bool `json:"idleShutdownDisabled,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookRuntime
type NotebookRuntime struct {

	// Required. The user email of the NotebookRuntime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.runtime_user
	RuntimeUser *string `json:"runtimeUser,omitempty"`

	// Required. The display name of the NotebookRuntime.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
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

// +kcc:proto=google.cloud.aiplatform.v1.NotebookRuntime
type NotebookRuntimeObservedState struct {
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

	// Output only. The service account that the NotebookRuntime workload runs as.
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

	// Output only. The idle shutdown configuration of the notebook runtime.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookRuntime.idle_shutdown_config
	IdleShutdownConfig *NotebookIdleShutdownConfig `json:"idleShutdownConfig,omitempty"`

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
