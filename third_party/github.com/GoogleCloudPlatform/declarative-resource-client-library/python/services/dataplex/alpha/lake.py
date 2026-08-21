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
from google3.cloud.graphite.mmv2.services.google.dataplex import lake_pb2
from google3.cloud.graphite.mmv2.services.google.dataplex import lake_pb2_grpc

from typing import List


class Lake(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        uid: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        description: str = None,
        state: str = None,
        service_account: str = None,
        metastore: dict = None,
        asset_status: dict = None,
        metastore_status: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.labels = labels
        self.description = description
        self.metastore = metastore
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = lake_pb2_grpc.DataplexAlphaLakeServiceStub(channel.Channel())
        request = lake_pb2.ApplyDataplexAlphaLakeRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if LakeMetastore.to_proto(self.metastore):
            request.resource.metastore.CopyFrom(LakeMetastore.to_proto(self.metastore))
        else:
            request.resource.ClearField("metastore")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDataplexAlphaLake(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.uid = Primitive.from_proto(response.uid)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.description = Primitive.from_proto(response.description)
        self.state = LakeStateEnum.from_proto(response.state)
        self.service_account = Primitive.from_proto(response.service_account)
        self.metastore = LakeMetastore.from_proto(response.metastore)
        self.asset_status = LakeAssetStatus.from_proto(response.asset_status)
        self.metastore_status = LakeMetastoreStatus.from_proto(
            response.metastore_status
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = lake_pb2_grpc.DataplexAlphaLakeServiceStub(channel.Channel())
        request = lake_pb2.DeleteDataplexAlphaLakeRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if LakeMetastore.to_proto(self.metastore):
            request.resource.metastore.CopyFrom(LakeMetastore.to_proto(self.metastore))
        else:
            request.resource.ClearField("metastore")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteDataplexAlphaLake(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = lake_pb2_grpc.DataplexAlphaLakeServiceStub(channel.Channel())
        request = lake_pb2.ListDataplexAlphaLakeRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListDataplexAlphaLake(request).items

    def to_proto(self):
        resource = lake_pb2.DataplexAlphaLake()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if LakeMetastore.to_proto(self.metastore):
            resource.metastore.CopyFrom(LakeMetastore.to_proto(self.metastore))
        else:
            resource.ClearField("metastore")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class LakeMetastore(object):
    def __init__(self, service: str = None):
        self.service = service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = lake_pb2.DataplexAlphaLakeMetastore()
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return LakeMetastore(
            service=Primitive.from_proto(resource.service),
        )


class LakeMetastoreArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [LakeMetastore.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [LakeMetastore.from_proto(i) for i in resources]


class LakeAssetStatus(object):
    def __init__(
        self,
        update_time: str = None,
        active_assets: int = None,
        security_policy_applying_assets: int = None,
    ):
        self.update_time = update_time
        self.active_assets = active_assets
        self.security_policy_applying_assets = security_policy_applying_assets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = lake_pb2.DataplexAlphaLakeAssetStatus()
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        if Primitive.to_proto(resource.active_assets):
            res.active_assets = Primitive.to_proto(resource.active_assets)
        if Primitive.to_proto(resource.security_policy_applying_assets):
            res.security_policy_applying_assets = Primitive.to_proto(
                resource.security_policy_applying_assets
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return LakeAssetStatus(
            update_time=Primitive.from_proto(resource.update_time),
            active_assets=Primitive.from_proto(resource.active_assets),
            security_policy_applying_assets=Primitive.from_proto(
                resource.security_policy_applying_assets
            ),
        )


class LakeAssetStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [LakeAssetStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [LakeAssetStatus.from_proto(i) for i in resources]


class LakeMetastoreStatus(object):
    def __init__(
        self,
        state: str = None,
        message: str = None,
        update_time: str = None,
        endpoint: str = None,
    ):
        self.state = state
        self.message = message
        self.update_time = update_time
        self.endpoint = endpoint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = lake_pb2.DataplexAlphaLakeMetastoreStatus()
        if LakeMetastoreStatusStateEnum.to_proto(resource.state):
            res.state = LakeMetastoreStatusStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        if Primitive.to_proto(resource.endpoint):
            res.endpoint = Primitive.to_proto(resource.endpoint)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return LakeMetastoreStatus(
            state=LakeMetastoreStatusStateEnum.from_proto(resource.state),
            message=Primitive.from_proto(resource.message),
            update_time=Primitive.from_proto(resource.update_time),
            endpoint=Primitive.from_proto(resource.endpoint),
        )


class LakeMetastoreStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [LakeMetastoreStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [LakeMetastoreStatus.from_proto(i) for i in resources]


class LakeStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return lake_pb2.DataplexAlphaLakeStateEnum.Value(
            "DataplexAlphaLakeStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return lake_pb2.DataplexAlphaLakeStateEnum.Name(resource)[
            len("DataplexAlphaLakeStateEnum") :
        ]


class LakeMetastoreStatusStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return lake_pb2.DataplexAlphaLakeMetastoreStatusStateEnum.Value(
            "DataplexAlphaLakeMetastoreStatusStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return lake_pb2.DataplexAlphaLakeMetastoreStatusStateEnum.Name(resource)[
            len("DataplexAlphaLakeMetastoreStatusStateEnum") :
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
