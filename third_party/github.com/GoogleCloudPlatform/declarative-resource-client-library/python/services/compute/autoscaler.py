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
from google3.cloud.graphite.mmv2.services.google.compute import autoscaler_pb2
from google3.cloud.graphite.mmv2.services.google.compute import autoscaler_pb2_grpc

from typing import List


class Autoscaler(object):
    def __init__(
        self,
        id: int = None,
        name: str = None,
        description: str = None,
        target: str = None,
        autoscaling_policy: dict = None,
        zone: str = None,
        region: str = None,
        self_link: str = None,
        status: str = None,
        status_details: list = None,
        recommended_size: int = None,
        self_link_with_id: str = None,
        scaling_schedule_status: dict = None,
        project: str = None,
        creation_timestamp: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.target = target
        self.autoscaling_policy = autoscaling_policy
        self.zone = zone
        self.region = region
        self.self_link_with_id = self_link_with_id
        self.scaling_schedule_status = scaling_schedule_status
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = autoscaler_pb2_grpc.ComputeAutoscalerServiceStub(channel.Channel())
        request = autoscaler_pb2.ApplyComputeAutoscalerRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.target):
            request.resource.target = Primitive.to_proto(self.target)

        if AutoscalerAutoscalingPolicy.to_proto(self.autoscaling_policy):
            request.resource.autoscaling_policy.CopyFrom(
                AutoscalerAutoscalingPolicy.to_proto(self.autoscaling_policy)
            )
        else:
            request.resource.ClearField("autoscaling_policy")
        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.self_link_with_id):
            request.resource.self_link_with_id = Primitive.to_proto(
                self.self_link_with_id
            )

        if Primitive.to_proto(self.scaling_schedule_status):
            request.resource.scaling_schedule_status = Primitive.to_proto(
                self.scaling_schedule_status
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeAutoscaler(request)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.target = Primitive.from_proto(response.target)
        self.autoscaling_policy = AutoscalerAutoscalingPolicy.from_proto(
            response.autoscaling_policy
        )
        self.zone = Primitive.from_proto(response.zone)
        self.region = Primitive.from_proto(response.region)
        self.self_link = Primitive.from_proto(response.self_link)
        self.status = AutoscalerStatusEnum.from_proto(response.status)
        self.status_details = AutoscalerStatusDetailsArray.from_proto(
            response.status_details
        )
        self.recommended_size = Primitive.from_proto(response.recommended_size)
        self.self_link_with_id = Primitive.from_proto(response.self_link_with_id)
        self.scaling_schedule_status = Primitive.from_proto(
            response.scaling_schedule_status
        )
        self.project = Primitive.from_proto(response.project)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = autoscaler_pb2_grpc.ComputeAutoscalerServiceStub(channel.Channel())
        request = autoscaler_pb2.DeleteComputeAutoscalerRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.target):
            request.resource.target = Primitive.to_proto(self.target)

        if AutoscalerAutoscalingPolicy.to_proto(self.autoscaling_policy):
            request.resource.autoscaling_policy.CopyFrom(
                AutoscalerAutoscalingPolicy.to_proto(self.autoscaling_policy)
            )
        else:
            request.resource.ClearField("autoscaling_policy")
        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.self_link_with_id):
            request.resource.self_link_with_id = Primitive.to_proto(
                self.self_link_with_id
            )

        if Primitive.to_proto(self.scaling_schedule_status):
            request.resource.scaling_schedule_status = Primitive.to_proto(
                self.scaling_schedule_status
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteComputeAutoscaler(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = autoscaler_pb2_grpc.ComputeAutoscalerServiceStub(channel.Channel())
        request = autoscaler_pb2.ListComputeAutoscalerRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeAutoscaler(request).items

    def to_proto(self):
        resource = autoscaler_pb2.ComputeAutoscaler()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.target):
            resource.target = Primitive.to_proto(self.target)
        if AutoscalerAutoscalingPolicy.to_proto(self.autoscaling_policy):
            resource.autoscaling_policy.CopyFrom(
                AutoscalerAutoscalingPolicy.to_proto(self.autoscaling_policy)
            )
        else:
            resource.ClearField("autoscaling_policy")
        if Primitive.to_proto(self.zone):
            resource.zone = Primitive.to_proto(self.zone)
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.self_link_with_id):
            resource.self_link_with_id = Primitive.to_proto(self.self_link_with_id)
        if Primitive.to_proto(self.scaling_schedule_status):
            resource.scaling_schedule_status = Primitive.to_proto(
                self.scaling_schedule_status
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class AutoscalerAutoscalingPolicy(object):
    def __init__(
        self,
        min_num_replicas: int = None,
        max_num_replicas: int = None,
        scale_in_control: dict = None,
        cool_down_period_sec: int = None,
        cpu_utilization: dict = None,
        custom_metric_utilizations: list = None,
        load_balancing_utilization: dict = None,
        mode: str = None,
    ):
        self.min_num_replicas = min_num_replicas
        self.max_num_replicas = max_num_replicas
        self.scale_in_control = scale_in_control
        self.cool_down_period_sec = cool_down_period_sec
        self.cpu_utilization = cpu_utilization
        self.custom_metric_utilizations = custom_metric_utilizations
        self.load_balancing_utilization = load_balancing_utilization
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = autoscaler_pb2.ComputeAutoscalerAutoscalingPolicy()
        if Primitive.to_proto(resource.min_num_replicas):
            res.min_num_replicas = Primitive.to_proto(resource.min_num_replicas)
        if Primitive.to_proto(resource.max_num_replicas):
            res.max_num_replicas = Primitive.to_proto(resource.max_num_replicas)
        if AutoscalerAutoscalingPolicyScaleInControl.to_proto(
            resource.scale_in_control
        ):
            res.scale_in_control.CopyFrom(
                AutoscalerAutoscalingPolicyScaleInControl.to_proto(
                    resource.scale_in_control
                )
            )
        else:
            res.ClearField("scale_in_control")
        if Primitive.to_proto(resource.cool_down_period_sec):
            res.cool_down_period_sec = Primitive.to_proto(resource.cool_down_period_sec)
        if AutoscalerAutoscalingPolicyCpuUtilization.to_proto(resource.cpu_utilization):
            res.cpu_utilization.CopyFrom(
                AutoscalerAutoscalingPolicyCpuUtilization.to_proto(
                    resource.cpu_utilization
                )
            )
        else:
            res.ClearField("cpu_utilization")
        if AutoscalerAutoscalingPolicyCustomMetricUtilizationsArray.to_proto(
            resource.custom_metric_utilizations
        ):
            res.custom_metric_utilizations.extend(
                AutoscalerAutoscalingPolicyCustomMetricUtilizationsArray.to_proto(
                    resource.custom_metric_utilizations
                )
            )
        if AutoscalerAutoscalingPolicyLoadBalancingUtilization.to_proto(
            resource.load_balancing_utilization
        ):
            res.load_balancing_utilization.CopyFrom(
                AutoscalerAutoscalingPolicyLoadBalancingUtilization.to_proto(
                    resource.load_balancing_utilization
                )
            )
        else:
            res.ClearField("load_balancing_utilization")
        if AutoscalerAutoscalingPolicyModeEnum.to_proto(resource.mode):
            res.mode = AutoscalerAutoscalingPolicyModeEnum.to_proto(resource.mode)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AutoscalerAutoscalingPolicy(
            min_num_replicas=Primitive.from_proto(resource.min_num_replicas),
            max_num_replicas=Primitive.from_proto(resource.max_num_replicas),
            scale_in_control=AutoscalerAutoscalingPolicyScaleInControl.from_proto(
                resource.scale_in_control
            ),
            cool_down_period_sec=Primitive.from_proto(resource.cool_down_period_sec),
            cpu_utilization=AutoscalerAutoscalingPolicyCpuUtilization.from_proto(
                resource.cpu_utilization
            ),
            custom_metric_utilizations=AutoscalerAutoscalingPolicyCustomMetricUtilizationsArray.from_proto(
                resource.custom_metric_utilizations
            ),
            load_balancing_utilization=AutoscalerAutoscalingPolicyLoadBalancingUtilization.from_proto(
                resource.load_balancing_utilization
            ),
            mode=AutoscalerAutoscalingPolicyModeEnum.from_proto(resource.mode),
        )


class AutoscalerAutoscalingPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AutoscalerAutoscalingPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AutoscalerAutoscalingPolicy.from_proto(i) for i in resources]


class AutoscalerAutoscalingPolicyScaleInControl(object):
    def __init__(
        self, max_scaled_in_replicas: dict = None, time_window_sec: int = None
    ):
        self.max_scaled_in_replicas = max_scaled_in_replicas
        self.time_window_sec = time_window_sec

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = autoscaler_pb2.ComputeAutoscalerAutoscalingPolicyScaleInControl()
        if AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas.to_proto(
            resource.max_scaled_in_replicas
        ):
            res.max_scaled_in_replicas.CopyFrom(
                AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas.to_proto(
                    resource.max_scaled_in_replicas
                )
            )
        else:
            res.ClearField("max_scaled_in_replicas")
        if Primitive.to_proto(resource.time_window_sec):
            res.time_window_sec = Primitive.to_proto(resource.time_window_sec)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AutoscalerAutoscalingPolicyScaleInControl(
            max_scaled_in_replicas=AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas.from_proto(
                resource.max_scaled_in_replicas
            ),
            time_window_sec=Primitive.from_proto(resource.time_window_sec),
        )


class AutoscalerAutoscalingPolicyScaleInControlArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AutoscalerAutoscalingPolicyScaleInControl.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AutoscalerAutoscalingPolicyScaleInControl.from_proto(i) for i in resources
        ]


class AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(object):
    def __init__(self, fixed: int = None, percent: int = None, calculated: int = None):
        self.fixed = fixed
        self.percent = percent
        self.calculated = calculated

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            autoscaler_pb2.ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas()
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

        return AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(
            fixed=Primitive.from_proto(resource.fixed),
            percent=Primitive.from_proto(resource.percent),
            calculated=Primitive.from_proto(resource.calculated),
        )


class AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas.from_proto(i)
            for i in resources
        ]


class AutoscalerAutoscalingPolicyCpuUtilization(object):
    def __init__(self, utilization_target: float = None):
        self.utilization_target = utilization_target

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = autoscaler_pb2.ComputeAutoscalerAutoscalingPolicyCpuUtilization()
        if Primitive.to_proto(resource.utilization_target):
            res.utilization_target = Primitive.to_proto(resource.utilization_target)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AutoscalerAutoscalingPolicyCpuUtilization(
            utilization_target=Primitive.from_proto(resource.utilization_target),
        )


class AutoscalerAutoscalingPolicyCpuUtilizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AutoscalerAutoscalingPolicyCpuUtilization.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AutoscalerAutoscalingPolicyCpuUtilization.from_proto(i) for i in resources
        ]


class AutoscalerAutoscalingPolicyCustomMetricUtilizations(object):
    def __init__(
        self,
        metric: str = None,
        utilization_target: float = None,
        utilization_target_type: str = None,
    ):
        self.metric = metric
        self.utilization_target = utilization_target
        self.utilization_target_type = utilization_target_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            autoscaler_pb2.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizations()
        )
        if Primitive.to_proto(resource.metric):
            res.metric = Primitive.to_proto(resource.metric)
        if Primitive.to_proto(resource.utilization_target):
            res.utilization_target = Primitive.to_proto(resource.utilization_target)
        if AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum.to_proto(
            resource.utilization_target_type
        ):
            res.utilization_target_type = AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum.to_proto(
                resource.utilization_target_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AutoscalerAutoscalingPolicyCustomMetricUtilizations(
            metric=Primitive.from_proto(resource.metric),
            utilization_target=Primitive.from_proto(resource.utilization_target),
            utilization_target_type=AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum.from_proto(
                resource.utilization_target_type
            ),
        )


class AutoscalerAutoscalingPolicyCustomMetricUtilizationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AutoscalerAutoscalingPolicyCustomMetricUtilizations.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AutoscalerAutoscalingPolicyCustomMetricUtilizations.from_proto(i)
            for i in resources
        ]


class AutoscalerAutoscalingPolicyLoadBalancingUtilization(object):
    def __init__(self, utilization_target: float = None):
        self.utilization_target = utilization_target

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            autoscaler_pb2.ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization()
        )
        if Primitive.to_proto(resource.utilization_target):
            res.utilization_target = Primitive.to_proto(resource.utilization_target)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AutoscalerAutoscalingPolicyLoadBalancingUtilization(
            utilization_target=Primitive.from_proto(resource.utilization_target),
        )


class AutoscalerAutoscalingPolicyLoadBalancingUtilizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AutoscalerAutoscalingPolicyLoadBalancingUtilization.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AutoscalerAutoscalingPolicyLoadBalancingUtilization.from_proto(i)
            for i in resources
        ]


class AutoscalerStatusDetails(object):
    def __init__(self, message: str = None, type: str = None):
        self.message = message
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = autoscaler_pb2.ComputeAutoscalerStatusDetails()
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if AutoscalerStatusDetailsTypeEnum.to_proto(resource.type):
            res.type = AutoscalerStatusDetailsTypeEnum.to_proto(resource.type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AutoscalerStatusDetails(
            message=Primitive.from_proto(resource.message),
            type=AutoscalerStatusDetailsTypeEnum.from_proto(resource.type),
        )


class AutoscalerStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AutoscalerStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AutoscalerStatusDetails.from_proto(i) for i in resources]


class AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return autoscaler_pb2.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum.Value(
            "ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return autoscaler_pb2.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum.Name(
            resource
        )[
            len(
                "ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum"
            ) :
        ]


class AutoscalerAutoscalingPolicyModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return autoscaler_pb2.ComputeAutoscalerAutoscalingPolicyModeEnum.Value(
            "ComputeAutoscalerAutoscalingPolicyModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return autoscaler_pb2.ComputeAutoscalerAutoscalingPolicyModeEnum.Name(resource)[
            len("ComputeAutoscalerAutoscalingPolicyModeEnum") :
        ]


class AutoscalerStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return autoscaler_pb2.ComputeAutoscalerStatusEnum.Value(
            "ComputeAutoscalerStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return autoscaler_pb2.ComputeAutoscalerStatusEnum.Name(resource)[
            len("ComputeAutoscalerStatusEnum") :
        ]


class AutoscalerStatusDetailsTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return autoscaler_pb2.ComputeAutoscalerStatusDetailsTypeEnum.Value(
            "ComputeAutoscalerStatusDetailsTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return autoscaler_pb2.ComputeAutoscalerStatusDetailsTypeEnum.Name(resource)[
            len("ComputeAutoscalerStatusDetailsTypeEnum") :
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
