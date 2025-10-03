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
from google3.cloud.graphite.mmv2.services.google.cloudkms import ekm_connection_pb2
from google3.cloud.graphite.mmv2.services.google.cloudkms import ekm_connection_pb2_grpc

from typing import List


class EkmConnection(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        service_resolvers: list = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.service_resolvers = service_resolvers
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = ekm_connection_pb2_grpc.CloudkmsAlphaEkmConnectionServiceStub(
            channel.Channel()
        )
        request = ekm_connection_pb2.ApplyCloudkmsAlphaEkmConnectionRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if EkmConnectionServiceResolversArray.to_proto(self.service_resolvers):
            request.resource.service_resolvers.extend(
                EkmConnectionServiceResolversArray.to_proto(self.service_resolvers)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudkmsAlphaEkmConnection(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.service_resolvers = EkmConnectionServiceResolversArray.from_proto(
            response.service_resolvers
        )
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = ekm_connection_pb2_grpc.CloudkmsAlphaEkmConnectionServiceStub(
            channel.Channel()
        )
        request = ekm_connection_pb2.DeleteCloudkmsAlphaEkmConnectionRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if EkmConnectionServiceResolversArray.to_proto(self.service_resolvers):
            request.resource.service_resolvers.extend(
                EkmConnectionServiceResolversArray.to_proto(self.service_resolvers)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteCloudkmsAlphaEkmConnection(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = ekm_connection_pb2_grpc.CloudkmsAlphaEkmConnectionServiceStub(
            channel.Channel()
        )
        request = ekm_connection_pb2.ListCloudkmsAlphaEkmConnectionRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListCloudkmsAlphaEkmConnection(request).items

    def to_proto(self):
        resource = ekm_connection_pb2.CloudkmsAlphaEkmConnection()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if EkmConnectionServiceResolversArray.to_proto(self.service_resolvers):
            resource.service_resolvers.extend(
                EkmConnectionServiceResolversArray.to_proto(self.service_resolvers)
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class EkmConnectionServiceResolvers(object):
    def __init__(
        self,
        service_directory_service: str = None,
        endpoint_filter: str = None,
        hostname: str = None,
        server_certificates: list = None,
    ):
        self.service_directory_service = service_directory_service
        self.endpoint_filter = endpoint_filter
        self.hostname = hostname
        self.server_certificates = server_certificates

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ekm_connection_pb2.CloudkmsAlphaEkmConnectionServiceResolvers()
        if Primitive.to_proto(resource.service_directory_service):
            res.service_directory_service = Primitive.to_proto(
                resource.service_directory_service
            )
        if Primitive.to_proto(resource.endpoint_filter):
            res.endpoint_filter = Primitive.to_proto(resource.endpoint_filter)
        if Primitive.to_proto(resource.hostname):
            res.hostname = Primitive.to_proto(resource.hostname)
        if EkmConnectionServiceResolversServerCertificatesArray.to_proto(
            resource.server_certificates
        ):
            res.server_certificates.extend(
                EkmConnectionServiceResolversServerCertificatesArray.to_proto(
                    resource.server_certificates
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EkmConnectionServiceResolvers(
            service_directory_service=Primitive.from_proto(
                resource.service_directory_service
            ),
            endpoint_filter=Primitive.from_proto(resource.endpoint_filter),
            hostname=Primitive.from_proto(resource.hostname),
            server_certificates=EkmConnectionServiceResolversServerCertificatesArray.from_proto(
                resource.server_certificates
            ),
        )


class EkmConnectionServiceResolversArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EkmConnectionServiceResolvers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EkmConnectionServiceResolvers.from_proto(i) for i in resources]


class EkmConnectionServiceResolversServerCertificates(object):
    def __init__(
        self,
        raw_der: str = None,
        parsed: bool = None,
        issuer: str = None,
        subject: str = None,
        subject_alternative_dns_names: list = None,
        not_before_time: str = None,
        not_after_time: str = None,
        serial_number: str = None,
        sha256_fingerprint: str = None,
    ):
        self.raw_der = raw_der
        self.parsed = parsed
        self.issuer = issuer
        self.subject = subject
        self.subject_alternative_dns_names = subject_alternative_dns_names
        self.not_before_time = not_before_time
        self.not_after_time = not_after_time
        self.serial_number = serial_number
        self.sha256_fingerprint = sha256_fingerprint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            ekm_connection_pb2.CloudkmsAlphaEkmConnectionServiceResolversServerCertificates()
        )
        if Primitive.to_proto(resource.raw_der):
            res.raw_der = Primitive.to_proto(resource.raw_der)
        if Primitive.to_proto(resource.parsed):
            res.parsed = Primitive.to_proto(resource.parsed)
        if Primitive.to_proto(resource.issuer):
            res.issuer = Primitive.to_proto(resource.issuer)
        if Primitive.to_proto(resource.subject):
            res.subject = Primitive.to_proto(resource.subject)
        if Primitive.to_proto(resource.subject_alternative_dns_names):
            res.subject_alternative_dns_names.extend(
                Primitive.to_proto(resource.subject_alternative_dns_names)
            )
        if Primitive.to_proto(resource.not_before_time):
            res.not_before_time = Primitive.to_proto(resource.not_before_time)
        if Primitive.to_proto(resource.not_after_time):
            res.not_after_time = Primitive.to_proto(resource.not_after_time)
        if Primitive.to_proto(resource.serial_number):
            res.serial_number = Primitive.to_proto(resource.serial_number)
        if Primitive.to_proto(resource.sha256_fingerprint):
            res.sha256_fingerprint = Primitive.to_proto(resource.sha256_fingerprint)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EkmConnectionServiceResolversServerCertificates(
            raw_der=Primitive.from_proto(resource.raw_der),
            parsed=Primitive.from_proto(resource.parsed),
            issuer=Primitive.from_proto(resource.issuer),
            subject=Primitive.from_proto(resource.subject),
            subject_alternative_dns_names=Primitive.from_proto(
                resource.subject_alternative_dns_names
            ),
            not_before_time=Primitive.from_proto(resource.not_before_time),
            not_after_time=Primitive.from_proto(resource.not_after_time),
            serial_number=Primitive.from_proto(resource.serial_number),
            sha256_fingerprint=Primitive.from_proto(resource.sha256_fingerprint),
        )


class EkmConnectionServiceResolversServerCertificatesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EkmConnectionServiceResolversServerCertificates.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EkmConnectionServiceResolversServerCertificates.from_proto(i)
            for i in resources
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
