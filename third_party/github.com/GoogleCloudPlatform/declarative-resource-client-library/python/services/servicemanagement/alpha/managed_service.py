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
from google3.cloud.graphite.mmv2.services.google.servicemanagement import (
    managed_service_pb2,
)
from google3.cloud.graphite.mmv2.services.google.servicemanagement import (
    managed_service_pb2_grpc,
)

from typing import List


class ManagedService(object):
    def __init__(
        self, name: str = None, project: str = None, service_account_file: str = ""
    ):

        channel.initialize()
        self.name = name
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = managed_service_pb2_grpc.ServicemanagementAlphaManagedServiceServiceStub(
            channel.Channel()
        )
        request = managed_service_pb2.ApplyServicemanagementAlphaManagedServiceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyServicemanagementAlphaManagedService(request)
        self.name = Primitive.from_proto(response.name)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = managed_service_pb2_grpc.ServicemanagementAlphaManagedServiceServiceStub(
            channel.Channel()
        )
        request = (
            managed_service_pb2.DeleteServicemanagementAlphaManagedServiceRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteServicemanagementAlphaManagedService(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = managed_service_pb2_grpc.ServicemanagementAlphaManagedServiceServiceStub(
            channel.Channel()
        )
        request = managed_service_pb2.ListServicemanagementAlphaManagedServiceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListServicemanagementAlphaManagedService(request).items

    def to_proto(self):
        resource = managed_service_pb2.ServicemanagementAlphaManagedService()
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
