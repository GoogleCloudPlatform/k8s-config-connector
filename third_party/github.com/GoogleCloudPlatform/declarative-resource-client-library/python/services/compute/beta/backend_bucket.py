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
from google3.cloud.graphite.mmv2.services.google.compute import backend_bucket_pb2
from google3.cloud.graphite.mmv2.services.google.compute import backend_bucket_pb2_grpc

from typing import List


class BackendBucket(object):
    def __init__(
        self,
        bucket_name: str = None,
        cdn_policy: dict = None,
        description: str = None,
        enable_cdn: bool = None,
        name: str = None,
        project: str = None,
        self_link: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.bucket_name = bucket_name
        self.cdn_policy = cdn_policy
        self.description = description
        self.enable_cdn = enable_cdn
        self.name = name
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = backend_bucket_pb2_grpc.ComputeBetaBackendBucketServiceStub(
            channel.Channel()
        )
        request = backend_bucket_pb2.ApplyComputeBetaBackendBucketRequest()
        if Primitive.to_proto(self.bucket_name):
            request.resource.bucket_name = Primitive.to_proto(self.bucket_name)

        if BackendBucketCdnPolicy.to_proto(self.cdn_policy):
            request.resource.cdn_policy.CopyFrom(
                BackendBucketCdnPolicy.to_proto(self.cdn_policy)
            )
        else:
            request.resource.ClearField("cdn_policy")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.enable_cdn):
            request.resource.enable_cdn = Primitive.to_proto(self.enable_cdn)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaBackendBucket(request)
        self.bucket_name = Primitive.from_proto(response.bucket_name)
        self.cdn_policy = BackendBucketCdnPolicy.from_proto(response.cdn_policy)
        self.description = Primitive.from_proto(response.description)
        self.enable_cdn = Primitive.from_proto(response.enable_cdn)
        self.name = Primitive.from_proto(response.name)
        self.project = Primitive.from_proto(response.project)
        self.self_link = Primitive.from_proto(response.self_link)

    def delete(self):
        stub = backend_bucket_pb2_grpc.ComputeBetaBackendBucketServiceStub(
            channel.Channel()
        )
        request = backend_bucket_pb2.DeleteComputeBetaBackendBucketRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.bucket_name):
            request.resource.bucket_name = Primitive.to_proto(self.bucket_name)

        if BackendBucketCdnPolicy.to_proto(self.cdn_policy):
            request.resource.cdn_policy.CopyFrom(
                BackendBucketCdnPolicy.to_proto(self.cdn_policy)
            )
        else:
            request.resource.ClearField("cdn_policy")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.enable_cdn):
            request.resource.enable_cdn = Primitive.to_proto(self.enable_cdn)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeBetaBackendBucket(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = backend_bucket_pb2_grpc.ComputeBetaBackendBucketServiceStub(
            channel.Channel()
        )
        request = backend_bucket_pb2.ListComputeBetaBackendBucketRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeBetaBackendBucket(request).items

    def to_proto(self):
        resource = backend_bucket_pb2.ComputeBetaBackendBucket()
        if Primitive.to_proto(self.bucket_name):
            resource.bucket_name = Primitive.to_proto(self.bucket_name)
        if BackendBucketCdnPolicy.to_proto(self.cdn_policy):
            resource.cdn_policy.CopyFrom(
                BackendBucketCdnPolicy.to_proto(self.cdn_policy)
            )
        else:
            resource.ClearField("cdn_policy")
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.enable_cdn):
            resource.enable_cdn = Primitive.to_proto(self.enable_cdn)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class BackendBucketCdnPolicy(object):
    def __init__(
        self,
        signed_url_key_names: list = None,
        signed_url_cache_max_age_sec: int = None,
    ):
        self.signed_url_key_names = signed_url_key_names
        self.signed_url_cache_max_age_sec = signed_url_cache_max_age_sec

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_bucket_pb2.ComputeBetaBackendBucketCdnPolicy()
        if Primitive.to_proto(resource.signed_url_key_names):
            res.signed_url_key_names.extend(
                Primitive.to_proto(resource.signed_url_key_names)
            )
        if Primitive.to_proto(resource.signed_url_cache_max_age_sec):
            res.signed_url_cache_max_age_sec = Primitive.to_proto(
                resource.signed_url_cache_max_age_sec
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendBucketCdnPolicy(
            signed_url_key_names=Primitive.from_proto(resource.signed_url_key_names),
            signed_url_cache_max_age_sec=Primitive.from_proto(
                resource.signed_url_cache_max_age_sec
            ),
        )


class BackendBucketCdnPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendBucketCdnPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendBucketCdnPolicy.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
