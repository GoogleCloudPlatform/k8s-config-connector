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
from google3.cloud.graphite.mmv2.services.google.pubsub import topic_pb2
from google3.cloud.graphite.mmv2.services.google.pubsub import topic_pb2_grpc

from typing import List


class Topic(object):
    def __init__(
        self,
        name: str = None,
        kms_key_name: str = None,
        labels: dict = None,
        message_storage_policy: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.kms_key_name = kms_key_name
        self.labels = labels
        self.message_storage_policy = message_storage_policy
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = topic_pb2_grpc.PubsubBetaTopicServiceStub(channel.Channel())
        request = topic_pb2.ApplyPubsubBetaTopicRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.kms_key_name):
            request.resource.kms_key_name = Primitive.to_proto(self.kms_key_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if TopicMessageStoragePolicy.to_proto(self.message_storage_policy):
            request.resource.message_storage_policy.CopyFrom(
                TopicMessageStoragePolicy.to_proto(self.message_storage_policy)
            )
        else:
            request.resource.ClearField("message_storage_policy")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyPubsubBetaTopic(request)
        self.name = Primitive.from_proto(response.name)
        self.kms_key_name = Primitive.from_proto(response.kms_key_name)
        self.labels = Primitive.from_proto(response.labels)
        self.message_storage_policy = TopicMessageStoragePolicy.from_proto(
            response.message_storage_policy
        )
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = topic_pb2_grpc.PubsubBetaTopicServiceStub(channel.Channel())
        request = topic_pb2.DeletePubsubBetaTopicRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.kms_key_name):
            request.resource.kms_key_name = Primitive.to_proto(self.kms_key_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if TopicMessageStoragePolicy.to_proto(self.message_storage_policy):
            request.resource.message_storage_policy.CopyFrom(
                TopicMessageStoragePolicy.to_proto(self.message_storage_policy)
            )
        else:
            request.resource.ClearField("message_storage_policy")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeletePubsubBetaTopic(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = topic_pb2_grpc.PubsubBetaTopicServiceStub(channel.Channel())
        request = topic_pb2.ListPubsubBetaTopicRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListPubsubBetaTopic(request).items

    def to_proto(self):
        resource = topic_pb2.PubsubBetaTopic()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.kms_key_name):
            resource.kms_key_name = Primitive.to_proto(self.kms_key_name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if TopicMessageStoragePolicy.to_proto(self.message_storage_policy):
            resource.message_storage_policy.CopyFrom(
                TopicMessageStoragePolicy.to_proto(self.message_storage_policy)
            )
        else:
            resource.ClearField("message_storage_policy")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class TopicMessageStoragePolicy(object):
    def __init__(self, allowed_persistence_regions: list = None):
        self.allowed_persistence_regions = allowed_persistence_regions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = topic_pb2.PubsubBetaTopicMessageStoragePolicy()
        if Primitive.to_proto(resource.allowed_persistence_regions):
            res.allowed_persistence_regions.extend(
                Primitive.to_proto(resource.allowed_persistence_regions)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TopicMessageStoragePolicy(
            allowed_persistence_regions=Primitive.from_proto(
                resource.allowed_persistence_regions
            ),
        )


class TopicMessageStoragePolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TopicMessageStoragePolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TopicMessageStoragePolicy.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
