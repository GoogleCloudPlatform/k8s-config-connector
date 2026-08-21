# Copyright 2023 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.vmware import private_cloud_pb2
from google3.cloud.graphite.mmv2.services.google.vmware import private_cloud_pb2_grpc

from typing import List


class PrivateCloud(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        delete_time: str = None,
        expire_time: str = None,
        state: str = None,
        network_config: dict = None,
        management_cluster: dict = None,
        description: str = None,
        hcx: dict = None,
        nsx: dict = None,
        vcenter: dict = None,
        uid: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.network_config = network_config
        self.management_cluster = management_cluster
        self.description = description
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = private_cloud_pb2_grpc.VmwareAlphaPrivateCloudServiceStub(
            channel.Channel()
        )
        request = private_cloud_pb2.ApplyVmwareAlphaPrivateCloudRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

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

        response = stub.ApplyVmwareAlphaPrivateCloud(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.expire_time = Primitive.from_proto(response.expire_time)
        self.state = PrivateCloudStateEnum.from_proto(response.state)
        self.network_config = PrivateCloudNetworkConfig.from_proto(
            response.network_config
        )
        self.management_cluster = PrivateCloudManagementCluster.from_proto(
            response.management_cluster
        )
        self.description = Primitive.from_proto(response.description)
        self.hcx = PrivateCloudHcx.from_proto(response.hcx)
        self.nsx = PrivateCloudNsx.from_proto(response.nsx)
        self.vcenter = PrivateCloudVcenter.from_proto(response.vcenter)
        self.uid = Primitive.from_proto(response.uid)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = private_cloud_pb2_grpc.VmwareAlphaPrivateCloudServiceStub(
            channel.Channel()
        )
        request = private_cloud_pb2.DeleteVmwareAlphaPrivateCloudRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

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

        response = stub.DeleteVmwareAlphaPrivateCloud(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = private_cloud_pb2_grpc.VmwareAlphaPrivateCloudServiceStub(
            channel.Channel()
        )
        request = private_cloud_pb2.ListVmwareAlphaPrivateCloudRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListVmwareAlphaPrivateCloud(request).items

    def to_proto(self):
        resource = private_cloud_pb2.VmwareAlphaPrivateCloud()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
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
        management_cidr: str = None,
        vmware_engine_network: str = None,
        vmware_engine_network_canonical: str = None,
        management_ip_address_layout_version: int = None,
    ):
        self.management_cidr = management_cidr
        self.vmware_engine_network = vmware_engine_network
        self.vmware_engine_network_canonical = vmware_engine_network_canonical
        self.management_ip_address_layout_version = management_ip_address_layout_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = private_cloud_pb2.VmwareAlphaPrivateCloudNetworkConfig()
        if Primitive.to_proto(resource.management_cidr):
            res.management_cidr = Primitive.to_proto(resource.management_cidr)
        if Primitive.to_proto(resource.vmware_engine_network):
            res.vmware_engine_network = Primitive.to_proto(
                resource.vmware_engine_network
            )
        if Primitive.to_proto(resource.vmware_engine_network_canonical):
            res.vmware_engine_network_canonical = Primitive.to_proto(
                resource.vmware_engine_network_canonical
            )
        if Primitive.to_proto(resource.management_ip_address_layout_version):
            res.management_ip_address_layout_version = Primitive.to_proto(
                resource.management_ip_address_layout_version
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PrivateCloudNetworkConfig(
            management_cidr=Primitive.from_proto(resource.management_cidr),
            vmware_engine_network=Primitive.from_proto(resource.vmware_engine_network),
            vmware_engine_network_canonical=Primitive.from_proto(
                resource.vmware_engine_network_canonical
            ),
            management_ip_address_layout_version=Primitive.from_proto(
                resource.management_ip_address_layout_version
            ),
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
    def __init__(self, cluster_id: str = None):
        self.cluster_id = cluster_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = private_cloud_pb2.VmwareAlphaPrivateCloudManagementCluster()
        if Primitive.to_proto(resource.cluster_id):
            res.cluster_id = Primitive.to_proto(resource.cluster_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PrivateCloudManagementCluster(
            cluster_id=Primitive.from_proto(resource.cluster_id),
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


class PrivateCloudHcx(object):
    def __init__(
        self,
        internal_ip: str = None,
        version: str = None,
        state: str = None,
        fqdn: str = None,
    ):
        self.internal_ip = internal_ip
        self.version = version
        self.state = state
        self.fqdn = fqdn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = private_cloud_pb2.VmwareAlphaPrivateCloudHcx()
        if Primitive.to_proto(resource.internal_ip):
            res.internal_ip = Primitive.to_proto(resource.internal_ip)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        if PrivateCloudHcxStateEnum.to_proto(resource.state):
            res.state = PrivateCloudHcxStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.fqdn):
            res.fqdn = Primitive.to_proto(resource.fqdn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PrivateCloudHcx(
            internal_ip=Primitive.from_proto(resource.internal_ip),
            version=Primitive.from_proto(resource.version),
            state=PrivateCloudHcxStateEnum.from_proto(resource.state),
            fqdn=Primitive.from_proto(resource.fqdn),
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
        internal_ip: str = None,
        version: str = None,
        state: str = None,
        fqdn: str = None,
    ):
        self.internal_ip = internal_ip
        self.version = version
        self.state = state
        self.fqdn = fqdn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = private_cloud_pb2.VmwareAlphaPrivateCloudNsx()
        if Primitive.to_proto(resource.internal_ip):
            res.internal_ip = Primitive.to_proto(resource.internal_ip)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        if PrivateCloudNsxStateEnum.to_proto(resource.state):
            res.state = PrivateCloudNsxStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.fqdn):
            res.fqdn = Primitive.to_proto(resource.fqdn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PrivateCloudNsx(
            internal_ip=Primitive.from_proto(resource.internal_ip),
            version=Primitive.from_proto(resource.version),
            state=PrivateCloudNsxStateEnum.from_proto(resource.state),
            fqdn=Primitive.from_proto(resource.fqdn),
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
        internal_ip: str = None,
        version: str = None,
        state: str = None,
        fqdn: str = None,
    ):
        self.internal_ip = internal_ip
        self.version = version
        self.state = state
        self.fqdn = fqdn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = private_cloud_pb2.VmwareAlphaPrivateCloudVcenter()
        if Primitive.to_proto(resource.internal_ip):
            res.internal_ip = Primitive.to_proto(resource.internal_ip)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        if PrivateCloudVcenterStateEnum.to_proto(resource.state):
            res.state = PrivateCloudVcenterStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.fqdn):
            res.fqdn = Primitive.to_proto(resource.fqdn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PrivateCloudVcenter(
            internal_ip=Primitive.from_proto(resource.internal_ip),
            version=Primitive.from_proto(resource.version),
            state=PrivateCloudVcenterStateEnum.from_proto(resource.state),
            fqdn=Primitive.from_proto(resource.fqdn),
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
        return private_cloud_pb2.VmwareAlphaPrivateCloudStateEnum.Value(
            "VmwareAlphaPrivateCloudStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return private_cloud_pb2.VmwareAlphaPrivateCloudStateEnum.Name(resource)[
            len("VmwareAlphaPrivateCloudStateEnum") :
        ]


class PrivateCloudHcxStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return private_cloud_pb2.VmwareAlphaPrivateCloudHcxStateEnum.Value(
            "VmwareAlphaPrivateCloudHcxStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return private_cloud_pb2.VmwareAlphaPrivateCloudHcxStateEnum.Name(resource)[
            len("VmwareAlphaPrivateCloudHcxStateEnum") :
        ]


class PrivateCloudNsxStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return private_cloud_pb2.VmwareAlphaPrivateCloudNsxStateEnum.Value(
            "VmwareAlphaPrivateCloudNsxStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return private_cloud_pb2.VmwareAlphaPrivateCloudNsxStateEnum.Name(resource)[
            len("VmwareAlphaPrivateCloudNsxStateEnum") :
        ]


class PrivateCloudVcenterStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return private_cloud_pb2.VmwareAlphaPrivateCloudVcenterStateEnum.Value(
            "VmwareAlphaPrivateCloudVcenterStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return private_cloud_pb2.VmwareAlphaPrivateCloudVcenterStateEnum.Name(resource)[
            len("VmwareAlphaPrivateCloudVcenterStateEnum") :
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
