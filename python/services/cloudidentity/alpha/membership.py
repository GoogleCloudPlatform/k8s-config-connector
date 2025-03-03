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
from google3.cloud.graphite.mmv2.services.google.cloudidentity import membership_pb2
from google3.cloud.graphite.mmv2.services.google.cloudidentity import (
    membership_pb2_grpc,
)

from typing import List


class Membership(object):
    def __init__(
        self,
        name: str = None,
        preferred_member_key: dict = None,
        create_time: str = None,
        update_time: str = None,
        roles: list = None,
        type: str = None,
        delivery_setting: str = None,
        display_name: dict = None,
        member_key: dict = None,
        group: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.preferred_member_key = preferred_member_key
        self.roles = roles
        self.member_key = member_key
        self.group = group
        self.service_account_file = service_account_file

    def apply(self):
        stub = membership_pb2_grpc.CloudidentityAlphaMembershipServiceStub(
            channel.Channel()
        )
        request = membership_pb2.ApplyCloudidentityAlphaMembershipRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if MembershipPreferredMemberKey.to_proto(self.preferred_member_key):
            request.resource.preferred_member_key.CopyFrom(
                MembershipPreferredMemberKey.to_proto(self.preferred_member_key)
            )
        else:
            request.resource.ClearField("preferred_member_key")
        if MembershipRolesArray.to_proto(self.roles):
            request.resource.roles.extend(MembershipRolesArray.to_proto(self.roles))
        if MembershipMemberKey.to_proto(self.member_key):
            request.resource.member_key.CopyFrom(
                MembershipMemberKey.to_proto(self.member_key)
            )
        else:
            request.resource.ClearField("member_key")
        if Primitive.to_proto(self.group):
            request.resource.group = Primitive.to_proto(self.group)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudidentityAlphaMembership(request)
        self.name = Primitive.from_proto(response.name)
        self.preferred_member_key = MembershipPreferredMemberKey.from_proto(
            response.preferred_member_key
        )
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.roles = MembershipRolesArray.from_proto(response.roles)
        self.type = MembershipTypeEnum.from_proto(response.type)
        self.delivery_setting = MembershipDeliverySettingEnum.from_proto(
            response.delivery_setting
        )
        self.display_name = MembershipDisplayName.from_proto(response.display_name)
        self.member_key = MembershipMemberKey.from_proto(response.member_key)
        self.group = Primitive.from_proto(response.group)

    def delete(self):
        stub = membership_pb2_grpc.CloudidentityAlphaMembershipServiceStub(
            channel.Channel()
        )
        request = membership_pb2.DeleteCloudidentityAlphaMembershipRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if MembershipPreferredMemberKey.to_proto(self.preferred_member_key):
            request.resource.preferred_member_key.CopyFrom(
                MembershipPreferredMemberKey.to_proto(self.preferred_member_key)
            )
        else:
            request.resource.ClearField("preferred_member_key")
        if MembershipRolesArray.to_proto(self.roles):
            request.resource.roles.extend(MembershipRolesArray.to_proto(self.roles))
        if MembershipMemberKey.to_proto(self.member_key):
            request.resource.member_key.CopyFrom(
                MembershipMemberKey.to_proto(self.member_key)
            )
        else:
            request.resource.ClearField("member_key")
        if Primitive.to_proto(self.group):
            request.resource.group = Primitive.to_proto(self.group)

        response = stub.DeleteCloudidentityAlphaMembership(request)

    @classmethod
    def list(self, group, service_account_file=""):
        stub = membership_pb2_grpc.CloudidentityAlphaMembershipServiceStub(
            channel.Channel()
        )
        request = membership_pb2.ListCloudidentityAlphaMembershipRequest()
        request.service_account_file = service_account_file
        request.Group = group

        return stub.ListCloudidentityAlphaMembership(request).items

    def to_proto(self):
        resource = membership_pb2.CloudidentityAlphaMembership()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if MembershipPreferredMemberKey.to_proto(self.preferred_member_key):
            resource.preferred_member_key.CopyFrom(
                MembershipPreferredMemberKey.to_proto(self.preferred_member_key)
            )
        else:
            resource.ClearField("preferred_member_key")
        if MembershipRolesArray.to_proto(self.roles):
            resource.roles.extend(MembershipRolesArray.to_proto(self.roles))
        if MembershipMemberKey.to_proto(self.member_key):
            resource.member_key.CopyFrom(MembershipMemberKey.to_proto(self.member_key))
        else:
            resource.ClearField("member_key")
        if Primitive.to_proto(self.group):
            resource.group = Primitive.to_proto(self.group)
        return resource


class MembershipPreferredMemberKey(object):
    def __init__(self, id: str = None, namespace: str = None):
        self.id = id
        self.namespace = namespace

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.CloudidentityAlphaMembershipPreferredMemberKey()
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.namespace):
            res.namespace = Primitive.to_proto(resource.namespace)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipPreferredMemberKey(
            id=Primitive.from_proto(resource.id),
            namespace=Primitive.from_proto(resource.namespace),
        )


class MembershipPreferredMemberKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipPreferredMemberKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipPreferredMemberKey.from_proto(i) for i in resources]


class MembershipRoles(object):
    def __init__(
        self,
        name: str = None,
        expiry_detail: dict = None,
        restriction_evaluations: dict = None,
    ):
        self.name = name
        self.expiry_detail = expiry_detail
        self.restriction_evaluations = restriction_evaluations

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.CloudidentityAlphaMembershipRoles()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if MembershipRolesExpiryDetail.to_proto(resource.expiry_detail):
            res.expiry_detail.CopyFrom(
                MembershipRolesExpiryDetail.to_proto(resource.expiry_detail)
            )
        else:
            res.ClearField("expiry_detail")
        if MembershipRolesRestrictionEvaluations.to_proto(
            resource.restriction_evaluations
        ):
            res.restriction_evaluations.CopyFrom(
                MembershipRolesRestrictionEvaluations.to_proto(
                    resource.restriction_evaluations
                )
            )
        else:
            res.ClearField("restriction_evaluations")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipRoles(
            name=Primitive.from_proto(resource.name),
            expiry_detail=MembershipRolesExpiryDetail.from_proto(
                resource.expiry_detail
            ),
            restriction_evaluations=MembershipRolesRestrictionEvaluations.from_proto(
                resource.restriction_evaluations
            ),
        )


class MembershipRolesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipRoles.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipRoles.from_proto(i) for i in resources]


class MembershipRolesExpiryDetail(object):
    def __init__(self, expire_time: str = None):
        self.expire_time = expire_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.CloudidentityAlphaMembershipRolesExpiryDetail()
        if Primitive.to_proto(resource.expire_time):
            res.expire_time = Primitive.to_proto(resource.expire_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipRolesExpiryDetail(
            expire_time=Primitive.from_proto(resource.expire_time),
        )


class MembershipRolesExpiryDetailArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipRolesExpiryDetail.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipRolesExpiryDetail.from_proto(i) for i in resources]


class MembershipRolesRestrictionEvaluations(object):
    def __init__(self, member_restriction_evaluation: dict = None):
        self.member_restriction_evaluation = member_restriction_evaluation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.CloudidentityAlphaMembershipRolesRestrictionEvaluations()
        if MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation.to_proto(
            resource.member_restriction_evaluation
        ):
            res.member_restriction_evaluation.CopyFrom(
                MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation.to_proto(
                    resource.member_restriction_evaluation
                )
            )
        else:
            res.ClearField("member_restriction_evaluation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipRolesRestrictionEvaluations(
            member_restriction_evaluation=MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation.from_proto(
                resource.member_restriction_evaluation
            ),
        )


class MembershipRolesRestrictionEvaluationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipRolesRestrictionEvaluations.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipRolesRestrictionEvaluations.from_proto(i) for i in resources]


class MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(object):
    def __init__(self, state: str = None):
        self.state = state

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            membership_pb2.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation()
        )
        if MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum.to_proto(
            resource.state
        ):
            res.state = MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum.to_proto(
                resource.state
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(
            state=MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum.from_proto(
                resource.state
            ),
        )


class MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation.from_proto(
                i
            )
            for i in resources
        ]


class MembershipDisplayName(object):
    def __init__(
        self, given_name: str = None, family_name: str = None, full_name: str = None
    ):
        self.given_name = given_name
        self.family_name = family_name
        self.full_name = full_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.CloudidentityAlphaMembershipDisplayName()
        if Primitive.to_proto(resource.given_name):
            res.given_name = Primitive.to_proto(resource.given_name)
        if Primitive.to_proto(resource.family_name):
            res.family_name = Primitive.to_proto(resource.family_name)
        if Primitive.to_proto(resource.full_name):
            res.full_name = Primitive.to_proto(resource.full_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipDisplayName(
            given_name=Primitive.from_proto(resource.given_name),
            family_name=Primitive.from_proto(resource.family_name),
            full_name=Primitive.from_proto(resource.full_name),
        )


class MembershipDisplayNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipDisplayName.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipDisplayName.from_proto(i) for i in resources]


class MembershipMemberKey(object):
    def __init__(self, id: str = None, namespace: str = None):
        self.id = id
        self.namespace = namespace

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = membership_pb2.CloudidentityAlphaMembershipMemberKey()
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.namespace):
            res.namespace = Primitive.to_proto(resource.namespace)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MembershipMemberKey(
            id=Primitive.from_proto(resource.id),
            namespace=Primitive.from_proto(resource.namespace),
        )


class MembershipMemberKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MembershipMemberKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MembershipMemberKey.from_proto(i) for i in resources]


class MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return membership_pb2.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum.Value(
            "CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return membership_pb2.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum.Name(
            resource
        )[
            len(
                "CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum"
            ) :
        ]


class MembershipTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return membership_pb2.CloudidentityAlphaMembershipTypeEnum.Value(
            "CloudidentityAlphaMembershipTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return membership_pb2.CloudidentityAlphaMembershipTypeEnum.Name(resource)[
            len("CloudidentityAlphaMembershipTypeEnum") :
        ]


class MembershipDeliverySettingEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return membership_pb2.CloudidentityAlphaMembershipDeliverySettingEnum.Value(
            "CloudidentityAlphaMembershipDeliverySettingEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return membership_pb2.CloudidentityAlphaMembershipDeliverySettingEnum.Name(
            resource
        )[len("CloudidentityAlphaMembershipDeliverySettingEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
