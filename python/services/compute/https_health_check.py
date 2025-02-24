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
from google3.cloud.graphite.mmv2.services.google.compute import https_health_check_pb2
from google3.cloud.graphite.mmv2.services.google.compute import (
    https_health_check_pb2_grpc,
)

from typing import List


class HttpsHealthCheck(object):
    def __init__(
        self,
        check_interval_sec: int = None,
        description: str = None,
        healthy_threshold: int = None,
        host: str = None,
        name: str = None,
        port: int = None,
        request_path: str = None,
        timeout_sec: int = None,
        unhealthy_threshold: int = None,
        project: str = None,
        self_link: str = None,
        creation_timestamp: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.check_interval_sec = check_interval_sec
        self.description = description
        self.healthy_threshold = healthy_threshold
        self.host = host
        self.name = name
        self.port = port
        self.request_path = request_path
        self.timeout_sec = timeout_sec
        self.unhealthy_threshold = unhealthy_threshold
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = https_health_check_pb2_grpc.ComputeHttpsHealthCheckServiceStub(
            channel.Channel()
        )
        request = https_health_check_pb2.ApplyComputeHttpsHealthCheckRequest()
        if Primitive.to_proto(self.check_interval_sec):
            request.resource.check_interval_sec = Primitive.to_proto(
                self.check_interval_sec
            )

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.healthy_threshold):
            request.resource.healthy_threshold = Primitive.to_proto(
                self.healthy_threshold
            )

        if Primitive.to_proto(self.host):
            request.resource.host = Primitive.to_proto(self.host)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.port):
            request.resource.port = Primitive.to_proto(self.port)

        if Primitive.to_proto(self.request_path):
            request.resource.request_path = Primitive.to_proto(self.request_path)

        if Primitive.to_proto(self.timeout_sec):
            request.resource.timeout_sec = Primitive.to_proto(self.timeout_sec)

        if Primitive.to_proto(self.unhealthy_threshold):
            request.resource.unhealthy_threshold = Primitive.to_proto(
                self.unhealthy_threshold
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeHttpsHealthCheck(request)
        self.check_interval_sec = Primitive.from_proto(response.check_interval_sec)
        self.description = Primitive.from_proto(response.description)
        self.healthy_threshold = Primitive.from_proto(response.healthy_threshold)
        self.host = Primitive.from_proto(response.host)
        self.name = Primitive.from_proto(response.name)
        self.port = Primitive.from_proto(response.port)
        self.request_path = Primitive.from_proto(response.request_path)
        self.timeout_sec = Primitive.from_proto(response.timeout_sec)
        self.unhealthy_threshold = Primitive.from_proto(response.unhealthy_threshold)
        self.project = Primitive.from_proto(response.project)
        self.self_link = Primitive.from_proto(response.self_link)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)

    def delete(self):
        stub = https_health_check_pb2_grpc.ComputeHttpsHealthCheckServiceStub(
            channel.Channel()
        )
        request = https_health_check_pb2.DeleteComputeHttpsHealthCheckRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.check_interval_sec):
            request.resource.check_interval_sec = Primitive.to_proto(
                self.check_interval_sec
            )

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.healthy_threshold):
            request.resource.healthy_threshold = Primitive.to_proto(
                self.healthy_threshold
            )

        if Primitive.to_proto(self.host):
            request.resource.host = Primitive.to_proto(self.host)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.port):
            request.resource.port = Primitive.to_proto(self.port)

        if Primitive.to_proto(self.request_path):
            request.resource.request_path = Primitive.to_proto(self.request_path)

        if Primitive.to_proto(self.timeout_sec):
            request.resource.timeout_sec = Primitive.to_proto(self.timeout_sec)

        if Primitive.to_proto(self.unhealthy_threshold):
            request.resource.unhealthy_threshold = Primitive.to_proto(
                self.unhealthy_threshold
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeHttpsHealthCheck(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = https_health_check_pb2_grpc.ComputeHttpsHealthCheckServiceStub(
            channel.Channel()
        )
        request = https_health_check_pb2.ListComputeHttpsHealthCheckRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeHttpsHealthCheck(request).items

    def to_proto(self):
        resource = https_health_check_pb2.ComputeHttpsHealthCheck()
        if Primitive.to_proto(self.check_interval_sec):
            resource.check_interval_sec = Primitive.to_proto(self.check_interval_sec)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.healthy_threshold):
            resource.healthy_threshold = Primitive.to_proto(self.healthy_threshold)
        if Primitive.to_proto(self.host):
            resource.host = Primitive.to_proto(self.host)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.port):
            resource.port = Primitive.to_proto(self.port)
        if Primitive.to_proto(self.request_path):
            resource.request_path = Primitive.to_proto(self.request_path)
        if Primitive.to_proto(self.timeout_sec):
            resource.timeout_sec = Primitive.to_proto(self.timeout_sec)
        if Primitive.to_proto(self.unhealthy_threshold):
            resource.unhealthy_threshold = Primitive.to_proto(self.unhealthy_threshold)
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
