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
from google3.cloud.graphite.mmv2.services.google.vertex import model_deployment_pb2
from google3.cloud.graphite.mmv2.services.google.vertex import model_deployment_pb2_grpc

from typing import List


class ModelDeployment(object):
    def __init__(
        self,
        model: str = None,
        id: str = None,
        dedicated_resources: dict = None,
        endpoint: str = None,
        location: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.model = model
        self.dedicated_resources = dedicated_resources
        self.endpoint = endpoint
        self.location = location
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = model_deployment_pb2_grpc.VertexAlphaModelDeploymentServiceStub(
            channel.Channel()
        )
        request = model_deployment_pb2.ApplyVertexAlphaModelDeploymentRequest()
        if Primitive.to_proto(self.model):
            request.resource.model = Primitive.to_proto(self.model)

        if ModelDeploymentDedicatedResources.to_proto(self.dedicated_resources):
            request.resource.dedicated_resources.CopyFrom(
                ModelDeploymentDedicatedResources.to_proto(self.dedicated_resources)
            )
        else:
            request.resource.ClearField("dedicated_resources")
        if Primitive.to_proto(self.endpoint):
            request.resource.endpoint = Primitive.to_proto(self.endpoint)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyVertexAlphaModelDeployment(request)
        self.model = Primitive.from_proto(response.model)
        self.id = Primitive.from_proto(response.id)
        self.dedicated_resources = ModelDeploymentDedicatedResources.from_proto(
            response.dedicated_resources
        )
        self.endpoint = Primitive.from_proto(response.endpoint)
        self.location = Primitive.from_proto(response.location)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = model_deployment_pb2_grpc.VertexAlphaModelDeploymentServiceStub(
            channel.Channel()
        )
        request = model_deployment_pb2.DeleteVertexAlphaModelDeploymentRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.model):
            request.resource.model = Primitive.to_proto(self.model)

        if ModelDeploymentDedicatedResources.to_proto(self.dedicated_resources):
            request.resource.dedicated_resources.CopyFrom(
                ModelDeploymentDedicatedResources.to_proto(self.dedicated_resources)
            )
        else:
            request.resource.ClearField("dedicated_resources")
        if Primitive.to_proto(self.endpoint):
            request.resource.endpoint = Primitive.to_proto(self.endpoint)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteVertexAlphaModelDeployment(request)

    @classmethod
    def list(self, project, location, endpoint, service_account_file=""):
        stub = model_deployment_pb2_grpc.VertexAlphaModelDeploymentServiceStub(
            channel.Channel()
        )
        request = model_deployment_pb2.ListVertexAlphaModelDeploymentRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Endpoint = endpoint

        return stub.ListVertexAlphaModelDeployment(request).items

    def to_proto(self):
        resource = model_deployment_pb2.VertexAlphaModelDeployment()
        if Primitive.to_proto(self.model):
            resource.model = Primitive.to_proto(self.model)
        if ModelDeploymentDedicatedResources.to_proto(self.dedicated_resources):
            resource.dedicated_resources.CopyFrom(
                ModelDeploymentDedicatedResources.to_proto(self.dedicated_resources)
            )
        else:
            resource.ClearField("dedicated_resources")
        if Primitive.to_proto(self.endpoint):
            resource.endpoint = Primitive.to_proto(self.endpoint)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class ModelDeploymentDedicatedResources(object):
    def __init__(
        self,
        machine_spec: dict = None,
        min_replica_count: int = None,
        max_replica_count: int = None,
    ):
        self.machine_spec = machine_spec
        self.min_replica_count = min_replica_count
        self.max_replica_count = max_replica_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = model_deployment_pb2.VertexAlphaModelDeploymentDedicatedResources()
        if ModelDeploymentDedicatedResourcesMachineSpec.to_proto(resource.machine_spec):
            res.machine_spec.CopyFrom(
                ModelDeploymentDedicatedResourcesMachineSpec.to_proto(
                    resource.machine_spec
                )
            )
        else:
            res.ClearField("machine_spec")
        if Primitive.to_proto(resource.min_replica_count):
            res.min_replica_count = Primitive.to_proto(resource.min_replica_count)
        if Primitive.to_proto(resource.max_replica_count):
            res.max_replica_count = Primitive.to_proto(resource.max_replica_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ModelDeploymentDedicatedResources(
            machine_spec=ModelDeploymentDedicatedResourcesMachineSpec.from_proto(
                resource.machine_spec
            ),
            min_replica_count=Primitive.from_proto(resource.min_replica_count),
            max_replica_count=Primitive.from_proto(resource.max_replica_count),
        )


class ModelDeploymentDedicatedResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ModelDeploymentDedicatedResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ModelDeploymentDedicatedResources.from_proto(i) for i in resources]


class ModelDeploymentDedicatedResourcesMachineSpec(object):
    def __init__(self, machine_type: str = None):
        self.machine_type = machine_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            model_deployment_pb2.VertexAlphaModelDeploymentDedicatedResourcesMachineSpec()
        )
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ModelDeploymentDedicatedResourcesMachineSpec(
            machine_type=Primitive.from_proto(resource.machine_type),
        )


class ModelDeploymentDedicatedResourcesMachineSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ModelDeploymentDedicatedResourcesMachineSpec.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ModelDeploymentDedicatedResourcesMachineSpec.from_proto(i)
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
