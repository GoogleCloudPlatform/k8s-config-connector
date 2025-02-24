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
from google3.cloud.graphite.mmv2.services.google.network_connectivity import hub_pb2
from google3.cloud.graphite.mmv2.services.google.network_connectivity import (
    hub_pb2_grpc,
)

from typing import List


class Hub(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        description: str = None,
        unique_id: str = None,
        state: str = None,
        project: str = None,
        routing_vpcs: list = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.labels = labels
        self.description = description
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = hub_pb2_grpc.NetworkconnectivityHubServiceStub(channel.Channel())
        request = hub_pb2.ApplyNetworkconnectivityHubRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworkconnectivityHub(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.description = Primitive.from_proto(response.description)
        self.unique_id = Primitive.from_proto(response.unique_id)
        self.state = HubStateEnum.from_proto(response.state)
        self.project = Primitive.from_proto(response.project)
        self.routing_vpcs = HubRoutingVpcsArray.from_proto(response.routing_vpcs)

    def delete(self):
        stub = hub_pb2_grpc.NetworkconnectivityHubServiceStub(channel.Channel())
        request = hub_pb2.DeleteNetworkconnectivityHubRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteNetworkconnectivityHub(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = hub_pb2_grpc.NetworkconnectivityHubServiceStub(channel.Channel())
        request = hub_pb2.ListNetworkconnectivityHubRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListNetworkconnectivityHub(request).items

    def to_proto(self):
        resource = hub_pb2.NetworkconnectivityHub()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class HubRoutingVpcs(object):
    def __init__(self, uri: str = None):
        self.uri = uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = hub_pb2.NetworkconnectivityHubRoutingVpcs()
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HubRoutingVpcs(
            uri=Primitive.from_proto(resource.uri),
        )


class HubRoutingVpcsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HubRoutingVpcs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HubRoutingVpcs.from_proto(i) for i in resources]


class HubStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return hub_pb2.NetworkconnectivityHubStateEnum.Value(
            "NetworkconnectivityHubStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return hub_pb2.NetworkconnectivityHubStateEnum.Name(resource)[
            len("NetworkconnectivityHubStateEnum") :
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
