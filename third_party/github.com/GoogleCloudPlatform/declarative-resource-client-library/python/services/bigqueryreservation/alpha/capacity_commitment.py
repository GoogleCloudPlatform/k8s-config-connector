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
from google3.cloud.graphite.mmv2.services.google.bigquery_reservation import (
    capacity_commitment_pb2,
)
from google3.cloud.graphite.mmv2.services.google.bigquery_reservation import (
    capacity_commitment_pb2_grpc,
)

from typing import List


class CapacityCommitment(object):
    def __init__(
        self,
        name: str = None,
        slot_count: int = None,
        plan: str = None,
        state: str = None,
        commitment_start_time: str = None,
        commitment_end_time: str = None,
        failure_status: dict = None,
        renewal_plan: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.slot_count = slot_count
        self.plan = plan
        self.renewal_plan = renewal_plan
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = capacity_commitment_pb2_grpc.BigqueryreservationAlphaCapacityCommitmentServiceStub(
            channel.Channel()
        )
        request = (
            capacity_commitment_pb2.ApplyBigqueryreservationAlphaCapacityCommitmentRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.slot_count):
            request.resource.slot_count = Primitive.to_proto(self.slot_count)

        if CapacityCommitmentPlanEnum.to_proto(self.plan):
            request.resource.plan = CapacityCommitmentPlanEnum.to_proto(self.plan)

        if CapacityCommitmentRenewalPlanEnum.to_proto(self.renewal_plan):
            request.resource.renewal_plan = CapacityCommitmentRenewalPlanEnum.to_proto(
                self.renewal_plan
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyBigqueryreservationAlphaCapacityCommitment(request)
        self.name = Primitive.from_proto(response.name)
        self.slot_count = Primitive.from_proto(response.slot_count)
        self.plan = CapacityCommitmentPlanEnum.from_proto(response.plan)
        self.state = CapacityCommitmentStateEnum.from_proto(response.state)
        self.commitment_start_time = Primitive.from_proto(
            response.commitment_start_time
        )
        self.commitment_end_time = Primitive.from_proto(response.commitment_end_time)
        self.failure_status = CapacityCommitmentFailureStatus.from_proto(
            response.failure_status
        )
        self.renewal_plan = CapacityCommitmentRenewalPlanEnum.from_proto(
            response.renewal_plan
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = capacity_commitment_pb2_grpc.BigqueryreservationAlphaCapacityCommitmentServiceStub(
            channel.Channel()
        )
        request = (
            capacity_commitment_pb2.DeleteBigqueryreservationAlphaCapacityCommitmentRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.slot_count):
            request.resource.slot_count = Primitive.to_proto(self.slot_count)

        if CapacityCommitmentPlanEnum.to_proto(self.plan):
            request.resource.plan = CapacityCommitmentPlanEnum.to_proto(self.plan)

        if CapacityCommitmentRenewalPlanEnum.to_proto(self.renewal_plan):
            request.resource.renewal_plan = CapacityCommitmentRenewalPlanEnum.to_proto(
                self.renewal_plan
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteBigqueryreservationAlphaCapacityCommitment(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = capacity_commitment_pb2_grpc.BigqueryreservationAlphaCapacityCommitmentServiceStub(
            channel.Channel()
        )
        request = (
            capacity_commitment_pb2.ListBigqueryreservationAlphaCapacityCommitmentRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListBigqueryreservationAlphaCapacityCommitment(request).items

    def to_proto(self):
        resource = capacity_commitment_pb2.BigqueryreservationAlphaCapacityCommitment()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.slot_count):
            resource.slot_count = Primitive.to_proto(self.slot_count)
        if CapacityCommitmentPlanEnum.to_proto(self.plan):
            resource.plan = CapacityCommitmentPlanEnum.to_proto(self.plan)
        if CapacityCommitmentRenewalPlanEnum.to_proto(self.renewal_plan):
            resource.renewal_plan = CapacityCommitmentRenewalPlanEnum.to_proto(
                self.renewal_plan
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class CapacityCommitmentFailureStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            capacity_commitment_pb2.BigqueryreservationAlphaCapacityCommitmentFailureStatus()
        )
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if CapacityCommitmentFailureStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                CapacityCommitmentFailureStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CapacityCommitmentFailureStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=CapacityCommitmentFailureStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class CapacityCommitmentFailureStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CapacityCommitmentFailureStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CapacityCommitmentFailureStatus.from_proto(i) for i in resources]


class CapacityCommitmentFailureStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            capacity_commitment_pb2.BigqueryreservationAlphaCapacityCommitmentFailureStatusDetails()
        )
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CapacityCommitmentFailureStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class CapacityCommitmentFailureStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CapacityCommitmentFailureStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CapacityCommitmentFailureStatusDetails.from_proto(i) for i in resources]


class CapacityCommitmentPlanEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return capacity_commitment_pb2.BigqueryreservationAlphaCapacityCommitmentPlanEnum.Value(
            "BigqueryreservationAlphaCapacityCommitmentPlanEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return capacity_commitment_pb2.BigqueryreservationAlphaCapacityCommitmentPlanEnum.Name(
            resource
        )[
            len("BigqueryreservationAlphaCapacityCommitmentPlanEnum") :
        ]


class CapacityCommitmentStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return capacity_commitment_pb2.BigqueryreservationAlphaCapacityCommitmentStateEnum.Value(
            "BigqueryreservationAlphaCapacityCommitmentStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return capacity_commitment_pb2.BigqueryreservationAlphaCapacityCommitmentStateEnum.Name(
            resource
        )[
            len("BigqueryreservationAlphaCapacityCommitmentStateEnum") :
        ]


class CapacityCommitmentRenewalPlanEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return capacity_commitment_pb2.BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum.Value(
            "BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return capacity_commitment_pb2.BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum.Name(
            resource
        )[
            len("BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum") :
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
