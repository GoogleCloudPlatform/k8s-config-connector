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
from google3.cloud.graphite.mmv2.services.google.dlp import deidentify_template_pb2
from google3.cloud.graphite.mmv2.services.google.dlp import deidentify_template_pb2_grpc

from typing import List


class DeidentifyTemplate(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        deidentify_config: dict = None,
        location_id: str = None,
        parent: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.description = description
        self.deidentify_config = deidentify_config
        self.parent = parent
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = deidentify_template_pb2_grpc.DlpBetaDeidentifyTemplateServiceStub(
            channel.Channel()
        )
        request = deidentify_template_pb2.ApplyDlpBetaDeidentifyTemplateRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if DeidentifyTemplateDeidentifyConfig.to_proto(self.deidentify_config):
            request.resource.deidentify_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfig.to_proto(self.deidentify_config)
            )
        else:
            request.resource.ClearField("deidentify_config")
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDlpBetaDeidentifyTemplate(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.deidentify_config = DeidentifyTemplateDeidentifyConfig.from_proto(
            response.deidentify_config
        )
        self.location_id = Primitive.from_proto(response.location_id)
        self.parent = Primitive.from_proto(response.parent)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = deidentify_template_pb2_grpc.DlpBetaDeidentifyTemplateServiceStub(
            channel.Channel()
        )
        request = deidentify_template_pb2.DeleteDlpBetaDeidentifyTemplateRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if DeidentifyTemplateDeidentifyConfig.to_proto(self.deidentify_config):
            request.resource.deidentify_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfig.to_proto(self.deidentify_config)
            )
        else:
            request.resource.ClearField("deidentify_config")
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteDlpBetaDeidentifyTemplate(request)

    @classmethod
    def list(self, location, parent, service_account_file=""):
        stub = deidentify_template_pb2_grpc.DlpBetaDeidentifyTemplateServiceStub(
            channel.Channel()
        )
        request = deidentify_template_pb2.ListDlpBetaDeidentifyTemplateRequest()
        request.service_account_file = service_account_file
        request.Location = location

        request.Parent = parent

        return stub.ListDlpBetaDeidentifyTemplate(request).items

    def to_proto(self):
        resource = deidentify_template_pb2.DlpBetaDeidentifyTemplate()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if DeidentifyTemplateDeidentifyConfig.to_proto(self.deidentify_config):
            resource.deidentify_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfig.to_proto(self.deidentify_config)
            )
        else:
            resource.ClearField("deidentify_config")
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class DeidentifyTemplateDeidentifyConfig(object):
    def __init__(
        self,
        info_type_transformations: dict = None,
        record_transformations: dict = None,
        transformation_error_handling: dict = None,
    ):
        self.info_type_transformations = info_type_transformations
        self.record_transformations = record_transformations
        self.transformation_error_handling = transformation_error_handling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfig()
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformations.to_proto(
            resource.info_type_transformations
        ):
            res.info_type_transformations.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformations.to_proto(
                    resource.info_type_transformations
                )
            )
        else:
            res.ClearField("info_type_transformations")
        if DeidentifyTemplateDeidentifyConfigRecordTransformations.to_proto(
            resource.record_transformations
        ):
            res.record_transformations.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformations.to_proto(
                    resource.record_transformations
                )
            )
        else:
            res.ClearField("record_transformations")
        if DeidentifyTemplateDeidentifyConfigTransformationErrorHandling.to_proto(
            resource.transformation_error_handling
        ):
            res.transformation_error_handling.CopyFrom(
                DeidentifyTemplateDeidentifyConfigTransformationErrorHandling.to_proto(
                    resource.transformation_error_handling
                )
            )
        else:
            res.ClearField("transformation_error_handling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfig(
            info_type_transformations=DeidentifyTemplateDeidentifyConfigInfoTypeTransformations.from_proto(
                resource.info_type_transformations
            ),
            record_transformations=DeidentifyTemplateDeidentifyConfigRecordTransformations.from_proto(
                resource.record_transformations
            ),
            transformation_error_handling=DeidentifyTemplateDeidentifyConfigTransformationErrorHandling.from_proto(
                resource.transformation_error_handling
            ),
        )


class DeidentifyTemplateDeidentifyConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DeidentifyTemplateDeidentifyConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DeidentifyTemplateDeidentifyConfig.from_proto(i) for i in resources]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformations(object):
    def __init__(self, transformations: list = None):
        self.transformations = transformations

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformations()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsArray.to_proto(
            resource.transformations
        ):
            res.transformations.extend(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsArray.to_proto(
                    resource.transformations
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformations(
            transformations=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsArray.from_proto(
                resource.transformations
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformations.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformations.from_proto(i)
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformations(object):
    def __init__(self, info_types: list = None, primitive_transformation: dict = None):
        self.info_types = info_types
        self.primitive_transformation = primitive_transformation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformations()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypesArray.to_proto(
            resource.info_types
        ):
            res.info_types.extend(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypesArray.to_proto(
                    resource.info_types
                )
            )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation.to_proto(
            resource.primitive_transformation
        ):
            res.primitive_transformation.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation.to_proto(
                    resource.primitive_transformation
                )
            )
        else:
            res.ClearField("primitive_transformation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformations(
            info_types=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypesArray.from_proto(
                resource.info_types
            ),
            primitive_transformation=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation.from_proto(
                resource.primitive_transformation
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformations.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformations.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation(
    object
):
    def __init__(
        self,
        replace_config: dict = None,
        redact_config: dict = None,
        character_mask_config: dict = None,
        crypto_replace_ffx_fpe_config: dict = None,
        fixed_size_bucketing_config: dict = None,
        bucketing_config: dict = None,
        replace_with_info_type_config: dict = None,
        time_part_config: dict = None,
        crypto_hash_config: dict = None,
        date_shift_config: dict = None,
        crypto_deterministic_config: dict = None,
    ):
        self.replace_config = replace_config
        self.redact_config = redact_config
        self.character_mask_config = character_mask_config
        self.crypto_replace_ffx_fpe_config = crypto_replace_ffx_fpe_config
        self.fixed_size_bucketing_config = fixed_size_bucketing_config
        self.bucketing_config = bucketing_config
        self.replace_with_info_type_config = replace_with_info_type_config
        self.time_part_config = time_part_config
        self.crypto_hash_config = crypto_hash_config
        self.date_shift_config = date_shift_config
        self.crypto_deterministic_config = crypto_deterministic_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig.to_proto(
            resource.replace_config
        ):
            res.replace_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig.to_proto(
                    resource.replace_config
                )
            )
        else:
            res.ClearField("replace_config")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig.to_proto(
            resource.redact_config
        ):
            res.redact_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig.to_proto(
                    resource.redact_config
                )
            )
        else:
            res.ClearField("redact_config")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig.to_proto(
            resource.character_mask_config
        ):
            res.character_mask_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig.to_proto(
                    resource.character_mask_config
                )
            )
        else:
            res.ClearField("character_mask_config")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.to_proto(
            resource.crypto_replace_ffx_fpe_config
        ):
            res.crypto_replace_ffx_fpe_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.to_proto(
                    resource.crypto_replace_ffx_fpe_config
                )
            )
        else:
            res.ClearField("crypto_replace_ffx_fpe_config")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig.to_proto(
            resource.fixed_size_bucketing_config
        ):
            res.fixed_size_bucketing_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig.to_proto(
                    resource.fixed_size_bucketing_config
                )
            )
        else:
            res.ClearField("fixed_size_bucketing_config")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig.to_proto(
            resource.bucketing_config
        ):
            res.bucketing_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig.to_proto(
                    resource.bucketing_config
                )
            )
        else:
            res.ClearField("bucketing_config")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.to_proto(
            resource.replace_with_info_type_config
        ):
            res.replace_with_info_type_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.to_proto(
                    resource.replace_with_info_type_config
                )
            )
        else:
            res.ClearField("replace_with_info_type_config")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig.to_proto(
            resource.time_part_config
        ):
            res.time_part_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig.to_proto(
                    resource.time_part_config
                )
            )
        else:
            res.ClearField("time_part_config")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig.to_proto(
            resource.crypto_hash_config
        ):
            res.crypto_hash_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig.to_proto(
                    resource.crypto_hash_config
                )
            )
        else:
            res.ClearField("crypto_hash_config")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig.to_proto(
            resource.date_shift_config
        ):
            res.date_shift_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig.to_proto(
                    resource.date_shift_config
                )
            )
        else:
            res.ClearField("date_shift_config")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig.to_proto(
            resource.crypto_deterministic_config
        ):
            res.crypto_deterministic_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig.to_proto(
                    resource.crypto_deterministic_config
                )
            )
        else:
            res.ClearField("crypto_deterministic_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation(
            replace_config=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig.from_proto(
                resource.replace_config
            ),
            redact_config=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig.from_proto(
                resource.redact_config
            ),
            character_mask_config=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig.from_proto(
                resource.character_mask_config
            ),
            crypto_replace_ffx_fpe_config=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.from_proto(
                resource.crypto_replace_ffx_fpe_config
            ),
            fixed_size_bucketing_config=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig.from_proto(
                resource.fixed_size_bucketing_config
            ),
            bucketing_config=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig.from_proto(
                resource.bucketing_config
            ),
            replace_with_info_type_config=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.from_proto(
                resource.replace_with_info_type_config
            ),
            time_part_config=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig.from_proto(
                resource.time_part_config
            ),
            crypto_hash_config=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig.from_proto(
                resource.crypto_hash_config
            ),
            date_shift_config=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig.from_proto(
                resource.date_shift_config
            ),
            crypto_deterministic_config=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig.from_proto(
                resource.crypto_deterministic_config
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig(
    object
):
    def __init__(self, new_value: dict = None):
        self.new_value = new_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue.to_proto(
            resource.new_value
        ):
            res.new_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue.to_proto(
                    resource.new_value
                )
            )
        else:
            res.ClearField("new_value")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig(
            new_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue.from_proto(
                resource.new_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig()
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig(
    object
):
    def __init__(
        self,
        masking_character: str = None,
        number_to_mask: int = None,
        reverse_order: bool = None,
        characters_to_ignore: list = None,
    ):
        self.masking_character = masking_character
        self.number_to_mask = number_to_mask
        self.reverse_order = reverse_order
        self.characters_to_ignore = characters_to_ignore

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig()
        )
        if Primitive.to_proto(resource.masking_character):
            res.masking_character = Primitive.to_proto(resource.masking_character)
        if Primitive.to_proto(resource.number_to_mask):
            res.number_to_mask = Primitive.to_proto(resource.number_to_mask)
        if Primitive.to_proto(resource.reverse_order):
            res.reverse_order = Primitive.to_proto(resource.reverse_order)
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray.to_proto(
            resource.characters_to_ignore
        ):
            res.characters_to_ignore.extend(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray.to_proto(
                    resource.characters_to_ignore
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig(
            masking_character=Primitive.from_proto(resource.masking_character),
            number_to_mask=Primitive.from_proto(resource.number_to_mask),
            reverse_order=Primitive.from_proto(resource.reverse_order),
            characters_to_ignore=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray.from_proto(
                resource.characters_to_ignore
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore(
    object
):
    def __init__(
        self, characters_to_skip: str = None, common_characters_to_ignore: str = None
    ):
        self.characters_to_skip = characters_to_skip
        self.common_characters_to_ignore = common_characters_to_ignore

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore()
        )
        if Primitive.to_proto(resource.characters_to_skip):
            res.characters_to_skip = Primitive.to_proto(resource.characters_to_skip)
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.to_proto(
            resource.common_characters_to_ignore
        ):
            res.common_characters_to_ignore = DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.to_proto(
                resource.common_characters_to_ignore
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore(
            characters_to_skip=Primitive.from_proto(resource.characters_to_skip),
            common_characters_to_ignore=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.from_proto(
                resource.common_characters_to_ignore
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig(
    object
):
    def __init__(
        self,
        crypto_key: dict = None,
        context: dict = None,
        common_alphabet: str = None,
        custom_alphabet: str = None,
        radix: int = None,
        surrogate_info_type: dict = None,
    ):
        self.crypto_key = crypto_key
        self.context = context
        self.common_alphabet = common_alphabet
        self.custom_alphabet = custom_alphabet
        self.radix = radix
        self.surrogate_info_type = surrogate_info_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.to_proto(
            resource.context
        ):
            res.context.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.to_proto(
                    resource.context
                )
            )
        else:
            res.ClearField("context")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.to_proto(
            resource.common_alphabet
        ):
            res.common_alphabet = DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.to_proto(
                resource.common_alphabet
            )
        if Primitive.to_proto(resource.custom_alphabet):
            res.custom_alphabet = Primitive.to_proto(resource.custom_alphabet)
        if Primitive.to_proto(resource.radix):
            res.radix = Primitive.to_proto(resource.radix)
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.to_proto(
            resource.surrogate_info_type
        ):
            res.surrogate_info_type.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.to_proto(
                    resource.surrogate_info_type
                )
            )
        else:
            res.ClearField("surrogate_info_type")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig(
            crypto_key=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
            context=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.from_proto(
                resource.context
            ),
            common_alphabet=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.from_proto(
                resource.common_alphabet
            ),
            custom_alphabet=Primitive.from_proto(resource.custom_alphabet),
            radix=Primitive.from_proto(resource.radix),
            surrogate_info_type=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.from_proto(
                resource.surrogate_info_type
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContextArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoTypeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig(
    object
):
    def __init__(
        self,
        lower_bound: dict = None,
        upper_bound: dict = None,
        bucket_size: float = None,
    ):
        self.lower_bound = lower_bound
        self.upper_bound = upper_bound
        self.bucket_size = bucket_size

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.to_proto(
            resource.lower_bound
        ):
            res.lower_bound.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.to_proto(
                    resource.lower_bound
                )
            )
        else:
            res.ClearField("lower_bound")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.to_proto(
            resource.upper_bound
        ):
            res.upper_bound.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.to_proto(
                    resource.upper_bound
                )
            )
        else:
            res.ClearField("upper_bound")
        if Primitive.to_proto(resource.bucket_size):
            res.bucket_size = Primitive.to_proto(resource.bucket_size)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig(
            lower_bound=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.from_proto(
                resource.lower_bound
            ),
            upper_bound=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.from_proto(
                resource.upper_bound
            ),
            bucket_size=Primitive.from_proto(resource.bucket_size),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig(
    object
):
    def __init__(self, buckets: list = None):
        self.buckets = buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsArray.to_proto(
            resource.buckets
        ):
            res.buckets.extend(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsArray.to_proto(
                    resource.buckets
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig(
            buckets=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsArray.from_proto(
                resource.buckets
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets(
    object
):
    def __init__(
        self, min: dict = None, max: dict = None, replacement_value: dict = None
    ):
        self.min = min
        self.max = max
        self.replacement_value = replacement_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin.to_proto(
            resource.min
        ):
            res.min.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin.to_proto(
                    resource.min
                )
            )
        else:
            res.ClearField("min")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax.to_proto(
            resource.max
        ):
            res.max.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax.to_proto(
                    resource.max
                )
            )
        else:
            res.ClearField("max")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.to_proto(
            resource.replacement_value
        ):
            res.replacement_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.to_proto(
                    resource.replacement_value
                )
            )
        else:
            res.ClearField("replacement_value")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets(
            min=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin.from_proto(
                resource.min
            ),
            max=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax.from_proto(
                resource.max
            ),
            replacement_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.from_proto(
                resource.replacement_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig()
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig(
    object
):
    def __init__(self, part_to_extract: str = None):
        self.part_to_extract = part_to_extract

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.to_proto(
            resource.part_to_extract
        ):
            res.part_to_extract = DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.to_proto(
                resource.part_to_extract
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig(
            part_to_extract=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.from_proto(
                resource.part_to_extract
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig(
    object
):
    def __init__(self, crypto_key: dict = None):
        self.crypto_key = crypto_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig(
            crypto_key=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig(
    object
):
    def __init__(
        self,
        upper_bound_days: int = None,
        lower_bound_days: int = None,
        context: dict = None,
        crypto_key: dict = None,
    ):
        self.upper_bound_days = upper_bound_days
        self.lower_bound_days = lower_bound_days
        self.context = context
        self.crypto_key = crypto_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig()
        )
        if Primitive.to_proto(resource.upper_bound_days):
            res.upper_bound_days = Primitive.to_proto(resource.upper_bound_days)
        if Primitive.to_proto(resource.lower_bound_days):
            res.lower_bound_days = Primitive.to_proto(resource.lower_bound_days)
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext.to_proto(
            resource.context
        ):
            res.context.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext.to_proto(
                    resource.context
                )
            )
        else:
            res.ClearField("context")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig(
            upper_bound_days=Primitive.from_proto(resource.upper_bound_days),
            lower_bound_days=Primitive.from_proto(resource.lower_bound_days),
            context=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext.from_proto(
                resource.context
            ),
            crypto_key=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContextArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig(
    object
):
    def __init__(
        self,
        crypto_key: dict = None,
        surrogate_info_type: dict = None,
        context: dict = None,
    ):
        self.crypto_key = crypto_key
        self.surrogate_info_type = surrogate_info_type
        self.context = context

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.to_proto(
            resource.surrogate_info_type
        ):
            res.surrogate_info_type.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.to_proto(
                    resource.surrogate_info_type
                )
            )
        else:
            res.ClearField("surrogate_info_type")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.to_proto(
            resource.context
        ):
            res.context.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.to_proto(
                    resource.context
                )
            )
        else:
            res.ClearField("context")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig(
            crypto_key=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
            surrogate_info_type=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.from_proto(
                resource.surrogate_info_type
            ),
            context=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.from_proto(
                resource.context
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoTypeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContextArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformations(object):
    def __init__(
        self, field_transformations: list = None, record_suppressions: list = None
    ):
        self.field_transformations = field_transformations
        self.record_suppressions = record_suppressions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformations()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsArray.to_proto(
            resource.field_transformations
        ):
            res.field_transformations.extend(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsArray.to_proto(
                    resource.field_transformations
                )
            )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsArray.to_proto(
            resource.record_suppressions
        ):
            res.record_suppressions.extend(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsArray.to_proto(
                    resource.record_suppressions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformations(
            field_transformations=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsArray.from_proto(
                resource.field_transformations
            ),
            record_suppressions=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsArray.from_proto(
                resource.record_suppressions
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformations.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformations.from_proto(i)
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformations(
    object
):
    def __init__(
        self,
        fields: list = None,
        condition: dict = None,
        primitive_transformation: dict = None,
        info_type_transformations: dict = None,
    ):
        self.fields = fields
        self.condition = condition
        self.primitive_transformation = primitive_transformation
        self.info_type_transformations = info_type_transformations

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformations()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsFieldsArray.to_proto(
            resource.fields
        ):
            res.fields.extend(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsFieldsArray.to_proto(
                    resource.fields
                )
            )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition.to_proto(
            resource.condition
        ):
            res.condition.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition.to_proto(
                    resource.condition
                )
            )
        else:
            res.ClearField("condition")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation.to_proto(
            resource.primitive_transformation
        ):
            res.primitive_transformation.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation.to_proto(
                    resource.primitive_transformation
                )
            )
        else:
            res.ClearField("primitive_transformation")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations.to_proto(
            resource.info_type_transformations
        ):
            res.info_type_transformations.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations.to_proto(
                    resource.info_type_transformations
                )
            )
        else:
            res.ClearField("info_type_transformations")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformations(
            fields=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsFieldsArray.from_proto(
                resource.fields
            ),
            condition=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition.from_proto(
                resource.condition
            ),
            primitive_transformation=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation.from_proto(
                resource.primitive_transformation
            ),
            info_type_transformations=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations.from_proto(
                resource.info_type_transformations
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformations.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformations.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsFields(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsFields()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsFields(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsFieldsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsFields.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsFields.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition(
    object
):
    def __init__(self, expressions: dict = None):
        self.expressions = expressions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions.to_proto(
            resource.expressions
        ):
            res.expressions.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions.to_proto(
                    resource.expressions
                )
            )
        else:
            res.ClearField("expressions")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition(
            expressions=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions.from_proto(
                resource.expressions
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsCondition.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions(
    object
):
    def __init__(self, logical_operator: str = None, conditions: dict = None):
        self.logical_operator = logical_operator
        self.conditions = conditions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsLogicalOperatorEnum.to_proto(
            resource.logical_operator
        ):
            res.logical_operator = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsLogicalOperatorEnum.to_proto(
                resource.logical_operator
            )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions.to_proto(
            resource.conditions
        ):
            res.conditions.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions.to_proto(
                    resource.conditions
                )
            )
        else:
            res.ClearField("conditions")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions(
            logical_operator=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsLogicalOperatorEnum.from_proto(
                resource.logical_operator
            ),
            conditions=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions.from_proto(
                resource.conditions
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressions.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions(
    object
):
    def __init__(self, conditions: list = None):
        self.conditions = conditions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsArray.to_proto(
            resource.conditions
        ):
            res.conditions.extend(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsArray.to_proto(
                    resource.conditions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions(
            conditions=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsArray.from_proto(
                resource.conditions
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditions.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions(
    object
):
    def __init__(self, field: dict = None, operator: str = None, value: dict = None):
        self.field = field
        self.operator = operator
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField.to_proto(
            resource.field
        ):
            res.field.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField.to_proto(
                    resource.field
                )
            )
        else:
            res.ClearField("field")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOperatorEnum.to_proto(
            resource.operator
        ):
            res.operator = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOperatorEnum.to_proto(
                resource.operator
            )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue.to_proto(
            resource.value
        ):
            res.value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue.to_proto(
                    resource.value
                )
            )
        else:
            res.ClearField("value")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions(
            field=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField.from_proto(
                resource.field
            ),
            operator=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOperatorEnum.from_proto(
                resource.operator
            ),
            value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue.from_proto(
                resource.value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditions.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsFieldArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation(
    object
):
    def __init__(
        self,
        replace_config: dict = None,
        redact_config: dict = None,
        character_mask_config: dict = None,
        crypto_replace_ffx_fpe_config: dict = None,
        fixed_size_bucketing_config: dict = None,
        bucketing_config: dict = None,
        replace_with_info_type_config: dict = None,
        time_part_config: dict = None,
        crypto_hash_config: dict = None,
        date_shift_config: dict = None,
        crypto_deterministic_config: dict = None,
    ):
        self.replace_config = replace_config
        self.redact_config = redact_config
        self.character_mask_config = character_mask_config
        self.crypto_replace_ffx_fpe_config = crypto_replace_ffx_fpe_config
        self.fixed_size_bucketing_config = fixed_size_bucketing_config
        self.bucketing_config = bucketing_config
        self.replace_with_info_type_config = replace_with_info_type_config
        self.time_part_config = time_part_config
        self.crypto_hash_config = crypto_hash_config
        self.date_shift_config = date_shift_config
        self.crypto_deterministic_config = crypto_deterministic_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig.to_proto(
            resource.replace_config
        ):
            res.replace_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig.to_proto(
                    resource.replace_config
                )
            )
        else:
            res.ClearField("replace_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig.to_proto(
            resource.redact_config
        ):
            res.redact_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig.to_proto(
                    resource.redact_config
                )
            )
        else:
            res.ClearField("redact_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig.to_proto(
            resource.character_mask_config
        ):
            res.character_mask_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig.to_proto(
                    resource.character_mask_config
                )
            )
        else:
            res.ClearField("character_mask_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.to_proto(
            resource.crypto_replace_ffx_fpe_config
        ):
            res.crypto_replace_ffx_fpe_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.to_proto(
                    resource.crypto_replace_ffx_fpe_config
                )
            )
        else:
            res.ClearField("crypto_replace_ffx_fpe_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig.to_proto(
            resource.fixed_size_bucketing_config
        ):
            res.fixed_size_bucketing_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig.to_proto(
                    resource.fixed_size_bucketing_config
                )
            )
        else:
            res.ClearField("fixed_size_bucketing_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig.to_proto(
            resource.bucketing_config
        ):
            res.bucketing_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig.to_proto(
                    resource.bucketing_config
                )
            )
        else:
            res.ClearField("bucketing_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.to_proto(
            resource.replace_with_info_type_config
        ):
            res.replace_with_info_type_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.to_proto(
                    resource.replace_with_info_type_config
                )
            )
        else:
            res.ClearField("replace_with_info_type_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig.to_proto(
            resource.time_part_config
        ):
            res.time_part_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig.to_proto(
                    resource.time_part_config
                )
            )
        else:
            res.ClearField("time_part_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig.to_proto(
            resource.crypto_hash_config
        ):
            res.crypto_hash_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig.to_proto(
                    resource.crypto_hash_config
                )
            )
        else:
            res.ClearField("crypto_hash_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig.to_proto(
            resource.date_shift_config
        ):
            res.date_shift_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig.to_proto(
                    resource.date_shift_config
                )
            )
        else:
            res.ClearField("date_shift_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig.to_proto(
            resource.crypto_deterministic_config
        ):
            res.crypto_deterministic_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig.to_proto(
                    resource.crypto_deterministic_config
                )
            )
        else:
            res.ClearField("crypto_deterministic_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation(
            replace_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig.from_proto(
                resource.replace_config
            ),
            redact_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig.from_proto(
                resource.redact_config
            ),
            character_mask_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig.from_proto(
                resource.character_mask_config
            ),
            crypto_replace_ffx_fpe_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.from_proto(
                resource.crypto_replace_ffx_fpe_config
            ),
            fixed_size_bucketing_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig.from_proto(
                resource.fixed_size_bucketing_config
            ),
            bucketing_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig.from_proto(
                resource.bucketing_config
            ),
            replace_with_info_type_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.from_proto(
                resource.replace_with_info_type_config
            ),
            time_part_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig.from_proto(
                resource.time_part_config
            ),
            crypto_hash_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig.from_proto(
                resource.crypto_hash_config
            ),
            date_shift_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig.from_proto(
                resource.date_shift_config
            ),
            crypto_deterministic_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig.from_proto(
                resource.crypto_deterministic_config
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig(
    object
):
    def __init__(self, new_value: dict = None):
        self.new_value = new_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue.to_proto(
            resource.new_value
        ):
            res.new_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue.to_proto(
                    resource.new_value
                )
            )
        else:
            res.ClearField("new_value")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig(
            new_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue.from_proto(
                resource.new_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig()
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig(
    object
):
    def __init__(
        self,
        masking_character: str = None,
        number_to_mask: int = None,
        reverse_order: bool = None,
        characters_to_ignore: list = None,
    ):
        self.masking_character = masking_character
        self.number_to_mask = number_to_mask
        self.reverse_order = reverse_order
        self.characters_to_ignore = characters_to_ignore

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig()
        )
        if Primitive.to_proto(resource.masking_character):
            res.masking_character = Primitive.to_proto(resource.masking_character)
        if Primitive.to_proto(resource.number_to_mask):
            res.number_to_mask = Primitive.to_proto(resource.number_to_mask)
        if Primitive.to_proto(resource.reverse_order):
            res.reverse_order = Primitive.to_proto(resource.reverse_order)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray.to_proto(
            resource.characters_to_ignore
        ):
            res.characters_to_ignore.extend(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray.to_proto(
                    resource.characters_to_ignore
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig(
            masking_character=Primitive.from_proto(resource.masking_character),
            number_to_mask=Primitive.from_proto(resource.number_to_mask),
            reverse_order=Primitive.from_proto(resource.reverse_order),
            characters_to_ignore=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray.from_proto(
                resource.characters_to_ignore
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore(
    object
):
    def __init__(
        self, characters_to_skip: str = None, common_characters_to_ignore: str = None
    ):
        self.characters_to_skip = characters_to_skip
        self.common_characters_to_ignore = common_characters_to_ignore

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore()
        )
        if Primitive.to_proto(resource.characters_to_skip):
            res.characters_to_skip = Primitive.to_proto(resource.characters_to_skip)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.to_proto(
            resource.common_characters_to_ignore
        ):
            res.common_characters_to_ignore = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.to_proto(
                resource.common_characters_to_ignore
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore(
            characters_to_skip=Primitive.from_proto(resource.characters_to_skip),
            common_characters_to_ignore=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.from_proto(
                resource.common_characters_to_ignore
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig(
    object
):
    def __init__(
        self,
        crypto_key: dict = None,
        context: dict = None,
        common_alphabet: str = None,
        custom_alphabet: str = None,
        radix: int = None,
        surrogate_info_type: dict = None,
    ):
        self.crypto_key = crypto_key
        self.context = context
        self.common_alphabet = common_alphabet
        self.custom_alphabet = custom_alphabet
        self.radix = radix
        self.surrogate_info_type = surrogate_info_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.to_proto(
            resource.context
        ):
            res.context.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.to_proto(
                    resource.context
                )
            )
        else:
            res.ClearField("context")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.to_proto(
            resource.common_alphabet
        ):
            res.common_alphabet = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.to_proto(
                resource.common_alphabet
            )
        if Primitive.to_proto(resource.custom_alphabet):
            res.custom_alphabet = Primitive.to_proto(resource.custom_alphabet)
        if Primitive.to_proto(resource.radix):
            res.radix = Primitive.to_proto(resource.radix)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.to_proto(
            resource.surrogate_info_type
        ):
            res.surrogate_info_type.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.to_proto(
                    resource.surrogate_info_type
                )
            )
        else:
            res.ClearField("surrogate_info_type")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig(
            crypto_key=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
            context=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.from_proto(
                resource.context
            ),
            common_alphabet=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.from_proto(
                resource.common_alphabet
            ),
            custom_alphabet=Primitive.from_proto(resource.custom_alphabet),
            radix=Primitive.from_proto(resource.radix),
            surrogate_info_type=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.from_proto(
                resource.surrogate_info_type
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContextArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoTypeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig(
    object
):
    def __init__(
        self,
        lower_bound: dict = None,
        upper_bound: dict = None,
        bucket_size: float = None,
    ):
        self.lower_bound = lower_bound
        self.upper_bound = upper_bound
        self.bucket_size = bucket_size

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.to_proto(
            resource.lower_bound
        ):
            res.lower_bound.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.to_proto(
                    resource.lower_bound
                )
            )
        else:
            res.ClearField("lower_bound")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.to_proto(
            resource.upper_bound
        ):
            res.upper_bound.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.to_proto(
                    resource.upper_bound
                )
            )
        else:
            res.ClearField("upper_bound")
        if Primitive.to_proto(resource.bucket_size):
            res.bucket_size = Primitive.to_proto(resource.bucket_size)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig(
            lower_bound=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.from_proto(
                resource.lower_bound
            ),
            upper_bound=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.from_proto(
                resource.upper_bound
            ),
            bucket_size=Primitive.from_proto(resource.bucket_size),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig(
    object
):
    def __init__(self, buckets: list = None):
        self.buckets = buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsArray.to_proto(
            resource.buckets
        ):
            res.buckets.extend(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsArray.to_proto(
                    resource.buckets
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig(
            buckets=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsArray.from_proto(
                resource.buckets
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets(
    object
):
    def __init__(
        self, min: dict = None, max: dict = None, replacement_value: dict = None
    ):
        self.min = min
        self.max = max
        self.replacement_value = replacement_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin.to_proto(
            resource.min
        ):
            res.min.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin.to_proto(
                    resource.min
                )
            )
        else:
            res.ClearField("min")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax.to_proto(
            resource.max
        ):
            res.max.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax.to_proto(
                    resource.max
                )
            )
        else:
            res.ClearField("max")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.to_proto(
            resource.replacement_value
        ):
            res.replacement_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.to_proto(
                    resource.replacement_value
                )
            )
        else:
            res.ClearField("replacement_value")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets(
            min=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin.from_proto(
                resource.min
            ),
            max=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax.from_proto(
                resource.max
            ),
            replacement_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.from_proto(
                resource.replacement_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMin.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMax.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig()
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceWithInfoTypeConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig(
    object
):
    def __init__(self, part_to_extract: str = None):
        self.part_to_extract = part_to_extract

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.to_proto(
            resource.part_to_extract
        ):
            res.part_to_extract = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.to_proto(
                resource.part_to_extract
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig(
            part_to_extract=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.from_proto(
                resource.part_to_extract
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig(
    object
):
    def __init__(self, crypto_key: dict = None):
        self.crypto_key = crypto_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig(
            crypto_key=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig(
    object
):
    def __init__(
        self,
        upper_bound_days: int = None,
        lower_bound_days: int = None,
        context: dict = None,
        crypto_key: dict = None,
    ):
        self.upper_bound_days = upper_bound_days
        self.lower_bound_days = lower_bound_days
        self.context = context
        self.crypto_key = crypto_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig()
        )
        if Primitive.to_proto(resource.upper_bound_days):
            res.upper_bound_days = Primitive.to_proto(resource.upper_bound_days)
        if Primitive.to_proto(resource.lower_bound_days):
            res.lower_bound_days = Primitive.to_proto(resource.lower_bound_days)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext.to_proto(
            resource.context
        ):
            res.context.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext.to_proto(
                    resource.context
                )
            )
        else:
            res.ClearField("context")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig(
            upper_bound_days=Primitive.from_proto(resource.upper_bound_days),
            lower_bound_days=Primitive.from_proto(resource.lower_bound_days),
            context=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext.from_proto(
                resource.context
            ),
            crypto_key=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContextArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigContext.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig(
    object
):
    def __init__(
        self,
        crypto_key: dict = None,
        surrogate_info_type: dict = None,
        context: dict = None,
    ):
        self.crypto_key = crypto_key
        self.surrogate_info_type = surrogate_info_type
        self.context = context

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.to_proto(
            resource.surrogate_info_type
        ):
            res.surrogate_info_type.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.to_proto(
                    resource.surrogate_info_type
                )
            )
        else:
            res.ClearField("surrogate_info_type")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.to_proto(
            resource.context
        ):
            res.context.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.to_proto(
                    resource.context
                )
            )
        else:
            res.ClearField("context")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig(
            crypto_key=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
            surrogate_info_type=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.from_proto(
                resource.surrogate_info_type
            ),
            context=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.from_proto(
                resource.context
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoTypeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContextArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations(
    object
):
    def __init__(self, transformations: list = None):
        self.transformations = transformations

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsArray.to_proto(
            resource.transformations
        ):
            res.transformations.extend(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsArray.to_proto(
                    resource.transformations
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations(
            transformations=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsArray.from_proto(
                resource.transformations
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformations.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations(
    object
):
    def __init__(self, info_types: list = None, primitive_transformation: dict = None):
        self.info_types = info_types
        self.primitive_transformation = primitive_transformation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypesArray.to_proto(
            resource.info_types
        ):
            res.info_types.extend(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypesArray.to_proto(
                    resource.info_types
                )
            )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation.to_proto(
            resource.primitive_transformation
        ):
            res.primitive_transformation.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation.to_proto(
                    resource.primitive_transformation
                )
            )
        else:
            res.ClearField("primitive_transformation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations(
            info_types=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypesArray.from_proto(
                resource.info_types
            ),
            primitive_transformation=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation.from_proto(
                resource.primitive_transformation
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformations.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypes(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypes()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypes(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypes.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsInfoTypes.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation(
    object
):
    def __init__(
        self,
        replace_config: dict = None,
        redact_config: dict = None,
        character_mask_config: dict = None,
        crypto_replace_ffx_fpe_config: dict = None,
        fixed_size_bucketing_config: dict = None,
        bucketing_config: dict = None,
        replace_with_info_type_config: dict = None,
        time_part_config: dict = None,
        crypto_hash_config: dict = None,
        date_shift_config: dict = None,
        crypto_deterministic_config: dict = None,
    ):
        self.replace_config = replace_config
        self.redact_config = redact_config
        self.character_mask_config = character_mask_config
        self.crypto_replace_ffx_fpe_config = crypto_replace_ffx_fpe_config
        self.fixed_size_bucketing_config = fixed_size_bucketing_config
        self.bucketing_config = bucketing_config
        self.replace_with_info_type_config = replace_with_info_type_config
        self.time_part_config = time_part_config
        self.crypto_hash_config = crypto_hash_config
        self.date_shift_config = date_shift_config
        self.crypto_deterministic_config = crypto_deterministic_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig.to_proto(
            resource.replace_config
        ):
            res.replace_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig.to_proto(
                    resource.replace_config
                )
            )
        else:
            res.ClearField("replace_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig.to_proto(
            resource.redact_config
        ):
            res.redact_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig.to_proto(
                    resource.redact_config
                )
            )
        else:
            res.ClearField("redact_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig.to_proto(
            resource.character_mask_config
        ):
            res.character_mask_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig.to_proto(
                    resource.character_mask_config
                )
            )
        else:
            res.ClearField("character_mask_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.to_proto(
            resource.crypto_replace_ffx_fpe_config
        ):
            res.crypto_replace_ffx_fpe_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.to_proto(
                    resource.crypto_replace_ffx_fpe_config
                )
            )
        else:
            res.ClearField("crypto_replace_ffx_fpe_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig.to_proto(
            resource.fixed_size_bucketing_config
        ):
            res.fixed_size_bucketing_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig.to_proto(
                    resource.fixed_size_bucketing_config
                )
            )
        else:
            res.ClearField("fixed_size_bucketing_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig.to_proto(
            resource.bucketing_config
        ):
            res.bucketing_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig.to_proto(
                    resource.bucketing_config
                )
            )
        else:
            res.ClearField("bucketing_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.to_proto(
            resource.replace_with_info_type_config
        ):
            res.replace_with_info_type_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.to_proto(
                    resource.replace_with_info_type_config
                )
            )
        else:
            res.ClearField("replace_with_info_type_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig.to_proto(
            resource.time_part_config
        ):
            res.time_part_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig.to_proto(
                    resource.time_part_config
                )
            )
        else:
            res.ClearField("time_part_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig.to_proto(
            resource.crypto_hash_config
        ):
            res.crypto_hash_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig.to_proto(
                    resource.crypto_hash_config
                )
            )
        else:
            res.ClearField("crypto_hash_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig.to_proto(
            resource.date_shift_config
        ):
            res.date_shift_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig.to_proto(
                    resource.date_shift_config
                )
            )
        else:
            res.ClearField("date_shift_config")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig.to_proto(
            resource.crypto_deterministic_config
        ):
            res.crypto_deterministic_config.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig.to_proto(
                    resource.crypto_deterministic_config
                )
            )
        else:
            res.ClearField("crypto_deterministic_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation(
            replace_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig.from_proto(
                resource.replace_config
            ),
            redact_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig.from_proto(
                resource.redact_config
            ),
            character_mask_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig.from_proto(
                resource.character_mask_config
            ),
            crypto_replace_ffx_fpe_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.from_proto(
                resource.crypto_replace_ffx_fpe_config
            ),
            fixed_size_bucketing_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig.from_proto(
                resource.fixed_size_bucketing_config
            ),
            bucketing_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig.from_proto(
                resource.bucketing_config
            ),
            replace_with_info_type_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.from_proto(
                resource.replace_with_info_type_config
            ),
            time_part_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig.from_proto(
                resource.time_part_config
            ),
            crypto_hash_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig.from_proto(
                resource.crypto_hash_config
            ),
            date_shift_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig.from_proto(
                resource.date_shift_config
            ),
            crypto_deterministic_config=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig.from_proto(
                resource.crypto_deterministic_config
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig(
    object
):
    def __init__(self, new_value: dict = None):
        self.new_value = new_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue.to_proto(
            resource.new_value
        ):
            res.new_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue.to_proto(
                    resource.new_value
                )
            )
        else:
            res.ClearField("new_value")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig(
            new_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue.from_proto(
                resource.new_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig()
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig(
    object
):
    def __init__(
        self,
        masking_character: str = None,
        number_to_mask: int = None,
        reverse_order: bool = None,
        characters_to_ignore: list = None,
    ):
        self.masking_character = masking_character
        self.number_to_mask = number_to_mask
        self.reverse_order = reverse_order
        self.characters_to_ignore = characters_to_ignore

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig()
        )
        if Primitive.to_proto(resource.masking_character):
            res.masking_character = Primitive.to_proto(resource.masking_character)
        if Primitive.to_proto(resource.number_to_mask):
            res.number_to_mask = Primitive.to_proto(resource.number_to_mask)
        if Primitive.to_proto(resource.reverse_order):
            res.reverse_order = Primitive.to_proto(resource.reverse_order)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray.to_proto(
            resource.characters_to_ignore
        ):
            res.characters_to_ignore.extend(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray.to_proto(
                    resource.characters_to_ignore
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig(
            masking_character=Primitive.from_proto(resource.masking_character),
            number_to_mask=Primitive.from_proto(resource.number_to_mask),
            reverse_order=Primitive.from_proto(resource.reverse_order),
            characters_to_ignore=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray.from_proto(
                resource.characters_to_ignore
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore(
    object
):
    def __init__(
        self, characters_to_skip: str = None, common_characters_to_ignore: str = None
    ):
        self.characters_to_skip = characters_to_skip
        self.common_characters_to_ignore = common_characters_to_ignore

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore()
        )
        if Primitive.to_proto(resource.characters_to_skip):
            res.characters_to_skip = Primitive.to_proto(resource.characters_to_skip)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.to_proto(
            resource.common_characters_to_ignore
        ):
            res.common_characters_to_ignore = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.to_proto(
                resource.common_characters_to_ignore
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore(
            characters_to_skip=Primitive.from_proto(resource.characters_to_skip),
            common_characters_to_ignore=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.from_proto(
                resource.common_characters_to_ignore
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig(
    object
):
    def __init__(
        self,
        crypto_key: dict = None,
        context: dict = None,
        common_alphabet: str = None,
        custom_alphabet: str = None,
        radix: int = None,
        surrogate_info_type: dict = None,
    ):
        self.crypto_key = crypto_key
        self.context = context
        self.common_alphabet = common_alphabet
        self.custom_alphabet = custom_alphabet
        self.radix = radix
        self.surrogate_info_type = surrogate_info_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.to_proto(
            resource.context
        ):
            res.context.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.to_proto(
                    resource.context
                )
            )
        else:
            res.ClearField("context")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.to_proto(
            resource.common_alphabet
        ):
            res.common_alphabet = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.to_proto(
                resource.common_alphabet
            )
        if Primitive.to_proto(resource.custom_alphabet):
            res.custom_alphabet = Primitive.to_proto(resource.custom_alphabet)
        if Primitive.to_proto(resource.radix):
            res.radix = Primitive.to_proto(resource.radix)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.to_proto(
            resource.surrogate_info_type
        ):
            res.surrogate_info_type.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.to_proto(
                    resource.surrogate_info_type
                )
            )
        else:
            res.ClearField("surrogate_info_type")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig(
            crypto_key=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
            context=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.from_proto(
                resource.context
            ),
            common_alphabet=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.from_proto(
                resource.common_alphabet
            ),
            custom_alphabet=Primitive.from_proto(resource.custom_alphabet),
            radix=Primitive.from_proto(resource.radix),
            surrogate_info_type=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.from_proto(
                resource.surrogate_info_type
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContextArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoTypeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig(
    object
):
    def __init__(
        self,
        lower_bound: dict = None,
        upper_bound: dict = None,
        bucket_size: float = None,
    ):
        self.lower_bound = lower_bound
        self.upper_bound = upper_bound
        self.bucket_size = bucket_size

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.to_proto(
            resource.lower_bound
        ):
            res.lower_bound.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.to_proto(
                    resource.lower_bound
                )
            )
        else:
            res.ClearField("lower_bound")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.to_proto(
            resource.upper_bound
        ):
            res.upper_bound.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.to_proto(
                    resource.upper_bound
                )
            )
        else:
            res.ClearField("upper_bound")
        if Primitive.to_proto(resource.bucket_size):
            res.bucket_size = Primitive.to_proto(resource.bucket_size)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig(
            lower_bound=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.from_proto(
                resource.lower_bound
            ),
            upper_bound=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.from_proto(
                resource.upper_bound
            ),
            bucket_size=Primitive.from_proto(resource.bucket_size),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBound.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig(
    object
):
    def __init__(self, buckets: list = None):
        self.buckets = buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsArray.to_proto(
            resource.buckets
        ):
            res.buckets.extend(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsArray.to_proto(
                    resource.buckets
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig(
            buckets=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsArray.from_proto(
                resource.buckets
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets(
    object
):
    def __init__(
        self, min: dict = None, max: dict = None, replacement_value: dict = None
    ):
        self.min = min
        self.max = max
        self.replacement_value = replacement_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin.to_proto(
            resource.min
        ):
            res.min.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin.to_proto(
                    resource.min
                )
            )
        else:
            res.ClearField("min")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax.to_proto(
            resource.max
        ):
            res.max.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax.to_proto(
                    resource.max
                )
            )
        else:
            res.ClearField("max")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.to_proto(
            resource.replacement_value
        ):
            res.replacement_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.to_proto(
                    resource.replacement_value
                )
            )
        else:
            res.ClearField("replacement_value")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets(
            min=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin.from_proto(
                resource.min
            ),
            max=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax.from_proto(
                resource.max
            ),
            replacement_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.from_proto(
                resource.replacement_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig()
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig(
    object
):
    def __init__(self, part_to_extract: str = None):
        self.part_to_extract = part_to_extract

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.to_proto(
            resource.part_to_extract
        ):
            res.part_to_extract = DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.to_proto(
                resource.part_to_extract
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig(
            part_to_extract=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.from_proto(
                resource.part_to_extract
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig(
    object
):
    def __init__(self, crypto_key: dict = None):
        self.crypto_key = crypto_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig(
            crypto_key=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig(
    object
):
    def __init__(
        self,
        upper_bound_days: int = None,
        lower_bound_days: int = None,
        context: dict = None,
        crypto_key: dict = None,
    ):
        self.upper_bound_days = upper_bound_days
        self.lower_bound_days = lower_bound_days
        self.context = context
        self.crypto_key = crypto_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig()
        )
        if Primitive.to_proto(resource.upper_bound_days):
            res.upper_bound_days = Primitive.to_proto(resource.upper_bound_days)
        if Primitive.to_proto(resource.lower_bound_days):
            res.lower_bound_days = Primitive.to_proto(resource.lower_bound_days)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext.to_proto(
            resource.context
        ):
            res.context.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext.to_proto(
                    resource.context
                )
            )
        else:
            res.ClearField("context")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig(
            upper_bound_days=Primitive.from_proto(resource.upper_bound_days),
            lower_bound_days=Primitive.from_proto(resource.lower_bound_days),
            context=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext.from_proto(
                resource.context
            ),
            crypto_key=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContextArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig(
    object
):
    def __init__(
        self,
        crypto_key: dict = None,
        surrogate_info_type: dict = None,
        context: dict = None,
    ):
        self.crypto_key = crypto_key
        self.surrogate_info_type = surrogate_info_type
        self.context = context

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.to_proto(
            resource.crypto_key
        ):
            res.crypto_key.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.to_proto(
                    resource.crypto_key
                )
            )
        else:
            res.ClearField("crypto_key")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.to_proto(
            resource.surrogate_info_type
        ):
            res.surrogate_info_type.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.to_proto(
                    resource.surrogate_info_type
                )
            )
        else:
            res.ClearField("surrogate_info_type")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.to_proto(
            resource.context
        ):
            res.context.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.to_proto(
                    resource.context
                )
            )
        else:
            res.ClearField("context")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig(
            crypto_key=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.from_proto(
                resource.crypto_key
            ),
            surrogate_info_type=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.from_proto(
                resource.surrogate_info_type
            ),
            context=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.from_proto(
                resource.context
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey(
    object
):
    def __init__(
        self, transient: dict = None, unwrapped: dict = None, kms_wrapped: dict = None
    ):
        self.transient = transient
        self.unwrapped = unwrapped
        self.kms_wrapped = kms_wrapped

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.to_proto(
            resource.transient
        ):
            res.transient.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.to_proto(
                    resource.transient
                )
            )
        else:
            res.ClearField("transient")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.to_proto(
            resource.unwrapped
        ):
            res.unwrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.to_proto(
                    resource.unwrapped
                )
            )
        else:
            res.ClearField("unwrapped")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.to_proto(
            resource.kms_wrapped
        ):
            res.kms_wrapped.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.to_proto(
                    resource.kms_wrapped
                )
            )
        else:
            res.ClearField("kms_wrapped")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey(
            transient=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.from_proto(
                resource.transient
            ),
            unwrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.from_proto(
                resource.unwrapped
            ),
            kms_wrapped=DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.from_proto(
                resource.kms_wrapped
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransientArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyTransient.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped(
    object
):
    def __init__(self, key: str = None):
        self.key = key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped(
            key=Primitive.from_proto(resource.key),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped(
    object
):
    def __init__(self, wrapped_key: str = None, crypto_key_name: str = None):
        self.wrapped_key = wrapped_key
        self.crypto_key_name = crypto_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped()
        )
        if Primitive.to_proto(resource.wrapped_key):
            res.wrapped_key = Primitive.to_proto(resource.wrapped_key)
        if Primitive.to_proto(resource.crypto_key_name):
            res.crypto_key_name = Primitive.to_proto(resource.crypto_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped(
            wrapped_key=Primitive.from_proto(resource.wrapped_key),
            crypto_key_name=Primitive.from_proto(resource.crypto_key_name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrappedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyKmsWrapped.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoTypeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContextArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressions(object):
    def __init__(self, condition: dict = None):
        self.condition = condition

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressions()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsCondition.to_proto(
            resource.condition
        ):
            res.condition.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsCondition.to_proto(
                    resource.condition
                )
            )
        else:
            res.ClearField("condition")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressions(
            condition=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsCondition.from_proto(
                resource.condition
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressions.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsCondition(
    object
):
    def __init__(self, expressions: dict = None):
        self.expressions = expressions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsCondition()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions.to_proto(
            resource.expressions
        ):
            res.expressions.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions.to_proto(
                    resource.expressions
                )
            )
        else:
            res.ClearField("expressions")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsCondition(
            expressions=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions.from_proto(
                resource.expressions
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsCondition.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsCondition.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions(
    object
):
    def __init__(self, logical_operator: str = None, conditions: dict = None):
        self.logical_operator = logical_operator
        self.conditions = conditions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsLogicalOperatorEnum.to_proto(
            resource.logical_operator
        ):
            res.logical_operator = DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsLogicalOperatorEnum.to_proto(
                resource.logical_operator
            )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions.to_proto(
            resource.conditions
        ):
            res.conditions.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions.to_proto(
                    resource.conditions
                )
            )
        else:
            res.ClearField("conditions")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions(
            logical_operator=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsLogicalOperatorEnum.from_proto(
                resource.logical_operator
            ),
            conditions=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions.from_proto(
                resource.conditions
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressions.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions(
    object
):
    def __init__(self, conditions: list = None):
        self.conditions = conditions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsArray.to_proto(
            resource.conditions
        ):
            res.conditions.extend(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsArray.to_proto(
                    resource.conditions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions(
            conditions=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsArray.from_proto(
                resource.conditions
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditions.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions(
    object
):
    def __init__(self, field: dict = None, operator: str = None, value: dict = None):
        self.field = field
        self.operator = operator
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions()
        )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField.to_proto(
            resource.field
        ):
            res.field.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField.to_proto(
                    resource.field
                )
            )
        else:
            res.ClearField("field")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsOperatorEnum.to_proto(
            resource.operator
        ):
            res.operator = DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsOperatorEnum.to_proto(
                resource.operator
            )
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue.to_proto(
            resource.value
        ):
            res.value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue.to_proto(
                    resource.value
                )
            )
        else:
            res.ClearField("value")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions(
            field=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField.from_proto(
                resource.field
            ),
            operator=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsOperatorEnum.from_proto(
                resource.operator
            ),
            value=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue.from_proto(
                resource.value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField(
            name=Primitive.from_proto(resource.name),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsFieldArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue(
    object
):
    def __init__(
        self,
        integer_value: int = None,
        float_value: float = None,
        string_value: str = None,
        boolean_value: bool = None,
        timestamp_value: str = None,
        time_value: dict = None,
        date_value: dict = None,
        day_of_week_value: str = None,
    ):
        self.integer_value = integer_value
        self.float_value = float_value
        self.string_value = string_value
        self.boolean_value = boolean_value
        self.timestamp_value = timestamp_value
        self.time_value = time_value
        self.date_value = date_value
        self.day_of_week_value = day_of_week_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue()
        )
        if Primitive.to_proto(resource.integer_value):
            res.integer_value = Primitive.to_proto(resource.integer_value)
        if Primitive.to_proto(resource.float_value):
            res.float_value = Primitive.to_proto(resource.float_value)
        if Primitive.to_proto(resource.string_value):
            res.string_value = Primitive.to_proto(resource.string_value)
        if Primitive.to_proto(resource.boolean_value):
            res.boolean_value = Primitive.to_proto(resource.boolean_value)
        if Primitive.to_proto(resource.timestamp_value):
            res.timestamp_value = Primitive.to_proto(resource.timestamp_value)
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue.to_proto(
            resource.time_value
        ):
            res.time_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue.to_proto(
                    resource.time_value
                )
            )
        else:
            res.ClearField("time_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue.to_proto(
            resource.date_value
        ):
            res.date_value.CopyFrom(
                DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue.to_proto(
                    resource.date_value
                )
            )
        else:
            res.ClearField("date_value")
        if DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum.to_proto(
            resource.day_of_week_value
        ):
            res.day_of_week_value = DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum.to_proto(
                resource.day_of_week_value
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue(
            integer_value=Primitive.from_proto(resource.integer_value),
            float_value=Primitive.from_proto(resource.float_value),
            string_value=Primitive.from_proto(resource.string_value),
            boolean_value=Primitive.from_proto(resource.boolean_value),
            timestamp_value=Primitive.from_proto(resource.timestamp_value),
            time_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue.from_proto(
                resource.time_value
            ),
            date_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue.from_proto(
                resource.date_value
            ),
            day_of_week_value=DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum.from_proto(
                resource.day_of_week_value
            ),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue(
    object
):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueTimeValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue(
    object
):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue()
        )
        if Primitive.to_proto(resource.year):
            res.year = Primitive.to_proto(resource.year)
        if Primitive.to_proto(resource.month):
            res.month = Primitive.to_proto(resource.month)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValueArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDateValue.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigTransformationErrorHandling(object):
    def __init__(self, throw_error: dict = None, leave_untransformed: dict = None):
        self.throw_error = throw_error
        self.leave_untransformed = leave_untransformed

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigTransformationErrorHandling()
        )
        if DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError.to_proto(
            resource.throw_error
        ):
            res.throw_error.CopyFrom(
                DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError.to_proto(
                    resource.throw_error
                )
            )
        else:
            res.ClearField("throw_error")
        if DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed.to_proto(
            resource.leave_untransformed
        ):
            res.leave_untransformed.CopyFrom(
                DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed.to_proto(
                    resource.leave_untransformed
                )
            )
        else:
            res.ClearField("leave_untransformed")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigTransformationErrorHandling(
            throw_error=DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError.from_proto(
                resource.throw_error
            ),
            leave_untransformed=DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed.from_proto(
                resource.leave_untransformed
            ),
        )


class DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigTransformationErrorHandling.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigTransformationErrorHandling.from_proto(i)
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError()


class DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowErrorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed()
        )


class DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformedArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed.from_proto(
                i
            )
            for i in resources
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsLogicalOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsLogicalOperatorEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsLogicalOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsLogicalOperatorEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsLogicalOperatorEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOperatorEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOperatorEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOperatorEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnoreCommonCharactersToIgnoreEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCommonAlphabetEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigLowerBoundDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMinDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMaxDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValueDayOfWeekValueEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfigPartToExtractEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsLogicalOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsLogicalOperatorEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsLogicalOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsLogicalOperatorEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsLogicalOperatorEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsOperatorEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsOperatorEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsOperatorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsOperatorEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsOperatorEnum"
            ) :
        ]


class DeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum.Value(
            "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return deidentify_template_pb2.DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum.Name(
            resource
        )[
            len(
                "DlpBetaDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValueDayOfWeekValueEnum"
            ) :
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
