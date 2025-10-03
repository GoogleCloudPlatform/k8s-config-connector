# Copyright 2022 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.apigee import attachment_pb2
from google3.cloud.graphite.mmv2.services.google.apigee import attachment_pb2_grpc

from typing import List


class Attachment(object):
    def __init__(
        self,
        name: str = None,
        environment: str = None,
        created_at: int = None,
        envgroup: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.environment = environment
        self.envgroup = envgroup
        self.service_account_file = service_account_file

    def apply(self):
        stub = attachment_pb2_grpc.ApigeeBetaAttachmentServiceStub(channel.Channel())
        request = attachment_pb2.ApplyApigeeBetaAttachmentRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.environment):
            request.resource.environment = Primitive.to_proto(self.environment)

        if Primitive.to_proto(self.envgroup):
            request.resource.envgroup = Primitive.to_proto(self.envgroup)

        request.service_account_file = self.service_account_file

        response = stub.ApplyApigeeBetaAttachment(request)
        self.name = Primitive.from_proto(response.name)
        self.environment = Primitive.from_proto(response.environment)
        self.created_at = Primitive.from_proto(response.created_at)
        self.envgroup = Primitive.from_proto(response.envgroup)

    def delete(self):
        stub = attachment_pb2_grpc.ApigeeBetaAttachmentServiceStub(channel.Channel())
        request = attachment_pb2.DeleteApigeeBetaAttachmentRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.environment):
            request.resource.environment = Primitive.to_proto(self.environment)

        if Primitive.to_proto(self.envgroup):
            request.resource.envgroup = Primitive.to_proto(self.envgroup)

        response = stub.DeleteApigeeBetaAttachment(request)

    @classmethod
    def list(self, envgroup, service_account_file=""):
        stub = attachment_pb2_grpc.ApigeeBetaAttachmentServiceStub(channel.Channel())
        request = attachment_pb2.ListApigeeBetaAttachmentRequest()
        request.service_account_file = service_account_file
        request.Envgroup = envgroup

        return stub.ListApigeeBetaAttachment(request).items

    def to_proto(self):
        resource = attachment_pb2.ApigeeBetaAttachment()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.environment):
            resource.environment = Primitive.to_proto(self.environment)
        if Primitive.to_proto(self.envgroup):
            resource.envgroup = Primitive.to_proto(self.envgroup)
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
