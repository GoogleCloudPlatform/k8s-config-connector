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
from google3.cloud.graphite.mmv2.services.google.filestore import backup_pb2
from google3.cloud.graphite.mmv2.services.google.filestore import backup_pb2_grpc

from typing import List


class Backup(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        state: str = None,
        create_time: str = None,
        labels: dict = None,
        capacity_gb: int = None,
        storage_bytes: int = None,
        source_instance: str = None,
        source_file_share: str = None,
        source_instance_tier: str = None,
        download_bytes: int = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.labels = labels
        self.source_instance = source_instance
        self.source_file_share = source_file_share
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = backup_pb2_grpc.FilestoreAlphaBackupServiceStub(channel.Channel())
        request = backup_pb2.ApplyFilestoreAlphaBackupRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.source_instance):
            request.resource.source_instance = Primitive.to_proto(self.source_instance)

        if Primitive.to_proto(self.source_file_share):
            request.resource.source_file_share = Primitive.to_proto(
                self.source_file_share
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyFilestoreAlphaBackup(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.state = BackupStateEnum.from_proto(response.state)
        self.create_time = Primitive.from_proto(response.create_time)
        self.labels = Primitive.from_proto(response.labels)
        self.capacity_gb = Primitive.from_proto(response.capacity_gb)
        self.storage_bytes = Primitive.from_proto(response.storage_bytes)
        self.source_instance = Primitive.from_proto(response.source_instance)
        self.source_file_share = Primitive.from_proto(response.source_file_share)
        self.source_instance_tier = BackupSourceInstanceTierEnum.from_proto(
            response.source_instance_tier
        )
        self.download_bytes = Primitive.from_proto(response.download_bytes)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = backup_pb2_grpc.FilestoreAlphaBackupServiceStub(channel.Channel())
        request = backup_pb2.DeleteFilestoreAlphaBackupRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.source_instance):
            request.resource.source_instance = Primitive.to_proto(self.source_instance)

        if Primitive.to_proto(self.source_file_share):
            request.resource.source_file_share = Primitive.to_proto(
                self.source_file_share
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteFilestoreAlphaBackup(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = backup_pb2_grpc.FilestoreAlphaBackupServiceStub(channel.Channel())
        request = backup_pb2.ListFilestoreAlphaBackupRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListFilestoreAlphaBackup(request).items

    def to_proto(self):
        resource = backup_pb2.FilestoreAlphaBackup()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.source_instance):
            resource.source_instance = Primitive.to_proto(self.source_instance)
        if Primitive.to_proto(self.source_file_share):
            resource.source_file_share = Primitive.to_proto(self.source_file_share)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class BackupStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return backup_pb2.FilestoreAlphaBackupStateEnum.Value(
            "FilestoreAlphaBackupStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return backup_pb2.FilestoreAlphaBackupStateEnum.Name(resource)[
            len("FilestoreAlphaBackupStateEnum") :
        ]


class BackupSourceInstanceTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return backup_pb2.FilestoreAlphaBackupSourceInstanceTierEnum.Value(
            "FilestoreAlphaBackupSourceInstanceTierEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return backup_pb2.FilestoreAlphaBackupSourceInstanceTierEnum.Name(resource)[
            len("FilestoreAlphaBackupSourceInstanceTierEnum") :
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
