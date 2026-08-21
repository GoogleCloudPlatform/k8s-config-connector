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
    folder_pb2,
)
from google3.cloud.graphite.mmv2.services.google.cloud_resource_manager import (
    folder_pb2_grpc,
)

from typing import List


class Folder(object):
    def __init__(
        self,
        name: str = None,
        parent: str = None,
        display_name: str = None,
        state: str = None,
        create_time: str = None,
        update_time: str = None,
        delete_time: str = None,
        etag: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.parent = parent
        self.display_name = display_name
        self.service_account_file = service_account_file

    def apply(self):
        stub = folder_pb2_grpc.CloudresourcemanagerBetaFolderServiceStub(
            channel.Channel()
        )
        request = folder_pb2.ApplyCloudresourcemanagerBetaFolderRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudresourcemanagerBetaFolder(request)
        self.name = Primitive.from_proto(response.name)
        self.parent = Primitive.from_proto(response.parent)
        self.display_name = Primitive.from_proto(response.display_name)
        self.state = FolderStateEnum.from_proto(response.state)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.etag = Primitive.from_proto(response.etag)

    def delete(self):
        stub = folder_pb2_grpc.CloudresourcemanagerBetaFolderServiceStub(
            channel.Channel()
        )
        request = folder_pb2.DeleteCloudresourcemanagerBetaFolderRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        response = stub.DeleteCloudresourcemanagerBetaFolder(request)

    @classmethod
    def list(self, parent, service_account_file=""):
        stub = folder_pb2_grpc.CloudresourcemanagerBetaFolderServiceStub(
            channel.Channel()
        )
        request = folder_pb2.ListCloudresourcemanagerBetaFolderRequest()
        request.service_account_file = service_account_file
        request.Parent = parent

        return stub.ListCloudresourcemanagerBetaFolder(request).items

    def to_proto(self):
        resource = folder_pb2.CloudresourcemanagerBetaFolder()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        return resource


class FolderStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return folder_pb2.CloudresourcemanagerBetaFolderStateEnum.Value(
            "CloudresourcemanagerBetaFolderStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return folder_pb2.CloudresourcemanagerBetaFolderStateEnum.Name(resource)[
            len("CloudresourcemanagerBetaFolderStateEnum") :
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
