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
from google3.cloud.graphite.mmv2.services.google.iam import (
    workload_identity_pool_provider_pb2,
)
from google3.cloud.graphite.mmv2.services.google.iam import (
    workload_identity_pool_provider_pb2_grpc,
)

from typing import List


class WorkloadIdentityPoolProvider(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        description: str = None,
        state: str = None,
        disabled: bool = None,
        attribute_mapping: dict = None,
        attribute_condition: str = None,
        aws: dict = None,
        oidc: dict = None,
        project: str = None,
        location: str = None,
        workload_identity_pool: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.description = description
        self.disabled = disabled
        self.attribute_mapping = attribute_mapping
        self.attribute_condition = attribute_condition
        self.aws = aws
        self.oidc = oidc
        self.project = project
        self.location = location
        self.workload_identity_pool = workload_identity_pool
        self.service_account_file = service_account_file

    def apply(self):
        stub = workload_identity_pool_provider_pb2_grpc.IamWorkloadIdentityPoolProviderServiceStub(
            channel.Channel()
        )
        request = (
            workload_identity_pool_provider_pb2.ApplyIamWorkloadIdentityPoolProviderRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.attribute_mapping):
            request.resource.attribute_mapping = Primitive.to_proto(
                self.attribute_mapping
            )

        if Primitive.to_proto(self.attribute_condition):
            request.resource.attribute_condition = Primitive.to_proto(
                self.attribute_condition
            )

        if WorkloadIdentityPoolProviderAws.to_proto(self.aws):
            request.resource.aws.CopyFrom(
                WorkloadIdentityPoolProviderAws.to_proto(self.aws)
            )
        else:
            request.resource.ClearField("aws")
        if WorkloadIdentityPoolProviderOidc.to_proto(self.oidc):
            request.resource.oidc.CopyFrom(
                WorkloadIdentityPoolProviderOidc.to_proto(self.oidc)
            )
        else:
            request.resource.ClearField("oidc")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.workload_identity_pool):
            request.resource.workload_identity_pool = Primitive.to_proto(
                self.workload_identity_pool
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyIamWorkloadIdentityPoolProvider(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.state = WorkloadIdentityPoolProviderStateEnum.from_proto(response.state)
        self.disabled = Primitive.from_proto(response.disabled)
        self.attribute_mapping = Primitive.from_proto(response.attribute_mapping)
        self.attribute_condition = Primitive.from_proto(response.attribute_condition)
        self.aws = WorkloadIdentityPoolProviderAws.from_proto(response.aws)
        self.oidc = WorkloadIdentityPoolProviderOidc.from_proto(response.oidc)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.workload_identity_pool = Primitive.from_proto(
            response.workload_identity_pool
        )

    def delete(self):
        stub = workload_identity_pool_provider_pb2_grpc.IamWorkloadIdentityPoolProviderServiceStub(
            channel.Channel()
        )
        request = (
            workload_identity_pool_provider_pb2.DeleteIamWorkloadIdentityPoolProviderRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.attribute_mapping):
            request.resource.attribute_mapping = Primitive.to_proto(
                self.attribute_mapping
            )

        if Primitive.to_proto(self.attribute_condition):
            request.resource.attribute_condition = Primitive.to_proto(
                self.attribute_condition
            )

        if WorkloadIdentityPoolProviderAws.to_proto(self.aws):
            request.resource.aws.CopyFrom(
                WorkloadIdentityPoolProviderAws.to_proto(self.aws)
            )
        else:
            request.resource.ClearField("aws")
        if WorkloadIdentityPoolProviderOidc.to_proto(self.oidc):
            request.resource.oidc.CopyFrom(
                WorkloadIdentityPoolProviderOidc.to_proto(self.oidc)
            )
        else:
            request.resource.ClearField("oidc")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.workload_identity_pool):
            request.resource.workload_identity_pool = Primitive.to_proto(
                self.workload_identity_pool
            )

        response = stub.DeleteIamWorkloadIdentityPoolProvider(request)

    @classmethod
    def list(self, project, location, workloadIdentityPool, service_account_file=""):
        stub = workload_identity_pool_provider_pb2_grpc.IamWorkloadIdentityPoolProviderServiceStub(
            channel.Channel()
        )
        request = (
            workload_identity_pool_provider_pb2.ListIamWorkloadIdentityPoolProviderRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.WorkloadIdentityPool = workloadIdentityPool

        return stub.ListIamWorkloadIdentityPoolProvider(request).items

    def to_proto(self):
        resource = workload_identity_pool_provider_pb2.IamWorkloadIdentityPoolProvider()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if Primitive.to_proto(self.attribute_mapping):
            resource.attribute_mapping = Primitive.to_proto(self.attribute_mapping)
        if Primitive.to_proto(self.attribute_condition):
            resource.attribute_condition = Primitive.to_proto(self.attribute_condition)
        if WorkloadIdentityPoolProviderAws.to_proto(self.aws):
            resource.aws.CopyFrom(WorkloadIdentityPoolProviderAws.to_proto(self.aws))
        else:
            resource.ClearField("aws")
        if WorkloadIdentityPoolProviderOidc.to_proto(self.oidc):
            resource.oidc.CopyFrom(WorkloadIdentityPoolProviderOidc.to_proto(self.oidc))
        else:
            resource.ClearField("oidc")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.workload_identity_pool):
            resource.workload_identity_pool = Primitive.to_proto(
                self.workload_identity_pool
            )
        return resource


class WorkloadIdentityPoolProviderAws(object):
    def __init__(self, account_id: str = None, sts_uri: list = None):
        self.account_id = account_id
        self.sts_uri = sts_uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_identity_pool_provider_pb2.IamWorkloadIdentityPoolProviderAws()
        if Primitive.to_proto(resource.account_id):
            res.account_id = Primitive.to_proto(resource.account_id)
        if Primitive.to_proto(resource.sts_uri):
            res.sts_uri.extend(Primitive.to_proto(resource.sts_uri))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadIdentityPoolProviderAws(
            account_id=Primitive.from_proto(resource.account_id),
            sts_uri=Primitive.from_proto(resource.sts_uri),
        )


class WorkloadIdentityPoolProviderAwsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadIdentityPoolProviderAws.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadIdentityPoolProviderAws.from_proto(i) for i in resources]


class WorkloadIdentityPoolProviderOidc(object):
    def __init__(self, issuer_uri: str = None, allowed_audiences: list = None):
        self.issuer_uri = issuer_uri
        self.allowed_audiences = allowed_audiences

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_identity_pool_provider_pb2.IamWorkloadIdentityPoolProviderOidc()
        if Primitive.to_proto(resource.issuer_uri):
            res.issuer_uri = Primitive.to_proto(resource.issuer_uri)
        if Primitive.to_proto(resource.allowed_audiences):
            res.allowed_audiences.extend(Primitive.to_proto(resource.allowed_audiences))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadIdentityPoolProviderOidc(
            issuer_uri=Primitive.from_proto(resource.issuer_uri),
            allowed_audiences=Primitive.from_proto(resource.allowed_audiences),
        )


class WorkloadIdentityPoolProviderOidcArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadIdentityPoolProviderOidc.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadIdentityPoolProviderOidc.from_proto(i) for i in resources]


class WorkloadIdentityPoolProviderStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_identity_pool_provider_pb2.IamWorkloadIdentityPoolProviderStateEnum.Value(
            "IamWorkloadIdentityPoolProviderStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_identity_pool_provider_pb2.IamWorkloadIdentityPoolProviderStateEnum.Name(
            resource
        )[
            len("IamWorkloadIdentityPoolProviderStateEnum") :
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
