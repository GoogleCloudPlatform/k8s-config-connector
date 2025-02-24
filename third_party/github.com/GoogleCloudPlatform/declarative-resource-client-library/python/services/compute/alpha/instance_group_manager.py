# Copyright 2024 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.compute import (
    instance_group_manager_pb2,
)
from google3.cloud.graphite.mmv2.services.google.compute import (
    instance_group_manager_pb2_grpc,
)

from typing import List


class InstanceGroupManager(object):
    def __init__(
        self,
        id: int = None,
        creation_timestamp: str = None,
        name: str = None,
        description: str = None,
        zone: str = None,
        region: str = None,
        distribution_policy: dict = None,
        instance_template: str = None,
        versions: list = None,
        instance_group: str = None,
        target_pools: list = None,
        base_instance_name: str = None,
        fingerprint: str = None,
        current_actions: dict = None,
        status: dict = None,
        target_size: int = None,
        self_link: str = None,
        auto_healing_policies: list = None,
        update_policy: dict = None,
        named_ports: list = None,
        stateful_policy: dict = None,
        service_account: str = None,
        failover_action: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.distribution_policy = distribution_policy
        self.instance_template = instance_template
        self.versions = versions
        self.target_pools = target_pools
        self.base_instance_name = base_instance_name
        self.target_size = target_size
        self.auto_healing_policies = auto_healing_policies
        self.update_policy = update_policy
        self.named_ports = named_ports
        self.stateful_policy = stateful_policy
        self.service_account = service_account
        self.failover_action = failover_action
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = (
            instance_group_manager_pb2_grpc.ComputeAlphaInstanceGroupManagerServiceStub(
                channel.Channel()
            )
        )
        request = (
            instance_group_manager_pb2.ApplyComputeAlphaInstanceGroupManagerRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if InstanceGroupManagerDistributionPolicy.to_proto(self.distribution_policy):
            request.resource.distribution_policy.CopyFrom(
                InstanceGroupManagerDistributionPolicy.to_proto(
                    self.distribution_policy
                )
            )
        else:
            request.resource.ClearField("distribution_policy")
        if Primitive.to_proto(self.instance_template):
            request.resource.instance_template = Primitive.to_proto(
                self.instance_template
            )

        if InstanceGroupManagerVersionsArray.to_proto(self.versions):
            request.resource.versions.extend(
                InstanceGroupManagerVersionsArray.to_proto(self.versions)
            )
        if Primitive.to_proto(self.target_pools):
            request.resource.target_pools.extend(Primitive.to_proto(self.target_pools))
        if Primitive.to_proto(self.base_instance_name):
            request.resource.base_instance_name = Primitive.to_proto(
                self.base_instance_name
            )

        if Primitive.to_proto(self.target_size):
            request.resource.target_size = Primitive.to_proto(self.target_size)

        if InstanceGroupManagerAutoHealingPoliciesArray.to_proto(
            self.auto_healing_policies
        ):
            request.resource.auto_healing_policies.extend(
                InstanceGroupManagerAutoHealingPoliciesArray.to_proto(
                    self.auto_healing_policies
                )
            )
        if InstanceGroupManagerUpdatePolicy.to_proto(self.update_policy):
            request.resource.update_policy.CopyFrom(
                InstanceGroupManagerUpdatePolicy.to_proto(self.update_policy)
            )
        else:
            request.resource.ClearField("update_policy")
        if InstanceGroupManagerNamedPortsArray.to_proto(self.named_ports):
            request.resource.named_ports.extend(
                InstanceGroupManagerNamedPortsArray.to_proto(self.named_ports)
            )
        if InstanceGroupManagerStatefulPolicy.to_proto(self.stateful_policy):
            request.resource.stateful_policy.CopyFrom(
                InstanceGroupManagerStatefulPolicy.to_proto(self.stateful_policy)
            )
        else:
            request.resource.ClearField("stateful_policy")
        if Primitive.to_proto(self.service_account):
            request.resource.service_account = Primitive.to_proto(self.service_account)

        if InstanceGroupManagerFailoverActionEnum.to_proto(self.failover_action):
            request.resource.failover_action = (
                InstanceGroupManagerFailoverActionEnum.to_proto(self.failover_action)
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeAlphaInstanceGroupManager(request)
        self.id = Primitive.from_proto(response.id)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.zone = Primitive.from_proto(response.zone)
        self.region = Primitive.from_proto(response.region)
        self.distribution_policy = InstanceGroupManagerDistributionPolicy.from_proto(
            response.distribution_policy
        )
        self.instance_template = Primitive.from_proto(response.instance_template)
        self.versions = InstanceGroupManagerVersionsArray.from_proto(response.versions)
        self.instance_group = Primitive.from_proto(response.instance_group)
        self.target_pools = Primitive.from_proto(response.target_pools)
        self.base_instance_name = Primitive.from_proto(response.base_instance_name)
        self.fingerprint = Primitive.from_proto(response.fingerprint)
        self.current_actions = InstanceGroupManagerCurrentActions.from_proto(
            response.current_actions
        )
        self.status = InstanceGroupManagerStatus.from_proto(response.status)
        self.target_size = Primitive.from_proto(response.target_size)
        self.self_link = Primitive.from_proto(response.self_link)
        self.auto_healing_policies = (
            InstanceGroupManagerAutoHealingPoliciesArray.from_proto(
                response.auto_healing_policies
            )
        )
        self.update_policy = InstanceGroupManagerUpdatePolicy.from_proto(
            response.update_policy
        )
        self.named_ports = InstanceGroupManagerNamedPortsArray.from_proto(
            response.named_ports
        )
        self.stateful_policy = InstanceGroupManagerStatefulPolicy.from_proto(
            response.stateful_policy
        )
        self.service_account = Primitive.from_proto(response.service_account)
        self.failover_action = InstanceGroupManagerFailoverActionEnum.from_proto(
            response.failover_action
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = (
            instance_group_manager_pb2_grpc.ComputeAlphaInstanceGroupManagerServiceStub(
                channel.Channel()
            )
        )
        request = (
            instance_group_manager_pb2.DeleteComputeAlphaInstanceGroupManagerRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if InstanceGroupManagerDistributionPolicy.to_proto(self.distribution_policy):
            request.resource.distribution_policy.CopyFrom(
                InstanceGroupManagerDistributionPolicy.to_proto(
                    self.distribution_policy
                )
            )
        else:
            request.resource.ClearField("distribution_policy")
        if Primitive.to_proto(self.instance_template):
            request.resource.instance_template = Primitive.to_proto(
                self.instance_template
            )

        if InstanceGroupManagerVersionsArray.to_proto(self.versions):
            request.resource.versions.extend(
                InstanceGroupManagerVersionsArray.to_proto(self.versions)
            )
        if Primitive.to_proto(self.target_pools):
            request.resource.target_pools.extend(Primitive.to_proto(self.target_pools))
        if Primitive.to_proto(self.base_instance_name):
            request.resource.base_instance_name = Primitive.to_proto(
                self.base_instance_name
            )

        if Primitive.to_proto(self.target_size):
            request.resource.target_size = Primitive.to_proto(self.target_size)

        if InstanceGroupManagerAutoHealingPoliciesArray.to_proto(
            self.auto_healing_policies
        ):
            request.resource.auto_healing_policies.extend(
                InstanceGroupManagerAutoHealingPoliciesArray.to_proto(
                    self.auto_healing_policies
                )
            )
        if InstanceGroupManagerUpdatePolicy.to_proto(self.update_policy):
            request.resource.update_policy.CopyFrom(
                InstanceGroupManagerUpdatePolicy.to_proto(self.update_policy)
            )
        else:
            request.resource.ClearField("update_policy")
        if InstanceGroupManagerNamedPortsArray.to_proto(self.named_ports):
            request.resource.named_ports.extend(
                InstanceGroupManagerNamedPortsArray.to_proto(self.named_ports)
            )
        if InstanceGroupManagerStatefulPolicy.to_proto(self.stateful_policy):
            request.resource.stateful_policy.CopyFrom(
                InstanceGroupManagerStatefulPolicy.to_proto(self.stateful_policy)
            )
        else:
            request.resource.ClearField("stateful_policy")
        if Primitive.to_proto(self.service_account):
            request.resource.service_account = Primitive.to_proto(self.service_account)

        if InstanceGroupManagerFailoverActionEnum.to_proto(self.failover_action):
            request.resource.failover_action = (
                InstanceGroupManagerFailoverActionEnum.to_proto(self.failover_action)
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteComputeAlphaInstanceGroupManager(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = (
            instance_group_manager_pb2_grpc.ComputeAlphaInstanceGroupManagerServiceStub(
                channel.Channel()
            )
        )
        request = (
            instance_group_manager_pb2.ListComputeAlphaInstanceGroupManagerRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeAlphaInstanceGroupManager(request).items

    def to_proto(self):
        resource = instance_group_manager_pb2.ComputeAlphaInstanceGroupManager()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if InstanceGroupManagerDistributionPolicy.to_proto(self.distribution_policy):
            resource.distribution_policy.CopyFrom(
                InstanceGroupManagerDistributionPolicy.to_proto(
                    self.distribution_policy
                )
            )
        else:
            resource.ClearField("distribution_policy")
        if Primitive.to_proto(self.instance_template):
            resource.instance_template = Primitive.to_proto(self.instance_template)
        if InstanceGroupManagerVersionsArray.to_proto(self.versions):
            resource.versions.extend(
                InstanceGroupManagerVersionsArray.to_proto(self.versions)
            )
        if Primitive.to_proto(self.target_pools):
            resource.target_pools.extend(Primitive.to_proto(self.target_pools))
        if Primitive.to_proto(self.base_instance_name):
            resource.base_instance_name = Primitive.to_proto(self.base_instance_name)
        if Primitive.to_proto(self.target_size):
            resource.target_size = Primitive.to_proto(self.target_size)
        if InstanceGroupManagerAutoHealingPoliciesArray.to_proto(
            self.auto_healing_policies
        ):
            resource.auto_healing_policies.extend(
                InstanceGroupManagerAutoHealingPoliciesArray.to_proto(
                    self.auto_healing_policies
                )
            )
        if InstanceGroupManagerUpdatePolicy.to_proto(self.update_policy):
            resource.update_policy.CopyFrom(
                InstanceGroupManagerUpdatePolicy.to_proto(self.update_policy)
            )
        else:
            resource.ClearField("update_policy")
        if InstanceGroupManagerNamedPortsArray.to_proto(self.named_ports):
            resource.named_ports.extend(
                InstanceGroupManagerNamedPortsArray.to_proto(self.named_ports)
            )
        if InstanceGroupManagerStatefulPolicy.to_proto(self.stateful_policy):
            resource.stateful_policy.CopyFrom(
                InstanceGroupManagerStatefulPolicy.to_proto(self.stateful_policy)
            )
        else:
            resource.ClearField("stateful_policy")
        if Primitive.to_proto(self.service_account):
            resource.service_account = Primitive.to_proto(self.service_account)
        if InstanceGroupManagerFailoverActionEnum.to_proto(self.failover_action):
            resource.failover_action = InstanceGroupManagerFailoverActionEnum.to_proto(
                self.failover_action
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class InstanceGroupManagerDistributionPolicy(object):
    def __init__(self, zones: list = None, target_shape: str = None):
        self.zones = zones
        self.target_shape = target_shape

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerDistributionPolicy()
        )
        if InstanceGroupManagerDistributionPolicyZonesArray.to_proto(resource.zones):
            res.zones.extend(
                InstanceGroupManagerDistributionPolicyZonesArray.to_proto(
                    resource.zones
                )
            )
        if InstanceGroupManagerDistributionPolicyTargetShapeEnum.to_proto(
            resource.target_shape
        ):
            res.target_shape = (
                InstanceGroupManagerDistributionPolicyTargetShapeEnum.to_proto(
                    resource.target_shape
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerDistributionPolicy(
            zones=InstanceGroupManagerDistributionPolicyZonesArray.from_proto(
                resource.zones
            ),
            target_shape=InstanceGroupManagerDistributionPolicyTargetShapeEnum.from_proto(
                resource.target_shape
            ),
        )


class InstanceGroupManagerDistributionPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerDistributionPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerDistributionPolicy.from_proto(i) for i in resources]


class InstanceGroupManagerDistributionPolicyZones(object):
    def __init__(self, zone: str = None):
        self.zone = zone

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerDistributionPolicyZones()
        )
        if Primitive.to_proto(resource.zone):
            res.zone = Primitive.to_proto(resource.zone)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerDistributionPolicyZones(
            zone=Primitive.from_proto(resource.zone),
        )


class InstanceGroupManagerDistributionPolicyZonesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceGroupManagerDistributionPolicyZones.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerDistributionPolicyZones.from_proto(i) for i in resources
        ]


class InstanceGroupManagerVersions(object):
    def __init__(
        self, name: str = None, instance_template: str = None, target_size: dict = None
    ):
        self.name = name
        self.instance_template = instance_template
        self.target_size = target_size

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerVersions()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.instance_template):
            res.instance_template = Primitive.to_proto(resource.instance_template)
        if InstanceGroupManagerVersionsTargetSize.to_proto(resource.target_size):
            res.target_size.CopyFrom(
                InstanceGroupManagerVersionsTargetSize.to_proto(resource.target_size)
            )
        else:
            res.ClearField("target_size")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerVersions(
            name=Primitive.from_proto(resource.name),
            instance_template=Primitive.from_proto(resource.instance_template),
            target_size=InstanceGroupManagerVersionsTargetSize.from_proto(
                resource.target_size
            ),
        )


class InstanceGroupManagerVersionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerVersions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerVersions.from_proto(i) for i in resources]


class InstanceGroupManagerVersionsTargetSize(object):
    def __init__(self, fixed: int = None, percent: int = None, calculated: int = None):
        self.fixed = fixed
        self.percent = percent
        self.calculated = calculated

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerVersionsTargetSize()
        )
        if Primitive.to_proto(resource.fixed):
            res.fixed = Primitive.to_proto(resource.fixed)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        if Primitive.to_proto(resource.calculated):
            res.calculated = Primitive.to_proto(resource.calculated)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerVersionsTargetSize(
            fixed=Primitive.from_proto(resource.fixed),
            percent=Primitive.from_proto(resource.percent),
            calculated=Primitive.from_proto(resource.calculated),
        )


class InstanceGroupManagerVersionsTargetSizeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerVersionsTargetSize.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerVersionsTargetSize.from_proto(i) for i in resources]


class InstanceGroupManagerCurrentActions(object):
    def __init__(
        self,
        none: int = None,
        creating: int = None,
        creating_without_retries: int = None,
        verifying: int = None,
        recreating: int = None,
        deleting: int = None,
        abandoning: int = None,
        restarting: int = None,
        refreshing: int = None,
    ):
        self.none = none
        self.creating = creating
        self.creating_without_retries = creating_without_retries
        self.verifying = verifying
        self.recreating = recreating
        self.deleting = deleting
        self.abandoning = abandoning
        self.restarting = restarting
        self.refreshing = refreshing

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerCurrentActions()
        )
        if Primitive.to_proto(resource.none):
            res.none = Primitive.to_proto(resource.none)
        if Primitive.to_proto(resource.creating):
            res.creating = Primitive.to_proto(resource.creating)
        if Primitive.to_proto(resource.creating_without_retries):
            res.creating_without_retries = Primitive.to_proto(
                resource.creating_without_retries
            )
        if Primitive.to_proto(resource.verifying):
            res.verifying = Primitive.to_proto(resource.verifying)
        if Primitive.to_proto(resource.recreating):
            res.recreating = Primitive.to_proto(resource.recreating)
        if Primitive.to_proto(resource.deleting):
            res.deleting = Primitive.to_proto(resource.deleting)
        if Primitive.to_proto(resource.abandoning):
            res.abandoning = Primitive.to_proto(resource.abandoning)
        if Primitive.to_proto(resource.restarting):
            res.restarting = Primitive.to_proto(resource.restarting)
        if Primitive.to_proto(resource.refreshing):
            res.refreshing = Primitive.to_proto(resource.refreshing)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerCurrentActions(
            none=Primitive.from_proto(resource.none),
            creating=Primitive.from_proto(resource.creating),
            creating_without_retries=Primitive.from_proto(
                resource.creating_without_retries
            ),
            verifying=Primitive.from_proto(resource.verifying),
            recreating=Primitive.from_proto(resource.recreating),
            deleting=Primitive.from_proto(resource.deleting),
            abandoning=Primitive.from_proto(resource.abandoning),
            restarting=Primitive.from_proto(resource.restarting),
            refreshing=Primitive.from_proto(resource.refreshing),
        )


class InstanceGroupManagerCurrentActionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerCurrentActions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerCurrentActions.from_proto(i) for i in resources]


class InstanceGroupManagerStatus(object):
    def __init__(
        self,
        is_stable: bool = None,
        version_target: dict = None,
        stateful: dict = None,
        autoscaler: str = None,
    ):
        self.is_stable = is_stable
        self.version_target = version_target
        self.stateful = stateful
        self.autoscaler = autoscaler

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatus()
        if Primitive.to_proto(resource.is_stable):
            res.is_stable = Primitive.to_proto(resource.is_stable)
        if InstanceGroupManagerStatusVersionTarget.to_proto(resource.version_target):
            res.version_target.CopyFrom(
                InstanceGroupManagerStatusVersionTarget.to_proto(
                    resource.version_target
                )
            )
        else:
            res.ClearField("version_target")
        if InstanceGroupManagerStatusStateful.to_proto(resource.stateful):
            res.stateful.CopyFrom(
                InstanceGroupManagerStatusStateful.to_proto(resource.stateful)
            )
        else:
            res.ClearField("stateful")
        if Primitive.to_proto(resource.autoscaler):
            res.autoscaler = Primitive.to_proto(resource.autoscaler)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerStatus(
            is_stable=Primitive.from_proto(resource.is_stable),
            version_target=InstanceGroupManagerStatusVersionTarget.from_proto(
                resource.version_target
            ),
            stateful=InstanceGroupManagerStatusStateful.from_proto(resource.stateful),
            autoscaler=Primitive.from_proto(resource.autoscaler),
        )


class InstanceGroupManagerStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerStatus.from_proto(i) for i in resources]


class InstanceGroupManagerStatusVersionTarget(object):
    def __init__(self, is_reached: bool = None):
        self.is_reached = is_reached

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatusVersionTarget()
        )
        if Primitive.to_proto(resource.is_reached):
            res.is_reached = Primitive.to_proto(resource.is_reached)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerStatusVersionTarget(
            is_reached=Primitive.from_proto(resource.is_reached),
        )


class InstanceGroupManagerStatusVersionTargetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerStatusVersionTarget.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerStatusVersionTarget.from_proto(i) for i in resources
        ]


class InstanceGroupManagerStatusStateful(object):
    def __init__(
        self,
        has_stateful_config: bool = None,
        per_instance_configs: dict = None,
        is_stateful: bool = None,
    ):
        self.has_stateful_config = has_stateful_config
        self.per_instance_configs = per_instance_configs
        self.is_stateful = is_stateful

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatusStateful()
        )
        if Primitive.to_proto(resource.has_stateful_config):
            res.has_stateful_config = Primitive.to_proto(resource.has_stateful_config)
        if InstanceGroupManagerStatusStatefulPerInstanceConfigs.to_proto(
            resource.per_instance_configs
        ):
            res.per_instance_configs.CopyFrom(
                InstanceGroupManagerStatusStatefulPerInstanceConfigs.to_proto(
                    resource.per_instance_configs
                )
            )
        else:
            res.ClearField("per_instance_configs")
        if Primitive.to_proto(resource.is_stateful):
            res.is_stateful = Primitive.to_proto(resource.is_stateful)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerStatusStateful(
            has_stateful_config=Primitive.from_proto(resource.has_stateful_config),
            per_instance_configs=InstanceGroupManagerStatusStatefulPerInstanceConfigs.from_proto(
                resource.per_instance_configs
            ),
            is_stateful=Primitive.from_proto(resource.is_stateful),
        )


class InstanceGroupManagerStatusStatefulArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerStatusStateful.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerStatusStateful.from_proto(i) for i in resources]


class InstanceGroupManagerStatusStatefulPerInstanceConfigs(object):
    def __init__(self, all_effective: bool = None):
        self.all_effective = all_effective

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigs()
        )
        if Primitive.to_proto(resource.all_effective):
            res.all_effective = Primitive.to_proto(resource.all_effective)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerStatusStatefulPerInstanceConfigs(
            all_effective=Primitive.from_proto(resource.all_effective),
        )


class InstanceGroupManagerStatusStatefulPerInstanceConfigsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceGroupManagerStatusStatefulPerInstanceConfigs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerStatusStatefulPerInstanceConfigs.from_proto(i)
            for i in resources
        ]


class InstanceGroupManagerAutoHealingPolicies(object):
    def __init__(self, health_check: str = None, initial_delay_sec: int = None):
        self.health_check = health_check
        self.initial_delay_sec = initial_delay_sec

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerAutoHealingPolicies()
        )
        if Primitive.to_proto(resource.health_check):
            res.health_check = Primitive.to_proto(resource.health_check)
        if Primitive.to_proto(resource.initial_delay_sec):
            res.initial_delay_sec = Primitive.to_proto(resource.initial_delay_sec)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerAutoHealingPolicies(
            health_check=Primitive.from_proto(resource.health_check),
            initial_delay_sec=Primitive.from_proto(resource.initial_delay_sec),
        )


class InstanceGroupManagerAutoHealingPoliciesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerAutoHealingPolicies.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerAutoHealingPolicies.from_proto(i) for i in resources
        ]


class InstanceGroupManagerUpdatePolicy(object):
    def __init__(
        self,
        type: str = None,
        instance_redistribution_type: str = None,
        minimal_action: str = None,
        max_surge: dict = None,
        max_unavailable: dict = None,
        replacement_method: str = None,
        most_disruptive_allowed_action: str = None,
        min_ready_sec: int = None,
    ):
        self.type = type
        self.instance_redistribution_type = instance_redistribution_type
        self.minimal_action = minimal_action
        self.max_surge = max_surge
        self.max_unavailable = max_unavailable
        self.replacement_method = replacement_method
        self.most_disruptive_allowed_action = most_disruptive_allowed_action
        self.min_ready_sec = min_ready_sec

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicy()
        if InstanceGroupManagerUpdatePolicyTypeEnum.to_proto(resource.type):
            res.type = InstanceGroupManagerUpdatePolicyTypeEnum.to_proto(resource.type)
        if InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum.to_proto(
            resource.instance_redistribution_type
        ):
            res.instance_redistribution_type = (
                InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum.to_proto(
                    resource.instance_redistribution_type
                )
            )
        if InstanceGroupManagerUpdatePolicyMinimalActionEnum.to_proto(
            resource.minimal_action
        ):
            res.minimal_action = (
                InstanceGroupManagerUpdatePolicyMinimalActionEnum.to_proto(
                    resource.minimal_action
                )
            )
        if InstanceGroupManagerUpdatePolicyMaxSurge.to_proto(resource.max_surge):
            res.max_surge.CopyFrom(
                InstanceGroupManagerUpdatePolicyMaxSurge.to_proto(resource.max_surge)
            )
        else:
            res.ClearField("max_surge")
        if InstanceGroupManagerUpdatePolicyMaxUnavailable.to_proto(
            resource.max_unavailable
        ):
            res.max_unavailable.CopyFrom(
                InstanceGroupManagerUpdatePolicyMaxUnavailable.to_proto(
                    resource.max_unavailable
                )
            )
        else:
            res.ClearField("max_unavailable")
        if InstanceGroupManagerUpdatePolicyReplacementMethodEnum.to_proto(
            resource.replacement_method
        ):
            res.replacement_method = (
                InstanceGroupManagerUpdatePolicyReplacementMethodEnum.to_proto(
                    resource.replacement_method
                )
            )
        if InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum.to_proto(
            resource.most_disruptive_allowed_action
        ):
            res.most_disruptive_allowed_action = InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum.to_proto(
                resource.most_disruptive_allowed_action
            )
        if Primitive.to_proto(resource.min_ready_sec):
            res.min_ready_sec = Primitive.to_proto(resource.min_ready_sec)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerUpdatePolicy(
            type=InstanceGroupManagerUpdatePolicyTypeEnum.from_proto(resource.type),
            instance_redistribution_type=InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum.from_proto(
                resource.instance_redistribution_type
            ),
            minimal_action=InstanceGroupManagerUpdatePolicyMinimalActionEnum.from_proto(
                resource.minimal_action
            ),
            max_surge=InstanceGroupManagerUpdatePolicyMaxSurge.from_proto(
                resource.max_surge
            ),
            max_unavailable=InstanceGroupManagerUpdatePolicyMaxUnavailable.from_proto(
                resource.max_unavailable
            ),
            replacement_method=InstanceGroupManagerUpdatePolicyReplacementMethodEnum.from_proto(
                resource.replacement_method
            ),
            most_disruptive_allowed_action=InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum.from_proto(
                resource.most_disruptive_allowed_action
            ),
            min_ready_sec=Primitive.from_proto(resource.min_ready_sec),
        )


class InstanceGroupManagerUpdatePolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerUpdatePolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerUpdatePolicy.from_proto(i) for i in resources]


class InstanceGroupManagerUpdatePolicyMaxSurge(object):
    def __init__(self, fixed: int = None, percent: int = None, calculated: int = None):
        self.fixed = fixed
        self.percent = percent
        self.calculated = calculated

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyMaxSurge()
        )
        if Primitive.to_proto(resource.fixed):
            res.fixed = Primitive.to_proto(resource.fixed)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        if Primitive.to_proto(resource.calculated):
            res.calculated = Primitive.to_proto(resource.calculated)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerUpdatePolicyMaxSurge(
            fixed=Primitive.from_proto(resource.fixed),
            percent=Primitive.from_proto(resource.percent),
            calculated=Primitive.from_proto(resource.calculated),
        )


class InstanceGroupManagerUpdatePolicyMaxSurgeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerUpdatePolicyMaxSurge.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerUpdatePolicyMaxSurge.from_proto(i) for i in resources
        ]


class InstanceGroupManagerUpdatePolicyMaxUnavailable(object):
    def __init__(self, fixed: int = None, percent: int = None, calculated: int = None):
        self.fixed = fixed
        self.percent = percent
        self.calculated = calculated

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyMaxUnavailable()
        )
        if Primitive.to_proto(resource.fixed):
            res.fixed = Primitive.to_proto(resource.fixed)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        if Primitive.to_proto(resource.calculated):
            res.calculated = Primitive.to_proto(resource.calculated)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerUpdatePolicyMaxUnavailable(
            fixed=Primitive.from_proto(resource.fixed),
            percent=Primitive.from_proto(resource.percent),
            calculated=Primitive.from_proto(resource.calculated),
        )


class InstanceGroupManagerUpdatePolicyMaxUnavailableArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceGroupManagerUpdatePolicyMaxUnavailable.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerUpdatePolicyMaxUnavailable.from_proto(i)
            for i in resources
        ]


class InstanceGroupManagerNamedPorts(object):
    def __init__(self, name: str = None, port: int = None):
        self.name = name
        self.port = port

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerNamedPorts()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerNamedPorts(
            name=Primitive.from_proto(resource.name),
            port=Primitive.from_proto(resource.port),
        )


class InstanceGroupManagerNamedPortsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerNamedPorts.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerNamedPorts.from_proto(i) for i in resources]


class InstanceGroupManagerStatefulPolicy(object):
    def __init__(self, preserved_state: dict = None):
        self.preserved_state = preserved_state

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatefulPolicy()
        )
        if InstanceGroupManagerStatefulPolicyPreservedState.to_proto(
            resource.preserved_state
        ):
            res.preserved_state.CopyFrom(
                InstanceGroupManagerStatefulPolicyPreservedState.to_proto(
                    resource.preserved_state
                )
            )
        else:
            res.ClearField("preserved_state")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerStatefulPolicy(
            preserved_state=InstanceGroupManagerStatefulPolicyPreservedState.from_proto(
                resource.preserved_state
            ),
        )


class InstanceGroupManagerStatefulPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerStatefulPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerStatefulPolicy.from_proto(i) for i in resources]


class InstanceGroupManagerStatefulPolicyPreservedState(object):
    def __init__(
        self, disks: dict = None, internal_ips: dict = None, external_ips: dict = None
    ):
        self.disks = disks
        self.internal_ips = internal_ips
        self.external_ips = external_ips

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedState()
        )
        if Primitive.to_proto(resource.disks):
            res.disks = Primitive.to_proto(resource.disks)
        if Primitive.to_proto(resource.internal_ips):
            res.internal_ips = Primitive.to_proto(resource.internal_ips)
        if Primitive.to_proto(resource.external_ips):
            res.external_ips = Primitive.to_proto(resource.external_ips)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerStatefulPolicyPreservedState(
            disks=Primitive.from_proto(resource.disks),
            internal_ips=Primitive.from_proto(resource.internal_ips),
            external_ips=Primitive.from_proto(resource.external_ips),
        )


class InstanceGroupManagerStatefulPolicyPreservedStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceGroupManagerStatefulPolicyPreservedState.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerStatefulPolicyPreservedState.from_proto(i)
            for i in resources
        ]


class InstanceGroupManagerStatefulPolicyPreservedStateDisks(object):
    def __init__(self, auto_delete: str = None):
        self.auto_delete = auto_delete

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisks()
        )
        if InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum.to_proto(
            resource.auto_delete
        ):
            res.auto_delete = InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum.to_proto(
                resource.auto_delete
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerStatefulPolicyPreservedStateDisks(
            auto_delete=InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum.from_proto(
                resource.auto_delete
            ),
        )


class InstanceGroupManagerStatefulPolicyPreservedStateDisksArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceGroupManagerStatefulPolicyPreservedStateDisks.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerStatefulPolicyPreservedStateDisks.from_proto(i)
            for i in resources
        ]


class InstanceGroupManagerStatefulPolicyPreservedStateInternalIps(object):
    def __init__(self, auto_delete: str = None):
        self.auto_delete = auto_delete

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIps()
        )
        if InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum.to_proto(
            resource.auto_delete
        ):
            res.auto_delete = InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum.to_proto(
                resource.auto_delete
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerStatefulPolicyPreservedStateInternalIps(
            auto_delete=InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum.from_proto(
                resource.auto_delete
            ),
        )


class InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceGroupManagerStatefulPolicyPreservedStateInternalIps.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerStatefulPolicyPreservedStateInternalIps.from_proto(i)
            for i in resources
        ]


class InstanceGroupManagerStatefulPolicyPreservedStateExternalIps(object):
    def __init__(self, auto_delete: str = None):
        self.auto_delete = auto_delete

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIps()
        )
        if InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum.to_proto(
            resource.auto_delete
        ):
            res.auto_delete = InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum.to_proto(
                resource.auto_delete
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerStatefulPolicyPreservedStateExternalIps(
            auto_delete=InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum.from_proto(
                resource.auto_delete
            ),
        )


class InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceGroupManagerStatefulPolicyPreservedStateExternalIps.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerStatefulPolicyPreservedStateExternalIps.from_proto(i)
            for i in resources
        ]


class InstanceGroupManagerDistributionPolicyTargetShapeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum.Value(
            "ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum.Name(
            resource
        )[
            len("ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum") :
        ]


class InstanceGroupManagerUpdatePolicyTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum.Value(
            "ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum.Name(
            resource
        )[
            len("ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum") :
        ]


class InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum.Value(
            "ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum.Name(
            resource
        )[
            len(
                "ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"
            ) :
        ]


class InstanceGroupManagerUpdatePolicyMinimalActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum.Value(
            "ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum.Name(
            resource
        )[
            len("ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum") :
        ]


class InstanceGroupManagerUpdatePolicyReplacementMethodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum.Value(
            "ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum.Name(
            resource
        )[
            len("ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum") :
        ]


class InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum.Value(
            "ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum.Name(
            resource
        )[
            len(
                "ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum"
            ) :
        ]


class InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum.Value(
            "ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum.Name(
            resource
        )[
            len(
                "ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"
            ) :
        ]


class InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum.Value(
            "ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum.Name(
            resource
        )[
            len(
                "ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum"
            ) :
        ]


class InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum.Value(
            "ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum.Name(
            resource
        )[
            len(
                "ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum"
            ) :
        ]


class InstanceGroupManagerFailoverActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerFailoverActionEnum.Value(
            "ComputeAlphaInstanceGroupManagerFailoverActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeAlphaInstanceGroupManagerFailoverActionEnum.Name(
            resource
        )[
            len("ComputeAlphaInstanceGroupManagerFailoverActionEnum") :
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
