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
from google3.cloud.graphite.mmv2.services.google.dataplex import asset_pb2
from google3.cloud.graphite.mmv2.services.google.dataplex import asset_pb2_grpc

from typing import List


class Asset(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        uid: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        description: str = None,
        state: str = None,
        resource_spec: dict = None,
        resource_status: dict = None,
        security_status: dict = None,
        discovery_spec: dict = None,
        discovery_status: dict = None,
        project: str = None,
        location: str = None,
        lake: str = None,
        dataplex_zone: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.labels = labels
        self.description = description
        self.resource_spec = resource_spec
        self.discovery_spec = discovery_spec
        self.project = project
        self.location = location
        self.lake = lake
        self.dataplex_zone = dataplex_zone
        self.service_account_file = service_account_file

    def apply(self):
        stub = asset_pb2_grpc.DataplexBetaAssetServiceStub(channel.Channel())
        request = asset_pb2.ApplyDataplexBetaAssetRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if AssetResourceSpec.to_proto(self.resource_spec):
            request.resource.resource_spec.CopyFrom(
                AssetResourceSpec.to_proto(self.resource_spec)
            )
        else:
            request.resource.ClearField("resource_spec")
        if AssetDiscoverySpec.to_proto(self.discovery_spec):
            request.resource.discovery_spec.CopyFrom(
                AssetDiscoverySpec.to_proto(self.discovery_spec)
            )
        else:
            request.resource.ClearField("discovery_spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.lake):
            request.resource.lake = Primitive.to_proto(self.lake)

        if Primitive.to_proto(self.dataplex_zone):
            request.resource.dataplex_zone = Primitive.to_proto(self.dataplex_zone)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDataplexBetaAsset(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.uid = Primitive.from_proto(response.uid)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.description = Primitive.from_proto(response.description)
        self.state = AssetStateEnum.from_proto(response.state)
        self.resource_spec = AssetResourceSpec.from_proto(response.resource_spec)
        self.resource_status = AssetResourceStatus.from_proto(response.resource_status)
        self.security_status = AssetSecurityStatus.from_proto(response.security_status)
        self.discovery_spec = AssetDiscoverySpec.from_proto(response.discovery_spec)
        self.discovery_status = AssetDiscoveryStatus.from_proto(
            response.discovery_status
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.lake = Primitive.from_proto(response.lake)
        self.dataplex_zone = Primitive.from_proto(response.dataplex_zone)

    def delete(self):
        stub = asset_pb2_grpc.DataplexBetaAssetServiceStub(channel.Channel())
        request = asset_pb2.DeleteDataplexBetaAssetRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if AssetResourceSpec.to_proto(self.resource_spec):
            request.resource.resource_spec.CopyFrom(
                AssetResourceSpec.to_proto(self.resource_spec)
            )
        else:
            request.resource.ClearField("resource_spec")
        if AssetDiscoverySpec.to_proto(self.discovery_spec):
            request.resource.discovery_spec.CopyFrom(
                AssetDiscoverySpec.to_proto(self.discovery_spec)
            )
        else:
            request.resource.ClearField("discovery_spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.lake):
            request.resource.lake = Primitive.to_proto(self.lake)

        if Primitive.to_proto(self.dataplex_zone):
            request.resource.dataplex_zone = Primitive.to_proto(self.dataplex_zone)

        response = stub.DeleteDataplexBetaAsset(request)

    @classmethod
    def list(self, project, location, dataplexZone, lake, service_account_file=""):
        stub = asset_pb2_grpc.DataplexBetaAssetServiceStub(channel.Channel())
        request = asset_pb2.ListDataplexBetaAssetRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.DataplexZone = dataplexZone

        request.Lake = lake

        return stub.ListDataplexBetaAsset(request).items

    def to_proto(self):
        resource = asset_pb2.DataplexBetaAsset()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if AssetResourceSpec.to_proto(self.resource_spec):
            resource.resource_spec.CopyFrom(
                AssetResourceSpec.to_proto(self.resource_spec)
            )
        else:
            resource.ClearField("resource_spec")
        if AssetDiscoverySpec.to_proto(self.discovery_spec):
            resource.discovery_spec.CopyFrom(
                AssetDiscoverySpec.to_proto(self.discovery_spec)
            )
        else:
            resource.ClearField("discovery_spec")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.lake):
            resource.lake = Primitive.to_proto(self.lake)
        if Primitive.to_proto(self.dataplex_zone):
            resource.dataplex_zone = Primitive.to_proto(self.dataplex_zone)
        return resource


class AssetResourceSpec(object):
    def __init__(
        self, name: str = None, type: str = None, read_access_mode: str = None
    ):
        self.name = name
        self.type = type
        self.read_access_mode = read_access_mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = asset_pb2.DataplexBetaAssetResourceSpec()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if AssetResourceSpecTypeEnum.to_proto(resource.type):
            res.type = AssetResourceSpecTypeEnum.to_proto(resource.type)
        if AssetResourceSpecReadAccessModeEnum.to_proto(resource.read_access_mode):
            res.read_access_mode = AssetResourceSpecReadAccessModeEnum.to_proto(
                resource.read_access_mode
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AssetResourceSpec(
            name=Primitive.from_proto(resource.name),
            type=AssetResourceSpecTypeEnum.from_proto(resource.type),
            read_access_mode=AssetResourceSpecReadAccessModeEnum.from_proto(
                resource.read_access_mode
            ),
        )


class AssetResourceSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AssetResourceSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AssetResourceSpec.from_proto(i) for i in resources]


class AssetResourceStatus(object):
    def __init__(self, state: str = None, message: str = None, update_time: str = None):
        self.state = state
        self.message = message
        self.update_time = update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = asset_pb2.DataplexBetaAssetResourceStatus()
        if AssetResourceStatusStateEnum.to_proto(resource.state):
            res.state = AssetResourceStatusStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AssetResourceStatus(
            state=AssetResourceStatusStateEnum.from_proto(resource.state),
            message=Primitive.from_proto(resource.message),
            update_time=Primitive.from_proto(resource.update_time),
        )


class AssetResourceStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AssetResourceStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AssetResourceStatus.from_proto(i) for i in resources]


class AssetSecurityStatus(object):
    def __init__(self, state: str = None, message: str = None, update_time: str = None):
        self.state = state
        self.message = message
        self.update_time = update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = asset_pb2.DataplexBetaAssetSecurityStatus()
        if AssetSecurityStatusStateEnum.to_proto(resource.state):
            res.state = AssetSecurityStatusStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AssetSecurityStatus(
            state=AssetSecurityStatusStateEnum.from_proto(resource.state),
            message=Primitive.from_proto(resource.message),
            update_time=Primitive.from_proto(resource.update_time),
        )


class AssetSecurityStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AssetSecurityStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AssetSecurityStatus.from_proto(i) for i in resources]


class AssetDiscoverySpec(object):
    def __init__(
        self,
        enabled: bool = None,
        include_patterns: list = None,
        exclude_patterns: list = None,
        csv_options: dict = None,
        json_options: dict = None,
        schedule: str = None,
    ):
        self.enabled = enabled
        self.include_patterns = include_patterns
        self.exclude_patterns = exclude_patterns
        self.csv_options = csv_options
        self.json_options = json_options
        self.schedule = schedule

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = asset_pb2.DataplexBetaAssetDiscoverySpec()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.include_patterns):
            res.include_patterns.extend(Primitive.to_proto(resource.include_patterns))
        if Primitive.to_proto(resource.exclude_patterns):
            res.exclude_patterns.extend(Primitive.to_proto(resource.exclude_patterns))
        if AssetDiscoverySpecCsvOptions.to_proto(resource.csv_options):
            res.csv_options.CopyFrom(
                AssetDiscoverySpecCsvOptions.to_proto(resource.csv_options)
            )
        else:
            res.ClearField("csv_options")
        if AssetDiscoverySpecJsonOptions.to_proto(resource.json_options):
            res.json_options.CopyFrom(
                AssetDiscoverySpecJsonOptions.to_proto(resource.json_options)
            )
        else:
            res.ClearField("json_options")
        if Primitive.to_proto(resource.schedule):
            res.schedule = Primitive.to_proto(resource.schedule)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AssetDiscoverySpec(
            enabled=Primitive.from_proto(resource.enabled),
            include_patterns=Primitive.from_proto(resource.include_patterns),
            exclude_patterns=Primitive.from_proto(resource.exclude_patterns),
            csv_options=AssetDiscoverySpecCsvOptions.from_proto(resource.csv_options),
            json_options=AssetDiscoverySpecJsonOptions.from_proto(
                resource.json_options
            ),
            schedule=Primitive.from_proto(resource.schedule),
        )


class AssetDiscoverySpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AssetDiscoverySpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AssetDiscoverySpec.from_proto(i) for i in resources]


class AssetDiscoverySpecCsvOptions(object):
    def __init__(
        self,
        header_rows: int = None,
        delimiter: str = None,
        encoding: str = None,
        disable_type_inference: bool = None,
    ):
        self.header_rows = header_rows
        self.delimiter = delimiter
        self.encoding = encoding
        self.disable_type_inference = disable_type_inference

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = asset_pb2.DataplexBetaAssetDiscoverySpecCsvOptions()
        if Primitive.to_proto(resource.header_rows):
            res.header_rows = Primitive.to_proto(resource.header_rows)
        if Primitive.to_proto(resource.delimiter):
            res.delimiter = Primitive.to_proto(resource.delimiter)
        if Primitive.to_proto(resource.encoding):
            res.encoding = Primitive.to_proto(resource.encoding)
        if Primitive.to_proto(resource.disable_type_inference):
            res.disable_type_inference = Primitive.to_proto(
                resource.disable_type_inference
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AssetDiscoverySpecCsvOptions(
            header_rows=Primitive.from_proto(resource.header_rows),
            delimiter=Primitive.from_proto(resource.delimiter),
            encoding=Primitive.from_proto(resource.encoding),
            disable_type_inference=Primitive.from_proto(
                resource.disable_type_inference
            ),
        )


class AssetDiscoverySpecCsvOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AssetDiscoverySpecCsvOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AssetDiscoverySpecCsvOptions.from_proto(i) for i in resources]


class AssetDiscoverySpecJsonOptions(object):
    def __init__(self, encoding: str = None, disable_type_inference: bool = None):
        self.encoding = encoding
        self.disable_type_inference = disable_type_inference

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = asset_pb2.DataplexBetaAssetDiscoverySpecJsonOptions()
        if Primitive.to_proto(resource.encoding):
            res.encoding = Primitive.to_proto(resource.encoding)
        if Primitive.to_proto(resource.disable_type_inference):
            res.disable_type_inference = Primitive.to_proto(
                resource.disable_type_inference
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AssetDiscoverySpecJsonOptions(
            encoding=Primitive.from_proto(resource.encoding),
            disable_type_inference=Primitive.from_proto(
                resource.disable_type_inference
            ),
        )


class AssetDiscoverySpecJsonOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AssetDiscoverySpecJsonOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AssetDiscoverySpecJsonOptions.from_proto(i) for i in resources]


class AssetDiscoveryStatus(object):
    def __init__(
        self,
        state: str = None,
        message: str = None,
        update_time: str = None,
        last_run_time: str = None,
        stats: dict = None,
        last_run_duration: str = None,
    ):
        self.state = state
        self.message = message
        self.update_time = update_time
        self.last_run_time = last_run_time
        self.stats = stats
        self.last_run_duration = last_run_duration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = asset_pb2.DataplexBetaAssetDiscoveryStatus()
        if AssetDiscoveryStatusStateEnum.to_proto(resource.state):
            res.state = AssetDiscoveryStatusStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        if Primitive.to_proto(resource.last_run_time):
            res.last_run_time = Primitive.to_proto(resource.last_run_time)
        if AssetDiscoveryStatusStats.to_proto(resource.stats):
            res.stats.CopyFrom(AssetDiscoveryStatusStats.to_proto(resource.stats))
        else:
            res.ClearField("stats")
        if Primitive.to_proto(resource.last_run_duration):
            res.last_run_duration = Primitive.to_proto(resource.last_run_duration)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AssetDiscoveryStatus(
            state=AssetDiscoveryStatusStateEnum.from_proto(resource.state),
            message=Primitive.from_proto(resource.message),
            update_time=Primitive.from_proto(resource.update_time),
            last_run_time=Primitive.from_proto(resource.last_run_time),
            stats=AssetDiscoveryStatusStats.from_proto(resource.stats),
            last_run_duration=Primitive.from_proto(resource.last_run_duration),
        )


class AssetDiscoveryStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AssetDiscoveryStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AssetDiscoveryStatus.from_proto(i) for i in resources]


class AssetDiscoveryStatusStats(object):
    def __init__(
        self,
        data_items: int = None,
        data_size: int = None,
        tables: int = None,
        filesets: int = None,
    ):
        self.data_items = data_items
        self.data_size = data_size
        self.tables = tables
        self.filesets = filesets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = asset_pb2.DataplexBetaAssetDiscoveryStatusStats()
        if Primitive.to_proto(resource.data_items):
            res.data_items = Primitive.to_proto(resource.data_items)
        if Primitive.to_proto(resource.data_size):
            res.data_size = Primitive.to_proto(resource.data_size)
        if Primitive.to_proto(resource.tables):
            res.tables = Primitive.to_proto(resource.tables)
        if Primitive.to_proto(resource.filesets):
            res.filesets = Primitive.to_proto(resource.filesets)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AssetDiscoveryStatusStats(
            data_items=Primitive.from_proto(resource.data_items),
            data_size=Primitive.from_proto(resource.data_size),
            tables=Primitive.from_proto(resource.tables),
            filesets=Primitive.from_proto(resource.filesets),
        )


class AssetDiscoveryStatusStatsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AssetDiscoveryStatusStats.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AssetDiscoveryStatusStats.from_proto(i) for i in resources]


class AssetStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetStateEnum.Value(
            "DataplexBetaAssetStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetStateEnum.Name(resource)[
            len("DataplexBetaAssetStateEnum") :
        ]


class AssetResourceSpecTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetResourceSpecTypeEnum.Value(
            "DataplexBetaAssetResourceSpecTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetResourceSpecTypeEnum.Name(resource)[
            len("DataplexBetaAssetResourceSpecTypeEnum") :
        ]


class AssetResourceSpecReadAccessModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetResourceSpecReadAccessModeEnum.Value(
            "DataplexBetaAssetResourceSpecReadAccessModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetResourceSpecReadAccessModeEnum.Name(resource)[
            len("DataplexBetaAssetResourceSpecReadAccessModeEnum") :
        ]


class AssetResourceStatusStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetResourceStatusStateEnum.Value(
            "DataplexBetaAssetResourceStatusStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetResourceStatusStateEnum.Name(resource)[
            len("DataplexBetaAssetResourceStatusStateEnum") :
        ]


class AssetSecurityStatusStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetSecurityStatusStateEnum.Value(
            "DataplexBetaAssetSecurityStatusStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetSecurityStatusStateEnum.Name(resource)[
            len("DataplexBetaAssetSecurityStatusStateEnum") :
        ]


class AssetDiscoveryStatusStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetDiscoveryStatusStateEnum.Value(
            "DataplexBetaAssetDiscoveryStatusStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return asset_pb2.DataplexBetaAssetDiscoveryStatusStateEnum.Name(resource)[
            len("DataplexBetaAssetDiscoveryStatusStateEnum") :
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
