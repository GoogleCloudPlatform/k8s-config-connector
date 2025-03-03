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
from google3.cloud.graphite.mmv2.services.google.logging import log_view_pb2
from google3.cloud.graphite.mmv2.services.google.logging import log_view_pb2_grpc

from typing import List


class LogView(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        filter: str = None,
        parent: str = None,
        location: str = None,
        bucket: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.filter = filter
        self.parent = parent
        self.location = location
        self.bucket = bucket
        self.service_account_file = service_account_file

    def apply(self):
        stub = log_view_pb2_grpc.LoggingLogViewServiceStub(channel.Channel())
        request = log_view_pb2.ApplyLoggingLogViewRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.filter):
            request.resource.filter = Primitive.to_proto(self.filter)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.bucket):
            request.resource.bucket = Primitive.to_proto(self.bucket)

        request.service_account_file = self.service_account_file

        response = stub.ApplyLoggingLogView(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.filter = Primitive.from_proto(response.filter)
        self.parent = Primitive.from_proto(response.parent)
        self.location = Primitive.from_proto(response.location)
        self.bucket = Primitive.from_proto(response.bucket)

    def delete(self):
        stub = log_view_pb2_grpc.LoggingLogViewServiceStub(channel.Channel())
        request = log_view_pb2.DeleteLoggingLogViewRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.filter):
            request.resource.filter = Primitive.to_proto(self.filter)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.bucket):
            request.resource.bucket = Primitive.to_proto(self.bucket)

        response = stub.DeleteLoggingLogView(request)

    @classmethod
    def list(self, location, bucket, parent, service_account_file=""):
        stub = log_view_pb2_grpc.LoggingLogViewServiceStub(channel.Channel())
        request = log_view_pb2.ListLoggingLogViewRequest()
        request.service_account_file = service_account_file
        request.Location = location

        request.Bucket = bucket

        request.Parent = parent

        return stub.ListLoggingLogView(request).items

    def to_proto(self):
        resource = log_view_pb2.LoggingLogView()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.filter):
            resource.filter = Primitive.to_proto(self.filter)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.bucket):
            resource.bucket = Primitive.to_proto(self.bucket)
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
