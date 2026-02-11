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
	cloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	containerkrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	gkehubkrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDeployTargetGVK = GroupVersion.WithKind("CloudDeployTarget")

// CloudDeployTargetSpec defines the desired state of CloudDeployTarget
// +kcc:spec:proto=google.cloud.deploy.v1.Target
type CloudDeployTargetSpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// Immutable. The location where the Target should reside.
	// +required
	Location *string `json:"location,omitempty"`

	// Optional. Description of the `Target`. Max length is 255 characters.
	Description *string `json:"description,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	// the user, and not by Cloud Deploy. See
	// https://google.aip.dev/128#annotations for more details such as format and
	// size limitations.
	// Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Labels are attributes that can be set and used by both the
	// user and by Cloud Deploy. Labels must meet the following constraints:
	//
	// * Keys and values can contain only lowercase letters, numeric characters,
	// underscores, and dashes.
	// * All characters must use UTF-8 encoding, and international characters are
	// allowed.
	// * Keys must start with a lowercase letter or international character.
	// * Each resource is limited to a maximum of 64 labels.
	//
	// Both keys and values are additionally constrained to be <= 128 bytes.
	// Labels map[string]string `json:"labels,omitempty"`

	// Optional. Whether or not the `Target` requires approval.
	RequireApproval *bool `json:"requireApproval,omitempty"`

	// Information specifying a GKE Cluster.
	Gke *GKECluster `json:"gke,omitempty"`

	// Information specifying an Anthos Cluster.
	AnthosCluster *AnthosCluster `json:"anthosCluster,omitempty"`

	// Information specifying a Cloud Run deployment target.
	Run *CloudRunLocation `json:"run,omitempty"`

	// Information specifying a multi-target.
	MultiTarget *MultiTarget `json:"multiTarget,omitempty"`

	// Information specifying a custom target.
	CustomTarget *CustomTarget `json:"customTarget,omitempty"`

	// Optional. The associations between the Target and other entities.
	AssociatedEntities map[string]*AssociatedEntities `json:"associatedEntities,omitempty"`

	// Optional. Configurations for all execution that relates to this `Target`.
	// Each `ExecutionEnvironmentUsage` value may only be used in a single
	// configuration; using the same value multiple times is an error.
	// When one or more configurations are specified, they must include the
	// `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values.
	// When no configurations are specified, execution will use the default
	// specified in `DefaultPool`.
	ExecutionConfigs []ExecutionConfig `json:"executionConfigs,omitempty"`

	// Optional. The deploy parameters to use for this target.
	DeployParameters map[string]string `json:"deployParameters,omitempty"`
}

// CloudDeployTargetStatus defines the config connector machine state of CloudDeployTarget
type CloudDeployTargetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDeployTarget resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDeployTargetObservedState `json:"observedState,omitempty"`
}

// CloudDeployTargetObservedState is the state of the CloudDeployTarget resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.deploy.v1.Target
type CloudDeployTargetObservedState struct {
	// Output only. Resource id of the `Target`.
	TargetId *string `json:"targetId,omitempty"`

	// Output only. Unique identifier of the `Target`.
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the `Target` was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Most recent time at which the `Target` was updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	// other fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddeploytarget;gcpclouddeploytargets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDeployTarget is the Schema for the CloudDeployTarget API
// +k8s:openapi-gen=true
type CloudDeployTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDeployTargetSpec   `json:"spec,omitempty"`
	Status CloudDeployTargetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDeployTargetList contains a list of CloudDeployTarget
type CloudDeployTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDeployTarget `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDeployTarget{}, &CloudDeployTargetList{})
}

// Moved here since auto-generator creates this type as GkeCluster, which causes
// issues down the line in both zz_generated.deepcopy and the mapper

// +kcc:proto=google.cloud.deploy.v1.GkeCluster
type GKECluster struct {
	// Optional. Information specifying a GKE Cluster. Format is
	//  `projects/{project_id}/locations/{location_id}/clusters/{cluster_id}`.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.cluster
	ClusterRef *containerkrm.ContainerClusterRef `json:"clusterRef,omitempty"`

	// Optional. If true, `cluster` is accessed using the private IP address of
	//  the control plane endpoint. Otherwise, the default IP address of the
	//  control plane endpoint is used. The default IP address is the private IP
	//  address for clusters with private control-plane endpoints and the public IP
	//  address otherwise.
	//
	//  Only specify this option when `cluster` is a [private GKE
	//  cluster](https://cloud.google.com/kubernetes-engine/docs/concepts/private-cluster-concept).
	//  Note that `internal_ip` and `dns_endpoint` cannot both be set to true.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.internal_ip
	InternalIP *bool `json:"internalIP,omitempty"`

	// Optional. If set, used to configure a
	//  [proxy](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/#proxy)
	//  to the Kubernetes server.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.proxy_url
	ProxyURL *string `json:"proxyURL,omitempty"`

	// Optional. If set, the cluster will be accessed using the DNS endpoint. Note
	//  that both `dns_endpoint` and `internal_ip` cannot be set to true.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.dns_endpoint
	DNSEndpoint *bool `json:"dnsEndpoint,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.DefaultPool
type DefaultPool struct {
	// Optional. Google service account to use for execution. If unspecified,
	//  the project execution service account
	//  (<PROJECT_NUMBER>-compute@developer.gserviceaccount.com) will be used.
	// +kcc:proto:field=google.cloud.deploy.v1.DefaultPool.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. Cloud Storage location where execution outputs should be stored.
	//  This can either be a bucket ("gs://my-bucket") or a path within a bucket
	//  ("gs://my-bucket/my-dir").
	//  If unspecified, a default bucket located in the same region will be used.
	// +kcc:proto:field=google.cloud.deploy.v1.DefaultPool.artifact_storage
	ArtifactStorage *string `json:"artifactStorage,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.ExecutionConfig
type ExecutionConfig struct {
	// Required. Usages when this configuration should be applied.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.usages
	Usages []string `json:"usages,omitempty"`

	// Optional. Use default Cloud Build pool.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.default_pool
	DefaultPool *DefaultPool `json:"defaultPool,omitempty"`

	// Optional. Use private Cloud Build pool.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.private_pool
	PrivatePool *PrivatePool `json:"privatePool,omitempty"`

	// Optional. The resource name of the `WorkerPool`, with the format
	//  `projects/{project}/locations/{location}/workerPools/{worker_pool}`.
	//  If this optional field is unspecified, the default Cloud Build pool will be
	//  used.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.worker_pool
	WorkerPoolRef *cloudbuildv1beta1.CloudBuildWorkerPoolRef `json:"workerPoolRef,omitempty"`

	// Optional. Google service account to use for execution. If unspecified,
	//  the project execution service account
	//  (<PROJECT_NUMBER>-compute@developer.gserviceaccount.com) is used.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. Cloud Storage location in which to store execution outputs. This
	//  can either be a bucket ("gs://my-bucket") or a path within a bucket
	//  ("gs://my-bucket/my-dir").
	//  If unspecified, a default bucket located in the same region will be used.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.artifact_storage
	ArtifactStorage *string `json:"artifactStorage,omitempty"`

	// Optional. Execution timeout for a Cloud Build Execution. This must be
	//  between 10m and 24h in seconds format. If unspecified, a default timeout of
	//  1h is used.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.execution_timeout
	ExecutionTimeout *string `json:"executionTimeout,omitempty"`

	// Optional. If true, additional logging will be enabled when running builds
	//  in this execution environment.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.verbose
	Verbose *bool `json:"verbose,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PrivatePool
type PrivatePool struct {
	// Required. Resource name of the Cloud Build worker pool to use. The format
	//  is `projects/{project}/locations/{location}/workerPools/{pool}`.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.worker_pool
	WorkerPoolRef *cloudbuildv1beta1.CloudBuildWorkerPoolRef `json:"workerPoolRef,omitempty"`

	// Optional. Google service account to use for execution. If unspecified,
	//  the project execution service account
	//  (<PROJECT_NUMBER>-compute@developer.gserviceaccount.com) will be used.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. Cloud Storage location where execution outputs should be stored.
	//  This can either be a bucket ("gs://my-bucket") or a path within a bucket
	//  ("gs://my-bucket/my-dir").
	//  If unspecified, a default bucket located in the same region will be used.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.artifact_storage
	ArtifactStorage *string `json:"artifactStorage,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.MultiTarget
type MultiTarget struct {
	// Required. The target_ids of this multiTarget.
	// +kcc:proto:field=google.cloud.deploy.v1.MultiTarget.target_ids
	TargetIDs []string `json:"targetIds,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AnthosCluster
type AnthosCluster struct {
	// Optional. Membership of the GKE Hub-registered cluster to which to apply
	//  the Skaffold configuration. Format is
	//  `projects/{project}/locations/{location}/memberships/{membership_name}`.
	// +kcc:proto:field=google.cloud.deploy.v1.AnthosCluster.membership
	MembershipRef *gkehubkrm.GKEHubMembershipRef `json:"membershipRef,omitempty"`
}
