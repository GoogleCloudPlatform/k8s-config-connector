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
from google3.cloud.graphite.mmv2.services.google.gkemulticloud import (
    azure_node_pool_pb2,
)
from google3.cloud.graphite.mmv2.services.google.gkemulticloud import (
    azure_node_pool_pb2_grpc,
)

from typing import List


class AzureNodePool(object):
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
        azure_availability_zone: str = None,
        project: str = None,
        location: str = None,
        azure_cluster: str = None,
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
        self.azure_availability_zone = azure_availability_zone
        self.project = project
        self.location = location
        self.azure_cluster = azure_cluster
        self.service_account_file = service_account_file

    def apply(self):
        stub = azure_node_pool_pb2_grpc.GkemulticloudBetaAzureNodePoolServiceStub(
            channel.Channel()
        )
        request = azure_node_pool_pb2.ApplyGkemulticloudBetaAzureNodePoolRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if AzureNodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(AzureNodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.subnet_id):
            request.resource.subnet_id = Primitive.to_proto(self.subnet_id)

        if AzureNodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                AzureNodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if AzureNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                AzureNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if Primitive.to_proto(self.azure_availability_zone):
            request.resource.azure_availability_zone = Primitive.to_proto(
                self.azure_availability_zone
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.azure_cluster):
            request.resource.azure_cluster = Primitive.to_proto(self.azure_cluster)

        request.service_account_file = self.service_account_file

        response = stub.ApplyGkemulticloudBetaAzureNodePool(request)
        self.name = Primitive.from_proto(response.name)
        self.version = Primitive.from_proto(response.version)
        self.config = AzureNodePoolConfig.from_proto(response.config)
        self.subnet_id = Primitive.from_proto(response.subnet_id)
        self.autoscaling = AzureNodePoolAutoscaling.from_proto(response.autoscaling)
        self.state = AzureNodePoolStateEnum.from_proto(response.state)
        self.uid = Primitive.from_proto(response.uid)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.etag = Primitive.from_proto(response.etag)
        self.annotations = Primitive.from_proto(response.annotations)
        self.max_pods_constraint = AzureNodePoolMaxPodsConstraint.from_proto(
            response.max_pods_constraint
        )
        self.azure_availability_zone = Primitive.from_proto(
            response.azure_availability_zone
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.azure_cluster = Primitive.from_proto(response.azure_cluster)

    def delete(self):
        stub = azure_node_pool_pb2_grpc.GkemulticloudBetaAzureNodePoolServiceStub(
            channel.Channel()
        )
        request = azure_node_pool_pb2.DeleteGkemulticloudBetaAzureNodePoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if AzureNodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(AzureNodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.subnet_id):
            request.resource.subnet_id = Primitive.to_proto(self.subnet_id)

        if AzureNodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                AzureNodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if AzureNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                AzureNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if Primitive.to_proto(self.azure_availability_zone):
            request.resource.azure_availability_zone = Primitive.to_proto(
                self.azure_availability_zone
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.azure_cluster):
            request.resource.azure_cluster = Primitive.to_proto(self.azure_cluster)

        response = stub.DeleteGkemulticloudBetaAzureNodePool(request)

    def list(self):
        stub = azure_node_pool_pb2_grpc.GkemulticloudBetaAzureNodePoolServiceStub(
            channel.Channel()
        )
        request = azure_node_pool_pb2.ListGkemulticloudBetaAzureNodePoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if AzureNodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(AzureNodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.subnet_id):
            request.resource.subnet_id = Primitive.to_proto(self.subnet_id)

        if AzureNodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                AzureNodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if AzureNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                AzureNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if Primitive.to_proto(self.azure_availability_zone):
            request.resource.azure_availability_zone = Primitive.to_proto(
                self.azure_availability_zone
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.azure_cluster):
            request.resource.azure_cluster = Primitive.to_proto(self.azure_cluster)

        return stub.ListGkemulticloudBetaAzureNodePool(request).items

    def to_proto(self):
        resource = azure_node_pool_pb2.GkemulticloudBetaAzureNodePool()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.version):
            resource.version = Primitive.to_proto(self.version)
        if AzureNodePoolConfig.to_proto(self.config):
            resource.config.CopyFrom(AzureNodePoolConfig.to_proto(self.config))
        else:
            resource.ClearField("config")
        if Primitive.to_proto(self.subnet_id):
            resource.subnet_id = Primitive.to_proto(self.subnet_id)
        if AzureNodePoolAutoscaling.to_proto(self.autoscaling):
            resource.autoscaling.CopyFrom(
                AzureNodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            resource.ClearField("autoscaling")
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if AzureNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            resource.max_pods_constraint.CopyFrom(
                AzureNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            resource.ClearField("max_pods_constraint")
        if Primitive.to_proto(self.azure_availability_zone):
            resource.azure_availability_zone = Primitive.to_proto(
                self.azure_availability_zone
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.azure_cluster):
            resource.azure_cluster = Primitive.to_proto(self.azure_cluster)
        return resource


class AzureNodePoolConfig(object):
    def __init__(
        self,
        vm_size: str = None,
        root_volume: dict = None,
        tags: dict = None,
        ssh_config: dict = None,
    ):
        self.vm_size = vm_size
        self.root_volume = root_volume
        self.tags = tags
        self.ssh_config = ssh_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = azure_node_pool_pb2.GkemulticloudBetaAzureNodePoolConfig()
        if Primitive.to_proto(resource.vm_size):
            res.vm_size = Primitive.to_proto(resource.vm_size)
        if AzureNodePoolConfigRootVolume.to_proto(resource.root_volume):
            res.root_volume.CopyFrom(
                AzureNodePoolConfigRootVolume.to_proto(resource.root_volume)
            )
        else:
            res.ClearField("root_volume")
        if Primitive.to_proto(resource.tags):
            res.tags = Primitive.to_proto(resource.tags)
        if AzureNodePoolConfigSshConfig.to_proto(resource.ssh_config):
            res.ssh_config.CopyFrom(
                AzureNodePoolConfigSshConfig.to_proto(resource.ssh_config)
            )
        else:
            res.ClearField("ssh_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AzureNodePoolConfig(
            vm_size=Primitive.from_proto(resource.vm_size),
            root_volume=AzureNodePoolConfigRootVolume.from_proto(resource.root_volume),
            tags=Primitive.from_proto(resource.tags),
            ssh_config=AzureNodePoolConfigSshConfig.from_proto(resource.ssh_config),
        )


class AzureNodePoolConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AzureNodePoolConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AzureNodePoolConfig.from_proto(i) for i in resources]


class AzureNodePoolConfigRootVolume(object):
    def __init__(self, size_gib: int = None):
        self.size_gib = size_gib

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = azure_node_pool_pb2.GkemulticloudBetaAzureNodePoolConfigRootVolume()
        if Primitive.to_proto(resource.size_gib):
            res.size_gib = Primitive.to_proto(resource.size_gib)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AzureNodePoolConfigRootVolume(
            size_gib=Primitive.from_proto(resource.size_gib),
        )


class AzureNodePoolConfigRootVolumeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AzureNodePoolConfigRootVolume.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AzureNodePoolConfigRootVolume.from_proto(i) for i in resources]


class AzureNodePoolConfigSshConfig(object):
    def __init__(self, authorized_key: str = None):
        self.authorized_key = authorized_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = azure_node_pool_pb2.GkemulticloudBetaAzureNodePoolConfigSshConfig()
        if Primitive.to_proto(resource.authorized_key):
            res.authorized_key = Primitive.to_proto(resource.authorized_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AzureNodePoolConfigSshConfig(
            authorized_key=Primitive.from_proto(resource.authorized_key),
        )


class AzureNodePoolConfigSshConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AzureNodePoolConfigSshConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AzureNodePoolConfigSshConfig.from_proto(i) for i in resources]


class AzureNodePoolAutoscaling(object):
    def __init__(self, min_node_count: int = None, max_node_count: int = None):
        self.min_node_count = min_node_count
        self.max_node_count = max_node_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = azure_node_pool_pb2.GkemulticloudBetaAzureNodePoolAutoscaling()
        if Primitive.to_proto(resource.min_node_count):
            res.min_node_count = Primitive.to_proto(resource.min_node_count)
        if Primitive.to_proto(resource.max_node_count):
            res.max_node_count = Primitive.to_proto(resource.max_node_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AzureNodePoolAutoscaling(
            min_node_count=Primitive.from_proto(resource.min_node_count),
            max_node_count=Primitive.from_proto(resource.max_node_count),
        )


class AzureNodePoolAutoscalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AzureNodePoolAutoscaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AzureNodePoolAutoscaling.from_proto(i) for i in resources]


class AzureNodePoolMaxPodsConstraint(object):
    def __init__(self, max_pods_per_node: int = None):
        self.max_pods_per_node = max_pods_per_node

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = azure_node_pool_pb2.GkemulticloudBetaAzureNodePoolMaxPodsConstraint()
        if Primitive.to_proto(resource.max_pods_per_node):
            res.max_pods_per_node = Primitive.to_proto(resource.max_pods_per_node)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AzureNodePoolMaxPodsConstraint(
            max_pods_per_node=Primitive.from_proto(resource.max_pods_per_node),
        )


class AzureNodePoolMaxPodsConstraintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AzureNodePoolMaxPodsConstraint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AzureNodePoolMaxPodsConstraint.from_proto(i) for i in resources]


class AzureNodePoolStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return azure_node_pool_pb2.GkemulticloudBetaAzureNodePoolStateEnum.Value(
            "GkemulticloudBetaAzureNodePoolStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return azure_node_pool_pb2.GkemulticloudBetaAzureNodePoolStateEnum.Name(
            resource
        )[len("GkemulticloudBetaAzureNodePoolStateEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
