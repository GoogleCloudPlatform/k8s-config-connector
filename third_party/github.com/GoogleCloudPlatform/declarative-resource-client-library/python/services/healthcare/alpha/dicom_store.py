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
from google3.cloud.graphite.mmv2.services.google.healthcare import dicom_store_pb2
from google3.cloud.graphite.mmv2.services.google.healthcare import dicom_store_pb2_grpc

from typing import List


class DicomStore(object):
    def __init__(
        self,
        name: str = None,
        notification_config: dict = None,
        labels: dict = None,
        project: str = None,
        location: str = None,
        dataset: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.notification_config = notification_config
        self.labels = labels
        self.project = project
        self.location = location
        self.dataset = dataset
        self.service_account_file = service_account_file

    def apply(self):
        stub = dicom_store_pb2_grpc.HealthcareAlphaDicomStoreServiceStub(
            channel.Channel()
        )
        request = dicom_store_pb2.ApplyHealthcareAlphaDicomStoreRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if DicomStoreNotificationConfig.to_proto(self.notification_config):
            request.resource.notification_config.CopyFrom(
                DicomStoreNotificationConfig.to_proto(self.notification_config)
            )
        else:
            request.resource.ClearField("notification_config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.dataset):
            request.resource.dataset = Primitive.to_proto(self.dataset)

        request.service_account_file = self.service_account_file

        response = stub.ApplyHealthcareAlphaDicomStore(request)
        self.name = Primitive.from_proto(response.name)
        self.notification_config = DicomStoreNotificationConfig.from_proto(
            response.notification_config
        )
        self.labels = Primitive.from_proto(response.labels)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.dataset = Primitive.from_proto(response.dataset)

    def delete(self):
        stub = dicom_store_pb2_grpc.HealthcareAlphaDicomStoreServiceStub(
            channel.Channel()
        )
        request = dicom_store_pb2.DeleteHealthcareAlphaDicomStoreRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if DicomStoreNotificationConfig.to_proto(self.notification_config):
            request.resource.notification_config.CopyFrom(
                DicomStoreNotificationConfig.to_proto(self.notification_config)
            )
        else:
            request.resource.ClearField("notification_config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.dataset):
            request.resource.dataset = Primitive.to_proto(self.dataset)

        response = stub.DeleteHealthcareAlphaDicomStore(request)

    @classmethod
    def list(self, project, location, dataset, service_account_file=""):
        stub = dicom_store_pb2_grpc.HealthcareAlphaDicomStoreServiceStub(
            channel.Channel()
        )
        request = dicom_store_pb2.ListHealthcareAlphaDicomStoreRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Dataset = dataset

        return stub.ListHealthcareAlphaDicomStore(request).items

    def to_proto(self):
        resource = dicom_store_pb2.HealthcareAlphaDicomStore()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if DicomStoreNotificationConfig.to_proto(self.notification_config):
            resource.notification_config.CopyFrom(
                DicomStoreNotificationConfig.to_proto(self.notification_config)
            )
        else:
            resource.ClearField("notification_config")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.dataset):
            resource.dataset = Primitive.to_proto(self.dataset)
        return resource


class DicomStoreNotificationConfig(object):
    def __init__(self, pubsub_topic: str = None):
        self.pubsub_topic = pubsub_topic

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dicom_store_pb2.HealthcareAlphaDicomStoreNotificationConfig()
        if Primitive.to_proto(resource.pubsub_topic):
            res.pubsub_topic = Primitive.to_proto(resource.pubsub_topic)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DicomStoreNotificationConfig(
            pubsub_topic=Primitive.from_proto(resource.pubsub_topic),
        )


class DicomStoreNotificationConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DicomStoreNotificationConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DicomStoreNotificationConfig.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
