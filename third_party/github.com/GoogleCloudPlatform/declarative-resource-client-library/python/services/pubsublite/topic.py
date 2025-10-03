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
from google3.cloud.graphite.mmv2.services.google.pubsub_lite import topic_pb2
from google3.cloud.graphite.mmv2.services.google.pubsub_lite import topic_pb2_grpc

from typing import List


class Topic(object):
    def __init__(
        self,
        name: str = None,
        partition_config: dict = None,
        retention_config: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.partition_config = partition_config
        self.retention_config = retention_config
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = topic_pb2_grpc.PubsubliteTopicServiceStub(channel.Channel())
        request = topic_pb2.ApplyPubsubliteTopicRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if TopicPartitionConfig.to_proto(self.partition_config):
            request.resource.partition_config.CopyFrom(
                TopicPartitionConfig.to_proto(self.partition_config)
            )
        else:
            request.resource.ClearField("partition_config")
        if TopicRetentionConfig.to_proto(self.retention_config):
            request.resource.retention_config.CopyFrom(
                TopicRetentionConfig.to_proto(self.retention_config)
            )
        else:
            request.resource.ClearField("retention_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyPubsubliteTopic(request)
        self.name = Primitive.from_proto(response.name)
        self.partition_config = TopicPartitionConfig.from_proto(
            response.partition_config
        )
        self.retention_config = TopicRetentionConfig.from_proto(
            response.retention_config
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = topic_pb2_grpc.PubsubliteTopicServiceStub(channel.Channel())
        request = topic_pb2.DeletePubsubliteTopicRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if TopicPartitionConfig.to_proto(self.partition_config):
            request.resource.partition_config.CopyFrom(
                TopicPartitionConfig.to_proto(self.partition_config)
            )
        else:
            request.resource.ClearField("partition_config")
        if TopicRetentionConfig.to_proto(self.retention_config):
            request.resource.retention_config.CopyFrom(
                TopicRetentionConfig.to_proto(self.retention_config)
            )
        else:
            request.resource.ClearField("retention_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeletePubsubliteTopic(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = topic_pb2_grpc.PubsubliteTopicServiceStub(channel.Channel())
        request = topic_pb2.ListPubsubliteTopicRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListPubsubliteTopic(request).items

    def to_proto(self):
        resource = topic_pb2.PubsubliteTopic()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if TopicPartitionConfig.to_proto(self.partition_config):
            resource.partition_config.CopyFrom(
                TopicPartitionConfig.to_proto(self.partition_config)
            )
        else:
            resource.ClearField("partition_config")
        if TopicRetentionConfig.to_proto(self.retention_config):
            resource.retention_config.CopyFrom(
                TopicRetentionConfig.to_proto(self.retention_config)
            )
        else:
            resource.ClearField("retention_config")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class TopicPartitionConfig(object):
    def __init__(self, count: int = None, capacity: dict = None):
        self.count = count
        self.capacity = capacity

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = topic_pb2.PubsubliteTopicPartitionConfig()
        if Primitive.to_proto(resource.count):
            res.count = Primitive.to_proto(resource.count)
        if TopicPartitionConfigCapacity.to_proto(resource.capacity):
            res.capacity.CopyFrom(
                TopicPartitionConfigCapacity.to_proto(resource.capacity)
            )
        else:
            res.ClearField("capacity")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TopicPartitionConfig(
            count=Primitive.from_proto(resource.count),
            capacity=TopicPartitionConfigCapacity.from_proto(resource.capacity),
        )


class TopicPartitionConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TopicPartitionConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TopicPartitionConfig.from_proto(i) for i in resources]


class TopicPartitionConfigCapacity(object):
    def __init__(
        self, publish_mib_per_sec: int = None, subscribe_mib_per_sec: int = None
    ):
        self.publish_mib_per_sec = publish_mib_per_sec
        self.subscribe_mib_per_sec = subscribe_mib_per_sec

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = topic_pb2.PubsubliteTopicPartitionConfigCapacity()
        if Primitive.to_proto(resource.publish_mib_per_sec):
            res.publish_mib_per_sec = Primitive.to_proto(resource.publish_mib_per_sec)
        if Primitive.to_proto(resource.subscribe_mib_per_sec):
            res.subscribe_mib_per_sec = Primitive.to_proto(
                resource.subscribe_mib_per_sec
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TopicPartitionConfigCapacity(
            publish_mib_per_sec=Primitive.from_proto(resource.publish_mib_per_sec),
            subscribe_mib_per_sec=Primitive.from_proto(resource.subscribe_mib_per_sec),
        )


class TopicPartitionConfigCapacityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TopicPartitionConfigCapacity.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TopicPartitionConfigCapacity.from_proto(i) for i in resources]


class TopicRetentionConfig(object):
    def __init__(self, per_partition_bytes: int = None, period: str = None):
        self.per_partition_bytes = per_partition_bytes
        self.period = period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = topic_pb2.PubsubliteTopicRetentionConfig()
        if Primitive.to_proto(resource.per_partition_bytes):
            res.per_partition_bytes = Primitive.to_proto(resource.per_partition_bytes)
        if Primitive.to_proto(resource.period):
            res.period = Primitive.to_proto(resource.period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TopicRetentionConfig(
            per_partition_bytes=Primitive.from_proto(resource.per_partition_bytes),
            period=Primitive.from_proto(resource.period),
        )


class TopicRetentionConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TopicRetentionConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TopicRetentionConfig.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
