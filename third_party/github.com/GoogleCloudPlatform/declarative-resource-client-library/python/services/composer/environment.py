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
from google3.cloud.graphite.mmv2.services.google.composer import environment_pb2
from google3.cloud.graphite.mmv2.services.google.composer import environment_pb2_grpc

from typing import List


class Environment(object):
    def __init__(
        self,
        name: str = None,
        config: dict = None,
        uuid: str = None,
        state: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.config = config
        self.labels = labels
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = environment_pb2_grpc.ComposerEnvironmentServiceStub(channel.Channel())
        request = environment_pb2.ApplyComposerEnvironmentRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if EnvironmentConfig.to_proto(self.config):
            request.resource.config.CopyFrom(EnvironmentConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComposerEnvironment(request)
        self.name = Primitive.from_proto(response.name)
        self.config = EnvironmentConfig.from_proto(response.config)
        self.uuid = Primitive.from_proto(response.uuid)
        self.state = EnvironmentStateEnum.from_proto(response.state)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = environment_pb2_grpc.ComposerEnvironmentServiceStub(channel.Channel())
        request = environment_pb2.DeleteComposerEnvironmentRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if EnvironmentConfig.to_proto(self.config):
            request.resource.config.CopyFrom(EnvironmentConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteComposerEnvironment(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = environment_pb2_grpc.ComposerEnvironmentServiceStub(channel.Channel())
        request = environment_pb2.ListComposerEnvironmentRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComposerEnvironment(request).items

    def to_proto(self):
        resource = environment_pb2.ComposerEnvironment()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if EnvironmentConfig.to_proto(self.config):
            resource.config.CopyFrom(EnvironmentConfig.to_proto(self.config))
        else:
            resource.ClearField("config")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class EnvironmentConfig(object):
    def __init__(
        self,
        gke_cluster: str = None,
        dag_gcs_prefix: str = None,
        node_count: int = None,
        software_config: dict = None,
        node_config: dict = None,
        private_environment_config: dict = None,
        web_server_network_access_control: dict = None,
        database_config: dict = None,
        web_server_config: dict = None,
        encryption_config: dict = None,
        airflow_uri: str = None,
    ):
        self.gke_cluster = gke_cluster
        self.dag_gcs_prefix = dag_gcs_prefix
        self.node_count = node_count
        self.software_config = software_config
        self.node_config = node_config
        self.private_environment_config = private_environment_config
        self.web_server_network_access_control = web_server_network_access_control
        self.database_config = database_config
        self.web_server_config = web_server_config
        self.encryption_config = encryption_config
        self.airflow_uri = airflow_uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = environment_pb2.ComposerEnvironmentConfig()
        if Primitive.to_proto(resource.gke_cluster):
            res.gke_cluster = Primitive.to_proto(resource.gke_cluster)
        if Primitive.to_proto(resource.dag_gcs_prefix):
            res.dag_gcs_prefix = Primitive.to_proto(resource.dag_gcs_prefix)
        if Primitive.to_proto(resource.node_count):
            res.node_count = Primitive.to_proto(resource.node_count)
        if EnvironmentConfigSoftwareConfig.to_proto(resource.software_config):
            res.software_config.CopyFrom(
                EnvironmentConfigSoftwareConfig.to_proto(resource.software_config)
            )
        else:
            res.ClearField("software_config")
        if EnvironmentConfigNodeConfig.to_proto(resource.node_config):
            res.node_config.CopyFrom(
                EnvironmentConfigNodeConfig.to_proto(resource.node_config)
            )
        else:
            res.ClearField("node_config")
        if EnvironmentConfigPrivateEnvironmentConfig.to_proto(
            resource.private_environment_config
        ):
            res.private_environment_config.CopyFrom(
                EnvironmentConfigPrivateEnvironmentConfig.to_proto(
                    resource.private_environment_config
                )
            )
        else:
            res.ClearField("private_environment_config")
        if EnvironmentConfigWebServerNetworkAccessControl.to_proto(
            resource.web_server_network_access_control
        ):
            res.web_server_network_access_control.CopyFrom(
                EnvironmentConfigWebServerNetworkAccessControl.to_proto(
                    resource.web_server_network_access_control
                )
            )
        else:
            res.ClearField("web_server_network_access_control")
        if EnvironmentConfigDatabaseConfig.to_proto(resource.database_config):
            res.database_config.CopyFrom(
                EnvironmentConfigDatabaseConfig.to_proto(resource.database_config)
            )
        else:
            res.ClearField("database_config")
        if EnvironmentConfigWebServerConfig.to_proto(resource.web_server_config):
            res.web_server_config.CopyFrom(
                EnvironmentConfigWebServerConfig.to_proto(resource.web_server_config)
            )
        else:
            res.ClearField("web_server_config")
        if EnvironmentConfigEncryptionConfig.to_proto(resource.encryption_config):
            res.encryption_config.CopyFrom(
                EnvironmentConfigEncryptionConfig.to_proto(resource.encryption_config)
            )
        else:
            res.ClearField("encryption_config")
        if Primitive.to_proto(resource.airflow_uri):
            res.airflow_uri = Primitive.to_proto(resource.airflow_uri)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentConfig(
            gke_cluster=Primitive.from_proto(resource.gke_cluster),
            dag_gcs_prefix=Primitive.from_proto(resource.dag_gcs_prefix),
            node_count=Primitive.from_proto(resource.node_count),
            software_config=EnvironmentConfigSoftwareConfig.from_proto(
                resource.software_config
            ),
            node_config=EnvironmentConfigNodeConfig.from_proto(resource.node_config),
            private_environment_config=EnvironmentConfigPrivateEnvironmentConfig.from_proto(
                resource.private_environment_config
            ),
            web_server_network_access_control=EnvironmentConfigWebServerNetworkAccessControl.from_proto(
                resource.web_server_network_access_control
            ),
            database_config=EnvironmentConfigDatabaseConfig.from_proto(
                resource.database_config
            ),
            web_server_config=EnvironmentConfigWebServerConfig.from_proto(
                resource.web_server_config
            ),
            encryption_config=EnvironmentConfigEncryptionConfig.from_proto(
                resource.encryption_config
            ),
            airflow_uri=Primitive.from_proto(resource.airflow_uri),
        )


class EnvironmentConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EnvironmentConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EnvironmentConfig.from_proto(i) for i in resources]


class EnvironmentConfigSoftwareConfig(object):
    def __init__(
        self,
        image_version: str = None,
        airflow_config_overrides: dict = None,
        pypi_packages: dict = None,
        env_variables: dict = None,
        python_version: str = None,
    ):
        self.image_version = image_version
        self.airflow_config_overrides = airflow_config_overrides
        self.pypi_packages = pypi_packages
        self.env_variables = env_variables
        self.python_version = python_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = environment_pb2.ComposerEnvironmentConfigSoftwareConfig()
        if Primitive.to_proto(resource.image_version):
            res.image_version = Primitive.to_proto(resource.image_version)
        if Primitive.to_proto(resource.airflow_config_overrides):
            res.airflow_config_overrides = Primitive.to_proto(
                resource.airflow_config_overrides
            )
        if Primitive.to_proto(resource.pypi_packages):
            res.pypi_packages = Primitive.to_proto(resource.pypi_packages)
        if Primitive.to_proto(resource.env_variables):
            res.env_variables = Primitive.to_proto(resource.env_variables)
        if Primitive.to_proto(resource.python_version):
            res.python_version = Primitive.to_proto(resource.python_version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentConfigSoftwareConfig(
            image_version=Primitive.from_proto(resource.image_version),
            airflow_config_overrides=Primitive.from_proto(
                resource.airflow_config_overrides
            ),
            pypi_packages=Primitive.from_proto(resource.pypi_packages),
            env_variables=Primitive.from_proto(resource.env_variables),
            python_version=Primitive.from_proto(resource.python_version),
        )


class EnvironmentConfigSoftwareConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EnvironmentConfigSoftwareConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EnvironmentConfigSoftwareConfig.from_proto(i) for i in resources]


class EnvironmentConfigNodeConfig(object):
    def __init__(
        self,
        location: str = None,
        machine_type: str = None,
        network: str = None,
        subnetwork: str = None,
        disk_size_gb: int = None,
        oauth_scopes: list = None,
        service_account: str = None,
        tags: list = None,
        ip_allocation_policy: dict = None,
    ):
        self.location = location
        self.machine_type = machine_type
        self.network = network
        self.subnetwork = subnetwork
        self.disk_size_gb = disk_size_gb
        self.oauth_scopes = oauth_scopes
        self.service_account = service_account
        self.tags = tags
        self.ip_allocation_policy = ip_allocation_policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = environment_pb2.ComposerEnvironmentConfigNodeConfig()
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        if Primitive.to_proto(resource.subnetwork):
            res.subnetwork = Primitive.to_proto(resource.subnetwork)
        if Primitive.to_proto(resource.disk_size_gb):
            res.disk_size_gb = Primitive.to_proto(resource.disk_size_gb)
        if Primitive.to_proto(resource.oauth_scopes):
            res.oauth_scopes.extend(Primitive.to_proto(resource.oauth_scopes))
        if Primitive.to_proto(resource.service_account):
            res.service_account = Primitive.to_proto(resource.service_account)
        if Primitive.to_proto(resource.tags):
            res.tags.extend(Primitive.to_proto(resource.tags))
        if EnvironmentConfigNodeConfigIPAllocationPolicy.to_proto(
            resource.ip_allocation_policy
        ):
            res.ip_allocation_policy.CopyFrom(
                EnvironmentConfigNodeConfigIPAllocationPolicy.to_proto(
                    resource.ip_allocation_policy
                )
            )
        else:
            res.ClearField("ip_allocation_policy")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentConfigNodeConfig(
            location=Primitive.from_proto(resource.location),
            machine_type=Primitive.from_proto(resource.machine_type),
            network=Primitive.from_proto(resource.network),
            subnetwork=Primitive.from_proto(resource.subnetwork),
            disk_size_gb=Primitive.from_proto(resource.disk_size_gb),
            oauth_scopes=Primitive.from_proto(resource.oauth_scopes),
            service_account=Primitive.from_proto(resource.service_account),
            tags=Primitive.from_proto(resource.tags),
            ip_allocation_policy=EnvironmentConfigNodeConfigIPAllocationPolicy.from_proto(
                resource.ip_allocation_policy
            ),
        )


class EnvironmentConfigNodeConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EnvironmentConfigNodeConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EnvironmentConfigNodeConfig.from_proto(i) for i in resources]


class EnvironmentConfigNodeConfigIPAllocationPolicy(object):
    def __init__(
        self,
        use_ip_aliases: bool = None,
        cluster_secondary_range_name: str = None,
        cluster_ipv4_cidr_block: str = None,
        services_secondary_range_name: str = None,
        services_ipv4_cidr_block: str = None,
    ):
        self.use_ip_aliases = use_ip_aliases
        self.cluster_secondary_range_name = cluster_secondary_range_name
        self.cluster_ipv4_cidr_block = cluster_ipv4_cidr_block
        self.services_secondary_range_name = services_secondary_range_name
        self.services_ipv4_cidr_block = services_ipv4_cidr_block

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = environment_pb2.ComposerEnvironmentConfigNodeConfigIPAllocationPolicy()
        if Primitive.to_proto(resource.use_ip_aliases):
            res.use_ip_aliases = Primitive.to_proto(resource.use_ip_aliases)
        if Primitive.to_proto(resource.cluster_secondary_range_name):
            res.cluster_secondary_range_name = Primitive.to_proto(
                resource.cluster_secondary_range_name
            )
        if Primitive.to_proto(resource.cluster_ipv4_cidr_block):
            res.cluster_ipv4_cidr_block = Primitive.to_proto(
                resource.cluster_ipv4_cidr_block
            )
        if Primitive.to_proto(resource.services_secondary_range_name):
            res.services_secondary_range_name = Primitive.to_proto(
                resource.services_secondary_range_name
            )
        if Primitive.to_proto(resource.services_ipv4_cidr_block):
            res.services_ipv4_cidr_block = Primitive.to_proto(
                resource.services_ipv4_cidr_block
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentConfigNodeConfigIPAllocationPolicy(
            use_ip_aliases=Primitive.from_proto(resource.use_ip_aliases),
            cluster_secondary_range_name=Primitive.from_proto(
                resource.cluster_secondary_range_name
            ),
            cluster_ipv4_cidr_block=Primitive.from_proto(
                resource.cluster_ipv4_cidr_block
            ),
            services_secondary_range_name=Primitive.from_proto(
                resource.services_secondary_range_name
            ),
            services_ipv4_cidr_block=Primitive.from_proto(
                resource.services_ipv4_cidr_block
            ),
        )


class EnvironmentConfigNodeConfigIPAllocationPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EnvironmentConfigNodeConfigIPAllocationPolicy.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EnvironmentConfigNodeConfigIPAllocationPolicy.from_proto(i)
            for i in resources
        ]


class EnvironmentConfigPrivateEnvironmentConfig(object):
    def __init__(
        self,
        enable_private_environment: bool = None,
        private_cluster_config: dict = None,
        web_server_ipv4_cidr_block: str = None,
        cloud_sql_ipv4_cidr_block: str = None,
        web_server_ipv4_reserved_range: str = None,
    ):
        self.enable_private_environment = enable_private_environment
        self.private_cluster_config = private_cluster_config
        self.web_server_ipv4_cidr_block = web_server_ipv4_cidr_block
        self.cloud_sql_ipv4_cidr_block = cloud_sql_ipv4_cidr_block
        self.web_server_ipv4_reserved_range = web_server_ipv4_reserved_range

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = environment_pb2.ComposerEnvironmentConfigPrivateEnvironmentConfig()
        if Primitive.to_proto(resource.enable_private_environment):
            res.enable_private_environment = Primitive.to_proto(
                resource.enable_private_environment
            )
        if EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig.to_proto(
            resource.private_cluster_config
        ):
            res.private_cluster_config.CopyFrom(
                EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig.to_proto(
                    resource.private_cluster_config
                )
            )
        else:
            res.ClearField("private_cluster_config")
        if Primitive.to_proto(resource.web_server_ipv4_cidr_block):
            res.web_server_ipv4_cidr_block = Primitive.to_proto(
                resource.web_server_ipv4_cidr_block
            )
        if Primitive.to_proto(resource.cloud_sql_ipv4_cidr_block):
            res.cloud_sql_ipv4_cidr_block = Primitive.to_proto(
                resource.cloud_sql_ipv4_cidr_block
            )
        if Primitive.to_proto(resource.web_server_ipv4_reserved_range):
            res.web_server_ipv4_reserved_range = Primitive.to_proto(
                resource.web_server_ipv4_reserved_range
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentConfigPrivateEnvironmentConfig(
            enable_private_environment=Primitive.from_proto(
                resource.enable_private_environment
            ),
            private_cluster_config=EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig.from_proto(
                resource.private_cluster_config
            ),
            web_server_ipv4_cidr_block=Primitive.from_proto(
                resource.web_server_ipv4_cidr_block
            ),
            cloud_sql_ipv4_cidr_block=Primitive.from_proto(
                resource.cloud_sql_ipv4_cidr_block
            ),
            web_server_ipv4_reserved_range=Primitive.from_proto(
                resource.web_server_ipv4_reserved_range
            ),
        )


class EnvironmentConfigPrivateEnvironmentConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EnvironmentConfigPrivateEnvironmentConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EnvironmentConfigPrivateEnvironmentConfig.from_proto(i) for i in resources
        ]


class EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(object):
    def __init__(
        self,
        enable_private_endpoint: bool = None,
        master_ipv4_cidr_block: str = None,
        master_ipv4_reserved_range: str = None,
    ):
        self.enable_private_endpoint = enable_private_endpoint
        self.master_ipv4_cidr_block = master_ipv4_cidr_block
        self.master_ipv4_reserved_range = master_ipv4_reserved_range

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            environment_pb2.ComposerEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig()
        )
        if Primitive.to_proto(resource.enable_private_endpoint):
            res.enable_private_endpoint = Primitive.to_proto(
                resource.enable_private_endpoint
            )
        if Primitive.to_proto(resource.master_ipv4_cidr_block):
            res.master_ipv4_cidr_block = Primitive.to_proto(
                resource.master_ipv4_cidr_block
            )
        if Primitive.to_proto(resource.master_ipv4_reserved_range):
            res.master_ipv4_reserved_range = Primitive.to_proto(
                resource.master_ipv4_reserved_range
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(
            enable_private_endpoint=Primitive.from_proto(
                resource.enable_private_endpoint
            ),
            master_ipv4_cidr_block=Primitive.from_proto(
                resource.master_ipv4_cidr_block
            ),
            master_ipv4_reserved_range=Primitive.from_proto(
                resource.master_ipv4_reserved_range
            ),
        )


class EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig.from_proto(i)
            for i in resources
        ]


class EnvironmentConfigWebServerNetworkAccessControl(object):
    def __init__(self, allowed_ip_ranges: list = None):
        self.allowed_ip_ranges = allowed_ip_ranges

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = environment_pb2.ComposerEnvironmentConfigWebServerNetworkAccessControl()
        if EnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesArray.to_proto(
            resource.allowed_ip_ranges
        ):
            res.allowed_ip_ranges.extend(
                EnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesArray.to_proto(
                    resource.allowed_ip_ranges
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentConfigWebServerNetworkAccessControl(
            allowed_ip_ranges=EnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesArray.from_proto(
                resource.allowed_ip_ranges
            ),
        )


class EnvironmentConfigWebServerNetworkAccessControlArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EnvironmentConfigWebServerNetworkAccessControl.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EnvironmentConfigWebServerNetworkAccessControl.from_proto(i)
            for i in resources
        ]


class EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(object):
    def __init__(self, value: str = None, description: str = None):
        self.value = value
        self.description = description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            environment_pb2.ComposerEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges()
        )
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(
            value=Primitive.from_proto(resource.value),
            description=Primitive.from_proto(resource.description),
        )


class EnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges.from_proto(i)
            for i in resources
        ]


class EnvironmentConfigDatabaseConfig(object):
    def __init__(self, machine_type: str = None):
        self.machine_type = machine_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = environment_pb2.ComposerEnvironmentConfigDatabaseConfig()
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentConfigDatabaseConfig(
            machine_type=Primitive.from_proto(resource.machine_type),
        )


class EnvironmentConfigDatabaseConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EnvironmentConfigDatabaseConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EnvironmentConfigDatabaseConfig.from_proto(i) for i in resources]


class EnvironmentConfigWebServerConfig(object):
    def __init__(self, machine_type: str = None):
        self.machine_type = machine_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = environment_pb2.ComposerEnvironmentConfigWebServerConfig()
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentConfigWebServerConfig(
            machine_type=Primitive.from_proto(resource.machine_type),
        )


class EnvironmentConfigWebServerConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EnvironmentConfigWebServerConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EnvironmentConfigWebServerConfig.from_proto(i) for i in resources]


class EnvironmentConfigEncryptionConfig(object):
    def __init__(self, kms_key_name: str = None):
        self.kms_key_name = kms_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = environment_pb2.ComposerEnvironmentConfigEncryptionConfig()
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentConfigEncryptionConfig(
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
        )


class EnvironmentConfigEncryptionConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EnvironmentConfigEncryptionConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EnvironmentConfigEncryptionConfig.from_proto(i) for i in resources]


class EnvironmentStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return environment_pb2.ComposerEnvironmentStateEnum.Value(
            "ComposerEnvironmentStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return environment_pb2.ComposerEnvironmentStateEnum.Name(resource)[
            len("ComposerEnvironmentStateEnum") :
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
