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
from google3.cloud.graphite.mmv2.services.google.run import job_pb2
from google3.cloud.graphite.mmv2.services.google.run import job_pb2_grpc

from typing import List


class Job(object):
    def __init__(
        self,
        name: str = None,
        uid: str = None,
        generation: int = None,
        labels: dict = None,
        annotations: dict = None,
        create_time: str = None,
        update_time: str = None,
        delete_time: str = None,
        expire_time: str = None,
        creator: str = None,
        last_modifier: str = None,
        client: str = None,
        client_version: str = None,
        launch_stage: str = None,
        binary_authorization: dict = None,
        template: dict = None,
        observed_generation: int = None,
        terminal_condition: dict = None,
        conditions: list = None,
        execution_count: int = None,
        latest_succeeded_execution: dict = None,
        latest_created_execution: dict = None,
        reconciling: bool = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.annotations = annotations
        self.client = client
        self.client_version = client_version
        self.launch_stage = launch_stage
        self.binary_authorization = binary_authorization
        self.template = template
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = job_pb2_grpc.RunAlphaJobServiceStub(channel.Channel())
        request = job_pb2.ApplyRunAlphaJobRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.client):
            request.resource.client = Primitive.to_proto(self.client)

        if Primitive.to_proto(self.client_version):
            request.resource.client_version = Primitive.to_proto(self.client_version)

        if JobLaunchStageEnum.to_proto(self.launch_stage):
            request.resource.launch_stage = JobLaunchStageEnum.to_proto(
                self.launch_stage
            )

        if JobBinaryAuthorization.to_proto(self.binary_authorization):
            request.resource.binary_authorization.CopyFrom(
                JobBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            request.resource.ClearField("binary_authorization")
        if JobTemplate.to_proto(self.template):
            request.resource.template.CopyFrom(JobTemplate.to_proto(self.template))
        else:
            request.resource.ClearField("template")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyRunAlphaJob(request)
        self.name = Primitive.from_proto(response.name)
        self.uid = Primitive.from_proto(response.uid)
        self.generation = Primitive.from_proto(response.generation)
        self.labels = Primitive.from_proto(response.labels)
        self.annotations = Primitive.from_proto(response.annotations)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.expire_time = Primitive.from_proto(response.expire_time)
        self.creator = Primitive.from_proto(response.creator)
        self.last_modifier = Primitive.from_proto(response.last_modifier)
        self.client = Primitive.from_proto(response.client)
        self.client_version = Primitive.from_proto(response.client_version)
        self.launch_stage = JobLaunchStageEnum.from_proto(response.launch_stage)
        self.binary_authorization = JobBinaryAuthorization.from_proto(
            response.binary_authorization
        )
        self.template = JobTemplate.from_proto(response.template)
        self.observed_generation = Primitive.from_proto(response.observed_generation)
        self.terminal_condition = JobTerminalCondition.from_proto(
            response.terminal_condition
        )
        self.conditions = JobConditionsArray.from_proto(response.conditions)
        self.execution_count = Primitive.from_proto(response.execution_count)
        self.latest_succeeded_execution = JobLatestSucceededExecution.from_proto(
            response.latest_succeeded_execution
        )
        self.latest_created_execution = JobLatestCreatedExecution.from_proto(
            response.latest_created_execution
        )
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = job_pb2_grpc.RunAlphaJobServiceStub(channel.Channel())
        request = job_pb2.DeleteRunAlphaJobRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.client):
            request.resource.client = Primitive.to_proto(self.client)

        if Primitive.to_proto(self.client_version):
            request.resource.client_version = Primitive.to_proto(self.client_version)

        if JobLaunchStageEnum.to_proto(self.launch_stage):
            request.resource.launch_stage = JobLaunchStageEnum.to_proto(
                self.launch_stage
            )

        if JobBinaryAuthorization.to_proto(self.binary_authorization):
            request.resource.binary_authorization.CopyFrom(
                JobBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            request.resource.ClearField("binary_authorization")
        if JobTemplate.to_proto(self.template):
            request.resource.template.CopyFrom(JobTemplate.to_proto(self.template))
        else:
            request.resource.ClearField("template")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteRunAlphaJob(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = job_pb2_grpc.RunAlphaJobServiceStub(channel.Channel())
        request = job_pb2.ListRunAlphaJobRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListRunAlphaJob(request).items

    def to_proto(self):
        resource = job_pb2.RunAlphaJob()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.client):
            resource.client = Primitive.to_proto(self.client)
        if Primitive.to_proto(self.client_version):
            resource.client_version = Primitive.to_proto(self.client_version)
        if JobLaunchStageEnum.to_proto(self.launch_stage):
            resource.launch_stage = JobLaunchStageEnum.to_proto(self.launch_stage)
        if JobBinaryAuthorization.to_proto(self.binary_authorization):
            resource.binary_authorization.CopyFrom(
                JobBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            resource.ClearField("binary_authorization")
        if JobTemplate.to_proto(self.template):
            resource.template.CopyFrom(JobTemplate.to_proto(self.template))
        else:
            resource.ClearField("template")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class JobBinaryAuthorization(object):
    def __init__(self, use_default: bool = None, breakglass_justification: str = None):
        self.use_default = use_default
        self.breakglass_justification = breakglass_justification

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobBinaryAuthorization()
        if Primitive.to_proto(resource.use_default):
            res.use_default = Primitive.to_proto(resource.use_default)
        if Primitive.to_proto(resource.breakglass_justification):
            res.breakglass_justification = Primitive.to_proto(
                resource.breakglass_justification
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobBinaryAuthorization(
            use_default=Primitive.from_proto(resource.use_default),
            breakglass_justification=Primitive.from_proto(
                resource.breakglass_justification
            ),
        )


class JobBinaryAuthorizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobBinaryAuthorization.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobBinaryAuthorization.from_proto(i) for i in resources]


class JobTemplate(object):
    def __init__(
        self,
        labels: dict = None,
        annotations: dict = None,
        parallelism: int = None,
        task_count: int = None,
        template: dict = None,
    ):
        self.labels = labels
        self.annotations = annotations
        self.parallelism = parallelism
        self.task_count = task_count
        self.template = template

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplate()
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if Primitive.to_proto(resource.annotations):
            res.annotations = Primitive.to_proto(resource.annotations)
        if Primitive.to_proto(resource.parallelism):
            res.parallelism = Primitive.to_proto(resource.parallelism)
        if Primitive.to_proto(resource.task_count):
            res.task_count = Primitive.to_proto(resource.task_count)
        if JobTemplateTemplate.to_proto(resource.template):
            res.template.CopyFrom(JobTemplateTemplate.to_proto(resource.template))
        else:
            res.ClearField("template")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplate(
            labels=Primitive.from_proto(resource.labels),
            annotations=Primitive.from_proto(resource.annotations),
            parallelism=Primitive.from_proto(resource.parallelism),
            task_count=Primitive.from_proto(resource.task_count),
            template=JobTemplateTemplate.from_proto(resource.template),
        )


class JobTemplateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTemplate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTemplate.from_proto(i) for i in resources]


class JobTemplateTemplate(object):
    def __init__(
        self,
        containers: list = None,
        volumes: list = None,
        max_retries: int = None,
        timeout: str = None,
        service_account: str = None,
        execution_environment: str = None,
        encryption_key: str = None,
        vpc_access: dict = None,
    ):
        self.containers = containers
        self.volumes = volumes
        self.max_retries = max_retries
        self.timeout = timeout
        self.service_account = service_account
        self.execution_environment = execution_environment
        self.encryption_key = encryption_key
        self.vpc_access = vpc_access

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplate()
        if JobTemplateTemplateContainersArray.to_proto(resource.containers):
            res.containers.extend(
                JobTemplateTemplateContainersArray.to_proto(resource.containers)
            )
        if JobTemplateTemplateVolumesArray.to_proto(resource.volumes):
            res.volumes.extend(
                JobTemplateTemplateVolumesArray.to_proto(resource.volumes)
            )
        if Primitive.to_proto(resource.max_retries):
            res.max_retries = Primitive.to_proto(resource.max_retries)
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        if Primitive.to_proto(resource.service_account):
            res.service_account = Primitive.to_proto(resource.service_account)
        if JobTemplateTemplateExecutionEnvironmentEnum.to_proto(
            resource.execution_environment
        ):
            res.execution_environment = (
                JobTemplateTemplateExecutionEnvironmentEnum.to_proto(
                    resource.execution_environment
                )
            )
        if Primitive.to_proto(resource.encryption_key):
            res.encryption_key = Primitive.to_proto(resource.encryption_key)
        if JobTemplateTemplateVPCAccess.to_proto(resource.vpc_access):
            res.vpc_access.CopyFrom(
                JobTemplateTemplateVPCAccess.to_proto(resource.vpc_access)
            )
        else:
            res.ClearField("vpc_access")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplate(
            containers=JobTemplateTemplateContainersArray.from_proto(
                resource.containers
            ),
            volumes=JobTemplateTemplateVolumesArray.from_proto(resource.volumes),
            max_retries=Primitive.from_proto(resource.max_retries),
            timeout=Primitive.from_proto(resource.timeout),
            service_account=Primitive.from_proto(resource.service_account),
            execution_environment=JobTemplateTemplateExecutionEnvironmentEnum.from_proto(
                resource.execution_environment
            ),
            encryption_key=Primitive.from_proto(resource.encryption_key),
            vpc_access=JobTemplateTemplateVPCAccess.from_proto(resource.vpc_access),
        )


class JobTemplateTemplateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTemplateTemplate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTemplateTemplate.from_proto(i) for i in resources]


class JobTemplateTemplateContainers(object):
    def __init__(
        self,
        name: str = None,
        image: str = None,
        command: list = None,
        args: list = None,
        env: list = None,
        resources: dict = None,
        ports: list = None,
        volume_mounts: list = None,
    ):
        self.name = name
        self.image = image
        self.command = command
        self.args = args
        self.env = env
        self.resources = resources
        self.ports = ports
        self.volume_mounts = volume_mounts

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateContainers()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        if Primitive.to_proto(resource.command):
            res.command.extend(Primitive.to_proto(resource.command))
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if JobTemplateTemplateContainersEnvArray.to_proto(resource.env):
            res.env.extend(JobTemplateTemplateContainersEnvArray.to_proto(resource.env))
        if JobTemplateTemplateContainersResources.to_proto(resource.resources):
            res.resources.CopyFrom(
                JobTemplateTemplateContainersResources.to_proto(resource.resources)
            )
        else:
            res.ClearField("resources")
        if JobTemplateTemplateContainersPortsArray.to_proto(resource.ports):
            res.ports.extend(
                JobTemplateTemplateContainersPortsArray.to_proto(resource.ports)
            )
        if JobTemplateTemplateContainersVolumeMountsArray.to_proto(
            resource.volume_mounts
        ):
            res.volume_mounts.extend(
                JobTemplateTemplateContainersVolumeMountsArray.to_proto(
                    resource.volume_mounts
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateContainers(
            name=Primitive.from_proto(resource.name),
            image=Primitive.from_proto(resource.image),
            command=Primitive.from_proto(resource.command),
            args=Primitive.from_proto(resource.args),
            env=JobTemplateTemplateContainersEnvArray.from_proto(resource.env),
            resources=JobTemplateTemplateContainersResources.from_proto(
                resource.resources
            ),
            ports=JobTemplateTemplateContainersPortsArray.from_proto(resource.ports),
            volume_mounts=JobTemplateTemplateContainersVolumeMountsArray.from_proto(
                resource.volume_mounts
            ),
        )


class JobTemplateTemplateContainersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTemplateTemplateContainers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTemplateTemplateContainers.from_proto(i) for i in resources]


class JobTemplateTemplateContainersEnv(object):
    def __init__(self, name: str = None, value: str = None, value_source: dict = None):
        self.name = name
        self.value = value
        self.value_source = value_source

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateContainersEnv()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if JobTemplateTemplateContainersEnvValueSource.to_proto(resource.value_source):
            res.value_source.CopyFrom(
                JobTemplateTemplateContainersEnvValueSource.to_proto(
                    resource.value_source
                )
            )
        else:
            res.ClearField("value_source")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateContainersEnv(
            name=Primitive.from_proto(resource.name),
            value=Primitive.from_proto(resource.value),
            value_source=JobTemplateTemplateContainersEnvValueSource.from_proto(
                resource.value_source
            ),
        )


class JobTemplateTemplateContainersEnvArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTemplateTemplateContainersEnv.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTemplateTemplateContainersEnv.from_proto(i) for i in resources]


class JobTemplateTemplateContainersEnvValueSource(object):
    def __init__(self, secret_key_ref: dict = None):
        self.secret_key_ref = secret_key_ref

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateContainersEnvValueSource()
        if JobTemplateTemplateContainersEnvValueSourceSecretKeyRef.to_proto(
            resource.secret_key_ref
        ):
            res.secret_key_ref.CopyFrom(
                JobTemplateTemplateContainersEnvValueSourceSecretKeyRef.to_proto(
                    resource.secret_key_ref
                )
            )
        else:
            res.ClearField("secret_key_ref")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateContainersEnvValueSource(
            secret_key_ref=JobTemplateTemplateContainersEnvValueSourceSecretKeyRef.from_proto(
                resource.secret_key_ref
            ),
        )


class JobTemplateTemplateContainersEnvValueSourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTemplateTemplateContainersEnvValueSource.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTemplateTemplateContainersEnvValueSource.from_proto(i) for i in resources
        ]


class JobTemplateTemplateContainersEnvValueSourceSecretKeyRef(object):
    def __init__(self, secret: str = None, version: str = None):
        self.secret = secret
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateContainersEnvValueSourceSecretKeyRef()
        if Primitive.to_proto(resource.secret):
            res.secret = Primitive.to_proto(resource.secret)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateContainersEnvValueSourceSecretKeyRef(
            secret=Primitive.from_proto(resource.secret),
            version=Primitive.from_proto(resource.version),
        )


class JobTemplateTemplateContainersEnvValueSourceSecretKeyRefArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTemplateTemplateContainersEnvValueSourceSecretKeyRef.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTemplateTemplateContainersEnvValueSourceSecretKeyRef.from_proto(i)
            for i in resources
        ]


class JobTemplateTemplateContainersResources(object):
    def __init__(self, limits: dict = None, cpu_idle: bool = None):
        self.limits = limits
        self.cpu_idle = cpu_idle

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateContainersResources()
        if Primitive.to_proto(resource.limits):
            res.limits = Primitive.to_proto(resource.limits)
        if Primitive.to_proto(resource.cpu_idle):
            res.cpu_idle = Primitive.to_proto(resource.cpu_idle)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateContainersResources(
            limits=Primitive.from_proto(resource.limits),
            cpu_idle=Primitive.from_proto(resource.cpu_idle),
        )


class JobTemplateTemplateContainersResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTemplateTemplateContainersResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTemplateTemplateContainersResources.from_proto(i) for i in resources]


class JobTemplateTemplateContainersPorts(object):
    def __init__(self, name: str = None, container_port: int = None):
        self.name = name
        self.container_port = container_port

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateContainersPorts()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.container_port):
            res.container_port = Primitive.to_proto(resource.container_port)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateContainersPorts(
            name=Primitive.from_proto(resource.name),
            container_port=Primitive.from_proto(resource.container_port),
        )


class JobTemplateTemplateContainersPortsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTemplateTemplateContainersPorts.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTemplateTemplateContainersPorts.from_proto(i) for i in resources]


class JobTemplateTemplateContainersVolumeMounts(object):
    def __init__(self, name: str = None, mount_path: str = None):
        self.name = name
        self.mount_path = mount_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateContainersVolumeMounts()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.mount_path):
            res.mount_path = Primitive.to_proto(resource.mount_path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateContainersVolumeMounts(
            name=Primitive.from_proto(resource.name),
            mount_path=Primitive.from_proto(resource.mount_path),
        )


class JobTemplateTemplateContainersVolumeMountsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTemplateTemplateContainersVolumeMounts.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTemplateTemplateContainersVolumeMounts.from_proto(i) for i in resources
        ]


class JobTemplateTemplateVolumes(object):
    def __init__(
        self, name: str = None, secret: dict = None, cloud_sql_instance: dict = None
    ):
        self.name = name
        self.secret = secret
        self.cloud_sql_instance = cloud_sql_instance

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateVolumes()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if JobTemplateTemplateVolumesSecret.to_proto(resource.secret):
            res.secret.CopyFrom(
                JobTemplateTemplateVolumesSecret.to_proto(resource.secret)
            )
        else:
            res.ClearField("secret")
        if JobTemplateTemplateVolumesCloudSqlInstance.to_proto(
            resource.cloud_sql_instance
        ):
            res.cloud_sql_instance.CopyFrom(
                JobTemplateTemplateVolumesCloudSqlInstance.to_proto(
                    resource.cloud_sql_instance
                )
            )
        else:
            res.ClearField("cloud_sql_instance")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateVolumes(
            name=Primitive.from_proto(resource.name),
            secret=JobTemplateTemplateVolumesSecret.from_proto(resource.secret),
            cloud_sql_instance=JobTemplateTemplateVolumesCloudSqlInstance.from_proto(
                resource.cloud_sql_instance
            ),
        )


class JobTemplateTemplateVolumesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTemplateTemplateVolumes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTemplateTemplateVolumes.from_proto(i) for i in resources]


class JobTemplateTemplateVolumesSecret(object):
    def __init__(
        self, secret: str = None, items: list = None, default_mode: int = None
    ):
        self.secret = secret
        self.items = items
        self.default_mode = default_mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateVolumesSecret()
        if Primitive.to_proto(resource.secret):
            res.secret = Primitive.to_proto(resource.secret)
        if JobTemplateTemplateVolumesSecretItemsArray.to_proto(resource.items):
            res.items.extend(
                JobTemplateTemplateVolumesSecretItemsArray.to_proto(resource.items)
            )
        if Primitive.to_proto(resource.default_mode):
            res.default_mode = Primitive.to_proto(resource.default_mode)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateVolumesSecret(
            secret=Primitive.from_proto(resource.secret),
            items=JobTemplateTemplateVolumesSecretItemsArray.from_proto(resource.items),
            default_mode=Primitive.from_proto(resource.default_mode),
        )


class JobTemplateTemplateVolumesSecretArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTemplateTemplateVolumesSecret.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTemplateTemplateVolumesSecret.from_proto(i) for i in resources]


class JobTemplateTemplateVolumesSecretItems(object):
    def __init__(self, path: str = None, version: str = None, mode: int = None):
        self.path = path
        self.version = version
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateVolumesSecretItems()
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        if Primitive.to_proto(resource.mode):
            res.mode = Primitive.to_proto(resource.mode)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateVolumesSecretItems(
            path=Primitive.from_proto(resource.path),
            version=Primitive.from_proto(resource.version),
            mode=Primitive.from_proto(resource.mode),
        )


class JobTemplateTemplateVolumesSecretItemsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTemplateTemplateVolumesSecretItems.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTemplateTemplateVolumesSecretItems.from_proto(i) for i in resources]


class JobTemplateTemplateVolumesCloudSqlInstance(object):
    def __init__(self, instances: list = None):
        self.instances = instances

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateVolumesCloudSqlInstance()
        if Primitive.to_proto(resource.instances):
            res.instances.extend(Primitive.to_proto(resource.instances))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateVolumesCloudSqlInstance(
            instances=Primitive.from_proto(resource.instances),
        )


class JobTemplateTemplateVolumesCloudSqlInstanceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            JobTemplateTemplateVolumesCloudSqlInstance.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            JobTemplateTemplateVolumesCloudSqlInstance.from_proto(i) for i in resources
        ]


class JobTemplateTemplateVPCAccess(object):
    def __init__(self, connector: str = None, egress: str = None):
        self.connector = connector
        self.egress = egress

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTemplateTemplateVPCAccess()
        if Primitive.to_proto(resource.connector):
            res.connector = Primitive.to_proto(resource.connector)
        if JobTemplateTemplateVPCAccessEgressEnum.to_proto(resource.egress):
            res.egress = JobTemplateTemplateVPCAccessEgressEnum.to_proto(
                resource.egress
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTemplateTemplateVPCAccess(
            connector=Primitive.from_proto(resource.connector),
            egress=JobTemplateTemplateVPCAccessEgressEnum.from_proto(resource.egress),
        )


class JobTemplateTemplateVPCAccessArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTemplateTemplateVPCAccess.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTemplateTemplateVPCAccess.from_proto(i) for i in resources]


class JobTerminalCondition(object):
    def __init__(
        self,
        type: str = None,
        state: str = None,
        message: str = None,
        last_transition_time: str = None,
        severity: str = None,
        reason: str = None,
        internal_reason: str = None,
        domain_mapping_reason: str = None,
        revision_reason: str = None,
        execution_reason: str = None,
    ):
        self.type = type
        self.state = state
        self.message = message
        self.last_transition_time = last_transition_time
        self.severity = severity
        self.reason = reason
        self.internal_reason = internal_reason
        self.domain_mapping_reason = domain_mapping_reason
        self.revision_reason = revision_reason
        self.execution_reason = execution_reason

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobTerminalCondition()
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if JobTerminalConditionStateEnum.to_proto(resource.state):
            res.state = JobTerminalConditionStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if Primitive.to_proto(resource.last_transition_time):
            res.last_transition_time = Primitive.to_proto(resource.last_transition_time)
        if JobTerminalConditionSeverityEnum.to_proto(resource.severity):
            res.severity = JobTerminalConditionSeverityEnum.to_proto(resource.severity)
        if JobTerminalConditionReasonEnum.to_proto(resource.reason):
            res.reason = JobTerminalConditionReasonEnum.to_proto(resource.reason)
        if JobTerminalConditionInternalReasonEnum.to_proto(resource.internal_reason):
            res.internal_reason = JobTerminalConditionInternalReasonEnum.to_proto(
                resource.internal_reason
            )
        if JobTerminalConditionDomainMappingReasonEnum.to_proto(
            resource.domain_mapping_reason
        ):
            res.domain_mapping_reason = (
                JobTerminalConditionDomainMappingReasonEnum.to_proto(
                    resource.domain_mapping_reason
                )
            )
        if JobTerminalConditionRevisionReasonEnum.to_proto(resource.revision_reason):
            res.revision_reason = JobTerminalConditionRevisionReasonEnum.to_proto(
                resource.revision_reason
            )
        if JobTerminalConditionExecutionReasonEnum.to_proto(resource.execution_reason):
            res.execution_reason = JobTerminalConditionExecutionReasonEnum.to_proto(
                resource.execution_reason
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobTerminalCondition(
            type=Primitive.from_proto(resource.type),
            state=JobTerminalConditionStateEnum.from_proto(resource.state),
            message=Primitive.from_proto(resource.message),
            last_transition_time=Primitive.from_proto(resource.last_transition_time),
            severity=JobTerminalConditionSeverityEnum.from_proto(resource.severity),
            reason=JobTerminalConditionReasonEnum.from_proto(resource.reason),
            internal_reason=JobTerminalConditionInternalReasonEnum.from_proto(
                resource.internal_reason
            ),
            domain_mapping_reason=JobTerminalConditionDomainMappingReasonEnum.from_proto(
                resource.domain_mapping_reason
            ),
            revision_reason=JobTerminalConditionRevisionReasonEnum.from_proto(
                resource.revision_reason
            ),
            execution_reason=JobTerminalConditionExecutionReasonEnum.from_proto(
                resource.execution_reason
            ),
        )


class JobTerminalConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobTerminalCondition.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobTerminalCondition.from_proto(i) for i in resources]


class JobConditions(object):
    def __init__(
        self,
        type: str = None,
        state: str = None,
        message: str = None,
        last_transition_time: str = None,
        severity: str = None,
        reason: str = None,
        revision_reason: str = None,
        execution_reason: str = None,
    ):
        self.type = type
        self.state = state
        self.message = message
        self.last_transition_time = last_transition_time
        self.severity = severity
        self.reason = reason
        self.revision_reason = revision_reason
        self.execution_reason = execution_reason

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobConditions()
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if JobConditionsStateEnum.to_proto(resource.state):
            res.state = JobConditionsStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if Primitive.to_proto(resource.last_transition_time):
            res.last_transition_time = Primitive.to_proto(resource.last_transition_time)
        if JobConditionsSeverityEnum.to_proto(resource.severity):
            res.severity = JobConditionsSeverityEnum.to_proto(resource.severity)
        if JobConditionsReasonEnum.to_proto(resource.reason):
            res.reason = JobConditionsReasonEnum.to_proto(resource.reason)
        if JobConditionsRevisionReasonEnum.to_proto(resource.revision_reason):
            res.revision_reason = JobConditionsRevisionReasonEnum.to_proto(
                resource.revision_reason
            )
        if JobConditionsExecutionReasonEnum.to_proto(resource.execution_reason):
            res.execution_reason = JobConditionsExecutionReasonEnum.to_proto(
                resource.execution_reason
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobConditions(
            type=Primitive.from_proto(resource.type),
            state=JobConditionsStateEnum.from_proto(resource.state),
            message=Primitive.from_proto(resource.message),
            last_transition_time=Primitive.from_proto(resource.last_transition_time),
            severity=JobConditionsSeverityEnum.from_proto(resource.severity),
            reason=JobConditionsReasonEnum.from_proto(resource.reason),
            revision_reason=JobConditionsRevisionReasonEnum.from_proto(
                resource.revision_reason
            ),
            execution_reason=JobConditionsExecutionReasonEnum.from_proto(
                resource.execution_reason
            ),
        )


class JobConditionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobConditions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobConditions.from_proto(i) for i in resources]


class JobLatestSucceededExecution(object):
    def __init__(self, name: str = None, create_time: str = None):
        self.name = name
        self.create_time = create_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobLatestSucceededExecution()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.create_time):
            res.create_time = Primitive.to_proto(resource.create_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobLatestSucceededExecution(
            name=Primitive.from_proto(resource.name),
            create_time=Primitive.from_proto(resource.create_time),
        )


class JobLatestSucceededExecutionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobLatestSucceededExecution.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobLatestSucceededExecution.from_proto(i) for i in resources]


class JobLatestCreatedExecution(object):
    def __init__(self, name: str = None, create_time: str = None):
        self.name = name
        self.create_time = create_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = job_pb2.RunAlphaJobLatestCreatedExecution()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.create_time):
            res.create_time = Primitive.to_proto(resource.create_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return JobLatestCreatedExecution(
            name=Primitive.from_proto(resource.name),
            create_time=Primitive.from_proto(resource.create_time),
        )


class JobLatestCreatedExecutionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [JobLatestCreatedExecution.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [JobLatestCreatedExecution.from_proto(i) for i in resources]


class JobLaunchStageEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobLaunchStageEnum.Value(
            "RunAlphaJobLaunchStageEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobLaunchStageEnum.Name(resource)[
            len("RunAlphaJobLaunchStageEnum") :
        ]


class JobTemplateTemplateExecutionEnvironmentEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTemplateTemplateExecutionEnvironmentEnum.Value(
            "RunAlphaJobTemplateTemplateExecutionEnvironmentEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTemplateTemplateExecutionEnvironmentEnum.Name(
            resource
        )[len("RunAlphaJobTemplateTemplateExecutionEnvironmentEnum") :]


class JobTemplateTemplateVPCAccessEgressEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTemplateTemplateVPCAccessEgressEnum.Value(
            "RunAlphaJobTemplateTemplateVPCAccessEgressEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTemplateTemplateVPCAccessEgressEnum.Name(resource)[
            len("RunAlphaJobTemplateTemplateVPCAccessEgressEnum") :
        ]


class JobTerminalConditionStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionStateEnum.Value(
            "RunAlphaJobTerminalConditionStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionStateEnum.Name(resource)[
            len("RunAlphaJobTerminalConditionStateEnum") :
        ]


class JobTerminalConditionSeverityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionSeverityEnum.Value(
            "RunAlphaJobTerminalConditionSeverityEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionSeverityEnum.Name(resource)[
            len("RunAlphaJobTerminalConditionSeverityEnum") :
        ]


class JobTerminalConditionReasonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionReasonEnum.Value(
            "RunAlphaJobTerminalConditionReasonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionReasonEnum.Name(resource)[
            len("RunAlphaJobTerminalConditionReasonEnum") :
        ]


class JobTerminalConditionInternalReasonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionInternalReasonEnum.Value(
            "RunAlphaJobTerminalConditionInternalReasonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionInternalReasonEnum.Name(resource)[
            len("RunAlphaJobTerminalConditionInternalReasonEnum") :
        ]


class JobTerminalConditionDomainMappingReasonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionDomainMappingReasonEnum.Value(
            "RunAlphaJobTerminalConditionDomainMappingReasonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionDomainMappingReasonEnum.Name(
            resource
        )[len("RunAlphaJobTerminalConditionDomainMappingReasonEnum") :]


class JobTerminalConditionRevisionReasonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionRevisionReasonEnum.Value(
            "RunAlphaJobTerminalConditionRevisionReasonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionRevisionReasonEnum.Name(resource)[
            len("RunAlphaJobTerminalConditionRevisionReasonEnum") :
        ]


class JobTerminalConditionExecutionReasonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionExecutionReasonEnum.Value(
            "RunAlphaJobTerminalConditionExecutionReasonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobTerminalConditionExecutionReasonEnum.Name(resource)[
            len("RunAlphaJobTerminalConditionExecutionReasonEnum") :
        ]


class JobConditionsStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobConditionsStateEnum.Value(
            "RunAlphaJobConditionsStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobConditionsStateEnum.Name(resource)[
            len("RunAlphaJobConditionsStateEnum") :
        ]


class JobConditionsSeverityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobConditionsSeverityEnum.Value(
            "RunAlphaJobConditionsSeverityEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobConditionsSeverityEnum.Name(resource)[
            len("RunAlphaJobConditionsSeverityEnum") :
        ]


class JobConditionsReasonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobConditionsReasonEnum.Value(
            "RunAlphaJobConditionsReasonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobConditionsReasonEnum.Name(resource)[
            len("RunAlphaJobConditionsReasonEnum") :
        ]


class JobConditionsRevisionReasonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobConditionsRevisionReasonEnum.Value(
            "RunAlphaJobConditionsRevisionReasonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobConditionsRevisionReasonEnum.Name(resource)[
            len("RunAlphaJobConditionsRevisionReasonEnum") :
        ]


class JobConditionsExecutionReasonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobConditionsExecutionReasonEnum.Value(
            "RunAlphaJobConditionsExecutionReasonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return job_pb2.RunAlphaJobConditionsExecutionReasonEnum.Name(resource)[
            len("RunAlphaJobConditionsExecutionReasonEnum") :
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
