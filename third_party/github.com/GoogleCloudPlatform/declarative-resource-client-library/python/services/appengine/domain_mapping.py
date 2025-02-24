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
from google3.cloud.graphite.mmv2.services.google.app_engine import domain_mapping_pb2
from google3.cloud.graphite.mmv2.services.google.app_engine import (
    domain_mapping_pb2_grpc,
)

from typing import List


class DomainMapping(object):
    def __init__(
        self,
        self_link: str = None,
        name: str = None,
        ssl_settings: dict = None,
        resource_records: list = None,
        app: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.self_link = self_link
        self.name = name
        self.ssl_settings = ssl_settings
        self.app = app
        self.service_account_file = service_account_file

    def apply(self):
        stub = domain_mapping_pb2_grpc.AppengineDomainMappingServiceStub(
            channel.Channel()
        )
        request = domain_mapping_pb2.ApplyAppengineDomainMappingRequest()
        if Primitive.to_proto(self.self_link):
            request.resource.self_link = Primitive.to_proto(self.self_link)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if DomainMappingSslSettings.to_proto(self.ssl_settings):
            request.resource.ssl_settings.CopyFrom(
                DomainMappingSslSettings.to_proto(self.ssl_settings)
            )
        else:
            request.resource.ClearField("ssl_settings")
        if Primitive.to_proto(self.app):
            request.resource.app = Primitive.to_proto(self.app)

        request.service_account_file = self.service_account_file

        response = stub.ApplyAppengineDomainMapping(request)
        self.self_link = Primitive.from_proto(response.self_link)
        self.name = Primitive.from_proto(response.name)
        self.ssl_settings = DomainMappingSslSettings.from_proto(response.ssl_settings)
        self.resource_records = DomainMappingResourceRecordsArray.from_proto(
            response.resource_records
        )
        self.app = Primitive.from_proto(response.app)

    def delete(self):
        stub = domain_mapping_pb2_grpc.AppengineDomainMappingServiceStub(
            channel.Channel()
        )
        request = domain_mapping_pb2.DeleteAppengineDomainMappingRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.self_link):
            request.resource.self_link = Primitive.to_proto(self.self_link)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if DomainMappingSslSettings.to_proto(self.ssl_settings):
            request.resource.ssl_settings.CopyFrom(
                DomainMappingSslSettings.to_proto(self.ssl_settings)
            )
        else:
            request.resource.ClearField("ssl_settings")
        if Primitive.to_proto(self.app):
            request.resource.app = Primitive.to_proto(self.app)

        response = stub.DeleteAppengineDomainMapping(request)

    @classmethod
    def list(self, app, service_account_file=""):
        stub = domain_mapping_pb2_grpc.AppengineDomainMappingServiceStub(
            channel.Channel()
        )
        request = domain_mapping_pb2.ListAppengineDomainMappingRequest()
        request.service_account_file = service_account_file
        request.App = app

        return stub.ListAppengineDomainMapping(request).items

    def to_proto(self):
        resource = domain_mapping_pb2.AppengineDomainMapping()
        if Primitive.to_proto(self.self_link):
            resource.self_link = Primitive.to_proto(self.self_link)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if DomainMappingSslSettings.to_proto(self.ssl_settings):
            resource.ssl_settings.CopyFrom(
                DomainMappingSslSettings.to_proto(self.ssl_settings)
            )
        else:
            resource.ClearField("ssl_settings")
        if Primitive.to_proto(self.app):
            resource.app = Primitive.to_proto(self.app)
        return resource


class DomainMappingSslSettings(object):
    def __init__(
        self,
        certificate_id: str = None,
        ssl_management_type: str = None,
        pending_managed_certificate_id: str = None,
    ):
        self.certificate_id = certificate_id
        self.ssl_management_type = ssl_management_type
        self.pending_managed_certificate_id = pending_managed_certificate_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = domain_mapping_pb2.AppengineDomainMappingSslSettings()
        if Primitive.to_proto(resource.certificate_id):
            res.certificate_id = Primitive.to_proto(resource.certificate_id)
        if DomainMappingSslSettingsSslManagementTypeEnum.to_proto(
            resource.ssl_management_type
        ):
            res.ssl_management_type = DomainMappingSslSettingsSslManagementTypeEnum.to_proto(
                resource.ssl_management_type
            )
        if Primitive.to_proto(resource.pending_managed_certificate_id):
            res.pending_managed_certificate_id = Primitive.to_proto(
                resource.pending_managed_certificate_id
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DomainMappingSslSettings(
            certificate_id=Primitive.from_proto(resource.certificate_id),
            ssl_management_type=DomainMappingSslSettingsSslManagementTypeEnum.from_proto(
                resource.ssl_management_type
            ),
            pending_managed_certificate_id=Primitive.from_proto(
                resource.pending_managed_certificate_id
            ),
        )


class DomainMappingSslSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DomainMappingSslSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DomainMappingSslSettings.from_proto(i) for i in resources]


class DomainMappingResourceRecords(object):
    def __init__(self, name: str = None, rrdata: str = None, type: str = None):
        self.name = name
        self.rrdata = rrdata
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = domain_mapping_pb2.AppengineDomainMappingResourceRecords()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.rrdata):
            res.rrdata = Primitive.to_proto(resource.rrdata)
        if DomainMappingResourceRecordsTypeEnum.to_proto(resource.type):
            res.type = DomainMappingResourceRecordsTypeEnum.to_proto(resource.type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DomainMappingResourceRecords(
            name=Primitive.from_proto(resource.name),
            rrdata=Primitive.from_proto(resource.rrdata),
            type=DomainMappingResourceRecordsTypeEnum.from_proto(resource.type),
        )


class DomainMappingResourceRecordsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DomainMappingResourceRecords.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DomainMappingResourceRecords.from_proto(i) for i in resources]


class DomainMappingSslSettingsSslManagementTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return domain_mapping_pb2.AppengineDomainMappingSslSettingsSslManagementTypeEnum.Value(
            "AppengineDomainMappingSslSettingsSslManagementTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return domain_mapping_pb2.AppengineDomainMappingSslSettingsSslManagementTypeEnum.Name(
            resource
        )[
            len("AppengineDomainMappingSslSettingsSslManagementTypeEnum") :
        ]


class DomainMappingResourceRecordsTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return domain_mapping_pb2.AppengineDomainMappingResourceRecordsTypeEnum.Value(
            "AppengineDomainMappingResourceRecordsTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return domain_mapping_pb2.AppengineDomainMappingResourceRecordsTypeEnum.Name(
            resource
        )[len("AppengineDomainMappingResourceRecordsTypeEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
