# Copyright 2023 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.gke_hub import feature_pb2
from google3.cloud.graphite.mmv2.services.google.gke_hub import feature_pb2_grpc

from typing import List


class Feature(object):
    def __init__(
        self,
        name: str = None,
        labels: dict = None,
        resource_state: dict = None,
        spec: dict = None,
        state: dict = None,
        create_time: str = None,
        update_time: str = None,
        delete_time: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.labels = labels
        self.spec = spec
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = feature_pb2_grpc.GkehubFeatureServiceStub(channel.Channel())
        request = feature_pb2.ApplyGkehubFeatureRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if FeatureSpec.to_proto(self.spec):
            request.resource.spec.CopyFrom(FeatureSpec.to_proto(self.spec))
        else:
            request.resource.ClearField("spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyGkehubFeature(request)
        self.name = Primitive.from_proto(response.name)
        self.labels = Primitive.from_proto(response.labels)
        self.resource_state = FeatureResourceState.from_proto(response.resource_state)
        self.spec = FeatureSpec.from_proto(response.spec)
        self.state = FeatureState.from_proto(response.state)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = feature_pb2_grpc.GkehubFeatureServiceStub(channel.Channel())
        request = feature_pb2.DeleteGkehubFeatureRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if FeatureSpec.to_proto(self.spec):
            request.resource.spec.CopyFrom(FeatureSpec.to_proto(self.spec))
        else:
            request.resource.ClearField("spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteGkehubFeature(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = feature_pb2_grpc.GkehubFeatureServiceStub(channel.Channel())
        request = feature_pb2.ListGkehubFeatureRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListGkehubFeature(request).items

    def to_proto(self):
        resource = feature_pb2.GkehubFeature()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if FeatureSpec.to_proto(self.spec):
            resource.spec.CopyFrom(FeatureSpec.to_proto(self.spec))
        else:
            resource.ClearField("spec")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class FeatureResourceState(object):
    def __init__(self, state: str = None, has_resources: bool = None):
        self.state = state
        self.has_resources = has_resources

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubFeatureResourceState()
        if FeatureResourceStateStateEnum.to_proto(resource.state):
            res.state = FeatureResourceStateStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.has_resources):
            res.has_resources = Primitive.to_proto(resource.has_resources)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureResourceState(
            state=FeatureResourceStateStateEnum.from_proto(resource.state),
            has_resources=Primitive.from_proto(resource.has_resources),
        )


class FeatureResourceStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureResourceState.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureResourceState.from_proto(i) for i in resources]


class FeatureSpec(object):
    def __init__(self, multiclusteringress: dict = None):
        self.multiclusteringress = multiclusteringress

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubFeatureSpec()
        if FeatureSpecMulticlusteringress.to_proto(resource.multiclusteringress):
            res.multiclusteringress.CopyFrom(
                FeatureSpecMulticlusteringress.to_proto(resource.multiclusteringress)
            )
        else:
            res.ClearField("multiclusteringress")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureSpec(
            multiclusteringress=FeatureSpecMulticlusteringress.from_proto(
                resource.multiclusteringress
            ),
        )


class FeatureSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureSpec.from_proto(i) for i in resources]


class FeatureSpecMulticlusteringress(object):
    def __init__(self, config_membership: str = None):
        self.config_membership = config_membership

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubFeatureSpecMulticlusteringress()
        if Primitive.to_proto(resource.config_membership):
            res.config_membership = Primitive.to_proto(resource.config_membership)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureSpecMulticlusteringress(
            config_membership=Primitive.from_proto(resource.config_membership),
        )


class FeatureSpecMulticlusteringressArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureSpecMulticlusteringress.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureSpecMulticlusteringress.from_proto(i) for i in resources]


class FeatureState(object):
    def __init__(self, state: dict = None):
        self.state = state

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubFeatureState()
        if FeatureStateState.to_proto(resource.state):
            res.state.CopyFrom(FeatureStateState.to_proto(resource.state))
        else:
            res.ClearField("state")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureState(
            state=FeatureStateState.from_proto(resource.state),
        )


class FeatureStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureState.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureState.from_proto(i) for i in resources]


class FeatureStateState(object):
    def __init__(
        self, code: str = None, description: str = None, update_time: str = None
    ):
        self.code = code
        self.description = description
        self.update_time = update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubFeatureStateState()
        if FeatureStateStateCodeEnum.to_proto(resource.code):
            res.code = FeatureStateStateCodeEnum.to_proto(resource.code)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureStateState(
            code=FeatureStateStateCodeEnum.from_proto(resource.code),
            description=Primitive.from_proto(resource.description),
            update_time=Primitive.from_proto(resource.update_time),
        )


class FeatureStateStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureStateState.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureStateState.from_proto(i) for i in resources]


class FeatureResourceStateStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubFeatureResourceStateStateEnum.Value(
            "GkehubFeatureResourceStateStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubFeatureResourceStateStateEnum.Name(resource)[
            len("GkehubFeatureResourceStateStateEnum") :
        ]


class FeatureStateStateCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubFeatureStateStateCodeEnum.Value(
            "GkehubFeatureStateStateCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubFeatureStateStateCodeEnum.Name(resource)[
            len("GkehubFeatureStateStateCodeEnum") :
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
