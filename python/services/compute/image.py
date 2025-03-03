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
from google3.cloud.graphite.mmv2.services.google.compute import image_pb2
from google3.cloud.graphite.mmv2.services.google.compute import image_pb2_grpc

from typing import List


class Image(object):
    def __init__(
        self,
        archive_size_bytes: int = None,
        description: str = None,
        disk_size_gb: int = None,
        family: str = None,
        guest_os_feature: list = None,
        image_encryption_key: dict = None,
        labels: dict = None,
        license: list = None,
        name: str = None,
        raw_disk: dict = None,
        shielded_instance_initial_state: dict = None,
        self_link: str = None,
        source_disk: str = None,
        source_disk_encryption_key: dict = None,
        source_disk_id: str = None,
        source_image: str = None,
        source_image_encryption_key: dict = None,
        source_image_id: str = None,
        source_snapshot: str = None,
        source_snapshot_encryption_key: dict = None,
        source_snapshot_id: str = None,
        source_type: str = None,
        status: str = None,
        storage_location: list = None,
        deprecated: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.description = description
        self.disk_size_gb = disk_size_gb
        self.family = family
        self.guest_os_feature = guest_os_feature
        self.image_encryption_key = image_encryption_key
        self.labels = labels
        self.license = license
        self.name = name
        self.raw_disk = raw_disk
        self.shielded_instance_initial_state = shielded_instance_initial_state
        self.source_disk = source_disk
        self.source_disk_encryption_key = source_disk_encryption_key
        self.source_image = source_image
        self.source_image_encryption_key = source_image_encryption_key
        self.source_image_id = source_image_id
        self.source_snapshot = source_snapshot
        self.source_snapshot_encryption_key = source_snapshot_encryption_key
        self.source_snapshot_id = source_snapshot_id
        self.source_type = source_type
        self.storage_location = storage_location
        self.deprecated = deprecated
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = image_pb2_grpc.ComputeImageServiceStub(channel.Channel())
        request = image_pb2.ApplyComputeImageRequest()
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.disk_size_gb):
            request.resource.disk_size_gb = Primitive.to_proto(self.disk_size_gb)

        if Primitive.to_proto(self.family):
            request.resource.family = Primitive.to_proto(self.family)

        if ImageGuestOSFeatureArray.to_proto(self.guest_os_feature):
            request.resource.guest_os_feature.extend(
                ImageGuestOSFeatureArray.to_proto(self.guest_os_feature)
            )
        if ImageImageEncryptionKey.to_proto(self.image_encryption_key):
            request.resource.image_encryption_key.CopyFrom(
                ImageImageEncryptionKey.to_proto(self.image_encryption_key)
            )
        else:
            request.resource.ClearField("image_encryption_key")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.license):
            request.resource.license.extend(Primitive.to_proto(self.license))
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ImageRawDisk.to_proto(self.raw_disk):
            request.resource.raw_disk.CopyFrom(ImageRawDisk.to_proto(self.raw_disk))
        else:
            request.resource.ClearField("raw_disk")
        if ImageShieldedInstanceInitialState.to_proto(
            self.shielded_instance_initial_state
        ):
            request.resource.shielded_instance_initial_state.CopyFrom(
                ImageShieldedInstanceInitialState.to_proto(
                    self.shielded_instance_initial_state
                )
            )
        else:
            request.resource.ClearField("shielded_instance_initial_state")
        if Primitive.to_proto(self.source_disk):
            request.resource.source_disk = Primitive.to_proto(self.source_disk)

        if ImageSourceDiskEncryptionKey.to_proto(self.source_disk_encryption_key):
            request.resource.source_disk_encryption_key.CopyFrom(
                ImageSourceDiskEncryptionKey.to_proto(self.source_disk_encryption_key)
            )
        else:
            request.resource.ClearField("source_disk_encryption_key")
        if Primitive.to_proto(self.source_image):
            request.resource.source_image = Primitive.to_proto(self.source_image)

        if ImageSourceImageEncryptionKey.to_proto(self.source_image_encryption_key):
            request.resource.source_image_encryption_key.CopyFrom(
                ImageSourceImageEncryptionKey.to_proto(self.source_image_encryption_key)
            )
        else:
            request.resource.ClearField("source_image_encryption_key")
        if Primitive.to_proto(self.source_image_id):
            request.resource.source_image_id = Primitive.to_proto(self.source_image_id)

        if Primitive.to_proto(self.source_snapshot):
            request.resource.source_snapshot = Primitive.to_proto(self.source_snapshot)

        if ImageSourceSnapshotEncryptionKey.to_proto(
            self.source_snapshot_encryption_key
        ):
            request.resource.source_snapshot_encryption_key.CopyFrom(
                ImageSourceSnapshotEncryptionKey.to_proto(
                    self.source_snapshot_encryption_key
                )
            )
        else:
            request.resource.ClearField("source_snapshot_encryption_key")
        if Primitive.to_proto(self.source_snapshot_id):
            request.resource.source_snapshot_id = Primitive.to_proto(
                self.source_snapshot_id
            )

        if ImageSourceTypeEnum.to_proto(self.source_type):
            request.resource.source_type = ImageSourceTypeEnum.to_proto(
                self.source_type
            )

        if Primitive.to_proto(self.storage_location):
            request.resource.storage_location.extend(
                Primitive.to_proto(self.storage_location)
            )
        if ImageDeprecated.to_proto(self.deprecated):
            request.resource.deprecated.CopyFrom(
                ImageDeprecated.to_proto(self.deprecated)
            )
        else:
            request.resource.ClearField("deprecated")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeImage(request)
        self.archive_size_bytes = Primitive.from_proto(response.archive_size_bytes)
        self.description = Primitive.from_proto(response.description)
        self.disk_size_gb = Primitive.from_proto(response.disk_size_gb)
        self.family = Primitive.from_proto(response.family)
        self.guest_os_feature = ImageGuestOSFeatureArray.from_proto(
            response.guest_os_feature
        )
        self.image_encryption_key = ImageImageEncryptionKey.from_proto(
            response.image_encryption_key
        )
        self.labels = Primitive.from_proto(response.labels)
        self.license = Primitive.from_proto(response.license)
        self.name = Primitive.from_proto(response.name)
        self.raw_disk = ImageRawDisk.from_proto(response.raw_disk)
        self.shielded_instance_initial_state = ImageShieldedInstanceInitialState.from_proto(
            response.shielded_instance_initial_state
        )
        self.self_link = Primitive.from_proto(response.self_link)
        self.source_disk = Primitive.from_proto(response.source_disk)
        self.source_disk_encryption_key = ImageSourceDiskEncryptionKey.from_proto(
            response.source_disk_encryption_key
        )
        self.source_disk_id = Primitive.from_proto(response.source_disk_id)
        self.source_image = Primitive.from_proto(response.source_image)
        self.source_image_encryption_key = ImageSourceImageEncryptionKey.from_proto(
            response.source_image_encryption_key
        )
        self.source_image_id = Primitive.from_proto(response.source_image_id)
        self.source_snapshot = Primitive.from_proto(response.source_snapshot)
        self.source_snapshot_encryption_key = ImageSourceSnapshotEncryptionKey.from_proto(
            response.source_snapshot_encryption_key
        )
        self.source_snapshot_id = Primitive.from_proto(response.source_snapshot_id)
        self.source_type = ImageSourceTypeEnum.from_proto(response.source_type)
        self.status = ImageStatusEnum.from_proto(response.status)
        self.storage_location = Primitive.from_proto(response.storage_location)
        self.deprecated = ImageDeprecated.from_proto(response.deprecated)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = image_pb2_grpc.ComputeImageServiceStub(channel.Channel())
        request = image_pb2.DeleteComputeImageRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.disk_size_gb):
            request.resource.disk_size_gb = Primitive.to_proto(self.disk_size_gb)

        if Primitive.to_proto(self.family):
            request.resource.family = Primitive.to_proto(self.family)

        if ImageGuestOSFeatureArray.to_proto(self.guest_os_feature):
            request.resource.guest_os_feature.extend(
                ImageGuestOSFeatureArray.to_proto(self.guest_os_feature)
            )
        if ImageImageEncryptionKey.to_proto(self.image_encryption_key):
            request.resource.image_encryption_key.CopyFrom(
                ImageImageEncryptionKey.to_proto(self.image_encryption_key)
            )
        else:
            request.resource.ClearField("image_encryption_key")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.license):
            request.resource.license.extend(Primitive.to_proto(self.license))
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ImageRawDisk.to_proto(self.raw_disk):
            request.resource.raw_disk.CopyFrom(ImageRawDisk.to_proto(self.raw_disk))
        else:
            request.resource.ClearField("raw_disk")
        if ImageShieldedInstanceInitialState.to_proto(
            self.shielded_instance_initial_state
        ):
            request.resource.shielded_instance_initial_state.CopyFrom(
                ImageShieldedInstanceInitialState.to_proto(
                    self.shielded_instance_initial_state
                )
            )
        else:
            request.resource.ClearField("shielded_instance_initial_state")
        if Primitive.to_proto(self.source_disk):
            request.resource.source_disk = Primitive.to_proto(self.source_disk)

        if ImageSourceDiskEncryptionKey.to_proto(self.source_disk_encryption_key):
            request.resource.source_disk_encryption_key.CopyFrom(
                ImageSourceDiskEncryptionKey.to_proto(self.source_disk_encryption_key)
            )
        else:
            request.resource.ClearField("source_disk_encryption_key")
        if Primitive.to_proto(self.source_image):
            request.resource.source_image = Primitive.to_proto(self.source_image)

        if ImageSourceImageEncryptionKey.to_proto(self.source_image_encryption_key):
            request.resource.source_image_encryption_key.CopyFrom(
                ImageSourceImageEncryptionKey.to_proto(self.source_image_encryption_key)
            )
        else:
            request.resource.ClearField("source_image_encryption_key")
        if Primitive.to_proto(self.source_image_id):
            request.resource.source_image_id = Primitive.to_proto(self.source_image_id)

        if Primitive.to_proto(self.source_snapshot):
            request.resource.source_snapshot = Primitive.to_proto(self.source_snapshot)

        if ImageSourceSnapshotEncryptionKey.to_proto(
            self.source_snapshot_encryption_key
        ):
            request.resource.source_snapshot_encryption_key.CopyFrom(
                ImageSourceSnapshotEncryptionKey.to_proto(
                    self.source_snapshot_encryption_key
                )
            )
        else:
            request.resource.ClearField("source_snapshot_encryption_key")
        if Primitive.to_proto(self.source_snapshot_id):
            request.resource.source_snapshot_id = Primitive.to_proto(
                self.source_snapshot_id
            )

        if ImageSourceTypeEnum.to_proto(self.source_type):
            request.resource.source_type = ImageSourceTypeEnum.to_proto(
                self.source_type
            )

        if Primitive.to_proto(self.storage_location):
            request.resource.storage_location.extend(
                Primitive.to_proto(self.storage_location)
            )
        if ImageDeprecated.to_proto(self.deprecated):
            request.resource.deprecated.CopyFrom(
                ImageDeprecated.to_proto(self.deprecated)
            )
        else:
            request.resource.ClearField("deprecated")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeImage(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = image_pb2_grpc.ComputeImageServiceStub(channel.Channel())
        request = image_pb2.ListComputeImageRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeImage(request).items

    def to_proto(self):
        resource = image_pb2.ComputeImage()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.disk_size_gb):
            resource.disk_size_gb = Primitive.to_proto(self.disk_size_gb)
        if Primitive.to_proto(self.family):
            resource.family = Primitive.to_proto(self.family)
        if ImageGuestOSFeatureArray.to_proto(self.guest_os_feature):
            resource.guest_os_feature.extend(
                ImageGuestOSFeatureArray.to_proto(self.guest_os_feature)
            )
        if ImageImageEncryptionKey.to_proto(self.image_encryption_key):
            resource.image_encryption_key.CopyFrom(
                ImageImageEncryptionKey.to_proto(self.image_encryption_key)
            )
        else:
            resource.ClearField("image_encryption_key")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.license):
            resource.license.extend(Primitive.to_proto(self.license))
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if ImageRawDisk.to_proto(self.raw_disk):
            resource.raw_disk.CopyFrom(ImageRawDisk.to_proto(self.raw_disk))
        else:
            resource.ClearField("raw_disk")
        if ImageShieldedInstanceInitialState.to_proto(
            self.shielded_instance_initial_state
        ):
            resource.shielded_instance_initial_state.CopyFrom(
                ImageShieldedInstanceInitialState.to_proto(
                    self.shielded_instance_initial_state
                )
            )
        else:
            resource.ClearField("shielded_instance_initial_state")
        if Primitive.to_proto(self.source_disk):
            resource.source_disk = Primitive.to_proto(self.source_disk)
        if ImageSourceDiskEncryptionKey.to_proto(self.source_disk_encryption_key):
            resource.source_disk_encryption_key.CopyFrom(
                ImageSourceDiskEncryptionKey.to_proto(self.source_disk_encryption_key)
            )
        else:
            resource.ClearField("source_disk_encryption_key")
        if Primitive.to_proto(self.source_image):
            resource.source_image = Primitive.to_proto(self.source_image)
        if ImageSourceImageEncryptionKey.to_proto(self.source_image_encryption_key):
            resource.source_image_encryption_key.CopyFrom(
                ImageSourceImageEncryptionKey.to_proto(self.source_image_encryption_key)
            )
        else:
            resource.ClearField("source_image_encryption_key")
        if Primitive.to_proto(self.source_image_id):
            resource.source_image_id = Primitive.to_proto(self.source_image_id)
        if Primitive.to_proto(self.source_snapshot):
            resource.source_snapshot = Primitive.to_proto(self.source_snapshot)
        if ImageSourceSnapshotEncryptionKey.to_proto(
            self.source_snapshot_encryption_key
        ):
            resource.source_snapshot_encryption_key.CopyFrom(
                ImageSourceSnapshotEncryptionKey.to_proto(
                    self.source_snapshot_encryption_key
                )
            )
        else:
            resource.ClearField("source_snapshot_encryption_key")
        if Primitive.to_proto(self.source_snapshot_id):
            resource.source_snapshot_id = Primitive.to_proto(self.source_snapshot_id)
        if ImageSourceTypeEnum.to_proto(self.source_type):
            resource.source_type = ImageSourceTypeEnum.to_proto(self.source_type)
        if Primitive.to_proto(self.storage_location):
            resource.storage_location.extend(Primitive.to_proto(self.storage_location))
        if ImageDeprecated.to_proto(self.deprecated):
            resource.deprecated.CopyFrom(ImageDeprecated.to_proto(self.deprecated))
        else:
            resource.ClearField("deprecated")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class ImageGuestOSFeature(object):
    def __init__(self, type: str = None):
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageGuestOSFeature()
        if ImageGuestOSFeatureTypeEnum.to_proto(resource.type):
            res.type = ImageGuestOSFeatureTypeEnum.to_proto(resource.type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageGuestOSFeature(
            type=ImageGuestOSFeatureTypeEnum.from_proto(resource.type),
        )


class ImageGuestOSFeatureArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageGuestOSFeature.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageGuestOSFeature.from_proto(i) for i in resources]


class ImageImageEncryptionKey(object):
    def __init__(
        self,
        raw_key: str = None,
        kms_key_name: str = None,
        sha256: str = None,
        kms_key_service_account: str = None,
    ):
        self.raw_key = raw_key
        self.kms_key_name = kms_key_name
        self.sha256 = sha256
        self.kms_key_service_account = kms_key_service_account

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageImageEncryptionKey()
        if Primitive.to_proto(resource.raw_key):
            res.raw_key = Primitive.to_proto(resource.raw_key)
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        if Primitive.to_proto(resource.sha256):
            res.sha256 = Primitive.to_proto(resource.sha256)
        if Primitive.to_proto(resource.kms_key_service_account):
            res.kms_key_service_account = Primitive.to_proto(
                resource.kms_key_service_account
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageImageEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
            sha256=Primitive.from_proto(resource.sha256),
            kms_key_service_account=Primitive.from_proto(
                resource.kms_key_service_account
            ),
        )


class ImageImageEncryptionKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageImageEncryptionKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageImageEncryptionKey.from_proto(i) for i in resources]


class ImageRawDisk(object):
    def __init__(
        self, source: str = None, sha1_checksum: str = None, container_type: str = None
    ):
        self.source = source
        self.sha1_checksum = sha1_checksum
        self.container_type = container_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageRawDisk()
        if Primitive.to_proto(resource.source):
            res.source = Primitive.to_proto(resource.source)
        if Primitive.to_proto(resource.sha1_checksum):
            res.sha1_checksum = Primitive.to_proto(resource.sha1_checksum)
        if ImageRawDiskContainerTypeEnum.to_proto(resource.container_type):
            res.container_type = ImageRawDiskContainerTypeEnum.to_proto(
                resource.container_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageRawDisk(
            source=Primitive.from_proto(resource.source),
            sha1_checksum=Primitive.from_proto(resource.sha1_checksum),
            container_type=ImageRawDiskContainerTypeEnum.from_proto(
                resource.container_type
            ),
        )


class ImageRawDiskArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageRawDisk.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageRawDisk.from_proto(i) for i in resources]


class ImageShieldedInstanceInitialState(object):
    def __init__(
        self, pk: dict = None, kek: list = None, db: list = None, dbx: list = None
    ):
        self.pk = pk
        self.kek = kek
        self.db = db
        self.dbx = dbx

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageShieldedInstanceInitialState()
        if ImageShieldedInstanceInitialStatePk.to_proto(resource.pk):
            res.pk.CopyFrom(ImageShieldedInstanceInitialStatePk.to_proto(resource.pk))
        else:
            res.ClearField("pk")
        if ImageShieldedInstanceInitialStateKekArray.to_proto(resource.kek):
            res.kek.extend(
                ImageShieldedInstanceInitialStateKekArray.to_proto(resource.kek)
            )
        if ImageShieldedInstanceInitialStateDbArray.to_proto(resource.db):
            res.db.extend(
                ImageShieldedInstanceInitialStateDbArray.to_proto(resource.db)
            )
        if ImageShieldedInstanceInitialStateDbxArray.to_proto(resource.dbx):
            res.dbx.extend(
                ImageShieldedInstanceInitialStateDbxArray.to_proto(resource.dbx)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageShieldedInstanceInitialState(
            pk=ImageShieldedInstanceInitialStatePk.from_proto(resource.pk),
            kek=ImageShieldedInstanceInitialStateKekArray.from_proto(resource.kek),
            db=ImageShieldedInstanceInitialStateDbArray.from_proto(resource.db),
            dbx=ImageShieldedInstanceInitialStateDbxArray.from_proto(resource.dbx),
        )


class ImageShieldedInstanceInitialStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageShieldedInstanceInitialState.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageShieldedInstanceInitialState.from_proto(i) for i in resources]


class ImageShieldedInstanceInitialStatePk(object):
    def __init__(self, content: str = None, file_type: str = None):
        self.content = content
        self.file_type = file_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageShieldedInstanceInitialStatePk()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if ImageShieldedInstanceInitialStatePkFileTypeEnum.to_proto(resource.file_type):
            res.file_type = ImageShieldedInstanceInitialStatePkFileTypeEnum.to_proto(
                resource.file_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageShieldedInstanceInitialStatePk(
            content=Primitive.from_proto(resource.content),
            file_type=ImageShieldedInstanceInitialStatePkFileTypeEnum.from_proto(
                resource.file_type
            ),
        )


class ImageShieldedInstanceInitialStatePkArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageShieldedInstanceInitialStatePk.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageShieldedInstanceInitialStatePk.from_proto(i) for i in resources]


class ImageShieldedInstanceInitialStateKek(object):
    def __init__(self, content: str = None, file_type: str = None):
        self.content = content
        self.file_type = file_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageShieldedInstanceInitialStateKek()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if ImageShieldedInstanceInitialStateKekFileTypeEnum.to_proto(
            resource.file_type
        ):
            res.file_type = ImageShieldedInstanceInitialStateKekFileTypeEnum.to_proto(
                resource.file_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageShieldedInstanceInitialStateKek(
            content=Primitive.from_proto(resource.content),
            file_type=ImageShieldedInstanceInitialStateKekFileTypeEnum.from_proto(
                resource.file_type
            ),
        )


class ImageShieldedInstanceInitialStateKekArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageShieldedInstanceInitialStateKek.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageShieldedInstanceInitialStateKek.from_proto(i) for i in resources]


class ImageShieldedInstanceInitialStateDb(object):
    def __init__(self, content: str = None, file_type: str = None):
        self.content = content
        self.file_type = file_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageShieldedInstanceInitialStateDb()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if ImageShieldedInstanceInitialStateDbFileTypeEnum.to_proto(resource.file_type):
            res.file_type = ImageShieldedInstanceInitialStateDbFileTypeEnum.to_proto(
                resource.file_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageShieldedInstanceInitialStateDb(
            content=Primitive.from_proto(resource.content),
            file_type=ImageShieldedInstanceInitialStateDbFileTypeEnum.from_proto(
                resource.file_type
            ),
        )


class ImageShieldedInstanceInitialStateDbArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageShieldedInstanceInitialStateDb.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageShieldedInstanceInitialStateDb.from_proto(i) for i in resources]


class ImageShieldedInstanceInitialStateDbx(object):
    def __init__(self, content: str = None, file_type: str = None):
        self.content = content
        self.file_type = file_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageShieldedInstanceInitialStateDbx()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if ImageShieldedInstanceInitialStateDbxFileTypeEnum.to_proto(
            resource.file_type
        ):
            res.file_type = ImageShieldedInstanceInitialStateDbxFileTypeEnum.to_proto(
                resource.file_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageShieldedInstanceInitialStateDbx(
            content=Primitive.from_proto(resource.content),
            file_type=ImageShieldedInstanceInitialStateDbxFileTypeEnum.from_proto(
                resource.file_type
            ),
        )


class ImageShieldedInstanceInitialStateDbxArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageShieldedInstanceInitialStateDbx.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageShieldedInstanceInitialStateDbx.from_proto(i) for i in resources]


class ImageSourceDiskEncryptionKey(object):
    def __init__(
        self,
        raw_key: str = None,
        kms_key_name: str = None,
        sha256: str = None,
        kms_key_service_account: str = None,
    ):
        self.raw_key = raw_key
        self.kms_key_name = kms_key_name
        self.sha256 = sha256
        self.kms_key_service_account = kms_key_service_account

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageSourceDiskEncryptionKey()
        if Primitive.to_proto(resource.raw_key):
            res.raw_key = Primitive.to_proto(resource.raw_key)
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        if Primitive.to_proto(resource.sha256):
            res.sha256 = Primitive.to_proto(resource.sha256)
        if Primitive.to_proto(resource.kms_key_service_account):
            res.kms_key_service_account = Primitive.to_proto(
                resource.kms_key_service_account
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageSourceDiskEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
            sha256=Primitive.from_proto(resource.sha256),
            kms_key_service_account=Primitive.from_proto(
                resource.kms_key_service_account
            ),
        )


class ImageSourceDiskEncryptionKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageSourceDiskEncryptionKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageSourceDiskEncryptionKey.from_proto(i) for i in resources]


class ImageSourceImageEncryptionKey(object):
    def __init__(
        self,
        raw_key: str = None,
        kms_key_name: str = None,
        sha256: str = None,
        kms_key_service_account: str = None,
    ):
        self.raw_key = raw_key
        self.kms_key_name = kms_key_name
        self.sha256 = sha256
        self.kms_key_service_account = kms_key_service_account

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageSourceImageEncryptionKey()
        if Primitive.to_proto(resource.raw_key):
            res.raw_key = Primitive.to_proto(resource.raw_key)
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        if Primitive.to_proto(resource.sha256):
            res.sha256 = Primitive.to_proto(resource.sha256)
        if Primitive.to_proto(resource.kms_key_service_account):
            res.kms_key_service_account = Primitive.to_proto(
                resource.kms_key_service_account
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageSourceImageEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
            sha256=Primitive.from_proto(resource.sha256),
            kms_key_service_account=Primitive.from_proto(
                resource.kms_key_service_account
            ),
        )


class ImageSourceImageEncryptionKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageSourceImageEncryptionKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageSourceImageEncryptionKey.from_proto(i) for i in resources]


class ImageSourceSnapshotEncryptionKey(object):
    def __init__(
        self,
        raw_key: str = None,
        kms_key_name: str = None,
        sha256: str = None,
        kms_key_service_account: str = None,
    ):
        self.raw_key = raw_key
        self.kms_key_name = kms_key_name
        self.sha256 = sha256
        self.kms_key_service_account = kms_key_service_account

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageSourceSnapshotEncryptionKey()
        if Primitive.to_proto(resource.raw_key):
            res.raw_key = Primitive.to_proto(resource.raw_key)
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        if Primitive.to_proto(resource.sha256):
            res.sha256 = Primitive.to_proto(resource.sha256)
        if Primitive.to_proto(resource.kms_key_service_account):
            res.kms_key_service_account = Primitive.to_proto(
                resource.kms_key_service_account
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageSourceSnapshotEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
            sha256=Primitive.from_proto(resource.sha256),
            kms_key_service_account=Primitive.from_proto(
                resource.kms_key_service_account
            ),
        )


class ImageSourceSnapshotEncryptionKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageSourceSnapshotEncryptionKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageSourceSnapshotEncryptionKey.from_proto(i) for i in resources]


class ImageDeprecated(object):
    def __init__(
        self,
        state: str = None,
        replacement: str = None,
        deprecated: str = None,
        obsolete: str = None,
        deleted: str = None,
    ):
        self.state = state
        self.replacement = replacement
        self.deprecated = deprecated
        self.obsolete = obsolete
        self.deleted = deleted

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = image_pb2.ComputeImageDeprecated()
        if ImageDeprecatedStateEnum.to_proto(resource.state):
            res.state = ImageDeprecatedStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.replacement):
            res.replacement = Primitive.to_proto(resource.replacement)
        if Primitive.to_proto(resource.deprecated):
            res.deprecated = Primitive.to_proto(resource.deprecated)
        if Primitive.to_proto(resource.obsolete):
            res.obsolete = Primitive.to_proto(resource.obsolete)
        if Primitive.to_proto(resource.deleted):
            res.deleted = Primitive.to_proto(resource.deleted)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ImageDeprecated(
            state=ImageDeprecatedStateEnum.from_proto(resource.state),
            replacement=Primitive.from_proto(resource.replacement),
            deprecated=Primitive.from_proto(resource.deprecated),
            obsolete=Primitive.from_proto(resource.obsolete),
            deleted=Primitive.from_proto(resource.deleted),
        )


class ImageDeprecatedArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ImageDeprecated.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ImageDeprecated.from_proto(i) for i in resources]


class ImageGuestOSFeatureTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageGuestOSFeatureTypeEnum.Value(
            "ComputeImageGuestOSFeatureTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageGuestOSFeatureTypeEnum.Name(resource)[
            len("ComputeImageGuestOSFeatureTypeEnum") :
        ]


class ImageRawDiskContainerTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageRawDiskContainerTypeEnum.Value(
            "ComputeImageRawDiskContainerTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageRawDiskContainerTypeEnum.Name(resource)[
            len("ComputeImageRawDiskContainerTypeEnum") :
        ]


class ImageShieldedInstanceInitialStatePkFileTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageShieldedInstanceInitialStatePkFileTypeEnum.Value(
            "ComputeImageShieldedInstanceInitialStatePkFileTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageShieldedInstanceInitialStatePkFileTypeEnum.Name(
            resource
        )[len("ComputeImageShieldedInstanceInitialStatePkFileTypeEnum") :]


class ImageShieldedInstanceInitialStateKekFileTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageShieldedInstanceInitialStateKekFileTypeEnum.Value(
            "ComputeImageShieldedInstanceInitialStateKekFileTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageShieldedInstanceInitialStateKekFileTypeEnum.Name(
            resource
        )[len("ComputeImageShieldedInstanceInitialStateKekFileTypeEnum") :]


class ImageShieldedInstanceInitialStateDbFileTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageShieldedInstanceInitialStateDbFileTypeEnum.Value(
            "ComputeImageShieldedInstanceInitialStateDbFileTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageShieldedInstanceInitialStateDbFileTypeEnum.Name(
            resource
        )[len("ComputeImageShieldedInstanceInitialStateDbFileTypeEnum") :]


class ImageShieldedInstanceInitialStateDbxFileTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum.Value(
            "ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum.Name(
            resource
        )[len("ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum") :]


class ImageSourceTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageSourceTypeEnum.Value(
            "ComputeImageSourceTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageSourceTypeEnum.Name(resource)[
            len("ComputeImageSourceTypeEnum") :
        ]


class ImageStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageStatusEnum.Value(
            "ComputeImageStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageStatusEnum.Name(resource)[
            len("ComputeImageStatusEnum") :
        ]


class ImageDeprecatedStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageDeprecatedStateEnum.Value(
            "ComputeImageDeprecatedStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return image_pb2.ComputeImageDeprecatedStateEnum.Name(resource)[
            len("ComputeImageDeprecatedStateEnum") :
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
