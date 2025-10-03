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
from google3.cloud.graphite.mmv2.services.google.compute import vpn_tunnel_pb2
from google3.cloud.graphite.mmv2.services.google.compute import vpn_tunnel_pb2_grpc

from typing import List


class VpnTunnel(object):
    def __init__(
        self,
        labels: dict = None,
        id: int = None,
        name: str = None,
        description: str = None,
        location: str = None,
        target_vpn_gateway: str = None,
        vpn_gateway: str = None,
        vpn_gateway_interface: int = None,
        peer_external_gateway: str = None,
        peer_external_gateway_interface: int = None,
        peer_gcp_gateway: str = None,
        router: str = None,
        peer_ip: str = None,
        shared_secret: str = None,
        shared_secret_hash: str = None,
        status: str = None,
        self_link: str = None,
        ike_version: int = None,
        detailed_status: str = None,
        local_traffic_selector: list = None,
        remote_traffic_selector: list = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.labels = labels
        self.name = name
        self.description = description
        self.location = location
        self.target_vpn_gateway = target_vpn_gateway
        self.vpn_gateway = vpn_gateway
        self.vpn_gateway_interface = vpn_gateway_interface
        self.peer_external_gateway = peer_external_gateway
        self.peer_external_gateway_interface = peer_external_gateway_interface
        self.peer_gcp_gateway = peer_gcp_gateway
        self.router = router
        self.peer_ip = peer_ip
        self.shared_secret = shared_secret
        self.ike_version = ike_version
        self.local_traffic_selector = local_traffic_selector
        self.remote_traffic_selector = remote_traffic_selector
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = vpn_tunnel_pb2_grpc.ComputeAlphaVpnTunnelServiceStub(channel.Channel())
        request = vpn_tunnel_pb2.ApplyComputeAlphaVpnTunnelRequest()
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.target_vpn_gateway):
            request.resource.target_vpn_gateway = Primitive.to_proto(
                self.target_vpn_gateway
            )

        if Primitive.to_proto(self.vpn_gateway):
            request.resource.vpn_gateway = Primitive.to_proto(self.vpn_gateway)

        if Primitive.to_proto(self.vpn_gateway_interface):
            request.resource.vpn_gateway_interface = Primitive.to_proto(
                self.vpn_gateway_interface
            )

        if Primitive.to_proto(self.peer_external_gateway):
            request.resource.peer_external_gateway = Primitive.to_proto(
                self.peer_external_gateway
            )

        if Primitive.to_proto(self.peer_external_gateway_interface):
            request.resource.peer_external_gateway_interface = Primitive.to_proto(
                self.peer_external_gateway_interface
            )

        if Primitive.to_proto(self.peer_gcp_gateway):
            request.resource.peer_gcp_gateway = Primitive.to_proto(
                self.peer_gcp_gateway
            )

        if Primitive.to_proto(self.router):
            request.resource.router = Primitive.to_proto(self.router)

        if Primitive.to_proto(self.peer_ip):
            request.resource.peer_ip = Primitive.to_proto(self.peer_ip)

        if Primitive.to_proto(self.shared_secret):
            request.resource.shared_secret = Primitive.to_proto(self.shared_secret)

        if Primitive.to_proto(self.ike_version):
            request.resource.ike_version = Primitive.to_proto(self.ike_version)

        if Primitive.to_proto(self.local_traffic_selector):
            request.resource.local_traffic_selector.extend(
                Primitive.to_proto(self.local_traffic_selector)
            )
        if Primitive.to_proto(self.remote_traffic_selector):
            request.resource.remote_traffic_selector.extend(
                Primitive.to_proto(self.remote_traffic_selector)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeAlphaVpnTunnel(request)
        self.labels = Primitive.from_proto(response.labels)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.location = Primitive.from_proto(response.location)
        self.target_vpn_gateway = Primitive.from_proto(response.target_vpn_gateway)
        self.vpn_gateway = Primitive.from_proto(response.vpn_gateway)
        self.vpn_gateway_interface = Primitive.from_proto(
            response.vpn_gateway_interface
        )
        self.peer_external_gateway = Primitive.from_proto(
            response.peer_external_gateway
        )
        self.peer_external_gateway_interface = Primitive.from_proto(
            response.peer_external_gateway_interface
        )
        self.peer_gcp_gateway = Primitive.from_proto(response.peer_gcp_gateway)
        self.router = Primitive.from_proto(response.router)
        self.peer_ip = Primitive.from_proto(response.peer_ip)
        self.shared_secret = Primitive.from_proto(response.shared_secret)
        self.shared_secret_hash = Primitive.from_proto(response.shared_secret_hash)
        self.status = VpnTunnelStatusEnum.from_proto(response.status)
        self.self_link = Primitive.from_proto(response.self_link)
        self.ike_version = Primitive.from_proto(response.ike_version)
        self.detailed_status = Primitive.from_proto(response.detailed_status)
        self.local_traffic_selector = Primitive.from_proto(
            response.local_traffic_selector
        )
        self.remote_traffic_selector = Primitive.from_proto(
            response.remote_traffic_selector
        )
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = vpn_tunnel_pb2_grpc.ComputeAlphaVpnTunnelServiceStub(channel.Channel())
        request = vpn_tunnel_pb2.DeleteComputeAlphaVpnTunnelRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.target_vpn_gateway):
            request.resource.target_vpn_gateway = Primitive.to_proto(
                self.target_vpn_gateway
            )

        if Primitive.to_proto(self.vpn_gateway):
            request.resource.vpn_gateway = Primitive.to_proto(self.vpn_gateway)

        if Primitive.to_proto(self.vpn_gateway_interface):
            request.resource.vpn_gateway_interface = Primitive.to_proto(
                self.vpn_gateway_interface
            )

        if Primitive.to_proto(self.peer_external_gateway):
            request.resource.peer_external_gateway = Primitive.to_proto(
                self.peer_external_gateway
            )

        if Primitive.to_proto(self.peer_external_gateway_interface):
            request.resource.peer_external_gateway_interface = Primitive.to_proto(
                self.peer_external_gateway_interface
            )

        if Primitive.to_proto(self.peer_gcp_gateway):
            request.resource.peer_gcp_gateway = Primitive.to_proto(
                self.peer_gcp_gateway
            )

        if Primitive.to_proto(self.router):
            request.resource.router = Primitive.to_proto(self.router)

        if Primitive.to_proto(self.peer_ip):
            request.resource.peer_ip = Primitive.to_proto(self.peer_ip)

        if Primitive.to_proto(self.shared_secret):
            request.resource.shared_secret = Primitive.to_proto(self.shared_secret)

        if Primitive.to_proto(self.ike_version):
            request.resource.ike_version = Primitive.to_proto(self.ike_version)

        if Primitive.to_proto(self.local_traffic_selector):
            request.resource.local_traffic_selector.extend(
                Primitive.to_proto(self.local_traffic_selector)
            )
        if Primitive.to_proto(self.remote_traffic_selector):
            request.resource.remote_traffic_selector.extend(
                Primitive.to_proto(self.remote_traffic_selector)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeAlphaVpnTunnel(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = vpn_tunnel_pb2_grpc.ComputeAlphaVpnTunnelServiceStub(channel.Channel())
        request = vpn_tunnel_pb2.ListComputeAlphaVpnTunnelRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeAlphaVpnTunnel(request).items

    def to_proto(self):
        resource = vpn_tunnel_pb2.ComputeAlphaVpnTunnel()
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.target_vpn_gateway):
            resource.target_vpn_gateway = Primitive.to_proto(self.target_vpn_gateway)
        if Primitive.to_proto(self.vpn_gateway):
            resource.vpn_gateway = Primitive.to_proto(self.vpn_gateway)
        if Primitive.to_proto(self.vpn_gateway_interface):
            resource.vpn_gateway_interface = Primitive.to_proto(
                self.vpn_gateway_interface
            )
        if Primitive.to_proto(self.peer_external_gateway):
            resource.peer_external_gateway = Primitive.to_proto(
                self.peer_external_gateway
            )
        if Primitive.to_proto(self.peer_external_gateway_interface):
            resource.peer_external_gateway_interface = Primitive.to_proto(
                self.peer_external_gateway_interface
            )
        if Primitive.to_proto(self.peer_gcp_gateway):
            resource.peer_gcp_gateway = Primitive.to_proto(self.peer_gcp_gateway)
        if Primitive.to_proto(self.router):
            resource.router = Primitive.to_proto(self.router)
        if Primitive.to_proto(self.peer_ip):
            resource.peer_ip = Primitive.to_proto(self.peer_ip)
        if Primitive.to_proto(self.shared_secret):
            resource.shared_secret = Primitive.to_proto(self.shared_secret)
        if Primitive.to_proto(self.ike_version):
            resource.ike_version = Primitive.to_proto(self.ike_version)
        if Primitive.to_proto(self.local_traffic_selector):
            resource.local_traffic_selector.extend(
                Primitive.to_proto(self.local_traffic_selector)
            )
        if Primitive.to_proto(self.remote_traffic_selector):
            resource.remote_traffic_selector.extend(
                Primitive.to_proto(self.remote_traffic_selector)
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class VpnTunnelStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return vpn_tunnel_pb2.ComputeAlphaVpnTunnelStatusEnum.Value(
            "ComputeAlphaVpnTunnelStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return vpn_tunnel_pb2.ComputeAlphaVpnTunnelStatusEnum.Name(resource)[
            len("ComputeAlphaVpnTunnelStatusEnum") :
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
