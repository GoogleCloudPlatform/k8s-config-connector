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
from google3.cloud.graphite.mmv2.services.google.dlp import inspect_template_pb2
from google3.cloud.graphite.mmv2.services.google.dlp import inspect_template_pb2_grpc

from typing import List


class InspectTemplate(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        inspect_config: dict = None,
        location_id: str = None,
        parent: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.description = description
        self.inspect_config = inspect_config
        self.parent = parent
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = inspect_template_pb2_grpc.DlpAlphaInspectTemplateServiceStub(
            channel.Channel()
        )
        request = inspect_template_pb2.ApplyDlpAlphaInspectTemplateRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if InspectTemplateInspectConfig.to_proto(self.inspect_config):
            request.resource.inspect_config.CopyFrom(
                InspectTemplateInspectConfig.to_proto(self.inspect_config)
            )
        else:
            request.resource.ClearField("inspect_config")
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDlpAlphaInspectTemplate(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.inspect_config = InspectTemplateInspectConfig.from_proto(
            response.inspect_config
        )
        self.location_id = Primitive.from_proto(response.location_id)
        self.parent = Primitive.from_proto(response.parent)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = inspect_template_pb2_grpc.DlpAlphaInspectTemplateServiceStub(
            channel.Channel()
        )
        request = inspect_template_pb2.DeleteDlpAlphaInspectTemplateRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if InspectTemplateInspectConfig.to_proto(self.inspect_config):
            request.resource.inspect_config.CopyFrom(
                InspectTemplateInspectConfig.to_proto(self.inspect_config)
            )
        else:
            request.resource.ClearField("inspect_config")
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteDlpAlphaInspectTemplate(request)

    @classmethod
    def list(self, location, parent, service_account_file=""):
        stub = inspect_template_pb2_grpc.DlpAlphaInspectTemplateServiceStub(
            channel.Channel()
        )
        request = inspect_template_pb2.ListDlpAlphaInspectTemplateRequest()
        request.service_account_file = service_account_file
        request.Location = location

        request.Parent = parent

        return stub.ListDlpAlphaInspectTemplate(request).items

    def to_proto(self):
        resource = inspect_template_pb2.DlpAlphaInspectTemplate()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if InspectTemplateInspectConfig.to_proto(self.inspect_config):
            resource.inspect_config.CopyFrom(
                InspectTemplateInspectConfig.to_proto(self.inspect_config)
            )
        else:
            resource.ClearField("inspect_config")
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class InspectTemplateInspectConfig(object):
    def __init__(
        self,
        info_types: list = None,
        min_likelihood: str = None,
        limits: dict = None,
        include_quote: bool = None,
        exclude_info_types: bool = None,
        custom_info_types: list = None,
        content_options: list = None,
        rule_set: list = None,
    ):
        self.info_types = info_types
        self.min_likelihood = min_likelihood
        self.limits = limits
        self.include_quote = include_quote
        self.exclude_info_types = exclude_info_types
        self.custom_info_types = custom_info_types
        self.content_options = content_options
        self.rule_set = rule_set

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = inspect_template_pb2.DlpAlphaInspectTemplateInspectConfig()
        if InspectTemplateInspectConfigInfoTypesArray.to_proto(resource.info_types):
            res.info_types.extend(
                InspectTemplateInspectConfigInfoTypesArray.to_proto(resource.info_types)
            )
        if InspectTemplateInspectConfigMinLikelihoodEnum.to_proto(
            resource.min_likelihood
        ):
            res.min_likelihood = InspectTemplateInspectConfigMinLikelihoodEnum.to_proto(
                resource.min_likelihood
            )
        if InspectTemplateInspectConfigLimits.to_proto(resource.limits):
            res.limits.CopyFrom(
                InspectTemplateInspectConfigLimits.to_proto(resource.limits)
            )
        else:
            res.ClearField("limits")
        if Primitive.to_proto(resource.include_quote):
            res.include_quote = Primitive.to_proto(resource.include_quote)
        if Primitive.to_proto(resource.exclude_info_types):
            res.exclude_info_types = Primitive.to_proto(resource.exclude_info_types)
        if InspectTemplateInspectConfigCustomInfoTypesArray.to_proto(
            resource.custom_info_types
        ):
            res.custom_info_types.extend(
                InspectTemplateInspectConfigCustomInfoTypesArray.to_proto(
                    resource.custom_info_types
                )
            )
        if InspectTemplateInspectConfigContentOptionsEnumArray.to_proto(
            resource.content_options
        ):
            res.content_options.extend(
                InspectTemplateInspectConfigContentOptionsEnumArray.to_proto(
                    resource.content_options
                )
            )
        if InspectTemplateInspectConfigRuleSetArray.to_proto(resource.rule_set):
            res.rule_set.extend(
                InspectTemplateInspectConfigRuleSetArray.to_proto(resource.rule_set)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfig(
            info_types=InspectTemplateInspectConfigInfoTypesArray.from_proto(
                resource.info_types
            ),
            min_likelihood=InspectTemplateInspectConfigMinLikelihoodEnum.from_proto(
                resource.min_likelihood
            ),
            limits=InspectTemplateInspectConfigLimits.from_proto(resource.limits),
            include_quote=Primitive.from_proto(resource.include_quote),
            exclude_info_types=Primitive.from_proto(resource.exclude_info_types),
            custom_info_types=InspectTemplateInspectConfigCustomInfoTypesArray.from_proto(
                resource.custom_info_types
            ),
            content_options=InspectTemplateInspectConfigContentOptionsEnumArray.from_proto(
                resource.content_options
            ),
            rule_set=InspectTemplateInspectConfigRuleSetArray.from_proto(
                resource.rule_set
            ),
        )


class InspectTemplateInspectConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InspectTemplateInspectConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InspectTemplateInspectConfig.from_proto(i) for i in resources]


class InspectTemplateInspectConfigInfoTypes(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigInfoTypes()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigInfoTypes(
            name=Primitive.from_proto(resource.name),
        )


class InspectTemplateInspectConfigInfoTypesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InspectTemplateInspectConfigInfoTypes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InspectTemplateInspectConfigInfoTypes.from_proto(i) for i in resources]


class InspectTemplateInspectConfigLimits(object):
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

        res = inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigLimits()
        if Primitive.to_proto(resource.max_findings_per_item):
            res.max_findings_per_item = Primitive.to_proto(
                resource.max_findings_per_item
            )
        if Primitive.to_proto(resource.max_findings_per_request):
            res.max_findings_per_request = Primitive.to_proto(
                resource.max_findings_per_request
            )
        if InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeArray.to_proto(
            resource.max_findings_per_info_type
        ):
            res.max_findings_per_info_type.extend(
                InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeArray.to_proto(
                    resource.max_findings_per_info_type
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigLimits(
            max_findings_per_item=Primitive.from_proto(resource.max_findings_per_item),
            max_findings_per_request=Primitive.from_proto(
                resource.max_findings_per_request
            ),
            max_findings_per_info_type=InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeArray.from_proto(
                resource.max_findings_per_info_type
            ),
        )


class InspectTemplateInspectConfigLimitsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InspectTemplateInspectConfigLimits.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InspectTemplateInspectConfigLimits.from_proto(i) for i in resources]


class InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(object):
    def __init__(self, info_type: dict = None, max_findings: int = None):
        self.info_type = info_type
        self.max_findings = max_findings

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType()
        )
        if InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType.to_proto(
            resource.info_type
        ):
            res.info_type.CopyFrom(
                InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType.to_proto(
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

        return InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(
            info_type=InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType.from_proto(
                resource.info_type
            ),
            max_findings=Primitive.from_proto(resource.max_findings),
        )


class InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(
            name=Primitive.from_proto(resource.name),
        )


class InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType.from_proto(
                i
            )
            for i in resources
        ]


class InspectTemplateInspectConfigCustomInfoTypes(object):
    def __init__(
        self,
        info_type: dict = None,
        likelihood: str = None,
        dictionary: dict = None,
        regex: dict = None,
        surrogate_type: dict = None,
        stored_type: dict = None,
        exclusion_type: str = None,
    ):
        self.info_type = info_type
        self.likelihood = likelihood
        self.dictionary = dictionary
        self.regex = regex
        self.surrogate_type = surrogate_type
        self.stored_type = stored_type
        self.exclusion_type = exclusion_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypes()
        if InspectTemplateInspectConfigCustomInfoTypesInfoType.to_proto(
            resource.info_type
        ):
            res.info_type.CopyFrom(
                InspectTemplateInspectConfigCustomInfoTypesInfoType.to_proto(
                    resource.info_type
                )
            )
        else:
            res.ClearField("info_type")
        if InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum.to_proto(
            resource.likelihood
        ):
            res.likelihood = (
                InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum.to_proto(
                    resource.likelihood
                )
            )
        if InspectTemplateInspectConfigCustomInfoTypesDictionary.to_proto(
            resource.dictionary
        ):
            res.dictionary.CopyFrom(
                InspectTemplateInspectConfigCustomInfoTypesDictionary.to_proto(
                    resource.dictionary
                )
            )
        else:
            res.ClearField("dictionary")
        if InspectTemplateInspectConfigCustomInfoTypesRegex.to_proto(resource.regex):
            res.regex.CopyFrom(
                InspectTemplateInspectConfigCustomInfoTypesRegex.to_proto(
                    resource.regex
                )
            )
        else:
            res.ClearField("regex")
        if InspectTemplateInspectConfigCustomInfoTypesSurrogateType.to_proto(
            resource.surrogate_type
        ):
            res.surrogate_type.CopyFrom(
                InspectTemplateInspectConfigCustomInfoTypesSurrogateType.to_proto(
                    resource.surrogate_type
                )
            )
        else:
            res.ClearField("surrogate_type")
        if InspectTemplateInspectConfigCustomInfoTypesStoredType.to_proto(
            resource.stored_type
        ):
            res.stored_type.CopyFrom(
                InspectTemplateInspectConfigCustomInfoTypesStoredType.to_proto(
                    resource.stored_type
                )
            )
        else:
            res.ClearField("stored_type")
        if InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum.to_proto(
            resource.exclusion_type
        ):
            res.exclusion_type = (
                InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum.to_proto(
                    resource.exclusion_type
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigCustomInfoTypes(
            info_type=InspectTemplateInspectConfigCustomInfoTypesInfoType.from_proto(
                resource.info_type
            ),
            likelihood=InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum.from_proto(
                resource.likelihood
            ),
            dictionary=InspectTemplateInspectConfigCustomInfoTypesDictionary.from_proto(
                resource.dictionary
            ),
            regex=InspectTemplateInspectConfigCustomInfoTypesRegex.from_proto(
                resource.regex
            ),
            surrogate_type=InspectTemplateInspectConfigCustomInfoTypesSurrogateType.from_proto(
                resource.surrogate_type
            ),
            stored_type=InspectTemplateInspectConfigCustomInfoTypesStoredType.from_proto(
                resource.stored_type
            ),
            exclusion_type=InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum.from_proto(
                resource.exclusion_type
            ),
        )


class InspectTemplateInspectConfigCustomInfoTypesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigCustomInfoTypes.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigCustomInfoTypes.from_proto(i) for i in resources
        ]


class InspectTemplateInspectConfigCustomInfoTypesInfoType(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesInfoType()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigCustomInfoTypesInfoType(
            name=Primitive.from_proto(resource.name),
        )


class InspectTemplateInspectConfigCustomInfoTypesInfoTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigCustomInfoTypesInfoType.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigCustomInfoTypesInfoType.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigCustomInfoTypesDictionary(object):
    def __init__(self, word_list: dict = None, cloud_storage_path: dict = None):
        self.word_list = word_list
        self.cloud_storage_path = cloud_storage_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionary()
        )
        if InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList.to_proto(
            resource.word_list
        ):
            res.word_list.CopyFrom(
                InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList.to_proto(
                    resource.word_list
                )
            )
        else:
            res.ClearField("word_list")
        if InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath.to_proto(
            resource.cloud_storage_path
        ):
            res.cloud_storage_path.CopyFrom(
                InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath.to_proto(
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

        return InspectTemplateInspectConfigCustomInfoTypesDictionary(
            word_list=InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList.from_proto(
                resource.word_list
            ),
            cloud_storage_path=InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath.from_proto(
                resource.cloud_storage_path
            ),
        )


class InspectTemplateInspectConfigCustomInfoTypesDictionaryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigCustomInfoTypesDictionary.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigCustomInfoTypesDictionary.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(object):
    def __init__(self, words: list = None):
        self.words = words

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList()
        )
        if Primitive.to_proto(resource.words):
            res.words.extend(Primitive.to_proto(resource.words))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(
            words=Primitive.from_proto(resource.words),
        )


class InspectTemplateInspectConfigCustomInfoTypesDictionaryWordListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(object):
    def __init__(self, path: str = None):
        self.path = path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath()
        )
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(
            path=Primitive.from_proto(resource.path),
        )


class InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath.from_proto(
                i
            )
            for i in resources
        ]


class InspectTemplateInspectConfigCustomInfoTypesRegex(object):
    def __init__(self, pattern: str = None, group_indexes: list = None):
        self.pattern = pattern
        self.group_indexes = group_indexes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesRegex()
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

        return InspectTemplateInspectConfigCustomInfoTypesRegex(
            pattern=Primitive.from_proto(resource.pattern),
            group_indexes=int64Array.from_proto(resource.group_indexes),
        )


class InspectTemplateInspectConfigCustomInfoTypesRegexArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigCustomInfoTypesRegex.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigCustomInfoTypesRegex.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigCustomInfoTypesSurrogateType(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesSurrogateType()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigCustomInfoTypesSurrogateType()


class InspectTemplateInspectConfigCustomInfoTypesSurrogateTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigCustomInfoTypesSurrogateType.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigCustomInfoTypesSurrogateType.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigCustomInfoTypesStoredType(object):
    def __init__(self, name: str = None, create_time: str = None):
        self.name = name
        self.create_time = create_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesStoredType()
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

        return InspectTemplateInspectConfigCustomInfoTypesStoredType(
            name=Primitive.from_proto(resource.name),
            create_time=Primitive.from_proto(resource.create_time),
        )


class InspectTemplateInspectConfigCustomInfoTypesStoredTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigCustomInfoTypesStoredType.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigCustomInfoTypesStoredType.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSet(object):
    def __init__(self, info_types: list = None, rules: list = None):
        self.info_types = info_types
        self.rules = rules

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSet()
        if InspectTemplateInspectConfigRuleSetInfoTypesArray.to_proto(
            resource.info_types
        ):
            res.info_types.extend(
                InspectTemplateInspectConfigRuleSetInfoTypesArray.to_proto(
                    resource.info_types
                )
            )
        if InspectTemplateInspectConfigRuleSetRulesArray.to_proto(resource.rules):
            res.rules.extend(
                InspectTemplateInspectConfigRuleSetRulesArray.to_proto(resource.rules)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigRuleSet(
            info_types=InspectTemplateInspectConfigRuleSetInfoTypesArray.from_proto(
                resource.info_types
            ),
            rules=InspectTemplateInspectConfigRuleSetRulesArray.from_proto(
                resource.rules
            ),
        )


class InspectTemplateInspectConfigRuleSetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InspectTemplateInspectConfigRuleSet.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InspectTemplateInspectConfigRuleSet.from_proto(i) for i in resources]


class InspectTemplateInspectConfigRuleSetInfoTypes(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetInfoTypes()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigRuleSetInfoTypes(
            name=Primitive.from_proto(resource.name),
        )


class InspectTemplateInspectConfigRuleSetInfoTypesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetInfoTypes.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetInfoTypes.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRules(object):
    def __init__(self, hotword_rule: dict = None, exclusion_rule: dict = None):
        self.hotword_rule = hotword_rule
        self.exclusion_rule = exclusion_rule

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRules()
        if InspectTemplateInspectConfigRuleSetRulesHotwordRule.to_proto(
            resource.hotword_rule
        ):
            res.hotword_rule.CopyFrom(
                InspectTemplateInspectConfigRuleSetRulesHotwordRule.to_proto(
                    resource.hotword_rule
                )
            )
        else:
            res.ClearField("hotword_rule")
        if InspectTemplateInspectConfigRuleSetRulesExclusionRule.to_proto(
            resource.exclusion_rule
        ):
            res.exclusion_rule.CopyFrom(
                InspectTemplateInspectConfigRuleSetRulesExclusionRule.to_proto(
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

        return InspectTemplateInspectConfigRuleSetRules(
            hotword_rule=InspectTemplateInspectConfigRuleSetRulesHotwordRule.from_proto(
                resource.hotword_rule
            ),
            exclusion_rule=InspectTemplateInspectConfigRuleSetRulesExclusionRule.from_proto(
                resource.exclusion_rule
            ),
        )


class InspectTemplateInspectConfigRuleSetRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InspectTemplateInspectConfigRuleSetRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRules.from_proto(i) for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRulesHotwordRule(object):
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
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRule()
        )
        if InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex.to_proto(
            resource.hotword_regex
        ):
            res.hotword_regex.CopyFrom(
                InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex.to_proto(
                    resource.hotword_regex
                )
            )
        else:
            res.ClearField("hotword_regex")
        if InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity.to_proto(
            resource.proximity
        ):
            res.proximity.CopyFrom(
                InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity.to_proto(
                    resource.proximity
                )
            )
        else:
            res.ClearField("proximity")
        if InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment.to_proto(
            resource.likelihood_adjustment
        ):
            res.likelihood_adjustment.CopyFrom(
                InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment.to_proto(
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

        return InspectTemplateInspectConfigRuleSetRulesHotwordRule(
            hotword_regex=InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex.from_proto(
                resource.hotword_regex
            ),
            proximity=InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity.from_proto(
                resource.proximity
            ),
            likelihood_adjustment=InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment.from_proto(
                resource.likelihood_adjustment
            ),
        )


class InspectTemplateInspectConfigRuleSetRulesHotwordRuleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetRulesHotwordRule.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRulesHotwordRule.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(object):
    def __init__(self, pattern: str = None, group_indexes: list = None):
        self.pattern = pattern
        self.group_indexes = group_indexes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex()
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

        return InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(
            pattern=Primitive.from_proto(resource.pattern),
            group_indexes=int64Array.from_proto(resource.group_indexes),
        )


class InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex.from_proto(
                i
            )
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(object):
    def __init__(self, window_before: int = None, window_after: int = None):
        self.window_before = window_before
        self.window_after = window_after

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity()
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

        return InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(
            window_before=Primitive.from_proto(resource.window_before),
            window_after=Primitive.from_proto(resource.window_after),
        )


class InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(object):
    def __init__(self, fixed_likelihood: str = None, relative_likelihood: int = None):
        self.fixed_likelihood = fixed_likelihood
        self.relative_likelihood = relative_likelihood

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment()
        )
        if InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.to_proto(
            resource.fixed_likelihood
        ):
            res.fixed_likelihood = InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.to_proto(
                resource.fixed_likelihood
            )
        if Primitive.to_proto(resource.relative_likelihood):
            res.relative_likelihood = Primitive.to_proto(resource.relative_likelihood)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(
            fixed_likelihood=InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.from_proto(
                resource.fixed_likelihood
            ),
            relative_likelihood=Primitive.from_proto(resource.relative_likelihood),
        )


class InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment.from_proto(
                i
            )
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRulesExclusionRule(object):
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
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRule()
        )
        if InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary.to_proto(
            resource.dictionary
        ):
            res.dictionary.CopyFrom(
                InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary.to_proto(
                    resource.dictionary
                )
            )
        else:
            res.ClearField("dictionary")
        if InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex.to_proto(
            resource.regex
        ):
            res.regex.CopyFrom(
                InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex.to_proto(
                    resource.regex
                )
            )
        else:
            res.ClearField("regex")
        if InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes.to_proto(
            resource.exclude_info_types
        ):
            res.exclude_info_types.CopyFrom(
                InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes.to_proto(
                    resource.exclude_info_types
                )
            )
        else:
            res.ClearField("exclude_info_types")
        if InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.to_proto(
            resource.matching_type
        ):
            res.matching_type = InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.to_proto(
                resource.matching_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigRuleSetRulesExclusionRule(
            dictionary=InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary.from_proto(
                resource.dictionary
            ),
            regex=InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex.from_proto(
                resource.regex
            ),
            exclude_info_types=InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes.from_proto(
                resource.exclude_info_types
            ),
            matching_type=InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.from_proto(
                resource.matching_type
            ),
        )


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRule.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRule.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(object):
    def __init__(self, word_list: dict = None, cloud_storage_path: dict = None):
        self.word_list = word_list
        self.cloud_storage_path = cloud_storage_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary()
        )
        if InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList.to_proto(
            resource.word_list
        ):
            res.word_list.CopyFrom(
                InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList.to_proto(
                    resource.word_list
                )
            )
        else:
            res.ClearField("word_list")
        if InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath.to_proto(
            resource.cloud_storage_path
        ):
            res.cloud_storage_path.CopyFrom(
                InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath.to_proto(
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

        return InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(
            word_list=InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList.from_proto(
                resource.word_list
            ),
            cloud_storage_path=InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath.from_proto(
                resource.cloud_storage_path
            ),
        )


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary.from_proto(
                i
            )
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(object):
    def __init__(self, words: list = None):
        self.words = words

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList()
        )
        if Primitive.to_proto(resource.words):
            res.words.extend(Primitive.to_proto(resource.words))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(
            words=Primitive.from_proto(resource.words),
        )


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList.from_proto(
                i
            )
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(
    object
):
    def __init__(self, path: str = None):
        self.path = path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath()
        )
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(
            path=Primitive.from_proto(resource.path),
        )


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath.from_proto(
                i
            )
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(object):
    def __init__(self, pattern: str = None, group_indexes: list = None):
        self.pattern = pattern
        self.group_indexes = group_indexes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex()
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

        return InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(
            pattern=Primitive.from_proto(resource.pattern),
            group_indexes=int64Array.from_proto(resource.group_indexes),
        )


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex.from_proto(i)
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(object):
    def __init__(self, info_types: list = None):
        self.info_types = info_types

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes()
        )
        if InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesArray.to_proto(
            resource.info_types
        ):
            res.info_types.extend(
                InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesArray.to_proto(
                    resource.info_types
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(
            info_types=InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesArray.from_proto(
                resource.info_types
            ),
        )


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes.from_proto(
                i
            )
            for i in resources
        ]


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(
            name=Primitive.from_proto(resource.name),
        )


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes.from_proto(
                i
            )
            for i in resources
        ]


class InspectTemplateInspectConfigMinLikelihoodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum.Value(
            "DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum.Name(
            resource
        )[
            len("DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum") :
        ]


class InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum.Value(
            "DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum.Name(
            resource
        )[
            len("DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum") :
        ]


class InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum.Value(
            "DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum.Name(
            resource
        )[
            len(
                "DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum"
            ) :
        ]


class InspectTemplateInspectConfigContentOptionsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum.Value(
            "DlpAlphaInspectTemplateInspectConfigContentOptionsEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum.Name(
            resource
        )[
            len("DlpAlphaInspectTemplateInspectConfigContentOptionsEnum") :
        ]


class InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.Value(
            "DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.Name(
            resource
        )[
            len(
                "DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"
            ) :
        ]


class InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.Value(
            "DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return inspect_template_pb2.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.Name(
            resource
        )[
            len(
                "DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"
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
