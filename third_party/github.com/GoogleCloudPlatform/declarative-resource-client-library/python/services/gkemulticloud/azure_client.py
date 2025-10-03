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
from google3.cloud.graphite.mmv2.services.google.gkemulticloud import azure_client_pb2
from google3.cloud.graphite.mmv2.services.google.gkemulticloud import (
    azure_client_pb2_grpc,
)

from typing import List


class AzureClient(object):
    def __init__(
        self,
        name: str = None,
        tenant_id: str = None,
        application_id: str = None,
        certificate: str = None,
        uid: str = None,
        create_time: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.tenant_id = tenant_id
        self.application_id = application_id
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = azure_client_pb2_grpc.GkemulticloudAzureClientServiceStub(
            channel.Channel()
        )
        request = azure_client_pb2.ApplyGkemulticloudAzureClientRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.tenant_id):
            request.resource.tenant_id = Primitive.to_proto(self.tenant_id)

        if Primitive.to_proto(self.application_id):
            request.resource.application_id = Primitive.to_proto(self.application_id)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyGkemulticloudAzureClient(request)
        self.name = Primitive.from_proto(response.name)
        self.tenant_id = Primitive.from_proto(response.tenant_id)
        self.application_id = Primitive.from_proto(response.application_id)
        self.certificate = Primitive.from_proto(response.certificate)
        self.uid = Primitive.from_proto(response.uid)
        self.create_time = Primitive.from_proto(response.create_time)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = azure_client_pb2_grpc.GkemulticloudAzureClientServiceStub(
            channel.Channel()
        )
        request = azure_client_pb2.DeleteGkemulticloudAzureClientRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.tenant_id):
            request.resource.tenant_id = Primitive.to_proto(self.tenant_id)

        if Primitive.to_proto(self.application_id):
            request.resource.application_id = Primitive.to_proto(self.application_id)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteGkemulticloudAzureClient(request)

    def list(self):
        stub = azure_client_pb2_grpc.GkemulticloudAzureClientServiceStub(
            channel.Channel()
        )
        request = azure_client_pb2.ListGkemulticloudAzureClientRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.tenant_id):
            request.resource.tenant_id = Primitive.to_proto(self.tenant_id)

        if Primitive.to_proto(self.application_id):
            request.resource.application_id = Primitive.to_proto(self.application_id)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        return stub.ListGkemulticloudAzureClient(request).items

    def to_proto(self):
        resource = azure_client_pb2.GkemulticloudAzureClient()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.tenant_id):
            resource.tenant_id = Primitive.to_proto(self.tenant_id)
        if Primitive.to_proto(self.application_id):
            resource.application_id = Primitive.to_proto(self.application_id)
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
