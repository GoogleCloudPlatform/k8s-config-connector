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
from google3.cloud.graphite.mmv2.services.google.container_azure import node_pool_pb2
from google3.cloud.graphite.mmv2.services.google.container_azure import (
    node_pool_pb2_grpc,
)

from typing import List


class NodePool(object):
    def __init__(
        self,
        name: str = None,
        version: str = None,
        config: dict = None,
        subnet_id: str = None,
        autoscaling: dict = None,
        state: str = None,
        uid: str = None,
        reconciling: bool = None,
        create_time: str = None,
        update_time: str = None,
        etag: str = None,
        annotations: dict = None,
        max_pods_constraint: dict = None,
        management: dict = None,
        azure_availability_zone: str = None,
        project: str = None,
        location: str = None,
        cluster: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.version = version
        self.config = config
        self.subnet_id = subnet_id
        self.autoscaling = autoscaling
        self.annotations = annotations
        self.max_pods_constraint = max_pods_constraint
        self.management = management
        self.azure_availability_zone = azure_availability_zone
        self.project = project
        self.location = location
        self.cluster = cluster
        self.service_account_file = service_account_file

    def apply(self):
        stub = node_pool_pb2_grpc.ContainerazureNodePoolServiceStub(channel.Channel())
        request = node_pool_pb2.ApplyContainerazureNodePoolRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if NodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(NodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.subnet_id):
            request.resource.subnet_id = Primitive.to_proto(self.subnet_id)

        if NodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                NodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if NodePoolManagement.to_proto(self.management):
            request.resource.management.CopyFrom(
                NodePoolManagement.to_proto(self.management)
            )
        else:
            request.resource.ClearField("management")
        if Primitive.to_proto(self.azure_availability_zone):
            request.resource.azure_availability_zone = Primitive.to_proto(
                self.azure_availability_zone
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.cluster):
            request.resource.cluster = Primitive.to_proto(self.cluster)

        request.service_account_file = self.service_account_file

        response = stub.ApplyContainerazureNodePool(request)
        self.name = Primitive.from_proto(response.name)
        self.version = Primitive.from_proto(response.version)
        self.config = NodePoolConfig.from_proto(response.config)
        self.subnet_id = Primitive.from_proto(response.subnet_id)
        self.autoscaling = NodePoolAutoscaling.from_proto(response.autoscaling)
        self.state = NodePoolStateEnum.from_proto(response.state)
        self.uid = Primitive.from_proto(response.uid)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.etag = Primitive.from_proto(response.etag)
        self.annotations = Primitive.from_proto(response.annotations)
        self.max_pods_constraint = NodePoolMaxPodsConstraint.from_proto(
            response.max_pods_constraint
        )
        self.management = NodePoolManagement.from_proto(response.management)
        self.azure_availability_zone = Primitive.from_proto(
            response.azure_availability_zone
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.cluster = Primitive.from_proto(response.cluster)

    def delete(self):
        stub = node_pool_pb2_grpc.ContainerazureNodePoolServiceStub(channel.Channel())
        request = node_pool_pb2.DeleteContainerazureNodePoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if NodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(NodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.subnet_id):
            request.resource.subnet_id = Primitive.to_proto(self.subnet_id)

        if NodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                NodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if NodePoolManagement.to_proto(self.management):
            request.resource.management.CopyFrom(
                NodePoolManagement.to_proto(self.management)
            )
        else:
            request.resource.ClearField("management")
        if Primitive.to_proto(self.azure_availability_zone):
            request.resource.azure_availability_zone = Primitive.to_proto(
                self.azure_availability_zone
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.cluster):
            request.resource.cluster = Primitive.to_proto(self.cluster)

        response = stub.DeleteContainerazureNodePool(request)

    @classmethod
    def list(self, project, location, cluster, service_account_file=""):
        stub = node_pool_pb2_grpc.ContainerazureNodePoolServiceStub(channel.Channel())
        request = node_pool_pb2.ListContainerazureNodePoolRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Cluster = cluster

        return stub.ListContainerazureNodePool(request).items

    def to_proto(self):
        resource = node_pool_pb2.ContainerazureNodePool()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.version):
            resource.version = Primitive.to_proto(self.version)
        if NodePoolConfig.to_proto(self.config):
            resource.config.CopyFrom(NodePoolConfig.to_proto(self.config))
        else:
            resource.ClearField("config")
        if Primitive.to_proto(self.subnet_id):
            resource.subnet_id = Primitive.to_proto(self.subnet_id)
        if NodePoolAutoscaling.to_proto(self.autoscaling):
            resource.autoscaling.CopyFrom(
                NodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            resource.ClearField("autoscaling")
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            resource.max_pods_constraint.CopyFrom(
                NodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            resource.ClearField("max_pods_constraint")
        if NodePoolManagement.to_proto(self.management):
            resource.management.CopyFrom(NodePoolManagement.to_proto(self.management))
        else:
            resource.ClearField("management")
        if Primitive.to_proto(self.azure_availability_zone):
            resource.azure_availability_zone = Primitive.to_proto(
                self.azure_availability_zone
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.cluster):
            resource.cluster = Primitive.to_proto(self.cluster)
        return resource


class NodePoolConfig(object):
    def __init__(
        self,
        vm_size: str = None,
        root_volume: dict = None,
        tags: dict = None,
        labels: dict = None,
        ssh_config: dict = None,
        proxy_config: dict = None,
    ):
        self.vm_size = vm_size
        self.root_volume = root_volume
        self.tags = tags
        self.labels = labels
        self.ssh_config = ssh_config
        self.proxy_config = proxy_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerazureNodePoolConfig()
        if Primitive.to_proto(resource.vm_size):
            res.vm_size = Primitive.to_proto(resource.vm_size)
        if NodePoolConfigRootVolume.to_proto(resource.root_volume):
            res.root_volume.CopyFrom(
                NodePoolConfigRootVolume.to_proto(resource.root_volume)
            )
        else:
            res.ClearField("root_volume")
        if Primitive.to_proto(resource.tags):
            res.tags = Primitive.to_proto(resource.tags)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if NodePoolConfigSshConfig.to_proto(resource.ssh_config):
            res.ssh_config.CopyFrom(
                NodePoolConfigSshConfig.to_proto(resource.ssh_config)
            )
        else:
            res.ClearField("ssh_config")
        if NodePoolConfigProxyConfig.to_proto(resource.proxy_config):
            res.proxy_config.CopyFrom(
                NodePoolConfigProxyConfig.to_proto(resource.proxy_config)
            )
        else:
            res.ClearField("proxy_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfig(
            vm_size=Primitive.from_proto(resource.vm_size),
            root_volume=NodePoolConfigRootVolume.from_proto(resource.root_volume),
            tags=Primitive.from_proto(resource.tags),
            labels=Primitive.from_proto(resource.labels),
            ssh_config=NodePoolConfigSshConfig.from_proto(resource.ssh_config),
            proxy_config=NodePoolConfigProxyConfig.from_proto(resource.proxy_config),
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


class NodePoolConfigRootVolume(object):
    def __init__(self, size_gib: int = None):
        self.size_gib = size_gib

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerazureNodePoolConfigRootVolume()
        if Primitive.to_proto(resource.size_gib):
            res.size_gib = Primitive.to_proto(resource.size_gib)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigRootVolume(
            size_gib=Primitive.from_proto(resource.size_gib),
        )


class NodePoolConfigRootVolumeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigRootVolume.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigRootVolume.from_proto(i) for i in resources]


class NodePoolConfigSshConfig(object):
    def __init__(self, authorized_key: str = None):
        self.authorized_key = authorized_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerazureNodePoolConfigSshConfig()
        if Primitive.to_proto(resource.authorized_key):
            res.authorized_key = Primitive.to_proto(resource.authorized_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigSshConfig(
            authorized_key=Primitive.from_proto(resource.authorized_key),
        )


class NodePoolConfigSshConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigSshConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigSshConfig.from_proto(i) for i in resources]


class NodePoolConfigProxyConfig(object):
    def __init__(self, resource_group_id: str = None, secret_id: str = None):
        self.resource_group_id = resource_group_id
        self.secret_id = secret_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerazureNodePoolConfigProxyConfig()
        if Primitive.to_proto(resource.resource_group_id):
            res.resource_group_id = Primitive.to_proto(resource.resource_group_id)
        if Primitive.to_proto(resource.secret_id):
            res.secret_id = Primitive.to_proto(resource.secret_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolConfigProxyConfig(
            resource_group_id=Primitive.from_proto(resource.resource_group_id),
            secret_id=Primitive.from_proto(resource.secret_id),
        )


class NodePoolConfigProxyConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodePoolConfigProxyConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodePoolConfigProxyConfig.from_proto(i) for i in resources]


class NodePoolAutoscaling(object):
    def __init__(self, min_node_count: int = None, max_node_count: int = None):
        self.min_node_count = min_node_count
        self.max_node_count = max_node_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerazureNodePoolAutoscaling()
        if Primitive.to_proto(resource.min_node_count):
            res.min_node_count = Primitive.to_proto(resource.min_node_count)
        if Primitive.to_proto(resource.max_node_count):
            res.max_node_count = Primitive.to_proto(resource.max_node_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolAutoscaling(
            min_node_count=Primitive.from_proto(resource.min_node_count),
            max_node_count=Primitive.from_proto(resource.max_node_count),
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


class NodePoolMaxPodsConstraint(object):
    def __init__(self, max_pods_per_node: int = None):
        self.max_pods_per_node = max_pods_per_node

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerazureNodePoolMaxPodsConstraint()
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


class NodePoolManagement(object):
    def __init__(self, auto_repair: bool = None):
        self.auto_repair = auto_repair

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pool_pb2.ContainerazureNodePoolManagement()
        if Primitive.to_proto(resource.auto_repair):
            res.auto_repair = Primitive.to_proto(resource.auto_repair)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodePoolManagement(
            auto_repair=Primitive.from_proto(resource.auto_repair),
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


class NodePoolStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerazureNodePoolStateEnum.Value(
            "ContainerazureNodePoolStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return node_pool_pb2.ContainerazureNodePoolStateEnum.Name(resource)[
            len("ContainerazureNodePoolStateEnum") :
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
