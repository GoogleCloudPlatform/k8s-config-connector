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
from google3.cloud.graphite.mmv2.services.google.compute import interconnect_pb2
from google3.cloud.graphite.mmv2.services.google.compute import interconnect_pb2_grpc

from typing import List


class Interconnect(object):
    def __init__(
        self,
        description: str = None,
        self_link: str = None,
        id: int = None,
        name: str = None,
        location: str = None,
        link_type: str = None,
        requested_link_count: int = None,
        interconnect_type: str = None,
        admin_enabled: bool = None,
        noc_contact_email: str = None,
        customer_name: str = None,
        operational_status: str = None,
        provisioned_link_count: int = None,
        interconnect_attachments: list = None,
        peer_ip_address: str = None,
        google_ip_address: str = None,
        google_reference_id: str = None,
        expected_outages: list = None,
        circuit_infos: list = None,
        state: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.description = description
        self.name = name
        self.location = location
        self.link_type = link_type
        self.requested_link_count = requested_link_count
        self.interconnect_type = interconnect_type
        self.admin_enabled = admin_enabled
        self.noc_contact_email = noc_contact_email
        self.customer_name = customer_name
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = interconnect_pb2_grpc.ComputeInterconnectServiceStub(channel.Channel())
        request = interconnect_pb2.ApplyComputeInterconnectRequest()
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if InterconnectLinkTypeEnum.to_proto(self.link_type):
            request.resource.link_type = InterconnectLinkTypeEnum.to_proto(
                self.link_type
            )

        if Primitive.to_proto(self.requested_link_count):
            request.resource.requested_link_count = Primitive.to_proto(
                self.requested_link_count
            )

        if InterconnectInterconnectTypeEnum.to_proto(self.interconnect_type):
            request.resource.interconnect_type = InterconnectInterconnectTypeEnum.to_proto(
                self.interconnect_type
            )

        if Primitive.to_proto(self.admin_enabled):
            request.resource.admin_enabled = Primitive.to_proto(self.admin_enabled)

        if Primitive.to_proto(self.noc_contact_email):
            request.resource.noc_contact_email = Primitive.to_proto(
                self.noc_contact_email
            )

        if Primitive.to_proto(self.customer_name):
            request.resource.customer_name = Primitive.to_proto(self.customer_name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeInterconnect(request)
        self.description = Primitive.from_proto(response.description)
        self.self_link = Primitive.from_proto(response.self_link)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.location = Primitive.from_proto(response.location)
        self.link_type = InterconnectLinkTypeEnum.from_proto(response.link_type)
        self.requested_link_count = Primitive.from_proto(response.requested_link_count)
        self.interconnect_type = InterconnectInterconnectTypeEnum.from_proto(
            response.interconnect_type
        )
        self.admin_enabled = Primitive.from_proto(response.admin_enabled)
        self.noc_contact_email = Primitive.from_proto(response.noc_contact_email)
        self.customer_name = Primitive.from_proto(response.customer_name)
        self.operational_status = InterconnectOperationalStatusEnum.from_proto(
            response.operational_status
        )
        self.provisioned_link_count = Primitive.from_proto(
            response.provisioned_link_count
        )
        self.interconnect_attachments = Primitive.from_proto(
            response.interconnect_attachments
        )
        self.peer_ip_address = Primitive.from_proto(response.peer_ip_address)
        self.google_ip_address = Primitive.from_proto(response.google_ip_address)
        self.google_reference_id = Primitive.from_proto(response.google_reference_id)
        self.expected_outages = InterconnectExpectedOutagesArray.from_proto(
            response.expected_outages
        )
        self.circuit_infos = InterconnectCircuitInfosArray.from_proto(
            response.circuit_infos
        )
        self.state = InterconnectStateEnum.from_proto(response.state)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = interconnect_pb2_grpc.ComputeInterconnectServiceStub(channel.Channel())
        request = interconnect_pb2.DeleteComputeInterconnectRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if InterconnectLinkTypeEnum.to_proto(self.link_type):
            request.resource.link_type = InterconnectLinkTypeEnum.to_proto(
                self.link_type
            )

        if Primitive.to_proto(self.requested_link_count):
            request.resource.requested_link_count = Primitive.to_proto(
                self.requested_link_count
            )

        if InterconnectInterconnectTypeEnum.to_proto(self.interconnect_type):
            request.resource.interconnect_type = InterconnectInterconnectTypeEnum.to_proto(
                self.interconnect_type
            )

        if Primitive.to_proto(self.admin_enabled):
            request.resource.admin_enabled = Primitive.to_proto(self.admin_enabled)

        if Primitive.to_proto(self.noc_contact_email):
            request.resource.noc_contact_email = Primitive.to_proto(
                self.noc_contact_email
            )

        if Primitive.to_proto(self.customer_name):
            request.resource.customer_name = Primitive.to_proto(self.customer_name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeInterconnect(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = interconnect_pb2_grpc.ComputeInterconnectServiceStub(channel.Channel())
        request = interconnect_pb2.ListComputeInterconnectRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeInterconnect(request).items

    def to_proto(self):
        resource = interconnect_pb2.ComputeInterconnect()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if InterconnectLinkTypeEnum.to_proto(self.link_type):
            resource.link_type = InterconnectLinkTypeEnum.to_proto(self.link_type)
        if Primitive.to_proto(self.requested_link_count):
            resource.requested_link_count = Primitive.to_proto(
                self.requested_link_count
            )
        if InterconnectInterconnectTypeEnum.to_proto(self.interconnect_type):
            resource.interconnect_type = InterconnectInterconnectTypeEnum.to_proto(
                self.interconnect_type
            )
        if Primitive.to_proto(self.admin_enabled):
            resource.admin_enabled = Primitive.to_proto(self.admin_enabled)
        if Primitive.to_proto(self.noc_contact_email):
            resource.noc_contact_email = Primitive.to_proto(self.noc_contact_email)
        if Primitive.to_proto(self.customer_name):
            resource.customer_name = Primitive.to_proto(self.customer_name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class InterconnectExpectedOutages(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        source: str = None,
        state: str = None,
        issue_type: str = None,
        affected_circuits: list = None,
        start_time: int = None,
        end_time: int = None,
    ):
        self.name = name
        self.description = description
        self.source = source
        self.state = state
        self.issue_type = issue_type
        self.affected_circuits = affected_circuits
        self.start_time = start_time
        self.end_time = end_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = interconnect_pb2.ComputeInterconnectExpectedOutages()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if InterconnectExpectedOutagesSourceEnum.to_proto(resource.source):
            res.source = InterconnectExpectedOutagesSourceEnum.to_proto(resource.source)
        if InterconnectExpectedOutagesStateEnum.to_proto(resource.state):
            res.state = InterconnectExpectedOutagesStateEnum.to_proto(resource.state)
        if InterconnectExpectedOutagesIssueTypeEnum.to_proto(resource.issue_type):
            res.issue_type = InterconnectExpectedOutagesIssueTypeEnum.to_proto(
                resource.issue_type
            )
        if Primitive.to_proto(resource.affected_circuits):
            res.affected_circuits.extend(Primitive.to_proto(resource.affected_circuits))
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.end_time):
            res.end_time = Primitive.to_proto(resource.end_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InterconnectExpectedOutages(
            name=Primitive.from_proto(resource.name),
            description=Primitive.from_proto(resource.description),
            source=InterconnectExpectedOutagesSourceEnum.from_proto(resource.source),
            state=InterconnectExpectedOutagesStateEnum.from_proto(resource.state),
            issue_type=InterconnectExpectedOutagesIssueTypeEnum.from_proto(
                resource.issue_type
            ),
            affected_circuits=Primitive.from_proto(resource.affected_circuits),
            start_time=Primitive.from_proto(resource.start_time),
            end_time=Primitive.from_proto(resource.end_time),
        )


class InterconnectExpectedOutagesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InterconnectExpectedOutages.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InterconnectExpectedOutages.from_proto(i) for i in resources]


class InterconnectCircuitInfos(object):
    def __init__(
        self,
        google_circuit_id: str = None,
        google_demarc_id: str = None,
        customer_demarc_id: str = None,
    ):
        self.google_circuit_id = google_circuit_id
        self.google_demarc_id = google_demarc_id
        self.customer_demarc_id = customer_demarc_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = interconnect_pb2.ComputeInterconnectCircuitInfos()
        if Primitive.to_proto(resource.google_circuit_id):
            res.google_circuit_id = Primitive.to_proto(resource.google_circuit_id)
        if Primitive.to_proto(resource.google_demarc_id):
            res.google_demarc_id = Primitive.to_proto(resource.google_demarc_id)
        if Primitive.to_proto(resource.customer_demarc_id):
            res.customer_demarc_id = Primitive.to_proto(resource.customer_demarc_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InterconnectCircuitInfos(
            google_circuit_id=Primitive.from_proto(resource.google_circuit_id),
            google_demarc_id=Primitive.from_proto(resource.google_demarc_id),
            customer_demarc_id=Primitive.from_proto(resource.customer_demarc_id),
        )


class InterconnectCircuitInfosArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InterconnectCircuitInfos.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InterconnectCircuitInfos.from_proto(i) for i in resources]


class InterconnectLinkTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectLinkTypeEnum.Value(
            "ComputeInterconnectLinkTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectLinkTypeEnum.Name(resource)[
            len("ComputeInterconnectLinkTypeEnum") :
        ]


class InterconnectInterconnectTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectInterconnectTypeEnum.Value(
            "ComputeInterconnectInterconnectTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectInterconnectTypeEnum.Name(resource)[
            len("ComputeInterconnectInterconnectTypeEnum") :
        ]


class InterconnectOperationalStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectOperationalStatusEnum.Value(
            "ComputeInterconnectOperationalStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectOperationalStatusEnum.Name(resource)[
            len("ComputeInterconnectOperationalStatusEnum") :
        ]


class InterconnectExpectedOutagesSourceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectExpectedOutagesSourceEnum.Value(
            "ComputeInterconnectExpectedOutagesSourceEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectExpectedOutagesSourceEnum.Name(
            resource
        )[len("ComputeInterconnectExpectedOutagesSourceEnum") :]


class InterconnectExpectedOutagesStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectExpectedOutagesStateEnum.Value(
            "ComputeInterconnectExpectedOutagesStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectExpectedOutagesStateEnum.Name(
            resource
        )[len("ComputeInterconnectExpectedOutagesStateEnum") :]


class InterconnectExpectedOutagesIssueTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectExpectedOutagesIssueTypeEnum.Value(
            "ComputeInterconnectExpectedOutagesIssueTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectExpectedOutagesIssueTypeEnum.Name(
            resource
        )[len("ComputeInterconnectExpectedOutagesIssueTypeEnum") :]


class InterconnectStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectStateEnum.Value(
            "ComputeInterconnectStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_pb2.ComputeInterconnectStateEnum.Name(resource)[
            len("ComputeInterconnectStateEnum") :
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
