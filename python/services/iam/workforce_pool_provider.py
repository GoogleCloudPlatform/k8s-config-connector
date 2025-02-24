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
from google3.cloud.graphite.mmv2.services.google.iam import workforce_pool_provider_pb2
from google3.cloud.graphite.mmv2.services.google.iam import (
    workforce_pool_provider_pb2_grpc,
)

from typing import List


class WorkforcePoolProvider(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        description: str = None,
        state: str = None,
        disabled: bool = None,
        attribute_mapping: dict = None,
        attribute_condition: str = None,
        saml: dict = None,
        oidc: dict = None,
        location: str = None,
        workforce_pool: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.description = description
        self.disabled = disabled
        self.attribute_mapping = attribute_mapping
        self.attribute_condition = attribute_condition
        self.saml = saml
        self.oidc = oidc
        self.location = location
        self.workforce_pool = workforce_pool
        self.service_account_file = service_account_file

    def apply(self):
        stub = workforce_pool_provider_pb2_grpc.IamWorkforcePoolProviderServiceStub(
            channel.Channel()
        )
        request = workforce_pool_provider_pb2.ApplyIamWorkforcePoolProviderRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.attribute_mapping):
            request.resource.attribute_mapping = Primitive.to_proto(
                self.attribute_mapping
            )

        if Primitive.to_proto(self.attribute_condition):
            request.resource.attribute_condition = Primitive.to_proto(
                self.attribute_condition
            )

        if WorkforcePoolProviderSaml.to_proto(self.saml):
            request.resource.saml.CopyFrom(
                WorkforcePoolProviderSaml.to_proto(self.saml)
            )
        else:
            request.resource.ClearField("saml")
        if WorkforcePoolProviderOidc.to_proto(self.oidc):
            request.resource.oidc.CopyFrom(
                WorkforcePoolProviderOidc.to_proto(self.oidc)
            )
        else:
            request.resource.ClearField("oidc")
        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.workforce_pool):
            request.resource.workforce_pool = Primitive.to_proto(self.workforce_pool)

        request.service_account_file = self.service_account_file

        response = stub.ApplyIamWorkforcePoolProvider(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.state = WorkforcePoolProviderStateEnum.from_proto(response.state)
        self.disabled = Primitive.from_proto(response.disabled)
        self.attribute_mapping = Primitive.from_proto(response.attribute_mapping)
        self.attribute_condition = Primitive.from_proto(response.attribute_condition)
        self.saml = WorkforcePoolProviderSaml.from_proto(response.saml)
        self.oidc = WorkforcePoolProviderOidc.from_proto(response.oidc)
        self.location = Primitive.from_proto(response.location)
        self.workforce_pool = Primitive.from_proto(response.workforce_pool)

    def delete(self):
        stub = workforce_pool_provider_pb2_grpc.IamWorkforcePoolProviderServiceStub(
            channel.Channel()
        )
        request = workforce_pool_provider_pb2.DeleteIamWorkforcePoolProviderRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.attribute_mapping):
            request.resource.attribute_mapping = Primitive.to_proto(
                self.attribute_mapping
            )

        if Primitive.to_proto(self.attribute_condition):
            request.resource.attribute_condition = Primitive.to_proto(
                self.attribute_condition
            )

        if WorkforcePoolProviderSaml.to_proto(self.saml):
            request.resource.saml.CopyFrom(
                WorkforcePoolProviderSaml.to_proto(self.saml)
            )
        else:
            request.resource.ClearField("saml")
        if WorkforcePoolProviderOidc.to_proto(self.oidc):
            request.resource.oidc.CopyFrom(
                WorkforcePoolProviderOidc.to_proto(self.oidc)
            )
        else:
            request.resource.ClearField("oidc")
        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.workforce_pool):
            request.resource.workforce_pool = Primitive.to_proto(self.workforce_pool)

        response = stub.DeleteIamWorkforcePoolProvider(request)

    @classmethod
    def list(self, location, workforcePool, service_account_file=""):
        stub = workforce_pool_provider_pb2_grpc.IamWorkforcePoolProviderServiceStub(
            channel.Channel()
        )
        request = workforce_pool_provider_pb2.ListIamWorkforcePoolProviderRequest()
        request.service_account_file = service_account_file
        request.Location = location

        request.WorkforcePool = workforcePool

        return stub.ListIamWorkforcePoolProvider(request).items

    def to_proto(self):
        resource = workforce_pool_provider_pb2.IamWorkforcePoolProvider()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if Primitive.to_proto(self.attribute_mapping):
            resource.attribute_mapping = Primitive.to_proto(self.attribute_mapping)
        if Primitive.to_proto(self.attribute_condition):
            resource.attribute_condition = Primitive.to_proto(self.attribute_condition)
        if WorkforcePoolProviderSaml.to_proto(self.saml):
            resource.saml.CopyFrom(WorkforcePoolProviderSaml.to_proto(self.saml))
        else:
            resource.ClearField("saml")
        if WorkforcePoolProviderOidc.to_proto(self.oidc):
            resource.oidc.CopyFrom(WorkforcePoolProviderOidc.to_proto(self.oidc))
        else:
            resource.ClearField("oidc")
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.workforce_pool):
            resource.workforce_pool = Primitive.to_proto(self.workforce_pool)
        return resource


class WorkforcePoolProviderSaml(object):
    def __init__(self, idp_metadata_xml: str = None):
        self.idp_metadata_xml = idp_metadata_xml

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workforce_pool_provider_pb2.IamWorkforcePoolProviderSaml()
        if Primitive.to_proto(resource.idp_metadata_xml):
            res.idp_metadata_xml = Primitive.to_proto(resource.idp_metadata_xml)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkforcePoolProviderSaml(
            idp_metadata_xml=Primitive.from_proto(resource.idp_metadata_xml),
        )


class WorkforcePoolProviderSamlArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkforcePoolProviderSaml.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkforcePoolProviderSaml.from_proto(i) for i in resources]


class WorkforcePoolProviderOidc(object):
    def __init__(
        self,
        issuer_uri: str = None,
        client_id: str = None,
        jwks_json: str = None,
        web_sso_config: dict = None,
        client_secret: dict = None,
    ):
        self.issuer_uri = issuer_uri
        self.client_id = client_id
        self.jwks_json = jwks_json
        self.web_sso_config = web_sso_config
        self.client_secret = client_secret

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workforce_pool_provider_pb2.IamWorkforcePoolProviderOidc()
        if Primitive.to_proto(resource.issuer_uri):
            res.issuer_uri = Primitive.to_proto(resource.issuer_uri)
        if Primitive.to_proto(resource.client_id):
            res.client_id = Primitive.to_proto(resource.client_id)
        if Primitive.to_proto(resource.jwks_json):
            res.jwks_json = Primitive.to_proto(resource.jwks_json)
        if WorkforcePoolProviderOidcWebSsoConfig.to_proto(resource.web_sso_config):
            res.web_sso_config.CopyFrom(
                WorkforcePoolProviderOidcWebSsoConfig.to_proto(resource.web_sso_config)
            )
        else:
            res.ClearField("web_sso_config")
        if WorkforcePoolProviderOidcClientSecret.to_proto(resource.client_secret):
            res.client_secret.CopyFrom(
                WorkforcePoolProviderOidcClientSecret.to_proto(resource.client_secret)
            )
        else:
            res.ClearField("client_secret")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkforcePoolProviderOidc(
            issuer_uri=Primitive.from_proto(resource.issuer_uri),
            client_id=Primitive.from_proto(resource.client_id),
            jwks_json=Primitive.from_proto(resource.jwks_json),
            web_sso_config=WorkforcePoolProviderOidcWebSsoConfig.from_proto(
                resource.web_sso_config
            ),
            client_secret=WorkforcePoolProviderOidcClientSecret.from_proto(
                resource.client_secret
            ),
        )


class WorkforcePoolProviderOidcArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkforcePoolProviderOidc.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkforcePoolProviderOidc.from_proto(i) for i in resources]


class WorkforcePoolProviderOidcWebSsoConfig(object):
    def __init__(
        self,
        response_type: str = None,
        assertion_claims_behavior: str = None,
        additional_scopes: list = None,
    ):
        self.response_type = response_type
        self.assertion_claims_behavior = assertion_claims_behavior
        self.additional_scopes = additional_scopes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workforce_pool_provider_pb2.IamWorkforcePoolProviderOidcWebSsoConfig()
        if WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum.to_proto(
            resource.response_type
        ):
            res.response_type = (
                WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum.to_proto(
                    resource.response_type
                )
            )
        if WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum.to_proto(
            resource.assertion_claims_behavior
        ):
            res.assertion_claims_behavior = WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum.to_proto(
                resource.assertion_claims_behavior
            )
        if Primitive.to_proto(resource.additional_scopes):
            res.additional_scopes.extend(Primitive.to_proto(resource.additional_scopes))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkforcePoolProviderOidcWebSsoConfig(
            response_type=WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum.from_proto(
                resource.response_type
            ),
            assertion_claims_behavior=WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum.from_proto(
                resource.assertion_claims_behavior
            ),
            additional_scopes=Primitive.from_proto(resource.additional_scopes),
        )


class WorkforcePoolProviderOidcWebSsoConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkforcePoolProviderOidcWebSsoConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkforcePoolProviderOidcWebSsoConfig.from_proto(i) for i in resources]


class WorkforcePoolProviderOidcClientSecret(object):
    def __init__(self, value: dict = None):
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workforce_pool_provider_pb2.IamWorkforcePoolProviderOidcClientSecret()
        if WorkforcePoolProviderOidcClientSecretValue.to_proto(resource.value):
            res.value.CopyFrom(
                WorkforcePoolProviderOidcClientSecretValue.to_proto(resource.value)
            )
        else:
            res.ClearField("value")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkforcePoolProviderOidcClientSecret(
            value=WorkforcePoolProviderOidcClientSecretValue.from_proto(resource.value),
        )


class WorkforcePoolProviderOidcClientSecretArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkforcePoolProviderOidcClientSecret.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkforcePoolProviderOidcClientSecret.from_proto(i) for i in resources]


class WorkforcePoolProviderOidcClientSecretValue(object):
    def __init__(self, plain_text: str = None, thumbprint: str = None):
        self.plain_text = plain_text
        self.thumbprint = thumbprint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workforce_pool_provider_pb2.IamWorkforcePoolProviderOidcClientSecretValue()
        )
        if Primitive.to_proto(resource.plain_text):
            res.plain_text = Primitive.to_proto(resource.plain_text)
        if Primitive.to_proto(resource.thumbprint):
            res.thumbprint = Primitive.to_proto(resource.thumbprint)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkforcePoolProviderOidcClientSecretValue(
            plain_text=Primitive.from_proto(resource.plain_text),
            thumbprint=Primitive.from_proto(resource.thumbprint),
        )


class WorkforcePoolProviderOidcClientSecretValueArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkforcePoolProviderOidcClientSecretValue.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkforcePoolProviderOidcClientSecretValue.from_proto(i) for i in resources
        ]


class WorkforcePoolProviderStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workforce_pool_provider_pb2.IamWorkforcePoolProviderStateEnum.Value(
            "IamWorkforcePoolProviderStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workforce_pool_provider_pb2.IamWorkforcePoolProviderStateEnum.Name(
            resource
        )[len("IamWorkforcePoolProviderStateEnum") :]


class WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workforce_pool_provider_pb2.IamWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum.Value(
            "IamWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workforce_pool_provider_pb2.IamWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum.Name(
            resource
        )[
            len("IamWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum") :
        ]


class WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workforce_pool_provider_pb2.IamWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum.Value(
            "IamWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workforce_pool_provider_pb2.IamWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum.Name(
            resource
        )[
            len("IamWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum") :
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
