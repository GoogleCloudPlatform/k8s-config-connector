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
from google3.cloud.graphite.mmv2.services.google.compute import firewall_policy_pb2
from google3.cloud.graphite.mmv2.services.google.compute import firewall_policy_pb2_grpc

from typing import List


class FirewallPolicy(object):
    def __init__(
        self,
        name: str = None,
        id: str = None,
        creation_timestamp: str = None,
        description: str = None,
        fingerprint: str = None,
        self_link: str = None,
        self_link_with_id: str = None,
        rule_tuple_count: int = None,
        short_name: str = None,
        parent: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.short_name = short_name
        self.parent = parent
        self.service_account_file = service_account_file

    def apply(self):
        stub = firewall_policy_pb2_grpc.ComputeFirewallPolicyServiceStub(
            channel.Channel()
        )
        request = firewall_policy_pb2.ApplyComputeFirewallPolicyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.short_name):
            request.resource.short_name = Primitive.to_proto(self.short_name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeFirewallPolicy(request)
        self.name = Primitive.from_proto(response.name)
        self.id = Primitive.from_proto(response.id)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.description = Primitive.from_proto(response.description)
        self.fingerprint = Primitive.from_proto(response.fingerprint)
        self.self_link = Primitive.from_proto(response.self_link)
        self.self_link_with_id = Primitive.from_proto(response.self_link_with_id)
        self.rule_tuple_count = Primitive.from_proto(response.rule_tuple_count)
        self.short_name = Primitive.from_proto(response.short_name)
        self.parent = Primitive.from_proto(response.parent)

    def delete(self):
        stub = firewall_policy_pb2_grpc.ComputeFirewallPolicyServiceStub(
            channel.Channel()
        )
        request = firewall_policy_pb2.DeleteComputeFirewallPolicyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.short_name):
            request.resource.short_name = Primitive.to_proto(self.short_name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        response = stub.DeleteComputeFirewallPolicy(request)

    @classmethod
    def list(self, parent, service_account_file=""):
        stub = firewall_policy_pb2_grpc.ComputeFirewallPolicyServiceStub(
            channel.Channel()
        )
        request = firewall_policy_pb2.ListComputeFirewallPolicyRequest()
        request.service_account_file = service_account_file
        request.Parent = parent

        return stub.ListComputeFirewallPolicy(request).items

    def to_proto(self):
        resource = firewall_policy_pb2.ComputeFirewallPolicy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.short_name):
            resource.short_name = Primitive.to_proto(self.short_name)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
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
