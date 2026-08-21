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
from google3.cloud.graphite.mmv2.services.google.cloudfunctions import (
    cloud_function_pb2,
)
from google3.cloud.graphite.mmv2.services.google.cloudfunctions import (
    cloud_function_pb2_grpc,
)

from typing import List


class CloudFunction(object):
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
        timeout: int = None,
        available_memory_mb: int = None,
        service_account_email: str = None,
        update_time: str = None,
        version_id: int = None,
        labels: dict = None,
        environment_variables: dict = None,
        network: str = None,
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
        self.network = network
        self.max_instances = max_instances
        self.vpc_connector = vpc_connector
        self.vpc_connector_egress_settings = vpc_connector_egress_settings
        self.ingress_settings = ingress_settings
        self.region = region
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = cloud_function_pb2_grpc.CloudfunctionsCloudFunctionServiceStub(
            channel.Channel()
        )
        request = cloud_function_pb2.ApplyCloudfunctionsCloudFunctionRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.source_archive_url):
            request.resource.source_archive_url = Primitive.to_proto(
                self.source_archive_url
            )

        if CloudFunctionSourceRepository.to_proto(self.source_repository):
            request.resource.source_repository.CopyFrom(
                CloudFunctionSourceRepository.to_proto(self.source_repository)
            )
        else:
            request.resource.ClearField("source_repository")
        if CloudFunctionHttpsTrigger.to_proto(self.https_trigger):
            request.resource.https_trigger.CopyFrom(
                CloudFunctionHttpsTrigger.to_proto(self.https_trigger)
            )
        else:
            request.resource.ClearField("https_trigger")
        if CloudFunctionEventTrigger.to_proto(self.event_trigger):
            request.resource.event_trigger.CopyFrom(
                CloudFunctionEventTrigger.to_proto(self.event_trigger)
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

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.max_instances):
            request.resource.max_instances = Primitive.to_proto(self.max_instances)

        if Primitive.to_proto(self.vpc_connector):
            request.resource.vpc_connector = Primitive.to_proto(self.vpc_connector)

        if CloudFunctionVPCConnectorEgressSettingsEnum.to_proto(
            self.vpc_connector_egress_settings
        ):
            request.resource.vpc_connector_egress_settings = CloudFunctionVPCConnectorEgressSettingsEnum.to_proto(
                self.vpc_connector_egress_settings
            )

        if CloudFunctionIngressSettingsEnum.to_proto(self.ingress_settings):
            request.resource.ingress_settings = CloudFunctionIngressSettingsEnum.to_proto(
                self.ingress_settings
            )

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudfunctionsCloudFunction(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.source_archive_url = Primitive.from_proto(response.source_archive_url)
        self.source_repository = CloudFunctionSourceRepository.from_proto(
            response.source_repository
        )
        self.https_trigger = CloudFunctionHttpsTrigger.from_proto(
            response.https_trigger
        )
        self.event_trigger = CloudFunctionEventTrigger.from_proto(
            response.event_trigger
        )
        self.status = CloudFunctionStatusEnum.from_proto(response.status)
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
        self.network = Primitive.from_proto(response.network)
        self.max_instances = Primitive.from_proto(response.max_instances)
        self.vpc_connector = Primitive.from_proto(response.vpc_connector)
        self.vpc_connector_egress_settings = CloudFunctionVPCConnectorEgressSettingsEnum.from_proto(
            response.vpc_connector_egress_settings
        )
        self.ingress_settings = CloudFunctionIngressSettingsEnum.from_proto(
            response.ingress_settings
        )
        self.region = Primitive.from_proto(response.region)
        self.project = Primitive.from_proto(response.project)

    @classmethod
    def delete(self, project, region, name, service_account_file=""):
        stub = cloud_function_pb2_grpc.CloudfunctionsCloudFunctionServiceStub(
            channel.Channel()
        )
        request = cloud_function_pb2.DeleteCloudfunctionsCloudFunctionRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        request.Name = name

        response = stub.DeleteCloudfunctionsCloudFunction(request)

    @classmethod
    def list(self, project, region, service_account_file=""):
        stub = cloud_function_pb2_grpc.CloudfunctionsCloudFunctionServiceStub(
            channel.Channel()
        )
        request = cloud_function_pb2.ListCloudfunctionsCloudFunctionRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        return stub.ListCloudfunctionsCloudFunction(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = cloud_function_pb2.CloudfunctionsCloudFunction()
        any_proto.Unpack(res_proto)

        res = CloudFunction()
        res.name = Primitive.from_proto(res_proto.name)
        res.description = Primitive.from_proto(res_proto.description)
        res.source_archive_url = Primitive.from_proto(res_proto.source_archive_url)
        res.source_repository = CloudFunctionSourceRepository.from_proto(
            res_proto.source_repository
        )
        res.https_trigger = CloudFunctionHttpsTrigger.from_proto(
            res_proto.https_trigger
        )
        res.event_trigger = CloudFunctionEventTrigger.from_proto(
            res_proto.event_trigger
        )
        res.status = CloudFunctionStatusEnum.from_proto(res_proto.status)
        res.entry_point = Primitive.from_proto(res_proto.entry_point)
        res.runtime = Primitive.from_proto(res_proto.runtime)
        res.timeout = Primitive.from_proto(res_proto.timeout)
        res.available_memory_mb = Primitive.from_proto(res_proto.available_memory_mb)
        res.service_account_email = Primitive.from_proto(
            res_proto.service_account_email
        )
        res.update_time = Primitive.from_proto(res_proto.update_time)
        res.version_id = Primitive.from_proto(res_proto.version_id)
        res.labels = Primitive.from_proto(res_proto.labels)
        res.environment_variables = Primitive.from_proto(
            res_proto.environment_variables
        )
        res.network = Primitive.from_proto(res_proto.network)
        res.max_instances = Primitive.from_proto(res_proto.max_instances)
        res.vpc_connector = Primitive.from_proto(res_proto.vpc_connector)
        res.vpc_connector_egress_settings = CloudFunctionVPCConnectorEgressSettingsEnum.from_proto(
            res_proto.vpc_connector_egress_settings
        )
        res.ingress_settings = CloudFunctionIngressSettingsEnum.from_proto(
            res_proto.ingress_settings
        )
        res.region = Primitive.from_proto(res_proto.region)
        res.project = Primitive.from_proto(res_proto.project)
        return res


class CloudFunctionSourceRepository(object):
    def __init__(self, url: str = None, deployed_url: str = None):
        self.url = url
        self.deployed_url = deployed_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cloud_function_pb2.CloudfunctionsCloudFunctionSourceRepository()
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if Primitive.to_proto(resource.deployed_url):
            res.deployed_url = Primitive.to_proto(resource.deployed_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CloudFunctionSourceRepository(
            url=resource.url, deployed_url=resource.deployed_url,
        )


class CloudFunctionSourceRepositoryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CloudFunctionSourceRepository.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CloudFunctionSourceRepository.from_proto(i) for i in resources]


class CloudFunctionHttpsTrigger(object):
    def __init__(self, url: str = None):
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cloud_function_pb2.CloudfunctionsCloudFunctionHttpsTrigger()
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CloudFunctionHttpsTrigger(url=resource.url,)


class CloudFunctionHttpsTriggerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CloudFunctionHttpsTrigger.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CloudFunctionHttpsTrigger.from_proto(i) for i in resources]


class CloudFunctionEventTrigger(object):
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

        res = cloud_function_pb2.CloudfunctionsCloudFunctionEventTrigger()
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

        return CloudFunctionEventTrigger(
            event_type=resource.event_type,
            resource=resource.resource,
            service=resource.service,
            failure_policy=resource.failure_policy,
        )


class CloudFunctionEventTriggerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CloudFunctionEventTrigger.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CloudFunctionEventTrigger.from_proto(i) for i in resources]


class CloudFunctionStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cloud_function_pb2.CloudfunctionsCloudFunctionStatusEnum.Value(
            "CloudfunctionsCloudFunctionStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cloud_function_pb2.CloudfunctionsCloudFunctionStatusEnum.Name(resource)[
            len("CloudfunctionsCloudFunctionStatusEnum") :
        ]


class CloudFunctionVPCConnectorEgressSettingsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cloud_function_pb2.CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum.Value(
            "CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cloud_function_pb2.CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum.Name(
            resource
        )[
            len("CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum") :
        ]


class CloudFunctionIngressSettingsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cloud_function_pb2.CloudfunctionsCloudFunctionIngressSettingsEnum.Value(
            "CloudfunctionsCloudFunctionIngressSettingsEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cloud_function_pb2.CloudfunctionsCloudFunctionIngressSettingsEnum.Name(
            resource
        )[len("CloudfunctionsCloudFunctionIngressSettingsEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
