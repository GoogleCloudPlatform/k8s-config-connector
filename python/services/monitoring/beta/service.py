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
from google3.cloud.graphite.mmv2.services.google.monitoring import service_pb2
from google3.cloud.graphite.mmv2.services.google.monitoring import service_pb2_grpc

from typing import List


class Service(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        custom: dict = None,
        telemetry: dict = None,
        user_labels: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.custom = custom
        self.telemetry = telemetry
        self.user_labels = user_labels
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = service_pb2_grpc.MonitoringBetaServiceServiceStub(channel.Channel())
        request = service_pb2.ApplyMonitoringBetaServiceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if ServiceCustom.to_proto(self.custom):
            request.resource.custom.CopyFrom(ServiceCustom.to_proto(self.custom))
        else:
            request.resource.ClearField("custom")
        if ServiceTelemetry.to_proto(self.telemetry):
            request.resource.telemetry.CopyFrom(
                ServiceTelemetry.to_proto(self.telemetry)
            )
        else:
            request.resource.ClearField("telemetry")
        if Primitive.to_proto(self.user_labels):
            request.resource.user_labels = Primitive.to_proto(self.user_labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringBetaService(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.custom = ServiceCustom.from_proto(response.custom)
        self.telemetry = ServiceTelemetry.from_proto(response.telemetry)
        self.user_labels = Primitive.from_proto(response.user_labels)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = service_pb2_grpc.MonitoringBetaServiceServiceStub(channel.Channel())
        request = service_pb2.DeleteMonitoringBetaServiceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if ServiceCustom.to_proto(self.custom):
            request.resource.custom.CopyFrom(ServiceCustom.to_proto(self.custom))
        else:
            request.resource.ClearField("custom")
        if ServiceTelemetry.to_proto(self.telemetry):
            request.resource.telemetry.CopyFrom(
                ServiceTelemetry.to_proto(self.telemetry)
            )
        else:
            request.resource.ClearField("telemetry")
        if Primitive.to_proto(self.user_labels):
            request.resource.user_labels = Primitive.to_proto(self.user_labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteMonitoringBetaService(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = service_pb2_grpc.MonitoringBetaServiceServiceStub(channel.Channel())
        request = service_pb2.ListMonitoringBetaServiceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListMonitoringBetaService(request).items

    def to_proto(self):
        resource = service_pb2.MonitoringBetaService()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if ServiceCustom.to_proto(self.custom):
            resource.custom.CopyFrom(ServiceCustom.to_proto(self.custom))
        else:
            resource.ClearField("custom")
        if ServiceTelemetry.to_proto(self.telemetry):
            resource.telemetry.CopyFrom(ServiceTelemetry.to_proto(self.telemetry))
        else:
            resource.ClearField("telemetry")
        if Primitive.to_proto(self.user_labels):
            resource.user_labels = Primitive.to_proto(self.user_labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class ServiceCustom(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringBetaServiceCustom()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceCustom()


class ServiceCustomArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceCustom.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceCustom.from_proto(i) for i in resources]


class ServiceTelemetry(object):
    def __init__(self, resource_name: str = None):
        self.resource_name = resource_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringBetaServiceTelemetry()
        if Primitive.to_proto(resource.resource_name):
            res.resource_name = Primitive.to_proto(resource.resource_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTelemetry(
            resource_name=Primitive.from_proto(resource.resource_name),
        )


class ServiceTelemetryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTelemetry.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTelemetry.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
