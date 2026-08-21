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
from google3.cloud.graphite.mmv2.services.google.cloudkms import crypto_key_pb2
from google3.cloud.graphite.mmv2.services.google.cloudkms import crypto_key_pb2_grpc

from typing import List


class CryptoKey(object):
    def __init__(
        self,
        name: str = None,
        primary: dict = None,
        purpose: str = None,
        create_time: str = None,
        next_rotation_time: str = None,
        rotation_period: str = None,
        version_template: dict = None,
        labels: dict = None,
        import_only: bool = None,
        destroy_scheduled_duration: str = None,
        project: str = None,
        location: str = None,
        key_ring: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.purpose = purpose
        self.next_rotation_time = next_rotation_time
        self.rotation_period = rotation_period
        self.version_template = version_template
        self.labels = labels
        self.import_only = import_only
        self.destroy_scheduled_duration = destroy_scheduled_duration
        self.project = project
        self.location = location
        self.key_ring = key_ring
        self.service_account_file = service_account_file

    def apply(self):
        stub = crypto_key_pb2_grpc.CloudkmsCryptoKeyServiceStub(channel.Channel())
        request = crypto_key_pb2.ApplyCloudkmsCryptoKeyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if CryptoKeyPurposeEnum.to_proto(self.purpose):
            request.resource.purpose = CryptoKeyPurposeEnum.to_proto(self.purpose)

        if Primitive.to_proto(self.next_rotation_time):
            request.resource.next_rotation_time = Primitive.to_proto(
                self.next_rotation_time
            )

        if Primitive.to_proto(self.rotation_period):
            request.resource.rotation_period = Primitive.to_proto(self.rotation_period)

        if CryptoKeyVersionTemplate.to_proto(self.version_template):
            request.resource.version_template.CopyFrom(
                CryptoKeyVersionTemplate.to_proto(self.version_template)
            )
        else:
            request.resource.ClearField("version_template")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.import_only):
            request.resource.import_only = Primitive.to_proto(self.import_only)

        if Primitive.to_proto(self.destroy_scheduled_duration):
            request.resource.destroy_scheduled_duration = Primitive.to_proto(
                self.destroy_scheduled_duration
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.key_ring):
            request.resource.key_ring = Primitive.to_proto(self.key_ring)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudkmsCryptoKey(request)
        self.name = Primitive.from_proto(response.name)
        self.primary = CryptoKeyPrimary.from_proto(response.primary)
        self.purpose = CryptoKeyPurposeEnum.from_proto(response.purpose)
        self.create_time = Primitive.from_proto(response.create_time)
        self.next_rotation_time = Primitive.from_proto(response.next_rotation_time)
        self.rotation_period = Primitive.from_proto(response.rotation_period)
        self.version_template = CryptoKeyVersionTemplate.from_proto(
            response.version_template
        )
        self.labels = Primitive.from_proto(response.labels)
        self.import_only = Primitive.from_proto(response.import_only)
        self.destroy_scheduled_duration = Primitive.from_proto(
            response.destroy_scheduled_duration
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.key_ring = Primitive.from_proto(response.key_ring)

    def delete(self):
        stub = crypto_key_pb2_grpc.CloudkmsCryptoKeyServiceStub(channel.Channel())
        request = crypto_key_pb2.DeleteCloudkmsCryptoKeyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if CryptoKeyPurposeEnum.to_proto(self.purpose):
            request.resource.purpose = CryptoKeyPurposeEnum.to_proto(self.purpose)

        if Primitive.to_proto(self.next_rotation_time):
            request.resource.next_rotation_time = Primitive.to_proto(
                self.next_rotation_time
            )

        if Primitive.to_proto(self.rotation_period):
            request.resource.rotation_period = Primitive.to_proto(self.rotation_period)

        if CryptoKeyVersionTemplate.to_proto(self.version_template):
            request.resource.version_template.CopyFrom(
                CryptoKeyVersionTemplate.to_proto(self.version_template)
            )
        else:
            request.resource.ClearField("version_template")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.import_only):
            request.resource.import_only = Primitive.to_proto(self.import_only)

        if Primitive.to_proto(self.destroy_scheduled_duration):
            request.resource.destroy_scheduled_duration = Primitive.to_proto(
                self.destroy_scheduled_duration
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.key_ring):
            request.resource.key_ring = Primitive.to_proto(self.key_ring)

        response = stub.DeleteCloudkmsCryptoKey(request)

    @classmethod
    def list(self, project, location, keyRing, service_account_file=""):
        stub = crypto_key_pb2_grpc.CloudkmsCryptoKeyServiceStub(channel.Channel())
        request = crypto_key_pb2.ListCloudkmsCryptoKeyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.KeyRing = keyRing

        return stub.ListCloudkmsCryptoKey(request).items

    def to_proto(self):
        resource = crypto_key_pb2.CloudkmsCryptoKey()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if CryptoKeyPurposeEnum.to_proto(self.purpose):
            resource.purpose = CryptoKeyPurposeEnum.to_proto(self.purpose)
        if Primitive.to_proto(self.next_rotation_time):
            resource.next_rotation_time = Primitive.to_proto(self.next_rotation_time)
        if Primitive.to_proto(self.rotation_period):
            resource.rotation_period = Primitive.to_proto(self.rotation_period)
        if CryptoKeyVersionTemplate.to_proto(self.version_template):
            resource.version_template.CopyFrom(
                CryptoKeyVersionTemplate.to_proto(self.version_template)
            )
        else:
            resource.ClearField("version_template")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.import_only):
            resource.import_only = Primitive.to_proto(self.import_only)
        if Primitive.to_proto(self.destroy_scheduled_duration):
            resource.destroy_scheduled_duration = Primitive.to_proto(
                self.destroy_scheduled_duration
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.key_ring):
            resource.key_ring = Primitive.to_proto(self.key_ring)
        return resource


class CryptoKeyPrimary(object):
    def __init__(
        self,
        name: str = None,
        state: str = None,
        protection_level: str = None,
        algorithm: str = None,
        attestation: dict = None,
        create_time: str = None,
        generate_time: str = None,
        destroy_time: str = None,
        destroy_event_time: str = None,
        import_job: str = None,
        import_time: str = None,
        import_failure_reason: str = None,
        external_protection_level_options: dict = None,
        reimport_eligible: bool = None,
    ):
        self.name = name
        self.state = state
        self.protection_level = protection_level
        self.algorithm = algorithm
        self.attestation = attestation
        self.create_time = create_time
        self.generate_time = generate_time
        self.destroy_time = destroy_time
        self.destroy_event_time = destroy_event_time
        self.import_job = import_job
        self.import_time = import_time
        self.import_failure_reason = import_failure_reason
        self.external_protection_level_options = external_protection_level_options
        self.reimport_eligible = reimport_eligible

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = crypto_key_pb2.CloudkmsCryptoKeyPrimary()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if CryptoKeyPrimaryStateEnum.to_proto(resource.state):
            res.state = CryptoKeyPrimaryStateEnum.to_proto(resource.state)
        if CryptoKeyPrimaryProtectionLevelEnum.to_proto(resource.protection_level):
            res.protection_level = CryptoKeyPrimaryProtectionLevelEnum.to_proto(
                resource.protection_level
            )
        if CryptoKeyPrimaryAlgorithmEnum.to_proto(resource.algorithm):
            res.algorithm = CryptoKeyPrimaryAlgorithmEnum.to_proto(resource.algorithm)
        if CryptoKeyPrimaryAttestation.to_proto(resource.attestation):
            res.attestation.CopyFrom(
                CryptoKeyPrimaryAttestation.to_proto(resource.attestation)
            )
        else:
            res.ClearField("attestation")
        if Primitive.to_proto(resource.create_time):
            res.create_time = Primitive.to_proto(resource.create_time)
        if Primitive.to_proto(resource.generate_time):
            res.generate_time = Primitive.to_proto(resource.generate_time)
        if Primitive.to_proto(resource.destroy_time):
            res.destroy_time = Primitive.to_proto(resource.destroy_time)
        if Primitive.to_proto(resource.destroy_event_time):
            res.destroy_event_time = Primitive.to_proto(resource.destroy_event_time)
        if Primitive.to_proto(resource.import_job):
            res.import_job = Primitive.to_proto(resource.import_job)
        if Primitive.to_proto(resource.import_time):
            res.import_time = Primitive.to_proto(resource.import_time)
        if Primitive.to_proto(resource.import_failure_reason):
            res.import_failure_reason = Primitive.to_proto(
                resource.import_failure_reason
            )
        if CryptoKeyPrimaryExternalProtectionLevelOptions.to_proto(
            resource.external_protection_level_options
        ):
            res.external_protection_level_options.CopyFrom(
                CryptoKeyPrimaryExternalProtectionLevelOptions.to_proto(
                    resource.external_protection_level_options
                )
            )
        else:
            res.ClearField("external_protection_level_options")
        if Primitive.to_proto(resource.reimport_eligible):
            res.reimport_eligible = Primitive.to_proto(resource.reimport_eligible)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CryptoKeyPrimary(
            name=Primitive.from_proto(resource.name),
            state=CryptoKeyPrimaryStateEnum.from_proto(resource.state),
            protection_level=CryptoKeyPrimaryProtectionLevelEnum.from_proto(
                resource.protection_level
            ),
            algorithm=CryptoKeyPrimaryAlgorithmEnum.from_proto(resource.algorithm),
            attestation=CryptoKeyPrimaryAttestation.from_proto(resource.attestation),
            create_time=Primitive.from_proto(resource.create_time),
            generate_time=Primitive.from_proto(resource.generate_time),
            destroy_time=Primitive.from_proto(resource.destroy_time),
            destroy_event_time=Primitive.from_proto(resource.destroy_event_time),
            import_job=Primitive.from_proto(resource.import_job),
            import_time=Primitive.from_proto(resource.import_time),
            import_failure_reason=Primitive.from_proto(resource.import_failure_reason),
            external_protection_level_options=CryptoKeyPrimaryExternalProtectionLevelOptions.from_proto(
                resource.external_protection_level_options
            ),
            reimport_eligible=Primitive.from_proto(resource.reimport_eligible),
        )


class CryptoKeyPrimaryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CryptoKeyPrimary.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CryptoKeyPrimary.from_proto(i) for i in resources]


class CryptoKeyPrimaryAttestation(object):
    def __init__(
        self, format: str = None, content: str = None, cert_chains: dict = None
    ):
        self.format = format
        self.content = content
        self.cert_chains = cert_chains

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = crypto_key_pb2.CloudkmsCryptoKeyPrimaryAttestation()
        if CryptoKeyPrimaryAttestationFormatEnum.to_proto(resource.format):
            res.format = CryptoKeyPrimaryAttestationFormatEnum.to_proto(resource.format)
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if CryptoKeyPrimaryAttestationCertChains.to_proto(resource.cert_chains):
            res.cert_chains.CopyFrom(
                CryptoKeyPrimaryAttestationCertChains.to_proto(resource.cert_chains)
            )
        else:
            res.ClearField("cert_chains")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CryptoKeyPrimaryAttestation(
            format=CryptoKeyPrimaryAttestationFormatEnum.from_proto(resource.format),
            content=Primitive.from_proto(resource.content),
            cert_chains=CryptoKeyPrimaryAttestationCertChains.from_proto(
                resource.cert_chains
            ),
        )


class CryptoKeyPrimaryAttestationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CryptoKeyPrimaryAttestation.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CryptoKeyPrimaryAttestation.from_proto(i) for i in resources]


class CryptoKeyPrimaryAttestationCertChains(object):
    def __init__(
        self,
        cavium_certs: list = None,
        google_card_certs: list = None,
        google_partition_certs: list = None,
    ):
        self.cavium_certs = cavium_certs
        self.google_card_certs = google_card_certs
        self.google_partition_certs = google_partition_certs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = crypto_key_pb2.CloudkmsCryptoKeyPrimaryAttestationCertChains()
        if Primitive.to_proto(resource.cavium_certs):
            res.cavium_certs.extend(Primitive.to_proto(resource.cavium_certs))
        if Primitive.to_proto(resource.google_card_certs):
            res.google_card_certs.extend(Primitive.to_proto(resource.google_card_certs))
        if Primitive.to_proto(resource.google_partition_certs):
            res.google_partition_certs.extend(
                Primitive.to_proto(resource.google_partition_certs)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CryptoKeyPrimaryAttestationCertChains(
            cavium_certs=Primitive.from_proto(resource.cavium_certs),
            google_card_certs=Primitive.from_proto(resource.google_card_certs),
            google_partition_certs=Primitive.from_proto(
                resource.google_partition_certs
            ),
        )


class CryptoKeyPrimaryAttestationCertChainsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CryptoKeyPrimaryAttestationCertChains.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CryptoKeyPrimaryAttestationCertChains.from_proto(i) for i in resources]


class CryptoKeyPrimaryExternalProtectionLevelOptions(object):
    def __init__(self, external_key_uri: str = None):
        self.external_key_uri = external_key_uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = crypto_key_pb2.CloudkmsCryptoKeyPrimaryExternalProtectionLevelOptions()
        if Primitive.to_proto(resource.external_key_uri):
            res.external_key_uri = Primitive.to_proto(resource.external_key_uri)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CryptoKeyPrimaryExternalProtectionLevelOptions(
            external_key_uri=Primitive.from_proto(resource.external_key_uri),
        )


class CryptoKeyPrimaryExternalProtectionLevelOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CryptoKeyPrimaryExternalProtectionLevelOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CryptoKeyPrimaryExternalProtectionLevelOptions.from_proto(i)
            for i in resources
        ]


class CryptoKeyVersionTemplate(object):
    def __init__(self, protection_level: str = None, algorithm: str = None):
        self.protection_level = protection_level
        self.algorithm = algorithm

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = crypto_key_pb2.CloudkmsCryptoKeyVersionTemplate()
        if CryptoKeyVersionTemplateProtectionLevelEnum.to_proto(
            resource.protection_level
        ):
            res.protection_level = CryptoKeyVersionTemplateProtectionLevelEnum.to_proto(
                resource.protection_level
            )
        if CryptoKeyVersionTemplateAlgorithmEnum.to_proto(resource.algorithm):
            res.algorithm = CryptoKeyVersionTemplateAlgorithmEnum.to_proto(
                resource.algorithm
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CryptoKeyVersionTemplate(
            protection_level=CryptoKeyVersionTemplateProtectionLevelEnum.from_proto(
                resource.protection_level
            ),
            algorithm=CryptoKeyVersionTemplateAlgorithmEnum.from_proto(
                resource.algorithm
            ),
        )


class CryptoKeyVersionTemplateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CryptoKeyVersionTemplate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CryptoKeyVersionTemplate.from_proto(i) for i in resources]


class CryptoKeyPrimaryStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyPrimaryStateEnum.Value(
            "CloudkmsCryptoKeyPrimaryStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyPrimaryStateEnum.Name(resource)[
            len("CloudkmsCryptoKeyPrimaryStateEnum") :
        ]


class CryptoKeyPrimaryProtectionLevelEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyPrimaryProtectionLevelEnum.Value(
            "CloudkmsCryptoKeyPrimaryProtectionLevelEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyPrimaryProtectionLevelEnum.Name(
            resource
        )[len("CloudkmsCryptoKeyPrimaryProtectionLevelEnum") :]


class CryptoKeyPrimaryAlgorithmEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyPrimaryAlgorithmEnum.Value(
            "CloudkmsCryptoKeyPrimaryAlgorithmEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyPrimaryAlgorithmEnum.Name(resource)[
            len("CloudkmsCryptoKeyPrimaryAlgorithmEnum") :
        ]


class CryptoKeyPrimaryAttestationFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyPrimaryAttestationFormatEnum.Value(
            "CloudkmsCryptoKeyPrimaryAttestationFormatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyPrimaryAttestationFormatEnum.Name(
            resource
        )[len("CloudkmsCryptoKeyPrimaryAttestationFormatEnum") :]


class CryptoKeyPurposeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyPurposeEnum.Value(
            "CloudkmsCryptoKeyPurposeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyPurposeEnum.Name(resource)[
            len("CloudkmsCryptoKeyPurposeEnum") :
        ]


class CryptoKeyVersionTemplateProtectionLevelEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum.Value(
            "CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum.Name(
            resource
        )[len("CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum") :]


class CryptoKeyVersionTemplateAlgorithmEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum.Value(
            "CloudkmsCryptoKeyVersionTemplateAlgorithmEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return crypto_key_pb2.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum.Name(
            resource
        )[len("CloudkmsCryptoKeyVersionTemplateAlgorithmEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
