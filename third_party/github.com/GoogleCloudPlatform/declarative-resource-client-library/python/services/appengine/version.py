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
from google3.cloud.graphite.mmv2.services.google.app_engine import version_pb2
from google3.cloud.graphite.mmv2.services.google.app_engine import version_pb2_grpc

from typing import List


class Version(object):
    def __init__(
        self,
        consumer_name: str = None,
        name: str = None,
        automatic_scaling: dict = None,
        basic_scaling: dict = None,
        manual_scaling: dict = None,
        inbound_services: list = None,
        instance_class: str = None,
        network: dict = None,
        zones: list = None,
        resources: dict = None,
        runtime: str = None,
        runtime_channel: str = None,
        threadsafe: bool = None,
        vm: bool = None,
        beta_settings: dict = None,
        env: str = None,
        serving_status: str = None,
        created_by: str = None,
        create_time: str = None,
        disk_usage_bytes: int = None,
        runtime_api_version: str = None,
        runtime_main_executable_path: str = None,
        handlers: list = None,
        error_handlers: list = None,
        libraries: list = None,
        api_config: dict = None,
        env_variables: dict = None,
        default_expiration: str = None,
        deployment: dict = None,
        health_check: dict = None,
        readiness_check: dict = None,
        liveness_check: dict = None,
        nobuild_files_regex: str = None,
        version_url: str = None,
        entrypoint: dict = None,
        vpc_access_connector: dict = None,
        app: str = None,
        service: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.consumer_name = consumer_name
        self.name = name
        self.automatic_scaling = automatic_scaling
        self.basic_scaling = basic_scaling
        self.manual_scaling = manual_scaling
        self.inbound_services = inbound_services
        self.instance_class = instance_class
        self.network = network
        self.zones = zones
        self.resources = resources
        self.runtime = runtime
        self.runtime_channel = runtime_channel
        self.threadsafe = threadsafe
        self.vm = vm
        self.beta_settings = beta_settings
        self.env = env
        self.serving_status = serving_status
        self.runtime_api_version = runtime_api_version
        self.runtime_main_executable_path = runtime_main_executable_path
        self.handlers = handlers
        self.error_handlers = error_handlers
        self.libraries = libraries
        self.api_config = api_config
        self.env_variables = env_variables
        self.default_expiration = default_expiration
        self.deployment = deployment
        self.health_check = health_check
        self.readiness_check = readiness_check
        self.liveness_check = liveness_check
        self.nobuild_files_regex = nobuild_files_regex
        self.entrypoint = entrypoint
        self.vpc_access_connector = vpc_access_connector
        self.app = app
        self.service = service
        self.service_account_file = service_account_file

    def apply(self):
        stub = version_pb2_grpc.AppengineVersionServiceStub(channel.Channel())
        request = version_pb2.ApplyAppengineVersionRequest()
        if Primitive.to_proto(self.consumer_name):
            request.resource.consumer_name = Primitive.to_proto(self.consumer_name)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if VersionAutomaticScaling.to_proto(self.automatic_scaling):
            request.resource.automatic_scaling.CopyFrom(
                VersionAutomaticScaling.to_proto(self.automatic_scaling)
            )
        else:
            request.resource.ClearField("automatic_scaling")
        if VersionBasicScaling.to_proto(self.basic_scaling):
            request.resource.basic_scaling.CopyFrom(
                VersionBasicScaling.to_proto(self.basic_scaling)
            )
        else:
            request.resource.ClearField("basic_scaling")
        if VersionManualScaling.to_proto(self.manual_scaling):
            request.resource.manual_scaling.CopyFrom(
                VersionManualScaling.to_proto(self.manual_scaling)
            )
        else:
            request.resource.ClearField("manual_scaling")
        if VersionInboundServicesEnumArray.to_proto(self.inbound_services):
            request.resource.inbound_services.extend(
                VersionInboundServicesEnumArray.to_proto(self.inbound_services)
            )
        if Primitive.to_proto(self.instance_class):
            request.resource.instance_class = Primitive.to_proto(self.instance_class)

        if VersionNetwork.to_proto(self.network):
            request.resource.network.CopyFrom(VersionNetwork.to_proto(self.network))
        else:
            request.resource.ClearField("network")
        if Primitive.to_proto(self.zones):
            request.resource.zones.extend(Primitive.to_proto(self.zones))
        if VersionResources.to_proto(self.resources):
            request.resource.resources.CopyFrom(
                VersionResources.to_proto(self.resources)
            )
        else:
            request.resource.ClearField("resources")
        if Primitive.to_proto(self.runtime):
            request.resource.runtime = Primitive.to_proto(self.runtime)

        if Primitive.to_proto(self.runtime_channel):
            request.resource.runtime_channel = Primitive.to_proto(self.runtime_channel)

        if Primitive.to_proto(self.threadsafe):
            request.resource.threadsafe = Primitive.to_proto(self.threadsafe)

        if Primitive.to_proto(self.vm):
            request.resource.vm = Primitive.to_proto(self.vm)

        if Primitive.to_proto(self.beta_settings):
            request.resource.beta_settings = Primitive.to_proto(self.beta_settings)

        if Primitive.to_proto(self.env):
            request.resource.env = Primitive.to_proto(self.env)

        if VersionServingStatusEnum.to_proto(self.serving_status):
            request.resource.serving_status = VersionServingStatusEnum.to_proto(
                self.serving_status
            )

        if Primitive.to_proto(self.runtime_api_version):
            request.resource.runtime_api_version = Primitive.to_proto(
                self.runtime_api_version
            )

        if Primitive.to_proto(self.runtime_main_executable_path):
            request.resource.runtime_main_executable_path = Primitive.to_proto(
                self.runtime_main_executable_path
            )

        if VersionHandlersArray.to_proto(self.handlers):
            request.resource.handlers.extend(
                VersionHandlersArray.to_proto(self.handlers)
            )
        if VersionErrorHandlersArray.to_proto(self.error_handlers):
            request.resource.error_handlers.extend(
                VersionErrorHandlersArray.to_proto(self.error_handlers)
            )
        if VersionLibrariesArray.to_proto(self.libraries):
            request.resource.libraries.extend(
                VersionLibrariesArray.to_proto(self.libraries)
            )
        if VersionApiConfig.to_proto(self.api_config):
            request.resource.api_config.CopyFrom(
                VersionApiConfig.to_proto(self.api_config)
            )
        else:
            request.resource.ClearField("api_config")
        if Primitive.to_proto(self.env_variables):
            request.resource.env_variables = Primitive.to_proto(self.env_variables)

        if Primitive.to_proto(self.default_expiration):
            request.resource.default_expiration = Primitive.to_proto(
                self.default_expiration
            )

        if VersionDeployment.to_proto(self.deployment):
            request.resource.deployment.CopyFrom(
                VersionDeployment.to_proto(self.deployment)
            )
        else:
            request.resource.ClearField("deployment")
        if VersionHealthCheck.to_proto(self.health_check):
            request.resource.health_check.CopyFrom(
                VersionHealthCheck.to_proto(self.health_check)
            )
        else:
            request.resource.ClearField("health_check")
        if VersionReadinessCheck.to_proto(self.readiness_check):
            request.resource.readiness_check.CopyFrom(
                VersionReadinessCheck.to_proto(self.readiness_check)
            )
        else:
            request.resource.ClearField("readiness_check")
        if VersionLivenessCheck.to_proto(self.liveness_check):
            request.resource.liveness_check.CopyFrom(
                VersionLivenessCheck.to_proto(self.liveness_check)
            )
        else:
            request.resource.ClearField("liveness_check")
        if Primitive.to_proto(self.nobuild_files_regex):
            request.resource.nobuild_files_regex = Primitive.to_proto(
                self.nobuild_files_regex
            )

        if VersionEntrypoint.to_proto(self.entrypoint):
            request.resource.entrypoint.CopyFrom(
                VersionEntrypoint.to_proto(self.entrypoint)
            )
        else:
            request.resource.ClearField("entrypoint")
        if VersionVPCAccessConnector.to_proto(self.vpc_access_connector):
            request.resource.vpc_access_connector.CopyFrom(
                VersionVPCAccessConnector.to_proto(self.vpc_access_connector)
            )
        else:
            request.resource.ClearField("vpc_access_connector")
        if Primitive.to_proto(self.app):
            request.resource.app = Primitive.to_proto(self.app)

        if Primitive.to_proto(self.service):
            request.resource.service = Primitive.to_proto(self.service)

        request.service_account_file = self.service_account_file

        response = stub.ApplyAppengineVersion(request)
        self.consumer_name = Primitive.from_proto(response.consumer_name)
        self.name = Primitive.from_proto(response.name)
        self.automatic_scaling = VersionAutomaticScaling.from_proto(
            response.automatic_scaling
        )
        self.basic_scaling = VersionBasicScaling.from_proto(response.basic_scaling)
        self.manual_scaling = VersionManualScaling.from_proto(response.manual_scaling)
        self.inbound_services = VersionInboundServicesEnumArray.from_proto(
            response.inbound_services
        )
        self.instance_class = Primitive.from_proto(response.instance_class)
        self.network = VersionNetwork.from_proto(response.network)
        self.zones = Primitive.from_proto(response.zones)
        self.resources = VersionResources.from_proto(response.resources)
        self.runtime = Primitive.from_proto(response.runtime)
        self.runtime_channel = Primitive.from_proto(response.runtime_channel)
        self.threadsafe = Primitive.from_proto(response.threadsafe)
        self.vm = Primitive.from_proto(response.vm)
        self.beta_settings = Primitive.from_proto(response.beta_settings)
        self.env = Primitive.from_proto(response.env)
        self.serving_status = VersionServingStatusEnum.from_proto(
            response.serving_status
        )
        self.created_by = Primitive.from_proto(response.created_by)
        self.create_time = Primitive.from_proto(response.create_time)
        self.disk_usage_bytes = Primitive.from_proto(response.disk_usage_bytes)
        self.runtime_api_version = Primitive.from_proto(response.runtime_api_version)
        self.runtime_main_executable_path = Primitive.from_proto(
            response.runtime_main_executable_path
        )
        self.handlers = VersionHandlersArray.from_proto(response.handlers)
        self.error_handlers = VersionErrorHandlersArray.from_proto(
            response.error_handlers
        )
        self.libraries = VersionLibrariesArray.from_proto(response.libraries)
        self.api_config = VersionApiConfig.from_proto(response.api_config)
        self.env_variables = Primitive.from_proto(response.env_variables)
        self.default_expiration = Primitive.from_proto(response.default_expiration)
        self.deployment = VersionDeployment.from_proto(response.deployment)
        self.health_check = VersionHealthCheck.from_proto(response.health_check)
        self.readiness_check = VersionReadinessCheck.from_proto(
            response.readiness_check
        )
        self.liveness_check = VersionLivenessCheck.from_proto(response.liveness_check)
        self.nobuild_files_regex = Primitive.from_proto(response.nobuild_files_regex)
        self.version_url = Primitive.from_proto(response.version_url)
        self.entrypoint = VersionEntrypoint.from_proto(response.entrypoint)
        self.vpc_access_connector = VersionVPCAccessConnector.from_proto(
            response.vpc_access_connector
        )
        self.app = Primitive.from_proto(response.app)
        self.service = Primitive.from_proto(response.service)

    def delete(self):
        stub = version_pb2_grpc.AppengineVersionServiceStub(channel.Channel())
        request = version_pb2.DeleteAppengineVersionRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.consumer_name):
            request.resource.consumer_name = Primitive.to_proto(self.consumer_name)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if VersionAutomaticScaling.to_proto(self.automatic_scaling):
            request.resource.automatic_scaling.CopyFrom(
                VersionAutomaticScaling.to_proto(self.automatic_scaling)
            )
        else:
            request.resource.ClearField("automatic_scaling")
        if VersionBasicScaling.to_proto(self.basic_scaling):
            request.resource.basic_scaling.CopyFrom(
                VersionBasicScaling.to_proto(self.basic_scaling)
            )
        else:
            request.resource.ClearField("basic_scaling")
        if VersionManualScaling.to_proto(self.manual_scaling):
            request.resource.manual_scaling.CopyFrom(
                VersionManualScaling.to_proto(self.manual_scaling)
            )
        else:
            request.resource.ClearField("manual_scaling")
        if VersionInboundServicesEnumArray.to_proto(self.inbound_services):
            request.resource.inbound_services.extend(
                VersionInboundServicesEnumArray.to_proto(self.inbound_services)
            )
        if Primitive.to_proto(self.instance_class):
            request.resource.instance_class = Primitive.to_proto(self.instance_class)

        if VersionNetwork.to_proto(self.network):
            request.resource.network.CopyFrom(VersionNetwork.to_proto(self.network))
        else:
            request.resource.ClearField("network")
        if Primitive.to_proto(self.zones):
            request.resource.zones.extend(Primitive.to_proto(self.zones))
        if VersionResources.to_proto(self.resources):
            request.resource.resources.CopyFrom(
                VersionResources.to_proto(self.resources)
            )
        else:
            request.resource.ClearField("resources")
        if Primitive.to_proto(self.runtime):
            request.resource.runtime = Primitive.to_proto(self.runtime)

        if Primitive.to_proto(self.runtime_channel):
            request.resource.runtime_channel = Primitive.to_proto(self.runtime_channel)

        if Primitive.to_proto(self.threadsafe):
            request.resource.threadsafe = Primitive.to_proto(self.threadsafe)

        if Primitive.to_proto(self.vm):
            request.resource.vm = Primitive.to_proto(self.vm)

        if Primitive.to_proto(self.beta_settings):
            request.resource.beta_settings = Primitive.to_proto(self.beta_settings)

        if Primitive.to_proto(self.env):
            request.resource.env = Primitive.to_proto(self.env)

        if VersionServingStatusEnum.to_proto(self.serving_status):
            request.resource.serving_status = VersionServingStatusEnum.to_proto(
                self.serving_status
            )

        if Primitive.to_proto(self.runtime_api_version):
            request.resource.runtime_api_version = Primitive.to_proto(
                self.runtime_api_version
            )

        if Primitive.to_proto(self.runtime_main_executable_path):
            request.resource.runtime_main_executable_path = Primitive.to_proto(
                self.runtime_main_executable_path
            )

        if VersionHandlersArray.to_proto(self.handlers):
            request.resource.handlers.extend(
                VersionHandlersArray.to_proto(self.handlers)
            )
        if VersionErrorHandlersArray.to_proto(self.error_handlers):
            request.resource.error_handlers.extend(
                VersionErrorHandlersArray.to_proto(self.error_handlers)
            )
        if VersionLibrariesArray.to_proto(self.libraries):
            request.resource.libraries.extend(
                VersionLibrariesArray.to_proto(self.libraries)
            )
        if VersionApiConfig.to_proto(self.api_config):
            request.resource.api_config.CopyFrom(
                VersionApiConfig.to_proto(self.api_config)
            )
        else:
            request.resource.ClearField("api_config")
        if Primitive.to_proto(self.env_variables):
            request.resource.env_variables = Primitive.to_proto(self.env_variables)

        if Primitive.to_proto(self.default_expiration):
            request.resource.default_expiration = Primitive.to_proto(
                self.default_expiration
            )

        if VersionDeployment.to_proto(self.deployment):
            request.resource.deployment.CopyFrom(
                VersionDeployment.to_proto(self.deployment)
            )
        else:
            request.resource.ClearField("deployment")
        if VersionHealthCheck.to_proto(self.health_check):
            request.resource.health_check.CopyFrom(
                VersionHealthCheck.to_proto(self.health_check)
            )
        else:
            request.resource.ClearField("health_check")
        if VersionReadinessCheck.to_proto(self.readiness_check):
            request.resource.readiness_check.CopyFrom(
                VersionReadinessCheck.to_proto(self.readiness_check)
            )
        else:
            request.resource.ClearField("readiness_check")
        if VersionLivenessCheck.to_proto(self.liveness_check):
            request.resource.liveness_check.CopyFrom(
                VersionLivenessCheck.to_proto(self.liveness_check)
            )
        else:
            request.resource.ClearField("liveness_check")
        if Primitive.to_proto(self.nobuild_files_regex):
            request.resource.nobuild_files_regex = Primitive.to_proto(
                self.nobuild_files_regex
            )

        if VersionEntrypoint.to_proto(self.entrypoint):
            request.resource.entrypoint.CopyFrom(
                VersionEntrypoint.to_proto(self.entrypoint)
            )
        else:
            request.resource.ClearField("entrypoint")
        if VersionVPCAccessConnector.to_proto(self.vpc_access_connector):
            request.resource.vpc_access_connector.CopyFrom(
                VersionVPCAccessConnector.to_proto(self.vpc_access_connector)
            )
        else:
            request.resource.ClearField("vpc_access_connector")
        if Primitive.to_proto(self.app):
            request.resource.app = Primitive.to_proto(self.app)

        if Primitive.to_proto(self.service):
            request.resource.service = Primitive.to_proto(self.service)

        response = stub.DeleteAppengineVersion(request)

    @classmethod
    def list(self, app, service, service_account_file=""):
        stub = version_pb2_grpc.AppengineVersionServiceStub(channel.Channel())
        request = version_pb2.ListAppengineVersionRequest()
        request.service_account_file = service_account_file
        request.App = app

        request.Service = service

        return stub.ListAppengineVersion(request).items

    def to_proto(self):
        resource = version_pb2.AppengineVersion()
        if Primitive.to_proto(self.consumer_name):
            resource.consumer_name = Primitive.to_proto(self.consumer_name)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if VersionAutomaticScaling.to_proto(self.automatic_scaling):
            resource.automatic_scaling.CopyFrom(
                VersionAutomaticScaling.to_proto(self.automatic_scaling)
            )
        else:
            resource.ClearField("automatic_scaling")
        if VersionBasicScaling.to_proto(self.basic_scaling):
            resource.basic_scaling.CopyFrom(
                VersionBasicScaling.to_proto(self.basic_scaling)
            )
        else:
            resource.ClearField("basic_scaling")
        if VersionManualScaling.to_proto(self.manual_scaling):
            resource.manual_scaling.CopyFrom(
                VersionManualScaling.to_proto(self.manual_scaling)
            )
        else:
            resource.ClearField("manual_scaling")
        if VersionInboundServicesEnumArray.to_proto(self.inbound_services):
            resource.inbound_services.extend(
                VersionInboundServicesEnumArray.to_proto(self.inbound_services)
            )
        if Primitive.to_proto(self.instance_class):
            resource.instance_class = Primitive.to_proto(self.instance_class)
        if VersionNetwork.to_proto(self.network):
            resource.network.CopyFrom(VersionNetwork.to_proto(self.network))
        else:
            resource.ClearField("network")
        if Primitive.to_proto(self.zones):
            resource.zones.extend(Primitive.to_proto(self.zones))
        if VersionResources.to_proto(self.resources):
            resource.resources.CopyFrom(VersionResources.to_proto(self.resources))
        else:
            resource.ClearField("resources")
        if Primitive.to_proto(self.runtime):
            resource.runtime = Primitive.to_proto(self.runtime)
        if Primitive.to_proto(self.runtime_channel):
            resource.runtime_channel = Primitive.to_proto(self.runtime_channel)
        if Primitive.to_proto(self.threadsafe):
            resource.threadsafe = Primitive.to_proto(self.threadsafe)
        if Primitive.to_proto(self.vm):
            resource.vm = Primitive.to_proto(self.vm)
        if Primitive.to_proto(self.beta_settings):
            resource.beta_settings = Primitive.to_proto(self.beta_settings)
        if Primitive.to_proto(self.env):
            resource.env = Primitive.to_proto(self.env)
        if VersionServingStatusEnum.to_proto(self.serving_status):
            resource.serving_status = VersionServingStatusEnum.to_proto(
                self.serving_status
            )
        if Primitive.to_proto(self.runtime_api_version):
            resource.runtime_api_version = Primitive.to_proto(self.runtime_api_version)
        if Primitive.to_proto(self.runtime_main_executable_path):
            resource.runtime_main_executable_path = Primitive.to_proto(
                self.runtime_main_executable_path
            )
        if VersionHandlersArray.to_proto(self.handlers):
            resource.handlers.extend(VersionHandlersArray.to_proto(self.handlers))
        if VersionErrorHandlersArray.to_proto(self.error_handlers):
            resource.error_handlers.extend(
                VersionErrorHandlersArray.to_proto(self.error_handlers)
            )
        if VersionLibrariesArray.to_proto(self.libraries):
            resource.libraries.extend(VersionLibrariesArray.to_proto(self.libraries))
        if VersionApiConfig.to_proto(self.api_config):
            resource.api_config.CopyFrom(VersionApiConfig.to_proto(self.api_config))
        else:
            resource.ClearField("api_config")
        if Primitive.to_proto(self.env_variables):
            resource.env_variables = Primitive.to_proto(self.env_variables)
        if Primitive.to_proto(self.default_expiration):
            resource.default_expiration = Primitive.to_proto(self.default_expiration)
        if VersionDeployment.to_proto(self.deployment):
            resource.deployment.CopyFrom(VersionDeployment.to_proto(self.deployment))
        else:
            resource.ClearField("deployment")
        if VersionHealthCheck.to_proto(self.health_check):
            resource.health_check.CopyFrom(
                VersionHealthCheck.to_proto(self.health_check)
            )
        else:
            resource.ClearField("health_check")
        if VersionReadinessCheck.to_proto(self.readiness_check):
            resource.readiness_check.CopyFrom(
                VersionReadinessCheck.to_proto(self.readiness_check)
            )
        else:
            resource.ClearField("readiness_check")
        if VersionLivenessCheck.to_proto(self.liveness_check):
            resource.liveness_check.CopyFrom(
                VersionLivenessCheck.to_proto(self.liveness_check)
            )
        else:
            resource.ClearField("liveness_check")
        if Primitive.to_proto(self.nobuild_files_regex):
            resource.nobuild_files_regex = Primitive.to_proto(self.nobuild_files_regex)
        if VersionEntrypoint.to_proto(self.entrypoint):
            resource.entrypoint.CopyFrom(VersionEntrypoint.to_proto(self.entrypoint))
        else:
            resource.ClearField("entrypoint")
        if VersionVPCAccessConnector.to_proto(self.vpc_access_connector):
            resource.vpc_access_connector.CopyFrom(
                VersionVPCAccessConnector.to_proto(self.vpc_access_connector)
            )
        else:
            resource.ClearField("vpc_access_connector")
        if Primitive.to_proto(self.app):
            resource.app = Primitive.to_proto(self.app)
        if Primitive.to_proto(self.service):
            resource.service = Primitive.to_proto(self.service)
        return resource


class VersionAutomaticScaling(object):
    def __init__(
        self,
        cool_down_period: str = None,
        cpu_utilization: dict = None,
        max_concurrent_requests: int = None,
        max_idle_instances: int = None,
        max_total_instances: int = None,
        max_pending_latency: str = None,
        min_idle_instances: int = None,
        min_total_instances: int = None,
        min_pending_latency: str = None,
        request_utilization: dict = None,
        disk_utilization: dict = None,
        network_utilization: dict = None,
        standard_scheduler_settings: dict = None,
    ):
        self.cool_down_period = cool_down_period
        self.cpu_utilization = cpu_utilization
        self.max_concurrent_requests = max_concurrent_requests
        self.max_idle_instances = max_idle_instances
        self.max_total_instances = max_total_instances
        self.max_pending_latency = max_pending_latency
        self.min_idle_instances = min_idle_instances
        self.min_total_instances = min_total_instances
        self.min_pending_latency = min_pending_latency
        self.request_utilization = request_utilization
        self.disk_utilization = disk_utilization
        self.network_utilization = network_utilization
        self.standard_scheduler_settings = standard_scheduler_settings

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScaling()
        if Primitive.to_proto(resource.cool_down_period):
            res.cool_down_period = Primitive.to_proto(resource.cool_down_period)
        if VersionAutomaticScalingCpuUtilization.to_proto(resource.cpu_utilization):
            res.cpu_utilization.CopyFrom(
                VersionAutomaticScalingCpuUtilization.to_proto(resource.cpu_utilization)
            )
        else:
            res.ClearField("cpu_utilization")
        if Primitive.to_proto(resource.max_concurrent_requests):
            res.max_concurrent_requests = Primitive.to_proto(
                resource.max_concurrent_requests
            )
        if Primitive.to_proto(resource.max_idle_instances):
            res.max_idle_instances = Primitive.to_proto(resource.max_idle_instances)
        if Primitive.to_proto(resource.max_total_instances):
            res.max_total_instances = Primitive.to_proto(resource.max_total_instances)
        if Primitive.to_proto(resource.max_pending_latency):
            res.max_pending_latency = Primitive.to_proto(resource.max_pending_latency)
        if Primitive.to_proto(resource.min_idle_instances):
            res.min_idle_instances = Primitive.to_proto(resource.min_idle_instances)
        if Primitive.to_proto(resource.min_total_instances):
            res.min_total_instances = Primitive.to_proto(resource.min_total_instances)
        if Primitive.to_proto(resource.min_pending_latency):
            res.min_pending_latency = Primitive.to_proto(resource.min_pending_latency)
        if VersionAutomaticScalingRequestUtilization.to_proto(
            resource.request_utilization
        ):
            res.request_utilization.CopyFrom(
                VersionAutomaticScalingRequestUtilization.to_proto(
                    resource.request_utilization
                )
            )
        else:
            res.ClearField("request_utilization")
        if VersionAutomaticScalingDiskUtilization.to_proto(resource.disk_utilization):
            res.disk_utilization.CopyFrom(
                VersionAutomaticScalingDiskUtilization.to_proto(
                    resource.disk_utilization
                )
            )
        else:
            res.ClearField("disk_utilization")
        if VersionAutomaticScalingNetworkUtilization.to_proto(
            resource.network_utilization
        ):
            res.network_utilization.CopyFrom(
                VersionAutomaticScalingNetworkUtilization.to_proto(
                    resource.network_utilization
                )
            )
        else:
            res.ClearField("network_utilization")
        if VersionAutomaticScalingStandardSchedulerSettings.to_proto(
            resource.standard_scheduler_settings
        ):
            res.standard_scheduler_settings.CopyFrom(
                VersionAutomaticScalingStandardSchedulerSettings.to_proto(
                    resource.standard_scheduler_settings
                )
            )
        else:
            res.ClearField("standard_scheduler_settings")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScaling(
            cool_down_period=Primitive.from_proto(resource.cool_down_period),
            cpu_utilization=VersionAutomaticScalingCpuUtilization.from_proto(
                resource.cpu_utilization
            ),
            max_concurrent_requests=Primitive.from_proto(
                resource.max_concurrent_requests
            ),
            max_idle_instances=Primitive.from_proto(resource.max_idle_instances),
            max_total_instances=Primitive.from_proto(resource.max_total_instances),
            max_pending_latency=Primitive.from_proto(resource.max_pending_latency),
            min_idle_instances=Primitive.from_proto(resource.min_idle_instances),
            min_total_instances=Primitive.from_proto(resource.min_total_instances),
            min_pending_latency=Primitive.from_proto(resource.min_pending_latency),
            request_utilization=VersionAutomaticScalingRequestUtilization.from_proto(
                resource.request_utilization
            ),
            disk_utilization=VersionAutomaticScalingDiskUtilization.from_proto(
                resource.disk_utilization
            ),
            network_utilization=VersionAutomaticScalingNetworkUtilization.from_proto(
                resource.network_utilization
            ),
            standard_scheduler_settings=VersionAutomaticScalingStandardSchedulerSettings.from_proto(
                resource.standard_scheduler_settings
            ),
        )


class VersionAutomaticScalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionAutomaticScaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionAutomaticScaling.from_proto(i) for i in resources]


class VersionAutomaticScalingCpuUtilization(object):
    def __init__(
        self, aggregation_window_length: str = None, target_utilization: float = None
    ):
        self.aggregation_window_length = aggregation_window_length
        self.target_utilization = target_utilization

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScalingCpuUtilization()
        if Primitive.to_proto(resource.aggregation_window_length):
            res.aggregation_window_length = Primitive.to_proto(
                resource.aggregation_window_length
            )
        if Primitive.to_proto(resource.target_utilization):
            res.target_utilization = Primitive.to_proto(resource.target_utilization)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScalingCpuUtilization(
            aggregation_window_length=Primitive.from_proto(
                resource.aggregation_window_length
            ),
            target_utilization=Primitive.from_proto(resource.target_utilization),
        )


class VersionAutomaticScalingCpuUtilizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionAutomaticScalingCpuUtilization.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionAutomaticScalingCpuUtilization.from_proto(i) for i in resources]


class VersionAutomaticScalingRequestUtilization(object):
    def __init__(
        self,
        target_request_count_per_second: int = None,
        target_concurrent_requests: int = None,
    ):
        self.target_request_count_per_second = target_request_count_per_second
        self.target_concurrent_requests = target_concurrent_requests

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScalingRequestUtilization()
        if Primitive.to_proto(resource.target_request_count_per_second):
            res.target_request_count_per_second = Primitive.to_proto(
                resource.target_request_count_per_second
            )
        if Primitive.to_proto(resource.target_concurrent_requests):
            res.target_concurrent_requests = Primitive.to_proto(
                resource.target_concurrent_requests
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScalingRequestUtilization(
            target_request_count_per_second=Primitive.from_proto(
                resource.target_request_count_per_second
            ),
            target_concurrent_requests=Primitive.from_proto(
                resource.target_concurrent_requests
            ),
        )


class VersionAutomaticScalingRequestUtilizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            VersionAutomaticScalingRequestUtilization.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            VersionAutomaticScalingRequestUtilization.from_proto(i) for i in resources
        ]


class VersionAutomaticScalingDiskUtilization(object):
    def __init__(
        self,
        target_write_bytes_per_second: int = None,
        target_write_ops_per_second: int = None,
        target_read_bytes_per_second: int = None,
        target_read_ops_per_second: int = None,
    ):
        self.target_write_bytes_per_second = target_write_bytes_per_second
        self.target_write_ops_per_second = target_write_ops_per_second
        self.target_read_bytes_per_second = target_read_bytes_per_second
        self.target_read_ops_per_second = target_read_ops_per_second

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScalingDiskUtilization()
        if Primitive.to_proto(resource.target_write_bytes_per_second):
            res.target_write_bytes_per_second = Primitive.to_proto(
                resource.target_write_bytes_per_second
            )
        if Primitive.to_proto(resource.target_write_ops_per_second):
            res.target_write_ops_per_second = Primitive.to_proto(
                resource.target_write_ops_per_second
            )
        if Primitive.to_proto(resource.target_read_bytes_per_second):
            res.target_read_bytes_per_second = Primitive.to_proto(
                resource.target_read_bytes_per_second
            )
        if Primitive.to_proto(resource.target_read_ops_per_second):
            res.target_read_ops_per_second = Primitive.to_proto(
                resource.target_read_ops_per_second
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScalingDiskUtilization(
            target_write_bytes_per_second=Primitive.from_proto(
                resource.target_write_bytes_per_second
            ),
            target_write_ops_per_second=Primitive.from_proto(
                resource.target_write_ops_per_second
            ),
            target_read_bytes_per_second=Primitive.from_proto(
                resource.target_read_bytes_per_second
            ),
            target_read_ops_per_second=Primitive.from_proto(
                resource.target_read_ops_per_second
            ),
        )


class VersionAutomaticScalingDiskUtilizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionAutomaticScalingDiskUtilization.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionAutomaticScalingDiskUtilization.from_proto(i) for i in resources]


class VersionAutomaticScalingNetworkUtilization(object):
    def __init__(
        self,
        target_sent_bytes_per_second: int = None,
        target_sent_packets_per_second: int = None,
        target_received_bytes_per_second: int = None,
        target_received_packets_per_second: int = None,
    ):
        self.target_sent_bytes_per_second = target_sent_bytes_per_second
        self.target_sent_packets_per_second = target_sent_packets_per_second
        self.target_received_bytes_per_second = target_received_bytes_per_second
        self.target_received_packets_per_second = target_received_packets_per_second

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScalingNetworkUtilization()
        if Primitive.to_proto(resource.target_sent_bytes_per_second):
            res.target_sent_bytes_per_second = Primitive.to_proto(
                resource.target_sent_bytes_per_second
            )
        if Primitive.to_proto(resource.target_sent_packets_per_second):
            res.target_sent_packets_per_second = Primitive.to_proto(
                resource.target_sent_packets_per_second
            )
        if Primitive.to_proto(resource.target_received_bytes_per_second):
            res.target_received_bytes_per_second = Primitive.to_proto(
                resource.target_received_bytes_per_second
            )
        if Primitive.to_proto(resource.target_received_packets_per_second):
            res.target_received_packets_per_second = Primitive.to_proto(
                resource.target_received_packets_per_second
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScalingNetworkUtilization(
            target_sent_bytes_per_second=Primitive.from_proto(
                resource.target_sent_bytes_per_second
            ),
            target_sent_packets_per_second=Primitive.from_proto(
                resource.target_sent_packets_per_second
            ),
            target_received_bytes_per_second=Primitive.from_proto(
                resource.target_received_bytes_per_second
            ),
            target_received_packets_per_second=Primitive.from_proto(
                resource.target_received_packets_per_second
            ),
        )


class VersionAutomaticScalingNetworkUtilizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            VersionAutomaticScalingNetworkUtilization.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            VersionAutomaticScalingNetworkUtilization.from_proto(i) for i in resources
        ]


class VersionAutomaticScalingStandardSchedulerSettings(object):
    def __init__(
        self,
        target_cpu_utilization: float = None,
        target_throughput_utilization: float = None,
        min_instances: int = None,
        max_instances: int = None,
    ):
        self.target_cpu_utilization = target_cpu_utilization
        self.target_throughput_utilization = target_throughput_utilization
        self.min_instances = min_instances
        self.max_instances = max_instances

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScalingStandardSchedulerSettings()
        if Primitive.to_proto(resource.target_cpu_utilization):
            res.target_cpu_utilization = Primitive.to_proto(
                resource.target_cpu_utilization
            )
        if Primitive.to_proto(resource.target_throughput_utilization):
            res.target_throughput_utilization = Primitive.to_proto(
                resource.target_throughput_utilization
            )
        if Primitive.to_proto(resource.min_instances):
            res.min_instances = Primitive.to_proto(resource.min_instances)
        if Primitive.to_proto(resource.max_instances):
            res.max_instances = Primitive.to_proto(resource.max_instances)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScalingStandardSchedulerSettings(
            target_cpu_utilization=Primitive.from_proto(
                resource.target_cpu_utilization
            ),
            target_throughput_utilization=Primitive.from_proto(
                resource.target_throughput_utilization
            ),
            min_instances=Primitive.from_proto(resource.min_instances),
            max_instances=Primitive.from_proto(resource.max_instances),
        )


class VersionAutomaticScalingStandardSchedulerSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            VersionAutomaticScalingStandardSchedulerSettings.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            VersionAutomaticScalingStandardSchedulerSettings.from_proto(i)
            for i in resources
        ]


class VersionBasicScaling(object):
    def __init__(self, idle_timeout: str = None, max_instances: int = None):
        self.idle_timeout = idle_timeout
        self.max_instances = max_instances

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionBasicScaling()
        if Primitive.to_proto(resource.idle_timeout):
            res.idle_timeout = Primitive.to_proto(resource.idle_timeout)
        if Primitive.to_proto(resource.max_instances):
            res.max_instances = Primitive.to_proto(resource.max_instances)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionBasicScaling(
            idle_timeout=Primitive.from_proto(resource.idle_timeout),
            max_instances=Primitive.from_proto(resource.max_instances),
        )


class VersionBasicScalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionBasicScaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionBasicScaling.from_proto(i) for i in resources]


class VersionManualScaling(object):
    def __init__(self, instances: int = None):
        self.instances = instances

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionManualScaling()
        if Primitive.to_proto(resource.instances):
            res.instances = Primitive.to_proto(resource.instances)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionManualScaling(instances=Primitive.from_proto(resource.instances),)


class VersionManualScalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionManualScaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionManualScaling.from_proto(i) for i in resources]


class VersionNetwork(object):
    def __init__(
        self,
        forwarded_ports: list = None,
        instance_tag: str = None,
        name: str = None,
        subnetwork_name: str = None,
        session_affinity: bool = None,
    ):
        self.forwarded_ports = forwarded_ports
        self.instance_tag = instance_tag
        self.name = name
        self.subnetwork_name = subnetwork_name
        self.session_affinity = session_affinity

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionNetwork()
        if Primitive.to_proto(resource.forwarded_ports):
            res.forwarded_ports.extend(Primitive.to_proto(resource.forwarded_ports))
        if Primitive.to_proto(resource.instance_tag):
            res.instance_tag = Primitive.to_proto(resource.instance_tag)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.subnetwork_name):
            res.subnetwork_name = Primitive.to_proto(resource.subnetwork_name)
        if Primitive.to_proto(resource.session_affinity):
            res.session_affinity = Primitive.to_proto(resource.session_affinity)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionNetwork(
            forwarded_ports=Primitive.from_proto(resource.forwarded_ports),
            instance_tag=Primitive.from_proto(resource.instance_tag),
            name=Primitive.from_proto(resource.name),
            subnetwork_name=Primitive.from_proto(resource.subnetwork_name),
            session_affinity=Primitive.from_proto(resource.session_affinity),
        )


class VersionNetworkArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionNetwork.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionNetwork.from_proto(i) for i in resources]


class VersionResources(object):
    def __init__(
        self,
        cpu: float = None,
        disk_gb: float = None,
        memory_gb: float = None,
        volumes: list = None,
    ):
        self.cpu = cpu
        self.disk_gb = disk_gb
        self.memory_gb = memory_gb
        self.volumes = volumes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionResources()
        if Primitive.to_proto(resource.cpu):
            res.cpu = Primitive.to_proto(resource.cpu)
        if Primitive.to_proto(resource.disk_gb):
            res.disk_gb = Primitive.to_proto(resource.disk_gb)
        if Primitive.to_proto(resource.memory_gb):
            res.memory_gb = Primitive.to_proto(resource.memory_gb)
        if VersionResourcesVolumesArray.to_proto(resource.volumes):
            res.volumes.extend(VersionResourcesVolumesArray.to_proto(resource.volumes))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionResources(
            cpu=Primitive.from_proto(resource.cpu),
            disk_gb=Primitive.from_proto(resource.disk_gb),
            memory_gb=Primitive.from_proto(resource.memory_gb),
            volumes=VersionResourcesVolumesArray.from_proto(resource.volumes),
        )


class VersionResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionResources.from_proto(i) for i in resources]


class VersionResourcesVolumes(object):
    def __init__(
        self, name: str = None, volume_type: str = None, size_gb: float = None
    ):
        self.name = name
        self.volume_type = volume_type
        self.size_gb = size_gb

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionResourcesVolumes()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.volume_type):
            res.volume_type = Primitive.to_proto(resource.volume_type)
        if Primitive.to_proto(resource.size_gb):
            res.size_gb = Primitive.to_proto(resource.size_gb)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionResourcesVolumes(
            name=Primitive.from_proto(resource.name),
            volume_type=Primitive.from_proto(resource.volume_type),
            size_gb=Primitive.from_proto(resource.size_gb),
        )


class VersionResourcesVolumesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionResourcesVolumes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionResourcesVolumes.from_proto(i) for i in resources]


class VersionHandlers(object):
    def __init__(
        self,
        url_regex: str = None,
        static_files: dict = None,
        script: dict = None,
        api_endpoint: dict = None,
        security_level: str = None,
        login: str = None,
        auth_fail_action: str = None,
        redirect_http_response_code: str = None,
    ):
        self.url_regex = url_regex
        self.static_files = static_files
        self.script = script
        self.api_endpoint = api_endpoint
        self.security_level = security_level
        self.login = login
        self.auth_fail_action = auth_fail_action
        self.redirect_http_response_code = redirect_http_response_code

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionHandlers()
        if Primitive.to_proto(resource.url_regex):
            res.url_regex = Primitive.to_proto(resource.url_regex)
        if VersionHandlersStaticFiles.to_proto(resource.static_files):
            res.static_files.CopyFrom(
                VersionHandlersStaticFiles.to_proto(resource.static_files)
            )
        else:
            res.ClearField("static_files")
        if VersionHandlersScript.to_proto(resource.script):
            res.script.CopyFrom(VersionHandlersScript.to_proto(resource.script))
        else:
            res.ClearField("script")
        if VersionHandlersApiEndpoint.to_proto(resource.api_endpoint):
            res.api_endpoint.CopyFrom(
                VersionHandlersApiEndpoint.to_proto(resource.api_endpoint)
            )
        else:
            res.ClearField("api_endpoint")
        if VersionHandlersSecurityLevelEnum.to_proto(resource.security_level):
            res.security_level = VersionHandlersSecurityLevelEnum.to_proto(
                resource.security_level
            )
        if VersionHandlersLoginEnum.to_proto(resource.login):
            res.login = VersionHandlersLoginEnum.to_proto(resource.login)
        if VersionHandlersAuthFailActionEnum.to_proto(resource.auth_fail_action):
            res.auth_fail_action = VersionHandlersAuthFailActionEnum.to_proto(
                resource.auth_fail_action
            )
        if VersionHandlersRedirectHttpResponseCodeEnum.to_proto(
            resource.redirect_http_response_code
        ):
            res.redirect_http_response_code = VersionHandlersRedirectHttpResponseCodeEnum.to_proto(
                resource.redirect_http_response_code
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionHandlers(
            url_regex=Primitive.from_proto(resource.url_regex),
            static_files=VersionHandlersStaticFiles.from_proto(resource.static_files),
            script=VersionHandlersScript.from_proto(resource.script),
            api_endpoint=VersionHandlersApiEndpoint.from_proto(resource.api_endpoint),
            security_level=VersionHandlersSecurityLevelEnum.from_proto(
                resource.security_level
            ),
            login=VersionHandlersLoginEnum.from_proto(resource.login),
            auth_fail_action=VersionHandlersAuthFailActionEnum.from_proto(
                resource.auth_fail_action
            ),
            redirect_http_response_code=VersionHandlersRedirectHttpResponseCodeEnum.from_proto(
                resource.redirect_http_response_code
            ),
        )


class VersionHandlersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionHandlers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionHandlers.from_proto(i) for i in resources]


class VersionHandlersStaticFiles(object):
    def __init__(
        self,
        path: str = None,
        upload_path_regex: str = None,
        http_headers: dict = None,
        mime_type: str = None,
        expiration: str = None,
        require_matching_file: bool = None,
        application_readable: bool = None,
    ):
        self.path = path
        self.upload_path_regex = upload_path_regex
        self.http_headers = http_headers
        self.mime_type = mime_type
        self.expiration = expiration
        self.require_matching_file = require_matching_file
        self.application_readable = application_readable

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionHandlersStaticFiles()
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.upload_path_regex):
            res.upload_path_regex = Primitive.to_proto(resource.upload_path_regex)
        if Primitive.to_proto(resource.http_headers):
            res.http_headers = Primitive.to_proto(resource.http_headers)
        if Primitive.to_proto(resource.mime_type):
            res.mime_type = Primitive.to_proto(resource.mime_type)
        if Primitive.to_proto(resource.expiration):
            res.expiration = Primitive.to_proto(resource.expiration)
        if Primitive.to_proto(resource.require_matching_file):
            res.require_matching_file = Primitive.to_proto(
                resource.require_matching_file
            )
        if Primitive.to_proto(resource.application_readable):
            res.application_readable = Primitive.to_proto(resource.application_readable)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionHandlersStaticFiles(
            path=Primitive.from_proto(resource.path),
            upload_path_regex=Primitive.from_proto(resource.upload_path_regex),
            http_headers=Primitive.from_proto(resource.http_headers),
            mime_type=Primitive.from_proto(resource.mime_type),
            expiration=Primitive.from_proto(resource.expiration),
            require_matching_file=Primitive.from_proto(resource.require_matching_file),
            application_readable=Primitive.from_proto(resource.application_readable),
        )


class VersionHandlersStaticFilesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionHandlersStaticFiles.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionHandlersStaticFiles.from_proto(i) for i in resources]


class VersionHandlersScript(object):
    def __init__(self, script_path: str = None):
        self.script_path = script_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionHandlersScript()
        if Primitive.to_proto(resource.script_path):
            res.script_path = Primitive.to_proto(resource.script_path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionHandlersScript(
            script_path=Primitive.from_proto(resource.script_path),
        )


class VersionHandlersScriptArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionHandlersScript.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionHandlersScript.from_proto(i) for i in resources]


class VersionHandlersApiEndpoint(object):
    def __init__(self, script_path: str = None):
        self.script_path = script_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionHandlersApiEndpoint()
        if Primitive.to_proto(resource.script_path):
            res.script_path = Primitive.to_proto(resource.script_path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionHandlersApiEndpoint(
            script_path=Primitive.from_proto(resource.script_path),
        )


class VersionHandlersApiEndpointArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionHandlersApiEndpoint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionHandlersApiEndpoint.from_proto(i) for i in resources]


class VersionErrorHandlers(object):
    def __init__(
        self, error_code: str = None, static_file: str = None, mime_type: str = None
    ):
        self.error_code = error_code
        self.static_file = static_file
        self.mime_type = mime_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionErrorHandlers()
        if VersionErrorHandlersErrorCodeEnum.to_proto(resource.error_code):
            res.error_code = VersionErrorHandlersErrorCodeEnum.to_proto(
                resource.error_code
            )
        if Primitive.to_proto(resource.static_file):
            res.static_file = Primitive.to_proto(resource.static_file)
        if Primitive.to_proto(resource.mime_type):
            res.mime_type = Primitive.to_proto(resource.mime_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionErrorHandlers(
            error_code=VersionErrorHandlersErrorCodeEnum.from_proto(
                resource.error_code
            ),
            static_file=Primitive.from_proto(resource.static_file),
            mime_type=Primitive.from_proto(resource.mime_type),
        )


class VersionErrorHandlersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionErrorHandlers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionErrorHandlers.from_proto(i) for i in resources]


class VersionLibraries(object):
    def __init__(self, name: str = None, version: str = None):
        self.name = name
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionLibraries()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionLibraries(
            name=Primitive.from_proto(resource.name),
            version=Primitive.from_proto(resource.version),
        )


class VersionLibrariesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionLibraries.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionLibraries.from_proto(i) for i in resources]


class VersionApiConfig(object):
    def __init__(
        self,
        auth_fail_action: str = None,
        login: str = None,
        script: str = None,
        security_level: str = None,
        url: str = None,
    ):
        self.auth_fail_action = auth_fail_action
        self.login = login
        self.script = script
        self.security_level = security_level
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionApiConfig()
        if VersionApiConfigAuthFailActionEnum.to_proto(resource.auth_fail_action):
            res.auth_fail_action = VersionApiConfigAuthFailActionEnum.to_proto(
                resource.auth_fail_action
            )
        if VersionApiConfigLoginEnum.to_proto(resource.login):
            res.login = VersionApiConfigLoginEnum.to_proto(resource.login)
        if Primitive.to_proto(resource.script):
            res.script = Primitive.to_proto(resource.script)
        if VersionApiConfigSecurityLevelEnum.to_proto(resource.security_level):
            res.security_level = VersionApiConfigSecurityLevelEnum.to_proto(
                resource.security_level
            )
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionApiConfig(
            auth_fail_action=VersionApiConfigAuthFailActionEnum.from_proto(
                resource.auth_fail_action
            ),
            login=VersionApiConfigLoginEnum.from_proto(resource.login),
            script=Primitive.from_proto(resource.script),
            security_level=VersionApiConfigSecurityLevelEnum.from_proto(
                resource.security_level
            ),
            url=Primitive.from_proto(resource.url),
        )


class VersionApiConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionApiConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionApiConfig.from_proto(i) for i in resources]


class VersionDeployment(object):
    def __init__(
        self,
        files: dict = None,
        container: dict = None,
        zip: dict = None,
        cloud_build_options: dict = None,
    ):
        self.files = files
        self.container = container
        self.zip = zip
        self.cloud_build_options = cloud_build_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionDeployment()
        if Primitive.to_proto(resource.files):
            res.files = Primitive.to_proto(resource.files)
        if VersionDeploymentContainer.to_proto(resource.container):
            res.container.CopyFrom(
                VersionDeploymentContainer.to_proto(resource.container)
            )
        else:
            res.ClearField("container")
        if VersionDeploymentZip.to_proto(resource.zip):
            res.zip.CopyFrom(VersionDeploymentZip.to_proto(resource.zip))
        else:
            res.ClearField("zip")
        if VersionDeploymentCloudBuildOptions.to_proto(resource.cloud_build_options):
            res.cloud_build_options.CopyFrom(
                VersionDeploymentCloudBuildOptions.to_proto(
                    resource.cloud_build_options
                )
            )
        else:
            res.ClearField("cloud_build_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionDeployment(
            files=Primitive.from_proto(resource.files),
            container=VersionDeploymentContainer.from_proto(resource.container),
            zip=VersionDeploymentZip.from_proto(resource.zip),
            cloud_build_options=VersionDeploymentCloudBuildOptions.from_proto(
                resource.cloud_build_options
            ),
        )


class VersionDeploymentArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionDeployment.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionDeployment.from_proto(i) for i in resources]


class VersionDeploymentFiles(object):
    def __init__(
        self, source_url: str = None, sha1_sum: str = None, mime_type: str = None
    ):
        self.source_url = source_url
        self.sha1_sum = sha1_sum
        self.mime_type = mime_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionDeploymentFiles()
        if Primitive.to_proto(resource.source_url):
            res.source_url = Primitive.to_proto(resource.source_url)
        if Primitive.to_proto(resource.sha1_sum):
            res.sha1_sum = Primitive.to_proto(resource.sha1_sum)
        if Primitive.to_proto(resource.mime_type):
            res.mime_type = Primitive.to_proto(resource.mime_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionDeploymentFiles(
            source_url=Primitive.from_proto(resource.source_url),
            sha1_sum=Primitive.from_proto(resource.sha1_sum),
            mime_type=Primitive.from_proto(resource.mime_type),
        )


class VersionDeploymentFilesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionDeploymentFiles.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionDeploymentFiles.from_proto(i) for i in resources]


class VersionDeploymentContainer(object):
    def __init__(self, image: str = None):
        self.image = image

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionDeploymentContainer()
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionDeploymentContainer(image=Primitive.from_proto(resource.image),)


class VersionDeploymentContainerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionDeploymentContainer.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionDeploymentContainer.from_proto(i) for i in resources]


class VersionDeploymentZip(object):
    def __init__(self, source_url: str = None, files_count: int = None):
        self.source_url = source_url
        self.files_count = files_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionDeploymentZip()
        if Primitive.to_proto(resource.source_url):
            res.source_url = Primitive.to_proto(resource.source_url)
        if Primitive.to_proto(resource.files_count):
            res.files_count = Primitive.to_proto(resource.files_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionDeploymentZip(
            source_url=Primitive.from_proto(resource.source_url),
            files_count=Primitive.from_proto(resource.files_count),
        )


class VersionDeploymentZipArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionDeploymentZip.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionDeploymentZip.from_proto(i) for i in resources]


class VersionDeploymentCloudBuildOptions(object):
    def __init__(self, app_yaml_path: str = None, cloud_build_timeout: str = None):
        self.app_yaml_path = app_yaml_path
        self.cloud_build_timeout = cloud_build_timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionDeploymentCloudBuildOptions()
        if Primitive.to_proto(resource.app_yaml_path):
            res.app_yaml_path = Primitive.to_proto(resource.app_yaml_path)
        if Primitive.to_proto(resource.cloud_build_timeout):
            res.cloud_build_timeout = Primitive.to_proto(resource.cloud_build_timeout)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionDeploymentCloudBuildOptions(
            app_yaml_path=Primitive.from_proto(resource.app_yaml_path),
            cloud_build_timeout=Primitive.from_proto(resource.cloud_build_timeout),
        )


class VersionDeploymentCloudBuildOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionDeploymentCloudBuildOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionDeploymentCloudBuildOptions.from_proto(i) for i in resources]


class VersionHealthCheck(object):
    def __init__(
        self,
        disable_health_check: bool = None,
        host: str = None,
        healthy_threshold: int = None,
        unhealthy_threshold: int = None,
        restart_threshold: int = None,
        check_interval: str = None,
        timeout: str = None,
    ):
        self.disable_health_check = disable_health_check
        self.host = host
        self.healthy_threshold = healthy_threshold
        self.unhealthy_threshold = unhealthy_threshold
        self.restart_threshold = restart_threshold
        self.check_interval = check_interval
        self.timeout = timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionHealthCheck()
        if Primitive.to_proto(resource.disable_health_check):
            res.disable_health_check = Primitive.to_proto(resource.disable_health_check)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.healthy_threshold):
            res.healthy_threshold = Primitive.to_proto(resource.healthy_threshold)
        if Primitive.to_proto(resource.unhealthy_threshold):
            res.unhealthy_threshold = Primitive.to_proto(resource.unhealthy_threshold)
        if Primitive.to_proto(resource.restart_threshold):
            res.restart_threshold = Primitive.to_proto(resource.restart_threshold)
        if Primitive.to_proto(resource.check_interval):
            res.check_interval = Primitive.to_proto(resource.check_interval)
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionHealthCheck(
            disable_health_check=Primitive.from_proto(resource.disable_health_check),
            host=Primitive.from_proto(resource.host),
            healthy_threshold=Primitive.from_proto(resource.healthy_threshold),
            unhealthy_threshold=Primitive.from_proto(resource.unhealthy_threshold),
            restart_threshold=Primitive.from_proto(resource.restart_threshold),
            check_interval=Primitive.from_proto(resource.check_interval),
            timeout=Primitive.from_proto(resource.timeout),
        )


class VersionHealthCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionHealthCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionHealthCheck.from_proto(i) for i in resources]


class VersionReadinessCheck(object):
    def __init__(
        self,
        path: str = None,
        host: str = None,
        failure_threshold: int = None,
        success_threshold: int = None,
        check_interval: str = None,
        timeout: str = None,
        app_start_timeout: str = None,
    ):
        self.path = path
        self.host = host
        self.failure_threshold = failure_threshold
        self.success_threshold = success_threshold
        self.check_interval = check_interval
        self.timeout = timeout
        self.app_start_timeout = app_start_timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionReadinessCheck()
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.failure_threshold):
            res.failure_threshold = Primitive.to_proto(resource.failure_threshold)
        if Primitive.to_proto(resource.success_threshold):
            res.success_threshold = Primitive.to_proto(resource.success_threshold)
        if Primitive.to_proto(resource.check_interval):
            res.check_interval = Primitive.to_proto(resource.check_interval)
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        if Primitive.to_proto(resource.app_start_timeout):
            res.app_start_timeout = Primitive.to_proto(resource.app_start_timeout)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionReadinessCheck(
            path=Primitive.from_proto(resource.path),
            host=Primitive.from_proto(resource.host),
            failure_threshold=Primitive.from_proto(resource.failure_threshold),
            success_threshold=Primitive.from_proto(resource.success_threshold),
            check_interval=Primitive.from_proto(resource.check_interval),
            timeout=Primitive.from_proto(resource.timeout),
            app_start_timeout=Primitive.from_proto(resource.app_start_timeout),
        )


class VersionReadinessCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionReadinessCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionReadinessCheck.from_proto(i) for i in resources]


class VersionLivenessCheck(object):
    def __init__(
        self,
        path: str = None,
        host: str = None,
        failure_threshold: int = None,
        success_threshold: int = None,
        check_interval: str = None,
        timeout: str = None,
        initial_delay: str = None,
    ):
        self.path = path
        self.host = host
        self.failure_threshold = failure_threshold
        self.success_threshold = success_threshold
        self.check_interval = check_interval
        self.timeout = timeout
        self.initial_delay = initial_delay

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionLivenessCheck()
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.failure_threshold):
            res.failure_threshold = Primitive.to_proto(resource.failure_threshold)
        if Primitive.to_proto(resource.success_threshold):
            res.success_threshold = Primitive.to_proto(resource.success_threshold)
        if Primitive.to_proto(resource.check_interval):
            res.check_interval = Primitive.to_proto(resource.check_interval)
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        if Primitive.to_proto(resource.initial_delay):
            res.initial_delay = Primitive.to_proto(resource.initial_delay)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionLivenessCheck(
            path=Primitive.from_proto(resource.path),
            host=Primitive.from_proto(resource.host),
            failure_threshold=Primitive.from_proto(resource.failure_threshold),
            success_threshold=Primitive.from_proto(resource.success_threshold),
            check_interval=Primitive.from_proto(resource.check_interval),
            timeout=Primitive.from_proto(resource.timeout),
            initial_delay=Primitive.from_proto(resource.initial_delay),
        )


class VersionLivenessCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionLivenessCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionLivenessCheck.from_proto(i) for i in resources]


class VersionEntrypoint(object):
    def __init__(self, shell: str = None):
        self.shell = shell

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionEntrypoint()
        if Primitive.to_proto(resource.shell):
            res.shell = Primitive.to_proto(resource.shell)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionEntrypoint(shell=Primitive.from_proto(resource.shell),)


class VersionEntrypointArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionEntrypoint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionEntrypoint.from_proto(i) for i in resources]


class VersionVPCAccessConnector(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionVPCAccessConnector()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionVPCAccessConnector(name=Primitive.from_proto(resource.name),)


class VersionVPCAccessConnectorArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionVPCAccessConnector.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionVPCAccessConnector.from_proto(i) for i in resources]


class VersionInboundServicesEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionInboundServicesEnum.Value(
            "AppengineVersionInboundServicesEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionInboundServicesEnum.Name(resource)[
            len("AppengineVersionInboundServicesEnum") :
        ]


class VersionServingStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionServingStatusEnum.Value(
            "AppengineVersionServingStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionServingStatusEnum.Name(resource)[
            len("AppengineVersionServingStatusEnum") :
        ]


class VersionHandlersSecurityLevelEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersSecurityLevelEnum.Value(
            "AppengineVersionHandlersSecurityLevelEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersSecurityLevelEnum.Name(resource)[
            len("AppengineVersionHandlersSecurityLevelEnum") :
        ]


class VersionHandlersLoginEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersLoginEnum.Value(
            "AppengineVersionHandlersLoginEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersLoginEnum.Name(resource)[
            len("AppengineVersionHandlersLoginEnum") :
        ]


class VersionHandlersAuthFailActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersAuthFailActionEnum.Value(
            "AppengineVersionHandlersAuthFailActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersAuthFailActionEnum.Name(resource)[
            len("AppengineVersionHandlersAuthFailActionEnum") :
        ]


class VersionHandlersRedirectHttpResponseCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersRedirectHttpResponseCodeEnum.Value(
            "AppengineVersionHandlersRedirectHttpResponseCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersRedirectHttpResponseCodeEnum.Name(
            resource
        )[len("AppengineVersionHandlersRedirectHttpResponseCodeEnum") :]


class VersionErrorHandlersErrorCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionErrorHandlersErrorCodeEnum.Value(
            "AppengineVersionErrorHandlersErrorCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionErrorHandlersErrorCodeEnum.Name(resource)[
            len("AppengineVersionErrorHandlersErrorCodeEnum") :
        ]


class VersionApiConfigAuthFailActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigAuthFailActionEnum.Value(
            "AppengineVersionApiConfigAuthFailActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigAuthFailActionEnum.Name(resource)[
            len("AppengineVersionApiConfigAuthFailActionEnum") :
        ]


class VersionApiConfigLoginEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigLoginEnum.Value(
            "AppengineVersionApiConfigLoginEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigLoginEnum.Name(resource)[
            len("AppengineVersionApiConfigLoginEnum") :
        ]


class VersionApiConfigSecurityLevelEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigSecurityLevelEnum.Value(
            "AppengineVersionApiConfigSecurityLevelEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigSecurityLevelEnum.Name(resource)[
            len("AppengineVersionApiConfigSecurityLevelEnum") :
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
