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
from google3.cloud.graphite.mmv2.services.google.storage import bucket_pb2
from google3.cloud.graphite.mmv2.services.google.storage import bucket_pb2_grpc

from typing import List


class Bucket(object):
    def __init__(
        self,
        project: str = None,
        location: str = None,
        name: str = None,
        cors: list = None,
        lifecycle: dict = None,
        logging: dict = None,
        storage_class: str = None,
        versioning: dict = None,
        website: dict = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.project = project
        self.location = location
        self.name = name
        self.cors = cors
        self.lifecycle = lifecycle
        self.logging = logging
        self.storage_class = storage_class
        self.versioning = versioning
        self.website = website
        self.service_account_file = service_account_file

    def apply(self):
        stub = bucket_pb2_grpc.StorageBucketServiceStub(channel.Channel())
        request = bucket_pb2.ApplyStorageBucketRequest()
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if BucketCorsArray.to_proto(self.cors):
            request.resource.cors.extend(BucketCorsArray.to_proto(self.cors))
        if BucketLifecycle.to_proto(self.lifecycle):
            request.resource.lifecycle.CopyFrom(
                BucketLifecycle.to_proto(self.lifecycle)
            )
        else:
            request.resource.ClearField("lifecycle")
        if BucketLogging.to_proto(self.logging):
            request.resource.logging.CopyFrom(BucketLogging.to_proto(self.logging))
        else:
            request.resource.ClearField("logging")
        if BucketStorageClassEnum.to_proto(self.storage_class):
            request.resource.storage_class = BucketStorageClassEnum.to_proto(
                self.storage_class
            )

        if BucketVersioning.to_proto(self.versioning):
            request.resource.versioning.CopyFrom(
                BucketVersioning.to_proto(self.versioning)
            )
        else:
            request.resource.ClearField("versioning")
        if BucketWebsite.to_proto(self.website):
            request.resource.website.CopyFrom(BucketWebsite.to_proto(self.website))
        else:
            request.resource.ClearField("website")
        request.service_account_file = self.service_account_file

        response = stub.ApplyStorageBucket(request)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.name = Primitive.from_proto(response.name)
        self.cors = BucketCorsArray.from_proto(response.cors)
        self.lifecycle = BucketLifecycle.from_proto(response.lifecycle)
        self.logging = BucketLogging.from_proto(response.logging)
        self.storage_class = BucketStorageClassEnum.from_proto(response.storage_class)
        self.versioning = BucketVersioning.from_proto(response.versioning)
        self.website = BucketWebsite.from_proto(response.website)

    def delete(self):
        stub = bucket_pb2_grpc.StorageBucketServiceStub(channel.Channel())
        request = bucket_pb2.DeleteStorageBucketRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if BucketCorsArray.to_proto(self.cors):
            request.resource.cors.extend(BucketCorsArray.to_proto(self.cors))
        if BucketLifecycle.to_proto(self.lifecycle):
            request.resource.lifecycle.CopyFrom(
                BucketLifecycle.to_proto(self.lifecycle)
            )
        else:
            request.resource.ClearField("lifecycle")
        if BucketLogging.to_proto(self.logging):
            request.resource.logging.CopyFrom(BucketLogging.to_proto(self.logging))
        else:
            request.resource.ClearField("logging")
        if BucketStorageClassEnum.to_proto(self.storage_class):
            request.resource.storage_class = BucketStorageClassEnum.to_proto(
                self.storage_class
            )

        if BucketVersioning.to_proto(self.versioning):
            request.resource.versioning.CopyFrom(
                BucketVersioning.to_proto(self.versioning)
            )
        else:
            request.resource.ClearField("versioning")
        if BucketWebsite.to_proto(self.website):
            request.resource.website.CopyFrom(BucketWebsite.to_proto(self.website))
        else:
            request.resource.ClearField("website")
        response = stub.DeleteStorageBucket(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = bucket_pb2_grpc.StorageBucketServiceStub(channel.Channel())
        request = bucket_pb2.ListStorageBucketRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListStorageBucket(request).items

    def to_proto(self):
        resource = bucket_pb2.StorageBucket()
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if BucketCorsArray.to_proto(self.cors):
            resource.cors.extend(BucketCorsArray.to_proto(self.cors))
        if BucketLifecycle.to_proto(self.lifecycle):
            resource.lifecycle.CopyFrom(BucketLifecycle.to_proto(self.lifecycle))
        else:
            resource.ClearField("lifecycle")
        if BucketLogging.to_proto(self.logging):
            resource.logging.CopyFrom(BucketLogging.to_proto(self.logging))
        else:
            resource.ClearField("logging")
        if BucketStorageClassEnum.to_proto(self.storage_class):
            resource.storage_class = BucketStorageClassEnum.to_proto(self.storage_class)
        if BucketVersioning.to_proto(self.versioning):
            resource.versioning.CopyFrom(BucketVersioning.to_proto(self.versioning))
        else:
            resource.ClearField("versioning")
        if BucketWebsite.to_proto(self.website):
            resource.website.CopyFrom(BucketWebsite.to_proto(self.website))
        else:
            resource.ClearField("website")
        return resource


class BucketCors(object):
    def __init__(
        self,
        max_age_seconds: int = None,
        method: list = None,
        origin: list = None,
        response_header: list = None,
    ):
        self.max_age_seconds = max_age_seconds
        self.method = method
        self.origin = origin
        self.response_header = response_header

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = bucket_pb2.StorageBucketCors()
        if Primitive.to_proto(resource.max_age_seconds):
            res.max_age_seconds = Primitive.to_proto(resource.max_age_seconds)
        if Primitive.to_proto(resource.method):
            res.method.extend(Primitive.to_proto(resource.method))
        if Primitive.to_proto(resource.origin):
            res.origin.extend(Primitive.to_proto(resource.origin))
        if Primitive.to_proto(resource.response_header):
            res.response_header.extend(Primitive.to_proto(resource.response_header))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BucketCors(
            max_age_seconds=Primitive.from_proto(resource.max_age_seconds),
            method=Primitive.from_proto(resource.method),
            origin=Primitive.from_proto(resource.origin),
            response_header=Primitive.from_proto(resource.response_header),
        )


class BucketCorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BucketCors.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BucketCors.from_proto(i) for i in resources]


class BucketLifecycle(object):
    def __init__(self, rule: list = None):
        self.rule = rule

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = bucket_pb2.StorageBucketLifecycle()
        if BucketLifecycleRuleArray.to_proto(resource.rule):
            res.rule.extend(BucketLifecycleRuleArray.to_proto(resource.rule))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BucketLifecycle(
            rule=BucketLifecycleRuleArray.from_proto(resource.rule),
        )


class BucketLifecycleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BucketLifecycle.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BucketLifecycle.from_proto(i) for i in resources]


class BucketLifecycleRule(object):
    def __init__(self, action: dict = None, condition: dict = None):
        self.action = action
        self.condition = condition

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = bucket_pb2.StorageBucketLifecycleRule()
        if BucketLifecycleRuleAction.to_proto(resource.action):
            res.action.CopyFrom(BucketLifecycleRuleAction.to_proto(resource.action))
        else:
            res.ClearField("action")
        if BucketLifecycleRuleCondition.to_proto(resource.condition):
            res.condition.CopyFrom(
                BucketLifecycleRuleCondition.to_proto(resource.condition)
            )
        else:
            res.ClearField("condition")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BucketLifecycleRule(
            action=BucketLifecycleRuleAction.from_proto(resource.action),
            condition=BucketLifecycleRuleCondition.from_proto(resource.condition),
        )


class BucketLifecycleRuleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BucketLifecycleRule.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BucketLifecycleRule.from_proto(i) for i in resources]


class BucketLifecycleRuleAction(object):
    def __init__(self, storage_class: str = None, type: str = None):
        self.storage_class = storage_class
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = bucket_pb2.StorageBucketLifecycleRuleAction()
        if Primitive.to_proto(resource.storage_class):
            res.storage_class = Primitive.to_proto(resource.storage_class)
        if BucketLifecycleRuleActionTypeEnum.to_proto(resource.type):
            res.type = BucketLifecycleRuleActionTypeEnum.to_proto(resource.type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BucketLifecycleRuleAction(
            storage_class=Primitive.from_proto(resource.storage_class),
            type=BucketLifecycleRuleActionTypeEnum.from_proto(resource.type),
        )


class BucketLifecycleRuleActionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BucketLifecycleRuleAction.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BucketLifecycleRuleAction.from_proto(i) for i in resources]


class BucketLifecycleRuleCondition(object):
    def __init__(
        self,
        age: int = None,
        created_before: str = None,
        with_state: str = None,
        matches_storage_class: list = None,
        num_newer_versions: int = None,
    ):
        self.age = age
        self.created_before = created_before
        self.with_state = with_state
        self.matches_storage_class = matches_storage_class
        self.num_newer_versions = num_newer_versions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = bucket_pb2.StorageBucketLifecycleRuleCondition()
        if Primitive.to_proto(resource.age):
            res.age = Primitive.to_proto(resource.age)
        if Primitive.to_proto(resource.created_before):
            res.created_before = Primitive.to_proto(resource.created_before)
        if BucketLifecycleRuleConditionWithStateEnum.to_proto(resource.with_state):
            res.with_state = BucketLifecycleRuleConditionWithStateEnum.to_proto(
                resource.with_state
            )
        if Primitive.to_proto(resource.matches_storage_class):
            res.matches_storage_class.extend(
                Primitive.to_proto(resource.matches_storage_class)
            )
        if Primitive.to_proto(resource.num_newer_versions):
            res.num_newer_versions = Primitive.to_proto(resource.num_newer_versions)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BucketLifecycleRuleCondition(
            age=Primitive.from_proto(resource.age),
            created_before=Primitive.from_proto(resource.created_before),
            with_state=BucketLifecycleRuleConditionWithStateEnum.from_proto(
                resource.with_state
            ),
            matches_storage_class=Primitive.from_proto(resource.matches_storage_class),
            num_newer_versions=Primitive.from_proto(resource.num_newer_versions),
        )


class BucketLifecycleRuleConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BucketLifecycleRuleCondition.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BucketLifecycleRuleCondition.from_proto(i) for i in resources]


class BucketLogging(object):
    def __init__(self, log_bucket: str = None, log_object_prefix: str = None):
        self.log_bucket = log_bucket
        self.log_object_prefix = log_object_prefix

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = bucket_pb2.StorageBucketLogging()
        if Primitive.to_proto(resource.log_bucket):
            res.log_bucket = Primitive.to_proto(resource.log_bucket)
        if Primitive.to_proto(resource.log_object_prefix):
            res.log_object_prefix = Primitive.to_proto(resource.log_object_prefix)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BucketLogging(
            log_bucket=Primitive.from_proto(resource.log_bucket),
            log_object_prefix=Primitive.from_proto(resource.log_object_prefix),
        )


class BucketLoggingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BucketLogging.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BucketLogging.from_proto(i) for i in resources]


class BucketVersioning(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = bucket_pb2.StorageBucketVersioning()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BucketVersioning(
            enabled=Primitive.from_proto(resource.enabled),
        )


class BucketVersioningArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BucketVersioning.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BucketVersioning.from_proto(i) for i in resources]


class BucketWebsite(object):
    def __init__(self, main_page_suffix: str = None, not_found_page: str = None):
        self.main_page_suffix = main_page_suffix
        self.not_found_page = not_found_page

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = bucket_pb2.StorageBucketWebsite()
        if Primitive.to_proto(resource.main_page_suffix):
            res.main_page_suffix = Primitive.to_proto(resource.main_page_suffix)
        if Primitive.to_proto(resource.not_found_page):
            res.not_found_page = Primitive.to_proto(resource.not_found_page)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BucketWebsite(
            main_page_suffix=Primitive.from_proto(resource.main_page_suffix),
            not_found_page=Primitive.from_proto(resource.not_found_page),
        )


class BucketWebsiteArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BucketWebsite.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BucketWebsite.from_proto(i) for i in resources]


class BucketLifecycleRuleActionTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return bucket_pb2.StorageBucketLifecycleRuleActionTypeEnum.Value(
            "StorageBucketLifecycleRuleActionTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return bucket_pb2.StorageBucketLifecycleRuleActionTypeEnum.Name(resource)[
            len("StorageBucketLifecycleRuleActionTypeEnum") :
        ]


class BucketLifecycleRuleConditionWithStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return bucket_pb2.StorageBucketLifecycleRuleConditionWithStateEnum.Value(
            "StorageBucketLifecycleRuleConditionWithStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return bucket_pb2.StorageBucketLifecycleRuleConditionWithStateEnum.Name(
            resource
        )[len("StorageBucketLifecycleRuleConditionWithStateEnum") :]


class BucketStorageClassEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return bucket_pb2.StorageBucketStorageClassEnum.Value(
            "StorageBucketStorageClassEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return bucket_pb2.StorageBucketStorageClassEnum.Name(resource)[
            len("StorageBucketStorageClassEnum") :
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
