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
from google3.cloud.graphite.mmv2.services.google.compute import target_vpn_gateway_pb2
from google3.cloud.graphite.mmv2.services.google.compute import (
    target_vpn_gateway_pb2_grpc,
)

from typing import List


class TargetVpnGateway(object):
    def __init__(
        self,
        id: int = None,
        name: str = None,
        description: str = None,
        region: str = None,
        network: str = None,
        tunnel: list = None,
        status: str = None,
        self_link: str = None,
        forwarding_rule: list = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.region = region
        self.network = network
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = target_vpn_gateway_pb2_grpc.ComputeTargetVpnGatewayServiceStub(
            channel.Channel()
        )
        request = target_vpn_gateway_pb2.ApplyComputeTargetVpnGatewayRequest()
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

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeTargetVpnGateway(request)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.region = Primitive.from_proto(response.region)
        self.network = Primitive.from_proto(response.network)
        self.tunnel = Primitive.from_proto(response.tunnel)
        self.status = TargetVpnGatewayStatusEnum.from_proto(response.status)
        self.self_link = Primitive.from_proto(response.self_link)
        self.forwarding_rule = Primitive.from_proto(response.forwarding_rule)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = target_vpn_gateway_pb2_grpc.ComputeTargetVpnGatewayServiceStub(
            channel.Channel()
        )
        request = target_vpn_gateway_pb2.DeleteComputeTargetVpnGatewayRequest()
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

        response = stub.DeleteComputeTargetVpnGateway(request)

    @classmethod
    def list(self, project, region, service_account_file=""):
        stub = target_vpn_gateway_pb2_grpc.ComputeTargetVpnGatewayServiceStub(
            channel.Channel()
        )
        request = target_vpn_gateway_pb2.ListComputeTargetVpnGatewayRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        return stub.ListComputeTargetVpnGateway(request).items

    def to_proto(self):
        resource = target_vpn_gateway_pb2.ComputeTargetVpnGateway()
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
        return resource


class TargetVpnGatewayStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return target_vpn_gateway_pb2.ComputeTargetVpnGatewayStatusEnum.Value(
            "ComputeTargetVpnGatewayStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return target_vpn_gateway_pb2.ComputeTargetVpnGatewayStatusEnum.Name(resource)[
            len("ComputeTargetVpnGatewayStatusEnum") :
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
