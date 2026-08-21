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
from google3.cloud.graphite.mmv2.services.google.dataplex import zone_pb2
from google3.cloud.graphite.mmv2.services.google.dataplex import zone_pb2_grpc

from typing import List


class Zone(object):
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
        type: str = None,
        discovery_spec: dict = None,
        resource_spec: dict = None,
        asset_status: dict = None,
        project: str = None,
        location: str = None,
        lake: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.labels = labels
        self.description = description
        self.type = type
        self.discovery_spec = discovery_spec
        self.resource_spec = resource_spec
        self.project = project
        self.location = location
        self.lake = lake
        self.service_account_file = service_account_file

    def apply(self):
        stub = zone_pb2_grpc.DataplexAlphaZoneServiceStub(channel.Channel())
        request = zone_pb2.ApplyDataplexAlphaZoneRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ZoneTypeEnum.to_proto(self.type):
            request.resource.type = ZoneTypeEnum.to_proto(self.type)

        if ZoneDiscoverySpec.to_proto(self.discovery_spec):
            request.resource.discovery_spec.CopyFrom(
                ZoneDiscoverySpec.to_proto(self.discovery_spec)
            )
        else:
            request.resource.ClearField("discovery_spec")
        if ZoneResourceSpec.to_proto(self.resource_spec):
            request.resource.resource_spec.CopyFrom(
                ZoneResourceSpec.to_proto(self.resource_spec)
            )
        else:
            request.resource.ClearField("resource_spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.lake):
            request.resource.lake = Primitive.to_proto(self.lake)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDataplexAlphaZone(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.uid = Primitive.from_proto(response.uid)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.description = Primitive.from_proto(response.description)
        self.state = ZoneStateEnum.from_proto(response.state)
        self.type = ZoneTypeEnum.from_proto(response.type)
        self.discovery_spec = ZoneDiscoverySpec.from_proto(response.discovery_spec)
        self.resource_spec = ZoneResourceSpec.from_proto(response.resource_spec)
        self.asset_status = ZoneAssetStatus.from_proto(response.asset_status)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.lake = Primitive.from_proto(response.lake)

    def delete(self):
        stub = zone_pb2_grpc.DataplexAlphaZoneServiceStub(channel.Channel())
        request = zone_pb2.DeleteDataplexAlphaZoneRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ZoneTypeEnum.to_proto(self.type):
            request.resource.type = ZoneTypeEnum.to_proto(self.type)

        if ZoneDiscoverySpec.to_proto(self.discovery_spec):
            request.resource.discovery_spec.CopyFrom(
                ZoneDiscoverySpec.to_proto(self.discovery_spec)
            )
        else:
            request.resource.ClearField("discovery_spec")
        if ZoneResourceSpec.to_proto(self.resource_spec):
            request.resource.resource_spec.CopyFrom(
                ZoneResourceSpec.to_proto(self.resource_spec)
            )
        else:
            request.resource.ClearField("resource_spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.lake):
            request.resource.lake = Primitive.to_proto(self.lake)

        response = stub.DeleteDataplexAlphaZone(request)

    @classmethod
    def list(self, project, location, lake, service_account_file=""):
        stub = zone_pb2_grpc.DataplexAlphaZoneServiceStub(channel.Channel())
        request = zone_pb2.ListDataplexAlphaZoneRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Lake = lake

        return stub.ListDataplexAlphaZone(request).items

    def to_proto(self):
        resource = zone_pb2.DataplexAlphaZone()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if ZoneTypeEnum.to_proto(self.type):
            resource.type = ZoneTypeEnum.to_proto(self.type)
        if ZoneDiscoverySpec.to_proto(self.discovery_spec):
            resource.discovery_spec.CopyFrom(
                ZoneDiscoverySpec.to_proto(self.discovery_spec)
            )
        else:
            resource.ClearField("discovery_spec")
        if ZoneResourceSpec.to_proto(self.resource_spec):
            resource.resource_spec.CopyFrom(
                ZoneResourceSpec.to_proto(self.resource_spec)
            )
        else:
            resource.ClearField("resource_spec")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.lake):
            resource.lake = Primitive.to_proto(self.lake)
        return resource


class ZoneDiscoverySpec(object):
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

        res = zone_pb2.DataplexAlphaZoneDiscoverySpec()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.include_patterns):
            res.include_patterns.extend(Primitive.to_proto(resource.include_patterns))
        if Primitive.to_proto(resource.exclude_patterns):
            res.exclude_patterns.extend(Primitive.to_proto(resource.exclude_patterns))
        if ZoneDiscoverySpecCsvOptions.to_proto(resource.csv_options):
            res.csv_options.CopyFrom(
                ZoneDiscoverySpecCsvOptions.to_proto(resource.csv_options)
            )
        else:
            res.ClearField("csv_options")
        if ZoneDiscoverySpecJsonOptions.to_proto(resource.json_options):
            res.json_options.CopyFrom(
                ZoneDiscoverySpecJsonOptions.to_proto(resource.json_options)
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

        return ZoneDiscoverySpec(
            enabled=Primitive.from_proto(resource.enabled),
            include_patterns=Primitive.from_proto(resource.include_patterns),
            exclude_patterns=Primitive.from_proto(resource.exclude_patterns),
            csv_options=ZoneDiscoverySpecCsvOptions.from_proto(resource.csv_options),
            json_options=ZoneDiscoverySpecJsonOptions.from_proto(resource.json_options),
            schedule=Primitive.from_proto(resource.schedule),
        )


class ZoneDiscoverySpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ZoneDiscoverySpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ZoneDiscoverySpec.from_proto(i) for i in resources]


class ZoneDiscoverySpecCsvOptions(object):
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

        res = zone_pb2.DataplexAlphaZoneDiscoverySpecCsvOptions()
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

        return ZoneDiscoverySpecCsvOptions(
            header_rows=Primitive.from_proto(resource.header_rows),
            delimiter=Primitive.from_proto(resource.delimiter),
            encoding=Primitive.from_proto(resource.encoding),
            disable_type_inference=Primitive.from_proto(
                resource.disable_type_inference
            ),
        )


class ZoneDiscoverySpecCsvOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ZoneDiscoverySpecCsvOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ZoneDiscoverySpecCsvOptions.from_proto(i) for i in resources]


class ZoneDiscoverySpecJsonOptions(object):
    def __init__(self, encoding: str = None, disable_type_inference: bool = None):
        self.encoding = encoding
        self.disable_type_inference = disable_type_inference

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = zone_pb2.DataplexAlphaZoneDiscoverySpecJsonOptions()
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

        return ZoneDiscoverySpecJsonOptions(
            encoding=Primitive.from_proto(resource.encoding),
            disable_type_inference=Primitive.from_proto(
                resource.disable_type_inference
            ),
        )


class ZoneDiscoverySpecJsonOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ZoneDiscoverySpecJsonOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ZoneDiscoverySpecJsonOptions.from_proto(i) for i in resources]


class ZoneResourceSpec(object):
    def __init__(self, location_type: str = None):
        self.location_type = location_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = zone_pb2.DataplexAlphaZoneResourceSpec()
        if ZoneResourceSpecLocationTypeEnum.to_proto(resource.location_type):
            res.location_type = ZoneResourceSpecLocationTypeEnum.to_proto(
                resource.location_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ZoneResourceSpec(
            location_type=ZoneResourceSpecLocationTypeEnum.from_proto(
                resource.location_type
            ),
        )


class ZoneResourceSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ZoneResourceSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ZoneResourceSpec.from_proto(i) for i in resources]


class ZoneAssetStatus(object):
    def __init__(
        self,
        update_time: str = None,
        active_assets: int = None,
        security_policy_applying_assets: int = None,
    ):
        self.update_time = update_time
        self.active_assets = active_assets
        self.security_policy_applying_assets = security_policy_applying_assets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = zone_pb2.DataplexAlphaZoneAssetStatus()
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        if Primitive.to_proto(resource.active_assets):
            res.active_assets = Primitive.to_proto(resource.active_assets)
        if Primitive.to_proto(resource.security_policy_applying_assets):
            res.security_policy_applying_assets = Primitive.to_proto(
                resource.security_policy_applying_assets
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ZoneAssetStatus(
            update_time=Primitive.from_proto(resource.update_time),
            active_assets=Primitive.from_proto(resource.active_assets),
            security_policy_applying_assets=Primitive.from_proto(
                resource.security_policy_applying_assets
            ),
        )


class ZoneAssetStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ZoneAssetStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ZoneAssetStatus.from_proto(i) for i in resources]


class ZoneStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return zone_pb2.DataplexAlphaZoneStateEnum.Value(
            "DataplexAlphaZoneStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return zone_pb2.DataplexAlphaZoneStateEnum.Name(resource)[
            len("DataplexAlphaZoneStateEnum") :
        ]


class ZoneTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return zone_pb2.DataplexAlphaZoneTypeEnum.Value(
            "DataplexAlphaZoneTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return zone_pb2.DataplexAlphaZoneTypeEnum.Name(resource)[
            len("DataplexAlphaZoneTypeEnum") :
        ]


class ZoneResourceSpecLocationTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return zone_pb2.DataplexAlphaZoneResourceSpecLocationTypeEnum.Value(
            "DataplexAlphaZoneResourceSpecLocationTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return zone_pb2.DataplexAlphaZoneResourceSpecLocationTypeEnum.Name(resource)[
            len("DataplexAlphaZoneResourceSpecLocationTypeEnum") :
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
