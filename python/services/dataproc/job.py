# Copyright 2021 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.dataproc import job_pb2
from google3.cloud.graphite.mmv2.services.google.dataproc import job_pb2_grpc

from typing import List


class Job(object):
    def __init__(
        self,
        reference: dict = None,
        placement: dict = None,
        hadoop_job: dict = None,
        spark_job: dict = None,
        pyspark_job: dict = None,
        hive_job: dict = None,
        pig_job: dict = None,
        spark_r_job: dict = None,
        spark_sql_job: dict = None,
        presto_job: dict = None,
        status: dict = None,
        status_history: list = None,
        yarn_applications: list = None,
        driver_output_resource_uri: str = None,
        driver_control_files_uri: str = None,
        labels: dict = None,
        scheduling: dict = None,
        name: str = None,
        done: bool = None,
        region: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.reference = reference
        self.placement = placement
        self.hadoop_job = hadoop_job
        self.spark_job = spark_job
        self.pyspark_job = pyspark_job
        self.hive_job = hive_job
        self.pig_job = pig_job
        self.spark_r_job = spark_r_job
        self.spark_sql_job = spark_sql_job
        self.presto_job = presto_job
        self.labels = labels
        self.scheduling = scheduling
        self.region = region
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = job_pb2_grpc.DataprocJobServiceStub(channel.Channel())
        request = job_pb2.ApplyDataprocJobRequest()
        if JobReference.to_proto(self.reference):
            request.resource.reference.CopyFrom(JobReference.to_proto(self.reference))
        else:
            request.resource.ClearField("reference")
        if JobPlacement.to_proto(self.placement):
            request.resource.placement.CopyFrom(JobPlacement.to_proto(self.placement))
        else:
            request.resource.ClearField("placement")
        if JobHadoopJob.to_proto(self.hadoop_job):
            request.resource.hadoop_job.CopyFrom(JobHadoopJob.to_proto(self.hadoop_job))
        else:
            request.resource.ClearField("hadoop_job")
        if JobSparkJob.to_proto(self.spark_job):
            request.resource.spark_job.CopyFrom(JobSparkJob.to_proto(self.spark_job))
        else:
            request.resource.ClearField("spark_job")
        if JobPysparkJob.to_proto(self.pyspark_job):
            request.resource.pyspark_job.CopyFrom(
                JobPysparkJob.to_proto(self.pyspark_job)
            )
        else:
            request.resource.ClearField("pyspark_job")
        if JobHiveJob.to_proto(self.hive_job):
            request.resource.hive_job.CopyFrom(JobHiveJob.to_proto(self.hive_job))
        else:
            request.resource.ClearField("hive_job")
        if JobPigJob.to_proto(self.pig_job):
            request.resource.pig_job.CopyFrom(JobPigJob.to_proto(self.pig_job))
        else:
            request.resource.ClearField("pig_job")
        if JobSparkRJob.to_proto(self.spark_r_job):
            request.resource.spark_r_job.CopyFrom(
                JobSparkRJob.to_proto(self.spark_r_job)
            )
        else:
            request.resource.ClearField("spark_r_job")
        if JobSparkSqlJob.to_proto(self.spark_sql_job):
            request.resource.spark_sql_job.CopyFrom(
                JobSparkSqlJob.to_proto(self.spark_sql_job)
            )
        else:
            request.resource.ClearField("spark_sql_job")
        if JobPrestoJob.to_proto(self.presto_job):
            request.resource.presto_job.CopyFrom(JobPrestoJob.to_proto(self.presto_job))
        else:
            request.resource.ClearField("presto_job")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if JobScheduling.to_proto(self.scheduling):
            request.resource.scheduling.CopyFrom(
                JobScheduling.to_proto(self.scheduling)
            )
        else:
            request.resource.ClearField("scheduling")
        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDataprocJob(request)
        self.reference = JobReference.from_proto(response.reference)
        self.placement = JobPlacement.from_proto(response.placement)
        self.hadoop_job = JobHadoopJob.from_proto(response.hadoop_job)
        self.spark_job = JobSparkJob.from_proto(response.spark_job)
        self.pyspark_job = JobPysparkJob.from_proto(response.pyspark_job)
        self.hive_job = JobHiveJob.from_proto(response.hive_job)
        self.pig_job = JobPigJob.from_proto(response.pig_job)
        self.spark_r_job = JobSparkRJob.from_proto(response.spark_r_job)
        self.spark_sql_job = JobSparkSqlJob.from_proto(response.spark_sql_job)
        self.presto_job = JobPrestoJob.from_proto(response.presto_job)
        self.status = JobStatus.from_proto(response.status)
        self.status_history = JobStatusHistoryArray.from_proto(response.status_history)
        self.yarn_applications = JobYarnApplicationsArray.from_proto(
            response.yarn_applications
        )
        self.driver_output_resource_uri = Primitive.from_proto(
            response.driver_output_resource_uri
        )
        self.driver_control_files_uri = Primitive.from_proto(
            response.driver_control_files_uri
        )
        self.labels = Primitive.from_proto(response.labels)
        self.scheduling = JobScheduling.from_proto(response.scheduling)
        self.name = Primitive.from_proto(response.name)
        self.done = Primitive.from_proto(response.done)
        self.region = Primitive.from_proto(response.region)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = job_pb2_grpc.DataprocJobServiceStub(channel.Channel())
        request = job_pb2.DeleteDataprocJobRequest()
        request.service_account_file = self.service_account_file
        if JobReference.to_proto(self.reference):
            request.resource.reference.CopyFrom(JobReference.to_proto(self.reference))
        else:
            request.resource.ClearField("reference")
        if JobPlacement.to_proto(self.placement):
            request.resource.placement.CopyFrom(JobPlacement.to_proto(self.placement))
        else:
            request.resource.ClearField("placement")
        if JobHadoopJob.to_proto(self.hadoop_job):
            request.resource.hadoop_job.CopyFrom(JobHadoopJob.to_proto(self.hadoop_job))
        else:
            request.resource.ClearField("hadoop_job")
        if JobSparkJob.to_proto(self.spark_job):
            request.resource.spark_job.CopyFrom(JobSparkJob.to_proto(self.spark_job))
        else:
            request.resource.ClearField("spark_job")
        if JobPysparkJob.to_proto(self.pyspark_job):
            request.resource.pyspark_job.CopyFrom(
                JobPysparkJob.to_proto(self.pyspark_job)
            )
        else:
            request.resource.ClearField("pyspark_job")
        if JobHiveJob.to_proto(self.hive_job):
            request.resource.hive_job.CopyFrom(JobHiveJob.to_proto(self.hive_job))
        else:
            request.resource.ClearField("hive_job")
        if JobPigJob.to_proto(self.pig_job):
            request.resource.pig_job.CopyFrom(JobPigJob.to_proto(self.pig_job))
        else:
            request.resource.ClearField("pig_job")
        if JobSparkRJob.to_proto(self.spark_r_job):
            request.resource.spark_r_job.CopyFrom(
                JobSparkRJob.to_proto(self.spark_r_job)
            )
        else:
            request.resource.ClearField("spark_r_job")
        if JobSparkSqlJob.to_proto(self.spark_sql_job):
            request.resource.spark_sql_job.CopyFrom(
                JobSparkSqlJob.to_proto(self.spark_sql_job)
            )
        else:
            request.resource.ClearField("spark_sql_job")
        if JobPrestoJob.to_proto(self.presto_job):
            request.resource.presto_job.CopyFrom(JobPrestoJob.to_proto(self.presto_job))
        else:
            request.resource.ClearField("presto_job")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if JobScheduling.to_proto(self.scheduling):
            request.resource.scheduling.CopyFrom(
                JobScheduling.to_proto(self.scheduling)
            )
        else:
            request.resource.ClearField("scheduling")
        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteDataprocJob(request)

    @classmethod
    def list(self, project, region, service_account_file=""):
        stub = job_pb2_grpc.DataprocJobServiceStub(channel.Channel())
        request = job_pb2.ListDataprocJobRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        return stub.ListDataprocJob(request).items

    def to_proto(self):
        resource = job_pb2.DataprocJob()
        if JobReference.to_proto(self.reference):
            resource.reference.CopyFrom(JobReference.to_proto(self.reference))
        else:
            resource.ClearField("reference")
        if JobPlacement.to_proto(self.placement):
            resource.placement.CopyFrom(JobPlacement.to_proto(self.placement))
        else:
            resource.ClearField("placement")
        if JobHadoopJob.to_proto(self.hadoop_job):
            resource.hadoop_job.CopyFrom(JobHadoopJob.to_proto(self.hadoop_job))
        else:
            resource.ClearField("hadoop_job")
        if JobSparkJob.to_proto(self.spark_job):
            resource.spark_job.CopyFrom(JobSparkJob.to_proto(self.spark_job))
        else:
            resource.ClearField("spark_job")
        if JobPysparkJob.to_proto(self.pyspark_job):
            resource.pyspark_job.CopyFrom(JobPysparkJob.to_proto(self.pyspark_job))
        else:
            resource.ClearField("pyspark_job")
        if JobHiveJob.to_proto(self.hive_job):
            resource.hive_job.CopyFrom(JobHiveJob.to_proto(self.hive_job))
        else:
            resource.ClearField("hive_job")
        if JobPigJob.to_proto(self.pig_job):
            resource.pig_job.CopyFrom(JobPigJob.to_proto(self.pig_job))
        else:
            resource.ClearField("pig_job")
        if JobSparkRJob.to_proto(self.spark_r_job):
            resource.spark_r_job.CopyFrom(JobSparkRJob.to_proto(self.spark_r_job))
        else:
            resource.ClearField("spark_r_job")
        if JobSparkSqlJob.to_proto(self.spark_sql_job):
            resource.spark_sql_job.CopyFrom(JobSparkSqlJob.to_proto(self.spark_sql_job))
        else:
            resource.ClearField("spark_sql_job")
        if JobPrestoJob.to_proto(self.presto_job):
            resource.presto_job.CopyFrom(JobPrestoJob.to_proto(self.presto_job))
        else:
            resource.ClearField("presto_job")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if JobScheduling.to_proto(self.scheduling):
            resource.scheduling.CopyFrom(JobScheduling.to_proto(self.scheduling))
        else:
            resource.ClearField("scheduling")
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class JobReference(object):
    def __init__(self, project_id: str = None, job_id: str = None):
        self.project_id = project_id
        self.job_id = job_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobReference()
        if Primitive.to_proto(resource.project_id):
            res.project_id = Primitive.to_proto(resource.project_id)
        if Primitive.to_proto(resource.job_id):
            res.job_id = Primitive.to_proto(resource.job_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobReference(
            project_id=Primitive.from_proto(resource.project_id),
            job_id=Primitive.from_proto(resource.job_id),
        )


class JobReferenceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobReference.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobReference.from_proto(i) for i in resources]


class JobPlacement(object):
    def __init__(
        self,
        cluster_name: str = None,
        cluster_uuid: str = None,
        cluster_labels: dict = None,
    ):
        self.cluster_name = cluster_name
        self.cluster_uuid = cluster_uuid
        self.cluster_labels = cluster_labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobPlacement()
        if Primitive.to_proto(resource.cluster_name):
            res.cluster_name = Primitive.to_proto(resource.cluster_name)
        if Primitive.to_proto(resource.cluster_uuid):
            res.cluster_uuid = Primitive.to_proto(resource.cluster_uuid)
        if Primitive.to_proto(resource.cluster_labels):
            res.cluster_labels = Primitive.to_proto(resource.cluster_labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobPlacement(
            cluster_name=Primitive.from_proto(resource.cluster_name),
            cluster_uuid=Primitive.from_proto(resource.cluster_uuid),
            cluster_labels=Primitive.from_proto(resource.cluster_labels),
        )


class JobPlacementArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobPlacement.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobPlacement.from_proto(i) for i in resources]


class JobHadoopJob(object):
    def __init__(
        self,
        main_jar_file_uri: str = None,
        main_class: str = None,
        args: list = None,
        jar_file_uris: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_jar_file_uri = main_jar_file_uri
        self.main_class = main_class
        self.args = args
        self.jar_file_uris = jar_file_uris
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobHadoopJob()
        if Primitive.to_proto(resource.main_jar_file_uri):
            res.main_jar_file_uri = Primitive.to_proto(resource.main_jar_file_uri)
        if Primitive.to_proto(resource.main_class):
            res.main_class = Primitive.to_proto(resource.main_class)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if JobHadoopJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                JobHadoopJobLoggingConfig.to_proto(resource.logging_config)
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobHadoopJob(
            main_jar_file_uri=Primitive.from_proto(resource.main_jar_file_uri),
            main_class=Primitive.from_proto(resource.main_class),
            args=Primitive.from_proto(resource.args),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
            file_uris=Primitive.from_proto(resource.file_uris),
            archive_uris=Primitive.from_proto(resource.archive_uris),
            properties=Primitive.from_proto(resource.properties),
            logging_config=JobHadoopJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class JobHadoopJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobHadoopJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobHadoopJob.from_proto(i) for i in resources]


class JobHadoopJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobHadoopJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobHadoopJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class JobHadoopJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobHadoopJobLoggingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobHadoopJobLoggingConfig.from_proto(i) for i in resources]


class JobSparkJob(object):
    def __init__(
        self,
        main_jar_file_uri: str = None,
        main_class: str = None,
        args: list = None,
        jar_file_uris: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_jar_file_uri = main_jar_file_uri
        self.main_class = main_class
        self.args = args
        self.jar_file_uris = jar_file_uris
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobSparkJob()
        if Primitive.to_proto(resource.main_jar_file_uri):
            res.main_jar_file_uri = Primitive.to_proto(resource.main_jar_file_uri)
        if Primitive.to_proto(resource.main_class):
            res.main_class = Primitive.to_proto(resource.main_class)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if JobSparkJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                JobSparkJobLoggingConfig.to_proto(resource.logging_config)
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobSparkJob(
            main_jar_file_uri=Primitive.from_proto(resource.main_jar_file_uri),
            main_class=Primitive.from_proto(resource.main_class),
            args=Primitive.from_proto(resource.args),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
            file_uris=Primitive.from_proto(resource.file_uris),
            archive_uris=Primitive.from_proto(resource.archive_uris),
            properties=Primitive.from_proto(resource.properties),
            logging_config=JobSparkJobLoggingConfig.from_proto(resource.logging_config),
        )


class JobSparkJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobSparkJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobSparkJob.from_proto(i) for i in resources]


class JobSparkJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobSparkJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobSparkJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class JobSparkJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobSparkJobLoggingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobSparkJobLoggingConfig.from_proto(i) for i in resources]


class JobPysparkJob(object):
    def __init__(
        self,
        main_python_file_uri: str = None,
        args: list = None,
        python_file_uris: list = None,
        jar_file_uris: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_python_file_uri = main_python_file_uri
        self.args = args
        self.python_file_uris = python_file_uris
        self.jar_file_uris = jar_file_uris
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobPysparkJob()
        if Primitive.to_proto(resource.main_python_file_uri):
            res.main_python_file_uri = Primitive.to_proto(resource.main_python_file_uri)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.python_file_uris):
            res.python_file_uris.extend(Primitive.to_proto(resource.python_file_uris))
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if JobPysparkJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                JobPysparkJobLoggingConfig.to_proto(resource.logging_config)
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobPysparkJob(
            main_python_file_uri=Primitive.from_proto(resource.main_python_file_uri),
            args=Primitive.from_proto(resource.args),
            python_file_uris=Primitive.from_proto(resource.python_file_uris),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
            file_uris=Primitive.from_proto(resource.file_uris),
            archive_uris=Primitive.from_proto(resource.archive_uris),
            properties=Primitive.from_proto(resource.properties),
            logging_config=JobPysparkJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class JobPysparkJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobPysparkJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobPysparkJob.from_proto(i) for i in resources]


class JobPysparkJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobPysparkJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobPysparkJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class JobPysparkJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobPysparkJobLoggingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobPysparkJobLoggingConfig.from_proto(i) for i in resources]


class JobHiveJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        continue_on_failure: bool = None,
        script_variables: dict = None,
        properties: dict = None,
        jar_file_uris: list = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.continue_on_failure = continue_on_failure
        self.script_variables = script_variables
        self.properties = properties
        self.jar_file_uris = jar_file_uris

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobHiveJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if JobHiveJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(JobHiveJobQueryList.to_proto(resource.query_list))
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.continue_on_failure):
            res.continue_on_failure = Primitive.to_proto(resource.continue_on_failure)
        if Primitive.to_proto(resource.script_variables):
            res.script_variables = Primitive.to_proto(resource.script_variables)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobHiveJob(
            query_file_uri=Primitive.from_proto(resource.query_file_uri),
            query_list=JobHiveJobQueryList.from_proto(resource.query_list),
            continue_on_failure=Primitive.from_proto(resource.continue_on_failure),
            script_variables=Primitive.from_proto(resource.script_variables),
            properties=Primitive.from_proto(resource.properties),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
        )


class JobHiveJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobHiveJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobHiveJob.from_proto(i) for i in resources]


class JobHiveJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobHiveJobQueryList()
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobHiveJobQueryList(queries=Primitive.from_proto(resource.queries),)


class JobHiveJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobHiveJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobHiveJobQueryList.from_proto(i) for i in resources]


class JobPigJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        continue_on_failure: bool = None,
        script_variables: dict = None,
        properties: dict = None,
        jar_file_uris: list = None,
        logging_config: dict = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.continue_on_failure = continue_on_failure
        self.script_variables = script_variables
        self.properties = properties
        self.jar_file_uris = jar_file_uris
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobPigJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if JobPigJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(JobPigJobQueryList.to_proto(resource.query_list))
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.continue_on_failure):
            res.continue_on_failure = Primitive.to_proto(resource.continue_on_failure)
        if Primitive.to_proto(resource.script_variables):
            res.script_variables = Primitive.to_proto(resource.script_variables)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if JobPigJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                JobPigJobLoggingConfig.to_proto(resource.logging_config)
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobPigJob(
            query_file_uri=Primitive.from_proto(resource.query_file_uri),
            query_list=JobPigJobQueryList.from_proto(resource.query_list),
            continue_on_failure=Primitive.from_proto(resource.continue_on_failure),
            script_variables=Primitive.from_proto(resource.script_variables),
            properties=Primitive.from_proto(resource.properties),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
            logging_config=JobPigJobLoggingConfig.from_proto(resource.logging_config),
        )


class JobPigJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobPigJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobPigJob.from_proto(i) for i in resources]


class JobPigJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobPigJobQueryList()
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobPigJobQueryList(queries=Primitive.from_proto(resource.queries),)


class JobPigJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobPigJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobPigJobQueryList.from_proto(i) for i in resources]


class JobPigJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobPigJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobPigJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class JobPigJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobPigJobLoggingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobPigJobLoggingConfig.from_proto(i) for i in resources]


class JobSparkRJob(object):
    def __init__(
        self,
        main_r_file_uri: str = None,
        args: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_r_file_uri = main_r_file_uri
        self.args = args
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobSparkRJob()
        if Primitive.to_proto(resource.main_r_file_uri):
            res.main_r_file_uri = Primitive.to_proto(resource.main_r_file_uri)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if JobSparkRJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                JobSparkRJobLoggingConfig.to_proto(resource.logging_config)
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobSparkRJob(
            main_r_file_uri=Primitive.from_proto(resource.main_r_file_uri),
            args=Primitive.from_proto(resource.args),
            file_uris=Primitive.from_proto(resource.file_uris),
            archive_uris=Primitive.from_proto(resource.archive_uris),
            properties=Primitive.from_proto(resource.properties),
            logging_config=JobSparkRJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class JobSparkRJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobSparkRJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobSparkRJob.from_proto(i) for i in resources]


class JobSparkRJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobSparkRJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobSparkRJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class JobSparkRJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobSparkRJobLoggingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobSparkRJobLoggingConfig.from_proto(i) for i in resources]


class JobSparkSqlJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        script_variables: dict = None,
        properties: dict = None,
        jar_file_uris: list = None,
        logging_config: dict = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.script_variables = script_variables
        self.properties = properties
        self.jar_file_uris = jar_file_uris
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobSparkSqlJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if JobSparkSqlJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(
                JobSparkSqlJobQueryList.to_proto(resource.query_list)
            )
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.script_variables):
            res.script_variables = Primitive.to_proto(resource.script_variables)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if JobSparkSqlJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                JobSparkSqlJobLoggingConfig.to_proto(resource.logging_config)
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobSparkSqlJob(
            query_file_uri=Primitive.from_proto(resource.query_file_uri),
            query_list=JobSparkSqlJobQueryList.from_proto(resource.query_list),
            script_variables=Primitive.from_proto(resource.script_variables),
            properties=Primitive.from_proto(resource.properties),
            jar_file_uris=Primitive.from_proto(resource.jar_file_uris),
            logging_config=JobSparkSqlJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class JobSparkSqlJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobSparkSqlJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobSparkSqlJob.from_proto(i) for i in resources]


class JobSparkSqlJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobSparkSqlJobQueryList()
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobSparkSqlJobQueryList(queries=Primitive.from_proto(resource.queries),)


class JobSparkSqlJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobSparkSqlJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobSparkSqlJobQueryList.from_proto(i) for i in resources]


class JobSparkSqlJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobSparkSqlJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobSparkSqlJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class JobSparkSqlJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobSparkSqlJobLoggingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobSparkSqlJobLoggingConfig.from_proto(i) for i in resources]


class JobPrestoJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        continue_on_failure: bool = None,
        output_format: str = None,
        client_tags: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.continue_on_failure = continue_on_failure
        self.output_format = output_format
        self.client_tags = client_tags
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobPrestoJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if JobPrestoJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(JobPrestoJobQueryList.to_proto(resource.query_list))
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.continue_on_failure):
            res.continue_on_failure = Primitive.to_proto(resource.continue_on_failure)
        if Primitive.to_proto(resource.output_format):
            res.output_format = Primitive.to_proto(resource.output_format)
        if Primitive.to_proto(resource.client_tags):
            res.client_tags.extend(Primitive.to_proto(resource.client_tags))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if JobPrestoJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                JobPrestoJobLoggingConfig.to_proto(resource.logging_config)
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobPrestoJob(
            query_file_uri=Primitive.from_proto(resource.query_file_uri),
            query_list=JobPrestoJobQueryList.from_proto(resource.query_list),
            continue_on_failure=Primitive.from_proto(resource.continue_on_failure),
            output_format=Primitive.from_proto(resource.output_format),
            client_tags=Primitive.from_proto(resource.client_tags),
            properties=Primitive.from_proto(resource.properties),
            logging_config=JobPrestoJobLoggingConfig.from_proto(
                resource.logging_config
            ),
        )


class JobPrestoJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobPrestoJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobPrestoJob.from_proto(i) for i in resources]


class JobPrestoJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobPrestoJobQueryList()
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobPrestoJobQueryList(queries=Primitive.from_proto(resource.queries),)


class JobPrestoJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobPrestoJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobPrestoJobQueryList.from_proto(i) for i in resources]


class JobPrestoJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobPrestoJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobPrestoJobLoggingConfig(
            driver_log_levels=Primitive.from_proto(resource.driver_log_levels),
        )


class JobPrestoJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobPrestoJobLoggingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobPrestoJobLoggingConfig.from_proto(i) for i in resources]


class JobStatus(object):
    def __init__(
        self,
        state: str = None,
        details: str = None,
        state_start_time: str = None,
        substate: str = None,
    ):
        self.state = state
        self.details = details
        self.state_start_time = state_start_time
        self.substate = substate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobStatus()
        if JobStatusStateEnum.to_proto(resource.state):
            res.state = JobStatusStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.details):
            res.details = Primitive.to_proto(resource.details)
        if Primitive.to_proto(resource.state_start_time):
            res.state_start_time = Primitive.to_proto(resource.state_start_time)
        if JobStatusSubstateEnum.to_proto(resource.substate):
            res.substate = JobStatusSubstateEnum.to_proto(resource.substate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobStatus(
            state=JobStatusStateEnum.from_proto(resource.state),
            details=Primitive.from_proto(resource.details),
            state_start_time=Primitive.from_proto(resource.state_start_time),
            substate=JobStatusSubstateEnum.from_proto(resource.substate),
        )


class JobStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobStatus.from_proto(i) for i in resources]


class JobStatusHistory(object):
    def __init__(
        self,
        state: str = None,
        details: str = None,
        state_start_time: str = None,
        substate: str = None,
    ):
        self.state = state
        self.details = details
        self.state_start_time = state_start_time
        self.substate = substate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobStatusHistory()
        if JobStatusHistoryStateEnum.to_proto(resource.state):
            res.state = JobStatusHistoryStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.details):
            res.details = Primitive.to_proto(resource.details)
        if Primitive.to_proto(resource.state_start_time):
            res.state_start_time = Primitive.to_proto(resource.state_start_time)
        if JobStatusHistorySubstateEnum.to_proto(resource.substate):
            res.substate = JobStatusHistorySubstateEnum.to_proto(resource.substate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobStatusHistory(
            state=JobStatusHistoryStateEnum.from_proto(resource.state),
            details=Primitive.from_proto(resource.details),
            state_start_time=Primitive.from_proto(resource.state_start_time),
            substate=JobStatusHistorySubstateEnum.from_proto(resource.substate),
        )


class JobStatusHistoryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobStatusHistory.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobStatusHistory.from_proto(i) for i in resources]


class JobYarnApplications(object):
    def __init__(
        self,
        name: str = None,
        state: str = None,
        progress: float = None,
        tracking_url: str = None,
    ):
        self.name = name
        self.state = state
        self.progress = progress
        self.tracking_url = tracking_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobYarnApplications()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if JobYarnApplicationsStateEnum.to_proto(resource.state):
            res.state = JobYarnApplicationsStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.progress):
            res.progress = Primitive.to_proto(resource.progress)
        if Primitive.to_proto(resource.tracking_url):
            res.tracking_url = Primitive.to_proto(resource.tracking_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobYarnApplications(
            name=Primitive.from_proto(resource.name),
            state=JobYarnApplicationsStateEnum.from_proto(resource.state),
            progress=Primitive.from_proto(resource.progress),
            tracking_url=Primitive.from_proto(resource.tracking_url),
        )


class JobYarnApplicationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobYarnApplications.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobYarnApplications.from_proto(i) for i in resources]


class JobScheduling(object):
    def __init__(
        self, max_failures_per_hour: int = None, max_failures_total: int = None
    ):
        self.max_failures_per_hour = max_failures_per_hour
        self.max_failures_total = max_failures_total

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.DataprocJobScheduling()
        if Primitive.to_proto(resource.max_failures_per_hour):
            res.max_failures_per_hour = Primitive.to_proto(
                resource.max_failures_per_hour
            )
        if Primitive.to_proto(resource.max_failures_total):
            res.max_failures_total = Primitive.to_proto(resource.max_failures_total)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobScheduling(
            max_failures_per_hour=Primitive.from_proto(resource.max_failures_per_hour),
            max_failures_total=Primitive.from_proto(resource.max_failures_total),
        )


class JobSchedulingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobScheduling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobScheduling.from_proto(i) for i in resources]


class JobStatusStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.DataprocJobStatusStateEnum.Value(
            "DataprocJobStatusStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.DataprocJobStatusStateEnum.Name(resource)[
            len("DataprocJobStatusStateEnum") :
        ]


class JobStatusSubstateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.DataprocJobStatusSubstateEnum.Value(
            "DataprocJobStatusSubstateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.DataprocJobStatusSubstateEnum.Name(resource)[
            len("DataprocJobStatusSubstateEnum") :
        ]


class JobStatusHistoryStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.DataprocJobStatusHistoryStateEnum.Value(
            "DataprocJobStatusHistoryStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.DataprocJobStatusHistoryStateEnum.Name(resource)[
            len("DataprocJobStatusHistoryStateEnum") :
        ]


class JobStatusHistorySubstateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.DataprocJobStatusHistorySubstateEnum.Value(
            "DataprocJobStatusHistorySubstateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.DataprocJobStatusHistorySubstateEnum.Name(resource)[
            len("DataprocJobStatusHistorySubstateEnum") :
        ]


class JobYarnApplicationsStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.DataprocJobYarnApplicationsStateEnum.Value(
            "DataprocJobYarnApplicationsStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.DataprocJobYarnApplicationsStateEnum.Name(resource)[
            len("DataprocJobYarnApplicationsStateEnum") :
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
