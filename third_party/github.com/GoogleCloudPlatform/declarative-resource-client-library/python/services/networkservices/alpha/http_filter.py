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
from google3.cloud.graphite.mmv2.services.google.network_services import http_filter_pb2
from google3.cloud.graphite.mmv2.services.google.network_services import (
    http_filter_pb2_grpc,
)

from typing import List


class HttpFilter(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        filter_name: str = None,
        config_type_url: str = None,
        config: str = None,
        description: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.labels = labels
        self.filter_name = filter_name
        self.config_type_url = config_type_url
        self.config = config
        self.description = description
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = http_filter_pb2_grpc.NetworkservicesAlphaHttpFilterServiceStub(
            channel.Channel()
        )
        request = http_filter_pb2.ApplyNetworkservicesAlphaHttpFilterRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.filter_name):
            request.resource.filter_name = Primitive.to_proto(self.filter_name)

        if Primitive.to_proto(self.config_type_url):
            request.resource.config_type_url = Primitive.to_proto(self.config_type_url)

        if Primitive.to_proto(self.config):
            request.resource.config = Primitive.to_proto(self.config)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworkservicesAlphaHttpFilter(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.filter_name = Primitive.from_proto(response.filter_name)
        self.config_type_url = Primitive.from_proto(response.config_type_url)
        self.config = Primitive.from_proto(response.config)
        self.description = Primitive.from_proto(response.description)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = http_filter_pb2_grpc.NetworkservicesAlphaHttpFilterServiceStub(
            channel.Channel()
        )
        request = http_filter_pb2.DeleteNetworkservicesAlphaHttpFilterRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.filter_name):
            request.resource.filter_name = Primitive.to_proto(self.filter_name)

        if Primitive.to_proto(self.config_type_url):
            request.resource.config_type_url = Primitive.to_proto(self.config_type_url)

        if Primitive.to_proto(self.config):
            request.resource.config = Primitive.to_proto(self.config)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworkservicesAlphaHttpFilter(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = http_filter_pb2_grpc.NetworkservicesAlphaHttpFilterServiceStub(
            channel.Channel()
        )
        request = http_filter_pb2.ListNetworkservicesAlphaHttpFilterRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworkservicesAlphaHttpFilter(request).items

    def to_proto(self):
        resource = http_filter_pb2.NetworkservicesAlphaHttpFilter()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.filter_name):
            resource.filter_name = Primitive.to_proto(self.filter_name)
        if Primitive.to_proto(self.config_type_url):
            resource.config_type_url = Primitive.to_proto(self.config_type_url)
        if Primitive.to_proto(self.config):
            resource.config = Primitive.to_proto(self.config)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
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
