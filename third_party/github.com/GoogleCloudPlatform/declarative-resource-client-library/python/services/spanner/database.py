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
from google3.cloud.graphite.mmv2.services.google.spanner import database_pb2
from google3.cloud.graphite.mmv2.services.google.spanner import database_pb2_grpc

from typing import List


class Database(object):
    def __init__(
        self,
        name: str = None,
        instance: str = None,
        state: str = None,
        project: str = None,
        ddl: list = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.instance = instance
        self.project = project
        self.ddl = ddl
        self.service_account_file = service_account_file

    def apply(self):
        stub = database_pb2_grpc.SpannerDatabaseServiceStub(channel.Channel())
        request = database_pb2.ApplySpannerDatabaseRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.instance):
            request.resource.instance = Primitive.to_proto(self.instance)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.ddl):
            request.resource.ddl.extend(Primitive.to_proto(self.ddl))
        request.service_account_file = self.service_account_file

        response = stub.ApplySpannerDatabase(request)
        self.name = Primitive.from_proto(response.name)
        self.instance = Primitive.from_proto(response.instance)
        self.state = DatabaseStateEnum.from_proto(response.state)
        self.project = Primitive.from_proto(response.project)
        self.ddl = Primitive.from_proto(response.ddl)

    def delete(self):
        stub = database_pb2_grpc.SpannerDatabaseServiceStub(channel.Channel())
        request = database_pb2.DeleteSpannerDatabaseRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.instance):
            request.resource.instance = Primitive.to_proto(self.instance)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.ddl):
            request.resource.ddl.extend(Primitive.to_proto(self.ddl))
        response = stub.DeleteSpannerDatabase(request)

    @classmethod
    def list(self, project, instance, service_account_file=""):
        stub = database_pb2_grpc.SpannerDatabaseServiceStub(channel.Channel())
        request = database_pb2.ListSpannerDatabaseRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Instance = instance

        return stub.ListSpannerDatabase(request).items

    def to_proto(self):
        resource = database_pb2.SpannerDatabase()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.instance):
            resource.instance = Primitive.to_proto(self.instance)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.ddl):
            resource.ddl.extend(Primitive.to_proto(self.ddl))
        return resource


class DatabaseStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return database_pb2.SpannerDatabaseStateEnum.Value(
            "SpannerDatabaseStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return database_pb2.SpannerDatabaseStateEnum.Name(resource)[
            len("SpannerDatabaseStateEnum") :
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
