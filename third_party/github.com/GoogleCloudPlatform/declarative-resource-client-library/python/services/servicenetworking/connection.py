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
from google3.cloud.graphite.mmv2.services.google.service_networking import (
    connection_pb2,
)
from google3.cloud.graphite.mmv2.services.google.service_networking import (
    connection_pb2_grpc,
)

from typing import List


class Connection(object):
    def __init__(
        self,
        network: str = None,
        project: str = None,
        name: str = None,
        reserved_peering_ranges: list = None,
        service: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.network = network
        self.project = project
        self.reserved_peering_ranges = reserved_peering_ranges
        self.service = service
        self.service_account_file = service_account_file

    def apply(self):
        stub = connection_pb2_grpc.ServicenetworkingConnectionServiceStub(
            channel.Channel()
        )
        request = connection_pb2.ApplyServicenetworkingConnectionRequest()
        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.reserved_peering_ranges):
            request.resource.reserved_peering_ranges.extend(
                Primitive.to_proto(self.reserved_peering_ranges)
            )
        if Primitive.to_proto(self.service):
            request.resource.service = Primitive.to_proto(self.service)

        request.service_account_file = self.service_account_file

        response = stub.ApplyServicenetworkingConnection(request)
        self.network = Primitive.from_proto(response.network)
        self.project = Primitive.from_proto(response.project)
        self.name = Primitive.from_proto(response.name)
        self.reserved_peering_ranges = Primitive.from_proto(
            response.reserved_peering_ranges
        )
        self.service = Primitive.from_proto(response.service)

    def delete(self):
        stub = connection_pb2_grpc.ServicenetworkingConnectionServiceStub(
            channel.Channel()
        )
        request = connection_pb2.DeleteServicenetworkingConnectionRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.reserved_peering_ranges):
            request.resource.reserved_peering_ranges.extend(
                Primitive.to_proto(self.reserved_peering_ranges)
            )
        if Primitive.to_proto(self.service):
            request.resource.service = Primitive.to_proto(self.service)

        response = stub.DeleteServicenetworkingConnection(request)

    @classmethod
    def list(self, project, network, service, service_account_file=""):
        stub = connection_pb2_grpc.ServicenetworkingConnectionServiceStub(
            channel.Channel()
        )
        request = connection_pb2.ListServicenetworkingConnectionRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Network = network

        request.Service = service

        return stub.ListServicenetworkingConnection(request).items

    def to_proto(self):
        resource = connection_pb2.ServicenetworkingConnection()
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.reserved_peering_ranges):
            resource.reserved_peering_ranges.extend(
                Primitive.to_proto(self.reserved_peering_ranges)
            )
        if Primitive.to_proto(self.service):
            resource.service = Primitive.to_proto(self.service)
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
