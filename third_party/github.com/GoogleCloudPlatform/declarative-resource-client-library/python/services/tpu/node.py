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
from google3.cloud.graphite.mmv2.services.google.tpu import node_pb2
from google3.cloud.graphite.mmv2.services.google.tpu import node_pb2_grpc

from typing import List


class Node(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        accelerator_type: str = None,
        ip_address: str = None,
        port: str = None,
        state: str = None,
        health_description: str = None,
        tensorflow_version: str = None,
        network: str = None,
        cidr_block: str = None,
        service_account: str = None,
        create_time: dict = None,
        scheduling_config: dict = None,
        network_endpoints: list = None,
        health: str = None,
        labels: dict = None,
        use_service_networking: bool = None,
        symptoms: list = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.accelerator_type = accelerator_type
        self.tensorflow_version = tensorflow_version
        self.network = network
        self.cidr_block = cidr_block
        self.scheduling_config = scheduling_config
        self.labels = labels
        self.use_service_networking = use_service_networking
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = node_pb2_grpc.TPUNodeServiceStub(channel.Channel())
        request = node_pb2.ApplyTPUNodeRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.accelerator_type):
            request.resource.accelerator_type = Primitive.to_proto(
                self.accelerator_type
            )

        if Primitive.to_proto(self.tensorflow_version):
            request.resource.tensorflow_version = Primitive.to_proto(
                self.tensorflow_version
            )

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.cidr_block):
            request.resource.cidr_block = Primitive.to_proto(self.cidr_block)

        if NodeSchedulingConfig.to_proto(self.scheduling_config):
            request.resource.scheduling_config.CopyFrom(
                NodeSchedulingConfig.to_proto(self.scheduling_config)
            )
        else:
            request.resource.ClearField("scheduling_config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.use_service_networking):
            request.resource.use_service_networking = Primitive.to_proto(
                self.use_service_networking
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyTPUNode(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.accelerator_type = Primitive.from_proto(response.accelerator_type)
        self.ip_address = Primitive.from_proto(response.ip_address)
        self.port = Primitive.from_proto(response.port)
        self.state = NodeStateEnum.from_proto(response.state)
        self.health_description = Primitive.from_proto(response.health_description)
        self.tensorflow_version = Primitive.from_proto(response.tensorflow_version)
        self.network = Primitive.from_proto(response.network)
        self.cidr_block = Primitive.from_proto(response.cidr_block)
        self.service_account = Primitive.from_proto(response.service_account)
        self.create_time = NodeCreateTime.from_proto(response.create_time)
        self.scheduling_config = NodeSchedulingConfig.from_proto(
            response.scheduling_config
        )
        self.network_endpoints = NodeNetworkEndpointsArray.from_proto(
            response.network_endpoints
        )
        self.health = NodeHealthEnum.from_proto(response.health)
        self.labels = Primitive.from_proto(response.labels)
        self.use_service_networking = Primitive.from_proto(
            response.use_service_networking
        )
        self.symptoms = NodeSymptomsArray.from_proto(response.symptoms)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = node_pb2_grpc.TPUNodeServiceStub(channel.Channel())
        request = node_pb2.DeleteTPUNodeRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.accelerator_type):
            request.resource.accelerator_type = Primitive.to_proto(
                self.accelerator_type
            )

        if Primitive.to_proto(self.tensorflow_version):
            request.resource.tensorflow_version = Primitive.to_proto(
                self.tensorflow_version
            )

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.cidr_block):
            request.resource.cidr_block = Primitive.to_proto(self.cidr_block)

        if NodeSchedulingConfig.to_proto(self.scheduling_config):
            request.resource.scheduling_config.CopyFrom(
                NodeSchedulingConfig.to_proto(self.scheduling_config)
            )
        else:
            request.resource.ClearField("scheduling_config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.use_service_networking):
            request.resource.use_service_networking = Primitive.to_proto(
                self.use_service_networking
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteTPUNode(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = node_pb2_grpc.TPUNodeServiceStub(channel.Channel())
        request = node_pb2.ListTPUNodeRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListTPUNode(request).items

    def to_proto(self):
        resource = node_pb2.TPUNode()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.accelerator_type):
            resource.accelerator_type = Primitive.to_proto(self.accelerator_type)
        if Primitive.to_proto(self.tensorflow_version):
            resource.tensorflow_version = Primitive.to_proto(self.tensorflow_version)
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if Primitive.to_proto(self.cidr_block):
            resource.cidr_block = Primitive.to_proto(self.cidr_block)
        if NodeSchedulingConfig.to_proto(self.scheduling_config):
            resource.scheduling_config.CopyFrom(
                NodeSchedulingConfig.to_proto(self.scheduling_config)
            )
        else:
            resource.ClearField("scheduling_config")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.use_service_networking):
            resource.use_service_networking = Primitive.to_proto(
                self.use_service_networking
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class NodeCreateTime(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pb2.TPUNodeCreateTime()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodeCreateTime(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class NodeCreateTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodeCreateTime.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodeCreateTime.from_proto(i) for i in resources]


class NodeSchedulingConfig(object):
    def __init__(self, preemptible: bool = None, reserved: bool = None):
        self.preemptible = preemptible
        self.reserved = reserved

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pb2.TPUNodeSchedulingConfig()
        if Primitive.to_proto(resource.preemptible):
            res.preemptible = Primitive.to_proto(resource.preemptible)
        if Primitive.to_proto(resource.reserved):
            res.reserved = Primitive.to_proto(resource.reserved)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodeSchedulingConfig(
            preemptible=Primitive.from_proto(resource.preemptible),
            reserved=Primitive.from_proto(resource.reserved),
        )


class NodeSchedulingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodeSchedulingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodeSchedulingConfig.from_proto(i) for i in resources]


class NodeNetworkEndpoints(object):
    def __init__(self, ip_address: str = None, port: int = None):
        self.ip_address = ip_address
        self.port = port

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pb2.TPUNodeNetworkEndpoints()
        if Primitive.to_proto(resource.ip_address):
            res.ip_address = Primitive.to_proto(resource.ip_address)
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodeNetworkEndpoints(
            ip_address=Primitive.from_proto(resource.ip_address),
            port=Primitive.from_proto(resource.port),
        )


class NodeNetworkEndpointsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodeNetworkEndpoints.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodeNetworkEndpoints.from_proto(i) for i in resources]


class NodeSymptoms(object):
    def __init__(
        self,
        create_time: dict = None,
        symptom_type: str = None,
        details: str = None,
        worker_id: str = None,
    ):
        self.create_time = create_time
        self.symptom_type = symptom_type
        self.details = details
        self.worker_id = worker_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pb2.TPUNodeSymptoms()
        if NodeSymptomsCreateTime.to_proto(resource.create_time):
            res.create_time.CopyFrom(
                NodeSymptomsCreateTime.to_proto(resource.create_time)
            )
        else:
            res.ClearField("create_time")
        if NodeSymptomsSymptomTypeEnum.to_proto(resource.symptom_type):
            res.symptom_type = NodeSymptomsSymptomTypeEnum.to_proto(
                resource.symptom_type
            )
        if Primitive.to_proto(resource.details):
            res.details = Primitive.to_proto(resource.details)
        if Primitive.to_proto(resource.worker_id):
            res.worker_id = Primitive.to_proto(resource.worker_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodeSymptoms(
            create_time=NodeSymptomsCreateTime.from_proto(resource.create_time),
            symptom_type=NodeSymptomsSymptomTypeEnum.from_proto(resource.symptom_type),
            details=Primitive.from_proto(resource.details),
            worker_id=Primitive.from_proto(resource.worker_id),
        )


class NodeSymptomsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodeSymptoms.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodeSymptoms.from_proto(i) for i in resources]


class NodeSymptomsCreateTime(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = node_pb2.TPUNodeSymptomsCreateTime()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NodeSymptomsCreateTime(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class NodeSymptomsCreateTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NodeSymptomsCreateTime.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NodeSymptomsCreateTime.from_proto(i) for i in resources]


class NodeStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return node_pb2.TPUNodeStateEnum.Value("TPUNodeStateEnum%s" % resource)

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return node_pb2.TPUNodeStateEnum.Name(resource)[len("TPUNodeStateEnum") :]


class NodeHealthEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return node_pb2.TPUNodeHealthEnum.Value("TPUNodeHealthEnum%s" % resource)

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return node_pb2.TPUNodeHealthEnum.Name(resource)[len("TPUNodeHealthEnum") :]


class NodeSymptomsSymptomTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return node_pb2.TPUNodeSymptomsSymptomTypeEnum.Value(
            "TPUNodeSymptomsSymptomTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return node_pb2.TPUNodeSymptomsSymptomTypeEnum.Name(resource)[
            len("TPUNodeSymptomsSymptomTypeEnum") :
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
