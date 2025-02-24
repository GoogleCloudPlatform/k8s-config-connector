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
from google3.cloud.graphite.mmv2.services.google.monitoring import monitored_project_pb2
from google3.cloud.graphite.mmv2.services.google.monitoring import (
    monitored_project_pb2_grpc,
)

from typing import List


class MonitoredProject(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        metrics_scope: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.metrics_scope = metrics_scope
        self.service_account_file = service_account_file

    def apply(self):
        stub = monitored_project_pb2_grpc.MonitoringMonitoredProjectServiceStub(
            channel.Channel()
        )
        request = monitored_project_pb2.ApplyMonitoringMonitoredProjectRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.metrics_scope):
            request.resource.metrics_scope = Primitive.to_proto(self.metrics_scope)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringMonitoredProject(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.metrics_scope = Primitive.from_proto(response.metrics_scope)

    def delete(self):
        stub = monitored_project_pb2_grpc.MonitoringMonitoredProjectServiceStub(
            channel.Channel()
        )
        request = monitored_project_pb2.DeleteMonitoringMonitoredProjectRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.metrics_scope):
            request.resource.metrics_scope = Primitive.to_proto(self.metrics_scope)

        response = stub.DeleteMonitoringMonitoredProject(request)

    @classmethod
    def list(self, metricsScope, service_account_file=""):
        stub = monitored_project_pb2_grpc.MonitoringMonitoredProjectServiceStub(
            channel.Channel()
        )
        request = monitored_project_pb2.ListMonitoringMonitoredProjectRequest()
        request.service_account_file = service_account_file
        request.MetricsScope = metricsScope

        return stub.ListMonitoringMonitoredProject(request).items

    def to_proto(self):
        resource = monitored_project_pb2.MonitoringMonitoredProject()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.metrics_scope):
            resource.metrics_scope = Primitive.to_proto(self.metrics_scope)
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
