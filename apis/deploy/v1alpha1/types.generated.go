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

// +kcc:proto=google.cloud.deploy.v1.AnthosCluster
type AnthosCluster struct {
	// Optional. Membership of the GKE Hub-registered cluster to which to apply
	//  the Skaffold configuration. Format is
	//  `projects/{project}/locations/{location}/memberships/{membership_name}`.
	// +kcc:proto:field=google.cloud.deploy.v1.AnthosCluster.membership
	Membership *string `json:"membership,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AssociatedEntities
type AssociatedEntities struct {
	// Optional. Information specifying GKE clusters as associated entities.
	// +kcc:proto:field=google.cloud.deploy.v1.AssociatedEntities.gke_clusters
	GkeClusters []GkeCluster `json:"gkeClusters,omitempty"`

	// Optional. Information specifying Anthos clusters as associated entities.
	// +kcc:proto:field=google.cloud.deploy.v1.AssociatedEntities.anthos_clusters
	AnthosClusters []AnthosCluster `json:"anthosClusters,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CloudRunLocation
type CloudRunLocation struct {
	// Required. The location for the Cloud Run Service. Format must be
	//  `projects/{project}/locations/{location}`.
	// +kcc:proto:field=google.cloud.deploy.v1.CloudRunLocation.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CustomTarget
type CustomTarget struct {
	// Required. The name of the CustomTargetType. Format must be
	//  `projects/{project}/locations/{location}/customTargetTypes/{custom_target_type}`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTarget.custom_target_type
	CustomTargetType *string `json:"customTargetType,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.DefaultPool
type DefaultPool struct {
	// Optional. Google service account to use for execution. If unspecified,
	//  the project execution service account
	//  (<PROJECT_NUMBER>-compute@developer.gserviceaccount.com) will be used.
	// +kcc:proto:field=google.cloud.deploy.v1.DefaultPool.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

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
	WorkerPool *string `json:"workerPool,omitempty"`

	// Optional. Google service account to use for execution. If unspecified,
	//  the project execution service account
	//  (<PROJECT_NUMBER>-compute@developer.gserviceaccount.com) is used.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

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

// +kcc:proto=google.cloud.deploy.v1.GkeCluster
type GkeCluster struct {
	// Optional. Information specifying a GKE Cluster. Format is
	//  `projects/{project_id}/locations/{location_id}/clusters/{cluster_id}`.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.cluster
	Cluster *string `json:"cluster,omitempty"`

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
	//  that `dns_endpoint` and `internal_ip` cannot both be set to true.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.dns_endpoint
	DNSEndpoint *bool `json:"dnsEndpoint,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.MultiTarget
type MultiTarget struct {
	// Required. The target_ids of this multiTarget.
	// +kcc:proto:field=google.cloud.deploy.v1.MultiTarget.target_ids
	TargetIds []string `json:"targetIds,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PrivatePool
type PrivatePool struct {
	// Required. Resource name of the Cloud Build worker pool to use. The format
	//  is `projects/{project}/locations/{location}/workerPools/{pool}`.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.worker_pool
	WorkerPool *string `json:"workerPool,omitempty"`

	// Optional. Google service account to use for execution. If unspecified,
	//  the project execution service account
	//  (<PROJECT_NUMBER>-compute@developer.gserviceaccount.com) will be used.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. Cloud Storage location where execution outputs should be stored.
	//  This can either be a bucket ("gs://my-bucket") or a path within a bucket
	//  ("gs://my-bucket/my-dir").
	//  If unspecified, a default bucket located in the same region will be used.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.artifact_storage
	ArtifactStorage *string `json:"artifactStorage,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Target
type Target struct {
	// Optional. Name of the `Target`. Format is
	//  `projects/{project}/locations/{location}/targets/{target}`.
	//  The `target` component must match `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`
	// +kcc:proto:field=google.cloud.deploy.v1.Target.name
	Name *string `json:"name,omitempty"`

	// Optional. Description of the `Target`. Max length is 255 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.description
	Description *string `json:"description,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	//  the user, and not by Cloud Deploy. See
	//  https://google.aip.dev/128#annotations for more details such as format and
	//  size limitations.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Labels are attributes that can be set and used by both the
	//  user and by Cloud Deploy. Labels must meet the following constraints:
	//
	//  * Keys and values can contain only lowercase letters, numeric characters,
	//  underscores, and dashes.
	//  * All characters must use UTF-8 encoding, and international characters are
	//  allowed.
	//  * Keys must start with a lowercase letter or international character.
	//  * Each resource is limited to a maximum of 64 labels.
	//
	//  Both keys and values are additionally constrained to be <= 128 bytes.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Whether or not the `Target` requires approval.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.require_approval
	RequireApproval *bool `json:"requireApproval,omitempty"`

	// Optional. Information specifying a GKE Cluster.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.gke
	Gke *GkeCluster `json:"gke,omitempty"`

	// Optional. Information specifying an Anthos Cluster.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.anthos_cluster
	AnthosCluster *AnthosCluster `json:"anthosCluster,omitempty"`

	// Optional. Information specifying a Cloud Run deployment target.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.run
	Run *CloudRunLocation `json:"run,omitempty"`

	// Optional. Information specifying a multiTarget.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.multi_target
	MultiTarget *MultiTarget `json:"multiTarget,omitempty"`

	// Optional. Information specifying a Custom Target.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.custom_target
	CustomTarget *CustomTarget `json:"customTarget,omitempty"`

	// TODO: unsupported map type with key string and value message

	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.etag
	Etag *string `json:"etag,omitempty"`

	// Configurations for all execution that relates to this `Target`.
	//  Each `ExecutionEnvironmentUsage` value may only be used in a single
	//  configuration; using the same value multiple times is an error.
	//  When one or more configurations are specified, they must include the
	//  `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values.
	//  When no configurations are specified, execution will use the default
	//  specified in `DefaultPool`.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.execution_configs
	ExecutionConfigs []ExecutionConfig `json:"executionConfigs,omitempty"`

	// Optional. The deploy parameters to use for this target.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.deploy_parameters
	DeployParameters map[string]string `json:"deployParameters,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Target
type TargetObservedState struct {
	// Output only. Resource id of the `Target`.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.target_id
	TargetID *string `json:"targetID,omitempty"`

	// Output only. Unique identifier of the `Target`.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the `Target` was created.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Most recent time at which the `Target` was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
