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

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIReasoningEngineGVK = GroupVersion.WithKind("VertexAIReasoningEngine")

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.EnvVar
type EnvVar struct {
	// Required. Name of the environment variable. Must be a valid C identifier.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EnvVar.name
	// +required
	Name *string `json:"name,omitempty"`

	// Required. Variables that reference a $(VAR_NAME) are expanded
	//  using the previous defined environment variables in the container and
	//  any service environment variables. If a variable cannot be resolved,
	//  the reference in the input string will be unchanged. The $(VAR_NAME)
	//  syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
	//  references will never be expanded, regardless of whether the variable
	//  exists or not.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EnvVar.value
	// +required
	Value *string `json:"value,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.SecretRef
type SecretRef struct {
	// Required. The name of the secret in Cloud Secret Manager.
	//  Format: {secret_name}.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SecretRef.secret
	// +required
	Secret *string `json:"secret,omitempty"`

	// The Cloud Secret Manager secret version.
	//  Can be 'latest' for the latest version, an integer for a specific
	//  version, or a version alias.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SecretRef.version
	Version *string `json:"version,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.SecretEnvVar
type SecretEnvVar struct {
	// Required. Name of the secret environment variable.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SecretEnvVar.name
	// +required
	Name *string `json:"name,omitempty"`

	// Required. Reference to a secret stored in the Cloud Secret Manager that
	//  will provide the value for this environment variable.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SecretEnvVar.secret_ref
	// +required
	SecretRef *SecretRef `json:"secretRef,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.DnsPeeringConfig
type DNSPeeringConfig struct {
	// Required. The DNS name suffix of the zone being peered to, e.g.,
	//  "my-internal-domain.corp.". Must end with a dot.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DnsPeeringConfig.domain
	// +required
	Domain *string `json:"domain,omitempty"`

	// Required. The project ID hosting the Cloud DNS managed zone that
	//  contains the 'domain'. The Vertex AI Service Agent requires the
	//  dns.peer role on this project.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DnsPeeringConfig.target_project
	// +required
	TargetProject *string `json:"targetProject,omitempty"`

	// Required. The VPC network name
	//  in the target_project where the DNS zone specified by 'domain' is
	//  visible.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DnsPeeringConfig.target_network
	// +required
	TargetNetwork *string `json:"targetNetwork,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.PscInterfaceConfig
type PSCInterfaceConfig struct {
	// Optional. The name of the Compute Engine
	//  [network
	//  attachment](https://cloud.google.com/vpc/docs/about-network-attachments) to
	//  attach to the resource within the region and user project.
	//  To specify this field, you must have already [created a network attachment]
	//  (https://cloud.google.com/vpc/docs/create-manage-network-attachments#create-network-attachments).
	//  This field is only used for resources using PSC-I.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PscInterfaceConfig.network_attachment
	NetworkAttachment *string `json:"networkAttachment,omitempty"`

	// Optional. DNS peering configurations. When specified, Vertex AI will
	//  attempt to configure DNS peering zones in the tenant project VPC
	//  to resolve the specified domains using the target network's Cloud DNS.
	//  The user must grant the dns.peer role to the Vertex AI Service Agent
	//  on the target project.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PscInterfaceConfig.dns_peering_configs
	DNSPeeringConfigs []DNSPeeringConfig `json:"dnsPeeringConfigs,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.PackageSpec
type ReasoningEngineSpec_PackageSpec struct {
	// Optional. The Cloud Storage URI of the pickled python object.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.PackageSpec.pickle_object_gcs_uri
	PickleObjectGCSURI *string `json:"pickleObjectGCSURI,omitempty"`

	// Optional. The Cloud Storage URI of the dependency files in tar.gz format.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.PackageSpec.dependency_files_gcs_uri
	DependencyFilesGCSURI *string `json:"dependencyFilesGCSURI,omitempty"`

	// Optional. The Cloud Storage URI of the `requirements.txt` file
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.PackageSpec.requirements_gcs_uri
	RequirementsGCSURI *string `json:"requirementsGCSURI,omitempty"`

	// Optional. The Python version. Supported values
	//  are 3.9, 3.10, 3.11, 3.12, 3.13. If not specified, the default value
	//  is 3.10.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.PackageSpec.python_version
	PythonVersion *string `json:"pythonVersion,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.DeploymentSpec
type ReasoningEngineSpec_DeploymentSpec struct {
	// Optional. Environment variables to be set with the Reasoning Engine
	//  deployment. The environment variables can be updated through the
	//  UpdateReasoningEngine API.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.DeploymentSpec.env
	Env []EnvVar `json:"env,omitempty"`

	// Optional. Environment variables where the value is a secret in Cloud
	//  Secret Manager.
	//  To use this feature, add 'Secret Manager Secret Accessor' role
	//  (roles/secretmanager.secretAccessor) to AI Platform Reasoning Engine
	//  Service Agent.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.DeploymentSpec.secret_env
	SecretEnv []SecretEnvVar `json:"secretEnv,omitempty"`

	// Optional. Configuration for PSC-I.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.DeploymentSpec.psc_interface_config
	PSCInterfaceConfig *PSCInterfaceConfig `json:"pscInterfaceConfig,omitempty"`

	// Optional. The minimum number of application instances that will be kept
	//  running at all times. Defaults to 1. Range: [0, 10].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.DeploymentSpec.min_instances
	MinInstances *int32 `json:"minInstances,omitempty"`

	// Optional. The maximum number of application instances that can be
	//  launched to handle increased traffic. Defaults to 100. Range: [1, 1000].
	//
	//  If VPC-SC or PSC-I is enabled, the acceptable range is [1, 100].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.DeploymentSpec.max_instances
	MaxInstances *int32 `json:"maxInstances,omitempty"`

	// Optional. Resource limits for each container. Only 'cpu' and 'memory'
	//  keys are supported. Defaults to {"cpu": "4", "memory": "4Gi"}.
	//
	//    - The only supported values for CPU are '1', '2', '4', '6' and '8'. For
	//      more information, go to
	//      https://cloud.google.com/run/docs/configuring/cpu.
	//    - The only supported values for memory are '1Gi', '2Gi', ... '32 Gi'.
	//    - For required cpu on different memory values, go to
	//      https://cloud.google.com/run/docs/configuring/memory-limits
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.DeploymentSpec.resource_limits
	ResourceLimits map[string]string `json:"resourceLimits,omitempty"`

	// Optional. Concurrency for each container and agent server. Recommended
	//  value: 2 * cpu + 1. Defaults to 9.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.DeploymentSpec.container_concurrency
	ContainerConcurrency *int32 `json:"containerConcurrency,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.InlineSource
type ReasoningEngineSpec_SourceCodeSpec_InlineSource struct {
	// Required. Input only. The application source code archive, provided as
	//  a compressed tarball
	//  (.tar.gz) file.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.InlineSource.source_archive
	// +required
	SourceArchive []byte `json:"sourceArchive,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.DeveloperConnectConfig
type ReasoningEngineSpec_SourceCodeSpec_DeveloperConnectConfig struct {
	// Required. The Developer Connect Git repository link, formatted as
	//  `projects/*/locations/*/connections/*/gitRepositoryLink/*`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.DeveloperConnectConfig.git_repository_link
	// +required
	GitRepositoryLink *string `json:"gitRepositoryLink,omitempty"`

	// Required. Directory, relative to the source root, in which to run the
	//  build.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.DeveloperConnectConfig.dir
	// +required
	Dir *string `json:"dir,omitempty"`

	// Required. The revision to fetch from the Git repository such as a
	//  branch, a tag, a commit SHA, or any Git ref.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.DeveloperConnectConfig.revision
	// +required
	Revision *string `json:"revision,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.DeveloperConnectSource
type ReasoningEngineSpec_SourceCodeSpec_DeveloperConnectSource struct {
	// Required. The Developer Connect configuration that defines the
	//  specific repository, revision, and directory to use as the source code
	//  root.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.DeveloperConnectSource.config
	// +required
	Config *ReasoningEngineSpec_SourceCodeSpec_DeveloperConnectConfig `json:"config,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.PythonSpec
type ReasoningEngineSpec_SourceCodeSpec_PythonSpec struct {
	// Optional. The version of Python to use. Support version
	//  includes 3.9, 3.10, 3.11, 3.12, 3.13.
	//  If not specified, default value is 3.10.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.PythonSpec.version
	Version *string `json:"version,omitempty"`

	// Optional. The Python module to load as the entrypoint, specified as a
	//  fully qualified module name. For example: path.to.agent.
	//  If not specified, defaults to "agent".
	//
	//  The project root will be added to Python sys.path, allowing imports
	//  to be specified relative to the root.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.PythonSpec.entrypoint_module
	EntrypointModule *string `json:"entrypointModule,omitempty"`

	// Optional. The name of the callable object within the
	//  `entrypoint_module` to use as the application If not specified,
	//  defaults to "root_agent".
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.PythonSpec.entrypoint_object
	EntrypointObject *string `json:"entrypointObject,omitempty"`

	// Optional. The path to the requirements file, relative to the source
	//  root. If not specified, defaults to "requirements.txt".
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.PythonSpec.requirements_file
	RequirementsFile *string `json:"requirementsFile,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec
type ReasoningEngineSpec_SourceCodeSpec struct {
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.inline_source
	InlineSource *ReasoningEngineSpec_SourceCodeSpec_InlineSource `json:"inlineSource,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.developer_connect_source
	DeveloperConnectSource *ReasoningEngineSpec_SourceCodeSpec_DeveloperConnectSource `json:"developerConnectSource,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.SourceCodeSpec.python_spec
	PythonSpec *ReasoningEngineSpec_SourceCodeSpec_PythonSpec `json:"pythonSpec,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec
type VertexAIReasoningEngineSpec_ReasoningEngineSpec struct {
	// Optional. The service account that the Reasoning Engine artifact runs as.
	//  It should have "roles/storage.objectViewer" for reading the user project's
	//  Cloud Storage and "roles/aiplatform.user" for using Vertex extensions. If
	//  not specified, the Vertex AI Reasoning Engine Service Agent in the project
	//  will be used.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. User provided package spec of the ReasoningEngine.
	//  Ignored when users directly specify a deployment image through
	//  `deployment_spec.first_party_image_override`, but keeping the
	//  field_behavior to avoid introducing breaking changes.
	//  The `deployment_source` field should not be set if `package_spec` is
	//  specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.package_spec
	PackageSpec *ReasoningEngineSpec_PackageSpec `json:"packageSpec,omitempty"`

	// Optional. The specification of a Reasoning Engine deployment.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.deployment_spec
	DeploymentSpec *ReasoningEngineSpec_DeploymentSpec `json:"deploymentSpec,omitempty"`

	// Optional. Declarations for object class methods in OpenAPI specification
	//  format.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.class_methods
	ClassMethods []apiextensionsv1.JSON `json:"classMethods,omitempty"`

	// Optional. The OSS agent framework used to develop the agent.
	//  Currently supported values: "google-adk", "langchain", "langgraph", "ag2",
	//  "llama-index", "custom".
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.agent_framework
	AgentFramework *string `json:"agentFramework,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.source_code_spec
	SourceCodeSpec *ReasoningEngineSpec_SourceCodeSpec `json:"sourceCodeSpec,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig.GenerationConfig
type ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig struct {
	// Required. The model used to generate memories.
	//  Format:
	//  `projects/{project}/locations/{location}/publishers/google/models/{model}`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig.GenerationConfig.model
	// +required
	Model *string `json:"model,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig.SimilaritySearchConfig
type ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig struct {
	// Required. The model used to generate embeddings to lookup similar
	//  memories. Format:
	//  `projects/{project}/locations/{location}/publishers/google/models/{model}`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig.SimilaritySearchConfig.embedding_model
	// +required
	EmbeddingModel *string `json:"embeddingModel,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig.TtlConfig.GranularTtlConfig
type ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig struct {
	// TODO: Add fields if needed. For now keeping it empty as per proto exploration.
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig.TtlConfig
type ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig struct {
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig.TtlConfig.default_ttl
	DefaultTTL *string `json:"defaultTTL,omitempty"`

	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig.TtlConfig.granular_ttl_config
	GranularTTLConfig *ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig `json:"granularTTLConfig,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig
type ReasoningEngineContextSpec_MemoryBankConfig struct {
	// Optional. Configuration for how to generate memories for the Memory Bank.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig.generation_config
	GenerationConfig *ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig `json:"generationConfig,omitempty"`

	// Optional. Configuration for how to perform similarity search on memories.
	//  If not set, the Memory Bank will use the default embedding model
	//  `text-embedding-005`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig.similarity_search_config
	SimilaritySearchConfig *ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig `json:"similaritySearchConfig,omitempty"`

	// Optional. Configuration for automatic TTL ("time-to-live") of the
	//  memories in the Memory Bank. If not set, TTL will not be applied
	//  automatically. The TTL can be explicitly set by modifying the
	//  `expire_time` of each Memory resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.MemoryBankConfig.ttl_config
	TtlConfig *ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig `json:"ttlConfig,omitempty"`
}

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec
type ReasoningEngineContextSpec struct {
	// Optional. Specification for a Memory Bank, which manages memories for the
	//  Agent Engine.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngineContextSpec.memory_bank_config
	MemoryBankConfig *ReasoningEngineContextSpec_MemoryBankConfig `json:"memoryBankConfig,omitempty"`
}

// VertexAIReasoningEngineSpec defines the desired state of VertexAIReasoningEngine
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ReasoningEngine
type VertexAIReasoningEngineSpec struct {
	// The Project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location for the resource.
	// +required
	Location string `json:"location"`

	// The VertexAIReasoningEngine name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the ReasoningEngine.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngine.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the ReasoningEngine.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngine.description
	Description *string `json:"description,omitempty"`

	// Optional. Configurations of the ReasoningEngine
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngine.spec
	Spec *VertexAIReasoningEngineSpec_ReasoningEngineSpec `json:"spec,omitempty"`

	// Optional. Used to perform consistent read-modify-write updates. If not set,
	//  a blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngine.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Configuration for how Agent Engine sub-resources should manage
	//  context.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngine.context_spec
	ContextSpec *ReasoningEngineContextSpec `json:"contextSpec,omitempty"`

	// Customer-managed encryption key spec for a ReasoningEngine. If set, this
	//  ReasoningEngine and all sub-resources of this ReasoningEngine will be
	//  secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngine.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAIReasoningEngineStatus defines the config connector machine state of VertexAIReasoningEngine
type VertexAIReasoningEngineStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIReasoningEngine resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIReasoningEngineObservedState `json:"observedState,omitempty"`
}

// VertexAIReasoningEngineObservedState is the state of the VertexAIReasoningEngine resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.ReasoningEngine
type VertexAIReasoningEngineObservedState struct {
	// Output only. Timestamp when this ReasoningEngine was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngine.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ReasoningEngine was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReasoningEngine.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaireasoningengine;gcpvertexaireasoningengines
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIReasoningEngine is the Schema for the VertexAIReasoningEngine API
// +k8s:openapi-gen=true
type VertexAIReasoningEngine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIReasoningEngineSpec   `json:"spec,omitempty"`
	Status VertexAIReasoningEngineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIReasoningEngineList contains a list of VertexAIReasoningEngine
type VertexAIReasoningEngineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIReasoningEngine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIReasoningEngine{}, &VertexAIReasoningEngineList{})
}
