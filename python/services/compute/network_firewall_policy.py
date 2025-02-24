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
    network_firewall_policy_pb2,
)
from google3.cloud.graphite.mmv2.services.google.compute import (
    network_firewall_policy_pb2_grpc,
)

from typing import List


class NetworkFirewallPolicy(object):
    def __init__(
        self,
        location: str = None,
        creation_timestamp: str = None,
        name: str = None,
        id: str = None,
        description: str = None,
        fingerprint: str = None,
        self_link: str = None,
        self_link_with_id: str = None,
        rule_tuple_count: int = None,
        region: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.location = location
        self.name = name
        self.description = description
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = network_firewall_policy_pb2_grpc.ComputeNetworkFirewallPolicyServiceStub(
            channel.Channel()
        )
        request = network_firewall_policy_pb2.ApplyComputeNetworkFirewallPolicyRequest()
        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeNetworkFirewallPolicy(request)
        self.location = Primitive.from_proto(response.location)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.name = Primitive.from_proto(response.name)
        self.id = Primitive.from_proto(response.id)
        self.description = Primitive.from_proto(response.description)
        self.fingerprint = Primitive.from_proto(response.fingerprint)
        self.self_link = Primitive.from_proto(response.self_link)
        self.self_link_with_id = Primitive.from_proto(response.self_link_with_id)
        self.rule_tuple_count = Primitive.from_proto(response.rule_tuple_count)
        self.region = Primitive.from_proto(response.region)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = network_firewall_policy_pb2_grpc.ComputeNetworkFirewallPolicyServiceStub(
            channel.Channel()
        )
        request = (
            network_firewall_policy_pb2.DeleteComputeNetworkFirewallPolicyRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeNetworkFirewallPolicy(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = network_firewall_policy_pb2_grpc.ComputeNetworkFirewallPolicyServiceStub(
            channel.Channel()
        )
        request = network_firewall_policy_pb2.ListComputeNetworkFirewallPolicyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeNetworkFirewallPolicy(request).items

    def to_proto(self):
        resource = network_firewall_policy_pb2.ComputeNetworkFirewallPolicy()
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
