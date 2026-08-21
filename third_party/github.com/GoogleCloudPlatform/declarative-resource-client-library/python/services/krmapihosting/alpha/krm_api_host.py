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
from google3.cloud.graphite.mmv2.services.google.krmapihosting import krm_api_host_pb2
from google3.cloud.graphite.mmv2.services.google.krmapihosting import (
    krm_api_host_pb2_grpc,
)

from typing import List


class KrmApiHost(object):
    def __init__(
        self,
        name: str = None,
        labels: dict = None,
        bundles_config: dict = None,
        use_private_endpoint: bool = None,
        gke_resource_link: str = None,
        state: str = None,
        management_config: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.labels = labels
        self.bundles_config = bundles_config
        self.use_private_endpoint = use_private_endpoint
        self.management_config = management_config
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = krm_api_host_pb2_grpc.KrmapihostingAlphaKrmApiHostServiceStub(
            channel.Channel()
        )
        request = krm_api_host_pb2.ApplyKrmapihostingAlphaKrmApiHostRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if KrmApiHostBundlesConfig.to_proto(self.bundles_config):
            request.resource.bundles_config.CopyFrom(
                KrmApiHostBundlesConfig.to_proto(self.bundles_config)
            )
        else:
            request.resource.ClearField("bundles_config")
        if Primitive.to_proto(self.use_private_endpoint):
            request.resource.use_private_endpoint = Primitive.to_proto(
                self.use_private_endpoint
            )

        if KrmApiHostManagementConfig.to_proto(self.management_config):
            request.resource.management_config.CopyFrom(
                KrmApiHostManagementConfig.to_proto(self.management_config)
            )
        else:
            request.resource.ClearField("management_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyKrmapihostingAlphaKrmApiHost(request)
        self.name = Primitive.from_proto(response.name)
        self.labels = Primitive.from_proto(response.labels)
        self.bundles_config = KrmApiHostBundlesConfig.from_proto(
            response.bundles_config
        )
        self.use_private_endpoint = Primitive.from_proto(response.use_private_endpoint)
        self.gke_resource_link = Primitive.from_proto(response.gke_resource_link)
        self.state = KrmApiHostStateEnum.from_proto(response.state)
        self.management_config = KrmApiHostManagementConfig.from_proto(
            response.management_config
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = krm_api_host_pb2_grpc.KrmapihostingAlphaKrmApiHostServiceStub(
            channel.Channel()
        )
        request = krm_api_host_pb2.DeleteKrmapihostingAlphaKrmApiHostRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if KrmApiHostBundlesConfig.to_proto(self.bundles_config):
            request.resource.bundles_config.CopyFrom(
                KrmApiHostBundlesConfig.to_proto(self.bundles_config)
            )
        else:
            request.resource.ClearField("bundles_config")
        if Primitive.to_proto(self.use_private_endpoint):
            request.resource.use_private_endpoint = Primitive.to_proto(
                self.use_private_endpoint
            )

        if KrmApiHostManagementConfig.to_proto(self.management_config):
            request.resource.management_config.CopyFrom(
                KrmApiHostManagementConfig.to_proto(self.management_config)
            )
        else:
            request.resource.ClearField("management_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteKrmapihostingAlphaKrmApiHost(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = krm_api_host_pb2_grpc.KrmapihostingAlphaKrmApiHostServiceStub(
            channel.Channel()
        )
        request = krm_api_host_pb2.ListKrmapihostingAlphaKrmApiHostRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListKrmapihostingAlphaKrmApiHost(request).items

    def to_proto(self):
        resource = krm_api_host_pb2.KrmapihostingAlphaKrmApiHost()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if KrmApiHostBundlesConfig.to_proto(self.bundles_config):
            resource.bundles_config.CopyFrom(
                KrmApiHostBundlesConfig.to_proto(self.bundles_config)
            )
        else:
            resource.ClearField("bundles_config")
        if Primitive.to_proto(self.use_private_endpoint):
            resource.use_private_endpoint = Primitive.to_proto(
                self.use_private_endpoint
            )
        if KrmApiHostManagementConfig.to_proto(self.management_config):
            resource.management_config.CopyFrom(
                KrmApiHostManagementConfig.to_proto(self.management_config)
            )
        else:
            resource.ClearField("management_config")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class KrmApiHostBundlesConfig(object):
    def __init__(self, config_controller_config: dict = None):
        self.config_controller_config = config_controller_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = krm_api_host_pb2.KrmapihostingAlphaKrmApiHostBundlesConfig()
        if KrmApiHostBundlesConfigConfigControllerConfig.to_proto(
            resource.config_controller_config
        ):
            res.config_controller_config.CopyFrom(
                KrmApiHostBundlesConfigConfigControllerConfig.to_proto(
                    resource.config_controller_config
                )
            )
        else:
            res.ClearField("config_controller_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KrmApiHostBundlesConfig(
            config_controller_config=KrmApiHostBundlesConfigConfigControllerConfig.from_proto(
                resource.config_controller_config
            ),
        )


class KrmApiHostBundlesConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KrmApiHostBundlesConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KrmApiHostBundlesConfig.from_proto(i) for i in resources]


class KrmApiHostBundlesConfigConfigControllerConfig(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            krm_api_host_pb2.KrmapihostingAlphaKrmApiHostBundlesConfigConfigControllerConfig()
        )
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KrmApiHostBundlesConfigConfigControllerConfig(
            enabled=Primitive.from_proto(resource.enabled),
        )


class KrmApiHostBundlesConfigConfigControllerConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            KrmApiHostBundlesConfigConfigControllerConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            KrmApiHostBundlesConfigConfigControllerConfig.from_proto(i)
            for i in resources
        ]


class KrmApiHostManagementConfig(object):
    def __init__(self, standard_management_config: dict = None):
        self.standard_management_config = standard_management_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = krm_api_host_pb2.KrmapihostingAlphaKrmApiHostManagementConfig()
        if KrmApiHostManagementConfigStandardManagementConfig.to_proto(
            resource.standard_management_config
        ):
            res.standard_management_config.CopyFrom(
                KrmApiHostManagementConfigStandardManagementConfig.to_proto(
                    resource.standard_management_config
                )
            )
        else:
            res.ClearField("standard_management_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KrmApiHostManagementConfig(
            standard_management_config=KrmApiHostManagementConfigStandardManagementConfig.from_proto(
                resource.standard_management_config
            ),
        )


class KrmApiHostManagementConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KrmApiHostManagementConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KrmApiHostManagementConfig.from_proto(i) for i in resources]


class KrmApiHostManagementConfigStandardManagementConfig(object):
    def __init__(
        self,
        network: str = None,
        master_ipv4_cidr_block: str = None,
        man_block: str = None,
        cluster_cidr_block: str = None,
        services_cidr_block: str = None,
        cluster_named_range: str = None,
        services_named_range: str = None,
    ):
        self.network = network
        self.master_ipv4_cidr_block = master_ipv4_cidr_block
        self.man_block = man_block
        self.cluster_cidr_block = cluster_cidr_block
        self.services_cidr_block = services_cidr_block
        self.cluster_named_range = cluster_named_range
        self.services_named_range = services_named_range

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            krm_api_host_pb2.KrmapihostingAlphaKrmApiHostManagementConfigStandardManagementConfig()
        )
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        if Primitive.to_proto(resource.master_ipv4_cidr_block):
            res.master_ipv4_cidr_block = Primitive.to_proto(
                resource.master_ipv4_cidr_block
            )
        if Primitive.to_proto(resource.man_block):
            res.man_block = Primitive.to_proto(resource.man_block)
        if Primitive.to_proto(resource.cluster_cidr_block):
            res.cluster_cidr_block = Primitive.to_proto(resource.cluster_cidr_block)
        if Primitive.to_proto(resource.services_cidr_block):
            res.services_cidr_block = Primitive.to_proto(resource.services_cidr_block)
        if Primitive.to_proto(resource.cluster_named_range):
            res.cluster_named_range = Primitive.to_proto(resource.cluster_named_range)
        if Primitive.to_proto(resource.services_named_range):
            res.services_named_range = Primitive.to_proto(resource.services_named_range)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KrmApiHostManagementConfigStandardManagementConfig(
            network=Primitive.from_proto(resource.network),
            master_ipv4_cidr_block=Primitive.from_proto(
                resource.master_ipv4_cidr_block
            ),
            man_block=Primitive.from_proto(resource.man_block),
            cluster_cidr_block=Primitive.from_proto(resource.cluster_cidr_block),
            services_cidr_block=Primitive.from_proto(resource.services_cidr_block),
            cluster_named_range=Primitive.from_proto(resource.cluster_named_range),
            services_named_range=Primitive.from_proto(resource.services_named_range),
        )


class KrmApiHostManagementConfigStandardManagementConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            KrmApiHostManagementConfigStandardManagementConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            KrmApiHostManagementConfigStandardManagementConfig.from_proto(i)
            for i in resources
        ]


class KrmApiHostStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return krm_api_host_pb2.KrmapihostingAlphaKrmApiHostStateEnum.Value(
            "KrmapihostingAlphaKrmApiHostStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return krm_api_host_pb2.KrmapihostingAlphaKrmApiHostStateEnum.Name(resource)[
            len("KrmapihostingAlphaKrmApiHostStateEnum") :
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
