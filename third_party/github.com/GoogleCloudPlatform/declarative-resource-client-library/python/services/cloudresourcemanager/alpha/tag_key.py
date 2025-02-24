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
    tag_key_pb2,
)
from google3.cloud.graphite.mmv2.services.google.cloud_resource_manager import (
    tag_key_pb2_grpc,
)

from typing import List


class TagKey(object):
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
        purpose: str = None,
        purpose_data: dict = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.parent = parent
        self.short_name = short_name
        self.description = description
        self.purpose = purpose
        self.purpose_data = purpose_data
        self.service_account_file = service_account_file

    def apply(self):
        stub = tag_key_pb2_grpc.CloudresourcemanagerAlphaTagKeyServiceStub(
            channel.Channel()
        )
        request = tag_key_pb2.ApplyCloudresourcemanagerAlphaTagKeyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.short_name):
            request.resource.short_name = Primitive.to_proto(self.short_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if TagKeyPurposeEnum.to_proto(self.purpose):
            request.resource.purpose = TagKeyPurposeEnum.to_proto(self.purpose)

        if Primitive.to_proto(self.purpose_data):
            request.resource.purpose_data = Primitive.to_proto(self.purpose_data)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudresourcemanagerAlphaTagKey(request)
        self.name = Primitive.from_proto(response.name)
        self.parent = Primitive.from_proto(response.parent)
        self.short_name = Primitive.from_proto(response.short_name)
        self.namespaced_name = Primitive.from_proto(response.namespaced_name)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.etag = Primitive.from_proto(response.etag)
        self.purpose = TagKeyPurposeEnum.from_proto(response.purpose)
        self.purpose_data = Primitive.from_proto(response.purpose_data)

    def delete(self):
        stub = tag_key_pb2_grpc.CloudresourcemanagerAlphaTagKeyServiceStub(
            channel.Channel()
        )
        request = tag_key_pb2.DeleteCloudresourcemanagerAlphaTagKeyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.short_name):
            request.resource.short_name = Primitive.to_proto(self.short_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if TagKeyPurposeEnum.to_proto(self.purpose):
            request.resource.purpose = TagKeyPurposeEnum.to_proto(self.purpose)

        if Primitive.to_proto(self.purpose_data):
            request.resource.purpose_data = Primitive.to_proto(self.purpose_data)

        response = stub.DeleteCloudresourcemanagerAlphaTagKey(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = tag_key_pb2_grpc.CloudresourcemanagerAlphaTagKeyServiceStub(
            channel.Channel()
        )
        request = tag_key_pb2.ListCloudresourcemanagerAlphaTagKeyRequest()
        request.service_account_file = service_account_file
        return stub.ListCloudresourcemanagerAlphaTagKey(request).items

    def to_proto(self):
        resource = tag_key_pb2.CloudresourcemanagerAlphaTagKey()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.short_name):
            resource.short_name = Primitive.to_proto(self.short_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if TagKeyPurposeEnum.to_proto(self.purpose):
            resource.purpose = TagKeyPurposeEnum.to_proto(self.purpose)
        if Primitive.to_proto(self.purpose_data):
            resource.purpose_data = Primitive.to_proto(self.purpose_data)
        return resource


class TagKeyPurposeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return tag_key_pb2.CloudresourcemanagerAlphaTagKeyPurposeEnum.Value(
            "CloudresourcemanagerAlphaTagKeyPurposeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return tag_key_pb2.CloudresourcemanagerAlphaTagKeyPurposeEnum.Name(resource)[
            len("CloudresourcemanagerAlphaTagKeyPurposeEnum") :
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
