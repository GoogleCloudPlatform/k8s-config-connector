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
from google3.cloud.graphite.mmv2.services.google.billing_budgets import budget_pb2
from google3.cloud.graphite.mmv2.services.google.billing_budgets import budget_pb2_grpc

from typing import List


class Budget(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        budget_filter: dict = None,
        amount: dict = None,
        threshold_rules: list = None,
        etag: str = None,
        all_updates_rule: dict = None,
        billing_account: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.budget_filter = budget_filter
        self.amount = amount
        self.threshold_rules = threshold_rules
        self.all_updates_rule = all_updates_rule
        self.billing_account = billing_account
        self.service_account_file = service_account_file

    def apply(self):
        stub = budget_pb2_grpc.BillingbudgetsAlphaBudgetServiceStub(channel.Channel())
        request = budget_pb2.ApplyBillingbudgetsAlphaBudgetRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if BudgetBudgetFilter.to_proto(self.budget_filter):
            request.resource.budget_filter.CopyFrom(
                BudgetBudgetFilter.to_proto(self.budget_filter)
            )
        else:
            request.resource.ClearField("budget_filter")
        if BudgetAmount.to_proto(self.amount):
            request.resource.amount.CopyFrom(BudgetAmount.to_proto(self.amount))
        else:
            request.resource.ClearField("amount")
        if BudgetThresholdRulesArray.to_proto(self.threshold_rules):
            request.resource.threshold_rules.extend(
                BudgetThresholdRulesArray.to_proto(self.threshold_rules)
            )
        if BudgetAllUpdatesRule.to_proto(self.all_updates_rule):
            request.resource.all_updates_rule.CopyFrom(
                BudgetAllUpdatesRule.to_proto(self.all_updates_rule)
            )
        else:
            request.resource.ClearField("all_updates_rule")
        if Primitive.to_proto(self.billing_account):
            request.resource.billing_account = Primitive.to_proto(self.billing_account)

        request.service_account_file = self.service_account_file

        response = stub.ApplyBillingbudgetsAlphaBudget(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.budget_filter = BudgetBudgetFilter.from_proto(response.budget_filter)
        self.amount = BudgetAmount.from_proto(response.amount)
        self.threshold_rules = BudgetThresholdRulesArray.from_proto(
            response.threshold_rules
        )
        self.etag = Primitive.from_proto(response.etag)
        self.all_updates_rule = BudgetAllUpdatesRule.from_proto(
            response.all_updates_rule
        )
        self.billing_account = Primitive.from_proto(response.billing_account)

    def delete(self):
        stub = budget_pb2_grpc.BillingbudgetsAlphaBudgetServiceStub(channel.Channel())
        request = budget_pb2.DeleteBillingbudgetsAlphaBudgetRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if BudgetBudgetFilter.to_proto(self.budget_filter):
            request.resource.budget_filter.CopyFrom(
                BudgetBudgetFilter.to_proto(self.budget_filter)
            )
        else:
            request.resource.ClearField("budget_filter")
        if BudgetAmount.to_proto(self.amount):
            request.resource.amount.CopyFrom(BudgetAmount.to_proto(self.amount))
        else:
            request.resource.ClearField("amount")
        if BudgetThresholdRulesArray.to_proto(self.threshold_rules):
            request.resource.threshold_rules.extend(
                BudgetThresholdRulesArray.to_proto(self.threshold_rules)
            )
        if BudgetAllUpdatesRule.to_proto(self.all_updates_rule):
            request.resource.all_updates_rule.CopyFrom(
                BudgetAllUpdatesRule.to_proto(self.all_updates_rule)
            )
        else:
            request.resource.ClearField("all_updates_rule")
        if Primitive.to_proto(self.billing_account):
            request.resource.billing_account = Primitive.to_proto(self.billing_account)

        response = stub.DeleteBillingbudgetsAlphaBudget(request)

    @classmethod
    def list(self, billingAccount, service_account_file=""):
        stub = budget_pb2_grpc.BillingbudgetsAlphaBudgetServiceStub(channel.Channel())
        request = budget_pb2.ListBillingbudgetsAlphaBudgetRequest()
        request.service_account_file = service_account_file
        request.BillingAccount = billingAccount

        return stub.ListBillingbudgetsAlphaBudget(request).items

    def to_proto(self):
        resource = budget_pb2.BillingbudgetsAlphaBudget()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if BudgetBudgetFilter.to_proto(self.budget_filter):
            resource.budget_filter.CopyFrom(
                BudgetBudgetFilter.to_proto(self.budget_filter)
            )
        else:
            resource.ClearField("budget_filter")
        if BudgetAmount.to_proto(self.amount):
            resource.amount.CopyFrom(BudgetAmount.to_proto(self.amount))
        else:
            resource.ClearField("amount")
        if BudgetThresholdRulesArray.to_proto(self.threshold_rules):
            resource.threshold_rules.extend(
                BudgetThresholdRulesArray.to_proto(self.threshold_rules)
            )
        if BudgetAllUpdatesRule.to_proto(self.all_updates_rule):
            resource.all_updates_rule.CopyFrom(
                BudgetAllUpdatesRule.to_proto(self.all_updates_rule)
            )
        else:
            resource.ClearField("all_updates_rule")
        if Primitive.to_proto(self.billing_account):
            resource.billing_account = Primitive.to_proto(self.billing_account)
        return resource


class BudgetBudgetFilter(object):
    def __init__(
        self,
        projects: list = None,
        credit_types: list = None,
        credit_types_treatment: str = None,
        services: list = None,
        subaccounts: list = None,
        labels: dict = None,
        calendar_period: str = None,
        custom_period: dict = None,
    ):
        self.projects = projects
        self.credit_types = credit_types
        self.credit_types_treatment = credit_types_treatment
        self.services = services
        self.subaccounts = subaccounts
        self.labels = labels
        self.calendar_period = calendar_period
        self.custom_period = custom_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = budget_pb2.BillingbudgetsAlphaBudgetBudgetFilter()
        if Primitive.to_proto(resource.projects):
            res.projects.extend(Primitive.to_proto(resource.projects))
        if Primitive.to_proto(resource.credit_types):
            res.credit_types.extend(Primitive.to_proto(resource.credit_types))
        if BudgetBudgetFilterCreditTypesTreatmentEnum.to_proto(
            resource.credit_types_treatment
        ):
            res.credit_types_treatment = (
                BudgetBudgetFilterCreditTypesTreatmentEnum.to_proto(
                    resource.credit_types_treatment
                )
            )
        if Primitive.to_proto(resource.services):
            res.services.extend(Primitive.to_proto(resource.services))
        if Primitive.to_proto(resource.subaccounts):
            res.subaccounts.extend(Primitive.to_proto(resource.subaccounts))
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if BudgetBudgetFilterCalendarPeriodEnum.to_proto(resource.calendar_period):
            res.calendar_period = BudgetBudgetFilterCalendarPeriodEnum.to_proto(
                resource.calendar_period
            )
        if BudgetBudgetFilterCustomPeriod.to_proto(resource.custom_period):
            res.custom_period.CopyFrom(
                BudgetBudgetFilterCustomPeriod.to_proto(resource.custom_period)
            )
        else:
            res.ClearField("custom_period")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BudgetBudgetFilter(
            projects=Primitive.from_proto(resource.projects),
            credit_types=Primitive.from_proto(resource.credit_types),
            credit_types_treatment=BudgetBudgetFilterCreditTypesTreatmentEnum.from_proto(
                resource.credit_types_treatment
            ),
            services=Primitive.from_proto(resource.services),
            subaccounts=Primitive.from_proto(resource.subaccounts),
            labels=Primitive.from_proto(resource.labels),
            calendar_period=BudgetBudgetFilterCalendarPeriodEnum.from_proto(
                resource.calendar_period
            ),
            custom_period=BudgetBudgetFilterCustomPeriod.from_proto(
                resource.custom_period
            ),
        )


class BudgetBudgetFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BudgetBudgetFilter.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BudgetBudgetFilter.from_proto(i) for i in resources]


class BudgetBudgetFilterLabels(object):
    def __init__(self, values: list = None):
        self.values = values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = budget_pb2.BillingbudgetsAlphaBudgetBudgetFilterLabels()
        if Primitive.to_proto(resource.values):
            res.values.extend(Primitive.to_proto(resource.values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BudgetBudgetFilterLabels(
            values=Primitive.from_proto(resource.values),
        )


class BudgetBudgetFilterLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BudgetBudgetFilterLabels.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BudgetBudgetFilterLabels.from_proto(i) for i in resources]


class BudgetBudgetFilterCustomPeriod(object):
    def __init__(self, start_date: dict = None, end_date: dict = None):
        self.start_date = start_date
        self.end_date = end_date

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = budget_pb2.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriod()
        if BudgetBudgetFilterCustomPeriodStartDate.to_proto(resource.start_date):
            res.start_date.CopyFrom(
                BudgetBudgetFilterCustomPeriodStartDate.to_proto(resource.start_date)
            )
        else:
            res.ClearField("start_date")
        if BudgetBudgetFilterCustomPeriodEndDate.to_proto(resource.end_date):
            res.end_date.CopyFrom(
                BudgetBudgetFilterCustomPeriodEndDate.to_proto(resource.end_date)
            )
        else:
            res.ClearField("end_date")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BudgetBudgetFilterCustomPeriod(
            start_date=BudgetBudgetFilterCustomPeriodStartDate.from_proto(
                resource.start_date
            ),
            end_date=BudgetBudgetFilterCustomPeriodEndDate.from_proto(
                resource.end_date
            ),
        )


class BudgetBudgetFilterCustomPeriodArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BudgetBudgetFilterCustomPeriod.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BudgetBudgetFilterCustomPeriod.from_proto(i) for i in resources]


class BudgetBudgetFilterCustomPeriodStartDate(object):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = budget_pb2.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodStartDate()
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

        return BudgetBudgetFilterCustomPeriodStartDate(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class BudgetBudgetFilterCustomPeriodStartDateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BudgetBudgetFilterCustomPeriodStartDate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            BudgetBudgetFilterCustomPeriodStartDate.from_proto(i) for i in resources
        ]


class BudgetBudgetFilterCustomPeriodEndDate(object):
    def __init__(self, year: int = None, month: int = None, day: int = None):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = budget_pb2.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodEndDate()
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

        return BudgetBudgetFilterCustomPeriodEndDate(
            year=Primitive.from_proto(resource.year),
            month=Primitive.from_proto(resource.month),
            day=Primitive.from_proto(resource.day),
        )


class BudgetBudgetFilterCustomPeriodEndDateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BudgetBudgetFilterCustomPeriodEndDate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BudgetBudgetFilterCustomPeriodEndDate.from_proto(i) for i in resources]


class BudgetAmount(object):
    def __init__(self, specified_amount: dict = None, last_period_amount: dict = None):
        self.specified_amount = specified_amount
        self.last_period_amount = last_period_amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = budget_pb2.BillingbudgetsAlphaBudgetAmount()
        if BudgetAmountSpecifiedAmount.to_proto(resource.specified_amount):
            res.specified_amount.CopyFrom(
                BudgetAmountSpecifiedAmount.to_proto(resource.specified_amount)
            )
        else:
            res.ClearField("specified_amount")
        if BudgetAmountLastPeriodAmount.to_proto(resource.last_period_amount):
            res.last_period_amount.CopyFrom(
                BudgetAmountLastPeriodAmount.to_proto(resource.last_period_amount)
            )
        else:
            res.ClearField("last_period_amount")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BudgetAmount(
            specified_amount=BudgetAmountSpecifiedAmount.from_proto(
                resource.specified_amount
            ),
            last_period_amount=BudgetAmountLastPeriodAmount.from_proto(
                resource.last_period_amount
            ),
        )


class BudgetAmountArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BudgetAmount.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BudgetAmount.from_proto(i) for i in resources]


class BudgetAmountSpecifiedAmount(object):
    def __init__(self, currency_code: str = None, units: int = None, nanos: int = None):
        self.currency_code = currency_code
        self.units = units
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = budget_pb2.BillingbudgetsAlphaBudgetAmountSpecifiedAmount()
        if Primitive.to_proto(resource.currency_code):
            res.currency_code = Primitive.to_proto(resource.currency_code)
        if Primitive.to_proto(resource.units):
            res.units = Primitive.to_proto(resource.units)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BudgetAmountSpecifiedAmount(
            currency_code=Primitive.from_proto(resource.currency_code),
            units=Primitive.from_proto(resource.units),
            nanos=Primitive.from_proto(resource.nanos),
        )


class BudgetAmountSpecifiedAmountArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BudgetAmountSpecifiedAmount.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BudgetAmountSpecifiedAmount.from_proto(i) for i in resources]


class BudgetAmountLastPeriodAmount(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = budget_pb2.BillingbudgetsAlphaBudgetAmountLastPeriodAmount()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BudgetAmountLastPeriodAmount()


class BudgetAmountLastPeriodAmountArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BudgetAmountLastPeriodAmount.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BudgetAmountLastPeriodAmount.from_proto(i) for i in resources]


class BudgetThresholdRules(object):
    def __init__(self, threshold_percent: float = None, spend_basis: str = None):
        self.threshold_percent = threshold_percent
        self.spend_basis = spend_basis

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = budget_pb2.BillingbudgetsAlphaBudgetThresholdRules()
        if Primitive.to_proto(resource.threshold_percent):
            res.threshold_percent = Primitive.to_proto(resource.threshold_percent)
        if BudgetThresholdRulesSpendBasisEnum.to_proto(resource.spend_basis):
            res.spend_basis = BudgetThresholdRulesSpendBasisEnum.to_proto(
                resource.spend_basis
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BudgetThresholdRules(
            threshold_percent=Primitive.from_proto(resource.threshold_percent),
            spend_basis=BudgetThresholdRulesSpendBasisEnum.from_proto(
                resource.spend_basis
            ),
        )


class BudgetThresholdRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BudgetThresholdRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BudgetThresholdRules.from_proto(i) for i in resources]


class BudgetAllUpdatesRule(object):
    def __init__(
        self,
        pubsub_topic: str = None,
        schema_version: str = None,
        monitoring_notification_channels: list = None,
        disable_default_iam_recipients: bool = None,
    ):
        self.pubsub_topic = pubsub_topic
        self.schema_version = schema_version
        self.monitoring_notification_channels = monitoring_notification_channels
        self.disable_default_iam_recipients = disable_default_iam_recipients

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = budget_pb2.BillingbudgetsAlphaBudgetAllUpdatesRule()
        if Primitive.to_proto(resource.pubsub_topic):
            res.pubsub_topic = Primitive.to_proto(resource.pubsub_topic)
        if Primitive.to_proto(resource.schema_version):
            res.schema_version = Primitive.to_proto(resource.schema_version)
        if Primitive.to_proto(resource.monitoring_notification_channels):
            res.monitoring_notification_channels.extend(
                Primitive.to_proto(resource.monitoring_notification_channels)
            )
        if Primitive.to_proto(resource.disable_default_iam_recipients):
            res.disable_default_iam_recipients = Primitive.to_proto(
                resource.disable_default_iam_recipients
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BudgetAllUpdatesRule(
            pubsub_topic=Primitive.from_proto(resource.pubsub_topic),
            schema_version=Primitive.from_proto(resource.schema_version),
            monitoring_notification_channels=Primitive.from_proto(
                resource.monitoring_notification_channels
            ),
            disable_default_iam_recipients=Primitive.from_proto(
                resource.disable_default_iam_recipients
            ),
        )


class BudgetAllUpdatesRuleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BudgetAllUpdatesRule.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BudgetAllUpdatesRule.from_proto(i) for i in resources]


class BudgetBudgetFilterCreditTypesTreatmentEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return budget_pb2.BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum.Value(
            "BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return budget_pb2.BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum.Name(
            resource
        )[
            len("BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum") :
        ]


class BudgetBudgetFilterCalendarPeriodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return budget_pb2.BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum.Value(
            "BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return budget_pb2.BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum.Name(
            resource
        )[len("BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum") :]


class BudgetThresholdRulesSpendBasisEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return budget_pb2.BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum.Value(
            "BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return budget_pb2.BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum.Name(
            resource
        )[len("BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
