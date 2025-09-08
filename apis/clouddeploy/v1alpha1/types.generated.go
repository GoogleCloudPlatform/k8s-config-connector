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

// +generated:types
// krm.group: clouddeploy.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.deploy.v1
// resource: CloudDeployDeliveryPipeline:DeliveryPipeline
// resource: CloudDeployTarget:Target
// resource: CloudDeployPolicy:DeployPolicy
// resource: CloudDeployPolicy:CustomTargetType

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

// +kcc:proto=google.cloud.deploy.v1.Canary
type Canary struct {
	// Optional. Runtime specific configurations for the deployment strategy. The
	//  runtime configuration is used to determine how Cloud Deploy will split
	//  traffic to enable a progressive deployment.
	// +kcc:proto:field=google.cloud.deploy.v1.Canary.runtime_config
	RuntimeConfig *RuntimeConfig `json:"runtimeConfig,omitempty"`

	// Optional. Configures the progressive based deployment for a Target.
	// +kcc:proto:field=google.cloud.deploy.v1.Canary.canary_deployment
	CanaryDeployment *CanaryDeployment `json:"canaryDeployment,omitempty"`

	// Optional. Configures the progressive based deployment for a Target, but
	//  allows customizing at the phase level where a phase represents each of
	//  the percentage deployments.
	// +kcc:proto:field=google.cloud.deploy.v1.Canary.custom_canary_deployment
	CustomCanaryDeployment *CustomCanaryDeployment `json:"customCanaryDeployment,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CanaryDeployment
type CanaryDeployment struct {
	// Required. The percentage based deployments that will occur as a part of a
	//  `Rollout`. List is expected in ascending order and each integer n is
	//  0 <= n < 100.
	//  If the GatewayServiceMesh is configured for Kubernetes, then the range for
	//  n is 0 <= n <= 100.
	// +kcc:proto:field=google.cloud.deploy.v1.CanaryDeployment.percentages
	Percentages []int32 `json:"percentages,omitempty"`

	// Optional. Whether to run verify tests after each percentage deployment via
	//  `skaffold verify`.
	// +kcc:proto:field=google.cloud.deploy.v1.CanaryDeployment.verify
	Verify *bool `json:"verify,omitempty"`

	// Optional. Configuration for the predeploy job of the first phase. If this
	//  is not configured, there will be no predeploy job for this phase.
	// +kcc:proto:field=google.cloud.deploy.v1.CanaryDeployment.predeploy
	Predeploy *Predeploy `json:"predeploy,omitempty"`

	// Optional. Configuration for the postdeploy job of the last phase. If this
	//  is not configured, there will be no postdeploy job for this phase.
	// +kcc:proto:field=google.cloud.deploy.v1.CanaryDeployment.postdeploy
	Postdeploy *Postdeploy `json:"postdeploy,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CloudRunConfig
type CloudRunConfig struct {
	// Optional. Whether Cloud Deploy should update the traffic stanza in a Cloud
	//  Run Service on the user's behalf to facilitate traffic splitting. This is
	//  required to be true for CanaryDeployments, but optional for
	//  CustomCanaryDeployments.
	// +kcc:proto:field=google.cloud.deploy.v1.CloudRunConfig.automatic_traffic_control
	AutomaticTrafficControl *bool `json:"automaticTrafficControl,omitempty"`

	// Optional. A list of tags that are added to the canary revision while the
	//  canary phase is in progress.
	// +kcc:proto:field=google.cloud.deploy.v1.CloudRunConfig.canary_revision_tags
	CanaryRevisionTags []string `json:"canaryRevisionTags,omitempty"`

	// Optional. A list of tags that are added to the prior revision while the
	//  canary phase is in progress.
	// +kcc:proto:field=google.cloud.deploy.v1.CloudRunConfig.prior_revision_tags
	PriorRevisionTags []string `json:"priorRevisionTags,omitempty"`

	// Optional. A list of tags that are added to the final stable revision when
	//  the stable phase is applied.
	// +kcc:proto:field=google.cloud.deploy.v1.CloudRunConfig.stable_revision_tags
	StableRevisionTags []string `json:"stableRevisionTags,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CloudRunLocation
type CloudRunLocation struct {
	// Required. The location for the Cloud Run Service. Format must be
	//  `projects/{project}/locations/{location}`.
	// +kcc:proto:field=google.cloud.deploy.v1.CloudRunLocation.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CustomCanaryDeployment
type CustomCanaryDeployment struct {
	// Required. Configuration for each phase in the canary deployment in the
	//  order executed.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomCanaryDeployment.phase_configs
	PhaseConfigs []CustomCanaryDeployment_PhaseConfig `json:"phaseConfigs,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CustomCanaryDeployment.PhaseConfig
type CustomCanaryDeployment_PhaseConfig struct {
	// Required. The ID to assign to the `Rollout` phase.
	//  This value must consist of lower-case letters, numbers, and hyphens,
	//  start with a letter and end with a letter or a number, and have a max
	//  length of 63 characters. In other words, it must match the following
	//  regex: `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomCanaryDeployment.PhaseConfig.phase_id
	PhaseID *string `json:"phaseID,omitempty"`

	// Required. Percentage deployment for the phase.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomCanaryDeployment.PhaseConfig.percentage
	Percentage *int32 `json:"percentage,omitempty"`

	// Optional. Skaffold profiles to use when rendering the manifest for this
	//  phase. These are in addition to the profiles list specified in the
	//  `DeliveryPipeline` stage.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomCanaryDeployment.PhaseConfig.profiles
	Profiles []string `json:"profiles,omitempty"`

	// Optional. Whether to run verify tests after the deployment via `skaffold
	//  verify`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomCanaryDeployment.PhaseConfig.verify
	Verify *bool `json:"verify,omitempty"`

	// Optional. Configuration for the predeploy job of this phase. If this is
	//  not configured, there will be no predeploy job for this phase.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomCanaryDeployment.PhaseConfig.predeploy
	Predeploy *Predeploy `json:"predeploy,omitempty"`

	// Optional. Configuration for the postdeploy job of this phase. If this is
	//  not configured, there will be no postdeploy job for this phase.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomCanaryDeployment.PhaseConfig.postdeploy
	Postdeploy *Postdeploy `json:"postdeploy,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CustomTarget
type CustomTarget struct {
	// Required. The name of the CustomTargetType. Format must be
	//  `projects/{project}/locations/{location}/customTargetTypes/{custom_target_type}`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTarget.custom_target_type
	CustomTargetType *string `json:"customTargetType,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CustomTargetSkaffoldActions
type CustomTargetSkaffoldActions struct {
	// Optional. The Skaffold custom action responsible for render operations. If
	//  not provided then Cloud Deploy will perform the render operations via
	//  `skaffold render`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetSkaffoldActions.render_action
	RenderAction *string `json:"renderAction,omitempty"`

	// Required. The Skaffold custom action responsible for deploy operations.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetSkaffoldActions.deploy_action
	DeployAction *string `json:"deployAction,omitempty"`

	// Optional. List of Skaffold modules Cloud Deploy will include in the
	//  Skaffold Config as required before performing diagnose.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetSkaffoldActions.include_skaffold_modules
	IncludeSkaffoldModules []SkaffoldModules `json:"includeSkaffoldModules,omitempty"`
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

// +kcc:proto=google.cloud.deploy.v1.DeliveryPipelineAttribute
type DeliveryPipelineAttribute struct {
	// Optional. ID of the `DeliveryPipeline`. The value of this field could be
	//  one of the following:
	//
	//  * The last segment of a pipeline name
	//  * "*", all delivery pipelines in a location
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipelineAttribute.id
	ID *string `json:"id,omitempty"`

	// DeliveryPipeline labels.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipelineAttribute.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.DeployParameters
type DeployParameters struct {
	// Required. Values are deploy parameters in key-value pairs.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployParameters.values
	Values map[string]string `json:"values,omitempty"`

	// Optional. Deploy parameters are applied to targets with match labels.
	//  If unspecified, deploy parameters are applied to all targets (including
	//  child targets of a multi-target).
	// +kcc:proto:field=google.cloud.deploy.v1.DeployParameters.match_target_labels
	MatchTargetLabels map[string]string `json:"matchTargetLabels,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.DeployPolicyResourceSelector
type DeployPolicyResourceSelector struct {
	// Optional. Contains attributes about a delivery pipeline.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicyResourceSelector.delivery_pipeline
	DeliveryPipeline *DeliveryPipelineAttribute `json:"deliveryPipeline,omitempty"`

	// Optional. Contains attributes about a target.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicyResourceSelector.target
	Target *TargetAttribute `json:"target,omitempty"`
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
	//  that both `dns_endpoint` and `internal_ip` cannot be set to true.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.dns_endpoint
	DNSEndpoint *bool `json:"dnsEndpoint,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.KubernetesConfig
type KubernetesConfig struct {
	// Optional. Kubernetes Gateway API service mesh configuration.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.gateway_service_mesh
	GatewayServiceMesh *KubernetesConfig_GatewayServiceMesh `json:"gatewayServiceMesh,omitempty"`

	// Optional. Kubernetes Service networking configuration.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.service_networking
	ServiceNetworking *KubernetesConfig_ServiceNetworking `json:"serviceNetworking,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh
type KubernetesConfig_GatewayServiceMesh struct {
	// Required. Name of the Gateway API HTTPRoute.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.http_route
	HTTPRoute *string `json:"httpRoute,omitempty"`

	// Required. Name of the Kubernetes Service.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.service
	Service *string `json:"service,omitempty"`

	// Required. Name of the Kubernetes Deployment whose traffic is managed by
	//  the specified HTTPRoute and Service.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.deployment
	Deployment *string `json:"deployment,omitempty"`

	// Optional. The time to wait for route updates to propagate. The maximum
	//  configurable time is 3 hours, in seconds format. If unspecified, there is
	//  no wait time.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.route_update_wait_time
	RouteUpdateWaitTime *string `json:"routeUpdateWaitTime,omitempty"`

	// Optional. The amount of time to migrate traffic back from the canary
	//  Service to the original Service during the stable phase deployment. If
	//  specified, must be between 15s and 3600s. If unspecified, there is no
	//  cutback time.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.stable_cutback_duration
	StableCutbackDuration *string `json:"stableCutbackDuration,omitempty"`

	// Optional. The label to use when selecting Pods for the Deployment and
	//  Service resources. This label must already be present in both resources.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.pod_selector_label
	PodSelectorLabel *string `json:"podSelectorLabel,omitempty"`

	// Optional. Route destinations allow configuring the Gateway API HTTPRoute
	//  to be deployed to additional clusters. This option is available for
	//  multi-cluster service mesh set ups that require the route to exist in the
	//  clusters that call the service. If unspecified, the HTTPRoute will only
	//  be deployed to the Target cluster.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.route_destinations
	RouteDestinations *KubernetesConfig_GatewayServiceMesh_RouteDestinations `json:"routeDestinations,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.RouteDestinations
type KubernetesConfig_GatewayServiceMesh_RouteDestinations struct {
	// Required. The clusters where the Gateway API HTTPRoute resource will be
	//  deployed to. Valid entries include the associated entities IDs
	//  configured in the Target resource and "@self" to include the Target
	//  cluster.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.RouteDestinations.destination_ids
	DestinationIds []string `json:"destinationIds,omitempty"`

	// Optional. Whether to propagate the Kubernetes Service to the route
	//  destination clusters. The Service will always be deployed to the Target
	//  cluster even if the HTTPRoute is not. This option may be used to
	//  facilitate successful DNS lookup in the route destination clusters. Can
	//  only be set to true if destinations are specified.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.RouteDestinations.propagate_service
	PropagateService *bool `json:"propagateService,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.KubernetesConfig.ServiceNetworking
type KubernetesConfig_ServiceNetworking struct {
	// Required. Name of the Kubernetes Service.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.ServiceNetworking.service
	Service *string `json:"service,omitempty"`

	// Required. Name of the Kubernetes Deployment whose traffic is managed by
	//  the specified Service.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.ServiceNetworking.deployment
	Deployment *string `json:"deployment,omitempty"`

	// Optional. Whether to disable Pod overprovisioning. If Pod
	//  overprovisioning is disabled then Cloud Deploy will limit the number of
	//  total Pods used for the deployment strategy to the number of Pods the
	//  Deployment has on the cluster.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.ServiceNetworking.disable_pod_overprovisioning
	DisablePodOverprovisioning *bool `json:"disablePodOverprovisioning,omitempty"`

	// Optional. The label to use when selecting Pods for the Deployment
	//  resource. This label must already be present in the Deployment.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.ServiceNetworking.pod_selector_label
	PodSelectorLabel *string `json:"podSelectorLabel,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.MultiTarget
type MultiTarget struct {
	// Required. The target_ids of this multiTarget.
	// +kcc:proto:field=google.cloud.deploy.v1.MultiTarget.target_ids
	TargetIds []string `json:"targetIds,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.OneTimeWindow
type OneTimeWindow struct {
	// Required. Start date.
	// +kcc:proto:field=google.cloud.deploy.v1.OneTimeWindow.start_date
	StartDate *Date `json:"startDate,omitempty"`

	// Required. Start time (inclusive). Use 00:00 for the beginning of the day.
	// +kcc:proto:field=google.cloud.deploy.v1.OneTimeWindow.start_time
	StartTime *TimeOfDay `json:"startTime,omitempty"`

	// Required. End date.
	// +kcc:proto:field=google.cloud.deploy.v1.OneTimeWindow.end_date
	EndDate *Date `json:"endDate,omitempty"`

	// Required. End time (exclusive). You may use 24:00 for the end of the day.
	// +kcc:proto:field=google.cloud.deploy.v1.OneTimeWindow.end_time
	EndTime *TimeOfDay `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PipelineCondition
type PipelineCondition struct {
	// Details around the Pipeline's overall status.
	// +kcc:proto:field=google.cloud.deploy.v1.PipelineCondition.pipeline_ready_condition
	PipelineReadyCondition *PipelineReadyCondition `json:"pipelineReadyCondition,omitempty"`

	// Details around targets enumerated in the pipeline.
	// +kcc:proto:field=google.cloud.deploy.v1.PipelineCondition.targets_present_condition
	TargetsPresentCondition *TargetsPresentCondition `json:"targetsPresentCondition,omitempty"`

	// Details on the whether the targets enumerated in the pipeline are of the
	//  same type.
	// +kcc:proto:field=google.cloud.deploy.v1.PipelineCondition.targets_type_condition
	TargetsTypeCondition *TargetsTypeCondition `json:"targetsTypeCondition,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PipelineReadyCondition
type PipelineReadyCondition struct {
	// True if the Pipeline is in a valid state. Otherwise at least one condition
	//  in `PipelineCondition` is in an invalid state. Iterate over those
	//  conditions and see which condition(s) has status = false to find out what
	//  is wrong with the Pipeline.
	// +kcc:proto:field=google.cloud.deploy.v1.PipelineReadyCondition.status
	Status *bool `json:"status,omitempty"`

	// Last time the condition was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.PipelineReadyCondition.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PolicyRule
type PolicyRule struct {
	// Optional. Rollout restrictions.
	// +kcc:proto:field=google.cloud.deploy.v1.PolicyRule.rollout_restriction
	RolloutRestriction *RolloutRestriction `json:"rolloutRestriction,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Postdeploy
type Postdeploy struct {
	// Optional. A sequence of Skaffold custom actions to invoke during execution
	//  of the postdeploy job.
	// +kcc:proto:field=google.cloud.deploy.v1.Postdeploy.actions
	Actions []string `json:"actions,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Predeploy
type Predeploy struct {
	// Optional. A sequence of Skaffold custom actions to invoke during execution
	//  of the predeploy job.
	// +kcc:proto:field=google.cloud.deploy.v1.Predeploy.actions
	Actions []string `json:"actions,omitempty"`
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

// +kcc:proto=google.cloud.deploy.v1.RolloutRestriction
type RolloutRestriction struct {
	// Required. Restriction rule ID. Required and must be unique within a
	//  DeployPolicy. The format is `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
	// +kcc:proto:field=google.cloud.deploy.v1.RolloutRestriction.id
	ID *string `json:"id,omitempty"`

	// Optional. What invoked the action. If left empty, all invoker types will be
	//  restricted.
	// +kcc:proto:field=google.cloud.deploy.v1.RolloutRestriction.invokers
	Invokers []string `json:"invokers,omitempty"`

	// Optional. Rollout actions to be restricted as part of the policy. If left
	//  empty, all actions will be restricted.
	// +kcc:proto:field=google.cloud.deploy.v1.RolloutRestriction.actions
	Actions []string `json:"actions,omitempty"`

	// Required. Time window within which actions are restricted.
	// +kcc:proto:field=google.cloud.deploy.v1.RolloutRestriction.time_windows
	TimeWindows *TimeWindows `json:"timeWindows,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.RuntimeConfig
type RuntimeConfig struct {
	// Optional. Kubernetes runtime configuration.
	// +kcc:proto:field=google.cloud.deploy.v1.RuntimeConfig.kubernetes
	Kubernetes *KubernetesConfig `json:"kubernetes,omitempty"`

	// Optional. Cloud Run runtime configuration.
	// +kcc:proto:field=google.cloud.deploy.v1.RuntimeConfig.cloud_run
	CloudRun *CloudRunConfig `json:"cloudRun,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SerialPipeline
type SerialPipeline struct {
	// Optional. Each stage specifies configuration for a `Target`. The ordering
	//  of this list defines the promotion flow.
	// +kcc:proto:field=google.cloud.deploy.v1.SerialPipeline.stages
	Stages []Stage `json:"stages,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules
type SkaffoldModules struct {
	// Optional. The Skaffold Config modules to use from the specified source.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.configs
	Configs []string `json:"configs,omitempty"`

	// Optional. Remote git repository containing the Skaffold Config modules.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.git
	Git *SkaffoldModules_SkaffoldGitSource `json:"git,omitempty"`

	// Optional. Cloud Storage bucket containing the Skaffold Config modules.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.google_cloud_storage
	GoogleCloudStorage *SkaffoldModules_SkaffoldGCSSource `json:"googleCloudStorage,omitempty"`

	// Optional. Cloud Build V2 repository containing the Skaffold Config
	//  modules.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.google_cloud_build_repo
	GoogleCloudBuildRepo *SkaffoldModules_SkaffoldGcbRepoSource `json:"googleCloudBuildRepo,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCSSource
type SkaffoldModules_SkaffoldGCSSource struct {
	// Required. Cloud Storage source paths to copy recursively. For example,
	//  providing "gs://my-bucket/dir/configs/*" will result in Skaffold copying
	//  all files within the "dir/configs" directory in the bucket "my-bucket".
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCSSource.source
	Source *string `json:"source,omitempty"`

	// Optional. Relative path from the source to the Skaffold file.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCSSource.path
	Path *string `json:"path,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource
type SkaffoldModules_SkaffoldGitSource struct {
	// Required. Git repository the package should be cloned from.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource.repo
	Repo *string `json:"repo,omitempty"`

	// Optional. Relative path from the repository root to the Skaffold file.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource.path
	Path *string `json:"path,omitempty"`

	// Optional. Git branch or tag to use when cloning the repository.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource.ref
	Ref *string `json:"ref,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Stage
type Stage struct {
	// Optional. The target_id to which this stage points. This field refers
	//  exclusively to the last segment of a target name. For example, this field
	//  would just be `my-target` (rather than
	//  `projects/project/locations/location/targets/my-target`). The location of
	//  the `Target` is inferred to be the same as the location of the
	//  `DeliveryPipeline` that contains this `Stage`.
	// +kcc:proto:field=google.cloud.deploy.v1.Stage.target_id
	TargetID *string `json:"targetID,omitempty"`

	// Optional. Skaffold profiles to use when rendering the manifest for this
	//  stage's `Target`.
	// +kcc:proto:field=google.cloud.deploy.v1.Stage.profiles
	Profiles []string `json:"profiles,omitempty"`

	// Optional. The strategy to use for a `Rollout` to this stage.
	// +kcc:proto:field=google.cloud.deploy.v1.Stage.strategy
	Strategy *Strategy `json:"strategy,omitempty"`

	// Optional. The deploy parameters to use for the target in this stage.
	// +kcc:proto:field=google.cloud.deploy.v1.Stage.deploy_parameters
	DeployParameters []DeployParameters `json:"deployParameters,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Standard
type Standard struct {
	// Optional. Whether to verify a deployment via `skaffold verify`.
	// +kcc:proto:field=google.cloud.deploy.v1.Standard.verify
	Verify *bool `json:"verify,omitempty"`

	// Optional. Configuration for the predeploy job. If this is not configured,
	//  the predeploy job will not be present.
	// +kcc:proto:field=google.cloud.deploy.v1.Standard.predeploy
	Predeploy *Predeploy `json:"predeploy,omitempty"`

	// Optional. Configuration for the postdeploy job. If this is not configured,
	//  the postdeploy job will not be present.
	// +kcc:proto:field=google.cloud.deploy.v1.Standard.postdeploy
	Postdeploy *Postdeploy `json:"postdeploy,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Strategy
type Strategy struct {
	// Optional. Standard deployment strategy executes a single deploy and
	//  allows verifying the deployment.
	// +kcc:proto:field=google.cloud.deploy.v1.Strategy.standard
	Standard *Standard `json:"standard,omitempty"`

	// Optional. Canary deployment strategy provides progressive percentage
	//  based deployments to a Target.
	// +kcc:proto:field=google.cloud.deploy.v1.Strategy.canary
	Canary *Canary `json:"canary,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Target
type Target struct {
	// Identifier. Name of the `Target`. Format is
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
	GKE *GkeCluster `json:"gke,omitempty"`

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

	// Optional. Configurations for all execution that relates to this `Target`.
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

// +kcc:proto=google.cloud.deploy.v1.TargetAttribute
type TargetAttribute struct {
	// Optional. ID of the `Target`. The value of this field could be one of the
	//  following:
	//
	//  * The last segment of a target name
	//  * "*", all targets in a location
	// +kcc:proto:field=google.cloud.deploy.v1.TargetAttribute.id
	ID *string `json:"id,omitempty"`

	// Target labels.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetAttribute.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TargetsPresentCondition
type TargetsPresentCondition struct {
	// True if there aren't any missing Targets.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetsPresentCondition.status
	Status *bool `json:"status,omitempty"`

	// The list of Target names that do not exist. For example,
	//  `projects/{project_id}/locations/{location_name}/targets/{target_name}`.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetsPresentCondition.missing_targets
	MissingTargets []string `json:"missingTargets,omitempty"`

	// Last time the condition was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetsPresentCondition.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TargetsTypeCondition
type TargetsTypeCondition struct {
	// True if the targets are all a comparable type. For example this is true if
	//  all targets are GKE clusters. This is false if some targets are Cloud Run
	//  targets and others are GKE clusters.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetsTypeCondition.status
	Status *bool `json:"status,omitempty"`

	// Human readable error message.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetsTypeCondition.error_details
	ErrorDetails *string `json:"errorDetails,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimeWindows
type TimeWindows struct {
	// Required. The time zone in IANA format [IANA Time Zone
	//  Database](https://www.iana.org/time-zones) (e.g. America/New_York).
	// +kcc:proto:field=google.cloud.deploy.v1.TimeWindows.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Optional. One-time windows within which actions are restricted.
	// +kcc:proto:field=google.cloud.deploy.v1.TimeWindows.one_time_windows
	OneTimeWindows []OneTimeWindow `json:"oneTimeWindows,omitempty"`

	// Optional. Recurring weekly windows within which actions are restricted.
	// +kcc:proto:field=google.cloud.deploy.v1.TimeWindows.weekly_windows
	WeeklyWindows []WeeklyWindow `json:"weeklyWindows,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.WeeklyWindow
type WeeklyWindow struct {
	// Optional. Days of week. If left empty, all days of the week will be
	//  included.
	// +kcc:proto:field=google.cloud.deploy.v1.WeeklyWindow.days_of_week
	DaysOfWeek []string `json:"daysOfWeek,omitempty"`

	// Optional. Start time (inclusive). Use 00:00 for the beginning of the day.
	//  If you specify start_time you must also specify end_time. If left empty,
	//  this will block for the entire day for the days specified in days_of_week.
	// +kcc:proto:field=google.cloud.deploy.v1.WeeklyWindow.start_time
	StartTime *TimeOfDay `json:"startTime,omitempty"`

	// Optional. End time (exclusive). Use 24:00 to indicate midnight. If you
	//  specify end_time you must also specify start_time. If left empty, this will
	//  block for the entire day for the days specified in days_of_week.
	// +kcc:proto:field=google.cloud.deploy.v1.WeeklyWindow.end_time
	EndTime *TimeOfDay `json:"endTime,omitempty"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
}

// +kcc:proto=google.type.TimeOfDay
type TimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
	//  to allow the value "24:00:00" for scenarios like business closing time.
	// +kcc:proto:field=google.type.TimeOfDay.hours
	Hours *int32 `json:"hours,omitempty"`

	// Minutes of hour of day. Must be from 0 to 59.
	// +kcc:proto:field=google.type.TimeOfDay.minutes
	Minutes *int32 `json:"minutes,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may
	//  allow the value 60 if it allows leap-seconds.
	// +kcc:proto:field=google.type.TimeOfDay.seconds
	Seconds *int32 `json:"seconds,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	// +kcc:proto:field=google.type.TimeOfDay.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.deploy.v1.Target
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
