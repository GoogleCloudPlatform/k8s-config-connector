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
from google3.cloud.graphite.mmv2.services.google.compute import packet_mirroring_pb2
from google3.cloud.graphite.mmv2.services.google.compute import (
    packet_mirroring_pb2_grpc,
)

from typing import List


class PacketMirroring(object):
    def __init__(
        self,
        id: int = None,
        self_link: str = None,
        name: str = None,
        description: str = None,
        region: str = None,
        network: dict = None,
        priority: int = None,
        collector_ilb: dict = None,
        mirrored_resources: dict = None,
        filter: dict = None,
        enable: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.network = network
        self.priority = priority
        self.collector_ilb = collector_ilb
        self.mirrored_resources = mirrored_resources
        self.filter = filter
        self.enable = enable
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = packet_mirroring_pb2_grpc.ComputeBetaPacketMirroringServiceStub(
            channel.Channel()
        )
        request = packet_mirroring_pb2.ApplyComputeBetaPacketMirroringRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if PacketMirroringNetwork.to_proto(self.network):
            request.resource.network.CopyFrom(
                PacketMirroringNetwork.to_proto(self.network)
            )
        else:
            request.resource.ClearField("network")
        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if PacketMirroringCollectorIlb.to_proto(self.collector_ilb):
            request.resource.collector_ilb.CopyFrom(
                PacketMirroringCollectorIlb.to_proto(self.collector_ilb)
            )
        else:
            request.resource.ClearField("collector_ilb")
        if PacketMirroringMirroredResources.to_proto(self.mirrored_resources):
            request.resource.mirrored_resources.CopyFrom(
                PacketMirroringMirroredResources.to_proto(self.mirrored_resources)
            )
        else:
            request.resource.ClearField("mirrored_resources")
        if PacketMirroringFilter.to_proto(self.filter):
            request.resource.filter.CopyFrom(
                PacketMirroringFilter.to_proto(self.filter)
            )
        else:
            request.resource.ClearField("filter")
        if PacketMirroringEnableEnum.to_proto(self.enable):
            request.resource.enable = PacketMirroringEnableEnum.to_proto(self.enable)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaPacketMirroring(request)
        self.id = Primitive.from_proto(response.id)
        self.self_link = Primitive.from_proto(response.self_link)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.region = Primitive.from_proto(response.region)
        self.network = PacketMirroringNetwork.from_proto(response.network)
        self.priority = Primitive.from_proto(response.priority)
        self.collector_ilb = PacketMirroringCollectorIlb.from_proto(
            response.collector_ilb
        )
        self.mirrored_resources = PacketMirroringMirroredResources.from_proto(
            response.mirrored_resources
        )
        self.filter = PacketMirroringFilter.from_proto(response.filter)
        self.enable = PacketMirroringEnableEnum.from_proto(response.enable)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = packet_mirroring_pb2_grpc.ComputeBetaPacketMirroringServiceStub(
            channel.Channel()
        )
        request = packet_mirroring_pb2.DeleteComputeBetaPacketMirroringRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if PacketMirroringNetwork.to_proto(self.network):
            request.resource.network.CopyFrom(
                PacketMirroringNetwork.to_proto(self.network)
            )
        else:
            request.resource.ClearField("network")
        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if PacketMirroringCollectorIlb.to_proto(self.collector_ilb):
            request.resource.collector_ilb.CopyFrom(
                PacketMirroringCollectorIlb.to_proto(self.collector_ilb)
            )
        else:
            request.resource.ClearField("collector_ilb")
        if PacketMirroringMirroredResources.to_proto(self.mirrored_resources):
            request.resource.mirrored_resources.CopyFrom(
                PacketMirroringMirroredResources.to_proto(self.mirrored_resources)
            )
        else:
            request.resource.ClearField("mirrored_resources")
        if PacketMirroringFilter.to_proto(self.filter):
            request.resource.filter.CopyFrom(
                PacketMirroringFilter.to_proto(self.filter)
            )
        else:
            request.resource.ClearField("filter")
        if PacketMirroringEnableEnum.to_proto(self.enable):
            request.resource.enable = PacketMirroringEnableEnum.to_proto(self.enable)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteComputeBetaPacketMirroring(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = packet_mirroring_pb2_grpc.ComputeBetaPacketMirroringServiceStub(
            channel.Channel()
        )
        request = packet_mirroring_pb2.ListComputeBetaPacketMirroringRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeBetaPacketMirroring(request).items

    def to_proto(self):
        resource = packet_mirroring_pb2.ComputeBetaPacketMirroring()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if PacketMirroringNetwork.to_proto(self.network):
            resource.network.CopyFrom(PacketMirroringNetwork.to_proto(self.network))
        else:
            resource.ClearField("network")
        if Primitive.to_proto(self.priority):
            resource.priority = Primitive.to_proto(self.priority)
        if PacketMirroringCollectorIlb.to_proto(self.collector_ilb):
            resource.collector_ilb.CopyFrom(
                PacketMirroringCollectorIlb.to_proto(self.collector_ilb)
            )
        else:
            resource.ClearField("collector_ilb")
        if PacketMirroringMirroredResources.to_proto(self.mirrored_resources):
            resource.mirrored_resources.CopyFrom(
                PacketMirroringMirroredResources.to_proto(self.mirrored_resources)
            )
        else:
            resource.ClearField("mirrored_resources")
        if PacketMirroringFilter.to_proto(self.filter):
            resource.filter.CopyFrom(PacketMirroringFilter.to_proto(self.filter))
        else:
            resource.ClearField("filter")
        if PacketMirroringEnableEnum.to_proto(self.enable):
            resource.enable = PacketMirroringEnableEnum.to_proto(self.enable)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class PacketMirroringNetwork(object):
    def __init__(self, url: str = None, canonical_url: str = None):
        self.url = url
        self.canonical_url = canonical_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = packet_mirroring_pb2.ComputeBetaPacketMirroringNetwork()
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if Primitive.to_proto(resource.canonical_url):
            res.canonical_url = Primitive.to_proto(resource.canonical_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PacketMirroringNetwork(
            url=Primitive.from_proto(resource.url),
            canonical_url=Primitive.from_proto(resource.canonical_url),
        )


class PacketMirroringNetworkArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PacketMirroringNetwork.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PacketMirroringNetwork.from_proto(i) for i in resources]


class PacketMirroringCollectorIlb(object):
    def __init__(self, url: str = None, canonical_url: str = None):
        self.url = url
        self.canonical_url = canonical_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = packet_mirroring_pb2.ComputeBetaPacketMirroringCollectorIlb()
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if Primitive.to_proto(resource.canonical_url):
            res.canonical_url = Primitive.to_proto(resource.canonical_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PacketMirroringCollectorIlb(
            url=Primitive.from_proto(resource.url),
            canonical_url=Primitive.from_proto(resource.canonical_url),
        )


class PacketMirroringCollectorIlbArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PacketMirroringCollectorIlb.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PacketMirroringCollectorIlb.from_proto(i) for i in resources]


class PacketMirroringMirroredResources(object):
    def __init__(
        self, subnetworks: list = None, instances: list = None, tags: list = None
    ):
        self.subnetworks = subnetworks
        self.instances = instances
        self.tags = tags

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = packet_mirroring_pb2.ComputeBetaPacketMirroringMirroredResources()
        if PacketMirroringMirroredResourcesSubnetworksArray.to_proto(
            resource.subnetworks
        ):
            res.subnetworks.extend(
                PacketMirroringMirroredResourcesSubnetworksArray.to_proto(
                    resource.subnetworks
                )
            )
        if PacketMirroringMirroredResourcesInstancesArray.to_proto(resource.instances):
            res.instances.extend(
                PacketMirroringMirroredResourcesInstancesArray.to_proto(
                    resource.instances
                )
            )
        if Primitive.to_proto(resource.tags):
            res.tags.extend(Primitive.to_proto(resource.tags))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PacketMirroringMirroredResources(
            subnetworks=PacketMirroringMirroredResourcesSubnetworksArray.from_proto(
                resource.subnetworks
            ),
            instances=PacketMirroringMirroredResourcesInstancesArray.from_proto(
                resource.instances
            ),
            tags=Primitive.from_proto(resource.tags),
        )


class PacketMirroringMirroredResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PacketMirroringMirroredResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PacketMirroringMirroredResources.from_proto(i) for i in resources]


class PacketMirroringMirroredResourcesSubnetworks(object):
    def __init__(self, url: str = None, canonical_url: str = None):
        self.url = url
        self.canonical_url = canonical_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            packet_mirroring_pb2.ComputeBetaPacketMirroringMirroredResourcesSubnetworks()
        )
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if Primitive.to_proto(resource.canonical_url):
            res.canonical_url = Primitive.to_proto(resource.canonical_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PacketMirroringMirroredResourcesSubnetworks(
            url=Primitive.from_proto(resource.url),
            canonical_url=Primitive.from_proto(resource.canonical_url),
        )


class PacketMirroringMirroredResourcesSubnetworksArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PacketMirroringMirroredResourcesSubnetworks.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PacketMirroringMirroredResourcesSubnetworks.from_proto(i) for i in resources
        ]


class PacketMirroringMirroredResourcesInstances(object):
    def __init__(self, url: str = None, canonical_url: str = None):
        self.url = url
        self.canonical_url = canonical_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            packet_mirroring_pb2.ComputeBetaPacketMirroringMirroredResourcesInstances()
        )
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if Primitive.to_proto(resource.canonical_url):
            res.canonical_url = Primitive.to_proto(resource.canonical_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PacketMirroringMirroredResourcesInstances(
            url=Primitive.from_proto(resource.url),
            canonical_url=Primitive.from_proto(resource.canonical_url),
        )


class PacketMirroringMirroredResourcesInstancesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PacketMirroringMirroredResourcesInstances.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PacketMirroringMirroredResourcesInstances.from_proto(i) for i in resources
        ]


class PacketMirroringFilter(object):
    def __init__(
        self, cidr_ranges: list = None, ip_protocols: list = None, direction: str = None
    ):
        self.cidr_ranges = cidr_ranges
        self.ip_protocols = ip_protocols
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = packet_mirroring_pb2.ComputeBetaPacketMirroringFilter()
        if Primitive.to_proto(resource.cidr_ranges):
            res.cidr_ranges.extend(Primitive.to_proto(resource.cidr_ranges))
        if Primitive.to_proto(resource.ip_protocols):
            res.ip_protocols.extend(Primitive.to_proto(resource.ip_protocols))
        if PacketMirroringFilterDirectionEnum.to_proto(resource.direction):
            res.direction = PacketMirroringFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PacketMirroringFilter(
            cidr_ranges=Primitive.from_proto(resource.cidr_ranges),
            ip_protocols=Primitive.from_proto(resource.ip_protocols),
            direction=PacketMirroringFilterDirectionEnum.from_proto(resource.direction),
        )


class PacketMirroringFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PacketMirroringFilter.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PacketMirroringFilter.from_proto(i) for i in resources]


class PacketMirroringFilterDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return packet_mirroring_pb2.ComputeBetaPacketMirroringFilterDirectionEnum.Value(
            "ComputeBetaPacketMirroringFilterDirectionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return packet_mirroring_pb2.ComputeBetaPacketMirroringFilterDirectionEnum.Name(
            resource
        )[len("ComputeBetaPacketMirroringFilterDirectionEnum") :]


class PacketMirroringEnableEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return packet_mirroring_pb2.ComputeBetaPacketMirroringEnableEnum.Value(
            "ComputeBetaPacketMirroringEnableEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return packet_mirroring_pb2.ComputeBetaPacketMirroringEnableEnum.Name(resource)[
            len("ComputeBetaPacketMirroringEnableEnum") :
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
