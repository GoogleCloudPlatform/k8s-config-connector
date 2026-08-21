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
from google3.cloud.graphite.mmv2.services.google.dns import managed_zone_pb2
from google3.cloud.graphite.mmv2.services.google.dns import managed_zone_pb2_grpc

from typing import List


class ManagedZone(object):
    def __init__(
        self,
        description: str = None,
        dns_name: str = None,
        dnssec_config: dict = None,
        name: str = None,
        name_servers: list = None,
        labels: dict = None,
        visibility: str = None,
        private_visibility_config: dict = None,
        forwarding_config: dict = None,
        reverse_lookup: bool = None,
        peering_config: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.description = description
        self.dns_name = dns_name
        self.dnssec_config = dnssec_config
        self.name = name
        self.labels = labels
        self.visibility = visibility
        self.private_visibility_config = private_visibility_config
        self.forwarding_config = forwarding_config
        self.reverse_lookup = reverse_lookup
        self.peering_config = peering_config
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = managed_zone_pb2_grpc.DnsManagedZoneServiceStub(channel.Channel())
        request = managed_zone_pb2.ApplyDnsManagedZoneRequest()
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.dns_name):
            request.resource.dns_name = Primitive.to_proto(self.dns_name)

        if ManagedZoneDnssecConfig.to_proto(self.dnssec_config):
            request.resource.dnssec_config.CopyFrom(
                ManagedZoneDnssecConfig.to_proto(self.dnssec_config)
            )
        else:
            request.resource.ClearField("dnssec_config")
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if ManagedZoneVisibilityEnum.to_proto(self.visibility):
            request.resource.visibility = ManagedZoneVisibilityEnum.to_proto(
                self.visibility
            )

        if ManagedZonePrivateVisibilityConfig.to_proto(self.private_visibility_config):
            request.resource.private_visibility_config.CopyFrom(
                ManagedZonePrivateVisibilityConfig.to_proto(
                    self.private_visibility_config
                )
            )
        else:
            request.resource.ClearField("private_visibility_config")
        if ManagedZoneForwardingConfig.to_proto(self.forwarding_config):
            request.resource.forwarding_config.CopyFrom(
                ManagedZoneForwardingConfig.to_proto(self.forwarding_config)
            )
        else:
            request.resource.ClearField("forwarding_config")
        if Primitive.to_proto(self.reverse_lookup):
            request.resource.reverse_lookup = Primitive.to_proto(self.reverse_lookup)

        if ManagedZonePeeringConfig.to_proto(self.peering_config):
            request.resource.peering_config.CopyFrom(
                ManagedZonePeeringConfig.to_proto(self.peering_config)
            )
        else:
            request.resource.ClearField("peering_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDnsManagedZone(request)
        self.description = Primitive.from_proto(response.description)
        self.dns_name = Primitive.from_proto(response.dns_name)
        self.dnssec_config = ManagedZoneDnssecConfig.from_proto(response.dnssec_config)
        self.name = Primitive.from_proto(response.name)
        self.name_servers = Primitive.from_proto(response.name_servers)
        self.labels = Primitive.from_proto(response.labels)
        self.visibility = ManagedZoneVisibilityEnum.from_proto(response.visibility)
        self.private_visibility_config = ManagedZonePrivateVisibilityConfig.from_proto(
            response.private_visibility_config
        )
        self.forwarding_config = ManagedZoneForwardingConfig.from_proto(
            response.forwarding_config
        )
        self.reverse_lookup = Primitive.from_proto(response.reverse_lookup)
        self.peering_config = ManagedZonePeeringConfig.from_proto(
            response.peering_config
        )
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = managed_zone_pb2_grpc.DnsManagedZoneServiceStub(channel.Channel())
        request = managed_zone_pb2.DeleteDnsManagedZoneRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.dns_name):
            request.resource.dns_name = Primitive.to_proto(self.dns_name)

        if ManagedZoneDnssecConfig.to_proto(self.dnssec_config):
            request.resource.dnssec_config.CopyFrom(
                ManagedZoneDnssecConfig.to_proto(self.dnssec_config)
            )
        else:
            request.resource.ClearField("dnssec_config")
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if ManagedZoneVisibilityEnum.to_proto(self.visibility):
            request.resource.visibility = ManagedZoneVisibilityEnum.to_proto(
                self.visibility
            )

        if ManagedZonePrivateVisibilityConfig.to_proto(self.private_visibility_config):
            request.resource.private_visibility_config.CopyFrom(
                ManagedZonePrivateVisibilityConfig.to_proto(
                    self.private_visibility_config
                )
            )
        else:
            request.resource.ClearField("private_visibility_config")
        if ManagedZoneForwardingConfig.to_proto(self.forwarding_config):
            request.resource.forwarding_config.CopyFrom(
                ManagedZoneForwardingConfig.to_proto(self.forwarding_config)
            )
        else:
            request.resource.ClearField("forwarding_config")
        if Primitive.to_proto(self.reverse_lookup):
            request.resource.reverse_lookup = Primitive.to_proto(self.reverse_lookup)

        if ManagedZonePeeringConfig.to_proto(self.peering_config):
            request.resource.peering_config.CopyFrom(
                ManagedZonePeeringConfig.to_proto(self.peering_config)
            )
        else:
            request.resource.ClearField("peering_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteDnsManagedZone(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = managed_zone_pb2_grpc.DnsManagedZoneServiceStub(channel.Channel())
        request = managed_zone_pb2.ListDnsManagedZoneRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListDnsManagedZone(request).items

    def to_proto(self):
        resource = managed_zone_pb2.DnsManagedZone()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.dns_name):
            resource.dns_name = Primitive.to_proto(self.dns_name)
        if ManagedZoneDnssecConfig.to_proto(self.dnssec_config):
            resource.dnssec_config.CopyFrom(
                ManagedZoneDnssecConfig.to_proto(self.dnssec_config)
            )
        else:
            resource.ClearField("dnssec_config")
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if ManagedZoneVisibilityEnum.to_proto(self.visibility):
            resource.visibility = ManagedZoneVisibilityEnum.to_proto(self.visibility)
        if ManagedZonePrivateVisibilityConfig.to_proto(self.private_visibility_config):
            resource.private_visibility_config.CopyFrom(
                ManagedZonePrivateVisibilityConfig.to_proto(
                    self.private_visibility_config
                )
            )
        else:
            resource.ClearField("private_visibility_config")
        if ManagedZoneForwardingConfig.to_proto(self.forwarding_config):
            resource.forwarding_config.CopyFrom(
                ManagedZoneForwardingConfig.to_proto(self.forwarding_config)
            )
        else:
            resource.ClearField("forwarding_config")
        if Primitive.to_proto(self.reverse_lookup):
            resource.reverse_lookup = Primitive.to_proto(self.reverse_lookup)
        if ManagedZonePeeringConfig.to_proto(self.peering_config):
            resource.peering_config.CopyFrom(
                ManagedZonePeeringConfig.to_proto(self.peering_config)
            )
        else:
            resource.ClearField("peering_config")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class ManagedZoneDnssecConfig(object):
    def __init__(
        self,
        kind: str = None,
        non_existence: str = None,
        state: str = None,
        default_key_specs: list = None,
    ):
        self.kind = kind
        self.non_existence = non_existence
        self.state = state
        self.default_key_specs = default_key_specs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = managed_zone_pb2.DnsManagedZoneDnssecConfig()
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if ManagedZoneDnssecConfigNonExistenceEnum.to_proto(resource.non_existence):
            res.non_existence = ManagedZoneDnssecConfigNonExistenceEnum.to_proto(
                resource.non_existence
            )
        if ManagedZoneDnssecConfigStateEnum.to_proto(resource.state):
            res.state = ManagedZoneDnssecConfigStateEnum.to_proto(resource.state)
        if ManagedZoneDnssecConfigDefaultKeySpecsArray.to_proto(
            resource.default_key_specs
        ):
            res.default_key_specs.extend(
                ManagedZoneDnssecConfigDefaultKeySpecsArray.to_proto(
                    resource.default_key_specs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ManagedZoneDnssecConfig(
            kind=Primitive.from_proto(resource.kind),
            non_existence=ManagedZoneDnssecConfigNonExistenceEnum.from_proto(
                resource.non_existence
            ),
            state=ManagedZoneDnssecConfigStateEnum.from_proto(resource.state),
            default_key_specs=ManagedZoneDnssecConfigDefaultKeySpecsArray.from_proto(
                resource.default_key_specs
            ),
        )


class ManagedZoneDnssecConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ManagedZoneDnssecConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ManagedZoneDnssecConfig.from_proto(i) for i in resources]


class ManagedZoneDnssecConfigDefaultKeySpecs(object):
    def __init__(
        self,
        algorithm: str = None,
        key_length: int = None,
        key_type: str = None,
        kind: str = None,
    ):
        self.algorithm = algorithm
        self.key_length = key_length
        self.key_type = key_type
        self.kind = kind

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = managed_zone_pb2.DnsManagedZoneDnssecConfigDefaultKeySpecs()
        if ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum.to_proto(
            resource.algorithm
        ):
            res.algorithm = ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum.to_proto(
                resource.algorithm
            )
        if Primitive.to_proto(resource.key_length):
            res.key_length = Primitive.to_proto(resource.key_length)
        if ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum.to_proto(
            resource.key_type
        ):
            res.key_type = ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum.to_proto(
                resource.key_type
            )
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ManagedZoneDnssecConfigDefaultKeySpecs(
            algorithm=ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum.from_proto(
                resource.algorithm
            ),
            key_length=Primitive.from_proto(resource.key_length),
            key_type=ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum.from_proto(
                resource.key_type
            ),
            kind=Primitive.from_proto(resource.kind),
        )


class ManagedZoneDnssecConfigDefaultKeySpecsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ManagedZoneDnssecConfigDefaultKeySpecs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ManagedZoneDnssecConfigDefaultKeySpecs.from_proto(i) for i in resources]


class ManagedZonePrivateVisibilityConfig(object):
    def __init__(self, networks: list = None):
        self.networks = networks

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = managed_zone_pb2.DnsManagedZonePrivateVisibilityConfig()
        if ManagedZonePrivateVisibilityConfigNetworksArray.to_proto(resource.networks):
            res.networks.extend(
                ManagedZonePrivateVisibilityConfigNetworksArray.to_proto(
                    resource.networks
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ManagedZonePrivateVisibilityConfig(
            networks=ManagedZonePrivateVisibilityConfigNetworksArray.from_proto(
                resource.networks
            ),
        )


class ManagedZonePrivateVisibilityConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ManagedZonePrivateVisibilityConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ManagedZonePrivateVisibilityConfig.from_proto(i) for i in resources]


class ManagedZonePrivateVisibilityConfigNetworks(object):
    def __init__(self, network_url: str = None):
        self.network_url = network_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = managed_zone_pb2.DnsManagedZonePrivateVisibilityConfigNetworks()
        if Primitive.to_proto(resource.network_url):
            res.network_url = Primitive.to_proto(resource.network_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ManagedZonePrivateVisibilityConfigNetworks(
            network_url=Primitive.from_proto(resource.network_url),
        )


class ManagedZonePrivateVisibilityConfigNetworksArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ManagedZonePrivateVisibilityConfigNetworks.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ManagedZonePrivateVisibilityConfigNetworks.from_proto(i) for i in resources
        ]


class ManagedZoneForwardingConfig(object):
    def __init__(self, target_name_servers: list = None):
        self.target_name_servers = target_name_servers

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = managed_zone_pb2.DnsManagedZoneForwardingConfig()
        if ManagedZoneForwardingConfigTargetNameServersArray.to_proto(
            resource.target_name_servers
        ):
            res.target_name_servers.extend(
                ManagedZoneForwardingConfigTargetNameServersArray.to_proto(
                    resource.target_name_servers
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ManagedZoneForwardingConfig(
            target_name_servers=ManagedZoneForwardingConfigTargetNameServersArray.from_proto(
                resource.target_name_servers
            ),
        )


class ManagedZoneForwardingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ManagedZoneForwardingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ManagedZoneForwardingConfig.from_proto(i) for i in resources]


class ManagedZoneForwardingConfigTargetNameServers(object):
    def __init__(self, ipv4_address: str = None, forwarding_path: str = None):
        self.ipv4_address = ipv4_address
        self.forwarding_path = forwarding_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = managed_zone_pb2.DnsManagedZoneForwardingConfigTargetNameServers()
        if Primitive.to_proto(resource.ipv4_address):
            res.ipv4_address = Primitive.to_proto(resource.ipv4_address)
        if ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum.to_proto(
            resource.forwarding_path
        ):
            res.forwarding_path = ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum.to_proto(
                resource.forwarding_path
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ManagedZoneForwardingConfigTargetNameServers(
            ipv4_address=Primitive.from_proto(resource.ipv4_address),
            forwarding_path=ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum.from_proto(
                resource.forwarding_path
            ),
        )


class ManagedZoneForwardingConfigTargetNameServersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ManagedZoneForwardingConfigTargetNameServers.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ManagedZoneForwardingConfigTargetNameServers.from_proto(i)
            for i in resources
        ]


class ManagedZonePeeringConfig(object):
    def __init__(self, target_network: dict = None):
        self.target_network = target_network

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = managed_zone_pb2.DnsManagedZonePeeringConfig()
        if ManagedZonePeeringConfigTargetNetwork.to_proto(resource.target_network):
            res.target_network.CopyFrom(
                ManagedZonePeeringConfigTargetNetwork.to_proto(resource.target_network)
            )
        else:
            res.ClearField("target_network")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ManagedZonePeeringConfig(
            target_network=ManagedZonePeeringConfigTargetNetwork.from_proto(
                resource.target_network
            ),
        )


class ManagedZonePeeringConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ManagedZonePeeringConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ManagedZonePeeringConfig.from_proto(i) for i in resources]


class ManagedZonePeeringConfigTargetNetwork(object):
    def __init__(self, network_url: str = None):
        self.network_url = network_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = managed_zone_pb2.DnsManagedZonePeeringConfigTargetNetwork()
        if Primitive.to_proto(resource.network_url):
            res.network_url = Primitive.to_proto(resource.network_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ManagedZonePeeringConfigTargetNetwork(
            network_url=Primitive.from_proto(resource.network_url),
        )


class ManagedZonePeeringConfigTargetNetworkArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ManagedZonePeeringConfigTargetNetwork.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ManagedZonePeeringConfigTargetNetwork.from_proto(i) for i in resources]


class ManagedZoneDnssecConfigNonExistenceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneDnssecConfigNonExistenceEnum.Value(
            "DnsManagedZoneDnssecConfigNonExistenceEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneDnssecConfigNonExistenceEnum.Name(
            resource
        )[len("DnsManagedZoneDnssecConfigNonExistenceEnum") :]


class ManagedZoneDnssecConfigStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneDnssecConfigStateEnum.Value(
            "DnsManagedZoneDnssecConfigStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneDnssecConfigStateEnum.Name(resource)[
            len("DnsManagedZoneDnssecConfigStateEnum") :
        ]


class ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum.Value(
            "DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum.Name(
            resource
        )[
            len("DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum") :
        ]


class ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum.Value(
            "DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum.Name(
            resource
        )[
            len("DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum") :
        ]


class ManagedZoneVisibilityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneVisibilityEnum.Value(
            "DnsManagedZoneVisibilityEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneVisibilityEnum.Name(resource)[
            len("DnsManagedZoneVisibilityEnum") :
        ]


class ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum.Value(
            "DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return managed_zone_pb2.DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum.Name(
            resource
        )[
            len("DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum") :
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
