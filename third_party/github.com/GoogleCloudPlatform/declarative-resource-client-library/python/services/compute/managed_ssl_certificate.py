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
    managed_ssl_certificate_pb2,
)
from google3.cloud.graphite.mmv2.services.google.compute import (
    managed_ssl_certificate_pb2_grpc,
)

from typing import List


class ManagedSslCertificate(object):
    def __init__(
        self,
        name: str = None,
        id: int = None,
        creation_timestamp: str = None,
        description: str = None,
        self_link: str = None,
        managed: dict = None,
        type: str = None,
        subject_alternative_names: list = None,
        expire_time: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.managed = managed
        self.type = type
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = managed_ssl_certificate_pb2_grpc.ComputeManagedSslCertificateServiceStub(
            channel.Channel()
        )
        request = managed_ssl_certificate_pb2.ApplyComputeManagedSslCertificateRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ManagedSslCertificateManaged.to_proto(self.managed):
            request.resource.managed.CopyFrom(
                ManagedSslCertificateManaged.to_proto(self.managed)
            )
        else:
            request.resource.ClearField("managed")
        if ManagedSslCertificateTypeEnum.to_proto(self.type):
            request.resource.type = ManagedSslCertificateTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeManagedSslCertificate(request)
        self.name = Primitive.from_proto(response.name)
        self.id = Primitive.from_proto(response.id)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.description = Primitive.from_proto(response.description)
        self.self_link = Primitive.from_proto(response.self_link)
        self.managed = ManagedSslCertificateManaged.from_proto(response.managed)
        self.type = ManagedSslCertificateTypeEnum.from_proto(response.type)
        self.subject_alternative_names = Primitive.from_proto(
            response.subject_alternative_names
        )
        self.expire_time = Primitive.from_proto(response.expire_time)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = managed_ssl_certificate_pb2_grpc.ComputeManagedSslCertificateServiceStub(
            channel.Channel()
        )
        request = (
            managed_ssl_certificate_pb2.DeleteComputeManagedSslCertificateRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ManagedSslCertificateManaged.to_proto(self.managed):
            request.resource.managed.CopyFrom(
                ManagedSslCertificateManaged.to_proto(self.managed)
            )
        else:
            request.resource.ClearField("managed")
        if ManagedSslCertificateTypeEnum.to_proto(self.type):
            request.resource.type = ManagedSslCertificateTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeManagedSslCertificate(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = managed_ssl_certificate_pb2_grpc.ComputeManagedSslCertificateServiceStub(
            channel.Channel()
        )
        request = managed_ssl_certificate_pb2.ListComputeManagedSslCertificateRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeManagedSslCertificate(request).items

    def to_proto(self):
        resource = managed_ssl_certificate_pb2.ComputeManagedSslCertificate()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if ManagedSslCertificateManaged.to_proto(self.managed):
            resource.managed.CopyFrom(
                ManagedSslCertificateManaged.to_proto(self.managed)
            )
        else:
            resource.ClearField("managed")
        if ManagedSslCertificateTypeEnum.to_proto(self.type):
            resource.type = ManagedSslCertificateTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class ManagedSslCertificateManaged(object):
    def __init__(
        self, domains: list = None, status: str = None, domain_status: dict = None
    ):
        self.domains = domains
        self.status = status
        self.domain_status = domain_status

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = managed_ssl_certificate_pb2.ComputeManagedSslCertificateManaged()
        if Primitive.to_proto(resource.domains):
            res.domains.extend(Primitive.to_proto(resource.domains))
        if ManagedSslCertificateManagedStatusEnum.to_proto(resource.status):
            res.status = ManagedSslCertificateManagedStatusEnum.to_proto(
                resource.status
            )
        if Primitive.to_proto(resource.domain_status):
            res.domain_status = Primitive.to_proto(resource.domain_status)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ManagedSslCertificateManaged(
            domains=Primitive.from_proto(resource.domains),
            status=ManagedSslCertificateManagedStatusEnum.from_proto(resource.status),
            domain_status=Primitive.from_proto(resource.domain_status),
        )


class ManagedSslCertificateManagedArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ManagedSslCertificateManaged.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ManagedSslCertificateManaged.from_proto(i) for i in resources]


class ManagedSslCertificateManagedStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return managed_ssl_certificate_pb2.ComputeManagedSslCertificateManagedStatusEnum.Value(
            "ComputeManagedSslCertificateManagedStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return managed_ssl_certificate_pb2.ComputeManagedSslCertificateManagedStatusEnum.Name(
            resource
        )[
            len("ComputeManagedSslCertificateManagedStatusEnum") :
        ]


class ManagedSslCertificateManagedDomainStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return managed_ssl_certificate_pb2.ComputeManagedSslCertificateManagedDomainStatusEnum.Value(
            "ComputeManagedSslCertificateManagedDomainStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return managed_ssl_certificate_pb2.ComputeManagedSslCertificateManagedDomainStatusEnum.Name(
            resource
        )[
            len("ComputeManagedSslCertificateManagedDomainStatusEnum") :
        ]


class ManagedSslCertificateTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return managed_ssl_certificate_pb2.ComputeManagedSslCertificateTypeEnum.Value(
            "ComputeManagedSslCertificateTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return managed_ssl_certificate_pb2.ComputeManagedSslCertificateTypeEnum.Name(
            resource
        )[len("ComputeManagedSslCertificateTypeEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
