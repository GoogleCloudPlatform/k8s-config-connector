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
from google3.cloud.graphite.mmv2.services.google.iap import brand_pb2
from google3.cloud.graphite.mmv2.services.google.iap import brand_pb2_grpc

from typing import List


class Brand(object):
    def __init__(
        self,
        application_title: str = None,
        name: str = None,
        org_internal_only: bool = None,
        support_email: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.application_title = application_title
        self.name = name
        self.support_email = support_email
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = brand_pb2_grpc.IapAlphaBrandServiceStub(channel.Channel())
        request = brand_pb2.ApplyIapAlphaBrandRequest()
        if Primitive.to_proto(self.application_title):
            request.resource.application_title = Primitive.to_proto(
                self.application_title
            )

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.support_email):
            request.resource.support_email = Primitive.to_proto(self.support_email)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyIapAlphaBrand(request)
        self.application_title = Primitive.from_proto(response.application_title)
        self.name = Primitive.from_proto(response.name)
        self.org_internal_only = Primitive.from_proto(response.org_internal_only)
        self.support_email = Primitive.from_proto(response.support_email)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = brand_pb2_grpc.IapAlphaBrandServiceStub(channel.Channel())
        request = brand_pb2.DeleteIapAlphaBrandRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.application_title):
            request.resource.application_title = Primitive.to_proto(
                self.application_title
            )

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.support_email):
            request.resource.support_email = Primitive.to_proto(self.support_email)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteIapAlphaBrand(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = brand_pb2_grpc.IapAlphaBrandServiceStub(channel.Channel())
        request = brand_pb2.ListIapAlphaBrandRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListIapAlphaBrand(request).items

    def to_proto(self):
        resource = brand_pb2.IapAlphaBrand()
        if Primitive.to_proto(self.application_title):
            resource.application_title = Primitive.to_proto(self.application_title)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.support_email):
            resource.support_email = Primitive.to_proto(self.support_email)
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
