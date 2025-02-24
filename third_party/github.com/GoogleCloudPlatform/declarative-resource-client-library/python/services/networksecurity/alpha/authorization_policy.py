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
    authorization_policy_pb2,
)
from google3.cloud.graphite.mmv2.services.google.network_security import (
    authorization_policy_pb2_grpc,
)

from typing import List


class AuthorizationPolicy(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        action: str = None,
        rules: list = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.labels = labels
        self.action = action
        self.rules = rules
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = authorization_policy_pb2_grpc.NetworksecurityAlphaAuthorizationPolicyServiceStub(
            channel.Channel()
        )
        request = (
            authorization_policy_pb2.ApplyNetworksecurityAlphaAuthorizationPolicyRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if AuthorizationPolicyActionEnum.to_proto(self.action):
            request.resource.action = AuthorizationPolicyActionEnum.to_proto(
                self.action
            )

        if AuthorizationPolicyRulesArray.to_proto(self.rules):
            request.resource.rules.extend(
                AuthorizationPolicyRulesArray.to_proto(self.rules)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworksecurityAlphaAuthorizationPolicy(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.action = AuthorizationPolicyActionEnum.from_proto(response.action)
        self.rules = AuthorizationPolicyRulesArray.from_proto(response.rules)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = authorization_policy_pb2_grpc.NetworksecurityAlphaAuthorizationPolicyServiceStub(
            channel.Channel()
        )
        request = (
            authorization_policy_pb2.DeleteNetworksecurityAlphaAuthorizationPolicyRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if AuthorizationPolicyActionEnum.to_proto(self.action):
            request.resource.action = AuthorizationPolicyActionEnum.to_proto(
                self.action
            )

        if AuthorizationPolicyRulesArray.to_proto(self.rules):
            request.resource.rules.extend(
                AuthorizationPolicyRulesArray.to_proto(self.rules)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworksecurityAlphaAuthorizationPolicy(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = authorization_policy_pb2_grpc.NetworksecurityAlphaAuthorizationPolicyServiceStub(
            channel.Channel()
        )
        request = (
            authorization_policy_pb2.ListNetworksecurityAlphaAuthorizationPolicyRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworksecurityAlphaAuthorizationPolicy(request).items

    def to_proto(self):
        resource = authorization_policy_pb2.NetworksecurityAlphaAuthorizationPolicy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if AuthorizationPolicyActionEnum.to_proto(self.action):
            resource.action = AuthorizationPolicyActionEnum.to_proto(self.action)
        if AuthorizationPolicyRulesArray.to_proto(self.rules):
            resource.rules.extend(AuthorizationPolicyRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class AuthorizationPolicyRules(object):
    def __init__(self, sources: list = None, destinations: list = None):
        self.sources = sources
        self.destinations = destinations

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = authorization_policy_pb2.NetworksecurityAlphaAuthorizationPolicyRules()
        if AuthorizationPolicyRulesSourcesArray.to_proto(resource.sources):
            res.sources.extend(
                AuthorizationPolicyRulesSourcesArray.to_proto(resource.sources)
            )
        if AuthorizationPolicyRulesDestinationsArray.to_proto(resource.destinations):
            res.destinations.extend(
                AuthorizationPolicyRulesDestinationsArray.to_proto(
                    resource.destinations
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AuthorizationPolicyRules(
            sources=AuthorizationPolicyRulesSourcesArray.from_proto(resource.sources),
            destinations=AuthorizationPolicyRulesDestinationsArray.from_proto(
                resource.destinations
            ),
        )


class AuthorizationPolicyRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AuthorizationPolicyRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AuthorizationPolicyRules.from_proto(i) for i in resources]


class AuthorizationPolicyRulesSources(object):
    def __init__(self, principals: list = None, ip_blocks: list = None):
        self.principals = principals
        self.ip_blocks = ip_blocks

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            authorization_policy_pb2.NetworksecurityAlphaAuthorizationPolicyRulesSources()
        )
        if Primitive.to_proto(resource.principals):
            res.principals.extend(Primitive.to_proto(resource.principals))
        if Primitive.to_proto(resource.ip_blocks):
            res.ip_blocks.extend(Primitive.to_proto(resource.ip_blocks))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AuthorizationPolicyRulesSources(
            principals=Primitive.from_proto(resource.principals),
            ip_blocks=Primitive.from_proto(resource.ip_blocks),
        )


class AuthorizationPolicyRulesSourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AuthorizationPolicyRulesSources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AuthorizationPolicyRulesSources.from_proto(i) for i in resources]


class AuthorizationPolicyRulesDestinations(object):
    def __init__(
        self,
        hosts: list = None,
        ports: list = None,
        methods: list = None,
        http_header_match: dict = None,
    ):
        self.hosts = hosts
        self.ports = ports
        self.methods = methods
        self.http_header_match = http_header_match

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            authorization_policy_pb2.NetworksecurityAlphaAuthorizationPolicyRulesDestinations()
        )
        if Primitive.to_proto(resource.hosts):
            res.hosts.extend(Primitive.to_proto(resource.hosts))
        if int64Array.to_proto(resource.ports):
            res.ports.extend(int64Array.to_proto(resource.ports))
        if Primitive.to_proto(resource.methods):
            res.methods.extend(Primitive.to_proto(resource.methods))
        if AuthorizationPolicyRulesDestinationsHttpHeaderMatch.to_proto(
            resource.http_header_match
        ):
            res.http_header_match.CopyFrom(
                AuthorizationPolicyRulesDestinationsHttpHeaderMatch.to_proto(
                    resource.http_header_match
                )
            )
        else:
            res.ClearField("http_header_match")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AuthorizationPolicyRulesDestinations(
            hosts=Primitive.from_proto(resource.hosts),
            ports=int64Array.from_proto(resource.ports),
            methods=Primitive.from_proto(resource.methods),
            http_header_match=AuthorizationPolicyRulesDestinationsHttpHeaderMatch.from_proto(
                resource.http_header_match
            ),
        )


class AuthorizationPolicyRulesDestinationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AuthorizationPolicyRulesDestinations.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AuthorizationPolicyRulesDestinations.from_proto(i) for i in resources]


class AuthorizationPolicyRulesDestinationsHttpHeaderMatch(object):
    def __init__(self, header_name: str = None, regex_match: str = None):
        self.header_name = header_name
        self.regex_match = regex_match

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            authorization_policy_pb2.NetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatch()
        )
        if Primitive.to_proto(resource.header_name):
            res.header_name = Primitive.to_proto(resource.header_name)
        if Primitive.to_proto(resource.regex_match):
            res.regex_match = Primitive.to_proto(resource.regex_match)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AuthorizationPolicyRulesDestinationsHttpHeaderMatch(
            header_name=Primitive.from_proto(resource.header_name),
            regex_match=Primitive.from_proto(resource.regex_match),
        )


class AuthorizationPolicyRulesDestinationsHttpHeaderMatchArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AuthorizationPolicyRulesDestinationsHttpHeaderMatch.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AuthorizationPolicyRulesDestinationsHttpHeaderMatch.from_proto(i)
            for i in resources
        ]


class AuthorizationPolicyActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return authorization_policy_pb2.NetworksecurityAlphaAuthorizationPolicyActionEnum.Value(
            "NetworksecurityAlphaAuthorizationPolicyActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return authorization_policy_pb2.NetworksecurityAlphaAuthorizationPolicyActionEnum.Name(
            resource
        )[
            len("NetworksecurityAlphaAuthorizationPolicyActionEnum") :
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
