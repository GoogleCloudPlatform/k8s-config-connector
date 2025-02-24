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
from google3.cloud.graphite.mmv2.services.google.logging import log_exclusion_pb2
from google3.cloud.graphite.mmv2.services.google.logging import log_exclusion_pb2_grpc

from typing import List


class LogExclusion(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        filter: str = None,
        disabled: bool = None,
        create_time: str = None,
        update_time: str = None,
        parent: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.filter = filter
        self.disabled = disabled
        self.parent = parent
        self.service_account_file = service_account_file

    def apply(self):
        stub = log_exclusion_pb2_grpc.LoggingLogExclusionServiceStub(channel.Channel())
        request = log_exclusion_pb2.ApplyLoggingLogExclusionRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.filter):
            request.resource.filter = Primitive.to_proto(self.filter)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        request.service_account_file = self.service_account_file

        response = stub.ApplyLoggingLogExclusion(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.filter = Primitive.from_proto(response.filter)
        self.disabled = Primitive.from_proto(response.disabled)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.parent = Primitive.from_proto(response.parent)

    def delete(self):
        stub = log_exclusion_pb2_grpc.LoggingLogExclusionServiceStub(channel.Channel())
        request = log_exclusion_pb2.DeleteLoggingLogExclusionRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.filter):
            request.resource.filter = Primitive.to_proto(self.filter)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        response = stub.DeleteLoggingLogExclusion(request)

    @classmethod
    def list(self, parent, service_account_file=""):
        stub = log_exclusion_pb2_grpc.LoggingLogExclusionServiceStub(channel.Channel())
        request = log_exclusion_pb2.ListLoggingLogExclusionRequest()
        request.service_account_file = service_account_file
        request.Parent = parent

        return stub.ListLoggingLogExclusion(request).items

    def to_proto(self):
        resource = log_exclusion_pb2.LoggingLogExclusion()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.filter):
            resource.filter = Primitive.to_proto(self.filter)
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
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
