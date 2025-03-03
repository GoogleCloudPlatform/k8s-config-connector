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
    reservation_pb2,
)
from google3.cloud.graphite.mmv2.services.google.bigquery_reservation import (
    reservation_pb2_grpc,
)

from typing import List


class Reservation(object):
    def __init__(
        self,
        name: str = None,
        slot_capacity: int = None,
        ignore_idle_slots: bool = None,
        creation_time: str = None,
        update_time: str = None,
        max_concurrency: int = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.slot_capacity = slot_capacity
        self.ignore_idle_slots = ignore_idle_slots
        self.max_concurrency = max_concurrency
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = reservation_pb2_grpc.BigqueryreservationBetaReservationServiceStub(
            channel.Channel()
        )
        request = reservation_pb2.ApplyBigqueryreservationBetaReservationRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.slot_capacity):
            request.resource.slot_capacity = Primitive.to_proto(self.slot_capacity)

        if Primitive.to_proto(self.ignore_idle_slots):
            request.resource.ignore_idle_slots = Primitive.to_proto(
                self.ignore_idle_slots
            )

        if Primitive.to_proto(self.max_concurrency):
            request.resource.max_concurrency = Primitive.to_proto(self.max_concurrency)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyBigqueryreservationBetaReservation(request)
        self.name = Primitive.from_proto(response.name)
        self.slot_capacity = Primitive.from_proto(response.slot_capacity)
        self.ignore_idle_slots = Primitive.from_proto(response.ignore_idle_slots)
        self.creation_time = Primitive.from_proto(response.creation_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.max_concurrency = Primitive.from_proto(response.max_concurrency)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = reservation_pb2_grpc.BigqueryreservationBetaReservationServiceStub(
            channel.Channel()
        )
        request = reservation_pb2.DeleteBigqueryreservationBetaReservationRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.slot_capacity):
            request.resource.slot_capacity = Primitive.to_proto(self.slot_capacity)

        if Primitive.to_proto(self.ignore_idle_slots):
            request.resource.ignore_idle_slots = Primitive.to_proto(
                self.ignore_idle_slots
            )

        if Primitive.to_proto(self.max_concurrency):
            request.resource.max_concurrency = Primitive.to_proto(self.max_concurrency)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteBigqueryreservationBetaReservation(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = reservation_pb2_grpc.BigqueryreservationBetaReservationServiceStub(
            channel.Channel()
        )
        request = reservation_pb2.ListBigqueryreservationBetaReservationRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListBigqueryreservationBetaReservation(request).items

    def to_proto(self):
        resource = reservation_pb2.BigqueryreservationBetaReservation()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.slot_capacity):
            resource.slot_capacity = Primitive.to_proto(self.slot_capacity)
        if Primitive.to_proto(self.ignore_idle_slots):
            resource.ignore_idle_slots = Primitive.to_proto(self.ignore_idle_slots)
        if Primitive.to_proto(self.max_concurrency):
            resource.max_concurrency = Primitive.to_proto(self.max_concurrency)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
