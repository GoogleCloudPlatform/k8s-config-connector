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
from google3.cloud.graphite.mmv2.services.google.org_policy import policy_pb2
from google3.cloud.graphite.mmv2.services.google.org_policy import policy_pb2_grpc

from typing import List


class Policy(object):
    def __init__(
        self,
        name: str = None,
        spec: dict = None,
        dry_run_spec: dict = None,
        etag: str = None,
        parent: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.spec = spec
        self.dry_run_spec = dry_run_spec
        self.parent = parent
        self.service_account_file = service_account_file

    def apply(self):
        stub = policy_pb2_grpc.OrgpolicyPolicyServiceStub(channel.Channel())
        request = policy_pb2.ApplyOrgpolicyPolicyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if PolicySpec.to_proto(self.spec):
            request.resource.spec.CopyFrom(PolicySpec.to_proto(self.spec))
        else:
            request.resource.ClearField("spec")
        if PolicyDryRunSpec.to_proto(self.dry_run_spec):
            request.resource.dry_run_spec.CopyFrom(
                PolicyDryRunSpec.to_proto(self.dry_run_spec)
            )
        else:
            request.resource.ClearField("dry_run_spec")
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        request.service_account_file = self.service_account_file

        response = stub.ApplyOrgpolicyPolicy(request)
        self.name = Primitive.from_proto(response.name)
        self.spec = PolicySpec.from_proto(response.spec)
        self.dry_run_spec = PolicyDryRunSpec.from_proto(response.dry_run_spec)
        self.etag = Primitive.from_proto(response.etag)
        self.parent = Primitive.from_proto(response.parent)

    def delete(self):
        stub = policy_pb2_grpc.OrgpolicyPolicyServiceStub(channel.Channel())
        request = policy_pb2.DeleteOrgpolicyPolicyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if PolicySpec.to_proto(self.spec):
            request.resource.spec.CopyFrom(PolicySpec.to_proto(self.spec))
        else:
            request.resource.ClearField("spec")
        if PolicyDryRunSpec.to_proto(self.dry_run_spec):
            request.resource.dry_run_spec.CopyFrom(
                PolicyDryRunSpec.to_proto(self.dry_run_spec)
            )
        else:
            request.resource.ClearField("dry_run_spec")
        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        response = stub.DeleteOrgpolicyPolicy(request)

    @classmethod
    def list(self, parent, service_account_file=""):
        stub = policy_pb2_grpc.OrgpolicyPolicyServiceStub(channel.Channel())
        request = policy_pb2.ListOrgpolicyPolicyRequest()
        request.service_account_file = service_account_file
        request.Parent = parent

        return stub.ListOrgpolicyPolicy(request).items

    def to_proto(self):
        resource = policy_pb2.OrgpolicyPolicy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if PolicySpec.to_proto(self.spec):
            resource.spec.CopyFrom(PolicySpec.to_proto(self.spec))
        else:
            resource.ClearField("spec")
        if PolicyDryRunSpec.to_proto(self.dry_run_spec):
            resource.dry_run_spec.CopyFrom(PolicyDryRunSpec.to_proto(self.dry_run_spec))
        else:
            resource.ClearField("dry_run_spec")
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        return resource


class PolicySpec(object):
    def __init__(
        self,
        etag: str = None,
        update_time: str = None,
        rules: list = None,
        inherit_from_parent: bool = None,
        reset: bool = None,
    ):
        self.etag = etag
        self.update_time = update_time
        self.rules = rules
        self.inherit_from_parent = inherit_from_parent
        self.reset = reset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = policy_pb2.OrgpolicyPolicySpec()
        if Primitive.to_proto(resource.etag):
            res.etag = Primitive.to_proto(resource.etag)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        if PolicySpecRulesArray.to_proto(resource.rules):
            res.rules.extend(PolicySpecRulesArray.to_proto(resource.rules))
        if Primitive.to_proto(resource.inherit_from_parent):
            res.inherit_from_parent = Primitive.to_proto(resource.inherit_from_parent)
        if Primitive.to_proto(resource.reset):
            res.reset = Primitive.to_proto(resource.reset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicySpec(
            etag=Primitive.from_proto(resource.etag),
            update_time=Primitive.from_proto(resource.update_time),
            rules=PolicySpecRulesArray.from_proto(resource.rules),
            inherit_from_parent=Primitive.from_proto(resource.inherit_from_parent),
            reset=Primitive.from_proto(resource.reset),
        )


class PolicySpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicySpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PolicySpec.from_proto(i) for i in resources]


class PolicySpecRules(object):
    def __init__(
        self,
        values: dict = None,
        allow_all: bool = None,
        deny_all: bool = None,
        enforce: bool = None,
        condition: dict = None,
    ):
        self.values = values
        self.allow_all = allow_all
        self.deny_all = deny_all
        self.enforce = enforce
        self.condition = condition

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = policy_pb2.OrgpolicyPolicySpecRules()
        if PolicySpecRulesValues.to_proto(resource.values):
            res.values.CopyFrom(PolicySpecRulesValues.to_proto(resource.values))
        else:
            res.ClearField("values")
        if Primitive.to_proto(resource.allow_all):
            res.allow_all = Primitive.to_proto(resource.allow_all)
        if Primitive.to_proto(resource.deny_all):
            res.deny_all = Primitive.to_proto(resource.deny_all)
        if Primitive.to_proto(resource.enforce):
            res.enforce = Primitive.to_proto(resource.enforce)
        if PolicySpecRulesCondition.to_proto(resource.condition):
            res.condition.CopyFrom(
                PolicySpecRulesCondition.to_proto(resource.condition)
            )
        else:
            res.ClearField("condition")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicySpecRules(
            values=PolicySpecRulesValues.from_proto(resource.values),
            allow_all=Primitive.from_proto(resource.allow_all),
            deny_all=Primitive.from_proto(resource.deny_all),
            enforce=Primitive.from_proto(resource.enforce),
            condition=PolicySpecRulesCondition.from_proto(resource.condition),
        )


class PolicySpecRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicySpecRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PolicySpecRules.from_proto(i) for i in resources]


class PolicySpecRulesValues(object):
    def __init__(self, allowed_values: list = None, denied_values: list = None):
        self.allowed_values = allowed_values
        self.denied_values = denied_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = policy_pb2.OrgpolicyPolicySpecRulesValues()
        if Primitive.to_proto(resource.allowed_values):
            res.allowed_values.extend(Primitive.to_proto(resource.allowed_values))
        if Primitive.to_proto(resource.denied_values):
            res.denied_values.extend(Primitive.to_proto(resource.denied_values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicySpecRulesValues(
            allowed_values=Primitive.from_proto(resource.allowed_values),
            denied_values=Primitive.from_proto(resource.denied_values),
        )


class PolicySpecRulesValuesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicySpecRulesValues.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PolicySpecRulesValues.from_proto(i) for i in resources]


class PolicySpecRulesCondition(object):
    def __init__(
        self,
        expression: str = None,
        title: str = None,
        description: str = None,
        location: str = None,
    ):
        self.expression = expression
        self.title = title
        self.description = description
        self.location = location

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = policy_pb2.OrgpolicyPolicySpecRulesCondition()
        if Primitive.to_proto(resource.expression):
            res.expression = Primitive.to_proto(resource.expression)
        if Primitive.to_proto(resource.title):
            res.title = Primitive.to_proto(resource.title)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicySpecRulesCondition(
            expression=Primitive.from_proto(resource.expression),
            title=Primitive.from_proto(resource.title),
            description=Primitive.from_proto(resource.description),
            location=Primitive.from_proto(resource.location),
        )


class PolicySpecRulesConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicySpecRulesCondition.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PolicySpecRulesCondition.from_proto(i) for i in resources]


class PolicyDryRunSpec(object):
    def __init__(
        self,
        etag: str = None,
        update_time: str = None,
        rules: list = None,
        inherit_from_parent: bool = None,
        reset: bool = None,
    ):
        self.etag = etag
        self.update_time = update_time
        self.rules = rules
        self.inherit_from_parent = inherit_from_parent
        self.reset = reset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = policy_pb2.OrgpolicyPolicyDryRunSpec()
        if Primitive.to_proto(resource.etag):
            res.etag = Primitive.to_proto(resource.etag)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        if PolicyDryRunSpecRulesArray.to_proto(resource.rules):
            res.rules.extend(PolicyDryRunSpecRulesArray.to_proto(resource.rules))
        if Primitive.to_proto(resource.inherit_from_parent):
            res.inherit_from_parent = Primitive.to_proto(resource.inherit_from_parent)
        if Primitive.to_proto(resource.reset):
            res.reset = Primitive.to_proto(resource.reset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicyDryRunSpec(
            etag=Primitive.from_proto(resource.etag),
            update_time=Primitive.from_proto(resource.update_time),
            rules=PolicyDryRunSpecRulesArray.from_proto(resource.rules),
            inherit_from_parent=Primitive.from_proto(resource.inherit_from_parent),
            reset=Primitive.from_proto(resource.reset),
        )


class PolicyDryRunSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicyDryRunSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PolicyDryRunSpec.from_proto(i) for i in resources]


class PolicyDryRunSpecRules(object):
    def __init__(
        self,
        values: dict = None,
        allow_all: bool = None,
        deny_all: bool = None,
        enforce: bool = None,
        condition: dict = None,
    ):
        self.values = values
        self.allow_all = allow_all
        self.deny_all = deny_all
        self.enforce = enforce
        self.condition = condition

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = policy_pb2.OrgpolicyPolicyDryRunSpecRules()
        if PolicyDryRunSpecRulesValues.to_proto(resource.values):
            res.values.CopyFrom(PolicyDryRunSpecRulesValues.to_proto(resource.values))
        else:
            res.ClearField("values")
        if Primitive.to_proto(resource.allow_all):
            res.allow_all = Primitive.to_proto(resource.allow_all)
        if Primitive.to_proto(resource.deny_all):
            res.deny_all = Primitive.to_proto(resource.deny_all)
        if Primitive.to_proto(resource.enforce):
            res.enforce = Primitive.to_proto(resource.enforce)
        if PolicyDryRunSpecRulesCondition.to_proto(resource.condition):
            res.condition.CopyFrom(
                PolicyDryRunSpecRulesCondition.to_proto(resource.condition)
            )
        else:
            res.ClearField("condition")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicyDryRunSpecRules(
            values=PolicyDryRunSpecRulesValues.from_proto(resource.values),
            allow_all=Primitive.from_proto(resource.allow_all),
            deny_all=Primitive.from_proto(resource.deny_all),
            enforce=Primitive.from_proto(resource.enforce),
            condition=PolicyDryRunSpecRulesCondition.from_proto(resource.condition),
        )


class PolicyDryRunSpecRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicyDryRunSpecRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PolicyDryRunSpecRules.from_proto(i) for i in resources]


class PolicyDryRunSpecRulesValues(object):
    def __init__(self, allowed_values: list = None, denied_values: list = None):
        self.allowed_values = allowed_values
        self.denied_values = denied_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = policy_pb2.OrgpolicyPolicyDryRunSpecRulesValues()
        if Primitive.to_proto(resource.allowed_values):
            res.allowed_values.extend(Primitive.to_proto(resource.allowed_values))
        if Primitive.to_proto(resource.denied_values):
            res.denied_values.extend(Primitive.to_proto(resource.denied_values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicyDryRunSpecRulesValues(
            allowed_values=Primitive.from_proto(resource.allowed_values),
            denied_values=Primitive.from_proto(resource.denied_values),
        )


class PolicyDryRunSpecRulesValuesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicyDryRunSpecRulesValues.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PolicyDryRunSpecRulesValues.from_proto(i) for i in resources]


class PolicyDryRunSpecRulesCondition(object):
    def __init__(
        self,
        expression: str = None,
        title: str = None,
        description: str = None,
        location: str = None,
    ):
        self.expression = expression
        self.title = title
        self.description = description
        self.location = location

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = policy_pb2.OrgpolicyPolicyDryRunSpecRulesCondition()
        if Primitive.to_proto(resource.expression):
            res.expression = Primitive.to_proto(resource.expression)
        if Primitive.to_proto(resource.title):
            res.title = Primitive.to_proto(resource.title)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicyDryRunSpecRulesCondition(
            expression=Primitive.from_proto(resource.expression),
            title=Primitive.from_proto(resource.title),
            description=Primitive.from_proto(resource.description),
            location=Primitive.from_proto(resource.location),
        )


class PolicyDryRunSpecRulesConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicyDryRunSpecRulesCondition.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PolicyDryRunSpecRulesCondition.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
