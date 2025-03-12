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

// +kcc:proto=google.cloud.deploy.v1.Canary
type Canary struct {
	// Optional. Runtime specific configurations for the deployment strategy. The
	//  runtime configuration is used to determine how Cloud Deploy will split
	//  traffic to enable a progressive deployment.
	// +kcc:proto:field=google.cloud.deploy.v1.Canary.runtime_config
	RuntimeConfig *RuntimeConfig `json:"runtimeConfig,omitempty"`

	// Configures the progressive based deployment for a Target.
	// +kcc:proto:field=google.cloud.deploy.v1.Canary.canary_deployment
	CanaryDeployment *CanaryDeployment `json:"canaryDeployment,omitempty"`

	// Configures the progressive based deployment for a Target, but allows
	//  customizing at the phase level where a phase represents each of the
	//  percentage deployments.
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

	// Whether to run verify tests after each percentage deployment.
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
	// Whether Cloud Deploy should update the traffic stanza in a Cloud Run
	//  Service on the user's behalf to facilitate traffic splitting. This is
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

	// Skaffold profiles to use when rendering the manifest for this phase.
	//  These are in addition to the profiles list specified in the
	//  `DeliveryPipeline` stage.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomCanaryDeployment.PhaseConfig.profiles
	Profiles []string `json:"profiles,omitempty"`

	// Whether to run verify tests after the deployment.
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

// +kcc:proto=google.cloud.deploy.v1.DeliveryPipeline
type DeliveryPipeline struct {
	// Optional. Name of the `DeliveryPipeline`. Format is
	//  `projects/{project}/locations/{location}/deliveryPipelines/{deliveryPipeline}`.
	//  The `deliveryPipeline` component must match
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.name
	Name *string `json:"name,omitempty"`

	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.description
	Description *string `json:"description,omitempty"`

	// User annotations. These attributes can only be set and used by the
	//  user, and not by Cloud Deploy.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Labels are attributes that can be set and used by both the
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
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.labels
	Labels map[string]string `json:"labels,omitempty"`

	// SerialPipeline defines a sequential set of stages for a
	//  `DeliveryPipeline`.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.serial_pipeline
	SerialPipeline *SerialPipeline `json:"serialPipeline,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.etag
	Etag *string `json:"etag,omitempty"`

	// When suspended, no new releases or rollouts can be created,
	//  but in-progress ones will complete.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.suspended
	Suspended *bool `json:"suspended,omitempty"`
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

// +kcc:proto=google.cloud.deploy.v1.KubernetesConfig
type KubernetesConfig struct {
	// Kubernetes Gateway API service mesh configuration.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.gateway_service_mesh
	GatewayServiceMesh *KubernetesConfig_GatewayServiceMesh `json:"gatewayServiceMesh,omitempty"`

	// Kubernetes Service networking configuration.
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
	DestinationIDs []string `json:"destinationIDs,omitempty"`

	// Optional. Whether to propagate the Kubernetes Service to the route
	//  destination clusters. The Service will always be deployed to the Target
	//  cluster even if the HTTPRoute is not. This option may be used to
	//  facilitate successful DNS lookup in the route destination clusters.
	//  Can only be set to true if destinations are specified.
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

// +kcc:proto=google.cloud.deploy.v1.RuntimeConfig
type RuntimeConfig struct {
	// Kubernetes runtime configuration.
	// +kcc:proto:field=google.cloud.deploy.v1.RuntimeConfig.kubernetes
	Kubernetes *KubernetesConfig `json:"kubernetes,omitempty"`

	// Cloud Run runtime configuration.
	// +kcc:proto:field=google.cloud.deploy.v1.RuntimeConfig.cloud_run
	CloudRun *CloudRunConfig `json:"cloudRun,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SerialPipeline
type SerialPipeline struct {
	// Each stage specifies configuration for a `Target`. The ordering
	//  of this list defines the promotion flow.
	// +kcc:proto:field=google.cloud.deploy.v1.SerialPipeline.stages
	Stages []Stage `json:"stages,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Stage
type Stage struct {
	// The target_id to which this stage points. This field refers exclusively to
	//  the last segment of a target name. For example, this field would just be
	//  `my-target` (rather than
	//  `projects/project/locations/location/targets/my-target`). The location of
	//  the `Target` is inferred to be the same as the location of the
	//  `DeliveryPipeline` that contains this `Stage`.
	// +kcc:proto:field=google.cloud.deploy.v1.Stage.target_id
	TargetID *string `json:"targetID,omitempty"`

	// Skaffold profiles to use when rendering the manifest for this stage's
	//  `Target`.
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
	// Whether to verify a deployment.
	// +kcc:proto:field=google.cloud.deploy.v1.Standard.verify
	Verify *bool `json:"verify,omitempty"`

	// Optional. Configuration for the predeploy job. If this is not configured,
	//  predeploy job will not be present.
	// +kcc:proto:field=google.cloud.deploy.v1.Standard.predeploy
	Predeploy *Predeploy `json:"predeploy,omitempty"`

	// Optional. Configuration for the postdeploy job. If this is not configured,
	//  postdeploy job will not be present.
	// +kcc:proto:field=google.cloud.deploy.v1.Standard.postdeploy
	Postdeploy *Postdeploy `json:"postdeploy,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Strategy
type Strategy struct {
	// Standard deployment strategy executes a single deploy and allows
	//  verifying the deployment.
	// +kcc:proto:field=google.cloud.deploy.v1.Strategy.standard
	Standard *Standard `json:"standard,omitempty"`

	// Canary deployment strategy provides progressive percentage based
	//  deployments to a Target.
	// +kcc:proto:field=google.cloud.deploy.v1.Strategy.canary
	Canary *Canary `json:"canary,omitempty"`
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

// +kcc:proto=google.cloud.deploy.v1.DeliveryPipeline
type DeliveryPipelineObservedState struct {
	// Output only. Unique identifier of the `DeliveryPipeline`.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the pipeline was created.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Most recent time at which the pipeline was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Information around the state of the Delivery Pipeline.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.condition
	Condition *PipelineCondition `json:"condition,omitempty"`
}
