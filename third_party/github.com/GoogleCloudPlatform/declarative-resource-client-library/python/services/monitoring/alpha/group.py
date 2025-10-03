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
from google3.cloud.graphite.mmv2.services.google.monitoring import group_pb2
from google3.cloud.graphite.mmv2.services.google.monitoring import group_pb2_grpc

from typing import List


class Group(object):
    def __init__(
        self,
        display_name: str = None,
        filter: str = None,
        is_cluster: bool = None,
        name: str = None,
        parent_name: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.display_name = display_name
        self.filter = filter
        self.is_cluster = is_cluster
        self.name = name
        self.parent_name = parent_name
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = group_pb2_grpc.MonitoringAlphaGroupServiceStub(channel.Channel())
        request = group_pb2.ApplyMonitoringAlphaGroupRequest()
        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.filter):
            request.resource.filter = Primitive.to_proto(self.filter)

        if Primitive.to_proto(self.is_cluster):
            request.resource.is_cluster = Primitive.to_proto(self.is_cluster)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.parent_name):
            request.resource.parent_name = Primitive.to_proto(self.parent_name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringAlphaGroup(request)
        self.display_name = Primitive.from_proto(response.display_name)
        self.filter = Primitive.from_proto(response.filter)
        self.is_cluster = Primitive.from_proto(response.is_cluster)
        self.name = Primitive.from_proto(response.name)
        self.parent_name = Primitive.from_proto(response.parent_name)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = group_pb2_grpc.MonitoringAlphaGroupServiceStub(channel.Channel())
        request = group_pb2.DeleteMonitoringAlphaGroupRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.filter):
            request.resource.filter = Primitive.to_proto(self.filter)

        if Primitive.to_proto(self.is_cluster):
            request.resource.is_cluster = Primitive.to_proto(self.is_cluster)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.parent_name):
            request.resource.parent_name = Primitive.to_proto(self.parent_name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteMonitoringAlphaGroup(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = group_pb2_grpc.MonitoringAlphaGroupServiceStub(channel.Channel())
        request = group_pb2.ListMonitoringAlphaGroupRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListMonitoringAlphaGroup(request).items

    def to_proto(self):
        resource = group_pb2.MonitoringAlphaGroup()
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.filter):
            resource.filter = Primitive.to_proto(self.filter)
        if Primitive.to_proto(self.is_cluster):
            resource.is_cluster = Primitive.to_proto(self.is_cluster)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.parent_name):
            resource.parent_name = Primitive.to_proto(self.parent_name)
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
