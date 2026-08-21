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
from google3.cloud.graphite.mmv2.services.google.vmwareengine import private_cloud_pb2
from google3.cloud.graphite.mmv2.services.google.vmwareengine import (
    private_cloud_pb2_grpc,
)

from typing import List


class PrivateCloud(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        delete_time: str = None,
        expire_time: str = None,
        labels: dict = None,
        state: str = None,
        network_config: dict = None,
        management_cluster: dict = None,
        description: str = None,
        conditions: list = None,
        hcx: dict = None,
        nsx: dict = None,
        vcenter: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.labels = labels
        self.network_config = network_config
        self.management_cluster = management_cluster
        self.description = description
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = private_cloud_pb2_grpc.VmwareengineAlphaPrivateCloudServiceStub(
            channel.Channel()
        )
        request = private_cloud_pb2.ApplyVmwareengineAlphaPrivateCloudRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if PrivateCloudNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                PrivateCloudNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if PrivateCloudManagementCluster.to_proto(self.management_cluster):
            request.resource.management_cluster.CopyFrom(
                PrivateCloudManagementCluster.to_proto(self.management_cluster)
            )
        else:
            request.resource.ClearField("management_cluster")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyVmwareengineAlphaPrivateCloud(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.expire_time = Primitive.from_proto(response.expire_time)
        self.labels = Primitive.from_proto(response.labels)
        self.state = PrivateCloudStateEnum.from_proto(response.state)
        self.network_config = PrivateCloudNetworkConfig.from_proto(
            response.network_config
        )
        self.management_cluster = PrivateCloudManagementCluster.from_proto(
            response.management_cluster
        )
        self.description = Primitive.from_proto(response.description)
        self.conditions = PrivateCloudConditionsArray.from_proto(response.conditions)
        self.hcx = PrivateCloudHcx.from_proto(response.hcx)
        self.nsx = PrivateCloudNsx.from_proto(response.nsx)
        self.vcenter = PrivateCloudVcenter.from_proto(response.vcenter)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = private_cloud_pb2_grpc.VmwareengineAlphaPrivateCloudServiceStub(
            channel.Channel()
        )
        request = private_cloud_pb2.DeleteVmwareengineAlphaPrivateCloudRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if PrivateCloudNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                PrivateCloudNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if PrivateCloudManagementCluster.to_proto(self.management_cluster):
            request.resource.management_cluster.CopyFrom(
                PrivateCloudManagementCluster.to_proto(self.management_cluster)
            )
        else:
            request.resource.ClearField("management_cluster")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteVmwareengineAlphaPrivateCloud(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = private_cloud_pb2_grpc.VmwareengineAlphaPrivateCloudServiceStub(
            channel.Channel()
        )
        request = private_cloud_pb2.ListVmwareengineAlphaPrivateCloudRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListVmwareengineAlphaPrivateCloud(request).items

    def to_proto(self):
        resource = private_cloud_pb2.VmwareengineAlphaPrivateCloud()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if PrivateCloudNetworkConfig.to_proto(self.network_config):
            resource.network_config.CopyFrom(
                PrivateCloudNetworkConfig.to_proto(self.network_config)
            )
        else:
            resource.ClearField("network_config")
        if PrivateCloudManagementCluster.to_proto(self.management_cluster):
            resource.management_cluster.CopyFrom(
                PrivateCloudManagementCluster.to_proto(self.management_cluster)
            )
        else:
            resource.ClearField("management_cluster")
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class PrivateCloudNetworkConfig(object):
    def __init__(
        self,
        network: str = None,
        service_network: str = None,
        management_cidr: str = None,
    ):
        self.network = network
        self.service_network = service_network
        self.management_cidr = management_cidr

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = private_cloud_pb2.VmwareengineAlphaPrivateCloudNetworkConfig()
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        if Primitive.to_proto(resource.service_network):
            res.service_network = Primitive.to_proto(resource.service_network)
        if Primitive.to_proto(resource.management_cidr):
            res.management_cidr = Primitive.to_proto(resource.management_cidr)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PrivateCloudNetworkConfig(
            network=Primitive.from_proto(resource.network),
            service_network=Primitive.from_proto(resource.service_network),
            management_cidr=Primitive.from_proto(resource.management_cidr),
        )


class PrivateCloudNetworkConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PrivateCloudNetworkConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PrivateCloudNetworkConfig.from_proto(i) for i in resources]


class PrivateCloudManagementCluster(object):
    def __init__(
        self, cluster_id: str = None, node_type_id: str = None, node_count: int = None
    ):
        self.cluster_id = cluster_id
        self.node_type_id = node_type_id
        self.node_count = node_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = private_cloud_pb2.VmwareengineAlphaPrivateCloudManagementCluster()
        if Primitive.to_proto(resource.cluster_id):
            res.cluster_id = Primitive.to_proto(resource.cluster_id)
        if Primitive.to_proto(resource.node_type_id):
            res.node_type_id = Primitive.to_proto(resource.node_type_id)
        if Primitive.to_proto(resource.node_count):
            res.node_count = Primitive.to_proto(resource.node_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PrivateCloudManagementCluster(
            cluster_id=Primitive.from_proto(resource.cluster_id),
            node_type_id=Primitive.from_proto(resource.node_type_id),
            node_count=Primitive.from_proto(resource.node_count),
        )


class PrivateCloudManagementClusterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PrivateCloudManagementCluster.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PrivateCloudManagementCluster.from_proto(i) for i in resources]


class PrivateCloudConditions(object):
    def __init__(self, code: str = None, message: str = None):
        self.code = code
        self.message = message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = private_cloud_pb2.VmwareengineAlphaPrivateCloudConditions()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PrivateCloudConditions(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
        )


class PrivateCloudConditionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PrivateCloudConditions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PrivateCloudConditions.from_proto(i) for i in resources]


class PrivateCloudHcx(object):
    def __init__(
        self,
        fdqn: str = None,
        internal_ip: str = None,
        external_ip: str = None,
        version: str = None,
    ):
        self.fdqn = fdqn
        self.internal_ip = internal_ip
        self.external_ip = external_ip
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = private_cloud_pb2.VmwareengineAlphaPrivateCloudHcx()
        if Primitive.to_proto(resource.fdqn):
            res.fdqn = Primitive.to_proto(resource.fdqn)
        if Primitive.to_proto(resource.internal_ip):
            res.internal_ip = Primitive.to_proto(resource.internal_ip)
        if Primitive.to_proto(resource.external_ip):
            res.external_ip = Primitive.to_proto(resource.external_ip)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PrivateCloudHcx(
            fdqn=Primitive.from_proto(resource.fdqn),
            internal_ip=Primitive.from_proto(resource.internal_ip),
            external_ip=Primitive.from_proto(resource.external_ip),
            version=Primitive.from_proto(resource.version),
        )


class PrivateCloudHcxArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PrivateCloudHcx.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PrivateCloudHcx.from_proto(i) for i in resources]


class PrivateCloudNsx(object):
    def __init__(
        self,
        fdqn: str = None,
        internal_ip: str = None,
        external_ip: str = None,
        version: str = None,
    ):
        self.fdqn = fdqn
        self.internal_ip = internal_ip
        self.external_ip = external_ip
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = private_cloud_pb2.VmwareengineAlphaPrivateCloudNsx()
        if Primitive.to_proto(resource.fdqn):
            res.fdqn = Primitive.to_proto(resource.fdqn)
        if Primitive.to_proto(resource.internal_ip):
            res.internal_ip = Primitive.to_proto(resource.internal_ip)
        if Primitive.to_proto(resource.external_ip):
            res.external_ip = Primitive.to_proto(resource.external_ip)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PrivateCloudNsx(
            fdqn=Primitive.from_proto(resource.fdqn),
            internal_ip=Primitive.from_proto(resource.internal_ip),
            external_ip=Primitive.from_proto(resource.external_ip),
            version=Primitive.from_proto(resource.version),
        )


class PrivateCloudNsxArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PrivateCloudNsx.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PrivateCloudNsx.from_proto(i) for i in resources]


class PrivateCloudVcenter(object):
    def __init__(
        self,
        fdqn: str = None,
        internal_ip: str = None,
        external_ip: str = None,
        version: str = None,
    ):
        self.fdqn = fdqn
        self.internal_ip = internal_ip
        self.external_ip = external_ip
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = private_cloud_pb2.VmwareengineAlphaPrivateCloudVcenter()
        if Primitive.to_proto(resource.fdqn):
            res.fdqn = Primitive.to_proto(resource.fdqn)
        if Primitive.to_proto(resource.internal_ip):
            res.internal_ip = Primitive.to_proto(resource.internal_ip)
        if Primitive.to_proto(resource.external_ip):
            res.external_ip = Primitive.to_proto(resource.external_ip)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PrivateCloudVcenter(
            fdqn=Primitive.from_proto(resource.fdqn),
            internal_ip=Primitive.from_proto(resource.internal_ip),
            external_ip=Primitive.from_proto(resource.external_ip),
            version=Primitive.from_proto(resource.version),
        )


class PrivateCloudVcenterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PrivateCloudVcenter.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PrivateCloudVcenter.from_proto(i) for i in resources]


class PrivateCloudStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return private_cloud_pb2.VmwareengineAlphaPrivateCloudStateEnum.Value(
            "VmwareengineAlphaPrivateCloudStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return private_cloud_pb2.VmwareengineAlphaPrivateCloudStateEnum.Name(resource)[
            len("VmwareengineAlphaPrivateCloudStateEnum") :
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
