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
    client_tls_policy_pb2,
)
from google3.cloud.graphite.mmv2.services.google.network_security import (
    client_tls_policy_pb2_grpc,
)

from typing import List


class ClientTlsPolicy(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        sni: str = None,
        client_certificate: dict = None,
        server_validation_ca: list = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.labels = labels
        self.sni = sni
        self.client_certificate = client_certificate
        self.server_validation_ca = server_validation_ca
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = (
            client_tls_policy_pb2_grpc.NetworksecurityAlphaClientTlsPolicyServiceStub(
                channel.Channel()
            )
        )
        request = (
            client_tls_policy_pb2.ApplyNetworksecurityAlphaClientTlsPolicyRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.sni):
            request.resource.sni = Primitive.to_proto(self.sni)

        if ClientTlsPolicyClientCertificate.to_proto(self.client_certificate):
            request.resource.client_certificate.CopyFrom(
                ClientTlsPolicyClientCertificate.to_proto(self.client_certificate)
            )
        else:
            request.resource.ClearField("client_certificate")
        if ClientTlsPolicyServerValidationCaArray.to_proto(self.server_validation_ca):
            request.resource.server_validation_ca.extend(
                ClientTlsPolicyServerValidationCaArray.to_proto(
                    self.server_validation_ca
                )
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworksecurityAlphaClientTlsPolicy(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.sni = Primitive.from_proto(response.sni)
        self.client_certificate = ClientTlsPolicyClientCertificate.from_proto(
            response.client_certificate
        )
        self.server_validation_ca = ClientTlsPolicyServerValidationCaArray.from_proto(
            response.server_validation_ca
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = (
            client_tls_policy_pb2_grpc.NetworksecurityAlphaClientTlsPolicyServiceStub(
                channel.Channel()
            )
        )
        request = (
            client_tls_policy_pb2.DeleteNetworksecurityAlphaClientTlsPolicyRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.sni):
            request.resource.sni = Primitive.to_proto(self.sni)

        if ClientTlsPolicyClientCertificate.to_proto(self.client_certificate):
            request.resource.client_certificate.CopyFrom(
                ClientTlsPolicyClientCertificate.to_proto(self.client_certificate)
            )
        else:
            request.resource.ClearField("client_certificate")
        if ClientTlsPolicyServerValidationCaArray.to_proto(self.server_validation_ca):
            request.resource.server_validation_ca.extend(
                ClientTlsPolicyServerValidationCaArray.to_proto(
                    self.server_validation_ca
                )
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworksecurityAlphaClientTlsPolicy(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = (
            client_tls_policy_pb2_grpc.NetworksecurityAlphaClientTlsPolicyServiceStub(
                channel.Channel()
            )
        )
        request = client_tls_policy_pb2.ListNetworksecurityAlphaClientTlsPolicyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworksecurityAlphaClientTlsPolicy(request).items

    def to_proto(self):
        resource = client_tls_policy_pb2.NetworksecurityAlphaClientTlsPolicy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.sni):
            resource.sni = Primitive.to_proto(self.sni)
        if ClientTlsPolicyClientCertificate.to_proto(self.client_certificate):
            resource.client_certificate.CopyFrom(
                ClientTlsPolicyClientCertificate.to_proto(self.client_certificate)
            )
        else:
            resource.ClearField("client_certificate")
        if ClientTlsPolicyServerValidationCaArray.to_proto(self.server_validation_ca):
            resource.server_validation_ca.extend(
                ClientTlsPolicyServerValidationCaArray.to_proto(
                    self.server_validation_ca
                )
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ClientTlsPolicyClientCertificate(object):
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
            client_tls_policy_pb2.NetworksecurityAlphaClientTlsPolicyClientCertificate()
        )
        if ClientTlsPolicyClientCertificateLocalFilepath.to_proto(
            resource.local_filepath
        ):
            res.local_filepath.CopyFrom(
                ClientTlsPolicyClientCertificateLocalFilepath.to_proto(
                    resource.local_filepath
                )
            )
        else:
            res.ClearField("local_filepath")
        if ClientTlsPolicyClientCertificateGrpcEndpoint.to_proto(
            resource.grpc_endpoint
        ):
            res.grpc_endpoint.CopyFrom(
                ClientTlsPolicyClientCertificateGrpcEndpoint.to_proto(
                    resource.grpc_endpoint
                )
            )
        else:
            res.ClearField("grpc_endpoint")
        if ClientTlsPolicyClientCertificateCertificateProviderInstance.to_proto(
            resource.certificate_provider_instance
        ):
            res.certificate_provider_instance.CopyFrom(
                ClientTlsPolicyClientCertificateCertificateProviderInstance.to_proto(
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

        return ClientTlsPolicyClientCertificate(
            local_filepath=ClientTlsPolicyClientCertificateLocalFilepath.from_proto(
                resource.local_filepath
            ),
            grpc_endpoint=ClientTlsPolicyClientCertificateGrpcEndpoint.from_proto(
                resource.grpc_endpoint
            ),
            certificate_provider_instance=ClientTlsPolicyClientCertificateCertificateProviderInstance.from_proto(
                resource.certificate_provider_instance
            ),
        )


class ClientTlsPolicyClientCertificateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClientTlsPolicyClientCertificate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClientTlsPolicyClientCertificate.from_proto(i) for i in resources]


class ClientTlsPolicyClientCertificateLocalFilepath(object):
    def __init__(self, certificate_path: str = None, private_key_path: str = None):
        self.certificate_path = certificate_path
        self.private_key_path = private_key_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            client_tls_policy_pb2.NetworksecurityAlphaClientTlsPolicyClientCertificateLocalFilepath()
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

        return ClientTlsPolicyClientCertificateLocalFilepath(
            certificate_path=Primitive.from_proto(resource.certificate_path),
            private_key_path=Primitive.from_proto(resource.private_key_path),
        )


class ClientTlsPolicyClientCertificateLocalFilepathArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClientTlsPolicyClientCertificateLocalFilepath.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClientTlsPolicyClientCertificateLocalFilepath.from_proto(i)
            for i in resources
        ]


class ClientTlsPolicyClientCertificateGrpcEndpoint(object):
    def __init__(self, target_uri: str = None):
        self.target_uri = target_uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            client_tls_policy_pb2.NetworksecurityAlphaClientTlsPolicyClientCertificateGrpcEndpoint()
        )
        if Primitive.to_proto(resource.target_uri):
            res.target_uri = Primitive.to_proto(resource.target_uri)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClientTlsPolicyClientCertificateGrpcEndpoint(
            target_uri=Primitive.from_proto(resource.target_uri),
        )


class ClientTlsPolicyClientCertificateGrpcEndpointArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClientTlsPolicyClientCertificateGrpcEndpoint.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClientTlsPolicyClientCertificateGrpcEndpoint.from_proto(i)
            for i in resources
        ]


class ClientTlsPolicyClientCertificateCertificateProviderInstance(object):
    def __init__(self, plugin_instance: str = None):
        self.plugin_instance = plugin_instance

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            client_tls_policy_pb2.NetworksecurityAlphaClientTlsPolicyClientCertificateCertificateProviderInstance()
        )
        if Primitive.to_proto(resource.plugin_instance):
            res.plugin_instance = Primitive.to_proto(resource.plugin_instance)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClientTlsPolicyClientCertificateCertificateProviderInstance(
            plugin_instance=Primitive.from_proto(resource.plugin_instance),
        )


class ClientTlsPolicyClientCertificateCertificateProviderInstanceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClientTlsPolicyClientCertificateCertificateProviderInstance.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClientTlsPolicyClientCertificateCertificateProviderInstance.from_proto(i)
            for i in resources
        ]


class ClientTlsPolicyServerValidationCa(object):
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
            client_tls_policy_pb2.NetworksecurityAlphaClientTlsPolicyServerValidationCa()
        )
        if Primitive.to_proto(resource.ca_cert_path):
            res.ca_cert_path = Primitive.to_proto(resource.ca_cert_path)
        if ClientTlsPolicyServerValidationCaGrpcEndpoint.to_proto(
            resource.grpc_endpoint
        ):
            res.grpc_endpoint.CopyFrom(
                ClientTlsPolicyServerValidationCaGrpcEndpoint.to_proto(
                    resource.grpc_endpoint
                )
            )
        else:
            res.ClearField("grpc_endpoint")
        if ClientTlsPolicyServerValidationCaCertificateProviderInstance.to_proto(
            resource.certificate_provider_instance
        ):
            res.certificate_provider_instance.CopyFrom(
                ClientTlsPolicyServerValidationCaCertificateProviderInstance.to_proto(
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

        return ClientTlsPolicyServerValidationCa(
            ca_cert_path=Primitive.from_proto(resource.ca_cert_path),
            grpc_endpoint=ClientTlsPolicyServerValidationCaGrpcEndpoint.from_proto(
                resource.grpc_endpoint
            ),
            certificate_provider_instance=ClientTlsPolicyServerValidationCaCertificateProviderInstance.from_proto(
                resource.certificate_provider_instance
            ),
        )


class ClientTlsPolicyServerValidationCaArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClientTlsPolicyServerValidationCa.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClientTlsPolicyServerValidationCa.from_proto(i) for i in resources]


class ClientTlsPolicyServerValidationCaGrpcEndpoint(object):
    def __init__(self, target_uri: str = None):
        self.target_uri = target_uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            client_tls_policy_pb2.NetworksecurityAlphaClientTlsPolicyServerValidationCaGrpcEndpoint()
        )
        if Primitive.to_proto(resource.target_uri):
            res.target_uri = Primitive.to_proto(resource.target_uri)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClientTlsPolicyServerValidationCaGrpcEndpoint(
            target_uri=Primitive.from_proto(resource.target_uri),
        )


class ClientTlsPolicyServerValidationCaGrpcEndpointArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClientTlsPolicyServerValidationCaGrpcEndpoint.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClientTlsPolicyServerValidationCaGrpcEndpoint.from_proto(i)
            for i in resources
        ]


class ClientTlsPolicyServerValidationCaCertificateProviderInstance(object):
    def __init__(self, plugin_instance: str = None):
        self.plugin_instance = plugin_instance

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            client_tls_policy_pb2.NetworksecurityAlphaClientTlsPolicyServerValidationCaCertificateProviderInstance()
        )
        if Primitive.to_proto(resource.plugin_instance):
            res.plugin_instance = Primitive.to_proto(resource.plugin_instance)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClientTlsPolicyServerValidationCaCertificateProviderInstance(
            plugin_instance=Primitive.from_proto(resource.plugin_instance),
        )


class ClientTlsPolicyServerValidationCaCertificateProviderInstanceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClientTlsPolicyServerValidationCaCertificateProviderInstance.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClientTlsPolicyServerValidationCaCertificateProviderInstance.from_proto(i)
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
