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
from google3.cloud.graphite.mmv2.services.google.privateca import certificate_pb2
from google3.cloud.graphite.mmv2.services.google.privateca import certificate_pb2_grpc

from typing import List


class Certificate(object):
    def __init__(
        self,
        name: str = None,
        pem_csr: str = None,
        config: dict = None,
        issuer_certificate_authority: str = None,
        lifetime: str = None,
        certificate_template: str = None,
        subject_mode: str = None,
        revocation_details: dict = None,
        pem_certificate: str = None,
        certificate_description: dict = None,
        pem_certificate_chain: list = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        project: str = None,
        location: str = None,
        ca_pool: str = None,
        certificate_authority: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.pem_csr = pem_csr
        self.config = config
        self.lifetime = lifetime
        self.certificate_template = certificate_template
        self.subject_mode = subject_mode
        self.labels = labels
        self.project = project
        self.location = location
        self.ca_pool = ca_pool
        self.certificate_authority = certificate_authority
        self.service_account_file = service_account_file

    def apply(self):
        stub = certificate_pb2_grpc.PrivatecaCertificateServiceStub(channel.Channel())
        request = certificate_pb2.ApplyPrivatecaCertificateRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.pem_csr):
            request.resource.pem_csr = Primitive.to_proto(self.pem_csr)

        if CertificateConfig.to_proto(self.config):
            request.resource.config.CopyFrom(CertificateConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.lifetime):
            request.resource.lifetime = Primitive.to_proto(self.lifetime)

        if Primitive.to_proto(self.certificate_template):
            request.resource.certificate_template = Primitive.to_proto(
                self.certificate_template
            )

        if CertificateSubjectModeEnum.to_proto(self.subject_mode):
            request.resource.subject_mode = CertificateSubjectModeEnum.to_proto(
                self.subject_mode
            )

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.ca_pool):
            request.resource.ca_pool = Primitive.to_proto(self.ca_pool)

        if Primitive.to_proto(self.certificate_authority):
            request.resource.certificate_authority = Primitive.to_proto(
                self.certificate_authority
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyPrivatecaCertificate(request)
        self.name = Primitive.from_proto(response.name)
        self.pem_csr = Primitive.from_proto(response.pem_csr)
        self.config = CertificateConfig.from_proto(response.config)
        self.issuer_certificate_authority = Primitive.from_proto(
            response.issuer_certificate_authority
        )
        self.lifetime = Primitive.from_proto(response.lifetime)
        self.certificate_template = Primitive.from_proto(response.certificate_template)
        self.subject_mode = CertificateSubjectModeEnum.from_proto(response.subject_mode)
        self.revocation_details = CertificateRevocationDetails.from_proto(
            response.revocation_details
        )
        self.pem_certificate = Primitive.from_proto(response.pem_certificate)
        self.certificate_description = CertificateCertificateDescription.from_proto(
            response.certificate_description
        )
        self.pem_certificate_chain = Primitive.from_proto(
            response.pem_certificate_chain
        )
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.ca_pool = Primitive.from_proto(response.ca_pool)
        self.certificate_authority = Primitive.from_proto(
            response.certificate_authority
        )

    def delete(self):
        stub = certificate_pb2_grpc.PrivatecaCertificateServiceStub(channel.Channel())
        request = certificate_pb2.DeletePrivatecaCertificateRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.pem_csr):
            request.resource.pem_csr = Primitive.to_proto(self.pem_csr)

        if CertificateConfig.to_proto(self.config):
            request.resource.config.CopyFrom(CertificateConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.lifetime):
            request.resource.lifetime = Primitive.to_proto(self.lifetime)

        if Primitive.to_proto(self.certificate_template):
            request.resource.certificate_template = Primitive.to_proto(
                self.certificate_template
            )

        if CertificateSubjectModeEnum.to_proto(self.subject_mode):
            request.resource.subject_mode = CertificateSubjectModeEnum.to_proto(
                self.subject_mode
            )

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.ca_pool):
            request.resource.ca_pool = Primitive.to_proto(self.ca_pool)

        if Primitive.to_proto(self.certificate_authority):
            request.resource.certificate_authority = Primitive.to_proto(
                self.certificate_authority
            )

        response = stub.DeletePrivatecaCertificate(request)

    @classmethod
    def list(self, project, location, caPool, service_account_file=""):
        stub = certificate_pb2_grpc.PrivatecaCertificateServiceStub(channel.Channel())
        request = certificate_pb2.ListPrivatecaCertificateRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.CaPool = caPool

        return stub.ListPrivatecaCertificate(request).items

    def to_proto(self):
        resource = certificate_pb2.PrivatecaCertificate()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.pem_csr):
            resource.pem_csr = Primitive.to_proto(self.pem_csr)
        if CertificateConfig.to_proto(self.config):
            resource.config.CopyFrom(CertificateConfig.to_proto(self.config))
        else:
            resource.ClearField("config")
        if Primitive.to_proto(self.lifetime):
            resource.lifetime = Primitive.to_proto(self.lifetime)
        if Primitive.to_proto(self.certificate_template):
            resource.certificate_template = Primitive.to_proto(
                self.certificate_template
            )
        if CertificateSubjectModeEnum.to_proto(self.subject_mode):
            resource.subject_mode = CertificateSubjectModeEnum.to_proto(
                self.subject_mode
            )
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.ca_pool):
            resource.ca_pool = Primitive.to_proto(self.ca_pool)
        if Primitive.to_proto(self.certificate_authority):
            resource.certificate_authority = Primitive.to_proto(
                self.certificate_authority
            )
        return resource


class CertificateConfig(object):
    def __init__(
        self,
        subject_config: dict = None,
        x509_config: dict = None,
        public_key: dict = None,
    ):
        self.subject_config = subject_config
        self.x509_config = x509_config
        self.public_key = public_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateConfig()
        if CertificateConfigSubjectConfig.to_proto(resource.subject_config):
            res.subject_config.CopyFrom(
                CertificateConfigSubjectConfig.to_proto(resource.subject_config)
            )
        else:
            res.ClearField("subject_config")
        if CertificateConfigX509Config.to_proto(resource.x509_config):
            res.x509_config.CopyFrom(
                CertificateConfigX509Config.to_proto(resource.x509_config)
            )
        else:
            res.ClearField("x509_config")
        if CertificateConfigPublicKey.to_proto(resource.public_key):
            res.public_key.CopyFrom(
                CertificateConfigPublicKey.to_proto(resource.public_key)
            )
        else:
            res.ClearField("public_key")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfig(
            subject_config=CertificateConfigSubjectConfig.from_proto(
                resource.subject_config
            ),
            x509_config=CertificateConfigX509Config.from_proto(resource.x509_config),
            public_key=CertificateConfigPublicKey.from_proto(resource.public_key),
        )


class CertificateConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateConfig.from_proto(i) for i in resources]


class CertificateConfigSubjectConfig(object):
    def __init__(self, subject: dict = None, subject_alt_name: dict = None):
        self.subject = subject
        self.subject_alt_name = subject_alt_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateConfigSubjectConfig()
        if CertificateConfigSubjectConfigSubject.to_proto(resource.subject):
            res.subject.CopyFrom(
                CertificateConfigSubjectConfigSubject.to_proto(resource.subject)
            )
        else:
            res.ClearField("subject")
        if CertificateConfigSubjectConfigSubjectAltName.to_proto(
            resource.subject_alt_name
        ):
            res.subject_alt_name.CopyFrom(
                CertificateConfigSubjectConfigSubjectAltName.to_proto(
                    resource.subject_alt_name
                )
            )
        else:
            res.ClearField("subject_alt_name")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigSubjectConfig(
            subject=CertificateConfigSubjectConfigSubject.from_proto(resource.subject),
            subject_alt_name=CertificateConfigSubjectConfigSubjectAltName.from_proto(
                resource.subject_alt_name
            ),
        )


class CertificateConfigSubjectConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateConfigSubjectConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateConfigSubjectConfig.from_proto(i) for i in resources]


class CertificateConfigSubjectConfigSubject(object):
    def __init__(
        self,
        common_name: str = None,
        country_code: str = None,
        organization: str = None,
        organizational_unit: str = None,
        locality: str = None,
        province: str = None,
        street_address: str = None,
        postal_code: str = None,
    ):
        self.common_name = common_name
        self.country_code = country_code
        self.organization = organization
        self.organizational_unit = organizational_unit
        self.locality = locality
        self.province = province
        self.street_address = street_address
        self.postal_code = postal_code

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateConfigSubjectConfigSubject()
        if Primitive.to_proto(resource.common_name):
            res.common_name = Primitive.to_proto(resource.common_name)
        if Primitive.to_proto(resource.country_code):
            res.country_code = Primitive.to_proto(resource.country_code)
        if Primitive.to_proto(resource.organization):
            res.organization = Primitive.to_proto(resource.organization)
        if Primitive.to_proto(resource.organizational_unit):
            res.organizational_unit = Primitive.to_proto(resource.organizational_unit)
        if Primitive.to_proto(resource.locality):
            res.locality = Primitive.to_proto(resource.locality)
        if Primitive.to_proto(resource.province):
            res.province = Primitive.to_proto(resource.province)
        if Primitive.to_proto(resource.street_address):
            res.street_address = Primitive.to_proto(resource.street_address)
        if Primitive.to_proto(resource.postal_code):
            res.postal_code = Primitive.to_proto(resource.postal_code)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigSubjectConfigSubject(
            common_name=Primitive.from_proto(resource.common_name),
            country_code=Primitive.from_proto(resource.country_code),
            organization=Primitive.from_proto(resource.organization),
            organizational_unit=Primitive.from_proto(resource.organizational_unit),
            locality=Primitive.from_proto(resource.locality),
            province=Primitive.from_proto(resource.province),
            street_address=Primitive.from_proto(resource.street_address),
            postal_code=Primitive.from_proto(resource.postal_code),
        )


class CertificateConfigSubjectConfigSubjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateConfigSubjectConfigSubject.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateConfigSubjectConfigSubject.from_proto(i) for i in resources]


class CertificateConfigSubjectConfigSubjectAltName(object):
    def __init__(
        self,
        dns_names: list = None,
        uris: list = None,
        email_addresses: list = None,
        ip_addresses: list = None,
    ):
        self.dns_names = dns_names
        self.uris = uris
        self.email_addresses = email_addresses
        self.ip_addresses = ip_addresses

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateConfigSubjectConfigSubjectAltName()
        if Primitive.to_proto(resource.dns_names):
            res.dns_names.extend(Primitive.to_proto(resource.dns_names))
        if Primitive.to_proto(resource.uris):
            res.uris.extend(Primitive.to_proto(resource.uris))
        if Primitive.to_proto(resource.email_addresses):
            res.email_addresses.extend(Primitive.to_proto(resource.email_addresses))
        if Primitive.to_proto(resource.ip_addresses):
            res.ip_addresses.extend(Primitive.to_proto(resource.ip_addresses))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigSubjectConfigSubjectAltName(
            dns_names=Primitive.from_proto(resource.dns_names),
            uris=Primitive.from_proto(resource.uris),
            email_addresses=Primitive.from_proto(resource.email_addresses),
            ip_addresses=Primitive.from_proto(resource.ip_addresses),
        )


class CertificateConfigSubjectConfigSubjectAltNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateConfigSubjectConfigSubjectAltName.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateConfigSubjectConfigSubjectAltName.from_proto(i)
            for i in resources
        ]


class CertificateConfigX509Config(object):
    def __init__(
        self,
        key_usage: dict = None,
        ca_options: dict = None,
        policy_ids: list = None,
        aia_ocsp_servers: list = None,
        additional_extensions: list = None,
    ):
        self.key_usage = key_usage
        self.ca_options = ca_options
        self.policy_ids = policy_ids
        self.aia_ocsp_servers = aia_ocsp_servers
        self.additional_extensions = additional_extensions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateConfigX509Config()
        if CertificateConfigX509ConfigKeyUsage.to_proto(resource.key_usage):
            res.key_usage.CopyFrom(
                CertificateConfigX509ConfigKeyUsage.to_proto(resource.key_usage)
            )
        else:
            res.ClearField("key_usage")
        if CertificateConfigX509ConfigCaOptions.to_proto(resource.ca_options):
            res.ca_options.CopyFrom(
                CertificateConfigX509ConfigCaOptions.to_proto(resource.ca_options)
            )
        else:
            res.ClearField("ca_options")
        if CertificateConfigX509ConfigPolicyIdsArray.to_proto(resource.policy_ids):
            res.policy_ids.extend(
                CertificateConfigX509ConfigPolicyIdsArray.to_proto(resource.policy_ids)
            )
        if Primitive.to_proto(resource.aia_ocsp_servers):
            res.aia_ocsp_servers.extend(Primitive.to_proto(resource.aia_ocsp_servers))
        if CertificateConfigX509ConfigAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CertificateConfigX509ConfigAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigX509Config(
            key_usage=CertificateConfigX509ConfigKeyUsage.from_proto(
                resource.key_usage
            ),
            ca_options=CertificateConfigX509ConfigCaOptions.from_proto(
                resource.ca_options
            ),
            policy_ids=CertificateConfigX509ConfigPolicyIdsArray.from_proto(
                resource.policy_ids
            ),
            aia_ocsp_servers=Primitive.from_proto(resource.aia_ocsp_servers),
            additional_extensions=CertificateConfigX509ConfigAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CertificateConfigX509ConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateConfigX509Config.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateConfigX509Config.from_proto(i) for i in resources]


class CertificateConfigX509ConfigKeyUsage(object):
    def __init__(
        self,
        base_key_usage: dict = None,
        extended_key_usage: dict = None,
        unknown_extended_key_usages: list = None,
    ):
        self.base_key_usage = base_key_usage
        self.extended_key_usage = extended_key_usage
        self.unknown_extended_key_usages = unknown_extended_key_usages

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateConfigX509ConfigKeyUsage()
        if CertificateConfigX509ConfigKeyUsageBaseKeyUsage.to_proto(
            resource.base_key_usage
        ):
            res.base_key_usage.CopyFrom(
                CertificateConfigX509ConfigKeyUsageBaseKeyUsage.to_proto(
                    resource.base_key_usage
                )
            )
        else:
            res.ClearField("base_key_usage")
        if CertificateConfigX509ConfigKeyUsageExtendedKeyUsage.to_proto(
            resource.extended_key_usage
        ):
            res.extended_key_usage.CopyFrom(
                CertificateConfigX509ConfigKeyUsageExtendedKeyUsage.to_proto(
                    resource.extended_key_usage
                )
            )
        else:
            res.ClearField("extended_key_usage")
        if CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
            resource.unknown_extended_key_usages
        ):
            res.unknown_extended_key_usages.extend(
                CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
                    resource.unknown_extended_key_usages
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigX509ConfigKeyUsage(
            base_key_usage=CertificateConfigX509ConfigKeyUsageBaseKeyUsage.from_proto(
                resource.base_key_usage
            ),
            extended_key_usage=CertificateConfigX509ConfigKeyUsageExtendedKeyUsage.from_proto(
                resource.extended_key_usage
            ),
            unknown_extended_key_usages=CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesArray.from_proto(
                resource.unknown_extended_key_usages
            ),
        )


class CertificateConfigX509ConfigKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateConfigX509ConfigKeyUsage.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateConfigX509ConfigKeyUsage.from_proto(i) for i in resources]


class CertificateConfigX509ConfigKeyUsageBaseKeyUsage(object):
    def __init__(
        self,
        digital_signature: bool = None,
        content_commitment: bool = None,
        key_encipherment: bool = None,
        data_encipherment: bool = None,
        key_agreement: bool = None,
        cert_sign: bool = None,
        crl_sign: bool = None,
        encipher_only: bool = None,
        decipher_only: bool = None,
    ):
        self.digital_signature = digital_signature
        self.content_commitment = content_commitment
        self.key_encipherment = key_encipherment
        self.data_encipherment = data_encipherment
        self.key_agreement = key_agreement
        self.cert_sign = cert_sign
        self.crl_sign = crl_sign
        self.encipher_only = encipher_only
        self.decipher_only = decipher_only

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage()
        if Primitive.to_proto(resource.digital_signature):
            res.digital_signature = Primitive.to_proto(resource.digital_signature)
        if Primitive.to_proto(resource.content_commitment):
            res.content_commitment = Primitive.to_proto(resource.content_commitment)
        if Primitive.to_proto(resource.key_encipherment):
            res.key_encipherment = Primitive.to_proto(resource.key_encipherment)
        if Primitive.to_proto(resource.data_encipherment):
            res.data_encipherment = Primitive.to_proto(resource.data_encipherment)
        if Primitive.to_proto(resource.key_agreement):
            res.key_agreement = Primitive.to_proto(resource.key_agreement)
        if Primitive.to_proto(resource.cert_sign):
            res.cert_sign = Primitive.to_proto(resource.cert_sign)
        if Primitive.to_proto(resource.crl_sign):
            res.crl_sign = Primitive.to_proto(resource.crl_sign)
        if Primitive.to_proto(resource.encipher_only):
            res.encipher_only = Primitive.to_proto(resource.encipher_only)
        if Primitive.to_proto(resource.decipher_only):
            res.decipher_only = Primitive.to_proto(resource.decipher_only)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigX509ConfigKeyUsageBaseKeyUsage(
            digital_signature=Primitive.from_proto(resource.digital_signature),
            content_commitment=Primitive.from_proto(resource.content_commitment),
            key_encipherment=Primitive.from_proto(resource.key_encipherment),
            data_encipherment=Primitive.from_proto(resource.data_encipherment),
            key_agreement=Primitive.from_proto(resource.key_agreement),
            cert_sign=Primitive.from_proto(resource.cert_sign),
            crl_sign=Primitive.from_proto(resource.crl_sign),
            encipher_only=Primitive.from_proto(resource.encipher_only),
            decipher_only=Primitive.from_proto(resource.decipher_only),
        )


class CertificateConfigX509ConfigKeyUsageBaseKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateConfigX509ConfigKeyUsageBaseKeyUsage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateConfigX509ConfigKeyUsageBaseKeyUsage.from_proto(i)
            for i in resources
        ]


class CertificateConfigX509ConfigKeyUsageExtendedKeyUsage(object):
    def __init__(
        self,
        server_auth: bool = None,
        client_auth: bool = None,
        code_signing: bool = None,
        email_protection: bool = None,
        time_stamping: bool = None,
        ocsp_signing: bool = None,
    ):
        self.server_auth = server_auth
        self.client_auth = client_auth
        self.code_signing = code_signing
        self.email_protection = email_protection
        self.time_stamping = time_stamping
        self.ocsp_signing = ocsp_signing

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage()
        )
        if Primitive.to_proto(resource.server_auth):
            res.server_auth = Primitive.to_proto(resource.server_auth)
        if Primitive.to_proto(resource.client_auth):
            res.client_auth = Primitive.to_proto(resource.client_auth)
        if Primitive.to_proto(resource.code_signing):
            res.code_signing = Primitive.to_proto(resource.code_signing)
        if Primitive.to_proto(resource.email_protection):
            res.email_protection = Primitive.to_proto(resource.email_protection)
        if Primitive.to_proto(resource.time_stamping):
            res.time_stamping = Primitive.to_proto(resource.time_stamping)
        if Primitive.to_proto(resource.ocsp_signing):
            res.ocsp_signing = Primitive.to_proto(resource.ocsp_signing)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigX509ConfigKeyUsageExtendedKeyUsage(
            server_auth=Primitive.from_proto(resource.server_auth),
            client_auth=Primitive.from_proto(resource.client_auth),
            code_signing=Primitive.from_proto(resource.code_signing),
            email_protection=Primitive.from_proto(resource.email_protection),
            time_stamping=Primitive.from_proto(resource.time_stamping),
            ocsp_signing=Primitive.from_proto(resource.ocsp_signing),
        )


class CertificateConfigX509ConfigKeyUsageExtendedKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateConfigX509ConfigKeyUsageExtendedKeyUsage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateConfigX509ConfigKeyUsageExtendedKeyUsage.from_proto(i)
            for i in resources
        ]


class CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages.from_proto(i)
            for i in resources
        ]


class CertificateConfigX509ConfigCaOptions(object):
    def __init__(
        self,
        is_ca: bool = None,
        non_ca: bool = None,
        max_issuer_path_length: int = None,
        zero_max_issuer_path_length: bool = None,
    ):
        self.is_ca = is_ca
        self.non_ca = non_ca
        self.max_issuer_path_length = max_issuer_path_length
        self.zero_max_issuer_path_length = zero_max_issuer_path_length

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateConfigX509ConfigCaOptions()
        if Primitive.to_proto(resource.is_ca):
            res.is_ca = Primitive.to_proto(resource.is_ca)
        if Primitive.to_proto(resource.non_ca):
            res.non_ca = Primitive.to_proto(resource.non_ca)
        if Primitive.to_proto(resource.max_issuer_path_length):
            res.max_issuer_path_length = Primitive.to_proto(
                resource.max_issuer_path_length
            )
        if Primitive.to_proto(resource.zero_max_issuer_path_length):
            res.zero_max_issuer_path_length = Primitive.to_proto(
                resource.zero_max_issuer_path_length
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigX509ConfigCaOptions(
            is_ca=Primitive.from_proto(resource.is_ca),
            non_ca=Primitive.from_proto(resource.non_ca),
            max_issuer_path_length=Primitive.from_proto(
                resource.max_issuer_path_length
            ),
            zero_max_issuer_path_length=Primitive.from_proto(
                resource.zero_max_issuer_path_length
            ),
        )


class CertificateConfigX509ConfigCaOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateConfigX509ConfigCaOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateConfigX509ConfigCaOptions.from_proto(i) for i in resources]


class CertificateConfigX509ConfigPolicyIds(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateConfigX509ConfigPolicyIds()
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigX509ConfigPolicyIds(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateConfigX509ConfigPolicyIdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateConfigX509ConfigPolicyIds.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateConfigX509ConfigPolicyIds.from_proto(i) for i in resources]


class CertificateConfigX509ConfigAdditionalExtensions(object):
    def __init__(
        self, object_id: dict = None, critical: bool = None, value: str = None
    ):
        self.object_id = object_id
        self.critical = critical
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateConfigX509ConfigAdditionalExtensions()
        if CertificateConfigX509ConfigAdditionalExtensionsObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateConfigX509ConfigAdditionalExtensionsObjectId.to_proto(
                    resource.object_id
                )
            )
        else:
            res.ClearField("object_id")
        if Primitive.to_proto(resource.critical):
            res.critical = Primitive.to_proto(resource.critical)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigX509ConfigAdditionalExtensions(
            object_id=CertificateConfigX509ConfigAdditionalExtensionsObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateConfigX509ConfigAdditionalExtensionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateConfigX509ConfigAdditionalExtensions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateConfigX509ConfigAdditionalExtensions.from_proto(i)
            for i in resources
        ]


class CertificateConfigX509ConfigAdditionalExtensionsObjectId(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigX509ConfigAdditionalExtensionsObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateConfigX509ConfigAdditionalExtensionsObjectIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateConfigX509ConfigAdditionalExtensionsObjectId.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateConfigX509ConfigAdditionalExtensionsObjectId.from_proto(i)
            for i in resources
        ]


class CertificateConfigPublicKey(object):
    def __init__(self, key: str = None, format: str = None):
        self.key = key
        self.format = format

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateConfigPublicKey()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if CertificateConfigPublicKeyFormatEnum.to_proto(resource.format):
            res.format = CertificateConfigPublicKeyFormatEnum.to_proto(resource.format)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateConfigPublicKey(
            key=Primitive.from_proto(resource.key),
            format=CertificateConfigPublicKeyFormatEnum.from_proto(resource.format),
        )


class CertificateConfigPublicKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateConfigPublicKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateConfigPublicKey.from_proto(i) for i in resources]


class CertificateRevocationDetails(object):
    def __init__(self, revocation_state: str = None, revocation_time: str = None):
        self.revocation_state = revocation_state
        self.revocation_time = revocation_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateRevocationDetails()
        if CertificateRevocationDetailsRevocationStateEnum.to_proto(
            resource.revocation_state
        ):
            res.revocation_state = (
                CertificateRevocationDetailsRevocationStateEnum.to_proto(
                    resource.revocation_state
                )
            )
        if Primitive.to_proto(resource.revocation_time):
            res.revocation_time = Primitive.to_proto(resource.revocation_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateRevocationDetails(
            revocation_state=CertificateRevocationDetailsRevocationStateEnum.from_proto(
                resource.revocation_state
            ),
            revocation_time=Primitive.from_proto(resource.revocation_time),
        )


class CertificateRevocationDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateRevocationDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateRevocationDetails.from_proto(i) for i in resources]


class CertificateCertificateDescription(object):
    def __init__(
        self,
        subject_description: dict = None,
        x509_description: dict = None,
        public_key: dict = None,
        subject_key_id: dict = None,
        authority_key_id: dict = None,
        crl_distribution_points: list = None,
        aia_issuing_certificate_urls: list = None,
        cert_fingerprint: dict = None,
    ):
        self.subject_description = subject_description
        self.x509_description = x509_description
        self.public_key = public_key
        self.subject_key_id = subject_key_id
        self.authority_key_id = authority_key_id
        self.crl_distribution_points = crl_distribution_points
        self.aia_issuing_certificate_urls = aia_issuing_certificate_urls
        self.cert_fingerprint = cert_fingerprint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateCertificateDescription()
        if CertificateCertificateDescriptionSubjectDescription.to_proto(
            resource.subject_description
        ):
            res.subject_description.CopyFrom(
                CertificateCertificateDescriptionSubjectDescription.to_proto(
                    resource.subject_description
                )
            )
        else:
            res.ClearField("subject_description")
        if CertificateCertificateDescriptionX509Description.to_proto(
            resource.x509_description
        ):
            res.x509_description.CopyFrom(
                CertificateCertificateDescriptionX509Description.to_proto(
                    resource.x509_description
                )
            )
        else:
            res.ClearField("x509_description")
        if CertificateCertificateDescriptionPublicKey.to_proto(resource.public_key):
            res.public_key.CopyFrom(
                CertificateCertificateDescriptionPublicKey.to_proto(resource.public_key)
            )
        else:
            res.ClearField("public_key")
        if CertificateCertificateDescriptionSubjectKeyId.to_proto(
            resource.subject_key_id
        ):
            res.subject_key_id.CopyFrom(
                CertificateCertificateDescriptionSubjectKeyId.to_proto(
                    resource.subject_key_id
                )
            )
        else:
            res.ClearField("subject_key_id")
        if CertificateCertificateDescriptionAuthorityKeyId.to_proto(
            resource.authority_key_id
        ):
            res.authority_key_id.CopyFrom(
                CertificateCertificateDescriptionAuthorityKeyId.to_proto(
                    resource.authority_key_id
                )
            )
        else:
            res.ClearField("authority_key_id")
        if Primitive.to_proto(resource.crl_distribution_points):
            res.crl_distribution_points.extend(
                Primitive.to_proto(resource.crl_distribution_points)
            )
        if Primitive.to_proto(resource.aia_issuing_certificate_urls):
            res.aia_issuing_certificate_urls.extend(
                Primitive.to_proto(resource.aia_issuing_certificate_urls)
            )
        if CertificateCertificateDescriptionCertFingerprint.to_proto(
            resource.cert_fingerprint
        ):
            res.cert_fingerprint.CopyFrom(
                CertificateCertificateDescriptionCertFingerprint.to_proto(
                    resource.cert_fingerprint
                )
            )
        else:
            res.ClearField("cert_fingerprint")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescription(
            subject_description=CertificateCertificateDescriptionSubjectDescription.from_proto(
                resource.subject_description
            ),
            x509_description=CertificateCertificateDescriptionX509Description.from_proto(
                resource.x509_description
            ),
            public_key=CertificateCertificateDescriptionPublicKey.from_proto(
                resource.public_key
            ),
            subject_key_id=CertificateCertificateDescriptionSubjectKeyId.from_proto(
                resource.subject_key_id
            ),
            authority_key_id=CertificateCertificateDescriptionAuthorityKeyId.from_proto(
                resource.authority_key_id
            ),
            crl_distribution_points=Primitive.from_proto(
                resource.crl_distribution_points
            ),
            aia_issuing_certificate_urls=Primitive.from_proto(
                resource.aia_issuing_certificate_urls
            ),
            cert_fingerprint=CertificateCertificateDescriptionCertFingerprint.from_proto(
                resource.cert_fingerprint
            ),
        )


class CertificateCertificateDescriptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateCertificateDescription.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateCertificateDescription.from_proto(i) for i in resources]


class CertificateCertificateDescriptionSubjectDescription(object):
    def __init__(
        self,
        subject: dict = None,
        subject_alt_name: dict = None,
        hex_serial_number: str = None,
        lifetime: str = None,
        not_before_time: str = None,
        not_after_time: str = None,
    ):
        self.subject = subject
        self.subject_alt_name = subject_alt_name
        self.hex_serial_number = hex_serial_number
        self.lifetime = lifetime
        self.not_before_time = not_before_time
        self.not_after_time = not_after_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionSubjectDescription()
        )
        if CertificateCertificateDescriptionSubjectDescriptionSubject.to_proto(
            resource.subject
        ):
            res.subject.CopyFrom(
                CertificateCertificateDescriptionSubjectDescriptionSubject.to_proto(
                    resource.subject
                )
            )
        else:
            res.ClearField("subject")
        if CertificateCertificateDescriptionSubjectDescriptionSubjectAltName.to_proto(
            resource.subject_alt_name
        ):
            res.subject_alt_name.CopyFrom(
                CertificateCertificateDescriptionSubjectDescriptionSubjectAltName.to_proto(
                    resource.subject_alt_name
                )
            )
        else:
            res.ClearField("subject_alt_name")
        if Primitive.to_proto(resource.hex_serial_number):
            res.hex_serial_number = Primitive.to_proto(resource.hex_serial_number)
        if Primitive.to_proto(resource.lifetime):
            res.lifetime = Primitive.to_proto(resource.lifetime)
        if Primitive.to_proto(resource.not_before_time):
            res.not_before_time = Primitive.to_proto(resource.not_before_time)
        if Primitive.to_proto(resource.not_after_time):
            res.not_after_time = Primitive.to_proto(resource.not_after_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionSubjectDescription(
            subject=CertificateCertificateDescriptionSubjectDescriptionSubject.from_proto(
                resource.subject
            ),
            subject_alt_name=CertificateCertificateDescriptionSubjectDescriptionSubjectAltName.from_proto(
                resource.subject_alt_name
            ),
            hex_serial_number=Primitive.from_proto(resource.hex_serial_number),
            lifetime=Primitive.from_proto(resource.lifetime),
            not_before_time=Primitive.from_proto(resource.not_before_time),
            not_after_time=Primitive.from_proto(resource.not_after_time),
        )


class CertificateCertificateDescriptionSubjectDescriptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionSubjectDescription.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionSubjectDescription.from_proto(i)
            for i in resources
        ]


class CertificateCertificateDescriptionSubjectDescriptionSubject(object):
    def __init__(
        self,
        common_name: str = None,
        country_code: str = None,
        organization: str = None,
        organizational_unit: str = None,
        locality: str = None,
        province: str = None,
        street_address: str = None,
        postal_code: str = None,
    ):
        self.common_name = common_name
        self.country_code = country_code
        self.organization = organization
        self.organizational_unit = organizational_unit
        self.locality = locality
        self.province = province
        self.street_address = street_address
        self.postal_code = postal_code

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubject()
        )
        if Primitive.to_proto(resource.common_name):
            res.common_name = Primitive.to_proto(resource.common_name)
        if Primitive.to_proto(resource.country_code):
            res.country_code = Primitive.to_proto(resource.country_code)
        if Primitive.to_proto(resource.organization):
            res.organization = Primitive.to_proto(resource.organization)
        if Primitive.to_proto(resource.organizational_unit):
            res.organizational_unit = Primitive.to_proto(resource.organizational_unit)
        if Primitive.to_proto(resource.locality):
            res.locality = Primitive.to_proto(resource.locality)
        if Primitive.to_proto(resource.province):
            res.province = Primitive.to_proto(resource.province)
        if Primitive.to_proto(resource.street_address):
            res.street_address = Primitive.to_proto(resource.street_address)
        if Primitive.to_proto(resource.postal_code):
            res.postal_code = Primitive.to_proto(resource.postal_code)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionSubjectDescriptionSubject(
            common_name=Primitive.from_proto(resource.common_name),
            country_code=Primitive.from_proto(resource.country_code),
            organization=Primitive.from_proto(resource.organization),
            organizational_unit=Primitive.from_proto(resource.organizational_unit),
            locality=Primitive.from_proto(resource.locality),
            province=Primitive.from_proto(resource.province),
            street_address=Primitive.from_proto(resource.street_address),
            postal_code=Primitive.from_proto(resource.postal_code),
        )


class CertificateCertificateDescriptionSubjectDescriptionSubjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionSubjectDescriptionSubject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionSubjectDescriptionSubject.from_proto(i)
            for i in resources
        ]


class CertificateCertificateDescriptionSubjectDescriptionSubjectAltName(object):
    def __init__(
        self,
        dns_names: list = None,
        uris: list = None,
        email_addresses: list = None,
        ip_addresses: list = None,
        custom_sans: list = None,
    ):
        self.dns_names = dns_names
        self.uris = uris
        self.email_addresses = email_addresses
        self.ip_addresses = ip_addresses
        self.custom_sans = custom_sans

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName()
        )
        if Primitive.to_proto(resource.dns_names):
            res.dns_names.extend(Primitive.to_proto(resource.dns_names))
        if Primitive.to_proto(resource.uris):
            res.uris.extend(Primitive.to_proto(resource.uris))
        if Primitive.to_proto(resource.email_addresses):
            res.email_addresses.extend(Primitive.to_proto(resource.email_addresses))
        if Primitive.to_proto(resource.ip_addresses):
            res.ip_addresses.extend(Primitive.to_proto(resource.ip_addresses))
        if CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansArray.to_proto(
            resource.custom_sans
        ):
            res.custom_sans.extend(
                CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansArray.to_proto(
                    resource.custom_sans
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionSubjectDescriptionSubjectAltName(
            dns_names=Primitive.from_proto(resource.dns_names),
            uris=Primitive.from_proto(resource.uris),
            email_addresses=Primitive.from_proto(resource.email_addresses),
            ip_addresses=Primitive.from_proto(resource.ip_addresses),
            custom_sans=CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansArray.from_proto(
                resource.custom_sans
            ),
        )


class CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionSubjectDescriptionSubjectAltName.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionSubjectDescriptionSubjectAltName.from_proto(
                i
            )
            for i in resources
        ]


class CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans(
    object
):
    def __init__(
        self, object_id: dict = None, critical: bool = None, value: str = None
    ):
        self.object_id = object_id
        self.critical = critical
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans()
        )
        if CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId.to_proto(
                    resource.object_id
                )
            )
        else:
            res.ClearField("object_id")
        if Primitive.to_proto(resource.critical):
            res.critical = Primitive.to_proto(resource.critical)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans(
            object_id=CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans.from_proto(
                i
            )
            for i in resources
        ]


class CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectIdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateCertificateDescriptionX509Description(object):
    def __init__(
        self,
        key_usage: dict = None,
        ca_options: dict = None,
        policy_ids: list = None,
        aia_ocsp_servers: list = None,
        additional_extensions: list = None,
    ):
        self.key_usage = key_usage
        self.ca_options = ca_options
        self.policy_ids = policy_ids
        self.aia_ocsp_servers = aia_ocsp_servers
        self.additional_extensions = additional_extensions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionX509Description()
        )
        if CertificateCertificateDescriptionX509DescriptionKeyUsage.to_proto(
            resource.key_usage
        ):
            res.key_usage.CopyFrom(
                CertificateCertificateDescriptionX509DescriptionKeyUsage.to_proto(
                    resource.key_usage
                )
            )
        else:
            res.ClearField("key_usage")
        if CertificateCertificateDescriptionX509DescriptionCaOptions.to_proto(
            resource.ca_options
        ):
            res.ca_options.CopyFrom(
                CertificateCertificateDescriptionX509DescriptionCaOptions.to_proto(
                    resource.ca_options
                )
            )
        else:
            res.ClearField("ca_options")
        if CertificateCertificateDescriptionX509DescriptionPolicyIdsArray.to_proto(
            resource.policy_ids
        ):
            res.policy_ids.extend(
                CertificateCertificateDescriptionX509DescriptionPolicyIdsArray.to_proto(
                    resource.policy_ids
                )
            )
        if Primitive.to_proto(resource.aia_ocsp_servers):
            res.aia_ocsp_servers.extend(Primitive.to_proto(resource.aia_ocsp_servers))
        if CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionX509Description(
            key_usage=CertificateCertificateDescriptionX509DescriptionKeyUsage.from_proto(
                resource.key_usage
            ),
            ca_options=CertificateCertificateDescriptionX509DescriptionCaOptions.from_proto(
                resource.ca_options
            ),
            policy_ids=CertificateCertificateDescriptionX509DescriptionPolicyIdsArray.from_proto(
                resource.policy_ids
            ),
            aia_ocsp_servers=Primitive.from_proto(resource.aia_ocsp_servers),
            additional_extensions=CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CertificateCertificateDescriptionX509DescriptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionX509Description.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionX509Description.from_proto(i)
            for i in resources
        ]


class CertificateCertificateDescriptionX509DescriptionKeyUsage(object):
    def __init__(
        self,
        base_key_usage: dict = None,
        extended_key_usage: dict = None,
        unknown_extended_key_usages: list = None,
    ):
        self.base_key_usage = base_key_usage
        self.extended_key_usage = extended_key_usage
        self.unknown_extended_key_usages = unknown_extended_key_usages

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsage()
        )
        if CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage.to_proto(
            resource.base_key_usage
        ):
            res.base_key_usage.CopyFrom(
                CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage.to_proto(
                    resource.base_key_usage
                )
            )
        else:
            res.ClearField("base_key_usage")
        if CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage.to_proto(
            resource.extended_key_usage
        ):
            res.extended_key_usage.CopyFrom(
                CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage.to_proto(
                    resource.extended_key_usage
                )
            )
        else:
            res.ClearField("extended_key_usage")
        if CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
            resource.unknown_extended_key_usages
        ):
            res.unknown_extended_key_usages.extend(
                CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
                    resource.unknown_extended_key_usages
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionX509DescriptionKeyUsage(
            base_key_usage=CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage.from_proto(
                resource.base_key_usage
            ),
            extended_key_usage=CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage.from_proto(
                resource.extended_key_usage
            ),
            unknown_extended_key_usages=CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesArray.from_proto(
                resource.unknown_extended_key_usages
            ),
        )


class CertificateCertificateDescriptionX509DescriptionKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionX509DescriptionKeyUsage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionX509DescriptionKeyUsage.from_proto(i)
            for i in resources
        ]


class CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage(object):
    def __init__(
        self,
        digital_signature: bool = None,
        content_commitment: bool = None,
        key_encipherment: bool = None,
        data_encipherment: bool = None,
        key_agreement: bool = None,
        cert_sign: bool = None,
        crl_sign: bool = None,
        encipher_only: bool = None,
        decipher_only: bool = None,
    ):
        self.digital_signature = digital_signature
        self.content_commitment = content_commitment
        self.key_encipherment = key_encipherment
        self.data_encipherment = data_encipherment
        self.key_agreement = key_agreement
        self.cert_sign = cert_sign
        self.crl_sign = crl_sign
        self.encipher_only = encipher_only
        self.decipher_only = decipher_only

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage()
        )
        if Primitive.to_proto(resource.digital_signature):
            res.digital_signature = Primitive.to_proto(resource.digital_signature)
        if Primitive.to_proto(resource.content_commitment):
            res.content_commitment = Primitive.to_proto(resource.content_commitment)
        if Primitive.to_proto(resource.key_encipherment):
            res.key_encipherment = Primitive.to_proto(resource.key_encipherment)
        if Primitive.to_proto(resource.data_encipherment):
            res.data_encipherment = Primitive.to_proto(resource.data_encipherment)
        if Primitive.to_proto(resource.key_agreement):
            res.key_agreement = Primitive.to_proto(resource.key_agreement)
        if Primitive.to_proto(resource.cert_sign):
            res.cert_sign = Primitive.to_proto(resource.cert_sign)
        if Primitive.to_proto(resource.crl_sign):
            res.crl_sign = Primitive.to_proto(resource.crl_sign)
        if Primitive.to_proto(resource.encipher_only):
            res.encipher_only = Primitive.to_proto(resource.encipher_only)
        if Primitive.to_proto(resource.decipher_only):
            res.decipher_only = Primitive.to_proto(resource.decipher_only)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage(
            digital_signature=Primitive.from_proto(resource.digital_signature),
            content_commitment=Primitive.from_proto(resource.content_commitment),
            key_encipherment=Primitive.from_proto(resource.key_encipherment),
            data_encipherment=Primitive.from_proto(resource.data_encipherment),
            key_agreement=Primitive.from_proto(resource.key_agreement),
            cert_sign=Primitive.from_proto(resource.cert_sign),
            crl_sign=Primitive.from_proto(resource.crl_sign),
            encipher_only=Primitive.from_proto(resource.encipher_only),
            decipher_only=Primitive.from_proto(resource.decipher_only),
        )


class CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage(object):
    def __init__(
        self,
        server_auth: bool = None,
        client_auth: bool = None,
        code_signing: bool = None,
        email_protection: bool = None,
        time_stamping: bool = None,
        ocsp_signing: bool = None,
    ):
        self.server_auth = server_auth
        self.client_auth = client_auth
        self.code_signing = code_signing
        self.email_protection = email_protection
        self.time_stamping = time_stamping
        self.ocsp_signing = ocsp_signing

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage()
        )
        if Primitive.to_proto(resource.server_auth):
            res.server_auth = Primitive.to_proto(resource.server_auth)
        if Primitive.to_proto(resource.client_auth):
            res.client_auth = Primitive.to_proto(resource.client_auth)
        if Primitive.to_proto(resource.code_signing):
            res.code_signing = Primitive.to_proto(resource.code_signing)
        if Primitive.to_proto(resource.email_protection):
            res.email_protection = Primitive.to_proto(resource.email_protection)
        if Primitive.to_proto(resource.time_stamping):
            res.time_stamping = Primitive.to_proto(resource.time_stamping)
        if Primitive.to_proto(resource.ocsp_signing):
            res.ocsp_signing = Primitive.to_proto(resource.ocsp_signing)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage(
            server_auth=Primitive.from_proto(resource.server_auth),
            client_auth=Primitive.from_proto(resource.client_auth),
            code_signing=Primitive.from_proto(resource.code_signing),
            email_protection=Primitive.from_proto(resource.email_protection),
            time_stamping=Primitive.from_proto(resource.time_stamping),
            ocsp_signing=Primitive.from_proto(resource.ocsp_signing),
        )


class CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages.from_proto(
                i
            )
            for i in resources
        ]


class CertificateCertificateDescriptionX509DescriptionCaOptions(object):
    def __init__(self, is_ca: bool = None, max_issuer_path_length: int = None):
        self.is_ca = is_ca
        self.max_issuer_path_length = max_issuer_path_length

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionX509DescriptionCaOptions()
        )
        if Primitive.to_proto(resource.is_ca):
            res.is_ca = Primitive.to_proto(resource.is_ca)
        if Primitive.to_proto(resource.max_issuer_path_length):
            res.max_issuer_path_length = Primitive.to_proto(
                resource.max_issuer_path_length
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionX509DescriptionCaOptions(
            is_ca=Primitive.from_proto(resource.is_ca),
            max_issuer_path_length=Primitive.from_proto(
                resource.max_issuer_path_length
            ),
        )


class CertificateCertificateDescriptionX509DescriptionCaOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionX509DescriptionCaOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionX509DescriptionCaOptions.from_proto(i)
            for i in resources
        ]


class CertificateCertificateDescriptionX509DescriptionPolicyIds(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionX509DescriptionPolicyIds()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionX509DescriptionPolicyIds(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateCertificateDescriptionX509DescriptionPolicyIdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionX509DescriptionPolicyIds.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionX509DescriptionPolicyIds.from_proto(i)
            for i in resources
        ]


class CertificateCertificateDescriptionX509DescriptionAdditionalExtensions(object):
    def __init__(
        self, object_id: dict = None, critical: bool = None, value: str = None
    ):
        self.object_id = object_id
        self.critical = critical
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions()
        )
        if CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId.to_proto(
                    resource.object_id
                )
            )
        else:
            res.ClearField("object_id")
        if Primitive.to_proto(resource.critical):
            res.critical = Primitive.to_proto(resource.critical)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionX509DescriptionAdditionalExtensions(
            object_id=CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionX509DescriptionAdditionalExtensions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionX509DescriptionAdditionalExtensions.from_proto(
                i
            )
            for i in resources
        ]


class CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectIdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateCertificateDescriptionPublicKey(object):
    def __init__(self, key: str = None, format: str = None):
        self.key = key
        self.format = format

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateCertificateDescriptionPublicKey()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if CertificateCertificateDescriptionPublicKeyFormatEnum.to_proto(
            resource.format
        ):
            res.format = CertificateCertificateDescriptionPublicKeyFormatEnum.to_proto(
                resource.format
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionPublicKey(
            key=Primitive.from_proto(resource.key),
            format=CertificateCertificateDescriptionPublicKeyFormatEnum.from_proto(
                resource.format
            ),
        )


class CertificateCertificateDescriptionPublicKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionPublicKey.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionPublicKey.from_proto(i) for i in resources
        ]


class CertificateCertificateDescriptionSubjectKeyId(object):
    def __init__(self, key_id: str = None):
        self.key_id = key_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateCertificateDescriptionSubjectKeyId()
        if Primitive.to_proto(resource.key_id):
            res.key_id = Primitive.to_proto(resource.key_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionSubjectKeyId(
            key_id=Primitive.from_proto(resource.key_id),
        )


class CertificateCertificateDescriptionSubjectKeyIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionSubjectKeyId.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionSubjectKeyId.from_proto(i)
            for i in resources
        ]


class CertificateCertificateDescriptionAuthorityKeyId(object):
    def __init__(self, key_id: str = None):
        self.key_id = key_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_pb2.PrivatecaCertificateCertificateDescriptionAuthorityKeyId()
        if Primitive.to_proto(resource.key_id):
            res.key_id = Primitive.to_proto(resource.key_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionAuthorityKeyId(
            key_id=Primitive.from_proto(resource.key_id),
        )


class CertificateCertificateDescriptionAuthorityKeyIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionAuthorityKeyId.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionAuthorityKeyId.from_proto(i)
            for i in resources
        ]


class CertificateCertificateDescriptionCertFingerprint(object):
    def __init__(self, sha256_hash: str = None):
        self.sha256_hash = sha256_hash

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_pb2.PrivatecaCertificateCertificateDescriptionCertFingerprint()
        )
        if Primitive.to_proto(resource.sha256_hash):
            res.sha256_hash = Primitive.to_proto(resource.sha256_hash)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateCertificateDescriptionCertFingerprint(
            sha256_hash=Primitive.from_proto(resource.sha256_hash),
        )


class CertificateCertificateDescriptionCertFingerprintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateCertificateDescriptionCertFingerprint.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateCertificateDescriptionCertFingerprint.from_proto(i)
            for i in resources
        ]


class CertificateConfigPublicKeyFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_pb2.PrivatecaCertificateConfigPublicKeyFormatEnum.Value(
            "PrivatecaCertificateConfigPublicKeyFormatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_pb2.PrivatecaCertificateConfigPublicKeyFormatEnum.Name(
            resource
        )[len("PrivatecaCertificateConfigPublicKeyFormatEnum") :]


class CertificateSubjectModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_pb2.PrivatecaCertificateSubjectModeEnum.Value(
            "PrivatecaCertificateSubjectModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_pb2.PrivatecaCertificateSubjectModeEnum.Name(resource)[
            len("PrivatecaCertificateSubjectModeEnum") :
        ]


class CertificateRevocationDetailsRevocationStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_pb2.PrivatecaCertificateRevocationDetailsRevocationStateEnum.Value(
            "PrivatecaCertificateRevocationDetailsRevocationStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_pb2.PrivatecaCertificateRevocationDetailsRevocationStateEnum.Name(
            resource
        )[
            len("PrivatecaCertificateRevocationDetailsRevocationStateEnum") :
        ]


class CertificateCertificateDescriptionPublicKeyFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_pb2.PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum.Value(
            "PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_pb2.PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum.Name(
            resource
        )[
            len("PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum") :
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
