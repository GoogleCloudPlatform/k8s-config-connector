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
from google3.cloud.graphite.mmv2.services.google.sourcerepo import repo_pb2
from google3.cloud.graphite.mmv2.services.google.sourcerepo import repo_pb2_grpc

from typing import List


class Repo(object):
    def __init__(
        self,
        name: str = None,
        size: int = None,
        url: str = None,
        pubsub_configs: list = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.pubsub_configs = pubsub_configs
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = repo_pb2_grpc.SourcerepoRepoServiceStub(channel.Channel())
        request = repo_pb2.ApplySourcerepoRepoRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if RepoPubsubConfigsArray.to_proto(self.pubsub_configs):
            request.resource.pubsub_configs.extend(
                RepoPubsubConfigsArray.to_proto(self.pubsub_configs)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplySourcerepoRepo(request)
        self.name = Primitive.from_proto(response.name)
        self.size = Primitive.from_proto(response.size)
        self.url = Primitive.from_proto(response.url)
        self.pubsub_configs = RepoPubsubConfigsArray.from_proto(response.pubsub_configs)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = repo_pb2_grpc.SourcerepoRepoServiceStub(channel.Channel())
        request = repo_pb2.DeleteSourcerepoRepoRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if RepoPubsubConfigsArray.to_proto(self.pubsub_configs):
            request.resource.pubsub_configs.extend(
                RepoPubsubConfigsArray.to_proto(self.pubsub_configs)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteSourcerepoRepo(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = repo_pb2_grpc.SourcerepoRepoServiceStub(channel.Channel())
        request = repo_pb2.ListSourcerepoRepoRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListSourcerepoRepo(request).items

    def to_proto(self):
        resource = repo_pb2.SourcerepoRepo()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if RepoPubsubConfigsArray.to_proto(self.pubsub_configs):
            resource.pubsub_configs.extend(
                RepoPubsubConfigsArray.to_proto(self.pubsub_configs)
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class RepoPubsubConfigs(object):
    def __init__(
        self,
        topic: str = None,
        message_format: str = None,
        service_account_email: str = None,
    ):
        self.topic = topic
        self.message_format = message_format
        self.service_account_email = service_account_email

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = repo_pb2.SourcerepoRepoPubsubConfigs()
        if Primitive.to_proto(resource.topic):
            res.topic = Primitive.to_proto(resource.topic)
        if Primitive.to_proto(resource.message_format):
            res.message_format = Primitive.to_proto(resource.message_format)
        if Primitive.to_proto(resource.service_account_email):
            res.service_account_email = Primitive.to_proto(
                resource.service_account_email
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RepoPubsubConfigs(
            topic=Primitive.from_proto(resource.topic),
            message_format=Primitive.from_proto(resource.message_format),
            service_account_email=Primitive.from_proto(resource.service_account_email),
        )


class RepoPubsubConfigsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RepoPubsubConfigs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RepoPubsubConfigs.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
