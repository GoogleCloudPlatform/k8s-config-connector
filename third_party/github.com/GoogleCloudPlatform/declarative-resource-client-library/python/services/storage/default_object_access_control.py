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
    default_object_access_control_pb2,
)
from google3.cloud.graphite.mmv2.services.google.storage import (
    default_object_access_control_pb2_grpc,
)

from typing import List


class DefaultObjectAccessControl(object):
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
        service_account_file: str = "",
    ):

        channel.initialize()
        self.project = project
        self.bucket = bucket
        self.entity = entity
        self.role = role
        self.service_account_file = service_account_file

    def apply(self):
        stub = default_object_access_control_pb2_grpc.StorageDefaultObjectAccessControlServiceStub(
            channel.Channel()
        )
        request = (
            default_object_access_control_pb2.ApplyStorageDefaultObjectAccessControlRequest()
        )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.bucket):
            request.resource.bucket = Primitive.to_proto(self.bucket)

        if Primitive.to_proto(self.entity):
            request.resource.entity = Primitive.to_proto(self.entity)

        if DefaultObjectAccessControlRoleEnum.to_proto(self.role):
            request.resource.role = DefaultObjectAccessControlRoleEnum.to_proto(
                self.role
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyStorageDefaultObjectAccessControl(request)
        self.project = Primitive.from_proto(response.project)
        self.bucket = Primitive.from_proto(response.bucket)
        self.domain = Primitive.from_proto(response.domain)
        self.email = Primitive.from_proto(response.email)
        self.entity = Primitive.from_proto(response.entity)
        self.entity_id = Primitive.from_proto(response.entity_id)
        self.project_team = DefaultObjectAccessControlProjectTeam.from_proto(
            response.project_team
        )
        self.role = DefaultObjectAccessControlRoleEnum.from_proto(response.role)

    def delete(self):
        stub = default_object_access_control_pb2_grpc.StorageDefaultObjectAccessControlServiceStub(
            channel.Channel()
        )
        request = (
            default_object_access_control_pb2.DeleteStorageDefaultObjectAccessControlRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.bucket):
            request.resource.bucket = Primitive.to_proto(self.bucket)

        if Primitive.to_proto(self.entity):
            request.resource.entity = Primitive.to_proto(self.entity)

        if DefaultObjectAccessControlRoleEnum.to_proto(self.role):
            request.resource.role = DefaultObjectAccessControlRoleEnum.to_proto(
                self.role
            )

        response = stub.DeleteStorageDefaultObjectAccessControl(request)

    @classmethod
    def list(self, project, bucket, service_account_file=""):
        stub = default_object_access_control_pb2_grpc.StorageDefaultObjectAccessControlServiceStub(
            channel.Channel()
        )
        request = (
            default_object_access_control_pb2.ListStorageDefaultObjectAccessControlRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Bucket = bucket

        return stub.ListStorageDefaultObjectAccessControl(request).items

    def to_proto(self):
        resource = default_object_access_control_pb2.StorageDefaultObjectAccessControl()
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.bucket):
            resource.bucket = Primitive.to_proto(self.bucket)
        if Primitive.to_proto(self.entity):
            resource.entity = Primitive.to_proto(self.entity)
        if DefaultObjectAccessControlRoleEnum.to_proto(self.role):
            resource.role = DefaultObjectAccessControlRoleEnum.to_proto(self.role)
        return resource


class DefaultObjectAccessControlProjectTeam(object):
    def __init__(self, project_number: str = None, team: str = None):
        self.project_number = project_number
        self.team = team

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            default_object_access_control_pb2.StorageDefaultObjectAccessControlProjectTeam()
        )
        if Primitive.to_proto(resource.project_number):
            res.project_number = Primitive.to_proto(resource.project_number)
        if DefaultObjectAccessControlProjectTeamTeamEnum.to_proto(resource.team):
            res.team = DefaultObjectAccessControlProjectTeamTeamEnum.to_proto(
                resource.team
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DefaultObjectAccessControlProjectTeam(
            project_number=Primitive.from_proto(resource.project_number),
            team=DefaultObjectAccessControlProjectTeamTeamEnum.from_proto(
                resource.team
            ),
        )


class DefaultObjectAccessControlProjectTeamArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DefaultObjectAccessControlProjectTeam.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DefaultObjectAccessControlProjectTeam.from_proto(i) for i in resources]


class DefaultObjectAccessControlProjectTeamTeamEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return default_object_access_control_pb2.StorageDefaultObjectAccessControlProjectTeamTeamEnum.Value(
            "StorageDefaultObjectAccessControlProjectTeamTeamEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return default_object_access_control_pb2.StorageDefaultObjectAccessControlProjectTeamTeamEnum.Name(
            resource
        )[
            len("StorageDefaultObjectAccessControlProjectTeamTeamEnum") :
        ]


class DefaultObjectAccessControlRoleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return default_object_access_control_pb2.StorageDefaultObjectAccessControlRoleEnum.Value(
            "StorageDefaultObjectAccessControlRoleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return default_object_access_control_pb2.StorageDefaultObjectAccessControlRoleEnum.Name(
            resource
        )[
            len("StorageDefaultObjectAccessControlRoleEnum") :
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
