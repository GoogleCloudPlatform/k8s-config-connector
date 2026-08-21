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
from google3.cloud.graphite.mmv2.services.google.apikeys import key_pb2
from google3.cloud.graphite.mmv2.services.google.apikeys import key_pb2_grpc

from typing import List


class Key(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        key_string: str = None,
        uid: str = None,
        restrictions: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.restrictions = restrictions
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = key_pb2_grpc.ApikeysBetaKeyServiceStub(channel.Channel())
        request = key_pb2.ApplyApikeysBetaKeyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if KeyRestrictions.to_proto(self.restrictions):
            request.resource.restrictions.CopyFrom(
                KeyRestrictions.to_proto(self.restrictions)
            )
        else:
            request.resource.ClearField("restrictions")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyApikeysBetaKey(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.key_string = Primitive.from_proto(response.key_string)
        self.uid = Primitive.from_proto(response.uid)
        self.restrictions = KeyRestrictions.from_proto(response.restrictions)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = key_pb2_grpc.ApikeysBetaKeyServiceStub(channel.Channel())
        request = key_pb2.DeleteApikeysBetaKeyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if KeyRestrictions.to_proto(self.restrictions):
            request.resource.restrictions.CopyFrom(
                KeyRestrictions.to_proto(self.restrictions)
            )
        else:
            request.resource.ClearField("restrictions")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteApikeysBetaKey(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = key_pb2_grpc.ApikeysBetaKeyServiceStub(channel.Channel())
        request = key_pb2.ListApikeysBetaKeyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListApikeysBetaKey(request).items

    def to_proto(self):
        resource = key_pb2.ApikeysBetaKey()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if KeyRestrictions.to_proto(self.restrictions):
            resource.restrictions.CopyFrom(KeyRestrictions.to_proto(self.restrictions))
        else:
            resource.ClearField("restrictions")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class KeyRestrictions(object):
    def __init__(
        self,
        browser_key_restrictions: dict = None,
        server_key_restrictions: dict = None,
        android_key_restrictions: dict = None,
        ios_key_restrictions: dict = None,
        api_targets: list = None,
    ):
        self.browser_key_restrictions = browser_key_restrictions
        self.server_key_restrictions = server_key_restrictions
        self.android_key_restrictions = android_key_restrictions
        self.ios_key_restrictions = ios_key_restrictions
        self.api_targets = api_targets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = key_pb2.ApikeysBetaKeyRestrictions()
        if KeyRestrictionsBrowserKeyRestrictions.to_proto(
            resource.browser_key_restrictions
        ):
            res.browser_key_restrictions.CopyFrom(
                KeyRestrictionsBrowserKeyRestrictions.to_proto(
                    resource.browser_key_restrictions
                )
            )
        else:
            res.ClearField("browser_key_restrictions")
        if KeyRestrictionsServerKeyRestrictions.to_proto(
            resource.server_key_restrictions
        ):
            res.server_key_restrictions.CopyFrom(
                KeyRestrictionsServerKeyRestrictions.to_proto(
                    resource.server_key_restrictions
                )
            )
        else:
            res.ClearField("server_key_restrictions")
        if KeyRestrictionsAndroidKeyRestrictions.to_proto(
            resource.android_key_restrictions
        ):
            res.android_key_restrictions.CopyFrom(
                KeyRestrictionsAndroidKeyRestrictions.to_proto(
                    resource.android_key_restrictions
                )
            )
        else:
            res.ClearField("android_key_restrictions")
        if KeyRestrictionsIosKeyRestrictions.to_proto(resource.ios_key_restrictions):
            res.ios_key_restrictions.CopyFrom(
                KeyRestrictionsIosKeyRestrictions.to_proto(
                    resource.ios_key_restrictions
                )
            )
        else:
            res.ClearField("ios_key_restrictions")
        if KeyRestrictionsApiTargetsArray.to_proto(resource.api_targets):
            res.api_targets.extend(
                KeyRestrictionsApiTargetsArray.to_proto(resource.api_targets)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyRestrictions(
            browser_key_restrictions=KeyRestrictionsBrowserKeyRestrictions.from_proto(
                resource.browser_key_restrictions
            ),
            server_key_restrictions=KeyRestrictionsServerKeyRestrictions.from_proto(
                resource.server_key_restrictions
            ),
            android_key_restrictions=KeyRestrictionsAndroidKeyRestrictions.from_proto(
                resource.android_key_restrictions
            ),
            ios_key_restrictions=KeyRestrictionsIosKeyRestrictions.from_proto(
                resource.ios_key_restrictions
            ),
            api_targets=KeyRestrictionsApiTargetsArray.from_proto(resource.api_targets),
        )


class KeyRestrictionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KeyRestrictions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KeyRestrictions.from_proto(i) for i in resources]


class KeyRestrictionsBrowserKeyRestrictions(object):
    def __init__(self, allowed_referrers: list = None):
        self.allowed_referrers = allowed_referrers

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = key_pb2.ApikeysBetaKeyRestrictionsBrowserKeyRestrictions()
        if Primitive.to_proto(resource.allowed_referrers):
            res.allowed_referrers.extend(Primitive.to_proto(resource.allowed_referrers))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyRestrictionsBrowserKeyRestrictions(
            allowed_referrers=Primitive.from_proto(resource.allowed_referrers),
        )


class KeyRestrictionsBrowserKeyRestrictionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KeyRestrictionsBrowserKeyRestrictions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KeyRestrictionsBrowserKeyRestrictions.from_proto(i) for i in resources]


class KeyRestrictionsServerKeyRestrictions(object):
    def __init__(self, allowed_ips: list = None):
        self.allowed_ips = allowed_ips

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = key_pb2.ApikeysBetaKeyRestrictionsServerKeyRestrictions()
        if Primitive.to_proto(resource.allowed_ips):
            res.allowed_ips.extend(Primitive.to_proto(resource.allowed_ips))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyRestrictionsServerKeyRestrictions(
            allowed_ips=Primitive.from_proto(resource.allowed_ips),
        )


class KeyRestrictionsServerKeyRestrictionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KeyRestrictionsServerKeyRestrictions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KeyRestrictionsServerKeyRestrictions.from_proto(i) for i in resources]


class KeyRestrictionsAndroidKeyRestrictions(object):
    def __init__(self, allowed_applications: list = None):
        self.allowed_applications = allowed_applications

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = key_pb2.ApikeysBetaKeyRestrictionsAndroidKeyRestrictions()
        if KeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsArray.to_proto(
            resource.allowed_applications
        ):
            res.allowed_applications.extend(
                KeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsArray.to_proto(
                    resource.allowed_applications
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyRestrictionsAndroidKeyRestrictions(
            allowed_applications=KeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsArray.from_proto(
                resource.allowed_applications
            ),
        )


class KeyRestrictionsAndroidKeyRestrictionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KeyRestrictionsAndroidKeyRestrictions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KeyRestrictionsAndroidKeyRestrictions.from_proto(i) for i in resources]


class KeyRestrictionsAndroidKeyRestrictionsAllowedApplications(object):
    def __init__(self, sha1_fingerprint: str = None, package_name: str = None):
        self.sha1_fingerprint = sha1_fingerprint
        self.package_name = package_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            key_pb2.ApikeysBetaKeyRestrictionsAndroidKeyRestrictionsAllowedApplications()
        )
        if Primitive.to_proto(resource.sha1_fingerprint):
            res.sha1_fingerprint = Primitive.to_proto(resource.sha1_fingerprint)
        if Primitive.to_proto(resource.package_name):
            res.package_name = Primitive.to_proto(resource.package_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyRestrictionsAndroidKeyRestrictionsAllowedApplications(
            sha1_fingerprint=Primitive.from_proto(resource.sha1_fingerprint),
            package_name=Primitive.from_proto(resource.package_name),
        )


class KeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            KeyRestrictionsAndroidKeyRestrictionsAllowedApplications.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            KeyRestrictionsAndroidKeyRestrictionsAllowedApplications.from_proto(i)
            for i in resources
        ]


class KeyRestrictionsIosKeyRestrictions(object):
    def __init__(self, allowed_bundle_ids: list = None):
        self.allowed_bundle_ids = allowed_bundle_ids

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = key_pb2.ApikeysBetaKeyRestrictionsIosKeyRestrictions()
        if Primitive.to_proto(resource.allowed_bundle_ids):
            res.allowed_bundle_ids.extend(
                Primitive.to_proto(resource.allowed_bundle_ids)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyRestrictionsIosKeyRestrictions(
            allowed_bundle_ids=Primitive.from_proto(resource.allowed_bundle_ids),
        )


class KeyRestrictionsIosKeyRestrictionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KeyRestrictionsIosKeyRestrictions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KeyRestrictionsIosKeyRestrictions.from_proto(i) for i in resources]


class KeyRestrictionsApiTargets(object):
    def __init__(self, service: str = None, methods: list = None):
        self.service = service
        self.methods = methods

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = key_pb2.ApikeysBetaKeyRestrictionsApiTargets()
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.methods):
            res.methods.extend(Primitive.to_proto(resource.methods))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return KeyRestrictionsApiTargets(
            service=Primitive.from_proto(resource.service),
            methods=Primitive.from_proto(resource.methods),
        )


class KeyRestrictionsApiTargetsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [KeyRestrictionsApiTargets.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [KeyRestrictionsApiTargets.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
