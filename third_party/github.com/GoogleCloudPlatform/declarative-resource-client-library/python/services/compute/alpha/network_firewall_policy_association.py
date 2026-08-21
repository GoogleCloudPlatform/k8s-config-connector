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
    network_firewall_policy_association_pb2,
)
from google3.cloud.graphite.mmv2.services.google.compute import (
    network_firewall_policy_association_pb2_grpc,
)

from typing import List


class NetworkFirewallPolicyAssociation(object):
    def __init__(
        self,
        name: str = None,
        attachment_target: str = None,
        firewall_policy: str = None,
        short_name: str = None,
        location: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.attachment_target = attachment_target
        self.firewall_policy = firewall_policy
        self.location = location
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = network_firewall_policy_association_pb2_grpc.ComputeAlphaNetworkFirewallPolicyAssociationServiceStub(
            channel.Channel()
        )
        request = (
            network_firewall_policy_association_pb2.ApplyComputeAlphaNetworkFirewallPolicyAssociationRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.attachment_target):
            request.resource.attachment_target = Primitive.to_proto(
                self.attachment_target
            )

        if Primitive.to_proto(self.firewall_policy):
            request.resource.firewall_policy = Primitive.to_proto(self.firewall_policy)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeAlphaNetworkFirewallPolicyAssociation(request)
        self.name = Primitive.from_proto(response.name)
        self.attachment_target = Primitive.from_proto(response.attachment_target)
        self.firewall_policy = Primitive.from_proto(response.firewall_policy)
        self.short_name = Primitive.from_proto(response.short_name)
        self.location = Primitive.from_proto(response.location)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = network_firewall_policy_association_pb2_grpc.ComputeAlphaNetworkFirewallPolicyAssociationServiceStub(
            channel.Channel()
        )
        request = (
            network_firewall_policy_association_pb2.DeleteComputeAlphaNetworkFirewallPolicyAssociationRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.attachment_target):
            request.resource.attachment_target = Primitive.to_proto(
                self.attachment_target
            )

        if Primitive.to_proto(self.firewall_policy):
            request.resource.firewall_policy = Primitive.to_proto(self.firewall_policy)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeAlphaNetworkFirewallPolicyAssociation(request)

    @classmethod
    def list(self, project, location, firewallPolicy, service_account_file=""):
        stub = network_firewall_policy_association_pb2_grpc.ComputeAlphaNetworkFirewallPolicyAssociationServiceStub(
            channel.Channel()
        )
        request = (
            network_firewall_policy_association_pb2.ListComputeAlphaNetworkFirewallPolicyAssociationRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.FirewallPolicy = firewallPolicy

        return stub.ListComputeAlphaNetworkFirewallPolicyAssociation(request).items

    def to_proto(self):
        resource = (
            network_firewall_policy_association_pb2.ComputeAlphaNetworkFirewallPolicyAssociation()
        )
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.attachment_target):
            resource.attachment_target = Primitive.to_proto(self.attachment_target)
        if Primitive.to_proto(self.firewall_policy):
            resource.firewall_policy = Primitive.to_proto(self.firewall_policy)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
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
