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
from google3.cloud.graphite.mmv2.services.google.binary_authorization import (
    attestor_pb2,
)
from google3.cloud.graphite.mmv2.services.google.binary_authorization import (
    attestor_pb2_grpc,
)

from typing import List


class Attestor(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        user_owned_drydock_note: dict = None,
        update_time: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.user_owned_drydock_note = user_owned_drydock_note
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = attestor_pb2_grpc.BinaryauthorizationAttestorServiceStub(
            channel.Channel()
        )
        request = attestor_pb2.ApplyBinaryauthorizationAttestorRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if AttestorUserOwnedDrydockNote.to_proto(self.user_owned_drydock_note):
            request.resource.user_owned_drydock_note.CopyFrom(
                AttestorUserOwnedDrydockNote.to_proto(self.user_owned_drydock_note)
            )
        else:
            request.resource.ClearField("user_owned_drydock_note")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyBinaryauthorizationAttestor(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.user_owned_drydock_note = AttestorUserOwnedDrydockNote.from_proto(
            response.user_owned_drydock_note
        )
        self.update_time = Primitive.from_proto(response.update_time)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = attestor_pb2_grpc.BinaryauthorizationAttestorServiceStub(
            channel.Channel()
        )
        request = attestor_pb2.DeleteBinaryauthorizationAttestorRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if AttestorUserOwnedDrydockNote.to_proto(self.user_owned_drydock_note):
            request.resource.user_owned_drydock_note.CopyFrom(
                AttestorUserOwnedDrydockNote.to_proto(self.user_owned_drydock_note)
            )
        else:
            request.resource.ClearField("user_owned_drydock_note")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteBinaryauthorizationAttestor(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = attestor_pb2_grpc.BinaryauthorizationAttestorServiceStub(
            channel.Channel()
        )
        request = attestor_pb2.ListBinaryauthorizationAttestorRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListBinaryauthorizationAttestor(request).items

    def to_proto(self):
        resource = attestor_pb2.BinaryauthorizationAttestor()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if AttestorUserOwnedDrydockNote.to_proto(self.user_owned_drydock_note):
            resource.user_owned_drydock_note.CopyFrom(
                AttestorUserOwnedDrydockNote.to_proto(self.user_owned_drydock_note)
            )
        else:
            resource.ClearField("user_owned_drydock_note")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class AttestorUserOwnedDrydockNote(object):
    def __init__(
        self,
        note_reference: str = None,
        public_keys: list = None,
        delegation_service_account_email: str = None,
    ):
        self.note_reference = note_reference
        self.public_keys = public_keys
        self.delegation_service_account_email = delegation_service_account_email

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = attestor_pb2.BinaryauthorizationAttestorUserOwnedDrydockNote()
        if Primitive.to_proto(resource.note_reference):
            res.note_reference = Primitive.to_proto(resource.note_reference)
        if AttestorUserOwnedDrydockNotePublicKeysArray.to_proto(resource.public_keys):
            res.public_keys.extend(
                AttestorUserOwnedDrydockNotePublicKeysArray.to_proto(
                    resource.public_keys
                )
            )
        if Primitive.to_proto(resource.delegation_service_account_email):
            res.delegation_service_account_email = Primitive.to_proto(
                resource.delegation_service_account_email
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AttestorUserOwnedDrydockNote(
            note_reference=Primitive.from_proto(resource.note_reference),
            public_keys=AttestorUserOwnedDrydockNotePublicKeysArray.from_proto(
                resource.public_keys
            ),
            delegation_service_account_email=Primitive.from_proto(
                resource.delegation_service_account_email
            ),
        )


class AttestorUserOwnedDrydockNoteArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AttestorUserOwnedDrydockNote.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AttestorUserOwnedDrydockNote.from_proto(i) for i in resources]


class AttestorUserOwnedDrydockNotePublicKeys(object):
    def __init__(
        self,
        comment: str = None,
        id: str = None,
        ascii_armored_pgp_public_key: str = None,
        pkix_public_key: dict = None,
    ):
        self.comment = comment
        self.id = id
        self.ascii_armored_pgp_public_key = ascii_armored_pgp_public_key
        self.pkix_public_key = pkix_public_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = attestor_pb2.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys()
        if Primitive.to_proto(resource.comment):
            res.comment = Primitive.to_proto(resource.comment)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.ascii_armored_pgp_public_key):
            res.ascii_armored_pgp_public_key = Primitive.to_proto(
                resource.ascii_armored_pgp_public_key
            )
        if AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey.to_proto(
            resource.pkix_public_key
        ):
            res.pkix_public_key.CopyFrom(
                AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey.to_proto(
                    resource.pkix_public_key
                )
            )
        else:
            res.ClearField("pkix_public_key")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AttestorUserOwnedDrydockNotePublicKeys(
            comment=Primitive.from_proto(resource.comment),
            id=Primitive.from_proto(resource.id),
            ascii_armored_pgp_public_key=Primitive.from_proto(
                resource.ascii_armored_pgp_public_key
            ),
            pkix_public_key=AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey.from_proto(
                resource.pkix_public_key
            ),
        )


class AttestorUserOwnedDrydockNotePublicKeysArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AttestorUserOwnedDrydockNotePublicKeys.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AttestorUserOwnedDrydockNotePublicKeys.from_proto(i) for i in resources]


class AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(object):
    def __init__(self, public_key_pem: str = None, signature_algorithm: str = None):
        self.public_key_pem = public_key_pem
        self.signature_algorithm = signature_algorithm

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            attestor_pb2.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey()
        )
        if Primitive.to_proto(resource.public_key_pem):
            res.public_key_pem = Primitive.to_proto(resource.public_key_pem)
        if AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum.to_proto(
            resource.signature_algorithm
        ):
            res.signature_algorithm = AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum.to_proto(
                resource.signature_algorithm
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(
            public_key_pem=Primitive.from_proto(resource.public_key_pem),
            signature_algorithm=AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum.from_proto(
                resource.signature_algorithm
            ),
        )


class AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey.from_proto(i)
            for i in resources
        ]


class AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return attestor_pb2.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum.Value(
            "BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return attestor_pb2.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum.Name(
            resource
        )[
            len(
                "BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"
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
