# Copyright 2023 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.vmware import vmware_engine_network_pb2
from google3.cloud.graphite.mmv2.services.google.vmware import (
    vmware_engine_network_pb2_grpc,
)

from typing import List


class VmwareEngineNetwork(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        description: str = None,
        vpc_networks: list = None,
        state: str = None,
        type: str = None,
        uid: str = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.type = type
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = vmware_engine_network_pb2_grpc.VmwareAlphaVmwareEngineNetworkServiceStub(
            channel.Channel()
        )
        request = vmware_engine_network_pb2.ApplyVmwareAlphaVmwareEngineNetworkRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if VmwareEngineNetworkTypeEnum.to_proto(self.type):
            request.resource.type = VmwareEngineNetworkTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyVmwareAlphaVmwareEngineNetwork(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.description = Primitive.from_proto(response.description)
        self.vpc_networks = VmwareEngineNetworkVPCNetworksArray.from_proto(
            response.vpc_networks
        )
        self.state = VmwareEngineNetworkStateEnum.from_proto(response.state)
        self.type = VmwareEngineNetworkTypeEnum.from_proto(response.type)
        self.uid = Primitive.from_proto(response.uid)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = vmware_engine_network_pb2_grpc.VmwareAlphaVmwareEngineNetworkServiceStub(
            channel.Channel()
        )
        request = (
            vmware_engine_network_pb2.DeleteVmwareAlphaVmwareEngineNetworkRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if VmwareEngineNetworkTypeEnum.to_proto(self.type):
            request.resource.type = VmwareEngineNetworkTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteVmwareAlphaVmwareEngineNetwork(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = vmware_engine_network_pb2_grpc.VmwareAlphaVmwareEngineNetworkServiceStub(
            channel.Channel()
        )
        request = vmware_engine_network_pb2.ListVmwareAlphaVmwareEngineNetworkRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListVmwareAlphaVmwareEngineNetwork(request).items

    def to_proto(self):
        resource = vmware_engine_network_pb2.VmwareAlphaVmwareEngineNetwork()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if VmwareEngineNetworkTypeEnum.to_proto(self.type):
            resource.type = VmwareEngineNetworkTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class VmwareEngineNetworkVPCNetworks(object):
    def __init__(self, type: str = None, network: str = None):
        self.type = type
        self.network = network

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = vmware_engine_network_pb2.VmwareAlphaVmwareEngineNetworkVPCNetworks()
        if VmwareEngineNetworkVPCNetworksTypeEnum.to_proto(resource.type):
            res.type = VmwareEngineNetworkVPCNetworksTypeEnum.to_proto(resource.type)
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VmwareEngineNetworkVPCNetworks(
            type=VmwareEngineNetworkVPCNetworksTypeEnum.from_proto(resource.type),
            network=Primitive.from_proto(resource.network),
        )


class VmwareEngineNetworkVPCNetworksArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VmwareEngineNetworkVPCNetworks.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VmwareEngineNetworkVPCNetworks.from_proto(i) for i in resources]


class VmwareEngineNetworkVPCNetworksTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return vmware_engine_network_pb2.VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum.Value(
            "VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return vmware_engine_network_pb2.VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum.Name(
            resource
        )[
            len("VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum") :
        ]


class VmwareEngineNetworkStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return vmware_engine_network_pb2.VmwareAlphaVmwareEngineNetworkStateEnum.Value(
            "VmwareAlphaVmwareEngineNetworkStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return vmware_engine_network_pb2.VmwareAlphaVmwareEngineNetworkStateEnum.Name(
            resource
        )[len("VmwareAlphaVmwareEngineNetworkStateEnum") :]


class VmwareEngineNetworkTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return vmware_engine_network_pb2.VmwareAlphaVmwareEngineNetworkTypeEnum.Value(
            "VmwareAlphaVmwareEngineNetworkTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return vmware_engine_network_pb2.VmwareAlphaVmwareEngineNetworkTypeEnum.Name(
            resource
        )[len("VmwareAlphaVmwareEngineNetworkTypeEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
