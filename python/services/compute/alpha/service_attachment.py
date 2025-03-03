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
from google3.cloud.graphite.mmv2.services.google.compute import service_attachment_pb2
from google3.cloud.graphite.mmv2.services.google.compute import (
    service_attachment_pb2_grpc,
)

from typing import List


class ServiceAttachment(object):
    def __init__(
        self,
        id: int = None,
        name: str = None,
        description: str = None,
        self_link: str = None,
        region: str = None,
        target_service: str = None,
        connection_preference: str = None,
        connected_endpoints: list = None,
        nat_subnets: list = None,
        enable_proxy_protocol: bool = None,
        consumer_reject_lists: list = None,
        consumer_accept_lists: list = None,
        psc_service_attachment_id: dict = None,
        fingerprint: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.target_service = target_service
        self.connection_preference = connection_preference
        self.nat_subnets = nat_subnets
        self.enable_proxy_protocol = enable_proxy_protocol
        self.consumer_reject_lists = consumer_reject_lists
        self.consumer_accept_lists = consumer_accept_lists
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = service_attachment_pb2_grpc.ComputeAlphaServiceAttachmentServiceStub(
            channel.Channel()
        )
        request = service_attachment_pb2.ApplyComputeAlphaServiceAttachmentRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.target_service):
            request.resource.target_service = Primitive.to_proto(self.target_service)

        if ServiceAttachmentConnectionPreferenceEnum.to_proto(
            self.connection_preference
        ):
            request.resource.connection_preference = (
                ServiceAttachmentConnectionPreferenceEnum.to_proto(
                    self.connection_preference
                )
            )

        if Primitive.to_proto(self.nat_subnets):
            request.resource.nat_subnets.extend(Primitive.to_proto(self.nat_subnets))
        if Primitive.to_proto(self.enable_proxy_protocol):
            request.resource.enable_proxy_protocol = Primitive.to_proto(
                self.enable_proxy_protocol
            )

        if Primitive.to_proto(self.consumer_reject_lists):
            request.resource.consumer_reject_lists.extend(
                Primitive.to_proto(self.consumer_reject_lists)
            )
        if ServiceAttachmentConsumerAcceptListsArray.to_proto(
            self.consumer_accept_lists
        ):
            request.resource.consumer_accept_lists.extend(
                ServiceAttachmentConsumerAcceptListsArray.to_proto(
                    self.consumer_accept_lists
                )
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeAlphaServiceAttachment(request)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.self_link = Primitive.from_proto(response.self_link)
        self.region = Primitive.from_proto(response.region)
        self.target_service = Primitive.from_proto(response.target_service)
        self.connection_preference = (
            ServiceAttachmentConnectionPreferenceEnum.from_proto(
                response.connection_preference
            )
        )
        self.connected_endpoints = ServiceAttachmentConnectedEndpointsArray.from_proto(
            response.connected_endpoints
        )
        self.nat_subnets = Primitive.from_proto(response.nat_subnets)
        self.enable_proxy_protocol = Primitive.from_proto(
            response.enable_proxy_protocol
        )
        self.consumer_reject_lists = Primitive.from_proto(
            response.consumer_reject_lists
        )
        self.consumer_accept_lists = (
            ServiceAttachmentConsumerAcceptListsArray.from_proto(
                response.consumer_accept_lists
            )
        )
        self.psc_service_attachment_id = (
            ServiceAttachmentPscServiceAttachmentId.from_proto(
                response.psc_service_attachment_id
            )
        )
        self.fingerprint = Primitive.from_proto(response.fingerprint)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = service_attachment_pb2_grpc.ComputeAlphaServiceAttachmentServiceStub(
            channel.Channel()
        )
        request = service_attachment_pb2.DeleteComputeAlphaServiceAttachmentRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.target_service):
            request.resource.target_service = Primitive.to_proto(self.target_service)

        if ServiceAttachmentConnectionPreferenceEnum.to_proto(
            self.connection_preference
        ):
            request.resource.connection_preference = (
                ServiceAttachmentConnectionPreferenceEnum.to_proto(
                    self.connection_preference
                )
            )

        if Primitive.to_proto(self.nat_subnets):
            request.resource.nat_subnets.extend(Primitive.to_proto(self.nat_subnets))
        if Primitive.to_proto(self.enable_proxy_protocol):
            request.resource.enable_proxy_protocol = Primitive.to_proto(
                self.enable_proxy_protocol
            )

        if Primitive.to_proto(self.consumer_reject_lists):
            request.resource.consumer_reject_lists.extend(
                Primitive.to_proto(self.consumer_reject_lists)
            )
        if ServiceAttachmentConsumerAcceptListsArray.to_proto(
            self.consumer_accept_lists
        ):
            request.resource.consumer_accept_lists.extend(
                ServiceAttachmentConsumerAcceptListsArray.to_proto(
                    self.consumer_accept_lists
                )
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteComputeAlphaServiceAttachment(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = service_attachment_pb2_grpc.ComputeAlphaServiceAttachmentServiceStub(
            channel.Channel()
        )
        request = service_attachment_pb2.ListComputeAlphaServiceAttachmentRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeAlphaServiceAttachment(request).items

    def to_proto(self):
        resource = service_attachment_pb2.ComputeAlphaServiceAttachment()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.target_service):
            resource.target_service = Primitive.to_proto(self.target_service)
        if ServiceAttachmentConnectionPreferenceEnum.to_proto(
            self.connection_preference
        ):
            resource.connection_preference = (
                ServiceAttachmentConnectionPreferenceEnum.to_proto(
                    self.connection_preference
                )
            )
        if Primitive.to_proto(self.nat_subnets):
            resource.nat_subnets.extend(Primitive.to_proto(self.nat_subnets))
        if Primitive.to_proto(self.enable_proxy_protocol):
            resource.enable_proxy_protocol = Primitive.to_proto(
                self.enable_proxy_protocol
            )
        if Primitive.to_proto(self.consumer_reject_lists):
            resource.consumer_reject_lists.extend(
                Primitive.to_proto(self.consumer_reject_lists)
            )
        if ServiceAttachmentConsumerAcceptListsArray.to_proto(
            self.consumer_accept_lists
        ):
            resource.consumer_accept_lists.extend(
                ServiceAttachmentConsumerAcceptListsArray.to_proto(
                    self.consumer_accept_lists
                )
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ServiceAttachmentConnectedEndpoints(object):
    def __init__(
        self, status: str = None, psc_connection_id: int = None, endpoint: str = None
    ):
        self.status = status
        self.psc_connection_id = psc_connection_id
        self.endpoint = endpoint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_attachment_pb2.ComputeAlphaServiceAttachmentConnectedEndpoints()
        if ServiceAttachmentConnectedEndpointsStatusEnum.to_proto(resource.status):
            res.status = ServiceAttachmentConnectedEndpointsStatusEnum.to_proto(
                resource.status
            )
        if Primitive.to_proto(resource.psc_connection_id):
            res.psc_connection_id = Primitive.to_proto(resource.psc_connection_id)
        if Primitive.to_proto(resource.endpoint):
            res.endpoint = Primitive.to_proto(resource.endpoint)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceAttachmentConnectedEndpoints(
            status=ServiceAttachmentConnectedEndpointsStatusEnum.from_proto(
                resource.status
            ),
            psc_connection_id=Primitive.from_proto(resource.psc_connection_id),
            endpoint=Primitive.from_proto(resource.endpoint),
        )


class ServiceAttachmentConnectedEndpointsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceAttachmentConnectedEndpoints.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceAttachmentConnectedEndpoints.from_proto(i) for i in resources]


class ServiceAttachmentConsumerAcceptLists(object):
    def __init__(self, project_id_or_num: str = None, connection_limit: int = None):
        self.project_id_or_num = project_id_or_num
        self.connection_limit = connection_limit

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_attachment_pb2.ComputeAlphaServiceAttachmentConsumerAcceptLists()
        if Primitive.to_proto(resource.project_id_or_num):
            res.project_id_or_num = Primitive.to_proto(resource.project_id_or_num)
        if Primitive.to_proto(resource.connection_limit):
            res.connection_limit = Primitive.to_proto(resource.connection_limit)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceAttachmentConsumerAcceptLists(
            project_id_or_num=Primitive.from_proto(resource.project_id_or_num),
            connection_limit=Primitive.from_proto(resource.connection_limit),
        )


class ServiceAttachmentConsumerAcceptListsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceAttachmentConsumerAcceptLists.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceAttachmentConsumerAcceptLists.from_proto(i) for i in resources]


class ServiceAttachmentPscServiceAttachmentId(object):
    def __init__(self, high: int = None, low: int = None):
        self.high = high
        self.low = low

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_attachment_pb2.ComputeAlphaServiceAttachmentPscServiceAttachmentId()
        )
        if Primitive.to_proto(resource.high):
            res.high = Primitive.to_proto(resource.high)
        if Primitive.to_proto(resource.low):
            res.low = Primitive.to_proto(resource.low)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceAttachmentPscServiceAttachmentId(
            high=Primitive.from_proto(resource.high),
            low=Primitive.from_proto(resource.low),
        )


class ServiceAttachmentPscServiceAttachmentIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceAttachmentPscServiceAttachmentId.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceAttachmentPscServiceAttachmentId.from_proto(i) for i in resources
        ]


class ServiceAttachmentConnectionPreferenceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_attachment_pb2.ComputeAlphaServiceAttachmentConnectionPreferenceEnum.Value(
            "ComputeAlphaServiceAttachmentConnectionPreferenceEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_attachment_pb2.ComputeAlphaServiceAttachmentConnectionPreferenceEnum.Name(
            resource
        )[
            len("ComputeAlphaServiceAttachmentConnectionPreferenceEnum") :
        ]


class ServiceAttachmentConnectedEndpointsStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_attachment_pb2.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum.Value(
            "ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_attachment_pb2.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum.Name(
            resource
        )[
            len("ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum") :
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
