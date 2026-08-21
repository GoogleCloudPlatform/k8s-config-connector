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
from google3.cloud.graphite.mmv2.services.google.compute import target_ssl_proxy_pb2
from google3.cloud.graphite.mmv2.services.google.compute import (
    target_ssl_proxy_pb2_grpc,
)

from typing import List


class TargetSslProxy(object):
    def __init__(
        self,
        id: int = None,
        name: str = None,
        description: str = None,
        self_link: str = None,
        service: str = None,
        ssl_certificates: list = None,
        proxy_header: str = None,
        ssl_policy: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.id = id
        self.name = name
        self.description = description
        self.service = service
        self.ssl_certificates = ssl_certificates
        self.proxy_header = proxy_header
        self.ssl_policy = ssl_policy
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = target_ssl_proxy_pb2_grpc.ComputeTargetSslProxyServiceStub(
            channel.Channel()
        )
        request = target_ssl_proxy_pb2.ApplyComputeTargetSslProxyRequest()
        if Primitive.to_proto(self.id):
            request.resource.id = Primitive.to_proto(self.id)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.service):
            request.resource.service = Primitive.to_proto(self.service)

        if Primitive.to_proto(self.ssl_certificates):
            request.resource.ssl_certificates.extend(
                Primitive.to_proto(self.ssl_certificates)
            )
        if TargetSslProxyProxyHeaderEnum.to_proto(self.proxy_header):
            request.resource.proxy_header = TargetSslProxyProxyHeaderEnum.to_proto(
                self.proxy_header
            )

        if Primitive.to_proto(self.ssl_policy):
            request.resource.ssl_policy = Primitive.to_proto(self.ssl_policy)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeTargetSslProxy(request)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.self_link = Primitive.from_proto(response.self_link)
        self.service = Primitive.from_proto(response.service)
        self.ssl_certificates = Primitive.from_proto(response.ssl_certificates)
        self.proxy_header = TargetSslProxyProxyHeaderEnum.from_proto(
            response.proxy_header
        )
        self.ssl_policy = Primitive.from_proto(response.ssl_policy)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = target_ssl_proxy_pb2_grpc.ComputeTargetSslProxyServiceStub(
            channel.Channel()
        )
        request = target_ssl_proxy_pb2.DeleteComputeTargetSslProxyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.id):
            request.resource.id = Primitive.to_proto(self.id)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.service):
            request.resource.service = Primitive.to_proto(self.service)

        if Primitive.to_proto(self.ssl_certificates):
            request.resource.ssl_certificates.extend(
                Primitive.to_proto(self.ssl_certificates)
            )
        if TargetSslProxyProxyHeaderEnum.to_proto(self.proxy_header):
            request.resource.proxy_header = TargetSslProxyProxyHeaderEnum.to_proto(
                self.proxy_header
            )

        if Primitive.to_proto(self.ssl_policy):
            request.resource.ssl_policy = Primitive.to_proto(self.ssl_policy)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeTargetSslProxy(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = target_ssl_proxy_pb2_grpc.ComputeTargetSslProxyServiceStub(
            channel.Channel()
        )
        request = target_ssl_proxy_pb2.ListComputeTargetSslProxyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeTargetSslProxy(request).items

    def to_proto(self):
        resource = target_ssl_proxy_pb2.ComputeTargetSslProxy()
        if Primitive.to_proto(self.id):
            resource.id = Primitive.to_proto(self.id)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.service):
            resource.service = Primitive.to_proto(self.service)
        if Primitive.to_proto(self.ssl_certificates):
            resource.ssl_certificates.extend(Primitive.to_proto(self.ssl_certificates))
        if TargetSslProxyProxyHeaderEnum.to_proto(self.proxy_header):
            resource.proxy_header = TargetSslProxyProxyHeaderEnum.to_proto(
                self.proxy_header
            )
        if Primitive.to_proto(self.ssl_policy):
            resource.ssl_policy = Primitive.to_proto(self.ssl_policy)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class TargetSslProxyProxyHeaderEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return target_ssl_proxy_pb2.ComputeTargetSslProxyProxyHeaderEnum.Value(
            "ComputeTargetSslProxyProxyHeaderEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return target_ssl_proxy_pb2.ComputeTargetSslProxyProxyHeaderEnum.Name(resource)[
            len("ComputeTargetSslProxyProxyHeaderEnum") :
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
