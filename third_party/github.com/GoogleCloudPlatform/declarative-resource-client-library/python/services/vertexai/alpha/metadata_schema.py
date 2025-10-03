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
from google3.cloud.graphite.mmv2.services.google.vertex_ai import metadata_schema_pb2
from google3.cloud.graphite.mmv2.services.google.vertex_ai import (
    metadata_schema_pb2_grpc,
)

from typing import List


class MetadataSchema(object):
    def __init__(
        self,
        name: str = None,
        schema_version: str = None,
        schema: str = None,
        schema_type: str = None,
        create_time: str = None,
        project: str = None,
        location: str = None,
        metadata_store: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.schema_version = schema_version
        self.schema = schema
        self.schema_type = schema_type
        self.project = project
        self.location = location
        self.metadata_store = metadata_store
        self.service_account_file = service_account_file

    def apply(self):
        stub = metadata_schema_pb2_grpc.VertexaiAlphaMetadataSchemaServiceStub(
            channel.Channel()
        )
        request = metadata_schema_pb2.ApplyVertexaiAlphaMetadataSchemaRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.schema_version):
            request.resource.schema_version = Primitive.to_proto(self.schema_version)

        if Primitive.to_proto(self.schema):
            request.resource.schema = Primitive.to_proto(self.schema)

        if MetadataSchemaSchemaTypeEnum.to_proto(self.schema_type):
            request.resource.schema_type = MetadataSchemaSchemaTypeEnum.to_proto(
                self.schema_type
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.metadata_store):
            request.resource.metadata_store = Primitive.to_proto(self.metadata_store)

        request.service_account_file = self.service_account_file

        response = stub.ApplyVertexaiAlphaMetadataSchema(request)
        self.name = Primitive.from_proto(response.name)
        self.schema_version = Primitive.from_proto(response.schema_version)
        self.schema = Primitive.from_proto(response.schema)
        self.schema_type = MetadataSchemaSchemaTypeEnum.from_proto(response.schema_type)
        self.create_time = Primitive.from_proto(response.create_time)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.metadata_store = Primitive.from_proto(response.metadata_store)

    def delete(self):
        stub = metadata_schema_pb2_grpc.VertexaiAlphaMetadataSchemaServiceStub(
            channel.Channel()
        )
        request = metadata_schema_pb2.DeleteVertexaiAlphaMetadataSchemaRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.schema_version):
            request.resource.schema_version = Primitive.to_proto(self.schema_version)

        if Primitive.to_proto(self.schema):
            request.resource.schema = Primitive.to_proto(self.schema)

        if MetadataSchemaSchemaTypeEnum.to_proto(self.schema_type):
            request.resource.schema_type = MetadataSchemaSchemaTypeEnum.to_proto(
                self.schema_type
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.metadata_store):
            request.resource.metadata_store = Primitive.to_proto(self.metadata_store)

        response = stub.DeleteVertexaiAlphaMetadataSchema(request)

    @classmethod
    def list(self, project, location, metadataStore, service_account_file=""):
        stub = metadata_schema_pb2_grpc.VertexaiAlphaMetadataSchemaServiceStub(
            channel.Channel()
        )
        request = metadata_schema_pb2.ListVertexaiAlphaMetadataSchemaRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.MetadataStore = metadataStore

        return stub.ListVertexaiAlphaMetadataSchema(request).items

    def to_proto(self):
        resource = metadata_schema_pb2.VertexaiAlphaMetadataSchema()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.schema_version):
            resource.schema_version = Primitive.to_proto(self.schema_version)
        if Primitive.to_proto(self.schema):
            resource.schema = Primitive.to_proto(self.schema)
        if MetadataSchemaSchemaTypeEnum.to_proto(self.schema_type):
            resource.schema_type = MetadataSchemaSchemaTypeEnum.to_proto(
                self.schema_type
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.metadata_store):
            resource.metadata_store = Primitive.to_proto(self.metadata_store)
        return resource


class MetadataSchemaSchemaTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return metadata_schema_pb2.VertexaiAlphaMetadataSchemaSchemaTypeEnum.Value(
            "VertexaiAlphaMetadataSchemaSchemaTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return metadata_schema_pb2.VertexaiAlphaMetadataSchemaSchemaTypeEnum.Name(
            resource
        )[len("VertexaiAlphaMetadataSchemaSchemaTypeEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
