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
from google3.cloud.graphite.mmv2.services.google.compute import ssl_policy_pb2
from google3.cloud.graphite.mmv2.services.google.compute import ssl_policy_pb2_grpc

from typing import List


class SslPolicy(object):
    def __init__(
        self,
        id: int = None,
        self_link: str = None,
        name: str = None,
        description: str = None,
        profile: str = None,
        min_tls_version: str = None,
        enabled_feature: list = None,
        custom_feature: list = None,
        warning: list = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.profile = profile
        self.min_tls_version = min_tls_version
        self.custom_feature = custom_feature
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = ssl_policy_pb2_grpc.ComputeSslPolicyServiceStub(channel.Channel())
        request = ssl_policy_pb2.ApplyComputeSslPolicyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if SslPolicyProfileEnum.to_proto(self.profile):
            request.resource.profile = SslPolicyProfileEnum.to_proto(self.profile)

        if SslPolicyMinTlsVersionEnum.to_proto(self.min_tls_version):
            request.resource.min_tls_version = SslPolicyMinTlsVersionEnum.to_proto(
                self.min_tls_version
            )

        if Primitive.to_proto(self.custom_feature):
            request.resource.custom_feature.extend(
                Primitive.to_proto(self.custom_feature)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeSslPolicy(request)
        self.id = Primitive.from_proto(response.id)
        self.self_link = Primitive.from_proto(response.self_link)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.profile = SslPolicyProfileEnum.from_proto(response.profile)
        self.min_tls_version = SslPolicyMinTlsVersionEnum.from_proto(
            response.min_tls_version
        )
        self.enabled_feature = Primitive.from_proto(response.enabled_feature)
        self.custom_feature = Primitive.from_proto(response.custom_feature)
        self.warning = SslPolicyWarningArray.from_proto(response.warning)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = ssl_policy_pb2_grpc.ComputeSslPolicyServiceStub(channel.Channel())
        request = ssl_policy_pb2.DeleteComputeSslPolicyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if SslPolicyProfileEnum.to_proto(self.profile):
            request.resource.profile = SslPolicyProfileEnum.to_proto(self.profile)

        if SslPolicyMinTlsVersionEnum.to_proto(self.min_tls_version):
            request.resource.min_tls_version = SslPolicyMinTlsVersionEnum.to_proto(
                self.min_tls_version
            )

        if Primitive.to_proto(self.custom_feature):
            request.resource.custom_feature.extend(
                Primitive.to_proto(self.custom_feature)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeSslPolicy(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = ssl_policy_pb2_grpc.ComputeSslPolicyServiceStub(channel.Channel())
        request = ssl_policy_pb2.ListComputeSslPolicyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeSslPolicy(request).items

    def to_proto(self):
        resource = ssl_policy_pb2.ComputeSslPolicy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if SslPolicyProfileEnum.to_proto(self.profile):
            resource.profile = SslPolicyProfileEnum.to_proto(self.profile)
        if SslPolicyMinTlsVersionEnum.to_proto(self.min_tls_version):
            resource.min_tls_version = SslPolicyMinTlsVersionEnum.to_proto(
                self.min_tls_version
            )
        if Primitive.to_proto(self.custom_feature):
            resource.custom_feature.extend(Primitive.to_proto(self.custom_feature))
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class SslPolicyWarning(object):
    def __init__(self, code: str = None, message: str = None, data: list = None):
        self.code = code
        self.message = message
        self.data = data

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ssl_policy_pb2.ComputeSslPolicyWarning()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if SslPolicyWarningDataArray.to_proto(resource.data):
            res.data.extend(SslPolicyWarningDataArray.to_proto(resource.data))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SslPolicyWarning(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            data=SslPolicyWarningDataArray.from_proto(resource.data),
        )


class SslPolicyWarningArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SslPolicyWarning.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SslPolicyWarning.from_proto(i) for i in resources]


class SslPolicyWarningData(object):
    def __init__(self, key: str = None, value: str = None):
        self.key = key
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ssl_policy_pb2.ComputeSslPolicyWarningData()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SslPolicyWarningData(
            key=Primitive.from_proto(resource.key),
            value=Primitive.from_proto(resource.value),
        )


class SslPolicyWarningDataArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SslPolicyWarningData.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SslPolicyWarningData.from_proto(i) for i in resources]


class SslPolicyProfileEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return ssl_policy_pb2.ComputeSslPolicyProfileEnum.Value(
            "ComputeSslPolicyProfileEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return ssl_policy_pb2.ComputeSslPolicyProfileEnum.Name(resource)[
            len("ComputeSslPolicyProfileEnum") :
        ]


class SslPolicyMinTlsVersionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return ssl_policy_pb2.ComputeSslPolicyMinTlsVersionEnum.Value(
            "ComputeSslPolicyMinTlsVersionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return ssl_policy_pb2.ComputeSslPolicyMinTlsVersionEnum.Name(resource)[
            len("ComputeSslPolicyMinTlsVersionEnum") :
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
