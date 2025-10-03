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
from google3.cloud.graphite.mmv2.services.google.network_services import (
    endpoint_config_selector_pb2,
)
from google3.cloud.graphite.mmv2.services.google.network_services import (
    endpoint_config_selector_pb2_grpc,
)

from typing import List


class EndpointConfigSelector(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        type: str = None,
        authorization_policy: str = None,
        http_filters: dict = None,
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
        self.http_filters = http_filters
        self.endpoint_matcher = endpoint_matcher
        self.traffic_port_selector = traffic_port_selector
        self.description = description
        self.server_tls_policy = server_tls_policy
        self.client_tls_policy = client_tls_policy
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = endpoint_config_selector_pb2_grpc.NetworkservicesAlphaEndpointConfigSelectorServiceStub(
            channel.Channel()
        )
        request = (
            endpoint_config_selector_pb2.ApplyNetworkservicesAlphaEndpointConfigSelectorRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if EndpointConfigSelectorTypeEnum.to_proto(self.type):
            request.resource.type = EndpointConfigSelectorTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.authorization_policy):
            request.resource.authorization_policy = Primitive.to_proto(
                self.authorization_policy
            )

        if EndpointConfigSelectorHttpFilters.to_proto(self.http_filters):
            request.resource.http_filters.CopyFrom(
                EndpointConfigSelectorHttpFilters.to_proto(self.http_filters)
            )
        else:
            request.resource.ClearField("http_filters")
        if EndpointConfigSelectorEndpointMatcher.to_proto(self.endpoint_matcher):
            request.resource.endpoint_matcher.CopyFrom(
                EndpointConfigSelectorEndpointMatcher.to_proto(self.endpoint_matcher)
            )
        else:
            request.resource.ClearField("endpoint_matcher")
        if EndpointConfigSelectorTrafficPortSelector.to_proto(
            self.traffic_port_selector
        ):
            request.resource.traffic_port_selector.CopyFrom(
                EndpointConfigSelectorTrafficPortSelector.to_proto(
                    self.traffic_port_selector
                )
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

        response = stub.ApplyNetworkservicesAlphaEndpointConfigSelector(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.type = EndpointConfigSelectorTypeEnum.from_proto(response.type)
        self.authorization_policy = Primitive.from_proto(response.authorization_policy)
        self.http_filters = EndpointConfigSelectorHttpFilters.from_proto(
            response.http_filters
        )
        self.endpoint_matcher = EndpointConfigSelectorEndpointMatcher.from_proto(
            response.endpoint_matcher
        )
        self.traffic_port_selector = EndpointConfigSelectorTrafficPortSelector.from_proto(
            response.traffic_port_selector
        )
        self.description = Primitive.from_proto(response.description)
        self.server_tls_policy = Primitive.from_proto(response.server_tls_policy)
        self.client_tls_policy = Primitive.from_proto(response.client_tls_policy)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = endpoint_config_selector_pb2_grpc.NetworkservicesAlphaEndpointConfigSelectorServiceStub(
            channel.Channel()
        )
        request = (
            endpoint_config_selector_pb2.DeleteNetworkservicesAlphaEndpointConfigSelectorRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if EndpointConfigSelectorTypeEnum.to_proto(self.type):
            request.resource.type = EndpointConfigSelectorTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.authorization_policy):
            request.resource.authorization_policy = Primitive.to_proto(
                self.authorization_policy
            )

        if EndpointConfigSelectorHttpFilters.to_proto(self.http_filters):
            request.resource.http_filters.CopyFrom(
                EndpointConfigSelectorHttpFilters.to_proto(self.http_filters)
            )
        else:
            request.resource.ClearField("http_filters")
        if EndpointConfigSelectorEndpointMatcher.to_proto(self.endpoint_matcher):
            request.resource.endpoint_matcher.CopyFrom(
                EndpointConfigSelectorEndpointMatcher.to_proto(self.endpoint_matcher)
            )
        else:
            request.resource.ClearField("endpoint_matcher")
        if EndpointConfigSelectorTrafficPortSelector.to_proto(
            self.traffic_port_selector
        ):
            request.resource.traffic_port_selector.CopyFrom(
                EndpointConfigSelectorTrafficPortSelector.to_proto(
                    self.traffic_port_selector
                )
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

        response = stub.DeleteNetworkservicesAlphaEndpointConfigSelector(request)

    def list(self):
        stub = endpoint_config_selector_pb2_grpc.NetworkservicesAlphaEndpointConfigSelectorServiceStub(
            channel.Channel()
        )
        request = (
            endpoint_config_selector_pb2.ListNetworkservicesAlphaEndpointConfigSelectorRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if EndpointConfigSelectorTypeEnum.to_proto(self.type):
            request.resource.type = EndpointConfigSelectorTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.authorization_policy):
            request.resource.authorization_policy = Primitive.to_proto(
                self.authorization_policy
            )

        if EndpointConfigSelectorHttpFilters.to_proto(self.http_filters):
            request.resource.http_filters.CopyFrom(
                EndpointConfigSelectorHttpFilters.to_proto(self.http_filters)
            )
        else:
            request.resource.ClearField("http_filters")
        if EndpointConfigSelectorEndpointMatcher.to_proto(self.endpoint_matcher):
            request.resource.endpoint_matcher.CopyFrom(
                EndpointConfigSelectorEndpointMatcher.to_proto(self.endpoint_matcher)
            )
        else:
            request.resource.ClearField("endpoint_matcher")
        if EndpointConfigSelectorTrafficPortSelector.to_proto(
            self.traffic_port_selector
        ):
            request.resource.traffic_port_selector.CopyFrom(
                EndpointConfigSelectorTrafficPortSelector.to_proto(
                    self.traffic_port_selector
                )
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

        return stub.ListNetworkservicesAlphaEndpointConfigSelector(request).items

    def to_proto(self):
        resource = (
            endpoint_config_selector_pb2.NetworkservicesAlphaEndpointConfigSelector()
        )
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if EndpointConfigSelectorTypeEnum.to_proto(self.type):
            resource.type = EndpointConfigSelectorTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.authorization_policy):
            resource.authorization_policy = Primitive.to_proto(
                self.authorization_policy
            )
        if EndpointConfigSelectorHttpFilters.to_proto(self.http_filters):
            resource.http_filters.CopyFrom(
                EndpointConfigSelectorHttpFilters.to_proto(self.http_filters)
            )
        else:
            resource.ClearField("http_filters")
        if EndpointConfigSelectorEndpointMatcher.to_proto(self.endpoint_matcher):
            resource.endpoint_matcher.CopyFrom(
                EndpointConfigSelectorEndpointMatcher.to_proto(self.endpoint_matcher)
            )
        else:
            resource.ClearField("endpoint_matcher")
        if EndpointConfigSelectorTrafficPortSelector.to_proto(
            self.traffic_port_selector
        ):
            resource.traffic_port_selector.CopyFrom(
                EndpointConfigSelectorTrafficPortSelector.to_proto(
                    self.traffic_port_selector
                )
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


class EndpointConfigSelectorHttpFilters(object):
    def __init__(self, http_filters: list = None):
        self.http_filters = http_filters

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            endpoint_config_selector_pb2.NetworkservicesAlphaEndpointConfigSelectorHttpFilters()
        )
        if Primitive.to_proto(resource.http_filters):
            res.http_filters.extend(Primitive.to_proto(resource.http_filters))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointConfigSelectorHttpFilters(
            http_filters=Primitive.from_proto(resource.http_filters),
        )


class EndpointConfigSelectorHttpFiltersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EndpointConfigSelectorHttpFilters.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EndpointConfigSelectorHttpFilters.from_proto(i) for i in resources]


class EndpointConfigSelectorEndpointMatcher(object):
    def __init__(self, metadata_label_matcher: dict = None):
        self.metadata_label_matcher = metadata_label_matcher

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            endpoint_config_selector_pb2.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcher()
        )
        if EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher.to_proto(
            resource.metadata_label_matcher
        ):
            res.metadata_label_matcher.CopyFrom(
                EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher.to_proto(
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

        return EndpointConfigSelectorEndpointMatcher(
            metadata_label_matcher=EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher.from_proto(
                resource.metadata_label_matcher
            ),
        )


class EndpointConfigSelectorEndpointMatcherArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EndpointConfigSelectorEndpointMatcher.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EndpointConfigSelectorEndpointMatcher.from_proto(i) for i in resources]


class EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(object):
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
            endpoint_config_selector_pb2.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher()
        )
        if EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.to_proto(
            resource.metadata_label_match_criteria
        ):
            res.metadata_label_match_criteria = EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.to_proto(
                resource.metadata_label_match_criteria
            )
        if EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsArray.to_proto(
            resource.metadata_labels
        ):
            res.metadata_labels.extend(
                EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsArray.to_proto(
                    resource.metadata_labels
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(
            metadata_label_match_criteria=EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.from_proto(
                resource.metadata_label_match_criteria
            ),
            metadata_labels=EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsArray.from_proto(
                resource.metadata_labels
            ),
        )


class EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher.from_proto(i)
            for i in resources
        ]


class EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(object):
    def __init__(self, label_name: str = None, label_value: str = None):
        self.label_name = label_name
        self.label_value = label_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            endpoint_config_selector_pb2.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels()
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

        return EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(
            label_name=Primitive.from_proto(resource.label_name),
            label_value=Primitive.from_proto(resource.label_value),
        )


class EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels.from_proto(
                i
            )
            for i in resources
        ]


class EndpointConfigSelectorTrafficPortSelector(object):
    def __init__(self, ports: list = None):
        self.ports = ports

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            endpoint_config_selector_pb2.NetworkservicesAlphaEndpointConfigSelectorTrafficPortSelector()
        )
        if Primitive.to_proto(resource.ports):
            res.ports.extend(Primitive.to_proto(resource.ports))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointConfigSelectorTrafficPortSelector(
            ports=Primitive.from_proto(resource.ports),
        )


class EndpointConfigSelectorTrafficPortSelectorArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EndpointConfigSelectorTrafficPortSelector.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EndpointConfigSelectorTrafficPortSelector.from_proto(i) for i in resources
        ]


class EndpointConfigSelectorTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return endpoint_config_selector_pb2.NetworkservicesAlphaEndpointConfigSelectorTypeEnum.Value(
            "NetworkservicesAlphaEndpointConfigSelectorTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return endpoint_config_selector_pb2.NetworkservicesAlphaEndpointConfigSelectorTypeEnum.Name(
            resource
        )[
            len("NetworkservicesAlphaEndpointConfigSelectorTypeEnum") :
        ]


class EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return endpoint_config_selector_pb2.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.Value(
            "NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return endpoint_config_selector_pb2.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.Name(
            resource
        )[
            len(
                "NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"
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
