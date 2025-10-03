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
from google3.cloud.graphite.mmv2.services.google.cloudidentity import group_pb2
from google3.cloud.graphite.mmv2.services.google.cloudidentity import group_pb2_grpc

from typing import List


class Group(object):
    def __init__(
        self,
        name: str = None,
        group_key: dict = None,
        additional_group_keys: list = None,
        parent: str = None,
        display_name: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        direct_member_count: int = None,
        direct_member_count_per_type: dict = None,
        derived_aliases: list = None,
        dynamic_group_metadata: dict = None,
        posix_groups: list = None,
        initial_group_config: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.group_key = group_key
        self.additional_group_keys = additional_group_keys
        self.parent = parent
        self.display_name = display_name
        self.description = description
        self.labels = labels
        self.dynamic_group_metadata = dynamic_group_metadata
        self.posix_groups = posix_groups
        self.initial_group_config = initial_group_config
        self.service_account_file = service_account_file

    def apply(self):
        stub = group_pb2_grpc.CloudidentityAlphaGroupServiceStub(channel.Channel())
        request = group_pb2.ApplyCloudidentityAlphaGroupRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if GroupGroupKey.to_proto(self.group_key):
            request.resource.group_key.CopyFrom(GroupGroupKey.to_proto(self.group_key))
        else:
            request.resource.ClearField("group_key")
        if GroupAdditionalGroupKeysArray.to_proto(self.additional_group_keys):
            request.resource.additional_group_keys.extend(
                GroupAdditionalGroupKeysArray.to_proto(self.additional_group_keys)
            )
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if GroupDynamicGroupMetadata.to_proto(self.dynamic_group_metadata):
            request.resource.dynamic_group_metadata.CopyFrom(
                GroupDynamicGroupMetadata.to_proto(self.dynamic_group_metadata)
            )
        else:
            request.resource.ClearField("dynamic_group_metadata")
        if GroupPosixGroupsArray.to_proto(self.posix_groups):
            request.resource.posix_groups.extend(
                GroupPosixGroupsArray.to_proto(self.posix_groups)
            )
        if GroupInitialGroupConfigEnum.to_proto(self.initial_group_config):
            request.resource.initial_group_config = (
                GroupInitialGroupConfigEnum.to_proto(self.initial_group_config)
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudidentityAlphaGroup(request)
        self.name = Primitive.from_proto(response.name)
        self.group_key = GroupGroupKey.from_proto(response.group_key)
        self.additional_group_keys = GroupAdditionalGroupKeysArray.from_proto(
            response.additional_group_keys
        )
        self.parent = Primitive.from_proto(response.parent)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.direct_member_count = Primitive.from_proto(response.direct_member_count)
        self.direct_member_count_per_type = GroupDirectMemberCountPerType.from_proto(
            response.direct_member_count_per_type
        )
        self.derived_aliases = GroupDerivedAliasesArray.from_proto(
            response.derived_aliases
        )
        self.dynamic_group_metadata = GroupDynamicGroupMetadata.from_proto(
            response.dynamic_group_metadata
        )
        self.posix_groups = GroupPosixGroupsArray.from_proto(response.posix_groups)
        self.initial_group_config = GroupInitialGroupConfigEnum.from_proto(
            response.initial_group_config
        )

    def delete(self):
        stub = group_pb2_grpc.CloudidentityAlphaGroupServiceStub(channel.Channel())
        request = group_pb2.DeleteCloudidentityAlphaGroupRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if GroupGroupKey.to_proto(self.group_key):
            request.resource.group_key.CopyFrom(GroupGroupKey.to_proto(self.group_key))
        else:
            request.resource.ClearField("group_key")
        if GroupAdditionalGroupKeysArray.to_proto(self.additional_group_keys):
            request.resource.additional_group_keys.extend(
                GroupAdditionalGroupKeysArray.to_proto(self.additional_group_keys)
            )
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if GroupDynamicGroupMetadata.to_proto(self.dynamic_group_metadata):
            request.resource.dynamic_group_metadata.CopyFrom(
                GroupDynamicGroupMetadata.to_proto(self.dynamic_group_metadata)
            )
        else:
            request.resource.ClearField("dynamic_group_metadata")
        if GroupPosixGroupsArray.to_proto(self.posix_groups):
            request.resource.posix_groups.extend(
                GroupPosixGroupsArray.to_proto(self.posix_groups)
            )
        if GroupInitialGroupConfigEnum.to_proto(self.initial_group_config):
            request.resource.initial_group_config = (
                GroupInitialGroupConfigEnum.to_proto(self.initial_group_config)
            )

        response = stub.DeleteCloudidentityAlphaGroup(request)

    @classmethod
    def list(self, parent, service_account_file=""):
        stub = group_pb2_grpc.CloudidentityAlphaGroupServiceStub(channel.Channel())
        request = group_pb2.ListCloudidentityAlphaGroupRequest()
        request.service_account_file = service_account_file
        request.Parent = parent

        return stub.ListCloudidentityAlphaGroup(request).items

    def to_proto(self):
        resource = group_pb2.CloudidentityAlphaGroup()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if GroupGroupKey.to_proto(self.group_key):
            resource.group_key.CopyFrom(GroupGroupKey.to_proto(self.group_key))
        else:
            resource.ClearField("group_key")
        if GroupAdditionalGroupKeysArray.to_proto(self.additional_group_keys):
            resource.additional_group_keys.extend(
                GroupAdditionalGroupKeysArray.to_proto(self.additional_group_keys)
            )
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if GroupDynamicGroupMetadata.to_proto(self.dynamic_group_metadata):
            resource.dynamic_group_metadata.CopyFrom(
                GroupDynamicGroupMetadata.to_proto(self.dynamic_group_metadata)
            )
        else:
            resource.ClearField("dynamic_group_metadata")
        if GroupPosixGroupsArray.to_proto(self.posix_groups):
            resource.posix_groups.extend(
                GroupPosixGroupsArray.to_proto(self.posix_groups)
            )
        if GroupInitialGroupConfigEnum.to_proto(self.initial_group_config):
            resource.initial_group_config = GroupInitialGroupConfigEnum.to_proto(
                self.initial_group_config
            )
        return resource


class GroupGroupKey(object):
    def __init__(self, id: str = None, namespace: str = None):
        self.id = id
        self.namespace = namespace

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = group_pb2.CloudidentityAlphaGroupGroupKey()
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.namespace):
            res.namespace = Primitive.to_proto(resource.namespace)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GroupGroupKey(
            id=Primitive.from_proto(resource.id),
            namespace=Primitive.from_proto(resource.namespace),
        )


class GroupGroupKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GroupGroupKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GroupGroupKey.from_proto(i) for i in resources]


class GroupAdditionalGroupKeys(object):
    def __init__(self, id: str = None, namespace: str = None):
        self.id = id
        self.namespace = namespace

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = group_pb2.CloudidentityAlphaGroupAdditionalGroupKeys()
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.namespace):
            res.namespace = Primitive.to_proto(resource.namespace)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GroupAdditionalGroupKeys(
            id=Primitive.from_proto(resource.id),
            namespace=Primitive.from_proto(resource.namespace),
        )


class GroupAdditionalGroupKeysArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GroupAdditionalGroupKeys.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GroupAdditionalGroupKeys.from_proto(i) for i in resources]


class GroupDirectMemberCountPerType(object):
    def __init__(self, user_count: int = None, group_count: int = None):
        self.user_count = user_count
        self.group_count = group_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = group_pb2.CloudidentityAlphaGroupDirectMemberCountPerType()
        if Primitive.to_proto(resource.user_count):
            res.user_count = Primitive.to_proto(resource.user_count)
        if Primitive.to_proto(resource.group_count):
            res.group_count = Primitive.to_proto(resource.group_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GroupDirectMemberCountPerType(
            user_count=Primitive.from_proto(resource.user_count),
            group_count=Primitive.from_proto(resource.group_count),
        )


class GroupDirectMemberCountPerTypeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GroupDirectMemberCountPerType.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GroupDirectMemberCountPerType.from_proto(i) for i in resources]


class GroupDerivedAliases(object):
    def __init__(self, id: str = None, namespace: str = None):
        self.id = id
        self.namespace = namespace

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = group_pb2.CloudidentityAlphaGroupDerivedAliases()
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.namespace):
            res.namespace = Primitive.to_proto(resource.namespace)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GroupDerivedAliases(
            id=Primitive.from_proto(resource.id),
            namespace=Primitive.from_proto(resource.namespace),
        )


class GroupDerivedAliasesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GroupDerivedAliases.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GroupDerivedAliases.from_proto(i) for i in resources]


class GroupDynamicGroupMetadata(object):
    def __init__(self, queries: list = None, status: dict = None):
        self.queries = queries
        self.status = status

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = group_pb2.CloudidentityAlphaGroupDynamicGroupMetadata()
        if GroupDynamicGroupMetadataQueriesArray.to_proto(resource.queries):
            res.queries.extend(
                GroupDynamicGroupMetadataQueriesArray.to_proto(resource.queries)
            )
        if GroupDynamicGroupMetadataStatus.to_proto(resource.status):
            res.status.CopyFrom(
                GroupDynamicGroupMetadataStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GroupDynamicGroupMetadata(
            queries=GroupDynamicGroupMetadataQueriesArray.from_proto(resource.queries),
            status=GroupDynamicGroupMetadataStatus.from_proto(resource.status),
        )


class GroupDynamicGroupMetadataArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GroupDynamicGroupMetadata.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GroupDynamicGroupMetadata.from_proto(i) for i in resources]


class GroupDynamicGroupMetadataQueries(object):
    def __init__(self, resource_type: str = None, query: str = None):
        self.resource_type = resource_type
        self.query = query

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = group_pb2.CloudidentityAlphaGroupDynamicGroupMetadataQueries()
        if GroupDynamicGroupMetadataQueriesResourceTypeEnum.to_proto(
            resource.resource_type
        ):
            res.resource_type = (
                GroupDynamicGroupMetadataQueriesResourceTypeEnum.to_proto(
                    resource.resource_type
                )
            )
        if Primitive.to_proto(resource.query):
            res.query = Primitive.to_proto(resource.query)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GroupDynamicGroupMetadataQueries(
            resource_type=GroupDynamicGroupMetadataQueriesResourceTypeEnum.from_proto(
                resource.resource_type
            ),
            query=Primitive.from_proto(resource.query),
        )


class GroupDynamicGroupMetadataQueriesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GroupDynamicGroupMetadataQueries.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GroupDynamicGroupMetadataQueries.from_proto(i) for i in resources]


class GroupDynamicGroupMetadataStatus(object):
    def __init__(self, status: str = None, status_time: str = None):
        self.status = status
        self.status_time = status_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = group_pb2.CloudidentityAlphaGroupDynamicGroupMetadataStatus()
        if GroupDynamicGroupMetadataStatusStatusEnum.to_proto(resource.status):
            res.status = GroupDynamicGroupMetadataStatusStatusEnum.to_proto(
                resource.status
            )
        if Primitive.to_proto(resource.status_time):
            res.status_time = Primitive.to_proto(resource.status_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GroupDynamicGroupMetadataStatus(
            status=GroupDynamicGroupMetadataStatusStatusEnum.from_proto(
                resource.status
            ),
            status_time=Primitive.from_proto(resource.status_time),
        )


class GroupDynamicGroupMetadataStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GroupDynamicGroupMetadataStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GroupDynamicGroupMetadataStatus.from_proto(i) for i in resources]


class GroupPosixGroups(object):
    def __init__(self, name: str = None, gid: str = None, system_id: str = None):
        self.name = name
        self.gid = gid
        self.system_id = system_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = group_pb2.CloudidentityAlphaGroupPosixGroups()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.gid):
            res.gid = Primitive.to_proto(resource.gid)
        if Primitive.to_proto(resource.system_id):
            res.system_id = Primitive.to_proto(resource.system_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GroupPosixGroups(
            name=Primitive.from_proto(resource.name),
            gid=Primitive.from_proto(resource.gid),
            system_id=Primitive.from_proto(resource.system_id),
        )


class GroupPosixGroupsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GroupPosixGroups.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GroupPosixGroups.from_proto(i) for i in resources]


class GroupDynamicGroupMetadataQueriesResourceTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return group_pb2.CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum.Value(
            "CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return group_pb2.CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum.Name(
            resource
        )[
            len("CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum") :
        ]


class GroupDynamicGroupMetadataStatusStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            group_pb2.CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum.Value(
                "CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum%s"
                % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            group_pb2.CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum.Name(
                resource
            )[len("CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum") :]
        )


class GroupInitialGroupConfigEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return group_pb2.CloudidentityAlphaGroupInitialGroupConfigEnum.Value(
            "CloudidentityAlphaGroupInitialGroupConfigEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return group_pb2.CloudidentityAlphaGroupInitialGroupConfigEnum.Name(resource)[
            len("CloudidentityAlphaGroupInitialGroupConfigEnum") :
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
