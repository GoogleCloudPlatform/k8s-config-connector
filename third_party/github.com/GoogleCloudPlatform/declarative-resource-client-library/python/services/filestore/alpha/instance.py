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
from google3.cloud.graphite.mmv2.services.google.filestore import instance_pb2
from google3.cloud.graphite.mmv2.services.google.filestore import instance_pb2_grpc

from typing import List


class Instance(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        state: str = None,
        status_message: str = None,
        create_time: str = None,
        tier: str = None,
        labels: dict = None,
        file_shares: list = None,
        networks: list = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.tier = tier
        self.labels = labels
        self.file_shares = file_shares
        self.networks = networks
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_pb2_grpc.FilestoreAlphaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ApplyFilestoreAlphaInstanceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if InstanceTierEnum.to_proto(self.tier):
            request.resource.tier = InstanceTierEnum.to_proto(self.tier)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if InstanceFileSharesArray.to_proto(self.file_shares):
            request.resource.file_shares.extend(
                InstanceFileSharesArray.to_proto(self.file_shares)
            )
        if InstanceNetworksArray.to_proto(self.networks):
            request.resource.networks.extend(
                InstanceNetworksArray.to_proto(self.networks)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyFilestoreAlphaInstance(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.state = InstanceStateEnum.from_proto(response.state)
        self.status_message = Primitive.from_proto(response.status_message)
        self.create_time = Primitive.from_proto(response.create_time)
        self.tier = InstanceTierEnum.from_proto(response.tier)
        self.labels = Primitive.from_proto(response.labels)
        self.file_shares = InstanceFileSharesArray.from_proto(response.file_shares)
        self.networks = InstanceNetworksArray.from_proto(response.networks)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = instance_pb2_grpc.FilestoreAlphaInstanceServiceStub(channel.Channel())
        request = instance_pb2.DeleteFilestoreAlphaInstanceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if InstanceTierEnum.to_proto(self.tier):
            request.resource.tier = InstanceTierEnum.to_proto(self.tier)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if InstanceFileSharesArray.to_proto(self.file_shares):
            request.resource.file_shares.extend(
                InstanceFileSharesArray.to_proto(self.file_shares)
            )
        if InstanceNetworksArray.to_proto(self.networks):
            request.resource.networks.extend(
                InstanceNetworksArray.to_proto(self.networks)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteFilestoreAlphaInstance(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = instance_pb2_grpc.FilestoreAlphaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ListFilestoreAlphaInstanceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListFilestoreAlphaInstance(request).items

    def to_proto(self):
        resource = instance_pb2.FilestoreAlphaInstance()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if InstanceTierEnum.to_proto(self.tier):
            resource.tier = InstanceTierEnum.to_proto(self.tier)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if InstanceFileSharesArray.to_proto(self.file_shares):
            resource.file_shares.extend(
                InstanceFileSharesArray.to_proto(self.file_shares)
            )
        if InstanceNetworksArray.to_proto(self.networks):
            resource.networks.extend(InstanceNetworksArray.to_proto(self.networks))
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class InstanceFileShares(object):
    def __init__(
        self,
        name: str = None,
        capacity_gb: int = None,
        source_backup: str = None,
        nfs_export_options: list = None,
    ):
        self.name = name
        self.capacity_gb = capacity_gb
        self.source_backup = source_backup
        self.nfs_export_options = nfs_export_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.FilestoreAlphaInstanceFileShares()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.capacity_gb):
            res.capacity_gb = Primitive.to_proto(resource.capacity_gb)
        if Primitive.to_proto(resource.source_backup):
            res.source_backup = Primitive.to_proto(resource.source_backup)
        if InstanceFileSharesNfsExportOptionsArray.to_proto(
            resource.nfs_export_options
        ):
            res.nfs_export_options.extend(
                InstanceFileSharesNfsExportOptionsArray.to_proto(
                    resource.nfs_export_options
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFileShares(
            name=Primitive.from_proto(resource.name),
            capacity_gb=Primitive.from_proto(resource.capacity_gb),
            source_backup=Primitive.from_proto(resource.source_backup),
            nfs_export_options=InstanceFileSharesNfsExportOptionsArray.from_proto(
                resource.nfs_export_options
            ),
        )


class InstanceFileSharesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFileShares.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceFileShares.from_proto(i) for i in resources]


class InstanceFileSharesNfsExportOptions(object):
    def __init__(
        self,
        ip_ranges: list = None,
        access_mode: str = None,
        squash_mode: str = None,
        anon_uid: int = None,
        anon_gid: int = None,
    ):
        self.ip_ranges = ip_ranges
        self.access_mode = access_mode
        self.squash_mode = squash_mode
        self.anon_uid = anon_uid
        self.anon_gid = anon_gid

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.FilestoreAlphaInstanceFileSharesNfsExportOptions()
        if Primitive.to_proto(resource.ip_ranges):
            res.ip_ranges.extend(Primitive.to_proto(resource.ip_ranges))
        if InstanceFileSharesNfsExportOptionsAccessModeEnum.to_proto(
            resource.access_mode
        ):
            res.access_mode = InstanceFileSharesNfsExportOptionsAccessModeEnum.to_proto(
                resource.access_mode
            )
        if InstanceFileSharesNfsExportOptionsSquashModeEnum.to_proto(
            resource.squash_mode
        ):
            res.squash_mode = InstanceFileSharesNfsExportOptionsSquashModeEnum.to_proto(
                resource.squash_mode
            )
        if Primitive.to_proto(resource.anon_uid):
            res.anon_uid = Primitive.to_proto(resource.anon_uid)
        if Primitive.to_proto(resource.anon_gid):
            res.anon_gid = Primitive.to_proto(resource.anon_gid)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFileSharesNfsExportOptions(
            ip_ranges=Primitive.from_proto(resource.ip_ranges),
            access_mode=InstanceFileSharesNfsExportOptionsAccessModeEnum.from_proto(
                resource.access_mode
            ),
            squash_mode=InstanceFileSharesNfsExportOptionsSquashModeEnum.from_proto(
                resource.squash_mode
            ),
            anon_uid=Primitive.from_proto(resource.anon_uid),
            anon_gid=Primitive.from_proto(resource.anon_gid),
        )


class InstanceFileSharesNfsExportOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFileSharesNfsExportOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceFileSharesNfsExportOptions.from_proto(i) for i in resources]


class InstanceNetworks(object):
    def __init__(
        self,
        network: str = None,
        modes: list = None,
        reserved_ip_range: str = None,
        ip_addresses: list = None,
    ):
        self.network = network
        self.modes = modes
        self.reserved_ip_range = reserved_ip_range
        self.ip_addresses = ip_addresses

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.FilestoreAlphaInstanceNetworks()
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        if InstanceNetworksModesEnumArray.to_proto(resource.modes):
            res.modes.extend(InstanceNetworksModesEnumArray.to_proto(resource.modes))
        if Primitive.to_proto(resource.reserved_ip_range):
            res.reserved_ip_range = Primitive.to_proto(resource.reserved_ip_range)
        if Primitive.to_proto(resource.ip_addresses):
            res.ip_addresses.extend(Primitive.to_proto(resource.ip_addresses))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNetworks(
            network=Primitive.from_proto(resource.network),
            modes=InstanceNetworksModesEnumArray.from_proto(resource.modes),
            reserved_ip_range=Primitive.from_proto(resource.reserved_ip_range),
            ip_addresses=Primitive.from_proto(resource.ip_addresses),
        )


class InstanceNetworksArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceNetworks.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceNetworks.from_proto(i) for i in resources]


class InstanceStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.FilestoreAlphaInstanceStateEnum.Value(
            "FilestoreAlphaInstanceStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.FilestoreAlphaInstanceStateEnum.Name(resource)[
            len("FilestoreAlphaInstanceStateEnum") :
        ]


class InstanceTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.FilestoreAlphaInstanceTierEnum.Value(
            "FilestoreAlphaInstanceTierEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.FilestoreAlphaInstanceTierEnum.Name(resource)[
            len("FilestoreAlphaInstanceTierEnum") :
        ]


class InstanceFileSharesNfsExportOptionsAccessModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum.Value(
            "FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum.Name(
            resource
        )[
            len("FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum") :
        ]


class InstanceFileSharesNfsExportOptionsSquashModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum.Value(
            "FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum.Name(
            resource
        )[
            len("FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum") :
        ]


class InstanceNetworksModesEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.FilestoreAlphaInstanceNetworksModesEnum.Value(
            "FilestoreAlphaInstanceNetworksModesEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.FilestoreAlphaInstanceNetworksModesEnum.Name(resource)[
            len("FilestoreAlphaInstanceNetworksModesEnum") :
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
