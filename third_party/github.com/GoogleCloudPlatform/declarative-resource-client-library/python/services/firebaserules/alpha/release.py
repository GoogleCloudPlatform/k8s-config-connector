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
from google3.cloud.graphite.mmv2.services.google.firebaserules import release_pb2
from google3.cloud.graphite.mmv2.services.google.firebaserules import release_pb2_grpc

from typing import List


class Release(object):
    def __init__(
        self,
        name: str = None,
        ruleset_name: str = None,
        create_time: str = None,
        update_time: str = None,
        disabled: bool = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.ruleset_name = ruleset_name
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = release_pb2_grpc.FirebaserulesAlphaReleaseServiceStub(channel.Channel())
        request = release_pb2.ApplyFirebaserulesAlphaReleaseRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.ruleset_name):
            request.resource.ruleset_name = Primitive.to_proto(self.ruleset_name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyFirebaserulesAlphaRelease(request)
        self.name = Primitive.from_proto(response.name)
        self.ruleset_name = Primitive.from_proto(response.ruleset_name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.disabled = Primitive.from_proto(response.disabled)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = release_pb2_grpc.FirebaserulesAlphaReleaseServiceStub(channel.Channel())
        request = release_pb2.DeleteFirebaserulesAlphaReleaseRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.ruleset_name):
            request.resource.ruleset_name = Primitive.to_proto(self.ruleset_name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteFirebaserulesAlphaRelease(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = release_pb2_grpc.FirebaserulesAlphaReleaseServiceStub(channel.Channel())
        request = release_pb2.ListFirebaserulesAlphaReleaseRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListFirebaserulesAlphaRelease(request).items

    def to_proto(self):
        resource = release_pb2.FirebaserulesAlphaRelease()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.ruleset_name):
            resource.ruleset_name = Primitive.to_proto(self.ruleset_name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
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
