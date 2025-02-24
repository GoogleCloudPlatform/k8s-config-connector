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
from google3.cloud.graphite.mmv2.services.google.gke_hub import membership_pb2
from google3.cloud.graphite.mmv2.services.google.gke_hub import membership_pb2_grpc

from typing import List


class Membership(object):
    def __init__(
        self,
        endpoint: dict = None,
        name: str = None,
        labels: dict = None,
        description: str = None,
        state: dict = None,
        create_time: str = None,
        update_time: str = None,
        delete_time: str = None,
        external_id: str = None,
        last_connection_time: str = None,
        unique_id: str = None,
        authority: dict = None,
        infrastructure_type: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.endpoint = endpoint
        self.name = name
        self.labels = labels
        self.description = description
        self.external_id = external_id
        self.authority = authority
        self.infrastructure_type = infrastructure_type
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = membership_pb2_grpc.GkehubBetaMembershipServiceStub(channel.Channel())
        request = membership_pb2.ApplyGkehubBetaMembershipRequest()
        if MembershipEndpoint.to_proto(self.endpoint):
            request.resource.endpoint.CopyFrom(
                MembershipEndpoint.to_proto(self.endpoint)
            )
        else:
            request.resource.ClearField("endpoint")
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.external_id):
            request.resource.external_id = Primitive.to_proto(self.external_id)

        if MembershipAuthority.to_proto(self.authority):
            request.resource.authority.CopyFrom(
                MembershipAuthority.to_proto(self.authority)
            )
        else:
            request.resource.ClearField("authority")
        if MembershipInfrastructureTypeEnum.to_proto(self.infrastructure_type):
            request.resource.infrastructure_type = (
                MembershipInfrastructureTypeEnum.to_proto(self.infrastructure_type)
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyGkehubBetaMembership(request)
        self.endpoint = MembershipEndpoint.from_proto(response.endpoint)
        self.name = Primitive.from_proto(response.name)
        self.labels = Primitive.from_proto(response.labels)
        self.description = Primitive.from_proto(response.description)
        self.state = MembershipState.from_proto(response.state)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.external_id = Primitive.from_proto(response.external_id)
        self.last_connection_time = Primitive.from_proto(response.last_connection_time)
        self.unique_id = Primitive.from_proto(response.unique_id)
        self.authority = MembershipAuthority.from_proto(response.authority)
        self.infrastructure_type = MembershipInfrastructureTypeEnum.from_proto(
            response.infrastructure_type
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = membership_pb2_grpc.GkehubBetaMembershipServiceStub(channel.Channel())
        request = membership_pb2.DeleteGkehubBetaMembershipRequest()
        request.service_account_file = self.service_account_file
        if MembershipEndpoint.to_proto(self.endpoint):
            request.resource.endpoint.CopyFrom(
                MembershipEndpoint.to_proto(self.endpoint)
            )
        else:
            request.resource.ClearField("endpoint")
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.external_id):
            request.resource.external_id = Primitive.to_proto(self.external_id)

        if MembershipAuthority.to_proto(self.authority):
            request.resource.authority.CopyFrom(
                MembershipAuthority.to_proto(self.authority)
            )
        else:
            request.resource.ClearField("authority")
        if MembershipInfrastructureTypeEnum.to_proto(self.infrastructure_type):
            request.resource.infrastructure_type = (
                MembershipInfrastructureTypeEnum.to_proto(self.infrastructure_type)
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteGkehubBetaMembership(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = membership_pb2_grpc.GkehubBetaMembershipServiceStub(channel.Channel())
        request = membership_pb2.ListGkehubBetaMembershipRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListGkehubBetaMembership(request).items

    def to_proto(self):
        resource = membership_pb2.GkehubBetaMembership()
        if MembershipEndpoint.to_proto(self.endpoint):
            resource.endpoint.CopyFrom(MembershipEndpoint.to_proto(self.endpoint))
        else:
            resource.ClearField("endpoint")
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.external_id):
            resource.external_id = Primitive.to_proto(self.external_id)
        if MembershipAuthority.to_proto(self.authority):
            resource.authority.CopyFrom(MembershipAuthority.to_proto(self.authority))
        else:
            resource.ClearField("authority")
        if MembershipInfrastructureTypeEnum.to_proto(self.infrastructure_type):
            resource.infrastructure_type = MembershipInfrastructureTypeEnum.to_proto(
                self.infrastructure_type
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class MembershipEndpoint(object):
    def __init__(
        self,
        gke_cluster: dict = None,
        kubernetes_metadata: dict = None,
        kubernetes_resource: dict = None,
    ):
        self.gke_cluster = gke_cluster
        self.kubernetes_metadata = kubernetes_metadata
        self.kubernetes_resource = kubernetes_resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.GkehubBetaMembershipEndpoint()
        if MembershipEndpointGkeCluster.to_proto(resource.gke_cluster):
            res.gke_cluster.CopyFrom(
                MembershipEndpointGkeCluster.to_proto(resource.gke_cluster)
            )
        else:
            res.ClearField("gke_cluster")
        if MembershipEndpointKubernetesMetadata.to_proto(resource.kubernetes_metadata):
            res.kubernetes_metadata.CopyFrom(
                MembershipEndpointKubernetesMetadata.to_proto(
                    resource.kubernetes_metadata
                )
            )
        else:
            res.ClearField("kubernetes_metadata")
        if MembershipEndpointKubernetesResource.to_proto(resource.kubernetes_resource):
            res.kubernetes_resource.CopyFrom(
                MembershipEndpointKubernetesResource.to_proto(
                    resource.kubernetes_resource
                )
            )
        else:
            res.ClearField("kubernetes_resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipEndpoint(
            gke_cluster=MembershipEndpointGkeCluster.from_proto(resource.gke_cluster),
            kubernetes_metadata=MembershipEndpointKubernetesMetadata.from_proto(
                resource.kubernetes_metadata
            ),
            kubernetes_resource=MembershipEndpointKubernetesResource.from_proto(
                resource.kubernetes_resource
            ),
        )


class MembershipEndpointArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipEndpoint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipEndpoint.from_proto(i) for i in resources]


class MembershipEndpointGkeCluster(object):
    def __init__(self, resource_link: str = None):
        self.resource_link = resource_link

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.GkehubBetaMembershipEndpointGkeCluster()
        if Primitive.to_proto(resource.resource_link):
            res.resource_link = Primitive.to_proto(resource.resource_link)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipEndpointGkeCluster(
            resource_link=Primitive.from_proto(resource.resource_link),
        )


class MembershipEndpointGkeClusterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipEndpointGkeCluster.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipEndpointGkeCluster.from_proto(i) for i in resources]


class MembershipEndpointKubernetesMetadata(object):
    def __init__(
        self,
        kubernetes_api_server_version: str = None,
        node_provider_id: str = None,
        node_count: int = None,
        vcpu_count: int = None,
        memory_mb: int = None,
        update_time: str = None,
    ):
        self.kubernetes_api_server_version = kubernetes_api_server_version
        self.node_provider_id = node_provider_id
        self.node_count = node_count
        self.vcpu_count = vcpu_count
        self.memory_mb = memory_mb
        self.update_time = update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.GkehubBetaMembershipEndpointKubernetesMetadata()
        if Primitive.to_proto(resource.kubernetes_api_server_version):
            res.kubernetes_api_server_version = Primitive.to_proto(
                resource.kubernetes_api_server_version
            )
        if Primitive.to_proto(resource.node_provider_id):
            res.node_provider_id = Primitive.to_proto(resource.node_provider_id)
        if Primitive.to_proto(resource.node_count):
            res.node_count = Primitive.to_proto(resource.node_count)
        if Primitive.to_proto(resource.vcpu_count):
            res.vcpu_count = Primitive.to_proto(resource.vcpu_count)
        if Primitive.to_proto(resource.memory_mb):
            res.memory_mb = Primitive.to_proto(resource.memory_mb)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipEndpointKubernetesMetadata(
            kubernetes_api_server_version=Primitive.from_proto(
                resource.kubernetes_api_server_version
            ),
            node_provider_id=Primitive.from_proto(resource.node_provider_id),
            node_count=Primitive.from_proto(resource.node_count),
            vcpu_count=Primitive.from_proto(resource.vcpu_count),
            memory_mb=Primitive.from_proto(resource.memory_mb),
            update_time=Primitive.from_proto(resource.update_time),
        )


class MembershipEndpointKubernetesMetadataArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipEndpointKubernetesMetadata.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipEndpointKubernetesMetadata.from_proto(i) for i in resources]


class MembershipEndpointKubernetesResource(object):
    def __init__(
        self,
        membership_cr_manifest: str = None,
        membership_resources: list = None,
        connect_resources: list = None,
        resource_options: dict = None,
    ):
        self.membership_cr_manifest = membership_cr_manifest
        self.membership_resources = membership_resources
        self.connect_resources = connect_resources
        self.resource_options = resource_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.GkehubBetaMembershipEndpointKubernetesResource()
        if Primitive.to_proto(resource.membership_cr_manifest):
            res.membership_cr_manifest = Primitive.to_proto(
                resource.membership_cr_manifest
            )
        if MembershipEndpointKubernetesResourceMembershipResourcesArray.to_proto(
            resource.membership_resources
        ):
            res.membership_resources.extend(
                MembershipEndpointKubernetesResourceMembershipResourcesArray.to_proto(
                    resource.membership_resources
                )
            )
        if MembershipEndpointKubernetesResourceConnectResourcesArray.to_proto(
            resource.connect_resources
        ):
            res.connect_resources.extend(
                MembershipEndpointKubernetesResourceConnectResourcesArray.to_proto(
                    resource.connect_resources
                )
            )
        if MembershipEndpointKubernetesResourceResourceOptions.to_proto(
            resource.resource_options
        ):
            res.resource_options.CopyFrom(
                MembershipEndpointKubernetesResourceResourceOptions.to_proto(
                    resource.resource_options
                )
            )
        else:
            res.ClearField("resource_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipEndpointKubernetesResource(
            membership_cr_manifest=Primitive.from_proto(
                resource.membership_cr_manifest
            ),
            membership_resources=MembershipEndpointKubernetesResourceMembershipResourcesArray.from_proto(
                resource.membership_resources
            ),
            connect_resources=MembershipEndpointKubernetesResourceConnectResourcesArray.from_proto(
                resource.connect_resources
            ),
            resource_options=MembershipEndpointKubernetesResourceResourceOptions.from_proto(
                resource.resource_options
            ),
        )


class MembershipEndpointKubernetesResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipEndpointKubernetesResource.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipEndpointKubernetesResource.from_proto(i) for i in resources]


class MembershipEndpointKubernetesResourceMembershipResources(object):
    def __init__(self, manifest: str = None, cluster_scoped: bool = None):
        self.manifest = manifest
        self.cluster_scoped = cluster_scoped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            membership_pb2.GkehubBetaMembershipEndpointKubernetesResourceMembershipResources()
        )
        if Primitive.to_proto(resource.manifest):
            res.manifest = Primitive.to_proto(resource.manifest)
        if Primitive.to_proto(resource.cluster_scoped):
            res.cluster_scoped = Primitive.to_proto(resource.cluster_scoped)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipEndpointKubernetesResourceMembershipResources(
            manifest=Primitive.from_proto(resource.manifest),
            cluster_scoped=Primitive.from_proto(resource.cluster_scoped),
        )


class MembershipEndpointKubernetesResourceMembershipResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            MembershipEndpointKubernetesResourceMembershipResources.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            MembershipEndpointKubernetesResourceMembershipResources.from_proto(i)
            for i in resources
        ]


class MembershipEndpointKubernetesResourceConnectResources(object):
    def __init__(self, manifest: str = None, cluster_scoped: bool = None):
        self.manifest = manifest
        self.cluster_scoped = cluster_scoped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            membership_pb2.GkehubBetaMembershipEndpointKubernetesResourceConnectResources()
        )
        if Primitive.to_proto(resource.manifest):
            res.manifest = Primitive.to_proto(resource.manifest)
        if Primitive.to_proto(resource.cluster_scoped):
            res.cluster_scoped = Primitive.to_proto(resource.cluster_scoped)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipEndpointKubernetesResourceConnectResources(
            manifest=Primitive.from_proto(resource.manifest),
            cluster_scoped=Primitive.from_proto(resource.cluster_scoped),
        )


class MembershipEndpointKubernetesResourceConnectResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            MembershipEndpointKubernetesResourceConnectResources.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            MembershipEndpointKubernetesResourceConnectResources.from_proto(i)
            for i in resources
        ]


class MembershipEndpointKubernetesResourceResourceOptions(object):
    def __init__(self, connect_version: str = None, v1beta1_crd: bool = None):
        self.connect_version = connect_version
        self.v1beta1_crd = v1beta1_crd

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            membership_pb2.GkehubBetaMembershipEndpointKubernetesResourceResourceOptions()
        )
        if Primitive.to_proto(resource.connect_version):
            res.connect_version = Primitive.to_proto(resource.connect_version)
        if Primitive.to_proto(resource.v1beta1_crd):
            res.v1beta1_crd = Primitive.to_proto(resource.v1beta1_crd)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipEndpointKubernetesResourceResourceOptions(
            connect_version=Primitive.from_proto(resource.connect_version),
            v1beta1_crd=Primitive.from_proto(resource.v1beta1_crd),
        )


class MembershipEndpointKubernetesResourceResourceOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            MembershipEndpointKubernetesResourceResourceOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            MembershipEndpointKubernetesResourceResourceOptions.from_proto(i)
            for i in resources
        ]


class MembershipState(object):
    def __init__(self, code: str = None):
        self.code = code

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.GkehubBetaMembershipState()
        if MembershipStateCodeEnum.to_proto(resource.code):
            res.code = MembershipStateCodeEnum.to_proto(resource.code)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipState(
            code=MembershipStateCodeEnum.from_proto(resource.code),
        )


class MembershipStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipState.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipState.from_proto(i) for i in resources]


class MembershipAuthority(object):
    def __init__(
        self,
        issuer: str = None,
        workload_identity_pool: str = None,
        identity_provider: str = None,
    ):
        self.issuer = issuer
        self.workload_identity_pool = workload_identity_pool
        self.identity_provider = identity_provider

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.GkehubBetaMembershipAuthority()
        if Primitive.to_proto(resource.issuer):
            res.issuer = Primitive.to_proto(resource.issuer)
        if Primitive.to_proto(resource.workload_identity_pool):
            res.workload_identity_pool = Primitive.to_proto(
                resource.workload_identity_pool
            )
        if Primitive.to_proto(resource.identity_provider):
            res.identity_provider = Primitive.to_proto(resource.identity_provider)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipAuthority(
            issuer=Primitive.from_proto(resource.issuer),
            workload_identity_pool=Primitive.from_proto(
                resource.workload_identity_pool
            ),
            identity_provider=Primitive.from_proto(resource.identity_provider),
        )


class MembershipAuthorityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipAuthority.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipAuthority.from_proto(i) for i in resources]


class MembershipStateCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return membership_pb2.GkehubBetaMembershipStateCodeEnum.Value(
            "GkehubBetaMembershipStateCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return membership_pb2.GkehubBetaMembershipStateCodeEnum.Name(resource)[
            len("GkehubBetaMembershipStateCodeEnum") :
        ]


class MembershipInfrastructureTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return membership_pb2.GkehubBetaMembershipInfrastructureTypeEnum.Value(
            "GkehubBetaMembershipInfrastructureTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return membership_pb2.GkehubBetaMembershipInfrastructureTypeEnum.Name(resource)[
            len("GkehubBetaMembershipInfrastructureTypeEnum") :
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
