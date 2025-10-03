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
from google3.cloud.graphite.mmv2.services.google.compute import url_map_pb2
from google3.cloud.graphite.mmv2.services.google.compute import url_map_pb2_grpc

from typing import List


class UrlMap(object):
    def __init__(
        self,
        default_route_action: dict = None,
        default_service: str = None,
        default_url_redirect: dict = None,
        description: str = None,
        self_link: str = None,
        header_action: dict = None,
        host_rule: list = None,
        name: str = None,
        path_matcher: list = None,
        region: str = None,
        test: list = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.default_route_action = default_route_action
        self.default_service = default_service
        self.default_url_redirect = default_url_redirect
        self.description = description
        self.header_action = header_action
        self.host_rule = host_rule
        self.name = name
        self.path_matcher = path_matcher
        self.region = region
        self.test = test
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = url_map_pb2_grpc.ComputeUrlMapServiceStub(channel.Channel())
        request = url_map_pb2.ApplyComputeUrlMapRequest()
        if UrlMapDefaultRouteAction.to_proto(self.default_route_action):
            request.resource.default_route_action.CopyFrom(
                UrlMapDefaultRouteAction.to_proto(self.default_route_action)
            )
        else:
            request.resource.ClearField("default_route_action")
        if Primitive.to_proto(self.default_service):
            request.resource.default_service = Primitive.to_proto(self.default_service)

        if UrlMapDefaultUrlRedirect.to_proto(self.default_url_redirect):
            request.resource.default_url_redirect.CopyFrom(
                UrlMapDefaultUrlRedirect.to_proto(self.default_url_redirect)
            )
        else:
            request.resource.ClearField("default_url_redirect")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if UrlMapHeaderAction.to_proto(self.header_action):
            request.resource.header_action.CopyFrom(
                UrlMapHeaderAction.to_proto(self.header_action)
            )
        else:
            request.resource.ClearField("header_action")
        if UrlMapHostRuleArray.to_proto(self.host_rule):
            request.resource.host_rule.extend(
                UrlMapHostRuleArray.to_proto(self.host_rule)
            )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if UrlMapPathMatcherArray.to_proto(self.path_matcher):
            request.resource.path_matcher.extend(
                UrlMapPathMatcherArray.to_proto(self.path_matcher)
            )
        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if UrlMapTestArray.to_proto(self.test):
            request.resource.test.extend(UrlMapTestArray.to_proto(self.test))
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeUrlMap(request)
        self.default_route_action = UrlMapDefaultRouteAction.from_proto(
            response.default_route_action
        )
        self.default_service = Primitive.from_proto(response.default_service)
        self.default_url_redirect = UrlMapDefaultUrlRedirect.from_proto(
            response.default_url_redirect
        )
        self.description = Primitive.from_proto(response.description)
        self.self_link = Primitive.from_proto(response.self_link)
        self.header_action = UrlMapHeaderAction.from_proto(response.header_action)
        self.host_rule = UrlMapHostRuleArray.from_proto(response.host_rule)
        self.name = Primitive.from_proto(response.name)
        self.path_matcher = UrlMapPathMatcherArray.from_proto(response.path_matcher)
        self.region = Primitive.from_proto(response.region)
        self.test = UrlMapTestArray.from_proto(response.test)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = url_map_pb2_grpc.ComputeUrlMapServiceStub(channel.Channel())
        request = url_map_pb2.DeleteComputeUrlMapRequest()
        request.service_account_file = self.service_account_file
        if UrlMapDefaultRouteAction.to_proto(self.default_route_action):
            request.resource.default_route_action.CopyFrom(
                UrlMapDefaultRouteAction.to_proto(self.default_route_action)
            )
        else:
            request.resource.ClearField("default_route_action")
        if Primitive.to_proto(self.default_service):
            request.resource.default_service = Primitive.to_proto(self.default_service)

        if UrlMapDefaultUrlRedirect.to_proto(self.default_url_redirect):
            request.resource.default_url_redirect.CopyFrom(
                UrlMapDefaultUrlRedirect.to_proto(self.default_url_redirect)
            )
        else:
            request.resource.ClearField("default_url_redirect")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if UrlMapHeaderAction.to_proto(self.header_action):
            request.resource.header_action.CopyFrom(
                UrlMapHeaderAction.to_proto(self.header_action)
            )
        else:
            request.resource.ClearField("header_action")
        if UrlMapHostRuleArray.to_proto(self.host_rule):
            request.resource.host_rule.extend(
                UrlMapHostRuleArray.to_proto(self.host_rule)
            )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if UrlMapPathMatcherArray.to_proto(self.path_matcher):
            request.resource.path_matcher.extend(
                UrlMapPathMatcherArray.to_proto(self.path_matcher)
            )
        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if UrlMapTestArray.to_proto(self.test):
            request.resource.test.extend(UrlMapTestArray.to_proto(self.test))
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeUrlMap(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = url_map_pb2_grpc.ComputeUrlMapServiceStub(channel.Channel())
        request = url_map_pb2.ListComputeUrlMapRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeUrlMap(request).items

    def to_proto(self):
        resource = url_map_pb2.ComputeUrlMap()
        if UrlMapDefaultRouteAction.to_proto(self.default_route_action):
            resource.default_route_action.CopyFrom(
                UrlMapDefaultRouteAction.to_proto(self.default_route_action)
            )
        else:
            resource.ClearField("default_route_action")
        if Primitive.to_proto(self.default_service):
            resource.default_service = Primitive.to_proto(self.default_service)
        if UrlMapDefaultUrlRedirect.to_proto(self.default_url_redirect):
            resource.default_url_redirect.CopyFrom(
                UrlMapDefaultUrlRedirect.to_proto(self.default_url_redirect)
            )
        else:
            resource.ClearField("default_url_redirect")
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if UrlMapHeaderAction.to_proto(self.header_action):
            resource.header_action.CopyFrom(
                UrlMapHeaderAction.to_proto(self.header_action)
            )
        else:
            resource.ClearField("header_action")
        if UrlMapHostRuleArray.to_proto(self.host_rule):
            resource.host_rule.extend(UrlMapHostRuleArray.to_proto(self.host_rule))
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if UrlMapPathMatcherArray.to_proto(self.path_matcher):
            resource.path_matcher.extend(
                UrlMapPathMatcherArray.to_proto(self.path_matcher)
            )
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if UrlMapTestArray.to_proto(self.test):
            resource.test.extend(UrlMapTestArray.to_proto(self.test))
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class UrlMapDefaultRouteAction(object):
    def __init__(
        self,
        weighted_backend_service: list = None,
        url_rewrite: dict = None,
        timeout: dict = None,
        retry_policy: dict = None,
        request_mirror_policy: dict = None,
        cors_policy: dict = None,
        fault_injection_policy: dict = None,
    ):
        self.weighted_backend_service = weighted_backend_service
        self.url_rewrite = url_rewrite
        self.timeout = timeout
        self.retry_policy = retry_policy
        self.request_mirror_policy = request_mirror_policy
        self.cors_policy = cors_policy
        self.fault_injection_policy = fault_injection_policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultRouteAction()
        if UrlMapDefaultRouteActionWeightedBackendServiceArray.to_proto(
            resource.weighted_backend_service
        ):
            res.weighted_backend_service.extend(
                UrlMapDefaultRouteActionWeightedBackendServiceArray.to_proto(
                    resource.weighted_backend_service
                )
            )
        if UrlMapDefaultRouteActionUrlRewrite.to_proto(resource.url_rewrite):
            res.url_rewrite.CopyFrom(
                UrlMapDefaultRouteActionUrlRewrite.to_proto(resource.url_rewrite)
            )
        else:
            res.ClearField("url_rewrite")
        if UrlMapDefaultRouteActionTimeout.to_proto(resource.timeout):
            res.timeout.CopyFrom(
                UrlMapDefaultRouteActionTimeout.to_proto(resource.timeout)
            )
        else:
            res.ClearField("timeout")
        if UrlMapDefaultRouteActionRetryPolicy.to_proto(resource.retry_policy):
            res.retry_policy.CopyFrom(
                UrlMapDefaultRouteActionRetryPolicy.to_proto(resource.retry_policy)
            )
        else:
            res.ClearField("retry_policy")
        if UrlMapDefaultRouteActionRequestMirrorPolicy.to_proto(
            resource.request_mirror_policy
        ):
            res.request_mirror_policy.CopyFrom(
                UrlMapDefaultRouteActionRequestMirrorPolicy.to_proto(
                    resource.request_mirror_policy
                )
            )
        else:
            res.ClearField("request_mirror_policy")
        if UrlMapDefaultRouteActionCorsPolicy.to_proto(resource.cors_policy):
            res.cors_policy.CopyFrom(
                UrlMapDefaultRouteActionCorsPolicy.to_proto(resource.cors_policy)
            )
        else:
            res.ClearField("cors_policy")
        if UrlMapDefaultRouteActionFaultInjectionPolicy.to_proto(
            resource.fault_injection_policy
        ):
            res.fault_injection_policy.CopyFrom(
                UrlMapDefaultRouteActionFaultInjectionPolicy.to_proto(
                    resource.fault_injection_policy
                )
            )
        else:
            res.ClearField("fault_injection_policy")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteAction(
            weighted_backend_service=UrlMapDefaultRouteActionWeightedBackendServiceArray.from_proto(
                resource.weighted_backend_service
            ),
            url_rewrite=UrlMapDefaultRouteActionUrlRewrite.from_proto(
                resource.url_rewrite
            ),
            timeout=UrlMapDefaultRouteActionTimeout.from_proto(resource.timeout),
            retry_policy=UrlMapDefaultRouteActionRetryPolicy.from_proto(
                resource.retry_policy
            ),
            request_mirror_policy=UrlMapDefaultRouteActionRequestMirrorPolicy.from_proto(
                resource.request_mirror_policy
            ),
            cors_policy=UrlMapDefaultRouteActionCorsPolicy.from_proto(
                resource.cors_policy
            ),
            fault_injection_policy=UrlMapDefaultRouteActionFaultInjectionPolicy.from_proto(
                resource.fault_injection_policy
            ),
        )


class UrlMapDefaultRouteActionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapDefaultRouteAction.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapDefaultRouteAction.from_proto(i) for i in resources]


class UrlMapDefaultRouteActionWeightedBackendService(object):
    def __init__(
        self,
        backend_service: str = None,
        weight: int = None,
        header_action: dict = None,
    ):
        self.backend_service = backend_service
        self.weight = weight
        self.header_action = header_action

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultRouteActionWeightedBackendService()
        if Primitive.to_proto(resource.backend_service):
            res.backend_service = Primitive.to_proto(resource.backend_service)
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if UrlMapHeaderAction.to_proto(resource.header_action):
            res.header_action.CopyFrom(
                UrlMapHeaderAction.to_proto(resource.header_action)
            )
        else:
            res.ClearField("header_action")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteActionWeightedBackendService(
            backend_service=Primitive.from_proto(resource.backend_service),
            weight=Primitive.from_proto(resource.weight),
            header_action=UrlMapHeaderAction.from_proto(resource.header_action),
        )


class UrlMapDefaultRouteActionWeightedBackendServiceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapDefaultRouteActionWeightedBackendService.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapDefaultRouteActionWeightedBackendService.from_proto(i)
            for i in resources
        ]


class UrlMapHeaderAction(object):
    def __init__(
        self,
        request_headers_to_remove: list = None,
        request_headers_to_add: list = None,
        response_headers_to_remove: list = None,
        response_headers_to_add: list = None,
    ):
        self.request_headers_to_remove = request_headers_to_remove
        self.request_headers_to_add = request_headers_to_add
        self.response_headers_to_remove = response_headers_to_remove
        self.response_headers_to_add = response_headers_to_add

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapHeaderAction()
        if Primitive.to_proto(resource.request_headers_to_remove):
            res.request_headers_to_remove.extend(
                Primitive.to_proto(resource.request_headers_to_remove)
            )
        if UrlMapHeaderActionRequestHeadersToAddArray.to_proto(
            resource.request_headers_to_add
        ):
            res.request_headers_to_add.extend(
                UrlMapHeaderActionRequestHeadersToAddArray.to_proto(
                    resource.request_headers_to_add
                )
            )
        if Primitive.to_proto(resource.response_headers_to_remove):
            res.response_headers_to_remove.extend(
                Primitive.to_proto(resource.response_headers_to_remove)
            )
        if UrlMapHeaderActionResponseHeadersToAddArray.to_proto(
            resource.response_headers_to_add
        ):
            res.response_headers_to_add.extend(
                UrlMapHeaderActionResponseHeadersToAddArray.to_proto(
                    resource.response_headers_to_add
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapHeaderAction(
            request_headers_to_remove=Primitive.from_proto(
                resource.request_headers_to_remove
            ),
            request_headers_to_add=UrlMapHeaderActionRequestHeadersToAddArray.from_proto(
                resource.request_headers_to_add
            ),
            response_headers_to_remove=Primitive.from_proto(
                resource.response_headers_to_remove
            ),
            response_headers_to_add=UrlMapHeaderActionResponseHeadersToAddArray.from_proto(
                resource.response_headers_to_add
            ),
        )


class UrlMapHeaderActionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapHeaderAction.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapHeaderAction.from_proto(i) for i in resources]


class UrlMapHeaderActionRequestHeadersToAdd(object):
    def __init__(
        self, header_name: str = None, header_value: str = None, replace: bool = None
    ):
        self.header_name = header_name
        self.header_value = header_value
        self.replace = replace

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapHeaderActionRequestHeadersToAdd()
        if Primitive.to_proto(resource.header_name):
            res.header_name = Primitive.to_proto(resource.header_name)
        if Primitive.to_proto(resource.header_value):
            res.header_value = Primitive.to_proto(resource.header_value)
        if Primitive.to_proto(resource.replace):
            res.replace = Primitive.to_proto(resource.replace)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapHeaderActionRequestHeadersToAdd(
            header_name=Primitive.from_proto(resource.header_name),
            header_value=Primitive.from_proto(resource.header_value),
            replace=Primitive.from_proto(resource.replace),
        )


class UrlMapHeaderActionRequestHeadersToAddArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapHeaderActionRequestHeadersToAdd.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapHeaderActionRequestHeadersToAdd.from_proto(i) for i in resources]


class UrlMapHeaderActionResponseHeadersToAdd(object):
    def __init__(
        self, header_name: str = None, header_value: str = None, replace: bool = None
    ):
        self.header_name = header_name
        self.header_value = header_value
        self.replace = replace

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapHeaderActionResponseHeadersToAdd()
        if Primitive.to_proto(resource.header_name):
            res.header_name = Primitive.to_proto(resource.header_name)
        if Primitive.to_proto(resource.header_value):
            res.header_value = Primitive.to_proto(resource.header_value)
        if Primitive.to_proto(resource.replace):
            res.replace = Primitive.to_proto(resource.replace)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapHeaderActionResponseHeadersToAdd(
            header_name=Primitive.from_proto(resource.header_name),
            header_value=Primitive.from_proto(resource.header_value),
            replace=Primitive.from_proto(resource.replace),
        )


class UrlMapHeaderActionResponseHeadersToAddArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapHeaderActionResponseHeadersToAdd.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapHeaderActionResponseHeadersToAdd.from_proto(i) for i in resources]


class UrlMapDefaultRouteActionUrlRewrite(object):
    def __init__(self, path_prefix_rewrite: str = None, host_rewrite: str = None):
        self.path_prefix_rewrite = path_prefix_rewrite
        self.host_rewrite = host_rewrite

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultRouteActionUrlRewrite()
        if Primitive.to_proto(resource.path_prefix_rewrite):
            res.path_prefix_rewrite = Primitive.to_proto(resource.path_prefix_rewrite)
        if Primitive.to_proto(resource.host_rewrite):
            res.host_rewrite = Primitive.to_proto(resource.host_rewrite)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteActionUrlRewrite(
            path_prefix_rewrite=Primitive.from_proto(resource.path_prefix_rewrite),
            host_rewrite=Primitive.from_proto(resource.host_rewrite),
        )


class UrlMapDefaultRouteActionUrlRewriteArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapDefaultRouteActionUrlRewrite.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapDefaultRouteActionUrlRewrite.from_proto(i) for i in resources]


class UrlMapDefaultRouteActionTimeout(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultRouteActionTimeout()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteActionTimeout(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class UrlMapDefaultRouteActionTimeoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapDefaultRouteActionTimeout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapDefaultRouteActionTimeout.from_proto(i) for i in resources]


class UrlMapDefaultRouteActionRetryPolicy(object):
    def __init__(
        self,
        retry_condition: list = None,
        num_retries: int = None,
        per_try_timeout: dict = None,
    ):
        self.retry_condition = retry_condition
        self.num_retries = num_retries
        self.per_try_timeout = per_try_timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultRouteActionRetryPolicy()
        if Primitive.to_proto(resource.retry_condition):
            res.retry_condition.extend(Primitive.to_proto(resource.retry_condition))
        if Primitive.to_proto(resource.num_retries):
            res.num_retries = Primitive.to_proto(resource.num_retries)
        if UrlMapDefaultRouteActionRetryPolicyPerTryTimeout.to_proto(
            resource.per_try_timeout
        ):
            res.per_try_timeout.CopyFrom(
                UrlMapDefaultRouteActionRetryPolicyPerTryTimeout.to_proto(
                    resource.per_try_timeout
                )
            )
        else:
            res.ClearField("per_try_timeout")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteActionRetryPolicy(
            retry_condition=Primitive.from_proto(resource.retry_condition),
            num_retries=Primitive.from_proto(resource.num_retries),
            per_try_timeout=UrlMapDefaultRouteActionRetryPolicyPerTryTimeout.from_proto(
                resource.per_try_timeout
            ),
        )


class UrlMapDefaultRouteActionRetryPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapDefaultRouteActionRetryPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapDefaultRouteActionRetryPolicy.from_proto(i) for i in resources]


class UrlMapDefaultRouteActionRetryPolicyPerTryTimeout(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultRouteActionRetryPolicyPerTryTimeout()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteActionRetryPolicyPerTryTimeout(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class UrlMapDefaultRouteActionRetryPolicyPerTryTimeoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapDefaultRouteActionRetryPolicyPerTryTimeout.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapDefaultRouteActionRetryPolicyPerTryTimeout.from_proto(i)
            for i in resources
        ]


class UrlMapDefaultRouteActionRequestMirrorPolicy(object):
    def __init__(self, backend_service: str = None):
        self.backend_service = backend_service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultRouteActionRequestMirrorPolicy()
        if Primitive.to_proto(resource.backend_service):
            res.backend_service = Primitive.to_proto(resource.backend_service)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteActionRequestMirrorPolicy(
            backend_service=Primitive.from_proto(resource.backend_service),
        )


class UrlMapDefaultRouteActionRequestMirrorPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapDefaultRouteActionRequestMirrorPolicy.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapDefaultRouteActionRequestMirrorPolicy.from_proto(i) for i in resources
        ]


class UrlMapDefaultRouteActionCorsPolicy(object):
    def __init__(
        self,
        allow_origin: list = None,
        allow_origin_regex: list = None,
        allow_method: list = None,
        allow_header: list = None,
        expose_header: list = None,
        max_age: int = None,
        allow_credentials: bool = None,
        disabled: bool = None,
    ):
        self.allow_origin = allow_origin
        self.allow_origin_regex = allow_origin_regex
        self.allow_method = allow_method
        self.allow_header = allow_header
        self.expose_header = expose_header
        self.max_age = max_age
        self.allow_credentials = allow_credentials
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultRouteActionCorsPolicy()
        if Primitive.to_proto(resource.allow_origin):
            res.allow_origin.extend(Primitive.to_proto(resource.allow_origin))
        if Primitive.to_proto(resource.allow_origin_regex):
            res.allow_origin_regex.extend(
                Primitive.to_proto(resource.allow_origin_regex)
            )
        if Primitive.to_proto(resource.allow_method):
            res.allow_method.extend(Primitive.to_proto(resource.allow_method))
        if Primitive.to_proto(resource.allow_header):
            res.allow_header.extend(Primitive.to_proto(resource.allow_header))
        if Primitive.to_proto(resource.expose_header):
            res.expose_header.extend(Primitive.to_proto(resource.expose_header))
        if Primitive.to_proto(resource.max_age):
            res.max_age = Primitive.to_proto(resource.max_age)
        if Primitive.to_proto(resource.allow_credentials):
            res.allow_credentials = Primitive.to_proto(resource.allow_credentials)
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteActionCorsPolicy(
            allow_origin=Primitive.from_proto(resource.allow_origin),
            allow_origin_regex=Primitive.from_proto(resource.allow_origin_regex),
            allow_method=Primitive.from_proto(resource.allow_method),
            allow_header=Primitive.from_proto(resource.allow_header),
            expose_header=Primitive.from_proto(resource.expose_header),
            max_age=Primitive.from_proto(resource.max_age),
            allow_credentials=Primitive.from_proto(resource.allow_credentials),
            disabled=Primitive.from_proto(resource.disabled),
        )


class UrlMapDefaultRouteActionCorsPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapDefaultRouteActionCorsPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapDefaultRouteActionCorsPolicy.from_proto(i) for i in resources]


class UrlMapDefaultRouteActionFaultInjectionPolicy(object):
    def __init__(self, delay: dict = None, abort: dict = None):
        self.delay = delay
        self.abort = abort

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultRouteActionFaultInjectionPolicy()
        if UrlMapDefaultRouteActionFaultInjectionPolicyDelay.to_proto(resource.delay):
            res.delay.CopyFrom(
                UrlMapDefaultRouteActionFaultInjectionPolicyDelay.to_proto(
                    resource.delay
                )
            )
        else:
            res.ClearField("delay")
        if UrlMapDefaultRouteActionFaultInjectionPolicyAbort.to_proto(resource.abort):
            res.abort.CopyFrom(
                UrlMapDefaultRouteActionFaultInjectionPolicyAbort.to_proto(
                    resource.abort
                )
            )
        else:
            res.ClearField("abort")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteActionFaultInjectionPolicy(
            delay=UrlMapDefaultRouteActionFaultInjectionPolicyDelay.from_proto(
                resource.delay
            ),
            abort=UrlMapDefaultRouteActionFaultInjectionPolicyAbort.from_proto(
                resource.abort
            ),
        )


class UrlMapDefaultRouteActionFaultInjectionPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapDefaultRouteActionFaultInjectionPolicy.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapDefaultRouteActionFaultInjectionPolicy.from_proto(i)
            for i in resources
        ]


class UrlMapDefaultRouteActionFaultInjectionPolicyDelay(object):
    def __init__(self, fixed_delay: dict = None, percentage: float = None):
        self.fixed_delay = fixed_delay
        self.percentage = percentage

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelay()
        if UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay.to_proto(
            resource.fixed_delay
        ):
            res.fixed_delay.CopyFrom(
                UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay.to_proto(
                    resource.fixed_delay
                )
            )
        else:
            res.ClearField("fixed_delay")
        if Primitive.to_proto(resource.percentage):
            res.percentage = Primitive.to_proto(resource.percentage)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteActionFaultInjectionPolicyDelay(
            fixed_delay=UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay.from_proto(
                resource.fixed_delay
            ),
            percentage=Primitive.from_proto(resource.percentage),
        )


class UrlMapDefaultRouteActionFaultInjectionPolicyDelayArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapDefaultRouteActionFaultInjectionPolicyDelay.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapDefaultRouteActionFaultInjectionPolicyDelay.from_proto(i)
            for i in resources
        ]


class UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay()
        )
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay.from_proto(i)
            for i in resources
        ]


class UrlMapDefaultRouteActionFaultInjectionPolicyAbort(object):
    def __init__(self, http_status: int = None, percentage: float = None):
        self.http_status = http_status
        self.percentage = percentage

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyAbort()
        if Primitive.to_proto(resource.http_status):
            res.http_status = Primitive.to_proto(resource.http_status)
        if Primitive.to_proto(resource.percentage):
            res.percentage = Primitive.to_proto(resource.percentage)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultRouteActionFaultInjectionPolicyAbort(
            http_status=Primitive.from_proto(resource.http_status),
            percentage=Primitive.from_proto(resource.percentage),
        )


class UrlMapDefaultRouteActionFaultInjectionPolicyAbortArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapDefaultRouteActionFaultInjectionPolicyAbort.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapDefaultRouteActionFaultInjectionPolicyAbort.from_proto(i)
            for i in resources
        ]


class UrlMapDefaultUrlRedirect(object):
    def __init__(
        self,
        host_redirect: str = None,
        path_redirect: str = None,
        prefix_redirect: str = None,
        redirect_response_code: str = None,
        https_redirect: bool = None,
        strip_query: bool = None,
    ):
        self.host_redirect = host_redirect
        self.path_redirect = path_redirect
        self.prefix_redirect = prefix_redirect
        self.redirect_response_code = redirect_response_code
        self.https_redirect = https_redirect
        self.strip_query = strip_query

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapDefaultUrlRedirect()
        if Primitive.to_proto(resource.host_redirect):
            res.host_redirect = Primitive.to_proto(resource.host_redirect)
        if Primitive.to_proto(resource.path_redirect):
            res.path_redirect = Primitive.to_proto(resource.path_redirect)
        if Primitive.to_proto(resource.prefix_redirect):
            res.prefix_redirect = Primitive.to_proto(resource.prefix_redirect)
        if UrlMapDefaultUrlRedirectRedirectResponseCodeEnum.to_proto(
            resource.redirect_response_code
        ):
            res.redirect_response_code = UrlMapDefaultUrlRedirectRedirectResponseCodeEnum.to_proto(
                resource.redirect_response_code
            )
        if Primitive.to_proto(resource.https_redirect):
            res.https_redirect = Primitive.to_proto(resource.https_redirect)
        if Primitive.to_proto(resource.strip_query):
            res.strip_query = Primitive.to_proto(resource.strip_query)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapDefaultUrlRedirect(
            host_redirect=Primitive.from_proto(resource.host_redirect),
            path_redirect=Primitive.from_proto(resource.path_redirect),
            prefix_redirect=Primitive.from_proto(resource.prefix_redirect),
            redirect_response_code=UrlMapDefaultUrlRedirectRedirectResponseCodeEnum.from_proto(
                resource.redirect_response_code
            ),
            https_redirect=Primitive.from_proto(resource.https_redirect),
            strip_query=Primitive.from_proto(resource.strip_query),
        )


class UrlMapDefaultUrlRedirectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapDefaultUrlRedirect.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapDefaultUrlRedirect.from_proto(i) for i in resources]


class UrlMapHostRule(object):
    def __init__(
        self, description: str = None, host: list = None, path_matcher: str = None
    ):
        self.description = description
        self.host = host
        self.path_matcher = path_matcher

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapHostRule()
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.host):
            res.host.extend(Primitive.to_proto(resource.host))
        if Primitive.to_proto(resource.path_matcher):
            res.path_matcher = Primitive.to_proto(resource.path_matcher)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapHostRule(
            description=Primitive.from_proto(resource.description),
            host=Primitive.from_proto(resource.host),
            path_matcher=Primitive.from_proto(resource.path_matcher),
        )


class UrlMapHostRuleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapHostRule.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapHostRule.from_proto(i) for i in resources]


class UrlMapPathMatcher(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        default_service: str = None,
        default_route_action: dict = None,
        default_url_redirect: dict = None,
        path_rule: list = None,
        route_rule: list = None,
        header_action: dict = None,
    ):
        self.name = name
        self.description = description
        self.default_service = default_service
        self.default_route_action = default_route_action
        self.default_url_redirect = default_url_redirect
        self.path_rule = path_rule
        self.route_rule = route_rule
        self.header_action = header_action

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcher()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.default_service):
            res.default_service = Primitive.to_proto(resource.default_service)
        if UrlMapDefaultRouteAction.to_proto(resource.default_route_action):
            res.default_route_action.CopyFrom(
                UrlMapDefaultRouteAction.to_proto(resource.default_route_action)
            )
        else:
            res.ClearField("default_route_action")
        if UrlMapPathMatcherDefaultUrlRedirect.to_proto(resource.default_url_redirect):
            res.default_url_redirect.CopyFrom(
                UrlMapPathMatcherDefaultUrlRedirect.to_proto(
                    resource.default_url_redirect
                )
            )
        else:
            res.ClearField("default_url_redirect")
        if UrlMapPathMatcherPathRuleArray.to_proto(resource.path_rule):
            res.path_rule.extend(
                UrlMapPathMatcherPathRuleArray.to_proto(resource.path_rule)
            )
        if UrlMapPathMatcherRouteRuleArray.to_proto(resource.route_rule):
            res.route_rule.extend(
                UrlMapPathMatcherRouteRuleArray.to_proto(resource.route_rule)
            )
        if UrlMapHeaderAction.to_proto(resource.header_action):
            res.header_action.CopyFrom(
                UrlMapHeaderAction.to_proto(resource.header_action)
            )
        else:
            res.ClearField("header_action")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcher(
            name=Primitive.from_proto(resource.name),
            description=Primitive.from_proto(resource.description),
            default_service=Primitive.from_proto(resource.default_service),
            default_route_action=UrlMapDefaultRouteAction.from_proto(
                resource.default_route_action
            ),
            default_url_redirect=UrlMapPathMatcherDefaultUrlRedirect.from_proto(
                resource.default_url_redirect
            ),
            path_rule=UrlMapPathMatcherPathRuleArray.from_proto(resource.path_rule),
            route_rule=UrlMapPathMatcherRouteRuleArray.from_proto(resource.route_rule),
            header_action=UrlMapHeaderAction.from_proto(resource.header_action),
        )


class UrlMapPathMatcherArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapPathMatcher.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapPathMatcher.from_proto(i) for i in resources]


class UrlMapPathMatcherDefaultUrlRedirect(object):
    def __init__(
        self,
        host_redirect: str = None,
        path_redirect: str = None,
        prefix_redirect: str = None,
        redirect_response_code: str = None,
        https_redirect: bool = None,
        strip_query: bool = None,
    ):
        self.host_redirect = host_redirect
        self.path_redirect = path_redirect
        self.prefix_redirect = prefix_redirect
        self.redirect_response_code = redirect_response_code
        self.https_redirect = https_redirect
        self.strip_query = strip_query

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherDefaultUrlRedirect()
        if Primitive.to_proto(resource.host_redirect):
            res.host_redirect = Primitive.to_proto(resource.host_redirect)
        if Primitive.to_proto(resource.path_redirect):
            res.path_redirect = Primitive.to_proto(resource.path_redirect)
        if Primitive.to_proto(resource.prefix_redirect):
            res.prefix_redirect = Primitive.to_proto(resource.prefix_redirect)
        if UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum.to_proto(
            resource.redirect_response_code
        ):
            res.redirect_response_code = UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum.to_proto(
                resource.redirect_response_code
            )
        if Primitive.to_proto(resource.https_redirect):
            res.https_redirect = Primitive.to_proto(resource.https_redirect)
        if Primitive.to_proto(resource.strip_query):
            res.strip_query = Primitive.to_proto(resource.strip_query)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherDefaultUrlRedirect(
            host_redirect=Primitive.from_proto(resource.host_redirect),
            path_redirect=Primitive.from_proto(resource.path_redirect),
            prefix_redirect=Primitive.from_proto(resource.prefix_redirect),
            redirect_response_code=UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum.from_proto(
                resource.redirect_response_code
            ),
            https_redirect=Primitive.from_proto(resource.https_redirect),
            strip_query=Primitive.from_proto(resource.strip_query),
        )


class UrlMapPathMatcherDefaultUrlRedirectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapPathMatcherDefaultUrlRedirect.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapPathMatcherDefaultUrlRedirect.from_proto(i) for i in resources]


class UrlMapPathMatcherPathRule(object):
    def __init__(
        self,
        backend_service: str = None,
        route_action: dict = None,
        url_redirect: dict = None,
        path: list = None,
    ):
        self.backend_service = backend_service
        self.route_action = route_action
        self.url_redirect = url_redirect
        self.path = path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherPathRule()
        if Primitive.to_proto(resource.backend_service):
            res.backend_service = Primitive.to_proto(resource.backend_service)
        if UrlMapPathMatcherPathRuleRouteAction.to_proto(resource.route_action):
            res.route_action.CopyFrom(
                UrlMapPathMatcherPathRuleRouteAction.to_proto(resource.route_action)
            )
        else:
            res.ClearField("route_action")
        if UrlMapPathMatcherPathRuleUrlRedirect.to_proto(resource.url_redirect):
            res.url_redirect.CopyFrom(
                UrlMapPathMatcherPathRuleUrlRedirect.to_proto(resource.url_redirect)
            )
        else:
            res.ClearField("url_redirect")
        if Primitive.to_proto(resource.path):
            res.path.extend(Primitive.to_proto(resource.path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRule(
            backend_service=Primitive.from_proto(resource.backend_service),
            route_action=UrlMapPathMatcherPathRuleRouteAction.from_proto(
                resource.route_action
            ),
            url_redirect=UrlMapPathMatcherPathRuleUrlRedirect.from_proto(
                resource.url_redirect
            ),
            path=Primitive.from_proto(resource.path),
        )


class UrlMapPathMatcherPathRuleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapPathMatcherPathRule.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapPathMatcherPathRule.from_proto(i) for i in resources]


class UrlMapPathMatcherPathRuleRouteAction(object):
    def __init__(
        self,
        weighted_backend_service: list = None,
        url_rewrite: dict = None,
        timeout: dict = None,
        retry_policy: dict = None,
        request_mirror_policy: dict = None,
        cors_policy: dict = None,
        fault_injection_policy: dict = None,
    ):
        self.weighted_backend_service = weighted_backend_service
        self.url_rewrite = url_rewrite
        self.timeout = timeout
        self.retry_policy = retry_policy
        self.request_mirror_policy = request_mirror_policy
        self.cors_policy = cors_policy
        self.fault_injection_policy = fault_injection_policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteAction()
        if UrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceArray.to_proto(
            resource.weighted_backend_service
        ):
            res.weighted_backend_service.extend(
                UrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceArray.to_proto(
                    resource.weighted_backend_service
                )
            )
        if UrlMapPathMatcherPathRuleRouteActionUrlRewrite.to_proto(
            resource.url_rewrite
        ):
            res.url_rewrite.CopyFrom(
                UrlMapPathMatcherPathRuleRouteActionUrlRewrite.to_proto(
                    resource.url_rewrite
                )
            )
        else:
            res.ClearField("url_rewrite")
        if UrlMapPathMatcherPathRuleRouteActionTimeout.to_proto(resource.timeout):
            res.timeout.CopyFrom(
                UrlMapPathMatcherPathRuleRouteActionTimeout.to_proto(resource.timeout)
            )
        else:
            res.ClearField("timeout")
        if UrlMapPathMatcherPathRuleRouteActionRetryPolicy.to_proto(
            resource.retry_policy
        ):
            res.retry_policy.CopyFrom(
                UrlMapPathMatcherPathRuleRouteActionRetryPolicy.to_proto(
                    resource.retry_policy
                )
            )
        else:
            res.ClearField("retry_policy")
        if UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy.to_proto(
            resource.request_mirror_policy
        ):
            res.request_mirror_policy.CopyFrom(
                UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy.to_proto(
                    resource.request_mirror_policy
                )
            )
        else:
            res.ClearField("request_mirror_policy")
        if UrlMapPathMatcherPathRuleRouteActionCorsPolicy.to_proto(
            resource.cors_policy
        ):
            res.cors_policy.CopyFrom(
                UrlMapPathMatcherPathRuleRouteActionCorsPolicy.to_proto(
                    resource.cors_policy
                )
            )
        else:
            res.ClearField("cors_policy")
        if UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy.to_proto(
            resource.fault_injection_policy
        ):
            res.fault_injection_policy.CopyFrom(
                UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy.to_proto(
                    resource.fault_injection_policy
                )
            )
        else:
            res.ClearField("fault_injection_policy")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteAction(
            weighted_backend_service=UrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceArray.from_proto(
                resource.weighted_backend_service
            ),
            url_rewrite=UrlMapPathMatcherPathRuleRouteActionUrlRewrite.from_proto(
                resource.url_rewrite
            ),
            timeout=UrlMapPathMatcherPathRuleRouteActionTimeout.from_proto(
                resource.timeout
            ),
            retry_policy=UrlMapPathMatcherPathRuleRouteActionRetryPolicy.from_proto(
                resource.retry_policy
            ),
            request_mirror_policy=UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy.from_proto(
                resource.request_mirror_policy
            ),
            cors_policy=UrlMapPathMatcherPathRuleRouteActionCorsPolicy.from_proto(
                resource.cors_policy
            ),
            fault_injection_policy=UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy.from_proto(
                resource.fault_injection_policy
            ),
        )


class UrlMapPathMatcherPathRuleRouteActionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapPathMatcherPathRuleRouteAction.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapPathMatcherPathRuleRouteAction.from_proto(i) for i in resources]


class UrlMapPathMatcherPathRuleRouteActionWeightedBackendService(object):
    def __init__(
        self,
        backend_service: str = None,
        weight: int = None,
        header_action: dict = None,
    ):
        self.backend_service = backend_service
        self.weight = weight
        self.header_action = header_action

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendService()
        )
        if Primitive.to_proto(resource.backend_service):
            res.backend_service = Primitive.to_proto(resource.backend_service)
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if UrlMapHeaderAction.to_proto(resource.header_action):
            res.header_action.CopyFrom(
                UrlMapHeaderAction.to_proto(resource.header_action)
            )
        else:
            res.ClearField("header_action")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteActionWeightedBackendService(
            backend_service=Primitive.from_proto(resource.backend_service),
            weight=Primitive.from_proto(resource.weight),
            header_action=UrlMapHeaderAction.from_proto(resource.header_action),
        )


class UrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherPathRuleRouteActionWeightedBackendService.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherPathRuleRouteActionWeightedBackendService.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherPathRuleRouteActionUrlRewrite(object):
    def __init__(self, path_prefix_rewrite: str = None, host_rewrite: str = None):
        self.path_prefix_rewrite = path_prefix_rewrite
        self.host_rewrite = host_rewrite

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteActionUrlRewrite()
        if Primitive.to_proto(resource.path_prefix_rewrite):
            res.path_prefix_rewrite = Primitive.to_proto(resource.path_prefix_rewrite)
        if Primitive.to_proto(resource.host_rewrite):
            res.host_rewrite = Primitive.to_proto(resource.host_rewrite)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteActionUrlRewrite(
            path_prefix_rewrite=Primitive.from_proto(resource.path_prefix_rewrite),
            host_rewrite=Primitive.from_proto(resource.host_rewrite),
        )


class UrlMapPathMatcherPathRuleRouteActionUrlRewriteArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherPathRuleRouteActionUrlRewrite.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherPathRuleRouteActionUrlRewrite.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherPathRuleRouteActionTimeout(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteActionTimeout()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteActionTimeout(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class UrlMapPathMatcherPathRuleRouteActionTimeoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherPathRuleRouteActionTimeout.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherPathRuleRouteActionTimeout.from_proto(i) for i in resources
        ]


class UrlMapPathMatcherPathRuleRouteActionRetryPolicy(object):
    def __init__(
        self,
        retry_condition: list = None,
        num_retries: int = None,
        per_try_timeout: dict = None,
    ):
        self.retry_condition = retry_condition
        self.num_retries = num_retries
        self.per_try_timeout = per_try_timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicy()
        if Primitive.to_proto(resource.retry_condition):
            res.retry_condition.extend(Primitive.to_proto(resource.retry_condition))
        if Primitive.to_proto(resource.num_retries):
            res.num_retries = Primitive.to_proto(resource.num_retries)
        if UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout.to_proto(
            resource.per_try_timeout
        ):
            res.per_try_timeout.CopyFrom(
                UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout.to_proto(
                    resource.per_try_timeout
                )
            )
        else:
            res.ClearField("per_try_timeout")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteActionRetryPolicy(
            retry_condition=Primitive.from_proto(resource.retry_condition),
            num_retries=Primitive.from_proto(resource.num_retries),
            per_try_timeout=UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout.from_proto(
                resource.per_try_timeout
            ),
        )


class UrlMapPathMatcherPathRuleRouteActionRetryPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherPathRuleRouteActionRetryPolicy.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherPathRuleRouteActionRetryPolicy.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout()
        )
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(object):
    def __init__(self, backend_service: str = None):
        self.backend_service = backend_service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy()
        )
        if Primitive.to_proto(resource.backend_service):
            res.backend_service = Primitive.to_proto(resource.backend_service)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(
            backend_service=Primitive.from_proto(resource.backend_service),
        )


class UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherPathRuleRouteActionCorsPolicy(object):
    def __init__(
        self,
        allow_origin: list = None,
        allow_origin_regex: list = None,
        allow_method: list = None,
        allow_header: list = None,
        expose_header: list = None,
        max_age: int = None,
        allow_credentials: bool = None,
        disabled: bool = None,
    ):
        self.allow_origin = allow_origin
        self.allow_origin_regex = allow_origin_regex
        self.allow_method = allow_method
        self.allow_header = allow_header
        self.expose_header = expose_header
        self.max_age = max_age
        self.allow_credentials = allow_credentials
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicy()
        if Primitive.to_proto(resource.allow_origin):
            res.allow_origin.extend(Primitive.to_proto(resource.allow_origin))
        if Primitive.to_proto(resource.allow_origin_regex):
            res.allow_origin_regex.extend(
                Primitive.to_proto(resource.allow_origin_regex)
            )
        if Primitive.to_proto(resource.allow_method):
            res.allow_method.extend(Primitive.to_proto(resource.allow_method))
        if Primitive.to_proto(resource.allow_header):
            res.allow_header.extend(Primitive.to_proto(resource.allow_header))
        if Primitive.to_proto(resource.expose_header):
            res.expose_header.extend(Primitive.to_proto(resource.expose_header))
        if Primitive.to_proto(resource.max_age):
            res.max_age = Primitive.to_proto(resource.max_age)
        if Primitive.to_proto(resource.allow_credentials):
            res.allow_credentials = Primitive.to_proto(resource.allow_credentials)
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteActionCorsPolicy(
            allow_origin=Primitive.from_proto(resource.allow_origin),
            allow_origin_regex=Primitive.from_proto(resource.allow_origin_regex),
            allow_method=Primitive.from_proto(resource.allow_method),
            allow_header=Primitive.from_proto(resource.allow_header),
            expose_header=Primitive.from_proto(resource.expose_header),
            max_age=Primitive.from_proto(resource.max_age),
            allow_credentials=Primitive.from_proto(resource.allow_credentials),
            disabled=Primitive.from_proto(resource.disabled),
        )


class UrlMapPathMatcherPathRuleRouteActionCorsPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherPathRuleRouteActionCorsPolicy.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherPathRuleRouteActionCorsPolicy.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(object):
    def __init__(self, delay: dict = None, abort: dict = None):
        self.delay = delay
        self.abort = abort

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy()
        )
        if UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay.to_proto(
            resource.delay
        ):
            res.delay.CopyFrom(
                UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay.to_proto(
                    resource.delay
                )
            )
        else:
            res.ClearField("delay")
        if UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort.to_proto(
            resource.abort
        ):
            res.abort.CopyFrom(
                UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort.to_proto(
                    resource.abort
                )
            )
        else:
            res.ClearField("abort")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(
            delay=UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay.from_proto(
                resource.delay
            ),
            abort=UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort.from_proto(
                resource.abort
            ),
        )


class UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(object):
    def __init__(self, fixed_delay: dict = None, percentage: float = None):
        self.fixed_delay = fixed_delay
        self.percentage = percentage

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay()
        )
        if UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay.to_proto(
            resource.fixed_delay
        ):
            res.fixed_delay.CopyFrom(
                UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay.to_proto(
                    resource.fixed_delay
                )
            )
        else:
            res.ClearField("fixed_delay")
        if Primitive.to_proto(resource.percentage):
            res.percentage = Primitive.to_proto(resource.percentage)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(
            fixed_delay=UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay.from_proto(
                resource.fixed_delay
            ),
            percentage=Primitive.from_proto(resource.percentage),
        )


class UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay()
        )
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay.from_proto(
                i
            )
            for i in resources
        ]


class UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(object):
    def __init__(self, http_status: int = None, percentage: float = None):
        self.http_status = http_status
        self.percentage = percentage

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort()
        )
        if Primitive.to_proto(resource.http_status):
            res.http_status = Primitive.to_proto(resource.http_status)
        if Primitive.to_proto(resource.percentage):
            res.percentage = Primitive.to_proto(resource.percentage)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(
            http_status=Primitive.from_proto(resource.http_status),
            percentage=Primitive.from_proto(resource.percentage),
        )


class UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherPathRuleUrlRedirect(object):
    def __init__(
        self,
        host_redirect: str = None,
        path_redirect: str = None,
        prefix_redirect: str = None,
        redirect_response_code: str = None,
        https_redirect: bool = None,
        strip_query: bool = None,
    ):
        self.host_redirect = host_redirect
        self.path_redirect = path_redirect
        self.prefix_redirect = prefix_redirect
        self.redirect_response_code = redirect_response_code
        self.https_redirect = https_redirect
        self.strip_query = strip_query

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherPathRuleUrlRedirect()
        if Primitive.to_proto(resource.host_redirect):
            res.host_redirect = Primitive.to_proto(resource.host_redirect)
        if Primitive.to_proto(resource.path_redirect):
            res.path_redirect = Primitive.to_proto(resource.path_redirect)
        if Primitive.to_proto(resource.prefix_redirect):
            res.prefix_redirect = Primitive.to_proto(resource.prefix_redirect)
        if UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum.to_proto(
            resource.redirect_response_code
        ):
            res.redirect_response_code = UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum.to_proto(
                resource.redirect_response_code
            )
        if Primitive.to_proto(resource.https_redirect):
            res.https_redirect = Primitive.to_proto(resource.https_redirect)
        if Primitive.to_proto(resource.strip_query):
            res.strip_query = Primitive.to_proto(resource.strip_query)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherPathRuleUrlRedirect(
            host_redirect=Primitive.from_proto(resource.host_redirect),
            path_redirect=Primitive.from_proto(resource.path_redirect),
            prefix_redirect=Primitive.from_proto(resource.prefix_redirect),
            redirect_response_code=UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum.from_proto(
                resource.redirect_response_code
            ),
            https_redirect=Primitive.from_proto(resource.https_redirect),
            strip_query=Primitive.from_proto(resource.strip_query),
        )


class UrlMapPathMatcherPathRuleUrlRedirectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapPathMatcherPathRuleUrlRedirect.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapPathMatcherPathRuleUrlRedirect.from_proto(i) for i in resources]


class UrlMapPathMatcherRouteRule(object):
    def __init__(
        self,
        priority: int = None,
        description: str = None,
        match_rule: list = None,
        backend_service: str = None,
        route_action: dict = None,
        url_redirect: dict = None,
        header_action: dict = None,
    ):
        self.priority = priority
        self.description = description
        self.match_rule = match_rule
        self.backend_service = backend_service
        self.route_action = route_action
        self.url_redirect = url_redirect
        self.header_action = header_action

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherRouteRule()
        if Primitive.to_proto(resource.priority):
            res.priority = Primitive.to_proto(resource.priority)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if UrlMapPathMatcherRouteRuleMatchRuleArray.to_proto(resource.match_rule):
            res.match_rule.extend(
                UrlMapPathMatcherRouteRuleMatchRuleArray.to_proto(resource.match_rule)
            )
        if Primitive.to_proto(resource.backend_service):
            res.backend_service = Primitive.to_proto(resource.backend_service)
        if UrlMapPathMatcherRouteRuleRouteAction.to_proto(resource.route_action):
            res.route_action.CopyFrom(
                UrlMapPathMatcherRouteRuleRouteAction.to_proto(resource.route_action)
            )
        else:
            res.ClearField("route_action")
        if UrlMapPathMatcherRouteRuleUrlRedirect.to_proto(resource.url_redirect):
            res.url_redirect.CopyFrom(
                UrlMapPathMatcherRouteRuleUrlRedirect.to_proto(resource.url_redirect)
            )
        else:
            res.ClearField("url_redirect")
        if UrlMapHeaderAction.to_proto(resource.header_action):
            res.header_action.CopyFrom(
                UrlMapHeaderAction.to_proto(resource.header_action)
            )
        else:
            res.ClearField("header_action")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRule(
            priority=Primitive.from_proto(resource.priority),
            description=Primitive.from_proto(resource.description),
            match_rule=UrlMapPathMatcherRouteRuleMatchRuleArray.from_proto(
                resource.match_rule
            ),
            backend_service=Primitive.from_proto(resource.backend_service),
            route_action=UrlMapPathMatcherRouteRuleRouteAction.from_proto(
                resource.route_action
            ),
            url_redirect=UrlMapPathMatcherRouteRuleUrlRedirect.from_proto(
                resource.url_redirect
            ),
            header_action=UrlMapHeaderAction.from_proto(resource.header_action),
        )


class UrlMapPathMatcherRouteRuleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapPathMatcherRouteRule.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapPathMatcherRouteRule.from_proto(i) for i in resources]


class UrlMapPathMatcherRouteRuleMatchRule(object):
    def __init__(
        self,
        prefix_match: str = None,
        full_path_match: str = None,
        regex_match: str = None,
        ignore_case: bool = None,
        header_match: list = None,
        query_parameter_match: list = None,
        metadata_filter: list = None,
    ):
        self.prefix_match = prefix_match
        self.full_path_match = full_path_match
        self.regex_match = regex_match
        self.ignore_case = ignore_case
        self.header_match = header_match
        self.query_parameter_match = query_parameter_match
        self.metadata_filter = metadata_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherRouteRuleMatchRule()
        if Primitive.to_proto(resource.prefix_match):
            res.prefix_match = Primitive.to_proto(resource.prefix_match)
        if Primitive.to_proto(resource.full_path_match):
            res.full_path_match = Primitive.to_proto(resource.full_path_match)
        if Primitive.to_proto(resource.regex_match):
            res.regex_match = Primitive.to_proto(resource.regex_match)
        if Primitive.to_proto(resource.ignore_case):
            res.ignore_case = Primitive.to_proto(resource.ignore_case)
        if UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchArray.to_proto(
            resource.header_match
        ):
            res.header_match.extend(
                UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchArray.to_proto(
                    resource.header_match
                )
            )
        if UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArray.to_proto(
            resource.query_parameter_match
        ):
            res.query_parameter_match.extend(
                UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArray.to_proto(
                    resource.query_parameter_match
                )
            )
        if UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArray.to_proto(
            resource.metadata_filter
        ):
            res.metadata_filter.extend(
                UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArray.to_proto(
                    resource.metadata_filter
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleMatchRule(
            prefix_match=Primitive.from_proto(resource.prefix_match),
            full_path_match=Primitive.from_proto(resource.full_path_match),
            regex_match=Primitive.from_proto(resource.regex_match),
            ignore_case=Primitive.from_proto(resource.ignore_case),
            header_match=UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchArray.from_proto(
                resource.header_match
            ),
            query_parameter_match=UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArray.from_proto(
                resource.query_parameter_match
            ),
            metadata_filter=UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArray.from_proto(
                resource.metadata_filter
            ),
        )


class UrlMapPathMatcherRouteRuleMatchRuleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapPathMatcherRouteRuleMatchRule.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapPathMatcherRouteRuleMatchRule.from_proto(i) for i in resources]


class UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(object):
    def __init__(
        self,
        header_name: str = None,
        exact_match: str = None,
        regex_match: str = None,
        range_match: dict = None,
        present_match: bool = None,
        prefix_match: str = None,
        suffix_match: str = None,
        invert_match: bool = None,
    ):
        self.header_name = header_name
        self.exact_match = exact_match
        self.regex_match = regex_match
        self.range_match = range_match
        self.present_match = present_match
        self.prefix_match = prefix_match
        self.suffix_match = suffix_match
        self.invert_match = invert_match

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch()
        if Primitive.to_proto(resource.header_name):
            res.header_name = Primitive.to_proto(resource.header_name)
        if Primitive.to_proto(resource.exact_match):
            res.exact_match = Primitive.to_proto(resource.exact_match)
        if Primitive.to_proto(resource.regex_match):
            res.regex_match = Primitive.to_proto(resource.regex_match)
        if UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch.to_proto(
            resource.range_match
        ):
            res.range_match.CopyFrom(
                UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch.to_proto(
                    resource.range_match
                )
            )
        else:
            res.ClearField("range_match")
        if Primitive.to_proto(resource.present_match):
            res.present_match = Primitive.to_proto(resource.present_match)
        if Primitive.to_proto(resource.prefix_match):
            res.prefix_match = Primitive.to_proto(resource.prefix_match)
        if Primitive.to_proto(resource.suffix_match):
            res.suffix_match = Primitive.to_proto(resource.suffix_match)
        if Primitive.to_proto(resource.invert_match):
            res.invert_match = Primitive.to_proto(resource.invert_match)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(
            header_name=Primitive.from_proto(resource.header_name),
            exact_match=Primitive.from_proto(resource.exact_match),
            regex_match=Primitive.from_proto(resource.regex_match),
            range_match=UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch.from_proto(
                resource.range_match
            ),
            present_match=Primitive.from_proto(resource.present_match),
            prefix_match=Primitive.from_proto(resource.prefix_match),
            suffix_match=Primitive.from_proto(resource.suffix_match),
            invert_match=Primitive.from_proto(resource.invert_match),
        )


class UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(object):
    def __init__(self, range_start: int = None, range_end: int = None):
        self.range_start = range_start
        self.range_end = range_end

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch()
        )
        if Primitive.to_proto(resource.range_start):
            res.range_start = Primitive.to_proto(resource.range_start)
        if Primitive.to_proto(resource.range_end):
            res.range_end = Primitive.to_proto(resource.range_end)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(
            range_start=Primitive.from_proto(resource.range_start),
            range_end=Primitive.from_proto(resource.range_end),
        )


class UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(object):
    def __init__(
        self,
        name: str = None,
        present_match: bool = None,
        exact_match: str = None,
        regex_match: str = None,
    ):
        self.name = name
        self.present_match = present_match
        self.exact_match = exact_match
        self.regex_match = regex_match

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.present_match):
            res.present_match = Primitive.to_proto(resource.present_match)
        if Primitive.to_proto(resource.exact_match):
            res.exact_match = Primitive.to_proto(resource.exact_match)
        if Primitive.to_proto(resource.regex_match):
            res.regex_match = Primitive.to_proto(resource.regex_match)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(
            name=Primitive.from_proto(resource.name),
            present_match=Primitive.from_proto(resource.present_match),
            exact_match=Primitive.from_proto(resource.exact_match),
            regex_match=Primitive.from_proto(resource.regex_match),
        )


class UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(object):
    def __init__(self, filter_match_criteria: str = None, filter_label: list = None):
        self.filter_match_criteria = filter_match_criteria
        self.filter_label = filter_label

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter()
        if UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum.to_proto(
            resource.filter_match_criteria
        ):
            res.filter_match_criteria = UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum.to_proto(
                resource.filter_match_criteria
            )
        if UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArray.to_proto(
            resource.filter_label
        ):
            res.filter_label.extend(
                UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArray.to_proto(
                    resource.filter_label
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(
            filter_match_criteria=UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum.from_proto(
                resource.filter_match_criteria
            ),
            filter_label=UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArray.from_proto(
                resource.filter_label
            ),
        )


class UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(object):
    def __init__(self, name: str = None, value: str = None):
        self.name = name
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(
            name=Primitive.from_proto(resource.name),
            value=Primitive.from_proto(resource.value),
        )


class UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleRouteAction(object):
    def __init__(
        self,
        weighted_backend_service: list = None,
        url_rewrite: dict = None,
        timeout: dict = None,
        retry_policy: dict = None,
        request_mirror_policy: dict = None,
        cors_policy: dict = None,
        fault_injection_policy: dict = None,
    ):
        self.weighted_backend_service = weighted_backend_service
        self.url_rewrite = url_rewrite
        self.timeout = timeout
        self.retry_policy = retry_policy
        self.request_mirror_policy = request_mirror_policy
        self.cors_policy = cors_policy
        self.fault_injection_policy = fault_injection_policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteAction()
        if UrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceArray.to_proto(
            resource.weighted_backend_service
        ):
            res.weighted_backend_service.extend(
                UrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceArray.to_proto(
                    resource.weighted_backend_service
                )
            )
        if UrlMapPathMatcherRouteRuleRouteActionUrlRewrite.to_proto(
            resource.url_rewrite
        ):
            res.url_rewrite.CopyFrom(
                UrlMapPathMatcherRouteRuleRouteActionUrlRewrite.to_proto(
                    resource.url_rewrite
                )
            )
        else:
            res.ClearField("url_rewrite")
        if UrlMapPathMatcherRouteRuleRouteActionTimeout.to_proto(resource.timeout):
            res.timeout.CopyFrom(
                UrlMapPathMatcherRouteRuleRouteActionTimeout.to_proto(resource.timeout)
            )
        else:
            res.ClearField("timeout")
        if UrlMapPathMatcherRouteRuleRouteActionRetryPolicy.to_proto(
            resource.retry_policy
        ):
            res.retry_policy.CopyFrom(
                UrlMapPathMatcherRouteRuleRouteActionRetryPolicy.to_proto(
                    resource.retry_policy
                )
            )
        else:
            res.ClearField("retry_policy")
        if UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy.to_proto(
            resource.request_mirror_policy
        ):
            res.request_mirror_policy.CopyFrom(
                UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy.to_proto(
                    resource.request_mirror_policy
                )
            )
        else:
            res.ClearField("request_mirror_policy")
        if UrlMapPathMatcherRouteRuleRouteActionCorsPolicy.to_proto(
            resource.cors_policy
        ):
            res.cors_policy.CopyFrom(
                UrlMapPathMatcherRouteRuleRouteActionCorsPolicy.to_proto(
                    resource.cors_policy
                )
            )
        else:
            res.ClearField("cors_policy")
        if UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy.to_proto(
            resource.fault_injection_policy
        ):
            res.fault_injection_policy.CopyFrom(
                UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy.to_proto(
                    resource.fault_injection_policy
                )
            )
        else:
            res.ClearField("fault_injection_policy")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteAction(
            weighted_backend_service=UrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceArray.from_proto(
                resource.weighted_backend_service
            ),
            url_rewrite=UrlMapPathMatcherRouteRuleRouteActionUrlRewrite.from_proto(
                resource.url_rewrite
            ),
            timeout=UrlMapPathMatcherRouteRuleRouteActionTimeout.from_proto(
                resource.timeout
            ),
            retry_policy=UrlMapPathMatcherRouteRuleRouteActionRetryPolicy.from_proto(
                resource.retry_policy
            ),
            request_mirror_policy=UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy.from_proto(
                resource.request_mirror_policy
            ),
            cors_policy=UrlMapPathMatcherRouteRuleRouteActionCorsPolicy.from_proto(
                resource.cors_policy
            ),
            fault_injection_policy=UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy.from_proto(
                resource.fault_injection_policy
            ),
        )


class UrlMapPathMatcherRouteRuleRouteActionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapPathMatcherRouteRuleRouteAction.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapPathMatcherRouteRuleRouteAction.from_proto(i) for i in resources]


class UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(object):
    def __init__(
        self,
        backend_service: str = None,
        weight: int = None,
        header_action: dict = None,
    ):
        self.backend_service = backend_service
        self.weight = weight
        self.header_action = header_action

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService()
        )
        if Primitive.to_proto(resource.backend_service):
            res.backend_service = Primitive.to_proto(resource.backend_service)
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if UrlMapHeaderAction.to_proto(resource.header_action):
            res.header_action.CopyFrom(
                UrlMapHeaderAction.to_proto(resource.header_action)
            )
        else:
            res.ClearField("header_action")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(
            backend_service=Primitive.from_proto(resource.backend_service),
            weight=Primitive.from_proto(resource.weight),
            header_action=UrlMapHeaderAction.from_proto(resource.header_action),
        )


class UrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleRouteActionUrlRewrite(object):
    def __init__(self, path_prefix_rewrite: str = None, host_rewrite: str = None):
        self.path_prefix_rewrite = path_prefix_rewrite
        self.host_rewrite = host_rewrite

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteActionUrlRewrite()
        if Primitive.to_proto(resource.path_prefix_rewrite):
            res.path_prefix_rewrite = Primitive.to_proto(resource.path_prefix_rewrite)
        if Primitive.to_proto(resource.host_rewrite):
            res.host_rewrite = Primitive.to_proto(resource.host_rewrite)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteActionUrlRewrite(
            path_prefix_rewrite=Primitive.from_proto(resource.path_prefix_rewrite),
            host_rewrite=Primitive.from_proto(resource.host_rewrite),
        )


class UrlMapPathMatcherRouteRuleRouteActionUrlRewriteArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleRouteActionUrlRewrite.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleRouteActionUrlRewrite.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleRouteActionTimeout(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteActionTimeout()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteActionTimeout(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class UrlMapPathMatcherRouteRuleRouteActionTimeoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleRouteActionTimeout.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleRouteActionTimeout.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleRouteActionRetryPolicy(object):
    def __init__(
        self,
        retry_condition: list = None,
        num_retries: int = None,
        per_try_timeout: dict = None,
    ):
        self.retry_condition = retry_condition
        self.num_retries = num_retries
        self.per_try_timeout = per_try_timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicy()
        if Primitive.to_proto(resource.retry_condition):
            res.retry_condition.extend(Primitive.to_proto(resource.retry_condition))
        if Primitive.to_proto(resource.num_retries):
            res.num_retries = Primitive.to_proto(resource.num_retries)
        if UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout.to_proto(
            resource.per_try_timeout
        ):
            res.per_try_timeout.CopyFrom(
                UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout.to_proto(
                    resource.per_try_timeout
                )
            )
        else:
            res.ClearField("per_try_timeout")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteActionRetryPolicy(
            retry_condition=Primitive.from_proto(resource.retry_condition),
            num_retries=Primitive.from_proto(resource.num_retries),
            per_try_timeout=UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout.from_proto(
                resource.per_try_timeout
            ),
        )


class UrlMapPathMatcherRouteRuleRouteActionRetryPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleRouteActionRetryPolicy.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleRouteActionRetryPolicy.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout()
        )
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(object):
    def __init__(self, backend_service: str = None):
        self.backend_service = backend_service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy()
        )
        if Primitive.to_proto(resource.backend_service):
            res.backend_service = Primitive.to_proto(resource.backend_service)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(
            backend_service=Primitive.from_proto(resource.backend_service),
        )


class UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleRouteActionCorsPolicy(object):
    def __init__(
        self,
        allow_origin: list = None,
        allow_origin_regex: list = None,
        allow_method: list = None,
        allow_header: list = None,
        expose_header: list = None,
        max_age: int = None,
        allow_credentials: bool = None,
        disabled: bool = None,
    ):
        self.allow_origin = allow_origin
        self.allow_origin_regex = allow_origin_regex
        self.allow_method = allow_method
        self.allow_header = allow_header
        self.expose_header = expose_header
        self.max_age = max_age
        self.allow_credentials = allow_credentials
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteActionCorsPolicy()
        if Primitive.to_proto(resource.allow_origin):
            res.allow_origin.extend(Primitive.to_proto(resource.allow_origin))
        if Primitive.to_proto(resource.allow_origin_regex):
            res.allow_origin_regex.extend(
                Primitive.to_proto(resource.allow_origin_regex)
            )
        if Primitive.to_proto(resource.allow_method):
            res.allow_method.extend(Primitive.to_proto(resource.allow_method))
        if Primitive.to_proto(resource.allow_header):
            res.allow_header.extend(Primitive.to_proto(resource.allow_header))
        if Primitive.to_proto(resource.expose_header):
            res.expose_header.extend(Primitive.to_proto(resource.expose_header))
        if Primitive.to_proto(resource.max_age):
            res.max_age = Primitive.to_proto(resource.max_age)
        if Primitive.to_proto(resource.allow_credentials):
            res.allow_credentials = Primitive.to_proto(resource.allow_credentials)
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteActionCorsPolicy(
            allow_origin=Primitive.from_proto(resource.allow_origin),
            allow_origin_regex=Primitive.from_proto(resource.allow_origin_regex),
            allow_method=Primitive.from_proto(resource.allow_method),
            allow_header=Primitive.from_proto(resource.allow_header),
            expose_header=Primitive.from_proto(resource.expose_header),
            max_age=Primitive.from_proto(resource.max_age),
            allow_credentials=Primitive.from_proto(resource.allow_credentials),
            disabled=Primitive.from_proto(resource.disabled),
        )


class UrlMapPathMatcherRouteRuleRouteActionCorsPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleRouteActionCorsPolicy.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleRouteActionCorsPolicy.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(object):
    def __init__(self, delay: dict = None, abort: dict = None):
        self.delay = delay
        self.abort = abort

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy()
        )
        if UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay.to_proto(
            resource.delay
        ):
            res.delay.CopyFrom(
                UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay.to_proto(
                    resource.delay
                )
            )
        else:
            res.ClearField("delay")
        if UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort.to_proto(
            resource.abort
        ):
            res.abort.CopyFrom(
                UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort.to_proto(
                    resource.abort
                )
            )
        else:
            res.ClearField("abort")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(
            delay=UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay.from_proto(
                resource.delay
            ),
            abort=UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort.from_proto(
                resource.abort
            ),
        )


class UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(object):
    def __init__(self, fixed_delay: dict = None, percentage: float = None):
        self.fixed_delay = fixed_delay
        self.percentage = percentage

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay()
        )
        if UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay.to_proto(
            resource.fixed_delay
        ):
            res.fixed_delay.CopyFrom(
                UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay.to_proto(
                    resource.fixed_delay
                )
            )
        else:
            res.ClearField("fixed_delay")
        if Primitive.to_proto(resource.percentage):
            res.percentage = Primitive.to_proto(resource.percentage)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(
            fixed_delay=UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay.from_proto(
                resource.fixed_delay
            ),
            percentage=Primitive.from_proto(resource.percentage),
        )


class UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay()
        )
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay.from_proto(
                i
            )
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(object):
    def __init__(self, http_status: int = None, percentage: float = None):
        self.http_status = http_status
        self.percentage = percentage

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            url_map_pb2.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort()
        )
        if Primitive.to_proto(resource.http_status):
            res.http_status = Primitive.to_proto(resource.http_status)
        if Primitive.to_proto(resource.percentage):
            res.percentage = Primitive.to_proto(resource.percentage)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(
            http_status=Primitive.from_proto(resource.http_status),
            percentage=Primitive.from_proto(resource.percentage),
        )


class UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort.from_proto(i)
            for i in resources
        ]


class UrlMapPathMatcherRouteRuleUrlRedirect(object):
    def __init__(
        self,
        host_redirect: str = None,
        path_redirect: str = None,
        prefix_redirect: str = None,
        redirect_response_code: str = None,
        https_redirect: bool = None,
        strip_query: bool = None,
    ):
        self.host_redirect = host_redirect
        self.path_redirect = path_redirect
        self.prefix_redirect = prefix_redirect
        self.redirect_response_code = redirect_response_code
        self.https_redirect = https_redirect
        self.strip_query = strip_query

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapPathMatcherRouteRuleUrlRedirect()
        if Primitive.to_proto(resource.host_redirect):
            res.host_redirect = Primitive.to_proto(resource.host_redirect)
        if Primitive.to_proto(resource.path_redirect):
            res.path_redirect = Primitive.to_proto(resource.path_redirect)
        if Primitive.to_proto(resource.prefix_redirect):
            res.prefix_redirect = Primitive.to_proto(resource.prefix_redirect)
        if UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum.to_proto(
            resource.redirect_response_code
        ):
            res.redirect_response_code = UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum.to_proto(
                resource.redirect_response_code
            )
        if Primitive.to_proto(resource.https_redirect):
            res.https_redirect = Primitive.to_proto(resource.https_redirect)
        if Primitive.to_proto(resource.strip_query):
            res.strip_query = Primitive.to_proto(resource.strip_query)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapPathMatcherRouteRuleUrlRedirect(
            host_redirect=Primitive.from_proto(resource.host_redirect),
            path_redirect=Primitive.from_proto(resource.path_redirect),
            prefix_redirect=Primitive.from_proto(resource.prefix_redirect),
            redirect_response_code=UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum.from_proto(
                resource.redirect_response_code
            ),
            https_redirect=Primitive.from_proto(resource.https_redirect),
            strip_query=Primitive.from_proto(resource.strip_query),
        )


class UrlMapPathMatcherRouteRuleUrlRedirectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapPathMatcherRouteRuleUrlRedirect.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapPathMatcherRouteRuleUrlRedirect.from_proto(i) for i in resources]


class UrlMapTest(object):
    def __init__(
        self,
        description: str = None,
        host: str = None,
        path: str = None,
        expected_backend_service: str = None,
    ):
        self.description = description
        self.host = host
        self.path = path
        self.expected_backend_service = expected_backend_service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = url_map_pb2.ComputeUrlMapTest()
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.expected_backend_service):
            res.expected_backend_service = Primitive.to_proto(
                resource.expected_backend_service
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UrlMapTest(
            description=Primitive.from_proto(resource.description),
            host=Primitive.from_proto(resource.host),
            path=Primitive.from_proto(resource.path),
            expected_backend_service=Primitive.from_proto(
                resource.expected_backend_service
            ),
        )


class UrlMapTestArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UrlMapTest.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UrlMapTest.from_proto(i) for i in resources]


class UrlMapDefaultUrlRedirectRedirectResponseCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return url_map_pb2.ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum.Value(
            "ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return url_map_pb2.ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum.Name(
            resource
        )[len("ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum") :]


class UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return url_map_pb2.ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum.Value(
            "ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return url_map_pb2.ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum.Name(
            resource
        )[
            len("ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum") :
        ]


class UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return url_map_pb2.ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum.Value(
            "ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return url_map_pb2.ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum.Name(
            resource
        )[
            len("ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum") :
        ]


class UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return url_map_pb2.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum.Value(
            "ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return url_map_pb2.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum.Name(
            resource
        )[
            len(
                "ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum"
            ) :
        ]


class UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return url_map_pb2.ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum.Value(
            "ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return url_map_pb2.ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum.Name(
            resource
        )[
            len(
                "ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum"
            ) :
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
