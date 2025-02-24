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
from google3.cloud.graphite.mmv2.services.google.recaptcha_enterprise import key_pb2
from google3.cloud.graphite.mmv2.services.google.recaptcha_enterprise import (
    key_pb2_grpc,
)

from typing import List


class Key(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        web_settings: dict = None,
        android_settings: dict = None,
        ios_settings: dict = None,
        labels: dict = None,
        create_time: str = None,
        testing_options: dict = None,
        waf_settings: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.web_settings = web_settings
        self.android_settings = android_settings
        self.ios_settings = ios_settings
        self.labels = labels
        self.testing_options = testing_options
        self.waf_settings = waf_settings
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = key_pb2_grpc.RecaptchaenterpriseBetaKeyServiceStub(channel.Channel())
        request = key_pb2.ApplyRecaptchaenterpriseBetaKeyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if KeyWebSettings.to_proto(self.web_settings):
            request.resource.web_settings.CopyFrom(
                KeyWebSettings.to_proto(self.web_settings)
            )
        else:
            request.resource.ClearField("web_settings")
        if KeyAndroidSettings.to_proto(self.android_settings):
            request.resource.android_settings.CopyFrom(
                KeyAndroidSettings.to_proto(self.android_settings)
            )
        else:
            request.resource.ClearField("android_settings")
        if KeyIosSettings.to_proto(self.ios_settings):
            request.resource.ios_settings.CopyFrom(
                KeyIosSettings.to_proto(self.ios_settings)
            )
        else:
            request.resource.ClearField("ios_settings")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if KeyTestingOptions.to_proto(self.testing_options):
            request.resource.testing_options.CopyFrom(
                KeyTestingOptions.to_proto(self.testing_options)
            )
        else:
            request.resource.ClearField("testing_options")
        if KeyWafSettings.to_proto(self.waf_settings):
            request.resource.waf_settings.CopyFrom(
                KeyWafSettings.to_proto(self.waf_settings)
            )
        else:
            request.resource.ClearField("waf_settings")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyRecaptchaenterpriseBetaKey(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.web_settings = KeyWebSettings.from_proto(response.web_settings)
        self.android_settings = KeyAndroidSettings.from_proto(response.android_settings)
        self.ios_settings = KeyIosSettings.from_proto(response.ios_settings)
        self.labels = Primitive.from_proto(response.labels)
        self.create_time = Primitive.from_proto(response.create_time)
        self.testing_options = KeyTestingOptions.from_proto(response.testing_options)
        self.waf_settings = KeyWafSettings.from_proto(response.waf_settings)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = key_pb2_grpc.RecaptchaenterpriseBetaKeyServiceStub(channel.Channel())
        request = key_pb2.DeleteRecaptchaenterpriseBetaKeyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if KeyWebSettings.to_proto(self.web_settings):
            request.resource.web_settings.CopyFrom(
                KeyWebSettings.to_proto(self.web_settings)
            )
        else:
            request.resource.ClearField("web_settings")
        if KeyAndroidSettings.to_proto(self.android_settings):
            request.resource.android_settings.CopyFrom(
                KeyAndroidSettings.to_proto(self.android_settings)
            )
        else:
            request.resource.ClearField("android_settings")
        if KeyIosSettings.to_proto(self.ios_settings):
            request.resource.ios_settings.CopyFrom(
                KeyIosSettings.to_proto(self.ios_settings)
            )
        else:
            request.resource.ClearField("ios_settings")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if KeyTestingOptions.to_proto(self.testing_options):
            request.resource.testing_options.CopyFrom(
                KeyTestingOptions.to_proto(self.testing_options)
            )
        else:
            request.resource.ClearField("testing_options")
        if KeyWafSettings.to_proto(self.waf_settings):
            request.resource.waf_settings.CopyFrom(
                KeyWafSettings.to_proto(self.waf_settings)
            )
        else:
            request.resource.ClearField("waf_settings")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteRecaptchaenterpriseBetaKey(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = key_pb2_grpc.RecaptchaenterpriseBetaKeyServiceStub(channel.Channel())
        request = key_pb2.ListRecaptchaenterpriseBetaKeyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListRecaptchaenterpriseBetaKey(request).items

    def to_proto(self):
        resource = key_pb2.RecaptchaenterpriseBetaKey()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if KeyWebSettings.to_proto(self.web_settings):
            resource.web_settings.CopyFrom(KeyWebSettings.to_proto(self.web_settings))
        else:
            resource.ClearField("web_settings")
        if KeyAndroidSettings.to_proto(self.android_settings):
            resource.android_settings.CopyFrom(
                KeyAndroidSettings.to_proto(self.android_settings)
            )
        else:
            resource.ClearField("android_settings")
        if KeyIosSettings.to_proto(self.ios_settings):
            resource.ios_settings.CopyFrom(KeyIosSettings.to_proto(self.ios_settings))
        else:
            resource.ClearField("ios_settings")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if KeyTestingOptions.to_proto(self.testing_options):
            resource.testing_options.CopyFrom(
                KeyTestingOptions.to_proto(self.testing_options)
            )
        else:
            resource.ClearField("testing_options")
        if KeyWafSettings.to_proto(self.waf_settings):
            resource.waf_settings.CopyFrom(KeyWafSettings.to_proto(self.waf_settings))
        else:
            resource.ClearField("waf_settings")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class KeyWebSettings(object):
    def __init__(
        self,
        allow_all_domains: bool = None,
        allowed_domains: list = None,
        allow_amp_traffic: bool = None,
        integration_type: str = None,
        challenge_security_preference: str = None,
    ):
        self.allow_all_domains = allow_all_domains
        self.allowed_domains = allowed_domains
        self.allow_amp_traffic = allow_amp_traffic
        self.integration_type = integration_type
        self.challenge_security_preference = challenge_security_preference

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = key_pb2.RecaptchaenterpriseBetaKeyWebSettings()
        if Primitive.to_proto(resource.allow_all_domains):
            res.allow_all_domains = Primitive.to_proto(resource.allow_all_domains)
        if Primitive.to_proto(resource.allowed_domains):
            res.allowed_domains.extend(Primitive.to_proto(resource.allowed_domains))
        if Primitive.to_proto(resource.allow_amp_traffic):
            res.allow_amp_traffic = Primitive.to_proto(resource.allow_amp_traffic)
        if KeyWebSettingsIntegrationTypeEnum.to_proto(resource.integration_type):
            res.integration_type = KeyWebSettingsIntegrationTypeEnum.to_proto(
                resource.integration_type
            )
        if KeyWebSettingsChallengeSecurityPreferenceEnum.to_proto(
            resource.challenge_security_preference
        ):
            res.challenge_security_preference = (
                KeyWebSettingsChallengeSecurityPreferenceEnum.to_proto(
                    resource.challenge_security_preference
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyWebSettings(
            allow_all_domains=Primitive.from_proto(resource.allow_all_domains),
            allowed_domains=Primitive.from_proto(resource.allowed_domains),
            allow_amp_traffic=Primitive.from_proto(resource.allow_amp_traffic),
            integration_type=KeyWebSettingsIntegrationTypeEnum.from_proto(
                resource.integration_type
            ),
            challenge_security_preference=KeyWebSettingsChallengeSecurityPreferenceEnum.from_proto(
                resource.challenge_security_preference
            ),
        )


class KeyWebSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KeyWebSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KeyWebSettings.from_proto(i) for i in resources]


class KeyAndroidSettings(object):
    def __init__(
        self, allow_all_package_names: bool = None, allowed_package_names: list = None
    ):
        self.allow_all_package_names = allow_all_package_names
        self.allowed_package_names = allowed_package_names

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = key_pb2.RecaptchaenterpriseBetaKeyAndroidSettings()
        if Primitive.to_proto(resource.allow_all_package_names):
            res.allow_all_package_names = Primitive.to_proto(
                resource.allow_all_package_names
            )
        if Primitive.to_proto(resource.allowed_package_names):
            res.allowed_package_names.extend(
                Primitive.to_proto(resource.allowed_package_names)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyAndroidSettings(
            allow_all_package_names=Primitive.from_proto(
                resource.allow_all_package_names
            ),
            allowed_package_names=Primitive.from_proto(resource.allowed_package_names),
        )


class KeyAndroidSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KeyAndroidSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KeyAndroidSettings.from_proto(i) for i in resources]


class KeyIosSettings(object):
    def __init__(
        self, allow_all_bundle_ids: bool = None, allowed_bundle_ids: list = None
    ):
        self.allow_all_bundle_ids = allow_all_bundle_ids
        self.allowed_bundle_ids = allowed_bundle_ids

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = key_pb2.RecaptchaenterpriseBetaKeyIosSettings()
        if Primitive.to_proto(resource.allow_all_bundle_ids):
            res.allow_all_bundle_ids = Primitive.to_proto(resource.allow_all_bundle_ids)
        if Primitive.to_proto(resource.allowed_bundle_ids):
            res.allowed_bundle_ids.extend(
                Primitive.to_proto(resource.allowed_bundle_ids)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyIosSettings(
            allow_all_bundle_ids=Primitive.from_proto(resource.allow_all_bundle_ids),
            allowed_bundle_ids=Primitive.from_proto(resource.allowed_bundle_ids),
        )


class KeyIosSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KeyIosSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KeyIosSettings.from_proto(i) for i in resources]


class KeyTestingOptions(object):
    def __init__(self, testing_score: float = None, testing_challenge: str = None):
        self.testing_score = testing_score
        self.testing_challenge = testing_challenge

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = key_pb2.RecaptchaenterpriseBetaKeyTestingOptions()
        if Primitive.to_proto(resource.testing_score):
            res.testing_score = Primitive.to_proto(resource.testing_score)
        if KeyTestingOptionsTestingChallengeEnum.to_proto(resource.testing_challenge):
            res.testing_challenge = KeyTestingOptionsTestingChallengeEnum.to_proto(
                resource.testing_challenge
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyTestingOptions(
            testing_score=Primitive.from_proto(resource.testing_score),
            testing_challenge=KeyTestingOptionsTestingChallengeEnum.from_proto(
                resource.testing_challenge
            ),
        )


class KeyTestingOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KeyTestingOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KeyTestingOptions.from_proto(i) for i in resources]


class KeyWafSettings(object):
    def __init__(self, waf_service: str = None, waf_feature: str = None):
        self.waf_service = waf_service
        self.waf_feature = waf_feature

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = key_pb2.RecaptchaenterpriseBetaKeyWafSettings()
        if KeyWafSettingsWafServiceEnum.to_proto(resource.waf_service):
            res.waf_service = KeyWafSettingsWafServiceEnum.to_proto(
                resource.waf_service
            )
        if KeyWafSettingsWafFeatureEnum.to_proto(resource.waf_feature):
            res.waf_feature = KeyWafSettingsWafFeatureEnum.to_proto(
                resource.waf_feature
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyWafSettings(
            waf_service=KeyWafSettingsWafServiceEnum.from_proto(resource.waf_service),
            waf_feature=KeyWafSettingsWafFeatureEnum.from_proto(resource.waf_feature),
        )


class KeyWafSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KeyWafSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KeyWafSettings.from_proto(i) for i in resources]


class KeyWebSettingsIntegrationTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return key_pb2.RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum.Value(
            "RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return key_pb2.RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum.Name(
            resource
        )[len("RecaptchaenterpriseBetaKeyWebSettingsIntegrationTypeEnum") :]


class KeyWebSettingsChallengeSecurityPreferenceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return key_pb2.RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum.Value(
            "RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return key_pb2.RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum.Name(
            resource
        )[
            len(
                "RecaptchaenterpriseBetaKeyWebSettingsChallengeSecurityPreferenceEnum"
            ) :
        ]


class KeyTestingOptionsTestingChallengeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            key_pb2.RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum.Value(
                "RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum%s"
                % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            key_pb2.RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum.Name(
                resource
            )[len("RecaptchaenterpriseBetaKeyTestingOptionsTestingChallengeEnum") :]
        )


class KeyWafSettingsWafServiceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return key_pb2.RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum.Value(
            "RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return key_pb2.RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum.Name(
            resource
        )[len("RecaptchaenterpriseBetaKeyWafSettingsWafServiceEnum") :]


class KeyWafSettingsWafFeatureEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return key_pb2.RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum.Value(
            "RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return key_pb2.RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum.Name(
            resource
        )[len("RecaptchaenterpriseBetaKeyWafSettingsWafFeatureEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
