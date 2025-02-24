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
from google3.cloud.graphite.mmv2.services.google.cloud_functions import function_pb2
from google3.cloud.graphite.mmv2.services.google.cloud_functions import (
    function_pb2_grpc,
)

from typing import List


class Function(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        source_archive_url: str = None,
        source_repository: dict = None,
        https_trigger: dict = None,
        event_trigger: dict = None,
        status: str = None,
        entry_point: str = None,
        runtime: str = None,
        timeout: str = None,
        available_memory_mb: int = None,
        service_account_email: str = None,
        update_time: str = None,
        version_id: int = None,
        labels: dict = None,
        environment_variables: dict = None,
        max_instances: int = None,
        vpc_connector: str = None,
        vpc_connector_egress_settings: str = None,
        ingress_settings: str = None,
        region: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.source_archive_url = source_archive_url
        self.source_repository = source_repository
        self.https_trigger = https_trigger
        self.event_trigger = event_trigger
        self.entry_point = entry_point
        self.runtime = runtime
        self.timeout = timeout
        self.available_memory_mb = available_memory_mb
        self.service_account_email = service_account_email
        self.labels = labels
        self.environment_variables = environment_variables
        self.max_instances = max_instances
        self.vpc_connector = vpc_connector
        self.vpc_connector_egress_settings = vpc_connector_egress_settings
        self.ingress_settings = ingress_settings
        self.region = region
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = function_pb2_grpc.CloudfunctionsFunctionServiceStub(channel.Channel())
        request = function_pb2.ApplyCloudfunctionsFunctionRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.source_archive_url):
            request.resource.source_archive_url = Primitive.to_proto(
                self.source_archive_url
            )

        if FunctionSourceRepository.to_proto(self.source_repository):
            request.resource.source_repository.CopyFrom(
                FunctionSourceRepository.to_proto(self.source_repository)
            )
        else:
            request.resource.ClearField("source_repository")
        if FunctionHttpsTrigger.to_proto(self.https_trigger):
            request.resource.https_trigger.CopyFrom(
                FunctionHttpsTrigger.to_proto(self.https_trigger)
            )
        else:
            request.resource.ClearField("https_trigger")
        if FunctionEventTrigger.to_proto(self.event_trigger):
            request.resource.event_trigger.CopyFrom(
                FunctionEventTrigger.to_proto(self.event_trigger)
            )
        else:
            request.resource.ClearField("event_trigger")
        if Primitive.to_proto(self.entry_point):
            request.resource.entry_point = Primitive.to_proto(self.entry_point)

        if Primitive.to_proto(self.runtime):
            request.resource.runtime = Primitive.to_proto(self.runtime)

        if Primitive.to_proto(self.timeout):
            request.resource.timeout = Primitive.to_proto(self.timeout)

        if Primitive.to_proto(self.available_memory_mb):
            request.resource.available_memory_mb = Primitive.to_proto(
                self.available_memory_mb
            )

        if Primitive.to_proto(self.service_account_email):
            request.resource.service_account_email = Primitive.to_proto(
                self.service_account_email
            )

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.environment_variables):
            request.resource.environment_variables = Primitive.to_proto(
                self.environment_variables
            )

        if Primitive.to_proto(self.max_instances):
            request.resource.max_instances = Primitive.to_proto(self.max_instances)

        if Primitive.to_proto(self.vpc_connector):
            request.resource.vpc_connector = Primitive.to_proto(self.vpc_connector)

        if FunctionVPCConnectorEgressSettingsEnum.to_proto(
            self.vpc_connector_egress_settings
        ):
            request.resource.vpc_connector_egress_settings = (
                FunctionVPCConnectorEgressSettingsEnum.to_proto(
                    self.vpc_connector_egress_settings
                )
            )

        if FunctionIngressSettingsEnum.to_proto(self.ingress_settings):
            request.resource.ingress_settings = FunctionIngressSettingsEnum.to_proto(
                self.ingress_settings
            )

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudfunctionsFunction(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.source_archive_url = Primitive.from_proto(response.source_archive_url)
        self.source_repository = FunctionSourceRepository.from_proto(
            response.source_repository
        )
        self.https_trigger = FunctionHttpsTrigger.from_proto(response.https_trigger)
        self.event_trigger = FunctionEventTrigger.from_proto(response.event_trigger)
        self.status = FunctionStatusEnum.from_proto(response.status)
        self.entry_point = Primitive.from_proto(response.entry_point)
        self.runtime = Primitive.from_proto(response.runtime)
        self.timeout = Primitive.from_proto(response.timeout)
        self.available_memory_mb = Primitive.from_proto(response.available_memory_mb)
        self.service_account_email = Primitive.from_proto(
            response.service_account_email
        )
        self.update_time = Primitive.from_proto(response.update_time)
        self.version_id = Primitive.from_proto(response.version_id)
        self.labels = Primitive.from_proto(response.labels)
        self.environment_variables = Primitive.from_proto(
            response.environment_variables
        )
        self.max_instances = Primitive.from_proto(response.max_instances)
        self.vpc_connector = Primitive.from_proto(response.vpc_connector)
        self.vpc_connector_egress_settings = (
            FunctionVPCConnectorEgressSettingsEnum.from_proto(
                response.vpc_connector_egress_settings
            )
        )
        self.ingress_settings = FunctionIngressSettingsEnum.from_proto(
            response.ingress_settings
        )
        self.region = Primitive.from_proto(response.region)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = function_pb2_grpc.CloudfunctionsFunctionServiceStub(channel.Channel())
        request = function_pb2.DeleteCloudfunctionsFunctionRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.source_archive_url):
            request.resource.source_archive_url = Primitive.to_proto(
                self.source_archive_url
            )

        if FunctionSourceRepository.to_proto(self.source_repository):
            request.resource.source_repository.CopyFrom(
                FunctionSourceRepository.to_proto(self.source_repository)
            )
        else:
            request.resource.ClearField("source_repository")
        if FunctionHttpsTrigger.to_proto(self.https_trigger):
            request.resource.https_trigger.CopyFrom(
                FunctionHttpsTrigger.to_proto(self.https_trigger)
            )
        else:
            request.resource.ClearField("https_trigger")
        if FunctionEventTrigger.to_proto(self.event_trigger):
            request.resource.event_trigger.CopyFrom(
                FunctionEventTrigger.to_proto(self.event_trigger)
            )
        else:
            request.resource.ClearField("event_trigger")
        if Primitive.to_proto(self.entry_point):
            request.resource.entry_point = Primitive.to_proto(self.entry_point)

        if Primitive.to_proto(self.runtime):
            request.resource.runtime = Primitive.to_proto(self.runtime)

        if Primitive.to_proto(self.timeout):
            request.resource.timeout = Primitive.to_proto(self.timeout)

        if Primitive.to_proto(self.available_memory_mb):
            request.resource.available_memory_mb = Primitive.to_proto(
                self.available_memory_mb
            )

        if Primitive.to_proto(self.service_account_email):
            request.resource.service_account_email = Primitive.to_proto(
                self.service_account_email
            )

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.environment_variables):
            request.resource.environment_variables = Primitive.to_proto(
                self.environment_variables
            )

        if Primitive.to_proto(self.max_instances):
            request.resource.max_instances = Primitive.to_proto(self.max_instances)

        if Primitive.to_proto(self.vpc_connector):
            request.resource.vpc_connector = Primitive.to_proto(self.vpc_connector)

        if FunctionVPCConnectorEgressSettingsEnum.to_proto(
            self.vpc_connector_egress_settings
        ):
            request.resource.vpc_connector_egress_settings = (
                FunctionVPCConnectorEgressSettingsEnum.to_proto(
                    self.vpc_connector_egress_settings
                )
            )

        if FunctionIngressSettingsEnum.to_proto(self.ingress_settings):
            request.resource.ingress_settings = FunctionIngressSettingsEnum.to_proto(
                self.ingress_settings
            )

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteCloudfunctionsFunction(request)

    @classmethod
    def list(self, project, region, service_account_file=""):
        stub = function_pb2_grpc.CloudfunctionsFunctionServiceStub(channel.Channel())
        request = function_pb2.ListCloudfunctionsFunctionRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        return stub.ListCloudfunctionsFunction(request).items

    def to_proto(self):
        resource = function_pb2.CloudfunctionsFunction()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.source_archive_url):
            resource.source_archive_url = Primitive.to_proto(self.source_archive_url)
        if FunctionSourceRepository.to_proto(self.source_repository):
            resource.source_repository.CopyFrom(
                FunctionSourceRepository.to_proto(self.source_repository)
            )
        else:
            resource.ClearField("source_repository")
        if FunctionHttpsTrigger.to_proto(self.https_trigger):
            resource.https_trigger.CopyFrom(
                FunctionHttpsTrigger.to_proto(self.https_trigger)
            )
        else:
            resource.ClearField("https_trigger")
        if FunctionEventTrigger.to_proto(self.event_trigger):
            resource.event_trigger.CopyFrom(
                FunctionEventTrigger.to_proto(self.event_trigger)
            )
        else:
            resource.ClearField("event_trigger")
        if Primitive.to_proto(self.entry_point):
            resource.entry_point = Primitive.to_proto(self.entry_point)
        if Primitive.to_proto(self.runtime):
            resource.runtime = Primitive.to_proto(self.runtime)
        if Primitive.to_proto(self.timeout):
            resource.timeout = Primitive.to_proto(self.timeout)
        if Primitive.to_proto(self.available_memory_mb):
            resource.available_memory_mb = Primitive.to_proto(self.available_memory_mb)
        if Primitive.to_proto(self.service_account_email):
            resource.service_account_email = Primitive.to_proto(
                self.service_account_email
            )
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.environment_variables):
            resource.environment_variables = Primitive.to_proto(
                self.environment_variables
            )
        if Primitive.to_proto(self.max_instances):
            resource.max_instances = Primitive.to_proto(self.max_instances)
        if Primitive.to_proto(self.vpc_connector):
            resource.vpc_connector = Primitive.to_proto(self.vpc_connector)
        if FunctionVPCConnectorEgressSettingsEnum.to_proto(
            self.vpc_connector_egress_settings
        ):
            resource.vpc_connector_egress_settings = (
                FunctionVPCConnectorEgressSettingsEnum.to_proto(
                    self.vpc_connector_egress_settings
                )
            )
        if FunctionIngressSettingsEnum.to_proto(self.ingress_settings):
            resource.ingress_settings = FunctionIngressSettingsEnum.to_proto(
                self.ingress_settings
            )
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class FunctionSourceRepository(object):
    def __init__(self, url: str = None, deployed_url: str = None):
        self.url = url
        self.deployed_url = deployed_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = function_pb2.CloudfunctionsFunctionSourceRepository()
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if Primitive.to_proto(resource.deployed_url):
            res.deployed_url = Primitive.to_proto(resource.deployed_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FunctionSourceRepository(
            url=Primitive.from_proto(resource.url),
            deployed_url=Primitive.from_proto(resource.deployed_url),
        )


class FunctionSourceRepositoryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FunctionSourceRepository.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FunctionSourceRepository.from_proto(i) for i in resources]


class FunctionHttpsTrigger(object):
    def __init__(self, url: str = None, security_level: str = None):
        self.url = url
        self.security_level = security_level

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = function_pb2.CloudfunctionsFunctionHttpsTrigger()
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if FunctionHttpsTriggerSecurityLevelEnum.to_proto(resource.security_level):
            res.security_level = FunctionHttpsTriggerSecurityLevelEnum.to_proto(
                resource.security_level
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FunctionHttpsTrigger(
            url=Primitive.from_proto(resource.url),
            security_level=FunctionHttpsTriggerSecurityLevelEnum.from_proto(
                resource.security_level
            ),
        )


class FunctionHttpsTriggerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FunctionHttpsTrigger.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FunctionHttpsTrigger.from_proto(i) for i in resources]


class FunctionEventTrigger(object):
    def __init__(
        self,
        event_type: str = None,
        resource: str = None,
        service: str = None,
        failure_policy: bool = None,
    ):
        self.event_type = event_type
        self.resource = resource
        self.service = service
        self.failure_policy = failure_policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = function_pb2.CloudfunctionsFunctionEventTrigger()
        if Primitive.to_proto(resource.event_type):
            res.event_type = Primitive.to_proto(resource.event_type)
        if Primitive.to_proto(resource.resource):
            res.resource = Primitive.to_proto(resource.resource)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.failure_policy):
            res.failure_policy = Primitive.to_proto(resource.failure_policy)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FunctionEventTrigger(
            event_type=Primitive.from_proto(resource.event_type),
            resource=Primitive.from_proto(resource.resource),
            service=Primitive.from_proto(resource.service),
            failure_policy=Primitive.from_proto(resource.failure_policy),
        )


class FunctionEventTriggerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FunctionEventTrigger.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FunctionEventTrigger.from_proto(i) for i in resources]


class FunctionHttpsTriggerSecurityLevelEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return function_pb2.CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum.Value(
            "CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return function_pb2.CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum.Name(
            resource
        )[len("CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum") :]


class FunctionStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return function_pb2.CloudfunctionsFunctionStatusEnum.Value(
            "CloudfunctionsFunctionStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return function_pb2.CloudfunctionsFunctionStatusEnum.Name(resource)[
            len("CloudfunctionsFunctionStatusEnum") :
        ]


class FunctionVPCConnectorEgressSettingsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return function_pb2.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum.Value(
            "CloudfunctionsFunctionVPCConnectorEgressSettingsEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return function_pb2.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum.Name(
            resource
        )[len("CloudfunctionsFunctionVPCConnectorEgressSettingsEnum") :]


class FunctionIngressSettingsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return function_pb2.CloudfunctionsFunctionIngressSettingsEnum.Value(
            "CloudfunctionsFunctionIngressSettingsEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return function_pb2.CloudfunctionsFunctionIngressSettingsEnum.Name(resource)[
            len("CloudfunctionsFunctionIngressSettingsEnum") :
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
