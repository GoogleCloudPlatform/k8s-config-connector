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
from google3.cloud.graphite.mmv2.services.google.storage import hmac_key_pb2
from google3.cloud.graphite.mmv2.services.google.storage import hmac_key_pb2_grpc

from typing import List


class HmacKey(object):
    def __init__(
        self,
        name: str = None,
        time_created: str = None,
        updated: str = None,
        secret: str = None,
        state: str = None,
        project: str = None,
        service_account_email: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.state = state
        self.project = project
        self.service_account_email = service_account_email
        self.service_account_file = service_account_file

    def apply(self):
        stub = hmac_key_pb2_grpc.StorageHmacKeyServiceStub(channel.Channel())
        request = hmac_key_pb2.ApplyStorageHmacKeyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if HmacKeyStateEnum.to_proto(self.state):
            request.resource.state = HmacKeyStateEnum.to_proto(self.state)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.service_account_email):
            request.resource.service_account_email = Primitive.to_proto(
                self.service_account_email
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyStorageHmacKey(request)
        self.name = Primitive.from_proto(response.name)
        self.time_created = Primitive.from_proto(response.time_created)
        self.updated = Primitive.from_proto(response.updated)
        self.secret = Primitive.from_proto(response.secret)
        self.state = HmacKeyStateEnum.from_proto(response.state)
        self.project = Primitive.from_proto(response.project)
        self.service_account_email = Primitive.from_proto(
            response.service_account_email
        )

    def delete(self):
        stub = hmac_key_pb2_grpc.StorageHmacKeyServiceStub(channel.Channel())
        request = hmac_key_pb2.DeleteStorageHmacKeyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if HmacKeyStateEnum.to_proto(self.state):
            request.resource.state = HmacKeyStateEnum.to_proto(self.state)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.service_account_email):
            request.resource.service_account_email = Primitive.to_proto(
                self.service_account_email
            )

        response = stub.DeleteStorageHmacKey(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = hmac_key_pb2_grpc.StorageHmacKeyServiceStub(channel.Channel())
        request = hmac_key_pb2.ListStorageHmacKeyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListStorageHmacKey(request).items

    def to_proto(self):
        resource = hmac_key_pb2.StorageHmacKey()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if HmacKeyStateEnum.to_proto(self.state):
            resource.state = HmacKeyStateEnum.to_proto(self.state)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.service_account_email):
            resource.service_account_email = Primitive.to_proto(
                self.service_account_email
            )
        return resource


class HmacKeyStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return hmac_key_pb2.StorageHmacKeyStateEnum.Value(
            "StorageHmacKeyStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return hmac_key_pb2.StorageHmacKeyStateEnum.Name(resource)[
            len("StorageHmacKeyStateEnum") :
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
