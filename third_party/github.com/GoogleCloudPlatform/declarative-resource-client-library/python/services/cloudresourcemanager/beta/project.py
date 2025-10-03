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
from google3.cloud.graphite.mmv2.services.google.cloud_resource_manager import (
    project_pb2,
)
from google3.cloud.graphite.mmv2.services.google.cloud_resource_manager import (
    project_pb2_grpc,
)

from typing import List


class Project(object):
    def __init__(
        self,
        labels: dict = None,
        lifecycle_state: str = None,
        displayName: str = None,
        parent: str = None,
        name: str = None,
        project_number: int = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.labels = labels
        self.displayName = displayName
        self.parent = parent
        self.name = name
        self.service_account_file = service_account_file

    def apply(self):
        stub = project_pb2_grpc.CloudresourcemanagerBetaProjectServiceStub(
            channel.Channel()
        )
        request = project_pb2.ApplyCloudresourcemanagerBetaProjectRequest()
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.displayName):
            request.resource.displayName = Primitive.to_proto(self.displayName)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudresourcemanagerBetaProject(request)
        self.labels = Primitive.from_proto(response.labels)
        self.lifecycle_state = ProjectLifecycleStateEnum.from_proto(
            response.lifecycle_state
        )
        self.displayName = Primitive.from_proto(response.displayName)
        self.parent = Primitive.from_proto(response.parent)
        self.name = Primitive.from_proto(response.name)
        self.project_number = Primitive.from_proto(response.project_number)

    def delete(self):
        stub = project_pb2_grpc.CloudresourcemanagerBetaProjectServiceStub(
            channel.Channel()
        )
        request = project_pb2.DeleteCloudresourcemanagerBetaProjectRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.displayName):
            request.resource.displayName = Primitive.to_proto(self.displayName)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        response = stub.DeleteCloudresourcemanagerBetaProject(request)

    @classmethod
    def list(self, parent, service_account_file=""):
        stub = project_pb2_grpc.CloudresourcemanagerBetaProjectServiceStub(
            channel.Channel()
        )
        request = project_pb2.ListCloudresourcemanagerBetaProjectRequest()
        request.service_account_file = service_account_file
        request.Parent = parent

        return stub.ListCloudresourcemanagerBetaProject(request).items

    def to_proto(self):
        resource = project_pb2.CloudresourcemanagerBetaProject()
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.displayName):
            resource.displayName = Primitive.to_proto(self.displayName)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        return resource


class ProjectLifecycleStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return project_pb2.CloudresourcemanagerBetaProjectLifecycleStateEnum.Value(
            "CloudresourcemanagerBetaProjectLifecycleStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return project_pb2.CloudresourcemanagerBetaProjectLifecycleStateEnum.Name(
            resource
        )[len("CloudresourcemanagerBetaProjectLifecycleStateEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
