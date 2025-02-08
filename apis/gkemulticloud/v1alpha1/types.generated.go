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


// +kcc:proto=google.cloud.gkemulticloud.v1.AwsAutoscalingGroupMetricsCollection
type AwsAutoscalingGroupMetricsCollection struct {
	// Required. The frequency at which EC2 Auto Scaling sends aggregated data to
	//  AWS CloudWatch. The only valid value is "1Minute".
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsAutoscalingGroupMetricsCollection.granularity
	Granularity *string `json:"granularity,omitempty"`

	// Optional. The metrics to enable. For a list of valid metrics, see
	//  https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_EnableMetricsCollection.html.
	//  If you specify Granularity and don't specify any metrics, all metrics are
	//  enabled.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsAutoscalingGroupMetricsCollection.metrics
	Metrics []string `json:"metrics,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsConfigEncryption
type AwsConfigEncryption struct {
	// Required. The ARN of the AWS KMS key used to encrypt user data.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsConfigEncryption.kms_key_arn
	KMSKeyArn *string `json:"kmsKeyArn,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsInstancePlacement
type AwsInstancePlacement struct {
	// Required. The tenancy for instance.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsInstancePlacement.tenancy
	Tenancy *string `json:"tenancy,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsNodeConfig
type AwsNodeConfig struct {
	// Optional. The EC2 instance type when creating on-Demand instances.
	//
	//  If unspecified during node pool creation, a default will be chosen based on
	//  the node pool version, and assigned to this field.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.instance_type
	InstanceType *string `json:"instanceType,omitempty"`

	// Optional. Template for the root volume provisioned for node pool nodes.
	//  Volumes will be provisioned in the availability zone assigned
	//  to the node pool subnet.
	//
	//  When unspecified, it defaults to 32 GiB with the GP2 volume type.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.root_volume
	RootVolume *AwsVolumeTemplate `json:"rootVolume,omitempty"`

	// Optional. The initial taints assigned to nodes of this node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.taints
	Taints []NodeTaint `json:"taints,omitempty"`

	// Optional. The initial labels assigned to nodes of this node pool. An object
	//  containing a list of "key": value pairs. Example: { "name": "wrench",
	//  "mass": "1.3kg", "count": "3" }.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Key/value metadata to assign to each underlying AWS resource.
	//  Specify at most 50 pairs containing alphanumerics, spaces, and symbols
	//  (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to
	//  255 Unicode characters.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.tags
	Tags map[string]string `json:"tags,omitempty"`

	// Required. The name or ARN of the AWS IAM instance profile to assign to
	//  nodes in the pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.iam_instance_profile
	IamInstanceProfile *string `json:"iamInstanceProfile,omitempty"`

	// Optional. The OS image type to use on node pool instances.
	//  Can be unspecified, or have a value of `ubuntu`.
	//
	//  When unspecified, it defaults to `ubuntu`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.image_type
	ImageType *string `json:"imageType,omitempty"`

	// Optional. The SSH configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.ssh_config
	SSHConfig *AwsSshConfig `json:"sshConfig,omitempty"`

	// Optional. The IDs of additional security groups to add to nodes in this
	//  pool. The manager will automatically create security groups with minimum
	//  rules needed for a functioning cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.security_group_ids
	SecurityGroupIds []string `json:"securityGroupIds,omitempty"`

	// Optional. Proxy configuration for outbound HTTP(S) traffic.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.proxy_config
	ProxyConfig *AwsProxyConfig `json:"proxyConfig,omitempty"`

	// Required. Config encryption for user data.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.config_encryption
	ConfigEncryption *AwsConfigEncryption `json:"configEncryption,omitempty"`

	// Optional. Placement related info for this node.
	//  When unspecified, the VPC's default tenancy will be used.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.instance_placement
	InstancePlacement *AwsInstancePlacement `json:"instancePlacement,omitempty"`

	// Optional. Configuration related to CloudWatch metrics collection on the
	//  Auto Scaling group of the node pool.
	//
	//  When unspecified, metrics collection is disabled.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.autoscaling_metrics_collection
	AutoscalingMetricsCollection *AwsAutoscalingGroupMetricsCollection `json:"autoscalingMetricsCollection,omitempty"`

	// Optional. Configuration for provisioning EC2 Spot instances
	//
	//  When specified, the node pool will provision Spot instances from the set
	//  of spot_config.instance_types.
	//  This field is mutually exclusive with `instance_type`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeConfig.spot_config
	SpotConfig *SpotConfig `json:"spotConfig,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsNodeManagement
type AwsNodeManagement struct {
	// Optional. Whether or not the nodes will be automatically repaired. When set
	//  to true, the nodes in this node pool will be monitored and if they fail
	//  health checks consistently over a period of time, an automatic repair
	//  action will be triggered to replace them with new nodes.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodeManagement.auto_repair
	AutoRepair *bool `json:"autoRepair,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsNodePool
type AwsNodePool struct {
	// The name of this resource.
	//
	//  Node pool names are formatted as
	//  `projects/<project-number>/locations/<region>/awsClusters/<cluster-id>/awsNodePools/<node-pool-id>`.
	//
	//  For more details on Google Cloud resource names,
	//  see [Resource Names](https://cloud.google.com/apis/design/resource_names)
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.name
	Name *string `json:"name,omitempty"`

	// Required. The Kubernetes version to run on this node pool (e.g.
	//  `1.19.10-gke.1000`).
	//
	//  You can list all supported versions on a given Google Cloud region by
	//  calling
	//  [GetAwsServerConfig][google.cloud.gkemulticloud.v1.AwsClusters.GetAwsServerConfig].
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.version
	Version *string `json:"version,omitempty"`

	// Required. The configuration of the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.config
	Config *AwsNodeConfig `json:"config,omitempty"`

	// Required. Autoscaler configuration for this node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.autoscaling
	Autoscaling *AwsNodePoolAutoscaling `json:"autoscaling,omitempty"`

	// Required. The subnet where the node pool node run.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.subnet_id
	SubnetID *string `json:"subnetID,omitempty"`

	// Allows clients to perform consistent read-modify-writes
	//  through optimistic concurrency control.
	//
	//  Can be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Annotations on the node pool.
	//
	//  This field has the same restrictions as Kubernetes annotations.
	//  The total size of all keys and values combined is limited to 256k.
	//  Key can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Required. The constraint on the maximum number of pods that can be run
	//  simultaneously on a node in the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.max_pods_constraint
	MaxPodsConstraint *MaxPodsConstraint `json:"maxPodsConstraint,omitempty"`

	// Optional. The Management configuration for this node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.management
	Management *AwsNodeManagement `json:"management,omitempty"`

	// Optional. Node kubelet configs.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.kubelet_config
	KubeletConfig *NodeKubeletConfig `json:"kubeletConfig,omitempty"`

	// Optional. Update settings control the speed and disruption of the update.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.update_settings
	UpdateSettings *UpdateSettings `json:"updateSettings,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsNodePoolAutoscaling
type AwsNodePoolAutoscaling struct {
	// Required. Minimum number of nodes in the node pool. Must be greater than or
	//  equal to 1 and less than or equal to max_node_count.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePoolAutoscaling.min_node_count
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	// Required. Maximum number of nodes in the node pool. Must be greater than or
	//  equal to min_node_count and less than or equal to 50.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePoolAutoscaling.max_node_count
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsNodePoolError
type AwsNodePoolError struct {
	// Human-friendly description of the error.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePoolError.message
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsProxyConfig
type AwsProxyConfig struct {
	// The ARN of the AWS Secret Manager secret that contains the HTTP(S) proxy
	//  configuration.
	//
	//  The secret must be a JSON encoded proxy configuration
	//  as described in
	//  https://cloud.google.com/kubernetes-engine/multi-cloud/docs/aws/how-to/use-a-proxy#create_a_proxy_configuration_file
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsProxyConfig.secret_arn
	SecretArn *string `json:"secretArn,omitempty"`

	// The version string of the AWS Secret Manager secret that contains the
	//  HTTP(S) proxy configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsProxyConfig.secret_version
	SecretVersion *string `json:"secretVersion,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsSshConfig
type AwsSshConfig struct {
	// Required. The name of the EC2 key pair used to login into cluster machines.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsSshConfig.ec2_key_pair
	Ec2KeyPair *string `json:"ec2KeyPair,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsVolumeTemplate
type AwsVolumeTemplate struct {
	// Optional. The size of the volume, in GiBs.
	//
	//  When unspecified, a default value is provided. See the specific reference
	//  in the parent resource.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsVolumeTemplate.size_gib
	SizeGib *int32 `json:"sizeGib,omitempty"`

	// Optional. Type of the EBS volume.
	//
	//  When unspecified, it defaults to GP2 volume.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsVolumeTemplate.volume_type
	VolumeType *string `json:"volumeType,omitempty"`

	// Optional. The number of I/O operations per second (IOPS) to provision for
	//  GP3 volume.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsVolumeTemplate.iops
	Iops *int32 `json:"iops,omitempty"`

	// Optional. The throughput that the volume supports, in MiB/s. Only valid if
	//  volume_type is GP3.
	//
	//  If the volume_type is GP3 and this is not speficied, it defaults to 125.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsVolumeTemplate.throughput
	Throughput *int32 `json:"throughput,omitempty"`

	// Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK)
	//  used to encrypt AWS EBS volumes.
	//
	//  If not specified, the default Amazon managed key associated to
	//  the AWS region where this cluster runs will be used.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsVolumeTemplate.kms_key_arn
	KMSKeyArn *string `json:"kmsKeyArn,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.MaxPodsConstraint
type MaxPodsConstraint struct {
	// Required. The maximum number of pods to schedule on a single node.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.MaxPodsConstraint.max_pods_per_node
	MaxPodsPerNode *int64 `json:"maxPodsPerNode,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.NodeKubeletConfig
type NodeKubeletConfig struct {
	// Optional. Enable the insecure kubelet read only port.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.NodeKubeletConfig.insecure_kubelet_readonly_port_enabled
	InsecureKubeletReadonlyPortEnabled *bool `json:"insecureKubeletReadonlyPortEnabled,omitempty"`

	// Optional. Control the CPU management policy on the node.
	//  See
	//  https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/
	//
	//  The following values are allowed.
	//  * "none": the default, which represents the existing scheduling behavior.
	//  * "static": allows pods with certain resource characteristics to be granted
	//  increased CPU affinity and exclusivity on the node.
	//  The default value is 'none' if unspecified.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.NodeKubeletConfig.cpu_manager_policy
	CpuManagerPolicy *string `json:"cpuManagerPolicy,omitempty"`

	// Optional. Enable CPU CFS quota enforcement for containers that specify CPU
	//  limits.
	//
	//  This option is enabled by default which makes kubelet use CFS quota
	//  (https://www.kernel.org/doc/Documentation/scheduler/sched-bwc.txt) to
	//  enforce container CPU limits. Otherwise, CPU limits will not be enforced at
	//  all.
	//
	//  Disable this option to mitigate CPU throttling problems while still having
	//  your pods to be in Guaranteed QoS class by specifying the CPU limits.
	//
	//  The default value is 'true' if unspecified.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.NodeKubeletConfig.cpu_cfs_quota
	CpuCfsQuota *bool `json:"cpuCfsQuota,omitempty"`

	// Optional. Set the CPU CFS quota period value 'cpu.cfs_period_us'.
	//
	//  The string must be a sequence of decimal numbers, each with optional
	//  fraction and a unit suffix, such as "300ms".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	//  The value must be a positive duration.
	//
	//  The default value is '100ms' if unspecified.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.NodeKubeletConfig.cpu_cfs_quota_period
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty"`

	// Optional. Set the Pod PID limits. See
	//  https://kubernetes.io/docs/concepts/policy/pid-limiting/#pod-pid-limits
	//
	//  Controls the maximum number of processes allowed to run in a pod. The value
	//  must be greater than or equal to 1024 and less than 4194304.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.NodeKubeletConfig.pod_pids_limit
	PodPidsLimit *int64 `json:"podPidsLimit,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.NodeTaint
type NodeTaint struct {
	// Required. Key for the taint.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.NodeTaint.key
	Key *string `json:"key,omitempty"`

	// Required. Value for the taint.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.NodeTaint.value
	Value *string `json:"value,omitempty"`

	// Required. The taint effect.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.NodeTaint.effect
	Effect *string `json:"effect,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.SpotConfig
type SpotConfig struct {
	// Required. A list of instance types for creating spot node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.SpotConfig.instance_types
	InstanceTypes []string `json:"instanceTypes,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.SurgeSettings
type SurgeSettings struct {
	// Optional. The maximum number of nodes that can be created beyond the
	//  current size of the node pool during the update process.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.SurgeSettings.max_surge
	MaxSurge *int32 `json:"maxSurge,omitempty"`

	// Optional. The maximum number of nodes that can be simultaneously
	//  unavailable during the update process. A node is considered unavailable if
	//  its status is not Ready.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.SurgeSettings.max_unavailable
	MaxUnavailable *int32 `json:"maxUnavailable,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.UpdateSettings
type UpdateSettings struct {
	// Optional. Settings for surge update.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.UpdateSettings.surge_settings
	SurgeSettings *SurgeSettings `json:"surgeSettings,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsNodePool
type AwsNodePoolObservedState struct {
	// Output only. The lifecycle state of the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.state
	State *string `json:"state,omitempty"`

	// Output only. A globally unique identifier for the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. If set, there are currently changes in flight to the node
	//  pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The time at which this node pool was created.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this node pool was last updated.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A set of errors found in the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.errors
	Errors []AwsNodePoolError `json:"errors,omitempty"`
}
