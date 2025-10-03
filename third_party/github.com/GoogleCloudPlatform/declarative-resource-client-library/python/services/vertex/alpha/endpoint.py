# Copyright 2022 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.vertex import endpoint_pb2
from google3.cloud.graphite.mmv2.services.google.vertex import endpoint_pb2_grpc

from typing import List


class Endpoint(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        description: str = None,
        deployed_models: list = None,
        etag: str = None,
        labels: dict = None,
        create_time: str = None,
        update_time: str = None,
        encryption_spec: dict = None,
        network: str = None,
        model_deployment_monitoring_job: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.description = description
        self.labels = labels
        self.encryption_spec = encryption_spec
        self.network = network
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = endpoint_pb2_grpc.VertexAlphaEndpointServiceStub(channel.Channel())
        request = endpoint_pb2.ApplyVertexAlphaEndpointRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if EndpointEncryptionSpec.to_proto(self.encryption_spec):
            request.resource.encryption_spec.CopyFrom(
                EndpointEncryptionSpec.to_proto(self.encryption_spec)
            )
        else:
            request.resource.ClearField("encryption_spec")
        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyVertexAlphaEndpoint(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.deployed_models = EndpointDeployedModelsArray.from_proto(
            response.deployed_models
        )
        self.etag = Primitive.from_proto(response.etag)
        self.labels = Primitive.from_proto(response.labels)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.encryption_spec = EndpointEncryptionSpec.from_proto(
            response.encryption_spec
        )
        self.network = Primitive.from_proto(response.network)
        self.model_deployment_monitoring_job = Primitive.from_proto(
            response.model_deployment_monitoring_job
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = endpoint_pb2_grpc.VertexAlphaEndpointServiceStub(channel.Channel())
        request = endpoint_pb2.DeleteVertexAlphaEndpointRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if EndpointEncryptionSpec.to_proto(self.encryption_spec):
            request.resource.encryption_spec.CopyFrom(
                EndpointEncryptionSpec.to_proto(self.encryption_spec)
            )
        else:
            request.resource.ClearField("encryption_spec")
        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteVertexAlphaEndpoint(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = endpoint_pb2_grpc.VertexAlphaEndpointServiceStub(channel.Channel())
        request = endpoint_pb2.ListVertexAlphaEndpointRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListVertexAlphaEndpoint(request).items

    def to_proto(self):
        resource = endpoint_pb2.VertexAlphaEndpoint()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if EndpointEncryptionSpec.to_proto(self.encryption_spec):
            resource.encryption_spec.CopyFrom(
                EndpointEncryptionSpec.to_proto(self.encryption_spec)
            )
        else:
            resource.ClearField("encryption_spec")
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class EndpointDeployedModels(object):
    def __init__(
        self,
        dedicated_resources: dict = None,
        automatic_resources: dict = None,
        id: str = None,
        model: str = None,
        model_version_id: str = None,
        display_name: str = None,
        create_time: str = None,
        service_account: str = None,
        disable_container_logging: bool = None,
        enable_access_logging: bool = None,
        private_endpoints: dict = None,
        shared_resources: str = None,
        enable_container_logging: bool = None,
    ):
        self.dedicated_resources = dedicated_resources
        self.automatic_resources = automatic_resources
        self.id = id
        self.model = model
        self.model_version_id = model_version_id
        self.display_name = display_name
        self.create_time = create_time
        self.service_account = service_account
        self.disable_container_logging = disable_container_logging
        self.enable_access_logging = enable_access_logging
        self.private_endpoints = private_endpoints
        self.shared_resources = shared_resources
        self.enable_container_logging = enable_container_logging

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = endpoint_pb2.VertexAlphaEndpointDeployedModels()
        if EndpointDeployedModelsDedicatedResources.to_proto(
            resource.dedicated_resources
        ):
            res.dedicated_resources.CopyFrom(
                EndpointDeployedModelsDedicatedResources.to_proto(
                    resource.dedicated_resources
                )
            )
        else:
            res.ClearField("dedicated_resources")
        if EndpointDeployedModelsAutomaticResources.to_proto(
            resource.automatic_resources
        ):
            res.automatic_resources.CopyFrom(
                EndpointDeployedModelsAutomaticResources.to_proto(
                    resource.automatic_resources
                )
            )
        else:
            res.ClearField("automatic_resources")
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.model):
            res.model = Primitive.to_proto(resource.model)
        if Primitive.to_proto(resource.model_version_id):
            res.model_version_id = Primitive.to_proto(resource.model_version_id)
        if Primitive.to_proto(resource.display_name):
            res.display_name = Primitive.to_proto(resource.display_name)
        if Primitive.to_proto(resource.create_time):
            res.create_time = Primitive.to_proto(resource.create_time)
        if Primitive.to_proto(resource.service_account):
            res.service_account = Primitive.to_proto(resource.service_account)
        if Primitive.to_proto(resource.disable_container_logging):
            res.disable_container_logging = Primitive.to_proto(
                resource.disable_container_logging
            )
        if Primitive.to_proto(resource.enable_access_logging):
            res.enable_access_logging = Primitive.to_proto(
                resource.enable_access_logging
            )
        if EndpointDeployedModelsPrivateEndpoints.to_proto(resource.private_endpoints):
            res.private_endpoints.CopyFrom(
                EndpointDeployedModelsPrivateEndpoints.to_proto(
                    resource.private_endpoints
                )
            )
        else:
            res.ClearField("private_endpoints")
        if Primitive.to_proto(resource.shared_resources):
            res.shared_resources = Primitive.to_proto(resource.shared_resources)
        if Primitive.to_proto(resource.enable_container_logging):
            res.enable_container_logging = Primitive.to_proto(
                resource.enable_container_logging
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointDeployedModels(
            dedicated_resources=EndpointDeployedModelsDedicatedResources.from_proto(
                resource.dedicated_resources
            ),
            automatic_resources=EndpointDeployedModelsAutomaticResources.from_proto(
                resource.automatic_resources
            ),
            id=Primitive.from_proto(resource.id),
            model=Primitive.from_proto(resource.model),
            model_version_id=Primitive.from_proto(resource.model_version_id),
            display_name=Primitive.from_proto(resource.display_name),
            create_time=Primitive.from_proto(resource.create_time),
            service_account=Primitive.from_proto(resource.service_account),
            disable_container_logging=Primitive.from_proto(
                resource.disable_container_logging
            ),
            enable_access_logging=Primitive.from_proto(resource.enable_access_logging),
            private_endpoints=EndpointDeployedModelsPrivateEndpoints.from_proto(
                resource.private_endpoints
            ),
            shared_resources=Primitive.from_proto(resource.shared_resources),
            enable_container_logging=Primitive.from_proto(
                resource.enable_container_logging
            ),
        )


class EndpointDeployedModelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EndpointDeployedModels.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EndpointDeployedModels.from_proto(i) for i in resources]


class EndpointDeployedModelsDedicatedResources(object):
    def __init__(
        self,
        machine_spec: dict = None,
        min_replica_count: int = None,
        max_replica_count: int = None,
        autoscaling_metric_specs: list = None,
    ):
        self.machine_spec = machine_spec
        self.min_replica_count = min_replica_count
        self.max_replica_count = max_replica_count
        self.autoscaling_metric_specs = autoscaling_metric_specs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = endpoint_pb2.VertexAlphaEndpointDeployedModelsDedicatedResources()
        if EndpointDeployedModelsDedicatedResourcesMachineSpec.to_proto(
            resource.machine_spec
        ):
            res.machine_spec.CopyFrom(
                EndpointDeployedModelsDedicatedResourcesMachineSpec.to_proto(
                    resource.machine_spec
                )
            )
        else:
            res.ClearField("machine_spec")
        if Primitive.to_proto(resource.min_replica_count):
            res.min_replica_count = Primitive.to_proto(resource.min_replica_count)
        if Primitive.to_proto(resource.max_replica_count):
            res.max_replica_count = Primitive.to_proto(resource.max_replica_count)
        if EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsArray.to_proto(
            resource.autoscaling_metric_specs
        ):
            res.autoscaling_metric_specs.extend(
                EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsArray.to_proto(
                    resource.autoscaling_metric_specs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointDeployedModelsDedicatedResources(
            machine_spec=EndpointDeployedModelsDedicatedResourcesMachineSpec.from_proto(
                resource.machine_spec
            ),
            min_replica_count=Primitive.from_proto(resource.min_replica_count),
            max_replica_count=Primitive.from_proto(resource.max_replica_count),
            autoscaling_metric_specs=EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsArray.from_proto(
                resource.autoscaling_metric_specs
            ),
        )


class EndpointDeployedModelsDedicatedResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EndpointDeployedModelsDedicatedResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            EndpointDeployedModelsDedicatedResources.from_proto(i) for i in resources
        ]


class EndpointDeployedModelsDedicatedResourcesMachineSpec(object):
    def __init__(
        self,
        machine_type: str = None,
        accelerator_type: str = None,
        accelerator_count: int = None,
    ):
        self.machine_type = machine_type
        self.accelerator_type = accelerator_type
        self.accelerator_count = accelerator_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            endpoint_pb2.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpec()
        )
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum.to_proto(
            resource.accelerator_type
        ):
            res.accelerator_type = EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum.to_proto(
                resource.accelerator_type
            )
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointDeployedModelsDedicatedResourcesMachineSpec(
            machine_type=Primitive.from_proto(resource.machine_type),
            accelerator_type=EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum.from_proto(
                resource.accelerator_type
            ),
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
        )


class EndpointDeployedModelsDedicatedResourcesMachineSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EndpointDeployedModelsDedicatedResourcesMachineSpec.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EndpointDeployedModelsDedicatedResourcesMachineSpec.from_proto(i)
            for i in resources
        ]


class EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(object):
    def __init__(self, metric_name: str = None, target: int = None):
        self.metric_name = metric_name
        self.target = target

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            endpoint_pb2.VertexAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.target):
            res.target = Primitive.to_proto(resource.target)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(
            metric_name=Primitive.from_proto(resource.metric_name),
            target=Primitive.from_proto(resource.target),
        )


class EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs.from_proto(i)
            for i in resources
        ]


class EndpointDeployedModelsAutomaticResources(object):
    def __init__(self, min_replica_count: int = None, max_replica_count: int = None):
        self.min_replica_count = min_replica_count
        self.max_replica_count = max_replica_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = endpoint_pb2.VertexAlphaEndpointDeployedModelsAutomaticResources()
        if Primitive.to_proto(resource.min_replica_count):
            res.min_replica_count = Primitive.to_proto(resource.min_replica_count)
        if Primitive.to_proto(resource.max_replica_count):
            res.max_replica_count = Primitive.to_proto(resource.max_replica_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointDeployedModelsAutomaticResources(
            min_replica_count=Primitive.from_proto(resource.min_replica_count),
            max_replica_count=Primitive.from_proto(resource.max_replica_count),
        )


class EndpointDeployedModelsAutomaticResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EndpointDeployedModelsAutomaticResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            EndpointDeployedModelsAutomaticResources.from_proto(i) for i in resources
        ]


class EndpointDeployedModelsPrivateEndpoints(object):
    def __init__(
        self,
        predict_http_uri: str = None,
        explain_http_uri: str = None,
        health_http_uri: str = None,
        service_attachment: str = None,
    ):
        self.predict_http_uri = predict_http_uri
        self.explain_http_uri = explain_http_uri
        self.health_http_uri = health_http_uri
        self.service_attachment = service_attachment

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = endpoint_pb2.VertexAlphaEndpointDeployedModelsPrivateEndpoints()
        if Primitive.to_proto(resource.predict_http_uri):
            res.predict_http_uri = Primitive.to_proto(resource.predict_http_uri)
        if Primitive.to_proto(resource.explain_http_uri):
            res.explain_http_uri = Primitive.to_proto(resource.explain_http_uri)
        if Primitive.to_proto(resource.health_http_uri):
            res.health_http_uri = Primitive.to_proto(resource.health_http_uri)
        if Primitive.to_proto(resource.service_attachment):
            res.service_attachment = Primitive.to_proto(resource.service_attachment)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointDeployedModelsPrivateEndpoints(
            predict_http_uri=Primitive.from_proto(resource.predict_http_uri),
            explain_http_uri=Primitive.from_proto(resource.explain_http_uri),
            health_http_uri=Primitive.from_proto(resource.health_http_uri),
            service_attachment=Primitive.from_proto(resource.service_attachment),
        )


class EndpointDeployedModelsPrivateEndpointsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EndpointDeployedModelsPrivateEndpoints.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EndpointDeployedModelsPrivateEndpoints.from_proto(i) for i in resources]


class EndpointEncryptionSpec(object):
    def __init__(self, kms_key_name: str = None):
        self.kms_key_name = kms_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = endpoint_pb2.VertexAlphaEndpointEncryptionSpec()
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointEncryptionSpec(
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
        )


class EndpointEncryptionSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EndpointEncryptionSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EndpointEncryptionSpec.from_proto(i) for i in resources]


class EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return endpoint_pb2.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum.Value(
            "VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return endpoint_pb2.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum.Name(
            resource
        )[
            len(
                "VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum"
            ) :
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
