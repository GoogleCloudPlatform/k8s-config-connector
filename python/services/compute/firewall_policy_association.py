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
    firewall_policy_association_pb2,
)
from google3.cloud.graphite.mmv2.services.google.compute import (
    firewall_policy_association_pb2_grpc,
)

from typing import List


class FirewallPolicyAssociation(object):
    def __init__(
        self,
        name: str = None,
        attachment_target: str = None,
        firewall_policy: str = None,
        short_name: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.attachment_target = attachment_target
        self.firewall_policy = firewall_policy
        self.service_account_file = service_account_file

    def apply(self):
        stub = firewall_policy_association_pb2_grpc.ComputeFirewallPolicyAssociationServiceStub(
            channel.Channel()
        )
        request = (
            firewall_policy_association_pb2.ApplyComputeFirewallPolicyAssociationRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.attachment_target):
            request.resource.attachment_target = Primitive.to_proto(
                self.attachment_target
            )

        if Primitive.to_proto(self.firewall_policy):
            request.resource.firewall_policy = Primitive.to_proto(self.firewall_policy)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeFirewallPolicyAssociation(request)
        self.name = Primitive.from_proto(response.name)
        self.attachment_target = Primitive.from_proto(response.attachment_target)
        self.firewall_policy = Primitive.from_proto(response.firewall_policy)
        self.short_name = Primitive.from_proto(response.short_name)

    def delete(self):
        stub = firewall_policy_association_pb2_grpc.ComputeFirewallPolicyAssociationServiceStub(
            channel.Channel()
        )
        request = (
            firewall_policy_association_pb2.DeleteComputeFirewallPolicyAssociationRequest()
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

        response = stub.DeleteComputeFirewallPolicyAssociation(request)

    @classmethod
    def list(self, firewallPolicy, service_account_file=""):
        stub = firewall_policy_association_pb2_grpc.ComputeFirewallPolicyAssociationServiceStub(
            channel.Channel()
        )
        request = (
            firewall_policy_association_pb2.ListComputeFirewallPolicyAssociationRequest()
        )
        request.service_account_file = service_account_file
        request.FirewallPolicy = firewallPolicy

        return stub.ListComputeFirewallPolicyAssociation(request).items

    def to_proto(self):
        resource = firewall_policy_association_pb2.ComputeFirewallPolicyAssociation()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.attachment_target):
            resource.attachment_target = Primitive.to_proto(self.attachment_target)
        if Primitive.to_proto(self.firewall_policy):
            resource.firewall_policy = Primitive.to_proto(self.firewall_policy)
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
