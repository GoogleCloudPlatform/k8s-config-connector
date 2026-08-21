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
from google3.cloud.graphite.mmv2.services.google.identity_toolkit import config_pb2
from google3.cloud.graphite.mmv2.services.google.identity_toolkit import config_pb2_grpc

from typing import List


class Config(object):
    def __init__(
        self,
        sign_in: dict = None,
        notification: dict = None,
        quota: dict = None,
        monitoring: dict = None,
        multi_tenant: dict = None,
        authorized_domains: list = None,
        subtype: str = None,
        client: dict = None,
        mfa: dict = None,
        blocking_functions: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.sign_in = sign_in
        self.notification = notification
        self.quota = quota
        self.monitoring = monitoring
        self.multi_tenant = multi_tenant
        self.authorized_domains = authorized_domains
        self.client = client
        self.mfa = mfa
        self.blocking_functions = blocking_functions
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = config_pb2_grpc.IdentitytoolkitBetaConfigServiceStub(channel.Channel())
        request = config_pb2.ApplyIdentitytoolkitBetaConfigRequest()
        if ConfigSignIn.to_proto(self.sign_in):
            request.resource.sign_in.CopyFrom(ConfigSignIn.to_proto(self.sign_in))
        else:
            request.resource.ClearField("sign_in")
        if ConfigNotification.to_proto(self.notification):
            request.resource.notification.CopyFrom(
                ConfigNotification.to_proto(self.notification)
            )
        else:
            request.resource.ClearField("notification")
        if ConfigQuota.to_proto(self.quota):
            request.resource.quota.CopyFrom(ConfigQuota.to_proto(self.quota))
        else:
            request.resource.ClearField("quota")
        if ConfigMonitoring.to_proto(self.monitoring):
            request.resource.monitoring.CopyFrom(
                ConfigMonitoring.to_proto(self.monitoring)
            )
        else:
            request.resource.ClearField("monitoring")
        if ConfigMultiTenant.to_proto(self.multi_tenant):
            request.resource.multi_tenant.CopyFrom(
                ConfigMultiTenant.to_proto(self.multi_tenant)
            )
        else:
            request.resource.ClearField("multi_tenant")
        if Primitive.to_proto(self.authorized_domains):
            request.resource.authorized_domains.extend(
                Primitive.to_proto(self.authorized_domains)
            )
        if ConfigClient.to_proto(self.client):
            request.resource.client.CopyFrom(ConfigClient.to_proto(self.client))
        else:
            request.resource.ClearField("client")
        if ConfigMfa.to_proto(self.mfa):
            request.resource.mfa.CopyFrom(ConfigMfa.to_proto(self.mfa))
        else:
            request.resource.ClearField("mfa")
        if ConfigBlockingFunctions.to_proto(self.blocking_functions):
            request.resource.blocking_functions.CopyFrom(
                ConfigBlockingFunctions.to_proto(self.blocking_functions)
            )
        else:
            request.resource.ClearField("blocking_functions")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyIdentitytoolkitBetaConfig(request)
        self.sign_in = ConfigSignIn.from_proto(response.sign_in)
        self.notification = ConfigNotification.from_proto(response.notification)
        self.quota = ConfigQuota.from_proto(response.quota)
        self.monitoring = ConfigMonitoring.from_proto(response.monitoring)
        self.multi_tenant = ConfigMultiTenant.from_proto(response.multi_tenant)
        self.authorized_domains = Primitive.from_proto(response.authorized_domains)
        self.subtype = ConfigSubtypeEnum.from_proto(response.subtype)
        self.client = ConfigClient.from_proto(response.client)
        self.mfa = ConfigMfa.from_proto(response.mfa)
        self.blocking_functions = ConfigBlockingFunctions.from_proto(
            response.blocking_functions
        )
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = config_pb2_grpc.IdentitytoolkitBetaConfigServiceStub(channel.Channel())
        request = config_pb2.DeleteIdentitytoolkitBetaConfigRequest()
        request.service_account_file = self.service_account_file
        if ConfigSignIn.to_proto(self.sign_in):
            request.resource.sign_in.CopyFrom(ConfigSignIn.to_proto(self.sign_in))
        else:
            request.resource.ClearField("sign_in")
        if ConfigNotification.to_proto(self.notification):
            request.resource.notification.CopyFrom(
                ConfigNotification.to_proto(self.notification)
            )
        else:
            request.resource.ClearField("notification")
        if ConfigQuota.to_proto(self.quota):
            request.resource.quota.CopyFrom(ConfigQuota.to_proto(self.quota))
        else:
            request.resource.ClearField("quota")
        if ConfigMonitoring.to_proto(self.monitoring):
            request.resource.monitoring.CopyFrom(
                ConfigMonitoring.to_proto(self.monitoring)
            )
        else:
            request.resource.ClearField("monitoring")
        if ConfigMultiTenant.to_proto(self.multi_tenant):
            request.resource.multi_tenant.CopyFrom(
                ConfigMultiTenant.to_proto(self.multi_tenant)
            )
        else:
            request.resource.ClearField("multi_tenant")
        if Primitive.to_proto(self.authorized_domains):
            request.resource.authorized_domains.extend(
                Primitive.to_proto(self.authorized_domains)
            )
        if ConfigClient.to_proto(self.client):
            request.resource.client.CopyFrom(ConfigClient.to_proto(self.client))
        else:
            request.resource.ClearField("client")
        if ConfigMfa.to_proto(self.mfa):
            request.resource.mfa.CopyFrom(ConfigMfa.to_proto(self.mfa))
        else:
            request.resource.ClearField("mfa")
        if ConfigBlockingFunctions.to_proto(self.blocking_functions):
            request.resource.blocking_functions.CopyFrom(
                ConfigBlockingFunctions.to_proto(self.blocking_functions)
            )
        else:
            request.resource.ClearField("blocking_functions")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteIdentitytoolkitBetaConfig(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = config_pb2_grpc.IdentitytoolkitBetaConfigServiceStub(channel.Channel())
        request = config_pb2.ListIdentitytoolkitBetaConfigRequest()
        request.service_account_file = service_account_file
        return stub.ListIdentitytoolkitBetaConfig(request).items

    def to_proto(self):
        resource = config_pb2.IdentitytoolkitBetaConfig()
        if ConfigSignIn.to_proto(self.sign_in):
            resource.sign_in.CopyFrom(ConfigSignIn.to_proto(self.sign_in))
        else:
            resource.ClearField("sign_in")
        if ConfigNotification.to_proto(self.notification):
            resource.notification.CopyFrom(
                ConfigNotification.to_proto(self.notification)
            )
        else:
            resource.ClearField("notification")
        if ConfigQuota.to_proto(self.quota):
            resource.quota.CopyFrom(ConfigQuota.to_proto(self.quota))
        else:
            resource.ClearField("quota")
        if ConfigMonitoring.to_proto(self.monitoring):
            resource.monitoring.CopyFrom(ConfigMonitoring.to_proto(self.monitoring))
        else:
            resource.ClearField("monitoring")
        if ConfigMultiTenant.to_proto(self.multi_tenant):
            resource.multi_tenant.CopyFrom(
                ConfigMultiTenant.to_proto(self.multi_tenant)
            )
        else:
            resource.ClearField("multi_tenant")
        if Primitive.to_proto(self.authorized_domains):
            resource.authorized_domains.extend(
                Primitive.to_proto(self.authorized_domains)
            )
        if ConfigClient.to_proto(self.client):
            resource.client.CopyFrom(ConfigClient.to_proto(self.client))
        else:
            resource.ClearField("client")
        if ConfigMfa.to_proto(self.mfa):
            resource.mfa.CopyFrom(ConfigMfa.to_proto(self.mfa))
        else:
            resource.ClearField("mfa")
        if ConfigBlockingFunctions.to_proto(self.blocking_functions):
            resource.blocking_functions.CopyFrom(
                ConfigBlockingFunctions.to_proto(self.blocking_functions)
            )
        else:
            resource.ClearField("blocking_functions")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class ConfigSignIn(object):
    def __init__(
        self,
        email: dict = None,
        phone_number: dict = None,
        anonymous: dict = None,
        allow_duplicate_emails: bool = None,
        hash_config: dict = None,
    ):
        self.email = email
        self.phone_number = phone_number
        self.anonymous = anonymous
        self.allow_duplicate_emails = allow_duplicate_emails
        self.hash_config = hash_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigSignIn()
        if ConfigSignInEmail.to_proto(resource.email):
            res.email.CopyFrom(ConfigSignInEmail.to_proto(resource.email))
        else:
            res.ClearField("email")
        if ConfigSignInPhoneNumber.to_proto(resource.phone_number):
            res.phone_number.CopyFrom(
                ConfigSignInPhoneNumber.to_proto(resource.phone_number)
            )
        else:
            res.ClearField("phone_number")
        if ConfigSignInAnonymous.to_proto(resource.anonymous):
            res.anonymous.CopyFrom(ConfigSignInAnonymous.to_proto(resource.anonymous))
        else:
            res.ClearField("anonymous")
        if Primitive.to_proto(resource.allow_duplicate_emails):
            res.allow_duplicate_emails = Primitive.to_proto(
                resource.allow_duplicate_emails
            )
        if ConfigSignInHashConfig.to_proto(resource.hash_config):
            res.hash_config.CopyFrom(
                ConfigSignInHashConfig.to_proto(resource.hash_config)
            )
        else:
            res.ClearField("hash_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigSignIn(
            email=ConfigSignInEmail.from_proto(resource.email),
            phone_number=ConfigSignInPhoneNumber.from_proto(resource.phone_number),
            anonymous=ConfigSignInAnonymous.from_proto(resource.anonymous),
            allow_duplicate_emails=Primitive.from_proto(
                resource.allow_duplicate_emails
            ),
            hash_config=ConfigSignInHashConfig.from_proto(resource.hash_config),
        )


class ConfigSignInArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigSignIn.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigSignIn.from_proto(i) for i in resources]


class ConfigSignInEmail(object):
    def __init__(
        self,
        enabled: bool = None,
        password_required: bool = None,
        hash_config: dict = None,
    ):
        self.enabled = enabled
        self.password_required = password_required
        self.hash_config = hash_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigSignInEmail()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.password_required):
            res.password_required = Primitive.to_proto(resource.password_required)
        if ConfigSignInEmailHashConfig.to_proto(resource.hash_config):
            res.hash_config.CopyFrom(
                ConfigSignInEmailHashConfig.to_proto(resource.hash_config)
            )
        else:
            res.ClearField("hash_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigSignInEmail(
            enabled=Primitive.from_proto(resource.enabled),
            password_required=Primitive.from_proto(resource.password_required),
            hash_config=ConfigSignInEmailHashConfig.from_proto(resource.hash_config),
        )


class ConfigSignInEmailArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigSignInEmail.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigSignInEmail.from_proto(i) for i in resources]


class ConfigSignInEmailHashConfig(object):
    def __init__(
        self,
        algorithm: str = None,
        signer_key: str = None,
        salt_separator: str = None,
        rounds: int = None,
        memory_cost: int = None,
    ):
        self.algorithm = algorithm
        self.signer_key = signer_key
        self.salt_separator = salt_separator
        self.rounds = rounds
        self.memory_cost = memory_cost

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigSignInEmailHashConfig()
        if ConfigSignInEmailHashConfigAlgorithmEnum.to_proto(resource.algorithm):
            res.algorithm = ConfigSignInEmailHashConfigAlgorithmEnum.to_proto(
                resource.algorithm
            )
        if Primitive.to_proto(resource.signer_key):
            res.signer_key = Primitive.to_proto(resource.signer_key)
        if Primitive.to_proto(resource.salt_separator):
            res.salt_separator = Primitive.to_proto(resource.salt_separator)
        if Primitive.to_proto(resource.rounds):
            res.rounds = Primitive.to_proto(resource.rounds)
        if Primitive.to_proto(resource.memory_cost):
            res.memory_cost = Primitive.to_proto(resource.memory_cost)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigSignInEmailHashConfig(
            algorithm=ConfigSignInEmailHashConfigAlgorithmEnum.from_proto(
                resource.algorithm
            ),
            signer_key=Primitive.from_proto(resource.signer_key),
            salt_separator=Primitive.from_proto(resource.salt_separator),
            rounds=Primitive.from_proto(resource.rounds),
            memory_cost=Primitive.from_proto(resource.memory_cost),
        )


class ConfigSignInEmailHashConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigSignInEmailHashConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigSignInEmailHashConfig.from_proto(i) for i in resources]


class ConfigSignInPhoneNumber(object):
    def __init__(self, enabled: bool = None, test_phone_numbers: dict = None):
        self.enabled = enabled
        self.test_phone_numbers = test_phone_numbers

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigSignInPhoneNumber()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.test_phone_numbers):
            res.test_phone_numbers = Primitive.to_proto(resource.test_phone_numbers)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigSignInPhoneNumber(
            enabled=Primitive.from_proto(resource.enabled),
            test_phone_numbers=Primitive.from_proto(resource.test_phone_numbers),
        )


class ConfigSignInPhoneNumberArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigSignInPhoneNumber.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigSignInPhoneNumber.from_proto(i) for i in resources]


class ConfigSignInAnonymous(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigSignInAnonymous()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigSignInAnonymous(
            enabled=Primitive.from_proto(resource.enabled),
        )


class ConfigSignInAnonymousArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigSignInAnonymous.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigSignInAnonymous.from_proto(i) for i in resources]


class ConfigSignInHashConfig(object):
    def __init__(
        self,
        algorithm: str = None,
        signer_key: str = None,
        salt_separator: str = None,
        rounds: int = None,
        memory_cost: int = None,
    ):
        self.algorithm = algorithm
        self.signer_key = signer_key
        self.salt_separator = salt_separator
        self.rounds = rounds
        self.memory_cost = memory_cost

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigSignInHashConfig()
        if ConfigSignInHashConfigAlgorithmEnum.to_proto(resource.algorithm):
            res.algorithm = ConfigSignInHashConfigAlgorithmEnum.to_proto(
                resource.algorithm
            )
        if Primitive.to_proto(resource.signer_key):
            res.signer_key = Primitive.to_proto(resource.signer_key)
        if Primitive.to_proto(resource.salt_separator):
            res.salt_separator = Primitive.to_proto(resource.salt_separator)
        if Primitive.to_proto(resource.rounds):
            res.rounds = Primitive.to_proto(resource.rounds)
        if Primitive.to_proto(resource.memory_cost):
            res.memory_cost = Primitive.to_proto(resource.memory_cost)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigSignInHashConfig(
            algorithm=ConfigSignInHashConfigAlgorithmEnum.from_proto(
                resource.algorithm
            ),
            signer_key=Primitive.from_proto(resource.signer_key),
            salt_separator=Primitive.from_proto(resource.salt_separator),
            rounds=Primitive.from_proto(resource.rounds),
            memory_cost=Primitive.from_proto(resource.memory_cost),
        )


class ConfigSignInHashConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigSignInHashConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigSignInHashConfig.from_proto(i) for i in resources]


class ConfigNotification(object):
    def __init__(
        self, send_email: dict = None, send_sms: dict = None, default_locale: str = None
    ):
        self.send_email = send_email
        self.send_sms = send_sms
        self.default_locale = default_locale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigNotification()
        if ConfigNotificationSendEmail.to_proto(resource.send_email):
            res.send_email.CopyFrom(
                ConfigNotificationSendEmail.to_proto(resource.send_email)
            )
        else:
            res.ClearField("send_email")
        if ConfigNotificationSendSms.to_proto(resource.send_sms):
            res.send_sms.CopyFrom(ConfigNotificationSendSms.to_proto(resource.send_sms))
        else:
            res.ClearField("send_sms")
        if Primitive.to_proto(resource.default_locale):
            res.default_locale = Primitive.to_proto(resource.default_locale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigNotification(
            send_email=ConfigNotificationSendEmail.from_proto(resource.send_email),
            send_sms=ConfigNotificationSendSms.from_proto(resource.send_sms),
            default_locale=Primitive.from_proto(resource.default_locale),
        )


class ConfigNotificationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigNotification.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigNotification.from_proto(i) for i in resources]


class ConfigNotificationSendEmail(object):
    def __init__(
        self,
        method: str = None,
        smtp: dict = None,
        reset_password_template: dict = None,
        verify_email_template: dict = None,
        change_email_template: dict = None,
        callback_uri: str = None,
        dns_info: dict = None,
        revert_second_factor_addition_template: dict = None,
    ):
        self.method = method
        self.smtp = smtp
        self.reset_password_template = reset_password_template
        self.verify_email_template = verify_email_template
        self.change_email_template = change_email_template
        self.callback_uri = callback_uri
        self.dns_info = dns_info
        self.revert_second_factor_addition_template = (
            revert_second_factor_addition_template
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigNotificationSendEmail()
        if ConfigNotificationSendEmailMethodEnum.to_proto(resource.method):
            res.method = ConfigNotificationSendEmailMethodEnum.to_proto(resource.method)
        if ConfigNotificationSendEmailSmtp.to_proto(resource.smtp):
            res.smtp.CopyFrom(ConfigNotificationSendEmailSmtp.to_proto(resource.smtp))
        else:
            res.ClearField("smtp")
        if ConfigNotificationSendEmailResetPasswordTemplate.to_proto(
            resource.reset_password_template
        ):
            res.reset_password_template.CopyFrom(
                ConfigNotificationSendEmailResetPasswordTemplate.to_proto(
                    resource.reset_password_template
                )
            )
        else:
            res.ClearField("reset_password_template")
        if ConfigNotificationSendEmailVerifyEmailTemplate.to_proto(
            resource.verify_email_template
        ):
            res.verify_email_template.CopyFrom(
                ConfigNotificationSendEmailVerifyEmailTemplate.to_proto(
                    resource.verify_email_template
                )
            )
        else:
            res.ClearField("verify_email_template")
        if ConfigNotificationSendEmailChangeEmailTemplate.to_proto(
            resource.change_email_template
        ):
            res.change_email_template.CopyFrom(
                ConfigNotificationSendEmailChangeEmailTemplate.to_proto(
                    resource.change_email_template
                )
            )
        else:
            res.ClearField("change_email_template")
        if Primitive.to_proto(resource.callback_uri):
            res.callback_uri = Primitive.to_proto(resource.callback_uri)
        if ConfigNotificationSendEmailDnsInfo.to_proto(resource.dns_info):
            res.dns_info.CopyFrom(
                ConfigNotificationSendEmailDnsInfo.to_proto(resource.dns_info)
            )
        else:
            res.ClearField("dns_info")
        if ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate.to_proto(
            resource.revert_second_factor_addition_template
        ):
            res.revert_second_factor_addition_template.CopyFrom(
                ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate.to_proto(
                    resource.revert_second_factor_addition_template
                )
            )
        else:
            res.ClearField("revert_second_factor_addition_template")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigNotificationSendEmail(
            method=ConfigNotificationSendEmailMethodEnum.from_proto(resource.method),
            smtp=ConfigNotificationSendEmailSmtp.from_proto(resource.smtp),
            reset_password_template=ConfigNotificationSendEmailResetPasswordTemplate.from_proto(
                resource.reset_password_template
            ),
            verify_email_template=ConfigNotificationSendEmailVerifyEmailTemplate.from_proto(
                resource.verify_email_template
            ),
            change_email_template=ConfigNotificationSendEmailChangeEmailTemplate.from_proto(
                resource.change_email_template
            ),
            callback_uri=Primitive.from_proto(resource.callback_uri),
            dns_info=ConfigNotificationSendEmailDnsInfo.from_proto(resource.dns_info),
            revert_second_factor_addition_template=ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate.from_proto(
                resource.revert_second_factor_addition_template
            ),
        )


class ConfigNotificationSendEmailArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigNotificationSendEmail.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigNotificationSendEmail.from_proto(i) for i in resources]


class ConfigNotificationSendEmailSmtp(object):
    def __init__(
        self,
        sender_email: str = None,
        host: str = None,
        port: int = None,
        username: str = None,
        password: str = None,
        security_mode: str = None,
    ):
        self.sender_email = sender_email
        self.host = host
        self.port = port
        self.username = username
        self.password = password
        self.security_mode = security_mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailSmtp()
        if Primitive.to_proto(resource.sender_email):
            res.sender_email = Primitive.to_proto(resource.sender_email)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        if Primitive.to_proto(resource.password):
            res.password = Primitive.to_proto(resource.password)
        if ConfigNotificationSendEmailSmtpSecurityModeEnum.to_proto(
            resource.security_mode
        ):
            res.security_mode = (
                ConfigNotificationSendEmailSmtpSecurityModeEnum.to_proto(
                    resource.security_mode
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigNotificationSendEmailSmtp(
            sender_email=Primitive.from_proto(resource.sender_email),
            host=Primitive.from_proto(resource.host),
            port=Primitive.from_proto(resource.port),
            username=Primitive.from_proto(resource.username),
            password=Primitive.from_proto(resource.password),
            security_mode=ConfigNotificationSendEmailSmtpSecurityModeEnum.from_proto(
                resource.security_mode
            ),
        )


class ConfigNotificationSendEmailSmtpArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigNotificationSendEmailSmtp.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigNotificationSendEmailSmtp.from_proto(i) for i in resources]


class ConfigNotificationSendEmailResetPasswordTemplate(object):
    def __init__(
        self,
        sender_local_part: str = None,
        subject: str = None,
        sender_display_name: str = None,
        body: str = None,
        body_format: str = None,
        reply_to: str = None,
        customized: bool = None,
    ):
        self.sender_local_part = sender_local_part
        self.subject = subject
        self.sender_display_name = sender_display_name
        self.body = body
        self.body_format = body_format
        self.reply_to = reply_to
        self.customized = customized

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplate()
        )
        if Primitive.to_proto(resource.sender_local_part):
            res.sender_local_part = Primitive.to_proto(resource.sender_local_part)
        if Primitive.to_proto(resource.subject):
            res.subject = Primitive.to_proto(resource.subject)
        if Primitive.to_proto(resource.sender_display_name):
            res.sender_display_name = Primitive.to_proto(resource.sender_display_name)
        if Primitive.to_proto(resource.body):
            res.body = Primitive.to_proto(resource.body)
        if ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum.to_proto(
            resource.body_format
        ):
            res.body_format = (
                ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum.to_proto(
                    resource.body_format
                )
            )
        if Primitive.to_proto(resource.reply_to):
            res.reply_to = Primitive.to_proto(resource.reply_to)
        if Primitive.to_proto(resource.customized):
            res.customized = Primitive.to_proto(resource.customized)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigNotificationSendEmailResetPasswordTemplate(
            sender_local_part=Primitive.from_proto(resource.sender_local_part),
            subject=Primitive.from_proto(resource.subject),
            sender_display_name=Primitive.from_proto(resource.sender_display_name),
            body=Primitive.from_proto(resource.body),
            body_format=ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum.from_proto(
                resource.body_format
            ),
            reply_to=Primitive.from_proto(resource.reply_to),
            customized=Primitive.from_proto(resource.customized),
        )


class ConfigNotificationSendEmailResetPasswordTemplateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ConfigNotificationSendEmailResetPasswordTemplate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ConfigNotificationSendEmailResetPasswordTemplate.from_proto(i)
            for i in resources
        ]


class ConfigNotificationSendEmailVerifyEmailTemplate(object):
    def __init__(
        self,
        sender_local_part: str = None,
        subject: str = None,
        sender_display_name: str = None,
        body: str = None,
        body_format: str = None,
        reply_to: str = None,
        customized: bool = None,
    ):
        self.sender_local_part = sender_local_part
        self.subject = subject
        self.sender_display_name = sender_display_name
        self.body = body
        self.body_format = body_format
        self.reply_to = reply_to
        self.customized = customized

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplate()
        )
        if Primitive.to_proto(resource.sender_local_part):
            res.sender_local_part = Primitive.to_proto(resource.sender_local_part)
        if Primitive.to_proto(resource.subject):
            res.subject = Primitive.to_proto(resource.subject)
        if Primitive.to_proto(resource.sender_display_name):
            res.sender_display_name = Primitive.to_proto(resource.sender_display_name)
        if Primitive.to_proto(resource.body):
            res.body = Primitive.to_proto(resource.body)
        if ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum.to_proto(
            resource.body_format
        ):
            res.body_format = (
                ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum.to_proto(
                    resource.body_format
                )
            )
        if Primitive.to_proto(resource.reply_to):
            res.reply_to = Primitive.to_proto(resource.reply_to)
        if Primitive.to_proto(resource.customized):
            res.customized = Primitive.to_proto(resource.customized)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigNotificationSendEmailVerifyEmailTemplate(
            sender_local_part=Primitive.from_proto(resource.sender_local_part),
            subject=Primitive.from_proto(resource.subject),
            sender_display_name=Primitive.from_proto(resource.sender_display_name),
            body=Primitive.from_proto(resource.body),
            body_format=ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum.from_proto(
                resource.body_format
            ),
            reply_to=Primitive.from_proto(resource.reply_to),
            customized=Primitive.from_proto(resource.customized),
        )


class ConfigNotificationSendEmailVerifyEmailTemplateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ConfigNotificationSendEmailVerifyEmailTemplate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ConfigNotificationSendEmailVerifyEmailTemplate.from_proto(i)
            for i in resources
        ]


class ConfigNotificationSendEmailChangeEmailTemplate(object):
    def __init__(
        self,
        sender_local_part: str = None,
        subject: str = None,
        sender_display_name: str = None,
        body: str = None,
        body_format: str = None,
        reply_to: str = None,
        customized: bool = None,
    ):
        self.sender_local_part = sender_local_part
        self.subject = subject
        self.sender_display_name = sender_display_name
        self.body = body
        self.body_format = body_format
        self.reply_to = reply_to
        self.customized = customized

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplate()
        )
        if Primitive.to_proto(resource.sender_local_part):
            res.sender_local_part = Primitive.to_proto(resource.sender_local_part)
        if Primitive.to_proto(resource.subject):
            res.subject = Primitive.to_proto(resource.subject)
        if Primitive.to_proto(resource.sender_display_name):
            res.sender_display_name = Primitive.to_proto(resource.sender_display_name)
        if Primitive.to_proto(resource.body):
            res.body = Primitive.to_proto(resource.body)
        if ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum.to_proto(
            resource.body_format
        ):
            res.body_format = (
                ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum.to_proto(
                    resource.body_format
                )
            )
        if Primitive.to_proto(resource.reply_to):
            res.reply_to = Primitive.to_proto(resource.reply_to)
        if Primitive.to_proto(resource.customized):
            res.customized = Primitive.to_proto(resource.customized)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigNotificationSendEmailChangeEmailTemplate(
            sender_local_part=Primitive.from_proto(resource.sender_local_part),
            subject=Primitive.from_proto(resource.subject),
            sender_display_name=Primitive.from_proto(resource.sender_display_name),
            body=Primitive.from_proto(resource.body),
            body_format=ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum.from_proto(
                resource.body_format
            ),
            reply_to=Primitive.from_proto(resource.reply_to),
            customized=Primitive.from_proto(resource.customized),
        )


class ConfigNotificationSendEmailChangeEmailTemplateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ConfigNotificationSendEmailChangeEmailTemplate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ConfigNotificationSendEmailChangeEmailTemplate.from_proto(i)
            for i in resources
        ]


class ConfigNotificationSendEmailDnsInfo(object):
    def __init__(
        self,
        custom_domain: str = None,
        use_custom_domain: bool = None,
        pending_custom_domain: str = None,
        custom_domain_state: str = None,
        domain_verification_request_time: str = None,
    ):
        self.custom_domain = custom_domain
        self.use_custom_domain = use_custom_domain
        self.pending_custom_domain = pending_custom_domain
        self.custom_domain_state = custom_domain_state
        self.domain_verification_request_time = domain_verification_request_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfo()
        if Primitive.to_proto(resource.custom_domain):
            res.custom_domain = Primitive.to_proto(resource.custom_domain)
        if Primitive.to_proto(resource.use_custom_domain):
            res.use_custom_domain = Primitive.to_proto(resource.use_custom_domain)
        if Primitive.to_proto(resource.pending_custom_domain):
            res.pending_custom_domain = Primitive.to_proto(
                resource.pending_custom_domain
            )
        if ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum.to_proto(
            resource.custom_domain_state
        ):
            res.custom_domain_state = (
                ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum.to_proto(
                    resource.custom_domain_state
                )
            )
        if Primitive.to_proto(resource.domain_verification_request_time):
            res.domain_verification_request_time = Primitive.to_proto(
                resource.domain_verification_request_time
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigNotificationSendEmailDnsInfo(
            custom_domain=Primitive.from_proto(resource.custom_domain),
            use_custom_domain=Primitive.from_proto(resource.use_custom_domain),
            pending_custom_domain=Primitive.from_proto(resource.pending_custom_domain),
            custom_domain_state=ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum.from_proto(
                resource.custom_domain_state
            ),
            domain_verification_request_time=Primitive.from_proto(
                resource.domain_verification_request_time
            ),
        )


class ConfigNotificationSendEmailDnsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigNotificationSendEmailDnsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigNotificationSendEmailDnsInfo.from_proto(i) for i in resources]


class ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(object):
    def __init__(
        self,
        sender_local_part: str = None,
        subject: str = None,
        sender_display_name: str = None,
        body: str = None,
        body_format: str = None,
        reply_to: str = None,
        customized: bool = None,
    ):
        self.sender_local_part = sender_local_part
        self.subject = subject
        self.sender_display_name = sender_display_name
        self.body = body
        self.body_format = body_format
        self.reply_to = reply_to
        self.customized = customized

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplate()
        )
        if Primitive.to_proto(resource.sender_local_part):
            res.sender_local_part = Primitive.to_proto(resource.sender_local_part)
        if Primitive.to_proto(resource.subject):
            res.subject = Primitive.to_proto(resource.subject)
        if Primitive.to_proto(resource.sender_display_name):
            res.sender_display_name = Primitive.to_proto(resource.sender_display_name)
        if Primitive.to_proto(resource.body):
            res.body = Primitive.to_proto(resource.body)
        if ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum.to_proto(
            resource.body_format
        ):
            res.body_format = ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum.to_proto(
                resource.body_format
            )
        if Primitive.to_proto(resource.reply_to):
            res.reply_to = Primitive.to_proto(resource.reply_to)
        if Primitive.to_proto(resource.customized):
            res.customized = Primitive.to_proto(resource.customized)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(
            sender_local_part=Primitive.from_proto(resource.sender_local_part),
            subject=Primitive.from_proto(resource.subject),
            sender_display_name=Primitive.from_proto(resource.sender_display_name),
            body=Primitive.from_proto(resource.body),
            body_format=ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum.from_proto(
                resource.body_format
            ),
            reply_to=Primitive.from_proto(resource.reply_to),
            customized=Primitive.from_proto(resource.customized),
        )


class ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate.from_proto(i)
            for i in resources
        ]


class ConfigNotificationSendSms(object):
    def __init__(self, use_device_locale: bool = None, sms_template: dict = None):
        self.use_device_locale = use_device_locale
        self.sms_template = sms_template

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigNotificationSendSms()
        if Primitive.to_proto(resource.use_device_locale):
            res.use_device_locale = Primitive.to_proto(resource.use_device_locale)
        if ConfigNotificationSendSmsSmsTemplate.to_proto(resource.sms_template):
            res.sms_template.CopyFrom(
                ConfigNotificationSendSmsSmsTemplate.to_proto(resource.sms_template)
            )
        else:
            res.ClearField("sms_template")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigNotificationSendSms(
            use_device_locale=Primitive.from_proto(resource.use_device_locale),
            sms_template=ConfigNotificationSendSmsSmsTemplate.from_proto(
                resource.sms_template
            ),
        )


class ConfigNotificationSendSmsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigNotificationSendSms.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigNotificationSendSms.from_proto(i) for i in resources]


class ConfigNotificationSendSmsSmsTemplate(object):
    def __init__(self, content: str = None):
        self.content = content

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigNotificationSendSmsSmsTemplate()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigNotificationSendSmsSmsTemplate(
            content=Primitive.from_proto(resource.content),
        )


class ConfigNotificationSendSmsSmsTemplateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigNotificationSendSmsSmsTemplate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigNotificationSendSmsSmsTemplate.from_proto(i) for i in resources]


class ConfigQuota(object):
    def __init__(self, sign_up_quota_config: dict = None):
        self.sign_up_quota_config = sign_up_quota_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigQuota()
        if ConfigQuotaSignUpQuotaConfig.to_proto(resource.sign_up_quota_config):
            res.sign_up_quota_config.CopyFrom(
                ConfigQuotaSignUpQuotaConfig.to_proto(resource.sign_up_quota_config)
            )
        else:
            res.ClearField("sign_up_quota_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigQuota(
            sign_up_quota_config=ConfigQuotaSignUpQuotaConfig.from_proto(
                resource.sign_up_quota_config
            ),
        )


class ConfigQuotaArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigQuota.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigQuota.from_proto(i) for i in resources]


class ConfigQuotaSignUpQuotaConfig(object):
    def __init__(
        self, quota: int = None, start_time: str = None, quota_duration: str = None
    ):
        self.quota = quota
        self.start_time = start_time
        self.quota_duration = quota_duration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigQuotaSignUpQuotaConfig()
        if Primitive.to_proto(resource.quota):
            res.quota = Primitive.to_proto(resource.quota)
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.quota_duration):
            res.quota_duration = Primitive.to_proto(resource.quota_duration)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigQuotaSignUpQuotaConfig(
            quota=Primitive.from_proto(resource.quota),
            start_time=Primitive.from_proto(resource.start_time),
            quota_duration=Primitive.from_proto(resource.quota_duration),
        )


class ConfigQuotaSignUpQuotaConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigQuotaSignUpQuotaConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigQuotaSignUpQuotaConfig.from_proto(i) for i in resources]


class ConfigMonitoring(object):
    def __init__(self, request_logging: dict = None):
        self.request_logging = request_logging

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigMonitoring()
        if ConfigMonitoringRequestLogging.to_proto(resource.request_logging):
            res.request_logging.CopyFrom(
                ConfigMonitoringRequestLogging.to_proto(resource.request_logging)
            )
        else:
            res.ClearField("request_logging")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigMonitoring(
            request_logging=ConfigMonitoringRequestLogging.from_proto(
                resource.request_logging
            ),
        )


class ConfigMonitoringArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigMonitoring.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigMonitoring.from_proto(i) for i in resources]


class ConfigMonitoringRequestLogging(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigMonitoringRequestLogging()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigMonitoringRequestLogging(
            enabled=Primitive.from_proto(resource.enabled),
        )


class ConfigMonitoringRequestLoggingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigMonitoringRequestLogging.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigMonitoringRequestLogging.from_proto(i) for i in resources]


class ConfigMultiTenant(object):
    def __init__(self, allow_tenants: bool = None, default_tenant_location: str = None):
        self.allow_tenants = allow_tenants
        self.default_tenant_location = default_tenant_location

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigMultiTenant()
        if Primitive.to_proto(resource.allow_tenants):
            res.allow_tenants = Primitive.to_proto(resource.allow_tenants)
        if Primitive.to_proto(resource.default_tenant_location):
            res.default_tenant_location = Primitive.to_proto(
                resource.default_tenant_location
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigMultiTenant(
            allow_tenants=Primitive.from_proto(resource.allow_tenants),
            default_tenant_location=Primitive.from_proto(
                resource.default_tenant_location
            ),
        )


class ConfigMultiTenantArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigMultiTenant.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigMultiTenant.from_proto(i) for i in resources]


class ConfigClient(object):
    def __init__(
        self,
        api_key: str = None,
        permissions: dict = None,
        firebase_subdomain: str = None,
    ):
        self.api_key = api_key
        self.permissions = permissions
        self.firebase_subdomain = firebase_subdomain

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigClient()
        if Primitive.to_proto(resource.api_key):
            res.api_key = Primitive.to_proto(resource.api_key)
        if ConfigClientPermissions.to_proto(resource.permissions):
            res.permissions.CopyFrom(
                ConfigClientPermissions.to_proto(resource.permissions)
            )
        else:
            res.ClearField("permissions")
        if Primitive.to_proto(resource.firebase_subdomain):
            res.firebase_subdomain = Primitive.to_proto(resource.firebase_subdomain)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigClient(
            api_key=Primitive.from_proto(resource.api_key),
            permissions=ConfigClientPermissions.from_proto(resource.permissions),
            firebase_subdomain=Primitive.from_proto(resource.firebase_subdomain),
        )


class ConfigClientArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigClient.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigClient.from_proto(i) for i in resources]


class ConfigClientPermissions(object):
    def __init__(
        self, disabled_user_signup: bool = None, disabled_user_deletion: bool = None
    ):
        self.disabled_user_signup = disabled_user_signup
        self.disabled_user_deletion = disabled_user_deletion

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigClientPermissions()
        if Primitive.to_proto(resource.disabled_user_signup):
            res.disabled_user_signup = Primitive.to_proto(resource.disabled_user_signup)
        if Primitive.to_proto(resource.disabled_user_deletion):
            res.disabled_user_deletion = Primitive.to_proto(
                resource.disabled_user_deletion
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigClientPermissions(
            disabled_user_signup=Primitive.from_proto(resource.disabled_user_signup),
            disabled_user_deletion=Primitive.from_proto(
                resource.disabled_user_deletion
            ),
        )


class ConfigClientPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigClientPermissions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigClientPermissions.from_proto(i) for i in resources]


class ConfigMfa(object):
    def __init__(self, state: str = None):
        self.state = state

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigMfa()
        if ConfigMfaStateEnum.to_proto(resource.state):
            res.state = ConfigMfaStateEnum.to_proto(resource.state)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigMfa(
            state=ConfigMfaStateEnum.from_proto(resource.state),
        )


class ConfigMfaArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigMfa.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigMfa.from_proto(i) for i in resources]


class ConfigBlockingFunctions(object):
    def __init__(self, triggers: dict = None):
        self.triggers = triggers

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigBlockingFunctions()
        if Primitive.to_proto(resource.triggers):
            res.triggers = Primitive.to_proto(resource.triggers)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigBlockingFunctions(
            triggers=Primitive.from_proto(resource.triggers),
        )


class ConfigBlockingFunctionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigBlockingFunctions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigBlockingFunctions.from_proto(i) for i in resources]


class ConfigBlockingFunctionsTriggers(object):
    def __init__(self, function_uri: str = None, update_time: str = None):
        self.function_uri = function_uri
        self.update_time = update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = config_pb2.IdentitytoolkitBetaConfigBlockingFunctionsTriggers()
        if Primitive.to_proto(resource.function_uri):
            res.function_uri = Primitive.to_proto(resource.function_uri)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConfigBlockingFunctionsTriggers(
            function_uri=Primitive.from_proto(resource.function_uri),
            update_time=Primitive.from_proto(resource.update_time),
        )


class ConfigBlockingFunctionsTriggersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConfigBlockingFunctionsTriggers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConfigBlockingFunctionsTriggers.from_proto(i) for i in resources]


class ConfigSignInEmailHashConfigAlgorithmEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum.Value(
            "IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            config_pb2.IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum.Name(
                resource
            )[len("IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum") :]
        )


class ConfigSignInHashConfigAlgorithmEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum.Value(
            "IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum.Name(
            resource
        )[len("IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum") :]


class ConfigNotificationSendEmailMethodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum.Value(
                "IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum%s" % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum.Name(
            resource
        )[len("IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum") :]


class ConfigNotificationSendEmailSmtpSecurityModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum.Value(
            "IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum.Name(
            resource
        )[
            len("IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum") :
        ]


class ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum.Value(
            "IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum.Name(
            resource
        )[
            len(
                "IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum"
            ) :
        ]


class ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum.Value(
            "IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum.Name(
            resource
        )[
            len(
                "IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum"
            ) :
        ]


class ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum.Value(
            "IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum.Name(
            resource
        )[
            len(
                "IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum"
            ) :
        ]


class ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum.Value(
            "IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum.Name(
            resource
        )[
            len(
                "IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum"
            ) :
        ]


class ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum.Value(
            "IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum.Name(
            resource
        )[
            len(
                "IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum"
            ) :
        ]


class ConfigSubtypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigSubtypeEnum.Value(
            "IdentitytoolkitBetaConfigSubtypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigSubtypeEnum.Name(resource)[
            len("IdentitytoolkitBetaConfigSubtypeEnum") :
        ]


class ConfigMfaStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigMfaStateEnum.Value(
            "IdentitytoolkitBetaConfigMfaStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return config_pb2.IdentitytoolkitBetaConfigMfaStateEnum.Name(resource)[
            len("IdentitytoolkitBetaConfigMfaStateEnum") :
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
