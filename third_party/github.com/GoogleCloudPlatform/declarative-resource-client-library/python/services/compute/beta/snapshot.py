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
from google3.cloud.graphite.mmv2.services.google.compute import snapshot_pb2
from google3.cloud.graphite.mmv2.services.google.compute import snapshot_pb2_grpc

from typing import List


class Snapshot(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        source_disk: str = None,
        disk_size_gb: int = None,
        storage_bytes: int = None,
        license: list = None,
        snapshot_encryption_key: dict = None,
        source_disk_encryption_key: dict = None,
        self_link: str = None,
        labels: dict = None,
        project: str = None,
        zone: str = None,
        id: int = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.source_disk = source_disk
        self.snapshot_encryption_key = snapshot_encryption_key
        self.source_disk_encryption_key = source_disk_encryption_key
        self.labels = labels
        self.project = project
        self.zone = zone
        self.service_account_file = service_account_file

    def apply(self):
        stub = snapshot_pb2_grpc.ComputeBetaSnapshotServiceStub(channel.Channel())
        request = snapshot_pb2.ApplyComputeBetaSnapshotRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.source_disk):
            request.resource.source_disk = Primitive.to_proto(self.source_disk)

        if SnapshotSnapshotEncryptionKey.to_proto(self.snapshot_encryption_key):
            request.resource.snapshot_encryption_key.CopyFrom(
                SnapshotSnapshotEncryptionKey.to_proto(self.snapshot_encryption_key)
            )
        else:
            request.resource.ClearField("snapshot_encryption_key")
        if SnapshotSourceDiskEncryptionKey.to_proto(self.source_disk_encryption_key):
            request.resource.source_disk_encryption_key.CopyFrom(
                SnapshotSourceDiskEncryptionKey.to_proto(
                    self.source_disk_encryption_key
                )
            )
        else:
            request.resource.ClearField("source_disk_encryption_key")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaSnapshot(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.source_disk = Primitive.from_proto(response.source_disk)
        self.disk_size_gb = Primitive.from_proto(response.disk_size_gb)
        self.storage_bytes = Primitive.from_proto(response.storage_bytes)
        self.license = Primitive.from_proto(response.license)
        self.snapshot_encryption_key = SnapshotSnapshotEncryptionKey.from_proto(
            response.snapshot_encryption_key
        )
        self.source_disk_encryption_key = SnapshotSourceDiskEncryptionKey.from_proto(
            response.source_disk_encryption_key
        )
        self.self_link = Primitive.from_proto(response.self_link)
        self.labels = Primitive.from_proto(response.labels)
        self.project = Primitive.from_proto(response.project)
        self.zone = Primitive.from_proto(response.zone)
        self.id = Primitive.from_proto(response.id)

    def delete(self):
        stub = snapshot_pb2_grpc.ComputeBetaSnapshotServiceStub(channel.Channel())
        request = snapshot_pb2.DeleteComputeBetaSnapshotRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.source_disk):
            request.resource.source_disk = Primitive.to_proto(self.source_disk)

        if SnapshotSnapshotEncryptionKey.to_proto(self.snapshot_encryption_key):
            request.resource.snapshot_encryption_key.CopyFrom(
                SnapshotSnapshotEncryptionKey.to_proto(self.snapshot_encryption_key)
            )
        else:
            request.resource.ClearField("snapshot_encryption_key")
        if SnapshotSourceDiskEncryptionKey.to_proto(self.source_disk_encryption_key):
            request.resource.source_disk_encryption_key.CopyFrom(
                SnapshotSourceDiskEncryptionKey.to_proto(
                    self.source_disk_encryption_key
                )
            )
        else:
            request.resource.ClearField("source_disk_encryption_key")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        response = stub.DeleteComputeBetaSnapshot(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = snapshot_pb2_grpc.ComputeBetaSnapshotServiceStub(channel.Channel())
        request = snapshot_pb2.ListComputeBetaSnapshotRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeBetaSnapshot(request).items

    def to_proto(self):
        resource = snapshot_pb2.ComputeBetaSnapshot()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.source_disk):
            resource.source_disk = Primitive.to_proto(self.source_disk)
        if SnapshotSnapshotEncryptionKey.to_proto(self.snapshot_encryption_key):
            resource.snapshot_encryption_key.CopyFrom(
                SnapshotSnapshotEncryptionKey.to_proto(self.snapshot_encryption_key)
            )
        else:
            resource.ClearField("snapshot_encryption_key")
        if SnapshotSourceDiskEncryptionKey.to_proto(self.source_disk_encryption_key):
            resource.source_disk_encryption_key.CopyFrom(
                SnapshotSourceDiskEncryptionKey.to_proto(
                    self.source_disk_encryption_key
                )
            )
        else:
            resource.ClearField("source_disk_encryption_key")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.zone):
            resource.zone = Primitive.to_proto(self.zone)
        return resource


class SnapshotSnapshotEncryptionKey(object):
    def __init__(self, raw_key: str = None, sha256: str = None):
        self.raw_key = raw_key
        self.sha256 = sha256

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = snapshot_pb2.ComputeBetaSnapshotSnapshotEncryptionKey()
        if Primitive.to_proto(resource.raw_key):
            res.raw_key = Primitive.to_proto(resource.raw_key)
        if Primitive.to_proto(resource.sha256):
            res.sha256 = Primitive.to_proto(resource.sha256)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SnapshotSnapshotEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
            sha256=Primitive.from_proto(resource.sha256),
        )


class SnapshotSnapshotEncryptionKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SnapshotSnapshotEncryptionKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SnapshotSnapshotEncryptionKey.from_proto(i) for i in resources]


class SnapshotSourceDiskEncryptionKey(object):
    def __init__(self, raw_key: str = None):
        self.raw_key = raw_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = snapshot_pb2.ComputeBetaSnapshotSourceDiskEncryptionKey()
        if Primitive.to_proto(resource.raw_key):
            res.raw_key = Primitive.to_proto(resource.raw_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SnapshotSourceDiskEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
        )


class SnapshotSourceDiskEncryptionKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SnapshotSourceDiskEncryptionKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SnapshotSourceDiskEncryptionKey.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
