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
from google3.cloud.graphite.mmv2.services.google.pubsub_lite import subscription_pb2
from google3.cloud.graphite.mmv2.services.google.pubsub_lite import (
    subscription_pb2_grpc,
)

from typing import List


class Subscription(object):
    def __init__(
        self,
        name: str = None,
        topic: str = None,
        delivery_config: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.topic = topic
        self.delivery_config = delivery_config
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = subscription_pb2_grpc.PubsubliteSubscriptionServiceStub(
            channel.Channel()
        )
        request = subscription_pb2.ApplyPubsubliteSubscriptionRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.topic):
            request.resource.topic = Primitive.to_proto(self.topic)

        if SubscriptionDeliveryConfig.to_proto(self.delivery_config):
            request.resource.delivery_config.CopyFrom(
                SubscriptionDeliveryConfig.to_proto(self.delivery_config)
            )
        else:
            request.resource.ClearField("delivery_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyPubsubliteSubscription(request)
        self.name = Primitive.from_proto(response.name)
        self.topic = Primitive.from_proto(response.topic)
        self.delivery_config = SubscriptionDeliveryConfig.from_proto(
            response.delivery_config
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = subscription_pb2_grpc.PubsubliteSubscriptionServiceStub(
            channel.Channel()
        )
        request = subscription_pb2.DeletePubsubliteSubscriptionRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.topic):
            request.resource.topic = Primitive.to_proto(self.topic)

        if SubscriptionDeliveryConfig.to_proto(self.delivery_config):
            request.resource.delivery_config.CopyFrom(
                SubscriptionDeliveryConfig.to_proto(self.delivery_config)
            )
        else:
            request.resource.ClearField("delivery_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeletePubsubliteSubscription(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = subscription_pb2_grpc.PubsubliteSubscriptionServiceStub(
            channel.Channel()
        )
        request = subscription_pb2.ListPubsubliteSubscriptionRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListPubsubliteSubscription(request).items

    def to_proto(self):
        resource = subscription_pb2.PubsubliteSubscription()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.topic):
            resource.topic = Primitive.to_proto(self.topic)
        if SubscriptionDeliveryConfig.to_proto(self.delivery_config):
            resource.delivery_config.CopyFrom(
                SubscriptionDeliveryConfig.to_proto(self.delivery_config)
            )
        else:
            resource.ClearField("delivery_config")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class SubscriptionDeliveryConfig(object):
    def __init__(self, delivery_requirement: str = None):
        self.delivery_requirement = delivery_requirement

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = subscription_pb2.PubsubliteSubscriptionDeliveryConfig()
        if SubscriptionDeliveryConfigDeliveryRequirementEnum.to_proto(
            resource.delivery_requirement
        ):
            res.delivery_requirement = SubscriptionDeliveryConfigDeliveryRequirementEnum.to_proto(
                resource.delivery_requirement
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SubscriptionDeliveryConfig(
            delivery_requirement=SubscriptionDeliveryConfigDeliveryRequirementEnum.from_proto(
                resource.delivery_requirement
            ),
        )


class SubscriptionDeliveryConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SubscriptionDeliveryConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SubscriptionDeliveryConfig.from_proto(i) for i in resources]


class SubscriptionDeliveryConfigDeliveryRequirementEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return subscription_pb2.PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum.Value(
            "PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return subscription_pb2.PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum.Name(
            resource
        )[
            len("PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum") :
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
