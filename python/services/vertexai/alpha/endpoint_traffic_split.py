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
from google3.cloud.graphite.mmv2.services.google.vertex_ai import (
    endpoint_traffic_split_pb2,
)
from google3.cloud.graphite.mmv2.services.google.vertex_ai import (
    endpoint_traffic_split_pb2_grpc,
)

from typing import List


class EndpointTrafficSplit(object):
    def __init__(
        self,
        endpoint: str = None,
        project: str = None,
        location: str = None,
        etag: str = None,
        traffic_split: list = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.endpoint = endpoint
        self.project = project
        self.location = location
        self.traffic_split = traffic_split
        self.service_account_file = service_account_file

    def apply(self):
        stub = endpoint_traffic_split_pb2_grpc.VertexaiAlphaEndpointTrafficSplitServiceStub(
            channel.Channel()
        )
        request = (
            endpoint_traffic_split_pb2.ApplyVertexaiAlphaEndpointTrafficSplitRequest()
        )
        if Primitive.to_proto(self.endpoint):
            request.resource.endpoint = Primitive.to_proto(self.endpoint)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if EndpointTrafficSplitTrafficSplitArray.to_proto(self.traffic_split):
            request.resource.traffic_split.extend(
                EndpointTrafficSplitTrafficSplitArray.to_proto(self.traffic_split)
            )
        request.service_account_file = self.service_account_file

        response = stub.ApplyVertexaiAlphaEndpointTrafficSplit(request)
        self.endpoint = Primitive.from_proto(response.endpoint)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.etag = Primitive.from_proto(response.etag)
        self.traffic_split = EndpointTrafficSplitTrafficSplitArray.from_proto(
            response.traffic_split
        )

    def delete(self):
        stub = endpoint_traffic_split_pb2_grpc.VertexaiAlphaEndpointTrafficSplitServiceStub(
            channel.Channel()
        )
        request = (
            endpoint_traffic_split_pb2.DeleteVertexaiAlphaEndpointTrafficSplitRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.endpoint):
            request.resource.endpoint = Primitive.to_proto(self.endpoint)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if EndpointTrafficSplitTrafficSplitArray.to_proto(self.traffic_split):
            request.resource.traffic_split.extend(
                EndpointTrafficSplitTrafficSplitArray.to_proto(self.traffic_split)
            )
        response = stub.DeleteVertexaiAlphaEndpointTrafficSplit(request)

    @classmethod
    def list(self, location, service_account_file=""):
        stub = endpoint_traffic_split_pb2_grpc.VertexaiAlphaEndpointTrafficSplitServiceStub(
            channel.Channel()
        )
        request = (
            endpoint_traffic_split_pb2.ListVertexaiAlphaEndpointTrafficSplitRequest()
        )
        request.service_account_file = service_account_file
        request.Location = location

        return stub.ListVertexaiAlphaEndpointTrafficSplit(request).items

    def to_proto(self):
        resource = endpoint_traffic_split_pb2.VertexaiAlphaEndpointTrafficSplit()
        if Primitive.to_proto(self.endpoint):
            resource.endpoint = Primitive.to_proto(self.endpoint)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if EndpointTrafficSplitTrafficSplitArray.to_proto(self.traffic_split):
            resource.traffic_split.extend(
                EndpointTrafficSplitTrafficSplitArray.to_proto(self.traffic_split)
            )
        return resource


class EndpointTrafficSplitTrafficSplit(object):
    def __init__(self, deployed_model_id: str = None, traffic_percentage: int = None):
        self.deployed_model_id = deployed_model_id
        self.traffic_percentage = traffic_percentage

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = endpoint_traffic_split_pb2.VertexaiAlphaEndpointTrafficSplitTrafficSplit()
        if Primitive.to_proto(resource.deployed_model_id):
            res.deployed_model_id = Primitive.to_proto(resource.deployed_model_id)
        if Primitive.to_proto(resource.traffic_percentage):
            res.traffic_percentage = Primitive.to_proto(resource.traffic_percentage)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EndpointTrafficSplitTrafficSplit(
            deployed_model_id=Primitive.from_proto(resource.deployed_model_id),
            traffic_percentage=Primitive.from_proto(resource.traffic_percentage),
        )


class EndpointTrafficSplitTrafficSplitArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EndpointTrafficSplitTrafficSplit.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EndpointTrafficSplitTrafficSplit.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
