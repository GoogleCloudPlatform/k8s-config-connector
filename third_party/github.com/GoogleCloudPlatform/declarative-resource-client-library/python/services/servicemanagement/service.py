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
from google3.cloud.graphite.mmv2.services.google.servicemanagement import service_pb2
from google3.cloud.graphite.mmv2.services.google.servicemanagement import (
    service_pb2_grpc,
)

from typing import List


class Service(object):
    def __init__(
        self, name: str = None, project: str = None, service_account_file: str = ""
    ):

        channel.initialize()
        self.name = name
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = service_pb2_grpc.ServicemanagementServiceServiceStub(channel.Channel())
        request = service_pb2.ApplyServicemanagementServiceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyServicemanagementService(request)
        self.name = Primitive.from_proto(response.name)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = service_pb2_grpc.ServicemanagementServiceServiceStub(channel.Channel())
        request = service_pb2.DeleteServicemanagementServiceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteServicemanagementService(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = service_pb2_grpc.ServicemanagementServiceServiceStub(channel.Channel())
        request = service_pb2.ListServicemanagementServiceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListServicemanagementService(request).items

    def to_proto(self):
        resource = service_pb2.ServicemanagementService()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
