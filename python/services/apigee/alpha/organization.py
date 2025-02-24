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
from google3.cloud.graphite.mmv2.services.google.apigee import organization_pb2
from google3.cloud.graphite.mmv2.services.google.apigee import organization_pb2_grpc

from typing import List


class Organization(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        description: str = None,
        created_at: int = None,
        last_modified_at: int = None,
        expires_at: int = None,
        environments: list = None,
        properties: dict = None,
        analytics_region: str = None,
        authorized_network: str = None,
        runtime_type: str = None,
        subscription_type: str = None,
        billing_type: str = None,
        addons_config: dict = None,
        ca_certificate: str = None,
        runtime_database_encryption_key_name: str = None,
        project_id: str = None,
        state: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.description = description
        self.properties = properties
        self.analytics_region = analytics_region
        self.authorized_network = authorized_network
        self.runtime_type = runtime_type
        self.addons_config = addons_config
        self.runtime_database_encryption_key_name = runtime_database_encryption_key_name
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = organization_pb2_grpc.ApigeeAlphaOrganizationServiceStub(
            channel.Channel()
        )
        request = organization_pb2.ApplyApigeeAlphaOrganizationRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.properties):
            request.resource.properties = Primitive.to_proto(self.properties)

        if Primitive.to_proto(self.analytics_region):
            request.resource.analytics_region = Primitive.to_proto(
                self.analytics_region
            )

        if Primitive.to_proto(self.authorized_network):
            request.resource.authorized_network = Primitive.to_proto(
                self.authorized_network
            )

        if OrganizationRuntimeTypeEnum.to_proto(self.runtime_type):
            request.resource.runtime_type = OrganizationRuntimeTypeEnum.to_proto(
                self.runtime_type
            )

        if OrganizationAddonsConfig.to_proto(self.addons_config):
            request.resource.addons_config.CopyFrom(
                OrganizationAddonsConfig.to_proto(self.addons_config)
            )
        else:
            request.resource.ClearField("addons_config")
        if Primitive.to_proto(self.runtime_database_encryption_key_name):
            request.resource.runtime_database_encryption_key_name = Primitive.to_proto(
                self.runtime_database_encryption_key_name
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyApigeeAlphaOrganization(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.created_at = Primitive.from_proto(response.created_at)
        self.last_modified_at = Primitive.from_proto(response.last_modified_at)
        self.expires_at = Primitive.from_proto(response.expires_at)
        self.environments = Primitive.from_proto(response.environments)
        self.properties = Primitive.from_proto(response.properties)
        self.analytics_region = Primitive.from_proto(response.analytics_region)
        self.authorized_network = Primitive.from_proto(response.authorized_network)
        self.runtime_type = OrganizationRuntimeTypeEnum.from_proto(
            response.runtime_type
        )
        self.subscription_type = OrganizationSubscriptionTypeEnum.from_proto(
            response.subscription_type
        )
        self.billing_type = OrganizationBillingTypeEnum.from_proto(
            response.billing_type
        )
        self.addons_config = OrganizationAddonsConfig.from_proto(response.addons_config)
        self.ca_certificate = Primitive.from_proto(response.ca_certificate)
        self.runtime_database_encryption_key_name = Primitive.from_proto(
            response.runtime_database_encryption_key_name
        )
        self.project_id = Primitive.from_proto(response.project_id)
        self.state = OrganizationStateEnum.from_proto(response.state)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = organization_pb2_grpc.ApigeeAlphaOrganizationServiceStub(
            channel.Channel()
        )
        request = organization_pb2.DeleteApigeeAlphaOrganizationRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.properties):
            request.resource.properties = Primitive.to_proto(self.properties)

        if Primitive.to_proto(self.analytics_region):
            request.resource.analytics_region = Primitive.to_proto(
                self.analytics_region
            )

        if Primitive.to_proto(self.authorized_network):
            request.resource.authorized_network = Primitive.to_proto(
                self.authorized_network
            )

        if OrganizationRuntimeTypeEnum.to_proto(self.runtime_type):
            request.resource.runtime_type = OrganizationRuntimeTypeEnum.to_proto(
                self.runtime_type
            )

        if OrganizationAddonsConfig.to_proto(self.addons_config):
            request.resource.addons_config.CopyFrom(
                OrganizationAddonsConfig.to_proto(self.addons_config)
            )
        else:
            request.resource.ClearField("addons_config")
        if Primitive.to_proto(self.runtime_database_encryption_key_name):
            request.resource.runtime_database_encryption_key_name = Primitive.to_proto(
                self.runtime_database_encryption_key_name
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteApigeeAlphaOrganization(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = organization_pb2_grpc.ApigeeAlphaOrganizationServiceStub(
            channel.Channel()
        )
        request = organization_pb2.ListApigeeAlphaOrganizationRequest()
        request.service_account_file = service_account_file
        return stub.ListApigeeAlphaOrganization(request).items

    def to_proto(self):
        resource = organization_pb2.ApigeeAlphaOrganization()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.properties):
            resource.properties = Primitive.to_proto(self.properties)
        if Primitive.to_proto(self.analytics_region):
            resource.analytics_region = Primitive.to_proto(self.analytics_region)
        if Primitive.to_proto(self.authorized_network):
            resource.authorized_network = Primitive.to_proto(self.authorized_network)
        if OrganizationRuntimeTypeEnum.to_proto(self.runtime_type):
            resource.runtime_type = OrganizationRuntimeTypeEnum.to_proto(
                self.runtime_type
            )
        if OrganizationAddonsConfig.to_proto(self.addons_config):
            resource.addons_config.CopyFrom(
                OrganizationAddonsConfig.to_proto(self.addons_config)
            )
        else:
            resource.ClearField("addons_config")
        if Primitive.to_proto(self.runtime_database_encryption_key_name):
            resource.runtime_database_encryption_key_name = Primitive.to_proto(
                self.runtime_database_encryption_key_name
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class OrganizationAddonsConfig(object):
    def __init__(
        self, advanced_api_ops_config: dict = None, monetization_config: dict = None
    ):
        self.advanced_api_ops_config = advanced_api_ops_config
        self.monetization_config = monetization_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = organization_pb2.ApigeeAlphaOrganizationAddonsConfig()
        if OrganizationAddonsConfigAdvancedApiOpsConfig.to_proto(
            resource.advanced_api_ops_config
        ):
            res.advanced_api_ops_config.CopyFrom(
                OrganizationAddonsConfigAdvancedApiOpsConfig.to_proto(
                    resource.advanced_api_ops_config
                )
            )
        else:
            res.ClearField("advanced_api_ops_config")
        if OrganizationAddonsConfigMonetizationConfig.to_proto(
            resource.monetization_config
        ):
            res.monetization_config.CopyFrom(
                OrganizationAddonsConfigMonetizationConfig.to_proto(
                    resource.monetization_config
                )
            )
        else:
            res.ClearField("monetization_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OrganizationAddonsConfig(
            advanced_api_ops_config=OrganizationAddonsConfigAdvancedApiOpsConfig.from_proto(
                resource.advanced_api_ops_config
            ),
            monetization_config=OrganizationAddonsConfigMonetizationConfig.from_proto(
                resource.monetization_config
            ),
        )


class OrganizationAddonsConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OrganizationAddonsConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OrganizationAddonsConfig.from_proto(i) for i in resources]


class OrganizationAddonsConfigAdvancedApiOpsConfig(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = organization_pb2.ApigeeAlphaOrganizationAddonsConfigAdvancedApiOpsConfig()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OrganizationAddonsConfigAdvancedApiOpsConfig(
            enabled=Primitive.from_proto(resource.enabled),
        )


class OrganizationAddonsConfigAdvancedApiOpsConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OrganizationAddonsConfigAdvancedApiOpsConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OrganizationAddonsConfigAdvancedApiOpsConfig.from_proto(i)
            for i in resources
        ]


class OrganizationAddonsConfigMonetizationConfig(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = organization_pb2.ApigeeAlphaOrganizationAddonsConfigMonetizationConfig()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OrganizationAddonsConfigMonetizationConfig(
            enabled=Primitive.from_proto(resource.enabled),
        )


class OrganizationAddonsConfigMonetizationConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OrganizationAddonsConfigMonetizationConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OrganizationAddonsConfigMonetizationConfig.from_proto(i) for i in resources
        ]


class OrganizationRuntimeTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeAlphaOrganizationRuntimeTypeEnum.Value(
            "ApigeeAlphaOrganizationRuntimeTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeAlphaOrganizationRuntimeTypeEnum.Name(resource)[
            len("ApigeeAlphaOrganizationRuntimeTypeEnum") :
        ]


class OrganizationSubscriptionTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeAlphaOrganizationSubscriptionTypeEnum.Value(
            "ApigeeAlphaOrganizationSubscriptionTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeAlphaOrganizationSubscriptionTypeEnum.Name(
            resource
        )[len("ApigeeAlphaOrganizationSubscriptionTypeEnum") :]


class OrganizationBillingTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeAlphaOrganizationBillingTypeEnum.Value(
            "ApigeeAlphaOrganizationBillingTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeAlphaOrganizationBillingTypeEnum.Name(resource)[
            len("ApigeeAlphaOrganizationBillingTypeEnum") :
        ]


class OrganizationStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeAlphaOrganizationStateEnum.Value(
            "ApigeeAlphaOrganizationStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeAlphaOrganizationStateEnum.Name(resource)[
            len("ApigeeAlphaOrganizationStateEnum") :
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
