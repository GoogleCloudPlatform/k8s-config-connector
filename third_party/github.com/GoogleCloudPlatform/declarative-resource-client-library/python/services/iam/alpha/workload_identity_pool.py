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
from google3.cloud.graphite.mmv2.services.google.iam import workload_identity_pool_pb2
from google3.cloud.graphite.mmv2.services.google.iam import (
    workload_identity_pool_pb2_grpc,
)

from typing import List


class WorkloadIdentityPool(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        description: str = None,
        state: str = None,
        disabled: bool = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.description = description
        self.disabled = disabled
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = workload_identity_pool_pb2_grpc.IamAlphaWorkloadIdentityPoolServiceStub(
            channel.Channel()
        )
        request = workload_identity_pool_pb2.ApplyIamAlphaWorkloadIdentityPoolRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyIamAlphaWorkloadIdentityPool(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.state = WorkloadIdentityPoolStateEnum.from_proto(response.state)
        self.disabled = Primitive.from_proto(response.disabled)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = workload_identity_pool_pb2_grpc.IamAlphaWorkloadIdentityPoolServiceStub(
            channel.Channel()
        )
        request = workload_identity_pool_pb2.DeleteIamAlphaWorkloadIdentityPoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteIamAlphaWorkloadIdentityPool(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = workload_identity_pool_pb2_grpc.IamAlphaWorkloadIdentityPoolServiceStub(
            channel.Channel()
        )
        request = workload_identity_pool_pb2.ListIamAlphaWorkloadIdentityPoolRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListIamAlphaWorkloadIdentityPool(request).items

    def to_proto(self):
        resource = workload_identity_pool_pb2.IamAlphaWorkloadIdentityPool()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class WorkloadIdentityPoolStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_identity_pool_pb2.IamAlphaWorkloadIdentityPoolStateEnum.Value(
            "IamAlphaWorkloadIdentityPoolStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_identity_pool_pb2.IamAlphaWorkloadIdentityPoolStateEnum.Name(
            resource
        )[len("IamAlphaWorkloadIdentityPoolStateEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
