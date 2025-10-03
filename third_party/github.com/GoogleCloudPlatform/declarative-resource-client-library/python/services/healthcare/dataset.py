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
from google3.cloud.graphite.mmv2.services.google.healthcare import dataset_pb2
from google3.cloud.graphite.mmv2.services.google.healthcare import dataset_pb2_grpc

from typing import List


class Dataset(object):
    def __init__(
        self,
        name: str = None,
        time_zone: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.time_zone = time_zone
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = dataset_pb2_grpc.HealthcareDatasetServiceStub(channel.Channel())
        request = dataset_pb2.ApplyHealthcareDatasetRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.time_zone):
            request.resource.time_zone = Primitive.to_proto(self.time_zone)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyHealthcareDataset(request)
        self.name = Primitive.from_proto(response.name)
        self.time_zone = Primitive.from_proto(response.time_zone)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = dataset_pb2_grpc.HealthcareDatasetServiceStub(channel.Channel())
        request = dataset_pb2.DeleteHealthcareDatasetRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.time_zone):
            request.resource.time_zone = Primitive.to_proto(self.time_zone)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteHealthcareDataset(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = dataset_pb2_grpc.HealthcareDatasetServiceStub(channel.Channel())
        request = dataset_pb2.ListHealthcareDatasetRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListHealthcareDataset(request).items

    def to_proto(self):
        resource = dataset_pb2.HealthcareDataset()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.time_zone):
            resource.time_zone = Primitive.to_proto(self.time_zone)
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
