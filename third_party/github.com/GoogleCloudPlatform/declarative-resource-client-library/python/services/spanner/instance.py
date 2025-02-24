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
from google3.cloud.graphite.mmv2.services.google.spanner import instance_pb2
from google3.cloud.graphite.mmv2.services.google.spanner import instance_pb2_grpc

from typing import List


class Instance(object):
    def __init__(
        self,
        name: str = None,
        project: str = None,
        config: str = None,
        display_name: str = None,
        node_count: int = None,
        state: str = None,
        labels: dict = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.project = project
        self.config = config
        self.display_name = display_name
        self.node_count = node_count
        self.labels = labels
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_pb2_grpc.SpannerInstanceServiceStub(channel.Channel())
        request = instance_pb2.ApplySpannerInstanceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.config):
            request.resource.config = Primitive.to_proto(self.config)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.node_count):
            request.resource.node_count = Primitive.to_proto(self.node_count)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        request.service_account_file = self.service_account_file

        response = stub.ApplySpannerInstance(request)
        self.name = Primitive.from_proto(response.name)
        self.project = Primitive.from_proto(response.project)
        self.config = Primitive.from_proto(response.config)
        self.display_name = Primitive.from_proto(response.display_name)
        self.node_count = Primitive.from_proto(response.node_count)
        self.state = InstanceStateEnum.from_proto(response.state)
        self.labels = Primitive.from_proto(response.labels)

    def delete(self):
        stub = instance_pb2_grpc.SpannerInstanceServiceStub(channel.Channel())
        request = instance_pb2.DeleteSpannerInstanceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.config):
            request.resource.config = Primitive.to_proto(self.config)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.node_count):
            request.resource.node_count = Primitive.to_proto(self.node_count)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        response = stub.DeleteSpannerInstance(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = instance_pb2_grpc.SpannerInstanceServiceStub(channel.Channel())
        request = instance_pb2.ListSpannerInstanceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListSpannerInstance(request).items

    def to_proto(self):
        resource = instance_pb2.SpannerInstance()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.config):
            resource.config = Primitive.to_proto(self.config)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.node_count):
            resource.node_count = Primitive.to_proto(self.node_count)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        return resource


class InstanceStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SpannerInstanceStateEnum.Value(
            "SpannerInstanceStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SpannerInstanceStateEnum.Name(resource)[
            len("SpannerInstanceStateEnum") :
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
