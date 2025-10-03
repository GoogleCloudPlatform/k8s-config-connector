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
from google3.cloud.graphite.mmv2.services.google.compute import router_peer_pb2
from google3.cloud.graphite.mmv2.services.google.compute import router_peer_pb2_grpc

from typing import List


class RouterPeer(object):
    def __init__(
        self,
        creation_timestamp: str = None,
        router: str = None,
        name: str = None,
        interface_name: str = None,
        ip_address: str = None,
        peer_ip_address: str = None,
        peer_asn: int = None,
        advertised_route_priority: int = None,
        advertise_mode: str = None,
        management_type: str = None,
        advertised_groups: list = None,
        advertised_ip_ranges: list = None,
        region: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.router = router
        self.name = name
        self.interface_name = interface_name
        self.ip_address = ip_address
        self.peer_ip_address = peer_ip_address
        self.peer_asn = peer_asn
        self.advertised_route_priority = advertised_route_priority
        self.advertise_mode = advertise_mode
        self.advertised_groups = advertised_groups
        self.advertised_ip_ranges = advertised_ip_ranges
        self.region = region
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = router_peer_pb2_grpc.ComputeRouterPeerServiceStub(channel.Channel())
        request = router_peer_pb2.ApplyComputeRouterPeerRequest()
        if Primitive.to_proto(self.router):
            request.resource.router = Primitive.to_proto(self.router)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.interface_name):
            request.resource.interface_name = Primitive.to_proto(self.interface_name)

        if Primitive.to_proto(self.ip_address):
            request.resource.ip_address = Primitive.to_proto(self.ip_address)

        if Primitive.to_proto(self.peer_ip_address):
            request.resource.peer_ip_address = Primitive.to_proto(self.peer_ip_address)

        if Primitive.to_proto(self.peer_asn):
            request.resource.peer_asn = Primitive.to_proto(self.peer_asn)

        if Primitive.to_proto(self.advertised_route_priority):
            request.resource.advertised_route_priority = Primitive.to_proto(
                self.advertised_route_priority
            )

        if Primitive.to_proto(self.advertise_mode):
            request.resource.advertise_mode = Primitive.to_proto(self.advertise_mode)

        if Primitive.to_proto(self.advertised_groups):
            request.resource.advertised_groups.extend(
                Primitive.to_proto(self.advertised_groups)
            )
        if RouterPeerAdvertisedIPRangesArray.to_proto(self.advertised_ip_ranges):
            request.resource.advertised_ip_ranges.extend(
                RouterPeerAdvertisedIPRangesArray.to_proto(self.advertised_ip_ranges)
            )
        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeRouterPeer(request)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.router = Primitive.from_proto(response.router)
        self.name = Primitive.from_proto(response.name)
        self.interface_name = Primitive.from_proto(response.interface_name)
        self.ip_address = Primitive.from_proto(response.ip_address)
        self.peer_ip_address = Primitive.from_proto(response.peer_ip_address)
        self.peer_asn = Primitive.from_proto(response.peer_asn)
        self.advertised_route_priority = Primitive.from_proto(
            response.advertised_route_priority
        )
        self.advertise_mode = Primitive.from_proto(response.advertise_mode)
        self.management_type = Primitive.from_proto(response.management_type)
        self.advertised_groups = Primitive.from_proto(response.advertised_groups)
        self.advertised_ip_ranges = RouterPeerAdvertisedIPRangesArray.from_proto(
            response.advertised_ip_ranges
        )
        self.region = Primitive.from_proto(response.region)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = router_peer_pb2_grpc.ComputeRouterPeerServiceStub(channel.Channel())
        request = router_peer_pb2.DeleteComputeRouterPeerRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.router):
            request.resource.router = Primitive.to_proto(self.router)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.interface_name):
            request.resource.interface_name = Primitive.to_proto(self.interface_name)

        if Primitive.to_proto(self.ip_address):
            request.resource.ip_address = Primitive.to_proto(self.ip_address)

        if Primitive.to_proto(self.peer_ip_address):
            request.resource.peer_ip_address = Primitive.to_proto(self.peer_ip_address)

        if Primitive.to_proto(self.peer_asn):
            request.resource.peer_asn = Primitive.to_proto(self.peer_asn)

        if Primitive.to_proto(self.advertised_route_priority):
            request.resource.advertised_route_priority = Primitive.to_proto(
                self.advertised_route_priority
            )

        if Primitive.to_proto(self.advertise_mode):
            request.resource.advertise_mode = Primitive.to_proto(self.advertise_mode)

        if Primitive.to_proto(self.advertised_groups):
            request.resource.advertised_groups.extend(
                Primitive.to_proto(self.advertised_groups)
            )
        if RouterPeerAdvertisedIPRangesArray.to_proto(self.advertised_ip_ranges):
            request.resource.advertised_ip_ranges.extend(
                RouterPeerAdvertisedIPRangesArray.to_proto(self.advertised_ip_ranges)
            )
        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeRouterPeer(request)

    @classmethod
    def list(self, project, region, service_account_file=""):
        stub = router_peer_pb2_grpc.ComputeRouterPeerServiceStub(channel.Channel())
        request = router_peer_pb2.ListComputeRouterPeerRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        return stub.ListComputeRouterPeer(request).items

    def to_proto(self):
        resource = router_peer_pb2.ComputeRouterPeer()
        if Primitive.to_proto(self.router):
            resource.router = Primitive.to_proto(self.router)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.interface_name):
            resource.interface_name = Primitive.to_proto(self.interface_name)
        if Primitive.to_proto(self.ip_address):
            resource.ip_address = Primitive.to_proto(self.ip_address)
        if Primitive.to_proto(self.peer_ip_address):
            resource.peer_ip_address = Primitive.to_proto(self.peer_ip_address)
        if Primitive.to_proto(self.peer_asn):
            resource.peer_asn = Primitive.to_proto(self.peer_asn)
        if Primitive.to_proto(self.advertised_route_priority):
            resource.advertised_route_priority = Primitive.to_proto(
                self.advertised_route_priority
            )
        if Primitive.to_proto(self.advertise_mode):
            resource.advertise_mode = Primitive.to_proto(self.advertise_mode)
        if Primitive.to_proto(self.advertised_groups):
            resource.advertised_groups.extend(
                Primitive.to_proto(self.advertised_groups)
            )
        if RouterPeerAdvertisedIPRangesArray.to_proto(self.advertised_ip_ranges):
            resource.advertised_ip_ranges.extend(
                RouterPeerAdvertisedIPRangesArray.to_proto(self.advertised_ip_ranges)
            )
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class RouterPeerAdvertisedIPRanges(object):
    def __init__(self, range: str = None, description: str = None):
        self.range = range
        self.description = description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = router_peer_pb2.ComputeRouterPeerAdvertisedIPRanges()
        if Primitive.to_proto(resource.range):
            res.range = Primitive.to_proto(resource.range)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RouterPeerAdvertisedIPRanges(
            range=Primitive.from_proto(resource.range),
            description=Primitive.from_proto(resource.description),
        )


class RouterPeerAdvertisedIPRangesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RouterPeerAdvertisedIPRanges.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RouterPeerAdvertisedIPRanges.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
