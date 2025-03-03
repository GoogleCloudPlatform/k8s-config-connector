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
from google3.cloud.graphite.mmv2.services.google.serviceusage import (
    serviceusage_service_pb2,
)
from google3.cloud.graphite.mmv2.services.google.serviceusage import (
    serviceusage_service_pb2_grpc,
)

from typing import List


class ServiceusageService(object):
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
        self.service_account_file = service_account_file

    def apply(self):
        stub = serviceusage_service_pb2_grpc.ServiceusageServiceusageServiceServiceStub(
            channel.Channel()
        )
        request = serviceusage_service_pb2.ApplyServiceusageServiceusageServiceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ServiceusageServiceStateEnum.to_proto(self.state):
            request.resource.state = ServiceusageServiceStateEnum.to_proto(self.state)

        request.service_account_file = self.service_account_file

        response = stub.ApplyServiceusageServiceusageService(request)
        self.name = Primitive.from_proto(response.name)
        self.state = ServiceusageServiceStateEnum.from_proto(response.state)
        self.project = Primitive.from_proto(response.project)

    @classmethod
    def delete(self, project, name, service_account_file=""):
        stub = serviceusage_service_pb2_grpc.ServiceusageServiceusageServiceServiceStub(
            channel.Channel()
        )
        request = (
            serviceusage_service_pb2.DeleteServiceusageServiceusageServiceRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Name = name

        response = stub.DeleteServiceusageServiceusageService(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = serviceusage_service_pb2_grpc.ServiceusageServiceusageServiceServiceStub(
            channel.Channel()
        )
        request = serviceusage_service_pb2.ListServiceusageServiceusageServiceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListServiceusageServiceusageService(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = serviceusage_service_pb2.ServiceusageServiceusageService()
        any_proto.Unpack(res_proto)

        res = ServiceusageService()
        res.name = Primitive.from_proto(res_proto.name)
        res.state = ServiceusageServiceStateEnum.from_proto(res_proto.state)
        res.project = Primitive.from_proto(res_proto.project)
        return res


class ServiceusageServiceStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return serviceusage_service_pb2.ServiceusageServiceusageServiceStateEnum.Value(
            "ServiceusageServiceusageServiceStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return serviceusage_service_pb2.ServiceusageServiceusageServiceStateEnum.Name(
            resource
        )[len("ServiceusageServiceusageServiceStateEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
