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
from google3.cloud.graphite.mmv2.services.google.run import service_pb2
from google3.cloud.graphite.mmv2.services.google.run import service_pb2_grpc

from typing import List


class Service(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
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
        ingress: str = None,
        launch_stage: str = None,
        binary_authorization: dict = None,
        template: dict = None,
        traffic: list = None,
        terminal_condition: dict = None,
        latest_ready_revision: str = None,
        latest_created_revision: str = None,
        traffic_statuses: list = None,
        uri: str = None,
        reconciling: bool = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.annotations = annotations
        self.client = client
        self.client_version = client_version
        self.ingress = ingress
        self.launch_stage = launch_stage
        self.binary_authorization = binary_authorization
        self.template = template
        self.traffic = traffic
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = service_pb2_grpc.RunAlphaServiceServiceStub(channel.Channel())
        request = service_pb2.ApplyRunAlphaServiceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.client):
            request.resource.client = Primitive.to_proto(self.client)

        if Primitive.to_proto(self.client_version):
            request.resource.client_version = Primitive.to_proto(self.client_version)

        if ServiceIngressEnum.to_proto(self.ingress):
            request.resource.ingress = ServiceIngressEnum.to_proto(self.ingress)

        if ServiceLaunchStageEnum.to_proto(self.launch_stage):
            request.resource.launch_stage = ServiceLaunchStageEnum.to_proto(
                self.launch_stage
            )

        if ServiceBinaryAuthorization.to_proto(self.binary_authorization):
            request.resource.binary_authorization.CopyFrom(
                ServiceBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            request.resource.ClearField("binary_authorization")
        if ServiceTemplate.to_proto(self.template):
            request.resource.template.CopyFrom(ServiceTemplate.to_proto(self.template))
        else:
            request.resource.ClearField("template")
        if ServiceTrafficArray.to_proto(self.traffic):
            request.resource.traffic.extend(ServiceTrafficArray.to_proto(self.traffic))
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyRunAlphaService(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
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
        self.ingress = ServiceIngressEnum.from_proto(response.ingress)
        self.launch_stage = ServiceLaunchStageEnum.from_proto(response.launch_stage)
        self.binary_authorization = ServiceBinaryAuthorization.from_proto(
            response.binary_authorization
        )
        self.template = ServiceTemplate.from_proto(response.template)
        self.traffic = ServiceTrafficArray.from_proto(response.traffic)
        self.terminal_condition = ServiceTerminalCondition.from_proto(
            response.terminal_condition
        )
        self.latest_ready_revision = Primitive.from_proto(
            response.latest_ready_revision
        )
        self.latest_created_revision = Primitive.from_proto(
            response.latest_created_revision
        )
        self.traffic_statuses = ServiceTrafficStatusesArray.from_proto(
            response.traffic_statuses
        )
        self.uri = Primitive.from_proto(response.uri)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = service_pb2_grpc.RunAlphaServiceServiceStub(channel.Channel())
        request = service_pb2.DeleteRunAlphaServiceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.client):
            request.resource.client = Primitive.to_proto(self.client)

        if Primitive.to_proto(self.client_version):
            request.resource.client_version = Primitive.to_proto(self.client_version)

        if ServiceIngressEnum.to_proto(self.ingress):
            request.resource.ingress = ServiceIngressEnum.to_proto(self.ingress)

        if ServiceLaunchStageEnum.to_proto(self.launch_stage):
            request.resource.launch_stage = ServiceLaunchStageEnum.to_proto(
                self.launch_stage
            )

        if ServiceBinaryAuthorization.to_proto(self.binary_authorization):
            request.resource.binary_authorization.CopyFrom(
                ServiceBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            request.resource.ClearField("binary_authorization")
        if ServiceTemplate.to_proto(self.template):
            request.resource.template.CopyFrom(ServiceTemplate.to_proto(self.template))
        else:
            request.resource.ClearField("template")
        if ServiceTrafficArray.to_proto(self.traffic):
            request.resource.traffic.extend(ServiceTrafficArray.to_proto(self.traffic))
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteRunAlphaService(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = service_pb2_grpc.RunAlphaServiceServiceStub(channel.Channel())
        request = service_pb2.ListRunAlphaServiceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListRunAlphaService(request).items

    def to_proto(self):
        resource = service_pb2.RunAlphaService()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.client):
            resource.client = Primitive.to_proto(self.client)
        if Primitive.to_proto(self.client_version):
            resource.client_version = Primitive.to_proto(self.client_version)
        if ServiceIngressEnum.to_proto(self.ingress):
            resource.ingress = ServiceIngressEnum.to_proto(self.ingress)
        if ServiceLaunchStageEnum.to_proto(self.launch_stage):
            resource.launch_stage = ServiceLaunchStageEnum.to_proto(self.launch_stage)
        if ServiceBinaryAuthorization.to_proto(self.binary_authorization):
            resource.binary_authorization.CopyFrom(
                ServiceBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            resource.ClearField("binary_authorization")
        if ServiceTemplate.to_proto(self.template):
            resource.template.CopyFrom(ServiceTemplate.to_proto(self.template))
        else:
            resource.ClearField("template")
        if ServiceTrafficArray.to_proto(self.traffic):
            resource.traffic.extend(ServiceTrafficArray.to_proto(self.traffic))
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ServiceBinaryAuthorization(object):
    def __init__(self, use_default: bool = None, breakglass_justification: str = None):
        self.use_default = use_default
        self.breakglass_justification = breakglass_justification

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceBinaryAuthorization()
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

        return ServiceBinaryAuthorization(
            use_default=Primitive.from_proto(resource.use_default),
            breakglass_justification=Primitive.from_proto(
                resource.breakglass_justification
            ),
        )


class ServiceBinaryAuthorizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceBinaryAuthorization.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceBinaryAuthorization.from_proto(i) for i in resources]


class ServiceTemplate(object):
    def __init__(
        self,
        revision: str = None,
        labels: dict = None,
        annotations: dict = None,
        scaling: dict = None,
        vpc_access: dict = None,
        container_concurrency: int = None,
        timeout: str = None,
        service_account: str = None,
        containers: list = None,
        volumes: list = None,
        execution_environment: str = None,
    ):
        self.revision = revision
        self.labels = labels
        self.annotations = annotations
        self.scaling = scaling
        self.vpc_access = vpc_access
        self.container_concurrency = container_concurrency
        self.timeout = timeout
        self.service_account = service_account
        self.containers = containers
        self.volumes = volumes
        self.execution_environment = execution_environment

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTemplate()
        if Primitive.to_proto(resource.revision):
            res.revision = Primitive.to_proto(resource.revision)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if Primitive.to_proto(resource.annotations):
            res.annotations = Primitive.to_proto(resource.annotations)
        if ServiceTemplateScaling.to_proto(resource.scaling):
            res.scaling.CopyFrom(ServiceTemplateScaling.to_proto(resource.scaling))
        else:
            res.ClearField("scaling")
        if ServiceTemplateVPCAccess.to_proto(resource.vpc_access):
            res.vpc_access.CopyFrom(
                ServiceTemplateVPCAccess.to_proto(resource.vpc_access)
            )
        else:
            res.ClearField("vpc_access")
        if Primitive.to_proto(resource.container_concurrency):
            res.container_concurrency = Primitive.to_proto(
                resource.container_concurrency
            )
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        if Primitive.to_proto(resource.service_account):
            res.service_account = Primitive.to_proto(resource.service_account)
        if ServiceTemplateContainersArray.to_proto(resource.containers):
            res.containers.extend(
                ServiceTemplateContainersArray.to_proto(resource.containers)
            )
        if ServiceTemplateVolumesArray.to_proto(resource.volumes):
            res.volumes.extend(ServiceTemplateVolumesArray.to_proto(resource.volumes))
        if ServiceTemplateExecutionEnvironmentEnum.to_proto(
            resource.execution_environment
        ):
            res.execution_environment = (
                ServiceTemplateExecutionEnvironmentEnum.to_proto(
                    resource.execution_environment
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTemplate(
            revision=Primitive.from_proto(resource.revision),
            labels=Primitive.from_proto(resource.labels),
            annotations=Primitive.from_proto(resource.annotations),
            scaling=ServiceTemplateScaling.from_proto(resource.scaling),
            vpc_access=ServiceTemplateVPCAccess.from_proto(resource.vpc_access),
            container_concurrency=Primitive.from_proto(resource.container_concurrency),
            timeout=Primitive.from_proto(resource.timeout),
            service_account=Primitive.from_proto(resource.service_account),
            containers=ServiceTemplateContainersArray.from_proto(resource.containers),
            volumes=ServiceTemplateVolumesArray.from_proto(resource.volumes),
            execution_environment=ServiceTemplateExecutionEnvironmentEnum.from_proto(
                resource.execution_environment
            ),
        )


class ServiceTemplateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplate.from_proto(i) for i in resources]


class ServiceTemplateScaling(object):
    def __init__(self, min_instance_count: int = None, max_instance_count: int = None):
        self.min_instance_count = min_instance_count
        self.max_instance_count = max_instance_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTemplateScaling()
        if Primitive.to_proto(resource.min_instance_count):
            res.min_instance_count = Primitive.to_proto(resource.min_instance_count)
        if Primitive.to_proto(resource.max_instance_count):
            res.max_instance_count = Primitive.to_proto(resource.max_instance_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTemplateScaling(
            min_instance_count=Primitive.from_proto(resource.min_instance_count),
            max_instance_count=Primitive.from_proto(resource.max_instance_count),
        )


class ServiceTemplateScalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateScaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplateScaling.from_proto(i) for i in resources]


class ServiceTemplateVPCAccess(object):
    def __init__(self, connector: str = None, egress: str = None):
        self.connector = connector
        self.egress = egress

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTemplateVPCAccess()
        if Primitive.to_proto(resource.connector):
            res.connector = Primitive.to_proto(resource.connector)
        if ServiceTemplateVPCAccessEgressEnum.to_proto(resource.egress):
            res.egress = ServiceTemplateVPCAccessEgressEnum.to_proto(resource.egress)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTemplateVPCAccess(
            connector=Primitive.from_proto(resource.connector),
            egress=ServiceTemplateVPCAccessEgressEnum.from_proto(resource.egress),
        )


class ServiceTemplateVPCAccessArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateVPCAccess.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplateVPCAccess.from_proto(i) for i in resources]


class ServiceTemplateContainers(object):
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

        res = service_pb2.RunAlphaServiceTemplateContainers()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        if Primitive.to_proto(resource.command):
            res.command.extend(Primitive.to_proto(resource.command))
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if ServiceTemplateContainersEnvArray.to_proto(resource.env):
            res.env.extend(ServiceTemplateContainersEnvArray.to_proto(resource.env))
        if ServiceTemplateContainersResources.to_proto(resource.resources):
            res.resources.CopyFrom(
                ServiceTemplateContainersResources.to_proto(resource.resources)
            )
        else:
            res.ClearField("resources")
        if ServiceTemplateContainersPortsArray.to_proto(resource.ports):
            res.ports.extend(
                ServiceTemplateContainersPortsArray.to_proto(resource.ports)
            )
        if ServiceTemplateContainersVolumeMountsArray.to_proto(resource.volume_mounts):
            res.volume_mounts.extend(
                ServiceTemplateContainersVolumeMountsArray.to_proto(
                    resource.volume_mounts
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTemplateContainers(
            name=Primitive.from_proto(resource.name),
            image=Primitive.from_proto(resource.image),
            command=Primitive.from_proto(resource.command),
            args=Primitive.from_proto(resource.args),
            env=ServiceTemplateContainersEnvArray.from_proto(resource.env),
            resources=ServiceTemplateContainersResources.from_proto(resource.resources),
            ports=ServiceTemplateContainersPortsArray.from_proto(resource.ports),
            volume_mounts=ServiceTemplateContainersVolumeMountsArray.from_proto(
                resource.volume_mounts
            ),
        )


class ServiceTemplateContainersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateContainers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplateContainers.from_proto(i) for i in resources]


class ServiceTemplateContainersEnv(object):
    def __init__(self, name: str = None, value: str = None, value_source: dict = None):
        self.name = name
        self.value = value
        self.value_source = value_source

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTemplateContainersEnv()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if ServiceTemplateContainersEnvValueSource.to_proto(resource.value_source):
            res.value_source.CopyFrom(
                ServiceTemplateContainersEnvValueSource.to_proto(resource.value_source)
            )
        else:
            res.ClearField("value_source")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTemplateContainersEnv(
            name=Primitive.from_proto(resource.name),
            value=Primitive.from_proto(resource.value),
            value_source=ServiceTemplateContainersEnvValueSource.from_proto(
                resource.value_source
            ),
        )


class ServiceTemplateContainersEnvArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateContainersEnv.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplateContainersEnv.from_proto(i) for i in resources]


class ServiceTemplateContainersEnvValueSource(object):
    def __init__(self, secret_key_ref: dict = None):
        self.secret_key_ref = secret_key_ref

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTemplateContainersEnvValueSource()
        if ServiceTemplateContainersEnvValueSourceSecretKeyRef.to_proto(
            resource.secret_key_ref
        ):
            res.secret_key_ref.CopyFrom(
                ServiceTemplateContainersEnvValueSourceSecretKeyRef.to_proto(
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

        return ServiceTemplateContainersEnvValueSource(
            secret_key_ref=ServiceTemplateContainersEnvValueSourceSecretKeyRef.from_proto(
                resource.secret_key_ref
            ),
        )


class ServiceTemplateContainersEnvValueSourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateContainersEnvValueSource.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceTemplateContainersEnvValueSource.from_proto(i) for i in resources
        ]


class ServiceTemplateContainersEnvValueSourceSecretKeyRef(object):
    def __init__(self, secret: str = None, version: str = None):
        self.secret = secret
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTemplateContainersEnvValueSourceSecretKeyRef()
        if Primitive.to_proto(resource.secret):
            res.secret = Primitive.to_proto(resource.secret)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTemplateContainersEnvValueSourceSecretKeyRef(
            secret=Primitive.from_proto(resource.secret),
            version=Primitive.from_proto(resource.version),
        )


class ServiceTemplateContainersEnvValueSourceSecretKeyRefArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceTemplateContainersEnvValueSourceSecretKeyRef.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceTemplateContainersEnvValueSourceSecretKeyRef.from_proto(i)
            for i in resources
        ]


class ServiceTemplateContainersResources(object):
    def __init__(self, limits: dict = None, cpu_idle: bool = None):
        self.limits = limits
        self.cpu_idle = cpu_idle

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTemplateContainersResources()
        if Primitive.to_proto(resource.limits):
            res.limits = Primitive.to_proto(resource.limits)
        if Primitive.to_proto(resource.cpu_idle):
            res.cpu_idle = Primitive.to_proto(resource.cpu_idle)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTemplateContainersResources(
            limits=Primitive.from_proto(resource.limits),
            cpu_idle=Primitive.from_proto(resource.cpu_idle),
        )


class ServiceTemplateContainersResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateContainersResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplateContainersResources.from_proto(i) for i in resources]


class ServiceTemplateContainersPorts(object):
    def __init__(self, name: str = None, container_port: int = None):
        self.name = name
        self.container_port = container_port

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTemplateContainersPorts()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.container_port):
            res.container_port = Primitive.to_proto(resource.container_port)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTemplateContainersPorts(
            name=Primitive.from_proto(resource.name),
            container_port=Primitive.from_proto(resource.container_port),
        )


class ServiceTemplateContainersPortsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateContainersPorts.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplateContainersPorts.from_proto(i) for i in resources]


class ServiceTemplateContainersVolumeMounts(object):
    def __init__(self, name: str = None, mount_path: str = None):
        self.name = name
        self.mount_path = mount_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTemplateContainersVolumeMounts()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.mount_path):
            res.mount_path = Primitive.to_proto(resource.mount_path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTemplateContainersVolumeMounts(
            name=Primitive.from_proto(resource.name),
            mount_path=Primitive.from_proto(resource.mount_path),
        )


class ServiceTemplateContainersVolumeMountsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateContainersVolumeMounts.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplateContainersVolumeMounts.from_proto(i) for i in resources]


class ServiceTemplateVolumes(object):
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

        res = service_pb2.RunAlphaServiceTemplateVolumes()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if ServiceTemplateVolumesSecret.to_proto(resource.secret):
            res.secret.CopyFrom(ServiceTemplateVolumesSecret.to_proto(resource.secret))
        else:
            res.ClearField("secret")
        if ServiceTemplateVolumesCloudSqlInstance.to_proto(resource.cloud_sql_instance):
            res.cloud_sql_instance.CopyFrom(
                ServiceTemplateVolumesCloudSqlInstance.to_proto(
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

        return ServiceTemplateVolumes(
            name=Primitive.from_proto(resource.name),
            secret=ServiceTemplateVolumesSecret.from_proto(resource.secret),
            cloud_sql_instance=ServiceTemplateVolumesCloudSqlInstance.from_proto(
                resource.cloud_sql_instance
            ),
        )


class ServiceTemplateVolumesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateVolumes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplateVolumes.from_proto(i) for i in resources]


class ServiceTemplateVolumesSecret(object):
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

        res = service_pb2.RunAlphaServiceTemplateVolumesSecret()
        if Primitive.to_proto(resource.secret):
            res.secret = Primitive.to_proto(resource.secret)
        if ServiceTemplateVolumesSecretItemsArray.to_proto(resource.items):
            res.items.extend(
                ServiceTemplateVolumesSecretItemsArray.to_proto(resource.items)
            )
        if Primitive.to_proto(resource.default_mode):
            res.default_mode = Primitive.to_proto(resource.default_mode)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTemplateVolumesSecret(
            secret=Primitive.from_proto(resource.secret),
            items=ServiceTemplateVolumesSecretItemsArray.from_proto(resource.items),
            default_mode=Primitive.from_proto(resource.default_mode),
        )


class ServiceTemplateVolumesSecretArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateVolumesSecret.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplateVolumesSecret.from_proto(i) for i in resources]


class ServiceTemplateVolumesSecretItems(object):
    def __init__(self, path: str = None, version: str = None, mode: int = None):
        self.path = path
        self.version = version
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTemplateVolumesSecretItems()
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

        return ServiceTemplateVolumesSecretItems(
            path=Primitive.from_proto(resource.path),
            version=Primitive.from_proto(resource.version),
            mode=Primitive.from_proto(resource.mode),
        )


class ServiceTemplateVolumesSecretItemsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateVolumesSecretItems.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplateVolumesSecretItems.from_proto(i) for i in resources]


class ServiceTemplateVolumesCloudSqlInstance(object):
    def __init__(self, instances: list = None):
        self.instances = instances

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTemplateVolumesCloudSqlInstance()
        if Primitive.to_proto(resource.instances):
            res.instances.extend(Primitive.to_proto(resource.instances))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTemplateVolumesCloudSqlInstance(
            instances=Primitive.from_proto(resource.instances),
        )


class ServiceTemplateVolumesCloudSqlInstanceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTemplateVolumesCloudSqlInstance.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTemplateVolumesCloudSqlInstance.from_proto(i) for i in resources]


class ServiceTraffic(object):
    def __init__(
        self,
        type: str = None,
        revision: str = None,
        percent: int = None,
        tag: str = None,
    ):
        self.type = type
        self.revision = revision
        self.percent = percent
        self.tag = tag

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTraffic()
        if ServiceTrafficTypeEnum.to_proto(resource.type):
            res.type = ServiceTrafficTypeEnum.to_proto(resource.type)
        if Primitive.to_proto(resource.revision):
            res.revision = Primitive.to_proto(resource.revision)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTraffic(
            type=ServiceTrafficTypeEnum.from_proto(resource.type),
            revision=Primitive.from_proto(resource.revision),
            percent=Primitive.from_proto(resource.percent),
            tag=Primitive.from_proto(resource.tag),
        )


class ServiceTrafficArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTraffic.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTraffic.from_proto(i) for i in resources]


class ServiceTerminalCondition(object):
    def __init__(
        self,
        type: str = None,
        state: str = None,
        message: str = None,
        last_transition_time: str = None,
        severity: str = None,
        reason: str = None,
        revision_reason: str = None,
        job_reason: str = None,
    ):
        self.type = type
        self.state = state
        self.message = message
        self.last_transition_time = last_transition_time
        self.severity = severity
        self.reason = reason
        self.revision_reason = revision_reason
        self.job_reason = job_reason

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTerminalCondition()
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if ServiceTerminalConditionStateEnum.to_proto(resource.state):
            res.state = ServiceTerminalConditionStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if Primitive.to_proto(resource.last_transition_time):
            res.last_transition_time = Primitive.to_proto(resource.last_transition_time)
        if ServiceTerminalConditionSeverityEnum.to_proto(resource.severity):
            res.severity = ServiceTerminalConditionSeverityEnum.to_proto(
                resource.severity
            )
        if ServiceTerminalConditionReasonEnum.to_proto(resource.reason):
            res.reason = ServiceTerminalConditionReasonEnum.to_proto(resource.reason)
        if ServiceTerminalConditionRevisionReasonEnum.to_proto(
            resource.revision_reason
        ):
            res.revision_reason = ServiceTerminalConditionRevisionReasonEnum.to_proto(
                resource.revision_reason
            )
        if ServiceTerminalConditionJobReasonEnum.to_proto(resource.job_reason):
            res.job_reason = ServiceTerminalConditionJobReasonEnum.to_proto(
                resource.job_reason
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTerminalCondition(
            type=Primitive.from_proto(resource.type),
            state=ServiceTerminalConditionStateEnum.from_proto(resource.state),
            message=Primitive.from_proto(resource.message),
            last_transition_time=Primitive.from_proto(resource.last_transition_time),
            severity=ServiceTerminalConditionSeverityEnum.from_proto(resource.severity),
            reason=ServiceTerminalConditionReasonEnum.from_proto(resource.reason),
            revision_reason=ServiceTerminalConditionRevisionReasonEnum.from_proto(
                resource.revision_reason
            ),
            job_reason=ServiceTerminalConditionJobReasonEnum.from_proto(
                resource.job_reason
            ),
        )


class ServiceTerminalConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTerminalCondition.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTerminalCondition.from_proto(i) for i in resources]


class ServiceTrafficStatuses(object):
    def __init__(
        self,
        type: str = None,
        revision: str = None,
        percent: int = None,
        tag: str = None,
        uri: str = None,
    ):
        self.type = type
        self.revision = revision
        self.percent = percent
        self.tag = tag
        self.uri = uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunAlphaServiceTrafficStatuses()
        if ServiceTrafficStatusesTypeEnum.to_proto(resource.type):
            res.type = ServiceTrafficStatusesTypeEnum.to_proto(resource.type)
        if Primitive.to_proto(resource.revision):
            res.revision = Primitive.to_proto(resource.revision)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTrafficStatuses(
            type=ServiceTrafficStatusesTypeEnum.from_proto(resource.type),
            revision=Primitive.from_proto(resource.revision),
            percent=Primitive.from_proto(resource.percent),
            tag=Primitive.from_proto(resource.tag),
            uri=Primitive.from_proto(resource.uri),
        )


class ServiceTrafficStatusesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTrafficStatuses.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTrafficStatuses.from_proto(i) for i in resources]


class ServiceIngressEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceIngressEnum.Value(
            "RunAlphaServiceIngressEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceIngressEnum.Name(resource)[
            len("RunAlphaServiceIngressEnum") :
        ]


class ServiceLaunchStageEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceLaunchStageEnum.Value(
            "RunAlphaServiceLaunchStageEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceLaunchStageEnum.Name(resource)[
            len("RunAlphaServiceLaunchStageEnum") :
        ]


class ServiceTemplateVPCAccessEgressEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTemplateVPCAccessEgressEnum.Value(
            "RunAlphaServiceTemplateVPCAccessEgressEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTemplateVPCAccessEgressEnum.Name(resource)[
            len("RunAlphaServiceTemplateVPCAccessEgressEnum") :
        ]


class ServiceTemplateExecutionEnvironmentEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTemplateExecutionEnvironmentEnum.Value(
            "RunAlphaServiceTemplateExecutionEnvironmentEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTemplateExecutionEnvironmentEnum.Name(
            resource
        )[len("RunAlphaServiceTemplateExecutionEnvironmentEnum") :]


class ServiceTrafficTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTrafficTypeEnum.Value(
            "RunAlphaServiceTrafficTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTrafficTypeEnum.Name(resource)[
            len("RunAlphaServiceTrafficTypeEnum") :
        ]


class ServiceTerminalConditionStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTerminalConditionStateEnum.Value(
            "RunAlphaServiceTerminalConditionStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTerminalConditionStateEnum.Name(resource)[
            len("RunAlphaServiceTerminalConditionStateEnum") :
        ]


class ServiceTerminalConditionSeverityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTerminalConditionSeverityEnum.Value(
            "RunAlphaServiceTerminalConditionSeverityEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTerminalConditionSeverityEnum.Name(resource)[
            len("RunAlphaServiceTerminalConditionSeverityEnum") :
        ]


class ServiceTerminalConditionReasonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTerminalConditionReasonEnum.Value(
            "RunAlphaServiceTerminalConditionReasonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTerminalConditionReasonEnum.Name(resource)[
            len("RunAlphaServiceTerminalConditionReasonEnum") :
        ]


class ServiceTerminalConditionRevisionReasonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTerminalConditionRevisionReasonEnum.Value(
            "RunAlphaServiceTerminalConditionRevisionReasonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTerminalConditionRevisionReasonEnum.Name(
            resource
        )[len("RunAlphaServiceTerminalConditionRevisionReasonEnum") :]


class ServiceTerminalConditionJobReasonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTerminalConditionJobReasonEnum.Value(
            "RunAlphaServiceTerminalConditionJobReasonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTerminalConditionJobReasonEnum.Name(resource)[
            len("RunAlphaServiceTerminalConditionJobReasonEnum") :
        ]


class ServiceTrafficStatusesTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTrafficStatusesTypeEnum.Value(
            "RunAlphaServiceTrafficStatusesTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_pb2.RunAlphaServiceTrafficStatusesTypeEnum.Name(resource)[
            len("RunAlphaServiceTrafficStatusesTypeEnum") :
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
