# Copyright 2022 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.vertex_ai import model_pb2
from google3.cloud.graphite.mmv2.services.google.vertex_ai import model_pb2_grpc

from typing import List


class Model(object):
    def __init__(
        self,
        name: str = None,
        version_id: str = None,
        version_create_time: str = None,
        version_update_time: str = None,
        display_name: str = None,
        description: str = None,
        version_description: str = None,
        supported_export_formats: list = None,
        training_pipeline: str = None,
        original_model_info: dict = None,
        container_spec: dict = None,
        artifact_uri: str = None,
        supported_deployment_resources_types: list = None,
        supported_input_storage_formats: list = None,
        supported_output_storage_formats: list = None,
        create_time: str = None,
        update_time: str = None,
        deployed_models: list = None,
        etag: str = None,
        labels: dict = None,
        encryption_spec: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.description = description
        self.version_description = version_description
        self.container_spec = container_spec
        self.artifact_uri = artifact_uri
        self.labels = labels
        self.encryption_spec = encryption_spec
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = model_pb2_grpc.VertexaiModelServiceStub(channel.Channel())
        request = model_pb2.ApplyVertexaiModelRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.version_description):
            request.resource.version_description = Primitive.to_proto(
                self.version_description
            )

        if ModelContainerSpec.to_proto(self.container_spec):
            request.resource.container_spec.CopyFrom(
                ModelContainerSpec.to_proto(self.container_spec)
            )
        else:
            request.resource.ClearField("container_spec")
        if Primitive.to_proto(self.artifact_uri):
            request.resource.artifact_uri = Primitive.to_proto(self.artifact_uri)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if ModelEncryptionSpec.to_proto(self.encryption_spec):
            request.resource.encryption_spec.CopyFrom(
                ModelEncryptionSpec.to_proto(self.encryption_spec)
            )
        else:
            request.resource.ClearField("encryption_spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyVertexaiModel(request)
        self.name = Primitive.from_proto(response.name)
        self.version_id = Primitive.from_proto(response.version_id)
        self.version_create_time = Primitive.from_proto(response.version_create_time)
        self.version_update_time = Primitive.from_proto(response.version_update_time)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.version_description = Primitive.from_proto(response.version_description)
        self.supported_export_formats = ModelSupportedExportFormatsArray.from_proto(
            response.supported_export_formats
        )
        self.training_pipeline = Primitive.from_proto(response.training_pipeline)
        self.original_model_info = ModelOriginalModelInfo.from_proto(
            response.original_model_info
        )
        self.container_spec = ModelContainerSpec.from_proto(response.container_spec)
        self.artifact_uri = Primitive.from_proto(response.artifact_uri)
        self.supported_deployment_resources_types = (
            ModelSupportedDeploymentResourcesTypesEnumArray.from_proto(
                response.supported_deployment_resources_types
            )
        )
        self.supported_input_storage_formats = Primitive.from_proto(
            response.supported_input_storage_formats
        )
        self.supported_output_storage_formats = Primitive.from_proto(
            response.supported_output_storage_formats
        )
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.deployed_models = ModelDeployedModelsArray.from_proto(
            response.deployed_models
        )
        self.etag = Primitive.from_proto(response.etag)
        self.labels = Primitive.from_proto(response.labels)
        self.encryption_spec = ModelEncryptionSpec.from_proto(response.encryption_spec)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = model_pb2_grpc.VertexaiModelServiceStub(channel.Channel())
        request = model_pb2.DeleteVertexaiModelRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.version_description):
            request.resource.version_description = Primitive.to_proto(
                self.version_description
            )

        if ModelContainerSpec.to_proto(self.container_spec):
            request.resource.container_spec.CopyFrom(
                ModelContainerSpec.to_proto(self.container_spec)
            )
        else:
            request.resource.ClearField("container_spec")
        if Primitive.to_proto(self.artifact_uri):
            request.resource.artifact_uri = Primitive.to_proto(self.artifact_uri)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if ModelEncryptionSpec.to_proto(self.encryption_spec):
            request.resource.encryption_spec.CopyFrom(
                ModelEncryptionSpec.to_proto(self.encryption_spec)
            )
        else:
            request.resource.ClearField("encryption_spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteVertexaiModel(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = model_pb2_grpc.VertexaiModelServiceStub(channel.Channel())
        request = model_pb2.ListVertexaiModelRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListVertexaiModel(request).items

    def to_proto(self):
        resource = model_pb2.VertexaiModel()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.version_description):
            resource.version_description = Primitive.to_proto(self.version_description)
        if ModelContainerSpec.to_proto(self.container_spec):
            resource.container_spec.CopyFrom(
                ModelContainerSpec.to_proto(self.container_spec)
            )
        else:
            resource.ClearField("container_spec")
        if Primitive.to_proto(self.artifact_uri):
            resource.artifact_uri = Primitive.to_proto(self.artifact_uri)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if ModelEncryptionSpec.to_proto(self.encryption_spec):
            resource.encryption_spec.CopyFrom(
                ModelEncryptionSpec.to_proto(self.encryption_spec)
            )
        else:
            resource.ClearField("encryption_spec")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ModelSupportedExportFormats(object):
    def __init__(self, id: str = None, exportable_contents: list = None):
        self.id = id
        self.exportable_contents = exportable_contents

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = model_pb2.VertexaiModelSupportedExportFormats()
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if ModelSupportedExportFormatsExportableContentsEnumArray.to_proto(
            resource.exportable_contents
        ):
            res.exportable_contents.extend(
                ModelSupportedExportFormatsExportableContentsEnumArray.to_proto(
                    resource.exportable_contents
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ModelSupportedExportFormats(
            id=Primitive.from_proto(resource.id),
            exportable_contents=ModelSupportedExportFormatsExportableContentsEnumArray.from_proto(
                resource.exportable_contents
            ),
        )


class ModelSupportedExportFormatsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ModelSupportedExportFormats.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ModelSupportedExportFormats.from_proto(i) for i in resources]


class ModelOriginalModelInfo(object):
    def __init__(self, model: str = None):
        self.model = model

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = model_pb2.VertexaiModelOriginalModelInfo()
        if Primitive.to_proto(resource.model):
            res.model = Primitive.to_proto(resource.model)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ModelOriginalModelInfo(
            model=Primitive.from_proto(resource.model),
        )


class ModelOriginalModelInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ModelOriginalModelInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ModelOriginalModelInfo.from_proto(i) for i in resources]


class ModelContainerSpec(object):
    def __init__(
        self,
        image_uri: str = None,
        command: list = None,
        args: list = None,
        env: list = None,
        ports: list = None,
        predict_route: str = None,
        health_route: str = None,
    ):
        self.image_uri = image_uri
        self.command = command
        self.args = args
        self.env = env
        self.ports = ports
        self.predict_route = predict_route
        self.health_route = health_route

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = model_pb2.VertexaiModelContainerSpec()
        if Primitive.to_proto(resource.image_uri):
            res.image_uri = Primitive.to_proto(resource.image_uri)
        if Primitive.to_proto(resource.command):
            res.command.extend(Primitive.to_proto(resource.command))
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if ModelContainerSpecEnvArray.to_proto(resource.env):
            res.env.extend(ModelContainerSpecEnvArray.to_proto(resource.env))
        if ModelContainerSpecPortsArray.to_proto(resource.ports):
            res.ports.extend(ModelContainerSpecPortsArray.to_proto(resource.ports))
        if Primitive.to_proto(resource.predict_route):
            res.predict_route = Primitive.to_proto(resource.predict_route)
        if Primitive.to_proto(resource.health_route):
            res.health_route = Primitive.to_proto(resource.health_route)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ModelContainerSpec(
            image_uri=Primitive.from_proto(resource.image_uri),
            command=Primitive.from_proto(resource.command),
            args=Primitive.from_proto(resource.args),
            env=ModelContainerSpecEnvArray.from_proto(resource.env),
            ports=ModelContainerSpecPortsArray.from_proto(resource.ports),
            predict_route=Primitive.from_proto(resource.predict_route),
            health_route=Primitive.from_proto(resource.health_route),
        )


class ModelContainerSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ModelContainerSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ModelContainerSpec.from_proto(i) for i in resources]


class ModelContainerSpecEnv(object):
    def __init__(self, name: str = None, value: str = None):
        self.name = name
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = model_pb2.VertexaiModelContainerSpecEnv()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ModelContainerSpecEnv(
            name=Primitive.from_proto(resource.name),
            value=Primitive.from_proto(resource.value),
        )


class ModelContainerSpecEnvArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ModelContainerSpecEnv.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ModelContainerSpecEnv.from_proto(i) for i in resources]


class ModelContainerSpecPorts(object):
    def __init__(self, container_port: int = None):
        self.container_port = container_port

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = model_pb2.VertexaiModelContainerSpecPorts()
        if Primitive.to_proto(resource.container_port):
            res.container_port = Primitive.to_proto(resource.container_port)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ModelContainerSpecPorts(
            container_port=Primitive.from_proto(resource.container_port),
        )


class ModelContainerSpecPortsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ModelContainerSpecPorts.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ModelContainerSpecPorts.from_proto(i) for i in resources]


class ModelDeployedModels(object):
    def __init__(self, endpoint: str = None, deployed_model_id: str = None):
        self.endpoint = endpoint
        self.deployed_model_id = deployed_model_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = model_pb2.VertexaiModelDeployedModels()
        if Primitive.to_proto(resource.endpoint):
            res.endpoint = Primitive.to_proto(resource.endpoint)
        if Primitive.to_proto(resource.deployed_model_id):
            res.deployed_model_id = Primitive.to_proto(resource.deployed_model_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ModelDeployedModels(
            endpoint=Primitive.from_proto(resource.endpoint),
            deployed_model_id=Primitive.from_proto(resource.deployed_model_id),
        )


class ModelDeployedModelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ModelDeployedModels.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ModelDeployedModels.from_proto(i) for i in resources]


class ModelEncryptionSpec(object):
    def __init__(self, kms_key_name: str = None):
        self.kms_key_name = kms_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = model_pb2.VertexaiModelEncryptionSpec()
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ModelEncryptionSpec(
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
        )


class ModelEncryptionSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ModelEncryptionSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ModelEncryptionSpec.from_proto(i) for i in resources]


class ModelSupportedExportFormatsExportableContentsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            model_pb2.VertexaiModelSupportedExportFormatsExportableContentsEnum.Value(
                "VertexaiModelSupportedExportFormatsExportableContentsEnum%s" % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return model_pb2.VertexaiModelSupportedExportFormatsExportableContentsEnum.Name(
            resource
        )[len("VertexaiModelSupportedExportFormatsExportableContentsEnum") :]


class ModelSupportedDeploymentResourcesTypesEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return model_pb2.VertexaiModelSupportedDeploymentResourcesTypesEnum.Value(
            "VertexaiModelSupportedDeploymentResourcesTypesEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return model_pb2.VertexaiModelSupportedDeploymentResourcesTypesEnum.Name(
            resource
        )[len("VertexaiModelSupportedDeploymentResourcesTypesEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
