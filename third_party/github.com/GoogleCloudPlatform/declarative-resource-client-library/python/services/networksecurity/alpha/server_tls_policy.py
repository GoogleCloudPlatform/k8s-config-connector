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
from google3.cloud.graphite.mmv2.services.google.network_security import (
    server_tls_policy_pb2,
)
from google3.cloud.graphite.mmv2.services.google.network_security import (
    server_tls_policy_pb2_grpc,
)

from typing import List


class ServerTlsPolicy(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        allow_open: bool = None,
        server_certificate: dict = None,
        mtls_policy: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.labels = labels
        self.allow_open = allow_open
        self.server_certificate = server_certificate
        self.mtls_policy = mtls_policy
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = (
            server_tls_policy_pb2_grpc.NetworksecurityAlphaServerTlsPolicyServiceStub(
                channel.Channel()
            )
        )
        request = (
            server_tls_policy_pb2.ApplyNetworksecurityAlphaServerTlsPolicyRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.allow_open):
            request.resource.allow_open = Primitive.to_proto(self.allow_open)

        if ServerTlsPolicyServerCertificate.to_proto(self.server_certificate):
            request.resource.server_certificate.CopyFrom(
                ServerTlsPolicyServerCertificate.to_proto(self.server_certificate)
            )
        else:
            request.resource.ClearField("server_certificate")
        if ServerTlsPolicyMtlsPolicy.to_proto(self.mtls_policy):
            request.resource.mtls_policy.CopyFrom(
                ServerTlsPolicyMtlsPolicy.to_proto(self.mtls_policy)
            )
        else:
            request.resource.ClearField("mtls_policy")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworksecurityAlphaServerTlsPolicy(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.allow_open = Primitive.from_proto(response.allow_open)
        self.server_certificate = ServerTlsPolicyServerCertificate.from_proto(
            response.server_certificate
        )
        self.mtls_policy = ServerTlsPolicyMtlsPolicy.from_proto(response.mtls_policy)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = (
            server_tls_policy_pb2_grpc.NetworksecurityAlphaServerTlsPolicyServiceStub(
                channel.Channel()
            )
        )
        request = (
            server_tls_policy_pb2.DeleteNetworksecurityAlphaServerTlsPolicyRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.allow_open):
            request.resource.allow_open = Primitive.to_proto(self.allow_open)

        if ServerTlsPolicyServerCertificate.to_proto(self.server_certificate):
            request.resource.server_certificate.CopyFrom(
                ServerTlsPolicyServerCertificate.to_proto(self.server_certificate)
            )
        else:
            request.resource.ClearField("server_certificate")
        if ServerTlsPolicyMtlsPolicy.to_proto(self.mtls_policy):
            request.resource.mtls_policy.CopyFrom(
                ServerTlsPolicyMtlsPolicy.to_proto(self.mtls_policy)
            )
        else:
            request.resource.ClearField("mtls_policy")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworksecurityAlphaServerTlsPolicy(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = (
            server_tls_policy_pb2_grpc.NetworksecurityAlphaServerTlsPolicyServiceStub(
                channel.Channel()
            )
        )
        request = server_tls_policy_pb2.ListNetworksecurityAlphaServerTlsPolicyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworksecurityAlphaServerTlsPolicy(request).items

    def to_proto(self):
        resource = server_tls_policy_pb2.NetworksecurityAlphaServerTlsPolicy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.allow_open):
            resource.allow_open = Primitive.to_proto(self.allow_open)
        if ServerTlsPolicyServerCertificate.to_proto(self.server_certificate):
            resource.server_certificate.CopyFrom(
                ServerTlsPolicyServerCertificate.to_proto(self.server_certificate)
            )
        else:
            resource.ClearField("server_certificate")
        if ServerTlsPolicyMtlsPolicy.to_proto(self.mtls_policy):
            resource.mtls_policy.CopyFrom(
                ServerTlsPolicyMtlsPolicy.to_proto(self.mtls_policy)
            )
        else:
            resource.ClearField("mtls_policy")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ServerTlsPolicyServerCertificate(object):
    def __init__(
        self,
        local_filepath: dict = None,
        grpc_endpoint: dict = None,
        certificate_provider_instance: dict = None,
    ):
        self.local_filepath = local_filepath
        self.grpc_endpoint = grpc_endpoint
        self.certificate_provider_instance = certificate_provider_instance

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            server_tls_policy_pb2.NetworksecurityAlphaServerTlsPolicyServerCertificate()
        )
        if ServerTlsPolicyServerCertificateLocalFilepath.to_proto(
            resource.local_filepath
        ):
            res.local_filepath.CopyFrom(
                ServerTlsPolicyServerCertificateLocalFilepath.to_proto(
                    resource.local_filepath
                )
            )
        else:
            res.ClearField("local_filepath")
        if ServerTlsPolicyServerCertificateGrpcEndpoint.to_proto(
            resource.grpc_endpoint
        ):
            res.grpc_endpoint.CopyFrom(
                ServerTlsPolicyServerCertificateGrpcEndpoint.to_proto(
                    resource.grpc_endpoint
                )
            )
        else:
            res.ClearField("grpc_endpoint")
        if ServerTlsPolicyServerCertificateCertificateProviderInstance.to_proto(
            resource.certificate_provider_instance
        ):
            res.certificate_provider_instance.CopyFrom(
                ServerTlsPolicyServerCertificateCertificateProviderInstance.to_proto(
                    resource.certificate_provider_instance
                )
            )
        else:
            res.ClearField("certificate_provider_instance")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServerTlsPolicyServerCertificate(
            local_filepath=ServerTlsPolicyServerCertificateLocalFilepath.from_proto(
                resource.local_filepath
            ),
            grpc_endpoint=ServerTlsPolicyServerCertificateGrpcEndpoint.from_proto(
                resource.grpc_endpoint
            ),
            certificate_provider_instance=ServerTlsPolicyServerCertificateCertificateProviderInstance.from_proto(
                resource.certificate_provider_instance
            ),
        )


class ServerTlsPolicyServerCertificateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServerTlsPolicyServerCertificate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServerTlsPolicyServerCertificate.from_proto(i) for i in resources]


class ServerTlsPolicyServerCertificateLocalFilepath(object):
    def __init__(self, certificate_path: str = None, private_key_path: str = None):
        self.certificate_path = certificate_path
        self.private_key_path = private_key_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            server_tls_policy_pb2.NetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepath()
        )
        if Primitive.to_proto(resource.certificate_path):
            res.certificate_path = Primitive.to_proto(resource.certificate_path)
        if Primitive.to_proto(resource.private_key_path):
            res.private_key_path = Primitive.to_proto(resource.private_key_path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServerTlsPolicyServerCertificateLocalFilepath(
            certificate_path=Primitive.from_proto(resource.certificate_path),
            private_key_path=Primitive.from_proto(resource.private_key_path),
        )


class ServerTlsPolicyServerCertificateLocalFilepathArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServerTlsPolicyServerCertificateLocalFilepath.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServerTlsPolicyServerCertificateLocalFilepath.from_proto(i)
            for i in resources
        ]


class ServerTlsPolicyServerCertificateGrpcEndpoint(object):
    def __init__(self, target_uri: str = None):
        self.target_uri = target_uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            server_tls_policy_pb2.NetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpoint()
        )
        if Primitive.to_proto(resource.target_uri):
            res.target_uri = Primitive.to_proto(resource.target_uri)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServerTlsPolicyServerCertificateGrpcEndpoint(
            target_uri=Primitive.from_proto(resource.target_uri),
        )


class ServerTlsPolicyServerCertificateGrpcEndpointArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServerTlsPolicyServerCertificateGrpcEndpoint.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServerTlsPolicyServerCertificateGrpcEndpoint.from_proto(i)
            for i in resources
        ]


class ServerTlsPolicyServerCertificateCertificateProviderInstance(object):
    def __init__(self, plugin_instance: str = None):
        self.plugin_instance = plugin_instance

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            server_tls_policy_pb2.NetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstance()
        )
        if Primitive.to_proto(resource.plugin_instance):
            res.plugin_instance = Primitive.to_proto(resource.plugin_instance)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServerTlsPolicyServerCertificateCertificateProviderInstance(
            plugin_instance=Primitive.from_proto(resource.plugin_instance),
        )


class ServerTlsPolicyServerCertificateCertificateProviderInstanceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServerTlsPolicyServerCertificateCertificateProviderInstance.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServerTlsPolicyServerCertificateCertificateProviderInstance.from_proto(i)
            for i in resources
        ]


class ServerTlsPolicyMtlsPolicy(object):
    def __init__(self, client_validation_ca: list = None):
        self.client_validation_ca = client_validation_ca

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = server_tls_policy_pb2.NetworksecurityAlphaServerTlsPolicyMtlsPolicy()
        if ServerTlsPolicyMtlsPolicyClientValidationCaArray.to_proto(
            resource.client_validation_ca
        ):
            res.client_validation_ca.extend(
                ServerTlsPolicyMtlsPolicyClientValidationCaArray.to_proto(
                    resource.client_validation_ca
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServerTlsPolicyMtlsPolicy(
            client_validation_ca=ServerTlsPolicyMtlsPolicyClientValidationCaArray.from_proto(
                resource.client_validation_ca
            ),
        )


class ServerTlsPolicyMtlsPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServerTlsPolicyMtlsPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServerTlsPolicyMtlsPolicy.from_proto(i) for i in resources]


class ServerTlsPolicyMtlsPolicyClientValidationCa(object):
    def __init__(
        self,
        ca_cert_path: str = None,
        grpc_endpoint: dict = None,
        certificate_provider_instance: dict = None,
    ):
        self.ca_cert_path = ca_cert_path
        self.grpc_endpoint = grpc_endpoint
        self.certificate_provider_instance = certificate_provider_instance

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            server_tls_policy_pb2.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa()
        )
        if Primitive.to_proto(resource.ca_cert_path):
            res.ca_cert_path = Primitive.to_proto(resource.ca_cert_path)
        if ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint.to_proto(
            resource.grpc_endpoint
        ):
            res.grpc_endpoint.CopyFrom(
                ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint.to_proto(
                    resource.grpc_endpoint
                )
            )
        else:
            res.ClearField("grpc_endpoint")
        if ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance.to_proto(
            resource.certificate_provider_instance
        ):
            res.certificate_provider_instance.CopyFrom(
                ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance.to_proto(
                    resource.certificate_provider_instance
                )
            )
        else:
            res.ClearField("certificate_provider_instance")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServerTlsPolicyMtlsPolicyClientValidationCa(
            ca_cert_path=Primitive.from_proto(resource.ca_cert_path),
            grpc_endpoint=ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint.from_proto(
                resource.grpc_endpoint
            ),
            certificate_provider_instance=ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance.from_proto(
                resource.certificate_provider_instance
            ),
        )


class ServerTlsPolicyMtlsPolicyClientValidationCaArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServerTlsPolicyMtlsPolicyClientValidationCa.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServerTlsPolicyMtlsPolicyClientValidationCa.from_proto(i) for i in resources
        ]


class ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(object):
    def __init__(self, target_uri: str = None):
        self.target_uri = target_uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            server_tls_policy_pb2.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint()
        )
        if Primitive.to_proto(resource.target_uri):
            res.target_uri = Primitive.to_proto(resource.target_uri)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(
            target_uri=Primitive.from_proto(resource.target_uri),
        )


class ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint.from_proto(i)
            for i in resources
        ]


class ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(object):
    def __init__(self, plugin_instance: str = None):
        self.plugin_instance = plugin_instance

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            server_tls_policy_pb2.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance()
        )
        if Primitive.to_proto(resource.plugin_instance):
            res.plugin_instance = Primitive.to_proto(resource.plugin_instance)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(
            plugin_instance=Primitive.from_proto(resource.plugin_instance),
        )


class ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance.from_proto(
                i
            )
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
