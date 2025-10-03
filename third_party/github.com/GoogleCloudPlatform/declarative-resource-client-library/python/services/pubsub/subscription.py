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
from google3.cloud.graphite.mmv2.services.google.pubsub import subscription_pb2
from google3.cloud.graphite.mmv2.services.google.pubsub import subscription_pb2_grpc

from typing import List


class Subscription(object):
    def __init__(
        self,
        name: str = None,
        topic: str = None,
        labels: dict = None,
        message_retention_duration: str = None,
        retain_acked_messages: bool = None,
        expiration_policy: dict = None,
        project: str = None,
        dead_letter_policy: dict = None,
        push_config: dict = None,
        ack_deadline_seconds: int = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.topic = topic
        self.labels = labels
        self.message_retention_duration = message_retention_duration
        self.retain_acked_messages = retain_acked_messages
        self.expiration_policy = expiration_policy
        self.project = project
        self.dead_letter_policy = dead_letter_policy
        self.push_config = push_config
        self.ack_deadline_seconds = ack_deadline_seconds
        self.service_account_file = service_account_file

    def apply(self):
        stub = subscription_pb2_grpc.PubsubSubscriptionServiceStub(channel.Channel())
        request = subscription_pb2.ApplyPubsubSubscriptionRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.topic):
            request.resource.topic = Primitive.to_proto(self.topic)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.message_retention_duration):
            request.resource.message_retention_duration = Primitive.to_proto(
                self.message_retention_duration
            )

        if Primitive.to_proto(self.retain_acked_messages):
            request.resource.retain_acked_messages = Primitive.to_proto(
                self.retain_acked_messages
            )

        if SubscriptionExpirationPolicy.to_proto(self.expiration_policy):
            request.resource.expiration_policy.CopyFrom(
                SubscriptionExpirationPolicy.to_proto(self.expiration_policy)
            )
        else:
            request.resource.ClearField("expiration_policy")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if SubscriptionDeadLetterPolicy.to_proto(self.dead_letter_policy):
            request.resource.dead_letter_policy.CopyFrom(
                SubscriptionDeadLetterPolicy.to_proto(self.dead_letter_policy)
            )
        else:
            request.resource.ClearField("dead_letter_policy")
        if SubscriptionPushConfig.to_proto(self.push_config):
            request.resource.push_config.CopyFrom(
                SubscriptionPushConfig.to_proto(self.push_config)
            )
        else:
            request.resource.ClearField("push_config")
        if Primitive.to_proto(self.ack_deadline_seconds):
            request.resource.ack_deadline_seconds = Primitive.to_proto(
                self.ack_deadline_seconds
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyPubsubSubscription(request)
        self.name = Primitive.from_proto(response.name)
        self.topic = Primitive.from_proto(response.topic)
        self.labels = Primitive.from_proto(response.labels)
        self.message_retention_duration = Primitive.from_proto(
            response.message_retention_duration
        )
        self.retain_acked_messages = Primitive.from_proto(
            response.retain_acked_messages
        )
        self.expiration_policy = SubscriptionExpirationPolicy.from_proto(
            response.expiration_policy
        )
        self.project = Primitive.from_proto(response.project)
        self.dead_letter_policy = SubscriptionDeadLetterPolicy.from_proto(
            response.dead_letter_policy
        )
        self.push_config = SubscriptionPushConfig.from_proto(response.push_config)
        self.ack_deadline_seconds = Primitive.from_proto(response.ack_deadline_seconds)

    def delete(self):
        stub = subscription_pb2_grpc.PubsubSubscriptionServiceStub(channel.Channel())
        request = subscription_pb2.DeletePubsubSubscriptionRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.topic):
            request.resource.topic = Primitive.to_proto(self.topic)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.message_retention_duration):
            request.resource.message_retention_duration = Primitive.to_proto(
                self.message_retention_duration
            )

        if Primitive.to_proto(self.retain_acked_messages):
            request.resource.retain_acked_messages = Primitive.to_proto(
                self.retain_acked_messages
            )

        if SubscriptionExpirationPolicy.to_proto(self.expiration_policy):
            request.resource.expiration_policy.CopyFrom(
                SubscriptionExpirationPolicy.to_proto(self.expiration_policy)
            )
        else:
            request.resource.ClearField("expiration_policy")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if SubscriptionDeadLetterPolicy.to_proto(self.dead_letter_policy):
            request.resource.dead_letter_policy.CopyFrom(
                SubscriptionDeadLetterPolicy.to_proto(self.dead_letter_policy)
            )
        else:
            request.resource.ClearField("dead_letter_policy")
        if SubscriptionPushConfig.to_proto(self.push_config):
            request.resource.push_config.CopyFrom(
                SubscriptionPushConfig.to_proto(self.push_config)
            )
        else:
            request.resource.ClearField("push_config")
        if Primitive.to_proto(self.ack_deadline_seconds):
            request.resource.ack_deadline_seconds = Primitive.to_proto(
                self.ack_deadline_seconds
            )

        response = stub.DeletePubsubSubscription(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = subscription_pb2_grpc.PubsubSubscriptionServiceStub(channel.Channel())
        request = subscription_pb2.ListPubsubSubscriptionRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListPubsubSubscription(request).items

    def to_proto(self):
        resource = subscription_pb2.PubsubSubscription()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.topic):
            resource.topic = Primitive.to_proto(self.topic)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.message_retention_duration):
            resource.message_retention_duration = Primitive.to_proto(
                self.message_retention_duration
            )
        if Primitive.to_proto(self.retain_acked_messages):
            resource.retain_acked_messages = Primitive.to_proto(
                self.retain_acked_messages
            )
        if SubscriptionExpirationPolicy.to_proto(self.expiration_policy):
            resource.expiration_policy.CopyFrom(
                SubscriptionExpirationPolicy.to_proto(self.expiration_policy)
            )
        else:
            resource.ClearField("expiration_policy")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if SubscriptionDeadLetterPolicy.to_proto(self.dead_letter_policy):
            resource.dead_letter_policy.CopyFrom(
                SubscriptionDeadLetterPolicy.to_proto(self.dead_letter_policy)
            )
        else:
            resource.ClearField("dead_letter_policy")
        if SubscriptionPushConfig.to_proto(self.push_config):
            resource.push_config.CopyFrom(
                SubscriptionPushConfig.to_proto(self.push_config)
            )
        else:
            resource.ClearField("push_config")
        if Primitive.to_proto(self.ack_deadline_seconds):
            resource.ack_deadline_seconds = Primitive.to_proto(
                self.ack_deadline_seconds
            )
        return resource


class SubscriptionExpirationPolicy(object):
    def __init__(self, ttl: str = None):
        self.ttl = ttl

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = subscription_pb2.PubsubSubscriptionExpirationPolicy()
        if Primitive.to_proto(resource.ttl):
            res.ttl = Primitive.to_proto(resource.ttl)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SubscriptionExpirationPolicy(ttl=Primitive.from_proto(resource.ttl),)


class SubscriptionExpirationPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SubscriptionExpirationPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SubscriptionExpirationPolicy.from_proto(i) for i in resources]


class SubscriptionDeadLetterPolicy(object):
    def __init__(
        self, dead_letter_topic: str = None, max_delivery_attempts: int = None
    ):
        self.dead_letter_topic = dead_letter_topic
        self.max_delivery_attempts = max_delivery_attempts

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = subscription_pb2.PubsubSubscriptionDeadLetterPolicy()
        if Primitive.to_proto(resource.dead_letter_topic):
            res.dead_letter_topic = Primitive.to_proto(resource.dead_letter_topic)
        if Primitive.to_proto(resource.max_delivery_attempts):
            res.max_delivery_attempts = Primitive.to_proto(
                resource.max_delivery_attempts
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SubscriptionDeadLetterPolicy(
            dead_letter_topic=Primitive.from_proto(resource.dead_letter_topic),
            max_delivery_attempts=Primitive.from_proto(resource.max_delivery_attempts),
        )


class SubscriptionDeadLetterPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SubscriptionDeadLetterPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SubscriptionDeadLetterPolicy.from_proto(i) for i in resources]


class SubscriptionPushConfig(object):
    def __init__(
        self,
        push_endpoint: str = None,
        attributes: dict = None,
        oidc_token: dict = None,
    ):
        self.push_endpoint = push_endpoint
        self.attributes = attributes
        self.oidc_token = oidc_token

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = subscription_pb2.PubsubSubscriptionPushConfig()
        if Primitive.to_proto(resource.push_endpoint):
            res.push_endpoint = Primitive.to_proto(resource.push_endpoint)
        if Primitive.to_proto(resource.attributes):
            res.attributes = Primitive.to_proto(resource.attributes)
        if SubscriptionPushConfigOidcToken.to_proto(resource.oidc_token):
            res.oidc_token.CopyFrom(
                SubscriptionPushConfigOidcToken.to_proto(resource.oidc_token)
            )
        else:
            res.ClearField("oidc_token")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SubscriptionPushConfig(
            push_endpoint=Primitive.from_proto(resource.push_endpoint),
            attributes=Primitive.from_proto(resource.attributes),
            oidc_token=SubscriptionPushConfigOidcToken.from_proto(resource.oidc_token),
        )


class SubscriptionPushConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SubscriptionPushConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SubscriptionPushConfig.from_proto(i) for i in resources]


class SubscriptionPushConfigOidcToken(object):
    def __init__(self, service_account_email: str = None, audience: str = None):
        self.service_account_email = service_account_email
        self.audience = audience

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = subscription_pb2.PubsubSubscriptionPushConfigOidcToken()
        if Primitive.to_proto(resource.service_account_email):
            res.service_account_email = Primitive.to_proto(
                resource.service_account_email
            )
        if Primitive.to_proto(resource.audience):
            res.audience = Primitive.to_proto(resource.audience)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SubscriptionPushConfigOidcToken(
            service_account_email=Primitive.from_proto(resource.service_account_email),
            audience=Primitive.from_proto(resource.audience),
        )


class SubscriptionPushConfigOidcTokenArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SubscriptionPushConfigOidcToken.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SubscriptionPushConfigOidcToken.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
