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
from google3.cloud.graphite.mmv2.services.google.vmware import cluster_pb2
from google3.cloud.graphite.mmv2.services.google.vmware import cluster_pb2_grpc

from typing import List


class Cluster(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        state: str = None,
        management: bool = None,
        uid: str = None,
        project: str = None,
        location: str = None,
        private_cloud: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.project = project
        self.location = location
        self.private_cloud = private_cloud
        self.service_account_file = service_account_file

    def apply(self):
        stub = cluster_pb2_grpc.VmwareAlphaClusterServiceStub(channel.Channel())
        request = cluster_pb2.ApplyVmwareAlphaClusterRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.private_cloud):
            request.resource.private_cloud = Primitive.to_proto(self.private_cloud)

        request.service_account_file = self.service_account_file

        response = stub.ApplyVmwareAlphaCluster(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.state = ClusterStateEnum.from_proto(response.state)
        self.management = Primitive.from_proto(response.management)
        self.uid = Primitive.from_proto(response.uid)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.private_cloud = Primitive.from_proto(response.private_cloud)

    def delete(self):
        stub = cluster_pb2_grpc.VmwareAlphaClusterServiceStub(channel.Channel())
        request = cluster_pb2.DeleteVmwareAlphaClusterRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.private_cloud):
            request.resource.private_cloud = Primitive.to_proto(self.private_cloud)

        response = stub.DeleteVmwareAlphaCluster(request)

    @classmethod
    def list(self, project, location, privateCloud, service_account_file=""):
        stub = cluster_pb2_grpc.VmwareAlphaClusterServiceStub(channel.Channel())
        request = cluster_pb2.ListVmwareAlphaClusterRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.PrivateCloud = privateCloud

        return stub.ListVmwareAlphaCluster(request).items

    def to_proto(self):
        resource = cluster_pb2.VmwareAlphaCluster()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.private_cloud):
            resource.private_cloud = Primitive.to_proto(self.private_cloud)
        return resource


class ClusterStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.VmwareAlphaClusterStateEnum.Value(
            "VmwareAlphaClusterStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.VmwareAlphaClusterStateEnum.Name(resource)[
            len("VmwareAlphaClusterStateEnum") :
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
