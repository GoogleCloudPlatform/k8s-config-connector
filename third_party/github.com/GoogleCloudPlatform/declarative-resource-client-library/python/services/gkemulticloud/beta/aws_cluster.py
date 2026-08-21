# Copyright 2021 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.gkemulticloud import aws_cluster_pb2
from google3.cloud.graphite.mmv2.services.google.gkemulticloud import (
    aws_cluster_pb2_grpc,
)

from typing import List


class AwsCluster(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        networking: dict = None,
        aws_region: str = None,
        control_plane: dict = None,
        authorization: dict = None,
        state: str = None,
        endpoint: str = None,
        uid: str = None,
        reconciling: bool = None,
        create_time: str = None,
        update_time: str = None,
        etag: str = None,
        annotations: dict = None,
        workload_identity_config: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.networking = networking
        self.aws_region = aws_region
        self.control_plane = control_plane
        self.authorization = authorization
        self.annotations = annotations
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = aws_cluster_pb2_grpc.GkemulticloudBetaAwsClusterServiceStub(
            channel.Channel()
        )
        request = aws_cluster_pb2.ApplyGkemulticloudBetaAwsClusterRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if AwsClusterNetworking.to_proto(self.networking):
            request.resource.networking.CopyFrom(
                AwsClusterNetworking.to_proto(self.networking)
            )
        else:
            request.resource.ClearField("networking")
        if Primitive.to_proto(self.aws_region):
            request.resource.aws_region = Primitive.to_proto(self.aws_region)

        if AwsClusterControlPlane.to_proto(self.control_plane):
            request.resource.control_plane.CopyFrom(
                AwsClusterControlPlane.to_proto(self.control_plane)
            )
        else:
            request.resource.ClearField("control_plane")
        if AwsClusterAuthorization.to_proto(self.authorization):
            request.resource.authorization.CopyFrom(
                AwsClusterAuthorization.to_proto(self.authorization)
            )
        else:
            request.resource.ClearField("authorization")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyGkemulticloudBetaAwsCluster(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.networking = AwsClusterNetworking.from_proto(response.networking)
        self.aws_region = Primitive.from_proto(response.aws_region)
        self.control_plane = AwsClusterControlPlane.from_proto(response.control_plane)
        self.authorization = AwsClusterAuthorization.from_proto(response.authorization)
        self.state = AwsClusterStateEnum.from_proto(response.state)
        self.endpoint = Primitive.from_proto(response.endpoint)
        self.uid = Primitive.from_proto(response.uid)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.etag = Primitive.from_proto(response.etag)
        self.annotations = Primitive.from_proto(response.annotations)
        self.workload_identity_config = AwsClusterWorkloadIdentityConfig.from_proto(
            response.workload_identity_config
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = aws_cluster_pb2_grpc.GkemulticloudBetaAwsClusterServiceStub(
            channel.Channel()
        )
        request = aws_cluster_pb2.DeleteGkemulticloudBetaAwsClusterRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if AwsClusterNetworking.to_proto(self.networking):
            request.resource.networking.CopyFrom(
                AwsClusterNetworking.to_proto(self.networking)
            )
        else:
            request.resource.ClearField("networking")
        if Primitive.to_proto(self.aws_region):
            request.resource.aws_region = Primitive.to_proto(self.aws_region)

        if AwsClusterControlPlane.to_proto(self.control_plane):
            request.resource.control_plane.CopyFrom(
                AwsClusterControlPlane.to_proto(self.control_plane)
            )
        else:
            request.resource.ClearField("control_plane")
        if AwsClusterAuthorization.to_proto(self.authorization):
            request.resource.authorization.CopyFrom(
                AwsClusterAuthorization.to_proto(self.authorization)
            )
        else:
            request.resource.ClearField("authorization")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteGkemulticloudBetaAwsCluster(request)

    def list(self):
        stub = aws_cluster_pb2_grpc.GkemulticloudBetaAwsClusterServiceStub(
            channel.Channel()
        )
        request = aws_cluster_pb2.ListGkemulticloudBetaAwsClusterRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if AwsClusterNetworking.to_proto(self.networking):
            request.resource.networking.CopyFrom(
                AwsClusterNetworking.to_proto(self.networking)
            )
        else:
            request.resource.ClearField("networking")
        if Primitive.to_proto(self.aws_region):
            request.resource.aws_region = Primitive.to_proto(self.aws_region)

        if AwsClusterControlPlane.to_proto(self.control_plane):
            request.resource.control_plane.CopyFrom(
                AwsClusterControlPlane.to_proto(self.control_plane)
            )
        else:
            request.resource.ClearField("control_plane")
        if AwsClusterAuthorization.to_proto(self.authorization):
            request.resource.authorization.CopyFrom(
                AwsClusterAuthorization.to_proto(self.authorization)
            )
        else:
            request.resource.ClearField("authorization")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        return stub.ListGkemulticloudBetaAwsCluster(request).items

    def to_proto(self):
        resource = aws_cluster_pb2.GkemulticloudBetaAwsCluster()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if AwsClusterNetworking.to_proto(self.networking):
            resource.networking.CopyFrom(AwsClusterNetworking.to_proto(self.networking))
        else:
            resource.ClearField("networking")
        if Primitive.to_proto(self.aws_region):
            resource.aws_region = Primitive.to_proto(self.aws_region)
        if AwsClusterControlPlane.to_proto(self.control_plane):
            resource.control_plane.CopyFrom(
                AwsClusterControlPlane.to_proto(self.control_plane)
            )
        else:
            resource.ClearField("control_plane")
        if AwsClusterAuthorization.to_proto(self.authorization):
            resource.authorization.CopyFrom(
                AwsClusterAuthorization.to_proto(self.authorization)
            )
        else:
            resource.ClearField("authorization")
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class AwsClusterNetworking(object):
    def __init__(
        self,
        vpc_id: str = None,
        pod_address_cidr_blocks: list = None,
        service_address_cidr_blocks: list = None,
        service_load_balancer_subnet_ids: list = None,
    ):
        self.vpc_id = vpc_id
        self.pod_address_cidr_blocks = pod_address_cidr_blocks
        self.service_address_cidr_blocks = service_address_cidr_blocks
        self.service_load_balancer_subnet_ids = service_load_balancer_subnet_ids

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_cluster_pb2.GkemulticloudBetaAwsClusterNetworking()
        if Primitive.to_proto(resource.vpc_id):
            res.vpc_id = Primitive.to_proto(resource.vpc_id)
        if Primitive.to_proto(resource.pod_address_cidr_blocks):
            res.pod_address_cidr_blocks.extend(
                Primitive.to_proto(resource.pod_address_cidr_blocks)
            )
        if Primitive.to_proto(resource.service_address_cidr_blocks):
            res.service_address_cidr_blocks.extend(
                Primitive.to_proto(resource.service_address_cidr_blocks)
            )
        if Primitive.to_proto(resource.service_load_balancer_subnet_ids):
            res.service_load_balancer_subnet_ids.extend(
                Primitive.to_proto(resource.service_load_balancer_subnet_ids)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsClusterNetworking(
            vpc_id=Primitive.from_proto(resource.vpc_id),
            pod_address_cidr_blocks=Primitive.from_proto(
                resource.pod_address_cidr_blocks
            ),
            service_address_cidr_blocks=Primitive.from_proto(
                resource.service_address_cidr_blocks
            ),
            service_load_balancer_subnet_ids=Primitive.from_proto(
                resource.service_load_balancer_subnet_ids
            ),
        )


class AwsClusterNetworkingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsClusterNetworking.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsClusterNetworking.from_proto(i) for i in resources]


class AwsClusterControlPlane(object):
    def __init__(
        self,
        version: str = None,
        instance_type: str = None,
        ssh_config: dict = None,
        subnet_ids: list = None,
        security_group_ids: list = None,
        iam_instance_profile: str = None,
        root_volume: dict = None,
        main_volume: dict = None,
        database_encryption: dict = None,
        tags: dict = None,
        aws_services_authentication: dict = None,
    ):
        self.version = version
        self.instance_type = instance_type
        self.ssh_config = ssh_config
        self.subnet_ids = subnet_ids
        self.security_group_ids = security_group_ids
        self.iam_instance_profile = iam_instance_profile
        self.root_volume = root_volume
        self.main_volume = main_volume
        self.database_encryption = database_encryption
        self.tags = tags
        self.aws_services_authentication = aws_services_authentication

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_cluster_pb2.GkemulticloudBetaAwsClusterControlPlane()
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        if Primitive.to_proto(resource.instance_type):
            res.instance_type = Primitive.to_proto(resource.instance_type)
        if AwsClusterControlPlaneSshConfig.to_proto(resource.ssh_config):
            res.ssh_config.CopyFrom(
                AwsClusterControlPlaneSshConfig.to_proto(resource.ssh_config)
            )
        else:
            res.ClearField("ssh_config")
        if Primitive.to_proto(resource.subnet_ids):
            res.subnet_ids.extend(Primitive.to_proto(resource.subnet_ids))
        if Primitive.to_proto(resource.security_group_ids):
            res.security_group_ids.extend(
                Primitive.to_proto(resource.security_group_ids)
            )
        if Primitive.to_proto(resource.iam_instance_profile):
            res.iam_instance_profile = Primitive.to_proto(resource.iam_instance_profile)
        if AwsClusterControlPlaneRootVolume.to_proto(resource.root_volume):
            res.root_volume.CopyFrom(
                AwsClusterControlPlaneRootVolume.to_proto(resource.root_volume)
            )
        else:
            res.ClearField("root_volume")
        if AwsClusterControlPlaneMainVolume.to_proto(resource.main_volume):
            res.main_volume.CopyFrom(
                AwsClusterControlPlaneMainVolume.to_proto(resource.main_volume)
            )
        else:
            res.ClearField("main_volume")
        if AwsClusterControlPlaneDatabaseEncryption.to_proto(
            resource.database_encryption
        ):
            res.database_encryption.CopyFrom(
                AwsClusterControlPlaneDatabaseEncryption.to_proto(
                    resource.database_encryption
                )
            )
        else:
            res.ClearField("database_encryption")
        if Primitive.to_proto(resource.tags):
            res.tags = Primitive.to_proto(resource.tags)
        if AwsClusterControlPlaneAwsServicesAuthentication.to_proto(
            resource.aws_services_authentication
        ):
            res.aws_services_authentication.CopyFrom(
                AwsClusterControlPlaneAwsServicesAuthentication.to_proto(
                    resource.aws_services_authentication
                )
            )
        else:
            res.ClearField("aws_services_authentication")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsClusterControlPlane(
            version=Primitive.from_proto(resource.version),
            instance_type=Primitive.from_proto(resource.instance_type),
            ssh_config=AwsClusterControlPlaneSshConfig.from_proto(resource.ssh_config),
            subnet_ids=Primitive.from_proto(resource.subnet_ids),
            security_group_ids=Primitive.from_proto(resource.security_group_ids),
            iam_instance_profile=Primitive.from_proto(resource.iam_instance_profile),
            root_volume=AwsClusterControlPlaneRootVolume.from_proto(
                resource.root_volume
            ),
            main_volume=AwsClusterControlPlaneMainVolume.from_proto(
                resource.main_volume
            ),
            database_encryption=AwsClusterControlPlaneDatabaseEncryption.from_proto(
                resource.database_encryption
            ),
            tags=Primitive.from_proto(resource.tags),
            aws_services_authentication=AwsClusterControlPlaneAwsServicesAuthentication.from_proto(
                resource.aws_services_authentication
            ),
        )


class AwsClusterControlPlaneArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsClusterControlPlane.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsClusterControlPlane.from_proto(i) for i in resources]


class AwsClusterControlPlaneSshConfig(object):
    def __init__(self, ec2_key_pair: str = None):
        self.ec2_key_pair = ec2_key_pair

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_cluster_pb2.GkemulticloudBetaAwsClusterControlPlaneSshConfig()
        if Primitive.to_proto(resource.ec2_key_pair):
            res.ec2_key_pair = Primitive.to_proto(resource.ec2_key_pair)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsClusterControlPlaneSshConfig(
            ec2_key_pair=Primitive.from_proto(resource.ec2_key_pair),
        )


class AwsClusterControlPlaneSshConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsClusterControlPlaneSshConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsClusterControlPlaneSshConfig.from_proto(i) for i in resources]


class AwsClusterControlPlaneRootVolume(object):
    def __init__(
        self,
        size_gib: int = None,
        volume_type: str = None,
        iops: int = None,
        kms_key_arn: str = None,
    ):
        self.size_gib = size_gib
        self.volume_type = volume_type
        self.iops = iops
        self.kms_key_arn = kms_key_arn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_cluster_pb2.GkemulticloudBetaAwsClusterControlPlaneRootVolume()
        if Primitive.to_proto(resource.size_gib):
            res.size_gib = Primitive.to_proto(resource.size_gib)
        if AwsClusterControlPlaneRootVolumeVolumeTypeEnum.to_proto(
            resource.volume_type
        ):
            res.volume_type = AwsClusterControlPlaneRootVolumeVolumeTypeEnum.to_proto(
                resource.volume_type
            )
        if Primitive.to_proto(resource.iops):
            res.iops = Primitive.to_proto(resource.iops)
        if Primitive.to_proto(resource.kms_key_arn):
            res.kms_key_arn = Primitive.to_proto(resource.kms_key_arn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsClusterControlPlaneRootVolume(
            size_gib=Primitive.from_proto(resource.size_gib),
            volume_type=AwsClusterControlPlaneRootVolumeVolumeTypeEnum.from_proto(
                resource.volume_type
            ),
            iops=Primitive.from_proto(resource.iops),
            kms_key_arn=Primitive.from_proto(resource.kms_key_arn),
        )


class AwsClusterControlPlaneRootVolumeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsClusterControlPlaneRootVolume.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsClusterControlPlaneRootVolume.from_proto(i) for i in resources]


class AwsClusterControlPlaneMainVolume(object):
    def __init__(
        self,
        size_gib: int = None,
        volume_type: str = None,
        iops: int = None,
        kms_key_arn: str = None,
    ):
        self.size_gib = size_gib
        self.volume_type = volume_type
        self.iops = iops
        self.kms_key_arn = kms_key_arn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_cluster_pb2.GkemulticloudBetaAwsClusterControlPlaneMainVolume()
        if Primitive.to_proto(resource.size_gib):
            res.size_gib = Primitive.to_proto(resource.size_gib)
        if AwsClusterControlPlaneMainVolumeVolumeTypeEnum.to_proto(
            resource.volume_type
        ):
            res.volume_type = AwsClusterControlPlaneMainVolumeVolumeTypeEnum.to_proto(
                resource.volume_type
            )
        if Primitive.to_proto(resource.iops):
            res.iops = Primitive.to_proto(resource.iops)
        if Primitive.to_proto(resource.kms_key_arn):
            res.kms_key_arn = Primitive.to_proto(resource.kms_key_arn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsClusterControlPlaneMainVolume(
            size_gib=Primitive.from_proto(resource.size_gib),
            volume_type=AwsClusterControlPlaneMainVolumeVolumeTypeEnum.from_proto(
                resource.volume_type
            ),
            iops=Primitive.from_proto(resource.iops),
            kms_key_arn=Primitive.from_proto(resource.kms_key_arn),
        )


class AwsClusterControlPlaneMainVolumeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsClusterControlPlaneMainVolume.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsClusterControlPlaneMainVolume.from_proto(i) for i in resources]


class AwsClusterControlPlaneDatabaseEncryption(object):
    def __init__(self, kms_key_arn: str = None):
        self.kms_key_arn = kms_key_arn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            aws_cluster_pb2.GkemulticloudBetaAwsClusterControlPlaneDatabaseEncryption()
        )
        if Primitive.to_proto(resource.kms_key_arn):
            res.kms_key_arn = Primitive.to_proto(resource.kms_key_arn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsClusterControlPlaneDatabaseEncryption(
            kms_key_arn=Primitive.from_proto(resource.kms_key_arn),
        )


class AwsClusterControlPlaneDatabaseEncryptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsClusterControlPlaneDatabaseEncryption.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            AwsClusterControlPlaneDatabaseEncryption.from_proto(i) for i in resources
        ]


class AwsClusterControlPlaneAwsServicesAuthentication(object):
    def __init__(self, role_arn: str = None, role_session_name: str = None):
        self.role_arn = role_arn
        self.role_session_name = role_session_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            aws_cluster_pb2.GkemulticloudBetaAwsClusterControlPlaneAwsServicesAuthentication()
        )
        if Primitive.to_proto(resource.role_arn):
            res.role_arn = Primitive.to_proto(resource.role_arn)
        if Primitive.to_proto(resource.role_session_name):
            res.role_session_name = Primitive.to_proto(resource.role_session_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsClusterControlPlaneAwsServicesAuthentication(
            role_arn=Primitive.from_proto(resource.role_arn),
            role_session_name=Primitive.from_proto(resource.role_session_name),
        )


class AwsClusterControlPlaneAwsServicesAuthenticationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AwsClusterControlPlaneAwsServicesAuthentication.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AwsClusterControlPlaneAwsServicesAuthentication.from_proto(i)
            for i in resources
        ]


class AwsClusterAuthorization(object):
    def __init__(self, admin_users: list = None):
        self.admin_users = admin_users

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_cluster_pb2.GkemulticloudBetaAwsClusterAuthorization()
        if AwsClusterAuthorizationAdminUsersArray.to_proto(resource.admin_users):
            res.admin_users.extend(
                AwsClusterAuthorizationAdminUsersArray.to_proto(resource.admin_users)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsClusterAuthorization(
            admin_users=AwsClusterAuthorizationAdminUsersArray.from_proto(
                resource.admin_users
            ),
        )


class AwsClusterAuthorizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsClusterAuthorization.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsClusterAuthorization.from_proto(i) for i in resources]


class AwsClusterAuthorizationAdminUsers(object):
    def __init__(self, username: str = None):
        self.username = username

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_cluster_pb2.GkemulticloudBetaAwsClusterAuthorizationAdminUsers()
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsClusterAuthorizationAdminUsers(
            username=Primitive.from_proto(resource.username),
        )


class AwsClusterAuthorizationAdminUsersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsClusterAuthorizationAdminUsers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsClusterAuthorizationAdminUsers.from_proto(i) for i in resources]


class AwsClusterWorkloadIdentityConfig(object):
    def __init__(
        self,
        issuer_uri: str = None,
        workload_pool: str = None,
        identity_provider: str = None,
    ):
        self.issuer_uri = issuer_uri
        self.workload_pool = workload_pool
        self.identity_provider = identity_provider

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = aws_cluster_pb2.GkemulticloudBetaAwsClusterWorkloadIdentityConfig()
        if Primitive.to_proto(resource.issuer_uri):
            res.issuer_uri = Primitive.to_proto(resource.issuer_uri)
        if Primitive.to_proto(resource.workload_pool):
            res.workload_pool = Primitive.to_proto(resource.workload_pool)
        if Primitive.to_proto(resource.identity_provider):
            res.identity_provider = Primitive.to_proto(resource.identity_provider)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AwsClusterWorkloadIdentityConfig(
            issuer_uri=Primitive.from_proto(resource.issuer_uri),
            workload_pool=Primitive.from_proto(resource.workload_pool),
            identity_provider=Primitive.from_proto(resource.identity_provider),
        )


class AwsClusterWorkloadIdentityConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AwsClusterWorkloadIdentityConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AwsClusterWorkloadIdentityConfig.from_proto(i) for i in resources]


class AwsClusterControlPlaneRootVolumeVolumeTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return aws_cluster_pb2.GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum.Value(
            "GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return aws_cluster_pb2.GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum.Name(
            resource
        )[
            len("GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum") :
        ]


class AwsClusterControlPlaneMainVolumeVolumeTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return aws_cluster_pb2.GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum.Value(
            "GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return aws_cluster_pb2.GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum.Name(
            resource
        )[
            len("GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum") :
        ]


class AwsClusterStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return aws_cluster_pb2.GkemulticloudBetaAwsClusterStateEnum.Value(
            "GkemulticloudBetaAwsClusterStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return aws_cluster_pb2.GkemulticloudBetaAwsClusterStateEnum.Name(resource)[
            len("GkemulticloudBetaAwsClusterStateEnum") :
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
