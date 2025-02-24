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
    tag_value_pb2,
)
from google3.cloud.graphite.mmv2.services.google.cloud_resource_manager import (
    tag_value_pb2_grpc,
)

from typing import List


class TagValue(object):
    def __init__(
        self,
        name: str = None,
        parent: str = None,
        short_name: str = None,
        namespaced_name: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        etag: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.parent = parent
        self.short_name = short_name
        self.description = description
        self.service_account_file = service_account_file

    def apply(self):
        stub = tag_value_pb2_grpc.CloudresourcemanagerBetaTagValueServiceStub(
            channel.Channel()
        )
        request = tag_value_pb2.ApplyCloudresourcemanagerBetaTagValueRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.short_name):
            request.resource.short_name = Primitive.to_proto(self.short_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudresourcemanagerBetaTagValue(request)
        self.name = Primitive.from_proto(response.name)
        self.parent = Primitive.from_proto(response.parent)
        self.short_name = Primitive.from_proto(response.short_name)
        self.namespaced_name = Primitive.from_proto(response.namespaced_name)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.etag = Primitive.from_proto(response.etag)

    def delete(self):
        stub = tag_value_pb2_grpc.CloudresourcemanagerBetaTagValueServiceStub(
            channel.Channel()
        )
        request = tag_value_pb2.DeleteCloudresourcemanagerBetaTagValueRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.short_name):
            request.resource.short_name = Primitive.to_proto(self.short_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        response = stub.DeleteCloudresourcemanagerBetaTagValue(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = tag_value_pb2_grpc.CloudresourcemanagerBetaTagValueServiceStub(
            channel.Channel()
        )
        request = tag_value_pb2.ListCloudresourcemanagerBetaTagValueRequest()
        request.service_account_file = service_account_file
        return stub.ListCloudresourcemanagerBetaTagValue(request).items

    def to_proto(self):
        resource = tag_value_pb2.CloudresourcemanagerBetaTagValue()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.short_name):
            resource.short_name = Primitive.to_proto(self.short_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
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
