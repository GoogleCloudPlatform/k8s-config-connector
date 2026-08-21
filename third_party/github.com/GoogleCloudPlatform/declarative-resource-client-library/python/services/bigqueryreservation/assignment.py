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
from google3.cloud.graphite.mmv2.services.google.bigquery_reservation import (
    assignment_pb2,
)
from google3.cloud.graphite.mmv2.services.google.bigquery_reservation import (
    assignment_pb2_grpc,
)

from typing import List


class Assignment(object):
    def __init__(
        self,
        name: str = None,
        assignee: str = None,
        job_type: str = None,
        state: str = None,
        project: str = None,
        location: str = None,
        reservation: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.assignee = assignee
        self.job_type = job_type
        self.project = project
        self.location = location
        self.reservation = reservation
        self.service_account_file = service_account_file

    def apply(self):
        stub = assignment_pb2_grpc.BigqueryreservationAssignmentServiceStub(
            channel.Channel()
        )
        request = assignment_pb2.ApplyBigqueryreservationAssignmentRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.assignee):
            request.resource.assignee = Primitive.to_proto(self.assignee)

        if AssignmentJobTypeEnum.to_proto(self.job_type):
            request.resource.job_type = AssignmentJobTypeEnum.to_proto(self.job_type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.reservation):
            request.resource.reservation = Primitive.to_proto(self.reservation)

        request.service_account_file = self.service_account_file

        response = stub.ApplyBigqueryreservationAssignment(request)
        self.name = Primitive.from_proto(response.name)
        self.assignee = Primitive.from_proto(response.assignee)
        self.job_type = AssignmentJobTypeEnum.from_proto(response.job_type)
        self.state = AssignmentStateEnum.from_proto(response.state)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.reservation = Primitive.from_proto(response.reservation)

    def delete(self):
        stub = assignment_pb2_grpc.BigqueryreservationAssignmentServiceStub(
            channel.Channel()
        )
        request = assignment_pb2.DeleteBigqueryreservationAssignmentRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.assignee):
            request.resource.assignee = Primitive.to_proto(self.assignee)

        if AssignmentJobTypeEnum.to_proto(self.job_type):
            request.resource.job_type = AssignmentJobTypeEnum.to_proto(self.job_type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.reservation):
            request.resource.reservation = Primitive.to_proto(self.reservation)

        response = stub.DeleteBigqueryreservationAssignment(request)

    @classmethod
    def list(self, project, location, reservation, service_account_file=""):
        stub = assignment_pb2_grpc.BigqueryreservationAssignmentServiceStub(
            channel.Channel()
        )
        request = assignment_pb2.ListBigqueryreservationAssignmentRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Reservation = reservation

        return stub.ListBigqueryreservationAssignment(request).items

    def to_proto(self):
        resource = assignment_pb2.BigqueryreservationAssignment()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.assignee):
            resource.assignee = Primitive.to_proto(self.assignee)
        if AssignmentJobTypeEnum.to_proto(self.job_type):
            resource.job_type = AssignmentJobTypeEnum.to_proto(self.job_type)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.reservation):
            resource.reservation = Primitive.to_proto(self.reservation)
        return resource


class AssignmentJobTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return assignment_pb2.BigqueryreservationAssignmentJobTypeEnum.Value(
            "BigqueryreservationAssignmentJobTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return assignment_pb2.BigqueryreservationAssignmentJobTypeEnum.Name(resource)[
            len("BigqueryreservationAssignmentJobTypeEnum") :
        ]


class AssignmentStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return assignment_pb2.BigqueryreservationAssignmentStateEnum.Value(
            "BigqueryreservationAssignmentStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return assignment_pb2.BigqueryreservationAssignmentStateEnum.Name(resource)[
            len("BigqueryreservationAssignmentStateEnum") :
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
