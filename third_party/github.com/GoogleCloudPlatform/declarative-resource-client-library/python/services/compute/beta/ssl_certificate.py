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
from google3.cloud.graphite.mmv2.services.google.compute import ssl_certificate_pb2
from google3.cloud.graphite.mmv2.services.google.compute import ssl_certificate_pb2_grpc

from typing import List


class SslCertificate(object):
    def __init__(
        self,
        name: str = None,
        id: int = None,
        creation_timestamp: str = None,
        description: str = None,
        self_link: str = None,
        self_managed: dict = None,
        type: str = None,
        subject_alternative_names: list = None,
        expire_time: str = None,
        region: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.self_managed = self_managed
        self.type = type
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = ssl_certificate_pb2_grpc.ComputeBetaSslCertificateServiceStub(
            channel.Channel()
        )
        request = ssl_certificate_pb2.ApplyComputeBetaSslCertificateRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if SslCertificateSelfManaged.to_proto(self.self_managed):
            request.resource.self_managed.CopyFrom(
                SslCertificateSelfManaged.to_proto(self.self_managed)
            )
        else:
            request.resource.ClearField("self_managed")
        if SslCertificateTypeEnum.to_proto(self.type):
            request.resource.type = SslCertificateTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaSslCertificate(request)
        self.name = Primitive.from_proto(response.name)
        self.id = Primitive.from_proto(response.id)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.description = Primitive.from_proto(response.description)
        self.self_link = Primitive.from_proto(response.self_link)
        self.self_managed = SslCertificateSelfManaged.from_proto(response.self_managed)
        self.type = SslCertificateTypeEnum.from_proto(response.type)
        self.subject_alternative_names = Primitive.from_proto(
            response.subject_alternative_names
        )
        self.expire_time = Primitive.from_proto(response.expire_time)
        self.region = Primitive.from_proto(response.region)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = ssl_certificate_pb2_grpc.ComputeBetaSslCertificateServiceStub(
            channel.Channel()
        )
        request = ssl_certificate_pb2.DeleteComputeBetaSslCertificateRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if SslCertificateSelfManaged.to_proto(self.self_managed):
            request.resource.self_managed.CopyFrom(
                SslCertificateSelfManaged.to_proto(self.self_managed)
            )
        else:
            request.resource.ClearField("self_managed")
        if SslCertificateTypeEnum.to_proto(self.type):
            request.resource.type = SslCertificateTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteComputeBetaSslCertificate(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = ssl_certificate_pb2_grpc.ComputeBetaSslCertificateServiceStub(
            channel.Channel()
        )
        request = ssl_certificate_pb2.ListComputeBetaSslCertificateRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeBetaSslCertificate(request).items

    def to_proto(self):
        resource = ssl_certificate_pb2.ComputeBetaSslCertificate()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if SslCertificateSelfManaged.to_proto(self.self_managed):
            resource.self_managed.CopyFrom(
                SslCertificateSelfManaged.to_proto(self.self_managed)
            )
        else:
            resource.ClearField("self_managed")
        if SslCertificateTypeEnum.to_proto(self.type):
            resource.type = SslCertificateTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class SslCertificateSelfManaged(object):
    def __init__(self, certificate: str = None, private_key: str = None):
        self.certificate = certificate
        self.private_key = private_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ssl_certificate_pb2.ComputeBetaSslCertificateSelfManaged()
        if Primitive.to_proto(resource.certificate):
            res.certificate = Primitive.to_proto(resource.certificate)
        if Primitive.to_proto(resource.private_key):
            res.private_key = Primitive.to_proto(resource.private_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SslCertificateSelfManaged(
            certificate=Primitive.from_proto(resource.certificate),
            private_key=Primitive.from_proto(resource.private_key),
        )


class SslCertificateSelfManagedArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SslCertificateSelfManaged.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SslCertificateSelfManaged.from_proto(i) for i in resources]


class SslCertificateTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return ssl_certificate_pb2.ComputeBetaSslCertificateTypeEnum.Value(
            "ComputeBetaSslCertificateTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return ssl_certificate_pb2.ComputeBetaSslCertificateTypeEnum.Name(resource)[
            len("ComputeBetaSslCertificateTypeEnum") :
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
