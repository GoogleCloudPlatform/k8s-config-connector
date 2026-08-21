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
from google3.cloud.graphite.mmv2.services.google.network_services import mesh_pb2
from google3.cloud.graphite.mmv2.services.google.network_services import mesh_pb2_grpc

from typing import List


class Mesh(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        description: str = None,
        interception_port: int = None,
        project: str = None,
        location: str = None,
        self_link: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.labels = labels
        self.description = description
        self.interception_port = interception_port
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = mesh_pb2_grpc.NetworkservicesMeshServiceStub(channel.Channel())
        request = mesh_pb2.ApplyNetworkservicesMeshRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.interception_port):
            request.resource.interception_port = Primitive.to_proto(
                self.interception_port
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworkservicesMesh(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.description = Primitive.from_proto(response.description)
        self.interception_port = Primitive.from_proto(response.interception_port)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.self_link = Primitive.from_proto(response.self_link)

    def delete(self):
        stub = mesh_pb2_grpc.NetworkservicesMeshServiceStub(channel.Channel())
        request = mesh_pb2.DeleteNetworkservicesMeshRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.interception_port):
            request.resource.interception_port = Primitive.to_proto(
                self.interception_port
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworkservicesMesh(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = mesh_pb2_grpc.NetworkservicesMeshServiceStub(channel.Channel())
        request = mesh_pb2.ListNetworkservicesMeshRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworkservicesMesh(request).items

    def to_proto(self):
        resource = mesh_pb2.NetworkservicesMesh()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.interception_port):
            resource.interception_port = Primitive.to_proto(self.interception_port)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
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
