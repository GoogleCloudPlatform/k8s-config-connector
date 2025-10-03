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
from google3.cloud.graphite.mmv2.services.google.cloud_scheduler import job_pb2
from google3.cloud.graphite.mmv2.services.google.cloud_scheduler import job_pb2_grpc

from typing import List


class Job(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        pubsub_target: dict = None,
        app_engine_http_target: dict = None,
        http_target: dict = None,
        schedule: str = None,
        time_zone: str = None,
        user_update_time: str = None,
        state: str = None,
        status: dict = None,
        schedule_time: str = None,
        last_attempt_time: str = None,
        retry_config: dict = None,
        attempt_deadline: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.pubsub_target = pubsub_target
        self.app_engine_http_target = app_engine_http_target
        self.http_target = http_target
        self.schedule = schedule
        self.time_zone = time_zone
        self.retry_config = retry_config
        self.attempt_deadline = attempt_deadline
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = job_pb2_grpc.CloudschedulerJobServiceStub(channel.Channel())
        request = job_pb2.ApplyCloudschedulerJobRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if JobPubsubTarget.to_proto(self.pubsub_target):
            request.resource.pubsub_target.CopyFrom(
                JobPubsubTarget.to_proto(self.pubsub_target)
            )
        else:
            request.resource.ClearField("pubsub_target")
        if JobAppEngineHttpTarget.to_proto(self.app_engine_http_target):
            request.resource.app_engine_http_target.CopyFrom(
                JobAppEngineHttpTarget.to_proto(self.app_engine_http_target)
            )
        else:
            request.resource.ClearField("app_engine_http_target")
        if JobHttpTarget.to_proto(self.http_target):
            request.resource.http_target.CopyFrom(
                JobHttpTarget.to_proto(self.http_target)
            )
        else:
            request.resource.ClearField("http_target")
        if Primitive.to_proto(self.schedule):
            request.resource.schedule = Primitive.to_proto(self.schedule)

        if Primitive.to_proto(self.time_zone):
            request.resource.time_zone = Primitive.to_proto(self.time_zone)

        if JobRetryConfig.to_proto(self.retry_config):
            request.resource.retry_config.CopyFrom(
                JobRetryConfig.to_proto(self.retry_config)
            )
        else:
            request.resource.ClearField("retry_config")
        if Primitive.to_proto(self.attempt_deadline):
            request.resource.attempt_deadline = Primitive.to_proto(
                self.attempt_deadline
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudschedulerJob(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.pubsub_target = JobPubsubTarget.from_proto(response.pubsub_target)
        self.app_engine_http_target = JobAppEngineHttpTarget.from_proto(
            response.app_engine_http_target
        )
        self.http_target = JobHttpTarget.from_proto(response.http_target)
        self.schedule = Primitive.from_proto(response.schedule)
        self.time_zone = Primitive.from_proto(response.time_zone)
        self.user_update_time = Primitive.from_proto(response.user_update_time)
        self.state = JobStateEnum.from_proto(response.state)
        self.status = JobStatus.from_proto(response.status)
        self.schedule_time = Primitive.from_proto(response.schedule_time)
        self.last_attempt_time = Primitive.from_proto(response.last_attempt_time)
        self.retry_config = JobRetryConfig.from_proto(response.retry_config)
        self.attempt_deadline = Primitive.from_proto(response.attempt_deadline)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = job_pb2_grpc.CloudschedulerJobServiceStub(channel.Channel())
        request = job_pb2.DeleteCloudschedulerJobRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if JobPubsubTarget.to_proto(self.pubsub_target):
            request.resource.pubsub_target.CopyFrom(
                JobPubsubTarget.to_proto(self.pubsub_target)
            )
        else:
            request.resource.ClearField("pubsub_target")
        if JobAppEngineHttpTarget.to_proto(self.app_engine_http_target):
            request.resource.app_engine_http_target.CopyFrom(
                JobAppEngineHttpTarget.to_proto(self.app_engine_http_target)
            )
        else:
            request.resource.ClearField("app_engine_http_target")
        if JobHttpTarget.to_proto(self.http_target):
            request.resource.http_target.CopyFrom(
                JobHttpTarget.to_proto(self.http_target)
            )
        else:
            request.resource.ClearField("http_target")
        if Primitive.to_proto(self.schedule):
            request.resource.schedule = Primitive.to_proto(self.schedule)

        if Primitive.to_proto(self.time_zone):
            request.resource.time_zone = Primitive.to_proto(self.time_zone)

        if JobRetryConfig.to_proto(self.retry_config):
            request.resource.retry_config.CopyFrom(
                JobRetryConfig.to_proto(self.retry_config)
            )
        else:
            request.resource.ClearField("retry_config")
        if Primitive.to_proto(self.attempt_deadline):
            request.resource.attempt_deadline = Primitive.to_proto(
                self.attempt_deadline
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteCloudschedulerJob(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = job_pb2_grpc.CloudschedulerJobServiceStub(channel.Channel())
        request = job_pb2.ListCloudschedulerJobRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListCloudschedulerJob(request).items

    def to_proto(self):
        resource = job_pb2.CloudschedulerJob()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if JobPubsubTarget.to_proto(self.pubsub_target):
            resource.pubsub_target.CopyFrom(
                JobPubsubTarget.to_proto(self.pubsub_target)
            )
        else:
            resource.ClearField("pubsub_target")
        if JobAppEngineHttpTarget.to_proto(self.app_engine_http_target):
            resource.app_engine_http_target.CopyFrom(
                JobAppEngineHttpTarget.to_proto(self.app_engine_http_target)
            )
        else:
            resource.ClearField("app_engine_http_target")
        if JobHttpTarget.to_proto(self.http_target):
            resource.http_target.CopyFrom(JobHttpTarget.to_proto(self.http_target))
        else:
            resource.ClearField("http_target")
        if Primitive.to_proto(self.schedule):
            resource.schedule = Primitive.to_proto(self.schedule)
        if Primitive.to_proto(self.time_zone):
            resource.time_zone = Primitive.to_proto(self.time_zone)
        if JobRetryConfig.to_proto(self.retry_config):
            resource.retry_config.CopyFrom(JobRetryConfig.to_proto(self.retry_config))
        else:
            resource.ClearField("retry_config")
        if Primitive.to_proto(self.attempt_deadline):
            resource.attempt_deadline = Primitive.to_proto(self.attempt_deadline)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class JobPubsubTarget(object):
    def __init__(
        self, topic_name: str = None, data: str = None, attributes: dict = None
    ):
        self.topic_name = topic_name
        self.data = data
        self.attributes = attributes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.CloudschedulerJobPubsubTarget()
        if Primitive.to_proto(resource.topic_name):
            res.topic_name = Primitive.to_proto(resource.topic_name)
        if Primitive.to_proto(resource.data):
            res.data = Primitive.to_proto(resource.data)
        if Primitive.to_proto(resource.attributes):
            res.attributes = Primitive.to_proto(resource.attributes)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobPubsubTarget(
            topic_name=Primitive.from_proto(resource.topic_name),
            data=Primitive.from_proto(resource.data),
            attributes=Primitive.from_proto(resource.attributes),
        )


class JobPubsubTargetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobPubsubTarget.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobPubsubTarget.from_proto(i) for i in resources]


class JobAppEngineHttpTarget(object):
    def __init__(
        self,
        http_method: str = None,
        app_engine_routing: dict = None,
        relative_uri: str = None,
        headers: dict = None,
        body: str = None,
    ):
        self.http_method = http_method
        self.app_engine_routing = app_engine_routing
        self.relative_uri = relative_uri
        self.headers = headers
        self.body = body

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.CloudschedulerJobAppEngineHttpTarget()
        if JobAppEngineHttpTargetHttpMethodEnum.to_proto(resource.http_method):
            res.http_method = JobAppEngineHttpTargetHttpMethodEnum.to_proto(
                resource.http_method
            )
        if JobAppEngineHttpTargetAppEngineRouting.to_proto(resource.app_engine_routing):
            res.app_engine_routing.CopyFrom(
                JobAppEngineHttpTargetAppEngineRouting.to_proto(
                    resource.app_engine_routing
                )
            )
        else:
            res.ClearField("app_engine_routing")
        if Primitive.to_proto(resource.relative_uri):
            res.relative_uri = Primitive.to_proto(resource.relative_uri)
        if Primitive.to_proto(resource.headers):
            res.headers = Primitive.to_proto(resource.headers)
        if Primitive.to_proto(resource.body):
            res.body = Primitive.to_proto(resource.body)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobAppEngineHttpTarget(
            http_method=JobAppEngineHttpTargetHttpMethodEnum.from_proto(
                resource.http_method
            ),
            app_engine_routing=JobAppEngineHttpTargetAppEngineRouting.from_proto(
                resource.app_engine_routing
            ),
            relative_uri=Primitive.from_proto(resource.relative_uri),
            headers=Primitive.from_proto(resource.headers),
            body=Primitive.from_proto(resource.body),
        )


class JobAppEngineHttpTargetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobAppEngineHttpTarget.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobAppEngineHttpTarget.from_proto(i) for i in resources]


class JobAppEngineHttpTargetAppEngineRouting(object):
    def __init__(
        self,
        service: str = None,
        version: str = None,
        instance: str = None,
        host: str = None,
    ):
        self.service = service
        self.version = version
        self.instance = instance
        self.host = host

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.CloudschedulerJobAppEngineHttpTargetAppEngineRouting()
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        if Primitive.to_proto(resource.instance):
            res.instance = Primitive.to_proto(resource.instance)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobAppEngineHttpTargetAppEngineRouting(
            service=Primitive.from_proto(resource.service),
            version=Primitive.from_proto(resource.version),
            instance=Primitive.from_proto(resource.instance),
            host=Primitive.from_proto(resource.host),
        )


class JobAppEngineHttpTargetAppEngineRoutingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobAppEngineHttpTargetAppEngineRouting.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobAppEngineHttpTargetAppEngineRouting.from_proto(i) for i in resources]


class JobHttpTarget(object):
    def __init__(
        self,
        uri: str = None,
        http_method: str = None,
        headers: dict = None,
        body: str = None,
        oauth_token: dict = None,
        oidc_token: dict = None,
    ):
        self.uri = uri
        self.http_method = http_method
        self.headers = headers
        self.body = body
        self.oauth_token = oauth_token
        self.oidc_token = oidc_token

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.CloudschedulerJobHttpTarget()
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if JobHttpTargetHttpMethodEnum.to_proto(resource.http_method):
            res.http_method = JobHttpTargetHttpMethodEnum.to_proto(resource.http_method)
        if Primitive.to_proto(resource.headers):
            res.headers = Primitive.to_proto(resource.headers)
        if Primitive.to_proto(resource.body):
            res.body = Primitive.to_proto(resource.body)
        if JobHttpTargetOAuthToken.to_proto(resource.oauth_token):
            res.oauth_token.CopyFrom(
                JobHttpTargetOAuthToken.to_proto(resource.oauth_token)
            )
        else:
            res.ClearField("oauth_token")
        if JobHttpTargetOidcToken.to_proto(resource.oidc_token):
            res.oidc_token.CopyFrom(
                JobHttpTargetOidcToken.to_proto(resource.oidc_token)
            )
        else:
            res.ClearField("oidc_token")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobHttpTarget(
            uri=Primitive.from_proto(resource.uri),
            http_method=JobHttpTargetHttpMethodEnum.from_proto(resource.http_method),
            headers=Primitive.from_proto(resource.headers),
            body=Primitive.from_proto(resource.body),
            oauth_token=JobHttpTargetOAuthToken.from_proto(resource.oauth_token),
            oidc_token=JobHttpTargetOidcToken.from_proto(resource.oidc_token),
        )


class JobHttpTargetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobHttpTarget.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobHttpTarget.from_proto(i) for i in resources]


class JobHttpTargetOAuthToken(object):
    def __init__(self, service_account_email: str = None, scope: str = None):
        self.service_account_email = service_account_email
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.CloudschedulerJobHttpTargetOAuthToken()
        if Primitive.to_proto(resource.service_account_email):
            res.service_account_email = Primitive.to_proto(
                resource.service_account_email
            )
        if Primitive.to_proto(resource.scope):
            res.scope = Primitive.to_proto(resource.scope)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobHttpTargetOAuthToken(
            service_account_email=Primitive.from_proto(resource.service_account_email),
            scope=Primitive.from_proto(resource.scope),
        )


class JobHttpTargetOAuthTokenArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobHttpTargetOAuthToken.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobHttpTargetOAuthToken.from_proto(i) for i in resources]


class JobHttpTargetOidcToken(object):
    def __init__(self, service_account_email: str = None, audience: str = None):
        self.service_account_email = service_account_email
        self.audience = audience

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.CloudschedulerJobHttpTargetOidcToken()
        if Primitive.to_proto(resource.service_account_email):
            res.service_account_email = Primitive.to_proto(
                resource.service_account_email
            )
        if Primitive.to_proto(resource.audience):
            res.audience = Primitive.to_proto(resource.audience)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobHttpTargetOidcToken(
            service_account_email=Primitive.from_proto(resource.service_account_email),
            audience=Primitive.from_proto(resource.audience),
        )


class JobHttpTargetOidcTokenArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobHttpTargetOidcToken.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobHttpTargetOidcToken.from_proto(i) for i in resources]


class JobStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.CloudschedulerJobStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if JobStatusDetailsArray.to_proto(resource.details):
            res.details.extend(JobStatusDetailsArray.to_proto(resource.details))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=JobStatusDetailsArray.from_proto(resource.details),
        )


class JobStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobStatus.from_proto(i) for i in resources]


class JobStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.CloudschedulerJobStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class JobStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobStatusDetails.from_proto(i) for i in resources]


class JobRetryConfig(object):
    def __init__(
        self,
        retry_count: int = None,
        max_retry_duration: str = None,
        min_backoff_duration: str = None,
        max_backoff_duration: str = None,
        max_doublings: int = None,
    ):
        self.retry_count = retry_count
        self.max_retry_duration = max_retry_duration
        self.min_backoff_duration = min_backoff_duration
        self.max_backoff_duration = max_backoff_duration
        self.max_doublings = max_doublings

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.CloudschedulerJobRetryConfig()
        if Primitive.to_proto(resource.retry_count):
            res.retry_count = Primitive.to_proto(resource.retry_count)
        if Primitive.to_proto(resource.max_retry_duration):
            res.max_retry_duration = Primitive.to_proto(resource.max_retry_duration)
        if Primitive.to_proto(resource.min_backoff_duration):
            res.min_backoff_duration = Primitive.to_proto(resource.min_backoff_duration)
        if Primitive.to_proto(resource.max_backoff_duration):
            res.max_backoff_duration = Primitive.to_proto(resource.max_backoff_duration)
        if Primitive.to_proto(resource.max_doublings):
            res.max_doublings = Primitive.to_proto(resource.max_doublings)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobRetryConfig(
            retry_count=Primitive.from_proto(resource.retry_count),
            max_retry_duration=Primitive.from_proto(resource.max_retry_duration),
            min_backoff_duration=Primitive.from_proto(resource.min_backoff_duration),
            max_backoff_duration=Primitive.from_proto(resource.max_backoff_duration),
            max_doublings=Primitive.from_proto(resource.max_doublings),
        )


class JobRetryConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobRetryConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobRetryConfig.from_proto(i) for i in resources]


class JobAppEngineHttpTargetHttpMethodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum.Value(
            "CloudschedulerJobAppEngineHttpTargetHttpMethodEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum.Name(
            resource
        )[len("CloudschedulerJobAppEngineHttpTargetHttpMethodEnum") :]


class JobHttpTargetHttpMethodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.CloudschedulerJobHttpTargetHttpMethodEnum.Value(
            "CloudschedulerJobHttpTargetHttpMethodEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.CloudschedulerJobHttpTargetHttpMethodEnum.Name(resource)[
            len("CloudschedulerJobHttpTargetHttpMethodEnum") :
        ]


class JobStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.CloudschedulerJobStateEnum.Value(
            "CloudschedulerJobStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.CloudschedulerJobStateEnum.Name(resource)[
            len("CloudschedulerJobStateEnum") :
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
