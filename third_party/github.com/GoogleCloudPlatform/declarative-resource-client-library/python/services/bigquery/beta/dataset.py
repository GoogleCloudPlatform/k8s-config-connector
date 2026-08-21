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
from google3.cloud.graphite.mmv2.services.google.bigquery import dataset_pb2
from google3.cloud.graphite.mmv2.services.google.bigquery import dataset_pb2_grpc

from typing import List


class Dataset(object):
    def __init__(
        self,
        etag: str = None,
        id: str = None,
        self_link: str = None,
        name: str = None,
        project: str = None,
        friendly_name: str = None,
        description: str = None,
        default_table_expiration_ms: str = None,
        default_partition_expiration_ms: str = None,
        labels: dict = None,
        access: list = None,
        creation_time: int = None,
        last_modified_time: int = None,
        location: str = None,
        published: bool = None,
        default_encryption_configuration: dict = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.project = project
        self.friendly_name = friendly_name
        self.description = description
        self.default_table_expiration_ms = default_table_expiration_ms
        self.default_partition_expiration_ms = default_partition_expiration_ms
        self.labels = labels
        self.access = access
        self.location = location
        self.published = published
        self.default_encryption_configuration = default_encryption_configuration
        self.service_account_file = service_account_file

    def apply(self):
        stub = dataset_pb2_grpc.BigqueryBetaDatasetServiceStub(channel.Channel())
        request = dataset_pb2.ApplyBigqueryBetaDatasetRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.friendly_name):
            request.resource.friendly_name = Primitive.to_proto(self.friendly_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.default_table_expiration_ms):
            request.resource.default_table_expiration_ms = Primitive.to_proto(
                self.default_table_expiration_ms
            )

        if Primitive.to_proto(self.default_partition_expiration_ms):
            request.resource.default_partition_expiration_ms = Primitive.to_proto(
                self.default_partition_expiration_ms
            )

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if DatasetAccessArray.to_proto(self.access):
            request.resource.access.extend(DatasetAccessArray.to_proto(self.access))
        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.published):
            request.resource.published = Primitive.to_proto(self.published)

        if DatasetDefaultEncryptionConfiguration.to_proto(
            self.default_encryption_configuration
        ):
            request.resource.default_encryption_configuration.CopyFrom(
                DatasetDefaultEncryptionConfiguration.to_proto(
                    self.default_encryption_configuration
                )
            )
        else:
            request.resource.ClearField("default_encryption_configuration")
        request.service_account_file = self.service_account_file

        response = stub.ApplyBigqueryBetaDataset(request)
        self.etag = Primitive.from_proto(response.etag)
        self.id = Primitive.from_proto(response.id)
        self.self_link = Primitive.from_proto(response.self_link)
        self.name = Primitive.from_proto(response.name)
        self.project = Primitive.from_proto(response.project)
        self.friendly_name = Primitive.from_proto(response.friendly_name)
        self.description = Primitive.from_proto(response.description)
        self.default_table_expiration_ms = Primitive.from_proto(
            response.default_table_expiration_ms
        )
        self.default_partition_expiration_ms = Primitive.from_proto(
            response.default_partition_expiration_ms
        )
        self.labels = Primitive.from_proto(response.labels)
        self.access = DatasetAccessArray.from_proto(response.access)
        self.creation_time = Primitive.from_proto(response.creation_time)
        self.last_modified_time = Primitive.from_proto(response.last_modified_time)
        self.location = Primitive.from_proto(response.location)
        self.published = Primitive.from_proto(response.published)
        self.default_encryption_configuration = (
            DatasetDefaultEncryptionConfiguration.from_proto(
                response.default_encryption_configuration
            )
        )

    def delete(self):
        stub = dataset_pb2_grpc.BigqueryBetaDatasetServiceStub(channel.Channel())
        request = dataset_pb2.DeleteBigqueryBetaDatasetRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.friendly_name):
            request.resource.friendly_name = Primitive.to_proto(self.friendly_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.default_table_expiration_ms):
            request.resource.default_table_expiration_ms = Primitive.to_proto(
                self.default_table_expiration_ms
            )

        if Primitive.to_proto(self.default_partition_expiration_ms):
            request.resource.default_partition_expiration_ms = Primitive.to_proto(
                self.default_partition_expiration_ms
            )

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if DatasetAccessArray.to_proto(self.access):
            request.resource.access.extend(DatasetAccessArray.to_proto(self.access))
        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.published):
            request.resource.published = Primitive.to_proto(self.published)

        if DatasetDefaultEncryptionConfiguration.to_proto(
            self.default_encryption_configuration
        ):
            request.resource.default_encryption_configuration.CopyFrom(
                DatasetDefaultEncryptionConfiguration.to_proto(
                    self.default_encryption_configuration
                )
            )
        else:
            request.resource.ClearField("default_encryption_configuration")
        response = stub.DeleteBigqueryBetaDataset(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = dataset_pb2_grpc.BigqueryBetaDatasetServiceStub(channel.Channel())
        request = dataset_pb2.ListBigqueryBetaDatasetRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListBigqueryBetaDataset(request).items

    def to_proto(self):
        resource = dataset_pb2.BigqueryBetaDataset()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.friendly_name):
            resource.friendly_name = Primitive.to_proto(self.friendly_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.default_table_expiration_ms):
            resource.default_table_expiration_ms = Primitive.to_proto(
                self.default_table_expiration_ms
            )
        if Primitive.to_proto(self.default_partition_expiration_ms):
            resource.default_partition_expiration_ms = Primitive.to_proto(
                self.default_partition_expiration_ms
            )
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if DatasetAccessArray.to_proto(self.access):
            resource.access.extend(DatasetAccessArray.to_proto(self.access))
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.published):
            resource.published = Primitive.to_proto(self.published)
        if DatasetDefaultEncryptionConfiguration.to_proto(
            self.default_encryption_configuration
        ):
            resource.default_encryption_configuration.CopyFrom(
                DatasetDefaultEncryptionConfiguration.to_proto(
                    self.default_encryption_configuration
                )
            )
        else:
            resource.ClearField("default_encryption_configuration")
        return resource


class DatasetAccess(object):
    def __init__(
        self,
        role: str = None,
        user_by_email: str = None,
        group_by_email: str = None,
        domain: str = None,
        special_group: str = None,
        iam_member: str = None,
        view: dict = None,
        routine: dict = None,
    ):
        self.role = role
        self.user_by_email = user_by_email
        self.group_by_email = group_by_email
        self.domain = domain
        self.special_group = special_group
        self.iam_member = iam_member
        self.view = view
        self.routine = routine

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dataset_pb2.BigqueryBetaDatasetAccess()
        if Primitive.to_proto(resource.role):
            res.role = Primitive.to_proto(resource.role)
        if Primitive.to_proto(resource.user_by_email):
            res.user_by_email = Primitive.to_proto(resource.user_by_email)
        if Primitive.to_proto(resource.group_by_email):
            res.group_by_email = Primitive.to_proto(resource.group_by_email)
        if Primitive.to_proto(resource.domain):
            res.domain = Primitive.to_proto(resource.domain)
        if Primitive.to_proto(resource.special_group):
            res.special_group = Primitive.to_proto(resource.special_group)
        if Primitive.to_proto(resource.iam_member):
            res.iam_member = Primitive.to_proto(resource.iam_member)
        if DatasetAccessView.to_proto(resource.view):
            res.view.CopyFrom(DatasetAccessView.to_proto(resource.view))
        else:
            res.ClearField("view")
        if DatasetAccessRoutine.to_proto(resource.routine):
            res.routine.CopyFrom(DatasetAccessRoutine.to_proto(resource.routine))
        else:
            res.ClearField("routine")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DatasetAccess(
            role=Primitive.from_proto(resource.role),
            user_by_email=Primitive.from_proto(resource.user_by_email),
            group_by_email=Primitive.from_proto(resource.group_by_email),
            domain=Primitive.from_proto(resource.domain),
            special_group=Primitive.from_proto(resource.special_group),
            iam_member=Primitive.from_proto(resource.iam_member),
            view=DatasetAccessView.from_proto(resource.view),
            routine=DatasetAccessRoutine.from_proto(resource.routine),
        )


class DatasetAccessArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DatasetAccess.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DatasetAccess.from_proto(i) for i in resources]


class DatasetAccessView(object):
    def __init__(
        self, project_id: str = None, dataset_id: str = None, table_id: str = None
    ):
        self.project_id = project_id
        self.dataset_id = dataset_id
        self.table_id = table_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dataset_pb2.BigqueryBetaDatasetAccessView()
        if Primitive.to_proto(resource.project_id):
            res.project_id = Primitive.to_proto(resource.project_id)
        if Primitive.to_proto(resource.dataset_id):
            res.dataset_id = Primitive.to_proto(resource.dataset_id)
        if Primitive.to_proto(resource.table_id):
            res.table_id = Primitive.to_proto(resource.table_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DatasetAccessView(
            project_id=Primitive.from_proto(resource.project_id),
            dataset_id=Primitive.from_proto(resource.dataset_id),
            table_id=Primitive.from_proto(resource.table_id),
        )


class DatasetAccessViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DatasetAccessView.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DatasetAccessView.from_proto(i) for i in resources]


class DatasetAccessRoutine(object):
    def __init__(
        self, project_id: str = None, dataset_id: str = None, routine_id: str = None
    ):
        self.project_id = project_id
        self.dataset_id = dataset_id
        self.routine_id = routine_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dataset_pb2.BigqueryBetaDatasetAccessRoutine()
        if Primitive.to_proto(resource.project_id):
            res.project_id = Primitive.to_proto(resource.project_id)
        if Primitive.to_proto(resource.dataset_id):
            res.dataset_id = Primitive.to_proto(resource.dataset_id)
        if Primitive.to_proto(resource.routine_id):
            res.routine_id = Primitive.to_proto(resource.routine_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DatasetAccessRoutine(
            project_id=Primitive.from_proto(resource.project_id),
            dataset_id=Primitive.from_proto(resource.dataset_id),
            routine_id=Primitive.from_proto(resource.routine_id),
        )


class DatasetAccessRoutineArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DatasetAccessRoutine.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DatasetAccessRoutine.from_proto(i) for i in resources]


class DatasetDefaultEncryptionConfiguration(object):
    def __init__(self, kms_key_name: str = None):
        self.kms_key_name = kms_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dataset_pb2.BigqueryBetaDatasetDefaultEncryptionConfiguration()
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DatasetDefaultEncryptionConfiguration(
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
        )


class DatasetDefaultEncryptionConfigurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DatasetDefaultEncryptionConfiguration.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DatasetDefaultEncryptionConfiguration.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
