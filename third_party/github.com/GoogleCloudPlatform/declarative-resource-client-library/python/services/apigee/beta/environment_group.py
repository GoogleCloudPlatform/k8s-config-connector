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
from google3.cloud.graphite.mmv2.services.google.apigee import environment_group_pb2
from google3.cloud.graphite.mmv2.services.google.apigee import (
    environment_group_pb2_grpc,
)

from typing import List


class EnvironmentGroup(object):
    def __init__(
        self,
        name: str = None,
        hostnames: list = None,
        created_at: int = None,
        last_modified_at: int = None,
        state: str = None,
        apigee_organization: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.hostnames = hostnames
        self.apigee_organization = apigee_organization
        self.service_account_file = service_account_file

    def apply(self):
        stub = environment_group_pb2_grpc.ApigeeBetaEnvironmentGroupServiceStub(
            channel.Channel()
        )
        request = environment_group_pb2.ApplyApigeeBetaEnvironmentGroupRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.hostnames):
            request.resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.apigee_organization):
            request.resource.apigee_organization = Primitive.to_proto(
                self.apigee_organization
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyApigeeBetaEnvironmentGroup(request)
        self.name = Primitive.from_proto(response.name)
        self.hostnames = Primitive.from_proto(response.hostnames)
        self.created_at = Primitive.from_proto(response.created_at)
        self.last_modified_at = Primitive.from_proto(response.last_modified_at)
        self.state = EnvironmentGroupStateEnum.from_proto(response.state)
        self.apigee_organization = Primitive.from_proto(response.apigee_organization)

    def delete(self):
        stub = environment_group_pb2_grpc.ApigeeBetaEnvironmentGroupServiceStub(
            channel.Channel()
        )
        request = environment_group_pb2.DeleteApigeeBetaEnvironmentGroupRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.hostnames):
            request.resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.apigee_organization):
            request.resource.apigee_organization = Primitive.to_proto(
                self.apigee_organization
            )

        response = stub.DeleteApigeeBetaEnvironmentGroup(request)

    @classmethod
    def list(self, apigeeOrganization, service_account_file=""):
        stub = environment_group_pb2_grpc.ApigeeBetaEnvironmentGroupServiceStub(
            channel.Channel()
        )
        request = environment_group_pb2.ListApigeeBetaEnvironmentGroupRequest()
        request.service_account_file = service_account_file
        request.ApigeeOrganization = apigeeOrganization

        return stub.ListApigeeBetaEnvironmentGroup(request).items

    def to_proto(self):
        resource = environment_group_pb2.ApigeeBetaEnvironmentGroup()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.hostnames):
            resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.apigee_organization):
            resource.apigee_organization = Primitive.to_proto(self.apigee_organization)
        return resource


class EnvironmentGroupStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return environment_group_pb2.ApigeeBetaEnvironmentGroupStateEnum.Value(
            "ApigeeBetaEnvironmentGroupStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return environment_group_pb2.ApigeeBetaEnvironmentGroupStateEnum.Name(resource)[
            len("ApigeeBetaEnvironmentGroupStateEnum") :
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
