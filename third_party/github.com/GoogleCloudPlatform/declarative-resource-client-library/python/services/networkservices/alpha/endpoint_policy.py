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
from google3.cloud.graphite.mmv2.services.google.network_services import (
    endpoint_policy_pb2,
)
from google3.cloud.graphite.mmv2.services.google.network_services import (
    endpoint_policy_pb2_grpc,
)

from typing import List


class EndpointPolicy(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        type: str = None,
        authorization_policy: str = None,
        endpoint_matcher: dict = None,
        traffic_port_selector: dict = None,
        description: str = None,
        server_tls_policy: str = None,
        client_tls_policy: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.labels = labels
        self.type = type
        self.authorization_policy = authorization_policy
        self.endpoint_matcher = endpoint_matcher
        self.traffic_port_selector = traffic_port_selector
        self.description = description
        self.server_tls_policy = server_tls_policy
        self.client_tls_policy = client_tls_policy
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = endpoint_policy_pb2_grpc.NetworkservicesAlphaEndpointPolicyServiceStub(
            channel.Channel()
        )
        request = endpoint_policy_pb2.ApplyNetworkservicesAlphaEndpointPolicyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if EndpointPolicyTypeEnum.to_proto(self.type):
            request.resource.type = EndpointPolicyTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.authorization_policy):
            request.resource.authorization_policy = Primitive.to_proto(
                self.authorization_policy
            )

        if EndpointPolicyEndpointMatcher.to_proto(self.endpoint_matcher):
            request.resource.endpoint_matcher.CopyFrom(
                EndpointPolicyEndpointMatcher.to_proto(self.endpoint_matcher)
            )
        else:
            request.resource.ClearField("endpoint_matcher")
        if EndpointPolicyTrafficPortSelector.to_proto(self.traffic_port_selector):
            request.resource.traffic_port_selector.CopyFrom(
                EndpointPolicyTrafficPortSelector.to_proto(self.traffic_port_selector)
            )
        else:
            request.resource.ClearField("traffic_port_selector")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.server_tls_policy):
            request.resource.server_tls_policy = Primitive.to_proto(
                self.server_tls_policy
            )

        if Primitive.to_proto(self.client_tls_policy):
            request.resource.client_tls_policy = Primitive.to_proto(
                self.client_tls_policy
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworkservicesAlphaEndpointPolicy(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.type = EndpointPolicyTypeEnum.from_proto(response.type)
        self.authorization_policy = Primitive.from_proto(response.authorization_policy)
        self.endpoint_matcher = EndpointPolicyEndpointMatcher.from_proto(
            response.endpoint_matcher
        )
        self.traffic_port_selector = EndpointPolicyTrafficPortSelector.from_proto(
            response.traffic_port_selector
        )
        self.description = Primitive.from_proto(response.description)
        self.server_tls_policy = Primitive.from_proto(response.server_tls_policy)
        self.client_tls_policy = Primitive.from_proto(response.client_tls_policy)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = endpoint_policy_pb2_grpc.NetworkservicesAlphaEndpointPolicyServiceStub(
            channel.Channel()
        )
        request = endpoint_policy_pb2.DeleteNetworkservicesAlphaEndpointPolicyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if EndpointPolicyTypeEnum.to_proto(self.type):
            request.resource.type = EndpointPolicyTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.authorization_policy):
            request.resource.authorization_policy = Primitive.to_proto(
                self.authorization_policy
            )

        if EndpointPolicyEndpointMatcher.to_proto(self.endpoint_matcher):
            request.resource.endpoint_matcher.CopyFrom(
                EndpointPolicyEndpointMatcher.to_proto(self.endpoint_matcher)
            )
        else:
            request.resource.ClearField("endpoint_matcher")
        if EndpointPolicyTrafficPortSelector.to_proto(self.traffic_port_selector):
            request.resource.traffic_port_selector.CopyFrom(
                EndpointPolicyTrafficPortSelector.to_proto(self.traffic_port_selector)
            )
        else:
            request.resource.ClearField("traffic_port_selector")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.server_tls_policy):
            request.resource.server_tls_policy = Primitive.to_proto(
                self.server_tls_policy
            )

        if Primitive.to_proto(self.client_tls_policy):
            request.resource.client_tls_policy = Primitive.to_proto(
                self.client_tls_policy
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworkservicesAlphaEndpointPolicy(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = endpoint_policy_pb2_grpc.NetworkservicesAlphaEndpointPolicyServiceStub(
            channel.Channel()
        )
        request = endpoint_policy_pb2.ListNetworkservicesAlphaEndpointPolicyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworkservicesAlphaEndpointPolicy(request).items

    def to_proto(self):
        resource = endpoint_policy_pb2.NetworkservicesAlphaEndpointPolicy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if EndpointPolicyTypeEnum.to_proto(self.type):
            resource.type = EndpointPolicyTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.authorization_policy):
            resource.authorization_policy = Primitive.to_proto(
                self.authorization_policy
            )
        if EndpointPolicyEndpointMatcher.to_proto(self.endpoint_matcher):
            resource.endpoint_matcher.CopyFrom(
                EndpointPolicyEndpointMatcher.to_proto(self.endpoint_matcher)
            )
        else:
            resource.ClearField("endpoint_matcher")
        if EndpointPolicyTrafficPortSelector.to_proto(self.traffic_port_selector):
            resource.traffic_port_selector.CopyFrom(
                EndpointPolicyTrafficPortSelector.to_proto(self.traffic_port_selector)
            )
        else:
            resource.ClearField("traffic_port_selector")
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.server_tls_policy):
            resource.server_tls_policy = Primitive.to_proto(self.server_tls_policy)
        if Primitive.to_proto(self.client_tls_policy):
            resource.client_tls_policy = Primitive.to_proto(self.client_tls_policy)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class EndpointPolicyEndpointMatcher(object):
    def __init__(self, metadata_label_matcher: dict = None):
        self.metadata_label_matcher = metadata_label_matcher

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = endpoint_policy_pb2.NetworkservicesAlphaEndpointPolicyEndpointMatcher()
        if EndpointPolicyEndpointMatcherMetadataLabelMatcher.to_proto(
            resource.metadata_label_matcher
        ):
            res.metadata_label_matcher.CopyFrom(
                EndpointPolicyEndpointMatcherMetadataLabelMatcher.to_proto(
                    resource.metadata_label_matcher
                )
            )
        else:
            res.ClearField("metadata_label_matcher")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointPolicyEndpointMatcher(
            metadata_label_matcher=EndpointPolicyEndpointMatcherMetadataLabelMatcher.from_proto(
                resource.metadata_label_matcher
            ),
        )


class EndpointPolicyEndpointMatcherArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EndpointPolicyEndpointMatcher.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EndpointPolicyEndpointMatcher.from_proto(i) for i in resources]


class EndpointPolicyEndpointMatcherMetadataLabelMatcher(object):
    def __init__(
        self, metadata_label_match_criteria: str = None, metadata_labels: list = None
    ):
        self.metadata_label_match_criteria = metadata_label_match_criteria
        self.metadata_labels = metadata_labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            endpoint_policy_pb2.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcher()
        )
        if EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.to_proto(
            resource.metadata_label_match_criteria
        ):
            res.metadata_label_match_criteria = EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.to_proto(
                resource.metadata_label_match_criteria
            )
        if EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsArray.to_proto(
            resource.metadata_labels
        ):
            res.metadata_labels.extend(
                EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsArray.to_proto(
                    resource.metadata_labels
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointPolicyEndpointMatcherMetadataLabelMatcher(
            metadata_label_match_criteria=EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.from_proto(
                resource.metadata_label_match_criteria
            ),
            metadata_labels=EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsArray.from_proto(
                resource.metadata_labels
            ),
        )


class EndpointPolicyEndpointMatcherMetadataLabelMatcherArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EndpointPolicyEndpointMatcherMetadataLabelMatcher.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EndpointPolicyEndpointMatcherMetadataLabelMatcher.from_proto(i)
            for i in resources
        ]


class EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(object):
    def __init__(self, label_name: str = None, label_value: str = None):
        self.label_name = label_name
        self.label_value = label_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            endpoint_policy_pb2.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels()
        )
        if Primitive.to_proto(resource.label_name):
            res.label_name = Primitive.to_proto(resource.label_name)
        if Primitive.to_proto(resource.label_value):
            res.label_value = Primitive.to_proto(resource.label_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(
            label_name=Primitive.from_proto(resource.label_name),
            label_value=Primitive.from_proto(resource.label_value),
        )


class EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels.from_proto(
                i
            )
            for i in resources
        ]


class EndpointPolicyTrafficPortSelector(object):
    def __init__(self, ports: list = None):
        self.ports = ports

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            endpoint_policy_pb2.NetworkservicesAlphaEndpointPolicyTrafficPortSelector()
        )
        if Primitive.to_proto(resource.ports):
            res.ports.extend(Primitive.to_proto(resource.ports))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointPolicyTrafficPortSelector(
            ports=Primitive.from_proto(resource.ports),
        )


class EndpointPolicyTrafficPortSelectorArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EndpointPolicyTrafficPortSelector.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EndpointPolicyTrafficPortSelector.from_proto(i) for i in resources]


class EndpointPolicyTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return endpoint_policy_pb2.NetworkservicesAlphaEndpointPolicyTypeEnum.Value(
            "NetworkservicesAlphaEndpointPolicyTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return endpoint_policy_pb2.NetworkservicesAlphaEndpointPolicyTypeEnum.Name(
            resource
        )[len("NetworkservicesAlphaEndpointPolicyTypeEnum") :]


class EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return endpoint_policy_pb2.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.Value(
            "NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return endpoint_policy_pb2.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.Name(
            resource
        )[
            len(
                "NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"
            ) :
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
