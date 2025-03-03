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
from google3.cloud.graphite.mmv2.services.google.compute import disk_pb2
from google3.cloud.graphite.mmv2.services.google.compute import disk_pb2_grpc

from typing import List


class Disk(object):
    def __init__(
        self,
        self_link: str = None,
        description: str = None,
        disk_encryption_key: dict = None,
        guest_os_feature: list = None,
        labels: dict = None,
        label_fingerprint: str = None,
        license: list = None,
        name: str = None,
        region: str = None,
        replica_zones: list = None,
        resource_policy: list = None,
        size_gb: int = None,
        source_image: str = None,
        source_image_encryption_key: dict = None,
        source_image_id: str = None,
        source_snapshot: str = None,
        source_snapshot_encryption_key: dict = None,
        source_snapshot_id: str = None,
        type: str = None,
        zone: str = None,
        project: str = None,
        id: int = None,
        status: str = None,
        options: str = None,
        licenses: list = None,
        guest_os_features: list = None,
        last_attach_timestamp: str = None,
        last_detach_timestamp: str = None,
        users: list = None,
        license_codes: list = None,
        physical_block_size_bytes: int = None,
        resource_policies: list = None,
        source_disk: str = None,
        source_disk_id: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.description = description
        self.disk_encryption_key = disk_encryption_key
        self.guest_os_feature = guest_os_feature
        self.labels = labels
        self.license = license
        self.name = name
        self.region = region
        self.replica_zones = replica_zones
        self.resource_policy = resource_policy
        self.size_gb = size_gb
        self.source_image = source_image
        self.source_image_encryption_key = source_image_encryption_key
        self.source_snapshot = source_snapshot
        self.source_snapshot_encryption_key = source_snapshot_encryption_key
        self.type = type
        self.project = project
        self.id = id
        self.options = options
        self.licenses = licenses
        self.guest_os_features = guest_os_features
        self.license_codes = license_codes
        self.physical_block_size_bytes = physical_block_size_bytes
        self.resource_policies = resource_policies
        self.source_disk = source_disk
        self.source_disk_id = source_disk_id
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = disk_pb2_grpc.ComputeBetaDiskServiceStub(channel.Channel())
        request = disk_pb2.ApplyComputeBetaDiskRequest()
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if DiskEncryptionKey.to_proto(self.disk_encryption_key):
            request.resource.disk_encryption_key.CopyFrom(
                DiskEncryptionKey.to_proto(self.disk_encryption_key)
            )
        else:
            request.resource.ClearField("disk_encryption_key")
        if DiskGuestOSFeatureArray.to_proto(self.guest_os_feature):
            request.resource.guest_os_feature.extend(
                DiskGuestOSFeatureArray.to_proto(self.guest_os_feature)
            )
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.license):
            request.resource.license.extend(Primitive.to_proto(self.license))
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.replica_zones):
            request.resource.replica_zones.extend(
                Primitive.to_proto(self.replica_zones)
            )
        if Primitive.to_proto(self.resource_policy):
            request.resource.resource_policy.extend(
                Primitive.to_proto(self.resource_policy)
            )
        if Primitive.to_proto(self.size_gb):
            request.resource.size_gb = Primitive.to_proto(self.size_gb)

        if Primitive.to_proto(self.source_image):
            request.resource.source_image = Primitive.to_proto(self.source_image)

        if DiskEncryptionKey.to_proto(self.source_image_encryption_key):
            request.resource.source_image_encryption_key.CopyFrom(
                DiskEncryptionKey.to_proto(self.source_image_encryption_key)
            )
        else:
            request.resource.ClearField("source_image_encryption_key")
        if Primitive.to_proto(self.source_snapshot):
            request.resource.source_snapshot = Primitive.to_proto(self.source_snapshot)

        if DiskEncryptionKey.to_proto(self.source_snapshot_encryption_key):
            request.resource.source_snapshot_encryption_key.CopyFrom(
                DiskEncryptionKey.to_proto(self.source_snapshot_encryption_key)
            )
        else:
            request.resource.ClearField("source_snapshot_encryption_key")
        if Primitive.to_proto(self.type):
            request.resource.type = Primitive.to_proto(self.type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.id):
            request.resource.id = Primitive.to_proto(self.id)

        if Primitive.to_proto(self.options):
            request.resource.options = Primitive.to_proto(self.options)

        if Primitive.to_proto(self.licenses):
            request.resource.licenses.extend(Primitive.to_proto(self.licenses))
        if DiskGuestOSFeaturesArray.to_proto(self.guest_os_features):
            request.resource.guest_os_features.extend(
                DiskGuestOSFeaturesArray.to_proto(self.guest_os_features)
            )
        if int64Array.to_proto(self.license_codes):
            request.resource.license_codes.extend(
                int64Array.to_proto(self.license_codes)
            )
        if Primitive.to_proto(self.physical_block_size_bytes):
            request.resource.physical_block_size_bytes = Primitive.to_proto(
                self.physical_block_size_bytes
            )

        if Primitive.to_proto(self.resource_policies):
            request.resource.resource_policies.extend(
                Primitive.to_proto(self.resource_policies)
            )
        if Primitive.to_proto(self.source_disk):
            request.resource.source_disk = Primitive.to_proto(self.source_disk)

        if Primitive.to_proto(self.source_disk_id):
            request.resource.source_disk_id = Primitive.to_proto(self.source_disk_id)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaDisk(request)
        self.self_link = Primitive.from_proto(response.self_link)
        self.description = Primitive.from_proto(response.description)
        self.disk_encryption_key = DiskEncryptionKey.from_proto(
            response.disk_encryption_key
        )
        self.guest_os_feature = DiskGuestOSFeatureArray.from_proto(
            response.guest_os_feature
        )
        self.labels = Primitive.from_proto(response.labels)
        self.label_fingerprint = Primitive.from_proto(response.label_fingerprint)
        self.license = Primitive.from_proto(response.license)
        self.name = Primitive.from_proto(response.name)
        self.region = Primitive.from_proto(response.region)
        self.replica_zones = Primitive.from_proto(response.replica_zones)
        self.resource_policy = Primitive.from_proto(response.resource_policy)
        self.size_gb = Primitive.from_proto(response.size_gb)
        self.source_image = Primitive.from_proto(response.source_image)
        self.source_image_encryption_key = DiskEncryptionKey.from_proto(
            response.source_image_encryption_key
        )
        self.source_image_id = Primitive.from_proto(response.source_image_id)
        self.source_snapshot = Primitive.from_proto(response.source_snapshot)
        self.source_snapshot_encryption_key = DiskEncryptionKey.from_proto(
            response.source_snapshot_encryption_key
        )
        self.source_snapshot_id = Primitive.from_proto(response.source_snapshot_id)
        self.type = Primitive.from_proto(response.type)
        self.zone = Primitive.from_proto(response.zone)
        self.project = Primitive.from_proto(response.project)
        self.id = Primitive.from_proto(response.id)
        self.status = DiskStatusEnum.from_proto(response.status)
        self.options = Primitive.from_proto(response.options)
        self.licenses = Primitive.from_proto(response.licenses)
        self.guest_os_features = DiskGuestOSFeaturesArray.from_proto(
            response.guest_os_features
        )
        self.last_attach_timestamp = Primitive.from_proto(
            response.last_attach_timestamp
        )
        self.last_detach_timestamp = Primitive.from_proto(
            response.last_detach_timestamp
        )
        self.users = Primitive.from_proto(response.users)
        self.license_codes = int64Array.from_proto(response.license_codes)
        self.physical_block_size_bytes = Primitive.from_proto(
            response.physical_block_size_bytes
        )
        self.resource_policies = Primitive.from_proto(response.resource_policies)
        self.source_disk = Primitive.from_proto(response.source_disk)
        self.source_disk_id = Primitive.from_proto(response.source_disk_id)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = disk_pb2_grpc.ComputeBetaDiskServiceStub(channel.Channel())
        request = disk_pb2.DeleteComputeBetaDiskRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if DiskEncryptionKey.to_proto(self.disk_encryption_key):
            request.resource.disk_encryption_key.CopyFrom(
                DiskEncryptionKey.to_proto(self.disk_encryption_key)
            )
        else:
            request.resource.ClearField("disk_encryption_key")
        if DiskGuestOSFeatureArray.to_proto(self.guest_os_feature):
            request.resource.guest_os_feature.extend(
                DiskGuestOSFeatureArray.to_proto(self.guest_os_feature)
            )
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.license):
            request.resource.license.extend(Primitive.to_proto(self.license))
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.replica_zones):
            request.resource.replica_zones.extend(
                Primitive.to_proto(self.replica_zones)
            )
        if Primitive.to_proto(self.resource_policy):
            request.resource.resource_policy.extend(
                Primitive.to_proto(self.resource_policy)
            )
        if Primitive.to_proto(self.size_gb):
            request.resource.size_gb = Primitive.to_proto(self.size_gb)

        if Primitive.to_proto(self.source_image):
            request.resource.source_image = Primitive.to_proto(self.source_image)

        if DiskEncryptionKey.to_proto(self.source_image_encryption_key):
            request.resource.source_image_encryption_key.CopyFrom(
                DiskEncryptionKey.to_proto(self.source_image_encryption_key)
            )
        else:
            request.resource.ClearField("source_image_encryption_key")
        if Primitive.to_proto(self.source_snapshot):
            request.resource.source_snapshot = Primitive.to_proto(self.source_snapshot)

        if DiskEncryptionKey.to_proto(self.source_snapshot_encryption_key):
            request.resource.source_snapshot_encryption_key.CopyFrom(
                DiskEncryptionKey.to_proto(self.source_snapshot_encryption_key)
            )
        else:
            request.resource.ClearField("source_snapshot_encryption_key")
        if Primitive.to_proto(self.type):
            request.resource.type = Primitive.to_proto(self.type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.id):
            request.resource.id = Primitive.to_proto(self.id)

        if Primitive.to_proto(self.options):
            request.resource.options = Primitive.to_proto(self.options)

        if Primitive.to_proto(self.licenses):
            request.resource.licenses.extend(Primitive.to_proto(self.licenses))
        if DiskGuestOSFeaturesArray.to_proto(self.guest_os_features):
            request.resource.guest_os_features.extend(
                DiskGuestOSFeaturesArray.to_proto(self.guest_os_features)
            )
        if int64Array.to_proto(self.license_codes):
            request.resource.license_codes.extend(
                int64Array.to_proto(self.license_codes)
            )
        if Primitive.to_proto(self.physical_block_size_bytes):
            request.resource.physical_block_size_bytes = Primitive.to_proto(
                self.physical_block_size_bytes
            )

        if Primitive.to_proto(self.resource_policies):
            request.resource.resource_policies.extend(
                Primitive.to_proto(self.resource_policies)
            )
        if Primitive.to_proto(self.source_disk):
            request.resource.source_disk = Primitive.to_proto(self.source_disk)

        if Primitive.to_proto(self.source_disk_id):
            request.resource.source_disk_id = Primitive.to_proto(self.source_disk_id)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteComputeBetaDisk(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = disk_pb2_grpc.ComputeBetaDiskServiceStub(channel.Channel())
        request = disk_pb2.ListComputeBetaDiskRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeBetaDisk(request).items

    def to_proto(self):
        resource = disk_pb2.ComputeBetaDisk()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if DiskEncryptionKey.to_proto(self.disk_encryption_key):
            resource.disk_encryption_key.CopyFrom(
                DiskEncryptionKey.to_proto(self.disk_encryption_key)
            )
        else:
            resource.ClearField("disk_encryption_key")
        if DiskGuestOSFeatureArray.to_proto(self.guest_os_feature):
            resource.guest_os_feature.extend(
                DiskGuestOSFeatureArray.to_proto(self.guest_os_feature)
            )
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.license):
            resource.license.extend(Primitive.to_proto(self.license))
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.replica_zones):
            resource.replica_zones.extend(Primitive.to_proto(self.replica_zones))
        if Primitive.to_proto(self.resource_policy):
            resource.resource_policy.extend(Primitive.to_proto(self.resource_policy))
        if Primitive.to_proto(self.size_gb):
            resource.size_gb = Primitive.to_proto(self.size_gb)
        if Primitive.to_proto(self.source_image):
            resource.source_image = Primitive.to_proto(self.source_image)
        if DiskEncryptionKey.to_proto(self.source_image_encryption_key):
            resource.source_image_encryption_key.CopyFrom(
                DiskEncryptionKey.to_proto(self.source_image_encryption_key)
            )
        else:
            resource.ClearField("source_image_encryption_key")
        if Primitive.to_proto(self.source_snapshot):
            resource.source_snapshot = Primitive.to_proto(self.source_snapshot)
        if DiskEncryptionKey.to_proto(self.source_snapshot_encryption_key):
            resource.source_snapshot_encryption_key.CopyFrom(
                DiskEncryptionKey.to_proto(self.source_snapshot_encryption_key)
            )
        else:
            resource.ClearField("source_snapshot_encryption_key")
        if Primitive.to_proto(self.type):
            resource.type = Primitive.to_proto(self.type)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.id):
            resource.id = Primitive.to_proto(self.id)
        if Primitive.to_proto(self.options):
            resource.options = Primitive.to_proto(self.options)
        if Primitive.to_proto(self.licenses):
            resource.licenses.extend(Primitive.to_proto(self.licenses))
        if DiskGuestOSFeaturesArray.to_proto(self.guest_os_features):
            resource.guest_os_features.extend(
                DiskGuestOSFeaturesArray.to_proto(self.guest_os_features)
            )
        if int64Array.to_proto(self.license_codes):
            resource.license_codes.extend(int64Array.to_proto(self.license_codes))
        if Primitive.to_proto(self.physical_block_size_bytes):
            resource.physical_block_size_bytes = Primitive.to_proto(
                self.physical_block_size_bytes
            )
        if Primitive.to_proto(self.resource_policies):
            resource.resource_policies.extend(
                Primitive.to_proto(self.resource_policies)
            )
        if Primitive.to_proto(self.source_disk):
            resource.source_disk = Primitive.to_proto(self.source_disk)
        if Primitive.to_proto(self.source_disk_id):
            resource.source_disk_id = Primitive.to_proto(self.source_disk_id)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class DiskGuestOSFeature(object):
    def __init__(self, type: str = None, type_alt: list = None):
        self.type = type
        self.type_alt = type_alt

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = disk_pb2.ComputeBetaDiskGuestOSFeature()
        if DiskGuestOSFeatureTypeEnum.to_proto(resource.type):
            res.type = DiskGuestOSFeatureTypeEnum.to_proto(resource.type)
        if DiskGuestOSFeatureTypeAltEnumArray.to_proto(resource.type_alt):
            res.type_alt.extend(
                DiskGuestOSFeatureTypeAltEnumArray.to_proto(resource.type_alt)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DiskGuestOSFeature(
            type=DiskGuestOSFeatureTypeEnum.from_proto(resource.type),
            type_alt=DiskGuestOSFeatureTypeAltEnumArray.from_proto(resource.type_alt),
        )


class DiskGuestOSFeatureArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DiskGuestOSFeature.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DiskGuestOSFeature.from_proto(i) for i in resources]


class DiskEncryptionKey(object):
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

        res = disk_pb2.ComputeBetaDiskEncryptionKey()
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

        return DiskEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
            sha256=Primitive.from_proto(resource.sha256),
            kms_key_service_account=Primitive.from_proto(
                resource.kms_key_service_account
            ),
        )


class DiskEncryptionKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DiskEncryptionKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DiskEncryptionKey.from_proto(i) for i in resources]


class DiskGuestOSFeatures(object):
    def __init__(self, type: str = None, type_alts: list = None):
        self.type = type
        self.type_alts = type_alts

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = disk_pb2.ComputeBetaDiskGuestOSFeatures()
        if DiskGuestOSFeaturesTypeEnum.to_proto(resource.type):
            res.type = DiskGuestOSFeaturesTypeEnum.to_proto(resource.type)
        if DiskGuestOSFeaturesTypeAltsEnumArray.to_proto(resource.type_alts):
            res.type_alts.extend(
                DiskGuestOSFeaturesTypeAltsEnumArray.to_proto(resource.type_alts)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DiskGuestOSFeatures(
            type=DiskGuestOSFeaturesTypeEnum.from_proto(resource.type),
            type_alts=DiskGuestOSFeaturesTypeAltsEnumArray.from_proto(
                resource.type_alts
            ),
        )


class DiskGuestOSFeaturesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DiskGuestOSFeatures.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DiskGuestOSFeatures.from_proto(i) for i in resources]


class DiskGuestOSFeatureTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return disk_pb2.ComputeBetaDiskGuestOSFeatureTypeEnum.Value(
            "ComputeBetaDiskGuestOSFeatureTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return disk_pb2.ComputeBetaDiskGuestOSFeatureTypeEnum.Name(resource)[
            len("ComputeBetaDiskGuestOSFeatureTypeEnum") :
        ]


class DiskGuestOSFeatureTypeAltEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return disk_pb2.ComputeBetaDiskGuestOSFeatureTypeAltEnum.Value(
            "ComputeBetaDiskGuestOSFeatureTypeAltEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return disk_pb2.ComputeBetaDiskGuestOSFeatureTypeAltEnum.Name(resource)[
            len("ComputeBetaDiskGuestOSFeatureTypeAltEnum") :
        ]


class DiskStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return disk_pb2.ComputeBetaDiskStatusEnum.Value(
            "ComputeBetaDiskStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return disk_pb2.ComputeBetaDiskStatusEnum.Name(resource)[
            len("ComputeBetaDiskStatusEnum") :
        ]


class DiskGuestOSFeaturesTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return disk_pb2.ComputeBetaDiskGuestOSFeaturesTypeEnum.Value(
            "ComputeBetaDiskGuestOSFeaturesTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return disk_pb2.ComputeBetaDiskGuestOSFeaturesTypeEnum.Name(resource)[
            len("ComputeBetaDiskGuestOSFeaturesTypeEnum") :
        ]


class DiskGuestOSFeaturesTypeAltsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return disk_pb2.ComputeBetaDiskGuestOSFeaturesTypeAltsEnum.Value(
            "ComputeBetaDiskGuestOSFeaturesTypeAltsEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return disk_pb2.ComputeBetaDiskGuestOSFeaturesTypeAltsEnum.Name(resource)[
            len("ComputeBetaDiskGuestOSFeaturesTypeAltsEnum") :
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
