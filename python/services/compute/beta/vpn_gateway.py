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
from google3.cloud.graphite.mmv2.services.google.compute import vpn_gateway_pb2
from google3.cloud.graphite.mmv2.services.google.compute import vpn_gateway_pb2_grpc

from typing import List


class VpnGateway(object):
    def __init__(
        self,
        id: int = None,
        name: str = None,
        description: str = None,
        region: str = None,
        network: str = None,
        self_link: str = None,
        project: str = None,
        labels: dict = None,
        vpn_interface: list = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.region = region
        self.network = network
        self.project = project
        self.labels = labels
        self.service_account_file = service_account_file

    def apply(self):
        stub = vpn_gateway_pb2_grpc.ComputeBetaVpnGatewayServiceStub(channel.Channel())
        request = vpn_gateway_pb2.ApplyComputeBetaVpnGatewayRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaVpnGateway(request)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.region = Primitive.from_proto(response.region)
        self.network = Primitive.from_proto(response.network)
        self.self_link = Primitive.from_proto(response.self_link)
        self.project = Primitive.from_proto(response.project)
        self.labels = Primitive.from_proto(response.labels)
        self.vpn_interface = VpnGatewayVpnInterfaceArray.from_proto(
            response.vpn_interface
        )

    def delete(self):
        stub = vpn_gateway_pb2_grpc.ComputeBetaVpnGatewayServiceStub(channel.Channel())
        request = vpn_gateway_pb2.DeleteComputeBetaVpnGatewayRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        response = stub.DeleteComputeBetaVpnGateway(request)

    @classmethod
    def list(self, project, region, service_account_file=""):
        stub = vpn_gateway_pb2_grpc.ComputeBetaVpnGatewayServiceStub(channel.Channel())
        request = vpn_gateway_pb2.ListComputeBetaVpnGatewayRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        return stub.ListComputeBetaVpnGateway(request).items

    def to_proto(self):
        resource = vpn_gateway_pb2.ComputeBetaVpnGateway()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        return resource


class VpnGatewayVpnInterface(object):
    def __init__(self, id: int = None, ip_address: str = None):
        self.id = id
        self.ip_address = ip_address

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = vpn_gateway_pb2.ComputeBetaVpnGatewayVpnInterface()
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.ip_address):
            res.ip_address = Primitive.to_proto(resource.ip_address)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VpnGatewayVpnInterface(
            id=Primitive.from_proto(resource.id),
            ip_address=Primitive.from_proto(resource.ip_address),
        )


class VpnGatewayVpnInterfaceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VpnGatewayVpnInterface.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VpnGatewayVpnInterface.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
