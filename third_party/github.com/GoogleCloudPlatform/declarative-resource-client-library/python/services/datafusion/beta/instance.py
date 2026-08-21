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
from google3.cloud.graphite.mmv2.services.google.data_fusion import instance_pb2
from google3.cloud.graphite.mmv2.services.google.data_fusion import instance_pb2_grpc

from typing import List


class Instance(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        type: str = None,
        enable_stackdriver_logging: bool = None,
        enable_stackdriver_monitoring: bool = None,
        private_instance: bool = None,
        network_config: dict = None,
        labels: dict = None,
        options: dict = None,
        create_time: str = None,
        update_time: str = None,
        state: str = None,
        state_message: str = None,
        service_endpoint: str = None,
        zone: str = None,
        version: str = None,
        display_name: str = None,
        available_version: list = None,
        api_endpoint: str = None,
        gcs_bucket: str = None,
        p4_service_account: str = None,
        tenant_project_id: str = None,
        dataproc_service_account: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.type = type
        self.enable_stackdriver_logging = enable_stackdriver_logging
        self.enable_stackdriver_monitoring = enable_stackdriver_monitoring
        self.private_instance = private_instance
        self.network_config = network_config
        self.labels = labels
        self.options = options
        self.zone = zone
        self.version = version
        self.display_name = display_name
        self.dataproc_service_account = dataproc_service_account
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_pb2_grpc.DatafusionBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ApplyDatafusionBetaInstanceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if InstanceTypeEnum.to_proto(self.type):
            request.resource.type = InstanceTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.enable_stackdriver_logging):
            request.resource.enable_stackdriver_logging = Primitive.to_proto(
                self.enable_stackdriver_logging
            )

        if Primitive.to_proto(self.enable_stackdriver_monitoring):
            request.resource.enable_stackdriver_monitoring = Primitive.to_proto(
                self.enable_stackdriver_monitoring
            )

        if Primitive.to_proto(self.private_instance):
            request.resource.private_instance = Primitive.to_proto(
                self.private_instance
            )

        if InstanceNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                InstanceNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.options):
            request.resource.options = Primitive.to_proto(self.options)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.dataproc_service_account):
            request.resource.dataproc_service_account = Primitive.to_proto(
                self.dataproc_service_account
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDatafusionBetaInstance(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.type = InstanceTypeEnum.from_proto(response.type)
        self.enable_stackdriver_logging = Primitive.from_proto(
            response.enable_stackdriver_logging
        )
        self.enable_stackdriver_monitoring = Primitive.from_proto(
            response.enable_stackdriver_monitoring
        )
        self.private_instance = Primitive.from_proto(response.private_instance)
        self.network_config = InstanceNetworkConfig.from_proto(response.network_config)
        self.labels = Primitive.from_proto(response.labels)
        self.options = Primitive.from_proto(response.options)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.state = InstanceStateEnum.from_proto(response.state)
        self.state_message = Primitive.from_proto(response.state_message)
        self.service_endpoint = Primitive.from_proto(response.service_endpoint)
        self.zone = Primitive.from_proto(response.zone)
        self.version = Primitive.from_proto(response.version)
        self.display_name = Primitive.from_proto(response.display_name)
        self.available_version = InstanceAvailableVersionArray.from_proto(
            response.available_version
        )
        self.api_endpoint = Primitive.from_proto(response.api_endpoint)
        self.gcs_bucket = Primitive.from_proto(response.gcs_bucket)
        self.p4_service_account = Primitive.from_proto(response.p4_service_account)
        self.tenant_project_id = Primitive.from_proto(response.tenant_project_id)
        self.dataproc_service_account = Primitive.from_proto(
            response.dataproc_service_account
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = instance_pb2_grpc.DatafusionBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.DeleteDatafusionBetaInstanceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if InstanceTypeEnum.to_proto(self.type):
            request.resource.type = InstanceTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.enable_stackdriver_logging):
            request.resource.enable_stackdriver_logging = Primitive.to_proto(
                self.enable_stackdriver_logging
            )

        if Primitive.to_proto(self.enable_stackdriver_monitoring):
            request.resource.enable_stackdriver_monitoring = Primitive.to_proto(
                self.enable_stackdriver_monitoring
            )

        if Primitive.to_proto(self.private_instance):
            request.resource.private_instance = Primitive.to_proto(
                self.private_instance
            )

        if InstanceNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                InstanceNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.options):
            request.resource.options = Primitive.to_proto(self.options)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.dataproc_service_account):
            request.resource.dataproc_service_account = Primitive.to_proto(
                self.dataproc_service_account
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteDatafusionBetaInstance(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = instance_pb2_grpc.DatafusionBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ListDatafusionBetaInstanceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListDatafusionBetaInstance(request).items

    def to_proto(self):
        resource = instance_pb2.DatafusionBetaInstance()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if InstanceTypeEnum.to_proto(self.type):
            resource.type = InstanceTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.enable_stackdriver_logging):
            resource.enable_stackdriver_logging = Primitive.to_proto(
                self.enable_stackdriver_logging
            )
        if Primitive.to_proto(self.enable_stackdriver_monitoring):
            resource.enable_stackdriver_monitoring = Primitive.to_proto(
                self.enable_stackdriver_monitoring
            )
        if Primitive.to_proto(self.private_instance):
            resource.private_instance = Primitive.to_proto(self.private_instance)
        if InstanceNetworkConfig.to_proto(self.network_config):
            resource.network_config.CopyFrom(
                InstanceNetworkConfig.to_proto(self.network_config)
            )
        else:
            resource.ClearField("network_config")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.options):
            resource.options = Primitive.to_proto(self.options)
        if Primitive.to_proto(self.zone):
            resource.zone = Primitive.to_proto(self.zone)
        if Primitive.to_proto(self.version):
            resource.version = Primitive.to_proto(self.version)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.dataproc_service_account):
            resource.dataproc_service_account = Primitive.to_proto(
                self.dataproc_service_account
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class InstanceNetworkConfig(object):
    def __init__(self, network: str = None, ip_allocation: str = None):
        self.network = network
        self.ip_allocation = ip_allocation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.DatafusionBetaInstanceNetworkConfig()
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        if Primitive.to_proto(resource.ip_allocation):
            res.ip_allocation = Primitive.to_proto(resource.ip_allocation)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNetworkConfig(
            network=Primitive.from_proto(resource.network),
            ip_allocation=Primitive.from_proto(resource.ip_allocation),
        )


class InstanceNetworkConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceNetworkConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceNetworkConfig.from_proto(i) for i in resources]


class InstanceAvailableVersion(object):
    def __init__(
        self,
        version_number: str = None,
        default_version: bool = None,
        available_features: list = None,
    ):
        self.version_number = version_number
        self.default_version = default_version
        self.available_features = available_features

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.DatafusionBetaInstanceAvailableVersion()
        if Primitive.to_proto(resource.version_number):
            res.version_number = Primitive.to_proto(resource.version_number)
        if Primitive.to_proto(resource.default_version):
            res.default_version = Primitive.to_proto(resource.default_version)
        if Primitive.to_proto(resource.available_features):
            res.available_features.extend(
                Primitive.to_proto(resource.available_features)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceAvailableVersion(
            version_number=Primitive.from_proto(resource.version_number),
            default_version=Primitive.from_proto(resource.default_version),
            available_features=Primitive.from_proto(resource.available_features),
        )


class InstanceAvailableVersionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceAvailableVersion.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceAvailableVersion.from_proto(i) for i in resources]


class InstanceTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.DatafusionBetaInstanceTypeEnum.Value(
            "DatafusionBetaInstanceTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.DatafusionBetaInstanceTypeEnum.Name(resource)[
            len("DatafusionBetaInstanceTypeEnum") :
        ]


class InstanceStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.DatafusionBetaInstanceStateEnum.Value(
            "DatafusionBetaInstanceStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.DatafusionBetaInstanceStateEnum.Name(resource)[
            len("DatafusionBetaInstanceStateEnum") :
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
