// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package beta

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLClusterSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Dataproc/Cluster",
			Description: "The Dataproc Cluster resource",
			StructName:  "Cluster",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Cluster",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Cluster",
						Required:    true,
						Description: "A full instance of a Cluster",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Cluster",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Cluster",
						Required:    true,
						Description: "A full instance of a Cluster",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Cluster",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Cluster",
						Required:    true,
						Description: "A full instance of a Cluster",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Cluster",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Cluster",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"Cluster": &dcl.Component{
					Title:           "Cluster",
					ID:              "projects/{{project}}/regions/{{location}}/clusters/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					HasIAM:          true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"project",
							"name",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"clusterUuid": &dcl.Property{
								Type:        "string",
								GoName:      "ClusterUuid",
								ReadOnly:    true,
								Description: "Output only. A cluster UUID (Unique Universal Identifier). Dataproc generates this value when it creates the cluster.",
								Immutable:   true,
							},
							"config": &dcl.Property{
								Type:        "object",
								GoName:      "Config",
								GoType:      "ClusterConfig",
								Description: "Required. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"autoscalingConfig": &dcl.Property{
										Type:        "object",
										GoName:      "AutoscalingConfig",
										GoType:      "ClusterConfigAutoscalingConfig",
										Description: "Optional. Autoscaling config for the policy associated with the cluster. Cluster does not autoscale if this field is unset.",
										Immutable:   true,
										Properties: map[string]*dcl.Property{
											"policy": &dcl.Property{
												Type:        "string",
												GoName:      "Policy",
												Description: "Optional. The autoscaling policy used by the cluster. Only resource names including projectid and location (region) are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` * `projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` Note that the policy must be in the same project and Dataproc region.",
												Immutable:   true,
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Dataproc/AutoscalingPolicy",
														Field:    "name",
													},
												},
											},
										},
									},
									"encryptionConfig": &dcl.Property{
										Type:        "object",
										GoName:      "EncryptionConfig",
										GoType:      "ClusterConfigEncryptionConfig",
										Description: "Optional. Encryption settings for the cluster.",
										Immutable:   true,
										Properties: map[string]*dcl.Property{
											"gcePdKmsKeyName": &dcl.Property{
												Type:        "string",
												GoName:      "GcePdKmsKeyName",
												Description: "Optional. The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.",
												Immutable:   true,
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Cloudkms/CryptoKey",
														Field:    "selfLink",
													},
												},
											},
										},
									},
									"endpointConfig": &dcl.Property{
										Type:          "object",
										GoName:        "EndpointConfig",
										GoType:        "ClusterConfigEndpointConfig",
										Description:   "Optional. Port/endpoint configuration for this cluster",
										Immutable:     true,
										ServerDefault: true,
										Properties: map[string]*dcl.Property{
											"enableHttpPortAccess": &dcl.Property{
												Type:        "boolean",
												GoName:      "EnableHttpPortAccess",
												Description: "Optional. If true, enable http access to specific ports on the cluster from external sources. Defaults to false.",
												Immutable:   true,
											},
											"httpPorts": &dcl.Property{
												Type: "object",
												AdditionalProperties: &dcl.Property{
													Type: "string",
												},
												GoName:      "HttpPorts",
												ReadOnly:    true,
												Description: "Output only. The map of port descriptions to URLs. Will only be populated if enable_http_port_access is true.",
												Immutable:   true,
											},
										},
									},
									"gceClusterConfig": &dcl.Property{
										Type:          "object",
										GoName:        "GceClusterConfig",
										GoType:        "ClusterConfigGceClusterConfig",
										Description:   "Optional. The shared Compute Engine config settings for all instances in a cluster.",
										Immutable:     true,
										ServerDefault: true,
										Properties: map[string]*dcl.Property{
											"internalIPOnly": &dcl.Property{
												Type:          "boolean",
												GoName:        "InternalIPOnly",
												Description:   "Optional. If true, all instances in the cluster will only have internal IP addresses. By default, clusters are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each instance. This `internal_ip_only` restriction can only be enabled for subnetwork enabled networks, and all off-cluster dependencies must be configured to be accessible without external IP addresses.",
												Immutable:     true,
												ServerDefault: true,
											},
											"metadata": &dcl.Property{
												Type: "object",
												AdditionalProperties: &dcl.Property{
													Type: "string",
												},
												GoName:      "Metadata",
												Description: "The Compute Engine metadata entries to add to all instances (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).",
												Immutable:   true,
											},
											"network": &dcl.Property{
												Type:          "string",
												GoName:        "Network",
												Description:   "Optional. The Compute Engine network to be used for machine communications. Cannot be specified with subnetwork_uri. If neither `network_uri` nor `subnetwork_uri` is specified, the \"default\" network of the project is used, if it exists. Cannot be a \"Custom Subnet Network\" (see [Using Subnetworks](https://cloud.google.com/compute/docs/subnetworks) for more information). A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/global/default` * `projects/[project_id]/regions/global/default` * `default`",
												Immutable:     true,
												ServerDefault: true,
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Compute/Network",
														Field:    "selfLink",
													},
												},
											},
											"nodeGroupAffinity": &dcl.Property{
												Type:        "object",
												GoName:      "NodeGroupAffinity",
												GoType:      "ClusterConfigGceClusterConfigNodeGroupAffinity",
												Description: "Optional. Node Group Affinity for sole-tenant clusters.",
												Immutable:   true,
												Required: []string{
													"nodeGroup",
												},
												Properties: map[string]*dcl.Property{
													"nodeGroup": &dcl.Property{
														Type:        "string",
														GoName:      "NodeGroup",
														Description: "Required. The URI of a sole-tenant [node group resource](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups) that the cluster will be created on. A full URL, partial URI, or node group name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` * `projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` * `node-group-1`",
														Immutable:   true,
														ResourceReferences: []*dcl.PropertyResourceReference{
															&dcl.PropertyResourceReference{
																Resource: "Compute/NodeGroup",
																Field:    "selfLink",
															},
														},
													},
												},
											},
											"privateIPv6GoogleAccess": &dcl.Property{
												Type:        "string",
												GoName:      "PrivateIPv6GoogleAccess",
												GoType:      "ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum",
												Description: "Optional. The type of IPv6 access for a cluster. Possible values: PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED, INHERIT_FROM_SUBNETWORK, OUTBOUND, BIDIRECTIONAL",
												Immutable:   true,
												Enum: []string{
													"PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED",
													"INHERIT_FROM_SUBNETWORK",
													"OUTBOUND",
													"BIDIRECTIONAL",
												},
											},
											"reservationAffinity": &dcl.Property{
												Type:        "object",
												GoName:      "ReservationAffinity",
												GoType:      "ClusterConfigGceClusterConfigReservationAffinity",
												Description: "Optional. Reservation Affinity for consuming Zonal reservation.",
												Immutable:   true,
												Properties: map[string]*dcl.Property{
													"consumeReservationType": &dcl.Property{
														Type:        "string",
														GoName:      "ConsumeReservationType",
														GoType:      "ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum",
														Description: "Optional. Type of reservation to consume Possible values: TYPE_UNSPECIFIED, NO_RESERVATION, ANY_RESERVATION, SPECIFIC_RESERVATION",
														Immutable:   true,
														Enum: []string{
															"TYPE_UNSPECIFIED",
															"NO_RESERVATION",
															"ANY_RESERVATION",
															"SPECIFIC_RESERVATION",
														},
													},
													"key": &dcl.Property{
														Type:        "string",
														GoName:      "Key",
														Description: "Optional. Corresponds to the label key of reservation resource.",
														Immutable:   true,
													},
													"values": &dcl.Property{
														Type:        "array",
														GoName:      "Values",
														Description: "Optional. Corresponds to the label values of reservation resource.",
														Immutable:   true,
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
													},
												},
											},
											"serviceAccount": &dcl.Property{
												Type:        "string",
												GoName:      "ServiceAccount",
												Description: "Optional. The [Dataproc service account](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/service-accounts#service_accounts_in_dataproc) (also see [VM Data Plane identity](https://cloud.google.com/dataproc/docs/concepts/iam/dataproc-principals#vm_service_account_data_plane_identity)) used by Dataproc cluster VM instances to access Google Cloud Platform services. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.",
												Immutable:   true,
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Iam/ServiceAccount",
														Field:    "email",
													},
												},
											},
											"serviceAccountScopes": &dcl.Property{
												Type:          "array",
												GoName:        "ServiceAccountScopes",
												Description:   "Optional. The URIs of service account scopes to be included in Compute Engine instances. The following base set of scopes is always included: * https://www.googleapis.com/auth/cloud.useraccounts.readonly * https://www.googleapis.com/auth/devstorage.read_write * https://www.googleapis.com/auth/logging.write If no scopes are specified, the following defaults are also provided: * https://www.googleapis.com/auth/bigquery * https://www.googleapis.com/auth/bigtable.admin.table * https://www.googleapis.com/auth/bigtable.data * https://www.googleapis.com/auth/devstorage.full_control",
												Immutable:     true,
												ServerDefault: true,
												SendEmpty:     true,
												ListType:      "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"subnetwork": &dcl.Property{
												Type:        "string",
												GoName:      "Subnetwork",
												Description: "Optional. The Compute Engine subnetwork to be used for machine communications. Cannot be specified with network_uri. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/us-east1/subnetworks/sub0` * `projects/[project_id]/regions/us-east1/subnetworks/sub0` * `sub0`",
												Immutable:   true,
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Compute/Subnetwork",
														Field:    "selfLink",
													},
												},
											},
											"tags": &dcl.Property{
												Type:        "array",
												GoName:      "Tags",
												Description: "The Compute Engine tags to add to all instances (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).",
												Immutable:   true,
												SendEmpty:   true,
												ListType:    "set",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"zone": &dcl.Property{
												Type:        "string",
												GoName:      "Zone",
												Description: "Optional. The zone where the Compute Engine cluster will be located. On a create request, it is required in the \"global\" region. If omitted in a non-global Dataproc region, the service will pick a zone in the corresponding Compute Engine region. On a get request, zone will always be present. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]` * `projects/[project_id]/zones/[zone]` * `us-central1-f`",
												Immutable:   true,
											},
										},
									},
									"gkeClusterConfig": &dcl.Property{
										Type:        "object",
										GoName:      "GkeClusterConfig",
										GoType:      "ClusterConfigGkeClusterConfig",
										Description: "Optional. BETA. The Kubernetes Engine config for Dataproc clusters deployed to Kubernetes. Setting this is considered mutually exclusive with Compute Engine-based options such as `gce_cluster_config`, `master_config`, `worker_config`, `secondary_worker_config`, and `autoscaling_config`.",
										Immutable:   true,
										Properties: map[string]*dcl.Property{
											"namespacedGkeDeploymentTarget": &dcl.Property{
												Type:        "object",
												GoName:      "NamespacedGkeDeploymentTarget",
												GoType:      "ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget",
												Description: "Optional. A target for the deployment.",
												Immutable:   true,
												Properties: map[string]*dcl.Property{
													"clusterNamespace": &dcl.Property{
														Type:        "string",
														GoName:      "ClusterNamespace",
														Description: "Optional. A namespace within the GKE cluster to deploy into.",
														Immutable:   true,
													},
													"targetGkeCluster": &dcl.Property{
														Type:        "string",
														GoName:      "TargetGkeCluster",
														Description: "Optional. The target GKE cluster to deploy to. Format: 'projects/{project}/locations/{location}/clusters/{cluster_id}'",
														Immutable:   true,
														ResourceReferences: []*dcl.PropertyResourceReference{
															&dcl.PropertyResourceReference{
																Resource: "Container/Cluster",
																Field:    "name",
															},
														},
													},
												},
											},
										},
									},
									"initializationActions": &dcl.Property{
										Type:        "array",
										GoName:      "InitializationActions",
										Description: "Optional. Commands to execute on each node after config is completed. By default, executables are run on master and all worker nodes. You can test a node's `role` metadata to run an executable on a master or worker node, as shown below using `curl` (you can also use `wget`): ROLE=$(curl -H Metadata-Flavor:Google http://metadata/computeMetadata/v1/instance/attributes/dataproc-role) if [[ \"${ROLE}\" == 'Master' ]]; then ... master specific actions ... else ... worker specific actions ... fi",
										Immutable:   true,
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "ClusterConfigInitializationActions",
											Required: []string{
												"executableFile",
											},
											Properties: map[string]*dcl.Property{
												"executableFile": &dcl.Property{
													Type:        "string",
													GoName:      "ExecutableFile",
													Description: "Required. Cloud Storage URI of executable file.",
													Immutable:   true,
												},
												"executionTimeout": &dcl.Property{
													Type:        "string",
													GoName:      "ExecutionTimeout",
													Description: "Optional. Amount of time executable has to complete. Default is 10 minutes (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). Cluster creation fails with an explanatory error message (the name of the executable that caused the error and the exceeded timeout period) if the executable is not completed at end of the timeout period.",
													Immutable:   true,
												},
											},
										},
									},
									"lifecycleConfig": &dcl.Property{
										Type:        "object",
										GoName:      "LifecycleConfig",
										GoType:      "ClusterConfigLifecycleConfig",
										Description: "Optional. Lifecycle setting for the cluster.",
										Immutable:   true,
										Properties: map[string]*dcl.Property{
											"autoDeleteTime": &dcl.Property{
												Type:        "string",
												Format:      "date-time",
												GoName:      "AutoDeleteTime",
												Description: "Optional. The time when cluster will be auto-deleted (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
												Immutable:   true,
											},
											"autoDeleteTtl": &dcl.Property{
												Type:        "string",
												GoName:      "AutoDeleteTtl",
												Description: "Optional. The lifetime duration of cluster. The cluster will be auto-deleted at the end of this period. Minimum value is 10 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
												Immutable:   true,
											},
											"idleDeleteTtl": &dcl.Property{
												Type:        "string",
												GoName:      "IdleDeleteTtl",
												Description: "Optional. The duration to keep the cluster alive while idling (when no jobs are running). Passing this threshold will cause the cluster to be deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
												Immutable:   true,
											},
											"idleStartTime": &dcl.Property{
												Type:        "string",
												Format:      "date-time",
												GoName:      "IdleStartTime",
												ReadOnly:    true,
												Description: "Output only. The time when cluster became idle (most recent job finished) and became eligible for deletion due to idleness (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
												Immutable:   true,
											},
										},
									},
									"masterConfig": &dcl.Property{
										Type:          "object",
										GoName:        "MasterConfig",
										GoType:        "ClusterConfigMasterConfig",
										Description:   "Optional. The Compute Engine config settings for the master instance in a cluster.",
										Immutable:     true,
										ServerDefault: true,
										Properties: map[string]*dcl.Property{
											"accelerators": &dcl.Property{
												Type:          "array",
												GoName:        "Accelerators",
												Description:   "Optional. The Compute Engine accelerator configuration for these instances.",
												Immutable:     true,
												ServerDefault: true,
												SendEmpty:     true,
												ListType:      "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "ClusterConfigMasterConfigAccelerators",
													Properties: map[string]*dcl.Property{
														"acceleratorCount": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "AcceleratorCount",
															Description: "The number of the accelerator cards of this type exposed to this instance.",
															Immutable:   true,
														},
														"acceleratorType": &dcl.Property{
															Type:        "string",
															GoName:      "AcceleratorType",
															Description: "Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes). Examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `nvidia-tesla-k80` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, `nvidia-tesla-k80`.",
															Immutable:   true,
														},
													},
												},
											},
											"diskConfig": &dcl.Property{
												Type:          "object",
												GoName:        "DiskConfig",
												GoType:        "ClusterConfigMasterConfigDiskConfig",
												Description:   "Optional. Disk option config settings.",
												Immutable:     true,
												ServerDefault: true,
												Properties: map[string]*dcl.Property{
													"bootDiskSizeGb": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "BootDiskSizeGb",
														Description: "Optional. Size in GB of the boot disk (default is 500GB).",
														Immutable:   true,
													},
													"bootDiskType": &dcl.Property{
														Type:        "string",
														GoName:      "BootDiskType",
														Description: "Optional. Type of the boot disk (default is \"pd-standard\"). Valid values: \"pd-balanced\" (Persistent Disk Balanced Solid State Drive), \"pd-ssd\" (Persistent Disk Solid State Drive), or \"pd-standard\" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).",
														Immutable:   true,
													},
													"numLocalSsds": &dcl.Property{
														Type:          "integer",
														Format:        "int64",
														GoName:        "NumLocalSsds",
														Description:   "Optional. Number of attached SSDs, from 0 to 4 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries.",
														Immutable:     true,
														ServerDefault: true,
													},
												},
											},
											"image": &dcl.Property{
												Type:        "string",
												GoName:      "Image",
												Description: "Optional. The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/[image-id]` * `projects/[project_id]/global/images/[image-id]` * `image-id` Image family examples. Dataproc will use the most recent image from the family: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/family/[custom-image-family-name]` * `projects/[project_id]/global/images/family/[custom-image-family-name]` If the URI is unspecified, it will be inferred from `SoftwareConfig.image_version` or the system default.",
												Immutable:   true,
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Compute/Image",
														Field:    "selfLink",
													},
												},
											},
											"instanceNames": &dcl.Property{
												Type:          "array",
												GoName:        "InstanceNames",
												ReadOnly:      true,
												Description:   "Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group.",
												Immutable:     true,
												ServerDefault: true,
												ListType:      "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
													ResourceReferences: []*dcl.PropertyResourceReference{
														&dcl.PropertyResourceReference{
															Resource: "Compute/Instance",
															Field:    "selfLink",
														},
													},
												},
											},
											"isPreemptible": &dcl.Property{
												Type:        "boolean",
												GoName:      "IsPreemptible",
												ReadOnly:    true,
												Description: "Output only. Specifies that this instance group contains preemptible instances.",
												Immutable:   true,
											},
											"machineType": &dcl.Property{
												Type:        "string",
												GoName:      "MachineType",
												Description: "Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `n1-standard-2` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`.",
												Immutable:   true,
											},
											"managedGroupConfig": &dcl.Property{
												Type:          "object",
												GoName:        "ManagedGroupConfig",
												GoType:        "ClusterConfigMasterConfigManagedGroupConfig",
												ReadOnly:      true,
												Description:   "Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.",
												Immutable:     true,
												ServerDefault: true,
												Properties: map[string]*dcl.Property{
													"instanceGroupManagerName": &dcl.Property{
														Type:        "string",
														GoName:      "InstanceGroupManagerName",
														ReadOnly:    true,
														Description: "Output only. The name of the Instance Group Manager for this group.",
														Immutable:   true,
													},
													"instanceTemplateName": &dcl.Property{
														Type:        "string",
														GoName:      "InstanceTemplateName",
														ReadOnly:    true,
														Description: "Output only. The name of the Instance Template used for the Managed Instance Group.",
														Immutable:   true,
													},
												},
											},
											"minCpuPlatform": &dcl.Property{
												Type:          "string",
												GoName:        "MinCpuPlatform",
												Description:   "Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).",
												Immutable:     true,
												ServerDefault: true,
											},
											"numInstances": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "NumInstances",
												Description: "Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**.",
												Immutable:   true,
											},
											"preemptibility": &dcl.Property{
												Type:        "string",
												GoName:      "Preemptibility",
												GoType:      "ClusterConfigMasterConfigPreemptibilityEnum",
												Description: "Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE",
												Immutable:   true,
												Enum: []string{
													"PREEMPTIBILITY_UNSPECIFIED",
													"NON_PREEMPTIBLE",
													"PREEMPTIBLE",
												},
											},
										},
									},
									"metastoreConfig": &dcl.Property{
										Type:        "object",
										GoName:      "MetastoreConfig",
										GoType:      "ClusterConfigMetastoreConfig",
										Description: "Optional. Metastore configuration.",
										Immutable:   true,
										Required: []string{
											"dataprocMetastoreService",
										},
										Properties: map[string]*dcl.Property{
											"dataprocMetastoreService": &dcl.Property{
												Type:        "string",
												GoName:      "DataprocMetastoreService",
												Description: "Required. Resource name of an existing Dataproc Metastore service. Example: * `projects/[project_id]/locations/[dataproc_region]/services/[service-name]`",
												Immutable:   true,
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Metastore/Service",
														Field:    "selfLink",
													},
												},
											},
										},
									},
									"secondaryWorkerConfig": &dcl.Property{
										Type:          "object",
										GoName:        "SecondaryWorkerConfig",
										GoType:        "ClusterConfigSecondaryWorkerConfig",
										Description:   "Optional. The Compute Engine config settings for additional worker instances in a cluster.",
										Immutable:     true,
										ServerDefault: true,
										Properties: map[string]*dcl.Property{
											"accelerators": &dcl.Property{
												Type:          "array",
												GoName:        "Accelerators",
												Description:   "Optional. The Compute Engine accelerator configuration for these instances.",
												Immutable:     true,
												ServerDefault: true,
												SendEmpty:     true,
												ListType:      "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "ClusterConfigSecondaryWorkerConfigAccelerators",
													Properties: map[string]*dcl.Property{
														"acceleratorCount": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "AcceleratorCount",
															Description: "The number of the accelerator cards of this type exposed to this instance.",
															Immutable:   true,
														},
														"acceleratorType": &dcl.Property{
															Type:        "string",
															GoName:      "AcceleratorType",
															Description: "Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes). Examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `nvidia-tesla-k80` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, `nvidia-tesla-k80`.",
															Immutable:   true,
														},
													},
												},
											},
											"diskConfig": &dcl.Property{
												Type:          "object",
												GoName:        "DiskConfig",
												GoType:        "ClusterConfigSecondaryWorkerConfigDiskConfig",
												Description:   "Optional. Disk option config settings.",
												Immutable:     true,
												ServerDefault: true,
												Properties: map[string]*dcl.Property{
													"bootDiskSizeGb": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "BootDiskSizeGb",
														Description: "Optional. Size in GB of the boot disk (default is 500GB).",
														Immutable:   true,
													},
													"bootDiskType": &dcl.Property{
														Type:        "string",
														GoName:      "BootDiskType",
														Description: "Optional. Type of the boot disk (default is \"pd-standard\"). Valid values: \"pd-balanced\" (Persistent Disk Balanced Solid State Drive), \"pd-ssd\" (Persistent Disk Solid State Drive), or \"pd-standard\" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).",
														Immutable:   true,
													},
													"numLocalSsds": &dcl.Property{
														Type:          "integer",
														Format:        "int64",
														GoName:        "NumLocalSsds",
														Description:   "Optional. Number of attached SSDs, from 0 to 4 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries.",
														Immutable:     true,
														ServerDefault: true,
													},
												},
											},
											"image": &dcl.Property{
												Type:        "string",
												GoName:      "Image",
												Description: "Optional. The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/[image-id]` * `projects/[project_id]/global/images/[image-id]` * `image-id` Image family examples. Dataproc will use the most recent image from the family: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/family/[custom-image-family-name]` * `projects/[project_id]/global/images/family/[custom-image-family-name]` If the URI is unspecified, it will be inferred from `SoftwareConfig.image_version` or the system default.",
												Immutable:   true,
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Compute/Image",
														Field:    "selfLink",
													},
												},
											},
											"instanceNames": &dcl.Property{
												Type:          "array",
												GoName:        "InstanceNames",
												ReadOnly:      true,
												Description:   "Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group.",
												Immutable:     true,
												ServerDefault: true,
												ListType:      "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
													ResourceReferences: []*dcl.PropertyResourceReference{
														&dcl.PropertyResourceReference{
															Resource: "Compute/Instance",
															Field:    "selfLink",
														},
													},
												},
											},
											"isPreemptible": &dcl.Property{
												Type:        "boolean",
												GoName:      "IsPreemptible",
												ReadOnly:    true,
												Description: "Output only. Specifies that this instance group contains preemptible instances.",
												Immutable:   true,
											},
											"machineType": &dcl.Property{
												Type:        "string",
												GoName:      "MachineType",
												Description: "Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `n1-standard-2` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`.",
												Immutable:   true,
											},
											"managedGroupConfig": &dcl.Property{
												Type:          "object",
												GoName:        "ManagedGroupConfig",
												GoType:        "ClusterConfigSecondaryWorkerConfigManagedGroupConfig",
												ReadOnly:      true,
												Description:   "Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.",
												Immutable:     true,
												ServerDefault: true,
												Properties: map[string]*dcl.Property{
													"instanceGroupManagerName": &dcl.Property{
														Type:        "string",
														GoName:      "InstanceGroupManagerName",
														ReadOnly:    true,
														Description: "Output only. The name of the Instance Group Manager for this group.",
														Immutable:   true,
													},
													"instanceTemplateName": &dcl.Property{
														Type:        "string",
														GoName:      "InstanceTemplateName",
														ReadOnly:    true,
														Description: "Output only. The name of the Instance Template used for the Managed Instance Group.",
														Immutable:   true,
													},
												},
											},
											"minCpuPlatform": &dcl.Property{
												Type:          "string",
												GoName:        "MinCpuPlatform",
												Description:   "Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).",
												Immutable:     true,
												ServerDefault: true,
											},
											"numInstances": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "NumInstances",
												Description: "Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**.",
												Immutable:   true,
											},
											"preemptibility": &dcl.Property{
												Type:        "string",
												GoName:      "Preemptibility",
												GoType:      "ClusterConfigSecondaryWorkerConfigPreemptibilityEnum",
												Description: "Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE",
												Immutable:   true,
												Enum: []string{
													"PREEMPTIBILITY_UNSPECIFIED",
													"NON_PREEMPTIBLE",
													"PREEMPTIBLE",
												},
											},
										},
									},
									"securityConfig": &dcl.Property{
										Type:        "object",
										GoName:      "SecurityConfig",
										GoType:      "ClusterConfigSecurityConfig",
										Description: "Optional. Security settings for the cluster.",
										Immutable:   true,
										Properties: map[string]*dcl.Property{
											"kerberosConfig": &dcl.Property{
												Type:        "object",
												GoName:      "KerberosConfig",
												GoType:      "ClusterConfigSecurityConfigKerberosConfig",
												Description: "Optional. Kerberos related configuration.",
												Immutable:   true,
												Properties: map[string]*dcl.Property{
													"crossRealmTrustAdminServer": &dcl.Property{
														Type:        "string",
														GoName:      "CrossRealmTrustAdminServer",
														Description: "Optional. The admin server (IP or hostname) for the remote trusted realm in a cross realm trust relationship.",
														Immutable:   true,
													},
													"crossRealmTrustKdc": &dcl.Property{
														Type:        "string",
														GoName:      "CrossRealmTrustKdc",
														Description: "Optional. The KDC (IP or hostname) for the remote trusted realm in a cross realm trust relationship.",
														Immutable:   true,
													},
													"crossRealmTrustRealm": &dcl.Property{
														Type:        "string",
														GoName:      "CrossRealmTrustRealm",
														Description: "Optional. The remote realm the Dataproc on-cluster KDC will trust, should the user enable cross realm trust.",
														Immutable:   true,
													},
													"crossRealmTrustSharedPassword": &dcl.Property{
														Type:        "string",
														GoName:      "CrossRealmTrustSharedPassword",
														Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the shared password between the on-cluster Kerberos realm and the remote trusted realm, in a cross realm trust relationship.",
														Immutable:   true,
													},
													"enableKerberos": &dcl.Property{
														Type:        "boolean",
														GoName:      "EnableKerberos",
														Description: "Optional. Flag to indicate whether to Kerberize the cluster (default: false). Set this field to true to enable Kerberos on a cluster.",
														Immutable:   true,
													},
													"kdcDbKey": &dcl.Property{
														Type:        "string",
														GoName:      "KdcDbKey",
														Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the master key of the KDC database.",
														Immutable:   true,
													},
													"keyPassword": &dcl.Property{
														Type:        "string",
														GoName:      "KeyPassword",
														Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided key. For the self-signed certificate, this password is generated by Dataproc.",
														Immutable:   true,
													},
													"keystore": &dcl.Property{
														Type:        "string",
														GoName:      "Keystore",
														Description: "Optional. The Cloud Storage URI of the keystore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.",
														Immutable:   true,
													},
													"keystorePassword": &dcl.Property{
														Type:        "string",
														GoName:      "KeystorePassword",
														Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided keystore. For the self-signed certificate, this password is generated by Dataproc.",
														Immutable:   true,
													},
													"kmsKey": &dcl.Property{
														Type:        "string",
														GoName:      "KmsKey",
														Description: "Optional. The uri of the KMS key used to encrypt various sensitive files.",
														Immutable:   true,
														ResourceReferences: []*dcl.PropertyResourceReference{
															&dcl.PropertyResourceReference{
																Resource: "Cloudkms/CryptoKey",
																Field:    "selfLink",
															},
														},
													},
													"realm": &dcl.Property{
														Type:        "string",
														GoName:      "Realm",
														Description: "Optional. The name of the on-cluster Kerberos realm. If not specified, the uppercased domain of hostnames will be the realm.",
														Immutable:   true,
													},
													"rootPrincipalPassword": &dcl.Property{
														Type:        "string",
														GoName:      "RootPrincipalPassword",
														Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the root principal password.",
														Immutable:   true,
													},
													"tgtLifetimeHours": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "TgtLifetimeHours",
														Description: "Optional. The lifetime of the ticket granting ticket, in hours. If not specified, or user specifies 0, then default value 10 will be used.",
														Immutable:   true,
													},
													"truststore": &dcl.Property{
														Type:        "string",
														GoName:      "Truststore",
														Description: "Optional. The Cloud Storage URI of the truststore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.",
														Immutable:   true,
													},
													"truststorePassword": &dcl.Property{
														Type:        "string",
														GoName:      "TruststorePassword",
														Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided truststore. For the self-signed certificate, this password is generated by Dataproc.",
														Immutable:   true,
													},
												},
											},
										},
									},
									"softwareConfig": &dcl.Property{
										Type:          "object",
										GoName:        "SoftwareConfig",
										GoType:        "ClusterConfigSoftwareConfig",
										Description:   "Optional. The config settings for software inside the cluster.",
										Immutable:     true,
										ServerDefault: true,
										Properties: map[string]*dcl.Property{
											"imageVersion": &dcl.Property{
												Type:        "string",
												GoName:      "ImageVersion",
												Description: "Optional. The version of software inside the cluster. It must be one of the supported [Dataproc Versions](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#supported_dataproc_versions), such as \"1.2\" (including a subminor version, such as \"1.2.29\"), or the [\"preview\" version](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#other_versions). If unspecified, it defaults to the latest Debian version.",
												Immutable:   true,
											},
											"optionalComponents": &dcl.Property{
												Type:        "array",
												GoName:      "OptionalComponents",
												Description: "Optional. The set of components to activate on the cluster.",
												Immutable:   true,
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "ClusterConfigSoftwareConfigOptionalComponentsEnum",
													Enum: []string{
														"COMPONENT_UNSPECIFIED",
														"ANACONDA",
														"DOCKER",
														"DRUID",
														"FLINK",
														"HBASE",
														"HIVE_WEBHCAT",
														"JUPYTER",
														"KERBEROS",
														"PRESTO",
														"RANGER",
														"SOLR",
														"ZEPPELIN",
														"ZOOKEEPER",
													},
												},
											},
											"properties": &dcl.Property{
												Type: "object",
												AdditionalProperties: &dcl.Property{
													Type: "string",
												},
												GoName:      "Properties",
												Description: "Optional. The properties to set on daemon config files. Property keys are specified in `prefix:property` format, for example `core:hadoop.tmp.dir`. The following are supported prefixes and their mappings: * capacity-scheduler: `capacity-scheduler.xml` * core: `core-site.xml` * distcp: `distcp-default.xml` * hdfs: `hdfs-site.xml` * hive: `hive-site.xml` * mapred: `mapred-site.xml` * pig: `pig.properties` * spark: `spark-defaults.conf` * yarn: `yarn-site.xml` For more information, see [Cluster properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties).",
												Immutable:   true,
											},
										},
									},
									"stagingBucket": &dcl.Property{
										Type:          "string",
										GoName:        "StagingBucket",
										Description:   "Optional. A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket (see [Dataproc staging bucket](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)). **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**",
										Immutable:     true,
										ServerDefault: true,
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Storage/Bucket",
												Field:    "name",
											},
										},
									},
									"tempBucket": &dcl.Property{
										Type:          "string",
										GoName:        "TempBucket",
										Description:   "Optional. A Cloud Storage bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. If you do not specify a temp bucket, Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's temp bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket. The default bucket has a TTL of 90 days, but you can use any TTL (or none) if you specify a bucket. **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**",
										Immutable:     true,
										ServerDefault: true,
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Storage/Bucket",
												Field:    "name",
											},
										},
									},
									"workerConfig": &dcl.Property{
										Type:          "object",
										GoName:        "WorkerConfig",
										GoType:        "ClusterConfigWorkerConfig",
										Description:   "Optional. The Compute Engine config settings for worker instances in a cluster.",
										Immutable:     true,
										ServerDefault: true,
										Properties: map[string]*dcl.Property{
											"accelerators": &dcl.Property{
												Type:          "array",
												GoName:        "Accelerators",
												Description:   "Optional. The Compute Engine accelerator configuration for these instances.",
												Immutable:     true,
												ServerDefault: true,
												SendEmpty:     true,
												ListType:      "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "ClusterConfigWorkerConfigAccelerators",
													Properties: map[string]*dcl.Property{
														"acceleratorCount": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "AcceleratorCount",
															Description: "The number of the accelerator cards of this type exposed to this instance.",
															Immutable:   true,
														},
														"acceleratorType": &dcl.Property{
															Type:        "string",
															GoName:      "AcceleratorType",
															Description: "Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes). Examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `nvidia-tesla-k80` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, `nvidia-tesla-k80`.",
															Immutable:   true,
														},
													},
												},
											},
											"diskConfig": &dcl.Property{
												Type:          "object",
												GoName:        "DiskConfig",
												GoType:        "ClusterConfigWorkerConfigDiskConfig",
												Description:   "Optional. Disk option config settings.",
												Immutable:     true,
												ServerDefault: true,
												Properties: map[string]*dcl.Property{
													"bootDiskSizeGb": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "BootDiskSizeGb",
														Description: "Optional. Size in GB of the boot disk (default is 500GB).",
														Immutable:   true,
													},
													"bootDiskType": &dcl.Property{
														Type:        "string",
														GoName:      "BootDiskType",
														Description: "Optional. Type of the boot disk (default is \"pd-standard\"). Valid values: \"pd-balanced\" (Persistent Disk Balanced Solid State Drive), \"pd-ssd\" (Persistent Disk Solid State Drive), or \"pd-standard\" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).",
														Immutable:   true,
													},
													"numLocalSsds": &dcl.Property{
														Type:          "integer",
														Format:        "int64",
														GoName:        "NumLocalSsds",
														Description:   "Optional. Number of attached SSDs, from 0 to 4 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries.",
														Immutable:     true,
														ServerDefault: true,
													},
												},
											},
											"image": &dcl.Property{
												Type:        "string",
												GoName:      "Image",
												Description: "Optional. The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/[image-id]` * `projects/[project_id]/global/images/[image-id]` * `image-id` Image family examples. Dataproc will use the most recent image from the family: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/family/[custom-image-family-name]` * `projects/[project_id]/global/images/family/[custom-image-family-name]` If the URI is unspecified, it will be inferred from `SoftwareConfig.image_version` or the system default.",
												Immutable:   true,
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Compute/Image",
														Field:    "selfLink",
													},
												},
											},
											"instanceNames": &dcl.Property{
												Type:          "array",
												GoName:        "InstanceNames",
												ReadOnly:      true,
												Description:   "Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group.",
												Immutable:     true,
												ServerDefault: true,
												ListType:      "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
													ResourceReferences: []*dcl.PropertyResourceReference{
														&dcl.PropertyResourceReference{
															Resource: "Compute/Instance",
															Field:    "selfLink",
														},
													},
												},
											},
											"isPreemptible": &dcl.Property{
												Type:        "boolean",
												GoName:      "IsPreemptible",
												ReadOnly:    true,
												Description: "Output only. Specifies that this instance group contains preemptible instances.",
												Immutable:   true,
											},
											"machineType": &dcl.Property{
												Type:        "string",
												GoName:      "MachineType",
												Description: "Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `n1-standard-2` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`.",
												Immutable:   true,
											},
											"managedGroupConfig": &dcl.Property{
												Type:          "object",
												GoName:        "ManagedGroupConfig",
												GoType:        "ClusterConfigWorkerConfigManagedGroupConfig",
												ReadOnly:      true,
												Description:   "Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.",
												Immutable:     true,
												ServerDefault: true,
												Properties: map[string]*dcl.Property{
													"instanceGroupManagerName": &dcl.Property{
														Type:        "string",
														GoName:      "InstanceGroupManagerName",
														ReadOnly:    true,
														Description: "Output only. The name of the Instance Group Manager for this group.",
														Immutable:   true,
													},
													"instanceTemplateName": &dcl.Property{
														Type:        "string",
														GoName:      "InstanceTemplateName",
														ReadOnly:    true,
														Description: "Output only. The name of the Instance Template used for the Managed Instance Group.",
														Immutable:   true,
													},
												},
											},
											"minCpuPlatform": &dcl.Property{
												Type:          "string",
												GoName:        "MinCpuPlatform",
												Description:   "Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).",
												Immutable:     true,
												ServerDefault: true,
											},
											"numInstances": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "NumInstances",
												Description: "Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**.",
												Immutable:   true,
											},
											"preemptibility": &dcl.Property{
												Type:        "string",
												GoName:      "Preemptibility",
												GoType:      "ClusterConfigWorkerConfigPreemptibilityEnum",
												Description: "Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE",
												Immutable:   true,
												Enum: []string{
													"PREEMPTIBILITY_UNSPECIFIED",
													"NON_PREEMPTIBLE",
													"PREEMPTIBLE",
												},
											},
										},
									},
								},
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Optional. The labels to associate with this cluster. Label **keys** must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.",
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource, usually a GCP region.",
								Immutable:   true,
							},
							"metrics": &dcl.Property{
								Type:        "object",
								GoName:      "Metrics",
								GoType:      "ClusterMetrics",
								ReadOnly:    true,
								Description: "Output only. Contains cluster daemon metrics such as HDFS and YARN stats. **Beta Feature**: This report is available for testing purposes only. It may be changed before final release.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"hdfsMetrics": &dcl.Property{
										Type: "object",
										AdditionalProperties: &dcl.Property{
											Type: "string",
										},
										GoName:      "HdfsMetrics",
										Description: "The HDFS metrics.",
										Immutable:   true,
									},
									"yarnMetrics": &dcl.Property{
										Type: "object",
										AdditionalProperties: &dcl.Property{
											Type: "string",
										},
										GoName:      "YarnMetrics",
										Description: "The YARN metrics.",
										Immutable:   true,
									},
								},
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Required. The cluster name. Cluster names within a project must be unique. Names of deleted clusters can be reused.",
								Immutable:   true,
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "Required. The Google Cloud Platform project ID that the cluster belongs to.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"status": &dcl.Property{
								Type:        "object",
								GoName:      "Status",
								GoType:      "ClusterStatus",
								ReadOnly:    true,
								Description: "Output only. Cluster status.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"detail": &dcl.Property{
										Type:        "string",
										GoName:      "Detail",
										ReadOnly:    true,
										Description: "Optional. Output only. Details of cluster's state.",
										Immutable:   true,
									},
									"state": &dcl.Property{
										Type:        "string",
										GoName:      "State",
										GoType:      "ClusterStatusStateEnum",
										ReadOnly:    true,
										Description: "Output only. The cluster's state. Possible values: UNKNOWN, CREATING, RUNNING, ERROR, DELETING, UPDATING, STOPPING, STOPPED, STARTING",
										Immutable:   true,
										Enum: []string{
											"UNKNOWN",
											"CREATING",
											"RUNNING",
											"ERROR",
											"DELETING",
											"UPDATING",
											"STOPPING",
											"STOPPED",
											"STARTING",
										},
									},
									"stateStartTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "StateStartTime",
										ReadOnly:    true,
										Description: "Output only. Time when this state was entered (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
										Immutable:   true,
									},
									"substate": &dcl.Property{
										Type:        "string",
										GoName:      "Substate",
										GoType:      "ClusterStatusSubstateEnum",
										ReadOnly:    true,
										Description: "Output only. Additional state information that includes status reported by the agent. Possible values: UNSPECIFIED, UNHEALTHY, STALE_STATUS",
										Immutable:   true,
										Enum: []string{
											"UNSPECIFIED",
											"UNHEALTHY",
											"STALE_STATUS",
										},
									},
								},
							},
							"statusHistory": &dcl.Property{
								Type:        "array",
								GoName:      "StatusHistory",
								ReadOnly:    true,
								Description: "Output only. The previous cluster status.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "ClusterStatusHistory",
									Properties: map[string]*dcl.Property{
										"detail": &dcl.Property{
											Type:        "string",
											GoName:      "Detail",
											ReadOnly:    true,
											Description: "Optional. Output only. Details of cluster's state.",
											Immutable:   true,
										},
										"state": &dcl.Property{
											Type:        "string",
											GoName:      "State",
											GoType:      "ClusterStatusHistoryStateEnum",
											ReadOnly:    true,
											Description: "Output only. The cluster's state. Possible values: UNKNOWN, CREATING, RUNNING, ERROR, DELETING, UPDATING, STOPPING, STOPPED, STARTING",
											Immutable:   true,
											Enum: []string{
												"UNKNOWN",
												"CREATING",
												"RUNNING",
												"ERROR",
												"DELETING",
												"UPDATING",
												"STOPPING",
												"STOPPED",
												"STARTING",
											},
										},
										"stateStartTime": &dcl.Property{
											Type:        "string",
											Format:      "date-time",
											GoName:      "StateStartTime",
											ReadOnly:    true,
											Description: "Output only. Time when this state was entered (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
											Immutable:   true,
										},
										"substate": &dcl.Property{
											Type:        "string",
											GoName:      "Substate",
											GoType:      "ClusterStatusHistorySubstateEnum",
											ReadOnly:    true,
											Description: "Output only. Additional state information that includes status reported by the agent. Possible values: UNSPECIFIED, UNHEALTHY, STALE_STATUS",
											Immutable:   true,
											Enum: []string{
												"UNSPECIFIED",
												"UNHEALTHY",
												"STALE_STATUS",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
