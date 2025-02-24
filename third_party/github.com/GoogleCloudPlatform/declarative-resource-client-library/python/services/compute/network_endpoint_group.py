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
from google3.cloud.graphite.mmv2.services.google.compute import (
    network_endpoint_group_pb2,
)
from google3.cloud.graphite.mmv2.services.google.compute import (
    network_endpoint_group_pb2_grpc,
)

from typing import List


class NetworkEndpointGroup(object):
    def __init__(
        self,
        id: int = None,
        self_link: str = None,
        self_link_with_id: str = None,
        name: str = None,
        description: str = None,
        network_endpoint_type: str = None,
        size: int = None,
        location: str = None,
        network: str = None,
        subnetwork: str = None,
        default_port: int = None,
        annotations: dict = None,
        cloud_run: dict = None,
        app_engine: dict = None,
        cloud_function: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.network_endpoint_type = network_endpoint_type
        self.size = size
        self.location = location
        self.network = network
        self.subnetwork = subnetwork
        self.default_port = default_port
        self.annotations = annotations
        self.cloud_run = cloud_run
        self.app_engine = app_engine
        self.cloud_function = cloud_function
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = network_endpoint_group_pb2_grpc.ComputeNetworkEndpointGroupServiceStub(
            channel.Channel()
        )
        request = network_endpoint_group_pb2.ApplyComputeNetworkEndpointGroupRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if NetworkEndpointGroupNetworkEndpointTypeEnum.to_proto(
            self.network_endpoint_type
        ):
            request.resource.network_endpoint_type = NetworkEndpointGroupNetworkEndpointTypeEnum.to_proto(
                self.network_endpoint_type
            )

        if Primitive.to_proto(self.size):
            request.resource.size = Primitive.to_proto(self.size)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.subnetwork):
            request.resource.subnetwork = Primitive.to_proto(self.subnetwork)

        if Primitive.to_proto(self.default_port):
            request.resource.default_port = Primitive.to_proto(self.default_port)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if NetworkEndpointGroupCloudRun.to_proto(self.cloud_run):
            request.resource.cloud_run.CopyFrom(
                NetworkEndpointGroupCloudRun.to_proto(self.cloud_run)
            )
        else:
            request.resource.ClearField("cloud_run")
        if NetworkEndpointGroupAppEngine.to_proto(self.app_engine):
            request.resource.app_engine.CopyFrom(
                NetworkEndpointGroupAppEngine.to_proto(self.app_engine)
            )
        else:
            request.resource.ClearField("app_engine")
        if NetworkEndpointGroupCloudFunction.to_proto(self.cloud_function):
            request.resource.cloud_function.CopyFrom(
                NetworkEndpointGroupCloudFunction.to_proto(self.cloud_function)
            )
        else:
            request.resource.ClearField("cloud_function")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeNetworkEndpointGroup(request)
        self.id = Primitive.from_proto(response.id)
        self.self_link = Primitive.from_proto(response.self_link)
        self.self_link_with_id = Primitive.from_proto(response.self_link_with_id)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.network_endpoint_type = NetworkEndpointGroupNetworkEndpointTypeEnum.from_proto(
            response.network_endpoint_type
        )
        self.size = Primitive.from_proto(response.size)
        self.location = Primitive.from_proto(response.location)
        self.network = Primitive.from_proto(response.network)
        self.subnetwork = Primitive.from_proto(response.subnetwork)
        self.default_port = Primitive.from_proto(response.default_port)
        self.annotations = Primitive.from_proto(response.annotations)
        self.cloud_run = NetworkEndpointGroupCloudRun.from_proto(response.cloud_run)
        self.app_engine = NetworkEndpointGroupAppEngine.from_proto(response.app_engine)
        self.cloud_function = NetworkEndpointGroupCloudFunction.from_proto(
            response.cloud_function
        )
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = network_endpoint_group_pb2_grpc.ComputeNetworkEndpointGroupServiceStub(
            channel.Channel()
        )
        request = network_endpoint_group_pb2.DeleteComputeNetworkEndpointGroupRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if NetworkEndpointGroupNetworkEndpointTypeEnum.to_proto(
            self.network_endpoint_type
        ):
            request.resource.network_endpoint_type = NetworkEndpointGroupNetworkEndpointTypeEnum.to_proto(
                self.network_endpoint_type
            )

        if Primitive.to_proto(self.size):
            request.resource.size = Primitive.to_proto(self.size)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.subnetwork):
            request.resource.subnetwork = Primitive.to_proto(self.subnetwork)

        if Primitive.to_proto(self.default_port):
            request.resource.default_port = Primitive.to_proto(self.default_port)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if NetworkEndpointGroupCloudRun.to_proto(self.cloud_run):
            request.resource.cloud_run.CopyFrom(
                NetworkEndpointGroupCloudRun.to_proto(self.cloud_run)
            )
        else:
            request.resource.ClearField("cloud_run")
        if NetworkEndpointGroupAppEngine.to_proto(self.app_engine):
            request.resource.app_engine.CopyFrom(
                NetworkEndpointGroupAppEngine.to_proto(self.app_engine)
            )
        else:
            request.resource.ClearField("app_engine")
        if NetworkEndpointGroupCloudFunction.to_proto(self.cloud_function):
            request.resource.cloud_function.CopyFrom(
                NetworkEndpointGroupCloudFunction.to_proto(self.cloud_function)
            )
        else:
            request.resource.ClearField("cloud_function")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeNetworkEndpointGroup(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = network_endpoint_group_pb2_grpc.ComputeNetworkEndpointGroupServiceStub(
            channel.Channel()
        )
        request = network_endpoint_group_pb2.ListComputeNetworkEndpointGroupRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeNetworkEndpointGroup(request).items

    def to_proto(self):
        resource = network_endpoint_group_pb2.ComputeNetworkEndpointGroup()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if NetworkEndpointGroupNetworkEndpointTypeEnum.to_proto(
            self.network_endpoint_type
        ):
            resource.network_endpoint_type = NetworkEndpointGroupNetworkEndpointTypeEnum.to_proto(
                self.network_endpoint_type
            )
        if Primitive.to_proto(self.size):
            resource.size = Primitive.to_proto(self.size)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if Primitive.to_proto(self.subnetwork):
            resource.subnetwork = Primitive.to_proto(self.subnetwork)
        if Primitive.to_proto(self.default_port):
            resource.default_port = Primitive.to_proto(self.default_port)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if NetworkEndpointGroupCloudRun.to_proto(self.cloud_run):
            resource.cloud_run.CopyFrom(
                NetworkEndpointGroupCloudRun.to_proto(self.cloud_run)
            )
        else:
            resource.ClearField("cloud_run")
        if NetworkEndpointGroupAppEngine.to_proto(self.app_engine):
            resource.app_engine.CopyFrom(
                NetworkEndpointGroupAppEngine.to_proto(self.app_engine)
            )
        else:
            resource.ClearField("app_engine")
        if NetworkEndpointGroupCloudFunction.to_proto(self.cloud_function):
            resource.cloud_function.CopyFrom(
                NetworkEndpointGroupCloudFunction.to_proto(self.cloud_function)
            )
        else:
            resource.ClearField("cloud_function")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class NetworkEndpointGroupCloudRun(object):
    def __init__(self, service: str = None, tag: str = None, url_mask: str = None):
        self.service = service
        self.tag = tag
        self.url_mask = url_mask

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = network_endpoint_group_pb2.ComputeNetworkEndpointGroupCloudRun()
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.url_mask):
            res.url_mask = Primitive.to_proto(resource.url_mask)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NetworkEndpointGroupCloudRun(
            service=Primitive.from_proto(resource.service),
            tag=Primitive.from_proto(resource.tag),
            url_mask=Primitive.from_proto(resource.url_mask),
        )


class NetworkEndpointGroupCloudRunArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NetworkEndpointGroupCloudRun.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NetworkEndpointGroupCloudRun.from_proto(i) for i in resources]


class NetworkEndpointGroupAppEngine(object):
    def __init__(self, service: str = None, version: str = None, url_mask: str = None):
        self.service = service
        self.version = version
        self.url_mask = url_mask

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = network_endpoint_group_pb2.ComputeNetworkEndpointGroupAppEngine()
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        if Primitive.to_proto(resource.url_mask):
            res.url_mask = Primitive.to_proto(resource.url_mask)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NetworkEndpointGroupAppEngine(
            service=Primitive.from_proto(resource.service),
            version=Primitive.from_proto(resource.version),
            url_mask=Primitive.from_proto(resource.url_mask),
        )


class NetworkEndpointGroupAppEngineArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NetworkEndpointGroupAppEngine.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NetworkEndpointGroupAppEngine.from_proto(i) for i in resources]


class NetworkEndpointGroupCloudFunction(object):
    def __init__(self, function: str = None, url_mask: str = None):
        self.function = function
        self.url_mask = url_mask

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = network_endpoint_group_pb2.ComputeNetworkEndpointGroupCloudFunction()
        if Primitive.to_proto(resource.function):
            res.function = Primitive.to_proto(resource.function)
        if Primitive.to_proto(resource.url_mask):
            res.url_mask = Primitive.to_proto(resource.url_mask)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NetworkEndpointGroupCloudFunction(
            function=Primitive.from_proto(resource.function),
            url_mask=Primitive.from_proto(resource.url_mask),
        )


class NetworkEndpointGroupCloudFunctionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NetworkEndpointGroupCloudFunction.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NetworkEndpointGroupCloudFunction.from_proto(i) for i in resources]


class NetworkEndpointGroupNetworkEndpointTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return network_endpoint_group_pb2.ComputeNetworkEndpointGroupNetworkEndpointTypeEnum.Value(
            "ComputeNetworkEndpointGroupNetworkEndpointTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return network_endpoint_group_pb2.ComputeNetworkEndpointGroupNetworkEndpointTypeEnum.Name(
            resource
        )[
            len("ComputeNetworkEndpointGroupNetworkEndpointTypeEnum") :
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
