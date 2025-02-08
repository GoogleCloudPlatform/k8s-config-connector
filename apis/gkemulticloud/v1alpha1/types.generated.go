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


// +kcc:proto=google.cloud.gkemulticloud.v1.AwsAuthorization
type AwsAuthorization struct {
	// Optional. Users that can perform operations as a cluster admin. A managed
	//  ClusterRoleBinding will be created to grant the `cluster-admin` ClusterRole
	//  to the users. Up to ten admin users can be provided.
	//
	//  For more info on RBAC, see
	//  https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsAuthorization.admin_users
	AdminUsers []AwsClusterUser `json:"adminUsers,omitempty"`

	// Optional. Groups of users that can perform operations as a cluster admin. A
	//  managed ClusterRoleBinding will be created to grant the `cluster-admin`
	//  ClusterRole to the groups. Up to ten admin groups can be provided.
	//
	//  For more info on RBAC, see
	//  https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsAuthorization.admin_groups
	AdminGroups []AwsClusterGroup `json:"adminGroups,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsCluster
type AwsCluster struct {
	// The name of this resource.
	//
	//  Cluster names are formatted as
	//  `projects/<project-number>/locations/<region>/awsClusters/<cluster-id>`.
	//
	//  See [Resource Names](https://cloud.google.com/apis/design/resource_names)
	//  for more details on Google Cloud Platform resource names.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.name
	Name *string `json:"name,omitempty"`

	// Optional. A human readable description of this cluster.
	//  Cannot be longer than 255 UTF-8 encoded bytes.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.description
	Description *string `json:"description,omitempty"`

	// Required. Cluster-wide networking configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.networking
	Networking *AwsClusterNetworking `json:"networking,omitempty"`

	// Required. The AWS region where the cluster runs.
	//
	//  Each Google Cloud region supports a subset of nearby AWS regions.
	//  You can call
	//  [GetAwsServerConfig][google.cloud.gkemulticloud.v1.AwsClusters.GetAwsServerConfig]
	//  to list all supported AWS regions within a given Google Cloud region.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.aws_region
	AwsRegion *string `json:"awsRegion,omitempty"`

	// Required. Configuration related to the cluster control plane.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.control_plane
	ControlPlane *AwsControlPlane `json:"controlPlane,omitempty"`

	// Required. Configuration related to the cluster RBAC settings.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.authorization
	Authorization *AwsAuthorization `json:"authorization,omitempty"`

	// Allows clients to perform consistent read-modify-writes
	//  through optimistic concurrency control.
	//
	//  Can be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Annotations on the cluster.
	//
	//  This field has the same restrictions as Kubernetes annotations.
	//  The total size of all keys and values combined is limited to 256k.
	//  Key can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Required. Fleet configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.fleet
	Fleet *Fleet `json:"fleet,omitempty"`

	// Optional. Logging configuration for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Optional. Monitoring configuration for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.monitoring_config
	MonitoringConfig *MonitoringConfig `json:"monitoringConfig,omitempty"`

	// Optional. Binary Authorization configuration for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.binary_authorization
	BinaryAuthorization *BinaryAuthorization `json:"binaryAuthorization,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsClusterError
type AwsClusterError struct {
	// Human-friendly description of the error.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsClusterError.message
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsClusterGroup
type AwsClusterGroup struct {
	// Required. The name of the group, e.g. `my-group@domain.com`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsClusterGroup.group
	Group *string `json:"group,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsClusterNetworking
type AwsClusterNetworking struct {
	// Required. The VPC associated with the cluster. All component clusters
	//  (i.e. control plane and node pools) run on a single VPC.
	//
	//  This field cannot be changed after creation.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsClusterNetworking.vpc_id
	VpcID *string `json:"vpcID,omitempty"`

	// Required. All pods in the cluster are assigned an IPv4 address from these
	//  ranges. Only a single range is supported. This field cannot be changed
	//  after creation.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsClusterNetworking.pod_address_cidr_blocks
	PodAddressCidrBlocks []string `json:"podAddressCidrBlocks,omitempty"`

	// Required. All services in the cluster are assigned an IPv4 address from
	//  these ranges. Only a single range is supported. This field cannot be
	//  changed after creation.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsClusterNetworking.service_address_cidr_blocks
	ServiceAddressCidrBlocks []string `json:"serviceAddressCidrBlocks,omitempty"`

	// Optional. Disable the per node pool subnet security group rules on the
	//  control plane security group. When set to true, you must also provide one
	//  or more security groups that ensure node pools are able to send requests to
	//  the control plane on TCP/443 and TCP/8132. Failure to do so may result in
	//  unavailable node pools.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsClusterNetworking.per_node_pool_sg_rules_disabled
	PerNodePoolSgRulesDisabled *bool `json:"perNodePoolSgRulesDisabled,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsClusterUser
type AwsClusterUser struct {
	// Required. The name of the user, e.g. `my-gcp-id@gmail.com`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsClusterUser.username
	Username *string `json:"username,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsConfigEncryption
type AwsConfigEncryption struct {
	// Required. The ARN of the AWS KMS key used to encrypt user data.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsConfigEncryption.kms_key_arn
	KMSKeyArn *string `json:"kmsKeyArn,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsControlPlane
type AwsControlPlane struct {
	// Required. The Kubernetes version to run on control plane replicas
	//  (e.g. `1.19.10-gke.1000`).
	//
	//  You can list all supported versions on a given Google Cloud region by
	//  calling
	//  [GetAwsServerConfig][google.cloud.gkemulticloud.v1.AwsClusters.GetAwsServerConfig].
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.version
	Version *string `json:"version,omitempty"`

	// Optional. The AWS instance type.
	//
	//  When unspecified, it uses a default based on the cluster's version.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.instance_type
	InstanceType *string `json:"instanceType,omitempty"`

	// Optional. SSH configuration for how to access the underlying control plane
	//  machines.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.ssh_config
	SSHConfig *AwsSshConfig `json:"sshConfig,omitempty"`

	// Required. The list of subnets where control plane replicas will run.
	//  A replica will be provisioned on each subnet and up to three values
	//  can be provided.
	//  Each subnet must be in a different AWS Availability Zone (AZ).
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.subnet_ids
	SubnetIds []string `json:"subnetIds,omitempty"`

	// Optional. The IDs of additional security groups to add to control plane
	//  replicas. The Anthos Multi-Cloud API will automatically create and manage
	//  security groups with the minimum rules needed for a functioning cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.security_group_ids
	SecurityGroupIds []string `json:"securityGroupIds,omitempty"`

	// Required. The name or ARN of the AWS IAM instance profile to assign to each
	//  control plane replica.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.iam_instance_profile
	IamInstanceProfile *string `json:"iamInstanceProfile,omitempty"`

	// Optional. Configuration related to the root volume provisioned for each
	//  control plane replica.
	//
	//  Volumes will be provisioned in the availability zone associated
	//  with the corresponding subnet.
	//
	//  When unspecified, it defaults to 32 GiB with the GP2 volume type.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.root_volume
	RootVolume *AwsVolumeTemplate `json:"rootVolume,omitempty"`

	// Optional. Configuration related to the main volume provisioned for each
	//  control plane replica.
	//  The main volume is in charge of storing all of the cluster's etcd state.
	//
	//  Volumes will be provisioned in the availability zone associated
	//  with the corresponding subnet.
	//
	//  When unspecified, it defaults to 8 GiB with the GP2 volume type.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.main_volume
	MainVolume *AwsVolumeTemplate `json:"mainVolume,omitempty"`

	// Required. The ARN of the AWS KMS key used to encrypt cluster secrets.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.database_encryption
	DatabaseEncryption *AwsDatabaseEncryption `json:"databaseEncryption,omitempty"`

	// Optional. A set of AWS resource tags to propagate to all underlying managed
	//  AWS resources.
	//
	//  Specify at most 50 pairs containing alphanumerics, spaces, and symbols
	//  (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to
	//  255 Unicode characters.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.tags
	Tags map[string]string `json:"tags,omitempty"`

	// Required. Authentication configuration for management of AWS resources.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.aws_services_authentication
	AwsServicesAuthentication *AwsServicesAuthentication `json:"awsServicesAuthentication,omitempty"`

	// Optional. Proxy configuration for outbound HTTP(S) traffic.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.proxy_config
	ProxyConfig *AwsProxyConfig `json:"proxyConfig,omitempty"`

	// Required. Config encryption for user data.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.config_encryption
	ConfigEncryption *AwsConfigEncryption `json:"configEncryption,omitempty"`

	// Optional. The placement to use on control plane instances.
	//  When unspecified, the VPC's default tenancy will be used.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsControlPlane.instance_placement
	InstancePlacement *AwsInstancePlacement `json:"instancePlacement,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsDatabaseEncryption
type AwsDatabaseEncryption struct {
	// Required. The ARN of the AWS KMS key used to encrypt cluster secrets.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsDatabaseEncryption.kms_key_arn
	KMSKeyArn *string `json:"kmsKeyArn,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsInstancePlacement
type AwsInstancePlacement struct {
	// Required. The tenancy for instance.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsInstancePlacement.tenancy
	Tenancy *string `json:"tenancy,omitempty"`
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

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsServicesAuthentication
type AwsServicesAuthentication struct {
	// Required. The Amazon Resource Name (ARN) of the role that the Anthos
	//  Multi-Cloud API will assume when managing AWS resources on your account.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsServicesAuthentication.role_arn
	RoleArn *string `json:"roleArn,omitempty"`

	// Optional. An identifier for the assumed role session.
	//
	//  When unspecified, it defaults to `multicloud-service-agent`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsServicesAuthentication.role_session_name
	RoleSessionName *string `json:"roleSessionName,omitempty"`
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

// +kcc:proto=google.cloud.gkemulticloud.v1.BinaryAuthorization
type BinaryAuthorization struct {
	// Mode of operation for binauthz policy evaluation. If unspecified, defaults
	//  to DISABLED.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.BinaryAuthorization.evaluation_mode
	EvaluationMode *string `json:"evaluationMode,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.CloudMonitoringConfig
type CloudMonitoringConfig struct {
	// Enable GKE-native logging and metrics.
	//  Only for Attached Clusters.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.CloudMonitoringConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.Fleet
type Fleet struct {
	// Required. The name of the Fleet host project where this cluster will be
	//  registered.
	//
	//  Project names are formatted as
	//  `projects/<project-number>`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.Fleet.project
	Project *string `json:"project,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.LoggingComponentConfig
type LoggingComponentConfig struct {
	// The components to be enabled.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.LoggingComponentConfig.enable_components
	EnableComponents []string `json:"enableComponents,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.LoggingConfig
type LoggingConfig struct {
	// The configuration of the logging components;
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.LoggingConfig.component_config
	ComponentConfig *LoggingComponentConfig `json:"componentConfig,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.ManagedPrometheusConfig
type ManagedPrometheusConfig struct {
	// Enable Managed Collection.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.ManagedPrometheusConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.MonitoringConfig
type MonitoringConfig struct {
	// Enable Google Cloud Managed Service for Prometheus in the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.MonitoringConfig.managed_prometheus_config
	ManagedPrometheusConfig *ManagedPrometheusConfig `json:"managedPrometheusConfig,omitempty"`

	// Optionally enable GKE metrics.
	//  Only for Attached Clusters.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.MonitoringConfig.cloud_monitoring_config
	CloudMonitoringConfig *CloudMonitoringConfig `json:"cloudMonitoringConfig,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig
type WorkloadIdentityConfig struct {
	// The OIDC issuer URL for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig.issuer_uri
	IssuerURI *string `json:"issuerURI,omitempty"`

	// The Workload Identity Pool associated to the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig.workload_pool
	WorkloadPool *string `json:"workloadPool,omitempty"`

	// The ID of the OIDC Identity Provider (IdP) associated to the Workload
	//  Identity Pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig.identity_provider
	IdentityProvider *string `json:"identityProvider,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsCluster
type AwsClusterObservedState struct {
	// Output only. The current state of the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.state
	State *string `json:"state,omitempty"`

	// Output only. The endpoint of the cluster's API server.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.endpoint
	Endpoint *string `json:"endpoint,omitempty"`

	// Output only. A globally unique identifier for the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. If set, there are currently changes in flight to the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The time at which this cluster was created.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this cluster was last updated.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Workload Identity settings.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.workload_identity_config
	WorkloadIdentityConfig *WorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty"`

	// Output only. PEM encoded x509 certificate of the cluster root of trust.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.cluster_ca_certificate
	ClusterCaCertificate *string `json:"clusterCaCertificate,omitempty"`

	// Required. Fleet configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.fleet
	Fleet *FleetObservedState `json:"fleet,omitempty"`

	// Output only. A set of errors found in the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsCluster.errors
	Errors []AwsClusterError `json:"errors,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.Fleet
type FleetObservedState struct {
	// Output only. The name of the managed Hub Membership resource associated to
	//  this cluster.
	//
	//  Membership names are formatted as
	//  `projects/<project-number>/locations/global/membership/<cluster-id>`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.Fleet.membership
	Membership *string `json:"membership,omitempty"`
}
