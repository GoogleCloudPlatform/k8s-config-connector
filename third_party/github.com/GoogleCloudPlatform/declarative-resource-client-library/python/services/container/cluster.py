# Copyright 2021 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from connector import channel
from google3.cloud.graphite.mmv2.services.google.container import cluster_pb2
from google3.cloud.graphite.mmv2.services.google.container import cluster_pb2_grpc

from typing import List


class Cluster(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        initial_node_count: int = None,
        master_auth: dict = None,
        logging_service: str = None,
        monitoring_service: str = None,
        network: str = None,
        cluster_ipv4_cidr: str = None,
        addons_config: dict = None,
        subnetwork: str = None,
        node_pools: list = None,
        locations: list = None,
        enable_kubernetes_alpha: bool = None,
        resource_labels: dict = None,
        label_fingerprint: str = None,
        legacy_abac: dict = None,
        network_policy: dict = None,
        ip_allocation_policy: dict = None,
        master_authorized_networks_config: dict = None,
        binary_authorization: dict = None,
        autoscaling: dict = None,
        network_config: dict = None,
        maintenance_policy: dict = None,
        default_max_pods_constraint: dict = None,
        resource_usage_export_config: dict = None,
        authenticator_groups_config: dict = None,
        private_cluster_config: dict = None,
        database_encryption: dict = None,
        vertical_pod_autoscaling: dict = None,
        shielded_nodes: dict = None,
        endpoint: str = None,
        master_version: str = None,
        create_time: str = None,
        status: str = None,
        status_message: str = None,
        node_ipv4_cidr_size: int = None,
        services_ipv4_cidr: str = None,
        expire_time: str = None,
        location: str = None,
        enable_tpu: bool = None,
        tpu_ipv4_cidr_block: str = None,
        conditions: list = None,
        autopilot: dict = None,
        project: str = None,
        node_config: dict = None,
        release_channel: dict = None,
        workload_identity_config: dict = None,
        notification_config: dict = None,
        confidential_nodes: dict = None,
        self_link: str = None,
        zone: str = None,
        initial_cluster_version: str = None,
        current_master_version: str = None,
        current_node_version: str = None,
        instance_group_urls: list = None,
        current_node_count: int = None,
        id: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.initial_node_count = initial_node_count
        self.master_auth = master_auth
        self.logging_service = logging_service
        self.monitoring_service = monitoring_service
        self.network = network
        self.cluster_ipv4_cidr = cluster_ipv4_cidr
        self.addons_config = addons_config
        self.subnetwork = subnetwork
        self.node_pools = node_pools
        self.locations = locations
        self.enable_kubernetes_alpha = enable_kubernetes_alpha
        self.resource_labels = resource_labels
        self.label_fingerprint = label_fingerprint
        self.legacy_abac = legacy_abac
        self.network_policy = network_policy
        self.ip_allocation_policy = ip_allocation_policy
        self.master_authorized_networks_config = master_authorized_networks_config
        self.binary_authorization = binary_authorization
        self.autoscaling = autoscaling
        self.network_config = network_config
        self.maintenance_policy = maintenance_policy
        self.default_max_pods_constraint = default_max_pods_constraint
        self.resource_usage_export_config = resource_usage_export_config
        self.authenticator_groups_config = authenticator_groups_config
        self.private_cluster_config = private_cluster_config
        self.database_encryption = database_encryption
        self.vertical_pod_autoscaling = vertical_pod_autoscaling
        self.shielded_nodes = shielded_nodes
        self.master_version = master_version
        self.location = location
        self.enable_tpu = enable_tpu
        self.conditions = conditions
        self.autopilot = autopilot
        self.project = project
        self.node_config = node_config
        self.release_channel = release_channel
        self.workload_identity_config = workload_identity_config
        self.notification_config = notification_config
        self.confidential_nodes = confidential_nodes
        self.initial_cluster_version = initial_cluster_version
        self.instance_group_urls = instance_group_urls
        self.service_account_file = service_account_file

    def apply(self):
        stub = cluster_pb2_grpc.ContainerClusterServiceStub(channel.Channel())
        request = cluster_pb2.ApplyContainerClusterRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.initial_node_count):
            request.resource.initial_node_count = Primitive.to_proto(
                self.initial_node_count
            )

        if ClusterMasterAuth.to_proto(self.master_auth):
            request.resource.master_auth.CopyFrom(
                ClusterMasterAuth.to_proto(self.master_auth)
            )
        else:
            request.resource.ClearField("master_auth")
        if Primitive.to_proto(self.logging_service):
            request.resource.logging_service = Primitive.to_proto(self.logging_service)

        if Primitive.to_proto(self.monitoring_service):
            request.resource.monitoring_service = Primitive.to_proto(
                self.monitoring_service
            )

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.cluster_ipv4_cidr):
            request.resource.cluster_ipv4_cidr = Primitive.to_proto(
                self.cluster_ipv4_cidr
            )

        if ClusterAddonsConfig.to_proto(self.addons_config):
            request.resource.addons_config.CopyFrom(
                ClusterAddonsConfig.to_proto(self.addons_config)
            )
        else:
            request.resource.ClearField("addons_config")
        if Primitive.to_proto(self.subnetwork):
            request.resource.subnetwork = Primitive.to_proto(self.subnetwork)

        if ClusterNodePoolsArray.to_proto(self.node_pools):
            request.resource.node_pools.extend(
                ClusterNodePoolsArray.to_proto(self.node_pools)
            )
        if Primitive.to_proto(self.locations):
            request.resource.locations.extend(Primitive.to_proto(self.locations))
        if Primitive.to_proto(self.enable_kubernetes_alpha):
            request.resource.enable_kubernetes_alpha = Primitive.to_proto(
                self.enable_kubernetes_alpha
            )

        if Primitive.to_proto(self.resource_labels):
            request.resource.resource_labels = Primitive.to_proto(self.resource_labels)

        if Primitive.to_proto(self.label_fingerprint):
            request.resource.label_fingerprint = Primitive.to_proto(
                self.label_fingerprint
            )

        if ClusterLegacyAbac.to_proto(self.legacy_abac):
            request.resource.legacy_abac.CopyFrom(
                ClusterLegacyAbac.to_proto(self.legacy_abac)
            )
        else:
            request.resource.ClearField("legacy_abac")
        if ClusterNetworkPolicy.to_proto(self.network_policy):
            request.resource.network_policy.CopyFrom(
                ClusterNetworkPolicy.to_proto(self.network_policy)
            )
        else:
            request.resource.ClearField("network_policy")
        if ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy):
            request.resource.ip_allocation_policy.CopyFrom(
                ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy)
            )
        else:
            request.resource.ClearField("ip_allocation_policy")
        if ClusterMasterAuthorizedNetworksConfig.to_proto(
            self.master_authorized_networks_config
        ):
            request.resource.master_authorized_networks_config.CopyFrom(
                ClusterMasterAuthorizedNetworksConfig.to_proto(
                    self.master_authorized_networks_config
                )
            )
        else:
            request.resource.ClearField("master_authorized_networks_config")
        if ClusterBinaryAuthorization.to_proto(self.binary_authorization):
            request.resource.binary_authorization.CopyFrom(
                ClusterBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            request.resource.ClearField("binary_authorization")
        if ClusterAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                ClusterAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if ClusterNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                ClusterNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if ClusterMaintenancePolicy.to_proto(self.maintenance_policy):
            request.resource.maintenance_policy.CopyFrom(
                ClusterMaintenancePolicy.to_proto(self.maintenance_policy)
            )
        else:
            request.resource.ClearField("maintenance_policy")
        if ClusterDefaultMaxPodsConstraint.to_proto(self.default_max_pods_constraint):
            request.resource.default_max_pods_constraint.CopyFrom(
                ClusterDefaultMaxPodsConstraint.to_proto(
                    self.default_max_pods_constraint
                )
            )
        else:
            request.resource.ClearField("default_max_pods_constraint")
        if ClusterResourceUsageExportConfig.to_proto(self.resource_usage_export_config):
            request.resource.resource_usage_export_config.CopyFrom(
                ClusterResourceUsageExportConfig.to_proto(
                    self.resource_usage_export_config
                )
            )
        else:
            request.resource.ClearField("resource_usage_export_config")
        if ClusterAuthenticatorGroupsConfig.to_proto(self.authenticator_groups_config):
            request.resource.authenticator_groups_config.CopyFrom(
                ClusterAuthenticatorGroupsConfig.to_proto(
                    self.authenticator_groups_config
                )
            )
        else:
            request.resource.ClearField("authenticator_groups_config")
        if ClusterPrivateClusterConfig.to_proto(self.private_cluster_config):
            request.resource.private_cluster_config.CopyFrom(
                ClusterPrivateClusterConfig.to_proto(self.private_cluster_config)
            )
        else:
            request.resource.ClearField("private_cluster_config")
        if ClusterDatabaseEncryption.to_proto(self.database_encryption):
            request.resource.database_encryption.CopyFrom(
                ClusterDatabaseEncryption.to_proto(self.database_encryption)
            )
        else:
            request.resource.ClearField("database_encryption")
        if ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling):
            request.resource.vertical_pod_autoscaling.CopyFrom(
                ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling)
            )
        else:
            request.resource.ClearField("vertical_pod_autoscaling")
        if ClusterShieldedNodes.to_proto(self.shielded_nodes):
            request.resource.shielded_nodes.CopyFrom(
                ClusterShieldedNodes.to_proto(self.shielded_nodes)
            )
        else:
            request.resource.ClearField("shielded_nodes")
        if Primitive.to_proto(self.master_version):
            request.resource.master_version = Primitive.to_proto(self.master_version)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.enable_tpu):
            request.resource.enable_tpu = Primitive.to_proto(self.enable_tpu)

        if ClusterConditionsArray.to_proto(self.conditions):
            request.resource.conditions.extend(
                ClusterConditionsArray.to_proto(self.conditions)
            )
        if ClusterAutopilot.to_proto(self.autopilot):
            request.resource.autopilot.CopyFrom(
                ClusterAutopilot.to_proto(self.autopilot)
            )
        else:
            request.resource.ClearField("autopilot")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if ClusterNodeConfig.to_proto(self.node_config):
            request.resource.node_config.CopyFrom(
                ClusterNodeConfig.to_proto(self.node_config)
            )
        else:
            request.resource.ClearField("node_config")
        if ClusterReleaseChannel.to_proto(self.release_channel):
            request.resource.release_channel.CopyFrom(
                ClusterReleaseChannel.to_proto(self.release_channel)
            )
        else:
            request.resource.ClearField("release_channel")
        if ClusterWorkloadIdentityConfig.to_proto(self.workload_identity_config):
            request.resource.workload_identity_config.CopyFrom(
                ClusterWorkloadIdentityConfig.to_proto(self.workload_identity_config)
            )
        else:
            request.resource.ClearField("workload_identity_config")
        if ClusterNotificationConfig.to_proto(self.notification_config):
            request.resource.notification_config.CopyFrom(
                ClusterNotificationConfig.to_proto(self.notification_config)
            )
        else:
            request.resource.ClearField("notification_config")
        if ClusterConfidentialNodes.to_proto(self.confidential_nodes):
            request.resource.confidential_nodes.CopyFrom(
                ClusterConfidentialNodes.to_proto(self.confidential_nodes)
            )
        else:
            request.resource.ClearField("confidential_nodes")
        if Primitive.to_proto(self.initial_cluster_version):
            request.resource.initial_cluster_version = Primitive.to_proto(
                self.initial_cluster_version
            )

        if Primitive.to_proto(self.instance_group_urls):
            request.resource.instance_group_urls.extend(
                Primitive.to_proto(self.instance_group_urls)
            )
        request.service_account_file = self.service_account_file

        response = stub.ApplyContainerCluster(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.initial_node_count = Primitive.from_proto(response.initial_node_count)
        self.master_auth = ClusterMasterAuth.from_proto(response.master_auth)
        self.logging_service = Primitive.from_proto(response.logging_service)
        self.monitoring_service = Primitive.from_proto(response.monitoring_service)
        self.network = Primitive.from_proto(response.network)
        self.cluster_ipv4_cidr = Primitive.from_proto(response.cluster_ipv4_cidr)
        self.addons_config = ClusterAddonsConfig.from_proto(response.addons_config)
        self.subnetwork = Primitive.from_proto(response.subnetwork)
        self.node_pools = ClusterNodePoolsArray.from_proto(response.node_pools)
        self.locations = Primitive.from_proto(response.locations)
        self.enable_kubernetes_alpha = Primitive.from_proto(
            response.enable_kubernetes_alpha
        )
        self.resource_labels = Primitive.from_proto(response.resource_labels)
        self.label_fingerprint = Primitive.from_proto(response.label_fingerprint)
        self.legacy_abac = ClusterLegacyAbac.from_proto(response.legacy_abac)
        self.network_policy = ClusterNetworkPolicy.from_proto(response.network_policy)
        self.ip_allocation_policy = ClusterIPAllocationPolicy.from_proto(
            response.ip_allocation_policy
        )
        self.master_authorized_networks_config = ClusterMasterAuthorizedNetworksConfig.from_proto(
            response.master_authorized_networks_config
        )
        self.binary_authorization = ClusterBinaryAuthorization.from_proto(
            response.binary_authorization
        )
        self.autoscaling = ClusterAutoscaling.from_proto(response.autoscaling)
        self.network_config = ClusterNetworkConfig.from_proto(response.network_config)
        self.maintenance_policy = ClusterMaintenancePolicy.from_proto(
            response.maintenance_policy
        )
        self.default_max_pods_constraint = ClusterDefaultMaxPodsConstraint.from_proto(
            response.default_max_pods_constraint
        )
        self.resource_usage_export_config = ClusterResourceUsageExportConfig.from_proto(
            response.resource_usage_export_config
        )
        self.authenticator_groups_config = ClusterAuthenticatorGroupsConfig.from_proto(
            response.authenticator_groups_config
        )
        self.private_cluster_config = ClusterPrivateClusterConfig.from_proto(
            response.private_cluster_config
        )
        self.database_encryption = ClusterDatabaseEncryption.from_proto(
            response.database_encryption
        )
        self.vertical_pod_autoscaling = ClusterVerticalPodAutoscaling.from_proto(
            response.vertical_pod_autoscaling
        )
        self.shielded_nodes = ClusterShieldedNodes.from_proto(response.shielded_nodes)
        self.endpoint = Primitive.from_proto(response.endpoint)
        self.master_version = Primitive.from_proto(response.master_version)
        self.create_time = Primitive.from_proto(response.create_time)
        self.status = Primitive.from_proto(response.status)
        self.status_message = Primitive.from_proto(response.status_message)
        self.node_ipv4_cidr_size = Primitive.from_proto(response.node_ipv4_cidr_size)
        self.services_ipv4_cidr = Primitive.from_proto(response.services_ipv4_cidr)
        self.expire_time = Primitive.from_proto(response.expire_time)
        self.location = Primitive.from_proto(response.location)
        self.enable_tpu = Primitive.from_proto(response.enable_tpu)
        self.tpu_ipv4_cidr_block = Primitive.from_proto(response.tpu_ipv4_cidr_block)
        self.conditions = ClusterConditionsArray.from_proto(response.conditions)
        self.autopilot = ClusterAutopilot.from_proto(response.autopilot)
        self.project = Primitive.from_proto(response.project)
        self.node_config = ClusterNodeConfig.from_proto(response.node_config)
        self.release_channel = ClusterReleaseChannel.from_proto(
            response.release_channel
        )
        self.workload_identity_config = ClusterWorkloadIdentityConfig.from_proto(
            response.workload_identity_config
        )
        self.notification_config = ClusterNotificationConfig.from_proto(
            response.notification_config
        )
        self.confidential_nodes = ClusterConfidentialNodes.from_proto(
            response.confidential_nodes
        )
        self.self_link = Primitive.from_proto(response.self_link)
        self.zone = Primitive.from_proto(response.zone)
        self.initial_cluster_version = Primitive.from_proto(
            response.initial_cluster_version
        )
        self.current_master_version = Primitive.from_proto(
            response.current_master_version
        )
        self.current_node_version = Primitive.from_proto(response.current_node_version)
        self.instance_group_urls = Primitive.from_proto(response.instance_group_urls)
        self.current_node_count = Primitive.from_proto(response.current_node_count)
        self.id = Primitive.from_proto(response.id)

    def delete(self):
        stub = cluster_pb2_grpc.ContainerClusterServiceStub(channel.Channel())
        request = cluster_pb2.DeleteContainerClusterRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.initial_node_count):
            request.resource.initial_node_count = Primitive.to_proto(
                self.initial_node_count
            )

        if ClusterMasterAuth.to_proto(self.master_auth):
            request.resource.master_auth.CopyFrom(
                ClusterMasterAuth.to_proto(self.master_auth)
            )
        else:
            request.resource.ClearField("master_auth")
        if Primitive.to_proto(self.logging_service):
            request.resource.logging_service = Primitive.to_proto(self.logging_service)

        if Primitive.to_proto(self.monitoring_service):
            request.resource.monitoring_service = Primitive.to_proto(
                self.monitoring_service
            )

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.cluster_ipv4_cidr):
            request.resource.cluster_ipv4_cidr = Primitive.to_proto(
                self.cluster_ipv4_cidr
            )

        if ClusterAddonsConfig.to_proto(self.addons_config):
            request.resource.addons_config.CopyFrom(
                ClusterAddonsConfig.to_proto(self.addons_config)
            )
        else:
            request.resource.ClearField("addons_config")
        if Primitive.to_proto(self.subnetwork):
            request.resource.subnetwork = Primitive.to_proto(self.subnetwork)

        if ClusterNodePoolsArray.to_proto(self.node_pools):
            request.resource.node_pools.extend(
                ClusterNodePoolsArray.to_proto(self.node_pools)
            )
        if Primitive.to_proto(self.locations):
            request.resource.locations.extend(Primitive.to_proto(self.locations))
        if Primitive.to_proto(self.enable_kubernetes_alpha):
            request.resource.enable_kubernetes_alpha = Primitive.to_proto(
                self.enable_kubernetes_alpha
            )

        if Primitive.to_proto(self.resource_labels):
            request.resource.resource_labels = Primitive.to_proto(self.resource_labels)

        if Primitive.to_proto(self.label_fingerprint):
            request.resource.label_fingerprint = Primitive.to_proto(
                self.label_fingerprint
            )

        if ClusterLegacyAbac.to_proto(self.legacy_abac):
            request.resource.legacy_abac.CopyFrom(
                ClusterLegacyAbac.to_proto(self.legacy_abac)
            )
        else:
            request.resource.ClearField("legacy_abac")
        if ClusterNetworkPolicy.to_proto(self.network_policy):
            request.resource.network_policy.CopyFrom(
                ClusterNetworkPolicy.to_proto(self.network_policy)
            )
        else:
            request.resource.ClearField("network_policy")
        if ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy):
            request.resource.ip_allocation_policy.CopyFrom(
                ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy)
            )
        else:
            request.resource.ClearField("ip_allocation_policy")
        if ClusterMasterAuthorizedNetworksConfig.to_proto(
            self.master_authorized_networks_config
        ):
            request.resource.master_authorized_networks_config.CopyFrom(
                ClusterMasterAuthorizedNetworksConfig.to_proto(
                    self.master_authorized_networks_config
                )
            )
        else:
            request.resource.ClearField("master_authorized_networks_config")
        if ClusterBinaryAuthorization.to_proto(self.binary_authorization):
            request.resource.binary_authorization.CopyFrom(
                ClusterBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            request.resource.ClearField("binary_authorization")
        if ClusterAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                ClusterAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if ClusterNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                ClusterNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if ClusterMaintenancePolicy.to_proto(self.maintenance_policy):
            request.resource.maintenance_policy.CopyFrom(
                ClusterMaintenancePolicy.to_proto(self.maintenance_policy)
            )
        else:
            request.resource.ClearField("maintenance_policy")
        if ClusterDefaultMaxPodsConstraint.to_proto(self.default_max_pods_constraint):
            request.resource.default_max_pods_constraint.CopyFrom(
                ClusterDefaultMaxPodsConstraint.to_proto(
                    self.default_max_pods_constraint
                )
            )
        else:
            request.resource.ClearField("default_max_pods_constraint")
        if ClusterResourceUsageExportConfig.to_proto(self.resource_usage_export_config):
            request.resource.resource_usage_export_config.CopyFrom(
                ClusterResourceUsageExportConfig.to_proto(
                    self.resource_usage_export_config
                )
            )
        else:
            request.resource.ClearField("resource_usage_export_config")
        if ClusterAuthenticatorGroupsConfig.to_proto(self.authenticator_groups_config):
            request.resource.authenticator_groups_config.CopyFrom(
                ClusterAuthenticatorGroupsConfig.to_proto(
                    self.authenticator_groups_config
                )
            )
        else:
            request.resource.ClearField("authenticator_groups_config")
        if ClusterPrivateClusterConfig.to_proto(self.private_cluster_config):
            request.resource.private_cluster_config.CopyFrom(
                ClusterPrivateClusterConfig.to_proto(self.private_cluster_config)
            )
        else:
            request.resource.ClearField("private_cluster_config")
        if ClusterDatabaseEncryption.to_proto(self.database_encryption):
            request.resource.database_encryption.CopyFrom(
                ClusterDatabaseEncryption.to_proto(self.database_encryption)
            )
        else:
            request.resource.ClearField("database_encryption")
        if ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling):
            request.resource.vertical_pod_autoscaling.CopyFrom(
                ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling)
            )
        else:
            request.resource.ClearField("vertical_pod_autoscaling")
        if ClusterShieldedNodes.to_proto(self.shielded_nodes):
            request.resource.shielded_nodes.CopyFrom(
                ClusterShieldedNodes.to_proto(self.shielded_nodes)
            )
        else:
            request.resource.ClearField("shielded_nodes")
        if Primitive.to_proto(self.master_version):
            request.resource.master_version = Primitive.to_proto(self.master_version)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.enable_tpu):
            request.resource.enable_tpu = Primitive.to_proto(self.enable_tpu)

        if ClusterConditionsArray.to_proto(self.conditions):
            request.resource.conditions.extend(
                ClusterConditionsArray.to_proto(self.conditions)
            )
        if ClusterAutopilot.to_proto(self.autopilot):
            request.resource.autopilot.CopyFrom(
                ClusterAutopilot.to_proto(self.autopilot)
            )
        else:
            request.resource.ClearField("autopilot")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if ClusterNodeConfig.to_proto(self.node_config):
            request.resource.node_config.CopyFrom(
                ClusterNodeConfig.to_proto(self.node_config)
            )
        else:
            request.resource.ClearField("node_config")
        if ClusterReleaseChannel.to_proto(self.release_channel):
            request.resource.release_channel.CopyFrom(
                ClusterReleaseChannel.to_proto(self.release_channel)
            )
        else:
            request.resource.ClearField("release_channel")
        if ClusterWorkloadIdentityConfig.to_proto(self.workload_identity_config):
            request.resource.workload_identity_config.CopyFrom(
                ClusterWorkloadIdentityConfig.to_proto(self.workload_identity_config)
            )
        else:
            request.resource.ClearField("workload_identity_config")
        if ClusterNotificationConfig.to_proto(self.notification_config):
            request.resource.notification_config.CopyFrom(
                ClusterNotificationConfig.to_proto(self.notification_config)
            )
        else:
            request.resource.ClearField("notification_config")
        if ClusterConfidentialNodes.to_proto(self.confidential_nodes):
            request.resource.confidential_nodes.CopyFrom(
                ClusterConfidentialNodes.to_proto(self.confidential_nodes)
            )
        else:
            request.resource.ClearField("confidential_nodes")
        if Primitive.to_proto(self.initial_cluster_version):
            request.resource.initial_cluster_version = Primitive.to_proto(
                self.initial_cluster_version
            )

        if Primitive.to_proto(self.instance_group_urls):
            request.resource.instance_group_urls.extend(
                Primitive.to_proto(self.instance_group_urls)
            )
        response = stub.DeleteContainerCluster(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = cluster_pb2_grpc.ContainerClusterServiceStub(channel.Channel())
        request = cluster_pb2.ListContainerClusterRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListContainerCluster(request).items

    def to_proto(self):
        resource = cluster_pb2.ContainerCluster()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.initial_node_count):
            resource.initial_node_count = Primitive.to_proto(self.initial_node_count)
        if ClusterMasterAuth.to_proto(self.master_auth):
            resource.master_auth.CopyFrom(ClusterMasterAuth.to_proto(self.master_auth))
        else:
            resource.ClearField("master_auth")
        if Primitive.to_proto(self.logging_service):
            resource.logging_service = Primitive.to_proto(self.logging_service)
        if Primitive.to_proto(self.monitoring_service):
            resource.monitoring_service = Primitive.to_proto(self.monitoring_service)
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if Primitive.to_proto(self.cluster_ipv4_cidr):
            resource.cluster_ipv4_cidr = Primitive.to_proto(self.cluster_ipv4_cidr)
        if ClusterAddonsConfig.to_proto(self.addons_config):
            resource.addons_config.CopyFrom(
                ClusterAddonsConfig.to_proto(self.addons_config)
            )
        else:
            resource.ClearField("addons_config")
        if Primitive.to_proto(self.subnetwork):
            resource.subnetwork = Primitive.to_proto(self.subnetwork)
        if ClusterNodePoolsArray.to_proto(self.node_pools):
            resource.node_pools.extend(ClusterNodePoolsArray.to_proto(self.node_pools))
        if Primitive.to_proto(self.locations):
            resource.locations.extend(Primitive.to_proto(self.locations))
        if Primitive.to_proto(self.enable_kubernetes_alpha):
            resource.enable_kubernetes_alpha = Primitive.to_proto(
                self.enable_kubernetes_alpha
            )
        if Primitive.to_proto(self.resource_labels):
            resource.resource_labels = Primitive.to_proto(self.resource_labels)
        if Primitive.to_proto(self.label_fingerprint):
            resource.label_fingerprint = Primitive.to_proto(self.label_fingerprint)
        if ClusterLegacyAbac.to_proto(self.legacy_abac):
            resource.legacy_abac.CopyFrom(ClusterLegacyAbac.to_proto(self.legacy_abac))
        else:
            resource.ClearField("legacy_abac")
        if ClusterNetworkPolicy.to_proto(self.network_policy):
            resource.network_policy.CopyFrom(
                ClusterNetworkPolicy.to_proto(self.network_policy)
            )
        else:
            resource.ClearField("network_policy")
        if ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy):
            resource.ip_allocation_policy.CopyFrom(
                ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy)
            )
        else:
            resource.ClearField("ip_allocation_policy")
        if ClusterMasterAuthorizedNetworksConfig.to_proto(
            self.master_authorized_networks_config
        ):
            resource.master_authorized_networks_config.CopyFrom(
                ClusterMasterAuthorizedNetworksConfig.to_proto(
                    self.master_authorized_networks_config
                )
            )
        else:
            resource.ClearField("master_authorized_networks_config")
        if ClusterBinaryAuthorization.to_proto(self.binary_authorization):
            resource.binary_authorization.CopyFrom(
                ClusterBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            resource.ClearField("binary_authorization")
        if ClusterAutoscaling.to_proto(self.autoscaling):
            resource.autoscaling.CopyFrom(ClusterAutoscaling.to_proto(self.autoscaling))
        else:
            resource.ClearField("autoscaling")
        if ClusterNetworkConfig.to_proto(self.network_config):
            resource.network_config.CopyFrom(
                ClusterNetworkConfig.to_proto(self.network_config)
            )
        else:
            resource.ClearField("network_config")
        if ClusterMaintenancePolicy.to_proto(self.maintenance_policy):
            resource.maintenance_policy.CopyFrom(
                ClusterMaintenancePolicy.to_proto(self.maintenance_policy)
            )
        else:
            resource.ClearField("maintenance_policy")
        if ClusterDefaultMaxPodsConstraint.to_proto(self.default_max_pods_constraint):
            resource.default_max_pods_constraint.CopyFrom(
                ClusterDefaultMaxPodsConstraint.to_proto(
                    self.default_max_pods_constraint
                )
            )
        else:
            resource.ClearField("default_max_pods_constraint")
        if ClusterResourceUsageExportConfig.to_proto(self.resource_usage_export_config):
            resource.resource_usage_export_config.CopyFrom(
                ClusterResourceUsageExportConfig.to_proto(
                    self.resource_usage_export_config
                )
            )
        else:
            resource.ClearField("resource_usage_export_config")
        if ClusterAuthenticatorGroupsConfig.to_proto(self.authenticator_groups_config):
            resource.authenticator_groups_config.CopyFrom(
                ClusterAuthenticatorGroupsConfig.to_proto(
                    self.authenticator_groups_config
                )
            )
        else:
            resource.ClearField("authenticator_groups_config")
        if ClusterPrivateClusterConfig.to_proto(self.private_cluster_config):
            resource.private_cluster_config.CopyFrom(
                ClusterPrivateClusterConfig.to_proto(self.private_cluster_config)
            )
        else:
            resource.ClearField("private_cluster_config")
        if ClusterDatabaseEncryption.to_proto(self.database_encryption):
            resource.database_encryption.CopyFrom(
                ClusterDatabaseEncryption.to_proto(self.database_encryption)
            )
        else:
            resource.ClearField("database_encryption")
        if ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling):
            resource.vertical_pod_autoscaling.CopyFrom(
                ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling)
            )
        else:
            resource.ClearField("vertical_pod_autoscaling")
        if ClusterShieldedNodes.to_proto(self.shielded_nodes):
            resource.shielded_nodes.CopyFrom(
                ClusterShieldedNodes.to_proto(self.shielded_nodes)
            )
        else:
            resource.ClearField("shielded_nodes")
        if Primitive.to_proto(self.master_version):
            resource.master_version = Primitive.to_proto(self.master_version)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.enable_tpu):
            resource.enable_tpu = Primitive.to_proto(self.enable_tpu)
        if ClusterConditionsArray.to_proto(self.conditions):
            resource.conditions.extend(ClusterConditionsArray.to_proto(self.conditions))
        if ClusterAutopilot.to_proto(self.autopilot):
            resource.autopilot.CopyFrom(ClusterAutopilot.to_proto(self.autopilot))
        else:
            resource.ClearField("autopilot")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if ClusterNodeConfig.to_proto(self.node_config):
            resource.node_config.CopyFrom(ClusterNodeConfig.to_proto(self.node_config))
        else:
            resource.ClearField("node_config")
        if ClusterReleaseChannel.to_proto(self.release_channel):
            resource.release_channel.CopyFrom(
                ClusterReleaseChannel.to_proto(self.release_channel)
            )
        else:
            resource.ClearField("release_channel")
        if ClusterWorkloadIdentityConfig.to_proto(self.workload_identity_config):
            resource.workload_identity_config.CopyFrom(
                ClusterWorkloadIdentityConfig.to_proto(self.workload_identity_config)
            )
        else:
            resource.ClearField("workload_identity_config")
        if ClusterNotificationConfig.to_proto(self.notification_config):
            resource.notification_config.CopyFrom(
                ClusterNotificationConfig.to_proto(self.notification_config)
            )
        else:
            resource.ClearField("notification_config")
        if ClusterConfidentialNodes.to_proto(self.confidential_nodes):
            resource.confidential_nodes.CopyFrom(
                ClusterConfidentialNodes.to_proto(self.confidential_nodes)
            )
        else:
            resource.ClearField("confidential_nodes")
        if Primitive.to_proto(self.initial_cluster_version):
            resource.initial_cluster_version = Primitive.to_proto(
                self.initial_cluster_version
            )
        if Primitive.to_proto(self.instance_group_urls):
            resource.instance_group_urls.extend(
                Primitive.to_proto(self.instance_group_urls)
            )
        return resource


class ClusterMasterAuth(object):
    def __init__(
        self,
        username: str = None,
        password: str = None,
        client_certificate_config: dict = None,
        cluster_ca_certificate: str = None,
        client_certificate: str = None,
        client_key: str = None,
    ):
        self.username = username
        self.password = password
        self.client_certificate_config = client_certificate_config
        self.cluster_ca_certificate = cluster_ca_certificate
        self.client_certificate = client_certificate
        self.client_key = client_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMasterAuth()
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        if Primitive.to_proto(resource.password):
            res.password = Primitive.to_proto(resource.password)
        if ClusterMasterAuthClientCertificateConfig.to_proto(
            resource.client_certificate_config
        ):
            res.client_certificate_config.CopyFrom(
                ClusterMasterAuthClientCertificateConfig.to_proto(
                    resource.client_certificate_config
                )
            )
        else:
            res.ClearField("client_certificate_config")
        if Primitive.to_proto(resource.cluster_ca_certificate):
            res.cluster_ca_certificate = Primitive.to_proto(
                resource.cluster_ca_certificate
            )
        if Primitive.to_proto(resource.client_certificate):
            res.client_certificate = Primitive.to_proto(resource.client_certificate)
        if Primitive.to_proto(resource.client_key):
            res.client_key = Primitive.to_proto(resource.client_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMasterAuth(
            username=Primitive.from_proto(resource.username),
            password=Primitive.from_proto(resource.password),
            client_certificate_config=ClusterMasterAuthClientCertificateConfig.from_proto(
                resource.client_certificate_config
            ),
            cluster_ca_certificate=Primitive.from_proto(
                resource.cluster_ca_certificate
            ),
            client_certificate=Primitive.from_proto(resource.client_certificate),
            client_key=Primitive.from_proto(resource.client_key),
        )


class ClusterMasterAuthArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMasterAuth.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterMasterAuth.from_proto(i) for i in resources]


class ClusterMasterAuthClientCertificateConfig(object):
    def __init__(self, issue_client_certificate: bool = None):
        self.issue_client_certificate = issue_client_certificate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMasterAuthClientCertificateConfig()
        if Primitive.to_proto(resource.issue_client_certificate):
            res.issue_client_certificate = Primitive.to_proto(
                resource.issue_client_certificate
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMasterAuthClientCertificateConfig(
            issue_client_certificate=Primitive.from_proto(
                resource.issue_client_certificate
            ),
        )


class ClusterMasterAuthClientCertificateConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMasterAuthClientCertificateConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterMasterAuthClientCertificateConfig.from_proto(i) for i in resources
        ]


class ClusterAddonsConfig(object):
    def __init__(
        self,
        http_load_balancing: dict = None,
        horizontal_pod_autoscaling: dict = None,
        kubernetes_dashboard: dict = None,
        network_policy_config: dict = None,
        cloud_run_config: dict = None,
        dns_cache_config: dict = None,
        config_connector_config: dict = None,
        gce_persistent_disk_csi_driver_config: dict = None,
    ):
        self.http_load_balancing = http_load_balancing
        self.horizontal_pod_autoscaling = horizontal_pod_autoscaling
        self.kubernetes_dashboard = kubernetes_dashboard
        self.network_policy_config = network_policy_config
        self.cloud_run_config = cloud_run_config
        self.dns_cache_config = dns_cache_config
        self.config_connector_config = config_connector_config
        self.gce_persistent_disk_csi_driver_config = (
            gce_persistent_disk_csi_driver_config
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfig()
        if ClusterAddonsConfigHttpLoadBalancing.to_proto(resource.http_load_balancing):
            res.http_load_balancing.CopyFrom(
                ClusterAddonsConfigHttpLoadBalancing.to_proto(
                    resource.http_load_balancing
                )
            )
        else:
            res.ClearField("http_load_balancing")
        if ClusterAddonsConfigHorizontalPodAutoscaling.to_proto(
            resource.horizontal_pod_autoscaling
        ):
            res.horizontal_pod_autoscaling.CopyFrom(
                ClusterAddonsConfigHorizontalPodAutoscaling.to_proto(
                    resource.horizontal_pod_autoscaling
                )
            )
        else:
            res.ClearField("horizontal_pod_autoscaling")
        if ClusterAddonsConfigKubernetesDashboard.to_proto(
            resource.kubernetes_dashboard
        ):
            res.kubernetes_dashboard.CopyFrom(
                ClusterAddonsConfigKubernetesDashboard.to_proto(
                    resource.kubernetes_dashboard
                )
            )
        else:
            res.ClearField("kubernetes_dashboard")
        if ClusterAddonsConfigNetworkPolicyConfig.to_proto(
            resource.network_policy_config
        ):
            res.network_policy_config.CopyFrom(
                ClusterAddonsConfigNetworkPolicyConfig.to_proto(
                    resource.network_policy_config
                )
            )
        else:
            res.ClearField("network_policy_config")
        if ClusterAddonsConfigCloudRunConfig.to_proto(resource.cloud_run_config):
            res.cloud_run_config.CopyFrom(
                ClusterAddonsConfigCloudRunConfig.to_proto(resource.cloud_run_config)
            )
        else:
            res.ClearField("cloud_run_config")
        if ClusterAddonsConfigDnsCacheConfig.to_proto(resource.dns_cache_config):
            res.dns_cache_config.CopyFrom(
                ClusterAddonsConfigDnsCacheConfig.to_proto(resource.dns_cache_config)
            )
        else:
            res.ClearField("dns_cache_config")
        if ClusterAddonsConfigConfigConnectorConfig.to_proto(
            resource.config_connector_config
        ):
            res.config_connector_config.CopyFrom(
                ClusterAddonsConfigConfigConnectorConfig.to_proto(
                    resource.config_connector_config
                )
            )
        else:
            res.ClearField("config_connector_config")
        if ClusterAddonsConfigGcePersistentDiskCsiDriverConfig.to_proto(
            resource.gce_persistent_disk_csi_driver_config
        ):
            res.gce_persistent_disk_csi_driver_config.CopyFrom(
                ClusterAddonsConfigGcePersistentDiskCsiDriverConfig.to_proto(
                    resource.gce_persistent_disk_csi_driver_config
                )
            )
        else:
            res.ClearField("gce_persistent_disk_csi_driver_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfig(
            http_load_balancing=ClusterAddonsConfigHttpLoadBalancing.from_proto(
                resource.http_load_balancing
            ),
            horizontal_pod_autoscaling=ClusterAddonsConfigHorizontalPodAutoscaling.from_proto(
                resource.horizontal_pod_autoscaling
            ),
            kubernetes_dashboard=ClusterAddonsConfigKubernetesDashboard.from_proto(
                resource.kubernetes_dashboard
            ),
            network_policy_config=ClusterAddonsConfigNetworkPolicyConfig.from_proto(
                resource.network_policy_config
            ),
            cloud_run_config=ClusterAddonsConfigCloudRunConfig.from_proto(
                resource.cloud_run_config
            ),
            dns_cache_config=ClusterAddonsConfigDnsCacheConfig.from_proto(
                resource.dns_cache_config
            ),
            config_connector_config=ClusterAddonsConfigConfigConnectorConfig.from_proto(
                resource.config_connector_config
            ),
            gce_persistent_disk_csi_driver_config=ClusterAddonsConfigGcePersistentDiskCsiDriverConfig.from_proto(
                resource.gce_persistent_disk_csi_driver_config
            ),
        )


class ClusterAddonsConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAddonsConfig.from_proto(i) for i in resources]


class ClusterAddonsConfigHttpLoadBalancing(object):
    def __init__(self, disabled: bool = None):
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigHttpLoadBalancing()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigHttpLoadBalancing(
            disabled=Primitive.from_proto(resource.disabled),
        )


class ClusterAddonsConfigHttpLoadBalancingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfigHttpLoadBalancing.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAddonsConfigHttpLoadBalancing.from_proto(i) for i in resources]


class ClusterAddonsConfigHorizontalPodAutoscaling(object):
    def __init__(self, disabled: bool = None):
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigHorizontalPodAutoscaling()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigHorizontalPodAutoscaling(
            disabled=Primitive.from_proto(resource.disabled),
        )


class ClusterAddonsConfigHorizontalPodAutoscalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterAddonsConfigHorizontalPodAutoscaling.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAddonsConfigHorizontalPodAutoscaling.from_proto(i) for i in resources
        ]


class ClusterAddonsConfigKubernetesDashboard(object):
    def __init__(self, disabled: bool = None):
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigKubernetesDashboard()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigKubernetesDashboard(
            disabled=Primitive.from_proto(resource.disabled),
        )


class ClusterAddonsConfigKubernetesDashboardArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfigKubernetesDashboard.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAddonsConfigKubernetesDashboard.from_proto(i) for i in resources]


class ClusterAddonsConfigNetworkPolicyConfig(object):
    def __init__(self, disabled: bool = None):
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigNetworkPolicyConfig()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigNetworkPolicyConfig(
            disabled=Primitive.from_proto(resource.disabled),
        )


class ClusterAddonsConfigNetworkPolicyConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfigNetworkPolicyConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAddonsConfigNetworkPolicyConfig.from_proto(i) for i in resources]


class ClusterAddonsConfigCloudRunConfig(object):
    def __init__(self, disabled: bool = None, load_balancer_type: str = None):
        self.disabled = disabled
        self.load_balancer_type = load_balancer_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigCloudRunConfig()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        if ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum.to_proto(
            resource.load_balancer_type
        ):
            res.load_balancer_type = ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum.to_proto(
                resource.load_balancer_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigCloudRunConfig(
            disabled=Primitive.from_proto(resource.disabled),
            load_balancer_type=ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum.from_proto(
                resource.load_balancer_type
            ),
        )


class ClusterAddonsConfigCloudRunConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfigCloudRunConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAddonsConfigCloudRunConfig.from_proto(i) for i in resources]


class ClusterAddonsConfigDnsCacheConfig(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigDnsCacheConfig()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigDnsCacheConfig(
            enabled=Primitive.from_proto(resource.enabled),
        )


class ClusterAddonsConfigDnsCacheConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfigDnsCacheConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAddonsConfigDnsCacheConfig.from_proto(i) for i in resources]


class ClusterAddonsConfigConfigConnectorConfig(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigConfigConnectorConfig()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigConfigConnectorConfig(
            enabled=Primitive.from_proto(resource.enabled),
        )


class ClusterAddonsConfigConfigConnectorConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfigConfigConnectorConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAddonsConfigConfigConnectorConfig.from_proto(i) for i in resources
        ]


class ClusterAddonsConfigGcePersistentDiskCsiDriverConfig(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigGcePersistentDiskCsiDriverConfig(
            enabled=Primitive.from_proto(resource.enabled),
        )


class ClusterAddonsConfigGcePersistentDiskCsiDriverConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterAddonsConfigGcePersistentDiskCsiDriverConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAddonsConfigGcePersistentDiskCsiDriverConfig.from_proto(i)
            for i in resources
        ]


class ClusterNodePools(object):
    def __init__(
        self,
        name: str = None,
        config: dict = None,
        initial_node_count: int = None,
        locations: list = None,
        self_link: str = None,
        version: str = None,
        instance_group_urls: list = None,
        status: str = None,
        status_message: str = None,
        autoscaling: dict = None,
        management: dict = None,
        max_pods_constraint: dict = None,
        conditions: list = None,
        pod_ipv4_cidr_size: int = None,
        upgrade_settings: dict = None,
    ):
        self.name = name
        self.config = config
        self.initial_node_count = initial_node_count
        self.locations = locations
        self.self_link = self_link
        self.version = version
        self.instance_group_urls = instance_group_urls
        self.status = status
        self.status_message = status_message
        self.autoscaling = autoscaling
        self.management = management
        self.max_pods_constraint = max_pods_constraint
        self.conditions = conditions
        self.pod_ipv4_cidr_size = pod_ipv4_cidr_size
        self.upgrade_settings = upgrade_settings

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePools()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if ClusterNodePoolsConfig.to_proto(resource.config):
            res.config.CopyFrom(ClusterNodePoolsConfig.to_proto(resource.config))
        else:
            res.ClearField("config")
        if Primitive.to_proto(resource.initial_node_count):
            res.initial_node_count = Primitive.to_proto(resource.initial_node_count)
        if Primitive.to_proto(resource.locations):
            res.locations.extend(Primitive.to_proto(resource.locations))
        if Primitive.to_proto(resource.self_link):
            res.self_link = Primitive.to_proto(resource.self_link)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        if Primitive.to_proto(resource.instance_group_urls):
            res.instance_group_urls.extend(
                Primitive.to_proto(resource.instance_group_urls)
            )
        if ClusterNodePoolsStatusEnum.to_proto(resource.status):
            res.status = ClusterNodePoolsStatusEnum.to_proto(resource.status)
        if Primitive.to_proto(resource.status_message):
            res.status_message = Primitive.to_proto(resource.status_message)
        if ClusterNodePoolsAutoscaling.to_proto(resource.autoscaling):
            res.autoscaling.CopyFrom(
                ClusterNodePoolsAutoscaling.to_proto(resource.autoscaling)
            )
        else:
            res.ClearField("autoscaling")
        if ClusterNodePoolsManagement.to_proto(resource.management):
            res.management.CopyFrom(
                ClusterNodePoolsManagement.to_proto(resource.management)
            )
        else:
            res.ClearField("management")
        if ClusterNodePoolsMaxPodsConstraint.to_proto(resource.max_pods_constraint):
            res.max_pods_constraint.CopyFrom(
                ClusterNodePoolsMaxPodsConstraint.to_proto(resource.max_pods_constraint)
            )
        else:
            res.ClearField("max_pods_constraint")
        if ClusterNodePoolsConditionsArray.to_proto(resource.conditions):
            res.conditions.extend(
                ClusterNodePoolsConditionsArray.to_proto(resource.conditions)
            )
        if Primitive.to_proto(resource.pod_ipv4_cidr_size):
            res.pod_ipv4_cidr_size = Primitive.to_proto(resource.pod_ipv4_cidr_size)
        if ClusterNodePoolsUpgradeSettings.to_proto(resource.upgrade_settings):
            res.upgrade_settings.CopyFrom(
                ClusterNodePoolsUpgradeSettings.to_proto(resource.upgrade_settings)
            )
        else:
            res.ClearField("upgrade_settings")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePools(
            name=Primitive.from_proto(resource.name),
            config=ClusterNodePoolsConfig.from_proto(resource.config),
            initial_node_count=Primitive.from_proto(resource.initial_node_count),
            locations=Primitive.from_proto(resource.locations),
            self_link=Primitive.from_proto(resource.self_link),
            version=Primitive.from_proto(resource.version),
            instance_group_urls=Primitive.from_proto(resource.instance_group_urls),
            status=ClusterNodePoolsStatusEnum.from_proto(resource.status),
            status_message=Primitive.from_proto(resource.status_message),
            autoscaling=ClusterNodePoolsAutoscaling.from_proto(resource.autoscaling),
            management=ClusterNodePoolsManagement.from_proto(resource.management),
            max_pods_constraint=ClusterNodePoolsMaxPodsConstraint.from_proto(
                resource.max_pods_constraint
            ),
            conditions=ClusterNodePoolsConditionsArray.from_proto(resource.conditions),
            pod_ipv4_cidr_size=Primitive.from_proto(resource.pod_ipv4_cidr_size),
            upgrade_settings=ClusterNodePoolsUpgradeSettings.from_proto(
                resource.upgrade_settings
            ),
        )


class ClusterNodePoolsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePools.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePools.from_proto(i) for i in resources]


class ClusterNodePoolsConfig(object):
    def __init__(
        self,
        machine_type: str = None,
        disk_size_gb: int = None,
        oauth_scopes: list = None,
        service_account: str = None,
        metadata: dict = None,
        image_type: str = None,
        labels: dict = None,
        local_ssd_count: int = None,
        tags: list = None,
        preemptible: bool = None,
        accelerators: list = None,
        disk_type: str = None,
        min_cpu_platform: str = None,
        workload_metadata_config: dict = None,
        taints: list = None,
        sandbox_config: dict = None,
        node_group: str = None,
        reservation_affinity: dict = None,
        shielded_instance_config: dict = None,
        linux_node_config: dict = None,
        kubelet_config: dict = None,
        boot_disk_kms_key: str = None,
    ):
        self.machine_type = machine_type
        self.disk_size_gb = disk_size_gb
        self.oauth_scopes = oauth_scopes
        self.service_account = service_account
        self.metadata = metadata
        self.image_type = image_type
        self.labels = labels
        self.local_ssd_count = local_ssd_count
        self.tags = tags
        self.preemptible = preemptible
        self.accelerators = accelerators
        self.disk_type = disk_type
        self.min_cpu_platform = min_cpu_platform
        self.workload_metadata_config = workload_metadata_config
        self.taints = taints
        self.sandbox_config = sandbox_config
        self.node_group = node_group
        self.reservation_affinity = reservation_affinity
        self.shielded_instance_config = shielded_instance_config
        self.linux_node_config = linux_node_config
        self.kubelet_config = kubelet_config
        self.boot_disk_kms_key = boot_disk_kms_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsConfig()
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if Primitive.to_proto(resource.disk_size_gb):
            res.disk_size_gb = Primitive.to_proto(resource.disk_size_gb)
        if Primitive.to_proto(resource.oauth_scopes):
            res.oauth_scopes.extend(Primitive.to_proto(resource.oauth_scopes))
        if Primitive.to_proto(resource.service_account):
            res.service_account = Primitive.to_proto(resource.service_account)
        if Primitive.to_proto(resource.metadata):
            res.metadata = Primitive.to_proto(resource.metadata)
        if Primitive.to_proto(resource.image_type):
            res.image_type = Primitive.to_proto(resource.image_type)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if Primitive.to_proto(resource.local_ssd_count):
            res.local_ssd_count = Primitive.to_proto(resource.local_ssd_count)
        if Primitive.to_proto(resource.tags):
            res.tags.extend(Primitive.to_proto(resource.tags))
        if Primitive.to_proto(resource.preemptible):
            res.preemptible = Primitive.to_proto(resource.preemptible)
        if ClusterNodePoolsConfigAcceleratorsArray.to_proto(resource.accelerators):
            res.accelerators.extend(
                ClusterNodePoolsConfigAcceleratorsArray.to_proto(resource.accelerators)
            )
        if Primitive.to_proto(resource.disk_type):
            res.disk_type = Primitive.to_proto(resource.disk_type)
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        if ClusterNodePoolsConfigWorkloadMetadataConfig.to_proto(
            resource.workload_metadata_config
        ):
            res.workload_metadata_config.CopyFrom(
                ClusterNodePoolsConfigWorkloadMetadataConfig.to_proto(
                    resource.workload_metadata_config
                )
            )
        else:
            res.ClearField("workload_metadata_config")
        if ClusterNodePoolsConfigTaintsArray.to_proto(resource.taints):
            res.taints.extend(
                ClusterNodePoolsConfigTaintsArray.to_proto(resource.taints)
            )
        if ClusterNodePoolsConfigSandboxConfig.to_proto(resource.sandbox_config):
            res.sandbox_config.CopyFrom(
                ClusterNodePoolsConfigSandboxConfig.to_proto(resource.sandbox_config)
            )
        else:
            res.ClearField("sandbox_config")
        if Primitive.to_proto(resource.node_group):
            res.node_group = Primitive.to_proto(resource.node_group)
        if ClusterNodePoolsConfigReservationAffinity.to_proto(
            resource.reservation_affinity
        ):
            res.reservation_affinity.CopyFrom(
                ClusterNodePoolsConfigReservationAffinity.to_proto(
                    resource.reservation_affinity
                )
            )
        else:
            res.ClearField("reservation_affinity")
        if ClusterNodePoolsConfigShieldedInstanceConfig.to_proto(
            resource.shielded_instance_config
        ):
            res.shielded_instance_config.CopyFrom(
                ClusterNodePoolsConfigShieldedInstanceConfig.to_proto(
                    resource.shielded_instance_config
                )
            )
        else:
            res.ClearField("shielded_instance_config")
        if ClusterNodePoolsConfigLinuxNodeConfig.to_proto(resource.linux_node_config):
            res.linux_node_config.CopyFrom(
                ClusterNodePoolsConfigLinuxNodeConfig.to_proto(
                    resource.linux_node_config
                )
            )
        else:
            res.ClearField("linux_node_config")
        if ClusterNodePoolsConfigKubeletConfig.to_proto(resource.kubelet_config):
            res.kubelet_config.CopyFrom(
                ClusterNodePoolsConfigKubeletConfig.to_proto(resource.kubelet_config)
            )
        else:
            res.ClearField("kubelet_config")
        if Primitive.to_proto(resource.boot_disk_kms_key):
            res.boot_disk_kms_key = Primitive.to_proto(resource.boot_disk_kms_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsConfig(
            machine_type=Primitive.from_proto(resource.machine_type),
            disk_size_gb=Primitive.from_proto(resource.disk_size_gb),
            oauth_scopes=Primitive.from_proto(resource.oauth_scopes),
            service_account=Primitive.from_proto(resource.service_account),
            metadata=Primitive.from_proto(resource.metadata),
            image_type=Primitive.from_proto(resource.image_type),
            labels=Primitive.from_proto(resource.labels),
            local_ssd_count=Primitive.from_proto(resource.local_ssd_count),
            tags=Primitive.from_proto(resource.tags),
            preemptible=Primitive.from_proto(resource.preemptible),
            accelerators=ClusterNodePoolsConfigAcceleratorsArray.from_proto(
                resource.accelerators
            ),
            disk_type=Primitive.from_proto(resource.disk_type),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
            workload_metadata_config=ClusterNodePoolsConfigWorkloadMetadataConfig.from_proto(
                resource.workload_metadata_config
            ),
            taints=ClusterNodePoolsConfigTaintsArray.from_proto(resource.taints),
            sandbox_config=ClusterNodePoolsConfigSandboxConfig.from_proto(
                resource.sandbox_config
            ),
            node_group=Primitive.from_proto(resource.node_group),
            reservation_affinity=ClusterNodePoolsConfigReservationAffinity.from_proto(
                resource.reservation_affinity
            ),
            shielded_instance_config=ClusterNodePoolsConfigShieldedInstanceConfig.from_proto(
                resource.shielded_instance_config
            ),
            linux_node_config=ClusterNodePoolsConfigLinuxNodeConfig.from_proto(
                resource.linux_node_config
            ),
            kubelet_config=ClusterNodePoolsConfigKubeletConfig.from_proto(
                resource.kubelet_config
            ),
            boot_disk_kms_key=Primitive.from_proto(resource.boot_disk_kms_key),
        )


class ClusterNodePoolsConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePoolsConfig.from_proto(i) for i in resources]


class ClusterNodePoolsConfigAccelerators(object):
    def __init__(self, accelerator_count: int = None, accelerator_type: str = None):
        self.accelerator_count = accelerator_count
        self.accelerator_type = accelerator_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsConfigAccelerators()
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsConfigAccelerators(
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
        )


class ClusterNodePoolsConfigAcceleratorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsConfigAccelerators.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePoolsConfigAccelerators.from_proto(i) for i in resources]


class ClusterNodePoolsConfigWorkloadMetadataConfig(object):
    def __init__(self, mode: str = None):
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsConfigWorkloadMetadataConfig()
        if ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum.to_proto(resource.mode):
            res.mode = ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum.to_proto(
                resource.mode
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsConfigWorkloadMetadataConfig(
            mode=ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum.from_proto(
                resource.mode
            ),
        )


class ClusterNodePoolsConfigWorkloadMetadataConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterNodePoolsConfigWorkloadMetadataConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterNodePoolsConfigWorkloadMetadataConfig.from_proto(i)
            for i in resources
        ]


class ClusterNodePoolsConfigTaints(object):
    def __init__(self, key: str = None, value: str = None, effect: str = None):
        self.key = key
        self.value = value
        self.effect = effect

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsConfigTaints()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if ClusterNodePoolsConfigTaintsEffectEnum.to_proto(resource.effect):
            res.effect = ClusterNodePoolsConfigTaintsEffectEnum.to_proto(
                resource.effect
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsConfigTaints(
            key=Primitive.from_proto(resource.key),
            value=Primitive.from_proto(resource.value),
            effect=ClusterNodePoolsConfigTaintsEffectEnum.from_proto(resource.effect),
        )


class ClusterNodePoolsConfigTaintsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsConfigTaints.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePoolsConfigTaints.from_proto(i) for i in resources]


class ClusterNodePoolsConfigSandboxConfig(object):
    def __init__(self, type: str = None):
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsConfigSandboxConfig()
        if ClusterNodePoolsConfigSandboxConfigTypeEnum.to_proto(resource.type):
            res.type = ClusterNodePoolsConfigSandboxConfigTypeEnum.to_proto(
                resource.type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsConfigSandboxConfig(
            type=ClusterNodePoolsConfigSandboxConfigTypeEnum.from_proto(resource.type),
        )


class ClusterNodePoolsConfigSandboxConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsConfigSandboxConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePoolsConfigSandboxConfig.from_proto(i) for i in resources]


class ClusterNodePoolsConfigReservationAffinity(object):
    def __init__(
        self, consume_reservation_type: str = None, key: str = None, values: list = None
    ):
        self.consume_reservation_type = consume_reservation_type
        self.key = key
        self.values = values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsConfigReservationAffinity()
        if ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
            resource.consume_reservation_type
        ):
            res.consume_reservation_type = ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
                resource.consume_reservation_type
            )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.values):
            res.values.extend(Primitive.to_proto(resource.values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsConfigReservationAffinity(
            consume_reservation_type=ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum.from_proto(
                resource.consume_reservation_type
            ),
            key=Primitive.from_proto(resource.key),
            values=Primitive.from_proto(resource.values),
        )


class ClusterNodePoolsConfigReservationAffinityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterNodePoolsConfigReservationAffinity.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterNodePoolsConfigReservationAffinity.from_proto(i) for i in resources
        ]


class ClusterNodePoolsConfigShieldedInstanceConfig(object):
    def __init__(
        self, enable_secure_boot: bool = None, enable_integrity_monitoring: bool = None
    ):
        self.enable_secure_boot = enable_secure_boot
        self.enable_integrity_monitoring = enable_integrity_monitoring

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsConfigShieldedInstanceConfig()
        if Primitive.to_proto(resource.enable_secure_boot):
            res.enable_secure_boot = Primitive.to_proto(resource.enable_secure_boot)
        if Primitive.to_proto(resource.enable_integrity_monitoring):
            res.enable_integrity_monitoring = Primitive.to_proto(
                resource.enable_integrity_monitoring
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsConfigShieldedInstanceConfig(
            enable_secure_boot=Primitive.from_proto(resource.enable_secure_boot),
            enable_integrity_monitoring=Primitive.from_proto(
                resource.enable_integrity_monitoring
            ),
        )


class ClusterNodePoolsConfigShieldedInstanceConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterNodePoolsConfigShieldedInstanceConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterNodePoolsConfigShieldedInstanceConfig.from_proto(i)
            for i in resources
        ]


class ClusterNodePoolsConfigLinuxNodeConfig(object):
    def __init__(self, sysctls: dict = None):
        self.sysctls = sysctls

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsConfigLinuxNodeConfig()
        if Primitive.to_proto(resource.sysctls):
            res.sysctls = Primitive.to_proto(resource.sysctls)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsConfigLinuxNodeConfig(
            sysctls=Primitive.from_proto(resource.sysctls),
        )


class ClusterNodePoolsConfigLinuxNodeConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsConfigLinuxNodeConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePoolsConfigLinuxNodeConfig.from_proto(i) for i in resources]


class ClusterNodePoolsConfigKubeletConfig(object):
    def __init__(
        self,
        cpu_manager_policy: str = None,
        cpu_cfs_quota: bool = None,
        cpu_cfs_quota_period: str = None,
    ):
        self.cpu_manager_policy = cpu_manager_policy
        self.cpu_cfs_quota = cpu_cfs_quota
        self.cpu_cfs_quota_period = cpu_cfs_quota_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsConfigKubeletConfig()
        if Primitive.to_proto(resource.cpu_manager_policy):
            res.cpu_manager_policy = Primitive.to_proto(resource.cpu_manager_policy)
        if Primitive.to_proto(resource.cpu_cfs_quota):
            res.cpu_cfs_quota = Primitive.to_proto(resource.cpu_cfs_quota)
        if Primitive.to_proto(resource.cpu_cfs_quota_period):
            res.cpu_cfs_quota_period = Primitive.to_proto(resource.cpu_cfs_quota_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsConfigKubeletConfig(
            cpu_manager_policy=Primitive.from_proto(resource.cpu_manager_policy),
            cpu_cfs_quota=Primitive.from_proto(resource.cpu_cfs_quota),
            cpu_cfs_quota_period=Primitive.from_proto(resource.cpu_cfs_quota_period),
        )


class ClusterNodePoolsConfigKubeletConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsConfigKubeletConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePoolsConfigKubeletConfig.from_proto(i) for i in resources]


class ClusterNodePoolsAutoscaling(object):
    def __init__(
        self,
        enabled: bool = None,
        min_node_count: int = None,
        max_node_count: int = None,
        autoprovisioned: bool = None,
    ):
        self.enabled = enabled
        self.min_node_count = min_node_count
        self.max_node_count = max_node_count
        self.autoprovisioned = autoprovisioned

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsAutoscaling()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.min_node_count):
            res.min_node_count = Primitive.to_proto(resource.min_node_count)
        if Primitive.to_proto(resource.max_node_count):
            res.max_node_count = Primitive.to_proto(resource.max_node_count)
        if Primitive.to_proto(resource.autoprovisioned):
            res.autoprovisioned = Primitive.to_proto(resource.autoprovisioned)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsAutoscaling(
            enabled=Primitive.from_proto(resource.enabled),
            min_node_count=Primitive.from_proto(resource.min_node_count),
            max_node_count=Primitive.from_proto(resource.max_node_count),
            autoprovisioned=Primitive.from_proto(resource.autoprovisioned),
        )


class ClusterNodePoolsAutoscalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsAutoscaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePoolsAutoscaling.from_proto(i) for i in resources]


class ClusterNodePoolsManagement(object):
    def __init__(
        self,
        auto_upgrade: bool = None,
        auto_repair: bool = None,
        upgrade_options: dict = None,
    ):
        self.auto_upgrade = auto_upgrade
        self.auto_repair = auto_repair
        self.upgrade_options = upgrade_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsManagement()
        if Primitive.to_proto(resource.auto_upgrade):
            res.auto_upgrade = Primitive.to_proto(resource.auto_upgrade)
        if Primitive.to_proto(resource.auto_repair):
            res.auto_repair = Primitive.to_proto(resource.auto_repair)
        if ClusterNodePoolsManagementUpgradeOptions.to_proto(resource.upgrade_options):
            res.upgrade_options.CopyFrom(
                ClusterNodePoolsManagementUpgradeOptions.to_proto(
                    resource.upgrade_options
                )
            )
        else:
            res.ClearField("upgrade_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsManagement(
            auto_upgrade=Primitive.from_proto(resource.auto_upgrade),
            auto_repair=Primitive.from_proto(resource.auto_repair),
            upgrade_options=ClusterNodePoolsManagementUpgradeOptions.from_proto(
                resource.upgrade_options
            ),
        )


class ClusterNodePoolsManagementArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsManagement.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePoolsManagement.from_proto(i) for i in resources]


class ClusterNodePoolsManagementUpgradeOptions(object):
    def __init__(self, auto_upgrade_start_time: str = None, description: str = None):
        self.auto_upgrade_start_time = auto_upgrade_start_time
        self.description = description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsManagementUpgradeOptions()
        if Primitive.to_proto(resource.auto_upgrade_start_time):
            res.auto_upgrade_start_time = Primitive.to_proto(
                resource.auto_upgrade_start_time
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsManagementUpgradeOptions(
            auto_upgrade_start_time=Primitive.from_proto(
                resource.auto_upgrade_start_time
            ),
            description=Primitive.from_proto(resource.description),
        )


class ClusterNodePoolsManagementUpgradeOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsManagementUpgradeOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterNodePoolsManagementUpgradeOptions.from_proto(i) for i in resources
        ]


class ClusterNodePoolsMaxPodsConstraint(object):
    def __init__(self, max_pods_per_node: int = None):
        self.max_pods_per_node = max_pods_per_node

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsMaxPodsConstraint()
        if Primitive.to_proto(resource.max_pods_per_node):
            res.max_pods_per_node = Primitive.to_proto(resource.max_pods_per_node)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsMaxPodsConstraint(
            max_pods_per_node=Primitive.from_proto(resource.max_pods_per_node),
        )


class ClusterNodePoolsMaxPodsConstraintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsMaxPodsConstraint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePoolsMaxPodsConstraint.from_proto(i) for i in resources]


class ClusterNodePoolsConditions(object):
    def __init__(
        self, code: str = None, message: str = None, canonical_code: str = None
    ):
        self.code = code
        self.message = message
        self.canonical_code = canonical_code

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsConditions()
        if ClusterNodePoolsConditionsCodeEnum.to_proto(resource.code):
            res.code = ClusterNodePoolsConditionsCodeEnum.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if ClusterNodePoolsConditionsCanonicalCodeEnum.to_proto(
            resource.canonical_code
        ):
            res.canonical_code = ClusterNodePoolsConditionsCanonicalCodeEnum.to_proto(
                resource.canonical_code
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsConditions(
            code=ClusterNodePoolsConditionsCodeEnum.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            canonical_code=ClusterNodePoolsConditionsCanonicalCodeEnum.from_proto(
                resource.canonical_code
            ),
        )


class ClusterNodePoolsConditionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsConditions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePoolsConditions.from_proto(i) for i in resources]


class ClusterNodePoolsUpgradeSettings(object):
    def __init__(self, max_surge: int = None, max_unavailable: int = None):
        self.max_surge = max_surge
        self.max_unavailable = max_unavailable

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePoolsUpgradeSettings()
        if Primitive.to_proto(resource.max_surge):
            res.max_surge = Primitive.to_proto(resource.max_surge)
        if Primitive.to_proto(resource.max_unavailable):
            res.max_unavailable = Primitive.to_proto(resource.max_unavailable)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePoolsUpgradeSettings(
            max_surge=Primitive.from_proto(resource.max_surge),
            max_unavailable=Primitive.from_proto(resource.max_unavailable),
        )


class ClusterNodePoolsUpgradeSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePoolsUpgradeSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePoolsUpgradeSettings.from_proto(i) for i in resources]


class ClusterLegacyAbac(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterLegacyAbac()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterLegacyAbac(enabled=Primitive.from_proto(resource.enabled),)


class ClusterLegacyAbacArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterLegacyAbac.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterLegacyAbac.from_proto(i) for i in resources]


class ClusterNetworkPolicy(object):
    def __init__(self, provider: str = None, enabled: bool = None):
        self.provider = provider
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNetworkPolicy()
        if ClusterNetworkPolicyProviderEnum.to_proto(resource.provider):
            res.provider = ClusterNetworkPolicyProviderEnum.to_proto(resource.provider)
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNetworkPolicy(
            provider=ClusterNetworkPolicyProviderEnum.from_proto(resource.provider),
            enabled=Primitive.from_proto(resource.enabled),
        )


class ClusterNetworkPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNetworkPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNetworkPolicy.from_proto(i) for i in resources]


class ClusterIPAllocationPolicy(object):
    def __init__(
        self,
        use_ip_aliases: bool = None,
        create_subnetwork: bool = None,
        subnetwork_name: str = None,
        cluster_secondary_range_name: str = None,
        services_secondary_range_name: str = None,
        cluster_ipv4_cidr_block: str = None,
        node_ipv4_cidr_block: str = None,
        services_ipv4_cidr_block: str = None,
        tpu_ipv4_cidr_block: str = None,
        cluster_ipv4_cidr: str = None,
        node_ipv4_cidr: str = None,
        services_ipv4_cidr: str = None,
        use_routes: bool = None,
    ):
        self.use_ip_aliases = use_ip_aliases
        self.create_subnetwork = create_subnetwork
        self.subnetwork_name = subnetwork_name
        self.cluster_secondary_range_name = cluster_secondary_range_name
        self.services_secondary_range_name = services_secondary_range_name
        self.cluster_ipv4_cidr_block = cluster_ipv4_cidr_block
        self.node_ipv4_cidr_block = node_ipv4_cidr_block
        self.services_ipv4_cidr_block = services_ipv4_cidr_block
        self.tpu_ipv4_cidr_block = tpu_ipv4_cidr_block
        self.cluster_ipv4_cidr = cluster_ipv4_cidr
        self.node_ipv4_cidr = node_ipv4_cidr
        self.services_ipv4_cidr = services_ipv4_cidr
        self.use_routes = use_routes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterIPAllocationPolicy()
        if Primitive.to_proto(resource.use_ip_aliases):
            res.use_ip_aliases = Primitive.to_proto(resource.use_ip_aliases)
        if Primitive.to_proto(resource.create_subnetwork):
            res.create_subnetwork = Primitive.to_proto(resource.create_subnetwork)
        if Primitive.to_proto(resource.subnetwork_name):
            res.subnetwork_name = Primitive.to_proto(resource.subnetwork_name)
        if Primitive.to_proto(resource.cluster_secondary_range_name):
            res.cluster_secondary_range_name = Primitive.to_proto(
                resource.cluster_secondary_range_name
            )
        if Primitive.to_proto(resource.services_secondary_range_name):
            res.services_secondary_range_name = Primitive.to_proto(
                resource.services_secondary_range_name
            )
        if Primitive.to_proto(resource.cluster_ipv4_cidr_block):
            res.cluster_ipv4_cidr_block = Primitive.to_proto(
                resource.cluster_ipv4_cidr_block
            )
        if Primitive.to_proto(resource.node_ipv4_cidr_block):
            res.node_ipv4_cidr_block = Primitive.to_proto(resource.node_ipv4_cidr_block)
        if Primitive.to_proto(resource.services_ipv4_cidr_block):
            res.services_ipv4_cidr_block = Primitive.to_proto(
                resource.services_ipv4_cidr_block
            )
        if Primitive.to_proto(resource.tpu_ipv4_cidr_block):
            res.tpu_ipv4_cidr_block = Primitive.to_proto(resource.tpu_ipv4_cidr_block)
        if Primitive.to_proto(resource.cluster_ipv4_cidr):
            res.cluster_ipv4_cidr = Primitive.to_proto(resource.cluster_ipv4_cidr)
        if Primitive.to_proto(resource.node_ipv4_cidr):
            res.node_ipv4_cidr = Primitive.to_proto(resource.node_ipv4_cidr)
        if Primitive.to_proto(resource.services_ipv4_cidr):
            res.services_ipv4_cidr = Primitive.to_proto(resource.services_ipv4_cidr)
        if Primitive.to_proto(resource.use_routes):
            res.use_routes = Primitive.to_proto(resource.use_routes)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterIPAllocationPolicy(
            use_ip_aliases=Primitive.from_proto(resource.use_ip_aliases),
            create_subnetwork=Primitive.from_proto(resource.create_subnetwork),
            subnetwork_name=Primitive.from_proto(resource.subnetwork_name),
            cluster_secondary_range_name=Primitive.from_proto(
                resource.cluster_secondary_range_name
            ),
            services_secondary_range_name=Primitive.from_proto(
                resource.services_secondary_range_name
            ),
            cluster_ipv4_cidr_block=Primitive.from_proto(
                resource.cluster_ipv4_cidr_block
            ),
            node_ipv4_cidr_block=Primitive.from_proto(resource.node_ipv4_cidr_block),
            services_ipv4_cidr_block=Primitive.from_proto(
                resource.services_ipv4_cidr_block
            ),
            tpu_ipv4_cidr_block=Primitive.from_proto(resource.tpu_ipv4_cidr_block),
            cluster_ipv4_cidr=Primitive.from_proto(resource.cluster_ipv4_cidr),
            node_ipv4_cidr=Primitive.from_proto(resource.node_ipv4_cidr),
            services_ipv4_cidr=Primitive.from_proto(resource.services_ipv4_cidr),
            use_routes=Primitive.from_proto(resource.use_routes),
        )


class ClusterIPAllocationPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterIPAllocationPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterIPAllocationPolicy.from_proto(i) for i in resources]


class ClusterMasterAuthorizedNetworksConfig(object):
    def __init__(self, enabled: bool = None, cidr_blocks: list = None):
        self.enabled = enabled
        self.cidr_blocks = cidr_blocks

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMasterAuthorizedNetworksConfig()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if ClusterMasterAuthorizedNetworksConfigCidrBlocksArray.to_proto(
            resource.cidr_blocks
        ):
            res.cidr_blocks.extend(
                ClusterMasterAuthorizedNetworksConfigCidrBlocksArray.to_proto(
                    resource.cidr_blocks
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMasterAuthorizedNetworksConfig(
            enabled=Primitive.from_proto(resource.enabled),
            cidr_blocks=ClusterMasterAuthorizedNetworksConfigCidrBlocksArray.from_proto(
                resource.cidr_blocks
            ),
        )


class ClusterMasterAuthorizedNetworksConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMasterAuthorizedNetworksConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterMasterAuthorizedNetworksConfig.from_proto(i) for i in resources]


class ClusterMasterAuthorizedNetworksConfigCidrBlocks(object):
    def __init__(self, display_name: str = None, cidr_block: str = None):
        self.display_name = display_name
        self.cidr_block = cidr_block

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMasterAuthorizedNetworksConfigCidrBlocks()
        if Primitive.to_proto(resource.display_name):
            res.display_name = Primitive.to_proto(resource.display_name)
        if Primitive.to_proto(resource.cidr_block):
            res.cidr_block = Primitive.to_proto(resource.cidr_block)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMasterAuthorizedNetworksConfigCidrBlocks(
            display_name=Primitive.from_proto(resource.display_name),
            cidr_block=Primitive.from_proto(resource.cidr_block),
        )


class ClusterMasterAuthorizedNetworksConfigCidrBlocksArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterMasterAuthorizedNetworksConfigCidrBlocks.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterMasterAuthorizedNetworksConfigCidrBlocks.from_proto(i)
            for i in resources
        ]


class ClusterBinaryAuthorization(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterBinaryAuthorization()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterBinaryAuthorization(
            enabled=Primitive.from_proto(resource.enabled),
        )


class ClusterBinaryAuthorizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterBinaryAuthorization.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterBinaryAuthorization.from_proto(i) for i in resources]


class ClusterAutoscaling(object):
    def __init__(
        self,
        enable_node_autoprovisioning: bool = None,
        resource_limits: list = None,
        autoprovisioning_node_pool_defaults: dict = None,
        autoprovisioning_locations: list = None,
    ):
        self.enable_node_autoprovisioning = enable_node_autoprovisioning
        self.resource_limits = resource_limits
        self.autoprovisioning_node_pool_defaults = autoprovisioning_node_pool_defaults
        self.autoprovisioning_locations = autoprovisioning_locations

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAutoscaling()
        if Primitive.to_proto(resource.enable_node_autoprovisioning):
            res.enable_node_autoprovisioning = Primitive.to_proto(
                resource.enable_node_autoprovisioning
            )
        if ClusterAutoscalingResourceLimitsArray.to_proto(resource.resource_limits):
            res.resource_limits.extend(
                ClusterAutoscalingResourceLimitsArray.to_proto(resource.resource_limits)
            )
        if ClusterAutoscalingAutoprovisioningNodePoolDefaults.to_proto(
            resource.autoprovisioning_node_pool_defaults
        ):
            res.autoprovisioning_node_pool_defaults.CopyFrom(
                ClusterAutoscalingAutoprovisioningNodePoolDefaults.to_proto(
                    resource.autoprovisioning_node_pool_defaults
                )
            )
        else:
            res.ClearField("autoprovisioning_node_pool_defaults")
        if Primitive.to_proto(resource.autoprovisioning_locations):
            res.autoprovisioning_locations.extend(
                Primitive.to_proto(resource.autoprovisioning_locations)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscaling(
            enable_node_autoprovisioning=Primitive.from_proto(
                resource.enable_node_autoprovisioning
            ),
            resource_limits=ClusterAutoscalingResourceLimitsArray.from_proto(
                resource.resource_limits
            ),
            autoprovisioning_node_pool_defaults=ClusterAutoscalingAutoprovisioningNodePoolDefaults.from_proto(
                resource.autoprovisioning_node_pool_defaults
            ),
            autoprovisioning_locations=Primitive.from_proto(
                resource.autoprovisioning_locations
            ),
        )


class ClusterAutoscalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAutoscaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAutoscaling.from_proto(i) for i in resources]


class ClusterAutoscalingResourceLimits(object):
    def __init__(
        self, resource_type: str = None, minimum: int = None, maximum: int = None
    ):
        self.resource_type = resource_type
        self.minimum = minimum
        self.maximum = maximum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAutoscalingResourceLimits()
        if Primitive.to_proto(resource.resource_type):
            res.resource_type = Primitive.to_proto(resource.resource_type)
        if Primitive.to_proto(resource.minimum):
            res.minimum = Primitive.to_proto(resource.minimum)
        if Primitive.to_proto(resource.maximum):
            res.maximum = Primitive.to_proto(resource.maximum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscalingResourceLimits(
            resource_type=Primitive.from_proto(resource.resource_type),
            minimum=Primitive.from_proto(resource.minimum),
            maximum=Primitive.from_proto(resource.maximum),
        )


class ClusterAutoscalingResourceLimitsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAutoscalingResourceLimits.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAutoscalingResourceLimits.from_proto(i) for i in resources]


class ClusterAutoscalingAutoprovisioningNodePoolDefaults(object):
    def __init__(
        self,
        oauth_scopes: list = None,
        service_account: str = None,
        upgrade_settings: dict = None,
        management: dict = None,
        min_cpu_platform: str = None,
        disk_size_gb: int = None,
        disk_type: str = None,
        shielded_instance_config: dict = None,
        boot_disk_kms_key: str = None,
    ):
        self.oauth_scopes = oauth_scopes
        self.service_account = service_account
        self.upgrade_settings = upgrade_settings
        self.management = management
        self.min_cpu_platform = min_cpu_platform
        self.disk_size_gb = disk_size_gb
        self.disk_type = disk_type
        self.shielded_instance_config = shielded_instance_config
        self.boot_disk_kms_key = boot_disk_kms_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaults()
        if Primitive.to_proto(resource.oauth_scopes):
            res.oauth_scopes.extend(Primitive.to_proto(resource.oauth_scopes))
        if Primitive.to_proto(resource.service_account):
            res.service_account = Primitive.to_proto(resource.service_account)
        if ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings.to_proto(
            resource.upgrade_settings
        ):
            res.upgrade_settings.CopyFrom(
                ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings.to_proto(
                    resource.upgrade_settings
                )
            )
        else:
            res.ClearField("upgrade_settings")
        if ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement.to_proto(
            resource.management
        ):
            res.management.CopyFrom(
                ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement.to_proto(
                    resource.management
                )
            )
        else:
            res.ClearField("management")
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        if Primitive.to_proto(resource.disk_size_gb):
            res.disk_size_gb = Primitive.to_proto(resource.disk_size_gb)
        if Primitive.to_proto(resource.disk_type):
            res.disk_type = Primitive.to_proto(resource.disk_type)
        if ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig.to_proto(
            resource.shielded_instance_config
        ):
            res.shielded_instance_config.CopyFrom(
                ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig.to_proto(
                    resource.shielded_instance_config
                )
            )
        else:
            res.ClearField("shielded_instance_config")
        if Primitive.to_proto(resource.boot_disk_kms_key):
            res.boot_disk_kms_key = Primitive.to_proto(resource.boot_disk_kms_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscalingAutoprovisioningNodePoolDefaults(
            oauth_scopes=Primitive.from_proto(resource.oauth_scopes),
            service_account=Primitive.from_proto(resource.service_account),
            upgrade_settings=ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings.from_proto(
                resource.upgrade_settings
            ),
            management=ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement.from_proto(
                resource.management
            ),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
            disk_size_gb=Primitive.from_proto(resource.disk_size_gb),
            disk_type=Primitive.from_proto(resource.disk_type),
            shielded_instance_config=ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig.from_proto(
                resource.shielded_instance_config
            ),
            boot_disk_kms_key=Primitive.from_proto(resource.boot_disk_kms_key),
        )


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaults.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaults.from_proto(i)
            for i in resources
        ]


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(object):
    def __init__(self, max_surge: int = None, max_unavailable: int = None):
        self.max_surge = max_surge
        self.max_unavailable = max_unavailable

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings()
        )
        if Primitive.to_proto(resource.max_surge):
            res.max_surge = Primitive.to_proto(resource.max_surge)
        if Primitive.to_proto(resource.max_unavailable):
            res.max_unavailable = Primitive.to_proto(resource.max_unavailable)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(
            max_surge=Primitive.from_proto(resource.max_surge),
            max_unavailable=Primitive.from_proto(resource.max_unavailable),
        )


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings.from_proto(
                i
            )
            for i in resources
        ]


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(object):
    def __init__(
        self,
        auto_upgrade: bool = None,
        auto_repair: bool = None,
        upgrade_options: dict = None,
    ):
        self.auto_upgrade = auto_upgrade
        self.auto_repair = auto_repair
        self.upgrade_options = upgrade_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement()
        )
        if Primitive.to_proto(resource.auto_upgrade):
            res.auto_upgrade = Primitive.to_proto(resource.auto_upgrade)
        if Primitive.to_proto(resource.auto_repair):
            res.auto_repair = Primitive.to_proto(resource.auto_repair)
        if ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions.to_proto(
            resource.upgrade_options
        ):
            res.upgrade_options.CopyFrom(
                ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions.to_proto(
                    resource.upgrade_options
                )
            )
        else:
            res.ClearField("upgrade_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(
            auto_upgrade=Primitive.from_proto(resource.auto_upgrade),
            auto_repair=Primitive.from_proto(resource.auto_repair),
            upgrade_options=ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions.from_proto(
                resource.upgrade_options
            ),
        )


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement.from_proto(i)
            for i in resources
        ]


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(
    object
):
    def __init__(self, auto_upgrade_start_time: str = None, description: str = None):
        self.auto_upgrade_start_time = auto_upgrade_start_time
        self.description = description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions()
        )
        if Primitive.to_proto(resource.auto_upgrade_start_time):
            res.auto_upgrade_start_time = Primitive.to_proto(
                resource.auto_upgrade_start_time
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(
            auto_upgrade_start_time=Primitive.from_proto(
                resource.auto_upgrade_start_time
            ),
            description=Primitive.from_proto(resource.description),
        )


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions.from_proto(
                i
            )
            for i in resources
        ]


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(object):
    def __init__(
        self, enable_secure_boot: bool = None, enable_integrity_monitoring: bool = None
    ):
        self.enable_secure_boot = enable_secure_boot
        self.enable_integrity_monitoring = enable_integrity_monitoring

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig()
        )
        if Primitive.to_proto(resource.enable_secure_boot):
            res.enable_secure_boot = Primitive.to_proto(resource.enable_secure_boot)
        if Primitive.to_proto(resource.enable_integrity_monitoring):
            res.enable_integrity_monitoring = Primitive.to_proto(
                resource.enable_integrity_monitoring
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(
            enable_secure_boot=Primitive.from_proto(resource.enable_secure_boot),
            enable_integrity_monitoring=Primitive.from_proto(
                resource.enable_integrity_monitoring
            ),
        )


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig.from_proto(
                i
            )
            for i in resources
        ]


class ClusterNetworkConfig(object):
    def __init__(
        self,
        network: str = None,
        subnetwork: str = None,
        enable_intra_node_visibility: bool = None,
        default_snat_status: dict = None,
        private_ipv6_google_access: str = None,
    ):
        self.network = network
        self.subnetwork = subnetwork
        self.enable_intra_node_visibility = enable_intra_node_visibility
        self.default_snat_status = default_snat_status
        self.private_ipv6_google_access = private_ipv6_google_access

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNetworkConfig()
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        if Primitive.to_proto(resource.subnetwork):
            res.subnetwork = Primitive.to_proto(resource.subnetwork)
        if Primitive.to_proto(resource.enable_intra_node_visibility):
            res.enable_intra_node_visibility = Primitive.to_proto(
                resource.enable_intra_node_visibility
            )
        if ClusterNetworkConfigDefaultSnatStatus.to_proto(resource.default_snat_status):
            res.default_snat_status.CopyFrom(
                ClusterNetworkConfigDefaultSnatStatus.to_proto(
                    resource.default_snat_status
                )
            )
        else:
            res.ClearField("default_snat_status")
        if ClusterNetworkConfigPrivateIPv6GoogleAccessEnum.to_proto(
            resource.private_ipv6_google_access
        ):
            res.private_ipv6_google_access = ClusterNetworkConfigPrivateIPv6GoogleAccessEnum.to_proto(
                resource.private_ipv6_google_access
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNetworkConfig(
            network=Primitive.from_proto(resource.network),
            subnetwork=Primitive.from_proto(resource.subnetwork),
            enable_intra_node_visibility=Primitive.from_proto(
                resource.enable_intra_node_visibility
            ),
            default_snat_status=ClusterNetworkConfigDefaultSnatStatus.from_proto(
                resource.default_snat_status
            ),
            private_ipv6_google_access=ClusterNetworkConfigPrivateIPv6GoogleAccessEnum.from_proto(
                resource.private_ipv6_google_access
            ),
        )


class ClusterNetworkConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNetworkConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNetworkConfig.from_proto(i) for i in resources]


class ClusterNetworkConfigDefaultSnatStatus(object):
    def __init__(self, disabled: bool = None):
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNetworkConfigDefaultSnatStatus()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNetworkConfigDefaultSnatStatus(
            disabled=Primitive.from_proto(resource.disabled),
        )


class ClusterNetworkConfigDefaultSnatStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNetworkConfigDefaultSnatStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNetworkConfigDefaultSnatStatus.from_proto(i) for i in resources]


class ClusterMaintenancePolicy(object):
    def __init__(self, window: dict = None, resource_version: str = None):
        self.window = window
        self.resource_version = resource_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMaintenancePolicy()
        if ClusterMaintenancePolicyWindow.to_proto(resource.window):
            res.window.CopyFrom(
                ClusterMaintenancePolicyWindow.to_proto(resource.window)
            )
        else:
            res.ClearField("window")
        if Primitive.to_proto(resource.resource_version):
            res.resource_version = Primitive.to_proto(resource.resource_version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMaintenancePolicy(
            window=ClusterMaintenancePolicyWindow.from_proto(resource.window),
            resource_version=Primitive.from_proto(resource.resource_version),
        )


class ClusterMaintenancePolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMaintenancePolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterMaintenancePolicy.from_proto(i) for i in resources]


class ClusterMaintenancePolicyWindow(object):
    def __init__(
        self,
        daily_maintenance_window: dict = None,
        recurring_window: dict = None,
        maintenance_exclusions: dict = None,
    ):
        self.daily_maintenance_window = daily_maintenance_window
        self.recurring_window = recurring_window
        self.maintenance_exclusions = maintenance_exclusions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMaintenancePolicyWindow()
        if ClusterMaintenancePolicyWindowDailyMaintenanceWindow.to_proto(
            resource.daily_maintenance_window
        ):
            res.daily_maintenance_window.CopyFrom(
                ClusterMaintenancePolicyWindowDailyMaintenanceWindow.to_proto(
                    resource.daily_maintenance_window
                )
            )
        else:
            res.ClearField("daily_maintenance_window")
        if ClusterMaintenancePolicyWindowRecurringWindow.to_proto(
            resource.recurring_window
        ):
            res.recurring_window.CopyFrom(
                ClusterMaintenancePolicyWindowRecurringWindow.to_proto(
                    resource.recurring_window
                )
            )
        else:
            res.ClearField("recurring_window")
        if Primitive.to_proto(resource.maintenance_exclusions):
            res.maintenance_exclusions = Primitive.to_proto(
                resource.maintenance_exclusions
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMaintenancePolicyWindow(
            daily_maintenance_window=ClusterMaintenancePolicyWindowDailyMaintenanceWindow.from_proto(
                resource.daily_maintenance_window
            ),
            recurring_window=ClusterMaintenancePolicyWindowRecurringWindow.from_proto(
                resource.recurring_window
            ),
            maintenance_exclusions=Primitive.from_proto(
                resource.maintenance_exclusions
            ),
        )


class ClusterMaintenancePolicyWindowArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMaintenancePolicyWindow.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterMaintenancePolicyWindow.from_proto(i) for i in resources]


class ClusterMaintenancePolicyWindowDailyMaintenanceWindow(object):
    def __init__(self, start_time: str = None, duration: str = None):
        self.start_time = start_time
        self.duration = duration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow()
        )
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.duration):
            res.duration = Primitive.to_proto(resource.duration)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMaintenancePolicyWindowDailyMaintenanceWindow(
            start_time=Primitive.from_proto(resource.start_time),
            duration=Primitive.from_proto(resource.duration),
        )


class ClusterMaintenancePolicyWindowDailyMaintenanceWindowArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterMaintenancePolicyWindowDailyMaintenanceWindow.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterMaintenancePolicyWindowDailyMaintenanceWindow.from_proto(i)
            for i in resources
        ]


class ClusterMaintenancePolicyWindowRecurringWindow(object):
    def __init__(self, window: dict = None, recurrence: str = None):
        self.window = window
        self.recurrence = recurrence

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMaintenancePolicyWindowRecurringWindow()
        if ClusterMaintenancePolicyWindowRecurringWindowWindow.to_proto(
            resource.window
        ):
            res.window.CopyFrom(
                ClusterMaintenancePolicyWindowRecurringWindowWindow.to_proto(
                    resource.window
                )
            )
        else:
            res.ClearField("window")
        if Primitive.to_proto(resource.recurrence):
            res.recurrence = Primitive.to_proto(resource.recurrence)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMaintenancePolicyWindowRecurringWindow(
            window=ClusterMaintenancePolicyWindowRecurringWindowWindow.from_proto(
                resource.window
            ),
            recurrence=Primitive.from_proto(resource.recurrence),
        )


class ClusterMaintenancePolicyWindowRecurringWindowArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterMaintenancePolicyWindowRecurringWindow.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterMaintenancePolicyWindowRecurringWindow.from_proto(i)
            for i in resources
        ]


class ClusterMaintenancePolicyWindowRecurringWindowWindow(object):
    def __init__(self, start_time: str = None, end_time: str = None):
        self.start_time = start_time
        self.end_time = end_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMaintenancePolicyWindowRecurringWindowWindow()
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.end_time):
            res.end_time = Primitive.to_proto(resource.end_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMaintenancePolicyWindowRecurringWindowWindow(
            start_time=Primitive.from_proto(resource.start_time),
            end_time=Primitive.from_proto(resource.end_time),
        )


class ClusterMaintenancePolicyWindowRecurringWindowWindowArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterMaintenancePolicyWindowRecurringWindowWindow.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterMaintenancePolicyWindowRecurringWindowWindow.from_proto(i)
            for i in resources
        ]


class ClusterDefaultMaxPodsConstraint(object):
    def __init__(self, max_pods_per_node: str = None):
        self.max_pods_per_node = max_pods_per_node

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterDefaultMaxPodsConstraint()
        if Primitive.to_proto(resource.max_pods_per_node):
            res.max_pods_per_node = Primitive.to_proto(resource.max_pods_per_node)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterDefaultMaxPodsConstraint(
            max_pods_per_node=Primitive.from_proto(resource.max_pods_per_node),
        )


class ClusterDefaultMaxPodsConstraintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterDefaultMaxPodsConstraint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterDefaultMaxPodsConstraint.from_proto(i) for i in resources]


class ClusterResourceUsageExportConfig(object):
    def __init__(
        self,
        bigquery_destination: dict = None,
        enable_network_egress_monitoring: bool = None,
        consumption_metering_config: dict = None,
        enable_network_egress_metering: bool = None,
    ):
        self.bigquery_destination = bigquery_destination
        self.enable_network_egress_monitoring = enable_network_egress_monitoring
        self.consumption_metering_config = consumption_metering_config
        self.enable_network_egress_metering = enable_network_egress_metering

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterResourceUsageExportConfig()
        if ClusterResourceUsageExportConfigBigqueryDestination.to_proto(
            resource.bigquery_destination
        ):
            res.bigquery_destination.CopyFrom(
                ClusterResourceUsageExportConfigBigqueryDestination.to_proto(
                    resource.bigquery_destination
                )
            )
        else:
            res.ClearField("bigquery_destination")
        if Primitive.to_proto(resource.enable_network_egress_monitoring):
            res.enable_network_egress_monitoring = Primitive.to_proto(
                resource.enable_network_egress_monitoring
            )
        if ClusterResourceUsageExportConfigConsumptionMeteringConfig.to_proto(
            resource.consumption_metering_config
        ):
            res.consumption_metering_config.CopyFrom(
                ClusterResourceUsageExportConfigConsumptionMeteringConfig.to_proto(
                    resource.consumption_metering_config
                )
            )
        else:
            res.ClearField("consumption_metering_config")
        if Primitive.to_proto(resource.enable_network_egress_metering):
            res.enable_network_egress_metering = Primitive.to_proto(
                resource.enable_network_egress_metering
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterResourceUsageExportConfig(
            bigquery_destination=ClusterResourceUsageExportConfigBigqueryDestination.from_proto(
                resource.bigquery_destination
            ),
            enable_network_egress_monitoring=Primitive.from_proto(
                resource.enable_network_egress_monitoring
            ),
            consumption_metering_config=ClusterResourceUsageExportConfigConsumptionMeteringConfig.from_proto(
                resource.consumption_metering_config
            ),
            enable_network_egress_metering=Primitive.from_proto(
                resource.enable_network_egress_metering
            ),
        )


class ClusterResourceUsageExportConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterResourceUsageExportConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterResourceUsageExportConfig.from_proto(i) for i in resources]


class ClusterResourceUsageExportConfigBigqueryDestination(object):
    def __init__(self, dataset_id: str = None):
        self.dataset_id = dataset_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterResourceUsageExportConfigBigqueryDestination()
        if Primitive.to_proto(resource.dataset_id):
            res.dataset_id = Primitive.to_proto(resource.dataset_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterResourceUsageExportConfigBigqueryDestination(
            dataset_id=Primitive.from_proto(resource.dataset_id),
        )


class ClusterResourceUsageExportConfigBigqueryDestinationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterResourceUsageExportConfigBigqueryDestination.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterResourceUsageExportConfigBigqueryDestination.from_proto(i)
            for i in resources
        ]


class ClusterResourceUsageExportConfigConsumptionMeteringConfig(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.ContainerClusterResourceUsageExportConfigConsumptionMeteringConfig()
        )
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterResourceUsageExportConfigConsumptionMeteringConfig(
            enabled=Primitive.from_proto(resource.enabled),
        )


class ClusterResourceUsageExportConfigConsumptionMeteringConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterResourceUsageExportConfigConsumptionMeteringConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterResourceUsageExportConfigConsumptionMeteringConfig.from_proto(i)
            for i in resources
        ]


class ClusterAuthenticatorGroupsConfig(object):
    def __init__(self, enabled: bool = None, security_group: str = None):
        self.enabled = enabled
        self.security_group = security_group

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAuthenticatorGroupsConfig()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.security_group):
            res.security_group = Primitive.to_proto(resource.security_group)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAuthenticatorGroupsConfig(
            enabled=Primitive.from_proto(resource.enabled),
            security_group=Primitive.from_proto(resource.security_group),
        )


class ClusterAuthenticatorGroupsConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAuthenticatorGroupsConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAuthenticatorGroupsConfig.from_proto(i) for i in resources]


class ClusterPrivateClusterConfig(object):
    def __init__(
        self,
        enable_private_nodes: bool = None,
        enable_private_endpoint: bool = None,
        master_ipv4_cidr_block: str = None,
        private_endpoint: str = None,
        public_endpoint: str = None,
        peering_name: str = None,
        master_global_access_config: dict = None,
    ):
        self.enable_private_nodes = enable_private_nodes
        self.enable_private_endpoint = enable_private_endpoint
        self.master_ipv4_cidr_block = master_ipv4_cidr_block
        self.private_endpoint = private_endpoint
        self.public_endpoint = public_endpoint
        self.peering_name = peering_name
        self.master_global_access_config = master_global_access_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterPrivateClusterConfig()
        if Primitive.to_proto(resource.enable_private_nodes):
            res.enable_private_nodes = Primitive.to_proto(resource.enable_private_nodes)
        if Primitive.to_proto(resource.enable_private_endpoint):
            res.enable_private_endpoint = Primitive.to_proto(
                resource.enable_private_endpoint
            )
        if Primitive.to_proto(resource.master_ipv4_cidr_block):
            res.master_ipv4_cidr_block = Primitive.to_proto(
                resource.master_ipv4_cidr_block
            )
        if Primitive.to_proto(resource.private_endpoint):
            res.private_endpoint = Primitive.to_proto(resource.private_endpoint)
        if Primitive.to_proto(resource.public_endpoint):
            res.public_endpoint = Primitive.to_proto(resource.public_endpoint)
        if Primitive.to_proto(resource.peering_name):
            res.peering_name = Primitive.to_proto(resource.peering_name)
        if ClusterPrivateClusterConfigMasterGlobalAccessConfig.to_proto(
            resource.master_global_access_config
        ):
            res.master_global_access_config.CopyFrom(
                ClusterPrivateClusterConfigMasterGlobalAccessConfig.to_proto(
                    resource.master_global_access_config
                )
            )
        else:
            res.ClearField("master_global_access_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterPrivateClusterConfig(
            enable_private_nodes=Primitive.from_proto(resource.enable_private_nodes),
            enable_private_endpoint=Primitive.from_proto(
                resource.enable_private_endpoint
            ),
            master_ipv4_cidr_block=Primitive.from_proto(
                resource.master_ipv4_cidr_block
            ),
            private_endpoint=Primitive.from_proto(resource.private_endpoint),
            public_endpoint=Primitive.from_proto(resource.public_endpoint),
            peering_name=Primitive.from_proto(resource.peering_name),
            master_global_access_config=ClusterPrivateClusterConfigMasterGlobalAccessConfig.from_proto(
                resource.master_global_access_config
            ),
        )


class ClusterPrivateClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterPrivateClusterConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterPrivateClusterConfig.from_proto(i) for i in resources]


class ClusterPrivateClusterConfigMasterGlobalAccessConfig(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterPrivateClusterConfigMasterGlobalAccessConfig(
            enabled=Primitive.from_proto(resource.enabled),
        )


class ClusterPrivateClusterConfigMasterGlobalAccessConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterPrivateClusterConfigMasterGlobalAccessConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterPrivateClusterConfigMasterGlobalAccessConfig.from_proto(i)
            for i in resources
        ]


class ClusterDatabaseEncryption(object):
    def __init__(self, state: str = None, key_name: str = None):
        self.state = state
        self.key_name = key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterDatabaseEncryption()
        if ClusterDatabaseEncryptionStateEnum.to_proto(resource.state):
            res.state = ClusterDatabaseEncryptionStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.key_name):
            res.key_name = Primitive.to_proto(resource.key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterDatabaseEncryption(
            state=ClusterDatabaseEncryptionStateEnum.from_proto(resource.state),
            key_name=Primitive.from_proto(resource.key_name),
        )


class ClusterDatabaseEncryptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterDatabaseEncryption.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterDatabaseEncryption.from_proto(i) for i in resources]


class ClusterVerticalPodAutoscaling(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterVerticalPodAutoscaling()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVerticalPodAutoscaling(
            enabled=Primitive.from_proto(resource.enabled),
        )


class ClusterVerticalPodAutoscalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterVerticalPodAutoscaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterVerticalPodAutoscaling.from_proto(i) for i in resources]


class ClusterShieldedNodes(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterShieldedNodes()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterShieldedNodes(enabled=Primitive.from_proto(resource.enabled),)


class ClusterShieldedNodesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterShieldedNodes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterShieldedNodes.from_proto(i) for i in resources]


class ClusterConditions(object):
    def __init__(
        self, code: str = None, message: str = None, canonical_code: str = None
    ):
        self.code = code
        self.message = message
        self.canonical_code = canonical_code

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterConditions()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if ClusterConditionsCanonicalCodeEnum.to_proto(resource.canonical_code):
            res.canonical_code = ClusterConditionsCanonicalCodeEnum.to_proto(
                resource.canonical_code
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConditions(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            canonical_code=ClusterConditionsCanonicalCodeEnum.from_proto(
                resource.canonical_code
            ),
        )


class ClusterConditionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConditions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConditions.from_proto(i) for i in resources]


class ClusterAutopilot(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAutopilot()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutopilot(enabled=Primitive.from_proto(resource.enabled),)


class ClusterAutopilotArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAutopilot.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAutopilot.from_proto(i) for i in resources]


class ClusterNodeConfig(object):
    def __init__(
        self,
        machine_type: str = None,
        disk_size_gb: int = None,
        oauth_scopes: list = None,
        service_account: str = None,
        metadata: dict = None,
        image_type: str = None,
        labels: dict = None,
        local_ssd_count: int = None,
        tags: list = None,
        preemptible: bool = None,
        accelerators: list = None,
        disk_type: str = None,
        min_cpu_platform: str = None,
        workload_metadata_config: dict = None,
        taints: list = None,
        sandbox_config: dict = None,
        node_group: str = None,
        reservation_affinity: dict = None,
        shielded_instance_config: dict = None,
        linux_node_config: dict = None,
        kubelet_config: dict = None,
        boot_disk_kms_key: str = None,
    ):
        self.machine_type = machine_type
        self.disk_size_gb = disk_size_gb
        self.oauth_scopes = oauth_scopes
        self.service_account = service_account
        self.metadata = metadata
        self.image_type = image_type
        self.labels = labels
        self.local_ssd_count = local_ssd_count
        self.tags = tags
        self.preemptible = preemptible
        self.accelerators = accelerators
        self.disk_type = disk_type
        self.min_cpu_platform = min_cpu_platform
        self.workload_metadata_config = workload_metadata_config
        self.taints = taints
        self.sandbox_config = sandbox_config
        self.node_group = node_group
        self.reservation_affinity = reservation_affinity
        self.shielded_instance_config = shielded_instance_config
        self.linux_node_config = linux_node_config
        self.kubelet_config = kubelet_config
        self.boot_disk_kms_key = boot_disk_kms_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodeConfig()
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if Primitive.to_proto(resource.disk_size_gb):
            res.disk_size_gb = Primitive.to_proto(resource.disk_size_gb)
        if Primitive.to_proto(resource.oauth_scopes):
            res.oauth_scopes.extend(Primitive.to_proto(resource.oauth_scopes))
        if Primitive.to_proto(resource.service_account):
            res.service_account = Primitive.to_proto(resource.service_account)
        if Primitive.to_proto(resource.metadata):
            res.metadata = Primitive.to_proto(resource.metadata)
        if Primitive.to_proto(resource.image_type):
            res.image_type = Primitive.to_proto(resource.image_type)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if Primitive.to_proto(resource.local_ssd_count):
            res.local_ssd_count = Primitive.to_proto(resource.local_ssd_count)
        if Primitive.to_proto(resource.tags):
            res.tags.extend(Primitive.to_proto(resource.tags))
        if Primitive.to_proto(resource.preemptible):
            res.preemptible = Primitive.to_proto(resource.preemptible)
        if ClusterNodeConfigAcceleratorsArray.to_proto(resource.accelerators):
            res.accelerators.extend(
                ClusterNodeConfigAcceleratorsArray.to_proto(resource.accelerators)
            )
        if Primitive.to_proto(resource.disk_type):
            res.disk_type = Primitive.to_proto(resource.disk_type)
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        if ClusterNodeConfigWorkloadMetadataConfig.to_proto(
            resource.workload_metadata_config
        ):
            res.workload_metadata_config.CopyFrom(
                ClusterNodeConfigWorkloadMetadataConfig.to_proto(
                    resource.workload_metadata_config
                )
            )
        else:
            res.ClearField("workload_metadata_config")
        if ClusterNodeConfigTaintsArray.to_proto(resource.taints):
            res.taints.extend(ClusterNodeConfigTaintsArray.to_proto(resource.taints))
        if ClusterNodeConfigSandboxConfig.to_proto(resource.sandbox_config):
            res.sandbox_config.CopyFrom(
                ClusterNodeConfigSandboxConfig.to_proto(resource.sandbox_config)
            )
        else:
            res.ClearField("sandbox_config")
        if Primitive.to_proto(resource.node_group):
            res.node_group = Primitive.to_proto(resource.node_group)
        if ClusterNodeConfigReservationAffinity.to_proto(resource.reservation_affinity):
            res.reservation_affinity.CopyFrom(
                ClusterNodeConfigReservationAffinity.to_proto(
                    resource.reservation_affinity
                )
            )
        else:
            res.ClearField("reservation_affinity")
        if ClusterNodeConfigShieldedInstanceConfig.to_proto(
            resource.shielded_instance_config
        ):
            res.shielded_instance_config.CopyFrom(
                ClusterNodeConfigShieldedInstanceConfig.to_proto(
                    resource.shielded_instance_config
                )
            )
        else:
            res.ClearField("shielded_instance_config")
        if ClusterNodeConfigLinuxNodeConfig.to_proto(resource.linux_node_config):
            res.linux_node_config.CopyFrom(
                ClusterNodeConfigLinuxNodeConfig.to_proto(resource.linux_node_config)
            )
        else:
            res.ClearField("linux_node_config")
        if ClusterNodeConfigKubeletConfig.to_proto(resource.kubelet_config):
            res.kubelet_config.CopyFrom(
                ClusterNodeConfigKubeletConfig.to_proto(resource.kubelet_config)
            )
        else:
            res.ClearField("kubelet_config")
        if Primitive.to_proto(resource.boot_disk_kms_key):
            res.boot_disk_kms_key = Primitive.to_proto(resource.boot_disk_kms_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodeConfig(
            machine_type=Primitive.from_proto(resource.machine_type),
            disk_size_gb=Primitive.from_proto(resource.disk_size_gb),
            oauth_scopes=Primitive.from_proto(resource.oauth_scopes),
            service_account=Primitive.from_proto(resource.service_account),
            metadata=Primitive.from_proto(resource.metadata),
            image_type=Primitive.from_proto(resource.image_type),
            labels=Primitive.from_proto(resource.labels),
            local_ssd_count=Primitive.from_proto(resource.local_ssd_count),
            tags=Primitive.from_proto(resource.tags),
            preemptible=Primitive.from_proto(resource.preemptible),
            accelerators=ClusterNodeConfigAcceleratorsArray.from_proto(
                resource.accelerators
            ),
            disk_type=Primitive.from_proto(resource.disk_type),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
            workload_metadata_config=ClusterNodeConfigWorkloadMetadataConfig.from_proto(
                resource.workload_metadata_config
            ),
            taints=ClusterNodeConfigTaintsArray.from_proto(resource.taints),
            sandbox_config=ClusterNodeConfigSandboxConfig.from_proto(
                resource.sandbox_config
            ),
            node_group=Primitive.from_proto(resource.node_group),
            reservation_affinity=ClusterNodeConfigReservationAffinity.from_proto(
                resource.reservation_affinity
            ),
            shielded_instance_config=ClusterNodeConfigShieldedInstanceConfig.from_proto(
                resource.shielded_instance_config
            ),
            linux_node_config=ClusterNodeConfigLinuxNodeConfig.from_proto(
                resource.linux_node_config
            ),
            kubelet_config=ClusterNodeConfigKubeletConfig.from_proto(
                resource.kubelet_config
            ),
            boot_disk_kms_key=Primitive.from_proto(resource.boot_disk_kms_key),
        )


class ClusterNodeConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodeConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodeConfig.from_proto(i) for i in resources]


class ClusterNodeConfigAccelerators(object):
    def __init__(self, accelerator_count: int = None, accelerator_type: str = None):
        self.accelerator_count = accelerator_count
        self.accelerator_type = accelerator_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodeConfigAccelerators()
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodeConfigAccelerators(
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
        )


class ClusterNodeConfigAcceleratorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodeConfigAccelerators.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodeConfigAccelerators.from_proto(i) for i in resources]


class ClusterNodeConfigWorkloadMetadataConfig(object):
    def __init__(self, mode: str = None):
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodeConfigWorkloadMetadataConfig()
        if ClusterNodeConfigWorkloadMetadataConfigModeEnum.to_proto(resource.mode):
            res.mode = ClusterNodeConfigWorkloadMetadataConfigModeEnum.to_proto(
                resource.mode
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodeConfigWorkloadMetadataConfig(
            mode=ClusterNodeConfigWorkloadMetadataConfigModeEnum.from_proto(
                resource.mode
            ),
        )


class ClusterNodeConfigWorkloadMetadataConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodeConfigWorkloadMetadataConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterNodeConfigWorkloadMetadataConfig.from_proto(i) for i in resources
        ]


class ClusterNodeConfigTaints(object):
    def __init__(self, key: str = None, value: str = None, effect: str = None):
        self.key = key
        self.value = value
        self.effect = effect

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodeConfigTaints()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if ClusterNodeConfigTaintsEffectEnum.to_proto(resource.effect):
            res.effect = ClusterNodeConfigTaintsEffectEnum.to_proto(resource.effect)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodeConfigTaints(
            key=Primitive.from_proto(resource.key),
            value=Primitive.from_proto(resource.value),
            effect=ClusterNodeConfigTaintsEffectEnum.from_proto(resource.effect),
        )


class ClusterNodeConfigTaintsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodeConfigTaints.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodeConfigTaints.from_proto(i) for i in resources]


class ClusterNodeConfigSandboxConfig(object):
    def __init__(self, type: str = None):
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodeConfigSandboxConfig()
        if ClusterNodeConfigSandboxConfigTypeEnum.to_proto(resource.type):
            res.type = ClusterNodeConfigSandboxConfigTypeEnum.to_proto(resource.type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodeConfigSandboxConfig(
            type=ClusterNodeConfigSandboxConfigTypeEnum.from_proto(resource.type),
        )


class ClusterNodeConfigSandboxConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodeConfigSandboxConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodeConfigSandboxConfig.from_proto(i) for i in resources]


class ClusterNodeConfigReservationAffinity(object):
    def __init__(
        self, consume_reservation_type: str = None, key: str = None, values: list = None
    ):
        self.consume_reservation_type = consume_reservation_type
        self.key = key
        self.values = values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodeConfigReservationAffinity()
        if ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
            resource.consume_reservation_type
        ):
            res.consume_reservation_type = ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
                resource.consume_reservation_type
            )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.values):
            res.values.extend(Primitive.to_proto(resource.values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodeConfigReservationAffinity(
            consume_reservation_type=ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum.from_proto(
                resource.consume_reservation_type
            ),
            key=Primitive.from_proto(resource.key),
            values=Primitive.from_proto(resource.values),
        )


class ClusterNodeConfigReservationAffinityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodeConfigReservationAffinity.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodeConfigReservationAffinity.from_proto(i) for i in resources]


class ClusterNodeConfigShieldedInstanceConfig(object):
    def __init__(
        self, enable_secure_boot: bool = None, enable_integrity_monitoring: bool = None
    ):
        self.enable_secure_boot = enable_secure_boot
        self.enable_integrity_monitoring = enable_integrity_monitoring

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodeConfigShieldedInstanceConfig()
        if Primitive.to_proto(resource.enable_secure_boot):
            res.enable_secure_boot = Primitive.to_proto(resource.enable_secure_boot)
        if Primitive.to_proto(resource.enable_integrity_monitoring):
            res.enable_integrity_monitoring = Primitive.to_proto(
                resource.enable_integrity_monitoring
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodeConfigShieldedInstanceConfig(
            enable_secure_boot=Primitive.from_proto(resource.enable_secure_boot),
            enable_integrity_monitoring=Primitive.from_proto(
                resource.enable_integrity_monitoring
            ),
        )


class ClusterNodeConfigShieldedInstanceConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodeConfigShieldedInstanceConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterNodeConfigShieldedInstanceConfig.from_proto(i) for i in resources
        ]


class ClusterNodeConfigLinuxNodeConfig(object):
    def __init__(self, sysctls: dict = None):
        self.sysctls = sysctls

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodeConfigLinuxNodeConfig()
        if Primitive.to_proto(resource.sysctls):
            res.sysctls = Primitive.to_proto(resource.sysctls)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodeConfigLinuxNodeConfig(
            sysctls=Primitive.from_proto(resource.sysctls),
        )


class ClusterNodeConfigLinuxNodeConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodeConfigLinuxNodeConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodeConfigLinuxNodeConfig.from_proto(i) for i in resources]


class ClusterNodeConfigKubeletConfig(object):
    def __init__(
        self,
        cpu_manager_policy: str = None,
        cpu_cfs_quota: bool = None,
        cpu_cfs_quota_period: str = None,
    ):
        self.cpu_manager_policy = cpu_manager_policy
        self.cpu_cfs_quota = cpu_cfs_quota
        self.cpu_cfs_quota_period = cpu_cfs_quota_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodeConfigKubeletConfig()
        if Primitive.to_proto(resource.cpu_manager_policy):
            res.cpu_manager_policy = Primitive.to_proto(resource.cpu_manager_policy)
        if Primitive.to_proto(resource.cpu_cfs_quota):
            res.cpu_cfs_quota = Primitive.to_proto(resource.cpu_cfs_quota)
        if Primitive.to_proto(resource.cpu_cfs_quota_period):
            res.cpu_cfs_quota_period = Primitive.to_proto(resource.cpu_cfs_quota_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodeConfigKubeletConfig(
            cpu_manager_policy=Primitive.from_proto(resource.cpu_manager_policy),
            cpu_cfs_quota=Primitive.from_proto(resource.cpu_cfs_quota),
            cpu_cfs_quota_period=Primitive.from_proto(resource.cpu_cfs_quota_period),
        )


class ClusterNodeConfigKubeletConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodeConfigKubeletConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodeConfigKubeletConfig.from_proto(i) for i in resources]


class ClusterReleaseChannel(object):
    def __init__(self, channel: str = None):
        self.channel = channel

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterReleaseChannel()
        if ClusterReleaseChannelChannelEnum.to_proto(resource.channel):
            res.channel = ClusterReleaseChannelChannelEnum.to_proto(resource.channel)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterReleaseChannel(
            channel=ClusterReleaseChannelChannelEnum.from_proto(resource.channel),
        )


class ClusterReleaseChannelArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterReleaseChannel.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterReleaseChannel.from_proto(i) for i in resources]


class ClusterWorkloadIdentityConfig(object):
    def __init__(self, workload_pool: str = None):
        self.workload_pool = workload_pool

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterWorkloadIdentityConfig()
        if Primitive.to_proto(resource.workload_pool):
            res.workload_pool = Primitive.to_proto(resource.workload_pool)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterWorkloadIdentityConfig(
            workload_pool=Primitive.from_proto(resource.workload_pool),
        )


class ClusterWorkloadIdentityConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterWorkloadIdentityConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterWorkloadIdentityConfig.from_proto(i) for i in resources]


class ClusterNotificationConfig(object):
    def __init__(self, pubsub: dict = None):
        self.pubsub = pubsub

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNotificationConfig()
        if ClusterNotificationConfigPubsub.to_proto(resource.pubsub):
            res.pubsub.CopyFrom(
                ClusterNotificationConfigPubsub.to_proto(resource.pubsub)
            )
        else:
            res.ClearField("pubsub")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNotificationConfig(
            pubsub=ClusterNotificationConfigPubsub.from_proto(resource.pubsub),
        )


class ClusterNotificationConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNotificationConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNotificationConfig.from_proto(i) for i in resources]


class ClusterNotificationConfigPubsub(object):
    def __init__(self, enabled: bool = None, topic: str = None):
        self.enabled = enabled
        self.topic = topic

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNotificationConfigPubsub()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.topic):
            res.topic = Primitive.to_proto(resource.topic)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNotificationConfigPubsub(
            enabled=Primitive.from_proto(resource.enabled),
            topic=Primitive.from_proto(resource.topic),
        )


class ClusterNotificationConfigPubsubArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNotificationConfigPubsub.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNotificationConfigPubsub.from_proto(i) for i in resources]


class ClusterConfidentialNodes(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterConfidentialNodes()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfidentialNodes(enabled=Primitive.from_proto(resource.enabled),)


class ClusterConfidentialNodesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfidentialNodes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfidentialNodes.from_proto(i) for i in resources]


class ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum.Value(
            "ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum.Name(
            resource
        )[
            len("ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum") :
        ]


class ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum.Value(
            "ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum.Name(
            resource
        )[
            len("ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum") :
        ]


class ClusterNodePoolsConfigTaintsEffectEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConfigTaintsEffectEnum.Value(
            "ContainerClusterNodePoolsConfigTaintsEffectEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConfigTaintsEffectEnum.Name(
            resource
        )[len("ContainerClusterNodePoolsConfigTaintsEffectEnum") :]


class ClusterNodePoolsConfigSandboxConfigTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConfigSandboxConfigTypeEnum.Value(
            "ContainerClusterNodePoolsConfigSandboxConfigTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConfigSandboxConfigTypeEnum.Name(
            resource
        )[len("ContainerClusterNodePoolsConfigSandboxConfigTypeEnum") :]


class ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum.Value(
            "ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum.Name(
            resource
        )[
            len(
                "ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum"
            ) :
        ]


class ClusterNodePoolsStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsStatusEnum.Value(
            "ContainerClusterNodePoolsStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsStatusEnum.Name(resource)[
            len("ContainerClusterNodePoolsStatusEnum") :
        ]


class ClusterNodePoolsConditionsCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConditionsCodeEnum.Value(
            "ContainerClusterNodePoolsConditionsCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConditionsCodeEnum.Name(resource)[
            len("ContainerClusterNodePoolsConditionsCodeEnum") :
        ]


class ClusterNodePoolsConditionsCanonicalCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConditionsCanonicalCodeEnum.Value(
            "ContainerClusterNodePoolsConditionsCanonicalCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodePoolsConditionsCanonicalCodeEnum.Name(
            resource
        )[len("ContainerClusterNodePoolsConditionsCanonicalCodeEnum") :]


class ClusterNetworkPolicyProviderEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNetworkPolicyProviderEnum.Value(
            "ContainerClusterNetworkPolicyProviderEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNetworkPolicyProviderEnum.Name(resource)[
            len("ContainerClusterNetworkPolicyProviderEnum") :
        ]


class ClusterNetworkConfigPrivateIPv6GoogleAccessEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum.Value(
            "ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum.Name(
            resource
        )[
            len("ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum") :
        ]


class ClusterDatabaseEncryptionStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterDatabaseEncryptionStateEnum.Value(
            "ContainerClusterDatabaseEncryptionStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterDatabaseEncryptionStateEnum.Name(resource)[
            len("ContainerClusterDatabaseEncryptionStateEnum") :
        ]


class ClusterConditionsCanonicalCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterConditionsCanonicalCodeEnum.Value(
            "ContainerClusterConditionsCanonicalCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterConditionsCanonicalCodeEnum.Name(resource)[
            len("ContainerClusterConditionsCanonicalCodeEnum") :
        ]


class ClusterNodeConfigWorkloadMetadataConfigModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum.Value(
            "ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum.Name(
            resource
        )[
            len("ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum") :
        ]


class ClusterNodeConfigTaintsEffectEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodeConfigTaintsEffectEnum.Value(
            "ContainerClusterNodeConfigTaintsEffectEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodeConfigTaintsEffectEnum.Name(resource)[
            len("ContainerClusterNodeConfigTaintsEffectEnum") :
        ]


class ClusterNodeConfigSandboxConfigTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodeConfigSandboxConfigTypeEnum.Value(
            "ContainerClusterNodeConfigSandboxConfigTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodeConfigSandboxConfigTypeEnum.Name(
            resource
        )[len("ContainerClusterNodeConfigSandboxConfigTypeEnum") :]


class ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum.Value(
            "ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum.Name(
            resource
        )[
            len(
                "ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum"
            ) :
        ]


class ClusterReleaseChannelChannelEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterReleaseChannelChannelEnum.Value(
            "ContainerClusterReleaseChannelChannelEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterReleaseChannelChannelEnum.Name(resource)[
            len("ContainerClusterReleaseChannelChannelEnum") :
        ]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
