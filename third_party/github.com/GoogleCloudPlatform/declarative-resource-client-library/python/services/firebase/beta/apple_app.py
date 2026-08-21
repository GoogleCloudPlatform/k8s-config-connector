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
from google3.cloud.graphite.mmv2.services.google.firebase import apple_app_pb2
from google3.cloud.graphite.mmv2.services.google.firebase import apple_app_pb2_grpc

from typing import List


class AppleApp(object):
    def __init__(
        self,
        name: str = None,
        app_id: str = None,
        display_name: str = None,
        project_id: str = None,
        bundle_id: str = None,
        app_store_id: str = None,
        team_id: str = None,
        api_key_id: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.bundle_id = bundle_id
        self.app_store_id = app_store_id
        self.team_id = team_id
        self.api_key_id = api_key_id
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = apple_app_pb2_grpc.FirebaseBetaAppleAppServiceStub(channel.Channel())
        request = apple_app_pb2.ApplyFirebaseBetaAppleAppRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.bundle_id):
            request.resource.bundle_id = Primitive.to_proto(self.bundle_id)

        if Primitive.to_proto(self.app_store_id):
            request.resource.app_store_id = Primitive.to_proto(self.app_store_id)

        if Primitive.to_proto(self.team_id):
            request.resource.team_id = Primitive.to_proto(self.team_id)

        if Primitive.to_proto(self.api_key_id):
            request.resource.api_key_id = Primitive.to_proto(self.api_key_id)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyFirebaseBetaAppleApp(request)
        self.name = Primitive.from_proto(response.name)
        self.app_id = Primitive.from_proto(response.app_id)
        self.display_name = Primitive.from_proto(response.display_name)
        self.project_id = Primitive.from_proto(response.project_id)
        self.bundle_id = Primitive.from_proto(response.bundle_id)
        self.app_store_id = Primitive.from_proto(response.app_store_id)
        self.team_id = Primitive.from_proto(response.team_id)
        self.api_key_id = Primitive.from_proto(response.api_key_id)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = apple_app_pb2_grpc.FirebaseBetaAppleAppServiceStub(channel.Channel())
        request = apple_app_pb2.DeleteFirebaseBetaAppleAppRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.bundle_id):
            request.resource.bundle_id = Primitive.to_proto(self.bundle_id)

        if Primitive.to_proto(self.app_store_id):
            request.resource.app_store_id = Primitive.to_proto(self.app_store_id)

        if Primitive.to_proto(self.team_id):
            request.resource.team_id = Primitive.to_proto(self.team_id)

        if Primitive.to_proto(self.api_key_id):
            request.resource.api_key_id = Primitive.to_proto(self.api_key_id)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteFirebaseBetaAppleApp(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = apple_app_pb2_grpc.FirebaseBetaAppleAppServiceStub(channel.Channel())
        request = apple_app_pb2.ListFirebaseBetaAppleAppRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListFirebaseBetaAppleApp(request).items

    def to_proto(self):
        resource = apple_app_pb2.FirebaseBetaAppleApp()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.bundle_id):
            resource.bundle_id = Primitive.to_proto(self.bundle_id)
        if Primitive.to_proto(self.app_store_id):
            resource.app_store_id = Primitive.to_proto(self.app_store_id)
        if Primitive.to_proto(self.team_id):
            resource.team_id = Primitive.to_proto(self.team_id)
        if Primitive.to_proto(self.api_key_id):
            resource.api_key_id = Primitive.to_proto(self.api_key_id)
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
