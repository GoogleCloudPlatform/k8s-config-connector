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
from google3.cloud.graphite.mmv2.services.google.vertex import metadata_store_pb2
from google3.cloud.graphite.mmv2.services.google.vertex import metadata_store_pb2_grpc

from typing import List


class MetadataStore(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        encryption_spec: dict = None,
        description: str = None,
        state: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.encryption_spec = encryption_spec
        self.description = description
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = metadata_store_pb2_grpc.VertexAlphaMetadataStoreServiceStub(
            channel.Channel()
        )
        request = metadata_store_pb2.ApplyVertexAlphaMetadataStoreRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if MetadataStoreEncryptionSpec.to_proto(self.encryption_spec):
            request.resource.encryption_spec.CopyFrom(
                MetadataStoreEncryptionSpec.to_proto(self.encryption_spec)
            )
        else:
            request.resource.ClearField("encryption_spec")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyVertexAlphaMetadataStore(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.encryption_spec = MetadataStoreEncryptionSpec.from_proto(
            response.encryption_spec
        )
        self.description = Primitive.from_proto(response.description)
        self.state = MetadataStoreState.from_proto(response.state)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = metadata_store_pb2_grpc.VertexAlphaMetadataStoreServiceStub(
            channel.Channel()
        )
        request = metadata_store_pb2.DeleteVertexAlphaMetadataStoreRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if MetadataStoreEncryptionSpec.to_proto(self.encryption_spec):
            request.resource.encryption_spec.CopyFrom(
                MetadataStoreEncryptionSpec.to_proto(self.encryption_spec)
            )
        else:
            request.resource.ClearField("encryption_spec")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteVertexAlphaMetadataStore(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = metadata_store_pb2_grpc.VertexAlphaMetadataStoreServiceStub(
            channel.Channel()
        )
        request = metadata_store_pb2.ListVertexAlphaMetadataStoreRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListVertexAlphaMetadataStore(request).items

    def to_proto(self):
        resource = metadata_store_pb2.VertexAlphaMetadataStore()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if MetadataStoreEncryptionSpec.to_proto(self.encryption_spec):
            resource.encryption_spec.CopyFrom(
                MetadataStoreEncryptionSpec.to_proto(self.encryption_spec)
            )
        else:
            resource.ClearField("encryption_spec")
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class MetadataStoreEncryptionSpec(object):
    def __init__(self, kms_key_name: str = None):
        self.kms_key_name = kms_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = metadata_store_pb2.VertexAlphaMetadataStoreEncryptionSpec()
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MetadataStoreEncryptionSpec(
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
        )


class MetadataStoreEncryptionSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MetadataStoreEncryptionSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MetadataStoreEncryptionSpec.from_proto(i) for i in resources]


class MetadataStoreState(object):
    def __init__(self, disk_utilization_bytes: int = None):
        self.disk_utilization_bytes = disk_utilization_bytes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = metadata_store_pb2.VertexAlphaMetadataStoreState()
        if Primitive.to_proto(resource.disk_utilization_bytes):
            res.disk_utilization_bytes = Primitive.to_proto(
                resource.disk_utilization_bytes
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MetadataStoreState(
            disk_utilization_bytes=Primitive.from_proto(
                resource.disk_utilization_bytes
            ),
        )


class MetadataStoreStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MetadataStoreState.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MetadataStoreState.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
