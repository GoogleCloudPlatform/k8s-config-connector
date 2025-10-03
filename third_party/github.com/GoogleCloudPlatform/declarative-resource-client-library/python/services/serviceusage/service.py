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
from google3.cloud.graphite.mmv2.services.google.service_usage import service_pb2
from google3.cloud.graphite.mmv2.services.google.service_usage import service_pb2_grpc

from typing import List


class Service(object):
    def __init__(
        self,
        name: str = None,
        state: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.state = state
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = service_pb2_grpc.ServiceusageServiceServiceStub(channel.Channel())
        request = service_pb2.ApplyServiceusageServiceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ServiceStateEnum.to_proto(self.state):
            request.resource.state = ServiceStateEnum.to_proto(self.state)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyServiceusageService(request)
        self.name = Primitive.from_proto(response.name)
        self.state = ServiceStateEnum.from_proto(response.state)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = service_pb2_grpc.ServiceusageServiceServiceStub(channel.Channel())
        request = service_pb2.DeleteServiceusageServiceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ServiceStateEnum.to_proto(self.state):
            request.resource.state = ServiceStateEnum.to_proto(self.state)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteServiceusageService(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = service_pb2_grpc.ServiceusageServiceServiceStub(channel.Channel())
        request = service_pb2.ListServiceusageServiceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListServiceusageService(request).items

    def to_proto(self):
        resource = service_pb2.ServiceusageService()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if ServiceStateEnum.to_proto(self.state):
            resource.state = ServiceStateEnum.to_proto(self.state)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class ServiceStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.ServiceusageServiceStateEnum.Value(
            "ServiceusageServiceStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.ServiceusageServiceStateEnum.Name(resource)[
            len("ServiceusageServiceStateEnum") :
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
