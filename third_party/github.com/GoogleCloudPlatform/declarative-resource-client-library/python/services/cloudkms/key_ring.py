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
from google3.cloud.graphite.mmv2.services.google.cloudkms import key_ring_pb2
from google3.cloud.graphite.mmv2.services.google.cloudkms import key_ring_pb2_grpc

from typing import List


class KeyRing(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = key_ring_pb2_grpc.CloudkmsKeyRingServiceStub(channel.Channel())
        request = key_ring_pb2.ApplyCloudkmsKeyRingRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudkmsKeyRing(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = key_ring_pb2_grpc.CloudkmsKeyRingServiceStub(channel.Channel())
        request = key_ring_pb2.DeleteCloudkmsKeyRingRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteCloudkmsKeyRing(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = key_ring_pb2_grpc.CloudkmsKeyRingServiceStub(channel.Channel())
        request = key_ring_pb2.ListCloudkmsKeyRingRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListCloudkmsKeyRing(request).items

    def to_proto(self):
        resource = key_ring_pb2.CloudkmsKeyRing()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
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
