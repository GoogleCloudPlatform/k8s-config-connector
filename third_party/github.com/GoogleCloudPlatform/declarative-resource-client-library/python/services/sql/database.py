# Copyright 2020 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.sql import database_pb2
from google3.cloud.graphite.mmv2.services.google.sql import database_pb2_grpc

from typing import List


class Database(object):
    def __init__(
        self,
        charset: str = None,
        collation: str = None,
        instance: str = None,
        name: str = None,
        project: str = None,
        self_link: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.charset = charset
        self.collation = collation
        self.instance = instance
        self.name = name
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = database_pb2_grpc.SqlDatabaseServiceStub(channel.Channel())
        request = database_pb2.ApplySqlDatabaseRequest()
        if Primitive.to_proto(self.charset):
            request.resource.charset = Primitive.to_proto(self.charset)

        if Primitive.to_proto(self.collation):
            request.resource.collation = Primitive.to_proto(self.collation)

        if Primitive.to_proto(self.instance):
            request.resource.instance = Primitive.to_proto(self.instance)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplySqlDatabase(request)
        self.charset = Primitive.from_proto(response.charset)
        self.collation = Primitive.from_proto(response.collation)
        self.instance = Primitive.from_proto(response.instance)
        self.name = Primitive.from_proto(response.name)
        self.project = Primitive.from_proto(response.project)
        self.self_link = Primitive.from_proto(response.self_link)

    @classmethod
    def delete(self, project, instance, name, service_account_file=""):
        stub = database_pb2_grpc.SqlDatabaseServiceStub(channel.Channel())
        request = database_pb2.DeleteSqlDatabaseRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Instance = instance

        request.Name = name

        response = stub.DeleteSqlDatabase(request)

    @classmethod
    def list(self, project, instance, service_account_file=""):
        stub = database_pb2_grpc.SqlDatabaseServiceStub(channel.Channel())
        request = database_pb2.ListSqlDatabaseRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Instance = instance

        return stub.ListSqlDatabase(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = database_pb2.SqlDatabase()
        any_proto.Unpack(res_proto)

        res = Database()
        res.charset = Primitive.from_proto(res_proto.charset)
        res.collation = Primitive.from_proto(res_proto.collation)
        res.instance = Primitive.from_proto(res_proto.instance)
        res.name = Primitive.from_proto(res_proto.name)
        res.project = Primitive.from_proto(res_proto.project)
        res.self_link = Primitive.from_proto(res_proto.self_link)
        return res


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
