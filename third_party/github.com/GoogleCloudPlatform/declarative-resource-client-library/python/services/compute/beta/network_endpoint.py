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
from google3.cloud.graphite.mmv2.services.google.compute import network_endpoint_pb2
from google3.cloud.graphite.mmv2.services.google.compute import (
    network_endpoint_pb2_grpc,
)

from typing import List


class NetworkEndpoint(object):
    def __init__(
        self,
        port: int = None,
        ip_address: str = None,
        fqdn: str = None,
        instance: str = None,
        annotations: dict = None,
        project: str = None,
        location: str = None,
        group: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.port = port
        self.ip_address = ip_address
        self.fqdn = fqdn
        self.instance = instance
        self.annotations = annotations
        self.project = project
        self.location = location
        self.group = group
        self.service_account_file = service_account_file

    def apply(self):
        stub = network_endpoint_pb2_grpc.ComputeBetaNetworkEndpointServiceStub(
            channel.Channel()
        )
        request = network_endpoint_pb2.ApplyComputeBetaNetworkEndpointRequest()
        if Primitive.to_proto(self.port):
            request.resource.port = Primitive.to_proto(self.port)

        if Primitive.to_proto(self.ip_address):
            request.resource.ip_address = Primitive.to_proto(self.ip_address)

        if Primitive.to_proto(self.fqdn):
            request.resource.fqdn = Primitive.to_proto(self.fqdn)

        if Primitive.to_proto(self.instance):
            request.resource.instance = Primitive.to_proto(self.instance)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.group):
            request.resource.group = Primitive.to_proto(self.group)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaNetworkEndpoint(request)
        self.port = Primitive.from_proto(response.port)
        self.ip_address = Primitive.from_proto(response.ip_address)
        self.fqdn = Primitive.from_proto(response.fqdn)
        self.instance = Primitive.from_proto(response.instance)
        self.annotations = Primitive.from_proto(response.annotations)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.group = Primitive.from_proto(response.group)

    def delete(self):
        stub = network_endpoint_pb2_grpc.ComputeBetaNetworkEndpointServiceStub(
            channel.Channel()
        )
        request = network_endpoint_pb2.DeleteComputeBetaNetworkEndpointRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.port):
            request.resource.port = Primitive.to_proto(self.port)

        if Primitive.to_proto(self.ip_address):
            request.resource.ip_address = Primitive.to_proto(self.ip_address)

        if Primitive.to_proto(self.fqdn):
            request.resource.fqdn = Primitive.to_proto(self.fqdn)

        if Primitive.to_proto(self.instance):
            request.resource.instance = Primitive.to_proto(self.instance)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.group):
            request.resource.group = Primitive.to_proto(self.group)

        response = stub.DeleteComputeBetaNetworkEndpoint(request)

    @classmethod
    def list(self, project, location, group, service_account_file=""):
        stub = network_endpoint_pb2_grpc.ComputeBetaNetworkEndpointServiceStub(
            channel.Channel()
        )
        request = network_endpoint_pb2.ListComputeBetaNetworkEndpointRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Group = group

        return stub.ListComputeBetaNetworkEndpoint(request).items

    def to_proto(self):
        resource = network_endpoint_pb2.ComputeBetaNetworkEndpoint()
        if Primitive.to_proto(self.port):
            resource.port = Primitive.to_proto(self.port)
        if Primitive.to_proto(self.ip_address):
            resource.ip_address = Primitive.to_proto(self.ip_address)
        if Primitive.to_proto(self.fqdn):
            resource.fqdn = Primitive.to_proto(self.fqdn)
        if Primitive.to_proto(self.instance):
            resource.instance = Primitive.to_proto(self.instance)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.group):
            resource.group = Primitive.to_proto(self.group)
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
