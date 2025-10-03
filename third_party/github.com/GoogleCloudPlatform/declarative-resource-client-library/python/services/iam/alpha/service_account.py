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
from google3.cloud.graphite.mmv2.services.google.iam import service_account_pb2
from google3.cloud.graphite.mmv2.services.google.iam import service_account_pb2_grpc

from typing import List


class ServiceAccount(object):
    def __init__(
        self,
        name: str = None,
        project: str = None,
        unique_id: str = None,
        email: str = None,
        display_name: str = None,
        description: str = None,
        oauth2_client_id: str = None,
        actas_resources: dict = None,
        disabled: bool = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.project = project
        self.display_name = display_name
        self.description = description
        self.actas_resources = actas_resources
        self.service_account_file = service_account_file

    def apply(self):
        stub = service_account_pb2_grpc.IamAlphaServiceAccountServiceStub(
            channel.Channel()
        )
        request = service_account_pb2.ApplyIamAlphaServiceAccountRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ServiceAccountActasResources.to_proto(self.actas_resources):
            request.resource.actas_resources.CopyFrom(
                ServiceAccountActasResources.to_proto(self.actas_resources)
            )
        else:
            request.resource.ClearField("actas_resources")
        request.service_account_file = self.service_account_file

        response = stub.ApplyIamAlphaServiceAccount(request)
        self.name = Primitive.from_proto(response.name)
        self.project = Primitive.from_proto(response.project)
        self.unique_id = Primitive.from_proto(response.unique_id)
        self.email = Primitive.from_proto(response.email)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.oauth2_client_id = Primitive.from_proto(response.oauth2_client_id)
        self.actas_resources = ServiceAccountActasResources.from_proto(
            response.actas_resources
        )
        self.disabled = Primitive.from_proto(response.disabled)

    def delete(self):
        stub = service_account_pb2_grpc.IamAlphaServiceAccountServiceStub(
            channel.Channel()
        )
        request = service_account_pb2.DeleteIamAlphaServiceAccountRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ServiceAccountActasResources.to_proto(self.actas_resources):
            request.resource.actas_resources.CopyFrom(
                ServiceAccountActasResources.to_proto(self.actas_resources)
            )
        else:
            request.resource.ClearField("actas_resources")
        response = stub.DeleteIamAlphaServiceAccount(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = service_account_pb2_grpc.IamAlphaServiceAccountServiceStub(
            channel.Channel()
        )
        request = service_account_pb2.ListIamAlphaServiceAccountRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListIamAlphaServiceAccount(request).items

    def to_proto(self):
        resource = service_account_pb2.IamAlphaServiceAccount()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if ServiceAccountActasResources.to_proto(self.actas_resources):
            resource.actas_resources.CopyFrom(
                ServiceAccountActasResources.to_proto(self.actas_resources)
            )
        else:
            resource.ClearField("actas_resources")
        return resource


class ServiceAccountActasResources(object):
    def __init__(self, resources: list = None):
        self.resources = resources

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_account_pb2.IamAlphaServiceAccountActasResources()
        if ServiceAccountActasResourcesResourcesArray.to_proto(resource.resources):
            res.resources.extend(
                ServiceAccountActasResourcesResourcesArray.to_proto(resource.resources)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceAccountActasResources(
            resources=ServiceAccountActasResourcesResourcesArray.from_proto(
                resource.resources
            ),
        )


class ServiceAccountActasResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceAccountActasResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceAccountActasResources.from_proto(i) for i in resources]


class ServiceAccountActasResourcesResources(object):
    def __init__(self, full_resource_name: str = None):
        self.full_resource_name = full_resource_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_account_pb2.IamAlphaServiceAccountActasResourcesResources()
        if Primitive.to_proto(resource.full_resource_name):
            res.full_resource_name = Primitive.to_proto(resource.full_resource_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceAccountActasResourcesResources(
            full_resource_name=Primitive.from_proto(resource.full_resource_name),
        )


class ServiceAccountActasResourcesResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceAccountActasResourcesResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceAccountActasResourcesResources.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
