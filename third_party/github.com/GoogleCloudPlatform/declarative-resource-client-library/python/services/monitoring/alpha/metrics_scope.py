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
from google3.cloud.graphite.mmv2.services.google.monitoring import metrics_scope_pb2
from google3.cloud.graphite.mmv2.services.google.monitoring import (
    metrics_scope_pb2_grpc,
)

from typing import List


class MetricsScope(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        monitored_projects: list = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.service_account_file = service_account_file

    def apply(self):
        stub = metrics_scope_pb2_grpc.MonitoringAlphaMetricsScopeServiceStub(
            channel.Channel()
        )
        request = metrics_scope_pb2.ApplyMonitoringAlphaMetricsScopeRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringAlphaMetricsScope(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.monitored_projects = MetricsScopeMonitoredProjectsArray.from_proto(
            response.monitored_projects
        )

    def delete(self):
        stub = metrics_scope_pb2_grpc.MonitoringAlphaMetricsScopeServiceStub(
            channel.Channel()
        )
        request = metrics_scope_pb2.DeleteMonitoringAlphaMetricsScopeRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        response = stub.DeleteMonitoringAlphaMetricsScope(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = metrics_scope_pb2_grpc.MonitoringAlphaMetricsScopeServiceStub(
            channel.Channel()
        )
        request = metrics_scope_pb2.ListMonitoringAlphaMetricsScopeRequest()
        request.service_account_file = service_account_file
        return stub.ListMonitoringAlphaMetricsScope(request).items

    def to_proto(self):
        resource = metrics_scope_pb2.MonitoringAlphaMetricsScope()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        return resource


class MetricsScopeMonitoredProjects(object):
    def __init__(self, name: str = None, create_time: str = None):
        self.name = name
        self.create_time = create_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = metrics_scope_pb2.MonitoringAlphaMetricsScopeMonitoredProjects()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.create_time):
            res.create_time = Primitive.to_proto(resource.create_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MetricsScopeMonitoredProjects(
            name=Primitive.from_proto(resource.name),
            create_time=Primitive.from_proto(resource.create_time),
        )


class MetricsScopeMonitoredProjectsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MetricsScopeMonitoredProjects.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MetricsScopeMonitoredProjects.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
