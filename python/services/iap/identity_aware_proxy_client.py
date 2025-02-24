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
from google3.cloud.graphite.mmv2.services.google.iap import (
    identity_aware_proxy_client_pb2,
)
from google3.cloud.graphite.mmv2.services.google.iap import (
    identity_aware_proxy_client_pb2_grpc,
)

from typing import List


class IdentityAwareProxyClient(object):
    def __init__(
        self,
        name: str = None,
        secret: str = None,
        display_name: str = None,
        project: str = None,
        brand: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.project = project
        self.brand = brand
        self.service_account_file = service_account_file

    def apply(self):
        stub = (
            identity_aware_proxy_client_pb2_grpc.IapIdentityAwareProxyClientServiceStub(
                channel.Channel()
            )
        )
        request = (
            identity_aware_proxy_client_pb2.ApplyIapIdentityAwareProxyClientRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.brand):
            request.resource.brand = Primitive.to_proto(self.brand)

        request.service_account_file = self.service_account_file

        response = stub.ApplyIapIdentityAwareProxyClient(request)
        self.name = Primitive.from_proto(response.name)
        self.secret = Primitive.from_proto(response.secret)
        self.display_name = Primitive.from_proto(response.display_name)
        self.project = Primitive.from_proto(response.project)
        self.brand = Primitive.from_proto(response.brand)

    def delete(self):
        stub = (
            identity_aware_proxy_client_pb2_grpc.IapIdentityAwareProxyClientServiceStub(
                channel.Channel()
            )
        )
        request = (
            identity_aware_proxy_client_pb2.DeleteIapIdentityAwareProxyClientRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.brand):
            request.resource.brand = Primitive.to_proto(self.brand)

        response = stub.DeleteIapIdentityAwareProxyClient(request)

    @classmethod
    def list(self, project, brand, service_account_file=""):
        stub = (
            identity_aware_proxy_client_pb2_grpc.IapIdentityAwareProxyClientServiceStub(
                channel.Channel()
            )
        )
        request = (
            identity_aware_proxy_client_pb2.ListIapIdentityAwareProxyClientRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Brand = brand

        return stub.ListIapIdentityAwareProxyClient(request).items

    def to_proto(self):
        resource = identity_aware_proxy_client_pb2.IapIdentityAwareProxyClient()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.brand):
            resource.brand = Primitive.to_proto(self.brand)
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
