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
from google3.cloud.graphite.mmv2.services.google.dlp import stored_info_type_pb2
from google3.cloud.graphite.mmv2.services.google.dlp import stored_info_type_pb2_grpc

from typing import List


class StoredInfoType(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        description: str = None,
        large_custom_dictionary: dict = None,
        dictionary: dict = None,
        regex: dict = None,
        parent: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.description = description
        self.large_custom_dictionary = large_custom_dictionary
        self.dictionary = dictionary
        self.regex = regex
        self.parent = parent
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = stored_info_type_pb2_grpc.DlpBetaStoredInfoTypeServiceStub(
            channel.Channel()
        )
        request = stored_info_type_pb2.ApplyDlpBetaStoredInfoTypeRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if StoredInfoTypeLargeCustomDictionary.to_proto(self.large_custom_dictionary):
            request.resource.large_custom_dictionary.CopyFrom(
                StoredInfoTypeLargeCustomDictionary.to_proto(
                    self.large_custom_dictionary
                )
            )
        else:
            request.resource.ClearField("large_custom_dictionary")
        if StoredInfoTypeDictionary.to_proto(self.dictionary):
            request.resource.dictionary.CopyFrom(
                StoredInfoTypeDictionary.to_proto(self.dictionary)
            )
        else:
            request.resource.ClearField("dictionary")
        if StoredInfoTypeRegex.to_proto(self.regex):
            request.resource.regex.CopyFrom(StoredInfoTypeRegex.to_proto(self.regex))
        else:
            request.resource.ClearField("regex")
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDlpBetaStoredInfoType(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.large_custom_dictionary = StoredInfoTypeLargeCustomDictionary.from_proto(
            response.large_custom_dictionary
        )
        self.dictionary = StoredInfoTypeDictionary.from_proto(response.dictionary)
        self.regex = StoredInfoTypeRegex.from_proto(response.regex)
        self.parent = Primitive.from_proto(response.parent)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = stored_info_type_pb2_grpc.DlpBetaStoredInfoTypeServiceStub(
            channel.Channel()
        )
        request = stored_info_type_pb2.DeleteDlpBetaStoredInfoTypeRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if StoredInfoTypeLargeCustomDictionary.to_proto(self.large_custom_dictionary):
            request.resource.large_custom_dictionary.CopyFrom(
                StoredInfoTypeLargeCustomDictionary.to_proto(
                    self.large_custom_dictionary
                )
            )
        else:
            request.resource.ClearField("large_custom_dictionary")
        if StoredInfoTypeDictionary.to_proto(self.dictionary):
            request.resource.dictionary.CopyFrom(
                StoredInfoTypeDictionary.to_proto(self.dictionary)
            )
        else:
            request.resource.ClearField("dictionary")
        if StoredInfoTypeRegex.to_proto(self.regex):
            request.resource.regex.CopyFrom(StoredInfoTypeRegex.to_proto(self.regex))
        else:
            request.resource.ClearField("regex")
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteDlpBetaStoredInfoType(request)

    @classmethod
    def list(self, location, parent, service_account_file=""):
        stub = stored_info_type_pb2_grpc.DlpBetaStoredInfoTypeServiceStub(
            channel.Channel()
        )
        request = stored_info_type_pb2.ListDlpBetaStoredInfoTypeRequest()
        request.service_account_file = service_account_file
        request.Location = location

        request.Parent = parent

        return stub.ListDlpBetaStoredInfoType(request).items

    def to_proto(self):
        resource = stored_info_type_pb2.DlpBetaStoredInfoType()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if StoredInfoTypeLargeCustomDictionary.to_proto(self.large_custom_dictionary):
            resource.large_custom_dictionary.CopyFrom(
                StoredInfoTypeLargeCustomDictionary.to_proto(
                    self.large_custom_dictionary
                )
            )
        else:
            resource.ClearField("large_custom_dictionary")
        if StoredInfoTypeDictionary.to_proto(self.dictionary):
            resource.dictionary.CopyFrom(
                StoredInfoTypeDictionary.to_proto(self.dictionary)
            )
        else:
            resource.ClearField("dictionary")
        if StoredInfoTypeRegex.to_proto(self.regex):
            resource.regex.CopyFrom(StoredInfoTypeRegex.to_proto(self.regex))
        else:
            resource.ClearField("regex")
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class StoredInfoTypeLargeCustomDictionary(object):
    def __init__(
        self,
        output_path: dict = None,
        cloud_storage_file_set: dict = None,
        big_query_field: dict = None,
    ):
        self.output_path = output_path
        self.cloud_storage_file_set = cloud_storage_file_set
        self.big_query_field = big_query_field

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = stored_info_type_pb2.DlpBetaStoredInfoTypeLargeCustomDictionary()
        if StoredInfoTypeLargeCustomDictionaryOutputPath.to_proto(resource.output_path):
            res.output_path.CopyFrom(
                StoredInfoTypeLargeCustomDictionaryOutputPath.to_proto(
                    resource.output_path
                )
            )
        else:
            res.ClearField("output_path")
        if StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet.to_proto(
            resource.cloud_storage_file_set
        ):
            res.cloud_storage_file_set.CopyFrom(
                StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet.to_proto(
                    resource.cloud_storage_file_set
                )
            )
        else:
            res.ClearField("cloud_storage_file_set")
        if StoredInfoTypeLargeCustomDictionaryBigQueryField.to_proto(
            resource.big_query_field
        ):
            res.big_query_field.CopyFrom(
                StoredInfoTypeLargeCustomDictionaryBigQueryField.to_proto(
                    resource.big_query_field
                )
            )
        else:
            res.ClearField("big_query_field")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return StoredInfoTypeLargeCustomDictionary(
            output_path=StoredInfoTypeLargeCustomDictionaryOutputPath.from_proto(
                resource.output_path
            ),
            cloud_storage_file_set=StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet.from_proto(
                resource.cloud_storage_file_set
            ),
            big_query_field=StoredInfoTypeLargeCustomDictionaryBigQueryField.from_proto(
                resource.big_query_field
            ),
        )


class StoredInfoTypeLargeCustomDictionaryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [StoredInfoTypeLargeCustomDictionary.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [StoredInfoTypeLargeCustomDictionary.from_proto(i) for i in resources]


class StoredInfoTypeLargeCustomDictionaryOutputPath(object):
    def __init__(self, path: str = None):
        self.path = path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            stored_info_type_pb2.DlpBetaStoredInfoTypeLargeCustomDictionaryOutputPath()
        )
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return StoredInfoTypeLargeCustomDictionaryOutputPath(
            path=Primitive.from_proto(resource.path),
        )


class StoredInfoTypeLargeCustomDictionaryOutputPathArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            StoredInfoTypeLargeCustomDictionaryOutputPath.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            StoredInfoTypeLargeCustomDictionaryOutputPath.from_proto(i)
            for i in resources
        ]


class StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(object):
    def __init__(self, url: str = None):
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            stored_info_type_pb2.DlpBetaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet()
        )
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(
            url=Primitive.from_proto(resource.url),
        )


class StoredInfoTypeLargeCustomDictionaryCloudStorageFileSetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet.from_proto(i)
            for i in resources
        ]


class StoredInfoTypeLargeCustomDictionaryBigQueryField(object):
    def __init__(self, table: dict = None, field: dict = None):
        self.table = table
        self.field = field

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            stored_info_type_pb2.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryField()
        )
        if StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable.to_proto(
            resource.table
        ):
            res.table.CopyFrom(
                StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable.to_proto(
                    resource.table
                )
            )
        else:
            res.ClearField("table")
        if StoredInfoTypeLargeCustomDictionaryBigQueryFieldField.to_proto(
            resource.field
        ):
            res.field.CopyFrom(
                StoredInfoTypeLargeCustomDictionaryBigQueryFieldField.to_proto(
                    resource.field
                )
            )
        else:
            res.ClearField("field")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return StoredInfoTypeLargeCustomDictionaryBigQueryField(
            table=StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable.from_proto(
                resource.table
            ),
            field=StoredInfoTypeLargeCustomDictionaryBigQueryFieldField.from_proto(
                resource.field
            ),
        )


class StoredInfoTypeLargeCustomDictionaryBigQueryFieldArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            StoredInfoTypeLargeCustomDictionaryBigQueryField.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            StoredInfoTypeLargeCustomDictionaryBigQueryField.from_proto(i)
            for i in resources
        ]


class StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(object):
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
            stored_info_type_pb2.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable()
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

        return StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(
            project_id=Primitive.from_proto(resource.project_id),
            dataset_id=Primitive.from_proto(resource.dataset_id),
            table_id=Primitive.from_proto(resource.table_id),
        )


class StoredInfoTypeLargeCustomDictionaryBigQueryFieldTableArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable.from_proto(i)
            for i in resources
        ]


class StoredInfoTypeLargeCustomDictionaryBigQueryFieldField(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            stored_info_type_pb2.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldField()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return StoredInfoTypeLargeCustomDictionaryBigQueryFieldField(
            name=Primitive.from_proto(resource.name),
        )


class StoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            StoredInfoTypeLargeCustomDictionaryBigQueryFieldField.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            StoredInfoTypeLargeCustomDictionaryBigQueryFieldField.from_proto(i)
            for i in resources
        ]


class StoredInfoTypeDictionary(object):
    def __init__(self, word_list: dict = None, cloud_storage_path: dict = None):
        self.word_list = word_list
        self.cloud_storage_path = cloud_storage_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = stored_info_type_pb2.DlpBetaStoredInfoTypeDictionary()
        if StoredInfoTypeDictionaryWordList.to_proto(resource.word_list):
            res.word_list.CopyFrom(
                StoredInfoTypeDictionaryWordList.to_proto(resource.word_list)
            )
        else:
            res.ClearField("word_list")
        if StoredInfoTypeDictionaryCloudStoragePath.to_proto(
            resource.cloud_storage_path
        ):
            res.cloud_storage_path.CopyFrom(
                StoredInfoTypeDictionaryCloudStoragePath.to_proto(
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

        return StoredInfoTypeDictionary(
            word_list=StoredInfoTypeDictionaryWordList.from_proto(resource.word_list),
            cloud_storage_path=StoredInfoTypeDictionaryCloudStoragePath.from_proto(
                resource.cloud_storage_path
            ),
        )


class StoredInfoTypeDictionaryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [StoredInfoTypeDictionary.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [StoredInfoTypeDictionary.from_proto(i) for i in resources]


class StoredInfoTypeDictionaryWordList(object):
    def __init__(self, words: list = None):
        self.words = words

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = stored_info_type_pb2.DlpBetaStoredInfoTypeDictionaryWordList()
        if Primitive.to_proto(resource.words):
            res.words.extend(Primitive.to_proto(resource.words))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return StoredInfoTypeDictionaryWordList(
            words=Primitive.from_proto(resource.words),
        )


class StoredInfoTypeDictionaryWordListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [StoredInfoTypeDictionaryWordList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [StoredInfoTypeDictionaryWordList.from_proto(i) for i in resources]


class StoredInfoTypeDictionaryCloudStoragePath(object):
    def __init__(self, path: str = None):
        self.path = path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = stored_info_type_pb2.DlpBetaStoredInfoTypeDictionaryCloudStoragePath()
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return StoredInfoTypeDictionaryCloudStoragePath(
            path=Primitive.from_proto(resource.path),
        )


class StoredInfoTypeDictionaryCloudStoragePathArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [StoredInfoTypeDictionaryCloudStoragePath.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            StoredInfoTypeDictionaryCloudStoragePath.from_proto(i) for i in resources
        ]


class StoredInfoTypeRegex(object):
    def __init__(self, pattern: str = None, group_indexes: list = None):
        self.pattern = pattern
        self.group_indexes = group_indexes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = stored_info_type_pb2.DlpBetaStoredInfoTypeRegex()
        if Primitive.to_proto(resource.pattern):
            res.pattern = Primitive.to_proto(resource.pattern)
        if int64Array.to_proto(resource.group_indexes):
            res.group_indexes.extend(int64Array.to_proto(resource.group_indexes))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return StoredInfoTypeRegex(
            pattern=Primitive.from_proto(resource.pattern),
            group_indexes=int64Array.from_proto(resource.group_indexes),
        )


class StoredInfoTypeRegexArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [StoredInfoTypeRegex.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [StoredInfoTypeRegex.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
