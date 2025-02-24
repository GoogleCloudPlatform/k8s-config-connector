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
    certificate_template_pb2,
)
from google3.cloud.graphite.mmv2.services.google.privateca import (
    certificate_template_pb2_grpc,
)

from typing import List


class CertificateTemplate(object):
    def __init__(
        self,
        name: str = None,
        predefined_values: dict = None,
        identity_constraints: dict = None,
        passthrough_extensions: dict = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.predefined_values = predefined_values
        self.identity_constraints = identity_constraints
        self.passthrough_extensions = passthrough_extensions
        self.description = description
        self.labels = labels
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = (
            certificate_template_pb2_grpc.PrivatecaAlphaCertificateTemplateServiceStub(
                channel.Channel()
            )
        )
        request = (
            certificate_template_pb2.ApplyPrivatecaAlphaCertificateTemplateRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if CertificateTemplatePredefinedValues.to_proto(self.predefined_values):
            request.resource.predefined_values.CopyFrom(
                CertificateTemplatePredefinedValues.to_proto(self.predefined_values)
            )
        else:
            request.resource.ClearField("predefined_values")
        if CertificateTemplateIdentityConstraints.to_proto(self.identity_constraints):
            request.resource.identity_constraints.CopyFrom(
                CertificateTemplateIdentityConstraints.to_proto(
                    self.identity_constraints
                )
            )
        else:
            request.resource.ClearField("identity_constraints")
        if CertificateTemplatePassthroughExtensions.to_proto(
            self.passthrough_extensions
        ):
            request.resource.passthrough_extensions.CopyFrom(
                CertificateTemplatePassthroughExtensions.to_proto(
                    self.passthrough_extensions
                )
            )
        else:
            request.resource.ClearField("passthrough_extensions")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyPrivatecaAlphaCertificateTemplate(request)
        self.name = Primitive.from_proto(response.name)
        self.predefined_values = CertificateTemplatePredefinedValues.from_proto(
            response.predefined_values
        )
        self.identity_constraints = CertificateTemplateIdentityConstraints.from_proto(
            response.identity_constraints
        )
        self.passthrough_extensions = (
            CertificateTemplatePassthroughExtensions.from_proto(
                response.passthrough_extensions
            )
        )
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = (
            certificate_template_pb2_grpc.PrivatecaAlphaCertificateTemplateServiceStub(
                channel.Channel()
            )
        )
        request = (
            certificate_template_pb2.DeletePrivatecaAlphaCertificateTemplateRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if CertificateTemplatePredefinedValues.to_proto(self.predefined_values):
            request.resource.predefined_values.CopyFrom(
                CertificateTemplatePredefinedValues.to_proto(self.predefined_values)
            )
        else:
            request.resource.ClearField("predefined_values")
        if CertificateTemplateIdentityConstraints.to_proto(self.identity_constraints):
            request.resource.identity_constraints.CopyFrom(
                CertificateTemplateIdentityConstraints.to_proto(
                    self.identity_constraints
                )
            )
        else:
            request.resource.ClearField("identity_constraints")
        if CertificateTemplatePassthroughExtensions.to_proto(
            self.passthrough_extensions
        ):
            request.resource.passthrough_extensions.CopyFrom(
                CertificateTemplatePassthroughExtensions.to_proto(
                    self.passthrough_extensions
                )
            )
        else:
            request.resource.ClearField("passthrough_extensions")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeletePrivatecaAlphaCertificateTemplate(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = (
            certificate_template_pb2_grpc.PrivatecaAlphaCertificateTemplateServiceStub(
                channel.Channel()
            )
        )
        request = (
            certificate_template_pb2.ListPrivatecaAlphaCertificateTemplateRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListPrivatecaAlphaCertificateTemplate(request).items

    def to_proto(self):
        resource = certificate_template_pb2.PrivatecaAlphaCertificateTemplate()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if CertificateTemplatePredefinedValues.to_proto(self.predefined_values):
            resource.predefined_values.CopyFrom(
                CertificateTemplatePredefinedValues.to_proto(self.predefined_values)
            )
        else:
            resource.ClearField("predefined_values")
        if CertificateTemplateIdentityConstraints.to_proto(self.identity_constraints):
            resource.identity_constraints.CopyFrom(
                CertificateTemplateIdentityConstraints.to_proto(
                    self.identity_constraints
                )
            )
        else:
            resource.ClearField("identity_constraints")
        if CertificateTemplatePassthroughExtensions.to_proto(
            self.passthrough_extensions
        ):
            resource.passthrough_extensions.CopyFrom(
                CertificateTemplatePassthroughExtensions.to_proto(
                    self.passthrough_extensions
                )
            )
        else:
            resource.ClearField("passthrough_extensions")
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class CertificateTemplatePredefinedValues(object):
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
            certificate_template_pb2.PrivatecaAlphaCertificateTemplatePredefinedValues()
        )
        if CertificateTemplatePredefinedValuesKeyUsage.to_proto(resource.key_usage):
            res.key_usage.CopyFrom(
                CertificateTemplatePredefinedValuesKeyUsage.to_proto(resource.key_usage)
            )
        else:
            res.ClearField("key_usage")
        if CertificateTemplatePredefinedValuesCaOptions.to_proto(resource.ca_options):
            res.ca_options.CopyFrom(
                CertificateTemplatePredefinedValuesCaOptions.to_proto(
                    resource.ca_options
                )
            )
        else:
            res.ClearField("ca_options")
        if CertificateTemplatePredefinedValuesPolicyIdsArray.to_proto(
            resource.policy_ids
        ):
            res.policy_ids.extend(
                CertificateTemplatePredefinedValuesPolicyIdsArray.to_proto(
                    resource.policy_ids
                )
            )
        if Primitive.to_proto(resource.aia_ocsp_servers):
            res.aia_ocsp_servers.extend(Primitive.to_proto(resource.aia_ocsp_servers))
        if CertificateTemplatePredefinedValuesAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CertificateTemplatePredefinedValuesAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateTemplatePredefinedValues(
            key_usage=CertificateTemplatePredefinedValuesKeyUsage.from_proto(
                resource.key_usage
            ),
            ca_options=CertificateTemplatePredefinedValuesCaOptions.from_proto(
                resource.ca_options
            ),
            policy_ids=CertificateTemplatePredefinedValuesPolicyIdsArray.from_proto(
                resource.policy_ids
            ),
            aia_ocsp_servers=Primitive.from_proto(resource.aia_ocsp_servers),
            additional_extensions=CertificateTemplatePredefinedValuesAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CertificateTemplatePredefinedValuesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateTemplatePredefinedValues.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateTemplatePredefinedValues.from_proto(i) for i in resources]


class CertificateTemplatePredefinedValuesKeyUsage(object):
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
            certificate_template_pb2.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsage()
        )
        if CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage.to_proto(
            resource.base_key_usage
        ):
            res.base_key_usage.CopyFrom(
                CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage.to_proto(
                    resource.base_key_usage
                )
            )
        else:
            res.ClearField("base_key_usage")
        if CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage.to_proto(
            resource.extended_key_usage
        ):
            res.extended_key_usage.CopyFrom(
                CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage.to_proto(
                    resource.extended_key_usage
                )
            )
        else:
            res.ClearField("extended_key_usage")
        if CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
            resource.unknown_extended_key_usages
        ):
            res.unknown_extended_key_usages.extend(
                CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
                    resource.unknown_extended_key_usages
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateTemplatePredefinedValuesKeyUsage(
            base_key_usage=CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage.from_proto(
                resource.base_key_usage
            ),
            extended_key_usage=CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage.from_proto(
                resource.extended_key_usage
            ),
            unknown_extended_key_usages=CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesArray.from_proto(
                resource.unknown_extended_key_usages
            ),
        )


class CertificateTemplatePredefinedValuesKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateTemplatePredefinedValuesKeyUsage.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateTemplatePredefinedValuesKeyUsage.from_proto(i) for i in resources
        ]


class CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(object):
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
            certificate_template_pb2.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage()
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

        return CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(
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


class CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage.from_proto(i)
            for i in resources
        ]


class CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(object):
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
            certificate_template_pb2.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage()
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

        return CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(
            server_auth=Primitive.from_proto(resource.server_auth),
            client_auth=Primitive.from_proto(resource.client_auth),
            code_signing=Primitive.from_proto(resource.code_signing),
            email_protection=Primitive.from_proto(resource.email_protection),
            time_stamping=Primitive.from_proto(resource.time_stamping),
            ocsp_signing=Primitive.from_proto(resource.ocsp_signing),
        )


class CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage.from_proto(i)
            for i in resources
        ]


class CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_template_pb2.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages.from_proto(
                i
            )
            for i in resources
        ]


class CertificateTemplatePredefinedValuesCaOptions(object):
    def __init__(self, is_ca: bool = None, max_issuer_path_length: int = None):
        self.is_ca = is_ca
        self.max_issuer_path_length = max_issuer_path_length

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_template_pb2.PrivatecaAlphaCertificateTemplatePredefinedValuesCaOptions()
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

        return CertificateTemplatePredefinedValuesCaOptions(
            is_ca=Primitive.from_proto(resource.is_ca),
            max_issuer_path_length=Primitive.from_proto(
                resource.max_issuer_path_length
            ),
        )


class CertificateTemplatePredefinedValuesCaOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateTemplatePredefinedValuesCaOptions.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateTemplatePredefinedValuesCaOptions.from_proto(i)
            for i in resources
        ]


class CertificateTemplatePredefinedValuesPolicyIds(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_template_pb2.PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateTemplatePredefinedValuesPolicyIds(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateTemplatePredefinedValuesPolicyIdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateTemplatePredefinedValuesPolicyIds.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateTemplatePredefinedValuesPolicyIds.from_proto(i)
            for i in resources
        ]


class CertificateTemplatePredefinedValuesAdditionalExtensions(object):
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
            certificate_template_pb2.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions()
        )
        if CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId.to_proto(
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

        return CertificateTemplatePredefinedValuesAdditionalExtensions(
            object_id=CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateTemplatePredefinedValuesAdditionalExtensionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateTemplatePredefinedValuesAdditionalExtensions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateTemplatePredefinedValuesAdditionalExtensions.from_proto(i)
            for i in resources
        ]


class CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_template_pb2.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateTemplateIdentityConstraints(object):
    def __init__(
        self,
        cel_expression: dict = None,
        allow_subject_passthrough: bool = None,
        allow_subject_alt_names_passthrough: bool = None,
    ):
        self.cel_expression = cel_expression
        self.allow_subject_passthrough = allow_subject_passthrough
        self.allow_subject_alt_names_passthrough = allow_subject_alt_names_passthrough

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_template_pb2.PrivatecaAlphaCertificateTemplateIdentityConstraints()
        )
        if CertificateTemplateIdentityConstraintsCelExpression.to_proto(
            resource.cel_expression
        ):
            res.cel_expression.CopyFrom(
                CertificateTemplateIdentityConstraintsCelExpression.to_proto(
                    resource.cel_expression
                )
            )
        else:
            res.ClearField("cel_expression")
        if Primitive.to_proto(resource.allow_subject_passthrough):
            res.allow_subject_passthrough = Primitive.to_proto(
                resource.allow_subject_passthrough
            )
        if Primitive.to_proto(resource.allow_subject_alt_names_passthrough):
            res.allow_subject_alt_names_passthrough = Primitive.to_proto(
                resource.allow_subject_alt_names_passthrough
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateTemplateIdentityConstraints(
            cel_expression=CertificateTemplateIdentityConstraintsCelExpression.from_proto(
                resource.cel_expression
            ),
            allow_subject_passthrough=Primitive.from_proto(
                resource.allow_subject_passthrough
            ),
            allow_subject_alt_names_passthrough=Primitive.from_proto(
                resource.allow_subject_alt_names_passthrough
            ),
        )


class CertificateTemplateIdentityConstraintsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateTemplateIdentityConstraints.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateTemplateIdentityConstraints.from_proto(i) for i in resources]


class CertificateTemplateIdentityConstraintsCelExpression(object):
    def __init__(
        self,
        expression: str = None,
        title: str = None,
        description: str = None,
        location: str = None,
    ):
        self.expression = expression
        self.title = title
        self.description = description
        self.location = location

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_template_pb2.PrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpression()
        )
        if Primitive.to_proto(resource.expression):
            res.expression = Primitive.to_proto(resource.expression)
        if Primitive.to_proto(resource.title):
            res.title = Primitive.to_proto(resource.title)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateTemplateIdentityConstraintsCelExpression(
            expression=Primitive.from_proto(resource.expression),
            title=Primitive.from_proto(resource.title),
            description=Primitive.from_proto(resource.description),
            location=Primitive.from_proto(resource.location),
        )


class CertificateTemplateIdentityConstraintsCelExpressionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateTemplateIdentityConstraintsCelExpression.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateTemplateIdentityConstraintsCelExpression.from_proto(i)
            for i in resources
        ]


class CertificateTemplatePassthroughExtensions(object):
    def __init__(
        self, known_extensions: list = None, additional_extensions: list = None
    ):
        self.known_extensions = known_extensions
        self.additional_extensions = additional_extensions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_template_pb2.PrivatecaAlphaCertificateTemplatePassthroughExtensions()
        )
        if CertificateTemplatePassthroughExtensionsKnownExtensionsEnumArray.to_proto(
            resource.known_extensions
        ):
            res.known_extensions.extend(
                CertificateTemplatePassthroughExtensionsKnownExtensionsEnumArray.to_proto(
                    resource.known_extensions
                )
            )
        if CertificateTemplatePassthroughExtensionsAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CertificateTemplatePassthroughExtensionsAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateTemplatePassthroughExtensions(
            known_extensions=CertificateTemplatePassthroughExtensionsKnownExtensionsEnumArray.from_proto(
                resource.known_extensions
            ),
            additional_extensions=CertificateTemplatePassthroughExtensionsAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CertificateTemplatePassthroughExtensionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateTemplatePassthroughExtensions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateTemplatePassthroughExtensions.from_proto(i) for i in resources
        ]


class CertificateTemplatePassthroughExtensionsAdditionalExtensions(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_template_pb2.PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateTemplatePassthroughExtensionsAdditionalExtensions(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateTemplatePassthroughExtensionsAdditionalExtensionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateTemplatePassthroughExtensionsAdditionalExtensions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateTemplatePassthroughExtensionsAdditionalExtensions.from_proto(i)
            for i in resources
        ]


class CertificateTemplatePassthroughExtensionsKnownExtensionsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_template_pb2.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum.Value(
            "PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_template_pb2.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum.Name(
            resource
        )[
            len(
                "PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum"
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
