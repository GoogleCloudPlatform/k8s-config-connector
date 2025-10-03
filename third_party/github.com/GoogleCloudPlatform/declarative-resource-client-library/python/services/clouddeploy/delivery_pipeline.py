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
from google3.cloud.graphite.mmv2.services.google.clouddeploy import (
    delivery_pipeline_pb2,
)
from google3.cloud.graphite.mmv2.services.google.clouddeploy import (
    delivery_pipeline_pb2_grpc,
)

from typing import List


class DeliveryPipeline(object):
    def __init__(
        self,
        name: str = None,
        uid: str = None,
        description: str = None,
        annotations: dict = None,
        labels: dict = None,
        create_time: str = None,
        update_time: str = None,
        serial_pipeline: dict = None,
        condition: dict = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        suspended: bool = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.annotations = annotations
        self.labels = labels
        self.serial_pipeline = serial_pipeline
        self.project = project
        self.location = location
        self.suspended = suspended
        self.service_account_file = service_account_file

    def apply(self):
        stub = delivery_pipeline_pb2_grpc.ClouddeployDeliveryPipelineServiceStub(
            channel.Channel()
        )
        request = delivery_pipeline_pb2.ApplyClouddeployDeliveryPipelineRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline):
            request.resource.serial_pipeline.CopyFrom(
                DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline)
            )
        else:
            request.resource.ClearField("serial_pipeline")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.suspended):
            request.resource.suspended = Primitive.to_proto(self.suspended)

        request.service_account_file = self.service_account_file

        response = stub.ApplyClouddeployDeliveryPipeline(request)
        self.name = Primitive.from_proto(response.name)
        self.uid = Primitive.from_proto(response.uid)
        self.description = Primitive.from_proto(response.description)
        self.annotations = Primitive.from_proto(response.annotations)
        self.labels = Primitive.from_proto(response.labels)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.serial_pipeline = DeliveryPipelineSerialPipeline.from_proto(
            response.serial_pipeline
        )
        self.condition = DeliveryPipelineCondition.from_proto(response.condition)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.suspended = Primitive.from_proto(response.suspended)

    def delete(self):
        stub = delivery_pipeline_pb2_grpc.ClouddeployDeliveryPipelineServiceStub(
            channel.Channel()
        )
        request = delivery_pipeline_pb2.DeleteClouddeployDeliveryPipelineRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline):
            request.resource.serial_pipeline.CopyFrom(
                DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline)
            )
        else:
            request.resource.ClearField("serial_pipeline")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.suspended):
            request.resource.suspended = Primitive.to_proto(self.suspended)

        response = stub.DeleteClouddeployDeliveryPipeline(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = delivery_pipeline_pb2_grpc.ClouddeployDeliveryPipelineServiceStub(
            channel.Channel()
        )
        request = delivery_pipeline_pb2.ListClouddeployDeliveryPipelineRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListClouddeployDeliveryPipeline(request).items

    def to_proto(self):
        resource = delivery_pipeline_pb2.ClouddeployDeliveryPipeline()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline):
            resource.serial_pipeline.CopyFrom(
                DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline)
            )
        else:
            resource.ClearField("serial_pipeline")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.suspended):
            resource.suspended = Primitive.to_proto(self.suspended)
        return resource


class DeliveryPipelineSerialPipeline(object):
    def __init__(self, stages: list = None):
        self.stages = stages

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipeline()
        if DeliveryPipelineSerialPipelineStagesArray.to_proto(resource.stages):
            res.stages.extend(
                DeliveryPipelineSerialPipelineStagesArray.to_proto(resource.stages)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipeline(
            stages=DeliveryPipelineSerialPipelineStagesArray.from_proto(
                resource.stages
            ),
        )


class DeliveryPipelineSerialPipelineArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DeliveryPipelineSerialPipeline.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DeliveryPipelineSerialPipeline.from_proto(i) for i in resources]


class DeliveryPipelineSerialPipelineStages(object):
    def __init__(
        self,
        target_id: str = None,
        profiles: list = None,
        strategy: dict = None,
        deploy_parameters: list = None,
    ):
        self.target_id = target_id
        self.profiles = profiles
        self.strategy = strategy
        self.deploy_parameters = deploy_parameters

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStages()
        if Primitive.to_proto(resource.target_id):
            res.target_id = Primitive.to_proto(resource.target_id)
        if Primitive.to_proto(resource.profiles):
            res.profiles.extend(Primitive.to_proto(resource.profiles))
        if DeliveryPipelineSerialPipelineStagesStrategy.to_proto(resource.strategy):
            res.strategy.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategy.to_proto(resource.strategy)
            )
        else:
            res.ClearField("strategy")
        if DeliveryPipelineSerialPipelineStagesDeployParametersArray.to_proto(
            resource.deploy_parameters
        ):
            res.deploy_parameters.extend(
                DeliveryPipelineSerialPipelineStagesDeployParametersArray.to_proto(
                    resource.deploy_parameters
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStages(
            target_id=Primitive.from_proto(resource.target_id),
            profiles=Primitive.from_proto(resource.profiles),
            strategy=DeliveryPipelineSerialPipelineStagesStrategy.from_proto(
                resource.strategy
            ),
            deploy_parameters=DeliveryPipelineSerialPipelineStagesDeployParametersArray.from_proto(
                resource.deploy_parameters
            ),
        )


class DeliveryPipelineSerialPipelineStagesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DeliveryPipelineSerialPipelineStages.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DeliveryPipelineSerialPipelineStages.from_proto(i) for i in resources]


class DeliveryPipelineSerialPipelineStagesStrategy(object):
    def __init__(self, standard: dict = None, canary: dict = None):
        self.standard = standard
        self.canary = canary

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategy()
        )
        if DeliveryPipelineSerialPipelineStagesStrategyStandard.to_proto(
            resource.standard
        ):
            res.standard.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyStandard.to_proto(
                    resource.standard
                )
            )
        else:
            res.ClearField("standard")
        if DeliveryPipelineSerialPipelineStagesStrategyCanary.to_proto(resource.canary):
            res.canary.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanary.to_proto(
                    resource.canary
                )
            )
        else:
            res.ClearField("canary")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategy(
            standard=DeliveryPipelineSerialPipelineStagesStrategyStandard.from_proto(
                resource.standard
            ),
            canary=DeliveryPipelineSerialPipelineStagesStrategyCanary.from_proto(
                resource.canary
            ),
        )


class DeliveryPipelineSerialPipelineStagesStrategyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategy.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategy.from_proto(i)
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyStandard(object):
    def __init__(
        self, verify: bool = None, predeploy: dict = None, postdeploy: dict = None
    ):
        self.verify = verify
        self.predeploy = predeploy
        self.postdeploy = postdeploy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard()
        )
        if Primitive.to_proto(resource.verify):
            res.verify = Primitive.to_proto(resource.verify)
        if DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy.to_proto(
            resource.predeploy
        ):
            res.predeploy.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy.to_proto(
                    resource.predeploy
                )
            )
        else:
            res.ClearField("predeploy")
        if DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy.to_proto(
            resource.postdeploy
        ):
            res.postdeploy.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy.to_proto(
                    resource.postdeploy
                )
            )
        else:
            res.ClearField("postdeploy")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyStandard(
            verify=Primitive.from_proto(resource.verify),
            predeploy=DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy.from_proto(
                resource.predeploy
            ),
            postdeploy=DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy.from_proto(
                resource.postdeploy
            ),
        )


class DeliveryPipelineSerialPipelineStagesStrategyStandardArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyStandard.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyStandard.from_proto(i)
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(object):
    def __init__(self, actions: list = None):
        self.actions = actions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy()
        )
        if Primitive.to_proto(resource.actions):
            res.actions.extend(Primitive.to_proto(resource.actions))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(
            actions=Primitive.from_proto(resource.actions),
        )


class DeliveryPipelineSerialPipelineStagesStrategyStandardPredeployArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy.from_proto(i)
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(object):
    def __init__(self, actions: list = None):
        self.actions = actions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy()
        )
        if Primitive.to_proto(resource.actions):
            res.actions.extend(Primitive.to_proto(resource.actions))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(
            actions=Primitive.from_proto(resource.actions),
        )


class DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy.from_proto(i)
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanary(object):
    def __init__(
        self,
        runtime_config: dict = None,
        canary_deployment: dict = None,
        custom_canary_deployment: dict = None,
    ):
        self.runtime_config = runtime_config
        self.canary_deployment = canary_deployment
        self.custom_canary_deployment = custom_canary_deployment

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary()
        )
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig.to_proto(
            resource.runtime_config
        ):
            res.runtime_config.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig.to_proto(
                    resource.runtime_config
                )
            )
        else:
            res.ClearField("runtime_config")
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment.to_proto(
            resource.canary_deployment
        ):
            res.canary_deployment.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment.to_proto(
                    resource.canary_deployment
                )
            )
        else:
            res.ClearField("canary_deployment")
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment.to_proto(
            resource.custom_canary_deployment
        ):
            res.custom_canary_deployment.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment.to_proto(
                    resource.custom_canary_deployment
                )
            )
        else:
            res.ClearField("custom_canary_deployment")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanary(
            runtime_config=DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig.from_proto(
                resource.runtime_config
            ),
            canary_deployment=DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment.from_proto(
                resource.canary_deployment
            ),
            custom_canary_deployment=DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment.from_proto(
                resource.custom_canary_deployment
            ),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanary.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanary.from_proto(i)
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(object):
    def __init__(self, kubernetes: dict = None, cloud_run: dict = None):
        self.kubernetes = kubernetes
        self.cloud_run = cloud_run

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig()
        )
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes.to_proto(
            resource.kubernetes
        ):
            res.kubernetes.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes.to_proto(
                    resource.kubernetes
                )
            )
        else:
            res.ClearField("kubernetes")
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun.to_proto(
            resource.cloud_run
        ):
            res.cloud_run.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun.to_proto(
                    resource.cloud_run
                )
            )
        else:
            res.ClearField("cloud_run")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(
            kubernetes=DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes.from_proto(
                resource.kubernetes
            ),
            cloud_run=DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun.from_proto(
                resource.cloud_run
            ),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(object):
    def __init__(
        self, gateway_service_mesh: dict = None, service_networking: dict = None
    ):
        self.gateway_service_mesh = gateway_service_mesh
        self.service_networking = service_networking

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes()
        )
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh.to_proto(
            resource.gateway_service_mesh
        ):
            res.gateway_service_mesh.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh.to_proto(
                    resource.gateway_service_mesh
                )
            )
        else:
            res.ClearField("gateway_service_mesh")
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking.to_proto(
            resource.service_networking
        ):
            res.service_networking.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking.to_proto(
                    resource.service_networking
                )
            )
        else:
            res.ClearField("service_networking")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(
            gateway_service_mesh=DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh.from_proto(
                resource.gateway_service_mesh
            ),
            service_networking=DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking.from_proto(
                resource.service_networking
            ),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(
    object
):
    def __init__(
        self,
        http_route: str = None,
        service: str = None,
        deployment: str = None,
        route_update_wait_time: str = None,
        stable_cutback_duration: str = None,
    ):
        self.http_route = http_route
        self.service = service
        self.deployment = deployment
        self.route_update_wait_time = route_update_wait_time
        self.stable_cutback_duration = stable_cutback_duration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh()
        )
        if Primitive.to_proto(resource.http_route):
            res.http_route = Primitive.to_proto(resource.http_route)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.deployment):
            res.deployment = Primitive.to_proto(resource.deployment)
        if Primitive.to_proto(resource.route_update_wait_time):
            res.route_update_wait_time = Primitive.to_proto(
                resource.route_update_wait_time
            )
        if Primitive.to_proto(resource.stable_cutback_duration):
            res.stable_cutback_duration = Primitive.to_proto(
                resource.stable_cutback_duration
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(
            http_route=Primitive.from_proto(resource.http_route),
            service=Primitive.from_proto(resource.service),
            deployment=Primitive.from_proto(resource.deployment),
            route_update_wait_time=Primitive.from_proto(
                resource.route_update_wait_time
            ),
            stable_cutback_duration=Primitive.from_proto(
                resource.stable_cutback_duration
            ),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(
    object
):
    def __init__(
        self,
        service: str = None,
        deployment: str = None,
        disable_pod_overprovisioning: bool = None,
    ):
        self.service = service
        self.deployment = deployment
        self.disable_pod_overprovisioning = disable_pod_overprovisioning

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking()
        )
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.deployment):
            res.deployment = Primitive.to_proto(resource.deployment)
        if Primitive.to_proto(resource.disable_pod_overprovisioning):
            res.disable_pod_overprovisioning = Primitive.to_proto(
                resource.disable_pod_overprovisioning
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(
            service=Primitive.from_proto(resource.service),
            deployment=Primitive.from_proto(resource.deployment),
            disable_pod_overprovisioning=Primitive.from_proto(
                resource.disable_pod_overprovisioning
            ),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(object):
    def __init__(
        self,
        automatic_traffic_control: bool = None,
        canary_revision_tags: list = None,
        prior_revision_tags: list = None,
        stable_revision_tags: list = None,
    ):
        self.automatic_traffic_control = automatic_traffic_control
        self.canary_revision_tags = canary_revision_tags
        self.prior_revision_tags = prior_revision_tags
        self.stable_revision_tags = stable_revision_tags

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun()
        )
        if Primitive.to_proto(resource.automatic_traffic_control):
            res.automatic_traffic_control = Primitive.to_proto(
                resource.automatic_traffic_control
            )
        if Primitive.to_proto(resource.canary_revision_tags):
            res.canary_revision_tags.extend(
                Primitive.to_proto(resource.canary_revision_tags)
            )
        if Primitive.to_proto(resource.prior_revision_tags):
            res.prior_revision_tags.extend(
                Primitive.to_proto(resource.prior_revision_tags)
            )
        if Primitive.to_proto(resource.stable_revision_tags):
            res.stable_revision_tags.extend(
                Primitive.to_proto(resource.stable_revision_tags)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(
            automatic_traffic_control=Primitive.from_proto(
                resource.automatic_traffic_control
            ),
            canary_revision_tags=Primitive.from_proto(resource.canary_revision_tags),
            prior_revision_tags=Primitive.from_proto(resource.prior_revision_tags),
            stable_revision_tags=Primitive.from_proto(resource.stable_revision_tags),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(object):
    def __init__(
        self,
        percentages: list = None,
        verify: bool = None,
        predeploy: dict = None,
        postdeploy: dict = None,
    ):
        self.percentages = percentages
        self.verify = verify
        self.predeploy = predeploy
        self.postdeploy = postdeploy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment()
        )
        if int64Array.to_proto(resource.percentages):
            res.percentages.extend(int64Array.to_proto(resource.percentages))
        if Primitive.to_proto(resource.verify):
            res.verify = Primitive.to_proto(resource.verify)
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy.to_proto(
            resource.predeploy
        ):
            res.predeploy.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy.to_proto(
                    resource.predeploy
                )
            )
        else:
            res.ClearField("predeploy")
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy.to_proto(
            resource.postdeploy
        ):
            res.postdeploy.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy.to_proto(
                    resource.postdeploy
                )
            )
        else:
            res.ClearField("postdeploy")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(
            percentages=int64Array.from_proto(resource.percentages),
            verify=Primitive.from_proto(resource.verify),
            predeploy=DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy.from_proto(
                resource.predeploy
            ),
            postdeploy=DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy.from_proto(
                resource.postdeploy
            ),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(
    object
):
    def __init__(self, actions: list = None):
        self.actions = actions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy()
        )
        if Primitive.to_proto(resource.actions):
            res.actions.extend(Primitive.to_proto(resource.actions))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(
                actions=Primitive.from_proto(resource.actions),
            )
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(
    object
):
    def __init__(self, actions: list = None):
        self.actions = actions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy()
        )
        if Primitive.to_proto(resource.actions):
            res.actions.extend(Primitive.to_proto(resource.actions))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(
            actions=Primitive.from_proto(resource.actions),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(object):
    def __init__(self, phase_configs: list = None):
        self.phase_configs = phase_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment()
        )
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsArray.to_proto(
            resource.phase_configs
        ):
            res.phase_configs.extend(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsArray.to_proto(
                    resource.phase_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(
            phase_configs=DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsArray.from_proto(
                resource.phase_configs
            ),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(
    object
):
    def __init__(
        self,
        phase_id: str = None,
        percentage: int = None,
        profiles: list = None,
        verify: bool = None,
        predeploy: dict = None,
        postdeploy: dict = None,
    ):
        self.phase_id = phase_id
        self.percentage = percentage
        self.profiles = profiles
        self.verify = verify
        self.predeploy = predeploy
        self.postdeploy = postdeploy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs()
        )
        if Primitive.to_proto(resource.phase_id):
            res.phase_id = Primitive.to_proto(resource.phase_id)
        if Primitive.to_proto(resource.percentage):
            res.percentage = Primitive.to_proto(resource.percentage)
        if Primitive.to_proto(resource.profiles):
            res.profiles.extend(Primitive.to_proto(resource.profiles))
        if Primitive.to_proto(resource.verify):
            res.verify = Primitive.to_proto(resource.verify)
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy.to_proto(
            resource.predeploy
        ):
            res.predeploy.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy.to_proto(
                    resource.predeploy
                )
            )
        else:
            res.ClearField("predeploy")
        if DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy.to_proto(
            resource.postdeploy
        ):
            res.postdeploy.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy.to_proto(
                    resource.postdeploy
                )
            )
        else:
            res.ClearField("postdeploy")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(
            phase_id=Primitive.from_proto(resource.phase_id),
            percentage=Primitive.from_proto(resource.percentage),
            profiles=Primitive.from_proto(resource.profiles),
            verify=Primitive.from_proto(resource.verify),
            predeploy=DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy.from_proto(
                resource.predeploy
            ),
            postdeploy=DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy.from_proto(
                resource.postdeploy
            ),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(
    object
):
    def __init__(self, actions: list = None):
        self.actions = actions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy()
        )
        if Primitive.to_proto(resource.actions):
            res.actions.extend(Primitive.to_proto(resource.actions))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(
            actions=Primitive.from_proto(resource.actions),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(
    object
):
    def __init__(self, actions: list = None):
        self.actions = actions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy()
        )
        if Primitive.to_proto(resource.actions):
            res.actions.extend(Primitive.to_proto(resource.actions))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(
            actions=Primitive.from_proto(resource.actions),
        )


class DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy.from_proto(
                i
            )
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesDeployParameters(object):
    def __init__(self, values: dict = None, match_target_labels: dict = None):
        self.values = values
        self.match_target_labels = match_target_labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineSerialPipelineStagesDeployParameters()
        )
        if Primitive.to_proto(resource.values):
            res.values = Primitive.to_proto(resource.values)
        if Primitive.to_proto(resource.match_target_labels):
            res.match_target_labels = Primitive.to_proto(resource.match_target_labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesDeployParameters(
            values=Primitive.from_proto(resource.values),
            match_target_labels=Primitive.from_proto(resource.match_target_labels),
        )


class DeliveryPipelineSerialPipelineStagesDeployParametersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesDeployParameters.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesDeployParameters.from_proto(i)
            for i in resources
        ]


class DeliveryPipelineCondition(object):
    def __init__(
        self,
        pipeline_ready_condition: dict = None,
        targets_present_condition: dict = None,
        targets_type_condition: dict = None,
    ):
        self.pipeline_ready_condition = pipeline_ready_condition
        self.targets_present_condition = targets_present_condition
        self.targets_type_condition = targets_type_condition

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = delivery_pipeline_pb2.ClouddeployDeliveryPipelineCondition()
        if DeliveryPipelineConditionPipelineReadyCondition.to_proto(
            resource.pipeline_ready_condition
        ):
            res.pipeline_ready_condition.CopyFrom(
                DeliveryPipelineConditionPipelineReadyCondition.to_proto(
                    resource.pipeline_ready_condition
                )
            )
        else:
            res.ClearField("pipeline_ready_condition")
        if DeliveryPipelineConditionTargetsPresentCondition.to_proto(
            resource.targets_present_condition
        ):
            res.targets_present_condition.CopyFrom(
                DeliveryPipelineConditionTargetsPresentCondition.to_proto(
                    resource.targets_present_condition
                )
            )
        else:
            res.ClearField("targets_present_condition")
        if DeliveryPipelineConditionTargetsTypeCondition.to_proto(
            resource.targets_type_condition
        ):
            res.targets_type_condition.CopyFrom(
                DeliveryPipelineConditionTargetsTypeCondition.to_proto(
                    resource.targets_type_condition
                )
            )
        else:
            res.ClearField("targets_type_condition")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineCondition(
            pipeline_ready_condition=DeliveryPipelineConditionPipelineReadyCondition.from_proto(
                resource.pipeline_ready_condition
            ),
            targets_present_condition=DeliveryPipelineConditionTargetsPresentCondition.from_proto(
                resource.targets_present_condition
            ),
            targets_type_condition=DeliveryPipelineConditionTargetsTypeCondition.from_proto(
                resource.targets_type_condition
            ),
        )


class DeliveryPipelineConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DeliveryPipelineCondition.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DeliveryPipelineCondition.from_proto(i) for i in resources]


class DeliveryPipelineConditionPipelineReadyCondition(object):
    def __init__(self, status: bool = None, update_time: str = None):
        self.status = status
        self.update_time = update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineConditionPipelineReadyCondition()
        )
        if Primitive.to_proto(resource.status):
            res.status = Primitive.to_proto(resource.status)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineConditionPipelineReadyCondition(
            status=Primitive.from_proto(resource.status),
            update_time=Primitive.from_proto(resource.update_time),
        )


class DeliveryPipelineConditionPipelineReadyConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineConditionPipelineReadyCondition.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineConditionPipelineReadyCondition.from_proto(i)
            for i in resources
        ]


class DeliveryPipelineConditionTargetsPresentCondition(object):
    def __init__(
        self, status: bool = None, missing_targets: list = None, update_time: str = None
    ):
        self.status = status
        self.missing_targets = missing_targets
        self.update_time = update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineConditionTargetsPresentCondition()
        )
        if Primitive.to_proto(resource.status):
            res.status = Primitive.to_proto(resource.status)
        if Primitive.to_proto(resource.missing_targets):
            res.missing_targets.extend(Primitive.to_proto(resource.missing_targets))
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineConditionTargetsPresentCondition(
            status=Primitive.from_proto(resource.status),
            missing_targets=Primitive.from_proto(resource.missing_targets),
            update_time=Primitive.from_proto(resource.update_time),
        )


class DeliveryPipelineConditionTargetsPresentConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineConditionTargetsPresentCondition.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineConditionTargetsPresentCondition.from_proto(i)
            for i in resources
        ]


class DeliveryPipelineConditionTargetsTypeCondition(object):
    def __init__(self, status: bool = None, error_details: str = None):
        self.status = status
        self.error_details = error_details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployDeliveryPipelineConditionTargetsTypeCondition()
        )
        if Primitive.to_proto(resource.status):
            res.status = Primitive.to_proto(resource.status)
        if Primitive.to_proto(resource.error_details):
            res.error_details = Primitive.to_proto(resource.error_details)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineConditionTargetsTypeCondition(
            status=Primitive.from_proto(resource.status),
            error_details=Primitive.from_proto(resource.error_details),
        )


class DeliveryPipelineConditionTargetsTypeConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineConditionTargetsTypeCondition.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineConditionTargetsTypeCondition.from_proto(i)
            for i in resources
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
