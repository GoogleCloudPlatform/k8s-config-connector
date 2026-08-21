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
from google3.cloud.graphite.mmv2.services.google.bigquery import routine_pb2
from google3.cloud.graphite.mmv2.services.google.bigquery import routine_pb2_grpc

from typing import List


class Routine(object):
    def __init__(
        self,
        etag: str = None,
        name: str = None,
        project: str = None,
        dataset: str = None,
        routine_type: str = None,
        creation_time: int = None,
        last_modified_time: int = None,
        language: str = None,
        arguments: list = None,
        return_type: dict = None,
        imported_libraries: list = None,
        definition_body: str = None,
        description: str = None,
        determinism_level: str = None,
        strict_mode: bool = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.project = project
        self.dataset = dataset
        self.routine_type = routine_type
        self.language = language
        self.arguments = arguments
        self.return_type = return_type
        self.imported_libraries = imported_libraries
        self.definition_body = definition_body
        self.description = description
        self.determinism_level = determinism_level
        self.strict_mode = strict_mode
        self.service_account_file = service_account_file

    def apply(self):
        stub = routine_pb2_grpc.BigqueryBetaRoutineServiceStub(channel.Channel())
        request = routine_pb2.ApplyBigqueryBetaRoutineRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.dataset):
            request.resource.dataset = Primitive.to_proto(self.dataset)

        if RoutineRoutineTypeEnum.to_proto(self.routine_type):
            request.resource.routine_type = RoutineRoutineTypeEnum.to_proto(
                self.routine_type
            )

        if RoutineLanguageEnum.to_proto(self.language):
            request.resource.language = RoutineLanguageEnum.to_proto(self.language)

        if RoutineArgumentsArray.to_proto(self.arguments):
            request.resource.arguments.extend(
                RoutineArgumentsArray.to_proto(self.arguments)
            )
        if RoutineArgumentsDataType.to_proto(self.return_type):
            request.resource.return_type.CopyFrom(
                RoutineArgumentsDataType.to_proto(self.return_type)
            )
        else:
            request.resource.ClearField("return_type")
        if Primitive.to_proto(self.imported_libraries):
            request.resource.imported_libraries.extend(
                Primitive.to_proto(self.imported_libraries)
            )
        if Primitive.to_proto(self.definition_body):
            request.resource.definition_body = Primitive.to_proto(self.definition_body)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if RoutineDeterminismLevelEnum.to_proto(self.determinism_level):
            request.resource.determinism_level = RoutineDeterminismLevelEnum.to_proto(
                self.determinism_level
            )

        if Primitive.to_proto(self.strict_mode):
            request.resource.strict_mode = Primitive.to_proto(self.strict_mode)

        request.service_account_file = self.service_account_file

        response = stub.ApplyBigqueryBetaRoutine(request)
        self.etag = Primitive.from_proto(response.etag)
        self.name = Primitive.from_proto(response.name)
        self.project = Primitive.from_proto(response.project)
        self.dataset = Primitive.from_proto(response.dataset)
        self.routine_type = RoutineRoutineTypeEnum.from_proto(response.routine_type)
        self.creation_time = Primitive.from_proto(response.creation_time)
        self.last_modified_time = Primitive.from_proto(response.last_modified_time)
        self.language = RoutineLanguageEnum.from_proto(response.language)
        self.arguments = RoutineArgumentsArray.from_proto(response.arguments)
        self.return_type = RoutineArgumentsDataType.from_proto(response.return_type)
        self.imported_libraries = Primitive.from_proto(response.imported_libraries)
        self.definition_body = Primitive.from_proto(response.definition_body)
        self.description = Primitive.from_proto(response.description)
        self.determinism_level = RoutineDeterminismLevelEnum.from_proto(
            response.determinism_level
        )
        self.strict_mode = Primitive.from_proto(response.strict_mode)

    def delete(self):
        stub = routine_pb2_grpc.BigqueryBetaRoutineServiceStub(channel.Channel())
        request = routine_pb2.DeleteBigqueryBetaRoutineRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.dataset):
            request.resource.dataset = Primitive.to_proto(self.dataset)

        if RoutineRoutineTypeEnum.to_proto(self.routine_type):
            request.resource.routine_type = RoutineRoutineTypeEnum.to_proto(
                self.routine_type
            )

        if RoutineLanguageEnum.to_proto(self.language):
            request.resource.language = RoutineLanguageEnum.to_proto(self.language)

        if RoutineArgumentsArray.to_proto(self.arguments):
            request.resource.arguments.extend(
                RoutineArgumentsArray.to_proto(self.arguments)
            )
        if RoutineArgumentsDataType.to_proto(self.return_type):
            request.resource.return_type.CopyFrom(
                RoutineArgumentsDataType.to_proto(self.return_type)
            )
        else:
            request.resource.ClearField("return_type")
        if Primitive.to_proto(self.imported_libraries):
            request.resource.imported_libraries.extend(
                Primitive.to_proto(self.imported_libraries)
            )
        if Primitive.to_proto(self.definition_body):
            request.resource.definition_body = Primitive.to_proto(self.definition_body)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if RoutineDeterminismLevelEnum.to_proto(self.determinism_level):
            request.resource.determinism_level = RoutineDeterminismLevelEnum.to_proto(
                self.determinism_level
            )

        if Primitive.to_proto(self.strict_mode):
            request.resource.strict_mode = Primitive.to_proto(self.strict_mode)

        response = stub.DeleteBigqueryBetaRoutine(request)

    @classmethod
    def list(self, project, dataset, service_account_file=""):
        stub = routine_pb2_grpc.BigqueryBetaRoutineServiceStub(channel.Channel())
        request = routine_pb2.ListBigqueryBetaRoutineRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Dataset = dataset

        return stub.ListBigqueryBetaRoutine(request).items

    def to_proto(self):
        resource = routine_pb2.BigqueryBetaRoutine()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.dataset):
            resource.dataset = Primitive.to_proto(self.dataset)
        if RoutineRoutineTypeEnum.to_proto(self.routine_type):
            resource.routine_type = RoutineRoutineTypeEnum.to_proto(self.routine_type)
        if RoutineLanguageEnum.to_proto(self.language):
            resource.language = RoutineLanguageEnum.to_proto(self.language)
        if RoutineArgumentsArray.to_proto(self.arguments):
            resource.arguments.extend(RoutineArgumentsArray.to_proto(self.arguments))
        if RoutineArgumentsDataType.to_proto(self.return_type):
            resource.return_type.CopyFrom(
                RoutineArgumentsDataType.to_proto(self.return_type)
            )
        else:
            resource.ClearField("return_type")
        if Primitive.to_proto(self.imported_libraries):
            resource.imported_libraries.extend(
                Primitive.to_proto(self.imported_libraries)
            )
        if Primitive.to_proto(self.definition_body):
            resource.definition_body = Primitive.to_proto(self.definition_body)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if RoutineDeterminismLevelEnum.to_proto(self.determinism_level):
            resource.determinism_level = RoutineDeterminismLevelEnum.to_proto(
                self.determinism_level
            )
        if Primitive.to_proto(self.strict_mode):
            resource.strict_mode = Primitive.to_proto(self.strict_mode)
        return resource


class RoutineArguments(object):
    def __init__(
        self,
        name: str = None,
        argument_kind: str = None,
        mode: str = None,
        data_type: dict = None,
    ):
        self.name = name
        self.argument_kind = argument_kind
        self.mode = mode
        self.data_type = data_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = routine_pb2.BigqueryBetaRoutineArguments()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if RoutineArgumentsArgumentKindEnum.to_proto(resource.argument_kind):
            res.argument_kind = RoutineArgumentsArgumentKindEnum.to_proto(
                resource.argument_kind
            )
        if RoutineArgumentsModeEnum.to_proto(resource.mode):
            res.mode = RoutineArgumentsModeEnum.to_proto(resource.mode)
        if RoutineArgumentsDataType.to_proto(resource.data_type):
            res.data_type.CopyFrom(
                RoutineArgumentsDataType.to_proto(resource.data_type)
            )
        else:
            res.ClearField("data_type")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RoutineArguments(
            name=Primitive.from_proto(resource.name),
            argument_kind=RoutineArgumentsArgumentKindEnum.from_proto(
                resource.argument_kind
            ),
            mode=RoutineArgumentsModeEnum.from_proto(resource.mode),
            data_type=RoutineArgumentsDataType.from_proto(resource.data_type),
        )


class RoutineArgumentsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RoutineArguments.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RoutineArguments.from_proto(i) for i in resources]


class RoutineArgumentsDataType(object):
    def __init__(
        self,
        type_kind: str = None,
        array_element_type: dict = None,
        struct_type: dict = None,
    ):
        self.type_kind = type_kind
        self.array_element_type = array_element_type
        self.struct_type = struct_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = routine_pb2.BigqueryBetaRoutineArgumentsDataType()
        if RoutineArgumentsDataTypeTypeKindEnum.to_proto(resource.type_kind):
            res.type_kind = RoutineArgumentsDataTypeTypeKindEnum.to_proto(
                resource.type_kind
            )
        if RoutineArgumentsDataType.to_proto(resource.array_element_type):
            res.array_element_type.CopyFrom(
                RoutineArgumentsDataType.to_proto(resource.array_element_type)
            )
        else:
            res.ClearField("array_element_type")
        if RoutineArgumentsDataTypeStructType.to_proto(resource.struct_type):
            res.struct_type.CopyFrom(
                RoutineArgumentsDataTypeStructType.to_proto(resource.struct_type)
            )
        else:
            res.ClearField("struct_type")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RoutineArgumentsDataType(
            type_kind=RoutineArgumentsDataTypeTypeKindEnum.from_proto(
                resource.type_kind
            ),
            array_element_type=RoutineArgumentsDataType.from_proto(
                resource.array_element_type
            ),
            struct_type=RoutineArgumentsDataTypeStructType.from_proto(
                resource.struct_type
            ),
        )


class RoutineArgumentsDataTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RoutineArgumentsDataType.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RoutineArgumentsDataType.from_proto(i) for i in resources]


class RoutineArgumentsDataTypeStructType(object):
    def __init__(self, fields: list = None):
        self.fields = fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = routine_pb2.BigqueryBetaRoutineArgumentsDataTypeStructType()
        if RoutineArgumentsDataTypeStructTypeFieldsArray.to_proto(resource.fields):
            res.fields.extend(
                RoutineArgumentsDataTypeStructTypeFieldsArray.to_proto(resource.fields)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RoutineArgumentsDataTypeStructType(
            fields=RoutineArgumentsDataTypeStructTypeFieldsArray.from_proto(
                resource.fields
            ),
        )


class RoutineArgumentsDataTypeStructTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RoutineArgumentsDataTypeStructType.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RoutineArgumentsDataTypeStructType.from_proto(i) for i in resources]


class RoutineArgumentsDataTypeStructTypeFields(object):
    def __init__(self, name: str = None, type: dict = None):
        self.name = name
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = routine_pb2.BigqueryBetaRoutineArgumentsDataTypeStructTypeFields()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if RoutineArgumentsDataType.to_proto(resource.type):
            res.type.CopyFrom(RoutineArgumentsDataType.to_proto(resource.type))
        else:
            res.ClearField("type")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RoutineArgumentsDataTypeStructTypeFields(
            name=Primitive.from_proto(resource.name),
            type=RoutineArgumentsDataType.from_proto(resource.type),
        )


class RoutineArgumentsDataTypeStructTypeFieldsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RoutineArgumentsDataTypeStructTypeFields.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            RoutineArgumentsDataTypeStructTypeFields.from_proto(i) for i in resources
        ]


class RoutineRoutineTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineRoutineTypeEnum.Value(
            "BigqueryBetaRoutineRoutineTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineRoutineTypeEnum.Name(resource)[
            len("BigqueryBetaRoutineRoutineTypeEnum") :
        ]


class RoutineLanguageEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineLanguageEnum.Value(
            "BigqueryBetaRoutineLanguageEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineLanguageEnum.Name(resource)[
            len("BigqueryBetaRoutineLanguageEnum") :
        ]


class RoutineArgumentsArgumentKindEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineArgumentsArgumentKindEnum.Value(
            "BigqueryBetaRoutineArgumentsArgumentKindEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineArgumentsArgumentKindEnum.Name(resource)[
            len("BigqueryBetaRoutineArgumentsArgumentKindEnum") :
        ]


class RoutineArgumentsModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineArgumentsModeEnum.Value(
            "BigqueryBetaRoutineArgumentsModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineArgumentsModeEnum.Name(resource)[
            len("BigqueryBetaRoutineArgumentsModeEnum") :
        ]


class RoutineArgumentsDataTypeTypeKindEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum.Value(
            "BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum.Name(
            resource
        )[len("BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum") :]


class RoutineDeterminismLevelEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineDeterminismLevelEnum.Value(
            "BigqueryBetaRoutineDeterminismLevelEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return routine_pb2.BigqueryBetaRoutineDeterminismLevelEnum.Name(resource)[
            len("BigqueryBetaRoutineDeterminismLevelEnum") :
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
