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
from google3.cloud.graphite.mmv2.services.google.sql import user_pb2
from google3.cloud.graphite.mmv2.services.google.sql import user_pb2_grpc

from typing import List


class User(object):
    def __init__(
        self,
        name: str = None,
        password: str = None,
        project: str = None,
        instance: str = None,
        sqlserver_user_details: dict = None,
        type: str = None,
        etag: str = None,
        host: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.password = password
        self.project = project
        self.instance = instance
        self.sqlserver_user_details = sqlserver_user_details
        self.type = type
        self.etag = etag
        self.host = host
        self.service_account_file = service_account_file

    def apply(self):
        stub = user_pb2_grpc.SqlBetaUserServiceStub(channel.Channel())
        request = user_pb2.ApplySqlBetaUserRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.password):
            request.resource.password = Primitive.to_proto(self.password)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.instance):
            request.resource.instance = Primitive.to_proto(self.instance)

        if UserSqlserverUserDetails.to_proto(self.sqlserver_user_details):
            request.resource.sqlserver_user_details.CopyFrom(
                UserSqlserverUserDetails.to_proto(self.sqlserver_user_details)
            )
        else:
            request.resource.ClearField("sqlserver_user_details")
        if UserTypeEnum.to_proto(self.type):
            request.resource.type = UserTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.etag):
            request.resource.etag = Primitive.to_proto(self.etag)

        if Primitive.to_proto(self.host):
            request.resource.host = Primitive.to_proto(self.host)

        request.service_account_file = self.service_account_file

        response = stub.ApplySqlBetaUser(request)
        self.name = Primitive.from_proto(response.name)
        self.password = Primitive.from_proto(response.password)
        self.project = Primitive.from_proto(response.project)
        self.instance = Primitive.from_proto(response.instance)
        self.sqlserver_user_details = UserSqlserverUserDetails.from_proto(
            response.sqlserver_user_details
        )
        self.type = UserTypeEnum.from_proto(response.type)
        self.etag = Primitive.from_proto(response.etag)
        self.host = Primitive.from_proto(response.host)

    def delete(self):
        stub = user_pb2_grpc.SqlBetaUserServiceStub(channel.Channel())
        request = user_pb2.DeleteSqlBetaUserRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.password):
            request.resource.password = Primitive.to_proto(self.password)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.instance):
            request.resource.instance = Primitive.to_proto(self.instance)

        if UserSqlserverUserDetails.to_proto(self.sqlserver_user_details):
            request.resource.sqlserver_user_details.CopyFrom(
                UserSqlserverUserDetails.to_proto(self.sqlserver_user_details)
            )
        else:
            request.resource.ClearField("sqlserver_user_details")
        if UserTypeEnum.to_proto(self.type):
            request.resource.type = UserTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.etag):
            request.resource.etag = Primitive.to_proto(self.etag)

        if Primitive.to_proto(self.host):
            request.resource.host = Primitive.to_proto(self.host)

        response = stub.DeleteSqlBetaUser(request)

    @classmethod
    def list(self, project, instance, service_account_file=""):
        stub = user_pb2_grpc.SqlBetaUserServiceStub(channel.Channel())
        request = user_pb2.ListSqlBetaUserRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Instance = instance

        return stub.ListSqlBetaUser(request).items

    def to_proto(self):
        resource = user_pb2.SqlBetaUser()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.password):
            resource.password = Primitive.to_proto(self.password)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.instance):
            resource.instance = Primitive.to_proto(self.instance)
        if UserSqlserverUserDetails.to_proto(self.sqlserver_user_details):
            resource.sqlserver_user_details.CopyFrom(
                UserSqlserverUserDetails.to_proto(self.sqlserver_user_details)
            )
        else:
            resource.ClearField("sqlserver_user_details")
        if UserTypeEnum.to_proto(self.type):
            resource.type = UserTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.etag):
            resource.etag = Primitive.to_proto(self.etag)
        if Primitive.to_proto(self.host):
            resource.host = Primitive.to_proto(self.host)
        return resource


class UserSqlserverUserDetails(object):
    def __init__(self, disabled: bool = None, server_roles: list = None):
        self.disabled = disabled
        self.server_roles = server_roles

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = user_pb2.SqlBetaUserSqlserverUserDetails()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        if Primitive.to_proto(resource.server_roles):
            res.server_roles.extend(Primitive.to_proto(resource.server_roles))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UserSqlserverUserDetails(
            disabled=Primitive.from_proto(resource.disabled),
            server_roles=Primitive.from_proto(resource.server_roles),
        )


class UserSqlserverUserDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UserSqlserverUserDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UserSqlserverUserDetails.from_proto(i) for i in resources]


class UserTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return user_pb2.SqlBetaUserTypeEnum.Value("SqlBetaUserTypeEnum%s" % resource)

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return user_pb2.SqlBetaUserTypeEnum.Name(resource)[len("SqlBetaUserTypeEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
