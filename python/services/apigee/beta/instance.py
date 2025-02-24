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
from google3.cloud.graphite.mmv2.services.google.apigee import instance_pb2
from google3.cloud.graphite.mmv2.services.google.apigee import instance_pb2_grpc

from typing import List


class Instance(object):
    def __init__(
        self,
        name: str = None,
        location: str = None,
        peering_cidr_range: str = None,
        host: str = None,
        port: str = None,
        description: str = None,
        display_name: str = None,
        created_at: int = None,
        last_modified_at: int = None,
        disk_encryption_key_name: str = None,
        state: str = None,
        apigee_organization: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.location = location
        self.peering_cidr_range = peering_cidr_range
        self.description = description
        self.display_name = display_name
        self.disk_encryption_key_name = disk_encryption_key_name
        self.apigee_organization = apigee_organization
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_pb2_grpc.ApigeeBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ApplyApigeeBetaInstanceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if InstancePeeringCidrRangeEnum.to_proto(self.peering_cidr_range):
            request.resource.peering_cidr_range = InstancePeeringCidrRangeEnum.to_proto(
                self.peering_cidr_range
            )

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.disk_encryption_key_name):
            request.resource.disk_encryption_key_name = Primitive.to_proto(
                self.disk_encryption_key_name
            )

        if Primitive.to_proto(self.apigee_organization):
            request.resource.apigee_organization = Primitive.to_proto(
                self.apigee_organization
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyApigeeBetaInstance(request)
        self.name = Primitive.from_proto(response.name)
        self.location = Primitive.from_proto(response.location)
        self.peering_cidr_range = InstancePeeringCidrRangeEnum.from_proto(
            response.peering_cidr_range
        )
        self.host = Primitive.from_proto(response.host)
        self.port = Primitive.from_proto(response.port)
        self.description = Primitive.from_proto(response.description)
        self.display_name = Primitive.from_proto(response.display_name)
        self.created_at = Primitive.from_proto(response.created_at)
        self.last_modified_at = Primitive.from_proto(response.last_modified_at)
        self.disk_encryption_key_name = Primitive.from_proto(
            response.disk_encryption_key_name
        )
        self.state = InstanceStateEnum.from_proto(response.state)
        self.apigee_organization = Primitive.from_proto(response.apigee_organization)

    def delete(self):
        stub = instance_pb2_grpc.ApigeeBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.DeleteApigeeBetaInstanceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if InstancePeeringCidrRangeEnum.to_proto(self.peering_cidr_range):
            request.resource.peering_cidr_range = InstancePeeringCidrRangeEnum.to_proto(
                self.peering_cidr_range
            )

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.disk_encryption_key_name):
            request.resource.disk_encryption_key_name = Primitive.to_proto(
                self.disk_encryption_key_name
            )

        if Primitive.to_proto(self.apigee_organization):
            request.resource.apigee_organization = Primitive.to_proto(
                self.apigee_organization
            )

        response = stub.DeleteApigeeBetaInstance(request)

    @classmethod
    def list(self, apigeeOrganization, service_account_file=""):
        stub = instance_pb2_grpc.ApigeeBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ListApigeeBetaInstanceRequest()
        request.service_account_file = service_account_file
        request.ApigeeOrganization = apigeeOrganization

        return stub.ListApigeeBetaInstance(request).items

    def to_proto(self):
        resource = instance_pb2.ApigeeBetaInstance()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if InstancePeeringCidrRangeEnum.to_proto(self.peering_cidr_range):
            resource.peering_cidr_range = InstancePeeringCidrRangeEnum.to_proto(
                self.peering_cidr_range
            )
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.disk_encryption_key_name):
            resource.disk_encryption_key_name = Primitive.to_proto(
                self.disk_encryption_key_name
            )
        if Primitive.to_proto(self.apigee_organization):
            resource.apigee_organization = Primitive.to_proto(self.apigee_organization)
        return resource


class InstancePeeringCidrRangeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ApigeeBetaInstancePeeringCidrRangeEnum.Value(
            "ApigeeBetaInstancePeeringCidrRangeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ApigeeBetaInstancePeeringCidrRangeEnum.Name(resource)[
            len("ApigeeBetaInstancePeeringCidrRangeEnum") :
        ]


class InstanceStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ApigeeBetaInstanceStateEnum.Value(
            "ApigeeBetaInstanceStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.ApigeeBetaInstanceStateEnum.Name(resource)[
            len("ApigeeBetaInstanceStateEnum") :
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
