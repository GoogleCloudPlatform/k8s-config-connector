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
from google3.cloud.graphite.mmv2.services.google.network_security import (
    address_group_pb2,
)
from google3.cloud.graphite.mmv2.services.google.network_security import (
    address_group_pb2_grpc,
)

from typing import List


class AddressGroup(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        type: str = None,
        items: list = None,
        capacity: int = None,
        parent: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.type = type
        self.items = items
        self.capacity = capacity
        self.parent = parent
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = address_group_pb2_grpc.NetworksecurityBetaAddressGroupServiceStub(
            channel.Channel()
        )
        request = address_group_pb2.ApplyNetworksecurityBetaAddressGroupRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if AddressGroupTypeEnum.to_proto(self.type):
            request.resource.type = AddressGroupTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.items):
            request.resource.items.extend(Primitive.to_proto(self.items))
        if Primitive.to_proto(self.capacity):
            request.resource.capacity = Primitive.to_proto(self.capacity)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworksecurityBetaAddressGroup(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.type = AddressGroupTypeEnum.from_proto(response.type)
        self.items = Primitive.from_proto(response.items)
        self.capacity = Primitive.from_proto(response.capacity)
        self.parent = Primitive.from_proto(response.parent)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = address_group_pb2_grpc.NetworksecurityBetaAddressGroupServiceStub(
            channel.Channel()
        )
        request = address_group_pb2.DeleteNetworksecurityBetaAddressGroupRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if AddressGroupTypeEnum.to_proto(self.type):
            request.resource.type = AddressGroupTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.items):
            request.resource.items.extend(Primitive.to_proto(self.items))
        if Primitive.to_proto(self.capacity):
            request.resource.capacity = Primitive.to_proto(self.capacity)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworksecurityBetaAddressGroup(request)

    @classmethod
    def list(self, location, parent, service_account_file=""):
        stub = address_group_pb2_grpc.NetworksecurityBetaAddressGroupServiceStub(
            channel.Channel()
        )
        request = address_group_pb2.ListNetworksecurityBetaAddressGroupRequest()
        request.service_account_file = service_account_file
        request.Location = location

        request.Parent = parent

        return stub.ListNetworksecurityBetaAddressGroup(request).items

    def to_proto(self):
        resource = address_group_pb2.NetworksecurityBetaAddressGroup()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if AddressGroupTypeEnum.to_proto(self.type):
            resource.type = AddressGroupTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.items):
            resource.items.extend(Primitive.to_proto(self.items))
        if Primitive.to_proto(self.capacity):
            resource.capacity = Primitive.to_proto(self.capacity)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class AddressGroupTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return address_group_pb2.NetworksecurityBetaAddressGroupTypeEnum.Value(
            "NetworksecurityBetaAddressGroupTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return address_group_pb2.NetworksecurityBetaAddressGroupTypeEnum.Name(resource)[
            len("NetworksecurityBetaAddressGroupTypeEnum") :
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
