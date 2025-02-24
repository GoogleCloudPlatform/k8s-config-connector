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
from google3.cloud.graphite.mmv2.services.google.cloudbuildv2 import repository_pb2
from google3.cloud.graphite.mmv2.services.google.cloudbuildv2 import repository_pb2_grpc

from typing import List


class Repository(object):
    def __init__(
        self,
        name: str = None,
        remote_uri: str = None,
        create_time: str = None,
        update_time: str = None,
        annotations: dict = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        connection: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.remote_uri = remote_uri
        self.annotations = annotations
        self.project = project
        self.location = location
        self.connection = connection
        self.service_account_file = service_account_file

    def apply(self):
        stub = repository_pb2_grpc.Cloudbuildv2BetaRepositoryServiceStub(
            channel.Channel()
        )
        request = repository_pb2.ApplyCloudbuildv2BetaRepositoryRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.remote_uri):
            request.resource.remote_uri = Primitive.to_proto(self.remote_uri)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.connection):
            request.resource.connection = Primitive.to_proto(self.connection)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudbuildv2BetaRepository(request)
        self.name = Primitive.from_proto(response.name)
        self.remote_uri = Primitive.from_proto(response.remote_uri)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.annotations = Primitive.from_proto(response.annotations)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.connection = Primitive.from_proto(response.connection)

    def delete(self):
        stub = repository_pb2_grpc.Cloudbuildv2BetaRepositoryServiceStub(
            channel.Channel()
        )
        request = repository_pb2.DeleteCloudbuildv2BetaRepositoryRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.remote_uri):
            request.resource.remote_uri = Primitive.to_proto(self.remote_uri)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.connection):
            request.resource.connection = Primitive.to_proto(self.connection)

        response = stub.DeleteCloudbuildv2BetaRepository(request)

    @classmethod
    def list(self, project, location, connection, service_account_file=""):
        stub = repository_pb2_grpc.Cloudbuildv2BetaRepositoryServiceStub(
            channel.Channel()
        )
        request = repository_pb2.ListCloudbuildv2BetaRepositoryRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Connection = connection

        return stub.ListCloudbuildv2BetaRepository(request).items

    def to_proto(self):
        resource = repository_pb2.Cloudbuildv2BetaRepository()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.remote_uri):
            resource.remote_uri = Primitive.to_proto(self.remote_uri)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.connection):
            resource.connection = Primitive.to_proto(self.connection)
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
