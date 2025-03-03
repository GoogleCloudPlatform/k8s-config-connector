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
from google3.cloud.graphite.mmv2.services.google.access_context_manager import (
    access_policy_pb2,
)
from google3.cloud.graphite.mmv2.services.google.access_context_manager import (
    access_policy_pb2_grpc,
)

from typing import List


class AccessPolicy(object):
    def __init__(
        self,
        name: str = None,
        parent: str = None,
        title: str = None,
        create_time: str = None,
        update_time: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.parent = parent
        self.title = title
        self.service_account_file = service_account_file

    def apply(self):
        stub = access_policy_pb2_grpc.AccesscontextmanagerAccessPolicyServiceStub(
            channel.Channel()
        )
        request = access_policy_pb2.ApplyAccesscontextmanagerAccessPolicyRequest()
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.title):
            request.resource.title = Primitive.to_proto(self.title)

        request.service_account_file = self.service_account_file

        response = stub.ApplyAccesscontextmanagerAccessPolicy(request)
        self.name = Primitive.from_proto(response.name)
        self.parent = Primitive.from_proto(response.parent)
        self.title = Primitive.from_proto(response.title)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)

    def delete(self):
        stub = access_policy_pb2_grpc.AccesscontextmanagerAccessPolicyServiceStub(
            channel.Channel()
        )
        request = access_policy_pb2.DeleteAccesscontextmanagerAccessPolicyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.title):
            request.resource.title = Primitive.to_proto(self.title)

        response = stub.DeleteAccesscontextmanagerAccessPolicy(request)

    @classmethod
    def list(self, parent, service_account_file=""):
        stub = access_policy_pb2_grpc.AccesscontextmanagerAccessPolicyServiceStub(
            channel.Channel()
        )
        request = access_policy_pb2.ListAccesscontextmanagerAccessPolicyRequest()
        request.service_account_file = service_account_file
        request.Parent = parent

        return stub.ListAccesscontextmanagerAccessPolicy(request).items

    def to_proto(self):
        resource = access_policy_pb2.AccesscontextmanagerAccessPolicy()
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.title):
            resource.title = Primitive.to_proto(self.title)
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
