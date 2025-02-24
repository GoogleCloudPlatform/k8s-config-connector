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
from google3.cloud.graphite.mmv2.services.google.firebaserules import ruleset_pb2
from google3.cloud.graphite.mmv2.services.google.firebaserules import ruleset_pb2_grpc

from typing import List


class Ruleset(object):
    def __init__(
        self,
        name: str = None,
        source: dict = None,
        create_time: str = None,
        metadata: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.source = source
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = ruleset_pb2_grpc.FirebaserulesAlphaRulesetServiceStub(channel.Channel())
        request = ruleset_pb2.ApplyFirebaserulesAlphaRulesetRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if RulesetSource.to_proto(self.source):
            request.resource.source.CopyFrom(RulesetSource.to_proto(self.source))
        else:
            request.resource.ClearField("source")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyFirebaserulesAlphaRuleset(request)
        self.name = Primitive.from_proto(response.name)
        self.source = RulesetSource.from_proto(response.source)
        self.create_time = Primitive.from_proto(response.create_time)
        self.metadata = RulesetMetadata.from_proto(response.metadata)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = ruleset_pb2_grpc.FirebaserulesAlphaRulesetServiceStub(channel.Channel())
        request = ruleset_pb2.DeleteFirebaserulesAlphaRulesetRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if RulesetSource.to_proto(self.source):
            request.resource.source.CopyFrom(RulesetSource.to_proto(self.source))
        else:
            request.resource.ClearField("source")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteFirebaserulesAlphaRuleset(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = ruleset_pb2_grpc.FirebaserulesAlphaRulesetServiceStub(channel.Channel())
        request = ruleset_pb2.ListFirebaserulesAlphaRulesetRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListFirebaserulesAlphaRuleset(request).items

    def to_proto(self):
        resource = ruleset_pb2.FirebaserulesAlphaRuleset()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if RulesetSource.to_proto(self.source):
            resource.source.CopyFrom(RulesetSource.to_proto(self.source))
        else:
            resource.ClearField("source")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class RulesetSource(object):
    def __init__(self, files: list = None, language: str = None):
        self.files = files
        self.language = language

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ruleset_pb2.FirebaserulesAlphaRulesetSource()
        if RulesetSourceFilesArray.to_proto(resource.files):
            res.files.extend(RulesetSourceFilesArray.to_proto(resource.files))
        if RulesetSourceLanguageEnum.to_proto(resource.language):
            res.language = RulesetSourceLanguageEnum.to_proto(resource.language)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RulesetSource(
            files=RulesetSourceFilesArray.from_proto(resource.files),
            language=RulesetSourceLanguageEnum.from_proto(resource.language),
        )


class RulesetSourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RulesetSource.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RulesetSource.from_proto(i) for i in resources]


class RulesetSourceFiles(object):
    def __init__(self, content: str = None, name: str = None, fingerprint: str = None):
        self.content = content
        self.name = name
        self.fingerprint = fingerprint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ruleset_pb2.FirebaserulesAlphaRulesetSourceFiles()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.fingerprint):
            res.fingerprint = Primitive.to_proto(resource.fingerprint)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RulesetSourceFiles(
            content=Primitive.from_proto(resource.content),
            name=Primitive.from_proto(resource.name),
            fingerprint=Primitive.from_proto(resource.fingerprint),
        )


class RulesetSourceFilesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RulesetSourceFiles.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RulesetSourceFiles.from_proto(i) for i in resources]


class RulesetMetadata(object):
    def __init__(self, services: list = None):
        self.services = services

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ruleset_pb2.FirebaserulesAlphaRulesetMetadata()
        if Primitive.to_proto(resource.services):
            res.services.extend(Primitive.to_proto(resource.services))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RulesetMetadata(
            services=Primitive.from_proto(resource.services),
        )


class RulesetMetadataArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RulesetMetadata.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RulesetMetadata.from_proto(i) for i in resources]


class RulesetSourceLanguageEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return ruleset_pb2.FirebaserulesAlphaRulesetSourceLanguageEnum.Value(
            "FirebaserulesAlphaRulesetSourceLanguageEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return ruleset_pb2.FirebaserulesAlphaRulesetSourceLanguageEnum.Name(resource)[
            len("FirebaserulesAlphaRulesetSourceLanguageEnum") :
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
