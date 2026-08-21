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
from google3.cloud.graphite.mmv2.services.google.identity_toolkit import tenant_pb2
from google3.cloud.graphite.mmv2.services.google.identity_toolkit import tenant_pb2_grpc

from typing import List


class Tenant(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        allow_password_signup: bool = None,
        enable_email_link_signin: bool = None,
        disable_auth: bool = None,
        enable_anonymous_user: bool = None,
        mfa_config: dict = None,
        test_phone_numbers: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.allow_password_signup = allow_password_signup
        self.enable_email_link_signin = enable_email_link_signin
        self.disable_auth = disable_auth
        self.enable_anonymous_user = enable_anonymous_user
        self.mfa_config = mfa_config
        self.test_phone_numbers = test_phone_numbers
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = tenant_pb2_grpc.IdentitytoolkitTenantServiceStub(channel.Channel())
        request = tenant_pb2.ApplyIdentitytoolkitTenantRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.allow_password_signup):
            request.resource.allow_password_signup = Primitive.to_proto(
                self.allow_password_signup
            )

        if Primitive.to_proto(self.enable_email_link_signin):
            request.resource.enable_email_link_signin = Primitive.to_proto(
                self.enable_email_link_signin
            )

        if Primitive.to_proto(self.disable_auth):
            request.resource.disable_auth = Primitive.to_proto(self.disable_auth)

        if Primitive.to_proto(self.enable_anonymous_user):
            request.resource.enable_anonymous_user = Primitive.to_proto(
                self.enable_anonymous_user
            )

        if TenantMfaConfig.to_proto(self.mfa_config):
            request.resource.mfa_config.CopyFrom(
                TenantMfaConfig.to_proto(self.mfa_config)
            )
        else:
            request.resource.ClearField("mfa_config")
        if Primitive.to_proto(self.test_phone_numbers):
            request.resource.test_phone_numbers = Primitive.to_proto(
                self.test_phone_numbers
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyIdentitytoolkitTenant(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.allow_password_signup = Primitive.from_proto(
            response.allow_password_signup
        )
        self.enable_email_link_signin = Primitive.from_proto(
            response.enable_email_link_signin
        )
        self.disable_auth = Primitive.from_proto(response.disable_auth)
        self.enable_anonymous_user = Primitive.from_proto(
            response.enable_anonymous_user
        )
        self.mfa_config = TenantMfaConfig.from_proto(response.mfa_config)
        self.test_phone_numbers = Primitive.from_proto(response.test_phone_numbers)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = tenant_pb2_grpc.IdentitytoolkitTenantServiceStub(channel.Channel())
        request = tenant_pb2.DeleteIdentitytoolkitTenantRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.allow_password_signup):
            request.resource.allow_password_signup = Primitive.to_proto(
                self.allow_password_signup
            )

        if Primitive.to_proto(self.enable_email_link_signin):
            request.resource.enable_email_link_signin = Primitive.to_proto(
                self.enable_email_link_signin
            )

        if Primitive.to_proto(self.disable_auth):
            request.resource.disable_auth = Primitive.to_proto(self.disable_auth)

        if Primitive.to_proto(self.enable_anonymous_user):
            request.resource.enable_anonymous_user = Primitive.to_proto(
                self.enable_anonymous_user
            )

        if TenantMfaConfig.to_proto(self.mfa_config):
            request.resource.mfa_config.CopyFrom(
                TenantMfaConfig.to_proto(self.mfa_config)
            )
        else:
            request.resource.ClearField("mfa_config")
        if Primitive.to_proto(self.test_phone_numbers):
            request.resource.test_phone_numbers = Primitive.to_proto(
                self.test_phone_numbers
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteIdentitytoolkitTenant(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = tenant_pb2_grpc.IdentitytoolkitTenantServiceStub(channel.Channel())
        request = tenant_pb2.ListIdentitytoolkitTenantRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListIdentitytoolkitTenant(request).items

    def to_proto(self):
        resource = tenant_pb2.IdentitytoolkitTenant()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.allow_password_signup):
            resource.allow_password_signup = Primitive.to_proto(
                self.allow_password_signup
            )
        if Primitive.to_proto(self.enable_email_link_signin):
            resource.enable_email_link_signin = Primitive.to_proto(
                self.enable_email_link_signin
            )
        if Primitive.to_proto(self.disable_auth):
            resource.disable_auth = Primitive.to_proto(self.disable_auth)
        if Primitive.to_proto(self.enable_anonymous_user):
            resource.enable_anonymous_user = Primitive.to_proto(
                self.enable_anonymous_user
            )
        if TenantMfaConfig.to_proto(self.mfa_config):
            resource.mfa_config.CopyFrom(TenantMfaConfig.to_proto(self.mfa_config))
        else:
            resource.ClearField("mfa_config")
        if Primitive.to_proto(self.test_phone_numbers):
            resource.test_phone_numbers = Primitive.to_proto(self.test_phone_numbers)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class TenantMfaConfig(object):
    def __init__(self, state: str = None, enabled_providers: list = None):
        self.state = state
        self.enabled_providers = enabled_providers

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = tenant_pb2.IdentitytoolkitTenantMfaConfig()
        if TenantMfaConfigStateEnum.to_proto(resource.state):
            res.state = TenantMfaConfigStateEnum.to_proto(resource.state)
        if TenantMfaConfigEnabledProvidersEnumArray.to_proto(
            resource.enabled_providers
        ):
            res.enabled_providers.extend(
                TenantMfaConfigEnabledProvidersEnumArray.to_proto(
                    resource.enabled_providers
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TenantMfaConfig(
            state=TenantMfaConfigStateEnum.from_proto(resource.state),
            enabled_providers=TenantMfaConfigEnabledProvidersEnumArray.from_proto(
                resource.enabled_providers
            ),
        )


class TenantMfaConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TenantMfaConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TenantMfaConfig.from_proto(i) for i in resources]


class TenantMfaConfigStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return tenant_pb2.IdentitytoolkitTenantMfaConfigStateEnum.Value(
            "IdentitytoolkitTenantMfaConfigStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return tenant_pb2.IdentitytoolkitTenantMfaConfigStateEnum.Name(resource)[
            len("IdentitytoolkitTenantMfaConfigStateEnum") :
        ]


class TenantMfaConfigEnabledProvidersEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return tenant_pb2.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum.Value(
            "IdentitytoolkitTenantMfaConfigEnabledProvidersEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return tenant_pb2.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum.Name(
            resource
        )[len("IdentitytoolkitTenantMfaConfigEnabledProvidersEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
