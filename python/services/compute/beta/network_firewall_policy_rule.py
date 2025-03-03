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
from google3.cloud.graphite.mmv2.services.google.compute import (
    network_firewall_policy_rule_pb2,
)
from google3.cloud.graphite.mmv2.services.google.compute import (
    network_firewall_policy_rule_pb2_grpc,
)

from typing import List


class NetworkFirewallPolicyRule(object):
    def __init__(
        self,
        description: str = None,
        rule_name: str = None,
        priority: int = None,
        location: str = None,
        match: dict = None,
        action: str = None,
        direction: str = None,
        enable_logging: bool = None,
        rule_tuple_count: int = None,
        target_service_accounts: list = None,
        target_secure_tags: list = None,
        disabled: bool = None,
        kind: str = None,
        firewall_policy: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.description = description
        self.rule_name = rule_name
        self.priority = priority
        self.location = location
        self.match = match
        self.action = action
        self.direction = direction
        self.enable_logging = enable_logging
        self.target_service_accounts = target_service_accounts
        self.target_secure_tags = target_secure_tags
        self.disabled = disabled
        self.firewall_policy = firewall_policy
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = network_firewall_policy_rule_pb2_grpc.ComputeBetaNetworkFirewallPolicyRuleServiceStub(
            channel.Channel()
        )
        request = (
            network_firewall_policy_rule_pb2.ApplyComputeBetaNetworkFirewallPolicyRuleRequest()
        )
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.rule_name):
            request.resource.rule_name = Primitive.to_proto(self.rule_name)

        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if NetworkFirewallPolicyRuleMatch.to_proto(self.match):
            request.resource.match.CopyFrom(
                NetworkFirewallPolicyRuleMatch.to_proto(self.match)
            )
        else:
            request.resource.ClearField("match")
        if Primitive.to_proto(self.action):
            request.resource.action = Primitive.to_proto(self.action)

        if NetworkFirewallPolicyRuleDirectionEnum.to_proto(self.direction):
            request.resource.direction = (
                NetworkFirewallPolicyRuleDirectionEnum.to_proto(self.direction)
            )

        if Primitive.to_proto(self.enable_logging):
            request.resource.enable_logging = Primitive.to_proto(self.enable_logging)

        if Primitive.to_proto(self.target_service_accounts):
            request.resource.target_service_accounts.extend(
                Primitive.to_proto(self.target_service_accounts)
            )
        if NetworkFirewallPolicyRuleTargetSecureTagsArray.to_proto(
            self.target_secure_tags
        ):
            request.resource.target_secure_tags.extend(
                NetworkFirewallPolicyRuleTargetSecureTagsArray.to_proto(
                    self.target_secure_tags
                )
            )
        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.firewall_policy):
            request.resource.firewall_policy = Primitive.to_proto(self.firewall_policy)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaNetworkFirewallPolicyRule(request)
        self.description = Primitive.from_proto(response.description)
        self.rule_name = Primitive.from_proto(response.rule_name)
        self.priority = Primitive.from_proto(response.priority)
        self.location = Primitive.from_proto(response.location)
        self.match = NetworkFirewallPolicyRuleMatch.from_proto(response.match)
        self.action = Primitive.from_proto(response.action)
        self.direction = NetworkFirewallPolicyRuleDirectionEnum.from_proto(
            response.direction
        )
        self.enable_logging = Primitive.from_proto(response.enable_logging)
        self.rule_tuple_count = Primitive.from_proto(response.rule_tuple_count)
        self.target_service_accounts = Primitive.from_proto(
            response.target_service_accounts
        )
        self.target_secure_tags = (
            NetworkFirewallPolicyRuleTargetSecureTagsArray.from_proto(
                response.target_secure_tags
            )
        )
        self.disabled = Primitive.from_proto(response.disabled)
        self.kind = Primitive.from_proto(response.kind)
        self.firewall_policy = Primitive.from_proto(response.firewall_policy)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = network_firewall_policy_rule_pb2_grpc.ComputeBetaNetworkFirewallPolicyRuleServiceStub(
            channel.Channel()
        )
        request = (
            network_firewall_policy_rule_pb2.DeleteComputeBetaNetworkFirewallPolicyRuleRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.rule_name):
            request.resource.rule_name = Primitive.to_proto(self.rule_name)

        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if NetworkFirewallPolicyRuleMatch.to_proto(self.match):
            request.resource.match.CopyFrom(
                NetworkFirewallPolicyRuleMatch.to_proto(self.match)
            )
        else:
            request.resource.ClearField("match")
        if Primitive.to_proto(self.action):
            request.resource.action = Primitive.to_proto(self.action)

        if NetworkFirewallPolicyRuleDirectionEnum.to_proto(self.direction):
            request.resource.direction = (
                NetworkFirewallPolicyRuleDirectionEnum.to_proto(self.direction)
            )

        if Primitive.to_proto(self.enable_logging):
            request.resource.enable_logging = Primitive.to_proto(self.enable_logging)

        if Primitive.to_proto(self.target_service_accounts):
            request.resource.target_service_accounts.extend(
                Primitive.to_proto(self.target_service_accounts)
            )
        if NetworkFirewallPolicyRuleTargetSecureTagsArray.to_proto(
            self.target_secure_tags
        ):
            request.resource.target_secure_tags.extend(
                NetworkFirewallPolicyRuleTargetSecureTagsArray.to_proto(
                    self.target_secure_tags
                )
            )
        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.firewall_policy):
            request.resource.firewall_policy = Primitive.to_proto(self.firewall_policy)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeBetaNetworkFirewallPolicyRule(request)

    @classmethod
    def list(self, project, location, firewallPolicy, service_account_file=""):
        stub = network_firewall_policy_rule_pb2_grpc.ComputeBetaNetworkFirewallPolicyRuleServiceStub(
            channel.Channel()
        )
        request = (
            network_firewall_policy_rule_pb2.ListComputeBetaNetworkFirewallPolicyRuleRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.FirewallPolicy = firewallPolicy

        return stub.ListComputeBetaNetworkFirewallPolicyRule(request).items

    def to_proto(self):
        resource = (
            network_firewall_policy_rule_pb2.ComputeBetaNetworkFirewallPolicyRule()
        )
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.rule_name):
            resource.rule_name = Primitive.to_proto(self.rule_name)
        if Primitive.to_proto(self.priority):
            resource.priority = Primitive.to_proto(self.priority)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if NetworkFirewallPolicyRuleMatch.to_proto(self.match):
            resource.match.CopyFrom(NetworkFirewallPolicyRuleMatch.to_proto(self.match))
        else:
            resource.ClearField("match")
        if Primitive.to_proto(self.action):
            resource.action = Primitive.to_proto(self.action)
        if NetworkFirewallPolicyRuleDirectionEnum.to_proto(self.direction):
            resource.direction = NetworkFirewallPolicyRuleDirectionEnum.to_proto(
                self.direction
            )
        if Primitive.to_proto(self.enable_logging):
            resource.enable_logging = Primitive.to_proto(self.enable_logging)
        if Primitive.to_proto(self.target_service_accounts):
            resource.target_service_accounts.extend(
                Primitive.to_proto(self.target_service_accounts)
            )
        if NetworkFirewallPolicyRuleTargetSecureTagsArray.to_proto(
            self.target_secure_tags
        ):
            resource.target_secure_tags.extend(
                NetworkFirewallPolicyRuleTargetSecureTagsArray.to_proto(
                    self.target_secure_tags
                )
            )
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if Primitive.to_proto(self.firewall_policy):
            resource.firewall_policy = Primitive.to_proto(self.firewall_policy)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class NetworkFirewallPolicyRuleMatch(object):
    def __init__(
        self,
        src_ip_ranges: list = None,
        dest_ip_ranges: list = None,
        layer4_configs: list = None,
        src_secure_tags: list = None,
        src_region_codes: list = None,
        dest_region_codes: list = None,
        src_threat_intelligences: list = None,
        dest_threat_intelligences: list = None,
        src_fqdns: list = None,
        dest_fqdns: list = None,
        src_address_groups: list = None,
        dest_address_groups: list = None,
    ):
        self.src_ip_ranges = src_ip_ranges
        self.dest_ip_ranges = dest_ip_ranges
        self.layer4_configs = layer4_configs
        self.src_secure_tags = src_secure_tags
        self.src_region_codes = src_region_codes
        self.dest_region_codes = dest_region_codes
        self.src_threat_intelligences = src_threat_intelligences
        self.dest_threat_intelligences = dest_threat_intelligences
        self.src_fqdns = src_fqdns
        self.dest_fqdns = dest_fqdns
        self.src_address_groups = src_address_groups
        self.dest_address_groups = dest_address_groups

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            network_firewall_policy_rule_pb2.ComputeBetaNetworkFirewallPolicyRuleMatch()
        )
        if Primitive.to_proto(resource.src_ip_ranges):
            res.src_ip_ranges.extend(Primitive.to_proto(resource.src_ip_ranges))
        if Primitive.to_proto(resource.dest_ip_ranges):
            res.dest_ip_ranges.extend(Primitive.to_proto(resource.dest_ip_ranges))
        if NetworkFirewallPolicyRuleMatchLayer4ConfigsArray.to_proto(
            resource.layer4_configs
        ):
            res.layer4_configs.extend(
                NetworkFirewallPolicyRuleMatchLayer4ConfigsArray.to_proto(
                    resource.layer4_configs
                )
            )
        if NetworkFirewallPolicyRuleMatchSrcSecureTagsArray.to_proto(
            resource.src_secure_tags
        ):
            res.src_secure_tags.extend(
                NetworkFirewallPolicyRuleMatchSrcSecureTagsArray.to_proto(
                    resource.src_secure_tags
                )
            )
        if Primitive.to_proto(resource.src_region_codes):
            res.src_region_codes.extend(Primitive.to_proto(resource.src_region_codes))
        if Primitive.to_proto(resource.dest_region_codes):
            res.dest_region_codes.extend(Primitive.to_proto(resource.dest_region_codes))
        if Primitive.to_proto(resource.src_threat_intelligences):
            res.src_threat_intelligences.extend(
                Primitive.to_proto(resource.src_threat_intelligences)
            )
        if Primitive.to_proto(resource.dest_threat_intelligences):
            res.dest_threat_intelligences.extend(
                Primitive.to_proto(resource.dest_threat_intelligences)
            )
        if Primitive.to_proto(resource.src_fqdns):
            res.src_fqdns.extend(Primitive.to_proto(resource.src_fqdns))
        if Primitive.to_proto(resource.dest_fqdns):
            res.dest_fqdns.extend(Primitive.to_proto(resource.dest_fqdns))
        if Primitive.to_proto(resource.src_address_groups):
            res.src_address_groups.extend(
                Primitive.to_proto(resource.src_address_groups)
            )
        if Primitive.to_proto(resource.dest_address_groups):
            res.dest_address_groups.extend(
                Primitive.to_proto(resource.dest_address_groups)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NetworkFirewallPolicyRuleMatch(
            src_ip_ranges=Primitive.from_proto(resource.src_ip_ranges),
            dest_ip_ranges=Primitive.from_proto(resource.dest_ip_ranges),
            layer4_configs=NetworkFirewallPolicyRuleMatchLayer4ConfigsArray.from_proto(
                resource.layer4_configs
            ),
            src_secure_tags=NetworkFirewallPolicyRuleMatchSrcSecureTagsArray.from_proto(
                resource.src_secure_tags
            ),
            src_region_codes=Primitive.from_proto(resource.src_region_codes),
            dest_region_codes=Primitive.from_proto(resource.dest_region_codes),
            src_threat_intelligences=Primitive.from_proto(
                resource.src_threat_intelligences
            ),
            dest_threat_intelligences=Primitive.from_proto(
                resource.dest_threat_intelligences
            ),
            src_fqdns=Primitive.from_proto(resource.src_fqdns),
            dest_fqdns=Primitive.from_proto(resource.dest_fqdns),
            src_address_groups=Primitive.from_proto(resource.src_address_groups),
            dest_address_groups=Primitive.from_proto(resource.dest_address_groups),
        )


class NetworkFirewallPolicyRuleMatchArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NetworkFirewallPolicyRuleMatch.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NetworkFirewallPolicyRuleMatch.from_proto(i) for i in resources]


class NetworkFirewallPolicyRuleMatchLayer4Configs(object):
    def __init__(self, ip_protocol: str = None, ports: list = None):
        self.ip_protocol = ip_protocol
        self.ports = ports

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            network_firewall_policy_rule_pb2.ComputeBetaNetworkFirewallPolicyRuleMatchLayer4Configs()
        )
        if Primitive.to_proto(resource.ip_protocol):
            res.ip_protocol = Primitive.to_proto(resource.ip_protocol)
        if Primitive.to_proto(resource.ports):
            res.ports.extend(Primitive.to_proto(resource.ports))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NetworkFirewallPolicyRuleMatchLayer4Configs(
            ip_protocol=Primitive.from_proto(resource.ip_protocol),
            ports=Primitive.from_proto(resource.ports),
        )


class NetworkFirewallPolicyRuleMatchLayer4ConfigsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            NetworkFirewallPolicyRuleMatchLayer4Configs.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            NetworkFirewallPolicyRuleMatchLayer4Configs.from_proto(i) for i in resources
        ]


class NetworkFirewallPolicyRuleMatchSrcSecureTags(object):
    def __init__(self, name: str = None, state: str = None):
        self.name = name
        self.state = state

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            network_firewall_policy_rule_pb2.ComputeBetaNetworkFirewallPolicyRuleMatchSrcSecureTags()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum.to_proto(
            resource.state
        ):
            res.state = NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum.to_proto(
                resource.state
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NetworkFirewallPolicyRuleMatchSrcSecureTags(
            name=Primitive.from_proto(resource.name),
            state=NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum.from_proto(
                resource.state
            ),
        )


class NetworkFirewallPolicyRuleMatchSrcSecureTagsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            NetworkFirewallPolicyRuleMatchSrcSecureTags.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            NetworkFirewallPolicyRuleMatchSrcSecureTags.from_proto(i) for i in resources
        ]


class NetworkFirewallPolicyRuleTargetSecureTags(object):
    def __init__(self, name: str = None, state: str = None):
        self.name = name
        self.state = state

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            network_firewall_policy_rule_pb2.ComputeBetaNetworkFirewallPolicyRuleTargetSecureTags()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if NetworkFirewallPolicyRuleTargetSecureTagsStateEnum.to_proto(resource.state):
            res.state = NetworkFirewallPolicyRuleTargetSecureTagsStateEnum.to_proto(
                resource.state
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NetworkFirewallPolicyRuleTargetSecureTags(
            name=Primitive.from_proto(resource.name),
            state=NetworkFirewallPolicyRuleTargetSecureTagsStateEnum.from_proto(
                resource.state
            ),
        )


class NetworkFirewallPolicyRuleTargetSecureTagsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            NetworkFirewallPolicyRuleTargetSecureTags.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            NetworkFirewallPolicyRuleTargetSecureTags.from_proto(i) for i in resources
        ]


class NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return network_firewall_policy_rule_pb2.ComputeBetaNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum.Value(
            "ComputeBetaNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return network_firewall_policy_rule_pb2.ComputeBetaNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum.Name(
            resource
        )[
            len("ComputeBetaNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum") :
        ]


class NetworkFirewallPolicyRuleDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return network_firewall_policy_rule_pb2.ComputeBetaNetworkFirewallPolicyRuleDirectionEnum.Value(
            "ComputeBetaNetworkFirewallPolicyRuleDirectionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return network_firewall_policy_rule_pb2.ComputeBetaNetworkFirewallPolicyRuleDirectionEnum.Name(
            resource
        )[
            len("ComputeBetaNetworkFirewallPolicyRuleDirectionEnum") :
        ]


class NetworkFirewallPolicyRuleTargetSecureTagsStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return network_firewall_policy_rule_pb2.ComputeBetaNetworkFirewallPolicyRuleTargetSecureTagsStateEnum.Value(
            "ComputeBetaNetworkFirewallPolicyRuleTargetSecureTagsStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return network_firewall_policy_rule_pb2.ComputeBetaNetworkFirewallPolicyRuleTargetSecureTagsStateEnum.Name(
            resource
        )[
            len("ComputeBetaNetworkFirewallPolicyRuleTargetSecureTagsStateEnum") :
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
