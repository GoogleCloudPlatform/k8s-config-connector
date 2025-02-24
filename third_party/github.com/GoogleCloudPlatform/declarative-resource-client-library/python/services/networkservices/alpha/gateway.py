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
from google3.cloud.graphite.mmv2.services.google.network_services import gateway_pb2
from google3.cloud.graphite.mmv2.services.google.network_services import (
    gateway_pb2_grpc,
)

from typing import List


class Gateway(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        description: str = None,
        type: str = None,
        addresses: list = None,
        ports: list = None,
        scope: str = None,
        server_tls_policy: str = None,
        project: str = None,
        location: str = None,
        self_link: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.labels = labels
        self.description = description
        self.type = type
        self.addresses = addresses
        self.ports = ports
        self.scope = scope
        self.server_tls_policy = server_tls_policy
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = gateway_pb2_grpc.NetworkservicesAlphaGatewayServiceStub(
            channel.Channel()
        )
        request = gateway_pb2.ApplyNetworkservicesAlphaGatewayRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if GatewayTypeEnum.to_proto(self.type):
            request.resource.type = GatewayTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.addresses):
            request.resource.addresses.extend(Primitive.to_proto(self.addresses))
        if int64Array.to_proto(self.ports):
            request.resource.ports.extend(int64Array.to_proto(self.ports))
        if Primitive.to_proto(self.scope):
            request.resource.scope = Primitive.to_proto(self.scope)

        if Primitive.to_proto(self.server_tls_policy):
            request.resource.server_tls_policy = Primitive.to_proto(
                self.server_tls_policy
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworkservicesAlphaGateway(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.description = Primitive.from_proto(response.description)
        self.type = GatewayTypeEnum.from_proto(response.type)
        self.addresses = Primitive.from_proto(response.addresses)
        self.ports = int64Array.from_proto(response.ports)
        self.scope = Primitive.from_proto(response.scope)
        self.server_tls_policy = Primitive.from_proto(response.server_tls_policy)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.self_link = Primitive.from_proto(response.self_link)

    def delete(self):
        stub = gateway_pb2_grpc.NetworkservicesAlphaGatewayServiceStub(
            channel.Channel()
        )
        request = gateway_pb2.DeleteNetworkservicesAlphaGatewayRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if GatewayTypeEnum.to_proto(self.type):
            request.resource.type = GatewayTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.addresses):
            request.resource.addresses.extend(Primitive.to_proto(self.addresses))
        if int64Array.to_proto(self.ports):
            request.resource.ports.extend(int64Array.to_proto(self.ports))
        if Primitive.to_proto(self.scope):
            request.resource.scope = Primitive.to_proto(self.scope)

        if Primitive.to_proto(self.server_tls_policy):
            request.resource.server_tls_policy = Primitive.to_proto(
                self.server_tls_policy
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworkservicesAlphaGateway(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = gateway_pb2_grpc.NetworkservicesAlphaGatewayServiceStub(
            channel.Channel()
        )
        request = gateway_pb2.ListNetworkservicesAlphaGatewayRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworkservicesAlphaGateway(request).items

    def to_proto(self):
        resource = gateway_pb2.NetworkservicesAlphaGateway()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if GatewayTypeEnum.to_proto(self.type):
            resource.type = GatewayTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.addresses):
            resource.addresses.extend(Primitive.to_proto(self.addresses))
        if int64Array.to_proto(self.ports):
            resource.ports.extend(int64Array.to_proto(self.ports))
        if Primitive.to_proto(self.scope):
            resource.scope = Primitive.to_proto(self.scope)
        if Primitive.to_proto(self.server_tls_policy):
            resource.server_tls_policy = Primitive.to_proto(self.server_tls_policy)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class GatewayTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return gateway_pb2.NetworkservicesAlphaGatewayTypeEnum.Value(
            "NetworkservicesAlphaGatewayTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return gateway_pb2.NetworkservicesAlphaGatewayTypeEnum.Name(resource)[
            len("NetworkservicesAlphaGatewayTypeEnum") :
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
