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
from google3.cloud.graphite.mmv2.services.google.compute import target_pool_pb2
from google3.cloud.graphite.mmv2.services.google.compute import target_pool_pb2_grpc

from typing import List


class TargetPool(object):
    def __init__(
        self,
        backup_pool: str = None,
        description: str = None,
        failover_ratio: float = None,
        health_checks: list = None,
        instances: list = None,
        name: str = None,
        region: str = None,
        self_link: str = None,
        session_affinity: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.backup_pool = backup_pool
        self.description = description
        self.failover_ratio = failover_ratio
        self.health_checks = health_checks
        self.instances = instances
        self.name = name
        self.region = region
        self.self_link = self_link
        self.session_affinity = session_affinity
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = target_pool_pb2_grpc.ComputeTargetPoolServiceStub(channel.Channel())
        request = target_pool_pb2.ApplyComputeTargetPoolRequest()
        if Primitive.to_proto(self.backup_pool):
            request.resource.backup_pool = Primitive.to_proto(self.backup_pool)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.failover_ratio):
            request.resource.failover_ratio = Primitive.to_proto(self.failover_ratio)

        if Primitive.to_proto(self.health_checks):
            request.resource.health_checks.extend(
                Primitive.to_proto(self.health_checks)
            )
        if Primitive.to_proto(self.instances):
            request.resource.instances.extend(Primitive.to_proto(self.instances))
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.self_link):
            request.resource.self_link = Primitive.to_proto(self.self_link)

        if TargetPoolSessionAffinityEnum.to_proto(self.session_affinity):
            request.resource.session_affinity = TargetPoolSessionAffinityEnum.to_proto(
                self.session_affinity
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeTargetPool(request)
        self.backup_pool = Primitive.from_proto(response.backup_pool)
        self.description = Primitive.from_proto(response.description)
        self.failover_ratio = Primitive.from_proto(response.failover_ratio)
        self.health_checks = Primitive.from_proto(response.health_checks)
        self.instances = Primitive.from_proto(response.instances)
        self.name = Primitive.from_proto(response.name)
        self.region = Primitive.from_proto(response.region)
        self.self_link = Primitive.from_proto(response.self_link)
        self.session_affinity = TargetPoolSessionAffinityEnum.from_proto(
            response.session_affinity
        )
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = target_pool_pb2_grpc.ComputeTargetPoolServiceStub(channel.Channel())
        request = target_pool_pb2.DeleteComputeTargetPoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.backup_pool):
            request.resource.backup_pool = Primitive.to_proto(self.backup_pool)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.failover_ratio):
            request.resource.failover_ratio = Primitive.to_proto(self.failover_ratio)

        if Primitive.to_proto(self.health_checks):
            request.resource.health_checks.extend(
                Primitive.to_proto(self.health_checks)
            )
        if Primitive.to_proto(self.instances):
            request.resource.instances.extend(Primitive.to_proto(self.instances))
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.self_link):
            request.resource.self_link = Primitive.to_proto(self.self_link)

        if TargetPoolSessionAffinityEnum.to_proto(self.session_affinity):
            request.resource.session_affinity = TargetPoolSessionAffinityEnum.to_proto(
                self.session_affinity
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeTargetPool(request)

    @classmethod
    def list(self, project, region, service_account_file=""):
        stub = target_pool_pb2_grpc.ComputeTargetPoolServiceStub(channel.Channel())
        request = target_pool_pb2.ListComputeTargetPoolRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        return stub.ListComputeTargetPool(request).items

    def to_proto(self):
        resource = target_pool_pb2.ComputeTargetPool()
        if Primitive.to_proto(self.backup_pool):
            resource.backup_pool = Primitive.to_proto(self.backup_pool)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.failover_ratio):
            resource.failover_ratio = Primitive.to_proto(self.failover_ratio)
        if Primitive.to_proto(self.health_checks):
            resource.health_checks.extend(Primitive.to_proto(self.health_checks))
        if Primitive.to_proto(self.instances):
            resource.instances.extend(Primitive.to_proto(self.instances))
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.self_link):
            resource.self_link = Primitive.to_proto(self.self_link)
        if TargetPoolSessionAffinityEnum.to_proto(self.session_affinity):
            resource.session_affinity = TargetPoolSessionAffinityEnum.to_proto(
                self.session_affinity
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class TargetPoolSessionAffinityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return target_pool_pb2.ComputeTargetPoolSessionAffinityEnum.Value(
            "ComputeTargetPoolSessionAffinityEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return target_pool_pb2.ComputeTargetPoolSessionAffinityEnum.Name(resource)[
            len("ComputeTargetPoolSessionAffinityEnum") :
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
