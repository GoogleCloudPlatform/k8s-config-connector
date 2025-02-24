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
from google3.cloud.graphite.mmv2.services.google.app_engine import application_pb2
from google3.cloud.graphite.mmv2.services.google.app_engine import application_pb2_grpc

from typing import List


class Application(object):
    def __init__(
        self,
        name: str = None,
        dispatch_rules: list = None,
        auth_domain: str = None,
        location: str = None,
        code_bucket: str = None,
        default_cookie_expiration: str = None,
        serving_status: str = None,
        default_hostname: str = None,
        default_bucket: str = None,
        iap: dict = None,
        gcr_domain: str = None,
        database_type: str = None,
        feature_settings: dict = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.dispatch_rules = dispatch_rules
        self.auth_domain = auth_domain
        self.location = location
        self.default_cookie_expiration = default_cookie_expiration
        self.serving_status = serving_status
        self.iap = iap
        self.gcr_domain = gcr_domain
        self.database_type = database_type
        self.feature_settings = feature_settings
        self.service_account_file = service_account_file

    def apply(self):
        stub = application_pb2_grpc.AppengineApplicationServiceStub(channel.Channel())
        request = application_pb2.ApplyAppengineApplicationRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ApplicationDispatchRulesArray.to_proto(self.dispatch_rules):
            request.resource.dispatch_rules.extend(
                ApplicationDispatchRulesArray.to_proto(self.dispatch_rules)
            )
        if Primitive.to_proto(self.auth_domain):
            request.resource.auth_domain = Primitive.to_proto(self.auth_domain)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.default_cookie_expiration):
            request.resource.default_cookie_expiration = Primitive.to_proto(
                self.default_cookie_expiration
            )

        if ApplicationServingStatusEnum.to_proto(self.serving_status):
            request.resource.serving_status = ApplicationServingStatusEnum.to_proto(
                self.serving_status
            )

        if ApplicationIap.to_proto(self.iap):
            request.resource.iap.CopyFrom(ApplicationIap.to_proto(self.iap))
        else:
            request.resource.ClearField("iap")
        if Primitive.to_proto(self.gcr_domain):
            request.resource.gcr_domain = Primitive.to_proto(self.gcr_domain)

        if ApplicationDatabaseTypeEnum.to_proto(self.database_type):
            request.resource.database_type = ApplicationDatabaseTypeEnum.to_proto(
                self.database_type
            )

        if ApplicationFeatureSettings.to_proto(self.feature_settings):
            request.resource.feature_settings.CopyFrom(
                ApplicationFeatureSettings.to_proto(self.feature_settings)
            )
        else:
            request.resource.ClearField("feature_settings")
        request.service_account_file = self.service_account_file

        response = stub.ApplyAppengineApplication(request)
        self.name = Primitive.from_proto(response.name)
        self.dispatch_rules = ApplicationDispatchRulesArray.from_proto(
            response.dispatch_rules
        )
        self.auth_domain = Primitive.from_proto(response.auth_domain)
        self.location = Primitive.from_proto(response.location)
        self.code_bucket = Primitive.from_proto(response.code_bucket)
        self.default_cookie_expiration = Primitive.from_proto(
            response.default_cookie_expiration
        )
        self.serving_status = ApplicationServingStatusEnum.from_proto(
            response.serving_status
        )
        self.default_hostname = Primitive.from_proto(response.default_hostname)
        self.default_bucket = Primitive.from_proto(response.default_bucket)
        self.iap = ApplicationIap.from_proto(response.iap)
        self.gcr_domain = Primitive.from_proto(response.gcr_domain)
        self.database_type = ApplicationDatabaseTypeEnum.from_proto(
            response.database_type
        )
        self.feature_settings = ApplicationFeatureSettings.from_proto(
            response.feature_settings
        )

    def delete(self):
        stub = application_pb2_grpc.AppengineApplicationServiceStub(channel.Channel())
        request = application_pb2.DeleteAppengineApplicationRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ApplicationDispatchRulesArray.to_proto(self.dispatch_rules):
            request.resource.dispatch_rules.extend(
                ApplicationDispatchRulesArray.to_proto(self.dispatch_rules)
            )
        if Primitive.to_proto(self.auth_domain):
            request.resource.auth_domain = Primitive.to_proto(self.auth_domain)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.default_cookie_expiration):
            request.resource.default_cookie_expiration = Primitive.to_proto(
                self.default_cookie_expiration
            )

        if ApplicationServingStatusEnum.to_proto(self.serving_status):
            request.resource.serving_status = ApplicationServingStatusEnum.to_proto(
                self.serving_status
            )

        if ApplicationIap.to_proto(self.iap):
            request.resource.iap.CopyFrom(ApplicationIap.to_proto(self.iap))
        else:
            request.resource.ClearField("iap")
        if Primitive.to_proto(self.gcr_domain):
            request.resource.gcr_domain = Primitive.to_proto(self.gcr_domain)

        if ApplicationDatabaseTypeEnum.to_proto(self.database_type):
            request.resource.database_type = ApplicationDatabaseTypeEnum.to_proto(
                self.database_type
            )

        if ApplicationFeatureSettings.to_proto(self.feature_settings):
            request.resource.feature_settings.CopyFrom(
                ApplicationFeatureSettings.to_proto(self.feature_settings)
            )
        else:
            request.resource.ClearField("feature_settings")
        response = stub.DeleteAppengineApplication(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = application_pb2_grpc.AppengineApplicationServiceStub(channel.Channel())
        request = application_pb2.ListAppengineApplicationRequest()
        request.service_account_file = service_account_file

        return stub.ListAppengineApplication(request).items

    def to_proto(self):
        resource = application_pb2.AppengineApplication()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if ApplicationDispatchRulesArray.to_proto(self.dispatch_rules):
            resource.dispatch_rules.extend(
                ApplicationDispatchRulesArray.to_proto(self.dispatch_rules)
            )
        if Primitive.to_proto(self.auth_domain):
            resource.auth_domain = Primitive.to_proto(self.auth_domain)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.default_cookie_expiration):
            resource.default_cookie_expiration = Primitive.to_proto(
                self.default_cookie_expiration
            )
        if ApplicationServingStatusEnum.to_proto(self.serving_status):
            resource.serving_status = ApplicationServingStatusEnum.to_proto(
                self.serving_status
            )
        if ApplicationIap.to_proto(self.iap):
            resource.iap.CopyFrom(ApplicationIap.to_proto(self.iap))
        else:
            resource.ClearField("iap")
        if Primitive.to_proto(self.gcr_domain):
            resource.gcr_domain = Primitive.to_proto(self.gcr_domain)
        if ApplicationDatabaseTypeEnum.to_proto(self.database_type):
            resource.database_type = ApplicationDatabaseTypeEnum.to_proto(
                self.database_type
            )
        if ApplicationFeatureSettings.to_proto(self.feature_settings):
            resource.feature_settings.CopyFrom(
                ApplicationFeatureSettings.to_proto(self.feature_settings)
            )
        else:
            resource.ClearField("feature_settings")
        return resource


class ApplicationDispatchRules(object):
    def __init__(self, domain: str = None, path: str = None, service: str = None):
        self.domain = domain
        self.path = path
        self.service = service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = application_pb2.AppengineApplicationDispatchRules()
        if Primitive.to_proto(resource.domain):
            res.domain = Primitive.to_proto(resource.domain)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ApplicationDispatchRules(
            domain=Primitive.from_proto(resource.domain),
            path=Primitive.from_proto(resource.path),
            service=Primitive.from_proto(resource.service),
        )


class ApplicationDispatchRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ApplicationDispatchRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ApplicationDispatchRules.from_proto(i) for i in resources]


class ApplicationIap(object):
    def __init__(
        self,
        enabled: bool = None,
        oauth2_client_id: str = None,
        oauth2_client_secret: str = None,
        oauth2_client_secret_sha256: str = None,
    ):
        self.enabled = enabled
        self.oauth2_client_id = oauth2_client_id
        self.oauth2_client_secret = oauth2_client_secret
        self.oauth2_client_secret_sha256 = oauth2_client_secret_sha256

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = application_pb2.AppengineApplicationIap()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.oauth2_client_id):
            res.oauth2_client_id = Primitive.to_proto(resource.oauth2_client_id)
        if Primitive.to_proto(resource.oauth2_client_secret):
            res.oauth2_client_secret = Primitive.to_proto(resource.oauth2_client_secret)
        if Primitive.to_proto(resource.oauth2_client_secret_sha256):
            res.oauth2_client_secret_sha256 = Primitive.to_proto(
                resource.oauth2_client_secret_sha256
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ApplicationIap(
            enabled=Primitive.from_proto(resource.enabled),
            oauth2_client_id=Primitive.from_proto(resource.oauth2_client_id),
            oauth2_client_secret=Primitive.from_proto(resource.oauth2_client_secret),
            oauth2_client_secret_sha256=Primitive.from_proto(
                resource.oauth2_client_secret_sha256
            ),
        )


class ApplicationIapArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ApplicationIap.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ApplicationIap.from_proto(i) for i in resources]


class ApplicationFeatureSettings(object):
    def __init__(
        self, split_health_checks: bool = None, use_container_optimized_os: bool = None
    ):
        self.split_health_checks = split_health_checks
        self.use_container_optimized_os = use_container_optimized_os

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = application_pb2.AppengineApplicationFeatureSettings()
        if Primitive.to_proto(resource.split_health_checks):
            res.split_health_checks = Primitive.to_proto(resource.split_health_checks)
        if Primitive.to_proto(resource.use_container_optimized_os):
            res.use_container_optimized_os = Primitive.to_proto(
                resource.use_container_optimized_os
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ApplicationFeatureSettings(
            split_health_checks=Primitive.from_proto(resource.split_health_checks),
            use_container_optimized_os=Primitive.from_proto(
                resource.use_container_optimized_os
            ),
        )


class ApplicationFeatureSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ApplicationFeatureSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ApplicationFeatureSettings.from_proto(i) for i in resources]


class ApplicationServingStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return application_pb2.AppengineApplicationServingStatusEnum.Value(
            "AppengineApplicationServingStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return application_pb2.AppengineApplicationServingStatusEnum.Name(resource)[
            len("AppengineApplicationServingStatusEnum") :
        ]


class ApplicationDatabaseTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return application_pb2.AppengineApplicationDatabaseTypeEnum.Value(
            "AppengineApplicationDatabaseTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return application_pb2.AppengineApplicationDatabaseTypeEnum.Name(resource)[
            len("AppengineApplicationDatabaseTypeEnum") :
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
