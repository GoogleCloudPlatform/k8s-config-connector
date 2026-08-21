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
from google3.cloud.graphite.mmv2.services.google.gkemulticloud import aws_node_pool_pb2
from google3.cloud.graphite.mmv2.services.google.gkemulticloud import (
    aws_node_pool_pb2_grpc,
)

from typing import List


class AwsNodePool(object):
    def __init__(
        self,
        name: str = None,
        version: str = None,
        config: dict = None,
        autoscaling: dict = None,
        subnet_id: str = None,
        state: str = None,
        uid: str = None,
        reconciling: bool = None,
        create_time: str = None,
        update_time: str = None,
        etag: str = None,
        annotations: dict = None,
        max_pods_constraint: dict = None,
        project: str = None,
        location: str = None,
        aws_cluster: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.version = version
        self.config = config
        self.autoscaling = autoscaling
        self.subnet_id = subnet_id
        self.annotations = annotations
        self.max_pods_constraint = max_pods_constraint
        self.project = project
        self.location = location
        self.aws_cluster = aws_cluster
        self.service_account_file = service_account_file

    def apply(self):
        stub = aws_node_pool_pb2_grpc.GkemulticloudBetaAwsNodePoolServiceStub(
            channel.Channel()
        )
        request = aws_node_pool_pb2.ApplyGkemulticloudBetaAwsNodePoolRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if AwsNodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(AwsNodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if AwsNodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                AwsNodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if Primitive.to_proto(self.subnet_id):
            request.resource.subnet_id = Primitive.to_proto(self.subnet_id)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if AwsNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                AwsNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.aws_cluster):
            request.resource.aws_cluster = Primitive.to_proto(self.aws_cluster)

        request.service_account_file = self.service_account_file

        response = stub.ApplyGkemulticloudBetaAwsNodePool(request)
        self.name = Primitive.from_proto(response.name)
        self.version = Primitive.from_proto(response.version)
        self.config = AwsNodePoolConfig.from_proto(response.config)
        self.autoscaling = AwsNodePoolAutoscaling.from_proto(response.autoscaling)
        self.subnet_id = Primitive.from_proto(response.subnet_id)
        self.state = AwsNodePoolStateEnum.from_proto(response.state)
        self.uid = Primitive.from_proto(response.uid)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.etag = Primitive.from_proto(response.etag)
        self.annotations = Primitive.from_proto(response.annotations)
        self.max_pods_constraint = AwsNodePoolMaxPodsConstraint.from_proto(
            response.max_pods_constraint
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.aws_cluster = Primitive.from_proto(response.aws_cluster)

    def delete(self):
        stub = aws_node_pool_pb2_grpc.GkemulticloudBetaAwsNodePoolServiceStub(
            channel.Channel()
        )
        request = aws_node_pool_pb2.DeleteGkemulticloudBetaAwsNodePoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if AwsNodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(AwsNodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if AwsNodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                AwsNodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if Primitive.to_proto(self.subnet_id):
            request.resource.subnet_id = Primitive.to_proto(self.subnet_id)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if AwsNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                AwsNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.aws_cluster):
            request.resource.aws_cluster = Primitive.to_proto(self.aws_cluster)

        response = stub.DeleteGkemulticloudBetaAwsNodePool(request)

    def list(self):
        stub = aws_node_pool_pb2_grpc.GkemulticloudBetaAwsNodePoolServiceStub(
            channel.Channel()
        )
        request = aws_node_pool_pb2.ListGkemulticloudBetaAwsNodePoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if AwsNodePoolConfig.to_proto(self.config):
            request.resource.config.CopyFrom(AwsNodePoolConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if AwsNodePoolAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                AwsNodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if Primitive.to_proto(self.subnet_id):
            request.resource.subnet_id = Primitive.to_proto(self.subnet_id)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if AwsNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            request.resource.max_pods_constraint.CopyFrom(
                AwsNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            request.resource.ClearField("max_pods_constraint")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.aws_cluster):
            request.resource.aws_cluster = Primitive.to_proto(self.aws_cluster)

        return stub.ListGkemulticloudBetaAwsNodePool(request).items

    def to_proto(self):
        resource = aws_node_pool_pb2.GkemulticloudBetaAwsNodePool()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.version):
            resource.version = Primitive.to_proto(self.version)
        if AwsNodePoolConfig.to_proto(self.config):
            resource.config.CopyFrom(AwsNodePoolConfig.to_proto(self.config))
        else:
            resource.ClearField("config")
        if AwsNodePoolAutoscaling.to_proto(self.autoscaling):
            resource.autoscaling.CopyFrom(
                AwsNodePoolAutoscaling.to_proto(self.autoscaling)
            )
        else:
            resource.ClearField("autoscaling")
        if Primitive.to_proto(self.subnet_id):
            resource.subnet_id = Primitive.to_proto(self.subnet_id)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if AwsNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint):
            resource.max_pods_constraint.CopyFrom(
                AwsNodePoolMaxPodsConstraint.to_proto(self.max_pods_constraint)
            )
        else:
            resource.ClearField("max_pods_constraint")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.aws_cluster):
            resource.aws_cluster = Primitive.to_proto(self.aws_cluster)
        return resource


class AwsNodePoolConfig(object):
    def __init__(
        self,
        instance_type: str = None,
        root_volume: dict = None,
        taints: list = None,
        labels: dict = None,
        tags: dict = None,
        iam_instance_profile: str = None,
        ssh_config: dict = None,
        security_group_ids: list = None,
    ):
        self.instance_type = instance_type
        self.root_volume = root_volume
        self.taints = taints
        self.labels = labels
        self.tags = tags
        self.iam_instance_profile = iam_instance_profile
        self.ssh_config = ssh_config
        self.security_group_ids = security_group_ids

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolConfig()
        if Primitive.to_proto(resource.instance_type):
            res.instance_type = Primitive.to_proto(resource.instance_type)
        if AwsNodePoolConfigRootVolume.to_proto(resource.root_volume):
            res.root_volume.CopyFrom(
                AwsNodePoolConfigRootVolume.to_proto(resource.root_volume)
            )
        else:
            res.ClearField("root_volume")
        if AwsNodePoolConfigTaintsArray.to_proto(resource.taints):
            res.taints.extend(AwsNodePoolConfigTaintsArray.to_proto(resource.taints))
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if Primitive.to_proto(resource.tags):
            res.tags = Primitive.to_proto(resource.tags)
        if Primitive.to_proto(resource.iam_instance_profile):
            res.iam_instance_profile = Primitive.to_proto(resource.iam_instance_profile)
        if AwsNodePoolConfigSshConfig.to_proto(resource.ssh_config):
            res.ssh_config.CopyFrom(
                AwsNodePoolConfigSshConfig.to_proto(resource.ssh_config)
            )
        else:
            res.ClearField("ssh_config")
        if Primitive.to_proto(resource.security_group_ids):
            res.security_group_ids.extend(
                Primitive.to_proto(resource.security_group_ids)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsNodePoolConfig(
            instance_type=Primitive.from_proto(resource.instance_type),
            root_volume=AwsNodePoolConfigRootVolume.from_proto(resource.root_volume),
            taints=AwsNodePoolConfigTaintsArray.from_proto(resource.taints),
            labels=Primitive.from_proto(resource.labels),
            tags=Primitive.from_proto(resource.tags),
            iam_instance_profile=Primitive.from_proto(resource.iam_instance_profile),
            ssh_config=AwsNodePoolConfigSshConfig.from_proto(resource.ssh_config),
            security_group_ids=Primitive.from_proto(resource.security_group_ids),
        )


class AwsNodePoolConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsNodePoolConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsNodePoolConfig.from_proto(i) for i in resources]


class AwsNodePoolConfigRootVolume(object):
    def __init__(
        self,
        size_gib: int = None,
        volume_type: str = None,
        iops: int = None,
        kms_key_arn: str = None,
    ):
        self.size_gib = size_gib
        self.volume_type = volume_type
        self.iops = iops
        self.kms_key_arn = kms_key_arn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolConfigRootVolume()
        if Primitive.to_proto(resource.size_gib):
            res.size_gib = Primitive.to_proto(resource.size_gib)
        if AwsNodePoolConfigRootVolumeVolumeTypeEnum.to_proto(resource.volume_type):
            res.volume_type = AwsNodePoolConfigRootVolumeVolumeTypeEnum.to_proto(
                resource.volume_type
            )
        if Primitive.to_proto(resource.iops):
            res.iops = Primitive.to_proto(resource.iops)
        if Primitive.to_proto(resource.kms_key_arn):
            res.kms_key_arn = Primitive.to_proto(resource.kms_key_arn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsNodePoolConfigRootVolume(
            size_gib=Primitive.from_proto(resource.size_gib),
            volume_type=AwsNodePoolConfigRootVolumeVolumeTypeEnum.from_proto(
                resource.volume_type
            ),
            iops=Primitive.from_proto(resource.iops),
            kms_key_arn=Primitive.from_proto(resource.kms_key_arn),
        )


class AwsNodePoolConfigRootVolumeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsNodePoolConfigRootVolume.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsNodePoolConfigRootVolume.from_proto(i) for i in resources]


class AwsNodePoolConfigTaints(object):
    def __init__(self, key: str = None, value: str = None, effect: str = None):
        self.key = key
        self.value = value
        self.effect = effect

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolConfigTaints()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if AwsNodePoolConfigTaintsEffectEnum.to_proto(resource.effect):
            res.effect = AwsNodePoolConfigTaintsEffectEnum.to_proto(resource.effect)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsNodePoolConfigTaints(
            key=Primitive.from_proto(resource.key),
            value=Primitive.from_proto(resource.value),
            effect=AwsNodePoolConfigTaintsEffectEnum.from_proto(resource.effect),
        )


class AwsNodePoolConfigTaintsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsNodePoolConfigTaints.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsNodePoolConfigTaints.from_proto(i) for i in resources]


class AwsNodePoolConfigSshConfig(object):
    def __init__(self, ec2_key_pair: str = None):
        self.ec2_key_pair = ec2_key_pair

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolConfigSshConfig()
        if Primitive.to_proto(resource.ec2_key_pair):
            res.ec2_key_pair = Primitive.to_proto(resource.ec2_key_pair)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsNodePoolConfigSshConfig(
            ec2_key_pair=Primitive.from_proto(resource.ec2_key_pair),
        )


class AwsNodePoolConfigSshConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsNodePoolConfigSshConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsNodePoolConfigSshConfig.from_proto(i) for i in resources]


class AwsNodePoolAutoscaling(object):
    def __init__(self, min_node_count: int = None, max_node_count: int = None):
        self.min_node_count = min_node_count
        self.max_node_count = max_node_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolAutoscaling()
        if Primitive.to_proto(resource.min_node_count):
            res.min_node_count = Primitive.to_proto(resource.min_node_count)
        if Primitive.to_proto(resource.max_node_count):
            res.max_node_count = Primitive.to_proto(resource.max_node_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsNodePoolAutoscaling(
            min_node_count=Primitive.from_proto(resource.min_node_count),
            max_node_count=Primitive.from_proto(resource.max_node_count),
        )


class AwsNodePoolAutoscalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsNodePoolAutoscaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsNodePoolAutoscaling.from_proto(i) for i in resources]


class AwsNodePoolMaxPodsConstraint(object):
    def __init__(self, max_pods_per_node: int = None):
        self.max_pods_per_node = max_pods_per_node

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolMaxPodsConstraint()
        if Primitive.to_proto(resource.max_pods_per_node):
            res.max_pods_per_node = Primitive.to_proto(resource.max_pods_per_node)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsNodePoolMaxPodsConstraint(
            max_pods_per_node=Primitive.from_proto(resource.max_pods_per_node),
        )


class AwsNodePoolMaxPodsConstraintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsNodePoolMaxPodsConstraint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsNodePoolMaxPodsConstraint.from_proto(i) for i in resources]


class AwsNodePoolConfigRootVolumeVolumeTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum.Value(
            "GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum.Name(
            resource
        )[
            len("GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum") :
        ]


class AwsNodePoolConfigTaintsEffectEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum.Value(
            "GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum.Name(
            resource
        )[
            len("GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum") :
        ]


class AwsNodePoolStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolStateEnum.Value(
            "GkemulticloudBetaAwsNodePoolStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return aws_node_pool_pb2.GkemulticloudBetaAwsNodePoolStateEnum.Name(resource)[
            len("GkemulticloudBetaAwsNodePoolStateEnum") :
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
