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
from google3.cloud.graphite.mmv2.services.google.compute import (
    interconnect_attachment_pb2,
)
from google3.cloud.graphite.mmv2.services.google.compute import (
    interconnect_attachment_pb2_grpc,
)

from typing import List


class InterconnectAttachment(object):
    def __init__(
        self,
        description: str = None,
        self_link: str = None,
        id: int = None,
        name: str = None,
        interconnect: str = None,
        router: str = None,
        region: str = None,
        mtu: int = None,
        private_interconnect_info: dict = None,
        operational_status: str = None,
        cloud_router_ip_address: str = None,
        customer_router_ip_address: str = None,
        type: str = None,
        pairing_key: str = None,
        admin_enabled: bool = None,
        vlan_tag8021q: int = None,
        edge_availability_domain: str = None,
        candidate_subnets: list = None,
        bandwidth: str = None,
        partner_metadata: dict = None,
        state: str = None,
        partner_asn: int = None,
        encryption: str = None,
        ipsec_internal_addresses: list = None,
        dataplane_version: int = None,
        satisfies_pzs: bool = None,
        labels: dict = None,
        label_fingerprint: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.description = description
        self.name = name
        self.interconnect = interconnect
        self.router = router
        self.region = region
        self.mtu = mtu
        self.type = type
        self.pairing_key = pairing_key
        self.admin_enabled = admin_enabled
        self.vlan_tag8021q = vlan_tag8021q
        self.edge_availability_domain = edge_availability_domain
        self.candidate_subnets = candidate_subnets
        self.bandwidth = bandwidth
        self.partner_metadata = partner_metadata
        self.partner_asn = partner_asn
        self.encryption = encryption
        self.ipsec_internal_addresses = ipsec_internal_addresses
        self.dataplane_version = dataplane_version
        self.labels = labels
        self.label_fingerprint = label_fingerprint
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = interconnect_attachment_pb2_grpc.ComputeAlphaInterconnectAttachmentServiceStub(
            channel.Channel()
        )
        request = (
            interconnect_attachment_pb2.ApplyComputeAlphaInterconnectAttachmentRequest()
        )
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.interconnect):
            request.resource.interconnect = Primitive.to_proto(self.interconnect)

        if Primitive.to_proto(self.router):
            request.resource.router = Primitive.to_proto(self.router)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.mtu):
            request.resource.mtu = Primitive.to_proto(self.mtu)

        if InterconnectAttachmentTypeEnum.to_proto(self.type):
            request.resource.type = InterconnectAttachmentTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.pairing_key):
            request.resource.pairing_key = Primitive.to_proto(self.pairing_key)

        if Primitive.to_proto(self.admin_enabled):
            request.resource.admin_enabled = Primitive.to_proto(self.admin_enabled)

        if Primitive.to_proto(self.vlan_tag8021q):
            request.resource.vlan_tag8021q = Primitive.to_proto(self.vlan_tag8021q)

        if InterconnectAttachmentEdgeAvailabilityDomainEnum.to_proto(
            self.edge_availability_domain
        ):
            request.resource.edge_availability_domain = (
                InterconnectAttachmentEdgeAvailabilityDomainEnum.to_proto(
                    self.edge_availability_domain
                )
            )

        if Primitive.to_proto(self.candidate_subnets):
            request.resource.candidate_subnets.extend(
                Primitive.to_proto(self.candidate_subnets)
            )
        if InterconnectAttachmentBandwidthEnum.to_proto(self.bandwidth):
            request.resource.bandwidth = InterconnectAttachmentBandwidthEnum.to_proto(
                self.bandwidth
            )

        if InterconnectAttachmentPartnerMetadata.to_proto(self.partner_metadata):
            request.resource.partner_metadata.CopyFrom(
                InterconnectAttachmentPartnerMetadata.to_proto(self.partner_metadata)
            )
        else:
            request.resource.ClearField("partner_metadata")
        if Primitive.to_proto(self.partner_asn):
            request.resource.partner_asn = Primitive.to_proto(self.partner_asn)

        if InterconnectAttachmentEncryptionEnum.to_proto(self.encryption):
            request.resource.encryption = InterconnectAttachmentEncryptionEnum.to_proto(
                self.encryption
            )

        if Primitive.to_proto(self.ipsec_internal_addresses):
            request.resource.ipsec_internal_addresses.extend(
                Primitive.to_proto(self.ipsec_internal_addresses)
            )
        if Primitive.to_proto(self.dataplane_version):
            request.resource.dataplane_version = Primitive.to_proto(
                self.dataplane_version
            )

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.label_fingerprint):
            request.resource.label_fingerprint = Primitive.to_proto(
                self.label_fingerprint
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeAlphaInterconnectAttachment(request)
        self.description = Primitive.from_proto(response.description)
        self.self_link = Primitive.from_proto(response.self_link)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.interconnect = Primitive.from_proto(response.interconnect)
        self.router = Primitive.from_proto(response.router)
        self.region = Primitive.from_proto(response.region)
        self.mtu = Primitive.from_proto(response.mtu)
        self.private_interconnect_info = (
            InterconnectAttachmentPrivateInterconnectInfo.from_proto(
                response.private_interconnect_info
            )
        )
        self.operational_status = (
            InterconnectAttachmentOperationalStatusEnum.from_proto(
                response.operational_status
            )
        )
        self.cloud_router_ip_address = Primitive.from_proto(
            response.cloud_router_ip_address
        )
        self.customer_router_ip_address = Primitive.from_proto(
            response.customer_router_ip_address
        )
        self.type = InterconnectAttachmentTypeEnum.from_proto(response.type)
        self.pairing_key = Primitive.from_proto(response.pairing_key)
        self.admin_enabled = Primitive.from_proto(response.admin_enabled)
        self.vlan_tag8021q = Primitive.from_proto(response.vlan_tag8021q)
        self.edge_availability_domain = (
            InterconnectAttachmentEdgeAvailabilityDomainEnum.from_proto(
                response.edge_availability_domain
            )
        )
        self.candidate_subnets = Primitive.from_proto(response.candidate_subnets)
        self.bandwidth = InterconnectAttachmentBandwidthEnum.from_proto(
            response.bandwidth
        )
        self.partner_metadata = InterconnectAttachmentPartnerMetadata.from_proto(
            response.partner_metadata
        )
        self.state = InterconnectAttachmentStateEnum.from_proto(response.state)
        self.partner_asn = Primitive.from_proto(response.partner_asn)
        self.encryption = InterconnectAttachmentEncryptionEnum.from_proto(
            response.encryption
        )
        self.ipsec_internal_addresses = Primitive.from_proto(
            response.ipsec_internal_addresses
        )
        self.dataplane_version = Primitive.from_proto(response.dataplane_version)
        self.satisfies_pzs = Primitive.from_proto(response.satisfies_pzs)
        self.labels = Primitive.from_proto(response.labels)
        self.label_fingerprint = Primitive.from_proto(response.label_fingerprint)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = interconnect_attachment_pb2_grpc.ComputeAlphaInterconnectAttachmentServiceStub(
            channel.Channel()
        )
        request = (
            interconnect_attachment_pb2.DeleteComputeAlphaInterconnectAttachmentRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.interconnect):
            request.resource.interconnect = Primitive.to_proto(self.interconnect)

        if Primitive.to_proto(self.router):
            request.resource.router = Primitive.to_proto(self.router)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.mtu):
            request.resource.mtu = Primitive.to_proto(self.mtu)

        if InterconnectAttachmentTypeEnum.to_proto(self.type):
            request.resource.type = InterconnectAttachmentTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.pairing_key):
            request.resource.pairing_key = Primitive.to_proto(self.pairing_key)

        if Primitive.to_proto(self.admin_enabled):
            request.resource.admin_enabled = Primitive.to_proto(self.admin_enabled)

        if Primitive.to_proto(self.vlan_tag8021q):
            request.resource.vlan_tag8021q = Primitive.to_proto(self.vlan_tag8021q)

        if InterconnectAttachmentEdgeAvailabilityDomainEnum.to_proto(
            self.edge_availability_domain
        ):
            request.resource.edge_availability_domain = (
                InterconnectAttachmentEdgeAvailabilityDomainEnum.to_proto(
                    self.edge_availability_domain
                )
            )

        if Primitive.to_proto(self.candidate_subnets):
            request.resource.candidate_subnets.extend(
                Primitive.to_proto(self.candidate_subnets)
            )
        if InterconnectAttachmentBandwidthEnum.to_proto(self.bandwidth):
            request.resource.bandwidth = InterconnectAttachmentBandwidthEnum.to_proto(
                self.bandwidth
            )

        if InterconnectAttachmentPartnerMetadata.to_proto(self.partner_metadata):
            request.resource.partner_metadata.CopyFrom(
                InterconnectAttachmentPartnerMetadata.to_proto(self.partner_metadata)
            )
        else:
            request.resource.ClearField("partner_metadata")
        if Primitive.to_proto(self.partner_asn):
            request.resource.partner_asn = Primitive.to_proto(self.partner_asn)

        if InterconnectAttachmentEncryptionEnum.to_proto(self.encryption):
            request.resource.encryption = InterconnectAttachmentEncryptionEnum.to_proto(
                self.encryption
            )

        if Primitive.to_proto(self.ipsec_internal_addresses):
            request.resource.ipsec_internal_addresses.extend(
                Primitive.to_proto(self.ipsec_internal_addresses)
            )
        if Primitive.to_proto(self.dataplane_version):
            request.resource.dataplane_version = Primitive.to_proto(
                self.dataplane_version
            )

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.label_fingerprint):
            request.resource.label_fingerprint = Primitive.to_proto(
                self.label_fingerprint
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeAlphaInterconnectAttachment(request)

    @classmethod
    def list(self, project, region, service_account_file=""):
        stub = interconnect_attachment_pb2_grpc.ComputeAlphaInterconnectAttachmentServiceStub(
            channel.Channel()
        )
        request = (
            interconnect_attachment_pb2.ListComputeAlphaInterconnectAttachmentRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        return stub.ListComputeAlphaInterconnectAttachment(request).items

    def to_proto(self):
        resource = interconnect_attachment_pb2.ComputeAlphaInterconnectAttachment()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.interconnect):
            resource.interconnect = Primitive.to_proto(self.interconnect)
        if Primitive.to_proto(self.router):
            resource.router = Primitive.to_proto(self.router)
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.mtu):
            resource.mtu = Primitive.to_proto(self.mtu)
        if InterconnectAttachmentTypeEnum.to_proto(self.type):
            resource.type = InterconnectAttachmentTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.pairing_key):
            resource.pairing_key = Primitive.to_proto(self.pairing_key)
        if Primitive.to_proto(self.admin_enabled):
            resource.admin_enabled = Primitive.to_proto(self.admin_enabled)
        if Primitive.to_proto(self.vlan_tag8021q):
            resource.vlan_tag8021q = Primitive.to_proto(self.vlan_tag8021q)
        if InterconnectAttachmentEdgeAvailabilityDomainEnum.to_proto(
            self.edge_availability_domain
        ):
            resource.edge_availability_domain = (
                InterconnectAttachmentEdgeAvailabilityDomainEnum.to_proto(
                    self.edge_availability_domain
                )
            )
        if Primitive.to_proto(self.candidate_subnets):
            resource.candidate_subnets.extend(
                Primitive.to_proto(self.candidate_subnets)
            )
        if InterconnectAttachmentBandwidthEnum.to_proto(self.bandwidth):
            resource.bandwidth = InterconnectAttachmentBandwidthEnum.to_proto(
                self.bandwidth
            )
        if InterconnectAttachmentPartnerMetadata.to_proto(self.partner_metadata):
            resource.partner_metadata.CopyFrom(
                InterconnectAttachmentPartnerMetadata.to_proto(self.partner_metadata)
            )
        else:
            resource.ClearField("partner_metadata")
        if Primitive.to_proto(self.partner_asn):
            resource.partner_asn = Primitive.to_proto(self.partner_asn)
        if InterconnectAttachmentEncryptionEnum.to_proto(self.encryption):
            resource.encryption = InterconnectAttachmentEncryptionEnum.to_proto(
                self.encryption
            )
        if Primitive.to_proto(self.ipsec_internal_addresses):
            resource.ipsec_internal_addresses.extend(
                Primitive.to_proto(self.ipsec_internal_addresses)
            )
        if Primitive.to_proto(self.dataplane_version):
            resource.dataplane_version = Primitive.to_proto(self.dataplane_version)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.label_fingerprint):
            resource.label_fingerprint = Primitive.to_proto(self.label_fingerprint)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class InterconnectAttachmentPrivateInterconnectInfo(object):
    def __init__(self, tag8021q: int = None):
        self.tag8021q = tag8021q

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentPrivateInterconnectInfo()
        )
        if Primitive.to_proto(resource.tag8021q):
            res.tag8021q = Primitive.to_proto(resource.tag8021q)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InterconnectAttachmentPrivateInterconnectInfo(
            tag8021q=Primitive.from_proto(resource.tag8021q),
        )


class InterconnectAttachmentPrivateInterconnectInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InterconnectAttachmentPrivateInterconnectInfo.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InterconnectAttachmentPrivateInterconnectInfo.from_proto(i)
            for i in resources
        ]


class InterconnectAttachmentPartnerMetadata(object):
    def __init__(
        self,
        partner_name: str = None,
        interconnect_name: str = None,
        portal_url: str = None,
    ):
        self.partner_name = partner_name
        self.interconnect_name = interconnect_name
        self.portal_url = portal_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentPartnerMetadata()
        )
        if Primitive.to_proto(resource.partner_name):
            res.partner_name = Primitive.to_proto(resource.partner_name)
        if Primitive.to_proto(resource.interconnect_name):
            res.interconnect_name = Primitive.to_proto(resource.interconnect_name)
        if Primitive.to_proto(resource.portal_url):
            res.portal_url = Primitive.to_proto(resource.portal_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InterconnectAttachmentPartnerMetadata(
            partner_name=Primitive.from_proto(resource.partner_name),
            interconnect_name=Primitive.from_proto(resource.interconnect_name),
            portal_url=Primitive.from_proto(resource.portal_url),
        )


class InterconnectAttachmentPartnerMetadataArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InterconnectAttachmentPartnerMetadata.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InterconnectAttachmentPartnerMetadata.from_proto(i) for i in resources]


class InterconnectAttachmentOperationalStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentOperationalStatusEnum.Value(
            "ComputeAlphaInterconnectAttachmentOperationalStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentOperationalStatusEnum.Name(
            resource
        )[
            len("ComputeAlphaInterconnectAttachmentOperationalStatusEnum") :
        ]


class InterconnectAttachmentTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentTypeEnum.Value(
            "ComputeAlphaInterconnectAttachmentTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentTypeEnum.Name(
                resource
            )[len("ComputeAlphaInterconnectAttachmentTypeEnum") :]
        )


class InterconnectAttachmentEdgeAvailabilityDomainEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum.Value(
            "ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum.Name(
            resource
        )[
            len("ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum") :
        ]


class InterconnectAttachmentBandwidthEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentBandwidthEnum.Value(
            "ComputeAlphaInterconnectAttachmentBandwidthEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentBandwidthEnum.Name(
            resource
        )[
            len("ComputeAlphaInterconnectAttachmentBandwidthEnum") :
        ]


class InterconnectAttachmentStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentStateEnum.Value(
            "ComputeAlphaInterconnectAttachmentStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentStateEnum.Name(
            resource
        )[
            len("ComputeAlphaInterconnectAttachmentStateEnum") :
        ]


class InterconnectAttachmentEncryptionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentEncryptionEnum.Value(
            "ComputeAlphaInterconnectAttachmentEncryptionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return interconnect_attachment_pb2.ComputeAlphaInterconnectAttachmentEncryptionEnum.Name(
            resource
        )[
            len("ComputeAlphaInterconnectAttachmentEncryptionEnum") :
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
