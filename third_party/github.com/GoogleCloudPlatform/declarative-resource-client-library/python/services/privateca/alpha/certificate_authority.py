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
from google3.cloud.graphite.mmv2.services.google.privateca import (
    certificate_authority_pb2,
)
from google3.cloud.graphite.mmv2.services.google.privateca import (
    certificate_authority_pb2_grpc,
)

from typing import List


class CertificateAuthority(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        config: dict = None,
        lifetime: str = None,
        key_spec: dict = None,
        subordinate_config: dict = None,
        tier: str = None,
        state: str = None,
        pem_ca_certificates: list = None,
        ca_certificate_descriptions: list = None,
        gcs_bucket: str = None,
        access_urls: dict = None,
        create_time: str = None,
        update_time: str = None,
        delete_time: str = None,
        expire_time: str = None,
        labels: dict = None,
        project: str = None,
        location: str = None,
        ca_pool: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.type = type
        self.config = config
        self.lifetime = lifetime
        self.key_spec = key_spec
        self.gcs_bucket = gcs_bucket
        self.labels = labels
        self.project = project
        self.location = location
        self.ca_pool = ca_pool
        self.service_account_file = service_account_file

    def apply(self):
        stub = certificate_authority_pb2_grpc.PrivatecaAlphaCertificateAuthorityServiceStub(
            channel.Channel()
        )
        request = (
            certificate_authority_pb2.ApplyPrivatecaAlphaCertificateAuthorityRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if CertificateAuthorityTypeEnum.to_proto(self.type):
            request.resource.type = CertificateAuthorityTypeEnum.to_proto(self.type)

        if CertificateAuthorityConfig.to_proto(self.config):
            request.resource.config.CopyFrom(
                CertificateAuthorityConfig.to_proto(self.config)
            )
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.lifetime):
            request.resource.lifetime = Primitive.to_proto(self.lifetime)

        if CertificateAuthorityKeySpec.to_proto(self.key_spec):
            request.resource.key_spec.CopyFrom(
                CertificateAuthorityKeySpec.to_proto(self.key_spec)
            )
        else:
            request.resource.ClearField("key_spec")
        if Primitive.to_proto(self.gcs_bucket):
            request.resource.gcs_bucket = Primitive.to_proto(self.gcs_bucket)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.ca_pool):
            request.resource.ca_pool = Primitive.to_proto(self.ca_pool)

        request.service_account_file = self.service_account_file

        response = stub.ApplyPrivatecaAlphaCertificateAuthority(request)
        self.name = Primitive.from_proto(response.name)
        self.type = CertificateAuthorityTypeEnum.from_proto(response.type)
        self.config = CertificateAuthorityConfig.from_proto(response.config)
        self.lifetime = Primitive.from_proto(response.lifetime)
        self.key_spec = CertificateAuthorityKeySpec.from_proto(response.key_spec)
        self.subordinate_config = CertificateAuthoritySubordinateConfig.from_proto(
            response.subordinate_config
        )
        self.tier = CertificateAuthorityTierEnum.from_proto(response.tier)
        self.state = CertificateAuthorityStateEnum.from_proto(response.state)
        self.pem_ca_certificates = Primitive.from_proto(response.pem_ca_certificates)
        self.ca_certificate_descriptions = (
            CertificateAuthorityCaCertificateDescriptionsArray.from_proto(
                response.ca_certificate_descriptions
            )
        )
        self.gcs_bucket = Primitive.from_proto(response.gcs_bucket)
        self.access_urls = CertificateAuthorityAccessUrls.from_proto(
            response.access_urls
        )
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.expire_time = Primitive.from_proto(response.expire_time)
        self.labels = Primitive.from_proto(response.labels)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.ca_pool = Primitive.from_proto(response.ca_pool)

    def delete(self):
        stub = certificate_authority_pb2_grpc.PrivatecaAlphaCertificateAuthorityServiceStub(
            channel.Channel()
        )
        request = (
            certificate_authority_pb2.DeletePrivatecaAlphaCertificateAuthorityRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if CertificateAuthorityTypeEnum.to_proto(self.type):
            request.resource.type = CertificateAuthorityTypeEnum.to_proto(self.type)

        if CertificateAuthorityConfig.to_proto(self.config):
            request.resource.config.CopyFrom(
                CertificateAuthorityConfig.to_proto(self.config)
            )
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.lifetime):
            request.resource.lifetime = Primitive.to_proto(self.lifetime)

        if CertificateAuthorityKeySpec.to_proto(self.key_spec):
            request.resource.key_spec.CopyFrom(
                CertificateAuthorityKeySpec.to_proto(self.key_spec)
            )
        else:
            request.resource.ClearField("key_spec")
        if Primitive.to_proto(self.gcs_bucket):
            request.resource.gcs_bucket = Primitive.to_proto(self.gcs_bucket)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.ca_pool):
            request.resource.ca_pool = Primitive.to_proto(self.ca_pool)

        response = stub.DeletePrivatecaAlphaCertificateAuthority(request)

    @classmethod
    def list(self, project, location, caPool, service_account_file=""):
        stub = certificate_authority_pb2_grpc.PrivatecaAlphaCertificateAuthorityServiceStub(
            channel.Channel()
        )
        request = (
            certificate_authority_pb2.ListPrivatecaAlphaCertificateAuthorityRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.CaPool = caPool

        return stub.ListPrivatecaAlphaCertificateAuthority(request).items

    def to_proto(self):
        resource = certificate_authority_pb2.PrivatecaAlphaCertificateAuthority()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if CertificateAuthorityTypeEnum.to_proto(self.type):
            resource.type = CertificateAuthorityTypeEnum.to_proto(self.type)
        if CertificateAuthorityConfig.to_proto(self.config):
            resource.config.CopyFrom(CertificateAuthorityConfig.to_proto(self.config))
        else:
            resource.ClearField("config")
        if Primitive.to_proto(self.lifetime):
            resource.lifetime = Primitive.to_proto(self.lifetime)
        if CertificateAuthorityKeySpec.to_proto(self.key_spec):
            resource.key_spec.CopyFrom(
                CertificateAuthorityKeySpec.to_proto(self.key_spec)
            )
        else:
            resource.ClearField("key_spec")
        if Primitive.to_proto(self.gcs_bucket):
            resource.gcs_bucket = Primitive.to_proto(self.gcs_bucket)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.ca_pool):
            resource.ca_pool = Primitive.to_proto(self.ca_pool)
        return resource


class CertificateAuthorityConfig(object):
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

        res = certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfig()
        if CertificateAuthorityConfigSubjectConfig.to_proto(resource.subject_config):
            res.subject_config.CopyFrom(
                CertificateAuthorityConfigSubjectConfig.to_proto(
                    resource.subject_config
                )
            )
        else:
            res.ClearField("subject_config")
        if CertificateAuthorityConfigX509Config.to_proto(resource.x509_config):
            res.x509_config.CopyFrom(
                CertificateAuthorityConfigX509Config.to_proto(resource.x509_config)
            )
        else:
            res.ClearField("x509_config")
        if CertificateAuthorityConfigPublicKey.to_proto(resource.public_key):
            res.public_key.CopyFrom(
                CertificateAuthorityConfigPublicKey.to_proto(resource.public_key)
            )
        else:
            res.ClearField("public_key")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfig(
            subject_config=CertificateAuthorityConfigSubjectConfig.from_proto(
                resource.subject_config
            ),
            x509_config=CertificateAuthorityConfigX509Config.from_proto(
                resource.x509_config
            ),
            public_key=CertificateAuthorityConfigPublicKey.from_proto(
                resource.public_key
            ),
        )


class CertificateAuthorityConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthorityConfig.from_proto(i) for i in resources]


class CertificateAuthorityConfigSubjectConfig(object):
    def __init__(self, subject: dict = None, subject_alt_name: dict = None):
        self.subject = subject
        self.subject_alt_name = subject_alt_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigSubjectConfig()
        )
        if CertificateAuthorityConfigSubjectConfigSubject.to_proto(resource.subject):
            res.subject.CopyFrom(
                CertificateAuthorityConfigSubjectConfigSubject.to_proto(
                    resource.subject
                )
            )
        else:
            res.ClearField("subject")
        if CertificateAuthorityConfigSubjectConfigSubjectAltName.to_proto(
            resource.subject_alt_name
        ):
            res.subject_alt_name.CopyFrom(
                CertificateAuthorityConfigSubjectConfigSubjectAltName.to_proto(
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

        return CertificateAuthorityConfigSubjectConfig(
            subject=CertificateAuthorityConfigSubjectConfigSubject.from_proto(
                resource.subject
            ),
            subject_alt_name=CertificateAuthorityConfigSubjectConfigSubjectAltName.from_proto(
                resource.subject_alt_name
            ),
        )


class CertificateAuthorityConfigSubjectConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityConfigSubjectConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigSubjectConfig.from_proto(i) for i in resources
        ]


class CertificateAuthorityConfigSubjectConfigSubject(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubject()
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

        return CertificateAuthorityConfigSubjectConfigSubject(
            common_name=Primitive.from_proto(resource.common_name),
            country_code=Primitive.from_proto(resource.country_code),
            organization=Primitive.from_proto(resource.organization),
            organizational_unit=Primitive.from_proto(resource.organizational_unit),
            locality=Primitive.from_proto(resource.locality),
            province=Primitive.from_proto(resource.province),
            street_address=Primitive.from_proto(resource.street_address),
            postal_code=Primitive.from_proto(resource.postal_code),
        )


class CertificateAuthorityConfigSubjectConfigSubjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigSubjectConfigSubject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigSubjectConfigSubject.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityConfigSubjectConfigSubjectAltName(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltName()
        )
        if Primitive.to_proto(resource.dns_names):
            res.dns_names.extend(Primitive.to_proto(resource.dns_names))
        if Primitive.to_proto(resource.uris):
            res.uris.extend(Primitive.to_proto(resource.uris))
        if Primitive.to_proto(resource.email_addresses):
            res.email_addresses.extend(Primitive.to_proto(resource.email_addresses))
        if Primitive.to_proto(resource.ip_addresses):
            res.ip_addresses.extend(Primitive.to_proto(resource.ip_addresses))
        if CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansArray.to_proto(
            resource.custom_sans
        ):
            res.custom_sans.extend(
                CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansArray.to_proto(
                    resource.custom_sans
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigSubjectConfigSubjectAltName(
            dns_names=Primitive.from_proto(resource.dns_names),
            uris=Primitive.from_proto(resource.uris),
            email_addresses=Primitive.from_proto(resource.email_addresses),
            ip_addresses=Primitive.from_proto(resource.ip_addresses),
            custom_sans=CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansArray.from_proto(
                resource.custom_sans
            ),
        )


class CertificateAuthorityConfigSubjectConfigSubjectAltNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltName.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans()
        )
        if CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId.to_proto(
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

        return CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(
            object_id=CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigX509Config(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigX509Config()
        )
        if CertificateAuthorityConfigX509ConfigKeyUsage.to_proto(resource.key_usage):
            res.key_usage.CopyFrom(
                CertificateAuthorityConfigX509ConfigKeyUsage.to_proto(
                    resource.key_usage
                )
            )
        else:
            res.ClearField("key_usage")
        if CertificateAuthorityConfigX509ConfigCaOptions.to_proto(resource.ca_options):
            res.ca_options.CopyFrom(
                CertificateAuthorityConfigX509ConfigCaOptions.to_proto(
                    resource.ca_options
                )
            )
        else:
            res.ClearField("ca_options")
        if CertificateAuthorityConfigX509ConfigPolicyIdsArray.to_proto(
            resource.policy_ids
        ):
            res.policy_ids.extend(
                CertificateAuthorityConfigX509ConfigPolicyIdsArray.to_proto(
                    resource.policy_ids
                )
            )
        if Primitive.to_proto(resource.aia_ocsp_servers):
            res.aia_ocsp_servers.extend(Primitive.to_proto(resource.aia_ocsp_servers))
        if CertificateAuthorityConfigX509ConfigAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CertificateAuthorityConfigX509ConfigAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigX509Config(
            key_usage=CertificateAuthorityConfigX509ConfigKeyUsage.from_proto(
                resource.key_usage
            ),
            ca_options=CertificateAuthorityConfigX509ConfigCaOptions.from_proto(
                resource.ca_options
            ),
            policy_ids=CertificateAuthorityConfigX509ConfigPolicyIdsArray.from_proto(
                resource.policy_ids
            ),
            aia_ocsp_servers=Primitive.from_proto(resource.aia_ocsp_servers),
            additional_extensions=CertificateAuthorityConfigX509ConfigAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CertificateAuthorityConfigX509ConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityConfigX509Config.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthorityConfigX509Config.from_proto(i) for i in resources]


class CertificateAuthorityConfigX509ConfigKeyUsage(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsage()
        )
        if CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage.to_proto(
            resource.base_key_usage
        ):
            res.base_key_usage.CopyFrom(
                CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage.to_proto(
                    resource.base_key_usage
                )
            )
        else:
            res.ClearField("base_key_usage")
        if CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage.to_proto(
            resource.extended_key_usage
        ):
            res.extended_key_usage.CopyFrom(
                CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage.to_proto(
                    resource.extended_key_usage
                )
            )
        else:
            res.ClearField("extended_key_usage")
        if CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
            resource.unknown_extended_key_usages
        ):
            res.unknown_extended_key_usages.extend(
                CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
                    resource.unknown_extended_key_usages
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigX509ConfigKeyUsage(
            base_key_usage=CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage.from_proto(
                resource.base_key_usage
            ),
            extended_key_usage=CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage.from_proto(
                resource.extended_key_usage
            ),
            unknown_extended_key_usages=CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesArray.from_proto(
                resource.unknown_extended_key_usages
            ),
        )


class CertificateAuthorityConfigX509ConfigKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigX509ConfigKeyUsage.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigX509ConfigKeyUsage.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage()
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

        return CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(
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


class CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage()
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

        return CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(
            server_auth=Primitive.from_proto(resource.server_auth),
            client_auth=Primitive.from_proto(resource.client_auth),
            code_signing=Primitive.from_proto(resource.code_signing),
            email_protection=Primitive.from_proto(resource.email_protection),
            time_stamping=Primitive.from_proto(resource.time_stamping),
            ocsp_signing=Primitive.from_proto(resource.ocsp_signing),
        )


class CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigX509ConfigCaOptions(object):
    def __init__(
        self,
        is_ca: bool = None,
        max_issuer_path_length: int = None,
        zero_max_issuer_path_length: bool = None,
    ):
        self.is_ca = is_ca
        self.max_issuer_path_length = max_issuer_path_length
        self.zero_max_issuer_path_length = zero_max_issuer_path_length

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigX509ConfigCaOptions()
        )
        if Primitive.to_proto(resource.is_ca):
            res.is_ca = Primitive.to_proto(resource.is_ca)
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

        return CertificateAuthorityConfigX509ConfigCaOptions(
            is_ca=Primitive.from_proto(resource.is_ca),
            max_issuer_path_length=Primitive.from_proto(
                resource.max_issuer_path_length
            ),
            zero_max_issuer_path_length=Primitive.from_proto(
                resource.zero_max_issuer_path_length
            ),
        )


class CertificateAuthorityConfigX509ConfigCaOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigX509ConfigCaOptions.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigX509ConfigCaOptions.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityConfigX509ConfigPolicyIds(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigX509ConfigPolicyIds()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigX509ConfigPolicyIds(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityConfigX509ConfigPolicyIdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigX509ConfigPolicyIds.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigX509ConfigPolicyIds.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityConfigX509ConfigAdditionalExtensions(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensions()
        )
        if CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId.to_proto(
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

        return CertificateAuthorityConfigX509ConfigAdditionalExtensions(
            object_id=CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateAuthorityConfigX509ConfigAdditionalExtensionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigX509ConfigAdditionalExtensions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigX509ConfigAdditionalExtensions.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigPublicKey(object):
    def __init__(self, key: str = None, format: str = None):
        self.key = key
        self.format = format

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigPublicKey()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if CertificateAuthorityConfigPublicKeyFormatEnum.to_proto(resource.format):
            res.format = CertificateAuthorityConfigPublicKeyFormatEnum.to_proto(
                resource.format
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigPublicKey(
            key=Primitive.from_proto(resource.key),
            format=CertificateAuthorityConfigPublicKeyFormatEnum.from_proto(
                resource.format
            ),
        )


class CertificateAuthorityConfigPublicKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityConfigPublicKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthorityConfigPublicKey.from_proto(i) for i in resources]


class CertificateAuthorityKeySpec(object):
    def __init__(self, cloud_kms_key_version: str = None, algorithm: str = None):
        self.cloud_kms_key_version = cloud_kms_key_version
        self.algorithm = algorithm

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityKeySpec()
        if Primitive.to_proto(resource.cloud_kms_key_version):
            res.cloud_kms_key_version = Primitive.to_proto(
                resource.cloud_kms_key_version
            )
        if CertificateAuthorityKeySpecAlgorithmEnum.to_proto(resource.algorithm):
            res.algorithm = CertificateAuthorityKeySpecAlgorithmEnum.to_proto(
                resource.algorithm
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityKeySpec(
            cloud_kms_key_version=Primitive.from_proto(resource.cloud_kms_key_version),
            algorithm=CertificateAuthorityKeySpecAlgorithmEnum.from_proto(
                resource.algorithm
            ),
        )


class CertificateAuthorityKeySpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityKeySpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthorityKeySpec.from_proto(i) for i in resources]


class CertificateAuthoritySubordinateConfig(object):
    def __init__(
        self, certificate_authority: str = None, pem_issuer_chain: dict = None
    ):
        self.certificate_authority = certificate_authority
        self.pem_issuer_chain = pem_issuer_chain

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthoritySubordinateConfig()
        )
        if Primitive.to_proto(resource.certificate_authority):
            res.certificate_authority = Primitive.to_proto(
                resource.certificate_authority
            )
        if CertificateAuthoritySubordinateConfigPemIssuerChain.to_proto(
            resource.pem_issuer_chain
        ):
            res.pem_issuer_chain.CopyFrom(
                CertificateAuthoritySubordinateConfigPemIssuerChain.to_proto(
                    resource.pem_issuer_chain
                )
            )
        else:
            res.ClearField("pem_issuer_chain")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthoritySubordinateConfig(
            certificate_authority=Primitive.from_proto(resource.certificate_authority),
            pem_issuer_chain=CertificateAuthoritySubordinateConfigPemIssuerChain.from_proto(
                resource.pem_issuer_chain
            ),
        )


class CertificateAuthoritySubordinateConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthoritySubordinateConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthoritySubordinateConfig.from_proto(i) for i in resources]


class CertificateAuthoritySubordinateConfigPemIssuerChain(object):
    def __init__(self, pem_certificates: list = None):
        self.pem_certificates = pem_certificates

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthoritySubordinateConfigPemIssuerChain()
        )
        if Primitive.to_proto(resource.pem_certificates):
            res.pem_certificates.extend(Primitive.to_proto(resource.pem_certificates))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthoritySubordinateConfigPemIssuerChain(
            pem_certificates=Primitive.from_proto(resource.pem_certificates),
        )


class CertificateAuthoritySubordinateConfigPemIssuerChainArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthoritySubordinateConfigPemIssuerChain.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthoritySubordinateConfigPemIssuerChain.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptions(object):
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

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptions()
        )
        if CertificateAuthorityCaCertificateDescriptionsSubjectDescription.to_proto(
            resource.subject_description
        ):
            res.subject_description.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsSubjectDescription.to_proto(
                    resource.subject_description
                )
            )
        else:
            res.ClearField("subject_description")
        if CertificateAuthorityCaCertificateDescriptionsX509Description.to_proto(
            resource.x509_description
        ):
            res.x509_description.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsX509Description.to_proto(
                    resource.x509_description
                )
            )
        else:
            res.ClearField("x509_description")
        if CertificateAuthorityCaCertificateDescriptionsPublicKey.to_proto(
            resource.public_key
        ):
            res.public_key.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsPublicKey.to_proto(
                    resource.public_key
                )
            )
        else:
            res.ClearField("public_key")
        if CertificateAuthorityCaCertificateDescriptionsSubjectKeyId.to_proto(
            resource.subject_key_id
        ):
            res.subject_key_id.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsSubjectKeyId.to_proto(
                    resource.subject_key_id
                )
            )
        else:
            res.ClearField("subject_key_id")
        if CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId.to_proto(
            resource.authority_key_id
        ):
            res.authority_key_id.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId.to_proto(
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
        if CertificateAuthorityCaCertificateDescriptionsCertFingerprint.to_proto(
            resource.cert_fingerprint
        ):
            res.cert_fingerprint.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsCertFingerprint.to_proto(
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

        return CertificateAuthorityCaCertificateDescriptions(
            subject_description=CertificateAuthorityCaCertificateDescriptionsSubjectDescription.from_proto(
                resource.subject_description
            ),
            x509_description=CertificateAuthorityCaCertificateDescriptionsX509Description.from_proto(
                resource.x509_description
            ),
            public_key=CertificateAuthorityCaCertificateDescriptionsPublicKey.from_proto(
                resource.public_key
            ),
            subject_key_id=CertificateAuthorityCaCertificateDescriptionsSubjectKeyId.from_proto(
                resource.subject_key_id
            ),
            authority_key_id=CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId.from_proto(
                resource.authority_key_id
            ),
            crl_distribution_points=Primitive.from_proto(
                resource.crl_distribution_points
            ),
            aia_issuing_certificate_urls=Primitive.from_proto(
                resource.aia_issuing_certificate_urls
            ),
            cert_fingerprint=CertificateAuthorityCaCertificateDescriptionsCertFingerprint.from_proto(
                resource.cert_fingerprint
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptions.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptions.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectDescription(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescription()
        )
        if CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject.to_proto(
            resource.subject
        ):
            res.subject.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject.to_proto(
                    resource.subject
                )
            )
        else:
            res.ClearField("subject")
        if CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName.to_proto(
            resource.subject_alt_name
        ):
            res.subject_alt_name.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName.to_proto(
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

        return CertificateAuthorityCaCertificateDescriptionsSubjectDescription(
            subject=CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject.from_proto(
                resource.subject
            ),
            subject_alt_name=CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName.from_proto(
                resource.subject_alt_name
            ),
            hex_serial_number=Primitive.from_proto(resource.hex_serial_number),
            lifetime=Primitive.from_proto(resource.lifetime),
            not_before_time=Primitive.from_proto(resource.not_before_time),
            not_after_time=Primitive.from_proto(resource.not_after_time),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescription.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescription.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject()
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

        return CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(
            common_name=Primitive.from_proto(resource.common_name),
            country_code=Primitive.from_proto(resource.country_code),
            organization=Primitive.from_proto(resource.organization),
            organizational_unit=Primitive.from_proto(resource.organizational_unit),
            locality=Primitive.from_proto(resource.locality),
            province=Primitive.from_proto(resource.province),
            street_address=Primitive.from_proto(resource.street_address),
            postal_code=Primitive.from_proto(resource.postal_code),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(
    object
):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName()
        )
        if Primitive.to_proto(resource.dns_names):
            res.dns_names.extend(Primitive.to_proto(resource.dns_names))
        if Primitive.to_proto(resource.uris):
            res.uris.extend(Primitive.to_proto(resource.uris))
        if Primitive.to_proto(resource.email_addresses):
            res.email_addresses.extend(Primitive.to_proto(resource.email_addresses))
        if Primitive.to_proto(resource.ip_addresses):
            res.ip_addresses.extend(Primitive.to_proto(resource.ip_addresses))
        if CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansArray.to_proto(
            resource.custom_sans
        ):
            res.custom_sans.extend(
                CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansArray.to_proto(
                    resource.custom_sans
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(
            dns_names=Primitive.from_proto(resource.dns_names),
            uris=Primitive.from_proto(resource.uris),
            email_addresses=Primitive.from_proto(resource.email_addresses),
            ip_addresses=Primitive.from_proto(resource.ip_addresses),
            custom_sans=CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansArray.from_proto(
                resource.custom_sans
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans()
        )
        if CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId.to_proto(
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

        return CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(
            object_id=CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsX509Description(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509Description()
        )
        if CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage.to_proto(
            resource.key_usage
        ):
            res.key_usage.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage.to_proto(
                    resource.key_usage
                )
            )
        else:
            res.ClearField("key_usage")
        if CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions.to_proto(
            resource.ca_options
        ):
            res.ca_options.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions.to_proto(
                    resource.ca_options
                )
            )
        else:
            res.ClearField("ca_options")
        if CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsArray.to_proto(
            resource.policy_ids
        ):
            res.policy_ids.extend(
                CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsArray.to_proto(
                    resource.policy_ids
                )
            )
        if Primitive.to_proto(resource.aia_ocsp_servers):
            res.aia_ocsp_servers.extend(Primitive.to_proto(resource.aia_ocsp_servers))
        if CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsX509Description(
            key_usage=CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage.from_proto(
                resource.key_usage
            ),
            ca_options=CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions.from_proto(
                resource.ca_options
            ),
            policy_ids=CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsArray.from_proto(
                resource.policy_ids
            ),
            aia_ocsp_servers=Primitive.from_proto(resource.aia_ocsp_servers),
            additional_extensions=CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsX509Description.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsX509Description.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage(object):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage()
        )
        if CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage.to_proto(
            resource.base_key_usage
        ):
            res.base_key_usage.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage.to_proto(
                    resource.base_key_usage
                )
            )
        else:
            res.ClearField("base_key_usage")
        if CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage.to_proto(
            resource.extended_key_usage
        ):
            res.extended_key_usage.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage.to_proto(
                    resource.extended_key_usage
                )
            )
        else:
            res.ClearField("extended_key_usage")
        if CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
            resource.unknown_extended_key_usages
        ):
            res.unknown_extended_key_usages.extend(
                CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
                    resource.unknown_extended_key_usages
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage(
            base_key_usage=CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage.from_proto(
                resource.base_key_usage
            ),
            extended_key_usage=CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage.from_proto(
                resource.extended_key_usage
            ),
            unknown_extended_key_usages=CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesArray.from_proto(
                resource.unknown_extended_key_usages
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage(
    object
):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage()
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

        return CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage(
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


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage(
    object
):
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage()
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

        return CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage(
            server_auth=Primitive.from_proto(resource.server_auth),
            client_auth=Primitive.from_proto(resource.client_auth),
            code_signing=Primitive.from_proto(resource.code_signing),
            email_protection=Primitive.from_proto(resource.email_protection),
            time_stamping=Primitive.from_proto(resource.time_stamping),
            ocsp_signing=Primitive.from_proto(resource.ocsp_signing),
        )


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions(object):
    def __init__(self, is_ca: bool = None, max_issuer_path_length: int = None):
        self.is_ca = is_ca
        self.max_issuer_path_length = max_issuer_path_length

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions()
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

        return CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions(
            is_ca=Primitive.from_proto(resource.is_ca),
            max_issuer_path_length=Primitive.from_proto(
                resource.max_issuer_path_length
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions(
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
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions()
        )
        if CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId.to_proto(
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

        return CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions(
            object_id=CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsPublicKey(object):
    def __init__(self, key: str = None, format: str = None):
        self.key = key
        self.format = format

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKey()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum.to_proto(
            resource.format
        ):
            res.format = CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum.to_proto(
                resource.format
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsPublicKey(
            key=Primitive.from_proto(resource.key),
            format=CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum.from_proto(
                resource.format
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsPublicKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsPublicKey.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsPublicKey.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectKeyId(object):
    def __init__(self, key_id: str = None):
        self.key_id = key_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId()
        )
        if Primitive.to_proto(resource.key_id):
            res.key_id = Primitive.to_proto(resource.key_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsSubjectKeyId(
            key_id=Primitive.from_proto(resource.key_id),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectKeyIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectKeyId.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectKeyId.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(object):
    def __init__(self, key_id: str = None):
        self.key_id = key_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId()
        )
        if Primitive.to_proto(resource.key_id):
            res.key_id = Primitive.to_proto(resource.key_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(
            key_id=Primitive.from_proto(resource.key_id),
        )


class CertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsCertFingerprint(object):
    def __init__(self, sha256_hash: str = None):
        self.sha256_hash = sha256_hash

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsCertFingerprint()
        )
        if Primitive.to_proto(resource.sha256_hash):
            res.sha256_hash = Primitive.to_proto(resource.sha256_hash)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsCertFingerprint(
            sha256_hash=Primitive.from_proto(resource.sha256_hash),
        )


class CertificateAuthorityCaCertificateDescriptionsCertFingerprintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsCertFingerprint.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsCertFingerprint.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityAccessUrls(object):
    def __init__(
        self, ca_certificate_access_url: str = None, crl_access_urls: list = None
    ):
        self.ca_certificate_access_url = ca_certificate_access_url
        self.crl_access_urls = crl_access_urls

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityAccessUrls()
        if Primitive.to_proto(resource.ca_certificate_access_url):
            res.ca_certificate_access_url = Primitive.to_proto(
                resource.ca_certificate_access_url
            )
        if Primitive.to_proto(resource.crl_access_urls):
            res.crl_access_urls.extend(Primitive.to_proto(resource.crl_access_urls))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityAccessUrls(
            ca_certificate_access_url=Primitive.from_proto(
                resource.ca_certificate_access_url
            ),
            crl_access_urls=Primitive.from_proto(resource.crl_access_urls),
        )


class CertificateAuthorityAccessUrlsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityAccessUrls.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthorityAccessUrls.from_proto(i) for i in resources]


class CertificateAuthorityTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityTypeEnum.Value(
                "PrivatecaAlphaCertificateAuthorityTypeEnum%s" % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityTypeEnum.Name(
                resource
            )[len("PrivatecaAlphaCertificateAuthorityTypeEnum") :]
        )


class CertificateAuthorityConfigPublicKeyFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum.Value(
            "PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum.Name(
            resource
        )[
            len("PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum") :
        ]


class CertificateAuthorityKeySpecAlgorithmEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum.Value(
            "PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum.Name(
            resource
        )[
            len("PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum") :
        ]


class CertificateAuthorityTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityTierEnum.Value(
                "PrivatecaAlphaCertificateAuthorityTierEnum%s" % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityTierEnum.Name(
                resource
            )[len("PrivatecaAlphaCertificateAuthorityTierEnum") :]
        )


class CertificateAuthorityStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityStateEnum.Value(
                "PrivatecaAlphaCertificateAuthorityStateEnum%s" % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityStateEnum.Name(
                resource
            )[len("PrivatecaAlphaCertificateAuthorityStateEnum") :]
        )


class CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum.Value(
            "PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum.Name(
            resource
        )[
            len(
                "PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum"
            ) :
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
