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
from google3.cloud.graphite.mmv2.services.google.compute import firewall_pb2
from google3.cloud.graphite.mmv2.services.google.compute import firewall_pb2_grpc

from typing import List


class Firewall(object):
    def __init__(
        self,
        creation_timestamp: str = None,
        description: str = None,
        direction: str = None,
        disabled: bool = None,
        id: str = None,
        log_config: dict = None,
        name: str = None,
        network: str = None,
        priority: int = None,
        self_link: str = None,
        project: str = None,
        allowed: list = None,
        denied: list = None,
        destination_ranges: list = None,
        source_ranges: list = None,
        source_service_accounts: list = None,
        source_tags: list = None,
        target_service_accounts: list = None,
        target_tags: list = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.description = description
        self.direction = direction
        self.disabled = disabled
        self.id = id
        self.log_config = log_config
        self.name = name
        self.network = network
        self.priority = priority
        self.project = project
        self.allowed = allowed
        self.denied = denied
        self.destination_ranges = destination_ranges
        self.source_ranges = source_ranges
        self.source_service_accounts = source_service_accounts
        self.source_tags = source_tags
        self.target_service_accounts = target_service_accounts
        self.target_tags = target_tags
        self.service_account_file = service_account_file

    def apply(self):
        stub = firewall_pb2_grpc.ComputeFirewallServiceStub(channel.Channel())
        request = firewall_pb2.ApplyComputeFirewallRequest()
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if FirewallDirectionEnum.to_proto(self.direction):
            request.resource.direction = FirewallDirectionEnum.to_proto(self.direction)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.id):
            request.resource.id = Primitive.to_proto(self.id)

        if FirewallLogConfig.to_proto(self.log_config):
            request.resource.log_config.CopyFrom(
                FirewallLogConfig.to_proto(self.log_config)
            )
        else:
            request.resource.ClearField("log_config")
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if FirewallAllowedArray.to_proto(self.allowed):
            request.resource.allowed.extend(FirewallAllowedArray.to_proto(self.allowed))
        if FirewallDeniedArray.to_proto(self.denied):
            request.resource.denied.extend(FirewallDeniedArray.to_proto(self.denied))
        if Primitive.to_proto(self.destination_ranges):
            request.resource.destination_ranges.extend(
                Primitive.to_proto(self.destination_ranges)
            )
        if Primitive.to_proto(self.source_ranges):
            request.resource.source_ranges.extend(
                Primitive.to_proto(self.source_ranges)
            )
        if Primitive.to_proto(self.source_service_accounts):
            request.resource.source_service_accounts.extend(
                Primitive.to_proto(self.source_service_accounts)
            )
        if Primitive.to_proto(self.source_tags):
            request.resource.source_tags.extend(Primitive.to_proto(self.source_tags))
        if Primitive.to_proto(self.target_service_accounts):
            request.resource.target_service_accounts.extend(
                Primitive.to_proto(self.target_service_accounts)
            )
        if Primitive.to_proto(self.target_tags):
            request.resource.target_tags.extend(Primitive.to_proto(self.target_tags))
        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeFirewall(request)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.description = Primitive.from_proto(response.description)
        self.direction = FirewallDirectionEnum.from_proto(response.direction)
        self.disabled = Primitive.from_proto(response.disabled)
        self.id = Primitive.from_proto(response.id)
        self.log_config = FirewallLogConfig.from_proto(response.log_config)
        self.name = Primitive.from_proto(response.name)
        self.network = Primitive.from_proto(response.network)
        self.priority = Primitive.from_proto(response.priority)
        self.self_link = Primitive.from_proto(response.self_link)
        self.project = Primitive.from_proto(response.project)
        self.allowed = FirewallAllowedArray.from_proto(response.allowed)
        self.denied = FirewallDeniedArray.from_proto(response.denied)
        self.destination_ranges = Primitive.from_proto(response.destination_ranges)
        self.source_ranges = Primitive.from_proto(response.source_ranges)
        self.source_service_accounts = Primitive.from_proto(
            response.source_service_accounts
        )
        self.source_tags = Primitive.from_proto(response.source_tags)
        self.target_service_accounts = Primitive.from_proto(
            response.target_service_accounts
        )
        self.target_tags = Primitive.from_proto(response.target_tags)

    def delete(self):
        stub = firewall_pb2_grpc.ComputeFirewallServiceStub(channel.Channel())
        request = firewall_pb2.DeleteComputeFirewallRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if FirewallDirectionEnum.to_proto(self.direction):
            request.resource.direction = FirewallDirectionEnum.to_proto(self.direction)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.id):
            request.resource.id = Primitive.to_proto(self.id)

        if FirewallLogConfig.to_proto(self.log_config):
            request.resource.log_config.CopyFrom(
                FirewallLogConfig.to_proto(self.log_config)
            )
        else:
            request.resource.ClearField("log_config")
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if FirewallAllowedArray.to_proto(self.allowed):
            request.resource.allowed.extend(FirewallAllowedArray.to_proto(self.allowed))
        if FirewallDeniedArray.to_proto(self.denied):
            request.resource.denied.extend(FirewallDeniedArray.to_proto(self.denied))
        if Primitive.to_proto(self.destination_ranges):
            request.resource.destination_ranges.extend(
                Primitive.to_proto(self.destination_ranges)
            )
        if Primitive.to_proto(self.source_ranges):
            request.resource.source_ranges.extend(
                Primitive.to_proto(self.source_ranges)
            )
        if Primitive.to_proto(self.source_service_accounts):
            request.resource.source_service_accounts.extend(
                Primitive.to_proto(self.source_service_accounts)
            )
        if Primitive.to_proto(self.source_tags):
            request.resource.source_tags.extend(Primitive.to_proto(self.source_tags))
        if Primitive.to_proto(self.target_service_accounts):
            request.resource.target_service_accounts.extend(
                Primitive.to_proto(self.target_service_accounts)
            )
        if Primitive.to_proto(self.target_tags):
            request.resource.target_tags.extend(Primitive.to_proto(self.target_tags))
        response = stub.DeleteComputeFirewall(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = firewall_pb2_grpc.ComputeFirewallServiceStub(channel.Channel())
        request = firewall_pb2.ListComputeFirewallRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeFirewall(request).items

    def to_proto(self):
        resource = firewall_pb2.ComputeFirewall()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if FirewallDirectionEnum.to_proto(self.direction):
            resource.direction = FirewallDirectionEnum.to_proto(self.direction)
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if Primitive.to_proto(self.id):
            resource.id = Primitive.to_proto(self.id)
        if FirewallLogConfig.to_proto(self.log_config):
            resource.log_config.CopyFrom(FirewallLogConfig.to_proto(self.log_config))
        else:
            resource.ClearField("log_config")
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if Primitive.to_proto(self.priority):
            resource.priority = Primitive.to_proto(self.priority)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if FirewallAllowedArray.to_proto(self.allowed):
            resource.allowed.extend(FirewallAllowedArray.to_proto(self.allowed))
        if FirewallDeniedArray.to_proto(self.denied):
            resource.denied.extend(FirewallDeniedArray.to_proto(self.denied))
        if Primitive.to_proto(self.destination_ranges):
            resource.destination_ranges.extend(
                Primitive.to_proto(self.destination_ranges)
            )
        if Primitive.to_proto(self.source_ranges):
            resource.source_ranges.extend(Primitive.to_proto(self.source_ranges))
        if Primitive.to_proto(self.source_service_accounts):
            resource.source_service_accounts.extend(
                Primitive.to_proto(self.source_service_accounts)
            )
        if Primitive.to_proto(self.source_tags):
            resource.source_tags.extend(Primitive.to_proto(self.source_tags))
        if Primitive.to_proto(self.target_service_accounts):
            resource.target_service_accounts.extend(
                Primitive.to_proto(self.target_service_accounts)
            )
        if Primitive.to_proto(self.target_tags):
            resource.target_tags.extend(Primitive.to_proto(self.target_tags))
        return resource


class FirewallLogConfig(object):
    def __init__(self, enable: bool = None):
        self.enable = enable

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = firewall_pb2.ComputeFirewallLogConfig()
        if Primitive.to_proto(resource.enable):
            res.enable = Primitive.to_proto(resource.enable)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FirewallLogConfig(enable=Primitive.from_proto(resource.enable),)


class FirewallLogConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FirewallLogConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FirewallLogConfig.from_proto(i) for i in resources]


class FirewallAllowed(object):
    def __init__(
        self, ip_protocol: str = None, ports: list = None, ip_protocol_alt: list = None
    ):
        self.ip_protocol = ip_protocol
        self.ports = ports
        self.ip_protocol_alt = ip_protocol_alt

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = firewall_pb2.ComputeFirewallAllowed()
        if Primitive.to_proto(resource.ip_protocol):
            res.ip_protocol = Primitive.to_proto(resource.ip_protocol)
        if Primitive.to_proto(resource.ports):
            res.ports.extend(Primitive.to_proto(resource.ports))
        if Primitive.to_proto(resource.ip_protocol_alt):
            res.ip_protocol_alt.extend(Primitive.to_proto(resource.ip_protocol_alt))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FirewallAllowed(
            ip_protocol=Primitive.from_proto(resource.ip_protocol),
            ports=Primitive.from_proto(resource.ports),
            ip_protocol_alt=Primitive.from_proto(resource.ip_protocol_alt),
        )


class FirewallAllowedArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FirewallAllowed.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FirewallAllowed.from_proto(i) for i in resources]


class FirewallDenied(object):
    def __init__(
        self, ip_protocol: str = None, ports: list = None, ip_protocol_alt: list = None
    ):
        self.ip_protocol = ip_protocol
        self.ports = ports
        self.ip_protocol_alt = ip_protocol_alt

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = firewall_pb2.ComputeFirewallDenied()
        if Primitive.to_proto(resource.ip_protocol):
            res.ip_protocol = Primitive.to_proto(resource.ip_protocol)
        if Primitive.to_proto(resource.ports):
            res.ports.extend(Primitive.to_proto(resource.ports))
        if Primitive.to_proto(resource.ip_protocol_alt):
            res.ip_protocol_alt.extend(Primitive.to_proto(resource.ip_protocol_alt))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FirewallDenied(
            ip_protocol=Primitive.from_proto(resource.ip_protocol),
            ports=Primitive.from_proto(resource.ports),
            ip_protocol_alt=Primitive.from_proto(resource.ip_protocol_alt),
        )


class FirewallDeniedArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FirewallDenied.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FirewallDenied.from_proto(i) for i in resources]


class FirewallDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return firewall_pb2.ComputeFirewallDirectionEnum.Value(
            "ComputeFirewallDirectionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return firewall_pb2.ComputeFirewallDirectionEnum.Name(resource)[
            len("ComputeFirewallDirectionEnum") :
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
