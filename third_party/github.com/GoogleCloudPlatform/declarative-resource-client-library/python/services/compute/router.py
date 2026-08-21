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
from google3.cloud.graphite.mmv2.services.google.compute import router_pb2
from google3.cloud.graphite.mmv2.services.google.compute import router_pb2_grpc

from typing import List


class Router(object):
    def __init__(
        self,
        creation_timestamp: str = None,
        nats: list = None,
        name: str = None,
        network: str = None,
        interfaces: list = None,
        description: str = None,
        bgp_peers: list = None,
        bgp: dict = None,
        region: str = None,
        project: str = None,
        self_link: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.nats = nats
        self.name = name
        self.network = network
        self.interfaces = interfaces
        self.description = description
        self.bgp_peers = bgp_peers
        self.bgp = bgp
        self.region = region
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = router_pb2_grpc.ComputeRouterServiceStub(channel.Channel())
        request = router_pb2.ApplyComputeRouterRequest()
        if RouterNatsArray.to_proto(self.nats):
            request.resource.nats.extend(RouterNatsArray.to_proto(self.nats))
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if RouterInterfacesArray.to_proto(self.interfaces):
            request.resource.interfaces.extend(
                RouterInterfacesArray.to_proto(self.interfaces)
            )
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if RouterBgpPeersArray.to_proto(self.bgp_peers):
            request.resource.bgp_peers.extend(
                RouterBgpPeersArray.to_proto(self.bgp_peers)
            )
        if RouterBgp.to_proto(self.bgp):
            request.resource.bgp.CopyFrom(RouterBgp.to_proto(self.bgp))
        else:
            request.resource.ClearField("bgp")
        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeRouter(request)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.nats = RouterNatsArray.from_proto(response.nats)
        self.name = Primitive.from_proto(response.name)
        self.network = Primitive.from_proto(response.network)
        self.interfaces = RouterInterfacesArray.from_proto(response.interfaces)
        self.description = Primitive.from_proto(response.description)
        self.bgp_peers = RouterBgpPeersArray.from_proto(response.bgp_peers)
        self.bgp = RouterBgp.from_proto(response.bgp)
        self.region = Primitive.from_proto(response.region)
        self.project = Primitive.from_proto(response.project)
        self.self_link = Primitive.from_proto(response.self_link)

    def delete(self):
        stub = router_pb2_grpc.ComputeRouterServiceStub(channel.Channel())
        request = router_pb2.DeleteComputeRouterRequest()
        request.service_account_file = self.service_account_file
        if RouterNatsArray.to_proto(self.nats):
            request.resource.nats.extend(RouterNatsArray.to_proto(self.nats))
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if RouterInterfacesArray.to_proto(self.interfaces):
            request.resource.interfaces.extend(
                RouterInterfacesArray.to_proto(self.interfaces)
            )
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if RouterBgpPeersArray.to_proto(self.bgp_peers):
            request.resource.bgp_peers.extend(
                RouterBgpPeersArray.to_proto(self.bgp_peers)
            )
        if RouterBgp.to_proto(self.bgp):
            request.resource.bgp.CopyFrom(RouterBgp.to_proto(self.bgp))
        else:
            request.resource.ClearField("bgp")
        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeRouter(request)

    @classmethod
    def list(self, project, region, service_account_file=""):
        stub = router_pb2_grpc.ComputeRouterServiceStub(channel.Channel())
        request = router_pb2.ListComputeRouterRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        return stub.ListComputeRouter(request).items

    def to_proto(self):
        resource = router_pb2.ComputeRouter()
        if RouterNatsArray.to_proto(self.nats):
            resource.nats.extend(RouterNatsArray.to_proto(self.nats))
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if RouterInterfacesArray.to_proto(self.interfaces):
            resource.interfaces.extend(RouterInterfacesArray.to_proto(self.interfaces))
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if RouterBgpPeersArray.to_proto(self.bgp_peers):
            resource.bgp_peers.extend(RouterBgpPeersArray.to_proto(self.bgp_peers))
        if RouterBgp.to_proto(self.bgp):
            resource.bgp.CopyFrom(RouterBgp.to_proto(self.bgp))
        else:
            resource.ClearField("bgp")
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class RouterNats(object):
    def __init__(
        self,
        name: str = None,
        log_config: dict = None,
        source_subnetwork_ip_ranges_to_nat: str = None,
        nat_ips: list = None,
        drain_nat_ips: list = None,
        nat_ip_allocate_option: list = None,
        min_ports_per_vm: int = None,
        udp_idle_timeout_sec: int = None,
        icmp_idle_timeout_sec: int = None,
        tcp_established_idle_timeout_sec: int = None,
        tcp_transitory_idle_timeout_sec: int = None,
        subnetworks: list = None,
    ):
        self.name = name
        self.log_config = log_config
        self.source_subnetwork_ip_ranges_to_nat = source_subnetwork_ip_ranges_to_nat
        self.nat_ips = nat_ips
        self.drain_nat_ips = drain_nat_ips
        self.nat_ip_allocate_option = nat_ip_allocate_option
        self.min_ports_per_vm = min_ports_per_vm
        self.udp_idle_timeout_sec = udp_idle_timeout_sec
        self.icmp_idle_timeout_sec = icmp_idle_timeout_sec
        self.tcp_established_idle_timeout_sec = tcp_established_idle_timeout_sec
        self.tcp_transitory_idle_timeout_sec = tcp_transitory_idle_timeout_sec
        self.subnetworks = subnetworks

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = router_pb2.ComputeRouterNats()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if RouterNatsLogConfig.to_proto(resource.log_config):
            res.log_config.CopyFrom(RouterNatsLogConfig.to_proto(resource.log_config))
        else:
            res.ClearField("log_config")
        if RouterNatsSourceSubnetworkIPRangesToNatEnum.to_proto(
            resource.source_subnetwork_ip_ranges_to_nat
        ):
            res.source_subnetwork_ip_ranges_to_nat = RouterNatsSourceSubnetworkIPRangesToNatEnum.to_proto(
                resource.source_subnetwork_ip_ranges_to_nat
            )
        if Primitive.to_proto(resource.nat_ips):
            res.nat_ips.extend(Primitive.to_proto(resource.nat_ips))
        if Primitive.to_proto(resource.drain_nat_ips):
            res.drain_nat_ips.extend(Primitive.to_proto(resource.drain_nat_ips))
        if RouterNatsNatIPAllocateOptionEnumArray.to_proto(
            resource.nat_ip_allocate_option
        ):
            res.nat_ip_allocate_option.extend(
                RouterNatsNatIPAllocateOptionEnumArray.to_proto(
                    resource.nat_ip_allocate_option
                )
            )
        if Primitive.to_proto(resource.min_ports_per_vm):
            res.min_ports_per_vm = Primitive.to_proto(resource.min_ports_per_vm)
        if Primitive.to_proto(resource.udp_idle_timeout_sec):
            res.udp_idle_timeout_sec = Primitive.to_proto(resource.udp_idle_timeout_sec)
        if Primitive.to_proto(resource.icmp_idle_timeout_sec):
            res.icmp_idle_timeout_sec = Primitive.to_proto(
                resource.icmp_idle_timeout_sec
            )
        if Primitive.to_proto(resource.tcp_established_idle_timeout_sec):
            res.tcp_established_idle_timeout_sec = Primitive.to_proto(
                resource.tcp_established_idle_timeout_sec
            )
        if Primitive.to_proto(resource.tcp_transitory_idle_timeout_sec):
            res.tcp_transitory_idle_timeout_sec = Primitive.to_proto(
                resource.tcp_transitory_idle_timeout_sec
            )
        if RouterNatsSubnetworksArray.to_proto(resource.subnetworks):
            res.subnetworks.extend(
                RouterNatsSubnetworksArray.to_proto(resource.subnetworks)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RouterNats(
            name=Primitive.from_proto(resource.name),
            log_config=RouterNatsLogConfig.from_proto(resource.log_config),
            source_subnetwork_ip_ranges_to_nat=RouterNatsSourceSubnetworkIPRangesToNatEnum.from_proto(
                resource.source_subnetwork_ip_ranges_to_nat
            ),
            nat_ips=Primitive.from_proto(resource.nat_ips),
            drain_nat_ips=Primitive.from_proto(resource.drain_nat_ips),
            nat_ip_allocate_option=RouterNatsNatIPAllocateOptionEnumArray.from_proto(
                resource.nat_ip_allocate_option
            ),
            min_ports_per_vm=Primitive.from_proto(resource.min_ports_per_vm),
            udp_idle_timeout_sec=Primitive.from_proto(resource.udp_idle_timeout_sec),
            icmp_idle_timeout_sec=Primitive.from_proto(resource.icmp_idle_timeout_sec),
            tcp_established_idle_timeout_sec=Primitive.from_proto(
                resource.tcp_established_idle_timeout_sec
            ),
            tcp_transitory_idle_timeout_sec=Primitive.from_proto(
                resource.tcp_transitory_idle_timeout_sec
            ),
            subnetworks=RouterNatsSubnetworksArray.from_proto(resource.subnetworks),
        )


class RouterNatsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RouterNats.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RouterNats.from_proto(i) for i in resources]


class RouterNatsLogConfig(object):
    def __init__(self, enable: bool = None, filter: str = None):
        self.enable = enable
        self.filter = filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = router_pb2.ComputeRouterNatsLogConfig()
        if Primitive.to_proto(resource.enable):
            res.enable = Primitive.to_proto(resource.enable)
        if RouterNatsLogConfigFilterEnum.to_proto(resource.filter):
            res.filter = RouterNatsLogConfigFilterEnum.to_proto(resource.filter)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RouterNatsLogConfig(
            enable=Primitive.from_proto(resource.enable),
            filter=RouterNatsLogConfigFilterEnum.from_proto(resource.filter),
        )


class RouterNatsLogConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RouterNatsLogConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RouterNatsLogConfig.from_proto(i) for i in resources]


class RouterNatsSubnetworks(object):
    def __init__(
        self,
        name: str = None,
        source_ip_ranges_to_nat: str = None,
        secondary_ip_range_names: str = None,
    ):
        self.name = name
        self.source_ip_ranges_to_nat = source_ip_ranges_to_nat
        self.secondary_ip_range_names = secondary_ip_range_names

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = router_pb2.ComputeRouterNatsSubnetworks()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.source_ip_ranges_to_nat):
            res.source_ip_ranges_to_nat = Primitive.to_proto(
                resource.source_ip_ranges_to_nat
            )
        if Primitive.to_proto(resource.secondary_ip_range_names):
            res.secondary_ip_range_names = Primitive.to_proto(
                resource.secondary_ip_range_names
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RouterNatsSubnetworks(
            name=Primitive.from_proto(resource.name),
            source_ip_ranges_to_nat=Primitive.from_proto(
                resource.source_ip_ranges_to_nat
            ),
            secondary_ip_range_names=Primitive.from_proto(
                resource.secondary_ip_range_names
            ),
        )


class RouterNatsSubnetworksArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RouterNatsSubnetworks.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RouterNatsSubnetworks.from_proto(i) for i in resources]


class RouterInterfaces(object):
    def __init__(
        self,
        name: str = None,
        linked_vpn_tunnel: str = None,
        ip_range: str = None,
        management_type: str = None,
    ):
        self.name = name
        self.linked_vpn_tunnel = linked_vpn_tunnel
        self.ip_range = ip_range
        self.management_type = management_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = router_pb2.ComputeRouterInterfaces()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.linked_vpn_tunnel):
            res.linked_vpn_tunnel = Primitive.to_proto(resource.linked_vpn_tunnel)
        if Primitive.to_proto(resource.ip_range):
            res.ip_range = Primitive.to_proto(resource.ip_range)
        if RouterInterfacesManagementTypeEnum.to_proto(resource.management_type):
            res.management_type = RouterInterfacesManagementTypeEnum.to_proto(
                resource.management_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RouterInterfaces(
            name=Primitive.from_proto(resource.name),
            linked_vpn_tunnel=Primitive.from_proto(resource.linked_vpn_tunnel),
            ip_range=Primitive.from_proto(resource.ip_range),
            management_type=RouterInterfacesManagementTypeEnum.from_proto(
                resource.management_type
            ),
        )


class RouterInterfacesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RouterInterfaces.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RouterInterfaces.from_proto(i) for i in resources]


class RouterBgpPeers(object):
    def __init__(
        self,
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
    ):
        self.name = name
        self.interface_name = interface_name
        self.ip_address = ip_address
        self.peer_ip_address = peer_ip_address
        self.peer_asn = peer_asn
        self.advertised_route_priority = advertised_route_priority
        self.advertise_mode = advertise_mode
        self.management_type = management_type
        self.advertised_groups = advertised_groups
        self.advertised_ip_ranges = advertised_ip_ranges

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = router_pb2.ComputeRouterBgpPeers()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.interface_name):
            res.interface_name = Primitive.to_proto(resource.interface_name)
        if Primitive.to_proto(resource.ip_address):
            res.ip_address = Primitive.to_proto(resource.ip_address)
        if Primitive.to_proto(resource.peer_ip_address):
            res.peer_ip_address = Primitive.to_proto(resource.peer_ip_address)
        if Primitive.to_proto(resource.peer_asn):
            res.peer_asn = Primitive.to_proto(resource.peer_asn)
        if Primitive.to_proto(resource.advertised_route_priority):
            res.advertised_route_priority = Primitive.to_proto(
                resource.advertised_route_priority
            )
        if Primitive.to_proto(resource.advertise_mode):
            res.advertise_mode = Primitive.to_proto(resource.advertise_mode)
        if Primitive.to_proto(resource.management_type):
            res.management_type = Primitive.to_proto(resource.management_type)
        if RouterBgpPeersAdvertisedGroupsEnumArray.to_proto(resource.advertised_groups):
            res.advertised_groups.extend(
                RouterBgpPeersAdvertisedGroupsEnumArray.to_proto(
                    resource.advertised_groups
                )
            )
        if RouterBgpPeersAdvertisedIPRangesArray.to_proto(
            resource.advertised_ip_ranges
        ):
            res.advertised_ip_ranges.extend(
                RouterBgpPeersAdvertisedIPRangesArray.to_proto(
                    resource.advertised_ip_ranges
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RouterBgpPeers(
            name=Primitive.from_proto(resource.name),
            interface_name=Primitive.from_proto(resource.interface_name),
            ip_address=Primitive.from_proto(resource.ip_address),
            peer_ip_address=Primitive.from_proto(resource.peer_ip_address),
            peer_asn=Primitive.from_proto(resource.peer_asn),
            advertised_route_priority=Primitive.from_proto(
                resource.advertised_route_priority
            ),
            advertise_mode=Primitive.from_proto(resource.advertise_mode),
            management_type=Primitive.from_proto(resource.management_type),
            advertised_groups=RouterBgpPeersAdvertisedGroupsEnumArray.from_proto(
                resource.advertised_groups
            ),
            advertised_ip_ranges=RouterBgpPeersAdvertisedIPRangesArray.from_proto(
                resource.advertised_ip_ranges
            ),
        )


class RouterBgpPeersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RouterBgpPeers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RouterBgpPeers.from_proto(i) for i in resources]


class RouterBgpPeersAdvertisedIPRanges(object):
    def __init__(self, range: str = None, description: str = None):
        self.range = range
        self.description = description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = router_pb2.ComputeRouterBgpPeersAdvertisedIPRanges()
        if Primitive.to_proto(resource.range):
            res.range = Primitive.to_proto(resource.range)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RouterBgpPeersAdvertisedIPRanges(
            range=Primitive.from_proto(resource.range),
            description=Primitive.from_proto(resource.description),
        )


class RouterBgpPeersAdvertisedIPRangesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RouterBgpPeersAdvertisedIPRanges.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RouterBgpPeersAdvertisedIPRanges.from_proto(i) for i in resources]


class RouterBgp(object):
    def __init__(
        self,
        asn: int = None,
        advertise_mode: str = None,
        advertised_groups: list = None,
        advertised_ip_ranges: list = None,
    ):
        self.asn = asn
        self.advertise_mode = advertise_mode
        self.advertised_groups = advertised_groups
        self.advertised_ip_ranges = advertised_ip_ranges

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = router_pb2.ComputeRouterBgp()
        if Primitive.to_proto(resource.asn):
            res.asn = Primitive.to_proto(resource.asn)
        if RouterBgpAdvertiseModeEnum.to_proto(resource.advertise_mode):
            res.advertise_mode = RouterBgpAdvertiseModeEnum.to_proto(
                resource.advertise_mode
            )
        if Primitive.to_proto(resource.advertised_groups):
            res.advertised_groups.extend(Primitive.to_proto(resource.advertised_groups))
        if RouterBgpAdvertisedIPRangesArray.to_proto(resource.advertised_ip_ranges):
            res.advertised_ip_ranges.extend(
                RouterBgpAdvertisedIPRangesArray.to_proto(resource.advertised_ip_ranges)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RouterBgp(
            asn=Primitive.from_proto(resource.asn),
            advertise_mode=RouterBgpAdvertiseModeEnum.from_proto(
                resource.advertise_mode
            ),
            advertised_groups=Primitive.from_proto(resource.advertised_groups),
            advertised_ip_ranges=RouterBgpAdvertisedIPRangesArray.from_proto(
                resource.advertised_ip_ranges
            ),
        )


class RouterBgpArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RouterBgp.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RouterBgp.from_proto(i) for i in resources]


class RouterBgpAdvertisedIPRanges(object):
    def __init__(self, range: str = None, description: str = None):
        self.range = range
        self.description = description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = router_pb2.ComputeRouterBgpAdvertisedIPRanges()
        if Primitive.to_proto(resource.range):
            res.range = Primitive.to_proto(resource.range)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RouterBgpAdvertisedIPRanges(
            range=Primitive.from_proto(resource.range),
            description=Primitive.from_proto(resource.description),
        )


class RouterBgpAdvertisedIPRangesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RouterBgpAdvertisedIPRanges.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RouterBgpAdvertisedIPRanges.from_proto(i) for i in resources]


class RouterNatsLogConfigFilterEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterNatsLogConfigFilterEnum.Value(
            "ComputeRouterNatsLogConfigFilterEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterNatsLogConfigFilterEnum.Name(resource)[
            len("ComputeRouterNatsLogConfigFilterEnum") :
        ]


class RouterNatsSourceSubnetworkIPRangesToNatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum.Value(
            "ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum.Name(
            resource
        )[len("ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum") :]


class RouterNatsNatIPAllocateOptionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterNatsNatIPAllocateOptionEnum.Value(
            "ComputeRouterNatsNatIPAllocateOptionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterNatsNatIPAllocateOptionEnum.Name(resource)[
            len("ComputeRouterNatsNatIPAllocateOptionEnum") :
        ]


class RouterInterfacesManagementTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterInterfacesManagementTypeEnum.Value(
            "ComputeRouterInterfacesManagementTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterInterfacesManagementTypeEnum.Name(resource)[
            len("ComputeRouterInterfacesManagementTypeEnum") :
        ]


class RouterBgpPeersAdvertisedGroupsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterBgpPeersAdvertisedGroupsEnum.Value(
            "ComputeRouterBgpPeersAdvertisedGroupsEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterBgpPeersAdvertisedGroupsEnum.Name(resource)[
            len("ComputeRouterBgpPeersAdvertisedGroupsEnum") :
        ]


class RouterBgpAdvertiseModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterBgpAdvertiseModeEnum.Value(
            "ComputeRouterBgpAdvertiseModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return router_pb2.ComputeRouterBgpAdvertiseModeEnum.Name(resource)[
            len("ComputeRouterBgpAdvertiseModeEnum") :
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
