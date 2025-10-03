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
from google3.cloud.graphite.mmv2.services.google.compute import instance_pb2
from google3.cloud.graphite.mmv2.services.google.compute import instance_pb2_grpc

from typing import List


class Instance(object):
    def __init__(
        self,
        can_ip_forward: bool = None,
        cpu_platform: str = None,
        creation_timestamp: str = None,
        deletion_protection: bool = None,
        description: str = None,
        disks: list = None,
        guest_accelerators: list = None,
        hostname: str = None,
        id: str = None,
        labels: dict = None,
        metadata: dict = None,
        machine_type: str = None,
        min_cpu_platform: str = None,
        name: str = None,
        network_interfaces: list = None,
        scheduling: dict = None,
        service_accounts: list = None,
        shielded_instance_config: dict = None,
        status: str = None,
        status_message: str = None,
        tags: list = None,
        zone: str = None,
        project: str = None,
        self_link: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.can_ip_forward = can_ip_forward
        self.deletion_protection = deletion_protection
        self.description = description
        self.disks = disks
        self.guest_accelerators = guest_accelerators
        self.hostname = hostname
        self.labels = labels
        self.metadata = metadata
        self.machine_type = machine_type
        self.min_cpu_platform = min_cpu_platform
        self.name = name
        self.network_interfaces = network_interfaces
        self.scheduling = scheduling
        self.service_accounts = service_accounts
        self.shielded_instance_config = shielded_instance_config
        self.status = status
        self.tags = tags
        self.zone = zone
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_pb2_grpc.ComputeInstanceServiceStub(channel.Channel())
        request = instance_pb2.ApplyComputeInstanceRequest()
        if Primitive.to_proto(self.can_ip_forward):
            request.resource.can_ip_forward = Primitive.to_proto(self.can_ip_forward)

        if Primitive.to_proto(self.deletion_protection):
            request.resource.deletion_protection = Primitive.to_proto(
                self.deletion_protection
            )

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if InstanceDisksArray.to_proto(self.disks):
            request.resource.disks.extend(InstanceDisksArray.to_proto(self.disks))
        if InstanceGuestAcceleratorsArray.to_proto(self.guest_accelerators):
            request.resource.guest_accelerators.extend(
                InstanceGuestAcceleratorsArray.to_proto(self.guest_accelerators)
            )
        if Primitive.to_proto(self.hostname):
            request.resource.hostname = Primitive.to_proto(self.hostname)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.metadata):
            request.resource.metadata = Primitive.to_proto(self.metadata)

        if Primitive.to_proto(self.machine_type):
            request.resource.machine_type = Primitive.to_proto(self.machine_type)

        if Primitive.to_proto(self.min_cpu_platform):
            request.resource.min_cpu_platform = Primitive.to_proto(
                self.min_cpu_platform
            )

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if InstanceNetworkInterfacesArray.to_proto(self.network_interfaces):
            request.resource.network_interfaces.extend(
                InstanceNetworkInterfacesArray.to_proto(self.network_interfaces)
            )
        if InstanceScheduling.to_proto(self.scheduling):
            request.resource.scheduling.CopyFrom(
                InstanceScheduling.to_proto(self.scheduling)
            )
        else:
            request.resource.ClearField("scheduling")
        if InstanceServiceAccountsArray.to_proto(self.service_accounts):
            request.resource.service_accounts.extend(
                InstanceServiceAccountsArray.to_proto(self.service_accounts)
            )
        if InstanceShieldedInstanceConfig.to_proto(self.shielded_instance_config):
            request.resource.shielded_instance_config.CopyFrom(
                InstanceShieldedInstanceConfig.to_proto(self.shielded_instance_config)
            )
        else:
            request.resource.ClearField("shielded_instance_config")
        if InstanceStatusEnum.to_proto(self.status):
            request.resource.status = InstanceStatusEnum.to_proto(self.status)

        if Primitive.to_proto(self.tags):
            request.resource.tags.extend(Primitive.to_proto(self.tags))
        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeInstance(request)
        self.can_ip_forward = Primitive.from_proto(response.can_ip_forward)
        self.cpu_platform = Primitive.from_proto(response.cpu_platform)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.deletion_protection = Primitive.from_proto(response.deletion_protection)
        self.description = Primitive.from_proto(response.description)
        self.disks = InstanceDisksArray.from_proto(response.disks)
        self.guest_accelerators = InstanceGuestAcceleratorsArray.from_proto(
            response.guest_accelerators
        )
        self.hostname = Primitive.from_proto(response.hostname)
        self.id = Primitive.from_proto(response.id)
        self.labels = Primitive.from_proto(response.labels)
        self.metadata = Primitive.from_proto(response.metadata)
        self.machine_type = Primitive.from_proto(response.machine_type)
        self.min_cpu_platform = Primitive.from_proto(response.min_cpu_platform)
        self.name = Primitive.from_proto(response.name)
        self.network_interfaces = InstanceNetworkInterfacesArray.from_proto(
            response.network_interfaces
        )
        self.scheduling = InstanceScheduling.from_proto(response.scheduling)
        self.service_accounts = InstanceServiceAccountsArray.from_proto(
            response.service_accounts
        )
        self.shielded_instance_config = InstanceShieldedInstanceConfig.from_proto(
            response.shielded_instance_config
        )
        self.status = InstanceStatusEnum.from_proto(response.status)
        self.status_message = Primitive.from_proto(response.status_message)
        self.tags = Primitive.from_proto(response.tags)
        self.zone = Primitive.from_proto(response.zone)
        self.project = Primitive.from_proto(response.project)
        self.self_link = Primitive.from_proto(response.self_link)

    def delete(self):
        stub = instance_pb2_grpc.ComputeInstanceServiceStub(channel.Channel())
        request = instance_pb2.DeleteComputeInstanceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.can_ip_forward):
            request.resource.can_ip_forward = Primitive.to_proto(self.can_ip_forward)

        if Primitive.to_proto(self.deletion_protection):
            request.resource.deletion_protection = Primitive.to_proto(
                self.deletion_protection
            )

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if InstanceDisksArray.to_proto(self.disks):
            request.resource.disks.extend(InstanceDisksArray.to_proto(self.disks))
        if InstanceGuestAcceleratorsArray.to_proto(self.guest_accelerators):
            request.resource.guest_accelerators.extend(
                InstanceGuestAcceleratorsArray.to_proto(self.guest_accelerators)
            )
        if Primitive.to_proto(self.hostname):
            request.resource.hostname = Primitive.to_proto(self.hostname)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.metadata):
            request.resource.metadata = Primitive.to_proto(self.metadata)

        if Primitive.to_proto(self.machine_type):
            request.resource.machine_type = Primitive.to_proto(self.machine_type)

        if Primitive.to_proto(self.min_cpu_platform):
            request.resource.min_cpu_platform = Primitive.to_proto(
                self.min_cpu_platform
            )

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if InstanceNetworkInterfacesArray.to_proto(self.network_interfaces):
            request.resource.network_interfaces.extend(
                InstanceNetworkInterfacesArray.to_proto(self.network_interfaces)
            )
        if InstanceScheduling.to_proto(self.scheduling):
            request.resource.scheduling.CopyFrom(
                InstanceScheduling.to_proto(self.scheduling)
            )
        else:
            request.resource.ClearField("scheduling")
        if InstanceServiceAccountsArray.to_proto(self.service_accounts):
            request.resource.service_accounts.extend(
                InstanceServiceAccountsArray.to_proto(self.service_accounts)
            )
        if InstanceShieldedInstanceConfig.to_proto(self.shielded_instance_config):
            request.resource.shielded_instance_config.CopyFrom(
                InstanceShieldedInstanceConfig.to_proto(self.shielded_instance_config)
            )
        else:
            request.resource.ClearField("shielded_instance_config")
        if InstanceStatusEnum.to_proto(self.status):
            request.resource.status = InstanceStatusEnum.to_proto(self.status)

        if Primitive.to_proto(self.tags):
            request.resource.tags.extend(Primitive.to_proto(self.tags))
        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeInstance(request)

    @classmethod
    def list(self, project, zone, service_account_file=""):
        stub = instance_pb2_grpc.ComputeInstanceServiceStub(channel.Channel())
        request = instance_pb2.ListComputeInstanceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Zone = zone

        return stub.ListComputeInstance(request).items

    def to_proto(self):
        resource = instance_pb2.ComputeInstance()
        if Primitive.to_proto(self.can_ip_forward):
            resource.can_ip_forward = Primitive.to_proto(self.can_ip_forward)
        if Primitive.to_proto(self.deletion_protection):
            resource.deletion_protection = Primitive.to_proto(self.deletion_protection)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if InstanceDisksArray.to_proto(self.disks):
            resource.disks.extend(InstanceDisksArray.to_proto(self.disks))
        if InstanceGuestAcceleratorsArray.to_proto(self.guest_accelerators):
            resource.guest_accelerators.extend(
                InstanceGuestAcceleratorsArray.to_proto(self.guest_accelerators)
            )
        if Primitive.to_proto(self.hostname):
            resource.hostname = Primitive.to_proto(self.hostname)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.metadata):
            resource.metadata = Primitive.to_proto(self.metadata)
        if Primitive.to_proto(self.machine_type):
            resource.machine_type = Primitive.to_proto(self.machine_type)
        if Primitive.to_proto(self.min_cpu_platform):
            resource.min_cpu_platform = Primitive.to_proto(self.min_cpu_platform)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if InstanceNetworkInterfacesArray.to_proto(self.network_interfaces):
            resource.network_interfaces.extend(
                InstanceNetworkInterfacesArray.to_proto(self.network_interfaces)
            )
        if InstanceScheduling.to_proto(self.scheduling):
            resource.scheduling.CopyFrom(InstanceScheduling.to_proto(self.scheduling))
        else:
            resource.ClearField("scheduling")
        if InstanceServiceAccountsArray.to_proto(self.service_accounts):
            resource.service_accounts.extend(
                InstanceServiceAccountsArray.to_proto(self.service_accounts)
            )
        if InstanceShieldedInstanceConfig.to_proto(self.shielded_instance_config):
            resource.shielded_instance_config.CopyFrom(
                InstanceShieldedInstanceConfig.to_proto(self.shielded_instance_config)
            )
        else:
            resource.ClearField("shielded_instance_config")
        if InstanceStatusEnum.to_proto(self.status):
            resource.status = InstanceStatusEnum.to_proto(self.status)
        if Primitive.to_proto(self.tags):
            resource.tags.extend(Primitive.to_proto(self.tags))
        if Primitive.to_proto(self.zone):
            resource.zone = Primitive.to_proto(self.zone)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class InstanceDisks(object):
    def __init__(
        self,
        auto_delete: bool = None,
        boot: bool = None,
        device_name: str = None,
        disk_encryption_key: dict = None,
        index: int = None,
        initialize_params: dict = None,
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
        self.interface = interface
        self.mode = mode
        self.source = source
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.ComputeInstanceDisks()
        if Primitive.to_proto(resource.auto_delete):
            res.auto_delete = Primitive.to_proto(resource.auto_delete)
        if Primitive.to_proto(resource.boot):
            res.boot = Primitive.to_proto(resource.boot)
        if Primitive.to_proto(resource.device_name):
            res.device_name = Primitive.to_proto(resource.device_name)
        if InstanceDisksDiskEncryptionKey.to_proto(resource.disk_encryption_key):
            res.disk_encryption_key.CopyFrom(
                InstanceDisksDiskEncryptionKey.to_proto(resource.disk_encryption_key)
            )
        else:
            res.ClearField("disk_encryption_key")
        if Primitive.to_proto(resource.index):
            res.index = Primitive.to_proto(resource.index)
        if InstanceDisksInitializeParams.to_proto(resource.initialize_params):
            res.initialize_params.CopyFrom(
                InstanceDisksInitializeParams.to_proto(resource.initialize_params)
            )
        else:
            res.ClearField("initialize_params")
        if InstanceDisksInterfaceEnum.to_proto(resource.interface):
            res.interface = InstanceDisksInterfaceEnum.to_proto(resource.interface)
        if InstanceDisksModeEnum.to_proto(resource.mode):
            res.mode = InstanceDisksModeEnum.to_proto(resource.mode)
        if Primitive.to_proto(resource.source):
            res.source = Primitive.to_proto(resource.source)
        if InstanceDisksTypeEnum.to_proto(resource.type):
            res.type = InstanceDisksTypeEnum.to_proto(resource.type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDisks(
            auto_delete=Primitive.from_proto(resource.auto_delete),
            boot=Primitive.from_proto(resource.boot),
            device_name=Primitive.from_proto(resource.device_name),
            disk_encryption_key=InstanceDisksDiskEncryptionKey.from_proto(
                resource.disk_encryption_key
            ),
            index=Primitive.from_proto(resource.index),
            initialize_params=InstanceDisksInitializeParams.from_proto(
                resource.initialize_params
            ),
            interface=InstanceDisksInterfaceEnum.from_proto(resource.interface),
            mode=InstanceDisksModeEnum.from_proto(resource.mode),
            source=Primitive.from_proto(resource.source),
            type=InstanceDisksTypeEnum.from_proto(resource.type),
        )


class InstanceDisksArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDisks.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDisks.from_proto(i) for i in resources]


class InstanceDisksDiskEncryptionKey(object):
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

        res = instance_pb2.ComputeInstanceDisksDiskEncryptionKey()
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

        return InstanceDisksDiskEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
            rsa_encrypted_key=Primitive.from_proto(resource.rsa_encrypted_key),
            sha256=Primitive.from_proto(resource.sha256),
        )


class InstanceDisksDiskEncryptionKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDisksDiskEncryptionKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDisksDiskEncryptionKey.from_proto(i) for i in resources]


class InstanceDisksInitializeParams(object):
    def __init__(
        self,
        disk_name: str = None,
        disk_size_gb: int = None,
        disk_type: str = None,
        source_image: str = None,
        source_image_encryption_key: dict = None,
    ):
        self.disk_name = disk_name
        self.disk_size_gb = disk_size_gb
        self.disk_type = disk_type
        self.source_image = source_image
        self.source_image_encryption_key = source_image_encryption_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.ComputeInstanceDisksInitializeParams()
        if Primitive.to_proto(resource.disk_name):
            res.disk_name = Primitive.to_proto(resource.disk_name)
        if Primitive.to_proto(resource.disk_size_gb):
            res.disk_size_gb = Primitive.to_proto(resource.disk_size_gb)
        if Primitive.to_proto(resource.disk_type):
            res.disk_type = Primitive.to_proto(resource.disk_type)
        if Primitive.to_proto(resource.source_image):
            res.source_image = Primitive.to_proto(resource.source_image)
        if InstanceDisksInitializeParamsSourceImageEncryptionKey.to_proto(
            resource.source_image_encryption_key
        ):
            res.source_image_encryption_key.CopyFrom(
                InstanceDisksInitializeParamsSourceImageEncryptionKey.to_proto(
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

        return InstanceDisksInitializeParams(
            disk_name=Primitive.from_proto(resource.disk_name),
            disk_size_gb=Primitive.from_proto(resource.disk_size_gb),
            disk_type=Primitive.from_proto(resource.disk_type),
            source_image=Primitive.from_proto(resource.source_image),
            source_image_encryption_key=InstanceDisksInitializeParamsSourceImageEncryptionKey.from_proto(
                resource.source_image_encryption_key
            ),
        )


class InstanceDisksInitializeParamsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDisksInitializeParams.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDisksInitializeParams.from_proto(i) for i in resources]


class InstanceDisksInitializeParamsSourceImageEncryptionKey(object):
    def __init__(self, raw_key: str = None, sha256: str = None):
        self.raw_key = raw_key
        self.sha256 = sha256

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.ComputeInstanceDisksInitializeParamsSourceImageEncryptionKey()
        )
        if Primitive.to_proto(resource.raw_key):
            res.raw_key = Primitive.to_proto(resource.raw_key)
        if Primitive.to_proto(resource.sha256):
            res.sha256 = Primitive.to_proto(resource.sha256)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDisksInitializeParamsSourceImageEncryptionKey(
            raw_key=Primitive.from_proto(resource.raw_key),
            sha256=Primitive.from_proto(resource.sha256),
        )


class InstanceDisksInitializeParamsSourceImageEncryptionKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDisksInitializeParamsSourceImageEncryptionKey.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDisksInitializeParamsSourceImageEncryptionKey.from_proto(i)
            for i in resources
        ]


class InstanceGuestAccelerators(object):
    def __init__(self, accelerator_count: int = None, accelerator_type: str = None):
        self.accelerator_count = accelerator_count
        self.accelerator_type = accelerator_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.ComputeInstanceGuestAccelerators()
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGuestAccelerators(
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
        )


class InstanceGuestAcceleratorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGuestAccelerators.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGuestAccelerators.from_proto(i) for i in resources]


class InstanceNetworkInterfaces(object):
    def __init__(
        self,
        access_configs: list = None,
        ipv6_access_configs: list = None,
        alias_ip_ranges: list = None,
        name: str = None,
        network: str = None,
        network_ip: str = None,
        subnetwork: str = None,
    ):
        self.access_configs = access_configs
        self.ipv6_access_configs = ipv6_access_configs
        self.alias_ip_ranges = alias_ip_ranges
        self.name = name
        self.network = network
        self.network_ip = network_ip
        self.subnetwork = subnetwork

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.ComputeInstanceNetworkInterfaces()
        if InstanceNetworkInterfacesAccessConfigsArray.to_proto(
            resource.access_configs
        ):
            res.access_configs.extend(
                InstanceNetworkInterfacesAccessConfigsArray.to_proto(
                    resource.access_configs
                )
            )
        if InstanceNetworkInterfacesIPv6AccessConfigsArray.to_proto(
            resource.ipv6_access_configs
        ):
            res.ipv6_access_configs.extend(
                InstanceNetworkInterfacesIPv6AccessConfigsArray.to_proto(
                    resource.ipv6_access_configs
                )
            )
        if InstanceNetworkInterfacesAliasIPRangesArray.to_proto(
            resource.alias_ip_ranges
        ):
            res.alias_ip_ranges.extend(
                InstanceNetworkInterfacesAliasIPRangesArray.to_proto(
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

        return InstanceNetworkInterfaces(
            access_configs=InstanceNetworkInterfacesAccessConfigsArray.from_proto(
                resource.access_configs
            ),
            ipv6_access_configs=InstanceNetworkInterfacesIPv6AccessConfigsArray.from_proto(
                resource.ipv6_access_configs
            ),
            alias_ip_ranges=InstanceNetworkInterfacesAliasIPRangesArray.from_proto(
                resource.alias_ip_ranges
            ),
            name=Primitive.from_proto(resource.name),
            network=Primitive.from_proto(resource.network),
            network_ip=Primitive.from_proto(resource.network_ip),
            subnetwork=Primitive.from_proto(resource.subnetwork),
        )


class InstanceNetworkInterfacesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceNetworkInterfaces.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceNetworkInterfaces.from_proto(i) for i in resources]


class InstanceNetworkInterfacesAccessConfigs(object):
    def __init__(
        self,
        name: str = None,
        nat_ip: str = None,
        external_ipv6: str = None,
        external_ipv6_prefix_length: str = None,
        set_public_ptr: bool = None,
        public_ptr_domain_name: str = None,
        network_tier: str = None,
        type: str = None,
    ):
        self.name = name
        self.nat_ip = nat_ip
        self.external_ipv6 = external_ipv6
        self.external_ipv6_prefix_length = external_ipv6_prefix_length
        self.set_public_ptr = set_public_ptr
        self.public_ptr_domain_name = public_ptr_domain_name
        self.network_tier = network_tier
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.ComputeInstanceNetworkInterfacesAccessConfigs()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.nat_ip):
            res.nat_ip = Primitive.to_proto(resource.nat_ip)
        if Primitive.to_proto(resource.external_ipv6):
            res.external_ipv6 = Primitive.to_proto(resource.external_ipv6)
        if Primitive.to_proto(resource.external_ipv6_prefix_length):
            res.external_ipv6_prefix_length = Primitive.to_proto(
                resource.external_ipv6_prefix_length
            )
        if Primitive.to_proto(resource.set_public_ptr):
            res.set_public_ptr = Primitive.to_proto(resource.set_public_ptr)
        if Primitive.to_proto(resource.public_ptr_domain_name):
            res.public_ptr_domain_name = Primitive.to_proto(
                resource.public_ptr_domain_name
            )
        if InstanceNetworkInterfacesAccessConfigsNetworkTierEnum.to_proto(
            resource.network_tier
        ):
            res.network_tier = (
                InstanceNetworkInterfacesAccessConfigsNetworkTierEnum.to_proto(
                    resource.network_tier
                )
            )
        if InstanceNetworkInterfacesAccessConfigsTypeEnum.to_proto(resource.type):
            res.type = InstanceNetworkInterfacesAccessConfigsTypeEnum.to_proto(
                resource.type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNetworkInterfacesAccessConfigs(
            name=Primitive.from_proto(resource.name),
            nat_ip=Primitive.from_proto(resource.nat_ip),
            external_ipv6=Primitive.from_proto(resource.external_ipv6),
            external_ipv6_prefix_length=Primitive.from_proto(
                resource.external_ipv6_prefix_length
            ),
            set_public_ptr=Primitive.from_proto(resource.set_public_ptr),
            public_ptr_domain_name=Primitive.from_proto(
                resource.public_ptr_domain_name
            ),
            network_tier=InstanceNetworkInterfacesAccessConfigsNetworkTierEnum.from_proto(
                resource.network_tier
            ),
            type=InstanceNetworkInterfacesAccessConfigsTypeEnum.from_proto(
                resource.type
            ),
        )


class InstanceNetworkInterfacesAccessConfigsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceNetworkInterfacesAccessConfigs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceNetworkInterfacesAccessConfigs.from_proto(i) for i in resources]


class InstanceNetworkInterfacesIPv6AccessConfigs(object):
    def __init__(
        self,
        name: str = None,
        nat_ip: str = None,
        external_ipv6: str = None,
        external_ipv6_prefix_length: str = None,
        set_public_ptr: bool = None,
        public_ptr_domain_name: str = None,
        network_tier: str = None,
        type: str = None,
    ):
        self.name = name
        self.nat_ip = nat_ip
        self.external_ipv6 = external_ipv6
        self.external_ipv6_prefix_length = external_ipv6_prefix_length
        self.set_public_ptr = set_public_ptr
        self.public_ptr_domain_name = public_ptr_domain_name
        self.network_tier = network_tier
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.ComputeInstanceNetworkInterfacesIPv6AccessConfigs()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.nat_ip):
            res.nat_ip = Primitive.to_proto(resource.nat_ip)
        if Primitive.to_proto(resource.external_ipv6):
            res.external_ipv6 = Primitive.to_proto(resource.external_ipv6)
        if Primitive.to_proto(resource.external_ipv6_prefix_length):
            res.external_ipv6_prefix_length = Primitive.to_proto(
                resource.external_ipv6_prefix_length
            )
        if Primitive.to_proto(resource.set_public_ptr):
            res.set_public_ptr = Primitive.to_proto(resource.set_public_ptr)
        if Primitive.to_proto(resource.public_ptr_domain_name):
            res.public_ptr_domain_name = Primitive.to_proto(
                resource.public_ptr_domain_name
            )
        if InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum.to_proto(
            resource.network_tier
        ):
            res.network_tier = (
                InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum.to_proto(
                    resource.network_tier
                )
            )
        if InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum.to_proto(resource.type):
            res.type = InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum.to_proto(
                resource.type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNetworkInterfacesIPv6AccessConfigs(
            name=Primitive.from_proto(resource.name),
            nat_ip=Primitive.from_proto(resource.nat_ip),
            external_ipv6=Primitive.from_proto(resource.external_ipv6),
            external_ipv6_prefix_length=Primitive.from_proto(
                resource.external_ipv6_prefix_length
            ),
            set_public_ptr=Primitive.from_proto(resource.set_public_ptr),
            public_ptr_domain_name=Primitive.from_proto(
                resource.public_ptr_domain_name
            ),
            network_tier=InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum.from_proto(
                resource.network_tier
            ),
            type=InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum.from_proto(
                resource.type
            ),
        )


class InstanceNetworkInterfacesIPv6AccessConfigsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNetworkInterfacesIPv6AccessConfigs.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNetworkInterfacesIPv6AccessConfigs.from_proto(i) for i in resources
        ]


class InstanceNetworkInterfacesAliasIPRanges(object):
    def __init__(self, ip_cidr_range: str = None, subnetwork_range_name: str = None):
        self.ip_cidr_range = ip_cidr_range
        self.subnetwork_range_name = subnetwork_range_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.ComputeInstanceNetworkInterfacesAliasIPRanges()
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

        return InstanceNetworkInterfacesAliasIPRanges(
            ip_cidr_range=Primitive.from_proto(resource.ip_cidr_range),
            subnetwork_range_name=Primitive.from_proto(resource.subnetwork_range_name),
        )


class InstanceNetworkInterfacesAliasIPRangesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceNetworkInterfacesAliasIPRanges.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceNetworkInterfacesAliasIPRanges.from_proto(i) for i in resources]


class InstanceScheduling(object):
    def __init__(
        self,
        automatic_restart: bool = None,
        on_host_maintenance: str = None,
        preemptible: bool = None,
    ):
        self.automatic_restart = automatic_restart
        self.on_host_maintenance = on_host_maintenance
        self.preemptible = preemptible

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.ComputeInstanceScheduling()
        if Primitive.to_proto(resource.automatic_restart):
            res.automatic_restart = Primitive.to_proto(resource.automatic_restart)
        if Primitive.to_proto(resource.on_host_maintenance):
            res.on_host_maintenance = Primitive.to_proto(resource.on_host_maintenance)
        if Primitive.to_proto(resource.preemptible):
            res.preemptible = Primitive.to_proto(resource.preemptible)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceScheduling(
            automatic_restart=Primitive.from_proto(resource.automatic_restart),
            on_host_maintenance=Primitive.from_proto(resource.on_host_maintenance),
            preemptible=Primitive.from_proto(resource.preemptible),
        )


class InstanceSchedulingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceScheduling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceScheduling.from_proto(i) for i in resources]


class InstanceServiceAccounts(object):
    def __init__(self, email: str = None, scopes: list = None):
        self.email = email
        self.scopes = scopes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.ComputeInstanceServiceAccounts()
        if Primitive.to_proto(resource.email):
            res.email = Primitive.to_proto(resource.email)
        if Primitive.to_proto(resource.scopes):
            res.scopes.extend(Primitive.to_proto(resource.scopes))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceServiceAccounts(
            email=Primitive.from_proto(resource.email),
            scopes=Primitive.from_proto(resource.scopes),
        )


class InstanceServiceAccountsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceServiceAccounts.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceServiceAccounts.from_proto(i) for i in resources]


class InstanceShieldedInstanceConfig(object):
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

        res = instance_pb2.ComputeInstanceShieldedInstanceConfig()
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

        return InstanceShieldedInstanceConfig(
            enable_secure_boot=Primitive.from_proto(resource.enable_secure_boot),
            enable_vtpm=Primitive.from_proto(resource.enable_vtpm),
            enable_integrity_monitoring=Primitive.from_proto(
                resource.enable_integrity_monitoring
            ),
        )


class InstanceShieldedInstanceConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceShieldedInstanceConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceShieldedInstanceConfig.from_proto(i) for i in resources]


class InstanceDisksInterfaceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceDisksInterfaceEnum.Value(
            "ComputeInstanceDisksInterfaceEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceDisksInterfaceEnum.Name(resource)[
            len("ComputeInstanceDisksInterfaceEnum") :
        ]


class InstanceDisksModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceDisksModeEnum.Value(
            "ComputeInstanceDisksModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceDisksModeEnum.Name(resource)[
            len("ComputeInstanceDisksModeEnum") :
        ]


class InstanceDisksTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceDisksTypeEnum.Value(
            "ComputeInstanceDisksTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceDisksTypeEnum.Name(resource)[
            len("ComputeInstanceDisksTypeEnum") :
        ]


class InstanceNetworkInterfacesAccessConfigsNetworkTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum.Value(
            "ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum.Name(
            resource
        )[
            len("ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum") :
        ]


class InstanceNetworkInterfacesAccessConfigsTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum.Value(
            "ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum.Name(
            resource
        )[len("ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum") :]


class InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum.Value(
            "ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum.Name(
            resource
        )[
            len("ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum") :
        ]


class InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum.Value(
            "ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            instance_pb2.ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum.Name(
                resource
            )[len("ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum") :]
        )


class InstanceStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceStatusEnum.Value(
            "ComputeInstanceStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ComputeInstanceStatusEnum.Name(resource)[
            len("ComputeInstanceStatusEnum") :
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
