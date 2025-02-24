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
from google3.cloud.graphite.mmv2.services.google.healthcare import fhir_store_pb2
from google3.cloud.graphite.mmv2.services.google.healthcare import fhir_store_pb2_grpc

from typing import List


class FhirStore(object):
    def __init__(
        self,
        name: str = None,
        enable_update_create: bool = None,
        notification_config: dict = None,
        disable_referential_integrity: bool = None,
        shard_num: int = None,
        disable_resource_versioning: bool = None,
        labels: dict = None,
        version: str = None,
        stream_configs: list = None,
        validation_config: dict = None,
        default_search_handling_strict: bool = None,
        complex_data_type_reference_parsing: str = None,
        project: str = None,
        location: str = None,
        dataset: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.enable_update_create = enable_update_create
        self.notification_config = notification_config
        self.disable_referential_integrity = disable_referential_integrity
        self.shard_num = shard_num
        self.disable_resource_versioning = disable_resource_versioning
        self.labels = labels
        self.version = version
        self.stream_configs = stream_configs
        self.validation_config = validation_config
        self.default_search_handling_strict = default_search_handling_strict
        self.complex_data_type_reference_parsing = complex_data_type_reference_parsing
        self.project = project
        self.location = location
        self.dataset = dataset
        self.service_account_file = service_account_file

    def apply(self):
        stub = fhir_store_pb2_grpc.HealthcareFhirStoreServiceStub(channel.Channel())
        request = fhir_store_pb2.ApplyHealthcareFhirStoreRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.enable_update_create):
            request.resource.enable_update_create = Primitive.to_proto(
                self.enable_update_create
            )

        if FhirStoreNotificationConfig.to_proto(self.notification_config):
            request.resource.notification_config.CopyFrom(
                FhirStoreNotificationConfig.to_proto(self.notification_config)
            )
        else:
            request.resource.ClearField("notification_config")
        if Primitive.to_proto(self.disable_referential_integrity):
            request.resource.disable_referential_integrity = Primitive.to_proto(
                self.disable_referential_integrity
            )

        if Primitive.to_proto(self.shard_num):
            request.resource.shard_num = Primitive.to_proto(self.shard_num)

        if Primitive.to_proto(self.disable_resource_versioning):
            request.resource.disable_resource_versioning = Primitive.to_proto(
                self.disable_resource_versioning
            )

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if FhirStoreVersionEnum.to_proto(self.version):
            request.resource.version = FhirStoreVersionEnum.to_proto(self.version)

        if FhirStoreStreamConfigsArray.to_proto(self.stream_configs):
            request.resource.stream_configs.extend(
                FhirStoreStreamConfigsArray.to_proto(self.stream_configs)
            )
        if FhirStoreValidationConfig.to_proto(self.validation_config):
            request.resource.validation_config.CopyFrom(
                FhirStoreValidationConfig.to_proto(self.validation_config)
            )
        else:
            request.resource.ClearField("validation_config")
        if Primitive.to_proto(self.default_search_handling_strict):
            request.resource.default_search_handling_strict = Primitive.to_proto(
                self.default_search_handling_strict
            )

        if FhirStoreComplexDataTypeReferenceParsingEnum.to_proto(
            self.complex_data_type_reference_parsing
        ):
            request.resource.complex_data_type_reference_parsing = (
                FhirStoreComplexDataTypeReferenceParsingEnum.to_proto(
                    self.complex_data_type_reference_parsing
                )
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.dataset):
            request.resource.dataset = Primitive.to_proto(self.dataset)

        request.service_account_file = self.service_account_file

        response = stub.ApplyHealthcareFhirStore(request)
        self.name = Primitive.from_proto(response.name)
        self.enable_update_create = Primitive.from_proto(response.enable_update_create)
        self.notification_config = FhirStoreNotificationConfig.from_proto(
            response.notification_config
        )
        self.disable_referential_integrity = Primitive.from_proto(
            response.disable_referential_integrity
        )
        self.shard_num = Primitive.from_proto(response.shard_num)
        self.disable_resource_versioning = Primitive.from_proto(
            response.disable_resource_versioning
        )
        self.labels = Primitive.from_proto(response.labels)
        self.version = FhirStoreVersionEnum.from_proto(response.version)
        self.stream_configs = FhirStoreStreamConfigsArray.from_proto(
            response.stream_configs
        )
        self.validation_config = FhirStoreValidationConfig.from_proto(
            response.validation_config
        )
        self.default_search_handling_strict = Primitive.from_proto(
            response.default_search_handling_strict
        )
        self.complex_data_type_reference_parsing = (
            FhirStoreComplexDataTypeReferenceParsingEnum.from_proto(
                response.complex_data_type_reference_parsing
            )
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.dataset = Primitive.from_proto(response.dataset)

    def delete(self):
        stub = fhir_store_pb2_grpc.HealthcareFhirStoreServiceStub(channel.Channel())
        request = fhir_store_pb2.DeleteHealthcareFhirStoreRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.enable_update_create):
            request.resource.enable_update_create = Primitive.to_proto(
                self.enable_update_create
            )

        if FhirStoreNotificationConfig.to_proto(self.notification_config):
            request.resource.notification_config.CopyFrom(
                FhirStoreNotificationConfig.to_proto(self.notification_config)
            )
        else:
            request.resource.ClearField("notification_config")
        if Primitive.to_proto(self.disable_referential_integrity):
            request.resource.disable_referential_integrity = Primitive.to_proto(
                self.disable_referential_integrity
            )

        if Primitive.to_proto(self.shard_num):
            request.resource.shard_num = Primitive.to_proto(self.shard_num)

        if Primitive.to_proto(self.disable_resource_versioning):
            request.resource.disable_resource_versioning = Primitive.to_proto(
                self.disable_resource_versioning
            )

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if FhirStoreVersionEnum.to_proto(self.version):
            request.resource.version = FhirStoreVersionEnum.to_proto(self.version)

        if FhirStoreStreamConfigsArray.to_proto(self.stream_configs):
            request.resource.stream_configs.extend(
                FhirStoreStreamConfigsArray.to_proto(self.stream_configs)
            )
        if FhirStoreValidationConfig.to_proto(self.validation_config):
            request.resource.validation_config.CopyFrom(
                FhirStoreValidationConfig.to_proto(self.validation_config)
            )
        else:
            request.resource.ClearField("validation_config")
        if Primitive.to_proto(self.default_search_handling_strict):
            request.resource.default_search_handling_strict = Primitive.to_proto(
                self.default_search_handling_strict
            )

        if FhirStoreComplexDataTypeReferenceParsingEnum.to_proto(
            self.complex_data_type_reference_parsing
        ):
            request.resource.complex_data_type_reference_parsing = (
                FhirStoreComplexDataTypeReferenceParsingEnum.to_proto(
                    self.complex_data_type_reference_parsing
                )
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.dataset):
            request.resource.dataset = Primitive.to_proto(self.dataset)

        response = stub.DeleteHealthcareFhirStore(request)

    @classmethod
    def list(self, project, location, dataset, service_account_file=""):
        stub = fhir_store_pb2_grpc.HealthcareFhirStoreServiceStub(channel.Channel())
        request = fhir_store_pb2.ListHealthcareFhirStoreRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Dataset = dataset

        return stub.ListHealthcareFhirStore(request).items

    def to_proto(self):
        resource = fhir_store_pb2.HealthcareFhirStore()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.enable_update_create):
            resource.enable_update_create = Primitive.to_proto(
                self.enable_update_create
            )
        if FhirStoreNotificationConfig.to_proto(self.notification_config):
            resource.notification_config.CopyFrom(
                FhirStoreNotificationConfig.to_proto(self.notification_config)
            )
        else:
            resource.ClearField("notification_config")
        if Primitive.to_proto(self.disable_referential_integrity):
            resource.disable_referential_integrity = Primitive.to_proto(
                self.disable_referential_integrity
            )
        if Primitive.to_proto(self.shard_num):
            resource.shard_num = Primitive.to_proto(self.shard_num)
        if Primitive.to_proto(self.disable_resource_versioning):
            resource.disable_resource_versioning = Primitive.to_proto(
                self.disable_resource_versioning
            )
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if FhirStoreVersionEnum.to_proto(self.version):
            resource.version = FhirStoreVersionEnum.to_proto(self.version)
        if FhirStoreStreamConfigsArray.to_proto(self.stream_configs):
            resource.stream_configs.extend(
                FhirStoreStreamConfigsArray.to_proto(self.stream_configs)
            )
        if FhirStoreValidationConfig.to_proto(self.validation_config):
            resource.validation_config.CopyFrom(
                FhirStoreValidationConfig.to_proto(self.validation_config)
            )
        else:
            resource.ClearField("validation_config")
        if Primitive.to_proto(self.default_search_handling_strict):
            resource.default_search_handling_strict = Primitive.to_proto(
                self.default_search_handling_strict
            )
        if FhirStoreComplexDataTypeReferenceParsingEnum.to_proto(
            self.complex_data_type_reference_parsing
        ):
            resource.complex_data_type_reference_parsing = (
                FhirStoreComplexDataTypeReferenceParsingEnum.to_proto(
                    self.complex_data_type_reference_parsing
                )
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.dataset):
            resource.dataset = Primitive.to_proto(self.dataset)
        return resource


class FhirStoreNotificationConfig(object):
    def __init__(self, pubsub_topic: str = None):
        self.pubsub_topic = pubsub_topic

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = fhir_store_pb2.HealthcareFhirStoreNotificationConfig()
        if Primitive.to_proto(resource.pubsub_topic):
            res.pubsub_topic = Primitive.to_proto(resource.pubsub_topic)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FhirStoreNotificationConfig(
            pubsub_topic=Primitive.from_proto(resource.pubsub_topic),
        )


class FhirStoreNotificationConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FhirStoreNotificationConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FhirStoreNotificationConfig.from_proto(i) for i in resources]


class FhirStoreStreamConfigs(object):
    def __init__(self, resource_types: list = None, bigquery_destination: dict = None):
        self.resource_types = resource_types
        self.bigquery_destination = bigquery_destination

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = fhir_store_pb2.HealthcareFhirStoreStreamConfigs()
        if Primitive.to_proto(resource.resource_types):
            res.resource_types.extend(Primitive.to_proto(resource.resource_types))
        if FhirStoreStreamConfigsBigqueryDestination.to_proto(
            resource.bigquery_destination
        ):
            res.bigquery_destination.CopyFrom(
                FhirStoreStreamConfigsBigqueryDestination.to_proto(
                    resource.bigquery_destination
                )
            )
        else:
            res.ClearField("bigquery_destination")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FhirStoreStreamConfigs(
            resource_types=Primitive.from_proto(resource.resource_types),
            bigquery_destination=FhirStoreStreamConfigsBigqueryDestination.from_proto(
                resource.bigquery_destination
            ),
        )


class FhirStoreStreamConfigsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FhirStoreStreamConfigs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FhirStoreStreamConfigs.from_proto(i) for i in resources]


class FhirStoreStreamConfigsBigqueryDestination(object):
    def __init__(self, dataset_uri: str = None, schema_config: dict = None):
        self.dataset_uri = dataset_uri
        self.schema_config = schema_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = fhir_store_pb2.HealthcareFhirStoreStreamConfigsBigqueryDestination()
        if Primitive.to_proto(resource.dataset_uri):
            res.dataset_uri = Primitive.to_proto(resource.dataset_uri)
        if FhirStoreStreamConfigsBigqueryDestinationSchemaConfig.to_proto(
            resource.schema_config
        ):
            res.schema_config.CopyFrom(
                FhirStoreStreamConfigsBigqueryDestinationSchemaConfig.to_proto(
                    resource.schema_config
                )
            )
        else:
            res.ClearField("schema_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FhirStoreStreamConfigsBigqueryDestination(
            dataset_uri=Primitive.from_proto(resource.dataset_uri),
            schema_config=FhirStoreStreamConfigsBigqueryDestinationSchemaConfig.from_proto(
                resource.schema_config
            ),
        )


class FhirStoreStreamConfigsBigqueryDestinationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            FhirStoreStreamConfigsBigqueryDestination.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            FhirStoreStreamConfigsBigqueryDestination.from_proto(i) for i in resources
        ]


class FhirStoreStreamConfigsBigqueryDestinationSchemaConfig(object):
    def __init__(self, schema_type: str = None, recursive_structure_depth: int = None):
        self.schema_type = schema_type
        self.recursive_structure_depth = recursive_structure_depth

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            fhir_store_pb2.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig()
        )
        if FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum.to_proto(
            resource.schema_type
        ):
            res.schema_type = FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum.to_proto(
                resource.schema_type
            )
        if Primitive.to_proto(resource.recursive_structure_depth):
            res.recursive_structure_depth = Primitive.to_proto(
                resource.recursive_structure_depth
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FhirStoreStreamConfigsBigqueryDestinationSchemaConfig(
            schema_type=FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum.from_proto(
                resource.schema_type
            ),
            recursive_structure_depth=Primitive.from_proto(
                resource.recursive_structure_depth
            ),
        )


class FhirStoreStreamConfigsBigqueryDestinationSchemaConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            FhirStoreStreamConfigsBigqueryDestinationSchemaConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            FhirStoreStreamConfigsBigqueryDestinationSchemaConfig.from_proto(i)
            for i in resources
        ]


class FhirStoreValidationConfig(object):
    def __init__(
        self,
        disable_profile_validation: bool = None,
        enabled_implementation_guides: list = None,
        disable_required_field_validation: bool = None,
        disable_reference_type_validation: bool = None,
        disable_fhirpath_validation: bool = None,
    ):
        self.disable_profile_validation = disable_profile_validation
        self.enabled_implementation_guides = enabled_implementation_guides
        self.disable_required_field_validation = disable_required_field_validation
        self.disable_reference_type_validation = disable_reference_type_validation
        self.disable_fhirpath_validation = disable_fhirpath_validation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = fhir_store_pb2.HealthcareFhirStoreValidationConfig()
        if Primitive.to_proto(resource.disable_profile_validation):
            res.disable_profile_validation = Primitive.to_proto(
                resource.disable_profile_validation
            )
        if Primitive.to_proto(resource.enabled_implementation_guides):
            res.enabled_implementation_guides.extend(
                Primitive.to_proto(resource.enabled_implementation_guides)
            )
        if Primitive.to_proto(resource.disable_required_field_validation):
            res.disable_required_field_validation = Primitive.to_proto(
                resource.disable_required_field_validation
            )
        if Primitive.to_proto(resource.disable_reference_type_validation):
            res.disable_reference_type_validation = Primitive.to_proto(
                resource.disable_reference_type_validation
            )
        if Primitive.to_proto(resource.disable_fhirpath_validation):
            res.disable_fhirpath_validation = Primitive.to_proto(
                resource.disable_fhirpath_validation
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FhirStoreValidationConfig(
            disable_profile_validation=Primitive.from_proto(
                resource.disable_profile_validation
            ),
            enabled_implementation_guides=Primitive.from_proto(
                resource.enabled_implementation_guides
            ),
            disable_required_field_validation=Primitive.from_proto(
                resource.disable_required_field_validation
            ),
            disable_reference_type_validation=Primitive.from_proto(
                resource.disable_reference_type_validation
            ),
            disable_fhirpath_validation=Primitive.from_proto(
                resource.disable_fhirpath_validation
            ),
        )


class FhirStoreValidationConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FhirStoreValidationConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FhirStoreValidationConfig.from_proto(i) for i in resources]


class FhirStoreVersionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return fhir_store_pb2.HealthcareFhirStoreVersionEnum.Value(
            "HealthcareFhirStoreVersionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return fhir_store_pb2.HealthcareFhirStoreVersionEnum.Name(resource)[
            len("HealthcareFhirStoreVersionEnum") :
        ]


class FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return fhir_store_pb2.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum.Value(
            "HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return fhir_store_pb2.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum.Name(
            resource
        )[
            len(
                "HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum"
            ) :
        ]


class FhirStoreComplexDataTypeReferenceParsingEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            fhir_store_pb2.HealthcareFhirStoreComplexDataTypeReferenceParsingEnum.Value(
                "HealthcareFhirStoreComplexDataTypeReferenceParsingEnum%s" % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            fhir_store_pb2.HealthcareFhirStoreComplexDataTypeReferenceParsingEnum.Name(
                resource
            )[len("HealthcareFhirStoreComplexDataTypeReferenceParsingEnum") :]
        )


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
