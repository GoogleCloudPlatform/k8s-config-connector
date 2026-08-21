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
from google3.cloud.graphite.mmv2.services.google.firebase import firebase_project_pb2
from google3.cloud.graphite.mmv2.services.google.firebase import (
    firebase_project_pb2_grpc,
)

from typing import List


class FirebaseProject(object):
    def __init__(
        self,
        project_id: str = None,
        project_number: int = None,
        display_name: str = None,
        resources: dict = None,
        state: str = None,
        annotations: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.display_name = display_name
        self.annotations = annotations
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = firebase_project_pb2_grpc.FirebaseBetaFirebaseProjectServiceStub(
            channel.Channel()
        )
        request = firebase_project_pb2.ApplyFirebaseBetaFirebaseProjectRequest()
        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyFirebaseBetaFirebaseProject(request)
        self.project_id = Primitive.from_proto(response.project_id)
        self.project_number = Primitive.from_proto(response.project_number)
        self.display_name = Primitive.from_proto(response.display_name)
        self.resources = FirebaseProjectResources.from_proto(response.resources)
        self.state = FirebaseProjectStateEnum.from_proto(response.state)
        self.annotations = Primitive.from_proto(response.annotations)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = firebase_project_pb2_grpc.FirebaseBetaFirebaseProjectServiceStub(
            channel.Channel()
        )
        request = firebase_project_pb2.DeleteFirebaseBetaFirebaseProjectRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteFirebaseBetaFirebaseProject(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = firebase_project_pb2_grpc.FirebaseBetaFirebaseProjectServiceStub(
            channel.Channel()
        )
        request = firebase_project_pb2.ListFirebaseBetaFirebaseProjectRequest()
        request.service_account_file = service_account_file
        return stub.ListFirebaseBetaFirebaseProject(request).items

    def to_proto(self):
        resource = firebase_project_pb2.FirebaseBetaFirebaseProject()
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class FirebaseProjectResources(object):
    def __init__(
        self,
        hosting_site: str = None,
        realtime_database_instance: str = None,
        storage_bucket: str = None,
        location_id: str = None,
    ):
        self.hosting_site = hosting_site
        self.realtime_database_instance = realtime_database_instance
        self.storage_bucket = storage_bucket
        self.location_id = location_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = firebase_project_pb2.FirebaseBetaFirebaseProjectResources()
        if Primitive.to_proto(resource.hosting_site):
            res.hosting_site = Primitive.to_proto(resource.hosting_site)
        if Primitive.to_proto(resource.realtime_database_instance):
            res.realtime_database_instance = Primitive.to_proto(
                resource.realtime_database_instance
            )
        if Primitive.to_proto(resource.storage_bucket):
            res.storage_bucket = Primitive.to_proto(resource.storage_bucket)
        if Primitive.to_proto(resource.location_id):
            res.location_id = Primitive.to_proto(resource.location_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FirebaseProjectResources(
            hosting_site=Primitive.from_proto(resource.hosting_site),
            realtime_database_instance=Primitive.from_proto(
                resource.realtime_database_instance
            ),
            storage_bucket=Primitive.from_proto(resource.storage_bucket),
            location_id=Primitive.from_proto(resource.location_id),
        )


class FirebaseProjectResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FirebaseProjectResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FirebaseProjectResources.from_proto(i) for i in resources]


class FirebaseProjectStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return firebase_project_pb2.FirebaseBetaFirebaseProjectStateEnum.Value(
            "FirebaseBetaFirebaseProjectStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return firebase_project_pb2.FirebaseBetaFirebaseProjectStateEnum.Name(resource)[
            len("FirebaseBetaFirebaseProjectStateEnum") :
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
