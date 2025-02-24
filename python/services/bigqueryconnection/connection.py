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
from google3.cloud.graphite.mmv2.services.google.bigquery_connection import (
    connection_pb2,
)
from google3.cloud.graphite.mmv2.services.google.bigquery_connection import (
    connection_pb2_grpc,
)

from typing import List


class Connection(object):
    def __init__(
        self,
        name: str = None,
        friendly_name: str = None,
        description: str = None,
        cloud_sql: dict = None,
        creation_time: int = None,
        last_modified_time: int = None,
        has_credential: bool = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.friendly_name = friendly_name
        self.description = description
        self.cloud_sql = cloud_sql
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = connection_pb2_grpc.BigqueryconnectionConnectionServiceStub(
            channel.Channel()
        )
        request = connection_pb2.ApplyBigqueryconnectionConnectionRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.friendly_name):
            request.resource.friendly_name = Primitive.to_proto(self.friendly_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ConnectionCloudSql.to_proto(self.cloud_sql):
            request.resource.cloud_sql.CopyFrom(
                ConnectionCloudSql.to_proto(self.cloud_sql)
            )
        else:
            request.resource.ClearField("cloud_sql")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyBigqueryconnectionConnection(request)
        self.name = Primitive.from_proto(response.name)
        self.friendly_name = Primitive.from_proto(response.friendly_name)
        self.description = Primitive.from_proto(response.description)
        self.cloud_sql = ConnectionCloudSql.from_proto(response.cloud_sql)
        self.creation_time = Primitive.from_proto(response.creation_time)
        self.last_modified_time = Primitive.from_proto(response.last_modified_time)
        self.has_credential = Primitive.from_proto(response.has_credential)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = connection_pb2_grpc.BigqueryconnectionConnectionServiceStub(
            channel.Channel()
        )
        request = connection_pb2.DeleteBigqueryconnectionConnectionRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.friendly_name):
            request.resource.friendly_name = Primitive.to_proto(self.friendly_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ConnectionCloudSql.to_proto(self.cloud_sql):
            request.resource.cloud_sql.CopyFrom(
                ConnectionCloudSql.to_proto(self.cloud_sql)
            )
        else:
            request.resource.ClearField("cloud_sql")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteBigqueryconnectionConnection(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = connection_pb2_grpc.BigqueryconnectionConnectionServiceStub(
            channel.Channel()
        )
        request = connection_pb2.ListBigqueryconnectionConnectionRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListBigqueryconnectionConnection(request).items

    def to_proto(self):
        resource = connection_pb2.BigqueryconnectionConnection()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.friendly_name):
            resource.friendly_name = Primitive.to_proto(self.friendly_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if ConnectionCloudSql.to_proto(self.cloud_sql):
            resource.cloud_sql.CopyFrom(ConnectionCloudSql.to_proto(self.cloud_sql))
        else:
            resource.ClearField("cloud_sql")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ConnectionCloudSql(object):
    def __init__(
        self,
        instance_id: str = None,
        database: str = None,
        type: str = None,
        credential: dict = None,
    ):
        self.instance_id = instance_id
        self.database = database
        self.type = type
        self.credential = credential

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = connection_pb2.BigqueryconnectionConnectionCloudSql()
        if Primitive.to_proto(resource.instance_id):
            res.instance_id = Primitive.to_proto(resource.instance_id)
        if Primitive.to_proto(resource.database):
            res.database = Primitive.to_proto(resource.database)
        if ConnectionCloudSqlTypeEnum.to_proto(resource.type):
            res.type = ConnectionCloudSqlTypeEnum.to_proto(resource.type)
        if ConnectionCloudSqlCredential.to_proto(resource.credential):
            res.credential.CopyFrom(
                ConnectionCloudSqlCredential.to_proto(resource.credential)
            )
        else:
            res.ClearField("credential")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConnectionCloudSql(
            instance_id=Primitive.from_proto(resource.instance_id),
            database=Primitive.from_proto(resource.database),
            type=ConnectionCloudSqlTypeEnum.from_proto(resource.type),
            credential=ConnectionCloudSqlCredential.from_proto(resource.credential),
        )


class ConnectionCloudSqlArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConnectionCloudSql.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConnectionCloudSql.from_proto(i) for i in resources]


class ConnectionCloudSqlCredential(object):
    def __init__(self, username: str = None, password: str = None):
        self.username = username
        self.password = password

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = connection_pb2.BigqueryconnectionConnectionCloudSqlCredential()
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        if Primitive.to_proto(resource.password):
            res.password = Primitive.to_proto(resource.password)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConnectionCloudSqlCredential(
            username=Primitive.from_proto(resource.username),
            password=Primitive.from_proto(resource.password),
        )


class ConnectionCloudSqlCredentialArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConnectionCloudSqlCredential.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConnectionCloudSqlCredential.from_proto(i) for i in resources]


class ConnectionCloudSqlTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return connection_pb2.BigqueryconnectionConnectionCloudSqlTypeEnum.Value(
            "BigqueryconnectionConnectionCloudSqlTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return connection_pb2.BigqueryconnectionConnectionCloudSqlTypeEnum.Name(
            resource
        )[len("BigqueryconnectionConnectionCloudSqlTypeEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
