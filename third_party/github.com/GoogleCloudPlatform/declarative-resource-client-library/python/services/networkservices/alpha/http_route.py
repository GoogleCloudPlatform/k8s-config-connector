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
from google3.cloud.graphite.mmv2.services.google.network_services import http_route_pb2
from google3.cloud.graphite.mmv2.services.google.network_services import (
    http_route_pb2_grpc,
)

from typing import List


class HttpRoute(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        hostnames: list = None,
        meshes: list = None,
        gateways: list = None,
        labels: dict = None,
        rules: list = None,
        project: str = None,
        location: str = None,
        self_link: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.hostnames = hostnames
        self.meshes = meshes
        self.gateways = gateways
        self.labels = labels
        self.rules = rules
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = http_route_pb2_grpc.NetworkservicesAlphaHttpRouteServiceStub(
            channel.Channel()
        )
        request = http_route_pb2.ApplyNetworkservicesAlphaHttpRouteRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.hostnames):
            request.resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.meshes):
            request.resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            request.resource.gateways.extend(Primitive.to_proto(self.gateways))
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if HttpRouteRulesArray.to_proto(self.rules):
            request.resource.rules.extend(HttpRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworkservicesAlphaHttpRoute(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.hostnames = Primitive.from_proto(response.hostnames)
        self.meshes = Primitive.from_proto(response.meshes)
        self.gateways = Primitive.from_proto(response.gateways)
        self.labels = Primitive.from_proto(response.labels)
        self.rules = HttpRouteRulesArray.from_proto(response.rules)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.self_link = Primitive.from_proto(response.self_link)

    def delete(self):
        stub = http_route_pb2_grpc.NetworkservicesAlphaHttpRouteServiceStub(
            channel.Channel()
        )
        request = http_route_pb2.DeleteNetworkservicesAlphaHttpRouteRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.hostnames):
            request.resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.meshes):
            request.resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            request.resource.gateways.extend(Primitive.to_proto(self.gateways))
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if HttpRouteRulesArray.to_proto(self.rules):
            request.resource.rules.extend(HttpRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworkservicesAlphaHttpRoute(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = http_route_pb2_grpc.NetworkservicesAlphaHttpRouteServiceStub(
            channel.Channel()
        )
        request = http_route_pb2.ListNetworkservicesAlphaHttpRouteRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworkservicesAlphaHttpRoute(request).items

    def to_proto(self):
        resource = http_route_pb2.NetworkservicesAlphaHttpRoute()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.hostnames):
            resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.meshes):
            resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            resource.gateways.extend(Primitive.to_proto(self.gateways))
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if HttpRouteRulesArray.to_proto(self.rules):
            resource.rules.extend(HttpRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class HttpRouteRules(object):
    def __init__(self, matches: list = None, action: dict = None):
        self.matches = matches
        self.action = action

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = http_route_pb2.NetworkservicesAlphaHttpRouteRules()
        if HttpRouteRulesMatchesArray.to_proto(resource.matches):
            res.matches.extend(HttpRouteRulesMatchesArray.to_proto(resource.matches))
        if HttpRouteRulesAction.to_proto(resource.action):
            res.action.CopyFrom(HttpRouteRulesAction.to_proto(resource.action))
        else:
            res.ClearField("action")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRules(
            matches=HttpRouteRulesMatchesArray.from_proto(resource.matches),
            action=HttpRouteRulesAction.from_proto(resource.action),
        )


class HttpRouteRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HttpRouteRules.from_proto(i) for i in resources]


class HttpRouteRulesMatches(object):
    def __init__(
        self,
        full_path_match: str = None,
        prefix_match: str = None,
        regex_match: str = None,
        ignore_case: bool = None,
        headers: list = None,
        query_parameters: list = None,
    ):
        self.full_path_match = full_path_match
        self.prefix_match = prefix_match
        self.regex_match = regex_match
        self.ignore_case = ignore_case
        self.headers = headers
        self.query_parameters = query_parameters

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = http_route_pb2.NetworkservicesAlphaHttpRouteRulesMatches()
        if Primitive.to_proto(resource.full_path_match):
            res.full_path_match = Primitive.to_proto(resource.full_path_match)
        if Primitive.to_proto(resource.prefix_match):
            res.prefix_match = Primitive.to_proto(resource.prefix_match)
        if Primitive.to_proto(resource.regex_match):
            res.regex_match = Primitive.to_proto(resource.regex_match)
        if Primitive.to_proto(resource.ignore_case):
            res.ignore_case = Primitive.to_proto(resource.ignore_case)
        if HttpRouteRulesMatchesHeadersArray.to_proto(resource.headers):
            res.headers.extend(
                HttpRouteRulesMatchesHeadersArray.to_proto(resource.headers)
            )
        if HttpRouteRulesMatchesQueryParametersArray.to_proto(
            resource.query_parameters
        ):
            res.query_parameters.extend(
                HttpRouteRulesMatchesQueryParametersArray.to_proto(
                    resource.query_parameters
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesMatches(
            full_path_match=Primitive.from_proto(resource.full_path_match),
            prefix_match=Primitive.from_proto(resource.prefix_match),
            regex_match=Primitive.from_proto(resource.regex_match),
            ignore_case=Primitive.from_proto(resource.ignore_case),
            headers=HttpRouteRulesMatchesHeadersArray.from_proto(resource.headers),
            query_parameters=HttpRouteRulesMatchesQueryParametersArray.from_proto(
                resource.query_parameters
            ),
        )


class HttpRouteRulesMatchesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesMatches.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HttpRouteRulesMatches.from_proto(i) for i in resources]


class HttpRouteRulesMatchesHeaders(object):
    def __init__(
        self,
        header: str = None,
        exact_match: str = None,
        regex_match: str = None,
        prefix_match: str = None,
        present_match: bool = None,
        suffix_match: str = None,
        range_match: dict = None,
        invert_match: bool = None,
    ):
        self.header = header
        self.exact_match = exact_match
        self.regex_match = regex_match
        self.prefix_match = prefix_match
        self.present_match = present_match
        self.suffix_match = suffix_match
        self.range_match = range_match
        self.invert_match = invert_match

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = http_route_pb2.NetworkservicesAlphaHttpRouteRulesMatchesHeaders()
        if Primitive.to_proto(resource.header):
            res.header = Primitive.to_proto(resource.header)
        if Primitive.to_proto(resource.exact_match):
            res.exact_match = Primitive.to_proto(resource.exact_match)
        if Primitive.to_proto(resource.regex_match):
            res.regex_match = Primitive.to_proto(resource.regex_match)
        if Primitive.to_proto(resource.prefix_match):
            res.prefix_match = Primitive.to_proto(resource.prefix_match)
        if Primitive.to_proto(resource.present_match):
            res.present_match = Primitive.to_proto(resource.present_match)
        if Primitive.to_proto(resource.suffix_match):
            res.suffix_match = Primitive.to_proto(resource.suffix_match)
        if HttpRouteRulesMatchesHeadersRangeMatch.to_proto(resource.range_match):
            res.range_match.CopyFrom(
                HttpRouteRulesMatchesHeadersRangeMatch.to_proto(resource.range_match)
            )
        else:
            res.ClearField("range_match")
        if Primitive.to_proto(resource.invert_match):
            res.invert_match = Primitive.to_proto(resource.invert_match)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesMatchesHeaders(
            header=Primitive.from_proto(resource.header),
            exact_match=Primitive.from_proto(resource.exact_match),
            regex_match=Primitive.from_proto(resource.regex_match),
            prefix_match=Primitive.from_proto(resource.prefix_match),
            present_match=Primitive.from_proto(resource.present_match),
            suffix_match=Primitive.from_proto(resource.suffix_match),
            range_match=HttpRouteRulesMatchesHeadersRangeMatch.from_proto(
                resource.range_match
            ),
            invert_match=Primitive.from_proto(resource.invert_match),
        )


class HttpRouteRulesMatchesHeadersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesMatchesHeaders.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HttpRouteRulesMatchesHeaders.from_proto(i) for i in resources]


class HttpRouteRulesMatchesHeadersRangeMatch(object):
    def __init__(self, start: int = None, end: int = None):
        self.start = start
        self.end = end

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            http_route_pb2.NetworkservicesAlphaHttpRouteRulesMatchesHeadersRangeMatch()
        )
        if Primitive.to_proto(resource.start):
            res.start = Primitive.to_proto(resource.start)
        if Primitive.to_proto(resource.end):
            res.end = Primitive.to_proto(resource.end)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesMatchesHeadersRangeMatch(
            start=Primitive.from_proto(resource.start),
            end=Primitive.from_proto(resource.end),
        )


class HttpRouteRulesMatchesHeadersRangeMatchArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesMatchesHeadersRangeMatch.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HttpRouteRulesMatchesHeadersRangeMatch.from_proto(i) for i in resources]


class HttpRouteRulesMatchesQueryParameters(object):
    def __init__(
        self,
        query_parameter: str = None,
        exact_match: str = None,
        regex_match: str = None,
        present_match: bool = None,
    ):
        self.query_parameter = query_parameter
        self.exact_match = exact_match
        self.regex_match = regex_match
        self.present_match = present_match

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = http_route_pb2.NetworkservicesAlphaHttpRouteRulesMatchesQueryParameters()
        if Primitive.to_proto(resource.query_parameter):
            res.query_parameter = Primitive.to_proto(resource.query_parameter)
        if Primitive.to_proto(resource.exact_match):
            res.exact_match = Primitive.to_proto(resource.exact_match)
        if Primitive.to_proto(resource.regex_match):
            res.regex_match = Primitive.to_proto(resource.regex_match)
        if Primitive.to_proto(resource.present_match):
            res.present_match = Primitive.to_proto(resource.present_match)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesMatchesQueryParameters(
            query_parameter=Primitive.from_proto(resource.query_parameter),
            exact_match=Primitive.from_proto(resource.exact_match),
            regex_match=Primitive.from_proto(resource.regex_match),
            present_match=Primitive.from_proto(resource.present_match),
        )


class HttpRouteRulesMatchesQueryParametersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesMatchesQueryParameters.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HttpRouteRulesMatchesQueryParameters.from_proto(i) for i in resources]


class HttpRouteRulesAction(object):
    def __init__(
        self,
        destinations: list = None,
        redirect: dict = None,
        fault_injection_policy: dict = None,
        request_header_modifier: dict = None,
        response_header_modifier: dict = None,
        url_rewrite: dict = None,
        timeout: str = None,
        retry_policy: dict = None,
        request_mirror_policy: dict = None,
        cors_policy: dict = None,
    ):
        self.destinations = destinations
        self.redirect = redirect
        self.fault_injection_policy = fault_injection_policy
        self.request_header_modifier = request_header_modifier
        self.response_header_modifier = response_header_modifier
        self.url_rewrite = url_rewrite
        self.timeout = timeout
        self.retry_policy = retry_policy
        self.request_mirror_policy = request_mirror_policy
        self.cors_policy = cors_policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = http_route_pb2.NetworkservicesAlphaHttpRouteRulesAction()
        if HttpRouteRulesActionDestinationsArray.to_proto(resource.destinations):
            res.destinations.extend(
                HttpRouteRulesActionDestinationsArray.to_proto(resource.destinations)
            )
        if HttpRouteRulesActionRedirect.to_proto(resource.redirect):
            res.redirect.CopyFrom(
                HttpRouteRulesActionRedirect.to_proto(resource.redirect)
            )
        else:
            res.ClearField("redirect")
        if HttpRouteRulesActionFaultInjectionPolicy.to_proto(
            resource.fault_injection_policy
        ):
            res.fault_injection_policy.CopyFrom(
                HttpRouteRulesActionFaultInjectionPolicy.to_proto(
                    resource.fault_injection_policy
                )
            )
        else:
            res.ClearField("fault_injection_policy")
        if HttpRouteRulesActionRequestHeaderModifier.to_proto(
            resource.request_header_modifier
        ):
            res.request_header_modifier.CopyFrom(
                HttpRouteRulesActionRequestHeaderModifier.to_proto(
                    resource.request_header_modifier
                )
            )
        else:
            res.ClearField("request_header_modifier")
        if HttpRouteRulesActionResponseHeaderModifier.to_proto(
            resource.response_header_modifier
        ):
            res.response_header_modifier.CopyFrom(
                HttpRouteRulesActionResponseHeaderModifier.to_proto(
                    resource.response_header_modifier
                )
            )
        else:
            res.ClearField("response_header_modifier")
        if HttpRouteRulesActionUrlRewrite.to_proto(resource.url_rewrite):
            res.url_rewrite.CopyFrom(
                HttpRouteRulesActionUrlRewrite.to_proto(resource.url_rewrite)
            )
        else:
            res.ClearField("url_rewrite")
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        if HttpRouteRulesActionRetryPolicy.to_proto(resource.retry_policy):
            res.retry_policy.CopyFrom(
                HttpRouteRulesActionRetryPolicy.to_proto(resource.retry_policy)
            )
        else:
            res.ClearField("retry_policy")
        if HttpRouteRulesActionRequestMirrorPolicy.to_proto(
            resource.request_mirror_policy
        ):
            res.request_mirror_policy.CopyFrom(
                HttpRouteRulesActionRequestMirrorPolicy.to_proto(
                    resource.request_mirror_policy
                )
            )
        else:
            res.ClearField("request_mirror_policy")
        if HttpRouteRulesActionCorsPolicy.to_proto(resource.cors_policy):
            res.cors_policy.CopyFrom(
                HttpRouteRulesActionCorsPolicy.to_proto(resource.cors_policy)
            )
        else:
            res.ClearField("cors_policy")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesAction(
            destinations=HttpRouteRulesActionDestinationsArray.from_proto(
                resource.destinations
            ),
            redirect=HttpRouteRulesActionRedirect.from_proto(resource.redirect),
            fault_injection_policy=HttpRouteRulesActionFaultInjectionPolicy.from_proto(
                resource.fault_injection_policy
            ),
            request_header_modifier=HttpRouteRulesActionRequestHeaderModifier.from_proto(
                resource.request_header_modifier
            ),
            response_header_modifier=HttpRouteRulesActionResponseHeaderModifier.from_proto(
                resource.response_header_modifier
            ),
            url_rewrite=HttpRouteRulesActionUrlRewrite.from_proto(resource.url_rewrite),
            timeout=Primitive.from_proto(resource.timeout),
            retry_policy=HttpRouteRulesActionRetryPolicy.from_proto(
                resource.retry_policy
            ),
            request_mirror_policy=HttpRouteRulesActionRequestMirrorPolicy.from_proto(
                resource.request_mirror_policy
            ),
            cors_policy=HttpRouteRulesActionCorsPolicy.from_proto(resource.cors_policy),
        )


class HttpRouteRulesActionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesAction.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HttpRouteRulesAction.from_proto(i) for i in resources]


class HttpRouteRulesActionDestinations(object):
    def __init__(self, weight: int = None, service_name: str = None):
        self.weight = weight
        self.service_name = service_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionDestinations()
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if Primitive.to_proto(resource.service_name):
            res.service_name = Primitive.to_proto(resource.service_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesActionDestinations(
            weight=Primitive.from_proto(resource.weight),
            service_name=Primitive.from_proto(resource.service_name),
        )


class HttpRouteRulesActionDestinationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesActionDestinations.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HttpRouteRulesActionDestinations.from_proto(i) for i in resources]


class HttpRouteRulesActionRedirect(object):
    def __init__(
        self,
        host_redirect: str = None,
        path_redirect: str = None,
        prefix_rewrite: str = None,
        response_code: str = None,
        https_redirect: bool = None,
        strip_query: bool = None,
        port_redirect: int = None,
    ):
        self.host_redirect = host_redirect
        self.path_redirect = path_redirect
        self.prefix_rewrite = prefix_rewrite
        self.response_code = response_code
        self.https_redirect = https_redirect
        self.strip_query = strip_query
        self.port_redirect = port_redirect

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionRedirect()
        if Primitive.to_proto(resource.host_redirect):
            res.host_redirect = Primitive.to_proto(resource.host_redirect)
        if Primitive.to_proto(resource.path_redirect):
            res.path_redirect = Primitive.to_proto(resource.path_redirect)
        if Primitive.to_proto(resource.prefix_rewrite):
            res.prefix_rewrite = Primitive.to_proto(resource.prefix_rewrite)
        if HttpRouteRulesActionRedirectResponseCodeEnum.to_proto(
            resource.response_code
        ):
            res.response_code = HttpRouteRulesActionRedirectResponseCodeEnum.to_proto(
                resource.response_code
            )
        if Primitive.to_proto(resource.https_redirect):
            res.https_redirect = Primitive.to_proto(resource.https_redirect)
        if Primitive.to_proto(resource.strip_query):
            res.strip_query = Primitive.to_proto(resource.strip_query)
        if Primitive.to_proto(resource.port_redirect):
            res.port_redirect = Primitive.to_proto(resource.port_redirect)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesActionRedirect(
            host_redirect=Primitive.from_proto(resource.host_redirect),
            path_redirect=Primitive.from_proto(resource.path_redirect),
            prefix_rewrite=Primitive.from_proto(resource.prefix_rewrite),
            response_code=HttpRouteRulesActionRedirectResponseCodeEnum.from_proto(
                resource.response_code
            ),
            https_redirect=Primitive.from_proto(resource.https_redirect),
            strip_query=Primitive.from_proto(resource.strip_query),
            port_redirect=Primitive.from_proto(resource.port_redirect),
        )


class HttpRouteRulesActionRedirectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesActionRedirect.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HttpRouteRulesActionRedirect.from_proto(i) for i in resources]


class HttpRouteRulesActionFaultInjectionPolicy(object):
    def __init__(self, delay: dict = None, abort: dict = None):
        self.delay = delay
        self.abort = abort

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicy()
        )
        if HttpRouteRulesActionFaultInjectionPolicyDelay.to_proto(resource.delay):
            res.delay.CopyFrom(
                HttpRouteRulesActionFaultInjectionPolicyDelay.to_proto(resource.delay)
            )
        else:
            res.ClearField("delay")
        if HttpRouteRulesActionFaultInjectionPolicyAbort.to_proto(resource.abort):
            res.abort.CopyFrom(
                HttpRouteRulesActionFaultInjectionPolicyAbort.to_proto(resource.abort)
            )
        else:
            res.ClearField("abort")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesActionFaultInjectionPolicy(
            delay=HttpRouteRulesActionFaultInjectionPolicyDelay.from_proto(
                resource.delay
            ),
            abort=HttpRouteRulesActionFaultInjectionPolicyAbort.from_proto(
                resource.abort
            ),
        )


class HttpRouteRulesActionFaultInjectionPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesActionFaultInjectionPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            HttpRouteRulesActionFaultInjectionPolicy.from_proto(i) for i in resources
        ]


class HttpRouteRulesActionFaultInjectionPolicyDelay(object):
    def __init__(self, fixed_delay: str = None, percentage: int = None):
        self.fixed_delay = fixed_delay
        self.percentage = percentage

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyDelay()
        )
        if Primitive.to_proto(resource.fixed_delay):
            res.fixed_delay = Primitive.to_proto(resource.fixed_delay)
        if Primitive.to_proto(resource.percentage):
            res.percentage = Primitive.to_proto(resource.percentage)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesActionFaultInjectionPolicyDelay(
            fixed_delay=Primitive.from_proto(resource.fixed_delay),
            percentage=Primitive.from_proto(resource.percentage),
        )


class HttpRouteRulesActionFaultInjectionPolicyDelayArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            HttpRouteRulesActionFaultInjectionPolicyDelay.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            HttpRouteRulesActionFaultInjectionPolicyDelay.from_proto(i)
            for i in resources
        ]


class HttpRouteRulesActionFaultInjectionPolicyAbort(object):
    def __init__(self, http_status: int = None, percentage: int = None):
        self.http_status = http_status
        self.percentage = percentage

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyAbort()
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

        return HttpRouteRulesActionFaultInjectionPolicyAbort(
            http_status=Primitive.from_proto(resource.http_status),
            percentage=Primitive.from_proto(resource.percentage),
        )


class HttpRouteRulesActionFaultInjectionPolicyAbortArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            HttpRouteRulesActionFaultInjectionPolicyAbort.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            HttpRouteRulesActionFaultInjectionPolicyAbort.from_proto(i)
            for i in resources
        ]


class HttpRouteRulesActionRequestHeaderModifier(object):
    def __init__(self, set: dict = None, add: dict = None, remove: list = None):
        self.set = set
        self.add = add
        self.remove = remove

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionRequestHeaderModifier()
        )
        if Primitive.to_proto(resource.set):
            res.set = Primitive.to_proto(resource.set)
        if Primitive.to_proto(resource.add):
            res.add = Primitive.to_proto(resource.add)
        if Primitive.to_proto(resource.remove):
            res.remove.extend(Primitive.to_proto(resource.remove))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesActionRequestHeaderModifier(
            set=Primitive.from_proto(resource.set),
            add=Primitive.from_proto(resource.add),
            remove=Primitive.from_proto(resource.remove),
        )


class HttpRouteRulesActionRequestHeaderModifierArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            HttpRouteRulesActionRequestHeaderModifier.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            HttpRouteRulesActionRequestHeaderModifier.from_proto(i) for i in resources
        ]


class HttpRouteRulesActionResponseHeaderModifier(object):
    def __init__(self, set: dict = None, add: dict = None, remove: list = None):
        self.set = set
        self.add = add
        self.remove = remove

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionResponseHeaderModifier()
        )
        if Primitive.to_proto(resource.set):
            res.set = Primitive.to_proto(resource.set)
        if Primitive.to_proto(resource.add):
            res.add = Primitive.to_proto(resource.add)
        if Primitive.to_proto(resource.remove):
            res.remove.extend(Primitive.to_proto(resource.remove))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesActionResponseHeaderModifier(
            set=Primitive.from_proto(resource.set),
            add=Primitive.from_proto(resource.add),
            remove=Primitive.from_proto(resource.remove),
        )


class HttpRouteRulesActionResponseHeaderModifierArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            HttpRouteRulesActionResponseHeaderModifier.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            HttpRouteRulesActionResponseHeaderModifier.from_proto(i) for i in resources
        ]


class HttpRouteRulesActionUrlRewrite(object):
    def __init__(self, path_prefix_rewrite: str = None, host_rewrite: str = None):
        self.path_prefix_rewrite = path_prefix_rewrite
        self.host_rewrite = host_rewrite

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionUrlRewrite()
        if Primitive.to_proto(resource.path_prefix_rewrite):
            res.path_prefix_rewrite = Primitive.to_proto(resource.path_prefix_rewrite)
        if Primitive.to_proto(resource.host_rewrite):
            res.host_rewrite = Primitive.to_proto(resource.host_rewrite)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesActionUrlRewrite(
            path_prefix_rewrite=Primitive.from_proto(resource.path_prefix_rewrite),
            host_rewrite=Primitive.from_proto(resource.host_rewrite),
        )


class HttpRouteRulesActionUrlRewriteArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesActionUrlRewrite.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HttpRouteRulesActionUrlRewrite.from_proto(i) for i in resources]


class HttpRouteRulesActionRetryPolicy(object):
    def __init__(
        self,
        retry_conditions: list = None,
        num_retries: int = None,
        per_try_timeout: str = None,
    ):
        self.retry_conditions = retry_conditions
        self.num_retries = num_retries
        self.per_try_timeout = per_try_timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionRetryPolicy()
        if Primitive.to_proto(resource.retry_conditions):
            res.retry_conditions.extend(Primitive.to_proto(resource.retry_conditions))
        if Primitive.to_proto(resource.num_retries):
            res.num_retries = Primitive.to_proto(resource.num_retries)
        if Primitive.to_proto(resource.per_try_timeout):
            res.per_try_timeout = Primitive.to_proto(resource.per_try_timeout)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesActionRetryPolicy(
            retry_conditions=Primitive.from_proto(resource.retry_conditions),
            num_retries=Primitive.from_proto(resource.num_retries),
            per_try_timeout=Primitive.from_proto(resource.per_try_timeout),
        )


class HttpRouteRulesActionRetryPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesActionRetryPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HttpRouteRulesActionRetryPolicy.from_proto(i) for i in resources]


class HttpRouteRulesActionRequestMirrorPolicy(object):
    def __init__(self, destination: dict = None):
        self.destination = destination

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicy()
        )
        if HttpRouteRulesActionRequestMirrorPolicyDestination.to_proto(
            resource.destination
        ):
            res.destination.CopyFrom(
                HttpRouteRulesActionRequestMirrorPolicyDestination.to_proto(
                    resource.destination
                )
            )
        else:
            res.ClearField("destination")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesActionRequestMirrorPolicy(
            destination=HttpRouteRulesActionRequestMirrorPolicyDestination.from_proto(
                resource.destination
            ),
        )


class HttpRouteRulesActionRequestMirrorPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesActionRequestMirrorPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            HttpRouteRulesActionRequestMirrorPolicy.from_proto(i) for i in resources
        ]


class HttpRouteRulesActionRequestMirrorPolicyDestination(object):
    def __init__(self, weight: int = None, service_name: str = None):
        self.weight = weight
        self.service_name = service_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicyDestination()
        )
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if Primitive.to_proto(resource.service_name):
            res.service_name = Primitive.to_proto(resource.service_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HttpRouteRulesActionRequestMirrorPolicyDestination(
            weight=Primitive.from_proto(resource.weight),
            service_name=Primitive.from_proto(resource.service_name),
        )


class HttpRouteRulesActionRequestMirrorPolicyDestinationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            HttpRouteRulesActionRequestMirrorPolicyDestination.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            HttpRouteRulesActionRequestMirrorPolicyDestination.from_proto(i)
            for i in resources
        ]


class HttpRouteRulesActionCorsPolicy(object):
    def __init__(
        self,
        allow_origins: list = None,
        allow_origin_regexes: list = None,
        allow_methods: list = None,
        allow_headers: list = None,
        expose_headers: list = None,
        max_age: str = None,
        allow_credentials: bool = None,
        disabled: bool = None,
    ):
        self.allow_origins = allow_origins
        self.allow_origin_regexes = allow_origin_regexes
        self.allow_methods = allow_methods
        self.allow_headers = allow_headers
        self.expose_headers = expose_headers
        self.max_age = max_age
        self.allow_credentials = allow_credentials
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionCorsPolicy()
        if Primitive.to_proto(resource.allow_origins):
            res.allow_origins.extend(Primitive.to_proto(resource.allow_origins))
        if Primitive.to_proto(resource.allow_origin_regexes):
            res.allow_origin_regexes.extend(
                Primitive.to_proto(resource.allow_origin_regexes)
            )
        if Primitive.to_proto(resource.allow_methods):
            res.allow_methods.extend(Primitive.to_proto(resource.allow_methods))
        if Primitive.to_proto(resource.allow_headers):
            res.allow_headers.extend(Primitive.to_proto(resource.allow_headers))
        if Primitive.to_proto(resource.expose_headers):
            res.expose_headers.extend(Primitive.to_proto(resource.expose_headers))
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

        return HttpRouteRulesActionCorsPolicy(
            allow_origins=Primitive.from_proto(resource.allow_origins),
            allow_origin_regexes=Primitive.from_proto(resource.allow_origin_regexes),
            allow_methods=Primitive.from_proto(resource.allow_methods),
            allow_headers=Primitive.from_proto(resource.allow_headers),
            expose_headers=Primitive.from_proto(resource.expose_headers),
            max_age=Primitive.from_proto(resource.max_age),
            allow_credentials=Primitive.from_proto(resource.allow_credentials),
            disabled=Primitive.from_proto(resource.disabled),
        )


class HttpRouteRulesActionCorsPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HttpRouteRulesActionCorsPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HttpRouteRulesActionCorsPolicy.from_proto(i) for i in resources]


class HttpRouteRulesActionRedirectResponseCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum.Value(
            "NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return http_route_pb2.NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum.Name(
            resource
        )[
            len("NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum") :
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
