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
from google3.cloud.graphite.mmv2.services.google.container import node_pool_pb2
from google3.cloud.graphite.mmv2.services.google.container import node_pool_pb2_grpc

from typing import List


class NodePool(object):
    def __init__(
        self,
        name: str = None,
        config: dict = None,
        node_count: int = None,
        version: str = None,
        status: str = None,
        status_message: str = None,
        locations: list = None,
        autoscaling: dict = None,
        management: dict = None,
        max_pods_constraint: dict = None,
        conditions: list = None,
        pod_ipv4_cidr_size: int = None,
        upgrade_settings: dict = None,
        cluster: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.config = config
        self.node_count = node_count
        self.version = version
        self.locations = locations
        self.autoscaling = autoscaling
        self.management = management
        self.max_pods_constraint = max_pods_constraint
        self.upgrade_settings = upgrade_settings
        self.cluster = cluster
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = node_pool_pb2_grpc.ContainerBetaNodePoolServiceStub(channel.Channel())
        request = node_pool_pb2.ApplyContainerBetaNodePoolRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if NodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(NodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.node_count):
            request.resource.node_count = Primitive.to_proto(self.node_count)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if Primitive.to_proto(self.locations):
            request.resource.locations.extend(Primitive.to_proto(self.locations))
        if NodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                NodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if NodePoolManagement.to_proto(self.management):
            request.resource.management.CopyFrom(
                NodePoolManagement.to_proto(self.management)
            )
        else:
            request.resource.ClearField("management")
        if NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if NodePoolUpgradeSettings.to_proto(self.upgrade_settings):
            request.resource.upgrade_settings.CopyFrom(
                NodePoolUpgradeSettings.to_proto(self.upgrade_settings)
            )
        else:
            request.resource.ClearField("upgrade_settings")
        if Primitive.to_proto(self.cluster):
            request.resource.cluster = Primitive.to_proto(self.cluster)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyContainerBetaNodePool(request)
        self.name = Primitive.from_proto(response.name)
        self.config = NodePoolConfig.from_proto(response.config)
        self.node_count = Primitive.from_proto(response.node_count)
        self.version = Primitive.from_proto(response.version)
        self.status = Primitive.from_proto(response.status)
        self.status_message = Primitive.from_proto(response.status_message)
        self.locations = Primitive.from_proto(response.locations)
        self.autoscaling = NodePoolAutoscaling.from_proto(response.autoscaling)
        self.management = NodePoolManagement.from_proto(response.management)
        self.max_pods_constraint = NodePoolMaxPodsConstraint.from_proto(
            response.max_pods_constraint
        )
        self.conditions = NodePoolConditionsArray.from_proto(response.conditions)
        self.pod_ipv4_cidr_size = Primitive.from_proto(response.pod_ipv4_cidr_size)
        self.upgrade_settings = NodePoolUpgradeSettings.from_proto(
            response.upgrade_settings
        )
        self.cluster = Primitive.from_proto(response.cluster)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = node_pool_pb2_grpc.ContainerBetaNodePoolServiceStub(channel.Channel())
        request = node_pool_pb2.DeleteContainerBetaNodePoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if NodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(NodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.node_count):
            request.resource.node_count = Primitive.to_proto(self.node_count)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if Primitive.to_proto(self.locations):
            request.resource.locations.extend(Primitive.to_proto(self.locations))
        if NodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                NodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if NodePoolManagement.to_proto(self.management):
            request.resource.management.CopyFrom(
                NodePoolManagement.to_proto(self.management)
            )
        else:
            request.resource.ClearField("management")
        if NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if NodePoolUpgradeSettings.to_proto(self.upgrade_settings):
            request.resource.upgrade_settings.CopyFrom(
                NodePoolUpgradeSettings.to_proto(self.upgrade_settings)
            )
        else:
            request.resource.ClearField("upgrade_settings")
        if Primitive.to_proto(self.cluster):
            request.resource.cluster = Primitive.to_proto(self.cluster)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteContainerBetaNodePool(request)

    @classmethod
    def list(self, project, location, cluster, service_account_file=""):
        stub = node_pool_pb2_grpc.ContainerBetaNodePoolServiceStub(channel.Channel())
        request = node_pool_pb2.ListContainerBetaNodePoolRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Cluster = cluster

        return stub.ListContainerBetaNodePool(request).items

    def to_proto(self):
        resource = node_pool_pb2.ContainerBetaNodePool()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if NodePoolConfig.to_proto(self.config):
            resource.config.CopyFrom(NodePoolConfig.to_proto(self.config))
        else:
            resource.ClearField("config")
        if Primitive.to_proto(self.node_count):
            resource.node_count = Primitive.to_proto(self.node_count)
        if Primitive.to_proto(self.version):
            resource.version = Primitive.to_proto(self.version)
        if Primitive.to_proto(self.locations):
            resource.locations.extend(Primitive.to_proto(self.locations))
        if NodePoolAutoscaling.to_proto(self.autoscaling):
            resource.autoscaling.CopyFrom(
                NodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            resource.ClearField("autoscaling")
        if NodePoolManagement.to_proto(self.management):
            resource.management.CopyFrom(NodePoolManagement.to_proto(self.management))
        else:
            resource.ClearField("management")
        if NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            resource.max_pods_constraint.CopyFrom(
                NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            resource.ClearField("max_pods_constraint")
        if NodePoolUpgradeSettings.to_proto(self.upgrade_settings):
            resource.upgrade_settings.CopyFrom(
                NodePoolUpgradeSettings.to_proto(self.upgrade_settings)
            )
        else:
            resource.ClearField("upgrade_settings")
        if Primitive.to_proto(self.cluster):
            resource.cluster = Primitive.to_proto(self.cluster)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class NodePoolConfig(object):
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
        taints: list = None,
        sandbox_config: dict = None,
        reservation_affinity: dict = None,
        shielded_instance_config: dict = None,
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
        self.taints = taints
        self.sandbox_config = sandbox_config
        self.reservation_affinity = reservation_affinity
        self.shielded_instance_config = shielded_instance_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerBetaNodePoolConfig()
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
        if NodePoolConfigAcceleratorsArray.to_proto(resource.accelerators):
            res.accelerators.extend(
                NodePoolConfigAcceleratorsArray.to_proto(resource.accelerators)
            )
        if Primitive.to_proto(resource.disk_type):
            res.disk_type = Primitive.to_proto(resource.disk_type)
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        if NodePoolConfigTaintsArray.to_proto(resource.taints):
            res.taints.extend(NodePoolConfigTaintsArray.to_proto(resource.taints))
        if NodePoolConfigSandboxConfig.to_proto(resource.sandbox_config):
            res.sandbox_config.CopyFrom(
                NodePoolConfigSandboxConfig.to_proto(resource.sandbox_config)
            )
        else:
            res.ClearField("sandbox_config")
        if NodePoolConfigReservationAffinity.to_proto(resource.reservation_affinity):
            res.reservation_affinity.CopyFrom(
                NodePoolConfigReservationAffinity.to_proto(
                    resource.reservation_affinity
                )
            )
        else:
            res.ClearField("reservation_affinity")
        if NodePoolConfigShieldedInstanceConfig.to_proto(
            resource.shielded_instance_config
        ):
            res.shielded_instance_config.CopyFrom(
                NodePoolConfigShieldedInstanceConfig.to_proto(
                    resource.shielded_instance_config
                )
            )
        else:
            res.ClearField("shielded_instance_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfig(
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
            accelerators=NodePoolConfigAcceleratorsArray.from_proto(
                resource.accelerators
            ),
            disk_type=Primitive.from_proto(resource.disk_type),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
            taints=NodePoolConfigTaintsArray.from_proto(resource.taints),
            sandbox_config=NodePoolConfigSandboxConfig.from_proto(
                resource.sandbox_config
            ),
            reservation_affinity=NodePoolConfigReservationAffinity.from_proto(
                resource.reservation_affinity
            ),
            shielded_instance_config=NodePoolConfigShieldedInstanceConfig.from_proto(
                resource.shielded_instance_config
            ),
        )


class NodePoolConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfig.from_proto(i) for i in resources]


class NodePoolConfigAccelerators(object):
    def __init__(self, accelerator_count: int = None, accelerator_type: str = None):
        self.accelerator_count = accelerator_count
        self.accelerator_type = accelerator_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerBetaNodePoolConfigAccelerators()
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigAccelerators(
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
        )


class NodePoolConfigAcceleratorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigAccelerators.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigAccelerators.from_proto(i) for i in resources]


class NodePoolConfigTaints(object):
    def __init__(self, key: str = None, value: str = None, effect: str = None):
        self.key = key
        self.value = value
        self.effect = effect

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerBetaNodePoolConfigTaints()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if Primitive.to_proto(resource.effect):
            res.effect = Primitive.to_proto(resource.effect)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigTaints(
            key=Primitive.from_proto(resource.key),
            value=Primitive.from_proto(resource.value),
            effect=Primitive.from_proto(resource.effect),
        )


class NodePoolConfigTaintsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigTaints.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigTaints.from_proto(i) for i in resources]


class NodePoolConfigSandboxConfig(object):
    def __init__(self, type: str = None):
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerBetaNodePoolConfigSandboxConfig()
        if NodePoolConfigSandboxConfigTypeEnum.to_proto(resource.type):
            res.type = NodePoolConfigSandboxConfigTypeEnum.to_proto(resource.type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigSandboxConfig(
            type=NodePoolConfigSandboxConfigTypeEnum.from_proto(resource.type),
        )


class NodePoolConfigSandboxConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigSandboxConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigSandboxConfig.from_proto(i) for i in resources]


class NodePoolConfigReservationAffinity(object):
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

        res = node_pool_pb2.ContainerBetaNodePoolConfigReservationAffinity()
        if NodePoolConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
            resource.consume_reservation_type
        ):
            res.consume_reservation_type = NodePoolConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
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

        return NodePoolConfigReservationAffinity(
            consume_reservation_type=NodePoolConfigReservationAffinityConsumeReservationTypeEnum.from_proto(
                resource.consume_reservation_type
            ),
            key=Primitive.from_proto(resource.key),
            values=Primitive.from_proto(resource.values),
        )


class NodePoolConfigReservationAffinityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigReservationAffinity.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigReservationAffinity.from_proto(i) for i in resources]


class NodePoolConfigShieldedInstanceConfig(object):
    def __init__(
        self, enable_secure_boot: bool = None, enable_integrity_monitoring: bool = None
    ):
        self.enable_secure_boot = enable_secure_boot
        self.enable_integrity_monitoring = enable_integrity_monitoring

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerBetaNodePoolConfigShieldedInstanceConfig()
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

        return NodePoolConfigShieldedInstanceConfig(
            enable_secure_boot=Primitive.from_proto(resource.enable_secure_boot),
            enable_integrity_monitoring=Primitive.from_proto(
                resource.enable_integrity_monitoring
            ),
        )


class NodePoolConfigShieldedInstanceConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigShieldedInstanceConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigShieldedInstanceConfig.from_proto(i) for i in resources]


class NodePoolAutoscaling(object):
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

        res = node_pool_pb2.ContainerBetaNodePoolAutoscaling()
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

        return NodePoolAutoscaling(
            enabled=Primitive.from_proto(resource.enabled),
            min_node_count=Primitive.from_proto(resource.min_node_count),
            max_node_count=Primitive.from_proto(resource.max_node_count),
            autoprovisioned=Primitive.from_proto(resource.autoprovisioned),
        )


class NodePoolAutoscalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolAutoscaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolAutoscaling.from_proto(i) for i in resources]


class NodePoolManagement(object):
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

        res = node_pool_pb2.ContainerBetaNodePoolManagement()
        if Primitive.to_proto(resource.auto_upgrade):
            res.auto_upgrade = Primitive.to_proto(resource.auto_upgrade)
        if Primitive.to_proto(resource.auto_repair):
            res.auto_repair = Primitive.to_proto(resource.auto_repair)
        if NodePoolManagementUpgradeOptions.to_proto(resource.upgrade_options):
            res.upgrade_options.CopyFrom(
                NodePoolManagementUpgradeOptions.to_proto(resource.upgrade_options)
            )
        else:
            res.ClearField("upgrade_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolManagement(
            auto_upgrade=Primitive.from_proto(resource.auto_upgrade),
            auto_repair=Primitive.from_proto(resource.auto_repair),
            upgrade_options=NodePoolManagementUpgradeOptions.from_proto(
                resource.upgrade_options
            ),
        )


class NodePoolManagementArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolManagement.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolManagement.from_proto(i) for i in resources]


class NodePoolManagementUpgradeOptions(object):
    def __init__(self, auto_upgrade_start_time: str = None, description: str = None):
        self.auto_upgrade_start_time = auto_upgrade_start_time
        self.description = description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerBetaNodePoolManagementUpgradeOptions()
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

        return NodePoolManagementUpgradeOptions(
            auto_upgrade_start_time=Primitive.from_proto(
                resource.auto_upgrade_start_time
            ),
            description=Primitive.from_proto(resource.description),
        )


class NodePoolManagementUpgradeOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolManagementUpgradeOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolManagementUpgradeOptions.from_proto(i) for i in resources]


class NodePoolMaxPodsConstraint(object):
    def __init__(self, max_pods_per_node: int = None):
        self.max_pods_per_node = max_pods_per_node

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerBetaNodePoolMaxPodsConstraint()
        if Primitive.to_proto(resource.max_pods_per_node):
            res.max_pods_per_node = Primitive.to_proto(resource.max_pods_per_node)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolMaxPodsConstraint(
            max_pods_per_node=Primitive.from_proto(resource.max_pods_per_node),
        )


class NodePoolMaxPodsConstraintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolMaxPodsConstraint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolMaxPodsConstraint.from_proto(i) for i in resources]


class NodePoolConditions(object):
    def __init__(self, code: str = None, message: str = None):
        self.code = code
        self.message = message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerBetaNodePoolConditions()
        if NodePoolConditionsCodeEnum.to_proto(resource.code):
            res.code = NodePoolConditionsCodeEnum.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConditions(
            code=NodePoolConditionsCodeEnum.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
        )


class NodePoolConditionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConditions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConditions.from_proto(i) for i in resources]


class NodePoolUpgradeSettings(object):
    def __init__(self, max_surge: int = None, max_unavailable: int = None):
        self.max_surge = max_surge
        self.max_unavailable = max_unavailable

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerBetaNodePoolUpgradeSettings()
        if Primitive.to_proto(resource.max_surge):
            res.max_surge = Primitive.to_proto(resource.max_surge)
        if Primitive.to_proto(resource.max_unavailable):
            res.max_unavailable = Primitive.to_proto(resource.max_unavailable)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolUpgradeSettings(
            max_surge=Primitive.from_proto(resource.max_surge),
            max_unavailable=Primitive.from_proto(resource.max_unavailable),
        )


class NodePoolUpgradeSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolUpgradeSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolUpgradeSettings.from_proto(i) for i in resources]


class NodePoolConfigSandboxConfigTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerBetaNodePoolConfigSandboxConfigTypeEnum.Value(
            "ContainerBetaNodePoolConfigSandboxConfigTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerBetaNodePoolConfigSandboxConfigTypeEnum.Name(
            resource
        )[len("ContainerBetaNodePoolConfigSandboxConfigTypeEnum") :]


class NodePoolConfigReservationAffinityConsumeReservationTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum.Value(
            "ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum.Name(
            resource
        )[
            len(
                "ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum"
            ) :
        ]


class NodePoolConditionsCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerBetaNodePoolConditionsCodeEnum.Value(
            "ContainerBetaNodePoolConditionsCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerBetaNodePoolConditionsCodeEnum.Name(resource)[
            len("ContainerBetaNodePoolConditionsCodeEnum") :
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
