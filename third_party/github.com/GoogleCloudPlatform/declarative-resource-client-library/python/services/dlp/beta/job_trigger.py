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
from google3.cloud.graphite.mmv2.services.google.dlp import job_trigger_pb2
from google3.cloud.graphite.mmv2.services.google.dlp import job_trigger_pb2_grpc

from typing import List


class JobTrigger(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        description: str = None,
        inspect_job: dict = None,
        triggers: list = None,
        errors: list = None,
        create_time: str = None,
        update_time: str = None,
        last_run_time: str = None,
        status: str = None,
        location_id: str = None,
        parent: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.description = description
        self.inspect_job = inspect_job
        self.triggers = triggers
        self.status = status
        self.parent = parent
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = job_trigger_pb2_grpc.DlpBetaJobTriggerServiceStub(channel.Channel())
        request = job_trigger_pb2.ApplyDlpBetaJobTriggerRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if JobTriggerInspectJob.to_proto(self.inspect_job):
            request.resource.inspect_job.CopyFrom(
                JobTriggerInspectJob.to_proto(self.inspect_job)
            )
        else:
            request.resource.ClearField("inspect_job")
        if JobTriggerTriggersArray.to_proto(self.triggers):
            request.resource.triggers.extend(
                JobTriggerTriggersArray.to_proto(self.triggers)
            )
        if JobTriggerStatusEnum.to_proto(self.status):
            request.resource.status = JobTriggerStatusEnum.to_proto(self.status)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDlpBetaJobTrigger(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.inspect_job = JobTriggerInspectJob.from_proto(response.inspect_job)
        self.triggers = JobTriggerTriggersArray.from_proto(response.triggers)
        self.errors = JobTriggerErrorsArray.from_proto(response.errors)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.last_run_time = Primitive.from_proto(response.last_run_time)
        self.status = JobTriggerStatusEnum.from_proto(response.status)
        self.location_id = Primitive.from_proto(response.location_id)
        self.parent = Primitive.from_proto(response.parent)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = job_trigger_pb2_grpc.DlpBetaJobTriggerServiceStub(channel.Channel())
        request = job_trigger_pb2.DeleteDlpBetaJobTriggerRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if JobTriggerInspectJob.to_proto(self.inspect_job):
            request.resource.inspect_job.CopyFrom(
                JobTriggerInspectJob.to_proto(self.inspect_job)
            )
        else:
            request.resource.ClearField("inspect_job")
        if JobTriggerTriggersArray.to_proto(self.triggers):
            request.resource.triggers.extend(
                JobTriggerTriggersArray.to_proto(self.triggers)
            )
        if JobTriggerStatusEnum.to_proto(self.status):
            request.resource.status = JobTriggerStatusEnum.to_proto(self.status)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteDlpBetaJobTrigger(request)

    @classmethod
    def list(self, location, parent, service_account_file=""):
        stub = job_trigger_pb2_grpc.DlpBetaJobTriggerServiceStub(channel.Channel())
        request = job_trigger_pb2.ListDlpBetaJobTriggerRequest()
        request.service_account_file = service_account_file
        request.Location = location

        request.Parent = parent

        return stub.ListDlpBetaJobTrigger(request).items

    def to_proto(self):
        resource = job_trigger_pb2.DlpBetaJobTrigger()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if JobTriggerInspectJob.to_proto(self.inspect_job):
            resource.inspect_job.CopyFrom(
                JobTriggerInspectJob.to_proto(self.inspect_job)
            )
        else:
            resource.ClearField("inspect_job")
        if JobTriggerTriggersArray.to_proto(self.triggers):
            resource.triggers.extend(JobTriggerTriggersArray.to_proto(self.triggers))
        if JobTriggerStatusEnum.to_proto(self.status):
            resource.status = JobTriggerStatusEnum.to_proto(self.status)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class JobTriggerInspectJob(object):
    def __init__(
        self,
        storage_config: dict = None,
        inspect_config: dict = None,
        inspect_template_name: str = None,
        actions: list = None,
    ):
        self.storage_config = storage_config
        self.inspect_config = inspect_config
        self.inspect_template_name = inspect_template_name
        self.actions = actions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJob()
        if JobTriggerInspectJobStorageConfig.to_proto(resource.storage_config):
            res.storage_config.CopyFrom(
                JobTriggerInspectJobStorageConfig.to_proto(resource.storage_config)
            )
        else:
            res.ClearField("storage_config")
        if JobTriggerInspectJobInspectConfig.to_proto(resource.inspect_config):
            res.inspect_config.CopyFrom(
                JobTriggerInspectJobInspectConfig.to_proto(resource.inspect_config)
            )
        else:
            res.ClearField("inspect_config")
        if Primitive.to_proto(resource.inspect_template_name):
            res.inspect_template_name = Primitive.to_proto(
                resource.inspect_template_name
            )
        if JobTriggerInspectJobActionsArray.to_proto(resource.actions):
            res.actions.extend(
                JobTriggerInspectJobActionsArray.to_proto(resource.actions)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJob(
            storage_config=JobTriggerInspectJobStorageConfig.from_proto(
                resource.storage_config
            ),
            inspect_config=JobTriggerInspectJobInspectConfig.from_proto(
                resource.inspect_config
            ),
            inspect_template_name=Primitive.from_proto(resource.inspect_template_name),
            actions=JobTriggerInspectJobActionsArray.from_proto(resource.actions),
        )


class JobTriggerInspectJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerInspectJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTriggerInspectJob.from_proto(i) for i in resources]


class JobTriggerInspectJobStorageConfig(object):
    def __init__(
        self,
        datastore_options: dict = None,
        cloud_storage_options: dict = None,
        big_query_options: dict = None,
        hybrid_options: dict = None,
        timespan_config: dict = None,
    ):
        self.datastore_options = datastore_options
        self.cloud_storage_options = cloud_storage_options
        self.big_query_options = big_query_options
        self.hybrid_options = hybrid_options
        self.timespan_config = timespan_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfig()
        if JobTriggerInspectJobStorageConfigDatastoreOptions.to_proto(
            resource.datastore_options
        ):
            res.datastore_options.CopyFrom(
                JobTriggerInspectJobStorageConfigDatastoreOptions.to_proto(
                    resource.datastore_options
                )
            )
        else:
            res.ClearField("datastore_options")
        if JobTriggerInspectJobStorageConfigCloudStorageOptions.to_proto(
            resource.cloud_storage_options
        ):
            res.cloud_storage_options.CopyFrom(
                JobTriggerInspectJobStorageConfigCloudStorageOptions.to_proto(
                    resource.cloud_storage_options
                )
            )
        else:
            res.ClearField("cloud_storage_options")
        if JobTriggerInspectJobStorageConfigBigQueryOptions.to_proto(
            resource.big_query_options
        ):
            res.big_query_options.CopyFrom(
                JobTriggerInspectJobStorageConfigBigQueryOptions.to_proto(
                    resource.big_query_options
                )
            )
        else:
            res.ClearField("big_query_options")
        if JobTriggerInspectJobStorageConfigHybridOptions.to_proto(
            resource.hybrid_options
        ):
            res.hybrid_options.CopyFrom(
                JobTriggerInspectJobStorageConfigHybridOptions.to_proto(
                    resource.hybrid_options
                )
            )
        else:
            res.ClearField("hybrid_options")
        if JobTriggerInspectJobStorageConfigTimespanConfig.to_proto(
            resource.timespan_config
        ):
            res.timespan_config.CopyFrom(
                JobTriggerInspectJobStorageConfigTimespanConfig.to_proto(
                    resource.timespan_config
                )
            )
        else:
            res.ClearField("timespan_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfig(
            datastore_options=JobTriggerInspectJobStorageConfigDatastoreOptions.from_proto(
                resource.datastore_options
            ),
            cloud_storage_options=JobTriggerInspectJobStorageConfigCloudStorageOptions.from_proto(
                resource.cloud_storage_options
            ),
            big_query_options=JobTriggerInspectJobStorageConfigBigQueryOptions.from_proto(
                resource.big_query_options
            ),
            hybrid_options=JobTriggerInspectJobStorageConfigHybridOptions.from_proto(
                resource.hybrid_options
            ),
            timespan_config=JobTriggerInspectJobStorageConfigTimespanConfig.from_proto(
                resource.timespan_config
            ),
        )


class JobTriggerInspectJobStorageConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerInspectJobStorageConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTriggerInspectJobStorageConfig.from_proto(i) for i in resources]


class JobTriggerInspectJobStorageConfigDatastoreOptions(object):
    def __init__(self, partition_id: dict = None, kind: dict = None):
        self.partition_id = partition_id
        self.kind = kind

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptions()
        if JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId.to_proto(
            resource.partition_id
        ):
            res.partition_id.CopyFrom(
                JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId.to_proto(
                    resource.partition_id
                )
            )
        else:
            res.ClearField("partition_id")
        if JobTriggerInspectJobStorageConfigDatastoreOptionsKind.to_proto(
            resource.kind
        ):
            res.kind.CopyFrom(
                JobTriggerInspectJobStorageConfigDatastoreOptionsKind.to_proto(
                    resource.kind
                )
            )
        else:
            res.ClearField("kind")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigDatastoreOptions(
            partition_id=JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId.from_proto(
                resource.partition_id
            ),
            kind=JobTriggerInspectJobStorageConfigDatastoreOptionsKind.from_proto(
                resource.kind
            ),
        )


class JobTriggerInspectJobStorageConfigDatastoreOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigDatastoreOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigDatastoreOptions.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(object):
    def __init__(self, project_id: str = None, namespace_id: str = None):
        self.project_id = project_id
        self.namespace_id = namespace_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId()
        )
        if Primitive.to_proto(resource.project_id):
            res.project_id = Primitive.to_proto(resource.project_id)
        if Primitive.to_proto(resource.namespace_id):
            res.namespace_id = Primitive.to_proto(resource.namespace_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(
            project_id=Primitive.from_proto(resource.project_id),
            namespace_id=Primitive.from_proto(resource.namespace_id),
        )


class JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigDatastoreOptionsKind(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsKind()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigDatastoreOptionsKind(
            name=Primitive.from_proto(resource.name),
        )


class JobTriggerInspectJobStorageConfigDatastoreOptionsKindArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigDatastoreOptionsKind.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigDatastoreOptionsKind.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigCloudStorageOptions(object):
    def __init__(
        self,
        file_set: dict = None,
        bytes_limit_per_file: int = None,
        bytes_limit_per_file_percent: int = None,
        file_types: list = None,
        sample_method: str = None,
        files_limit_percent: int = None,
    ):
        self.file_set = file_set
        self.bytes_limit_per_file = bytes_limit_per_file
        self.bytes_limit_per_file_percent = bytes_limit_per_file_percent
        self.file_types = file_types
        self.sample_method = sample_method
        self.files_limit_percent = files_limit_percent

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptions()
        )
        if JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet.to_proto(
            resource.file_set
        ):
            res.file_set.CopyFrom(
                JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet.to_proto(
                    resource.file_set
                )
            )
        else:
            res.ClearField("file_set")
        if Primitive.to_proto(resource.bytes_limit_per_file):
            res.bytes_limit_per_file = Primitive.to_proto(resource.bytes_limit_per_file)
        if Primitive.to_proto(resource.bytes_limit_per_file_percent):
            res.bytes_limit_per_file_percent = Primitive.to_proto(
                resource.bytes_limit_per_file_percent
            )
        if JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumArray.to_proto(
            resource.file_types
        ):
            res.file_types.extend(
                JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumArray.to_proto(
                    resource.file_types
                )
            )
        if JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum.to_proto(
            resource.sample_method
        ):
            res.sample_method = JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum.to_proto(
                resource.sample_method
            )
        if Primitive.to_proto(resource.files_limit_percent):
            res.files_limit_percent = Primitive.to_proto(resource.files_limit_percent)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigCloudStorageOptions(
            file_set=JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet.from_proto(
                resource.file_set
            ),
            bytes_limit_per_file=Primitive.from_proto(resource.bytes_limit_per_file),
            bytes_limit_per_file_percent=Primitive.from_proto(
                resource.bytes_limit_per_file_percent
            ),
            file_types=JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumArray.from_proto(
                resource.file_types
            ),
            sample_method=JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum.from_proto(
                resource.sample_method
            ),
            files_limit_percent=Primitive.from_proto(resource.files_limit_percent),
        )


class JobTriggerInspectJobStorageConfigCloudStorageOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigCloudStorageOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigCloudStorageOptions.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(object):
    def __init__(self, url: str = None, regex_file_set: dict = None):
        self.url = url
        self.regex_file_set = regex_file_set

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet()
        )
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet.to_proto(
            resource.regex_file_set
        ):
            res.regex_file_set.CopyFrom(
                JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet.to_proto(
                    resource.regex_file_set
                )
            )
        else:
            res.ClearField("regex_file_set")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(
            url=Primitive.from_proto(resource.url),
            regex_file_set=JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet.from_proto(
                resource.regex_file_set
            ),
        )


class JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(object):
    def __init__(
        self,
        bucket_name: str = None,
        include_regex: list = None,
        exclude_regex: list = None,
    ):
        self.bucket_name = bucket_name
        self.include_regex = include_regex
        self.exclude_regex = exclude_regex

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet()
        )
        if Primitive.to_proto(resource.bucket_name):
            res.bucket_name = Primitive.to_proto(resource.bucket_name)
        if Primitive.to_proto(resource.include_regex):
            res.include_regex.extend(Primitive.to_proto(resource.include_regex))
        if Primitive.to_proto(resource.exclude_regex):
            res.exclude_regex.extend(Primitive.to_proto(resource.exclude_regex))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(
            bucket_name=Primitive.from_proto(resource.bucket_name),
            include_regex=Primitive.from_proto(resource.include_regex),
            exclude_regex=Primitive.from_proto(resource.exclude_regex),
        )


class JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigBigQueryOptions(object):
    def __init__(
        self,
        table_reference: dict = None,
        identifying_fields: list = None,
        rows_limit: int = None,
        rows_limit_percent: int = None,
        sample_method: str = None,
        excluded_fields: list = None,
        included_fields: list = None,
    ):
        self.table_reference = table_reference
        self.identifying_fields = identifying_fields
        self.rows_limit = rows_limit
        self.rows_limit_percent = rows_limit_percent
        self.sample_method = sample_method
        self.excluded_fields = excluded_fields
        self.included_fields = included_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptions()
        if JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference.to_proto(
            resource.table_reference
        ):
            res.table_reference.CopyFrom(
                JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference.to_proto(
                    resource.table_reference
                )
            )
        else:
            res.ClearField("table_reference")
        if JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsArray.to_proto(
            resource.identifying_fields
        ):
            res.identifying_fields.extend(
                JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsArray.to_proto(
                    resource.identifying_fields
                )
            )
        if Primitive.to_proto(resource.rows_limit):
            res.rows_limit = Primitive.to_proto(resource.rows_limit)
        if Primitive.to_proto(resource.rows_limit_percent):
            res.rows_limit_percent = Primitive.to_proto(resource.rows_limit_percent)
        if JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum.to_proto(
            resource.sample_method
        ):
            res.sample_method = JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum.to_proto(
                resource.sample_method
            )
        if JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsArray.to_proto(
            resource.excluded_fields
        ):
            res.excluded_fields.extend(
                JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsArray.to_proto(
                    resource.excluded_fields
                )
            )
        if JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsArray.to_proto(
            resource.included_fields
        ):
            res.included_fields.extend(
                JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsArray.to_proto(
                    resource.included_fields
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigBigQueryOptions(
            table_reference=JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference.from_proto(
                resource.table_reference
            ),
            identifying_fields=JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsArray.from_proto(
                resource.identifying_fields
            ),
            rows_limit=Primitive.from_proto(resource.rows_limit),
            rows_limit_percent=Primitive.from_proto(resource.rows_limit_percent),
            sample_method=JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum.from_proto(
                resource.sample_method
            ),
            excluded_fields=JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsArray.from_proto(
                resource.excluded_fields
            ),
            included_fields=JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsArray.from_proto(
                resource.included_fields
            ),
        )


class JobTriggerInspectJobStorageConfigBigQueryOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigBigQueryOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigBigQueryOptions.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(object):
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

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference()
        )
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

        return JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(
            project_id=Primitive.from_proto(resource.project_id),
            dataset_id=Primitive.from_proto(resource.dataset_id),
            table_id=Primitive.from_proto(resource.table_id),
        )


class JobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(
            name=Primitive.from_proto(resource.name),
        )


class JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(
            name=Primitive.from_proto(resource.name),
        )


class JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(
            name=Primitive.from_proto(resource.name),
        )


class JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigHybridOptions(object):
    def __init__(
        self,
        description: str = None,
        required_finding_label_keys: list = None,
        labels: dict = None,
        table_options: dict = None,
    ):
        self.description = description
        self.required_finding_label_keys = required_finding_label_keys
        self.labels = labels
        self.table_options = table_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigHybridOptions()
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.required_finding_label_keys):
            res.required_finding_label_keys.extend(
                Primitive.to_proto(resource.required_finding_label_keys)
            )
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if JobTriggerInspectJobStorageConfigHybridOptionsTableOptions.to_proto(
            resource.table_options
        ):
            res.table_options.CopyFrom(
                JobTriggerInspectJobStorageConfigHybridOptionsTableOptions.to_proto(
                    resource.table_options
                )
            )
        else:
            res.ClearField("table_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigHybridOptions(
            description=Primitive.from_proto(resource.description),
            required_finding_label_keys=Primitive.from_proto(
                resource.required_finding_label_keys
            ),
            labels=Primitive.from_proto(resource.labels),
            table_options=JobTriggerInspectJobStorageConfigHybridOptionsTableOptions.from_proto(
                resource.table_options
            ),
        )


class JobTriggerInspectJobStorageConfigHybridOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigHybridOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigHybridOptions.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigHybridOptionsTableOptions(object):
    def __init__(self, identifying_fields: list = None):
        self.identifying_fields = identifying_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptions()
        )
        if JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsArray.to_proto(
            resource.identifying_fields
        ):
            res.identifying_fields.extend(
                JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsArray.to_proto(
                    resource.identifying_fields
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigHybridOptionsTableOptions(
            identifying_fields=JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsArray.from_proto(
                resource.identifying_fields
            ),
        )


class JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigHybridOptionsTableOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigHybridOptionsTableOptions.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(
                name=Primitive.from_proto(resource.name),
            )
        )


class JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigTimespanConfig(object):
    def __init__(
        self,
        start_time: str = None,
        end_time: str = None,
        timestamp_field: dict = None,
        enable_auto_population_of_timespan_config: bool = None,
    ):
        self.start_time = start_time
        self.end_time = end_time
        self.timestamp_field = timestamp_field
        self.enable_auto_population_of_timespan_config = (
            enable_auto_population_of_timespan_config
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigTimespanConfig()
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.end_time):
            res.end_time = Primitive.to_proto(resource.end_time)
        if JobTriggerInspectJobStorageConfigTimespanConfigTimestampField.to_proto(
            resource.timestamp_field
        ):
            res.timestamp_field.CopyFrom(
                JobTriggerInspectJobStorageConfigTimespanConfigTimestampField.to_proto(
                    resource.timestamp_field
                )
            )
        else:
            res.ClearField("timestamp_field")
        if Primitive.to_proto(resource.enable_auto_population_of_timespan_config):
            res.enable_auto_population_of_timespan_config = Primitive.to_proto(
                resource.enable_auto_population_of_timespan_config
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigTimespanConfig(
            start_time=Primitive.from_proto(resource.start_time),
            end_time=Primitive.from_proto(resource.end_time),
            timestamp_field=JobTriggerInspectJobStorageConfigTimespanConfigTimestampField.from_proto(
                resource.timestamp_field
            ),
            enable_auto_population_of_timespan_config=Primitive.from_proto(
                resource.enable_auto_population_of_timespan_config
            ),
        )


class JobTriggerInspectJobStorageConfigTimespanConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigTimespanConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigTimespanConfig.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobStorageConfigTimespanConfigTimestampField(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigTimespanConfigTimestampField()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobStorageConfigTimespanConfigTimestampField(
            name=Primitive.from_proto(resource.name),
        )


class JobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobStorageConfigTimespanConfigTimestampField.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobStorageConfigTimespanConfigTimestampField.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfig(object):
    def __init__(
        self,
        info_types: list = None,
        min_likelihood: str = None,
        limits: dict = None,
        include_quote: bool = None,
        exclude_info_types: bool = None,
        custom_info_types: list = None,
        rule_set: list = None,
    ):
        self.info_types = info_types
        self.min_likelihood = min_likelihood
        self.limits = limits
        self.include_quote = include_quote
        self.exclude_info_types = exclude_info_types
        self.custom_info_types = custom_info_types
        self.rule_set = rule_set

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfig()
        if JobTriggerInspectJobInspectConfigInfoTypesArray.to_proto(
            resource.info_types
        ):
            res.info_types.extend(
                JobTriggerInspectJobInspectConfigInfoTypesArray.to_proto(
                    resource.info_types
                )
            )
        if JobTriggerInspectJobInspectConfigMinLikelihoodEnum.to_proto(
            resource.min_likelihood
        ):
            res.min_likelihood = (
                JobTriggerInspectJobInspectConfigMinLikelihoodEnum.to_proto(
                    resource.min_likelihood
                )
            )
        if JobTriggerInspectJobInspectConfigLimits.to_proto(resource.limits):
            res.limits.CopyFrom(
                JobTriggerInspectJobInspectConfigLimits.to_proto(resource.limits)
            )
        else:
            res.ClearField("limits")
        if Primitive.to_proto(resource.include_quote):
            res.include_quote = Primitive.to_proto(resource.include_quote)
        if Primitive.to_proto(resource.exclude_info_types):
            res.exclude_info_types = Primitive.to_proto(resource.exclude_info_types)
        if JobTriggerInspectJobInspectConfigCustomInfoTypesArray.to_proto(
            resource.custom_info_types
        ):
            res.custom_info_types.extend(
                JobTriggerInspectJobInspectConfigCustomInfoTypesArray.to_proto(
                    resource.custom_info_types
                )
            )
        if JobTriggerInspectJobInspectConfigRuleSetArray.to_proto(resource.rule_set):
            res.rule_set.extend(
                JobTriggerInspectJobInspectConfigRuleSetArray.to_proto(
                    resource.rule_set
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfig(
            info_types=JobTriggerInspectJobInspectConfigInfoTypesArray.from_proto(
                resource.info_types
            ),
            min_likelihood=JobTriggerInspectJobInspectConfigMinLikelihoodEnum.from_proto(
                resource.min_likelihood
            ),
            limits=JobTriggerInspectJobInspectConfigLimits.from_proto(resource.limits),
            include_quote=Primitive.from_proto(resource.include_quote),
            exclude_info_types=Primitive.from_proto(resource.exclude_info_types),
            custom_info_types=JobTriggerInspectJobInspectConfigCustomInfoTypesArray.from_proto(
                resource.custom_info_types
            ),
            rule_set=JobTriggerInspectJobInspectConfigRuleSetArray.from_proto(
                resource.rule_set
            ),
        )


class JobTriggerInspectJobInspectConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerInspectJobInspectConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTriggerInspectJobInspectConfig.from_proto(i) for i in resources]


class JobTriggerInspectJobInspectConfigInfoTypes(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigInfoTypes()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigInfoTypes(
            name=Primitive.from_proto(resource.name),
        )


class JobTriggerInspectJobInspectConfigInfoTypesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigInfoTypes.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigInfoTypes.from_proto(i) for i in resources
        ]


class JobTriggerInspectJobInspectConfigLimits(object):
    def __init__(
        self,
        max_findings_per_item: int = None,
        max_findings_per_request: int = None,
        max_findings_per_info_type: list = None,
    ):
        self.max_findings_per_item = max_findings_per_item
        self.max_findings_per_request = max_findings_per_request
        self.max_findings_per_info_type = max_findings_per_info_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigLimits()
        if Primitive.to_proto(resource.max_findings_per_item):
            res.max_findings_per_item = Primitive.to_proto(
                resource.max_findings_per_item
            )
        if Primitive.to_proto(resource.max_findings_per_request):
            res.max_findings_per_request = Primitive.to_proto(
                resource.max_findings_per_request
            )
        if JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeArray.to_proto(
            resource.max_findings_per_info_type
        ):
            res.max_findings_per_info_type.extend(
                JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeArray.to_proto(
                    resource.max_findings_per_info_type
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigLimits(
            max_findings_per_item=Primitive.from_proto(resource.max_findings_per_item),
            max_findings_per_request=Primitive.from_proto(
                resource.max_findings_per_request
            ),
            max_findings_per_info_type=JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeArray.from_proto(
                resource.max_findings_per_info_type
            ),
        )


class JobTriggerInspectJobInspectConfigLimitsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerInspectJobInspectConfigLimits.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigLimits.from_proto(i) for i in resources
        ]


class JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(object):
    def __init__(self, info_type: dict = None, max_findings: int = None):
        self.info_type = info_type
        self.max_findings = max_findings

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType()
        )
        if JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType.to_proto(
            resource.info_type
        ):
            res.info_type.CopyFrom(
                JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType.to_proto(
                    resource.info_type
                )
            )
        else:
            res.ClearField("info_type")
        if Primitive.to_proto(resource.max_findings):
            res.max_findings = Primitive.to_proto(resource.max_findings)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(
            info_type=JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType.from_proto(
                resource.info_type
            ),
            max_findings=Primitive.from_proto(resource.max_findings),
        )


class JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(object):
    def __init__(self, name: str = None, version: str = None):
        self.name = name
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(
            name=Primitive.from_proto(resource.name),
            version=Primitive.from_proto(resource.version),
        )


class JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypes(object):
    def __init__(
        self,
        info_type: dict = None,
        likelihood: str = None,
        dictionary: dict = None,
        regex: dict = None,
        surrogate_type: dict = None,
        stored_type: dict = None,
        detection_rules: list = None,
        exclusion_type: str = None,
    ):
        self.info_type = info_type
        self.likelihood = likelihood
        self.dictionary = dictionary
        self.regex = regex
        self.surrogate_type = surrogate_type
        self.stored_type = stored_type
        self.detection_rules = detection_rules
        self.exclusion_type = exclusion_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypes()
        if JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType.to_proto(
            resource.info_type
        ):
            res.info_type.CopyFrom(
                JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType.to_proto(
                    resource.info_type
                )
            )
        else:
            res.ClearField("info_type")
        if JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum.to_proto(
            resource.likelihood
        ):
            res.likelihood = (
                JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum.to_proto(
                    resource.likelihood
                )
            )
        if JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary.to_proto(
            resource.dictionary
        ):
            res.dictionary.CopyFrom(
                JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary.to_proto(
                    resource.dictionary
                )
            )
        else:
            res.ClearField("dictionary")
        if JobTriggerInspectJobInspectConfigCustomInfoTypesRegex.to_proto(
            resource.regex
        ):
            res.regex.CopyFrom(
                JobTriggerInspectJobInspectConfigCustomInfoTypesRegex.to_proto(
                    resource.regex
                )
            )
        else:
            res.ClearField("regex")
        if JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType.to_proto(
            resource.surrogate_type
        ):
            res.surrogate_type.CopyFrom(
                JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType.to_proto(
                    resource.surrogate_type
                )
            )
        else:
            res.ClearField("surrogate_type")
        if JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType.to_proto(
            resource.stored_type
        ):
            res.stored_type.CopyFrom(
                JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType.to_proto(
                    resource.stored_type
                )
            )
        else:
            res.ClearField("stored_type")
        if JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesArray.to_proto(
            resource.detection_rules
        ):
            res.detection_rules.extend(
                JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesArray.to_proto(
                    resource.detection_rules
                )
            )
        if JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum.to_proto(
            resource.exclusion_type
        ):
            res.exclusion_type = JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum.to_proto(
                resource.exclusion_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypes(
            info_type=JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType.from_proto(
                resource.info_type
            ),
            likelihood=JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum.from_proto(
                resource.likelihood
            ),
            dictionary=JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary.from_proto(
                resource.dictionary
            ),
            regex=JobTriggerInspectJobInspectConfigCustomInfoTypesRegex.from_proto(
                resource.regex
            ),
            surrogate_type=JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType.from_proto(
                resource.surrogate_type
            ),
            stored_type=JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType.from_proto(
                resource.stored_type
            ),
            detection_rules=JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesArray.from_proto(
                resource.detection_rules
            ),
            exclusion_type=JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum.from_proto(
                resource.exclusion_type
            ),
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypes.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypes.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(object):
    def __init__(self, name: str = None, version: str = None):
        self.name = name
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(
            name=Primitive.from_proto(resource.name),
            version=Primitive.from_proto(resource.version),
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(object):
    def __init__(self, word_list: dict = None, cloud_storage_path: dict = None):
        self.word_list = word_list
        self.cloud_storage_path = cloud_storage_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary()
        )
        if JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList.to_proto(
            resource.word_list
        ):
            res.word_list.CopyFrom(
                JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList.to_proto(
                    resource.word_list
                )
            )
        else:
            res.ClearField("word_list")
        if JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath.to_proto(
            resource.cloud_storage_path
        ):
            res.cloud_storage_path.CopyFrom(
                JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath.to_proto(
                    resource.cloud_storage_path
                )
            )
        else:
            res.ClearField("cloud_storage_path")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(
            word_list=JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList.from_proto(
                resource.word_list
            ),
            cloud_storage_path=JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath.from_proto(
                resource.cloud_storage_path
            ),
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(object):
    def __init__(self, words: list = None):
        self.words = words

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList()
        )
        if Primitive.to_proto(resource.words):
            res.words.extend(Primitive.to_proto(resource.words))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(
            words=Primitive.from_proto(resource.words),
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(
    object
):
    def __init__(self, path: str = None):
        self.path = path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath()
        )
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(
                path=Primitive.from_proto(resource.path),
            )
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesRegex(object):
    def __init__(self, pattern: str = None, group_indexes: list = None):
        self.pattern = pattern
        self.group_indexes = group_indexes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesRegex()
        )
        if Primitive.to_proto(resource.pattern):
            res.pattern = Primitive.to_proto(resource.pattern)
        if int64Array.to_proto(resource.group_indexes):
            res.group_indexes.extend(int64Array.to_proto(resource.group_indexes))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypesRegex(
            pattern=Primitive.from_proto(resource.pattern),
            group_indexes=int64Array.from_proto(resource.group_indexes),
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesRegexArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesRegex.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesRegex.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType()


class JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(object):
    def __init__(self, name: str = None, create_time: str = None):
        self.name = name
        self.create_time = create_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.create_time):
            res.create_time = Primitive.to_proto(resource.create_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(
            name=Primitive.from_proto(resource.name),
            create_time=Primitive.from_proto(resource.create_time),
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(object):
    def __init__(self, hotword_rule: dict = None):
        self.hotword_rule = hotword_rule

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules()
        )
        if JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule.to_proto(
            resource.hotword_rule
        ):
            res.hotword_rule.CopyFrom(
                JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule.to_proto(
                    resource.hotword_rule
                )
            )
        else:
            res.ClearField("hotword_rule")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(
            hotword_rule=JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule.from_proto(
                resource.hotword_rule
            ),
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(object):
    def __init__(
        self,
        hotword_regex: dict = None,
        proximity: dict = None,
        likelihood_adjustment: dict = None,
    ):
        self.hotword_regex = hotword_regex
        self.proximity = proximity
        self.likelihood_adjustment = likelihood_adjustment

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule()
        )
        if JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex.to_proto(
            resource.hotword_regex
        ):
            res.hotword_regex.CopyFrom(
                JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex.to_proto(
                    resource.hotword_regex
                )
            )
        else:
            res.ClearField("hotword_regex")
        if JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity.to_proto(
            resource.proximity
        ):
            res.proximity.CopyFrom(
                JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity.to_proto(
                    resource.proximity
                )
            )
        else:
            res.ClearField("proximity")
        if JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment.to_proto(
            resource.likelihood_adjustment
        ):
            res.likelihood_adjustment.CopyFrom(
                JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment.to_proto(
                    resource.likelihood_adjustment
                )
            )
        else:
            res.ClearField("likelihood_adjustment")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(
            hotword_regex=JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex.from_proto(
                resource.hotword_regex
            ),
            proximity=JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity.from_proto(
                resource.proximity
            ),
            likelihood_adjustment=JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment.from_proto(
                resource.likelihood_adjustment
            ),
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(
    object
):
    def __init__(self, pattern: str = None, group_indexes: list = None):
        self.pattern = pattern
        self.group_indexes = group_indexes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex()
        )
        if Primitive.to_proto(resource.pattern):
            res.pattern = Primitive.to_proto(resource.pattern)
        if int64Array.to_proto(resource.group_indexes):
            res.group_indexes.extend(int64Array.to_proto(resource.group_indexes))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(
            pattern=Primitive.from_proto(resource.pattern),
            group_indexes=int64Array.from_proto(resource.group_indexes),
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(
    object
):
    def __init__(self, window_before: int = None, window_after: int = None):
        self.window_before = window_before
        self.window_after = window_after

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity()
        )
        if Primitive.to_proto(resource.window_before):
            res.window_before = Primitive.to_proto(resource.window_before)
        if Primitive.to_proto(resource.window_after):
            res.window_after = Primitive.to_proto(resource.window_after)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(
            window_before=Primitive.from_proto(resource.window_before),
            window_after=Primitive.from_proto(resource.window_after),
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(
    object
):
    def __init__(self, fixed_likelihood: str = None, relative_likelihood: int = None):
        self.fixed_likelihood = fixed_likelihood
        self.relative_likelihood = relative_likelihood

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment()
        )
        if JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.to_proto(
            resource.fixed_likelihood
        ):
            res.fixed_likelihood = JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.to_proto(
                resource.fixed_likelihood
            )
        if Primitive.to_proto(resource.relative_likelihood):
            res.relative_likelihood = Primitive.to_proto(resource.relative_likelihood)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(
            fixed_likelihood=JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.from_proto(
                resource.fixed_likelihood
            ),
            relative_likelihood=Primitive.from_proto(resource.relative_likelihood),
        )


class JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSet(object):
    def __init__(self, info_types: list = None, rules: list = None):
        self.info_types = info_types
        self.rules = rules

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSet()
        if JobTriggerInspectJobInspectConfigRuleSetInfoTypesArray.to_proto(
            resource.info_types
        ):
            res.info_types.extend(
                JobTriggerInspectJobInspectConfigRuleSetInfoTypesArray.to_proto(
                    resource.info_types
                )
            )
        if JobTriggerInspectJobInspectConfigRuleSetRulesArray.to_proto(resource.rules):
            res.rules.extend(
                JobTriggerInspectJobInspectConfigRuleSetRulesArray.to_proto(
                    resource.rules
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSet(
            info_types=JobTriggerInspectJobInspectConfigRuleSetInfoTypesArray.from_proto(
                resource.info_types
            ),
            rules=JobTriggerInspectJobInspectConfigRuleSetRulesArray.from_proto(
                resource.rules
            ),
        )


class JobTriggerInspectJobInspectConfigRuleSetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerInspectJobInspectConfigRuleSet.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSet.from_proto(i) for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetInfoTypes(object):
    def __init__(self, name: str = None, version: str = None):
        self.name = name
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetInfoTypes()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetInfoTypes(
            name=Primitive.from_proto(resource.name),
            version=Primitive.from_proto(resource.version),
        )


class JobTriggerInspectJobInspectConfigRuleSetInfoTypesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetInfoTypes.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetInfoTypes.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRules(object):
    def __init__(self, hotword_rule: dict = None, exclusion_rule: dict = None):
        self.hotword_rule = hotword_rule
        self.exclusion_rule = exclusion_rule

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRules()
        if JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule.to_proto(
            resource.hotword_rule
        ):
            res.hotword_rule.CopyFrom(
                JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule.to_proto(
                    resource.hotword_rule
                )
            )
        else:
            res.ClearField("hotword_rule")
        if JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule.to_proto(
            resource.exclusion_rule
        ):
            res.exclusion_rule.CopyFrom(
                JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule.to_proto(
                    resource.exclusion_rule
                )
            )
        else:
            res.ClearField("exclusion_rule")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRules(
            hotword_rule=JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule.from_proto(
                resource.hotword_rule
            ),
            exclusion_rule=JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule.from_proto(
                resource.exclusion_rule
            ),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRules.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRules.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(object):
    def __init__(
        self,
        hotword_regex: dict = None,
        proximity: dict = None,
        likelihood_adjustment: dict = None,
    ):
        self.hotword_regex = hotword_regex
        self.proximity = proximity
        self.likelihood_adjustment = likelihood_adjustment

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule()
        )
        if JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex.to_proto(
            resource.hotword_regex
        ):
            res.hotword_regex.CopyFrom(
                JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex.to_proto(
                    resource.hotword_regex
                )
            )
        else:
            res.ClearField("hotword_regex")
        if JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity.to_proto(
            resource.proximity
        ):
            res.proximity.CopyFrom(
                JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity.to_proto(
                    resource.proximity
                )
            )
        else:
            res.ClearField("proximity")
        if JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment.to_proto(
            resource.likelihood_adjustment
        ):
            res.likelihood_adjustment.CopyFrom(
                JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment.to_proto(
                    resource.likelihood_adjustment
                )
            )
        else:
            res.ClearField("likelihood_adjustment")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(
            hotword_regex=JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex.from_proto(
                resource.hotword_regex
            ),
            proximity=JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity.from_proto(
                resource.proximity
            ),
            likelihood_adjustment=JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment.from_proto(
                resource.likelihood_adjustment
            ),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(object):
    def __init__(self, pattern: str = None, group_indexes: list = None):
        self.pattern = pattern
        self.group_indexes = group_indexes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex()
        )
        if Primitive.to_proto(resource.pattern):
            res.pattern = Primitive.to_proto(resource.pattern)
        if int64Array.to_proto(resource.group_indexes):
            res.group_indexes.extend(int64Array.to_proto(resource.group_indexes))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(
            pattern=Primitive.from_proto(resource.pattern),
            group_indexes=int64Array.from_proto(resource.group_indexes),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(object):
    def __init__(self, window_before: int = None, window_after: int = None):
        self.window_before = window_before
        self.window_after = window_after

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity()
        )
        if Primitive.to_proto(resource.window_before):
            res.window_before = Primitive.to_proto(resource.window_before)
        if Primitive.to_proto(resource.window_after):
            res.window_after = Primitive.to_proto(resource.window_after)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(
            window_before=Primitive.from_proto(resource.window_before),
            window_after=Primitive.from_proto(resource.window_after),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(
    object
):
    def __init__(self, fixed_likelihood: str = None, relative_likelihood: int = None):
        self.fixed_likelihood = fixed_likelihood
        self.relative_likelihood = relative_likelihood

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment()
        )
        if JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.to_proto(
            resource.fixed_likelihood
        ):
            res.fixed_likelihood = JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.to_proto(
                resource.fixed_likelihood
            )
        if Primitive.to_proto(resource.relative_likelihood):
            res.relative_likelihood = Primitive.to_proto(resource.relative_likelihood)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(
            fixed_likelihood=JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.from_proto(
                resource.fixed_likelihood
            ),
            relative_likelihood=Primitive.from_proto(resource.relative_likelihood),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(object):
    def __init__(
        self,
        dictionary: dict = None,
        regex: dict = None,
        exclude_info_types: dict = None,
        matching_type: str = None,
    ):
        self.dictionary = dictionary
        self.regex = regex
        self.exclude_info_types = exclude_info_types
        self.matching_type = matching_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule()
        )
        if JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary.to_proto(
            resource.dictionary
        ):
            res.dictionary.CopyFrom(
                JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary.to_proto(
                    resource.dictionary
                )
            )
        else:
            res.ClearField("dictionary")
        if JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex.to_proto(
            resource.regex
        ):
            res.regex.CopyFrom(
                JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex.to_proto(
                    resource.regex
                )
            )
        else:
            res.ClearField("regex")
        if JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes.to_proto(
            resource.exclude_info_types
        ):
            res.exclude_info_types.CopyFrom(
                JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes.to_proto(
                    resource.exclude_info_types
                )
            )
        else:
            res.ClearField("exclude_info_types")
        if JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.to_proto(
            resource.matching_type
        ):
            res.matching_type = JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.to_proto(
                resource.matching_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(
            dictionary=JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary.from_proto(
                resource.dictionary
            ),
            regex=JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex.from_proto(
                resource.regex
            ),
            exclude_info_types=JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes.from_proto(
                resource.exclude_info_types
            ),
            matching_type=JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.from_proto(
                resource.matching_type
            ),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(object):
    def __init__(self, word_list: dict = None, cloud_storage_path: dict = None):
        self.word_list = word_list
        self.cloud_storage_path = cloud_storage_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary()
        )
        if JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList.to_proto(
            resource.word_list
        ):
            res.word_list.CopyFrom(
                JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList.to_proto(
                    resource.word_list
                )
            )
        else:
            res.ClearField("word_list")
        if JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath.to_proto(
            resource.cloud_storage_path
        ):
            res.cloud_storage_path.CopyFrom(
                JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath.to_proto(
                    resource.cloud_storage_path
                )
            )
        else:
            res.ClearField("cloud_storage_path")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(
            word_list=JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList.from_proto(
                resource.word_list
            ),
            cloud_storage_path=JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath.from_proto(
                resource.cloud_storage_path
            ),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(
    object
):
    def __init__(self, words: list = None):
        self.words = words

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList()
        )
        if Primitive.to_proto(resource.words):
            res.words.extend(Primitive.to_proto(resource.words))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(
            words=Primitive.from_proto(resource.words),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(
    object
):
    def __init__(self, path: str = None):
        self.path = path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath()
        )
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(
            path=Primitive.from_proto(resource.path),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(object):
    def __init__(self, pattern: str = None, group_indexes: list = None):
        self.pattern = pattern
        self.group_indexes = group_indexes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex()
        )
        if Primitive.to_proto(resource.pattern):
            res.pattern = Primitive.to_proto(resource.pattern)
        if int64Array.to_proto(resource.group_indexes):
            res.group_indexes.extend(int64Array.to_proto(resource.group_indexes))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(
            pattern=Primitive.from_proto(resource.pattern),
            group_indexes=int64Array.from_proto(resource.group_indexes),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(
    object
):
    def __init__(self, info_types: list = None):
        self.info_types = info_types

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes()
        )
        if JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesArray.to_proto(
            resource.info_types
        ):
            res.info_types.extend(
                JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesArray.to_proto(
                    resource.info_types
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(
            info_types=JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesArray.from_proto(
                resource.info_types
            ),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(
    object
):
    def __init__(self, name: str = None, version: str = None):
        self.name = name
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(
            name=Primitive.from_proto(resource.name),
            version=Primitive.from_proto(resource.version),
        )


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes.from_proto(
                i
            )
            for i in resources
        ]


class JobTriggerInspectJobActions(object):
    def __init__(
        self,
        save_findings: dict = None,
        pub_sub: dict = None,
        publish_summary_to_cscc: dict = None,
        publish_findings_to_cloud_data_catalog: dict = None,
        job_notification_emails: dict = None,
        publish_to_stackdriver: dict = None,
    ):
        self.save_findings = save_findings
        self.pub_sub = pub_sub
        self.publish_summary_to_cscc = publish_summary_to_cscc
        self.publish_findings_to_cloud_data_catalog = (
            publish_findings_to_cloud_data_catalog
        )
        self.job_notification_emails = job_notification_emails
        self.publish_to_stackdriver = publish_to_stackdriver

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobActions()
        if JobTriggerInspectJobActionsSaveFindings.to_proto(resource.save_findings):
            res.save_findings.CopyFrom(
                JobTriggerInspectJobActionsSaveFindings.to_proto(resource.save_findings)
            )
        else:
            res.ClearField("save_findings")
        if JobTriggerInspectJobActionsPubSub.to_proto(resource.pub_sub):
            res.pub_sub.CopyFrom(
                JobTriggerInspectJobActionsPubSub.to_proto(resource.pub_sub)
            )
        else:
            res.ClearField("pub_sub")
        if JobTriggerInspectJobActionsPublishSummaryToCscc.to_proto(
            resource.publish_summary_to_cscc
        ):
            res.publish_summary_to_cscc.CopyFrom(
                JobTriggerInspectJobActionsPublishSummaryToCscc.to_proto(
                    resource.publish_summary_to_cscc
                )
            )
        else:
            res.ClearField("publish_summary_to_cscc")
        if JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog.to_proto(
            resource.publish_findings_to_cloud_data_catalog
        ):
            res.publish_findings_to_cloud_data_catalog.CopyFrom(
                JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog.to_proto(
                    resource.publish_findings_to_cloud_data_catalog
                )
            )
        else:
            res.ClearField("publish_findings_to_cloud_data_catalog")
        if JobTriggerInspectJobActionsJobNotificationEmails.to_proto(
            resource.job_notification_emails
        ):
            res.job_notification_emails.CopyFrom(
                JobTriggerInspectJobActionsJobNotificationEmails.to_proto(
                    resource.job_notification_emails
                )
            )
        else:
            res.ClearField("job_notification_emails")
        if JobTriggerInspectJobActionsPublishToStackdriver.to_proto(
            resource.publish_to_stackdriver
        ):
            res.publish_to_stackdriver.CopyFrom(
                JobTriggerInspectJobActionsPublishToStackdriver.to_proto(
                    resource.publish_to_stackdriver
                )
            )
        else:
            res.ClearField("publish_to_stackdriver")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobActions(
            save_findings=JobTriggerInspectJobActionsSaveFindings.from_proto(
                resource.save_findings
            ),
            pub_sub=JobTriggerInspectJobActionsPubSub.from_proto(resource.pub_sub),
            publish_summary_to_cscc=JobTriggerInspectJobActionsPublishSummaryToCscc.from_proto(
                resource.publish_summary_to_cscc
            ),
            publish_findings_to_cloud_data_catalog=JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog.from_proto(
                resource.publish_findings_to_cloud_data_catalog
            ),
            job_notification_emails=JobTriggerInspectJobActionsJobNotificationEmails.from_proto(
                resource.job_notification_emails
            ),
            publish_to_stackdriver=JobTriggerInspectJobActionsPublishToStackdriver.from_proto(
                resource.publish_to_stackdriver
            ),
        )


class JobTriggerInspectJobActionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerInspectJobActions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTriggerInspectJobActions.from_proto(i) for i in resources]


class JobTriggerInspectJobActionsSaveFindings(object):
    def __init__(self, output_config: dict = None):
        self.output_config = output_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobActionsSaveFindings()
        if JobTriggerInspectJobActionsSaveFindingsOutputConfig.to_proto(
            resource.output_config
        ):
            res.output_config.CopyFrom(
                JobTriggerInspectJobActionsSaveFindingsOutputConfig.to_proto(
                    resource.output_config
                )
            )
        else:
            res.ClearField("output_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobActionsSaveFindings(
            output_config=JobTriggerInspectJobActionsSaveFindingsOutputConfig.from_proto(
                resource.output_config
            ),
        )


class JobTriggerInspectJobActionsSaveFindingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerInspectJobActionsSaveFindings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobActionsSaveFindings.from_proto(i) for i in resources
        ]


class JobTriggerInspectJobActionsSaveFindingsOutputConfig(object):
    def __init__(
        self, table: dict = None, dlp_storage: dict = None, output_schema: str = None
    ):
        self.table = table
        self.dlp_storage = dlp_storage
        self.output_schema = output_schema

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfig()
        )
        if JobTriggerInspectJobActionsSaveFindingsOutputConfigTable.to_proto(
            resource.table
        ):
            res.table.CopyFrom(
                JobTriggerInspectJobActionsSaveFindingsOutputConfigTable.to_proto(
                    resource.table
                )
            )
        else:
            res.ClearField("table")
        if JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage.to_proto(
            resource.dlp_storage
        ):
            res.dlp_storage.CopyFrom(
                JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage.to_proto(
                    resource.dlp_storage
                )
            )
        else:
            res.ClearField("dlp_storage")
        if JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum.to_proto(
            resource.output_schema
        ):
            res.output_schema = JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum.to_proto(
                resource.output_schema
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobActionsSaveFindingsOutputConfig(
            table=JobTriggerInspectJobActionsSaveFindingsOutputConfigTable.from_proto(
                resource.table
            ),
            dlp_storage=JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage.from_proto(
                resource.dlp_storage
            ),
            output_schema=JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum.from_proto(
                resource.output_schema
            ),
        )


class JobTriggerInspectJobActionsSaveFindingsOutputConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobActionsSaveFindingsOutputConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobActionsSaveFindingsOutputConfig.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobActionsSaveFindingsOutputConfigTable(object):
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

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigTable()
        )
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

        return JobTriggerInspectJobActionsSaveFindingsOutputConfigTable(
            project_id=Primitive.from_proto(resource.project_id),
            dataset_id=Primitive.from_proto(resource.dataset_id),
            table_id=Primitive.from_proto(resource.table_id),
        )


class JobTriggerInspectJobActionsSaveFindingsOutputConfigTableArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobActionsSaveFindingsOutputConfigTable.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobActionsSaveFindingsOutputConfigTable.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage()


class JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobActionsPubSub(object):
    def __init__(self, topic: str = None):
        self.topic = topic

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobActionsPubSub()
        if Primitive.to_proto(resource.topic):
            res.topic = Primitive.to_proto(resource.topic)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobActionsPubSub(
            topic=Primitive.from_proto(resource.topic),
        )


class JobTriggerInspectJobActionsPubSubArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerInspectJobActionsPubSub.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTriggerInspectJobActionsPubSub.from_proto(i) for i in resources]


class JobTriggerInspectJobActionsPublishSummaryToCscc(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobActionsPublishSummaryToCscc()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobActionsPublishSummaryToCscc()


class JobTriggerInspectJobActionsPublishSummaryToCsccArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobActionsPublishSummaryToCscc.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobActionsPublishSummaryToCscc.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            job_trigger_pb2.DlpBetaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog()


class JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobActionsJobNotificationEmails(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobActionsJobNotificationEmails()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobActionsJobNotificationEmails()


class JobTriggerInspectJobActionsJobNotificationEmailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobActionsJobNotificationEmails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobActionsJobNotificationEmails.from_proto(i)
            for i in resources
        ]


class JobTriggerInspectJobActionsPublishToStackdriver(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerInspectJobActionsPublishToStackdriver()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerInspectJobActionsPublishToStackdriver()


class JobTriggerInspectJobActionsPublishToStackdriverArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTriggerInspectJobActionsPublishToStackdriver.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTriggerInspectJobActionsPublishToStackdriver.from_proto(i)
            for i in resources
        ]


class JobTriggerTriggers(object):
    def __init__(self, schedule: dict = None, manual: dict = None):
        self.schedule = schedule
        self.manual = manual

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerTriggers()
        if JobTriggerTriggersSchedule.to_proto(resource.schedule):
            res.schedule.CopyFrom(
                JobTriggerTriggersSchedule.to_proto(resource.schedule)
            )
        else:
            res.ClearField("schedule")
        if JobTriggerTriggersManual.to_proto(resource.manual):
            res.manual.CopyFrom(JobTriggerTriggersManual.to_proto(resource.manual))
        else:
            res.ClearField("manual")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerTriggers(
            schedule=JobTriggerTriggersSchedule.from_proto(resource.schedule),
            manual=JobTriggerTriggersManual.from_proto(resource.manual),
        )


class JobTriggerTriggersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerTriggers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTriggerTriggers.from_proto(i) for i in resources]


class JobTriggerTriggersSchedule(object):
    def __init__(self, recurrence_period_duration: str = None):
        self.recurrence_period_duration = recurrence_period_duration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerTriggersSchedule()
        if Primitive.to_proto(resource.recurrence_period_duration):
            res.recurrence_period_duration = Primitive.to_proto(
                resource.recurrence_period_duration
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerTriggersSchedule(
            recurrence_period_duration=Primitive.from_proto(
                resource.recurrence_period_duration
            ),
        )


class JobTriggerTriggersScheduleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerTriggersSchedule.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTriggerTriggersSchedule.from_proto(i) for i in resources]


class JobTriggerTriggersManual(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerTriggersManual()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerTriggersManual()


class JobTriggerTriggersManualArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerTriggersManual.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTriggerTriggersManual.from_proto(i) for i in resources]


class JobTriggerErrors(object):
    def __init__(self, details: dict = None, timestamps: list = None):
        self.details = details
        self.timestamps = timestamps

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerErrors()
        if JobTriggerErrorsDetails.to_proto(resource.details):
            res.details.CopyFrom(JobTriggerErrorsDetails.to_proto(resource.details))
        else:
            res.ClearField("details")
        if Primitive.to_proto(resource.timestamps):
            res.timestamps.extend(Primitive.to_proto(resource.timestamps))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerErrors(
            details=JobTriggerErrorsDetails.from_proto(resource.details),
            timestamps=Primitive.from_proto(resource.timestamps),
        )


class JobTriggerErrorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerErrors.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTriggerErrors.from_proto(i) for i in resources]


class JobTriggerErrorsDetails(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerErrorsDetails()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if JobTriggerErrorsDetailsDetailsArray.to_proto(resource.details):
            res.details.extend(
                JobTriggerErrorsDetailsDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerErrorsDetails(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=JobTriggerErrorsDetailsDetailsArray.from_proto(resource.details),
        )


class JobTriggerErrorsDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerErrorsDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTriggerErrorsDetails.from_proto(i) for i in resources]


class JobTriggerErrorsDetailsDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_trigger_pb2.DlpBetaJobTriggerErrorsDetailsDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTriggerErrorsDetailsDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class JobTriggerErrorsDetailsDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTriggerErrorsDetailsDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTriggerErrorsDetailsDetails.from_proto(i) for i in resources]


class JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum.Value(
            "DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum.Name(
            resource
        )[
            len(
                "DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum"
            ) :
        ]


class JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum.Value(
            "DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum.Name(
            resource
        )[
            len(
                "DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum"
            ) :
        ]


class JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum.Value(
            "DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum.Name(
            resource
        )[
            len(
                "DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum"
            ) :
        ]


class JobTriggerInspectJobInspectConfigMinLikelihoodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum.Value(
            "DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum.Name(
            resource
        )[
            len("DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum") :
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum.Value(
            "DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum.Name(
            resource
        )[
            len(
                "DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum"
            ) :
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.Value(
            "DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.Name(
            resource
        )[
            len(
                "DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"
            ) :
        ]


class JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum.Value(
            "DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum.Name(
            resource
        )[
            len(
                "DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum"
            ) :
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.Value(
            "DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.Name(
            resource
        )[
            len(
                "DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"
            ) :
        ]


class JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.Value(
            "DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.Name(
            resource
        )[
            len(
                "DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"
            ) :
        ]


class JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum.Value(
            "DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum.Name(
            resource
        )[
            len(
                "DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum"
            ) :
        ]


class JobTriggerStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerStatusEnum.Value(
            "DlpBetaJobTriggerStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_trigger_pb2.DlpBetaJobTriggerStatusEnum.Name(resource)[
            len("DlpBetaJobTriggerStatusEnum") :
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
