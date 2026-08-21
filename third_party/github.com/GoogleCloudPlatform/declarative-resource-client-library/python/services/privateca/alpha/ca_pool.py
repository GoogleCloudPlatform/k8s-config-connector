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
from google3.cloud.graphite.mmv2.services.google.privateca import ca_pool_pb2
from google3.cloud.graphite.mmv2.services.google.privateca import ca_pool_pb2_grpc

from typing import List


class CaPool(object):
    def __init__(
        self,
        name: str = None,
        tier: str = None,
        issuance_policy: dict = None,
        publishing_options: dict = None,
        labels: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.tier = tier
        self.issuance_policy = issuance_policy
        self.publishing_options = publishing_options
        self.labels = labels
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = ca_pool_pb2_grpc.PrivatecaAlphaCaPoolServiceStub(channel.Channel())
        request = ca_pool_pb2.ApplyPrivatecaAlphaCaPoolRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if CaPoolTierEnum.to_proto(self.tier):
            request.resource.tier = CaPoolTierEnum.to_proto(self.tier)

        if CaPoolIssuancePolicy.to_proto(self.issuance_policy):
            request.resource.issuance_policy.CopyFrom(
                CaPoolIssuancePolicy.to_proto(self.issuance_policy)
            )
        else:
            request.resource.ClearField("issuance_policy")
        if CaPoolPublishingOptions.to_proto(self.publishing_options):
            request.resource.publishing_options.CopyFrom(
                CaPoolPublishingOptions.to_proto(self.publishing_options)
            )
        else:
            request.resource.ClearField("publishing_options")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyPrivatecaAlphaCaPool(request)
        self.name = Primitive.from_proto(response.name)
        self.tier = CaPoolTierEnum.from_proto(response.tier)
        self.issuance_policy = CaPoolIssuancePolicy.from_proto(response.issuance_policy)
        self.publishing_options = CaPoolPublishingOptions.from_proto(
            response.publishing_options
        )
        self.labels = Primitive.from_proto(response.labels)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = ca_pool_pb2_grpc.PrivatecaAlphaCaPoolServiceStub(channel.Channel())
        request = ca_pool_pb2.DeletePrivatecaAlphaCaPoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if CaPoolTierEnum.to_proto(self.tier):
            request.resource.tier = CaPoolTierEnum.to_proto(self.tier)

        if CaPoolIssuancePolicy.to_proto(self.issuance_policy):
            request.resource.issuance_policy.CopyFrom(
                CaPoolIssuancePolicy.to_proto(self.issuance_policy)
            )
        else:
            request.resource.ClearField("issuance_policy")
        if CaPoolPublishingOptions.to_proto(self.publishing_options):
            request.resource.publishing_options.CopyFrom(
                CaPoolPublishingOptions.to_proto(self.publishing_options)
            )
        else:
            request.resource.ClearField("publishing_options")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeletePrivatecaAlphaCaPool(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = ca_pool_pb2_grpc.PrivatecaAlphaCaPoolServiceStub(channel.Channel())
        request = ca_pool_pb2.ListPrivatecaAlphaCaPoolRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListPrivatecaAlphaCaPool(request).items

    def to_proto(self):
        resource = ca_pool_pb2.PrivatecaAlphaCaPool()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if CaPoolTierEnum.to_proto(self.tier):
            resource.tier = CaPoolTierEnum.to_proto(self.tier)
        if CaPoolIssuancePolicy.to_proto(self.issuance_policy):
            resource.issuance_policy.CopyFrom(
                CaPoolIssuancePolicy.to_proto(self.issuance_policy)
            )
        else:
            resource.ClearField("issuance_policy")
        if CaPoolPublishingOptions.to_proto(self.publishing_options):
            resource.publishing_options.CopyFrom(
                CaPoolPublishingOptions.to_proto(self.publishing_options)
            )
        else:
            resource.ClearField("publishing_options")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class CaPoolIssuancePolicy(object):
    def __init__(
        self,
        allowed_key_types: list = None,
        maximum_lifetime: str = None,
        allowed_issuance_modes: dict = None,
        baseline_values: dict = None,
        identity_constraints: dict = None,
        passthrough_extensions: dict = None,
    ):
        self.allowed_key_types = allowed_key_types
        self.maximum_lifetime = maximum_lifetime
        self.allowed_issuance_modes = allowed_issuance_modes
        self.baseline_values = baseline_values
        self.identity_constraints = identity_constraints
        self.passthrough_extensions = passthrough_extensions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicy()
        if CaPoolIssuancePolicyAllowedKeyTypesArray.to_proto(
            resource.allowed_key_types
        ):
            res.allowed_key_types.extend(
                CaPoolIssuancePolicyAllowedKeyTypesArray.to_proto(
                    resource.allowed_key_types
                )
            )
        if Primitive.to_proto(resource.maximum_lifetime):
            res.maximum_lifetime = Primitive.to_proto(resource.maximum_lifetime)
        if CaPoolIssuancePolicyAllowedIssuanceModes.to_proto(
            resource.allowed_issuance_modes
        ):
            res.allowed_issuance_modes.CopyFrom(
                CaPoolIssuancePolicyAllowedIssuanceModes.to_proto(
                    resource.allowed_issuance_modes
                )
            )
        else:
            res.ClearField("allowed_issuance_modes")
        if CaPoolIssuancePolicyBaselineValues.to_proto(resource.baseline_values):
            res.baseline_values.CopyFrom(
                CaPoolIssuancePolicyBaselineValues.to_proto(resource.baseline_values)
            )
        else:
            res.ClearField("baseline_values")
        if CaPoolIssuancePolicyIdentityConstraints.to_proto(
            resource.identity_constraints
        ):
            res.identity_constraints.CopyFrom(
                CaPoolIssuancePolicyIdentityConstraints.to_proto(
                    resource.identity_constraints
                )
            )
        else:
            res.ClearField("identity_constraints")
        if CaPoolIssuancePolicyPassthroughExtensions.to_proto(
            resource.passthrough_extensions
        ):
            res.passthrough_extensions.CopyFrom(
                CaPoolIssuancePolicyPassthroughExtensions.to_proto(
                    resource.passthrough_extensions
                )
            )
        else:
            res.ClearField("passthrough_extensions")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicy(
            allowed_key_types=CaPoolIssuancePolicyAllowedKeyTypesArray.from_proto(
                resource.allowed_key_types
            ),
            maximum_lifetime=Primitive.from_proto(resource.maximum_lifetime),
            allowed_issuance_modes=CaPoolIssuancePolicyAllowedIssuanceModes.from_proto(
                resource.allowed_issuance_modes
            ),
            baseline_values=CaPoolIssuancePolicyBaselineValues.from_proto(
                resource.baseline_values
            ),
            identity_constraints=CaPoolIssuancePolicyIdentityConstraints.from_proto(
                resource.identity_constraints
            ),
            passthrough_extensions=CaPoolIssuancePolicyPassthroughExtensions.from_proto(
                resource.passthrough_extensions
            ),
        )


class CaPoolIssuancePolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CaPoolIssuancePolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CaPoolIssuancePolicy.from_proto(i) for i in resources]


class CaPoolIssuancePolicyAllowedKeyTypes(object):
    def __init__(self, rsa: dict = None, elliptic_curve: dict = None):
        self.rsa = rsa
        self.elliptic_curve = elliptic_curve

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypes()
        if CaPoolIssuancePolicyAllowedKeyTypesRsa.to_proto(resource.rsa):
            res.rsa.CopyFrom(
                CaPoolIssuancePolicyAllowedKeyTypesRsa.to_proto(resource.rsa)
            )
        else:
            res.ClearField("rsa")
        if CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve.to_proto(
            resource.elliptic_curve
        ):
            res.elliptic_curve.CopyFrom(
                CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve.to_proto(
                    resource.elliptic_curve
                )
            )
        else:
            res.ClearField("elliptic_curve")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicyAllowedKeyTypes(
            rsa=CaPoolIssuancePolicyAllowedKeyTypesRsa.from_proto(resource.rsa),
            elliptic_curve=CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve.from_proto(
                resource.elliptic_curve
            ),
        )


class CaPoolIssuancePolicyAllowedKeyTypesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CaPoolIssuancePolicyAllowedKeyTypes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CaPoolIssuancePolicyAllowedKeyTypes.from_proto(i) for i in resources]


class CaPoolIssuancePolicyAllowedKeyTypesRsa(object):
    def __init__(self, min_modulus_size: int = None, max_modulus_size: int = None):
        self.min_modulus_size = min_modulus_size
        self.max_modulus_size = max_modulus_size

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesRsa()
        if Primitive.to_proto(resource.min_modulus_size):
            res.min_modulus_size = Primitive.to_proto(resource.min_modulus_size)
        if Primitive.to_proto(resource.max_modulus_size):
            res.max_modulus_size = Primitive.to_proto(resource.max_modulus_size)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicyAllowedKeyTypesRsa(
            min_modulus_size=Primitive.from_proto(resource.min_modulus_size),
            max_modulus_size=Primitive.from_proto(resource.max_modulus_size),
        )


class CaPoolIssuancePolicyAllowedKeyTypesRsaArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CaPoolIssuancePolicyAllowedKeyTypesRsa.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CaPoolIssuancePolicyAllowedKeyTypesRsa.from_proto(i) for i in resources]


class CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve(object):
    def __init__(self, signature_algorithm: str = None):
        self.signature_algorithm = signature_algorithm

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve()
        )
        if CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum.to_proto(
            resource.signature_algorithm
        ):
            res.signature_algorithm = CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum.to_proto(
                resource.signature_algorithm
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve(
            signature_algorithm=CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum.from_proto(
                resource.signature_algorithm
            ),
        )


class CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve.from_proto(i)
            for i in resources
        ]


class CaPoolIssuancePolicyAllowedIssuanceModes(object):
    def __init__(
        self,
        allow_csr_based_issuance: bool = None,
        allow_config_based_issuance: bool = None,
    ):
        self.allow_csr_based_issuance = allow_csr_based_issuance
        self.allow_config_based_issuance = allow_config_based_issuance

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyAllowedIssuanceModes()
        if Primitive.to_proto(resource.allow_csr_based_issuance):
            res.allow_csr_based_issuance = Primitive.to_proto(
                resource.allow_csr_based_issuance
            )
        if Primitive.to_proto(resource.allow_config_based_issuance):
            res.allow_config_based_issuance = Primitive.to_proto(
                resource.allow_config_based_issuance
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicyAllowedIssuanceModes(
            allow_csr_based_issuance=Primitive.from_proto(
                resource.allow_csr_based_issuance
            ),
            allow_config_based_issuance=Primitive.from_proto(
                resource.allow_config_based_issuance
            ),
        )


class CaPoolIssuancePolicyAllowedIssuanceModesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CaPoolIssuancePolicyAllowedIssuanceModes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyAllowedIssuanceModes.from_proto(i) for i in resources
        ]


class CaPoolIssuancePolicyBaselineValues(object):
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

        res = ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyBaselineValues()
        if CaPoolIssuancePolicyBaselineValuesKeyUsage.to_proto(resource.key_usage):
            res.key_usage.CopyFrom(
                CaPoolIssuancePolicyBaselineValuesKeyUsage.to_proto(resource.key_usage)
            )
        else:
            res.ClearField("key_usage")
        if CaPoolIssuancePolicyBaselineValuesCaOptions.to_proto(resource.ca_options):
            res.ca_options.CopyFrom(
                CaPoolIssuancePolicyBaselineValuesCaOptions.to_proto(
                    resource.ca_options
                )
            )
        else:
            res.ClearField("ca_options")
        if CaPoolIssuancePolicyBaselineValuesPolicyIdsArray.to_proto(
            resource.policy_ids
        ):
            res.policy_ids.extend(
                CaPoolIssuancePolicyBaselineValuesPolicyIdsArray.to_proto(
                    resource.policy_ids
                )
            )
        if Primitive.to_proto(resource.aia_ocsp_servers):
            res.aia_ocsp_servers.extend(Primitive.to_proto(resource.aia_ocsp_servers))
        if CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicyBaselineValues(
            key_usage=CaPoolIssuancePolicyBaselineValuesKeyUsage.from_proto(
                resource.key_usage
            ),
            ca_options=CaPoolIssuancePolicyBaselineValuesCaOptions.from_proto(
                resource.ca_options
            ),
            policy_ids=CaPoolIssuancePolicyBaselineValuesPolicyIdsArray.from_proto(
                resource.policy_ids
            ),
            aia_ocsp_servers=Primitive.from_proto(resource.aia_ocsp_servers),
            additional_extensions=CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CaPoolIssuancePolicyBaselineValuesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CaPoolIssuancePolicyBaselineValues.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CaPoolIssuancePolicyBaselineValues.from_proto(i) for i in resources]


class CaPoolIssuancePolicyBaselineValuesKeyUsage(object):
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

        res = ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsage()
        if CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage.to_proto(
            resource.base_key_usage
        ):
            res.base_key_usage.CopyFrom(
                CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage.to_proto(
                    resource.base_key_usage
                )
            )
        else:
            res.ClearField("base_key_usage")
        if CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage.to_proto(
            resource.extended_key_usage
        ):
            res.extended_key_usage.CopyFrom(
                CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage.to_proto(
                    resource.extended_key_usage
                )
            )
        else:
            res.ClearField("extended_key_usage")
        if CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
            resource.unknown_extended_key_usages
        ):
            res.unknown_extended_key_usages.extend(
                CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
                    resource.unknown_extended_key_usages
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicyBaselineValuesKeyUsage(
            base_key_usage=CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage.from_proto(
                resource.base_key_usage
            ),
            extended_key_usage=CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage.from_proto(
                resource.extended_key_usage
            ),
            unknown_extended_key_usages=CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesArray.from_proto(
                resource.unknown_extended_key_usages
            ),
        )


class CaPoolIssuancePolicyBaselineValuesKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyBaselineValuesKeyUsage.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyBaselineValuesKeyUsage.from_proto(i) for i in resources
        ]


class CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage(object):
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
            ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage()
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

        return CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage(
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


class CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage.from_proto(i)
            for i in resources
        ]


class CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage(object):
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
            ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage()
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

        return CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage(
            server_auth=Primitive.from_proto(resource.server_auth),
            client_auth=Primitive.from_proto(resource.client_auth),
            code_signing=Primitive.from_proto(resource.code_signing),
            email_protection=Primitive.from_proto(resource.email_protection),
            time_stamping=Primitive.from_proto(resource.time_stamping),
            ocsp_signing=Primitive.from_proto(resource.ocsp_signing),
        )


class CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage.from_proto(i)
            for i in resources
        ]


class CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages.from_proto(
                i
            )
            for i in resources
        ]


class CaPoolIssuancePolicyBaselineValuesCaOptions(object):
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

        res = ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesCaOptions()
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

        return CaPoolIssuancePolicyBaselineValuesCaOptions(
            is_ca=Primitive.from_proto(resource.is_ca),
            max_issuer_path_length=Primitive.from_proto(
                resource.max_issuer_path_length
            ),
            zero_max_issuer_path_length=Primitive.from_proto(
                resource.zero_max_issuer_path_length
            ),
        )


class CaPoolIssuancePolicyBaselineValuesCaOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyBaselineValuesCaOptions.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyBaselineValuesCaOptions.from_proto(i) for i in resources
        ]


class CaPoolIssuancePolicyBaselineValuesPolicyIds(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesPolicyIds()
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicyBaselineValuesPolicyIds(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CaPoolIssuancePolicyBaselineValuesPolicyIdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyBaselineValuesPolicyIds.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyBaselineValuesPolicyIds.from_proto(i) for i in resources
        ]


class CaPoolIssuancePolicyBaselineValuesAdditionalExtensions(object):
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
            ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions()
        )
        if CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId.to_proto(
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

        return CaPoolIssuancePolicyBaselineValuesAdditionalExtensions(
            object_id=CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyBaselineValuesAdditionalExtensions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyBaselineValuesAdditionalExtensions.from_proto(i)
            for i in resources
        ]


class CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId.from_proto(i)
            for i in resources
        ]


class CaPoolIssuancePolicyIdentityConstraints(object):
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

        res = ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraints()
        if CaPoolIssuancePolicyIdentityConstraintsCelExpression.to_proto(
            resource.cel_expression
        ):
            res.cel_expression.CopyFrom(
                CaPoolIssuancePolicyIdentityConstraintsCelExpression.to_proto(
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

        return CaPoolIssuancePolicyIdentityConstraints(
            cel_expression=CaPoolIssuancePolicyIdentityConstraintsCelExpression.from_proto(
                resource.cel_expression
            ),
            allow_subject_passthrough=Primitive.from_proto(
                resource.allow_subject_passthrough
            ),
            allow_subject_alt_names_passthrough=Primitive.from_proto(
                resource.allow_subject_alt_names_passthrough
            ),
        )


class CaPoolIssuancePolicyIdentityConstraintsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CaPoolIssuancePolicyIdentityConstraints.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyIdentityConstraints.from_proto(i) for i in resources
        ]


class CaPoolIssuancePolicyIdentityConstraintsCelExpression(object):
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
            ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyIdentityConstraintsCelExpression()
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

        return CaPoolIssuancePolicyIdentityConstraintsCelExpression(
            expression=Primitive.from_proto(resource.expression),
            title=Primitive.from_proto(resource.title),
            description=Primitive.from_proto(resource.description),
            location=Primitive.from_proto(resource.location),
        )


class CaPoolIssuancePolicyIdentityConstraintsCelExpressionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyIdentityConstraintsCelExpression.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyIdentityConstraintsCelExpression.from_proto(i)
            for i in resources
        ]


class CaPoolIssuancePolicyPassthroughExtensions(object):
    def __init__(
        self, known_extensions: list = None, additional_extensions: list = None
    ):
        self.known_extensions = known_extensions
        self.additional_extensions = additional_extensions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensions()
        if CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnumArray.to_proto(
            resource.known_extensions
        ):
            res.known_extensions.extend(
                CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnumArray.to_proto(
                    resource.known_extensions
                )
            )
        if CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicyPassthroughExtensions(
            known_extensions=CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnumArray.from_proto(
                resource.known_extensions
            ),
            additional_extensions=CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CaPoolIssuancePolicyPassthroughExtensionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyPassthroughExtensions.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyPassthroughExtensions.from_proto(i) for i in resources
        ]


class CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions.from_proto(i)
            for i in resources
        ]


class CaPoolPublishingOptions(object):
    def __init__(self, publish_ca_cert: bool = None, publish_crl: bool = None):
        self.publish_ca_cert = publish_ca_cert
        self.publish_crl = publish_crl

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ca_pool_pb2.PrivatecaAlphaCaPoolPublishingOptions()
        if Primitive.to_proto(resource.publish_ca_cert):
            res.publish_ca_cert = Primitive.to_proto(resource.publish_ca_cert)
        if Primitive.to_proto(resource.publish_crl):
            res.publish_crl = Primitive.to_proto(resource.publish_crl)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CaPoolPublishingOptions(
            publish_ca_cert=Primitive.from_proto(resource.publish_ca_cert),
            publish_crl=Primitive.from_proto(resource.publish_crl),
        )


class CaPoolPublishingOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CaPoolPublishingOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CaPoolPublishingOptions.from_proto(i) for i in resources]


class CaPoolTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return ca_pool_pb2.PrivatecaAlphaCaPoolTierEnum.Value(
            "PrivatecaAlphaCaPoolTierEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return ca_pool_pb2.PrivatecaAlphaCaPoolTierEnum.Name(resource)[
            len("PrivatecaAlphaCaPoolTierEnum") :
        ]


class CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum.Value(
            "PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum.Name(
            resource
        )[
            len(
                "PrivatecaAlphaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum"
            ) :
        ]


class CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum.Value(
            "PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return ca_pool_pb2.PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum.Name(
            resource
        )[
            len(
                "PrivatecaAlphaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum"
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
