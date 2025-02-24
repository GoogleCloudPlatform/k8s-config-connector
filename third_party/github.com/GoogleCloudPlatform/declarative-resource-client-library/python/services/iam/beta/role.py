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
from google3.cloud.graphite.mmv2.services.google.iam import role_pb2
from google3.cloud.graphite.mmv2.services.google.iam import role_pb2_grpc

from typing import List


class Role(object):
    def __init__(
        self,
        name: str = None,
        title: str = None,
        description: str = None,
        localized_values: dict = None,
        lifecycle_phase: str = None,
        group_name: str = None,
        group_title: str = None,
        included_permissions: list = None,
        stage: str = None,
        etag: str = None,
        deleted: bool = None,
        included_roles: list = None,
        parent: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.title = title
        self.description = description
        self.localized_values = localized_values
        self.lifecycle_phase = lifecycle_phase
        self.group_name = group_name
        self.group_title = group_title
        self.included_permissions = included_permissions
        self.stage = stage
        self.etag = etag
        self.deleted = deleted
        self.included_roles = included_roles
        self.parent = parent
        self.service_account_file = service_account_file

    def apply(self):
        stub = role_pb2_grpc.IamBetaRoleServiceStub(channel.Channel())
        request = role_pb2.ApplyIamBetaRoleRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.title):
            request.resource.title = Primitive.to_proto(self.title)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if RoleLocalizedValues.to_proto(self.localized_values):
            request.resource.localized_values.CopyFrom(
                RoleLocalizedValues.to_proto(self.localized_values)
            )
        else:
            request.resource.ClearField("localized_values")
        if Primitive.to_proto(self.lifecycle_phase):
            request.resource.lifecycle_phase = Primitive.to_proto(self.lifecycle_phase)

        if Primitive.to_proto(self.group_name):
            request.resource.group_name = Primitive.to_proto(self.group_name)

        if Primitive.to_proto(self.group_title):
            request.resource.group_title = Primitive.to_proto(self.group_title)

        if Primitive.to_proto(self.included_permissions):
            request.resource.included_permissions.extend(
                Primitive.to_proto(self.included_permissions)
            )
        if RoleStageEnum.to_proto(self.stage):
            request.resource.stage = RoleStageEnum.to_proto(self.stage)

        if Primitive.to_proto(self.etag):
            request.resource.etag = Primitive.to_proto(self.etag)

        if Primitive.to_proto(self.deleted):
            request.resource.deleted = Primitive.to_proto(self.deleted)

        if Primitive.to_proto(self.included_roles):
            request.resource.included_roles.extend(
                Primitive.to_proto(self.included_roles)
            )
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        request.service_account_file = self.service_account_file

        response = stub.ApplyIamBetaRole(request)
        self.name = Primitive.from_proto(response.name)
        self.title = Primitive.from_proto(response.title)
        self.description = Primitive.from_proto(response.description)
        self.localized_values = RoleLocalizedValues.from_proto(
            response.localized_values
        )
        self.lifecycle_phase = Primitive.from_proto(response.lifecycle_phase)
        self.group_name = Primitive.from_proto(response.group_name)
        self.group_title = Primitive.from_proto(response.group_title)
        self.included_permissions = Primitive.from_proto(response.included_permissions)
        self.stage = RoleStageEnum.from_proto(response.stage)
        self.etag = Primitive.from_proto(response.etag)
        self.deleted = Primitive.from_proto(response.deleted)
        self.included_roles = Primitive.from_proto(response.included_roles)
        self.parent = Primitive.from_proto(response.parent)

    def delete(self):
        stub = role_pb2_grpc.IamBetaRoleServiceStub(channel.Channel())
        request = role_pb2.DeleteIamBetaRoleRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.title):
            request.resource.title = Primitive.to_proto(self.title)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if RoleLocalizedValues.to_proto(self.localized_values):
            request.resource.localized_values.CopyFrom(
                RoleLocalizedValues.to_proto(self.localized_values)
            )
        else:
            request.resource.ClearField("localized_values")
        if Primitive.to_proto(self.lifecycle_phase):
            request.resource.lifecycle_phase = Primitive.to_proto(self.lifecycle_phase)

        if Primitive.to_proto(self.group_name):
            request.resource.group_name = Primitive.to_proto(self.group_name)

        if Primitive.to_proto(self.group_title):
            request.resource.group_title = Primitive.to_proto(self.group_title)

        if Primitive.to_proto(self.included_permissions):
            request.resource.included_permissions.extend(
                Primitive.to_proto(self.included_permissions)
            )
        if RoleStageEnum.to_proto(self.stage):
            request.resource.stage = RoleStageEnum.to_proto(self.stage)

        if Primitive.to_proto(self.etag):
            request.resource.etag = Primitive.to_proto(self.etag)

        if Primitive.to_proto(self.deleted):
            request.resource.deleted = Primitive.to_proto(self.deleted)

        if Primitive.to_proto(self.included_roles):
            request.resource.included_roles.extend(
                Primitive.to_proto(self.included_roles)
            )
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        response = stub.DeleteIamBetaRole(request)

    @classmethod
    def list(self, parent, service_account_file=""):
        stub = role_pb2_grpc.IamBetaRoleServiceStub(channel.Channel())
        request = role_pb2.ListIamBetaRoleRequest()
        request.service_account_file = service_account_file
        request.Parent = parent

        return stub.ListIamBetaRole(request).items

    def to_proto(self):
        resource = role_pb2.IamBetaRole()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.title):
            resource.title = Primitive.to_proto(self.title)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if RoleLocalizedValues.to_proto(self.localized_values):
            resource.localized_values.CopyFrom(
                RoleLocalizedValues.to_proto(self.localized_values)
            )
        else:
            resource.ClearField("localized_values")
        if Primitive.to_proto(self.lifecycle_phase):
            resource.lifecycle_phase = Primitive.to_proto(self.lifecycle_phase)
        if Primitive.to_proto(self.group_name):
            resource.group_name = Primitive.to_proto(self.group_name)
        if Primitive.to_proto(self.group_title):
            resource.group_title = Primitive.to_proto(self.group_title)
        if Primitive.to_proto(self.included_permissions):
            resource.included_permissions.extend(
                Primitive.to_proto(self.included_permissions)
            )
        if RoleStageEnum.to_proto(self.stage):
            resource.stage = RoleStageEnum.to_proto(self.stage)
        if Primitive.to_proto(self.etag):
            resource.etag = Primitive.to_proto(self.etag)
        if Primitive.to_proto(self.deleted):
            resource.deleted = Primitive.to_proto(self.deleted)
        if Primitive.to_proto(self.included_roles):
            resource.included_roles.extend(Primitive.to_proto(self.included_roles))
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        return resource


class RoleLocalizedValues(object):
    def __init__(self, localized_title: str = None, localized_description: str = None):
        self.localized_title = localized_title
        self.localized_description = localized_description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = role_pb2.IamBetaRoleLocalizedValues()
        if Primitive.to_proto(resource.localized_title):
            res.localized_title = Primitive.to_proto(resource.localized_title)
        if Primitive.to_proto(resource.localized_description):
            res.localized_description = Primitive.to_proto(
                resource.localized_description
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RoleLocalizedValues(
            localized_title=Primitive.from_proto(resource.localized_title),
            localized_description=Primitive.from_proto(resource.localized_description),
        )


class RoleLocalizedValuesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RoleLocalizedValues.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RoleLocalizedValues.from_proto(i) for i in resources]


class RoleStageEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return role_pb2.IamBetaRoleStageEnum.Value("IamBetaRoleStageEnum%s" % resource)

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return role_pb2.IamBetaRoleStageEnum.Name(resource)[
            len("IamBetaRoleStageEnum") :
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
