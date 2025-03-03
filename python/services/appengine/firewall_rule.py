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
from google3.cloud.graphite.mmv2.services.google.app_engine import firewall_rule_pb2
from google3.cloud.graphite.mmv2.services.google.app_engine import (
    firewall_rule_pb2_grpc,
)

from typing import List


class FirewallRule(object):
    def __init__(
        self,
        action: str = None,
        description: str = None,
        priority: int = None,
        source_range: str = None,
        app: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.action = action
        self.description = description
        self.priority = priority
        self.source_range = source_range
        self.app = app
        self.service_account_file = service_account_file

    def apply(self):
        stub = firewall_rule_pb2_grpc.AppengineFirewallRuleServiceStub(
            channel.Channel()
        )
        request = firewall_rule_pb2.ApplyAppengineFirewallRuleRequest()
        if FirewallRuleActionEnum.to_proto(self.action):
            request.resource.action = FirewallRuleActionEnum.to_proto(self.action)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if Primitive.to_proto(self.source_range):
            request.resource.source_range = Primitive.to_proto(self.source_range)

        if Primitive.to_proto(self.app):
            request.resource.app = Primitive.to_proto(self.app)

        request.service_account_file = self.service_account_file

        response = stub.ApplyAppengineFirewallRule(request)
        self.action = FirewallRuleActionEnum.from_proto(response.action)
        self.description = Primitive.from_proto(response.description)
        self.priority = Primitive.from_proto(response.priority)
        self.source_range = Primitive.from_proto(response.source_range)
        self.app = Primitive.from_proto(response.app)

    def delete(self):
        stub = firewall_rule_pb2_grpc.AppengineFirewallRuleServiceStub(
            channel.Channel()
        )
        request = firewall_rule_pb2.DeleteAppengineFirewallRuleRequest()
        request.service_account_file = self.service_account_file
        if FirewallRuleActionEnum.to_proto(self.action):
            request.resource.action = FirewallRuleActionEnum.to_proto(self.action)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if Primitive.to_proto(self.source_range):
            request.resource.source_range = Primitive.to_proto(self.source_range)

        if Primitive.to_proto(self.app):
            request.resource.app = Primitive.to_proto(self.app)

        response = stub.DeleteAppengineFirewallRule(request)

    @classmethod
    def list(self, app, service_account_file=""):
        stub = firewall_rule_pb2_grpc.AppengineFirewallRuleServiceStub(
            channel.Channel()
        )
        request = firewall_rule_pb2.ListAppengineFirewallRuleRequest()
        request.service_account_file = service_account_file
        request.App = app

        return stub.ListAppengineFirewallRule(request).items

    def to_proto(self):
        resource = firewall_rule_pb2.AppengineFirewallRule()
        if FirewallRuleActionEnum.to_proto(self.action):
            resource.action = FirewallRuleActionEnum.to_proto(self.action)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.priority):
            resource.priority = Primitive.to_proto(self.priority)
        if Primitive.to_proto(self.source_range):
            resource.source_range = Primitive.to_proto(self.source_range)
        if Primitive.to_proto(self.app):
            resource.app = Primitive.to_proto(self.app)
        return resource


class FirewallRuleActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return firewall_rule_pb2.AppengineFirewallRuleActionEnum.Value(
            "AppengineFirewallRuleActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return firewall_rule_pb2.AppengineFirewallRuleActionEnum.Name(resource)[
            len("AppengineFirewallRuleActionEnum") :
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
