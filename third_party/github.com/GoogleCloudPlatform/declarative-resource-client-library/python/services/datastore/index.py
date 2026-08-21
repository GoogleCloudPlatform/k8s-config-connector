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
from google3.cloud.graphite.mmv2.services.google.datastore import index_pb2
from google3.cloud.graphite.mmv2.services.google.datastore import index_pb2_grpc

from typing import List


class Index(object):
    def __init__(
        self,
        ancestor: str = None,
        index_id: str = None,
        kind: str = None,
        project: str = None,
        properties: list = None,
        state: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.ancestor = ancestor
        self.kind = kind
        self.project = project
        self.properties = properties
        self.service_account_file = service_account_file

    def apply(self):
        stub = index_pb2_grpc.DatastoreIndexServiceStub(channel.Channel())
        request = index_pb2.ApplyDatastoreIndexRequest()
        if IndexAncestorEnum.to_proto(self.ancestor):
            request.resource.ancestor = IndexAncestorEnum.to_proto(self.ancestor)

        if Primitive.to_proto(self.kind):
            request.resource.kind = Primitive.to_proto(self.kind)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if IndexPropertiesArray.to_proto(self.properties):
            request.resource.properties.extend(
                IndexPropertiesArray.to_proto(self.properties)
            )
        request.service_account_file = self.service_account_file

        response = stub.ApplyDatastoreIndex(request)
        self.ancestor = IndexAncestorEnum.from_proto(response.ancestor)
        self.index_id = Primitive.from_proto(response.index_id)
        self.kind = Primitive.from_proto(response.kind)
        self.project = Primitive.from_proto(response.project)
        self.properties = IndexPropertiesArray.from_proto(response.properties)
        self.state = IndexStateEnum.from_proto(response.state)

    def delete(self):
        stub = index_pb2_grpc.DatastoreIndexServiceStub(channel.Channel())
        request = index_pb2.DeleteDatastoreIndexRequest()
        request.service_account_file = self.service_account_file
        if IndexAncestorEnum.to_proto(self.ancestor):
            request.resource.ancestor = IndexAncestorEnum.to_proto(self.ancestor)

        if Primitive.to_proto(self.kind):
            request.resource.kind = Primitive.to_proto(self.kind)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if IndexPropertiesArray.to_proto(self.properties):
            request.resource.properties.extend(
                IndexPropertiesArray.to_proto(self.properties)
            )
        response = stub.DeleteDatastoreIndex(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = index_pb2_grpc.DatastoreIndexServiceStub(channel.Channel())
        request = index_pb2.ListDatastoreIndexRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListDatastoreIndex(request).items

    def to_proto(self):
        resource = index_pb2.DatastoreIndex()
        if IndexAncestorEnum.to_proto(self.ancestor):
            resource.ancestor = IndexAncestorEnum.to_proto(self.ancestor)
        if Primitive.to_proto(self.kind):
            resource.kind = Primitive.to_proto(self.kind)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if IndexPropertiesArray.to_proto(self.properties):
            resource.properties.extend(IndexPropertiesArray.to_proto(self.properties))
        return resource


class IndexProperties(object):
    def __init__(self, name: str = None, direction: str = None):
        self.name = name
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = index_pb2.DatastoreIndexProperties()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if IndexPropertiesDirectionEnum.to_proto(resource.direction):
            res.direction = IndexPropertiesDirectionEnum.to_proto(resource.direction)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return IndexProperties(
            name=Primitive.from_proto(resource.name),
            direction=IndexPropertiesDirectionEnum.from_proto(resource.direction),
        )


class IndexPropertiesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [IndexProperties.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [IndexProperties.from_proto(i) for i in resources]


class IndexAncestorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return index_pb2.DatastoreIndexAncestorEnum.Value(
            "DatastoreIndexAncestorEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return index_pb2.DatastoreIndexAncestorEnum.Name(resource)[
            len("DatastoreIndexAncestorEnum") :
        ]


class IndexPropertiesDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return index_pb2.DatastoreIndexPropertiesDirectionEnum.Value(
            "DatastoreIndexPropertiesDirectionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return index_pb2.DatastoreIndexPropertiesDirectionEnum.Name(resource)[
            len("DatastoreIndexPropertiesDirectionEnum") :
        ]


class IndexStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return index_pb2.DatastoreIndexStateEnum.Value(
            "DatastoreIndexStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return index_pb2.DatastoreIndexStateEnum.Name(resource)[
            len("DatastoreIndexStateEnum") :
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
