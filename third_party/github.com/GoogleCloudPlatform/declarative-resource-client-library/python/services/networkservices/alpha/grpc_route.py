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
from google3.cloud.graphite.mmv2.services.google.network_services import grpc_route_pb2
from google3.cloud.graphite.mmv2.services.google.network_services import (
    grpc_route_pb2_grpc,
)

from typing import List


class GrpcRoute(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        description: str = None,
        hostnames: list = None,
        meshes: list = None,
        gateways: list = None,
        rules: list = None,
        project: str = None,
        location: str = None,
        self_link: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.labels = labels
        self.description = description
        self.hostnames = hostnames
        self.meshes = meshes
        self.gateways = gateways
        self.rules = rules
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = grpc_route_pb2_grpc.NetworkservicesAlphaGrpcRouteServiceStub(
            channel.Channel()
        )
        request = grpc_route_pb2.ApplyNetworkservicesAlphaGrpcRouteRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.hostnames):
            request.resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.meshes):
            request.resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            request.resource.gateways.extend(Primitive.to_proto(self.gateways))
        if GrpcRouteRulesArray.to_proto(self.rules):
            request.resource.rules.extend(GrpcRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworkservicesAlphaGrpcRoute(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.description = Primitive.from_proto(response.description)
        self.hostnames = Primitive.from_proto(response.hostnames)
        self.meshes = Primitive.from_proto(response.meshes)
        self.gateways = Primitive.from_proto(response.gateways)
        self.rules = GrpcRouteRulesArray.from_proto(response.rules)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.self_link = Primitive.from_proto(response.self_link)

    def delete(self):
        stub = grpc_route_pb2_grpc.NetworkservicesAlphaGrpcRouteServiceStub(
            channel.Channel()
        )
        request = grpc_route_pb2.DeleteNetworkservicesAlphaGrpcRouteRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.hostnames):
            request.resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.meshes):
            request.resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            request.resource.gateways.extend(Primitive.to_proto(self.gateways))
        if GrpcRouteRulesArray.to_proto(self.rules):
            request.resource.rules.extend(GrpcRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworkservicesAlphaGrpcRoute(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = grpc_route_pb2_grpc.NetworkservicesAlphaGrpcRouteServiceStub(
            channel.Channel()
        )
        request = grpc_route_pb2.ListNetworkservicesAlphaGrpcRouteRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworkservicesAlphaGrpcRoute(request).items

    def to_proto(self):
        resource = grpc_route_pb2.NetworkservicesAlphaGrpcRoute()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.hostnames):
            resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.meshes):
            resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            resource.gateways.extend(Primitive.to_proto(self.gateways))
        if GrpcRouteRulesArray.to_proto(self.rules):
            resource.rules.extend(GrpcRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class GrpcRouteRules(object):
    def __init__(self, matches: list = None, action: dict = None):
        self.matches = matches
        self.action = action

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = grpc_route_pb2.NetworkservicesAlphaGrpcRouteRules()
        if GrpcRouteRulesMatchesArray.to_proto(resource.matches):
            res.matches.extend(GrpcRouteRulesMatchesArray.to_proto(resource.matches))
        if GrpcRouteRulesAction.to_proto(resource.action):
            res.action.CopyFrom(GrpcRouteRulesAction.to_proto(resource.action))
        else:
            res.ClearField("action")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GrpcRouteRules(
            matches=GrpcRouteRulesMatchesArray.from_proto(resource.matches),
            action=GrpcRouteRulesAction.from_proto(resource.action),
        )


class GrpcRouteRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GrpcRouteRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GrpcRouteRules.from_proto(i) for i in resources]


class GrpcRouteRulesMatches(object):
    def __init__(self, method: dict = None, headers: list = None):
        self.method = method
        self.headers = headers

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesMatches()
        if GrpcRouteRulesMatchesMethod.to_proto(resource.method):
            res.method.CopyFrom(GrpcRouteRulesMatchesMethod.to_proto(resource.method))
        else:
            res.ClearField("method")
        if GrpcRouteRulesMatchesHeadersArray.to_proto(resource.headers):
            res.headers.extend(
                GrpcRouteRulesMatchesHeadersArray.to_proto(resource.headers)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GrpcRouteRulesMatches(
            method=GrpcRouteRulesMatchesMethod.from_proto(resource.method),
            headers=GrpcRouteRulesMatchesHeadersArray.from_proto(resource.headers),
        )


class GrpcRouteRulesMatchesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GrpcRouteRulesMatches.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GrpcRouteRulesMatches.from_proto(i) for i in resources]


class GrpcRouteRulesMatchesMethod(object):
    def __init__(
        self,
        type: str = None,
        grpc_service: str = None,
        grpc_method: str = None,
        case_sensitive: bool = None,
    ):
        self.type = type
        self.grpc_service = grpc_service
        self.grpc_method = grpc_method
        self.case_sensitive = case_sensitive

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesMatchesMethod()
        if GrpcRouteRulesMatchesMethodTypeEnum.to_proto(resource.type):
            res.type = GrpcRouteRulesMatchesMethodTypeEnum.to_proto(resource.type)
        if Primitive.to_proto(resource.grpc_service):
            res.grpc_service = Primitive.to_proto(resource.grpc_service)
        if Primitive.to_proto(resource.grpc_method):
            res.grpc_method = Primitive.to_proto(resource.grpc_method)
        if Primitive.to_proto(resource.case_sensitive):
            res.case_sensitive = Primitive.to_proto(resource.case_sensitive)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GrpcRouteRulesMatchesMethod(
            type=GrpcRouteRulesMatchesMethodTypeEnum.from_proto(resource.type),
            grpc_service=Primitive.from_proto(resource.grpc_service),
            grpc_method=Primitive.from_proto(resource.grpc_method),
            case_sensitive=Primitive.from_proto(resource.case_sensitive),
        )


class GrpcRouteRulesMatchesMethodArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GrpcRouteRulesMatchesMethod.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GrpcRouteRulesMatchesMethod.from_proto(i) for i in resources]


class GrpcRouteRulesMatchesHeaders(object):
    def __init__(self, type: str = None, key: str = None, value: str = None):
        self.type = type
        self.key = key
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesMatchesHeaders()
        if GrpcRouteRulesMatchesHeadersTypeEnum.to_proto(resource.type):
            res.type = GrpcRouteRulesMatchesHeadersTypeEnum.to_proto(resource.type)
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GrpcRouteRulesMatchesHeaders(
            type=GrpcRouteRulesMatchesHeadersTypeEnum.from_proto(resource.type),
            key=Primitive.from_proto(resource.key),
            value=Primitive.from_proto(resource.value),
        )


class GrpcRouteRulesMatchesHeadersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GrpcRouteRulesMatchesHeaders.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GrpcRouteRulesMatchesHeaders.from_proto(i) for i in resources]


class GrpcRouteRulesAction(object):
    def __init__(
        self,
        destinations: list = None,
        fault_injection_policy: dict = None,
        timeout: str = None,
        retry_policy: dict = None,
    ):
        self.destinations = destinations
        self.fault_injection_policy = fault_injection_policy
        self.timeout = timeout
        self.retry_policy = retry_policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesAction()
        if GrpcRouteRulesActionDestinationsArray.to_proto(resource.destinations):
            res.destinations.extend(
                GrpcRouteRulesActionDestinationsArray.to_proto(resource.destinations)
            )
        if GrpcRouteRulesActionFaultInjectionPolicy.to_proto(
            resource.fault_injection_policy
        ):
            res.fault_injection_policy.CopyFrom(
                GrpcRouteRulesActionFaultInjectionPolicy.to_proto(
                    resource.fault_injection_policy
                )
            )
        else:
            res.ClearField("fault_injection_policy")
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        if GrpcRouteRulesActionRetryPolicy.to_proto(resource.retry_policy):
            res.retry_policy.CopyFrom(
                GrpcRouteRulesActionRetryPolicy.to_proto(resource.retry_policy)
            )
        else:
            res.ClearField("retry_policy")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GrpcRouteRulesAction(
            destinations=GrpcRouteRulesActionDestinationsArray.from_proto(
                resource.destinations
            ),
            fault_injection_policy=GrpcRouteRulesActionFaultInjectionPolicy.from_proto(
                resource.fault_injection_policy
            ),
            timeout=Primitive.from_proto(resource.timeout),
            retry_policy=GrpcRouteRulesActionRetryPolicy.from_proto(
                resource.retry_policy
            ),
        )


class GrpcRouteRulesActionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GrpcRouteRulesAction.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GrpcRouteRulesAction.from_proto(i) for i in resources]


class GrpcRouteRulesActionDestinations(object):
    def __init__(self, weight: int = None, service_name: str = None):
        self.weight = weight
        self.service_name = service_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesActionDestinations()
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if Primitive.to_proto(resource.service_name):
            res.service_name = Primitive.to_proto(resource.service_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GrpcRouteRulesActionDestinations(
            weight=Primitive.from_proto(resource.weight),
            service_name=Primitive.from_proto(resource.service_name),
        )


class GrpcRouteRulesActionDestinationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GrpcRouteRulesActionDestinations.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GrpcRouteRulesActionDestinations.from_proto(i) for i in resources]


class GrpcRouteRulesActionFaultInjectionPolicy(object):
    def __init__(self, delay: dict = None, abort: dict = None):
        self.delay = delay
        self.abort = abort

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicy()
        )
        if GrpcRouteRulesActionFaultInjectionPolicyDelay.to_proto(resource.delay):
            res.delay.CopyFrom(
                GrpcRouteRulesActionFaultInjectionPolicyDelay.to_proto(resource.delay)
            )
        else:
            res.ClearField("delay")
        if GrpcRouteRulesActionFaultInjectionPolicyAbort.to_proto(resource.abort):
            res.abort.CopyFrom(
                GrpcRouteRulesActionFaultInjectionPolicyAbort.to_proto(resource.abort)
            )
        else:
            res.ClearField("abort")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GrpcRouteRulesActionFaultInjectionPolicy(
            delay=GrpcRouteRulesActionFaultInjectionPolicyDelay.from_proto(
                resource.delay
            ),
            abort=GrpcRouteRulesActionFaultInjectionPolicyAbort.from_proto(
                resource.abort
            ),
        )


class GrpcRouteRulesActionFaultInjectionPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GrpcRouteRulesActionFaultInjectionPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            GrpcRouteRulesActionFaultInjectionPolicy.from_proto(i) for i in resources
        ]


class GrpcRouteRulesActionFaultInjectionPolicyDelay(object):
    def __init__(self, fixed_delay: str = None, percentage: int = None):
        self.fixed_delay = fixed_delay
        self.percentage = percentage

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyDelay()
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

        return GrpcRouteRulesActionFaultInjectionPolicyDelay(
            fixed_delay=Primitive.from_proto(resource.fixed_delay),
            percentage=Primitive.from_proto(resource.percentage),
        )


class GrpcRouteRulesActionFaultInjectionPolicyDelayArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            GrpcRouteRulesActionFaultInjectionPolicyDelay.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            GrpcRouteRulesActionFaultInjectionPolicyDelay.from_proto(i)
            for i in resources
        ]


class GrpcRouteRulesActionFaultInjectionPolicyAbort(object):
    def __init__(self, http_status: int = None, percentage: int = None):
        self.http_status = http_status
        self.percentage = percentage

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyAbort()
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

        return GrpcRouteRulesActionFaultInjectionPolicyAbort(
            http_status=Primitive.from_proto(resource.http_status),
            percentage=Primitive.from_proto(resource.percentage),
        )


class GrpcRouteRulesActionFaultInjectionPolicyAbortArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            GrpcRouteRulesActionFaultInjectionPolicyAbort.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            GrpcRouteRulesActionFaultInjectionPolicyAbort.from_proto(i)
            for i in resources
        ]


class GrpcRouteRulesActionRetryPolicy(object):
    def __init__(self, retry_conditions: list = None, num_retries: int = None):
        self.retry_conditions = retry_conditions
        self.num_retries = num_retries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesActionRetryPolicy()
        if Primitive.to_proto(resource.retry_conditions):
            res.retry_conditions.extend(Primitive.to_proto(resource.retry_conditions))
        if Primitive.to_proto(resource.num_retries):
            res.num_retries = Primitive.to_proto(resource.num_retries)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GrpcRouteRulesActionRetryPolicy(
            retry_conditions=Primitive.from_proto(resource.retry_conditions),
            num_retries=Primitive.from_proto(resource.num_retries),
        )


class GrpcRouteRulesActionRetryPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GrpcRouteRulesActionRetryPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GrpcRouteRulesActionRetryPolicy.from_proto(i) for i in resources]


class GrpcRouteRulesMatchesMethodTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum.Value(
            "NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum.Name(
                resource
            )[len("NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum") :]
        )


class GrpcRouteRulesMatchesHeadersTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum.Value(
            "NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return grpc_route_pb2.NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum.Name(
            resource
        )[
            len("NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum") :
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
