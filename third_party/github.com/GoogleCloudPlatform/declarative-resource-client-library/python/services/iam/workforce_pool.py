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
from google3.cloud.graphite.mmv2.services.google.iam import workforce_pool_pb2
from google3.cloud.graphite.mmv2.services.google.iam import workforce_pool_pb2_grpc

from typing import List


class WorkforcePool(object):
    def __init__(
        self,
        name: str = None,
        self_link: str = None,
        parent: str = None,
        display_name: str = None,
        description: str = None,
        state: str = None,
        disabled: bool = None,
        session_duration: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.parent = parent
        self.display_name = display_name
        self.description = description
        self.disabled = disabled
        self.session_duration = session_duration
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = workforce_pool_pb2_grpc.IamWorkforcePoolServiceStub(channel.Channel())
        request = workforce_pool_pb2.ApplyIamWorkforcePoolRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.session_duration):
            request.resource.session_duration = Primitive.to_proto(
                self.session_duration
            )

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyIamWorkforcePool(request)
        self.name = Primitive.from_proto(response.name)
        self.self_link = Primitive.from_proto(response.self_link)
        self.parent = Primitive.from_proto(response.parent)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.state = WorkforcePoolStateEnum.from_proto(response.state)
        self.disabled = Primitive.from_proto(response.disabled)
        self.session_duration = Primitive.from_proto(response.session_duration)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = workforce_pool_pb2_grpc.IamWorkforcePoolServiceStub(channel.Channel())
        request = workforce_pool_pb2.DeleteIamWorkforcePoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.session_duration):
            request.resource.session_duration = Primitive.to_proto(
                self.session_duration
            )

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteIamWorkforcePool(request)

    @classmethod
    def list(self, location, parent, service_account_file=""):
        stub = workforce_pool_pb2_grpc.IamWorkforcePoolServiceStub(channel.Channel())
        request = workforce_pool_pb2.ListIamWorkforcePoolRequest()
        request.service_account_file = service_account_file
        request.Location = location

        request.Parent = parent

        return stub.ListIamWorkforcePool(request).items

    def to_proto(self):
        resource = workforce_pool_pb2.IamWorkforcePool()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if Primitive.to_proto(self.session_duration):
            resource.session_duration = Primitive.to_proto(self.session_duration)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class WorkforcePoolStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workforce_pool_pb2.IamWorkforcePoolStateEnum.Value(
            "IamWorkforcePoolStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workforce_pool_pb2.IamWorkforcePoolStateEnum.Name(resource)[
            len("IamWorkforcePoolStateEnum") :
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
