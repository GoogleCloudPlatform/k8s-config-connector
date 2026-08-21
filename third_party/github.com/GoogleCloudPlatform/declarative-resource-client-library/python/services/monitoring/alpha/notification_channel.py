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
from google3.cloud.graphite.mmv2.services.google.monitoring import (
    notification_channel_pb2,
)
from google3.cloud.graphite.mmv2.services.google.monitoring import (
    notification_channel_pb2_grpc,
)

from typing import List


class NotificationChannel(object):
    def __init__(
        self,
        description: str = None,
        display_name: str = None,
        enabled: bool = None,
        labels: dict = None,
        name: str = None,
        type: str = None,
        user_labels: dict = None,
        verification_status: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.description = description
        self.display_name = display_name
        self.enabled = enabled
        self.labels = labels
        self.name = name
        self.type = type
        self.user_labels = user_labels
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = (
            notification_channel_pb2_grpc.MonitoringAlphaNotificationChannelServiceStub(
                channel.Channel()
            )
        )
        request = (
            notification_channel_pb2.ApplyMonitoringAlphaNotificationChannelRequest()
        )
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.enabled):
            request.resource.enabled = Primitive.to_proto(self.enabled)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.type):
            request.resource.type = Primitive.to_proto(self.type)

        if Primitive.to_proto(self.user_labels):
            request.resource.user_labels = Primitive.to_proto(self.user_labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringAlphaNotificationChannel(request)
        self.description = Primitive.from_proto(response.description)
        self.display_name = Primitive.from_proto(response.display_name)
        self.enabled = Primitive.from_proto(response.enabled)
        self.labels = Primitive.from_proto(response.labels)
        self.name = Primitive.from_proto(response.name)
        self.type = Primitive.from_proto(response.type)
        self.user_labels = Primitive.from_proto(response.user_labels)
        self.verification_status = NotificationChannelVerificationStatusEnum.from_proto(
            response.verification_status
        )
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = (
            notification_channel_pb2_grpc.MonitoringAlphaNotificationChannelServiceStub(
                channel.Channel()
            )
        )
        request = (
            notification_channel_pb2.DeleteMonitoringAlphaNotificationChannelRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.enabled):
            request.resource.enabled = Primitive.to_proto(self.enabled)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.type):
            request.resource.type = Primitive.to_proto(self.type)

        if Primitive.to_proto(self.user_labels):
            request.resource.user_labels = Primitive.to_proto(self.user_labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteMonitoringAlphaNotificationChannel(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = (
            notification_channel_pb2_grpc.MonitoringAlphaNotificationChannelServiceStub(
                channel.Channel()
            )
        )
        request = (
            notification_channel_pb2.ListMonitoringAlphaNotificationChannelRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListMonitoringAlphaNotificationChannel(request).items

    def to_proto(self):
        resource = notification_channel_pb2.MonitoringAlphaNotificationChannel()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.enabled):
            resource.enabled = Primitive.to_proto(self.enabled)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.type):
            resource.type = Primitive.to_proto(self.type)
        if Primitive.to_proto(self.user_labels):
            resource.user_labels = Primitive.to_proto(self.user_labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class NotificationChannelVerificationStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return notification_channel_pb2.MonitoringAlphaNotificationChannelVerificationStatusEnum.Value(
            "MonitoringAlphaNotificationChannelVerificationStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return notification_channel_pb2.MonitoringAlphaNotificationChannelVerificationStatusEnum.Name(
            resource
        )[
            len("MonitoringAlphaNotificationChannelVerificationStatusEnum") :
        ]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
