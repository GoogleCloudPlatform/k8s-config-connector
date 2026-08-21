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
from google3.cloud.graphite.mmv2.services.google.compute import instance_template_pb2
from google3.cloud.graphite.mmv2.services.google.compute import (
    instance_template_pb2_grpc,
)

from typing import List


class InstanceTemplate(object):
    def __init__(
        self,
        creation_timestamp: str = None,
        description: str = None,
        id: int = None,
        self_link: str = None,
        name: str = None,
        properties: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.description = description
        self.self_link = self_link
        self.name = name
        self.properties = properties
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_template_pb2_grpc.ComputeBetaInstanceTemplateServiceStub(
            channel.Channel()
        )
        request = instance_template_pb2.ApplyComputeBetaInstanceTemplateRequest()
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.self_link):
            request.resource.self_link = Primitive.to_proto(self.self_link)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if InstanceTemplateProperties.to_proto(self.properties):
            request.resource.properties.CopyFrom(
                InstanceTemplateProperties.to_proto(self.properties)
            )
        else:
            request.resource.ClearField("properties")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaInstanceTemplate(request)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.description = Primitive.from_proto(response.description)
        self.id = Primitive.from_proto(response.id)
        self.self_link = Primitive.from_proto(response.self_link)
        self.name = Primitive.from_proto(response.name)
        self.properties = InstanceTemplateProperties.from_proto(response.properties)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = instance_template_pb2_grpc.ComputeBetaInstanceTemplateServiceStub(
            channel.Channel()
        )
        request = instance_template_pb2.DeleteComputeBetaInstanceTemplateRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.self_link):
            request.resource.self_link = Primitive.to_proto(self.self_link)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if InstanceTemplateProperties.to_proto(self.properties):
            request.resource.properties.CopyFrom(
                InstanceTemplateProperties.to_proto(self.properties)
            )
        else:
            request.resource.ClearField("properties")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeBetaInstanceTemplate(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = instance_template_pb2_grpc.ComputeBetaInstanceTemplateServiceStub(
            channel.Channel()
        )
        request = instance_template_pb2.ListComputeBetaInstanceTemplateRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeBetaInstanceTemplate(request).items

    def to_proto(self):
        resource = instance_template_pb2.ComputeBetaInstanceTemplate()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.self_link):
            resource.self_link = Primitive.to_proto(self.self_link)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if InstanceTemplateProperties.to_proto(self.properties):
            resource.properties.CopyFrom(
                InstanceTemplateProperties.to_proto(self.properties)
            )
        else:
            resource.ClearField("properties")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class InstanceTemplateProperties(object):
    def __init__(
        self,
        can_ip_forward: bool = None,
        description: str = None,
        disks: list = None,
        labels: dict = None,
        machine_type: str = None,
        min_cpu_platform: str = None,
        metadata: dict = None,
        reservation_affinity: dict = None,
        guest_accelerators: list = None,
        network_interfaces: list = None,
        shielded_instance_config: dict = None,
        scheduling: dict = None,
        service_accounts: list = None,
        tags: list = None,
    ):
        self.can_ip_forward = can_ip_forward
        self.description = description
        self.disks = disks
        self.labels = labels
        self.machine_type = machine_type
        self.min_cpu_platform = min_cpu_platform
        self.metadata = metadata
        self.reservation_affinity = reservation_affinity
        self.guest_accelerators = guest_accelerators
        self.network_interfaces = network_interfaces
        self.shielded_instance_config = shielded_instance_config
        self.scheduling = scheduling
        self.service_accounts = service_accounts
        self.tags = tags

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_template_pb2.ComputeBetaInstanceTemplateProperties()
        if Primitive.to_proto(resource.can_ip_forward):
            res.can_ip_forward = Primitive.to_proto(resource.can_ip_forward)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if InstanceTemplatePropertiesDisksArray.to_proto(resource.disks):
            res.disks.extend(
                InstanceTemplatePropertiesDisksArray.to_proto(resource.disks)
            )
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        if Primitive.to_proto(resource.metadata):
            res.metadata = Primitive.to_proto(resource.metadata)
        if InstanceTemplatePropertiesReservationAffinity.to_proto(
            resource.reservation_affinity
        ):
            res.reservation_affinity.CopyFrom(
                InstanceTemplatePropertiesReservationAffinity.to_proto(
                    resource.reservation_affinity
                )
            )
        else:
            res.ClearField("reservation_affinity")
        if InstanceTemplatePropertiesGuestAcceleratorsArray.to_proto(
            resource.guest_accelerators
        ):
            res.guest_accelerators.extend(
                InstanceTemplatePropertiesGuestAcceleratorsArray.to_proto(
                    resource.guest_accelerators
                )
            )
        if InstanceTemplatePropertiesNetworkInterfacesArray.to_proto(
            resource.network_interfaces
        ):
            res.network_interfaces.extend(
                InstanceTemplatePropertiesNetworkInterfacesArray.to_proto(
                    resource.network_interfaces
                )
            )
        if InstanceTemplatePropertiesShieldedInstanceConfig.to_proto(
            resource.shielded_instance_config
        ):
            res.shielded_instance_config.CopyFrom(
                InstanceTemplatePropertiesShieldedInstanceConfig.to_proto(
                    resource.shielded_instance_config
                )
            )
        else:
            res.ClearField("shielded_instance_config")
        if InstanceTemplatePropertiesScheduling.to_proto(resource.scheduling):
            res.scheduling.CopyFrom(
                InstanceTemplatePropertiesScheduling.to_proto(resource.scheduling)
            )
        else:
            res.ClearField("scheduling")
        if InstanceTemplatePropertiesServiceAccountsArray.to_proto(
            resource.service_accounts
        ):
            res.service_accounts.extend(
                InstanceTemplatePropertiesServiceAccountsArray.to_proto(
                    resource.service_accounts
                )
            )
        if Primitive.to_proto(resource.tags):
            res.tags.extend(Primitive.to_proto(resource.tags))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplateProperties(
            can_ip_forward=Primitive.from_proto(resource.can_ip_forward),
            description=Primitive.from_proto(resource.description),
            disks=InstanceTemplatePropertiesDisksArray.from_proto(resource.disks),
            labels=Primitive.from_proto(resource.labels),
            machine_type=Primitive.from_proto(resource.machine_type),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
            metadata=Primitive.from_proto(resource.metadata),
            reservation_affinity=InstanceTemplatePropertiesReservationAffinity.from_proto(
                resource.reservation_affinity
            ),
            guest_accelerators=InstanceTemplatePropertiesGuestAcceleratorsArray.from_proto(
                resource.guest_accelerators
            ),
            network_interfaces=InstanceTemplatePropertiesNetworkInterfacesArray.from_proto(
                resource.network_interfaces
            ),
            shielded_instance_config=InstanceTemplatePropertiesShieldedInstanceConfig.from_proto(
                resource.shielded_instance_config
            ),
            scheduling=InstanceTemplatePropertiesScheduling.from_proto(
                resource.scheduling
            ),
            service_accounts=InstanceTemplatePropertiesServiceAccountsArray.from_proto(
                resource.service_accounts
            ),
            tags=Primitive.from_proto(resource.tags),
        )


class InstanceTemplatePropertiesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceTemplateProperties.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceTemplateProperties.from_proto(i) for i in resources]


class InstanceTemplatePropertiesDisks(object):
    def __init__(
        self,
        auto_delete: bool = None,
        boot: bool = None,
        device_name: str = None,
        disk_encryption_key: dict = None,
        index: int = None,
        initialize_params: dict = None,
        guest_os_features: list = None,
        interface: str = None,
        mode: str = None,
        source: str = None,
        type: str = None,
    ):
        self.auto_delete = auto_delete
        self.boot = boot
        self.device_name = device_name
        self.disk_encryption_key = disk_encryption_key
        self.index = index
        self.initialize_params = initialize_params
        self.guest_os_features = guest_os_features
        self.interface = interface
        self.mode = mode
        self.source = source
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisks()
        if Primitive.to_proto(resource.auto_delete):
            res.auto_delete = Primitive.to_proto(resource.auto_delete)
        if Primitive.to_proto(resource.boot):
            res.boot = Primitive.to_proto(resource.boot)
        if Primitive.to_proto(resource.device_name):
            res.device_name = Primitive.to_proto(resource.device_name)
        if InstanceTemplatePropertiesDisksDiskEncryptionKey.to_proto(
            resource.disk_encryption_key
        ):
            res.disk_encryption_key.CopyFrom(
                InstanceTemplatePropertiesDisksDiskEncryptionKey.to_proto(
                    resource.disk_encryption_key
                )
            )
        else:
            res.ClearField("disk_encryption_key")
        if Primitive.to_proto(resource.index):
            res.index = Primitive.to_proto(resource.index)
        if InstanceTemplatePropertiesDisksInitializeParams.to_proto(
            resource.initialize_params
        ):
            res.initialize_params.CopyFrom(
                InstanceTemplatePropertiesDisksInitializeParams.to_proto(
                    resource.initialize_params
                )
            )
        else:
            res.ClearField("initialize_params")
        if InstanceTemplatePropertiesDisksGuestOSFeaturesArray.to_proto(
            resource.guest_os_features
        ):
            res.guest_os_features.extend(
                InstanceTemplatePropertiesDisksGuestOSFeaturesArray.to_proto(
                    resource.guest_os_features
                )
            )
        if InstanceTemplatePropertiesDisksInterfaceEnum.to_proto(resource.interface):
            res.interface = InstanceTemplatePropertiesDisksInterfaceEnum.to_proto(
                resource.interface
            )
        if InstanceTemplatePropertiesDisksModeEnum.to_proto(resource.mode):
            res.mode = InstanceTemplatePropertiesDisksModeEnum.to_proto(resource.mode)
        if Primitive.to_proto(resource.source):
            res.source = Primitive.to_proto(resource.source)
        if InstanceTemplatePropertiesDisksTypeEnum.to_proto(resource.type):
            res.type = InstanceTemplatePropertiesDisksTypeEnum.to_proto(resource.type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesDisks(
            auto_delete=Primitive.from_proto(resource.auto_delete),
            boot=Primitive.from_proto(resource.boot),
            device_name=Primitive.from_proto(resource.device_name),
            disk_encryption_key=InstanceTemplatePropertiesDisksDiskEncryptionKey.from_proto(
                resource.disk_encryption_key
            ),
            index=Primitive.from_proto(resource.index),
            initialize_params=InstanceTemplatePropertiesDisksInitializeParams.from_proto(
                resource.initialize_params
            ),
            guest_os_features=InstanceTemplatePropertiesDisksGuestOSFeaturesArray.from_proto(
                resource.guest_os_features
            ),
            interface=InstanceTemplatePropertiesDisksInterfaceEnum.from_proto(
                resource.interface
            ),
            mode=InstanceTemplatePropertiesDisksModeEnum.from_proto(resource.mode),
            source=Primitive.from_proto(resource.source),
            type=InstanceTemplatePropertiesDisksTypeEnum.from_proto(resource.type),
        )


class InstanceTemplatePropertiesDisksArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceTemplatePropertiesDisks.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceTemplatePropertiesDisks.from_proto(i) for i in resources]


class InstanceTemplatePropertiesDisksDiskEncryptionKey(object):
    def __init__(
        self, raw_key: str = None, rsa_encrypted_key: str = None, sha256: str = None
    ):
        self.raw_key = raw_key
        self.rsa_encrypted_key = rsa_encrypted_key
        self.sha256 = sha256

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisksDiskEncryptionKey()
        )
        if Primitive.to_proto(resource.raw_key):
            res.raw_key = Primitive.to_proto(resource.raw_key)
        if Primitive.to_proto(resource.rsa_encrypted_key):
            res.rsa_encrypted_key = Primitive.to_proto(resource.rsa_encrypted_key)
        if Primitive.to_proto(resource.sha256):
            res.sha256 = Primitive.to_proto(resource.sha256)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesDisksDiskEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
            rsa_encrypted_key=Primitive.from_proto(resource.rsa_encrypted_key),
            sha256=Primitive.from_proto(resource.sha256),
        )


class InstanceTemplatePropertiesDisksDiskEncryptionKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesDisksDiskEncryptionKey.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesDisksDiskEncryptionKey.from_proto(i)
            for i in resources
        ]


class InstanceTemplatePropertiesDisksInitializeParams(object):
    def __init__(
        self,
        disk_name: str = None,
        disk_size_gb: int = None,
        disk_type: str = None,
        source_image: str = None,
        labels: dict = None,
        source_snapshot: str = None,
        source_snapshot_encryption_key: dict = None,
        description: str = None,
        resource_policies: list = None,
        on_update_action: str = None,
        source_image_encryption_key: dict = None,
    ):
        self.disk_name = disk_name
        self.disk_size_gb = disk_size_gb
        self.disk_type = disk_type
        self.source_image = source_image
        self.labels = labels
        self.source_snapshot = source_snapshot
        self.source_snapshot_encryption_key = source_snapshot_encryption_key
        self.description = description
        self.resource_policies = resource_policies
        self.on_update_action = on_update_action
        self.source_image_encryption_key = source_image_encryption_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisksInitializeParams()
        )
        if Primitive.to_proto(resource.disk_name):
            res.disk_name = Primitive.to_proto(resource.disk_name)
        if Primitive.to_proto(resource.disk_size_gb):
            res.disk_size_gb = Primitive.to_proto(resource.disk_size_gb)
        if Primitive.to_proto(resource.disk_type):
            res.disk_type = Primitive.to_proto(resource.disk_type)
        if Primitive.to_proto(resource.source_image):
            res.source_image = Primitive.to_proto(resource.source_image)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if Primitive.to_proto(resource.source_snapshot):
            res.source_snapshot = Primitive.to_proto(resource.source_snapshot)
        if InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey.to_proto(
            resource.source_snapshot_encryption_key
        ):
            res.source_snapshot_encryption_key.CopyFrom(
                InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey.to_proto(
                    resource.source_snapshot_encryption_key
                )
            )
        else:
            res.ClearField("source_snapshot_encryption_key")
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.resource_policies):
            res.resource_policies.extend(Primitive.to_proto(resource.resource_policies))
        if Primitive.to_proto(resource.on_update_action):
            res.on_update_action = Primitive.to_proto(resource.on_update_action)
        if InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey.to_proto(
            resource.source_image_encryption_key
        ):
            res.source_image_encryption_key.CopyFrom(
                InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey.to_proto(
                    resource.source_image_encryption_key
                )
            )
        else:
            res.ClearField("source_image_encryption_key")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesDisksInitializeParams(
            disk_name=Primitive.from_proto(resource.disk_name),
            disk_size_gb=Primitive.from_proto(resource.disk_size_gb),
            disk_type=Primitive.from_proto(resource.disk_type),
            source_image=Primitive.from_proto(resource.source_image),
            labels=Primitive.from_proto(resource.labels),
            source_snapshot=Primitive.from_proto(resource.source_snapshot),
            source_snapshot_encryption_key=InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey.from_proto(
                resource.source_snapshot_encryption_key
            ),
            description=Primitive.from_proto(resource.description),
            resource_policies=Primitive.from_proto(resource.resource_policies),
            on_update_action=Primitive.from_proto(resource.on_update_action),
            source_image_encryption_key=InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey.from_proto(
                resource.source_image_encryption_key
            ),
        )


class InstanceTemplatePropertiesDisksInitializeParamsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesDisksInitializeParams.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesDisksInitializeParams.from_proto(i)
            for i in resources
        ]


class InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(
    object
):
    def __init__(
        self, raw_key: str = None, sha256: str = None, kms_key_name: str = None
    ):
        self.raw_key = raw_key
        self.sha256 = sha256
        self.kms_key_name = kms_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey()
        )
        if Primitive.to_proto(resource.raw_key):
            res.raw_key = Primitive.to_proto(resource.raw_key)
        if Primitive.to_proto(resource.sha256):
            res.sha256 = Primitive.to_proto(resource.sha256)
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
            sha256=Primitive.from_proto(resource.sha256),
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
        )


class InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey.from_proto(
                i
            )
            for i in resources
        ]


class InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(object):
    def __init__(
        self, raw_key: str = None, sha256: str = None, kms_key_name: str = None
    ):
        self.raw_key = raw_key
        self.sha256 = sha256
        self.kms_key_name = kms_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey()
        )
        if Primitive.to_proto(resource.raw_key):
            res.raw_key = Primitive.to_proto(resource.raw_key)
        if Primitive.to_proto(resource.sha256):
            res.sha256 = Primitive.to_proto(resource.sha256)
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
            sha256=Primitive.from_proto(resource.sha256),
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
        )


class InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey.from_proto(
                i
            )
            for i in resources
        ]


class InstanceTemplatePropertiesDisksGuestOSFeatures(object):
    def __init__(self, type: str = None):
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisksGuestOSFeatures()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesDisksGuestOSFeatures(
            type=Primitive.from_proto(resource.type),
        )


class InstanceTemplatePropertiesDisksGuestOSFeaturesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesDisksGuestOSFeatures.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesDisksGuestOSFeatures.from_proto(i)
            for i in resources
        ]


class InstanceTemplatePropertiesReservationAffinity(object):
    def __init__(self, key: str = None, value: list = None):
        self.key = key
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesReservationAffinity()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.value):
            res.value.extend(Primitive.to_proto(resource.value))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesReservationAffinity(
            key=Primitive.from_proto(resource.key),
            value=Primitive.from_proto(resource.value),
        )


class InstanceTemplatePropertiesReservationAffinityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesReservationAffinity.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesReservationAffinity.from_proto(i)
            for i in resources
        ]


class InstanceTemplatePropertiesGuestAccelerators(object):
    def __init__(self, accelerator_count: int = None, accelerator_type: str = None):
        self.accelerator_count = accelerator_count
        self.accelerator_type = accelerator_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesGuestAccelerators()
        )
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesGuestAccelerators(
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
        )


class InstanceTemplatePropertiesGuestAcceleratorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesGuestAccelerators.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesGuestAccelerators.from_proto(i) for i in resources
        ]


class InstanceTemplatePropertiesNetworkInterfaces(object):
    def __init__(
        self,
        access_configs: list = None,
        alias_ip_ranges: list = None,
        name: str = None,
        network: str = None,
        network_ip: str = None,
        subnetwork: str = None,
    ):
        self.access_configs = access_configs
        self.alias_ip_ranges = alias_ip_ranges
        self.name = name
        self.network = network
        self.network_ip = network_ip
        self.subnetwork = subnetwork

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesNetworkInterfaces()
        )
        if InstanceTemplatePropertiesNetworkInterfacesAccessConfigsArray.to_proto(
            resource.access_configs
        ):
            res.access_configs.extend(
                InstanceTemplatePropertiesNetworkInterfacesAccessConfigsArray.to_proto(
                    resource.access_configs
                )
            )
        if InstanceTemplatePropertiesNetworkInterfacesAliasIPRangesArray.to_proto(
            resource.alias_ip_ranges
        ):
            res.alias_ip_ranges.extend(
                InstanceTemplatePropertiesNetworkInterfacesAliasIPRangesArray.to_proto(
                    resource.alias_ip_ranges
                )
            )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        if Primitive.to_proto(resource.network_ip):
            res.network_ip = Primitive.to_proto(resource.network_ip)
        if Primitive.to_proto(resource.subnetwork):
            res.subnetwork = Primitive.to_proto(resource.subnetwork)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesNetworkInterfaces(
            access_configs=InstanceTemplatePropertiesNetworkInterfacesAccessConfigsArray.from_proto(
                resource.access_configs
            ),
            alias_ip_ranges=InstanceTemplatePropertiesNetworkInterfacesAliasIPRangesArray.from_proto(
                resource.alias_ip_ranges
            ),
            name=Primitive.from_proto(resource.name),
            network=Primitive.from_proto(resource.network),
            network_ip=Primitive.from_proto(resource.network_ip),
            subnetwork=Primitive.from_proto(resource.subnetwork),
        )


class InstanceTemplatePropertiesNetworkInterfacesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesNetworkInterfaces.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesNetworkInterfaces.from_proto(i) for i in resources
        ]


class InstanceTemplatePropertiesNetworkInterfacesAccessConfigs(object):
    def __init__(
        self,
        name: str = None,
        nat_ip: str = None,
        type: str = None,
        set_public_ptr: bool = None,
        public_ptr_domain_name: str = None,
        network_tier: str = None,
    ):
        self.name = name
        self.nat_ip = nat_ip
        self.type = type
        self.set_public_ptr = set_public_ptr
        self.public_ptr_domain_name = public_ptr_domain_name
        self.network_tier = network_tier

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigs()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.nat_ip):
            res.nat_ip = Primitive.to_proto(resource.nat_ip)
        if InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum.to_proto(
            resource.type
        ):
            res.type = InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum.to_proto(
                resource.type
            )
        if Primitive.to_proto(resource.set_public_ptr):
            res.set_public_ptr = Primitive.to_proto(resource.set_public_ptr)
        if Primitive.to_proto(resource.public_ptr_domain_name):
            res.public_ptr_domain_name = Primitive.to_proto(
                resource.public_ptr_domain_name
            )
        if InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum.to_proto(
            resource.network_tier
        ):
            res.network_tier = InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum.to_proto(
                resource.network_tier
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesNetworkInterfacesAccessConfigs(
            name=Primitive.from_proto(resource.name),
            nat_ip=Primitive.from_proto(resource.nat_ip),
            type=InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum.from_proto(
                resource.type
            ),
            set_public_ptr=Primitive.from_proto(resource.set_public_ptr),
            public_ptr_domain_name=Primitive.from_proto(
                resource.public_ptr_domain_name
            ),
            network_tier=InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum.from_proto(
                resource.network_tier
            ),
        )


class InstanceTemplatePropertiesNetworkInterfacesAccessConfigsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesNetworkInterfacesAccessConfigs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesNetworkInterfacesAccessConfigs.from_proto(i)
            for i in resources
        ]


class InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(object):
    def __init__(self, ip_cidr_range: str = None, subnetwork_range_name: str = None):
        self.ip_cidr_range = ip_cidr_range
        self.subnetwork_range_name = subnetwork_range_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges()
        )
        if Primitive.to_proto(resource.ip_cidr_range):
            res.ip_cidr_range = Primitive.to_proto(resource.ip_cidr_range)
        if Primitive.to_proto(resource.subnetwork_range_name):
            res.subnetwork_range_name = Primitive.to_proto(
                resource.subnetwork_range_name
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(
            ip_cidr_range=Primitive.from_proto(resource.ip_cidr_range),
            subnetwork_range_name=Primitive.from_proto(resource.subnetwork_range_name),
        )


class InstanceTemplatePropertiesNetworkInterfacesAliasIPRangesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges.from_proto(i)
            for i in resources
        ]


class InstanceTemplatePropertiesShieldedInstanceConfig(object):
    def __init__(
        self,
        enable_secure_boot: bool = None,
        enable_vtpm: bool = None,
        enable_integrity_monitoring: bool = None,
    ):
        self.enable_secure_boot = enable_secure_boot
        self.enable_vtpm = enable_vtpm
        self.enable_integrity_monitoring = enable_integrity_monitoring

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesShieldedInstanceConfig()
        )
        if Primitive.to_proto(resource.enable_secure_boot):
            res.enable_secure_boot = Primitive.to_proto(resource.enable_secure_boot)
        if Primitive.to_proto(resource.enable_vtpm):
            res.enable_vtpm = Primitive.to_proto(resource.enable_vtpm)
        if Primitive.to_proto(resource.enable_integrity_monitoring):
            res.enable_integrity_monitoring = Primitive.to_proto(
                resource.enable_integrity_monitoring
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesShieldedInstanceConfig(
            enable_secure_boot=Primitive.from_proto(resource.enable_secure_boot),
            enable_vtpm=Primitive.from_proto(resource.enable_vtpm),
            enable_integrity_monitoring=Primitive.from_proto(
                resource.enable_integrity_monitoring
            ),
        )


class InstanceTemplatePropertiesShieldedInstanceConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesShieldedInstanceConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesShieldedInstanceConfig.from_proto(i)
            for i in resources
        ]


class InstanceTemplatePropertiesScheduling(object):
    def __init__(
        self,
        automatic_restart: bool = None,
        on_host_maintenance: str = None,
        preemptible: bool = None,
        node_affinities: list = None,
    ):
        self.automatic_restart = automatic_restart
        self.on_host_maintenance = on_host_maintenance
        self.preemptible = preemptible
        self.node_affinities = node_affinities

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_template_pb2.ComputeBetaInstanceTemplatePropertiesScheduling()
        if Primitive.to_proto(resource.automatic_restart):
            res.automatic_restart = Primitive.to_proto(resource.automatic_restart)
        if Primitive.to_proto(resource.on_host_maintenance):
            res.on_host_maintenance = Primitive.to_proto(resource.on_host_maintenance)
        if Primitive.to_proto(resource.preemptible):
            res.preemptible = Primitive.to_proto(resource.preemptible)
        if InstanceTemplatePropertiesSchedulingNodeAffinitiesArray.to_proto(
            resource.node_affinities
        ):
            res.node_affinities.extend(
                InstanceTemplatePropertiesSchedulingNodeAffinitiesArray.to_proto(
                    resource.node_affinities
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesScheduling(
            automatic_restart=Primitive.from_proto(resource.automatic_restart),
            on_host_maintenance=Primitive.from_proto(resource.on_host_maintenance),
            preemptible=Primitive.from_proto(resource.preemptible),
            node_affinities=InstanceTemplatePropertiesSchedulingNodeAffinitiesArray.from_proto(
                resource.node_affinities
            ),
        )


class InstanceTemplatePropertiesSchedulingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceTemplatePropertiesScheduling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceTemplatePropertiesScheduling.from_proto(i) for i in resources]


class InstanceTemplatePropertiesSchedulingNodeAffinities(object):
    def __init__(self, key: str = None, operator: str = None, values: list = None):
        self.key = key
        self.operator = operator
        self.values = values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinities()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum.to_proto(
            resource.operator
        ):
            res.operator = InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum.to_proto(
                resource.operator
            )
        if Primitive.to_proto(resource.values):
            res.values.extend(Primitive.to_proto(resource.values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesSchedulingNodeAffinities(
            key=Primitive.from_proto(resource.key),
            operator=InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum.from_proto(
                resource.operator
            ),
            values=Primitive.from_proto(resource.values),
        )


class InstanceTemplatePropertiesSchedulingNodeAffinitiesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesSchedulingNodeAffinities.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesSchedulingNodeAffinities.from_proto(i)
            for i in resources
        ]


class InstanceTemplatePropertiesServiceAccounts(object):
    def __init__(self, email: str = None, scopes: list = None):
        self.email = email
        self.scopes = scopes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_template_pb2.ComputeBetaInstanceTemplatePropertiesServiceAccounts()
        )
        if Primitive.to_proto(resource.email):
            res.email = Primitive.to_proto(resource.email)
        if Primitive.to_proto(resource.scopes):
            res.scopes.extend(Primitive.to_proto(resource.scopes))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceTemplatePropertiesServiceAccounts(
            email=Primitive.from_proto(resource.email),
            scopes=Primitive.from_proto(resource.scopes),
        )


class InstanceTemplatePropertiesServiceAccountsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceTemplatePropertiesServiceAccounts.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceTemplatePropertiesServiceAccounts.from_proto(i) for i in resources
        ]


class InstanceTemplatePropertiesDisksInterfaceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum.Value(
            "ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum.Name(
            resource
        )[
            len("ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum") :
        ]


class InstanceTemplatePropertiesDisksModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisksModeEnum.Value(
            "ComputeBetaInstanceTemplatePropertiesDisksModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisksModeEnum.Name(
            resource
        )[
            len("ComputeBetaInstanceTemplatePropertiesDisksModeEnum") :
        ]


class InstanceTemplatePropertiesDisksTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisksTypeEnum.Value(
            "ComputeBetaInstanceTemplatePropertiesDisksTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesDisksTypeEnum.Name(
            resource
        )[
            len("ComputeBetaInstanceTemplatePropertiesDisksTypeEnum") :
        ]


class InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum.Value(
            "ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum.Name(
            resource
        )[
            len(
                "ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum"
            ) :
        ]


class InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum.Value(
            "ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum.Name(
            resource
        )[
            len(
                "ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum"
            ) :
        ]


class InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum.Value(
            "ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_template_pb2.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum.Name(
            resource
        )[
            len(
                "ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum"
            ) :
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
