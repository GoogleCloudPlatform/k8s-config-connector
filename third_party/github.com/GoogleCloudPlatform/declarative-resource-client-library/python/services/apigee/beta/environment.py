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
from google3.cloud.graphite.mmv2.services.google.apigee import environment_pb2
from google3.cloud.graphite.mmv2.services.google.apigee import environment_pb2_grpc

from typing import List


class Environment(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        created_at: int = None,
        last_modified_at: int = None,
        properties: dict = None,
        display_name: str = None,
        state: str = None,
        apigee_organization: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.properties = properties
        self.display_name = display_name
        self.apigee_organization = apigee_organization
        self.service_account_file = service_account_file

    def apply(self):
        stub = environment_pb2_grpc.ApigeeBetaEnvironmentServiceStub(channel.Channel())
        request = environment_pb2.ApplyApigeeBetaEnvironmentRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.properties):
            request.resource.properties = Primitive.to_proto(self.properties)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.apigee_organization):
            request.resource.apigee_organization = Primitive.to_proto(
                self.apigee_organization
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyApigeeBetaEnvironment(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.created_at = Primitive.from_proto(response.created_at)
        self.last_modified_at = Primitive.from_proto(response.last_modified_at)
        self.properties = Primitive.from_proto(response.properties)
        self.display_name = Primitive.from_proto(response.display_name)
        self.state = EnvironmentStateEnum.from_proto(response.state)
        self.apigee_organization = Primitive.from_proto(response.apigee_organization)

    def delete(self):
        stub = environment_pb2_grpc.ApigeeBetaEnvironmentServiceStub(channel.Channel())
        request = environment_pb2.DeleteApigeeBetaEnvironmentRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.properties):
            request.resource.properties = Primitive.to_proto(self.properties)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.apigee_organization):
            request.resource.apigee_organization = Primitive.to_proto(
                self.apigee_organization
            )

        response = stub.DeleteApigeeBetaEnvironment(request)

    @classmethod
    def list(self, apigeeOrganization, service_account_file=""):
        stub = environment_pb2_grpc.ApigeeBetaEnvironmentServiceStub(channel.Channel())
        request = environment_pb2.ListApigeeBetaEnvironmentRequest()
        request.service_account_file = service_account_file
        request.ApigeeOrganization = apigeeOrganization

        return stub.ListApigeeBetaEnvironment(request).items

    def to_proto(self):
        resource = environment_pb2.ApigeeBetaEnvironment()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.properties):
            resource.properties = Primitive.to_proto(self.properties)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.apigee_organization):
            resource.apigee_organization = Primitive.to_proto(self.apigee_organization)
        return resource


class EnvironmentStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return environment_pb2.ApigeeBetaEnvironmentStateEnum.Value(
            "ApigeeBetaEnvironmentStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return environment_pb2.ApigeeBetaEnvironmentStateEnum.Name(resource)[
            len("ApigeeBetaEnvironmentStateEnum") :
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
