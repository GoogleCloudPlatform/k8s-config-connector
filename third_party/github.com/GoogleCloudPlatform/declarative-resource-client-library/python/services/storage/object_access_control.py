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
from google3.cloud.graphite.mmv2.services.google.storage import (
    object_access_control_pb2,
)
from google3.cloud.graphite.mmv2.services.google.storage import (
    object_access_control_pb2_grpc,
)

from typing import List


class ObjectAccessControl(object):
    def __init__(
        self,
        project: str = None,
        bucket: str = None,
        domain: str = None,
        email: str = None,
        entity: str = None,
        entity_id: str = None,
        project_team: dict = None,
        role: str = None,
        id: str = None,
        object: str = None,
        generation: int = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.project = project
        self.bucket = bucket
        self.entity = entity
        self.role = role
        self.object = object
        self.service_account_file = service_account_file

    def apply(self):
        stub = object_access_control_pb2_grpc.StorageObjectAccessControlServiceStub(
            channel.Channel()
        )
        request = object_access_control_pb2.ApplyStorageObjectAccessControlRequest()
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.bucket):
            request.resource.bucket = Primitive.to_proto(self.bucket)

        if Primitive.to_proto(self.entity):
            request.resource.entity = Primitive.to_proto(self.entity)

        if ObjectAccessControlRoleEnum.to_proto(self.role):
            request.resource.role = ObjectAccessControlRoleEnum.to_proto(self.role)

        if Primitive.to_proto(self.object):
            request.resource.object = Primitive.to_proto(self.object)

        request.service_account_file = self.service_account_file

        response = stub.ApplyStorageObjectAccessControl(request)
        self.project = Primitive.from_proto(response.project)
        self.bucket = Primitive.from_proto(response.bucket)
        self.domain = Primitive.from_proto(response.domain)
        self.email = Primitive.from_proto(response.email)
        self.entity = Primitive.from_proto(response.entity)
        self.entity_id = Primitive.from_proto(response.entity_id)
        self.project_team = ObjectAccessControlProjectTeam.from_proto(
            response.project_team
        )
        self.role = ObjectAccessControlRoleEnum.from_proto(response.role)
        self.id = Primitive.from_proto(response.id)
        self.object = Primitive.from_proto(response.object)
        self.generation = Primitive.from_proto(response.generation)

    def delete(self):
        stub = object_access_control_pb2_grpc.StorageObjectAccessControlServiceStub(
            channel.Channel()
        )
        request = object_access_control_pb2.DeleteStorageObjectAccessControlRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.bucket):
            request.resource.bucket = Primitive.to_proto(self.bucket)

        if Primitive.to_proto(self.entity):
            request.resource.entity = Primitive.to_proto(self.entity)

        if ObjectAccessControlRoleEnum.to_proto(self.role):
            request.resource.role = ObjectAccessControlRoleEnum.to_proto(self.role)

        if Primitive.to_proto(self.object):
            request.resource.object = Primitive.to_proto(self.object)

        response = stub.DeleteStorageObjectAccessControl(request)

    @classmethod
    def list(self, project, bucket, object, service_account_file=""):
        stub = object_access_control_pb2_grpc.StorageObjectAccessControlServiceStub(
            channel.Channel()
        )
        request = object_access_control_pb2.ListStorageObjectAccessControlRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Bucket = bucket

        request.Object = object

        return stub.ListStorageObjectAccessControl(request).items

    def to_proto(self):
        resource = object_access_control_pb2.StorageObjectAccessControl()
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.bucket):
            resource.bucket = Primitive.to_proto(self.bucket)
        if Primitive.to_proto(self.entity):
            resource.entity = Primitive.to_proto(self.entity)
        if ObjectAccessControlRoleEnum.to_proto(self.role):
            resource.role = ObjectAccessControlRoleEnum.to_proto(self.role)
        if Primitive.to_proto(self.object):
            resource.object = Primitive.to_proto(self.object)
        return resource


class ObjectAccessControlProjectTeam(object):
    def __init__(self, project_number: str = None, team: str = None):
        self.project_number = project_number
        self.team = team

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = object_access_control_pb2.StorageObjectAccessControlProjectTeam()
        if Primitive.to_proto(resource.project_number):
            res.project_number = Primitive.to_proto(resource.project_number)
        if ObjectAccessControlProjectTeamTeamEnum.to_proto(resource.team):
            res.team = ObjectAccessControlProjectTeamTeamEnum.to_proto(resource.team)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ObjectAccessControlProjectTeam(
            project_number=Primitive.from_proto(resource.project_number),
            team=ObjectAccessControlProjectTeamTeamEnum.from_proto(resource.team),
        )


class ObjectAccessControlProjectTeamArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ObjectAccessControlProjectTeam.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ObjectAccessControlProjectTeam.from_proto(i) for i in resources]


class ObjectAccessControlProjectTeamTeamEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return object_access_control_pb2.StorageObjectAccessControlProjectTeamTeamEnum.Value(
            "StorageObjectAccessControlProjectTeamTeamEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return object_access_control_pb2.StorageObjectAccessControlProjectTeamTeamEnum.Name(
            resource
        )[
            len("StorageObjectAccessControlProjectTeamTeamEnum") :
        ]


class ObjectAccessControlRoleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return object_access_control_pb2.StorageObjectAccessControlRoleEnum.Value(
            "StorageObjectAccessControlRoleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return object_access_control_pb2.StorageObjectAccessControlRoleEnum.Name(
            resource
        )[len("StorageObjectAccessControlRoleEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
