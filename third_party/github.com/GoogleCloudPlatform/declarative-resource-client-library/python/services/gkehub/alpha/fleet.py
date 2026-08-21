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
from google3.cloud.graphite.mmv2.services.google.gke_hub import fleet_pb2
from google3.cloud.graphite.mmv2.services.google.gke_hub import fleet_pb2_grpc

from typing import List


class Fleet(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        create_time: str = None,
        update_time: str = None,
        uid: str = None,
        managed_namespaces: bool = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.managed_namespaces = managed_namespaces
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = fleet_pb2_grpc.GkehubAlphaFleetServiceStub(channel.Channel())
        request = fleet_pb2.ApplyGkehubAlphaFleetRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.managed_namespaces):
            request.resource.managed_namespaces = Primitive.to_proto(
                self.managed_namespaces
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyGkehubAlphaFleet(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.uid = Primitive.from_proto(response.uid)
        self.managed_namespaces = Primitive.from_proto(response.managed_namespaces)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = fleet_pb2_grpc.GkehubAlphaFleetServiceStub(channel.Channel())
        request = fleet_pb2.DeleteGkehubAlphaFleetRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.managed_namespaces):
            request.resource.managed_namespaces = Primitive.to_proto(
                self.managed_namespaces
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteGkehubAlphaFleet(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = fleet_pb2_grpc.GkehubAlphaFleetServiceStub(channel.Channel())
        request = fleet_pb2.ListGkehubAlphaFleetRequest()
        request.service_account_file = service_account_file
        return stub.ListGkehubAlphaFleet(request).items

    def to_proto(self):
        resource = fleet_pb2.GkehubAlphaFleet()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.managed_namespaces):
            resource.managed_namespaces = Primitive.to_proto(self.managed_namespaces)
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
