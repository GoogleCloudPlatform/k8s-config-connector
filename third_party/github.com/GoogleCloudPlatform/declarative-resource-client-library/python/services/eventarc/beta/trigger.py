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
from google3.cloud.graphite.mmv2.services.google.eventarc import trigger_pb2
from google3.cloud.graphite.mmv2.services.google.eventarc import trigger_pb2_grpc

from typing import List


class Trigger(object):
    def __init__(
        self,
        name: str = None,
        uid: str = None,
        create_time: str = None,
        update_time: str = None,
        matching_criteria: list = None,
        service_account: str = None,
        destination: dict = None,
        transport: dict = None,
        labels: dict = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        channel: str = None,
        conditions: dict = None,
        event_data_content_type: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.matching_criteria = matching_criteria
        self.service_account = service_account
        self.destination = destination
        self.transport = transport
        self.labels = labels
        self.project = project
        self.location = location
        self.channel = channel
        self.event_data_content_type = event_data_content_type
        self.service_account_file = service_account_file

    def apply(self):
        stub = trigger_pb2_grpc.EventarcBetaTriggerServiceStub(channel.Channel())
        request = trigger_pb2.ApplyEventarcBetaTriggerRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if TriggerMatchingCriteriaArray.to_proto(self.matching_criteria):
            request.resource.matching_criteria.extend(
                TriggerMatchingCriteriaArray.to_proto(self.matching_criteria)
            )
        if Primitive.to_proto(self.service_account):
            request.resource.service_account = Primitive.to_proto(self.service_account)

        if TriggerDestination.to_proto(self.destination):
            request.resource.destination.CopyFrom(
                TriggerDestination.to_proto(self.destination)
            )
        else:
            request.resource.ClearField("destination")
        if TriggerTransport.to_proto(self.transport):
            request.resource.transport.CopyFrom(
                TriggerTransport.to_proto(self.transport)
            )
        else:
            request.resource.ClearField("transport")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.channel):
            request.resource.channel = Primitive.to_proto(self.channel)

        if Primitive.to_proto(self.event_data_content_type):
            request.resource.event_data_content_type = Primitive.to_proto(
                self.event_data_content_type
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyEventarcBetaTrigger(request)
        self.name = Primitive.from_proto(response.name)
        self.uid = Primitive.from_proto(response.uid)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.matching_criteria = TriggerMatchingCriteriaArray.from_proto(
            response.matching_criteria
        )
        self.service_account = Primitive.from_proto(response.service_account)
        self.destination = TriggerDestination.from_proto(response.destination)
        self.transport = TriggerTransport.from_proto(response.transport)
        self.labels = Primitive.from_proto(response.labels)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.channel = Primitive.from_proto(response.channel)
        self.conditions = Primitive.from_proto(response.conditions)
        self.event_data_content_type = Primitive.from_proto(
            response.event_data_content_type
        )

    def delete(self):
        stub = trigger_pb2_grpc.EventarcBetaTriggerServiceStub(channel.Channel())
        request = trigger_pb2.DeleteEventarcBetaTriggerRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if TriggerMatchingCriteriaArray.to_proto(self.matching_criteria):
            request.resource.matching_criteria.extend(
                TriggerMatchingCriteriaArray.to_proto(self.matching_criteria)
            )
        if Primitive.to_proto(self.service_account):
            request.resource.service_account = Primitive.to_proto(self.service_account)

        if TriggerDestination.to_proto(self.destination):
            request.resource.destination.CopyFrom(
                TriggerDestination.to_proto(self.destination)
            )
        else:
            request.resource.ClearField("destination")
        if TriggerTransport.to_proto(self.transport):
            request.resource.transport.CopyFrom(
                TriggerTransport.to_proto(self.transport)
            )
        else:
            request.resource.ClearField("transport")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.channel):
            request.resource.channel = Primitive.to_proto(self.channel)

        if Primitive.to_proto(self.event_data_content_type):
            request.resource.event_data_content_type = Primitive.to_proto(
                self.event_data_content_type
            )

        response = stub.DeleteEventarcBetaTrigger(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = trigger_pb2_grpc.EventarcBetaTriggerServiceStub(channel.Channel())
        request = trigger_pb2.ListEventarcBetaTriggerRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListEventarcBetaTrigger(request).items

    def to_proto(self):
        resource = trigger_pb2.EventarcBetaTrigger()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if TriggerMatchingCriteriaArray.to_proto(self.matching_criteria):
            resource.matching_criteria.extend(
                TriggerMatchingCriteriaArray.to_proto(self.matching_criteria)
            )
        if Primitive.to_proto(self.service_account):
            resource.service_account = Primitive.to_proto(self.service_account)
        if TriggerDestination.to_proto(self.destination):
            resource.destination.CopyFrom(TriggerDestination.to_proto(self.destination))
        else:
            resource.ClearField("destination")
        if TriggerTransport.to_proto(self.transport):
            resource.transport.CopyFrom(TriggerTransport.to_proto(self.transport))
        else:
            resource.ClearField("transport")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.channel):
            resource.channel = Primitive.to_proto(self.channel)
        if Primitive.to_proto(self.event_data_content_type):
            resource.event_data_content_type = Primitive.to_proto(
                self.event_data_content_type
            )
        return resource


class TriggerMatchingCriteria(object):
    def __init__(self, attribute: str = None, value: str = None, operator: str = None):
        self.attribute = attribute
        self.value = value
        self.operator = operator

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcBetaTriggerMatchingCriteria()
        if Primitive.to_proto(resource.attribute):
            res.attribute = Primitive.to_proto(resource.attribute)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if Primitive.to_proto(resource.operator):
            res.operator = Primitive.to_proto(resource.operator)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerMatchingCriteria(
            attribute=Primitive.from_proto(resource.attribute),
            value=Primitive.from_proto(resource.value),
            operator=Primitive.from_proto(resource.operator),
        )


class TriggerMatchingCriteriaArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerMatchingCriteria.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerMatchingCriteria.from_proto(i) for i in resources]


class TriggerDestination(object):
    def __init__(
        self,
        cloud_run_service: dict = None,
        cloud_function: str = None,
        gke: dict = None,
        workflow: str = None,
        http_endpoint: dict = None,
        network_config: dict = None,
    ):
        self.cloud_run_service = cloud_run_service
        self.cloud_function = cloud_function
        self.gke = gke
        self.workflow = workflow
        self.http_endpoint = http_endpoint
        self.network_config = network_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcBetaTriggerDestination()
        if TriggerDestinationCloudRunService.to_proto(resource.cloud_run_service):
            res.cloud_run_service.CopyFrom(
                TriggerDestinationCloudRunService.to_proto(resource.cloud_run_service)
            )
        else:
            res.ClearField("cloud_run_service")
        if Primitive.to_proto(resource.cloud_function):
            res.cloud_function = Primitive.to_proto(resource.cloud_function)
        if TriggerDestinationGke.to_proto(resource.gke):
            res.gke.CopyFrom(TriggerDestinationGke.to_proto(resource.gke))
        else:
            res.ClearField("gke")
        if Primitive.to_proto(resource.workflow):
            res.workflow = Primitive.to_proto(resource.workflow)
        if TriggerDestinationHttpEndpoint.to_proto(resource.http_endpoint):
            res.http_endpoint.CopyFrom(
                TriggerDestinationHttpEndpoint.to_proto(resource.http_endpoint)
            )
        else:
            res.ClearField("http_endpoint")
        if TriggerDestinationNetworkConfig.to_proto(resource.network_config):
            res.network_config.CopyFrom(
                TriggerDestinationNetworkConfig.to_proto(resource.network_config)
            )
        else:
            res.ClearField("network_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerDestination(
            cloud_run_service=TriggerDestinationCloudRunService.from_proto(
                resource.cloud_run_service
            ),
            cloud_function=Primitive.from_proto(resource.cloud_function),
            gke=TriggerDestinationGke.from_proto(resource.gke),
            workflow=Primitive.from_proto(resource.workflow),
            http_endpoint=TriggerDestinationHttpEndpoint.from_proto(
                resource.http_endpoint
            ),
            network_config=TriggerDestinationNetworkConfig.from_proto(
                resource.network_config
            ),
        )


class TriggerDestinationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerDestination.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerDestination.from_proto(i) for i in resources]


class TriggerDestinationCloudRunService(object):
    def __init__(self, service: str = None, path: str = None, region: str = None):
        self.service = service
        self.path = path
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcBetaTriggerDestinationCloudRunService()
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerDestinationCloudRunService(
            service=Primitive.from_proto(resource.service),
            path=Primitive.from_proto(resource.path),
            region=Primitive.from_proto(resource.region),
        )


class TriggerDestinationCloudRunServiceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerDestinationCloudRunService.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerDestinationCloudRunService.from_proto(i) for i in resources]


class TriggerDestinationGke(object):
    def __init__(
        self,
        cluster: str = None,
        location: str = None,
        namespace: str = None,
        service: str = None,
        path: str = None,
    ):
        self.cluster = cluster
        self.location = location
        self.namespace = namespace
        self.service = service
        self.path = path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcBetaTriggerDestinationGke()
        if Primitive.to_proto(resource.cluster):
            res.cluster = Primitive.to_proto(resource.cluster)
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        if Primitive.to_proto(resource.namespace):
            res.namespace = Primitive.to_proto(resource.namespace)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerDestinationGke(
            cluster=Primitive.from_proto(resource.cluster),
            location=Primitive.from_proto(resource.location),
            namespace=Primitive.from_proto(resource.namespace),
            service=Primitive.from_proto(resource.service),
            path=Primitive.from_proto(resource.path),
        )


class TriggerDestinationGkeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerDestinationGke.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerDestinationGke.from_proto(i) for i in resources]


class TriggerDestinationHttpEndpoint(object):
    def __init__(self, uri: str = None):
        self.uri = uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcBetaTriggerDestinationHttpEndpoint()
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerDestinationHttpEndpoint(
            uri=Primitive.from_proto(resource.uri),
        )


class TriggerDestinationHttpEndpointArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerDestinationHttpEndpoint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerDestinationHttpEndpoint.from_proto(i) for i in resources]


class TriggerDestinationNetworkConfig(object):
    def __init__(self, network_attachment: str = None):
        self.network_attachment = network_attachment

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcBetaTriggerDestinationNetworkConfig()
        if Primitive.to_proto(resource.network_attachment):
            res.network_attachment = Primitive.to_proto(resource.network_attachment)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerDestinationNetworkConfig(
            network_attachment=Primitive.from_proto(resource.network_attachment),
        )


class TriggerDestinationNetworkConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerDestinationNetworkConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerDestinationNetworkConfig.from_proto(i) for i in resources]


class TriggerTransport(object):
    def __init__(self, pubsub: dict = None):
        self.pubsub = pubsub

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcBetaTriggerTransport()
        if TriggerTransportPubsub.to_proto(resource.pubsub):
            res.pubsub.CopyFrom(TriggerTransportPubsub.to_proto(resource.pubsub))
        else:
            res.ClearField("pubsub")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerTransport(
            pubsub=TriggerTransportPubsub.from_proto(resource.pubsub),
        )


class TriggerTransportArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerTransport.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerTransport.from_proto(i) for i in resources]


class TriggerTransportPubsub(object):
    def __init__(self, topic: str = None, subscription: str = None):
        self.topic = topic
        self.subscription = subscription

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcBetaTriggerTransportPubsub()
        if Primitive.to_proto(resource.topic):
            res.topic = Primitive.to_proto(resource.topic)
        if Primitive.to_proto(resource.subscription):
            res.subscription = Primitive.to_proto(resource.subscription)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerTransportPubsub(
            topic=Primitive.from_proto(resource.topic),
            subscription=Primitive.from_proto(resource.subscription),
        )


class TriggerTransportPubsubArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerTransportPubsub.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerTransportPubsub.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
