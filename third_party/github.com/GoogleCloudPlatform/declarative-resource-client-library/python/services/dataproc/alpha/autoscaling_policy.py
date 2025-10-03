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
from google3.cloud.graphite.mmv2.services.google.dataproc import autoscaling_policy_pb2
from google3.cloud.graphite.mmv2.services.google.dataproc import (
    autoscaling_policy_pb2_grpc,
)

from typing import List


class AutoscalingPolicy(object):
    def __init__(
        self,
        name: str = None,
        basic_algorithm: dict = None,
        worker_config: dict = None,
        secondary_worker_config: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.basic_algorithm = basic_algorithm
        self.worker_config = worker_config
        self.secondary_worker_config = secondary_worker_config
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = autoscaling_policy_pb2_grpc.DataprocAlphaAutoscalingPolicyServiceStub(
            channel.Channel()
        )
        request = autoscaling_policy_pb2.ApplyDataprocAlphaAutoscalingPolicyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if AutoscalingPolicyBasicAlgorithm.to_proto(self.basic_algorithm):
            request.resource.basic_algorithm.CopyFrom(
                AutoscalingPolicyBasicAlgorithm.to_proto(self.basic_algorithm)
            )
        else:
            request.resource.ClearField("basic_algorithm")
        if AutoscalingPolicyWorkerConfig.to_proto(self.worker_config):
            request.resource.worker_config.CopyFrom(
                AutoscalingPolicyWorkerConfig.to_proto(self.worker_config)
            )
        else:
            request.resource.ClearField("worker_config")
        if AutoscalingPolicySecondaryWorkerConfig.to_proto(
            self.secondary_worker_config
        ):
            request.resource.secondary_worker_config.CopyFrom(
                AutoscalingPolicySecondaryWorkerConfig.to_proto(
                    self.secondary_worker_config
                )
            )
        else:
            request.resource.ClearField("secondary_worker_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDataprocAlphaAutoscalingPolicy(request)
        self.name = Primitive.from_proto(response.name)
        self.basic_algorithm = AutoscalingPolicyBasicAlgorithm.from_proto(
            response.basic_algorithm
        )
        self.worker_config = AutoscalingPolicyWorkerConfig.from_proto(
            response.worker_config
        )
        self.secondary_worker_config = (
            AutoscalingPolicySecondaryWorkerConfig.from_proto(
                response.secondary_worker_config
            )
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = autoscaling_policy_pb2_grpc.DataprocAlphaAutoscalingPolicyServiceStub(
            channel.Channel()
        )
        request = autoscaling_policy_pb2.DeleteDataprocAlphaAutoscalingPolicyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if AutoscalingPolicyBasicAlgorithm.to_proto(self.basic_algorithm):
            request.resource.basic_algorithm.CopyFrom(
                AutoscalingPolicyBasicAlgorithm.to_proto(self.basic_algorithm)
            )
        else:
            request.resource.ClearField("basic_algorithm")
        if AutoscalingPolicyWorkerConfig.to_proto(self.worker_config):
            request.resource.worker_config.CopyFrom(
                AutoscalingPolicyWorkerConfig.to_proto(self.worker_config)
            )
        else:
            request.resource.ClearField("worker_config")
        if AutoscalingPolicySecondaryWorkerConfig.to_proto(
            self.secondary_worker_config
        ):
            request.resource.secondary_worker_config.CopyFrom(
                AutoscalingPolicySecondaryWorkerConfig.to_proto(
                    self.secondary_worker_config
                )
            )
        else:
            request.resource.ClearField("secondary_worker_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteDataprocAlphaAutoscalingPolicy(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = autoscaling_policy_pb2_grpc.DataprocAlphaAutoscalingPolicyServiceStub(
            channel.Channel()
        )
        request = autoscaling_policy_pb2.ListDataprocAlphaAutoscalingPolicyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListDataprocAlphaAutoscalingPolicy(request).items

    def to_proto(self):
        resource = autoscaling_policy_pb2.DataprocAlphaAutoscalingPolicy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if AutoscalingPolicyBasicAlgorithm.to_proto(self.basic_algorithm):
            resource.basic_algorithm.CopyFrom(
                AutoscalingPolicyBasicAlgorithm.to_proto(self.basic_algorithm)
            )
        else:
            resource.ClearField("basic_algorithm")
        if AutoscalingPolicyWorkerConfig.to_proto(self.worker_config):
            resource.worker_config.CopyFrom(
                AutoscalingPolicyWorkerConfig.to_proto(self.worker_config)
            )
        else:
            resource.ClearField("worker_config")
        if AutoscalingPolicySecondaryWorkerConfig.to_proto(
            self.secondary_worker_config
        ):
            resource.secondary_worker_config.CopyFrom(
                AutoscalingPolicySecondaryWorkerConfig.to_proto(
                    self.secondary_worker_config
                )
            )
        else:
            resource.ClearField("secondary_worker_config")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class AutoscalingPolicyBasicAlgorithm(object):
    def __init__(self, yarn_config: dict = None, cooldown_period: str = None):
        self.yarn_config = yarn_config
        self.cooldown_period = cooldown_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = autoscaling_policy_pb2.DataprocAlphaAutoscalingPolicyBasicAlgorithm()
        if AutoscalingPolicyBasicAlgorithmYarnConfig.to_proto(resource.yarn_config):
            res.yarn_config.CopyFrom(
                AutoscalingPolicyBasicAlgorithmYarnConfig.to_proto(resource.yarn_config)
            )
        else:
            res.ClearField("yarn_config")
        if Primitive.to_proto(resource.cooldown_period):
            res.cooldown_period = Primitive.to_proto(resource.cooldown_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AutoscalingPolicyBasicAlgorithm(
            yarn_config=AutoscalingPolicyBasicAlgorithmYarnConfig.from_proto(
                resource.yarn_config
            ),
            cooldown_period=Primitive.from_proto(resource.cooldown_period),
        )


class AutoscalingPolicyBasicAlgorithmArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AutoscalingPolicyBasicAlgorithm.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AutoscalingPolicyBasicAlgorithm.from_proto(i) for i in resources]


class AutoscalingPolicyBasicAlgorithmYarnConfig(object):
    def __init__(
        self,
        graceful_decommission_timeout: str = None,
        scale_up_factor: float = None,
        scale_down_factor: float = None,
        scale_up_min_worker_fraction: float = None,
        scale_down_min_worker_fraction: float = None,
    ):
        self.graceful_decommission_timeout = graceful_decommission_timeout
        self.scale_up_factor = scale_up_factor
        self.scale_down_factor = scale_down_factor
        self.scale_up_min_worker_fraction = scale_up_min_worker_fraction
        self.scale_down_min_worker_fraction = scale_down_min_worker_fraction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            autoscaling_policy_pb2.DataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfig()
        )
        if Primitive.to_proto(resource.graceful_decommission_timeout):
            res.graceful_decommission_timeout = Primitive.to_proto(
                resource.graceful_decommission_timeout
            )
        if Primitive.to_proto(resource.scale_up_factor):
            res.scale_up_factor = Primitive.to_proto(resource.scale_up_factor)
        if Primitive.to_proto(resource.scale_down_factor):
            res.scale_down_factor = Primitive.to_proto(resource.scale_down_factor)
        if Primitive.to_proto(resource.scale_up_min_worker_fraction):
            res.scale_up_min_worker_fraction = Primitive.to_proto(
                resource.scale_up_min_worker_fraction
            )
        if Primitive.to_proto(resource.scale_down_min_worker_fraction):
            res.scale_down_min_worker_fraction = Primitive.to_proto(
                resource.scale_down_min_worker_fraction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AutoscalingPolicyBasicAlgorithmYarnConfig(
            graceful_decommission_timeout=Primitive.from_proto(
                resource.graceful_decommission_timeout
            ),
            scale_up_factor=Primitive.from_proto(resource.scale_up_factor),
            scale_down_factor=Primitive.from_proto(resource.scale_down_factor),
            scale_up_min_worker_fraction=Primitive.from_proto(
                resource.scale_up_min_worker_fraction
            ),
            scale_down_min_worker_fraction=Primitive.from_proto(
                resource.scale_down_min_worker_fraction
            ),
        )


class AutoscalingPolicyBasicAlgorithmYarnConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AutoscalingPolicyBasicAlgorithmYarnConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AutoscalingPolicyBasicAlgorithmYarnConfig.from_proto(i) for i in resources
        ]


class AutoscalingPolicyWorkerConfig(object):
    def __init__(
        self, min_instances: int = None, max_instances: int = None, weight: int = None
    ):
        self.min_instances = min_instances
        self.max_instances = max_instances
        self.weight = weight

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = autoscaling_policy_pb2.DataprocAlphaAutoscalingPolicyWorkerConfig()
        if Primitive.to_proto(resource.min_instances):
            res.min_instances = Primitive.to_proto(resource.min_instances)
        if Primitive.to_proto(resource.max_instances):
            res.max_instances = Primitive.to_proto(resource.max_instances)
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AutoscalingPolicyWorkerConfig(
            min_instances=Primitive.from_proto(resource.min_instances),
            max_instances=Primitive.from_proto(resource.max_instances),
            weight=Primitive.from_proto(resource.weight),
        )


class AutoscalingPolicyWorkerConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AutoscalingPolicyWorkerConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AutoscalingPolicyWorkerConfig.from_proto(i) for i in resources]


class AutoscalingPolicySecondaryWorkerConfig(object):
    def __init__(
        self, min_instances: int = None, max_instances: int = None, weight: int = None
    ):
        self.min_instances = min_instances
        self.max_instances = max_instances
        self.weight = weight

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            autoscaling_policy_pb2.DataprocAlphaAutoscalingPolicySecondaryWorkerConfig()
        )
        if Primitive.to_proto(resource.min_instances):
            res.min_instances = Primitive.to_proto(resource.min_instances)
        if Primitive.to_proto(resource.max_instances):
            res.max_instances = Primitive.to_proto(resource.max_instances)
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AutoscalingPolicySecondaryWorkerConfig(
            min_instances=Primitive.from_proto(resource.min_instances),
            max_instances=Primitive.from_proto(resource.max_instances),
            weight=Primitive.from_proto(resource.weight),
        )


class AutoscalingPolicySecondaryWorkerConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AutoscalingPolicySecondaryWorkerConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AutoscalingPolicySecondaryWorkerConfig.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
