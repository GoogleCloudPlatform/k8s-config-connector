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
from google3.cloud.graphite.mmv2.services.google.network_connectivity import spoke_pb2
from google3.cloud.graphite.mmv2.services.google.network_connectivity import (
    spoke_pb2_grpc,
)

from typing import List


class Spoke(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        description: str = None,
        hub: str = None,
        linked_vpn_tunnels: dict = None,
        linked_interconnect_attachments: dict = None,
        linked_router_appliance_instances: dict = None,
        linked_vpc_network: dict = None,
        unique_id: str = None,
        state: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.labels = labels
        self.description = description
        self.hub = hub
        self.linked_vpn_tunnels = linked_vpn_tunnels
        self.linked_interconnect_attachments = linked_interconnect_attachments
        self.linked_router_appliance_instances = linked_router_appliance_instances
        self.linked_vpc_network = linked_vpc_network
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = spoke_pb2_grpc.NetworkconnectivitySpokeServiceStub(channel.Channel())
        request = spoke_pb2.ApplyNetworkconnectivitySpokeRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.hub):
            request.resource.hub = Primitive.to_proto(self.hub)

        if SpokeLinkedVpnTunnels.to_proto(self.linked_vpn_tunnels):
            request.resource.linked_vpn_tunnels.CopyFrom(
                SpokeLinkedVpnTunnels.to_proto(self.linked_vpn_tunnels)
            )
        else:
            request.resource.ClearField("linked_vpn_tunnels")
        if SpokeLinkedInterconnectAttachments.to_proto(
            self.linked_interconnect_attachments
        ):
            request.resource.linked_interconnect_attachments.CopyFrom(
                SpokeLinkedInterconnectAttachments.to_proto(
                    self.linked_interconnect_attachments
                )
            )
        else:
            request.resource.ClearField("linked_interconnect_attachments")
        if SpokeLinkedRouterApplianceInstances.to_proto(
            self.linked_router_appliance_instances
        ):
            request.resource.linked_router_appliance_instances.CopyFrom(
                SpokeLinkedRouterApplianceInstances.to_proto(
                    self.linked_router_appliance_instances
                )
            )
        else:
            request.resource.ClearField("linked_router_appliance_instances")
        if SpokeLinkedVPCNetwork.to_proto(self.linked_vpc_network):
            request.resource.linked_vpc_network.CopyFrom(
                SpokeLinkedVPCNetwork.to_proto(self.linked_vpc_network)
            )
        else:
            request.resource.ClearField("linked_vpc_network")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworkconnectivitySpoke(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.description = Primitive.from_proto(response.description)
        self.hub = Primitive.from_proto(response.hub)
        self.linked_vpn_tunnels = SpokeLinkedVpnTunnels.from_proto(
            response.linked_vpn_tunnels
        )
        self.linked_interconnect_attachments = (
            SpokeLinkedInterconnectAttachments.from_proto(
                response.linked_interconnect_attachments
            )
        )
        self.linked_router_appliance_instances = (
            SpokeLinkedRouterApplianceInstances.from_proto(
                response.linked_router_appliance_instances
            )
        )
        self.linked_vpc_network = SpokeLinkedVPCNetwork.from_proto(
            response.linked_vpc_network
        )
        self.unique_id = Primitive.from_proto(response.unique_id)
        self.state = SpokeStateEnum.from_proto(response.state)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = spoke_pb2_grpc.NetworkconnectivitySpokeServiceStub(channel.Channel())
        request = spoke_pb2.DeleteNetworkconnectivitySpokeRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.hub):
            request.resource.hub = Primitive.to_proto(self.hub)

        if SpokeLinkedVpnTunnels.to_proto(self.linked_vpn_tunnels):
            request.resource.linked_vpn_tunnels.CopyFrom(
                SpokeLinkedVpnTunnels.to_proto(self.linked_vpn_tunnels)
            )
        else:
            request.resource.ClearField("linked_vpn_tunnels")
        if SpokeLinkedInterconnectAttachments.to_proto(
            self.linked_interconnect_attachments
        ):
            request.resource.linked_interconnect_attachments.CopyFrom(
                SpokeLinkedInterconnectAttachments.to_proto(
                    self.linked_interconnect_attachments
                )
            )
        else:
            request.resource.ClearField("linked_interconnect_attachments")
        if SpokeLinkedRouterApplianceInstances.to_proto(
            self.linked_router_appliance_instances
        ):
            request.resource.linked_router_appliance_instances.CopyFrom(
                SpokeLinkedRouterApplianceInstances.to_proto(
                    self.linked_router_appliance_instances
                )
            )
        else:
            request.resource.ClearField("linked_router_appliance_instances")
        if SpokeLinkedVPCNetwork.to_proto(self.linked_vpc_network):
            request.resource.linked_vpc_network.CopyFrom(
                SpokeLinkedVPCNetwork.to_proto(self.linked_vpc_network)
            )
        else:
            request.resource.ClearField("linked_vpc_network")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworkconnectivitySpoke(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = spoke_pb2_grpc.NetworkconnectivitySpokeServiceStub(channel.Channel())
        request = spoke_pb2.ListNetworkconnectivitySpokeRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworkconnectivitySpoke(request).items

    def to_proto(self):
        resource = spoke_pb2.NetworkconnectivitySpoke()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.hub):
            resource.hub = Primitive.to_proto(self.hub)
        if SpokeLinkedVpnTunnels.to_proto(self.linked_vpn_tunnels):
            resource.linked_vpn_tunnels.CopyFrom(
                SpokeLinkedVpnTunnels.to_proto(self.linked_vpn_tunnels)
            )
        else:
            resource.ClearField("linked_vpn_tunnels")
        if SpokeLinkedInterconnectAttachments.to_proto(
            self.linked_interconnect_attachments
        ):
            resource.linked_interconnect_attachments.CopyFrom(
                SpokeLinkedInterconnectAttachments.to_proto(
                    self.linked_interconnect_attachments
                )
            )
        else:
            resource.ClearField("linked_interconnect_attachments")
        if SpokeLinkedRouterApplianceInstances.to_proto(
            self.linked_router_appliance_instances
        ):
            resource.linked_router_appliance_instances.CopyFrom(
                SpokeLinkedRouterApplianceInstances.to_proto(
                    self.linked_router_appliance_instances
                )
            )
        else:
            resource.ClearField("linked_router_appliance_instances")
        if SpokeLinkedVPCNetwork.to_proto(self.linked_vpc_network):
            resource.linked_vpc_network.CopyFrom(
                SpokeLinkedVPCNetwork.to_proto(self.linked_vpc_network)
            )
        else:
            resource.ClearField("linked_vpc_network")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class SpokeLinkedVpnTunnels(object):
    def __init__(self, uris: list = None, site_to_site_data_transfer: bool = None):
        self.uris = uris
        self.site_to_site_data_transfer = site_to_site_data_transfer

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = spoke_pb2.NetworkconnectivitySpokeLinkedVpnTunnels()
        if Primitive.to_proto(resource.uris):
            res.uris.extend(Primitive.to_proto(resource.uris))
        if Primitive.to_proto(resource.site_to_site_data_transfer):
            res.site_to_site_data_transfer = Primitive.to_proto(
                resource.site_to_site_data_transfer
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SpokeLinkedVpnTunnels(
            uris=Primitive.from_proto(resource.uris),
            site_to_site_data_transfer=Primitive.from_proto(
                resource.site_to_site_data_transfer
            ),
        )


class SpokeLinkedVpnTunnelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SpokeLinkedVpnTunnels.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SpokeLinkedVpnTunnels.from_proto(i) for i in resources]


class SpokeLinkedInterconnectAttachments(object):
    def __init__(self, uris: list = None, site_to_site_data_transfer: bool = None):
        self.uris = uris
        self.site_to_site_data_transfer = site_to_site_data_transfer

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = spoke_pb2.NetworkconnectivitySpokeLinkedInterconnectAttachments()
        if Primitive.to_proto(resource.uris):
            res.uris.extend(Primitive.to_proto(resource.uris))
        if Primitive.to_proto(resource.site_to_site_data_transfer):
            res.site_to_site_data_transfer = Primitive.to_proto(
                resource.site_to_site_data_transfer
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SpokeLinkedInterconnectAttachments(
            uris=Primitive.from_proto(resource.uris),
            site_to_site_data_transfer=Primitive.from_proto(
                resource.site_to_site_data_transfer
            ),
        )


class SpokeLinkedInterconnectAttachmentsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SpokeLinkedInterconnectAttachments.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SpokeLinkedInterconnectAttachments.from_proto(i) for i in resources]


class SpokeLinkedRouterApplianceInstances(object):
    def __init__(self, instances: list = None, site_to_site_data_transfer: bool = None):
        self.instances = instances
        self.site_to_site_data_transfer = site_to_site_data_transfer

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = spoke_pb2.NetworkconnectivitySpokeLinkedRouterApplianceInstances()
        if SpokeLinkedRouterApplianceInstancesInstancesArray.to_proto(
            resource.instances
        ):
            res.instances.extend(
                SpokeLinkedRouterApplianceInstancesInstancesArray.to_proto(
                    resource.instances
                )
            )
        if Primitive.to_proto(resource.site_to_site_data_transfer):
            res.site_to_site_data_transfer = Primitive.to_proto(
                resource.site_to_site_data_transfer
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SpokeLinkedRouterApplianceInstances(
            instances=SpokeLinkedRouterApplianceInstancesInstancesArray.from_proto(
                resource.instances
            ),
            site_to_site_data_transfer=Primitive.from_proto(
                resource.site_to_site_data_transfer
            ),
        )


class SpokeLinkedRouterApplianceInstancesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SpokeLinkedRouterApplianceInstances.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SpokeLinkedRouterApplianceInstances.from_proto(i) for i in resources]


class SpokeLinkedRouterApplianceInstancesInstances(object):
    def __init__(self, virtual_machine: str = None, ip_address: str = None):
        self.virtual_machine = virtual_machine
        self.ip_address = ip_address

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            spoke_pb2.NetworkconnectivitySpokeLinkedRouterApplianceInstancesInstances()
        )
        if Primitive.to_proto(resource.virtual_machine):
            res.virtual_machine = Primitive.to_proto(resource.virtual_machine)
        if Primitive.to_proto(resource.ip_address):
            res.ip_address = Primitive.to_proto(resource.ip_address)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SpokeLinkedRouterApplianceInstancesInstances(
            virtual_machine=Primitive.from_proto(resource.virtual_machine),
            ip_address=Primitive.from_proto(resource.ip_address),
        )


class SpokeLinkedRouterApplianceInstancesInstancesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            SpokeLinkedRouterApplianceInstancesInstances.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            SpokeLinkedRouterApplianceInstancesInstances.from_proto(i)
            for i in resources
        ]


class SpokeLinkedVPCNetwork(object):
    def __init__(self, uri: str = None, exclude_export_ranges: list = None):
        self.uri = uri
        self.exclude_export_ranges = exclude_export_ranges

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = spoke_pb2.NetworkconnectivitySpokeLinkedVPCNetwork()
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.exclude_export_ranges):
            res.exclude_export_ranges.extend(
                Primitive.to_proto(resource.exclude_export_ranges)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SpokeLinkedVPCNetwork(
            uri=Primitive.from_proto(resource.uri),
            exclude_export_ranges=Primitive.from_proto(resource.exclude_export_ranges),
        )


class SpokeLinkedVPCNetworkArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SpokeLinkedVPCNetwork.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SpokeLinkedVPCNetwork.from_proto(i) for i in resources]


class SpokeStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return spoke_pb2.NetworkconnectivitySpokeStateEnum.Value(
            "NetworkconnectivitySpokeStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return spoke_pb2.NetworkconnectivitySpokeStateEnum.Name(resource)[
            len("NetworkconnectivitySpokeStateEnum") :
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
