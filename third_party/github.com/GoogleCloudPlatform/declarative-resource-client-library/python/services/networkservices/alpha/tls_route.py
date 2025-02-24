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
from google3.cloud.graphite.mmv2.services.google.network_services import tls_route_pb2
from google3.cloud.graphite.mmv2.services.google.network_services import (
    tls_route_pb2_grpc,
)

from typing import List


class TlsRoute(object):
    def __init__(
        self,
        name: str = None,
        self_link: str = None,
        create_time: str = None,
        update_time: str = None,
        description: str = None,
        rules: list = None,
        meshes: list = None,
        gateways: list = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.rules = rules
        self.meshes = meshes
        self.gateways = gateways
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = tls_route_pb2_grpc.NetworkservicesAlphaTlsRouteServiceStub(
            channel.Channel()
        )
        request = tls_route_pb2.ApplyNetworkservicesAlphaTlsRouteRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if TlsRouteRulesArray.to_proto(self.rules):
            request.resource.rules.extend(TlsRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.meshes):
            request.resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            request.resource.gateways.extend(Primitive.to_proto(self.gateways))
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworkservicesAlphaTlsRoute(request)
        self.name = Primitive.from_proto(response.name)
        self.self_link = Primitive.from_proto(response.self_link)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.description = Primitive.from_proto(response.description)
        self.rules = TlsRouteRulesArray.from_proto(response.rules)
        self.meshes = Primitive.from_proto(response.meshes)
        self.gateways = Primitive.from_proto(response.gateways)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = tls_route_pb2_grpc.NetworkservicesAlphaTlsRouteServiceStub(
            channel.Channel()
        )
        request = tls_route_pb2.DeleteNetworkservicesAlphaTlsRouteRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if TlsRouteRulesArray.to_proto(self.rules):
            request.resource.rules.extend(TlsRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.meshes):
            request.resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            request.resource.gateways.extend(Primitive.to_proto(self.gateways))
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworkservicesAlphaTlsRoute(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = tls_route_pb2_grpc.NetworkservicesAlphaTlsRouteServiceStub(
            channel.Channel()
        )
        request = tls_route_pb2.ListNetworkservicesAlphaTlsRouteRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworkservicesAlphaTlsRoute(request).items

    def to_proto(self):
        resource = tls_route_pb2.NetworkservicesAlphaTlsRoute()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if TlsRouteRulesArray.to_proto(self.rules):
            resource.rules.extend(TlsRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.meshes):
            resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            resource.gateways.extend(Primitive.to_proto(self.gateways))
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class TlsRouteRules(object):
    def __init__(self, matches: list = None, action: dict = None):
        self.matches = matches
        self.action = action

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = tls_route_pb2.NetworkservicesAlphaTlsRouteRules()
        if TlsRouteRulesMatchesArray.to_proto(resource.matches):
            res.matches.extend(TlsRouteRulesMatchesArray.to_proto(resource.matches))
        if TlsRouteRulesAction.to_proto(resource.action):
            res.action.CopyFrom(TlsRouteRulesAction.to_proto(resource.action))
        else:
            res.ClearField("action")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TlsRouteRules(
            matches=TlsRouteRulesMatchesArray.from_proto(resource.matches),
            action=TlsRouteRulesAction.from_proto(resource.action),
        )


class TlsRouteRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TlsRouteRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TlsRouteRules.from_proto(i) for i in resources]


class TlsRouteRulesMatches(object):
    def __init__(self, sni_host: list = None, alpn: list = None):
        self.sni_host = sni_host
        self.alpn = alpn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = tls_route_pb2.NetworkservicesAlphaTlsRouteRulesMatches()
        if Primitive.to_proto(resource.sni_host):
            res.sni_host.extend(Primitive.to_proto(resource.sni_host))
        if Primitive.to_proto(resource.alpn):
            res.alpn.extend(Primitive.to_proto(resource.alpn))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TlsRouteRulesMatches(
            sni_host=Primitive.from_proto(resource.sni_host),
            alpn=Primitive.from_proto(resource.alpn),
        )


class TlsRouteRulesMatchesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TlsRouteRulesMatches.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TlsRouteRulesMatches.from_proto(i) for i in resources]


class TlsRouteRulesAction(object):
    def __init__(self, destinations: list = None):
        self.destinations = destinations

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = tls_route_pb2.NetworkservicesAlphaTlsRouteRulesAction()
        if TlsRouteRulesActionDestinationsArray.to_proto(resource.destinations):
            res.destinations.extend(
                TlsRouteRulesActionDestinationsArray.to_proto(resource.destinations)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TlsRouteRulesAction(
            destinations=TlsRouteRulesActionDestinationsArray.from_proto(
                resource.destinations
            ),
        )


class TlsRouteRulesActionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TlsRouteRulesAction.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TlsRouteRulesAction.from_proto(i) for i in resources]


class TlsRouteRulesActionDestinations(object):
    def __init__(self, service_name: str = None, weight: int = None):
        self.service_name = service_name
        self.weight = weight

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = tls_route_pb2.NetworkservicesAlphaTlsRouteRulesActionDestinations()
        if Primitive.to_proto(resource.service_name):
            res.service_name = Primitive.to_proto(resource.service_name)
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TlsRouteRulesActionDestinations(
            service_name=Primitive.from_proto(resource.service_name),
            weight=Primitive.from_proto(resource.weight),
        )


class TlsRouteRulesActionDestinationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TlsRouteRulesActionDestinations.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TlsRouteRulesActionDestinations.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
