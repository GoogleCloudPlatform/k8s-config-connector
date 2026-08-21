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
from google3.cloud.graphite.mmv2.services.google.identity_toolkit import (
    tenant_oauth_idp_config_pb2,
)
from google3.cloud.graphite.mmv2.services.google.identity_toolkit import (
    tenant_oauth_idp_config_pb2_grpc,
)

from typing import List


class TenantOAuthIdpConfig(object):
    def __init__(
        self,
        name: str = None,
        client_id: str = None,
        issuer: str = None,
        display_name: str = None,
        enabled: bool = None,
        client_secret: str = None,
        response_type: dict = None,
        project: str = None,
        tenant: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.client_id = client_id
        self.issuer = issuer
        self.display_name = display_name
        self.enabled = enabled
        self.client_secret = client_secret
        self.response_type = response_type
        self.project = project
        self.tenant = tenant
        self.service_account_file = service_account_file

    def apply(self):
        stub = tenant_oauth_idp_config_pb2_grpc.IdentitytoolkitAlphaTenantOAuthIdpConfigServiceStub(
            channel.Channel()
        )
        request = (
            tenant_oauth_idp_config_pb2.ApplyIdentitytoolkitAlphaTenantOAuthIdpConfigRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.client_id):
            request.resource.client_id = Primitive.to_proto(self.client_id)

        if Primitive.to_proto(self.issuer):
            request.resource.issuer = Primitive.to_proto(self.issuer)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.enabled):
            request.resource.enabled = Primitive.to_proto(self.enabled)

        if Primitive.to_proto(self.client_secret):
            request.resource.client_secret = Primitive.to_proto(self.client_secret)

        if TenantOAuthIdpConfigResponseType.to_proto(self.response_type):
            request.resource.response_type.CopyFrom(
                TenantOAuthIdpConfigResponseType.to_proto(self.response_type)
            )
        else:
            request.resource.ClearField("response_type")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.tenant):
            request.resource.tenant = Primitive.to_proto(self.tenant)

        request.service_account_file = self.service_account_file

        response = stub.ApplyIdentitytoolkitAlphaTenantOAuthIdpConfig(request)
        self.name = Primitive.from_proto(response.name)
        self.client_id = Primitive.from_proto(response.client_id)
        self.issuer = Primitive.from_proto(response.issuer)
        self.display_name = Primitive.from_proto(response.display_name)
        self.enabled = Primitive.from_proto(response.enabled)
        self.client_secret = Primitive.from_proto(response.client_secret)
        self.response_type = TenantOAuthIdpConfigResponseType.from_proto(
            response.response_type
        )
        self.project = Primitive.from_proto(response.project)
        self.tenant = Primitive.from_proto(response.tenant)

    def delete(self):
        stub = tenant_oauth_idp_config_pb2_grpc.IdentitytoolkitAlphaTenantOAuthIdpConfigServiceStub(
            channel.Channel()
        )
        request = (
            tenant_oauth_idp_config_pb2.DeleteIdentitytoolkitAlphaTenantOAuthIdpConfigRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.client_id):
            request.resource.client_id = Primitive.to_proto(self.client_id)

        if Primitive.to_proto(self.issuer):
            request.resource.issuer = Primitive.to_proto(self.issuer)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.enabled):
            request.resource.enabled = Primitive.to_proto(self.enabled)

        if Primitive.to_proto(self.client_secret):
            request.resource.client_secret = Primitive.to_proto(self.client_secret)

        if TenantOAuthIdpConfigResponseType.to_proto(self.response_type):
            request.resource.response_type.CopyFrom(
                TenantOAuthIdpConfigResponseType.to_proto(self.response_type)
            )
        else:
            request.resource.ClearField("response_type")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.tenant):
            request.resource.tenant = Primitive.to_proto(self.tenant)

        response = stub.DeleteIdentitytoolkitAlphaTenantOAuthIdpConfig(request)

    @classmethod
    def list(self, project, tenant, service_account_file=""):
        stub = tenant_oauth_idp_config_pb2_grpc.IdentitytoolkitAlphaTenantOAuthIdpConfigServiceStub(
            channel.Channel()
        )
        request = (
            tenant_oauth_idp_config_pb2.ListIdentitytoolkitAlphaTenantOAuthIdpConfigRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Tenant = tenant

        return stub.ListIdentitytoolkitAlphaTenantOAuthIdpConfig(request).items

    def to_proto(self):
        resource = (
            tenant_oauth_idp_config_pb2.IdentitytoolkitAlphaTenantOAuthIdpConfig()
        )
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.client_id):
            resource.client_id = Primitive.to_proto(self.client_id)
        if Primitive.to_proto(self.issuer):
            resource.issuer = Primitive.to_proto(self.issuer)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.enabled):
            resource.enabled = Primitive.to_proto(self.enabled)
        if Primitive.to_proto(self.client_secret):
            resource.client_secret = Primitive.to_proto(self.client_secret)
        if TenantOAuthIdpConfigResponseType.to_proto(self.response_type):
            resource.response_type.CopyFrom(
                TenantOAuthIdpConfigResponseType.to_proto(self.response_type)
            )
        else:
            resource.ClearField("response_type")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.tenant):
            resource.tenant = Primitive.to_proto(self.tenant)
        return resource


class TenantOAuthIdpConfigResponseType(object):
    def __init__(self, id_token: bool = None, code: bool = None, token: bool = None):
        self.id_token = id_token
        self.code = code
        self.token = token

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            tenant_oauth_idp_config_pb2.IdentitytoolkitAlphaTenantOAuthIdpConfigResponseType()
        )
        if Primitive.to_proto(resource.id_token):
            res.id_token = Primitive.to_proto(resource.id_token)
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.token):
            res.token = Primitive.to_proto(resource.token)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TenantOAuthIdpConfigResponseType(
            id_token=Primitive.from_proto(resource.id_token),
            code=Primitive.from_proto(resource.code),
            token=Primitive.from_proto(resource.token),
        )


class TenantOAuthIdpConfigResponseTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TenantOAuthIdpConfigResponseType.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TenantOAuthIdpConfigResponseType.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
