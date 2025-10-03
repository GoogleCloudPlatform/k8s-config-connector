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
from google3.cloud.graphite.mmv2.services.google.dataproc import workflow_template_pb2
from google3.cloud.graphite.mmv2.services.google.dataproc import (
    workflow_template_pb2_grpc,
)

from typing import List


class WorkflowTemplate(object):
    def __init__(
        self,
        name: str = None,
        version: int = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        placement: dict = None,
        jobs: list = None,
        parameters: list = None,
        dag_timeout: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.labels = labels
        self.placement = placement
        self.jobs = jobs
        self.parameters = parameters
        self.dag_timeout = dag_timeout
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = workflow_template_pb2_grpc.DataprocAlphaWorkflowTemplateServiceStub(
            channel.Channel()
        )
        request = workflow_template_pb2.ApplyDataprocAlphaWorkflowTemplateRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if WorkflowTemplatePlacement.to_proto(self.placement):
            request.resource.placement.CopyFrom(
                WorkflowTemplatePlacement.to_proto(self.placement)
            )
        else:
            request.resource.ClearField("placement")
        if WorkflowTemplateJobsArray.to_proto(self.jobs):
            request.resource.jobs.extend(WorkflowTemplateJobsArray.to_proto(self.jobs))
        if WorkflowTemplateParametersArray.to_proto(self.parameters):
            request.resource.parameters.extend(
                WorkflowTemplateParametersArray.to_proto(self.parameters)
            )
        if Primitive.to_proto(self.dag_timeout):
            request.resource.dag_timeout = Primitive.to_proto(self.dag_timeout)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDataprocAlphaWorkflowTemplate(request)
        self.name = Primitive.from_proto(response.name)
        self.version = Primitive.from_proto(response.version)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.placement = WorkflowTemplatePlacement.from_proto(response.placement)
        self.jobs = WorkflowTemplateJobsArray.from_proto(response.jobs)
        self.parameters = WorkflowTemplateParametersArray.from_proto(
            response.parameters
        )
        self.dag_timeout = Primitive.from_proto(response.dag_timeout)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = workflow_template_pb2_grpc.DataprocAlphaWorkflowTemplateServiceStub(
            channel.Channel()
        )
        request = workflow_template_pb2.DeleteDataprocAlphaWorkflowTemplateRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if WorkflowTemplatePlacement.to_proto(self.placement):
            request.resource.placement.CopyFrom(
                WorkflowTemplatePlacement.to_proto(self.placement)
            )
        else:
            request.resource.ClearField("placement")
        if WorkflowTemplateJobsArray.to_proto(self.jobs):
            request.resource.jobs.extend(WorkflowTemplateJobsArray.to_proto(self.jobs))
        if WorkflowTemplateParametersArray.to_proto(self.parameters):
            request.resource.parameters.extend(
                WorkflowTemplateParametersArray.to_proto(self.parameters)
            )
        if Primitive.to_proto(self.dag_timeout):
            request.resource.dag_timeout = Primitive.to_proto(self.dag_timeout)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteDataprocAlphaWorkflowTemplate(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = workflow_template_pb2_grpc.DataprocAlphaWorkflowTemplateServiceStub(
            channel.Channel()
        )
        request = workflow_template_pb2.ListDataprocAlphaWorkflowTemplateRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListDataprocAlphaWorkflowTemplate(request).items

    def to_proto(self):
        resource = workflow_template_pb2.DataprocAlphaWorkflowTemplate()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if WorkflowTemplatePlacement.to_proto(self.placement):
            resource.placement.CopyFrom(
                WorkflowTemplatePlacement.to_proto(self.placement)
            )
        else:
            resource.ClearField("placement")
        if WorkflowTemplateJobsArray.to_proto(self.jobs):
            resource.jobs.extend(WorkflowTemplateJobsArray.to_proto(self.jobs))
        if WorkflowTemplateParametersArray.to_proto(self.parameters):
            resource.parameters.extend(
                WorkflowTemplateParametersArray.to_proto(self.parameters)
            )
        if Primitive.to_proto(self.dag_timeout):
            resource.dag_timeout = Primitive.to_proto(self.dag_timeout)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class WorkflowTemplatePlacement(object):
    def __init__(self, managed_cluster: dict = None, cluster_selector: dict = None):
        self.managed_cluster = managed_cluster
        self.cluster_selector = cluster_selector

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacement()
        if WorkflowTemplatePlacementManagedCluster.to_proto(resource.managed_cluster):
            res.managed_cluster.CopyFrom(
                WorkflowTemplatePlacementManagedCluster.to_proto(
                    resource.managed_cluster
                )
            )
        else:
            res.ClearField("managed_cluster")
        if WorkflowTemplatePlacementClusterSelector.to_proto(resource.cluster_selector):
            res.cluster_selector.CopyFrom(
                WorkflowTemplatePlacementClusterSelector.to_proto(
                    resource.cluster_selector
                )
            )
        else:
            res.ClearField("cluster_selector")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacement(
            managed_cluster=WorkflowTemplatePlacementManagedCluster.from_proto(
                resource.managed_cluster
            ),
            cluster_selector=WorkflowTemplatePlacementClusterSelector.from_proto(
                resource.cluster_selector
            ),
        )


class WorkflowTemplatePlacementArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplatePlacement.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplatePlacement.from_proto(i) for i in resources]


class WorkflowTemplatePlacementManagedCluster(object):
    def __init__(
        self, cluster_name: str = None, config: dict = None, labels: dict = None
    ):
        self.cluster_name = cluster_name
        self.config = config
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedCluster()
        )
        if Primitive.to_proto(resource.cluster_name):
            res.cluster_name = Primitive.to_proto(resource.cluster_name)
        if WorkflowTemplatePlacementManagedClusterConfig.to_proto(resource.config):
            res.config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfig.to_proto(resource.config)
            )
        else:
            res.ClearField("config")
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedCluster(
            cluster_name=Primitive.from_proto(resource.cluster_name),
            config=WorkflowTemplatePlacementManagedClusterConfig.from_proto(
                resource.config
            ),
            labels=Primitive.from_proto(resource.labels),
        )


class WorkflowTemplatePlacementManagedClusterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplatePlacementManagedCluster.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedCluster.from_proto(i) for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfig(object):
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

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfig()
        )
        if Primitive.to_proto(resource.staging_bucket):
            res.staging_bucket = Primitive.to_proto(resource.staging_bucket)
        if Primitive.to_proto(resource.temp_bucket):
            res.temp_bucket = Primitive.to_proto(resource.temp_bucket)
        if WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig.to_proto(
            resource.gce_cluster_config
        ):
            res.gce_cluster_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig.to_proto(
                    resource.gce_cluster_config
                )
            )
        else:
            res.ClearField("gce_cluster_config")
        if WorkflowTemplatePlacementManagedClusterConfigMasterConfig.to_proto(
            resource.master_config
        ):
            res.master_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigMasterConfig.to_proto(
                    resource.master_config
                )
            )
        else:
            res.ClearField("master_config")
        if WorkflowTemplatePlacementManagedClusterConfigWorkerConfig.to_proto(
            resource.worker_config
        ):
            res.worker_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigWorkerConfig.to_proto(
                    resource.worker_config
                )
            )
        else:
            res.ClearField("worker_config")
        if WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig.to_proto(
            resource.secondary_worker_config
        ):
            res.secondary_worker_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig.to_proto(
                    resource.secondary_worker_config
                )
            )
        else:
            res.ClearField("secondary_worker_config")
        if WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig.to_proto(
            resource.software_config
        ):
            res.software_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig.to_proto(
                    resource.software_config
                )
            )
        else:
            res.ClearField("software_config")
        if WorkflowTemplatePlacementManagedClusterConfigInitializationActionsArray.to_proto(
            resource.initialization_actions
        ):
            res.initialization_actions.extend(
                WorkflowTemplatePlacementManagedClusterConfigInitializationActionsArray.to_proto(
                    resource.initialization_actions
                )
            )
        if WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig.to_proto(
            resource.encryption_config
        ):
            res.encryption_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig.to_proto(
                    resource.encryption_config
                )
            )
        else:
            res.ClearField("encryption_config")
        if WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig.to_proto(
            resource.autoscaling_config
        ):
            res.autoscaling_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig.to_proto(
                    resource.autoscaling_config
                )
            )
        else:
            res.ClearField("autoscaling_config")
        if WorkflowTemplatePlacementManagedClusterConfigSecurityConfig.to_proto(
            resource.security_config
        ):
            res.security_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigSecurityConfig.to_proto(
                    resource.security_config
                )
            )
        else:
            res.ClearField("security_config")
        if WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig.to_proto(
            resource.lifecycle_config
        ):
            res.lifecycle_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig.to_proto(
                    resource.lifecycle_config
                )
            )
        else:
            res.ClearField("lifecycle_config")
        if WorkflowTemplatePlacementManagedClusterConfigEndpointConfig.to_proto(
            resource.endpoint_config
        ):
            res.endpoint_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigEndpointConfig.to_proto(
                    resource.endpoint_config
                )
            )
        else:
            res.ClearField("endpoint_config")
        if WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig.to_proto(
            resource.gke_cluster_config
        ):
            res.gke_cluster_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig.to_proto(
                    resource.gke_cluster_config
                )
            )
        else:
            res.ClearField("gke_cluster_config")
        if WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig.to_proto(
            resource.metastore_config
        ):
            res.metastore_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig.to_proto(
                    resource.metastore_config
                )
            )
        else:
            res.ClearField("metastore_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfig(
            staging_bucket=Primitive.from_proto(resource.staging_bucket),
            temp_bucket=Primitive.from_proto(resource.temp_bucket),
            gce_cluster_config=WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig.from_proto(
                resource.gce_cluster_config
            ),
            master_config=WorkflowTemplatePlacementManagedClusterConfigMasterConfig.from_proto(
                resource.master_config
            ),
            worker_config=WorkflowTemplatePlacementManagedClusterConfigWorkerConfig.from_proto(
                resource.worker_config
            ),
            secondary_worker_config=WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig.from_proto(
                resource.secondary_worker_config
            ),
            software_config=WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig.from_proto(
                resource.software_config
            ),
            initialization_actions=WorkflowTemplatePlacementManagedClusterConfigInitializationActionsArray.from_proto(
                resource.initialization_actions
            ),
            encryption_config=WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig.from_proto(
                resource.encryption_config
            ),
            autoscaling_config=WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig.from_proto(
                resource.autoscaling_config
            ),
            security_config=WorkflowTemplatePlacementManagedClusterConfigSecurityConfig.from_proto(
                resource.security_config
            ),
            lifecycle_config=WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig.from_proto(
                resource.lifecycle_config
            ),
            endpoint_config=WorkflowTemplatePlacementManagedClusterConfigEndpointConfig.from_proto(
                resource.endpoint_config
            ),
            gke_cluster_config=WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig.from_proto(
                resource.gke_cluster_config
            ),
            metastore_config=WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig.from_proto(
                resource.metastore_config
            ),
        )


class WorkflowTemplatePlacementManagedClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig(object):
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

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig()
        )
        if Primitive.to_proto(resource.zone):
            res.zone = Primitive.to_proto(resource.zone)
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        if Primitive.to_proto(resource.subnetwork):
            res.subnetwork = Primitive.to_proto(resource.subnetwork)
        if Primitive.to_proto(resource.internal_ip_only):
            res.internal_ip_only = Primitive.to_proto(resource.internal_ip_only)
        if WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.to_proto(
            resource.private_ipv6_google_access
        ):
            res.private_ipv6_google_access = WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.to_proto(
                resource.private_ipv6_google_access
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
        if WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity.to_proto(
            resource.reservation_affinity
        ):
            res.reservation_affinity.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity.to_proto(
                    resource.reservation_affinity
                )
            )
        else:
            res.ClearField("reservation_affinity")
        if WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity.to_proto(
            resource.node_group_affinity
        ):
            res.node_group_affinity.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity.to_proto(
                    resource.node_group_affinity
                )
            )
        else:
            res.ClearField("node_group_affinity")
        if WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig.to_proto(
            resource.shielded_instance_config
        ):
            res.shielded_instance_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig.to_proto(
                    resource.shielded_instance_config
                )
            )
        else:
            res.ClearField("shielded_instance_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig(
            zone=Primitive.from_proto(resource.zone),
            network=Primitive.from_proto(resource.network),
            subnetwork=Primitive.from_proto(resource.subnetwork),
            internal_ip_only=Primitive.from_proto(resource.internal_ip_only),
            private_ipv6_google_access=WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.from_proto(
                resource.private_ipv6_google_access
            ),
            service_account=Primitive.from_proto(resource.service_account),
            service_account_scopes=Primitive.from_proto(
                resource.service_account_scopes
            ),
            tags=Primitive.from_proto(resource.tags),
            metadata=Primitive.from_proto(resource.metadata),
            reservation_affinity=WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity.from_proto(
                resource.reservation_affinity
            ),
            node_group_affinity=WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity.from_proto(
                resource.node_group_affinity
            ),
            shielded_instance_config=WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig.from_proto(
                resource.shielded_instance_config
            ),
        )


class WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity(
    object
):
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
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity()
        )
        if WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
            resource.consume_reservation_type
        ):
            res.consume_reservation_type = WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
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

        return WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity(
            consume_reservation_type=WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.from_proto(
                resource.consume_reservation_type
            ),
            key=Primitive.from_proto(resource.key),
            values=Primitive.from_proto(resource.values),
        )


class WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity(
    object
):
    def __init__(self, node_group: str = None):
        self.node_group = node_group

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity()
        )
        if Primitive.to_proto(resource.node_group):
            res.node_group = Primitive.to_proto(resource.node_group)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity(
            node_group=Primitive.from_proto(resource.node_group),
        )


class WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig(
    object
):
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
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig()
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

        return WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig(
            enable_secure_boot=Primitive.from_proto(resource.enable_secure_boot),
            enable_vtpm=Primitive.from_proto(resource.enable_vtpm),
            enable_integrity_monitoring=Primitive.from_proto(
                resource.enable_integrity_monitoring
            ),
        )


class WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigMasterConfig(object):
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

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfig()
        )
        if Primitive.to_proto(resource.num_instances):
            res.num_instances = Primitive.to_proto(resource.num_instances)
        if Primitive.to_proto(resource.instance_names):
            res.instance_names.extend(Primitive.to_proto(resource.instance_names))
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig.to_proto(
            resource.disk_config
        ):
            res.disk_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig.to_proto(
                    resource.disk_config
                )
            )
        else:
            res.ClearField("disk_config")
        if Primitive.to_proto(resource.is_preemptible):
            res.is_preemptible = Primitive.to_proto(resource.is_preemptible)
        if WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum.to_proto(
            resource.preemptibility
        ):
            res.preemptibility = WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum.to_proto(
                resource.preemptibility
            )
        if WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig.to_proto(
            resource.managed_group_config
        ):
            res.managed_group_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig.to_proto(
                    resource.managed_group_config
                )
            )
        else:
            res.ClearField("managed_group_config")
        if WorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsArray.to_proto(
            resource.accelerators
        ):
            res.accelerators.extend(
                WorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsArray.to_proto(
                    resource.accelerators
                )
            )
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigMasterConfig(
            num_instances=Primitive.from_proto(resource.num_instances),
            instance_names=Primitive.from_proto(resource.instance_names),
            image=Primitive.from_proto(resource.image),
            machine_type=Primitive.from_proto(resource.machine_type),
            disk_config=WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig.from_proto(
                resource.disk_config
            ),
            is_preemptible=Primitive.from_proto(resource.is_preemptible),
            preemptibility=WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum.from_proto(
                resource.preemptibility
            ),
            managed_group_config=WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig.from_proto(
                resource.managed_group_config
            ),
            accelerators=WorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsArray.from_proto(
                resource.accelerators
            ),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
        )


class WorkflowTemplatePlacementManagedClusterConfigMasterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigMasterConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigMasterConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig(object):
    def __init__(
        self,
        boot_disk_type: str = None,
        boot_disk_size_gb: int = None,
        num_local_ssds: int = None,
    ):
        self.boot_disk_type = boot_disk_type
        self.boot_disk_size_gb = boot_disk_size_gb
        self.num_local_ssds = num_local_ssds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig()
        )
        if Primitive.to_proto(resource.boot_disk_type):
            res.boot_disk_type = Primitive.to_proto(resource.boot_disk_type)
        if Primitive.to_proto(resource.boot_disk_size_gb):
            res.boot_disk_size_gb = Primitive.to_proto(resource.boot_disk_size_gb)
        if Primitive.to_proto(resource.num_local_ssds):
            res.num_local_ssds = Primitive.to_proto(resource.num_local_ssds)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig(
            boot_disk_type=Primitive.from_proto(resource.boot_disk_type),
            boot_disk_size_gb=Primitive.from_proto(resource.boot_disk_size_gb),
            num_local_ssds=Primitive.from_proto(resource.num_local_ssds),
        )


class WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig(
    object
):
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
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig()
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

        return (
            WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig(
                instance_template_name=Primitive.from_proto(
                    resource.instance_template_name
                ),
                instance_group_manager_name=Primitive.from_proto(
                    resource.instance_group_manager_name
                ),
            )
        )


class WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators(object):
    def __init__(self, accelerator_type: str = None, accelerator_count: int = None):
        self.accelerator_type = accelerator_type
        self.accelerator_count = accelerator_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators()
        )
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators(
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
        )


class WorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigWorkerConfig(object):
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

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfig()
        )
        if Primitive.to_proto(resource.num_instances):
            res.num_instances = Primitive.to_proto(resource.num_instances)
        if Primitive.to_proto(resource.instance_names):
            res.instance_names.extend(Primitive.to_proto(resource.instance_names))
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig.to_proto(
            resource.disk_config
        ):
            res.disk_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig.to_proto(
                    resource.disk_config
                )
            )
        else:
            res.ClearField("disk_config")
        if Primitive.to_proto(resource.is_preemptible):
            res.is_preemptible = Primitive.to_proto(resource.is_preemptible)
        if WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum.to_proto(
            resource.preemptibility
        ):
            res.preemptibility = WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum.to_proto(
                resource.preemptibility
            )
        if WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig.to_proto(
            resource.managed_group_config
        ):
            res.managed_group_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig.to_proto(
                    resource.managed_group_config
                )
            )
        else:
            res.ClearField("managed_group_config")
        if WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsArray.to_proto(
            resource.accelerators
        ):
            res.accelerators.extend(
                WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsArray.to_proto(
                    resource.accelerators
                )
            )
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigWorkerConfig(
            num_instances=Primitive.from_proto(resource.num_instances),
            instance_names=Primitive.from_proto(resource.instance_names),
            image=Primitive.from_proto(resource.image),
            machine_type=Primitive.from_proto(resource.machine_type),
            disk_config=WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig.from_proto(
                resource.disk_config
            ),
            is_preemptible=Primitive.from_proto(resource.is_preemptible),
            preemptibility=WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum.from_proto(
                resource.preemptibility
            ),
            managed_group_config=WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig.from_proto(
                resource.managed_group_config
            ),
            accelerators=WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsArray.from_proto(
                resource.accelerators
            ),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
        )


class WorkflowTemplatePlacementManagedClusterConfigWorkerConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigWorkerConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigWorkerConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig(object):
    def __init__(
        self,
        boot_disk_type: str = None,
        boot_disk_size_gb: int = None,
        num_local_ssds: int = None,
    ):
        self.boot_disk_type = boot_disk_type
        self.boot_disk_size_gb = boot_disk_size_gb
        self.num_local_ssds = num_local_ssds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig()
        )
        if Primitive.to_proto(resource.boot_disk_type):
            res.boot_disk_type = Primitive.to_proto(resource.boot_disk_type)
        if Primitive.to_proto(resource.boot_disk_size_gb):
            res.boot_disk_size_gb = Primitive.to_proto(resource.boot_disk_size_gb)
        if Primitive.to_proto(resource.num_local_ssds):
            res.num_local_ssds = Primitive.to_proto(resource.num_local_ssds)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig(
            boot_disk_type=Primitive.from_proto(resource.boot_disk_type),
            boot_disk_size_gb=Primitive.from_proto(resource.boot_disk_size_gb),
            num_local_ssds=Primitive.from_proto(resource.num_local_ssds),
        )


class WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig(
    object
):
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
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig()
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

        return (
            WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig(
                instance_template_name=Primitive.from_proto(
                    resource.instance_template_name
                ),
                instance_group_manager_name=Primitive.from_proto(
                    resource.instance_group_manager_name
                ),
            )
        )


class WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators(object):
    def __init__(self, accelerator_type: str = None, accelerator_count: int = None):
        self.accelerator_type = accelerator_type
        self.accelerator_count = accelerator_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators()
        )
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators(
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
        )


class WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig(object):
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

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig()
        )
        if Primitive.to_proto(resource.num_instances):
            res.num_instances = Primitive.to_proto(resource.num_instances)
        if Primitive.to_proto(resource.instance_names):
            res.instance_names.extend(Primitive.to_proto(resource.instance_names))
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig.to_proto(
            resource.disk_config
        ):
            res.disk_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig.to_proto(
                    resource.disk_config
                )
            )
        else:
            res.ClearField("disk_config")
        if Primitive.to_proto(resource.is_preemptible):
            res.is_preemptible = Primitive.to_proto(resource.is_preemptible)
        if WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum.to_proto(
            resource.preemptibility
        ):
            res.preemptibility = WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum.to_proto(
                resource.preemptibility
            )
        if WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig.to_proto(
            resource.managed_group_config
        ):
            res.managed_group_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig.to_proto(
                    resource.managed_group_config
                )
            )
        else:
            res.ClearField("managed_group_config")
        if WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsArray.to_proto(
            resource.accelerators
        ):
            res.accelerators.extend(
                WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsArray.to_proto(
                    resource.accelerators
                )
            )
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig(
            num_instances=Primitive.from_proto(resource.num_instances),
            instance_names=Primitive.from_proto(resource.instance_names),
            image=Primitive.from_proto(resource.image),
            machine_type=Primitive.from_proto(resource.machine_type),
            disk_config=WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig.from_proto(
                resource.disk_config
            ),
            is_preemptible=Primitive.from_proto(resource.is_preemptible),
            preemptibility=WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum.from_proto(
                resource.preemptibility
            ),
            managed_group_config=WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig.from_proto(
                resource.managed_group_config
            ),
            accelerators=WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsArray.from_proto(
                resource.accelerators
            ),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
        )


class WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig(
    object
):
    def __init__(
        self,
        boot_disk_type: str = None,
        boot_disk_size_gb: int = None,
        num_local_ssds: int = None,
    ):
        self.boot_disk_type = boot_disk_type
        self.boot_disk_size_gb = boot_disk_size_gb
        self.num_local_ssds = num_local_ssds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig()
        )
        if Primitive.to_proto(resource.boot_disk_type):
            res.boot_disk_type = Primitive.to_proto(resource.boot_disk_type)
        if Primitive.to_proto(resource.boot_disk_size_gb):
            res.boot_disk_size_gb = Primitive.to_proto(resource.boot_disk_size_gb)
        if Primitive.to_proto(resource.num_local_ssds):
            res.num_local_ssds = Primitive.to_proto(resource.num_local_ssds)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig(
            boot_disk_type=Primitive.from_proto(resource.boot_disk_type),
            boot_disk_size_gb=Primitive.from_proto(resource.boot_disk_size_gb),
            num_local_ssds=Primitive.from_proto(resource.num_local_ssds),
        )


class WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig(
    object
):
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
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig()
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

        return WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig(
            instance_template_name=Primitive.from_proto(
                resource.instance_template_name
            ),
            instance_group_manager_name=Primitive.from_proto(
                resource.instance_group_manager_name
            ),
        )


class WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators(
    object
):
    def __init__(self, accelerator_type: str = None, accelerator_count: int = None):
        self.accelerator_type = accelerator_type
        self.accelerator_count = accelerator_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators()
        )
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators(
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
        )


class WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig(object):
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

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig()
        )
        if Primitive.to_proto(resource.image_version):
            res.image_version = Primitive.to_proto(resource.image_version)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnumArray.to_proto(
            resource.optional_components
        ):
            res.optional_components.extend(
                WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnumArray.to_proto(
                    resource.optional_components
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig(
            image_version=Primitive.from_proto(resource.image_version),
            properties=Primitive.from_proto(resource.properties),
            optional_components=WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnumArray.from_proto(
                resource.optional_components
            ),
        )


class WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigInitializationActions(object):
    def __init__(self, executable_file: str = None, execution_timeout: str = None):
        self.executable_file = executable_file
        self.execution_timeout = execution_timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigInitializationActions()
        )
        if Primitive.to_proto(resource.executable_file):
            res.executable_file = Primitive.to_proto(resource.executable_file)
        if Primitive.to_proto(resource.execution_timeout):
            res.execution_timeout = Primitive.to_proto(resource.execution_timeout)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigInitializationActions(
            executable_file=Primitive.from_proto(resource.executable_file),
            execution_timeout=Primitive.from_proto(resource.execution_timeout),
        )


class WorkflowTemplatePlacementManagedClusterConfigInitializationActionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigInitializationActions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigInitializationActions.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig(object):
    def __init__(self, gce_pd_kms_key_name: str = None):
        self.gce_pd_kms_key_name = gce_pd_kms_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig()
        )
        if Primitive.to_proto(resource.gce_pd_kms_key_name):
            res.gce_pd_kms_key_name = Primitive.to_proto(resource.gce_pd_kms_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig(
            gce_pd_kms_key_name=Primitive.from_proto(resource.gce_pd_kms_key_name),
        )


class WorkflowTemplatePlacementManagedClusterConfigEncryptionConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig(object):
    def __init__(self, policy: str = None):
        self.policy = policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig()
        )
        if Primitive.to_proto(resource.policy):
            res.policy = Primitive.to_proto(resource.policy)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig(
            policy=Primitive.from_proto(resource.policy),
        )


class WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigSecurityConfig(object):
    def __init__(self, kerberos_config: dict = None):
        self.kerberos_config = kerberos_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfig()
        )
        if WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig.to_proto(
            resource.kerberos_config
        ):
            res.kerberos_config.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig.to_proto(
                    resource.kerberos_config
                )
            )
        else:
            res.ClearField("kerberos_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedClusterConfigSecurityConfig(
            kerberos_config=WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig.from_proto(
                resource.kerberos_config
            ),
        )


class WorkflowTemplatePlacementManagedClusterConfigSecurityConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecurityConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecurityConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig(object):
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

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig()
        )
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

        return (
            WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig(
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
                cross_realm_trust_kdc=Primitive.from_proto(
                    resource.cross_realm_trust_kdc
                ),
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
        )


class WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig(object):
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

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig()
        )
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

        return WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig(
            idle_delete_ttl=Primitive.from_proto(resource.idle_delete_ttl),
            auto_delete_time=Primitive.from_proto(resource.auto_delete_time),
            auto_delete_ttl=Primitive.from_proto(resource.auto_delete_ttl),
            idle_start_time=Primitive.from_proto(resource.idle_start_time),
        )


class WorkflowTemplatePlacementManagedClusterConfigLifecycleConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigEndpointConfig(object):
    def __init__(self, http_ports: dict = None, enable_http_port_access: bool = None):
        self.http_ports = http_ports
        self.enable_http_port_access = enable_http_port_access

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEndpointConfig()
        )
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

        return WorkflowTemplatePlacementManagedClusterConfigEndpointConfig(
            http_ports=Primitive.from_proto(resource.http_ports),
            enable_http_port_access=Primitive.from_proto(
                resource.enable_http_port_access
            ),
        )


class WorkflowTemplatePlacementManagedClusterConfigEndpointConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigEndpointConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigEndpointConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig(object):
    def __init__(self, namespaced_gke_deployment_target: dict = None):
        self.namespaced_gke_deployment_target = namespaced_gke_deployment_target

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig()
        )
        if WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget.to_proto(
            resource.namespaced_gke_deployment_target
        ):
            res.namespaced_gke_deployment_target.CopyFrom(
                WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget.to_proto(
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

        return WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig(
            namespaced_gke_deployment_target=WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget.from_proto(
                resource.namespaced_gke_deployment_target
            ),
        )


class WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(
    object
):
    def __init__(self, target_gke_cluster: str = None, cluster_namespace: str = None):
        self.target_gke_cluster = target_gke_cluster
        self.cluster_namespace = cluster_namespace

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget()
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

        return WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(
            target_gke_cluster=Primitive.from_proto(resource.target_gke_cluster),
            cluster_namespace=Primitive.from_proto(resource.cluster_namespace),
        )


class WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget.from_proto(
                i
            )
            for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig(object):
    def __init__(self, dataproc_metastore_service: str = None):
        self.dataproc_metastore_service = dataproc_metastore_service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig()
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

        return WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig(
            dataproc_metastore_service=Primitive.from_proto(
                resource.dataproc_metastore_service
            ),
        )


class WorkflowTemplatePlacementManagedClusterConfigMetastoreConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplatePlacementClusterSelector(object):
    def __init__(self, zone: str = None, cluster_labels: dict = None):
        self.zone = zone
        self.cluster_labels = cluster_labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementClusterSelector()
        )
        if Primitive.to_proto(resource.zone):
            res.zone = Primitive.to_proto(resource.zone)
        if Primitive.to_proto(resource.cluster_labels):
            res.cluster_labels = Primitive.to_proto(resource.cluster_labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementClusterSelector(
            zone=Primitive.from_proto(resource.zone),
            cluster_labels=Primitive.from_proto(resource.cluster_labels),
        )


class WorkflowTemplatePlacementClusterSelectorArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplatePlacementClusterSelector.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementClusterSelector.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobs(object):
    def __init__(
        self,
        step_id: str = None,
        hadoop_job: dict = None,
        spark_job: dict = None,
        pyspark_job: dict = None,
        hive_job: dict = None,
        pig_job: dict = None,
        spark_r_job: dict = None,
        spark_sql_job: dict = None,
        presto_job: dict = None,
        labels: dict = None,
        scheduling: dict = None,
        prerequisite_step_ids: list = None,
    ):
        self.step_id = step_id
        self.hadoop_job = hadoop_job
        self.spark_job = spark_job
        self.pyspark_job = pyspark_job
        self.hive_job = hive_job
        self.pig_job = pig_job
        self.spark_r_job = spark_r_job
        self.spark_sql_job = spark_sql_job
        self.presto_job = presto_job
        self.labels = labels
        self.scheduling = scheduling
        self.prerequisite_step_ids = prerequisite_step_ids

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobs()
        if Primitive.to_proto(resource.step_id):
            res.step_id = Primitive.to_proto(resource.step_id)
        if WorkflowTemplateJobsHadoopJob.to_proto(resource.hadoop_job):
            res.hadoop_job.CopyFrom(
                WorkflowTemplateJobsHadoopJob.to_proto(resource.hadoop_job)
            )
        else:
            res.ClearField("hadoop_job")
        if WorkflowTemplateJobsSparkJob.to_proto(resource.spark_job):
            res.spark_job.CopyFrom(
                WorkflowTemplateJobsSparkJob.to_proto(resource.spark_job)
            )
        else:
            res.ClearField("spark_job")
        if WorkflowTemplateJobsPysparkJob.to_proto(resource.pyspark_job):
            res.pyspark_job.CopyFrom(
                WorkflowTemplateJobsPysparkJob.to_proto(resource.pyspark_job)
            )
        else:
            res.ClearField("pyspark_job")
        if WorkflowTemplateJobsHiveJob.to_proto(resource.hive_job):
            res.hive_job.CopyFrom(
                WorkflowTemplateJobsHiveJob.to_proto(resource.hive_job)
            )
        else:
            res.ClearField("hive_job")
        if WorkflowTemplateJobsPigJob.to_proto(resource.pig_job):
            res.pig_job.CopyFrom(WorkflowTemplateJobsPigJob.to_proto(resource.pig_job))
        else:
            res.ClearField("pig_job")
        if WorkflowTemplateJobsSparkRJob.to_proto(resource.spark_r_job):
            res.spark_r_job.CopyFrom(
                WorkflowTemplateJobsSparkRJob.to_proto(resource.spark_r_job)
            )
        else:
            res.ClearField("spark_r_job")
        if WorkflowTemplateJobsSparkSqlJob.to_proto(resource.spark_sql_job):
            res.spark_sql_job.CopyFrom(
                WorkflowTemplateJobsSparkSqlJob.to_proto(resource.spark_sql_job)
            )
        else:
            res.ClearField("spark_sql_job")
        if WorkflowTemplateJobsPrestoJob.to_proto(resource.presto_job):
            res.presto_job.CopyFrom(
                WorkflowTemplateJobsPrestoJob.to_proto(resource.presto_job)
            )
        else:
            res.ClearField("presto_job")
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if WorkflowTemplateJobsScheduling.to_proto(resource.scheduling):
            res.scheduling.CopyFrom(
                WorkflowTemplateJobsScheduling.to_proto(resource.scheduling)
            )
        else:
            res.ClearField("scheduling")
        if Primitive.to_proto(resource.prerequisite_step_ids):
            res.prerequisite_step_ids.extend(
                Primitive.to_proto(resource.prerequisite_step_ids)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobs(
            step_id=Primitive.from_proto(resource.step_id),
            hadoop_job=WorkflowTemplateJobsHadoopJob.from_proto(resource.hadoop_job),
            spark_job=WorkflowTemplateJobsSparkJob.from_proto(resource.spark_job),
            pyspark_job=WorkflowTemplateJobsPysparkJob.from_proto(resource.pyspark_job),
            hive_job=WorkflowTemplateJobsHiveJob.from_proto(resource.hive_job),
            pig_job=WorkflowTemplateJobsPigJob.from_proto(resource.pig_job),
            spark_r_job=WorkflowTemplateJobsSparkRJob.from_proto(resource.spark_r_job),
            spark_sql_job=WorkflowTemplateJobsSparkSqlJob.from_proto(
                resource.spark_sql_job
            ),
            presto_job=WorkflowTemplateJobsPrestoJob.from_proto(resource.presto_job),
            labels=Primitive.from_proto(resource.labels),
            scheduling=WorkflowTemplateJobsScheduling.from_proto(resource.scheduling),
            prerequisite_step_ids=Primitive.from_proto(resource.prerequisite_step_ids),
        )


class WorkflowTemplateJobsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobs.from_proto(i) for i in resources]


class WorkflowTemplateJobsHadoopJob(object):
    def __init__(
        self,
        main_jar_file_uri: str = None,
        main_class: str = None,
        args: list = None,
        jar_file_uris: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_jar_file_uri = main_jar_file_uri
        self.main_class = main_class
        self.args = args
        self.jar_file_uris = jar_file_uris
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsHadoopJob()
        if Primitive.to_proto(resource.main_jar_file_uri):
            res.main_jar_file_uri = Primitive.to_proto(resource.main_jar_file_uri)
        if Primitive.to_proto(resource.main_class):
            res.main_class = Primitive.to_proto(resource.main_class)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if WorkflowTemplateJobsHadoopJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsHadoopJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsHadoopJob(
            main_jar_file_uri=Primitive.from_proto(resource.main_jar_file_uri),
            main_class=Primitive.from_proto(resource.main_class),
            args=Primitive.from_proto(resource.args),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
            file_uris=Primitive.from_proto(resource.file_uris),
            archive_uris=Primitive.from_proto(resource.archive_uris),
            properties=Primitive.from_proto(resource.properties),
            logging_config=WorkflowTemplateJobsHadoopJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class WorkflowTemplateJobsHadoopJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsHadoopJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsHadoopJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsHadoopJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig()
        )
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsHadoopJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class WorkflowTemplateJobsHadoopJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsHadoopJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsHadoopJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsSparkJob(object):
    def __init__(
        self,
        main_jar_file_uri: str = None,
        main_class: str = None,
        args: list = None,
        jar_file_uris: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_jar_file_uri = main_jar_file_uri
        self.main_class = main_class
        self.args = args
        self.jar_file_uris = jar_file_uris
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsSparkJob()
        if Primitive.to_proto(resource.main_jar_file_uri):
            res.main_jar_file_uri = Primitive.to_proto(resource.main_jar_file_uri)
        if Primitive.to_proto(resource.main_class):
            res.main_class = Primitive.to_proto(resource.main_class)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if WorkflowTemplateJobsSparkJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsSparkJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkJob(
            main_jar_file_uri=Primitive.from_proto(resource.main_jar_file_uri),
            main_class=Primitive.from_proto(resource.main_class),
            args=Primitive.from_proto(resource.args),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
            file_uris=Primitive.from_proto(resource.file_uris),
            archive_uris=Primitive.from_proto(resource.archive_uris),
            properties=Primitive.from_proto(resource.properties),
            logging_config=WorkflowTemplateJobsSparkJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class WorkflowTemplateJobsSparkJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsSparkJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsSparkJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsSparkJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig()
        )
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class WorkflowTemplateJobsSparkJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsSparkJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsSparkJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsPysparkJob(object):
    def __init__(
        self,
        main_python_file_uri: str = None,
        args: list = None,
        python_file_uris: list = None,
        jar_file_uris: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_python_file_uri = main_python_file_uri
        self.args = args
        self.python_file_uris = python_file_uris
        self.jar_file_uris = jar_file_uris
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsPysparkJob()
        if Primitive.to_proto(resource.main_python_file_uri):
            res.main_python_file_uri = Primitive.to_proto(resource.main_python_file_uri)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.python_file_uris):
            res.python_file_uris.extend(Primitive.to_proto(resource.python_file_uris))
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if WorkflowTemplateJobsPysparkJobLoggingConfig.to_proto(
            resource.logging_config
        ):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsPysparkJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPysparkJob(
            main_python_file_uri=Primitive.from_proto(resource.main_python_file_uri),
            args=Primitive.from_proto(resource.args),
            python_file_uris=Primitive.from_proto(resource.python_file_uris),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
            file_uris=Primitive.from_proto(resource.file_uris),
            archive_uris=Primitive.from_proto(resource.archive_uris),
            properties=Primitive.from_proto(resource.properties),
            logging_config=WorkflowTemplateJobsPysparkJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class WorkflowTemplateJobsPysparkJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPysparkJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsPysparkJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsPysparkJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig()
        )
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPysparkJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class WorkflowTemplateJobsPysparkJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsPysparkJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsPysparkJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsHiveJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        continue_on_failure: bool = None,
        script_variables: dict = None,
        properties: dict = None,
        jar_file_uris: list = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.continue_on_failure = continue_on_failure
        self.script_variables = script_variables
        self.properties = properties
        self.jar_file_uris = jar_file_uris

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsHiveJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if WorkflowTemplateJobsHiveJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(
                WorkflowTemplateJobsHiveJobQueryList.to_proto(resource.query_list)
            )
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.continue_on_failure):
            res.continue_on_failure = Primitive.to_proto(resource.continue_on_failure)
        if Primitive.to_proto(resource.script_variables):
            res.script_variables = Primitive.to_proto(resource.script_variables)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsHiveJob(
            query_file_uri=Primitive.from_proto(resource.query_file_uri),
            query_list=WorkflowTemplateJobsHiveJobQueryList.from_proto(
                resource.query_list
            ),
            continue_on_failure=Primitive.from_proto(resource.continue_on_failure),
            script_variables=Primitive.from_proto(resource.script_variables),
            properties=Primitive.from_proto(resource.properties),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
        )


class WorkflowTemplateJobsHiveJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsHiveJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsHiveJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsHiveJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsHiveJobQueryList()
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsHiveJobQueryList(
            queries=Primitive.from_proto(resource.queries),
        )


class WorkflowTemplateJobsHiveJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsHiveJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsHiveJobQueryList.from_proto(i) for i in resources]


class WorkflowTemplateJobsPigJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        continue_on_failure: bool = None,
        script_variables: dict = None,
        properties: dict = None,
        jar_file_uris: list = None,
        logging_config: dict = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.continue_on_failure = continue_on_failure
        self.script_variables = script_variables
        self.properties = properties
        self.jar_file_uris = jar_file_uris
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsPigJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if WorkflowTemplateJobsPigJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(
                WorkflowTemplateJobsPigJobQueryList.to_proto(resource.query_list)
            )
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.continue_on_failure):
            res.continue_on_failure = Primitive.to_proto(resource.continue_on_failure)
        if Primitive.to_proto(resource.script_variables):
            res.script_variables = Primitive.to_proto(resource.script_variables)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if WorkflowTemplateJobsPigJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsPigJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPigJob(
            query_file_uri=Primitive.from_proto(resource.query_file_uri),
            query_list=WorkflowTemplateJobsPigJobQueryList.from_proto(
                resource.query_list
            ),
            continue_on_failure=Primitive.from_proto(resource.continue_on_failure),
            script_variables=Primitive.from_proto(resource.script_variables),
            properties=Primitive.from_proto(resource.properties),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
            logging_config=WorkflowTemplateJobsPigJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class WorkflowTemplateJobsPigJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPigJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsPigJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsPigJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsPigJobQueryList()
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPigJobQueryList(
            queries=Primitive.from_proto(resource.queries),
        )


class WorkflowTemplateJobsPigJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPigJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsPigJobQueryList.from_proto(i) for i in resources]


class WorkflowTemplateJobsPigJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig()
        )
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPigJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class WorkflowTemplateJobsPigJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPigJobLoggingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsPigJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsSparkRJob(object):
    def __init__(
        self,
        main_r_file_uri: str = None,
        args: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_r_file_uri = main_r_file_uri
        self.args = args
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsSparkRJob()
        if Primitive.to_proto(resource.main_r_file_uri):
            res.main_r_file_uri = Primitive.to_proto(resource.main_r_file_uri)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if WorkflowTemplateJobsSparkRJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsSparkRJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkRJob(
            main_r_file_uri=Primitive.from_proto(resource.main_r_file_uri),
            args=Primitive.from_proto(resource.args),
            file_uris=Primitive.from_proto(resource.file_uris),
            archive_uris=Primitive.from_proto(resource.archive_uris),
            properties=Primitive.from_proto(resource.properties),
            logging_config=WorkflowTemplateJobsSparkRJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class WorkflowTemplateJobsSparkRJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsSparkRJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsSparkRJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsSparkRJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig()
        )
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkRJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class WorkflowTemplateJobsSparkRJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsSparkRJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsSparkRJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsSparkSqlJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        script_variables: dict = None,
        properties: dict = None,
        jar_file_uris: list = None,
        logging_config: dict = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.script_variables = script_variables
        self.properties = properties
        self.jar_file_uris = jar_file_uris
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsSparkSqlJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if WorkflowTemplateJobsSparkSqlJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(
                WorkflowTemplateJobsSparkSqlJobQueryList.to_proto(resource.query_list)
            )
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.script_variables):
            res.script_variables = Primitive.to_proto(resource.script_variables)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if WorkflowTemplateJobsSparkSqlJobLoggingConfig.to_proto(
            resource.logging_config
        ):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsSparkSqlJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkSqlJob(
            query_file_uri=Primitive.from_proto(resource.query_file_uri),
            query_list=WorkflowTemplateJobsSparkSqlJobQueryList.from_proto(
                resource.query_list
            ),
            script_variables=Primitive.from_proto(resource.script_variables),
            properties=Primitive.from_proto(resource.properties),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
            logging_config=WorkflowTemplateJobsSparkSqlJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class WorkflowTemplateJobsSparkSqlJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsSparkSqlJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsSparkSqlJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsSparkSqlJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList()
        )
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkSqlJobQueryList(
            queries=Primitive.from_proto(resource.queries),
        )


class WorkflowTemplateJobsSparkSqlJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsSparkSqlJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsSparkSqlJobQueryList.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsSparkSqlJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig()
        )
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkSqlJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class WorkflowTemplateJobsSparkSqlJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsSparkSqlJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsSparkSqlJobLoggingConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplateJobsPrestoJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        continue_on_failure: bool = None,
        output_format: str = None,
        client_tags: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.continue_on_failure = continue_on_failure
        self.output_format = output_format
        self.client_tags = client_tags
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsPrestoJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if WorkflowTemplateJobsPrestoJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(
                WorkflowTemplateJobsPrestoJobQueryList.to_proto(resource.query_list)
            )
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.continue_on_failure):
            res.continue_on_failure = Primitive.to_proto(resource.continue_on_failure)
        if Primitive.to_proto(resource.output_format):
            res.output_format = Primitive.to_proto(resource.output_format)
        if Primitive.to_proto(resource.client_tags):
            res.client_tags.extend(Primitive.to_proto(resource.client_tags))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if WorkflowTemplateJobsPrestoJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsPrestoJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPrestoJob(
            query_file_uri=Primitive.from_proto(resource.query_file_uri),
            query_list=WorkflowTemplateJobsPrestoJobQueryList.from_proto(
                resource.query_list
            ),
            continue_on_failure=Primitive.from_proto(resource.continue_on_failure),
            output_format=Primitive.from_proto(resource.output_format),
            client_tags=Primitive.from_proto(resource.client_tags),
            properties=Primitive.from_proto(resource.properties),
            logging_config=WorkflowTemplateJobsPrestoJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class WorkflowTemplateJobsPrestoJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPrestoJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsPrestoJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsPrestoJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsPrestoJobQueryList()
        )
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPrestoJobQueryList(
            queries=Primitive.from_proto(resource.queries),
        )


class WorkflowTemplateJobsPrestoJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPrestoJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsPrestoJobQueryList.from_proto(i) for i in resources]


class WorkflowTemplateJobsPrestoJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig()
        )
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPrestoJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class WorkflowTemplateJobsPrestoJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsPrestoJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsPrestoJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsScheduling(object):
    def __init__(
        self, max_failures_per_hour: int = None, max_failures_total: int = None
    ):
        self.max_failures_per_hour = max_failures_per_hour
        self.max_failures_total = max_failures_total

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateJobsScheduling()
        if Primitive.to_proto(resource.max_failures_per_hour):
            res.max_failures_per_hour = Primitive.to_proto(
                resource.max_failures_per_hour
            )
        if Primitive.to_proto(resource.max_failures_total):
            res.max_failures_total = Primitive.to_proto(resource.max_failures_total)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsScheduling(
            max_failures_per_hour=Primitive.from_proto(resource.max_failures_per_hour),
            max_failures_total=Primitive.from_proto(resource.max_failures_total),
        )


class WorkflowTemplateJobsSchedulingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsScheduling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsScheduling.from_proto(i) for i in resources]


class WorkflowTemplateParameters(object):
    def __init__(
        self,
        name: str = None,
        fields: list = None,
        description: str = None,
        validation: dict = None,
    ):
        self.name = name
        self.fields = fields
        self.description = description
        self.validation = validation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateParameters()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.fields):
            res.fields.extend(Primitive.to_proto(resource.fields))
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if WorkflowTemplateParametersValidation.to_proto(resource.validation):
            res.validation.CopyFrom(
                WorkflowTemplateParametersValidation.to_proto(resource.validation)
            )
        else:
            res.ClearField("validation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateParameters(
            name=Primitive.from_proto(resource.name),
            fields=Primitive.from_proto(resource.fields),
            description=Primitive.from_proto(resource.description),
            validation=WorkflowTemplateParametersValidation.from_proto(
                resource.validation
            ),
        )


class WorkflowTemplateParametersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateParameters.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateParameters.from_proto(i) for i in resources]


class WorkflowTemplateParametersValidation(object):
    def __init__(self, regex: dict = None, values: dict = None):
        self.regex = regex
        self.values = values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocAlphaWorkflowTemplateParametersValidation()
        if WorkflowTemplateParametersValidationRegex.to_proto(resource.regex):
            res.regex.CopyFrom(
                WorkflowTemplateParametersValidationRegex.to_proto(resource.regex)
            )
        else:
            res.ClearField("regex")
        if WorkflowTemplateParametersValidationValues.to_proto(resource.values):
            res.values.CopyFrom(
                WorkflowTemplateParametersValidationValues.to_proto(resource.values)
            )
        else:
            res.ClearField("values")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateParametersValidation(
            regex=WorkflowTemplateParametersValidationRegex.from_proto(resource.regex),
            values=WorkflowTemplateParametersValidationValues.from_proto(
                resource.values
            ),
        )


class WorkflowTemplateParametersValidationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateParametersValidation.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateParametersValidation.from_proto(i) for i in resources]


class WorkflowTemplateParametersValidationRegex(object):
    def __init__(self, regexes: list = None):
        self.regexes = regexes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplateParametersValidationRegex()
        )
        if Primitive.to_proto(resource.regexes):
            res.regexes.extend(Primitive.to_proto(resource.regexes))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateParametersValidationRegex(
            regexes=Primitive.from_proto(resource.regexes),
        )


class WorkflowTemplateParametersValidationRegexArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateParametersValidationRegex.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateParametersValidationRegex.from_proto(i) for i in resources
        ]


class WorkflowTemplateParametersValidationValues(object):
    def __init__(self, values: list = None):
        self.values = values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocAlphaWorkflowTemplateParametersValidationValues()
        )
        if Primitive.to_proto(resource.values):
            res.values.extend(Primitive.to_proto(resource.values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateParametersValidationValues(
            values=Primitive.from_proto(resource.values),
        )


class WorkflowTemplateParametersValidationValuesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateParametersValidationValues.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateParametersValidationValues.from_proto(i) for i in resources
        ]


class WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.Value(
            "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.Name(
            resource
        )[
            len(
                "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"
            ) :
        ]


class WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.Value(
            "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.Name(
            resource
        )[
            len(
                "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"
            ) :
        ]


class WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum.Value(
            "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum.Name(
            resource
        )[
            len(
                "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum"
            ) :
        ]


class WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum.Value(
            "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum.Name(
            resource
        )[
            len(
                "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum"
            ) :
        ]


class WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum.Value(
            "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum.Name(
            resource
        )[
            len(
                "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum"
            ) :
        ]


class WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum.Value(
            "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workflow_template_pb2.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum.Name(
            resource
        )[
            len(
                "DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum"
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
