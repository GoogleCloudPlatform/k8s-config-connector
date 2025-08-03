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
// resource: DeployDeployPolicy:DeployPolicy

package v1alpha1

// +kcc:proto=google.cloud.deploy.v1.DeliveryPipelineAttribute
type DeliveryPipelineAttribute struct {
	// ID of the `DeliveryPipeline`. The value of this field could be one of the
	//  following:
	//
	//  * The last segment of a pipeline name
	//  * "*", all delivery pipelines in a location
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipelineAttribute.id
	ID *string `json:"id,omitempty"`

	// DeliveryPipeline labels.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipelineAttribute.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.DeployPolicy
type DeployPolicy struct {

	// Description of the `DeployPolicy`. Max length is 255 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.description
	Description *string `json:"description,omitempty"`

	// User annotations. These attributes can only be set and used by the
	//  user, and not by Cloud Deploy. Annotations must meet the following
	//  constraints:
	//
	//  * Annotations are key/value pairs.
	//  * Valid annotation keys have two segments: an optional prefix and name,
	//  separated by a slash (`/`).
	//  * The name segment is required and must be 63 characters or less,
	//  beginning and ending with an alphanumeric character (`[a-z0-9A-Z]`) with
	//  dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between.
	//  * The prefix is optional. If specified, the prefix must be a DNS subdomain:
	//  a series of DNS labels separated by dots(`.`), not longer than 253
	//  characters in total, followed by a slash (`/`).
	//
	//  See
	//  https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/#syntax-and-character-set
	//  for more details.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.annotations
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
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`

	// When suspended, the policy will not prevent actions from occurring, even
	//  if the action violates the policy.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.suspended
	Suspended *bool `json:"suspended,omitempty"`

	// Required. Selected resources to which the policy will be applied. At least
	//  one selector is required. If one selector matches the resource the policy
	//  applies. For example, if there are two selectors and the action being
	//  attempted matches one of them, the policy will apply to that action.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.selectors
	Selectors []DeployPolicyResourceSelector `json:"selectors,omitempty"`

	// Required. Rules to apply. At least one rule must be present.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.rules
	Rules []PolicyRule `json:"rules,omitempty"`

	// The weak etag of the `Automation` resource.
	//  This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.etag
	Etag *string `json:"etag,omitempty"`
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

// +kcc:proto=google.cloud.deploy.v1.PolicyRule
type PolicyRule struct {
	// Rollout restrictions.
	// +kcc:proto:field=google.cloud.deploy.v1.PolicyRule.rollout_restriction
	RolloutRestriction *RolloutRestriction `json:"rolloutRestriction,omitempty"`
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

// +kcc:proto=google.cloud.deploy.v1.TargetAttribute
type TargetAttribute struct {
	// ID of the `Target`. The value of this field could be one of the
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

// +kcc:proto=google.cloud.deploy.v1.CustomTargetType
type CustomTargetType struct {
	// Optional. Name of the `CustomTargetType`. Format is
	//  `projects/{project}/locations/{location}/customTargetTypes/{customTargetType}`.
	//  The `customTargetType` component must match
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.name
	Name *string `json:"name,omitempty"`

	// Optional. Description of the `CustomTargetType`. Max length is 255
	//  characters.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.description
	Description *string `json:"description,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	//  the user, and not by Cloud Deploy. See
	//  https://google.aip.dev/128#annotations for more details such as format and
	//  size limitations.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.annotations
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
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.etag
	Etag *string `json:"etag,omitempty"`

	// Configures render and deploy for the `CustomTargetType` using Skaffold
	//  custom actions.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.custom_actions
	CustomActions *CustomTargetSkaffoldActions `json:"customActions,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules
type SkaffoldModules struct {
	// Optional. The Skaffold Config modules to use from the specified source.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.configs
	Configs []string `json:"configs,omitempty"`

	// Remote git repository containing the Skaffold Config modules.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.git
	Git *SkaffoldModules_SkaffoldGitSource `json:"git,omitempty"`

	// Cloud Storage bucket containing the Skaffold Config modules.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.google_cloud_storage
	GoogleCloudStorage *SkaffoldModules_SkaffoldGCSSource `json:"googleCloudStorage,omitempty"`

	// Cloud Build V2 repository containing the Skaffold Config modules.
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

// +kcc:proto=google.cloud.deploy.v1.AdvanceRolloutRule
type AdvanceRolloutRule struct {
	// Required. ID of the rule. This id must be unique in the `Automation`
	//  resource to which this rule belongs. The format is
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.id
	ID *string `json:"id,omitempty"`

	// Optional. Proceeds only after phase name matched any one in the list.
	//  This value must consist of lower-case letters, numbers, and hyphens,
	//  start with a letter and end with a letter or a number, and have a max
	//  length of 63 characters. In other words, it must match the following
	//  regex: `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.source_phases
	SourcePhases []string `json:"sourcePhases,omitempty"`

	// Optional. How long to wait after a rollout is finished.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.wait
	Wait *string `json:"wait,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AutomationResourceSelector
type AutomationResourceSelector struct {
	// Contains attributes about a target.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationResourceSelector.targets
	Targets []TargetAttribute `json:"targets,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AutomationRule
type AutomationRule struct {
	// Optional. `PromoteReleaseRule` will automatically promote a release from
	//  the current target to a specified target.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.promote_release_rule
	PromoteReleaseRule *PromoteReleaseRule `json:"promoteReleaseRule,omitempty"`

	// Optional. The `AdvanceRolloutRule` will automatically advance a
	//  successful Rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.advance_rollout_rule
	AdvanceRolloutRule *AdvanceRolloutRule `json:"advanceRolloutRule,omitempty"`

	// Optional. The `RepairRolloutRule` will automatically repair a failed
	//  rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.repair_rollout_rule
	RepairRolloutRule *RepairRolloutRule `json:"repairRolloutRule,omitempty"`

	// Optional. The `TimedPromoteReleaseRule` will automatically promote a
	//  release from the current target(s) to the specified target(s) on a
	//  configured schedule.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.timed_promote_release_rule
	TimedPromoteReleaseRule *TimedPromoteReleaseRule `json:"timedPromoteReleaseRule,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AutomationRuleCondition
type AutomationRuleCondition struct {
	// Optional. Details around targets enumerated in the rule.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRuleCondition.targets_present_condition
	TargetsPresentCondition *TargetsPresentCondition `json:"targetsPresentCondition,omitempty"`

	// Optional. TimedPromoteReleaseCondition contains rule conditions specific
	//  to a an Automation with a timed promote release rule defined.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRuleCondition.timed_promote_release_condition
	TimedPromoteReleaseCondition *TimedPromoteReleaseCondition `json:"timedPromoteReleaseCondition,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PromoteReleaseRule
type PromoteReleaseRule struct {
	// Required. ID of the rule. This id must be unique in the `Automation`
	//  resource to which this rule belongs. The format is
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.id
	ID *string `json:"id,omitempty"`

	// Optional. How long the release need to be paused until being promoted to
	//  the next target.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.wait
	Wait *string `json:"wait,omitempty"`

	// Optional. The ID of the stage in the pipeline to which this `Release` is
	//  deploying. If unspecified, default it to the next stage in the promotion
	//  flow. The value of this field could be one of the following:
	//
	//  * The last segment of a target name
	//  * "@next", the next target in the promotion sequence
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.destination_target_id
	DestinationTargetID *string `json:"destinationTargetID,omitempty"`

	// Optional. The starting phase of the rollout created by this operation.
	//  Default to the first phase.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.destination_phase
	DestinationPhase *string `json:"destinationPhase,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.RepairPhaseConfig
type RepairPhaseConfig struct {
	// Optional. Retries a failed job.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairPhaseConfig.retry
	Retry *Retry `json:"retry,omitempty"`

	// Optional. Rolls back a `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairPhaseConfig.rollback
	Rollback *Rollback `json:"rollback,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.RepairRolloutRule
type RepairRolloutRule struct {
	// Required. ID of the rule. This id must be unique in the `Automation`
	//  resource to which this rule belongs. The format is
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.id
	ID *string `json:"id,omitempty"`

	// Optional. Phases within which jobs are subject to automatic repair actions
	//  on failure. Proceeds only after phase name matched any one in the list, or
	//  for all phases if unspecified. This value must consist of lower-case
	//  letters, numbers, and hyphens, start with a letter and end with a letter or
	//  a number, and have a max length of 63 characters. In other words, it must
	//  match the following regex: `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.phases
	Phases []string `json:"phases,omitempty"`

	// Optional. Jobs to repair. Proceeds only after job name matched any one in
	//  the list, or for all jobs if unspecified or empty. The phase that includes
	//  the job must match the phase ID specified in `source_phase`. This value
	//  must consist of lower-case letters, numbers, and hyphens, start with a
	//  letter and end with a letter or a number, and have a max length of 63
	//  characters. In other words, it must match the following regex:
	//  `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.jobs
	Jobs []string `json:"jobs,omitempty"`

	// Required. Defines the types of automatic repair phases for failed jobs.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.repair_phases
	RepairPhases []RepairPhaseConfig `json:"repairPhases,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Retry
type Retry struct {
	// Required. Total number of retries. Retry is skipped if set to 0; The
	//  minimum value is 1, and the maximum value is 10.
	// +kcc:proto:field=google.cloud.deploy.v1.Retry.attempts
	Attempts *int64 `json:"attempts,omitempty"`

	// Optional. How long to wait for the first retry. Default is 0, and the
	//  maximum value is 14d.
	// +kcc:proto:field=google.cloud.deploy.v1.Retry.wait
	Wait *string `json:"wait,omitempty"`

	// Optional. The pattern of how wait time will be increased. Default is
	//  linear. Backoff mode will be ignored if `wait` is 0.
	// +kcc:proto:field=google.cloud.deploy.v1.Retry.backoff_mode
	BackoffMode *string `json:"backoffMode,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Rollback
type Rollback struct {
	// Optional. The starting phase ID for the `Rollout`. If unspecified, the
	//  `Rollout` will start in the stable phase.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollback.destination_phase
	DestinationPhase *string `json:"destinationPhase,omitempty"`

	// Optional. If pending rollout exists on the target, the rollback operation
	//  will be aborted.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollback.disable_rollback_if_rollout_pending
	DisableRollbackIfRolloutPending *bool `json:"disableRollbackIfRolloutPending,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseCondition
type TimedPromoteReleaseCondition struct {
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseCondition.Targets
type TimedPromoteReleaseCondition_Targets struct {
	// Optional. The source target ID.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseCondition.Targets.source_target_id
	SourceTargetID *string `json:"sourceTargetID,omitempty"`

	// Optional. The destination target ID.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseCondition.Targets.destination_target_id
	DestinationTargetID *string `json:"destinationTargetID,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseRule
type TimedPromoteReleaseRule struct {
	// Required. ID of the rule. This ID must be unique in the `Automation`
	//  resource to which this rule belongs. The format is
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.id
	ID *string `json:"id,omitempty"`

	// Optional. The ID of the stage in the pipeline to which this `Release` is
	//  deploying. If unspecified, default it to the next stage in the promotion
	//  flow. The value of this field could be one of the following:
	//
	//  * The last segment of a target name
	//  * "@next", the next target in the promotion sequence
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.destination_target_id
	DestinationTargetID *string `json:"destinationTargetID,omitempty"`

	// Required. Schedule in crontab format. e.g. "0 9 * * 1" for every Monday at
	//  9am.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Required. The time zone in IANA format [IANA Time Zone
	//  Database](https://www.iana.org/time-zones) (e.g. America/New_York).
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Optional. The starting phase of the rollout created by this rule. Default
	//  to the first phase.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.destination_phase
	DestinationPhase *string `json:"destinationPhase,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AdvanceRolloutRule
type AdvanceRolloutRuleObservedState struct {
	// Output only. Information around the state of the Automation rule.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Automation
type AutomationObservedState struct {
	// Output only. Name of the `Automation`. Format is
	//  `projects/{project}/locations/{location}/deliveryPipelines/{delivery_pipeline}/automations/{automation}`.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.name
	Name *string `json:"name,omitempty"`

	// Output only. Unique identifier of the `Automation`.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the automation was created.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time at which the automation was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Required. List of Automation rules associated with the Automation resource.
	//  Must have at least one rule and limited to 250 rules per Delivery Pipeline.
	//  Note: the order of the rules here is not the same as the order of
	//  execution.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.rules
	Rules []AutomationRuleObservedState `json:"rules,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AutomationRule
type AutomationRuleObservedState struct {
	// Optional. `PromoteReleaseRule` will automatically promote a release from
	//  the current target to a specified target.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.promote_release_rule
	PromoteReleaseRule *PromoteReleaseRuleObservedState `json:"promoteReleaseRule,omitempty"`

	// Optional. The `AdvanceRolloutRule` will automatically advance a
	//  successful Rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.advance_rollout_rule
	AdvanceRolloutRule *AdvanceRolloutRuleObservedState `json:"advanceRolloutRule,omitempty"`

	// Optional. The `RepairRolloutRule` will automatically repair a failed
	//  rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.repair_rollout_rule
	RepairRolloutRule *RepairRolloutRuleObservedState `json:"repairRolloutRule,omitempty"`

	// Optional. The `TimedPromoteReleaseRule` will automatically promote a
	//  release from the current target(s) to the specified target(s) on a
	//  configured schedule.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.timed_promote_release_rule
	TimedPromoteReleaseRule *TimedPromoteReleaseRuleObservedState `json:"timedPromoteReleaseRule,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AutomationRuleCondition
type AutomationRuleConditionObservedState struct {
	// Optional. TimedPromoteReleaseCondition contains rule conditions specific
	//  to a an Automation with a timed promote release rule defined.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRuleCondition.timed_promote_release_condition
	TimedPromoteReleaseCondition *TimedPromoteReleaseConditionObservedState `json:"timedPromoteReleaseCondition,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PromoteReleaseRule
type PromoteReleaseRuleObservedState struct {
	// Output only. Information around the state of the Automation rule.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.RepairRolloutRule
type RepairRolloutRuleObservedState struct {
	// Output only. Information around the state of the 'Automation' rule.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseCondition
type TimedPromoteReleaseConditionObservedState struct {
	// Output only. When the next scheduled promotion(s) will occur.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseCondition.next_promotion_time
	NextPromotionTime *string `json:"nextPromotionTime,omitempty"`

	// Output only. A list of targets involved in the upcoming timed promotion(s).
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseCondition.targets_list
	TargetsList []TimedPromoteReleaseCondition_Targets `json:"targetsList,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseRule
type TimedPromoteReleaseRuleObservedState struct {
	// Output only. Information around the state of the Automation rule.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
}
