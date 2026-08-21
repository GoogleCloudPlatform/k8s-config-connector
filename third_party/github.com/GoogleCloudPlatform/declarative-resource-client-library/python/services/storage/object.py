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
from google3.cloud.graphite.mmv2.services.google.storage import object_pb2
from google3.cloud.graphite.mmv2.services.google.storage import object_pb2_grpc

from typing import List


class Object(object):
    def __init__(
        self,
        name: str = None,
        bucket: str = None,
        generation: int = None,
        metageneration: int = None,
        id: str = None,
        self_link: str = None,
        content_type: str = None,
        time_created: str = None,
        updated: str = None,
        custom_time: str = None,
        time_deleted: str = None,
        temporary_hold: bool = None,
        event_based_hold: bool = None,
        retention_expiration_time: str = None,
        storage_class: str = None,
        time_storage_class_updated: str = None,
        size: int = None,
        md5_hash: str = None,
        media_link: str = None,
        metadata: dict = None,
        owner: dict = None,
        crc32c: str = None,
        component_count: int = None,
        etag: str = None,
        customer_encryption: dict = None,
        kms_key_name: str = None,
        content: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.bucket = bucket
        self.content_type = content_type
        self.custom_time = custom_time
        self.temporary_hold = temporary_hold
        self.event_based_hold = event_based_hold
        self.storage_class = storage_class
        self.md5_hash = md5_hash
        self.metadata = metadata
        self.crc32c = crc32c
        self.customer_encryption = customer_encryption
        self.kms_key_name = kms_key_name
        self.content = content
        self.service_account_file = service_account_file

    def apply(self):
        stub = object_pb2_grpc.StorageObjectServiceStub(channel.Channel())
        request = object_pb2.ApplyStorageObjectRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.bucket):
            request.resource.bucket = Primitive.to_proto(self.bucket)

        if Primitive.to_proto(self.content_type):
            request.resource.content_type = Primitive.to_proto(self.content_type)

        if Primitive.to_proto(self.custom_time):
            request.resource.custom_time = Primitive.to_proto(self.custom_time)

        if Primitive.to_proto(self.temporary_hold):
            request.resource.temporary_hold = Primitive.to_proto(self.temporary_hold)

        if Primitive.to_proto(self.event_based_hold):
            request.resource.event_based_hold = Primitive.to_proto(
                self.event_based_hold
            )

        if Primitive.to_proto(self.storage_class):
            request.resource.storage_class = Primitive.to_proto(self.storage_class)

        if Primitive.to_proto(self.md5_hash):
            request.resource.md5_hash = Primitive.to_proto(self.md5_hash)

        if Primitive.to_proto(self.metadata):
            request.resource.metadata = Primitive.to_proto(self.metadata)

        if Primitive.to_proto(self.crc32c):
            request.resource.crc32c = Primitive.to_proto(self.crc32c)

        if ObjectCustomerEncryption.to_proto(self.customer_encryption):
            request.resource.customer_encryption.CopyFrom(
                ObjectCustomerEncryption.to_proto(self.customer_encryption)
            )
        else:
            request.resource.ClearField("customer_encryption")
        if Primitive.to_proto(self.kms_key_name):
            request.resource.kms_key_name = Primitive.to_proto(self.kms_key_name)

        if Primitive.to_proto(self.content):
            request.resource.content = Primitive.to_proto(self.content)

        request.service_account_file = self.service_account_file

        response = stub.ApplyStorageObject(request)
        self.name = Primitive.from_proto(response.name)
        self.bucket = Primitive.from_proto(response.bucket)
        self.generation = Primitive.from_proto(response.generation)
        self.metageneration = Primitive.from_proto(response.metageneration)
        self.id = Primitive.from_proto(response.id)
        self.self_link = Primitive.from_proto(response.self_link)
        self.content_type = Primitive.from_proto(response.content_type)
        self.time_created = Primitive.from_proto(response.time_created)
        self.updated = Primitive.from_proto(response.updated)
        self.custom_time = Primitive.from_proto(response.custom_time)
        self.time_deleted = Primitive.from_proto(response.time_deleted)
        self.temporary_hold = Primitive.from_proto(response.temporary_hold)
        self.event_based_hold = Primitive.from_proto(response.event_based_hold)
        self.retention_expiration_time = Primitive.from_proto(
            response.retention_expiration_time
        )
        self.storage_class = Primitive.from_proto(response.storage_class)
        self.time_storage_class_updated = Primitive.from_proto(
            response.time_storage_class_updated
        )
        self.size = Primitive.from_proto(response.size)
        self.md5_hash = Primitive.from_proto(response.md5_hash)
        self.media_link = Primitive.from_proto(response.media_link)
        self.metadata = Primitive.from_proto(response.metadata)
        self.owner = ObjectOwner.from_proto(response.owner)
        self.crc32c = Primitive.from_proto(response.crc32c)
        self.component_count = Primitive.from_proto(response.component_count)
        self.etag = Primitive.from_proto(response.etag)
        self.customer_encryption = ObjectCustomerEncryption.from_proto(
            response.customer_encryption
        )
        self.kms_key_name = Primitive.from_proto(response.kms_key_name)
        self.content = Primitive.from_proto(response.content)

    def delete(self):
        stub = object_pb2_grpc.StorageObjectServiceStub(channel.Channel())
        request = object_pb2.DeleteStorageObjectRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.bucket):
            request.resource.bucket = Primitive.to_proto(self.bucket)

        if Primitive.to_proto(self.content_type):
            request.resource.content_type = Primitive.to_proto(self.content_type)

        if Primitive.to_proto(self.custom_time):
            request.resource.custom_time = Primitive.to_proto(self.custom_time)

        if Primitive.to_proto(self.temporary_hold):
            request.resource.temporary_hold = Primitive.to_proto(self.temporary_hold)

        if Primitive.to_proto(self.event_based_hold):
            request.resource.event_based_hold = Primitive.to_proto(
                self.event_based_hold
            )

        if Primitive.to_proto(self.storage_class):
            request.resource.storage_class = Primitive.to_proto(self.storage_class)

        if Primitive.to_proto(self.md5_hash):
            request.resource.md5_hash = Primitive.to_proto(self.md5_hash)

        if Primitive.to_proto(self.metadata):
            request.resource.metadata = Primitive.to_proto(self.metadata)

        if Primitive.to_proto(self.crc32c):
            request.resource.crc32c = Primitive.to_proto(self.crc32c)

        if ObjectCustomerEncryption.to_proto(self.customer_encryption):
            request.resource.customer_encryption.CopyFrom(
                ObjectCustomerEncryption.to_proto(self.customer_encryption)
            )
        else:
            request.resource.ClearField("customer_encryption")
        if Primitive.to_proto(self.kms_key_name):
            request.resource.kms_key_name = Primitive.to_proto(self.kms_key_name)

        if Primitive.to_proto(self.content):
            request.resource.content = Primitive.to_proto(self.content)

        response = stub.DeleteStorageObject(request)

    @classmethod
    def list(self, bucket, service_account_file=""):
        stub = object_pb2_grpc.StorageObjectServiceStub(channel.Channel())
        request = object_pb2.ListStorageObjectRequest()
        request.service_account_file = service_account_file
        request.Bucket = bucket

        return stub.ListStorageObject(request).items

    def to_proto(self):
        resource = object_pb2.StorageObject()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.bucket):
            resource.bucket = Primitive.to_proto(self.bucket)
        if Primitive.to_proto(self.content_type):
            resource.content_type = Primitive.to_proto(self.content_type)
        if Primitive.to_proto(self.custom_time):
            resource.custom_time = Primitive.to_proto(self.custom_time)
        if Primitive.to_proto(self.temporary_hold):
            resource.temporary_hold = Primitive.to_proto(self.temporary_hold)
        if Primitive.to_proto(self.event_based_hold):
            resource.event_based_hold = Primitive.to_proto(self.event_based_hold)
        if Primitive.to_proto(self.storage_class):
            resource.storage_class = Primitive.to_proto(self.storage_class)
        if Primitive.to_proto(self.md5_hash):
            resource.md5_hash = Primitive.to_proto(self.md5_hash)
        if Primitive.to_proto(self.metadata):
            resource.metadata = Primitive.to_proto(self.metadata)
        if Primitive.to_proto(self.crc32c):
            resource.crc32c = Primitive.to_proto(self.crc32c)
        if ObjectCustomerEncryption.to_proto(self.customer_encryption):
            resource.customer_encryption.CopyFrom(
                ObjectCustomerEncryption.to_proto(self.customer_encryption)
            )
        else:
            resource.ClearField("customer_encryption")
        if Primitive.to_proto(self.kms_key_name):
            resource.kms_key_name = Primitive.to_proto(self.kms_key_name)
        if Primitive.to_proto(self.content):
            resource.content = Primitive.to_proto(self.content)
        return resource


class ObjectOwner(object):
    def __init__(self, entity: str = None, entity_id: str = None):
        self.entity = entity
        self.entity_id = entity_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = object_pb2.StorageObjectOwner()
        if Primitive.to_proto(resource.entity):
            res.entity = Primitive.to_proto(resource.entity)
        if Primitive.to_proto(resource.entity_id):
            res.entity_id = Primitive.to_proto(resource.entity_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ObjectOwner(
            entity=Primitive.from_proto(resource.entity),
            entity_id=Primitive.from_proto(resource.entity_id),
        )


class ObjectOwnerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ObjectOwner.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ObjectOwner.from_proto(i) for i in resources]


class ObjectCustomerEncryption(object):
    def __init__(
        self, encryption_algorithm: str = None, key_sha256: str = None, key: str = None
    ):
        self.encryption_algorithm = encryption_algorithm
        self.key_sha256 = key_sha256
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = object_pb2.StorageObjectCustomerEncryption()
        if Primitive.to_proto(resource.encryption_algorithm):
            res.encryption_algorithm = Primitive.to_proto(resource.encryption_algorithm)
        if Primitive.to_proto(resource.key_sha256):
            res.key_sha256 = Primitive.to_proto(resource.key_sha256)
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ObjectCustomerEncryption(
            encryption_algorithm=Primitive.from_proto(resource.encryption_algorithm),
            key_sha256=Primitive.from_proto(resource.key_sha256),
            key=Primitive.from_proto(resource.key),
        )


class ObjectCustomerEncryptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ObjectCustomerEncryption.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ObjectCustomerEncryption.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
