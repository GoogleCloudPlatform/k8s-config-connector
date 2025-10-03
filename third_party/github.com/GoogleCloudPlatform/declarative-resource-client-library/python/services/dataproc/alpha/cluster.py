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
from google3.cloud.graphite.mmv2.services.google.dataproc import cluster_pb2
from google3.cloud.graphite.mmv2.services.google.dataproc import cluster_pb2_grpc

from typing import List


class Cluster(object):
    def __init__(
        self,
        project: str = None,
        name: str = None,
        config: dict = None,
        labels: dict = None,
        status: dict = None,
        status_history: list = None,
        cluster_uuid: str = None,
        metrics: dict = None,
        location: str = None,
        virtual_cluster_config: dict = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.project = project
        self.name = name
        self.config = config
        self.labels = labels
        self.location = location
        self.virtual_cluster_config = virtual_cluster_config
        self.service_account_file = service_account_file

    def apply(self):
        stub = cluster_pb2_grpc.DataprocAlphaClusterServiceStub(channel.Channel())
        request = cluster_pb2.ApplyDataprocAlphaClusterRequest()
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ClusterConfig.to_proto(self.config):
            request.resource.config.CopyFrom(ClusterConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if ClusterVirtualClusterConfig.to_proto(self.virtual_cluster_config):
            request.resource.virtual_cluster_config.CopyFrom(
                ClusterVirtualClusterConfig.to_proto(self.virtual_cluster_config)
            )
        else:
            request.resource.ClearField("virtual_cluster_config")
        request.service_account_file = self.service_account_file

        response = stub.ApplyDataprocAlphaCluster(request)
        self.project = Primitive.from_proto(response.project)
        self.name = Primitive.from_proto(response.name)
        self.config = ClusterConfig.from_proto(response.config)
        self.labels = Primitive.from_proto(response.labels)
        self.status = ClusterStatus.from_proto(response.status)
        self.status_history = ClusterStatusHistoryArray.from_proto(
            response.status_history
        )
        self.cluster_uuid = Primitive.from_proto(response.cluster_uuid)
        self.metrics = ClusterMetrics.from_proto(response.metrics)
        self.location = Primitive.from_proto(response.location)
        self.virtual_cluster_config = ClusterVirtualClusterConfig.from_proto(
            response.virtual_cluster_config
        )

    def delete(self):
        stub = cluster_pb2_grpc.DataprocAlphaClusterServiceStub(channel.Channel())
        request = cluster_pb2.DeleteDataprocAlphaClusterRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ClusterConfig.to_proto(self.config):
            request.resource.config.CopyFrom(ClusterConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if ClusterVirtualClusterConfig.to_proto(self.virtual_cluster_config):
            request.resource.virtual_cluster_config.CopyFrom(
                ClusterVirtualClusterConfig.to_proto(self.virtual_cluster_config)
            )
        else:
            request.resource.ClearField("virtual_cluster_config")
        response = stub.DeleteDataprocAlphaCluster(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = cluster_pb2_grpc.DataprocAlphaClusterServiceStub(channel.Channel())
        request = cluster_pb2.ListDataprocAlphaClusterRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListDataprocAlphaCluster(request).items

    def to_proto(self):
        resource = cluster_pb2.DataprocAlphaCluster()
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if ClusterConfig.to_proto(self.config):
            resource.config.CopyFrom(ClusterConfig.to_proto(self.config))
        else:
            resource.ClearField("config")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if ClusterVirtualClusterConfig.to_proto(self.virtual_cluster_config):
            resource.virtual_cluster_config.CopyFrom(
                ClusterVirtualClusterConfig.to_proto(self.virtual_cluster_config)
            )
        else:
            resource.ClearField("virtual_cluster_config")
        return resource


class ClusterConfig(object):
    def __init__(
        self,
        staging_bucket: str = None,
        temp_bucket: str = None,
        gce_cluster_config: dict = None,
        master_config: dict = None,
        worker_config: dict = None,
        secondary_worker_config: dict = None,
        software_config: dict = None,
        initialization_actions: list = None,
        encryption_config: dict = None,
        autoscaling_config: dict = None,
        security_config: dict = None,
        lifecycle_config: dict = None,
        endpoint_config: dict = None,
        gke_cluster_config: dict = None,
        metastore_config: dict = None,
        dataproc_metric_config: dict = None,
    ):
        self.staging_bucket = staging_bucket
        self.temp_bucket = temp_bucket
        self.gce_cluster_config = gce_cluster_config
        self.master_config = master_config
        self.worker_config = worker_config
        self.secondary_worker_config = secondary_worker_config
        self.software_config = software_config
        self.initialization_actions = initialization_actions
        self.encryption_config = encryption_config
        self.autoscaling_config = autoscaling_config
        self.security_config = security_config
        self.lifecycle_config = lifecycle_config
        self.endpoint_config = endpoint_config
        self.gke_cluster_config = gke_cluster_config
        self.metastore_config = metastore_config
        self.dataproc_metric_config = dataproc_metric_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfig()
        if Primitive.to_proto(resource.staging_bucket):
            res.staging_bucket = Primitive.to_proto(resource.staging_bucket)
        if Primitive.to_proto(resource.temp_bucket):
            res.temp_bucket = Primitive.to_proto(resource.temp_bucket)
        if ClusterConfigGceClusterConfig.to_proto(resource.gce_cluster_config):
            res.gce_cluster_config.CopyFrom(
                ClusterConfigGceClusterConfig.to_proto(resource.gce_cluster_config)
            )
        else:
            res.ClearField("gce_cluster_config")
        if ClusterConfigMasterConfig.to_proto(resource.master_config):
            res.master_config.CopyFrom(
                ClusterConfigMasterConfig.to_proto(resource.master_config)
            )
        else:
            res.ClearField("master_config")
        if ClusterConfigWorkerConfig.to_proto(resource.worker_config):
            res.worker_config.CopyFrom(
                ClusterConfigWorkerConfig.to_proto(resource.worker_config)
            )
        else:
            res.ClearField("worker_config")
        if ClusterConfigSecondaryWorkerConfig.to_proto(
            resource.secondary_worker_config
        ):
            res.secondary_worker_config.CopyFrom(
                ClusterConfigSecondaryWorkerConfig.to_proto(
                    resource.secondary_worker_config
                )
            )
        else:
            res.ClearField("secondary_worker_config")
        if ClusterConfigSoftwareConfig.to_proto(resource.software_config):
            res.software_config.CopyFrom(
                ClusterConfigSoftwareConfig.to_proto(resource.software_config)
            )
        else:
            res.ClearField("software_config")
        if ClusterConfigInitializationActionsArray.to_proto(
            resource.initialization_actions
        ):
            res.initialization_actions.extend(
                ClusterConfigInitializationActionsArray.to_proto(
                    resource.initialization_actions
                )
            )
        if ClusterConfigEncryptionConfig.to_proto(resource.encryption_config):
            res.encryption_config.CopyFrom(
                ClusterConfigEncryptionConfig.to_proto(resource.encryption_config)
            )
        else:
            res.ClearField("encryption_config")
        if ClusterConfigAutoscalingConfig.to_proto(resource.autoscaling_config):
            res.autoscaling_config.CopyFrom(
                ClusterConfigAutoscalingConfig.to_proto(resource.autoscaling_config)
            )
        else:
            res.ClearField("autoscaling_config")
        if ClusterConfigSecurityConfig.to_proto(resource.security_config):
            res.security_config.CopyFrom(
                ClusterConfigSecurityConfig.to_proto(resource.security_config)
            )
        else:
            res.ClearField("security_config")
        if ClusterConfigLifecycleConfig.to_proto(resource.lifecycle_config):
            res.lifecycle_config.CopyFrom(
                ClusterConfigLifecycleConfig.to_proto(resource.lifecycle_config)
            )
        else:
            res.ClearField("lifecycle_config")
        if ClusterConfigEndpointConfig.to_proto(resource.endpoint_config):
            res.endpoint_config.CopyFrom(
                ClusterConfigEndpointConfig.to_proto(resource.endpoint_config)
            )
        else:
            res.ClearField("endpoint_config")
        if ClusterConfigGkeClusterConfig.to_proto(resource.gke_cluster_config):
            res.gke_cluster_config.CopyFrom(
                ClusterConfigGkeClusterConfig.to_proto(resource.gke_cluster_config)
            )
        else:
            res.ClearField("gke_cluster_config")
        if ClusterConfigMetastoreConfig.to_proto(resource.metastore_config):
            res.metastore_config.CopyFrom(
                ClusterConfigMetastoreConfig.to_proto(resource.metastore_config)
            )
        else:
            res.ClearField("metastore_config")
        if ClusterConfigDataprocMetricConfig.to_proto(resource.dataproc_metric_config):
            res.dataproc_metric_config.CopyFrom(
                ClusterConfigDataprocMetricConfig.to_proto(
                    resource.dataproc_metric_config
                )
            )
        else:
            res.ClearField("dataproc_metric_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfig(
            staging_bucket=Primitive.from_proto(resource.staging_bucket),
            temp_bucket=Primitive.from_proto(resource.temp_bucket),
            gce_cluster_config=ClusterConfigGceClusterConfig.from_proto(
                resource.gce_cluster_config
            ),
            master_config=ClusterConfigMasterConfig.from_proto(resource.master_config),
            worker_config=ClusterConfigWorkerConfig.from_proto(resource.worker_config),
            secondary_worker_config=ClusterConfigSecondaryWorkerConfig.from_proto(
                resource.secondary_worker_config
            ),
            software_config=ClusterConfigSoftwareConfig.from_proto(
                resource.software_config
            ),
            initialization_actions=ClusterConfigInitializationActionsArray.from_proto(
                resource.initialization_actions
            ),
            encryption_config=ClusterConfigEncryptionConfig.from_proto(
                resource.encryption_config
            ),
            autoscaling_config=ClusterConfigAutoscalingConfig.from_proto(
                resource.autoscaling_config
            ),
            security_config=ClusterConfigSecurityConfig.from_proto(
                resource.security_config
            ),
            lifecycle_config=ClusterConfigLifecycleConfig.from_proto(
                resource.lifecycle_config
            ),
            endpoint_config=ClusterConfigEndpointConfig.from_proto(
                resource.endpoint_config
            ),
            gke_cluster_config=ClusterConfigGkeClusterConfig.from_proto(
                resource.gke_cluster_config
            ),
            metastore_config=ClusterConfigMetastoreConfig.from_proto(
                resource.metastore_config
            ),
            dataproc_metric_config=ClusterConfigDataprocMetricConfig.from_proto(
                resource.dataproc_metric_config
            ),
        )


class ClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfig.from_proto(i) for i in resources]


class ClusterConfigGceClusterConfig(object):
    def __init__(
        self,
        zone: str = None,
        network: str = None,
        subnetwork: str = None,
        internal_ip_only: bool = None,
        private_ipv6_google_access: str = None,
        service_account: str = None,
        service_account_scopes: list = None,
        tags: list = None,
        metadata: dict = None,
        reservation_affinity: dict = None,
        node_group_affinity: dict = None,
        shielded_instance_config: dict = None,
        confidential_instance_config: dict = None,
    ):
        self.zone = zone
        self.network = network
        self.subnetwork = subnetwork
        self.internal_ip_only = internal_ip_only
        self.private_ipv6_google_access = private_ipv6_google_access
        self.service_account = service_account
        self.service_account_scopes = service_account_scopes
        self.tags = tags
        self.metadata = metadata
        self.reservation_affinity = reservation_affinity
        self.node_group_affinity = node_group_affinity
        self.shielded_instance_config = shielded_instance_config
        self.confidential_instance_config = confidential_instance_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigGceClusterConfig()
        if Primitive.to_proto(resource.zone):
            res.zone = Primitive.to_proto(resource.zone)
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        if Primitive.to_proto(resource.subnetwork):
            res.subnetwork = Primitive.to_proto(resource.subnetwork)
        if Primitive.to_proto(resource.internal_ip_only):
            res.internal_ip_only = Primitive.to_proto(resource.internal_ip_only)
        if ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.to_proto(
            resource.private_ipv6_google_access
        ):
            res.private_ipv6_google_access = (
                ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.to_proto(
                    resource.private_ipv6_google_access
                )
            )
        if Primitive.to_proto(resource.service_account):
            res.service_account = Primitive.to_proto(resource.service_account)
        if Primitive.to_proto(resource.service_account_scopes):
            res.service_account_scopes.extend(
                Primitive.to_proto(resource.service_account_scopes)
            )
        if Primitive.to_proto(resource.tags):
            res.tags.extend(Primitive.to_proto(resource.tags))
        if Primitive.to_proto(resource.metadata):
            res.metadata = Primitive.to_proto(resource.metadata)
        if ClusterConfigGceClusterConfigReservationAffinity.to_proto(
            resource.reservation_affinity
        ):
            res.reservation_affinity.CopyFrom(
                ClusterConfigGceClusterConfigReservationAffinity.to_proto(
                    resource.reservation_affinity
                )
            )
        else:
            res.ClearField("reservation_affinity")
        if ClusterConfigGceClusterConfigNodeGroupAffinity.to_proto(
            resource.node_group_affinity
        ):
            res.node_group_affinity.CopyFrom(
                ClusterConfigGceClusterConfigNodeGroupAffinity.to_proto(
                    resource.node_group_affinity
                )
            )
        else:
            res.ClearField("node_group_affinity")
        if ClusterConfigGceClusterConfigShieldedInstanceConfig.to_proto(
            resource.shielded_instance_config
        ):
            res.shielded_instance_config.CopyFrom(
                ClusterConfigGceClusterConfigShieldedInstanceConfig.to_proto(
                    resource.shielded_instance_config
                )
            )
        else:
            res.ClearField("shielded_instance_config")
        if ClusterConfigGceClusterConfigConfidentialInstanceConfig.to_proto(
            resource.confidential_instance_config
        ):
            res.confidential_instance_config.CopyFrom(
                ClusterConfigGceClusterConfigConfidentialInstanceConfig.to_proto(
                    resource.confidential_instance_config
                )
            )
        else:
            res.ClearField("confidential_instance_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigGceClusterConfig(
            zone=Primitive.from_proto(resource.zone),
            network=Primitive.from_proto(resource.network),
            subnetwork=Primitive.from_proto(resource.subnetwork),
            internal_ip_only=Primitive.from_proto(resource.internal_ip_only),
            private_ipv6_google_access=ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.from_proto(
                resource.private_ipv6_google_access
            ),
            service_account=Primitive.from_proto(resource.service_account),
            service_account_scopes=Primitive.from_proto(
                resource.service_account_scopes
            ),
            tags=Primitive.from_proto(resource.tags),
            metadata=Primitive.from_proto(resource.metadata),
            reservation_affinity=ClusterConfigGceClusterConfigReservationAffinity.from_proto(
                resource.reservation_affinity
            ),
            node_group_affinity=ClusterConfigGceClusterConfigNodeGroupAffinity.from_proto(
                resource.node_group_affinity
            ),
            shielded_instance_config=ClusterConfigGceClusterConfigShieldedInstanceConfig.from_proto(
                resource.shielded_instance_config
            ),
            confidential_instance_config=ClusterConfigGceClusterConfigConfidentialInstanceConfig.from_proto(
                resource.confidential_instance_config
            ),
        )


class ClusterConfigGceClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigGceClusterConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigGceClusterConfig.from_proto(i) for i in resources]


class ClusterConfigGceClusterConfigReservationAffinity(object):
    def __init__(
        self, consume_reservation_type: str = None, key: str = None, values: list = None
    ):
        self.consume_reservation_type = consume_reservation_type
        self.key = key
        self.values = values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterConfigGceClusterConfigReservationAffinity()
        )
        if ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
            resource.consume_reservation_type
        ):
            res.consume_reservation_type = ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
                resource.consume_reservation_type
            )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.values):
            res.values.extend(Primitive.to_proto(resource.values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigGceClusterConfigReservationAffinity(
            consume_reservation_type=ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.from_proto(
                resource.consume_reservation_type
            ),
            key=Primitive.from_proto(resource.key),
            values=Primitive.from_proto(resource.values),
        )


class ClusterConfigGceClusterConfigReservationAffinityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigGceClusterConfigReservationAffinity.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigGceClusterConfigReservationAffinity.from_proto(i)
            for i in resources
        ]


class ClusterConfigGceClusterConfigNodeGroupAffinity(object):
    def __init__(self, node_group: str = None):
        self.node_group = node_group

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigGceClusterConfigNodeGroupAffinity()
        if Primitive.to_proto(resource.node_group):
            res.node_group = Primitive.to_proto(resource.node_group)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigGceClusterConfigNodeGroupAffinity(
            node_group=Primitive.from_proto(resource.node_group),
        )


class ClusterConfigGceClusterConfigNodeGroupAffinityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigGceClusterConfigNodeGroupAffinity.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigGceClusterConfigNodeGroupAffinity.from_proto(i)
            for i in resources
        ]


class ClusterConfigGceClusterConfigShieldedInstanceConfig(object):
    def __init__(
        self,
        enable_secure_boot: bool = None,
        enable_vtpm: bool = None,
        enable_integrity_monitoring: bool = None,
    ):
        self.enable_secure_boot = enable_secure_boot
        self.enable_vtpm = enable_vtpm
        self.enable_integrity_monitoring = enable_integrity_monitoring

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterConfigGceClusterConfigShieldedInstanceConfig()
        )
        if Primitive.to_proto(resource.enable_secure_boot):
            res.enable_secure_boot = Primitive.to_proto(resource.enable_secure_boot)
        if Primitive.to_proto(resource.enable_vtpm):
            res.enable_vtpm = Primitive.to_proto(resource.enable_vtpm)
        if Primitive.to_proto(resource.enable_integrity_monitoring):
            res.enable_integrity_monitoring = Primitive.to_proto(
                resource.enable_integrity_monitoring
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigGceClusterConfigShieldedInstanceConfig(
            enable_secure_boot=Primitive.from_proto(resource.enable_secure_boot),
            enable_vtpm=Primitive.from_proto(resource.enable_vtpm),
            enable_integrity_monitoring=Primitive.from_proto(
                resource.enable_integrity_monitoring
            ),
        )


class ClusterConfigGceClusterConfigShieldedInstanceConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigGceClusterConfigShieldedInstanceConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigGceClusterConfigShieldedInstanceConfig.from_proto(i)
            for i in resources
        ]


class ClusterConfigGceClusterConfigConfidentialInstanceConfig(object):
    def __init__(self, enable_confidential_compute: bool = None):
        self.enable_confidential_compute = enable_confidential_compute

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterConfigGceClusterConfigConfidentialInstanceConfig()
        )
        if Primitive.to_proto(resource.enable_confidential_compute):
            res.enable_confidential_compute = Primitive.to_proto(
                resource.enable_confidential_compute
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigGceClusterConfigConfidentialInstanceConfig(
            enable_confidential_compute=Primitive.from_proto(
                resource.enable_confidential_compute
            ),
        )


class ClusterConfigGceClusterConfigConfidentialInstanceConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigGceClusterConfigConfidentialInstanceConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigGceClusterConfigConfidentialInstanceConfig.from_proto(i)
            for i in resources
        ]


class ClusterConfigMasterConfig(object):
    def __init__(
        self,
        num_instances: int = None,
        instance_names: list = None,
        image: str = None,
        machine_type: str = None,
        disk_config: dict = None,
        is_preemptible: bool = None,
        preemptibility: str = None,
        managed_group_config: dict = None,
        accelerators: list = None,
        min_cpu_platform: str = None,
        instance_references: list = None,
    ):
        self.num_instances = num_instances
        self.instance_names = instance_names
        self.image = image
        self.machine_type = machine_type
        self.disk_config = disk_config
        self.is_preemptible = is_preemptible
        self.preemptibility = preemptibility
        self.managed_group_config = managed_group_config
        self.accelerators = accelerators
        self.min_cpu_platform = min_cpu_platform
        self.instance_references = instance_references

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigMasterConfig()
        if Primitive.to_proto(resource.num_instances):
            res.num_instances = Primitive.to_proto(resource.num_instances)
        if Primitive.to_proto(resource.instance_names):
            res.instance_names.extend(Primitive.to_proto(resource.instance_names))
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if ClusterConfigMasterConfigDiskConfig.to_proto(resource.disk_config):
            res.disk_config.CopyFrom(
                ClusterConfigMasterConfigDiskConfig.to_proto(resource.disk_config)
            )
        else:
            res.ClearField("disk_config")
        if Primitive.to_proto(resource.is_preemptible):
            res.is_preemptible = Primitive.to_proto(resource.is_preemptible)
        if ClusterConfigMasterConfigPreemptibilityEnum.to_proto(
            resource.preemptibility
        ):
            res.preemptibility = ClusterConfigMasterConfigPreemptibilityEnum.to_proto(
                resource.preemptibility
            )
        if ClusterConfigMasterConfigManagedGroupConfig.to_proto(
            resource.managed_group_config
        ):
            res.managed_group_config.CopyFrom(
                ClusterConfigMasterConfigManagedGroupConfig.to_proto(
                    resource.managed_group_config
                )
            )
        else:
            res.ClearField("managed_group_config")
        if ClusterConfigMasterConfigAcceleratorsArray.to_proto(resource.accelerators):
            res.accelerators.extend(
                ClusterConfigMasterConfigAcceleratorsArray.to_proto(
                    resource.accelerators
                )
            )
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        if ClusterConfigMasterConfigInstanceReferencesArray.to_proto(
            resource.instance_references
        ):
            res.instance_references.extend(
                ClusterConfigMasterConfigInstanceReferencesArray.to_proto(
                    resource.instance_references
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigMasterConfig(
            num_instances=Primitive.from_proto(resource.num_instances),
            instance_names=Primitive.from_proto(resource.instance_names),
            image=Primitive.from_proto(resource.image),
            machine_type=Primitive.from_proto(resource.machine_type),
            disk_config=ClusterConfigMasterConfigDiskConfig.from_proto(
                resource.disk_config
            ),
            is_preemptible=Primitive.from_proto(resource.is_preemptible),
            preemptibility=ClusterConfigMasterConfigPreemptibilityEnum.from_proto(
                resource.preemptibility
            ),
            managed_group_config=ClusterConfigMasterConfigManagedGroupConfig.from_proto(
                resource.managed_group_config
            ),
            accelerators=ClusterConfigMasterConfigAcceleratorsArray.from_proto(
                resource.accelerators
            ),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
            instance_references=ClusterConfigMasterConfigInstanceReferencesArray.from_proto(
                resource.instance_references
            ),
        )


class ClusterConfigMasterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigMasterConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigMasterConfig.from_proto(i) for i in resources]


class ClusterConfigMasterConfigDiskConfig(object):
    def __init__(
        self,
        boot_disk_type: str = None,
        boot_disk_size_gb: int = None,
        num_local_ssds: int = None,
        local_ssd_interface: str = None,
    ):
        self.boot_disk_type = boot_disk_type
        self.boot_disk_size_gb = boot_disk_size_gb
        self.num_local_ssds = num_local_ssds
        self.local_ssd_interface = local_ssd_interface

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigMasterConfigDiskConfig()
        if Primitive.to_proto(resource.boot_disk_type):
            res.boot_disk_type = Primitive.to_proto(resource.boot_disk_type)
        if Primitive.to_proto(resource.boot_disk_size_gb):
            res.boot_disk_size_gb = Primitive.to_proto(resource.boot_disk_size_gb)
        if Primitive.to_proto(resource.num_local_ssds):
            res.num_local_ssds = Primitive.to_proto(resource.num_local_ssds)
        if Primitive.to_proto(resource.local_ssd_interface):
            res.local_ssd_interface = Primitive.to_proto(resource.local_ssd_interface)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigMasterConfigDiskConfig(
            boot_disk_type=Primitive.from_proto(resource.boot_disk_type),
            boot_disk_size_gb=Primitive.from_proto(resource.boot_disk_size_gb),
            num_local_ssds=Primitive.from_proto(resource.num_local_ssds),
            local_ssd_interface=Primitive.from_proto(resource.local_ssd_interface),
        )


class ClusterConfigMasterConfigDiskConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigMasterConfigDiskConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigMasterConfigDiskConfig.from_proto(i) for i in resources]


class ClusterConfigMasterConfigManagedGroupConfig(object):
    def __init__(
        self,
        instance_template_name: str = None,
        instance_group_manager_name: str = None,
    ):
        self.instance_template_name = instance_template_name
        self.instance_group_manager_name = instance_group_manager_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigMasterConfigManagedGroupConfig()
        if Primitive.to_proto(resource.instance_template_name):
            res.instance_template_name = Primitive.to_proto(
                resource.instance_template_name
            )
        if Primitive.to_proto(resource.instance_group_manager_name):
            res.instance_group_manager_name = Primitive.to_proto(
                resource.instance_group_manager_name
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigMasterConfigManagedGroupConfig(
            instance_template_name=Primitive.from_proto(
                resource.instance_template_name
            ),
            instance_group_manager_name=Primitive.from_proto(
                resource.instance_group_manager_name
            ),
        )


class ClusterConfigMasterConfigManagedGroupConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigMasterConfigManagedGroupConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigMasterConfigManagedGroupConfig.from_proto(i) for i in resources
        ]


class ClusterConfigMasterConfigAccelerators(object):
    def __init__(self, accelerator_type: str = None, accelerator_count: int = None):
        self.accelerator_type = accelerator_type
        self.accelerator_count = accelerator_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigMasterConfigAccelerators()
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigMasterConfigAccelerators(
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
        )


class ClusterConfigMasterConfigAcceleratorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigMasterConfigAccelerators.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigMasterConfigAccelerators.from_proto(i) for i in resources]


class ClusterConfigMasterConfigInstanceReferences(object):
    def __init__(
        self,
        instance_name: str = None,
        instance_id: str = None,
        public_key: str = None,
        public_ecies_key: str = None,
    ):
        self.instance_name = instance_name
        self.instance_id = instance_id
        self.public_key = public_key
        self.public_ecies_key = public_ecies_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigMasterConfigInstanceReferences()
        if Primitive.to_proto(resource.instance_name):
            res.instance_name = Primitive.to_proto(resource.instance_name)
        if Primitive.to_proto(resource.instance_id):
            res.instance_id = Primitive.to_proto(resource.instance_id)
        if Primitive.to_proto(resource.public_key):
            res.public_key = Primitive.to_proto(resource.public_key)
        if Primitive.to_proto(resource.public_ecies_key):
            res.public_ecies_key = Primitive.to_proto(resource.public_ecies_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigMasterConfigInstanceReferences(
            instance_name=Primitive.from_proto(resource.instance_name),
            instance_id=Primitive.from_proto(resource.instance_id),
            public_key=Primitive.from_proto(resource.public_key),
            public_ecies_key=Primitive.from_proto(resource.public_ecies_key),
        )


class ClusterConfigMasterConfigInstanceReferencesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigMasterConfigInstanceReferences.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigMasterConfigInstanceReferences.from_proto(i) for i in resources
        ]


class ClusterConfigWorkerConfig(object):
    def __init__(
        self,
        num_instances: int = None,
        instance_names: list = None,
        image: str = None,
        machine_type: str = None,
        disk_config: dict = None,
        is_preemptible: bool = None,
        preemptibility: str = None,
        managed_group_config: dict = None,
        accelerators: list = None,
        min_cpu_platform: str = None,
        instance_references: list = None,
    ):
        self.num_instances = num_instances
        self.instance_names = instance_names
        self.image = image
        self.machine_type = machine_type
        self.disk_config = disk_config
        self.is_preemptible = is_preemptible
        self.preemptibility = preemptibility
        self.managed_group_config = managed_group_config
        self.accelerators = accelerators
        self.min_cpu_platform = min_cpu_platform
        self.instance_references = instance_references

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigWorkerConfig()
        if Primitive.to_proto(resource.num_instances):
            res.num_instances = Primitive.to_proto(resource.num_instances)
        if Primitive.to_proto(resource.instance_names):
            res.instance_names.extend(Primitive.to_proto(resource.instance_names))
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if ClusterConfigWorkerConfigDiskConfig.to_proto(resource.disk_config):
            res.disk_config.CopyFrom(
                ClusterConfigWorkerConfigDiskConfig.to_proto(resource.disk_config)
            )
        else:
            res.ClearField("disk_config")
        if Primitive.to_proto(resource.is_preemptible):
            res.is_preemptible = Primitive.to_proto(resource.is_preemptible)
        if ClusterConfigWorkerConfigPreemptibilityEnum.to_proto(
            resource.preemptibility
        ):
            res.preemptibility = ClusterConfigWorkerConfigPreemptibilityEnum.to_proto(
                resource.preemptibility
            )
        if ClusterConfigWorkerConfigManagedGroupConfig.to_proto(
            resource.managed_group_config
        ):
            res.managed_group_config.CopyFrom(
                ClusterConfigWorkerConfigManagedGroupConfig.to_proto(
                    resource.managed_group_config
                )
            )
        else:
            res.ClearField("managed_group_config")
        if ClusterConfigWorkerConfigAcceleratorsArray.to_proto(resource.accelerators):
            res.accelerators.extend(
                ClusterConfigWorkerConfigAcceleratorsArray.to_proto(
                    resource.accelerators
                )
            )
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        if ClusterConfigWorkerConfigInstanceReferencesArray.to_proto(
            resource.instance_references
        ):
            res.instance_references.extend(
                ClusterConfigWorkerConfigInstanceReferencesArray.to_proto(
                    resource.instance_references
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigWorkerConfig(
            num_instances=Primitive.from_proto(resource.num_instances),
            instance_names=Primitive.from_proto(resource.instance_names),
            image=Primitive.from_proto(resource.image),
            machine_type=Primitive.from_proto(resource.machine_type),
            disk_config=ClusterConfigWorkerConfigDiskConfig.from_proto(
                resource.disk_config
            ),
            is_preemptible=Primitive.from_proto(resource.is_preemptible),
            preemptibility=ClusterConfigWorkerConfigPreemptibilityEnum.from_proto(
                resource.preemptibility
            ),
            managed_group_config=ClusterConfigWorkerConfigManagedGroupConfig.from_proto(
                resource.managed_group_config
            ),
            accelerators=ClusterConfigWorkerConfigAcceleratorsArray.from_proto(
                resource.accelerators
            ),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
            instance_references=ClusterConfigWorkerConfigInstanceReferencesArray.from_proto(
                resource.instance_references
            ),
        )


class ClusterConfigWorkerConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigWorkerConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigWorkerConfig.from_proto(i) for i in resources]


class ClusterConfigWorkerConfigDiskConfig(object):
    def __init__(
        self,
        boot_disk_type: str = None,
        boot_disk_size_gb: int = None,
        num_local_ssds: int = None,
        local_ssd_interface: str = None,
    ):
        self.boot_disk_type = boot_disk_type
        self.boot_disk_size_gb = boot_disk_size_gb
        self.num_local_ssds = num_local_ssds
        self.local_ssd_interface = local_ssd_interface

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigWorkerConfigDiskConfig()
        if Primitive.to_proto(resource.boot_disk_type):
            res.boot_disk_type = Primitive.to_proto(resource.boot_disk_type)
        if Primitive.to_proto(resource.boot_disk_size_gb):
            res.boot_disk_size_gb = Primitive.to_proto(resource.boot_disk_size_gb)
        if Primitive.to_proto(resource.num_local_ssds):
            res.num_local_ssds = Primitive.to_proto(resource.num_local_ssds)
        if Primitive.to_proto(resource.local_ssd_interface):
            res.local_ssd_interface = Primitive.to_proto(resource.local_ssd_interface)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigWorkerConfigDiskConfig(
            boot_disk_type=Primitive.from_proto(resource.boot_disk_type),
            boot_disk_size_gb=Primitive.from_proto(resource.boot_disk_size_gb),
            num_local_ssds=Primitive.from_proto(resource.num_local_ssds),
            local_ssd_interface=Primitive.from_proto(resource.local_ssd_interface),
        )


class ClusterConfigWorkerConfigDiskConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigWorkerConfigDiskConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigWorkerConfigDiskConfig.from_proto(i) for i in resources]


class ClusterConfigWorkerConfigManagedGroupConfig(object):
    def __init__(
        self,
        instance_template_name: str = None,
        instance_group_manager_name: str = None,
    ):
        self.instance_template_name = instance_template_name
        self.instance_group_manager_name = instance_group_manager_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigWorkerConfigManagedGroupConfig()
        if Primitive.to_proto(resource.instance_template_name):
            res.instance_template_name = Primitive.to_proto(
                resource.instance_template_name
            )
        if Primitive.to_proto(resource.instance_group_manager_name):
            res.instance_group_manager_name = Primitive.to_proto(
                resource.instance_group_manager_name
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigWorkerConfigManagedGroupConfig(
            instance_template_name=Primitive.from_proto(
                resource.instance_template_name
            ),
            instance_group_manager_name=Primitive.from_proto(
                resource.instance_group_manager_name
            ),
        )


class ClusterConfigWorkerConfigManagedGroupConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigWorkerConfigManagedGroupConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigWorkerConfigManagedGroupConfig.from_proto(i) for i in resources
        ]


class ClusterConfigWorkerConfigAccelerators(object):
    def __init__(self, accelerator_type: str = None, accelerator_count: int = None):
        self.accelerator_type = accelerator_type
        self.accelerator_count = accelerator_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigWorkerConfigAccelerators()
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigWorkerConfigAccelerators(
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
        )


class ClusterConfigWorkerConfigAcceleratorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigWorkerConfigAccelerators.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigWorkerConfigAccelerators.from_proto(i) for i in resources]


class ClusterConfigWorkerConfigInstanceReferences(object):
    def __init__(
        self,
        instance_name: str = None,
        instance_id: str = None,
        public_key: str = None,
        public_ecies_key: str = None,
    ):
        self.instance_name = instance_name
        self.instance_id = instance_id
        self.public_key = public_key
        self.public_ecies_key = public_ecies_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigWorkerConfigInstanceReferences()
        if Primitive.to_proto(resource.instance_name):
            res.instance_name = Primitive.to_proto(resource.instance_name)
        if Primitive.to_proto(resource.instance_id):
            res.instance_id = Primitive.to_proto(resource.instance_id)
        if Primitive.to_proto(resource.public_key):
            res.public_key = Primitive.to_proto(resource.public_key)
        if Primitive.to_proto(resource.public_ecies_key):
            res.public_ecies_key = Primitive.to_proto(resource.public_ecies_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigWorkerConfigInstanceReferences(
            instance_name=Primitive.from_proto(resource.instance_name),
            instance_id=Primitive.from_proto(resource.instance_id),
            public_key=Primitive.from_proto(resource.public_key),
            public_ecies_key=Primitive.from_proto(resource.public_ecies_key),
        )


class ClusterConfigWorkerConfigInstanceReferencesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigWorkerConfigInstanceReferences.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigWorkerConfigInstanceReferences.from_proto(i) for i in resources
        ]


class ClusterConfigSecondaryWorkerConfig(object):
    def __init__(
        self,
        num_instances: int = None,
        instance_names: list = None,
        image: str = None,
        machine_type: str = None,
        disk_config: dict = None,
        is_preemptible: bool = None,
        preemptibility: str = None,
        managed_group_config: dict = None,
        accelerators: list = None,
        min_cpu_platform: str = None,
        instance_references: list = None,
    ):
        self.num_instances = num_instances
        self.instance_names = instance_names
        self.image = image
        self.machine_type = machine_type
        self.disk_config = disk_config
        self.is_preemptible = is_preemptible
        self.preemptibility = preemptibility
        self.managed_group_config = managed_group_config
        self.accelerators = accelerators
        self.min_cpu_platform = min_cpu_platform
        self.instance_references = instance_references

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigSecondaryWorkerConfig()
        if Primitive.to_proto(resource.num_instances):
            res.num_instances = Primitive.to_proto(resource.num_instances)
        if Primitive.to_proto(resource.instance_names):
            res.instance_names.extend(Primitive.to_proto(resource.instance_names))
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if ClusterConfigSecondaryWorkerConfigDiskConfig.to_proto(resource.disk_config):
            res.disk_config.CopyFrom(
                ClusterConfigSecondaryWorkerConfigDiskConfig.to_proto(
                    resource.disk_config
                )
            )
        else:
            res.ClearField("disk_config")
        if Primitive.to_proto(resource.is_preemptible):
            res.is_preemptible = Primitive.to_proto(resource.is_preemptible)
        if ClusterConfigSecondaryWorkerConfigPreemptibilityEnum.to_proto(
            resource.preemptibility
        ):
            res.preemptibility = (
                ClusterConfigSecondaryWorkerConfigPreemptibilityEnum.to_proto(
                    resource.preemptibility
                )
            )
        if ClusterConfigSecondaryWorkerConfigManagedGroupConfig.to_proto(
            resource.managed_group_config
        ):
            res.managed_group_config.CopyFrom(
                ClusterConfigSecondaryWorkerConfigManagedGroupConfig.to_proto(
                    resource.managed_group_config
                )
            )
        else:
            res.ClearField("managed_group_config")
        if ClusterConfigSecondaryWorkerConfigAcceleratorsArray.to_proto(
            resource.accelerators
        ):
            res.accelerators.extend(
                ClusterConfigSecondaryWorkerConfigAcceleratorsArray.to_proto(
                    resource.accelerators
                )
            )
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        if ClusterConfigSecondaryWorkerConfigInstanceReferencesArray.to_proto(
            resource.instance_references
        ):
            res.instance_references.extend(
                ClusterConfigSecondaryWorkerConfigInstanceReferencesArray.to_proto(
                    resource.instance_references
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigSecondaryWorkerConfig(
            num_instances=Primitive.from_proto(resource.num_instances),
            instance_names=Primitive.from_proto(resource.instance_names),
            image=Primitive.from_proto(resource.image),
            machine_type=Primitive.from_proto(resource.machine_type),
            disk_config=ClusterConfigSecondaryWorkerConfigDiskConfig.from_proto(
                resource.disk_config
            ),
            is_preemptible=Primitive.from_proto(resource.is_preemptible),
            preemptibility=ClusterConfigSecondaryWorkerConfigPreemptibilityEnum.from_proto(
                resource.preemptibility
            ),
            managed_group_config=ClusterConfigSecondaryWorkerConfigManagedGroupConfig.from_proto(
                resource.managed_group_config
            ),
            accelerators=ClusterConfigSecondaryWorkerConfigAcceleratorsArray.from_proto(
                resource.accelerators
            ),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
            instance_references=ClusterConfigSecondaryWorkerConfigInstanceReferencesArray.from_proto(
                resource.instance_references
            ),
        )


class ClusterConfigSecondaryWorkerConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigSecondaryWorkerConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigSecondaryWorkerConfig.from_proto(i) for i in resources]


class ClusterConfigSecondaryWorkerConfigDiskConfig(object):
    def __init__(
        self,
        boot_disk_type: str = None,
        boot_disk_size_gb: int = None,
        num_local_ssds: int = None,
        local_ssd_interface: str = None,
    ):
        self.boot_disk_type = boot_disk_type
        self.boot_disk_size_gb = boot_disk_size_gb
        self.num_local_ssds = num_local_ssds
        self.local_ssd_interface = local_ssd_interface

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigSecondaryWorkerConfigDiskConfig()
        if Primitive.to_proto(resource.boot_disk_type):
            res.boot_disk_type = Primitive.to_proto(resource.boot_disk_type)
        if Primitive.to_proto(resource.boot_disk_size_gb):
            res.boot_disk_size_gb = Primitive.to_proto(resource.boot_disk_size_gb)
        if Primitive.to_proto(resource.num_local_ssds):
            res.num_local_ssds = Primitive.to_proto(resource.num_local_ssds)
        if Primitive.to_proto(resource.local_ssd_interface):
            res.local_ssd_interface = Primitive.to_proto(resource.local_ssd_interface)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigSecondaryWorkerConfigDiskConfig(
            boot_disk_type=Primitive.from_proto(resource.boot_disk_type),
            boot_disk_size_gb=Primitive.from_proto(resource.boot_disk_size_gb),
            num_local_ssds=Primitive.from_proto(resource.num_local_ssds),
            local_ssd_interface=Primitive.from_proto(resource.local_ssd_interface),
        )


class ClusterConfigSecondaryWorkerConfigDiskConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigSecondaryWorkerConfigDiskConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigSecondaryWorkerConfigDiskConfig.from_proto(i)
            for i in resources
        ]


class ClusterConfigSecondaryWorkerConfigManagedGroupConfig(object):
    def __init__(
        self,
        instance_template_name: str = None,
        instance_group_manager_name: str = None,
    ):
        self.instance_template_name = instance_template_name
        self.instance_group_manager_name = instance_group_manager_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterConfigSecondaryWorkerConfigManagedGroupConfig()
        )
        if Primitive.to_proto(resource.instance_template_name):
            res.instance_template_name = Primitive.to_proto(
                resource.instance_template_name
            )
        if Primitive.to_proto(resource.instance_group_manager_name):
            res.instance_group_manager_name = Primitive.to_proto(
                resource.instance_group_manager_name
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigSecondaryWorkerConfigManagedGroupConfig(
            instance_template_name=Primitive.from_proto(
                resource.instance_template_name
            ),
            instance_group_manager_name=Primitive.from_proto(
                resource.instance_group_manager_name
            ),
        )


class ClusterConfigSecondaryWorkerConfigManagedGroupConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigSecondaryWorkerConfigManagedGroupConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigSecondaryWorkerConfigManagedGroupConfig.from_proto(i)
            for i in resources
        ]


class ClusterConfigSecondaryWorkerConfigAccelerators(object):
    def __init__(self, accelerator_type: str = None, accelerator_count: int = None):
        self.accelerator_type = accelerator_type
        self.accelerator_count = accelerator_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigSecondaryWorkerConfigAccelerators()
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigSecondaryWorkerConfigAccelerators(
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
        )


class ClusterConfigSecondaryWorkerConfigAcceleratorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigSecondaryWorkerConfigAccelerators.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigSecondaryWorkerConfigAccelerators.from_proto(i)
            for i in resources
        ]


class ClusterConfigSecondaryWorkerConfigInstanceReferences(object):
    def __init__(
        self,
        instance_name: str = None,
        instance_id: str = None,
        public_key: str = None,
        public_ecies_key: str = None,
    ):
        self.instance_name = instance_name
        self.instance_id = instance_id
        self.public_key = public_key
        self.public_ecies_key = public_ecies_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterConfigSecondaryWorkerConfigInstanceReferences()
        )
        if Primitive.to_proto(resource.instance_name):
            res.instance_name = Primitive.to_proto(resource.instance_name)
        if Primitive.to_proto(resource.instance_id):
            res.instance_id = Primitive.to_proto(resource.instance_id)
        if Primitive.to_proto(resource.public_key):
            res.public_key = Primitive.to_proto(resource.public_key)
        if Primitive.to_proto(resource.public_ecies_key):
            res.public_ecies_key = Primitive.to_proto(resource.public_ecies_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigSecondaryWorkerConfigInstanceReferences(
            instance_name=Primitive.from_proto(resource.instance_name),
            instance_id=Primitive.from_proto(resource.instance_id),
            public_key=Primitive.from_proto(resource.public_key),
            public_ecies_key=Primitive.from_proto(resource.public_ecies_key),
        )


class ClusterConfigSecondaryWorkerConfigInstanceReferencesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigSecondaryWorkerConfigInstanceReferences.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigSecondaryWorkerConfigInstanceReferences.from_proto(i)
            for i in resources
        ]


class ClusterConfigSoftwareConfig(object):
    def __init__(
        self,
        image_version: str = None,
        properties: dict = None,
        optional_components: list = None,
    ):
        self.image_version = image_version
        self.properties = properties
        self.optional_components = optional_components

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigSoftwareConfig()
        if Primitive.to_proto(resource.image_version):
            res.image_version = Primitive.to_proto(resource.image_version)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if ClusterConfigSoftwareConfigOptionalComponentsEnumArray.to_proto(
            resource.optional_components
        ):
            res.optional_components.extend(
                ClusterConfigSoftwareConfigOptionalComponentsEnumArray.to_proto(
                    resource.optional_components
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigSoftwareConfig(
            image_version=Primitive.from_proto(resource.image_version),
            properties=Primitive.from_proto(resource.properties),
            optional_components=ClusterConfigSoftwareConfigOptionalComponentsEnumArray.from_proto(
                resource.optional_components
            ),
        )


class ClusterConfigSoftwareConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigSoftwareConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigSoftwareConfig.from_proto(i) for i in resources]


class ClusterConfigInitializationActions(object):
    def __init__(self, executable_file: str = None, execution_timeout: str = None):
        self.executable_file = executable_file
        self.execution_timeout = execution_timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigInitializationActions()
        if Primitive.to_proto(resource.executable_file):
            res.executable_file = Primitive.to_proto(resource.executable_file)
        if Primitive.to_proto(resource.execution_timeout):
            res.execution_timeout = Primitive.to_proto(resource.execution_timeout)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigInitializationActions(
            executable_file=Primitive.from_proto(resource.executable_file),
            execution_timeout=Primitive.from_proto(resource.execution_timeout),
        )


class ClusterConfigInitializationActionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigInitializationActions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigInitializationActions.from_proto(i) for i in resources]


class ClusterConfigEncryptionConfig(object):
    def __init__(self, gce_pd_kms_key_name: str = None):
        self.gce_pd_kms_key_name = gce_pd_kms_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigEncryptionConfig()
        if Primitive.to_proto(resource.gce_pd_kms_key_name):
            res.gce_pd_kms_key_name = Primitive.to_proto(resource.gce_pd_kms_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigEncryptionConfig(
            gce_pd_kms_key_name=Primitive.from_proto(resource.gce_pd_kms_key_name),
        )


class ClusterConfigEncryptionConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigEncryptionConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigEncryptionConfig.from_proto(i) for i in resources]


class ClusterConfigAutoscalingConfig(object):
    def __init__(self, policy: str = None):
        self.policy = policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigAutoscalingConfig()
        if Primitive.to_proto(resource.policy):
            res.policy = Primitive.to_proto(resource.policy)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigAutoscalingConfig(
            policy=Primitive.from_proto(resource.policy),
        )


class ClusterConfigAutoscalingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigAutoscalingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigAutoscalingConfig.from_proto(i) for i in resources]


class ClusterConfigSecurityConfig(object):
    def __init__(self, kerberos_config: dict = None, identity_config: dict = None):
        self.kerberos_config = kerberos_config
        self.identity_config = identity_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigSecurityConfig()
        if ClusterConfigSecurityConfigKerberosConfig.to_proto(resource.kerberos_config):
            res.kerberos_config.CopyFrom(
                ClusterConfigSecurityConfigKerberosConfig.to_proto(
                    resource.kerberos_config
                )
            )
        else:
            res.ClearField("kerberos_config")
        if ClusterConfigSecurityConfigIdentityConfig.to_proto(resource.identity_config):
            res.identity_config.CopyFrom(
                ClusterConfigSecurityConfigIdentityConfig.to_proto(
                    resource.identity_config
                )
            )
        else:
            res.ClearField("identity_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigSecurityConfig(
            kerberos_config=ClusterConfigSecurityConfigKerberosConfig.from_proto(
                resource.kerberos_config
            ),
            identity_config=ClusterConfigSecurityConfigIdentityConfig.from_proto(
                resource.identity_config
            ),
        )


class ClusterConfigSecurityConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigSecurityConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigSecurityConfig.from_proto(i) for i in resources]


class ClusterConfigSecurityConfigKerberosConfig(object):
    def __init__(
        self,
        enable_kerberos: bool = None,
        root_principal_password: str = None,
        kms_key: str = None,
        keystore: str = None,
        truststore: str = None,
        keystore_password: str = None,
        key_password: str = None,
        truststore_password: str = None,
        cross_realm_trust_realm: str = None,
        cross_realm_trust_kdc: str = None,
        cross_realm_trust_admin_server: str = None,
        cross_realm_trust_shared_password: str = None,
        kdc_db_key: str = None,
        tgt_lifetime_hours: int = None,
        realm: str = None,
    ):
        self.enable_kerberos = enable_kerberos
        self.root_principal_password = root_principal_password
        self.kms_key = kms_key
        self.keystore = keystore
        self.truststore = truststore
        self.keystore_password = keystore_password
        self.key_password = key_password
        self.truststore_password = truststore_password
        self.cross_realm_trust_realm = cross_realm_trust_realm
        self.cross_realm_trust_kdc = cross_realm_trust_kdc
        self.cross_realm_trust_admin_server = cross_realm_trust_admin_server
        self.cross_realm_trust_shared_password = cross_realm_trust_shared_password
        self.kdc_db_key = kdc_db_key
        self.tgt_lifetime_hours = tgt_lifetime_hours
        self.realm = realm

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigSecurityConfigKerberosConfig()
        if Primitive.to_proto(resource.enable_kerberos):
            res.enable_kerberos = Primitive.to_proto(resource.enable_kerberos)
        if Primitive.to_proto(resource.root_principal_password):
            res.root_principal_password = Primitive.to_proto(
                resource.root_principal_password
            )
        if Primitive.to_proto(resource.kms_key):
            res.kms_key = Primitive.to_proto(resource.kms_key)
        if Primitive.to_proto(resource.keystore):
            res.keystore = Primitive.to_proto(resource.keystore)
        if Primitive.to_proto(resource.truststore):
            res.truststore = Primitive.to_proto(resource.truststore)
        if Primitive.to_proto(resource.keystore_password):
            res.keystore_password = Primitive.to_proto(resource.keystore_password)
        if Primitive.to_proto(resource.key_password):
            res.key_password = Primitive.to_proto(resource.key_password)
        if Primitive.to_proto(resource.truststore_password):
            res.truststore_password = Primitive.to_proto(resource.truststore_password)
        if Primitive.to_proto(resource.cross_realm_trust_realm):
            res.cross_realm_trust_realm = Primitive.to_proto(
                resource.cross_realm_trust_realm
            )
        if Primitive.to_proto(resource.cross_realm_trust_kdc):
            res.cross_realm_trust_kdc = Primitive.to_proto(
                resource.cross_realm_trust_kdc
            )
        if Primitive.to_proto(resource.cross_realm_trust_admin_server):
            res.cross_realm_trust_admin_server = Primitive.to_proto(
                resource.cross_realm_trust_admin_server
            )
        if Primitive.to_proto(resource.cross_realm_trust_shared_password):
            res.cross_realm_trust_shared_password = Primitive.to_proto(
                resource.cross_realm_trust_shared_password
            )
        if Primitive.to_proto(resource.kdc_db_key):
            res.kdc_db_key = Primitive.to_proto(resource.kdc_db_key)
        if Primitive.to_proto(resource.tgt_lifetime_hours):
            res.tgt_lifetime_hours = Primitive.to_proto(resource.tgt_lifetime_hours)
        if Primitive.to_proto(resource.realm):
            res.realm = Primitive.to_proto(resource.realm)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigSecurityConfigKerberosConfig(
            enable_kerberos=Primitive.from_proto(resource.enable_kerberos),
            root_principal_password=Primitive.from_proto(
                resource.root_principal_password
            ),
            kms_key=Primitive.from_proto(resource.kms_key),
            keystore=Primitive.from_proto(resource.keystore),
            truststore=Primitive.from_proto(resource.truststore),
            keystore_password=Primitive.from_proto(resource.keystore_password),
            key_password=Primitive.from_proto(resource.key_password),
            truststore_password=Primitive.from_proto(resource.truststore_password),
            cross_realm_trust_realm=Primitive.from_proto(
                resource.cross_realm_trust_realm
            ),
            cross_realm_trust_kdc=Primitive.from_proto(resource.cross_realm_trust_kdc),
            cross_realm_trust_admin_server=Primitive.from_proto(
                resource.cross_realm_trust_admin_server
            ),
            cross_realm_trust_shared_password=Primitive.from_proto(
                resource.cross_realm_trust_shared_password
            ),
            kdc_db_key=Primitive.from_proto(resource.kdc_db_key),
            tgt_lifetime_hours=Primitive.from_proto(resource.tgt_lifetime_hours),
            realm=Primitive.from_proto(resource.realm),
        )


class ClusterConfigSecurityConfigKerberosConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigSecurityConfigKerberosConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigSecurityConfigKerberosConfig.from_proto(i) for i in resources
        ]


class ClusterConfigSecurityConfigIdentityConfig(object):
    def __init__(self, user_service_account_mapping: dict = None):
        self.user_service_account_mapping = user_service_account_mapping

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigSecurityConfigIdentityConfig()
        if Primitive.to_proto(resource.user_service_account_mapping):
            res.user_service_account_mapping = Primitive.to_proto(
                resource.user_service_account_mapping
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigSecurityConfigIdentityConfig(
            user_service_account_mapping=Primitive.from_proto(
                resource.user_service_account_mapping
            ),
        )


class ClusterConfigSecurityConfigIdentityConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigSecurityConfigIdentityConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigSecurityConfigIdentityConfig.from_proto(i) for i in resources
        ]


class ClusterConfigLifecycleConfig(object):
    def __init__(
        self,
        idle_delete_ttl: str = None,
        auto_delete_time: str = None,
        auto_delete_ttl: str = None,
        idle_start_time: str = None,
    ):
        self.idle_delete_ttl = idle_delete_ttl
        self.auto_delete_time = auto_delete_time
        self.auto_delete_ttl = auto_delete_ttl
        self.idle_start_time = idle_start_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigLifecycleConfig()
        if Primitive.to_proto(resource.idle_delete_ttl):
            res.idle_delete_ttl = Primitive.to_proto(resource.idle_delete_ttl)
        if Primitive.to_proto(resource.auto_delete_time):
            res.auto_delete_time = Primitive.to_proto(resource.auto_delete_time)
        if Primitive.to_proto(resource.auto_delete_ttl):
            res.auto_delete_ttl = Primitive.to_proto(resource.auto_delete_ttl)
        if Primitive.to_proto(resource.idle_start_time):
            res.idle_start_time = Primitive.to_proto(resource.idle_start_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigLifecycleConfig(
            idle_delete_ttl=Primitive.from_proto(resource.idle_delete_ttl),
            auto_delete_time=Primitive.from_proto(resource.auto_delete_time),
            auto_delete_ttl=Primitive.from_proto(resource.auto_delete_ttl),
            idle_start_time=Primitive.from_proto(resource.idle_start_time),
        )


class ClusterConfigLifecycleConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigLifecycleConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigLifecycleConfig.from_proto(i) for i in resources]


class ClusterConfigEndpointConfig(object):
    def __init__(self, http_ports: dict = None, enable_http_port_access: bool = None):
        self.http_ports = http_ports
        self.enable_http_port_access = enable_http_port_access

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigEndpointConfig()
        if Primitive.to_proto(resource.http_ports):
            res.http_ports = Primitive.to_proto(resource.http_ports)
        if Primitive.to_proto(resource.enable_http_port_access):
            res.enable_http_port_access = Primitive.to_proto(
                resource.enable_http_port_access
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigEndpointConfig(
            http_ports=Primitive.from_proto(resource.http_ports),
            enable_http_port_access=Primitive.from_proto(
                resource.enable_http_port_access
            ),
        )


class ClusterConfigEndpointConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigEndpointConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigEndpointConfig.from_proto(i) for i in resources]


class ClusterConfigGkeClusterConfig(object):
    def __init__(self, namespaced_gke_deployment_target: dict = None):
        self.namespaced_gke_deployment_target = namespaced_gke_deployment_target

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigGkeClusterConfig()
        if ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget.to_proto(
            resource.namespaced_gke_deployment_target
        ):
            res.namespaced_gke_deployment_target.CopyFrom(
                ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget.to_proto(
                    resource.namespaced_gke_deployment_target
                )
            )
        else:
            res.ClearField("namespaced_gke_deployment_target")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigGkeClusterConfig(
            namespaced_gke_deployment_target=ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget.from_proto(
                resource.namespaced_gke_deployment_target
            ),
        )


class ClusterConfigGkeClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigGkeClusterConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigGkeClusterConfig.from_proto(i) for i in resources]


class ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(object):
    def __init__(self, target_gke_cluster: str = None, cluster_namespace: str = None):
        self.target_gke_cluster = target_gke_cluster
        self.cluster_namespace = cluster_namespace

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget()
        )
        if Primitive.to_proto(resource.target_gke_cluster):
            res.target_gke_cluster = Primitive.to_proto(resource.target_gke_cluster)
        if Primitive.to_proto(resource.cluster_namespace):
            res.cluster_namespace = Primitive.to_proto(resource.cluster_namespace)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(
            target_gke_cluster=Primitive.from_proto(resource.target_gke_cluster),
            cluster_namespace=Primitive.from_proto(resource.cluster_namespace),
        )


class ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget.from_proto(i)
            for i in resources
        ]


class ClusterConfigMetastoreConfig(object):
    def __init__(self, dataproc_metastore_service: str = None):
        self.dataproc_metastore_service = dataproc_metastore_service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigMetastoreConfig()
        if Primitive.to_proto(resource.dataproc_metastore_service):
            res.dataproc_metastore_service = Primitive.to_proto(
                resource.dataproc_metastore_service
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigMetastoreConfig(
            dataproc_metastore_service=Primitive.from_proto(
                resource.dataproc_metastore_service
            ),
        )


class ClusterConfigMetastoreConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigMetastoreConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigMetastoreConfig.from_proto(i) for i in resources]


class ClusterConfigDataprocMetricConfig(object):
    def __init__(self, metrics: list = None):
        self.metrics = metrics

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigDataprocMetricConfig()
        if ClusterConfigDataprocMetricConfigMetricsArray.to_proto(resource.metrics):
            res.metrics.extend(
                ClusterConfigDataprocMetricConfigMetricsArray.to_proto(resource.metrics)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigDataprocMetricConfig(
            metrics=ClusterConfigDataprocMetricConfigMetricsArray.from_proto(
                resource.metrics
            ),
        )


class ClusterConfigDataprocMetricConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigDataprocMetricConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConfigDataprocMetricConfig.from_proto(i) for i in resources]


class ClusterConfigDataprocMetricConfigMetrics(object):
    def __init__(self, metric_source: str = None, metric_overrides: list = None):
        self.metric_source = metric_source
        self.metric_overrides = metric_overrides

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterConfigDataprocMetricConfigMetrics()
        if ClusterConfigDataprocMetricConfigMetricsMetricSourceEnum.to_proto(
            resource.metric_source
        ):
            res.metric_source = (
                ClusterConfigDataprocMetricConfigMetricsMetricSourceEnum.to_proto(
                    resource.metric_source
                )
            )
        if Primitive.to_proto(resource.metric_overrides):
            res.metric_overrides.extend(Primitive.to_proto(resource.metric_overrides))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConfigDataprocMetricConfigMetrics(
            metric_source=ClusterConfigDataprocMetricConfigMetricsMetricSourceEnum.from_proto(
                resource.metric_source
            ),
            metric_overrides=Primitive.from_proto(resource.metric_overrides),
        )


class ClusterConfigDataprocMetricConfigMetricsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConfigDataprocMetricConfigMetrics.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterConfigDataprocMetricConfigMetrics.from_proto(i) for i in resources
        ]


class ClusterStatus(object):
    def __init__(
        self,
        state: str = None,
        detail: str = None,
        state_start_time: str = None,
        substate: str = None,
    ):
        self.state = state
        self.detail = detail
        self.state_start_time = state_start_time
        self.substate = substate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterStatus()
        if ClusterStatusStateEnum.to_proto(resource.state):
            res.state = ClusterStatusStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.detail):
            res.detail = Primitive.to_proto(resource.detail)
        if Primitive.to_proto(resource.state_start_time):
            res.state_start_time = Primitive.to_proto(resource.state_start_time)
        if ClusterStatusSubstateEnum.to_proto(resource.substate):
            res.substate = ClusterStatusSubstateEnum.to_proto(resource.substate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterStatus(
            state=ClusterStatusStateEnum.from_proto(resource.state),
            detail=Primitive.from_proto(resource.detail),
            state_start_time=Primitive.from_proto(resource.state_start_time),
            substate=ClusterStatusSubstateEnum.from_proto(resource.substate),
        )


class ClusterStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterStatus.from_proto(i) for i in resources]


class ClusterStatusHistory(object):
    def __init__(
        self,
        state: str = None,
        detail: str = None,
        state_start_time: str = None,
        substate: str = None,
    ):
        self.state = state
        self.detail = detail
        self.state_start_time = state_start_time
        self.substate = substate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterStatusHistory()
        if ClusterStatusHistoryStateEnum.to_proto(resource.state):
            res.state = ClusterStatusHistoryStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.detail):
            res.detail = Primitive.to_proto(resource.detail)
        if Primitive.to_proto(resource.state_start_time):
            res.state_start_time = Primitive.to_proto(resource.state_start_time)
        if ClusterStatusHistorySubstateEnum.to_proto(resource.substate):
            res.substate = ClusterStatusHistorySubstateEnum.to_proto(resource.substate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterStatusHistory(
            state=ClusterStatusHistoryStateEnum.from_proto(resource.state),
            detail=Primitive.from_proto(resource.detail),
            state_start_time=Primitive.from_proto(resource.state_start_time),
            substate=ClusterStatusHistorySubstateEnum.from_proto(resource.substate),
        )


class ClusterStatusHistoryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterStatusHistory.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterStatusHistory.from_proto(i) for i in resources]


class ClusterMetrics(object):
    def __init__(self, hdfs_metrics: dict = None, yarn_metrics: dict = None):
        self.hdfs_metrics = hdfs_metrics
        self.yarn_metrics = yarn_metrics

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterMetrics()
        if Primitive.to_proto(resource.hdfs_metrics):
            res.hdfs_metrics = Primitive.to_proto(resource.hdfs_metrics)
        if Primitive.to_proto(resource.yarn_metrics):
            res.yarn_metrics = Primitive.to_proto(resource.yarn_metrics)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMetrics(
            hdfs_metrics=Primitive.from_proto(resource.hdfs_metrics),
            yarn_metrics=Primitive.from_proto(resource.yarn_metrics),
        )


class ClusterMetricsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMetrics.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterMetrics.from_proto(i) for i in resources]


class ClusterVirtualClusterConfig(object):
    def __init__(
        self,
        staging_bucket: str = None,
        kubernetes_cluster_config: dict = None,
        auxiliary_services_config: dict = None,
    ):
        self.staging_bucket = staging_bucket
        self.kubernetes_cluster_config = kubernetes_cluster_config
        self.auxiliary_services_config = auxiliary_services_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocAlphaClusterVirtualClusterConfig()
        if Primitive.to_proto(resource.staging_bucket):
            res.staging_bucket = Primitive.to_proto(resource.staging_bucket)
        if ClusterVirtualClusterConfigKubernetesClusterConfig.to_proto(
            resource.kubernetes_cluster_config
        ):
            res.kubernetes_cluster_config.CopyFrom(
                ClusterVirtualClusterConfigKubernetesClusterConfig.to_proto(
                    resource.kubernetes_cluster_config
                )
            )
        else:
            res.ClearField("kubernetes_cluster_config")
        if ClusterVirtualClusterConfigAuxiliaryServicesConfig.to_proto(
            resource.auxiliary_services_config
        ):
            res.auxiliary_services_config.CopyFrom(
                ClusterVirtualClusterConfigAuxiliaryServicesConfig.to_proto(
                    resource.auxiliary_services_config
                )
            )
        else:
            res.ClearField("auxiliary_services_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVirtualClusterConfig(
            staging_bucket=Primitive.from_proto(resource.staging_bucket),
            kubernetes_cluster_config=ClusterVirtualClusterConfigKubernetesClusterConfig.from_proto(
                resource.kubernetes_cluster_config
            ),
            auxiliary_services_config=ClusterVirtualClusterConfigAuxiliaryServicesConfig.from_proto(
                resource.auxiliary_services_config
            ),
        )


class ClusterVirtualClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterVirtualClusterConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterVirtualClusterConfig.from_proto(i) for i in resources]


class ClusterVirtualClusterConfigKubernetesClusterConfig(object):
    def __init__(
        self,
        kubernetes_namespace: str = None,
        gke_cluster_config: dict = None,
        kubernetes_software_config: dict = None,
    ):
        self.kubernetes_namespace = kubernetes_namespace
        self.gke_cluster_config = gke_cluster_config
        self.kubernetes_software_config = kubernetes_software_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfig()
        )
        if Primitive.to_proto(resource.kubernetes_namespace):
            res.kubernetes_namespace = Primitive.to_proto(resource.kubernetes_namespace)
        if ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig.to_proto(
            resource.gke_cluster_config
        ):
            res.gke_cluster_config.CopyFrom(
                ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig.to_proto(
                    resource.gke_cluster_config
                )
            )
        else:
            res.ClearField("gke_cluster_config")
        if ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig.to_proto(
            resource.kubernetes_software_config
        ):
            res.kubernetes_software_config.CopyFrom(
                ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig.to_proto(
                    resource.kubernetes_software_config
                )
            )
        else:
            res.ClearField("kubernetes_software_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVirtualClusterConfigKubernetesClusterConfig(
            kubernetes_namespace=Primitive.from_proto(resource.kubernetes_namespace),
            gke_cluster_config=ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig.from_proto(
                resource.gke_cluster_config
            ),
            kubernetes_software_config=ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig.from_proto(
                resource.kubernetes_software_config
            ),
        )


class ClusterVirtualClusterConfigKubernetesClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfig.from_proto(i)
            for i in resources
        ]


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig(object):
    def __init__(self, gke_cluster_target: str = None, node_pool_target: list = None):
        self.gke_cluster_target = gke_cluster_target
        self.node_pool_target = node_pool_target

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig()
        )
        if Primitive.to_proto(resource.gke_cluster_target):
            res.gke_cluster_target = Primitive.to_proto(resource.gke_cluster_target)
        if ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetArray.to_proto(
            resource.node_pool_target
        ):
            res.node_pool_target.extend(
                ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetArray.to_proto(
                    resource.node_pool_target
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig(
            gke_cluster_target=Primitive.from_proto(resource.gke_cluster_target),
            node_pool_target=ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetArray.from_proto(
                resource.node_pool_target
            ),
        )


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig.from_proto(
                i
            )
            for i in resources
        ]


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget(
    object
):
    def __init__(
        self, node_pool: str = None, roles: list = None, node_pool_config: dict = None
    ):
        self.node_pool = node_pool
        self.roles = roles
        self.node_pool_config = node_pool_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget()
        )
        if Primitive.to_proto(resource.node_pool):
            res.node_pool = Primitive.to_proto(resource.node_pool)
        if ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnumArray.to_proto(
            resource.roles
        ):
            res.roles.extend(
                ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnumArray.to_proto(
                    resource.roles
                )
            )
        if ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig.to_proto(
            resource.node_pool_config
        ):
            res.node_pool_config.CopyFrom(
                ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig.to_proto(
                    resource.node_pool_config
                )
            )
        else:
            res.ClearField("node_pool_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget(
            node_pool=Primitive.from_proto(resource.node_pool),
            roles=ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnumArray.from_proto(
                resource.roles
            ),
            node_pool_config=ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig.from_proto(
                resource.node_pool_config
            ),
        )


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.from_proto(
                i
            )
            for i in resources
        ]


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig(
    object
):
    def __init__(
        self, config: dict = None, locations: list = None, autoscaling: dict = None
    ):
        self.config = config
        self.locations = locations
        self.autoscaling = autoscaling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig()
        )
        if ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig.to_proto(
            resource.config
        ):
            res.config.CopyFrom(
                ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig.to_proto(
                    resource.config
                )
            )
        else:
            res.ClearField("config")
        if Primitive.to_proto(resource.locations):
            res.locations.extend(Primitive.to_proto(resource.locations))
        if ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling.to_proto(
            resource.autoscaling
        ):
            res.autoscaling.CopyFrom(
                ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling.to_proto(
                    resource.autoscaling
                )
            )
        else:
            res.ClearField("autoscaling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig(
            config=ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig.from_proto(
                resource.config
            ),
            locations=Primitive.from_proto(resource.locations),
            autoscaling=ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling.from_proto(
                resource.autoscaling
            ),
        )


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig.from_proto(
                i
            )
            for i in resources
        ]


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig(
    object
):
    def __init__(
        self,
        machine_type: str = None,
        local_ssd_count: int = None,
        preemptible: bool = None,
        accelerators: list = None,
        min_cpu_platform: str = None,
        boot_disk_kms_key: str = None,
        ephemeral_storage_config: dict = None,
        spot: bool = None,
    ):
        self.machine_type = machine_type
        self.local_ssd_count = local_ssd_count
        self.preemptible = preemptible
        self.accelerators = accelerators
        self.min_cpu_platform = min_cpu_platform
        self.boot_disk_kms_key = boot_disk_kms_key
        self.ephemeral_storage_config = ephemeral_storage_config
        self.spot = spot

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig()
        )
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if Primitive.to_proto(resource.local_ssd_count):
            res.local_ssd_count = Primitive.to_proto(resource.local_ssd_count)
        if Primitive.to_proto(resource.preemptible):
            res.preemptible = Primitive.to_proto(resource.preemptible)
        if ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAcceleratorsArray.to_proto(
            resource.accelerators
        ):
            res.accelerators.extend(
                ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAcceleratorsArray.to_proto(
                    resource.accelerators
                )
            )
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        if Primitive.to_proto(resource.boot_disk_kms_key):
            res.boot_disk_kms_key = Primitive.to_proto(resource.boot_disk_kms_key)
        if ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig.to_proto(
            resource.ephemeral_storage_config
        ):
            res.ephemeral_storage_config.CopyFrom(
                ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig.to_proto(
                    resource.ephemeral_storage_config
                )
            )
        else:
            res.ClearField("ephemeral_storage_config")
        if Primitive.to_proto(resource.spot):
            res.spot = Primitive.to_proto(resource.spot)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig(
            machine_type=Primitive.from_proto(resource.machine_type),
            local_ssd_count=Primitive.from_proto(resource.local_ssd_count),
            preemptible=Primitive.from_proto(resource.preemptible),
            accelerators=ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAcceleratorsArray.from_proto(
                resource.accelerators
            ),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
            boot_disk_kms_key=Primitive.from_proto(resource.boot_disk_kms_key),
            ephemeral_storage_config=ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig.from_proto(
                resource.ephemeral_storage_config
            ),
            spot=Primitive.from_proto(resource.spot),
        )


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig.from_proto(
                i
            )
            for i in resources
        ]


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators(
    object
):
    def __init__(
        self,
        accelerator_count: int = None,
        accelerator_type: str = None,
        gpu_partition_size: str = None,
    ):
        self.accelerator_count = accelerator_count
        self.accelerator_type = accelerator_type
        self.gpu_partition_size = gpu_partition_size

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators()
        )
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        if Primitive.to_proto(resource.gpu_partition_size):
            res.gpu_partition_size = Primitive.to_proto(resource.gpu_partition_size)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators(
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
            gpu_partition_size=Primitive.from_proto(resource.gpu_partition_size),
        )


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAcceleratorsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators.from_proto(
                i
            )
            for i in resources
        ]


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig(
    object
):
    def __init__(self, local_ssd_count: int = None):
        self.local_ssd_count = local_ssd_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig()
        )
        if Primitive.to_proto(resource.local_ssd_count):
            res.local_ssd_count = Primitive.to_proto(resource.local_ssd_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig(
            local_ssd_count=Primitive.from_proto(resource.local_ssd_count),
        )


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig.from_proto(
                i
            )
            for i in resources
        ]


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling(
    object
):
    def __init__(self, min_node_count: int = None, max_node_count: int = None):
        self.min_node_count = min_node_count
        self.max_node_count = max_node_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling()
        )
        if Primitive.to_proto(resource.min_node_count):
            res.min_node_count = Primitive.to_proto(resource.min_node_count)
        if Primitive.to_proto(resource.max_node_count):
            res.max_node_count = Primitive.to_proto(resource.max_node_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling(
            min_node_count=Primitive.from_proto(resource.min_node_count),
            max_node_count=Primitive.from_proto(resource.max_node_count),
        )


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscalingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling.from_proto(
                i
            )
            for i in resources
        ]


class ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig(
    object
):
    def __init__(self, component_version: dict = None, properties: dict = None):
        self.component_version = component_version
        self.properties = properties

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig()
        )
        if Primitive.to_proto(resource.component_version):
            res.component_version = Primitive.to_proto(resource.component_version)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig(
                component_version=Primitive.from_proto(resource.component_version),
                properties=Primitive.from_proto(resource.properties),
            )
        )


class ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig.from_proto(
                i
            )
            for i in resources
        ]


class ClusterVirtualClusterConfigAuxiliaryServicesConfig(object):
    def __init__(
        self, metastore_config: dict = None, spark_history_server_config: dict = None
    ):
        self.metastore_config = metastore_config
        self.spark_history_server_config = spark_history_server_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfig()
        )
        if ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig.to_proto(
            resource.metastore_config
        ):
            res.metastore_config.CopyFrom(
                ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig.to_proto(
                    resource.metastore_config
                )
            )
        else:
            res.ClearField("metastore_config")
        if ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig.to_proto(
            resource.spark_history_server_config
        ):
            res.spark_history_server_config.CopyFrom(
                ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig.to_proto(
                    resource.spark_history_server_config
                )
            )
        else:
            res.ClearField("spark_history_server_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVirtualClusterConfigAuxiliaryServicesConfig(
            metastore_config=ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig.from_proto(
                resource.metastore_config
            ),
            spark_history_server_config=ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig.from_proto(
                resource.spark_history_server_config
            ),
        )


class ClusterVirtualClusterConfigAuxiliaryServicesConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigAuxiliaryServicesConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigAuxiliaryServicesConfig.from_proto(i)
            for i in resources
        ]


class ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig(object):
    def __init__(self, dataproc_metastore_service: str = None):
        self.dataproc_metastore_service = dataproc_metastore_service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig()
        )
        if Primitive.to_proto(resource.dataproc_metastore_service):
            res.dataproc_metastore_service = Primitive.to_proto(
                resource.dataproc_metastore_service
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig(
            dataproc_metastore_service=Primitive.from_proto(
                resource.dataproc_metastore_service
            ),
        )


class ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig.from_proto(
                i
            )
            for i in resources
        ]


class ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig(
    object
):
    def __init__(self, dataproc_cluster: str = None):
        self.dataproc_cluster = dataproc_cluster

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig()
        )
        if Primitive.to_proto(resource.dataproc_cluster):
            res.dataproc_cluster = Primitive.to_proto(resource.dataproc_cluster)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig(
                dataproc_cluster=Primitive.from_proto(resource.dataproc_cluster),
            )
        )


class ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig.from_proto(
                i
            )
            for i in resources
        ]


class ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.Value(
            "DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.Name(
            resource
        )[
            len(
                "DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"
            ) :
        ]


class ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.Value(
            "DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.Name(
            resource
        )[
            len(
                "DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"
            ) :
        ]


class ClusterConfigMasterConfigPreemptibilityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            cluster_pb2.DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum.Value(
                "DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum%s" % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            cluster_pb2.DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum.Name(
                resource
            )[len("DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum") :]
        )


class ClusterConfigWorkerConfigPreemptibilityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            cluster_pb2.DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum.Value(
                "DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum%s" % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            cluster_pb2.DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum.Name(
                resource
            )[len("DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum") :]
        )


class ClusterConfigSecondaryWorkerConfigPreemptibilityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum.Value(
            "DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum.Name(
            resource
        )[
            len("DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum") :
        ]


class ClusterConfigSoftwareConfigOptionalComponentsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum.Value(
            "DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum.Name(
            resource
        )[
            len("DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum") :
        ]


class ClusterConfigDataprocMetricConfigMetricsMetricSourceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum.Value(
            "DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum.Name(
            resource
        )[
            len(
                "DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum"
            ) :
        ]


class ClusterStatusStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterStatusStateEnum.Value(
            "DataprocAlphaClusterStatusStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterStatusStateEnum.Name(resource)[
            len("DataprocAlphaClusterStatusStateEnum") :
        ]


class ClusterStatusSubstateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterStatusSubstateEnum.Value(
            "DataprocAlphaClusterStatusSubstateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterStatusSubstateEnum.Name(resource)[
            len("DataprocAlphaClusterStatusSubstateEnum") :
        ]


class ClusterStatusHistoryStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterStatusHistoryStateEnum.Value(
            "DataprocAlphaClusterStatusHistoryStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterStatusHistoryStateEnum.Name(resource)[
            len("DataprocAlphaClusterStatusHistoryStateEnum") :
        ]


class ClusterStatusHistorySubstateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterStatusHistorySubstateEnum.Value(
            "DataprocAlphaClusterStatusHistorySubstateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterStatusHistorySubstateEnum.Name(resource)[
            len("DataprocAlphaClusterStatusHistorySubstateEnum") :
        ]


class ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum.Value(
            "DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum.Name(
            resource
        )[
            len(
                "DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum"
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
