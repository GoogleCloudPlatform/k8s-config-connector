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
from google3.cloud.graphite.mmv2.services.google.network_services import tcp_route_pb2
from google3.cloud.graphite.mmv2.services.google.network_services import (
    tcp_route_pb2_grpc,
)

from typing import List


class TcpRoute(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        description: str = None,
        rules: list = None,
        meshes: list = None,
        gateways: list = None,
        labels: dict = None,
        project: str = None,
        location: str = None,
        self_link: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.rules = rules
        self.meshes = meshes
        self.gateways = gateways
        self.labels = labels
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = tcp_route_pb2_grpc.NetworkservicesAlphaTcpRouteServiceStub(
            channel.Channel()
        )
        request = tcp_route_pb2.ApplyNetworkservicesAlphaTcpRouteRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if TcpRouteRulesArray.to_proto(self.rules):
            request.resource.rules.extend(TcpRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.meshes):
            request.resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            request.resource.gateways.extend(Primitive.to_proto(self.gateways))
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworkservicesAlphaTcpRoute(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.description = Primitive.from_proto(response.description)
        self.rules = TcpRouteRulesArray.from_proto(response.rules)
        self.meshes = Primitive.from_proto(response.meshes)
        self.gateways = Primitive.from_proto(response.gateways)
        self.labels = Primitive.from_proto(response.labels)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.self_link = Primitive.from_proto(response.self_link)

    def delete(self):
        stub = tcp_route_pb2_grpc.NetworkservicesAlphaTcpRouteServiceStub(
            channel.Channel()
        )
        request = tcp_route_pb2.DeleteNetworkservicesAlphaTcpRouteRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if TcpRouteRulesArray.to_proto(self.rules):
            request.resource.rules.extend(TcpRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.meshes):
            request.resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            request.resource.gateways.extend(Primitive.to_proto(self.gateways))
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworkservicesAlphaTcpRoute(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = tcp_route_pb2_grpc.NetworkservicesAlphaTcpRouteServiceStub(
            channel.Channel()
        )
        request = tcp_route_pb2.ListNetworkservicesAlphaTcpRouteRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListNetworkservicesAlphaTcpRoute(request).items

    def to_proto(self):
        resource = tcp_route_pb2.NetworkservicesAlphaTcpRoute()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if TcpRouteRulesArray.to_proto(self.rules):
            resource.rules.extend(TcpRouteRulesArray.to_proto(self.rules))
        if Primitive.to_proto(self.meshes):
            resource.meshes.extend(Primitive.to_proto(self.meshes))
        if Primitive.to_proto(self.gateways):
            resource.gateways.extend(Primitive.to_proto(self.gateways))
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class TcpRouteRules(object):
    def __init__(self, matches: list = None, action: dict = None):
        self.matches = matches
        self.action = action

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = tcp_route_pb2.NetworkservicesAlphaTcpRouteRules()
        if TcpRouteRulesMatchesArray.to_proto(resource.matches):
            res.matches.extend(TcpRouteRulesMatchesArray.to_proto(resource.matches))
        if TcpRouteRulesAction.to_proto(resource.action):
            res.action.CopyFrom(TcpRouteRulesAction.to_proto(resource.action))
        else:
            res.ClearField("action")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TcpRouteRules(
            matches=TcpRouteRulesMatchesArray.from_proto(resource.matches),
            action=TcpRouteRulesAction.from_proto(resource.action),
        )


class TcpRouteRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TcpRouteRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TcpRouteRules.from_proto(i) for i in resources]


class TcpRouteRulesMatches(object):
    def __init__(self, address: str = None, port: str = None):
        self.address = address
        self.port = port

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = tcp_route_pb2.NetworkservicesAlphaTcpRouteRulesMatches()
        if Primitive.to_proto(resource.address):
            res.address = Primitive.to_proto(resource.address)
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TcpRouteRulesMatches(
            address=Primitive.from_proto(resource.address),
            port=Primitive.from_proto(resource.port),
        )


class TcpRouteRulesMatchesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TcpRouteRulesMatches.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TcpRouteRulesMatches.from_proto(i) for i in resources]


class TcpRouteRulesAction(object):
    def __init__(self, destinations: list = None, original_destination: bool = None):
        self.destinations = destinations
        self.original_destination = original_destination

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = tcp_route_pb2.NetworkservicesAlphaTcpRouteRulesAction()
        if TcpRouteRulesActionDestinationsArray.to_proto(resource.destinations):
            res.destinations.extend(
                TcpRouteRulesActionDestinationsArray.to_proto(resource.destinations)
            )
        if Primitive.to_proto(resource.original_destination):
            res.original_destination = Primitive.to_proto(resource.original_destination)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TcpRouteRulesAction(
            destinations=TcpRouteRulesActionDestinationsArray.from_proto(
                resource.destinations
            ),
            original_destination=Primitive.from_proto(resource.original_destination),
        )


class TcpRouteRulesActionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TcpRouteRulesAction.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TcpRouteRulesAction.from_proto(i) for i in resources]


class TcpRouteRulesActionDestinations(object):
    def __init__(self, weight: int = None, service_name: str = None):
        self.weight = weight
        self.service_name = service_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = tcp_route_pb2.NetworkservicesAlphaTcpRouteRulesActionDestinations()
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if Primitive.to_proto(resource.service_name):
            res.service_name = Primitive.to_proto(resource.service_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TcpRouteRulesActionDestinations(
            weight=Primitive.from_proto(resource.weight),
            service_name=Primitive.from_proto(resource.service_name),
        )


class TcpRouteRulesActionDestinationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TcpRouteRulesActionDestinations.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TcpRouteRulesActionDestinations.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
