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
from google3.cloud.graphite.mmv2.services.google.compute import target_http_proxy_pb2
from google3.cloud.graphite.mmv2.services.google.compute import (
    target_http_proxy_pb2_grpc,
)

from typing import List


class TargetHttpProxy(object):
    def __init__(
        self,
        id: int = None,
        name: str = None,
        description: str = None,
        self_link: str = None,
        url_map: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.url_map = url_map
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = target_http_proxy_pb2_grpc.ComputeBetaTargetHttpProxyServiceStub(
            channel.Channel()
        )
        request = target_http_proxy_pb2.ApplyComputeBetaTargetHttpProxyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.url_map):
            request.resource.url_map = Primitive.to_proto(self.url_map)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaTargetHttpProxy(request)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.self_link = Primitive.from_proto(response.self_link)
        self.url_map = Primitive.from_proto(response.url_map)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = target_http_proxy_pb2_grpc.ComputeBetaTargetHttpProxyServiceStub(
            channel.Channel()
        )
        request = target_http_proxy_pb2.DeleteComputeBetaTargetHttpProxyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.url_map):
            request.resource.url_map = Primitive.to_proto(self.url_map)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeBetaTargetHttpProxy(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = target_http_proxy_pb2_grpc.ComputeBetaTargetHttpProxyServiceStub(
            channel.Channel()
        )
        request = target_http_proxy_pb2.ListComputeBetaTargetHttpProxyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeBetaTargetHttpProxy(request).items

    def to_proto(self):
        resource = target_http_proxy_pb2.ComputeBetaTargetHttpProxy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.url_map):
            resource.url_map = Primitive.to_proto(self.url_map)
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
